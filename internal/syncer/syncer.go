package syncer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Khan/genqlient/graphql"
	"github.com/cockroachdb/pebble"
	"github.com/miscord-dev/annict-epgstation-connector/annict"
	"github.com/miscord-dev/annict-epgstation-connector/epgstation"
	"go.uber.org/multierr"
	"golang.org/x/exp/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

const (
	defaultAnnictEndpoint     = "https://api.annict.com/graphql"
	defaultEPGStationEndpoint = "http://localhost:8888/api"
	defaultDBPath             = ""
)

type Interface interface {
	Sync(context.Context) error
}

type syncer struct {
	annictClient graphql.Client
	esClient     *epgstation.Client
	db           *pebble.DB
}

type options struct {
	AnnictEndpoint     string
	AnnictAPIToken     string
	EPGStationEndpoint string
	DBPath             string
}

type Option func(*options)

func WithAnnictEndpoint(endpoint string) Option {
	return func(o *options) {
		o.AnnictEndpoint = endpoint
	}
}

func WithAnnictAPIToken(token string) Option {
	return func(o *options) {
		o.AnnictAPIToken = token
	}
}

func WithEPGStationEndpoint(endpoint string) Option {
	return func(o *options) {
		o.EPGStationEndpoint = endpoint
	}
}

func WithDBPath(path string) Option {
	return func(o *options) {
		o.DBPath = path
	}
}

func NewSyncer(opts ...Option) (Interface, error) {
	o := options{
		AnnictEndpoint:     defaultAnnictEndpoint,
		AnnictAPIToken:     "",
		EPGStationEndpoint: defaultEPGStationEndpoint,
		DBPath:             defaultDBPath,
	}
	for _, opt := range opts {
		opt(&o)
	}

	dbPath := filepath.Join(o.DBPath, "db")
	db, err := pebble.Open(dbPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize DB for Syncer: %w", err)
	}
	annictClient := graphql.NewClient(o.AnnictEndpoint,
		&http.Client{Transport: annict.NewAuthedTransport(o.AnnictAPIToken, http.DefaultTransport)})
	esClient, err := epgstation.NewClient(o.EPGStationEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Syncer: %w", err)
	}
	return &syncer{
		annictClient: annictClient,
		esClient:     esClient,
		db:           db,
	}, nil
}

func (s *syncer) TearDown() error {
	if err := s.db.Close(); err != nil {
		return fmt.Errorf("failed to tear down Syncer DB: %w", err)
	}
	return nil
}

func (s *syncer) Sync(ctx context.Context) error {
	start := time.Now()
	defer func() {
		syncerSyncDuration.WithLabelValues().Observe(time.Now().Sub(start).Seconds())
	}()

	if err := s.sync(ctx); err != nil {
		syncerSyncError.WithLabelValues().Inc()
		return fmt.Errorf("failed to sync: %w", err)
	}
	syncerSyncSuccess.WithLabelValues().Inc()
	return nil
}

func (s *syncer) sync(ctx context.Context) error {
	var titles []annictWork
	if ts, err := s.getWannaWatchWorks(ctx); err != nil {
		return err
	} else {
		titles = append(titles, ts...)
	}
	if ts, err := s.getWatchingWorks(ctx); err != nil {
		return err
	} else {
		titles = append(titles, ts...)
	}
	if ts, err := s.getOnHoldWorks(ctx); err != nil {
		return err
	} else {
		titles = append(titles, ts...)
	}

	if err := s.registerRulesToEPGStation(ctx, titles); err != nil {
		return err
	}
	return nil
}

func (s *syncer) registerRulesToEPGStation(ctx context.Context, works []annictWork) error {
	var errs error
	for _, work := range works {
		errs = multierr.Append(errs, s.registerRuleToEPGStation(ctx, work))
	}
	if errs != nil {
		return fmt.Errorf("failed to register rules into EPGStation: %w", errs)
	}
	return nil
}

