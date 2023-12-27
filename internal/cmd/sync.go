package cmd

import (
	"github.com/musaprg/annict-epgstation-connector/internal/syncer"
	"github.com/urfave/cli/v2"
)

const (
	annictEndpoint = "https://api.annict.com/graphql"
)

var (
	cleanupFlag = flagName("cleanup")
)

var syncCmd = &cli.Command{
	Name:  "sync",
	Usage: "sync",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  string(cleanupFlag),
			Usage: "Cleanup duplicated rules before sync",
		},
	},
	Action: func(c *cli.Context) error {
		s, err := syncer.NewSyncer(&syncer.SyncerOpt{
			AnnictEndpoint:     annictEndpoint,
			EPGStationEndpoint: c.String(string(epgstationEndpointFlag)),
			AnnictAPIToken:     c.String(string(annictAPITokenFlag)),
			Cleanup:            c.Bool(string(cleanupFlag)),
		})
		if err != nil {
			return err
		}
		if err := s.Sync(c.Context); err != nil {
			return err
		}
		return nil
	},
}
