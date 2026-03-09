package types

// Scheduling represents scheduling configuration for automation policies and SLO corrections
type Scheduling struct {
	Start      int64      `json:"start"`
	End        int64      `json:"end"`
	Recurrence Recurrence `json:"recurrence"`
	Duration   Duration   `json:"duration"`
}

// Recurrence represents recurrence configuration
type Recurrence struct {
	Type     string `json:"type"`
	Interval int    `json:"interval"`
}

// Duration represents duration configuration
type Duration struct {
	Amount int          `json:"amount"`
	Unit   DurationUnit `json:"unit"`
}

// DurationUnit represents the unit of duration
type DurationUnit string

const (
	// DurationUnitMinute represents minute duration unit
	DurationUnitMinute DurationUnit = "MINUTE"
	// DurationUnitHour represents hour duration unit
	DurationUnitHour DurationUnit = "HOUR"
	// DurationUnitDay represents day duration unit
	DurationUnitDay DurationUnit = "DAY"
)
