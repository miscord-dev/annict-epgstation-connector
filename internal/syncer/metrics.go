package syncer

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "annict_epgstation_connector"
)

func init() {
	prometheus.MustRegister(syncerSyncDuration)
	prometheus.MustRegister(syncerSyncSuccess)
	prometheus.MustRegister(syncerSyncError)
	prometheus.MustRegister(syncerRecordingRuleSynced)
	prometheus.MustRegister(syncerAnnictWorkStartedAt)
}

var (
	// syncerSyncDuration is a histogram of the time it takes to sync.
	syncerSyncDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: "syncer",
			Name:      "sync_duration_seconds",
			Help:      "The time it takes to sync.",
			Buckets:   prometheus.DefBuckets,
		},
		nil,
	)

	// syncerSyncSuccess is a counter of the number of successful syncs.
	syncerSyncSuccess = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "syncer",
			Name:      "sync_success_total",
			Help:      "The number of successful syncs.",
		},
		nil,
	)

	// syncerSyncError is a counter of the number of failed syncs.
	syncerSyncError = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: "syncer",
			Name:      "sync_error_total",
			Help:      "The number of failed syncs.",
		},
		nil,
	)

	// syncerRecordingRuleSynced is a gauge if the recording rule is synced.
	syncerRecordingRuleSynced = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "syncer",
			Name:      "recording_rule_synced",
			Help:      "Whether the recording rule is synced.",
		},
		[]string{
			"rule_id",
			"annict_work_id",
		},
	)

	// syncerAnnictWorkStartedAt is a gauge of the time when the annict work is going to start or started broadcasting.
	syncerAnnictWorkStartedAt = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: "syncer",
			Name:      "annict_work_started_at",
			Help:      "The time when the annict work is going to start or started broadcasting.",
		},
		[]string{
			"id",
			"title",
			"season_name",
			"season_year",
		},
	)
)
