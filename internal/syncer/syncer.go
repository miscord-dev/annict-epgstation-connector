package syncer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/musaprg/annict-epgstation-connector/annict"
	"github.com/musaprg/annict-epgstation-connector/epgstation"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

type Interface interface {
	Sync(context.Context) error
}

type syncer struct {
	annictClient graphql.Client
	esClient     *epgstation.Client
	cleanup      bool
	logger       *zap.Logger
}

type SyncerOpt struct {
	AnnictEndpoint     string
	AnnictAPIToken     string
	EPGStationEndpoint string
	Cleanup            bool
	Debug              bool
}

func NewSyncer(opts *SyncerOpt) (Interface, error) {
	annictClient := graphql.NewClient(opts.AnnictEndpoint,
		&http.Client{Transport: annict.NewAuthedTransport(opts.AnnictAPIToken, http.DefaultTransport)})
	esClient, err := epgstation.NewClient(opts.EPGStationEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Syncer: %w", err)
	}
	var logger *zap.Logger
	if opts.Debug {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize logger: %w", err)
		}
	} else {
		logger, _ = zap.NewProduction()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize logger: %w", err)
		}
	}
	return &syncer{
		annictClient: annictClient,
		esClient:     esClient,
		cleanup:      opts.Cleanup,
		logger:       logger,
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
	eg, egCtx := errgroup.WithContext(ctx)
	for _, title := range titles {
		title := title
		eg.Go(func() error {
			if rules, _ := s.getRulesByKeyword(egCtx, title); len(rules) != 0 {
				s.logger.Debug("rule already exists", zap.String("title", title))
				if s.cleanup {
					if err := s.deleteRules(egCtx, rules); err != nil {
						return fmt.Errorf("failed to delete rules: %w", err)
					}
				} else {
					return nil
				}
			}
			body := epgstation.PostRulesJSONRequestBody{
				SearchOption: epgstation.RuleSearchOption{
					GR: epgstation.NewTruePointer(),
					BS: epgstation.NewTruePointer(),

					// Only search by title
					Keyword:     &title,
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
			_, err := s.esClient.PostRules(ctx, body)
			if err != nil {
				return err
			}
			s.logger.Info("registered rule", zap.String("title", title))
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to register rules into EPGStation: %w", err)
	}
	return nil
}

func (s *syncer) deleteRules(ctx context.Context, rules []epgstation.RuleKeywordItem) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, rule := range rules {
		rule := rule
		eg.Go(func() error {
			_, err := s.esClient.DeleteRulesRuleId(ctx, rule.Id)
			if err != nil {
				return err
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to delete rules: %w", err)
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
