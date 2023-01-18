package syncer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/musaprg/annict-epgstation-connector/annict"
	"github.com/musaprg/annict-epgstation-connector/epgstation"
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
	EPGStationEndpoint string
}

func NewSyncer(opts SyncerOpt) (Interface, error) {
	annictClient := graphql.NewClient(opts.AnnictEndpoint, http.DefaultClient)
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
	titles, err := s.getWannaWatchWorks(ctx)
	if err != nil {
		return err
	}
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
			body := epgstation.PostRulesJSONRequestBody{
				SearchOption: epgstation.RuleSearchOption{
					SKY: epgstation.NewTruePointer(),
					BS:  epgstation.NewTruePointer(),

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
			res, err := s.esClient.PostRules(ctx, body)
			if err != nil {
				return err
			}
			if res.StatusCode != http.StatusOK {
				return fmt.Errorf("invalid status code: %d", res.StatusCode)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return fmt.Errorf("failed to register rules into EPGStation: %w", err)
	}
	return nil
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
