package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var syncCmd = &cli.Command{
	Name:  "sync",
	Usage: "sync",
	Action: func(c *cli.Context) error {
		fmt.Println("added task: ", c.Args().First())
		return nil
	},
}
