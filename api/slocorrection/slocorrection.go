package slocorrection

import (
	"github.com/instana/instana-go-client/shared/types"
)

const (
	//SloCorrectionConfigResourcePath path to slo correction config resource of Instana RESTful API
	SloCorrectionConfigResourcePath = "/api/settings" + "/correction"
)

// SloCorrectionConfig represents the REST resource of SLO Correction Configuration at Instana
type SloCorrectionConfig struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Active      bool             `json:"active"`
	Scheduling  types.Scheduling `json:"scheduling"`
	SloIds      []string         `json:"sloIds"`
	Tags        []string         `json:"tags"`
}

type DurationUnit string

const (
	DurationUnitMinute DurationUnit = "MINUTE"
	DurationUnitHour   DurationUnit = "HOUR"
	DurationUnitDay    DurationUnit = "DAY"
)

type Scheduling struct {
	StartTime     int64        `json:"startTime"` // Unix timestamp in milliseconds
	Duration      int          `json:"duration"`
	DurationUnit  DurationUnit `json:"durationUnit"`
	RecurrentRule string       `json:"recurrentRule,omitempty"`
	Recurrent     bool         `json:"recurrent"`
}

// GetIDForResourcePath implementation of the interface InstanaDataObject
func (s *SloCorrectionConfig) GetIDForResourcePath() string {
	return s.ID
}
