package types

// RuleWithThreshold represents a rule with threshold configuration
// This is a generic type that can be used with different rule types
type RuleWithThreshold[R any] struct {
	ThresholdOperator ThresholdOperator               `json:"thresholdOperator"`
	Rule              R                               `json:"rule"`
	Thresholds        map[AlertSeverity]ThresholdRule `json:"thresholds"`
}
