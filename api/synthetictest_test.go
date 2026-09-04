package api_test

import (
	"encoding/json"
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestSyntheticTestResourcePath(t *testing.T) {
	expected := "/api/synthetics/settings/tests"
	if SyntheticTestResourcePath != expected {
		t.Errorf("Expected SyntheticTestResourcePath to be %s, got %s", expected, SyntheticTestResourcePath)
	}
}

func TestSyntheticTestGetIDForResourcePath(t *testing.T) {
	id := "test-synthetic-id"
	test := SyntheticTest{ID: id}

	result := test.GetIDForResourcePath()

	if result != id {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", id, result)
	}
}

func TestSyntheticTestStructure(t *testing.T) {
	id := "synthetic-test-id"
	label := "Test Synthetic"
	active := true

	test := SyntheticTest{
		ID:     id,
		Label:  label,
		Active: active,
	}

	if test.ID != id {
		t.Errorf("Expected ID to be %s, got %s", id, test.ID)
	}
	if test.Label != label {
		t.Errorf("Expected Label to be %s, got %s", label, test.Label)
	}
	if test.Active != active {
		t.Errorf("Expected Active to be %v, got %v", active, test.Active)
	}
}

func TestApiTagStructure(t *testing.T) {
	name := "environment"
	value := "production"

	tag := ApiTag{
		Name:  name,
		Value: value,
	}

	if tag.Name != name {
		t.Errorf("Expected Name to be %s, got %s", name, tag.Name)
	}
	if tag.Value != value {
		t.Errorf("Expected Value to be %s, got %s", value, tag.Value)
	}
}

func TestDNSFilterQueryTimeStructure(t *testing.T) {
	key := "queryTime"
	operator := ">"
	value := int64(100)

	filter := DNSFilterQueryTime{
		Key:      key,
		Operator: operator,
		Value:    value,
	}

	if filter.Key != key {
		t.Errorf("Expected Key to be %s, got %s", key, filter.Key)
	}
	if filter.Operator != operator {
		t.Errorf("Expected Operator to be %s, got %s", operator, filter.Operator)
	}
	if filter.Value != value {
		t.Errorf("Expected Value to be %d, got %d", value, filter.Value)
	}
}

// ---------------------------------------------------------------------------
// ICMPValidation struct
// ---------------------------------------------------------------------------

func TestICMPValidationStructure(t *testing.T) {
	rule := ICMPValidation{
		Key:      "packetLoss",
		Operator: "LESS_THAN",
		Value:    10,
	}

	if rule.Key != "packetLoss" {
		t.Errorf("Expected Key to be packetLoss, got %s", rule.Key)
	}
	if rule.Operator != "LESS_THAN" {
		t.Errorf("Expected Operator to be LESS_THAN, got %s", rule.Operator)
	}
	if rule.Value != 10 {
		t.Errorf("Expected Value to be 10, got %d", rule.Value)
	}
}

func TestICMPValidationJSONRoundtrip(t *testing.T) {
	original := ICMPValidation{
		Key:      "rtt",
		Operator: "LESS_THAN_OR_EQUALS",
		Value:    200,
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal ICMPValidation: %v", err)
	}

	var decoded ICMPValidation
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal ICMPValidation: %v", err)
	}

	if decoded.Key != original.Key {
		t.Errorf("Expected Key %s, got %s", original.Key, decoded.Key)
	}
	if decoded.Operator != original.Operator {
		t.Errorf("Expected Operator %s, got %s", original.Operator, decoded.Operator)
	}
	if decoded.Value != original.Value {
		t.Errorf("Expected Value %d, got %d", original.Value, decoded.Value)
	}
}

// ---------------------------------------------------------------------------
// SyntheticTestConfig – ICMP marshal/unmarshal
// ---------------------------------------------------------------------------

