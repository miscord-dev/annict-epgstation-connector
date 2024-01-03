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
		s, err := syncer.NewSyncer(
			syncer.WithAnnictAPIToken(c.String(string(annictAPITokenFlag))),
			syncer.WithAnnictEndpoint(annictEndpoint),
			syncer.WithEPGStationEndpoint(c.String(string(epgstationEndpointFlag))),
		)
		if err != nil {
			return err
		}
		if err := s.Sync(c.Context); err != nil {
			return err
		}
		return nil
	},
}
