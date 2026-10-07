package types_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/instana/instana-go-client/shared/types"
)

func TestThresholdRuleJSONMarshalling(t *testing.T) {
	upper := 90.0
	lower := 20.0
	val := 50.0

	tests := []struct {
		name          string
		rule          types.ThresholdRule
		shouldContain []string
		shouldOmit    []string
	}{
		{
			name: "Range bounds present",
			rule: types.ThresholdRule{
				Type:       "static",
				UpperBound: &upper,
				LowerBound: &lower,
			},
			shouldContain: []string{`"upperBound":90`, `"lowerBound":20`},
			shouldOmit:    []string{`"upperBound":null`, `"lowerBound":null`},
		},
		{
			name: "Single value present without bounds",
			rule: types.ThresholdRule{
				Type:  "static",
				Value: &val,
			},
			shouldContain: []string{`"value":50`},
			shouldOmit:    []string{`"upperBound"`, `"lowerBound"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.rule)
			if err != nil {
				t.Fatalf("Failed to marshal ThresholdRule: %v", err)
			}
			jsonStr := string(data)

			for _, expected := range tt.shouldContain {
				if !strings.Contains(jsonStr, expected) {
					t.Errorf("Expected JSON to contain %s, got %s", expected, jsonStr)
				}
			}
			for _, omitted := range tt.shouldOmit {
				if strings.Contains(jsonStr, omitted) {
					t.Errorf("Expected JSON to omit %s, got %s", omitted, jsonStr)
				}
			}

			var unmarshalled types.ThresholdRule
			if err := json.Unmarshal(data, &unmarshalled); err != nil {
				t.Fatalf("Failed to unmarshal ThresholdRule: %v", err)
			}

			if tt.rule.UpperBound != nil {
				if unmarshalled.UpperBound == nil || *unmarshalled.UpperBound != *tt.rule.UpperBound {
					t.Errorf("Expected UpperBound %v, got %v", *tt.rule.UpperBound, unmarshalled.UpperBound)
				}
			}
			if tt.rule.LowerBound != nil {
				if unmarshalled.LowerBound == nil || *unmarshalled.LowerBound != *tt.rule.LowerBound {
					t.Errorf("Expected LowerBound %v, got %v", *tt.rule.LowerBound, unmarshalled.LowerBound)
				}
			}
		})
	}
}

func TestThresholdOperatorsToStringSlice(t *testing.T) {
	operators := types.SupportedThresholdOperators
	slice := operators.ToStringSlice()
	if len(slice) != 4 {
		t.Errorf("Expected 4 operators, got %d", len(slice))
	}
}