func (s *syncer) registerRuleToEPGStation(ctx context.Context, work annictWork) error {
	syncerAnnictWorkStartedAt.WithLabelValues(
		work.ID,
		work.Title,
		work.SeasonName,
		strconv.Itoa(work.SeasonYear),
	).Set(float64(work.StartedAt.Unix()))

	ruleIDs, err := s.getRecordingRuleIDsByAnnictWorkID(work.ID)
	switch {
	case err != nil && !errors.Is(err, pebble.ErrNotFound):
		return fmt.Errorf("failed to get recording rule IDs for Annict work ID %s: %w", work.ID, err)
	case err == nil:
		// recording rule IDs found for the given Annict work ID
		for _, id := range ruleIDs {
			syncerRecordingRuleSynced.WithLabelValues(strconv.Itoa(int(id)), work.ID).Set(1)
		}
		return nil
	}
	if rules, _ := s.getRulesByKeyword(ctx, work.Title); len(rules) != 0 {
		// recording rule with same keyword has already been registered
		// skip registration
		// TODO: Remove this logic after introducing cleanup logic
		slog.Debug("recording rule with same keyword has already been registered", slog.String("keyword", work.Title), slog.String("annict_work_id", work.ID), slog.Int("number_of_rules", len(rules)))
		for _, rule := range rules {
			syncerRecordingRuleSynced.WithLabelValues(strconv.Itoa(rule.Id), work.ID).Set(1)
			ruleIDs = append(ruleIDs, RecordingRuleID(rule.Id))
		}
		if err := s.setRecordingRuleIDsByAnnictWorkID(work.ID, ruleIDs); err != nil {
			return fmt.Errorf("failed to store recording rule IDs for Annict work ID %s: %w", work.ID, err)
		}
		slog.Debug("recording rule with same keyword has already been registered", slog.String("keyword", work.Title))
		return nil
	}
	body := epgstation.PostRulesJSONRequestBody{
		SearchOption: epgstation.RuleSearchOption{
			GR: epgstation.NewTruePointer(),
			BS: epgstation.NewTruePointer(),

			// Only search by work
			Keyword:     &work.Title,
			Name:        epgstation.NewTruePointer(),
			Description: epgstation.NewFalsePointer(),
			Extended:    epgstation.NewFalsePointer(),

			// https://github.com/l3tnun/EPGStation/blob/master/client/src/lib/event.ts
			Genres: &[]epgstation.Genre{
				{Genre: 0x6}, // 0x6 = 映画
				{Genre: 0x7}, // 0x7 = アニメ・特撮
			},

			Times: &[]epgstation.SearchTime{
				{
					// whole week
					Week: 0b1111111,
				},
			},

			IsFree: epgstation.NewTruePointer(), // TODO(musaprg): how about NHK?
		},
		IsTimeSpecification: false,
		SaveOption:          &epgstation.ReserveSaveOption{},
		EncodeOption:        &epgstation.ReserveEncodedOption{},
		ReserveOption: epgstation.RuleReserveOption{
			AvoidDuplicate: false,
			Enable:         true,
			AllowEndLack:   false,
		},
	}
	slog.Debug("registering rules into EPGStation", slog.String("annict_work_title", work.Title), slog.String("annict_work_id", work.ID))
	r, err := s.esClient.PostRules(ctx, body)
	if err != nil {
		return err
	}
	slog.Debug("got response from EPGStation", slog.String("annict_work_title", work.Title), slog.String("annict_work_id", work.ID), slog.Int("status_code", r.StatusCode))
	res, err := epgstation.ParsePostRulesResponse(r)
	if err != nil {
		return err
	}
	if res.JSON201 == nil {
		return fmt.Errorf("failed to register rules into EPGStation: %s", res.Body)
	}
	ids := RecordingRuleIDs{RecordingRuleID(res.JSON201.RuleId)}
	slog.Debug("store recording rule IDs into DB", slog.String("annict_work_title", work.Title), slog.String("annict_work_id", work.ID), slog.Int("number_of_rules", len(ids)), slog.Int("rule_id", int(ids[0])))
	if err := s.setRecordingRuleIDsByAnnictWorkID(work.ID, ids); err != nil {
		return err
	}
	slog.Debug("registered rules into EPGStation", slog.String("annict_work_title", work.Title), slog.String("annict_work_id", work.ID), slog.Int("number_of_rules", len(ids)))
	syncerRecordingRuleSynced.WithLabelValues(strconv.Itoa(int(ids[0])), work.ID).Set(1)
	return nil
}

