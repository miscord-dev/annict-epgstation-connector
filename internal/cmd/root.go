package cmd

import (
	"github.com/urfave/cli/v2"
)

var rootCmd = &cli.App{
	Name:  "annict-epgstation-connector",
	Usage: "generate recording rules based on the annict statuses",
	Commands: []*cli.Command{
		syncCmd,
	},
}

func Root() *cli.App {
	return rootCmd
}
