package types

// AlertSeverity represents the severity level of an alert
type AlertSeverity string

// AlertSeverities custom type for a slice of AlertSeverity
type AlertSeverities []AlertSeverity

const (
	// WarningSeverity constant value for warning severity
	WarningSeverity = AlertSeverity("WARNING")
	// CriticalSeverity constant value for critical severity
	CriticalSeverity = AlertSeverity("CRITICAL")
)

// SupportedAlertSeverities list of all supported alert severities
var SupportedAlertSeverities = AlertSeverities{WarningSeverity, CriticalSeverity}

// Made with Bob