func TestSyntheticTestConfigICMPMarshal(t *testing.T) {
	targetHost := "example.com"
	packetCount := int32(5)
	packetSize := int32(64)
	packetInterval := "1s"
	packetTimeout := "2s"
	useDNS := true
	useIPv6 := false

	cfg := SyntheticTestConfig{
		MarkSyntheticCall: true,
		SyntheticType:     SyntheticTypeICMP,
		TargetHost:        &targetHost,
		PacketCount:       &packetCount,
		PacketSize:        &packetSize,
		PacketInterval:    &packetInterval,
		PacketTimeout:     &packetTimeout,
		UseDNS:            &useDNS,
		UseIPv6:           &useIPv6,
		ICMPValidationRules: []ICMPValidation{
			{Key: "packetLoss", Operator: "LESS_THAN", Value: 5},
			{Key: "rtt", Operator: "LESS_THAN_OR_EQUALS", Value: 100},
		},
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("Failed to marshal ICMP SyntheticTestConfig: %v", err)
	}

	// Verify the JSON contains expected keys
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("Failed to re-parse marshaled JSON: %v", err)
	}

	if raw["syntheticType"] != SyntheticTypeICMP {
		t.Errorf("Expected syntheticType %s, got %v", SyntheticTypeICMP, raw["syntheticType"])
	}
	if raw["targetHost"] != targetHost {
		t.Errorf("Expected targetHost %s, got %v", targetHost, raw["targetHost"])
	}
	rules, ok := raw["validationRules"].([]interface{})
	if !ok {
		t.Fatalf("Expected validationRules to be an array, got %T", raw["validationRules"])
	}
	if len(rules) != 2 {
		t.Errorf("Expected 2 validation rules, got %d", len(rules))
	}

	// icmpValidationRules must NOT appear in the output
	if _, exists := raw["icmpValidationRules"]; exists {
		t.Error("icmpValidationRules should not appear in marshaled JSON; expected validationRules")
	}
}

func TestSyntheticTestConfigICMPUnmarshal(t *testing.T) {
	jsonStr := `{
		"markSyntheticCall": true,
		"syntheticType": "ICMP",
		"targetHost": "192.0.2.1",
		"packetCount": 3,
		"packetSize": 32,
		"packetInterval": "500ms",
		"packetTimeout": "1s",
		"useDNS": false,
		"useIPv6": true,
		"validationRules": [
			{"key": "packetLoss", "operator": "EQUALS", "value": 0},
			{"key": "rtt", "operator": "LESS_THAN", "value": 50}
		]
	}`

	var cfg SyntheticTestConfig
	if err := json.Unmarshal([]byte(jsonStr), &cfg); err != nil {
		t.Fatalf("Failed to unmarshal ICMP SyntheticTestConfig: %v", err)
	}

	if cfg.SyntheticType != SyntheticTypeICMP {
		t.Errorf("Expected SyntheticType %s, got %s", SyntheticTypeICMP, cfg.SyntheticType)
	}
	if cfg.TargetHost == nil || *cfg.TargetHost != "192.0.2.1" {
		t.Errorf("Expected TargetHost 192.0.2.1, got %v", cfg.TargetHost)
	}
	if cfg.PacketCount == nil || *cfg.PacketCount != 3 {
		t.Errorf("Expected PacketCount 3, got %v", cfg.PacketCount)
	}
	if cfg.PacketSize == nil || *cfg.PacketSize != 32 {
		t.Errorf("Expected PacketSize 32, got %v", cfg.PacketSize)
	}
	if cfg.PacketInterval == nil || *cfg.PacketInterval != "500ms" {
		t.Errorf("Expected PacketInterval 500ms, got %v", cfg.PacketInterval)
	}
	if cfg.PacketTimeout == nil || *cfg.PacketTimeout != "1s" {
		t.Errorf("Expected PacketTimeout 1s, got %v", cfg.PacketTimeout)
	}
	if cfg.UseDNS == nil || *cfg.UseDNS != false {
		t.Errorf("Expected UseDNS false, got %v", cfg.UseDNS)
	}
	if cfg.UseIPv6 == nil || *cfg.UseIPv6 != true {
		t.Errorf("Expected UseIPv6 true, got %v", cfg.UseIPv6)
	}
	if len(cfg.ICMPValidationRules) != 2 {
		t.Fatalf("Expected 2 ICMPValidationRules, got %d", len(cfg.ICMPValidationRules))
	}
	if cfg.ICMPValidationRules[0].Key != "packetLoss" {
		t.Errorf("Expected first rule key packetLoss, got %s", cfg.ICMPValidationRules[0].Key)
	}
	if cfg.ICMPValidationRules[0].Operator != "EQUALS" {
		t.Errorf("Expected first rule operator EQUALS, got %s", cfg.ICMPValidationRules[0].Operator)
	}
	if cfg.ICMPValidationRules[0].Value != 0 {
		t.Errorf("Expected first rule value 0, got %d", cfg.ICMPValidationRules[0].Value)
	}
	if cfg.ICMPValidationRules[1].Key != "rtt" {
		t.Errorf("Expected second rule key rtt, got %s", cfg.ICMPValidationRules[1].Key)
	}
	// SSL ValidationRules must be empty when type is ICMP
	if len(cfg.ValidationRules) != 0 {
		t.Errorf("Expected SSL ValidationRules to be empty for ICMP type, got %d entries", len(cfg.ValidationRules))
	}
}

