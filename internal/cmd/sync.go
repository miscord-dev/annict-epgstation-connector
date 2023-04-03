package cmd

import (
	"github.com/musaprg/annict-epgstation-connector/internal/syncer"
	"github.com/urfave/cli/v2"
)

const (
	annictEndpoint = "https://api.annict.com/graphql"
)

var syncCmd = &cli.Command{
	Name:  "sync",
	Usage: "sync",
	Action: func(c *cli.Context) error {
		s, err := syncer.NewSyncer(&syncer.SyncerOpt{
			AnnictEndpoint:     annictEndpoint,
			EPGStationEndpoint: c.String(string(epgstationEndpointFlag)),
			AnnictAPIToken:     c.String(string(annictAPITokenFlag)),
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
