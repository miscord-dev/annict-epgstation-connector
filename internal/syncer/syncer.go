package syncer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/musaprg/annict-epgstation-connector/annict"
	"github.com/musaprg/annict-epgstation-connector/epgstation"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

type Interface interface {
	Sync(context.Context) error
}

type syncer struct {
	annictClient graphql.Client
	esClient     *epgstation.Client
}

type SyncerOpt struct {
	AnnictEndpoint     string
	AnnictAPIToken     string
	EPGStationEndpoint string
}

func NewSyncer(opts *SyncerOpt) (Interface, error) {
	annictClient := graphql.NewClient(opts.AnnictEndpoint,
		&http.Client{Transport: annict.NewAuthedTransport(opts.AnnictAPIToken, http.DefaultTransport)})
	esClient, err := epgstation.NewClient(opts.EPGStationEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Syncer: %w", err)
	}
	return &syncer{
		annictClient: annictClient,
		esClient:     esClient,
	}, nil
}

func (s *syncer) Sync(ctx context.Context) error {
	titles := []string{}
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
	slices.Compact(titles)

	if err := s.registerRulesToEpgStation(ctx, titles); err != nil {
		return err
	}
	return nil
}

func (s *syncer) registerRulesToEpgStation(ctx context.Context, titles []string) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, title := range titles {
		title := title
		eg.Go(func() error {
			if rules, _ := s.getRulesByKeyword(ctx, title); len(rules) != 0 {
				// TODO(musaprg): output log message about skipping registeration of rule for this keyword
				return nil
			}
			body := epgstation.PostRulesJSONRequestBody{
				SearchOption: epgstation.RuleSearchOption{
					GR: epgstation.NewFalsePointer(),
					BS: epgstation.NewTruePointer(),

					Keyword:     &title,
					Description: epgstation.NewTruePointer(),
					Name:        epgstation.NewTruePointer(),
					Extended:    epgstation.NewTruePointer(), // TODO(musaprg): is this really needed?

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
					AvoidDuplicate: true,
					Enable:         true,
					AllowEndLack:   false,
				},
			}
			_, err := s.esClient.PostRules(ctx, body)
			if err != nil {
				return err
			}
			// TODO(musaprg): output response in the log message
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to register rules into EPGStation: %w", err)
	}
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

func (s *syncer) getWannaWatchWorks(ctx context.Context) ([]string, error) {
	var titles []string
	r, err := annict.GetWannaWatchWorks(ctx, s.annictClient)
	if err != nil {
		return titles, fmt.Errorf("failed to sync: %w", err)
	}
	for _, n := range r.Viewer.Works.Nodes {
		titles = append(titles, n.Title)
	}
	return titles, nil
}

func (s *syncer) getWatchingWorks(ctx context.Context) ([]string, error) {
	var titles []string
	r, err := annict.GetWatchingWorks(ctx, s.annictClient)
	if err != nil {
		return titles, fmt.Errorf("failed to sync: %w", err)
	}
	for _, n := range r.Viewer.Works.Nodes {
		titles = append(titles, n.Title)
	}
	return titles, nil
}

func (s *syncer) getOnHoldWorks(ctx context.Context) ([]string, error) {
	var titles []string
	r, err := annict.GetOnHoldWorks(ctx, s.annictClient)
	if err != nil {
		return titles, fmt.Errorf("failed to sync: %w", err)
	}
	for _, n := range r.Viewer.Works.Nodes {
		titles = append(titles, n.Title)
	}
	return titles, nil
}