func TestSyntheticTestConfigICMPRoundtrip(t *testing.T) {
	targetHost := "ping.example.com"
	packetCount := int32(10)

	original := SyntheticTestConfig{
		MarkSyntheticCall: true,
		SyntheticType:     SyntheticTypeICMP,
		TargetHost:        &targetHost,
		PacketCount:       &packetCount,
		ICMPValidationRules: []ICMPValidation{
			{Key: "packetLoss", Operator: "LESS_THAN", Value: 5},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var decoded SyntheticTestConfig
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if decoded.SyntheticType != original.SyntheticType {
		t.Errorf("SyntheticType mismatch: got %s", decoded.SyntheticType)
	}
	if decoded.TargetHost == nil || *decoded.TargetHost != targetHost {
		t.Errorf("TargetHost mismatch: got %v", decoded.TargetHost)
	}
	if decoded.PacketCount == nil || *decoded.PacketCount != packetCount {
		t.Errorf("PacketCount mismatch: got %v", decoded.PacketCount)
	}
	if len(decoded.ICMPValidationRules) != 1 {
		t.Fatalf("Expected 1 ICMPValidationRule, got %d", len(decoded.ICMPValidationRules))
	}
	if decoded.ICMPValidationRules[0].Key != "packetLoss" {
		t.Errorf("Rule key mismatch: got %s", decoded.ICMPValidationRules[0].Key)
	}
	if decoded.ICMPValidationRules[0].Value != 5 {
		t.Errorf("Rule value mismatch: got %d", decoded.ICMPValidationRules[0].Value)
	}
}

func TestSyntheticTestConfigICMPNoValidationRules(t *testing.T) {
	targetHost := "192.0.2.1"

	cfg := SyntheticTestConfig{
		MarkSyntheticCall: true,
		SyntheticType:     SyntheticTypeICMP,
		TargetHost:        &targetHost,
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("Failed to re-parse: %v", err)
	}

	if _, exists := raw["validationRules"]; exists {
		t.Error("validationRules should be absent when ICMPValidationRules is empty")
	}
}

// ---------------------------------------------------------------------------
// SyntheticTestConfig – SSL validationRules still works after the refactor
// ---------------------------------------------------------------------------

func TestSyntheticTestConfigSSLValidationRulesRoundtrip(t *testing.T) {
	hostname := "example.com"
	daysRemaining := int32(30)
	sslPort := int32(443)

	original := SyntheticTestConfig{
		MarkSyntheticCall:  true,
		SyntheticType:      SyntheticTypeSSLCertificate,
		Hostname:           &hostname,
		DaysRemainingCheck: &daysRemaining,
		SSLPort:            &sslPort,
		ValidationRules: []SSLCertificateValidation{
			{Key: "expiryDays", Operator: "GREATER_THAN", Value: "30"},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Failed to marshal SSL config: %v", err)
	}

	var decoded SyntheticTestConfig
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal SSL config: %v", err)
	}

	if decoded.SyntheticType != SyntheticTypeSSLCertificate {
		t.Errorf("SyntheticType mismatch: got %s", decoded.SyntheticType)
	}
	if len(decoded.ValidationRules) != 1 {
		t.Fatalf("Expected 1 SSL ValidationRule, got %d", len(decoded.ValidationRules))
	}
	if decoded.ValidationRules[0].Key != "expiryDays" {
		t.Errorf("SSL rule key mismatch: got %s", decoded.ValidationRules[0].Key)
	}
	if decoded.ValidationRules[0].Value != "30" {
		t.Errorf("SSL rule value mismatch: got %s", decoded.ValidationRules[0].Value)
	}
	// ICMP validation rules must be empty
	if len(decoded.ICMPValidationRules) != 0 {
		t.Errorf("Expected ICMPValidationRules to be empty for SSL type, got %d entries", len(decoded.ICMPValidationRules))
	}
}

// ---------------------------------------------------------------------------
// SyntheticTypeICMP / SyntheticTypeSSLCertificate constants
// ---------------------------------------------------------------------------

func TestSyntheticTypeConstants(t *testing.T) {
	if SyntheticTypeICMP != "ICMP" {
		t.Errorf("Expected SyntheticTypeICMP to be ICMP, got %s", SyntheticTypeICMP)
	}
	if SyntheticTypeSSLCertificate != "SSLCertificate" {
		t.Errorf("Expected SyntheticTypeSSLCertificate to be SSLCertificate, got %s", SyntheticTypeSSLCertificate)
	}
}
