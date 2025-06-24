package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/miscord-dev/annict-epgstation-connector/internal/syncer"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"log/slog"
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
		// exclude VOD services
		&cli.StringSliceFlag{
			Name:    "exclude-vod-services",
			Aliases: []string{"exclude-vod"},
			Usage:   "exclude recording rules for anime available on these VOD services (comma-separated list: netflix, amazon-prime, hulu, disney, abema, crunchyroll, funimation, dazn, bandai, nico, danime)",
		},
		// enable VOD fallback
		&cli.BoolFlag{
			Name:  "enable-vod-fallback",
			Usage: "enable fallback VOD detection when specific VOD section is not found (searches all page links with filtering)",
			Value: false,
		},
		// enable recording rule removal
		&cli.BoolFlag{
			Name:  "enable-rule-removal",
			Usage: "enable automatic removal of recording rules when anime is marked as STOP_WATCHING or WATCHED on Annict",
			Value: false,
		},
	},
	Action: func(c *cli.Context) error {
		// Parse excluded VOD services
		excludedVODServices := c.StringSlice("exclude-vod-services")
		enableVODFallback := c.Bool("enable-vod-fallback")
		enableRuleRemoval := c.Bool("enable-rule-removal")

		s, err := syncer.NewSyncer(
			syncer.WithAnnictAPIToken(c.String(string(annictAPITokenFlag))),
			syncer.WithAnnictEndpoint(annictEndpoint),
			syncer.WithEPGStationEndpoint(c.String(string(epgstationEndpointFlag))),
			syncer.WithDBPath(c.String("db-path")),
			syncer.WithExcludedVODServicesFromStrings(excludedVODServices),
			syncer.WithVODFallback(enableVODFallback),
			syncer.WithRuleRemoval(enableRuleRemoval),
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
					slog.Error("failed to sync", slog.String("error", err.Error()))
				} else {
					slog.Info("synced")
				}
				slog.Info("sleep", slog.Int("interval", interval))
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
