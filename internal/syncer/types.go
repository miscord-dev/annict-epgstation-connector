package syncer

import "time"

type annictWork struct {
	ID         string
	SeasonName string
	SeasonYear int
	Title      string
	// StartedAt represents the time when the work is going to start or started broadcasting.
	StartedAt time.Time
}

type RecordingRuleID int

type RecordingRuleIDs []RecordingRuleID
