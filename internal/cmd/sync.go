package cmd

import (
	"github.com/musaprg/annict-epgstation-connector/internal/syncer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
	"log"
	"net/http"
	"time"
)

const (
	annictEndpoint = "https://api.annict.com/graphql"
)

var syncCmd = &cli.Command{
	Name:  "sync",
	Usage: "sync",
	Flags: []cli.Flag{
		// daemon mode
		&cli.BoolFlag{
			Name:    "daemon",
			Aliases: []string{"d"},
			Usage:   "run as a daemon",
		},
		// interval
		&cli.IntFlag{
			Name:    "interval",
			Aliases: []string{"i"},
			Usage:   "interval of sync in daemon mode (seconds)",
			Value:   60,
		},
		// metrics listen address
		&cli.StringFlag{
			Name:  "metrics-listen-address",
			Usage: "listen address of metrics",
			Value: ":8080",
		},
		// db path
		&cli.StringFlag{
			Name:    "db-path",
			Aliases: []string{"db"},
			Usage:   "path to the database",
			Value:   "/tmp/annict-epgstation-connector",
		},
	},
	Action: func(c *cli.Context) error {
		s, err := syncer.NewSyncer(
			syncer.WithAnnictAPIToken(c.String(string(annictAPITokenFlag))),
			syncer.WithAnnictEndpoint(annictEndpoint),
			syncer.WithEPGStationEndpoint(c.String(string(epgstationEndpointFlag))),
			syncer.WithDBPath(c.String("db-path")),
		)
		if err != nil {
			return err
		}

		// daemon mode
		if c.Bool("daemon") {
			// metrics http server for prometheus
			// only in daemon mode
			go func() {
				http.Handle("/metrics", promhttp.Handler())
				if err := http.ListenAndServe(c.String("metrics-listen-address"), nil); err != nil {
					log.Fatal(err)
				}
			}()

			interval := c.Int("interval")
			slog.Info("start sync in daemon mode", slog.Int("interval", interval))
			for {
				slog.Info("syncing...")
				if err := s.Sync(c.Context); err != nil {
					return err
				}
				slog.Info("finish sync and sleep", slog.Int("interval", interval))
				time.Sleep(time.Duration(interval) * time.Second)
			}
		} else { // one shot mode
			slog.Info("start sync in one shot mode")
			if err := s.Sync(c.Context); err != nil {
				return err
			}
		}
		return nil
	},
}
