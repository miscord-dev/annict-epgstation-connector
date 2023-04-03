package cmd

import (
	"github.com/urfave/cli/v2"
)

type flagName string

var (
	epgstationEndpointFlag = flagName("epgstation-endpoint")
	annictAPITokenFlag     = flagName("annict-api-token")
)

var rootCmd = &cli.App{
	Name:  "annict-epgstation-connector",
	Usage: "generate recording rules based on the annict statuses",
	Commands: []*cli.Command{
		syncCmd,
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     string(annictAPITokenFlag),
			Usage:    "An API token of annict GraphQL API server",
			EnvVars:  []string{"ANNICT_API_TOKEN"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     string(epgstationEndpointFlag),
			Usage:    "An endpoint of EPGStation API server",
			EnvVars:  []string{"EPGSTATION_ENDPOINT"},
			Required: true,
		},
	},
}

func Root() *cli.App {
	return rootCmd
}
