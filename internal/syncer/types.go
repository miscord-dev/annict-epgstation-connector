package syncer

import "time"

type annictWork struct {
	ID         string
	SeasonName string
	SeasonYear int
	Title      string
	// StartedAt represents the time when the work is going to start or started broadcasting.
	StartedAt   time.Time
	VodServices []VodService
}

type RecordingRuleID int

type RecordingRuleIDs []RecordingRuleID

// VodService represents a VOD service.
type VodService struct {
	Name string
}