func (s *syncer) getRulesByKeyword(ctx context.Context, keyword string) ([]epgstation.RuleKeywordItem, error) {
	r, err := s.esClient.GetRulesKeyword(ctx, &epgstation.GetRulesKeywordParams{
		Keyword: &keyword,
	})
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, nil
	}
	res, err := epgstation.ParseGetRulesKeywordResponse(r)
	if err != nil {
		return nil, err
	}
	return res.JSON200.Items, nil
}

func (s *syncer) getWannaWatchWorks(ctx context.Context) ([]annictWork, error) {
	var titles []annictWork
	r, err := annict.GetWannaWatchWorks(ctx, s.annictClient)
	if err != nil {
		return titles, fmt.Errorf("failed to sync: %w", err)
	}
	for _, n := range r.Viewer.Works.Nodes {
		var startedAt time.Time
		if len(n.Programs.Nodes) != 0 {
			startedAt = n.Programs.Nodes[0].StartedAt
		}
		titles = append(titles, annictWork{
			ID:         strconv.Itoa(n.AnnictId),
			Title:      n.Title,
			SeasonName: string(n.SeasonName),
			SeasonYear: n.SeasonYear,
			StartedAt:  startedAt,
		})
	}
	return titles, nil
}

func (s *syncer) getWatchingWorks(ctx context.Context) ([]annictWork, error) {
	var titles []annictWork
	r, err := annict.GetWatchingWorks(ctx, s.annictClient)
	if err != nil {
		return titles, fmt.Errorf("failed to sync: %w", err)
	}
	for _, n := range r.Viewer.Works.Nodes {
		var startedAt time.Time
		if len(n.Programs.Nodes) != 0 {
			startedAt = n.Programs.Nodes[0].StartedAt
		}
		titles = append(titles, annictWork{
			ID:         strconv.Itoa(n.AnnictId),
			Title:      n.Title,
			SeasonName: string(n.SeasonName),
			SeasonYear: n.SeasonYear,
			StartedAt:  startedAt,
		})
	}
	return titles, nil
}

func (s *syncer) getOnHoldWorks(ctx context.Context) ([]annictWork, error) {
	var titles []annictWork
	r, err := annict.GetOnHoldWorks(ctx, s.annictClient)
	if err != nil {
		return titles, fmt.Errorf("failed to sync: %w", err)
	}
	for _, n := range r.Viewer.Works.Nodes {
		var startedAt time.Time
		if len(n.Programs.Nodes) != 0 {
			startedAt = n.Programs.Nodes[0].StartedAt
		}
		titles = append(titles, annictWork{
			ID:         strconv.Itoa(n.AnnictId),
			Title:      n.Title,
			SeasonName: string(n.SeasonName),
			SeasonYear: n.SeasonYear,
			StartedAt:  startedAt,
		})
	}
	return titles, nil
}

// TODO: move these functions to a separate package

// setRecordingRuleIDsByAnnictWorkID stores recording rule IDs for the given Annict work ID in the pebble DB
func (s *syncer) setRecordingRuleIDsByAnnictWorkID(annictWorkID string, ids []RecordingRuleID) error {
	value, err := json.Marshal(ids)
	if err != nil {
		return fmt.Errorf("failed to encode recording rule IDs for Annict work ID %s: %w", annictWorkID, err)
	}
	err = s.db.Set([]byte(annictWorkID), value, pebble.Sync)
	if err != nil {
		return fmt.Errorf("failed to store recording rule IDs for Annict work ID %s: %w", annictWorkID, err)
	}
	return nil
}

// getRecordingRuleIDsByAnnictWorkID returns recording rule IDs for the given Annict work ID stored in the pebble DB
func (s *syncer) getRecordingRuleIDsByAnnictWorkID(annictWorkID string) ([]RecordingRuleID, error) {
	var ids []RecordingRuleID
	value, closer, err := s.db.Get([]byte(annictWorkID))
	if err != nil {
		return ids, fmt.Errorf("failed to get recording rule IDs for Annict work ID %s: %w", annictWorkID, err)
	}
	defer closer.Close()
	err = json.NewDecoder(bytes.NewReader(value)).Decode(&ids)
	if err != nil {
		return ids, fmt.Errorf("failed to decode recording rule IDs for Annict work ID %s: %w", annictWorkID, err)
	}
	return ids, nil
}
