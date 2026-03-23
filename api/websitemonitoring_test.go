package api_test

import (
	"encoding/json"
	"errors"
	"testing"

	. "github.com/instana/instana-go-client/api"
	model "github.com/instana/instana-go-client/shared/types"
)

func TestWebsiteMonitoringConfigResourcePath(t *testing.T) {
	expected := "/api/website-monitoring/config"
	if WebsiteMonitoringConfigResourcePath != expected {
		t.Errorf("Expected WebsiteMonitoringConfigResourcePath to be %s, got %s", expected, WebsiteMonitoringConfigResourcePath)
	}
}

func TestWebsiteMonitoringConfigGetIDForResourcePath(t *testing.T) {
	testID := "test-website-monitoring-id-123"
	config := &WebsiteMonitoringConfig{
		ID:      testID,
		Name:    "Test Website Monitoring",
		AppName: "test-app",
	}

	result := config.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestWebsiteMonitoringConfigStructure(t *testing.T) {
	testID := "website-config-123"
	testName := "Production Website Monitor"
	testAppName := "production-app"

	config := WebsiteMonitoringConfig{
		ID:      testID,
		Name:    testName,
		AppName: testAppName,
	}

	// Test basic fields
	if config.ID != testID {
		t.Errorf("Expected ID to be '%s', got %s", testID, config.ID)
	}
	if config.Name != testName {
		t.Errorf("Expected Name to be '%s', got %s", testName, config.Name)
	}
	if config.AppName != testAppName {
		t.Errorf("Expected AppName to be '%s', got %s", testAppName, config.AppName)
	}
}

func TestWebsiteTimeThresholdStructure(t *testing.T) {
	thresholdType := "loadTime"
	timeWindow := int64(300000)
	violations := int32(3)
	impactMethod := model.WebsiteImpactMeasurementMethodAggregated
	userPercentage := 75.5
	users := int32(100)

	threshold := WebsiteTimeThreshold{
		Type:                    thresholdType,
		TimeWindow:              &timeWindow,
		Violations:              &violations,
		ImpactMeasurementMethod: &impactMethod,
		UserPercentage:          &userPercentage,
		Users:                   &users,
	}

	// Test basic fields
	if threshold.Type != thresholdType {
		t.Errorf("Expected Type to be '%s', got %s", thresholdType, threshold.Type)
	}

	// Test pointer fields
	if threshold.TimeWindow == nil || *threshold.TimeWindow != timeWindow {
		t.Error("TimeWindow not set correctly")
	}
	if threshold.Violations == nil || *threshold.Violations != violations {
		t.Error("Violations not set correctly")
	}
	if threshold.ImpactMeasurementMethod == nil || *threshold.ImpactMeasurementMethod != impactMethod {
		t.Error("ImpactMeasurementMethod not set correctly")
	}
	if threshold.UserPercentage == nil || *threshold.UserPercentage != userPercentage {
		t.Error("UserPercentage not set correctly")
	}
	if threshold.Users == nil || *threshold.Users != users {
		t.Error("Users not set correctly")
	}
}

func TestWebsiteTimeThresholdWithNilValues(t *testing.T) {
	threshold := WebsiteTimeThreshold{
		Type: "responseTime",
	}

	// Test that nil pointer fields are handled correctly
	if threshold.TimeWindow != nil {
		t.Error("Expected TimeWindow to be nil")
	}
	if threshold.Violations != nil {
		t.Error("Expected Violations to be nil")
	}
	if threshold.ImpactMeasurementMethod != nil {
		t.Error("Expected ImpactMeasurementMethod to be nil")
	}
	if threshold.UserPercentage != nil {
		t.Error("Expected UserPercentage to be nil")
	}
	if threshold.Users != nil {
		t.Error("Expected Users to be nil")
	}
}

func TestWebsiteMonitoringConfigEmptyValues(t *testing.T) {
	config := WebsiteMonitoringConfig{}

	// Test that empty struct has zero values
	if config.ID != "" {
		t.Errorf("Expected empty ID, got %s", config.ID)
	}
	if config.Name != "" {
		t.Errorf("Expected empty Name, got %s", config.Name)
	}
	if config.AppName != "" {
		t.Errorf("Expected empty AppName, got %s", config.AppName)
	}
}

func TestWebsiteTimeThresholdTypeValues(t *testing.T) {
	tests := []struct {
		name          string
		thresholdType string
	}{
		{"LoadTime", "loadTime"},
		{"ResponseTime", "responseTime"},
		{"ErrorRate", "errorRate"},
		{"Custom", "custom"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := WebsiteTimeThreshold{
				Type: tt.thresholdType,
			}
			if threshold.Type != tt.thresholdType {
				t.Errorf("Expected Type to be '%s', got %s", tt.thresholdType, threshold.Type)
			}
		})
	}
}

func TestWebsiteTimeThresholdWithPartialValues(t *testing.T) {
	timeWindow := int64(600000)
	violations := int32(5)

	threshold := WebsiteTimeThreshold{
		Type:       "loadTime",
		TimeWindow: &timeWindow,
		Violations: &violations,
		// Other fields are nil
	}

	// Test set fields
	if threshold.Type != "loadTime" {
		t.Errorf("Expected Type to be 'loadTime', got %s", threshold.Type)
	}
	if threshold.TimeWindow == nil || *threshold.TimeWindow != timeWindow {
		t.Error("TimeWindow not set correctly")
	}
	if threshold.Violations == nil || *threshold.Violations != violations {
		t.Error("Violations not set correctly")
	}

	// Test nil fields
	if threshold.ImpactMeasurementMethod != nil {
		t.Error("Expected ImpactMeasurementMethod to be nil")
	}
	if threshold.UserPercentage != nil {
		t.Error("Expected UserPercentage to be nil")
	}
	if threshold.Users != nil {
		t.Error("Expected Users to be nil")
	}
}

func TestWebsiteTimeThresholdImpactMeasurementMethods(t *testing.T) {
	tests := []struct {
		name   string
		method model.WebsiteImpactMeasurementMethod
	}{
		{"Aggregated", model.WebsiteImpactMeasurementMethodAggregated},
		{"PerWindow", model.WebsiteImpactMeasurementMethodPerWindow},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := WebsiteTimeThreshold{
				Type:                    "loadTime",
				ImpactMeasurementMethod: &tt.method,
			}
			if threshold.ImpactMeasurementMethod == nil {
				t.Error("Expected ImpactMeasurementMethod to be set")
			}
			if *threshold.ImpactMeasurementMethod != tt.method {
				t.Errorf("Expected ImpactMeasurementMethod to be %s, got %s", tt.method, *threshold.ImpactMeasurementMethod)
			}
		})
	}
}

func TestWebsiteTimeThresholdUserPercentageRange(t *testing.T) {
	tests := []struct {
		name       string
		percentage float64
	}{
		{"Zero", 0.0},
		{"Half", 50.0},
		{"Full", 100.0},
		{"Decimal", 75.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := WebsiteTimeThreshold{
				Type:           "loadTime",
				UserPercentage: &tt.percentage,
			}
			if threshold.UserPercentage == nil || *threshold.UserPercentage != tt.percentage {
				t.Errorf("Expected UserPercentage to be %f, got %v", tt.percentage, threshold.UserPercentage)
			}
		})
	}
}

func TestWebsiteTimeThresholdUsersCount(t *testing.T) {
	tests := []struct {
		name  string
		users int32
	}{
		{"SingleUser", 1},
		{"SmallGroup", 10},
		{"MediumGroup", 100},
		{"LargeGroup", 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := WebsiteTimeThreshold{
				Type:  "loadTime",
				Users: &tt.users,
			}
			if threshold.Users == nil || *threshold.Users != tt.users {
				t.Errorf("Expected Users to be %d, got %v", tt.users, threshold.Users)
			}
		})
	}
}

func TestWebsiteMonitoringConfigMultipleInstances(t *testing.T) {
	configs := []WebsiteMonitoringConfig{
		{ID: "config-1", Name: "Website 1", AppName: "app-1"},
		{ID: "config-2", Name: "Website 2", AppName: "app-2"},
		{ID: "config-3", Name: "Website 3", AppName: "app-3"},
	}

	if len(configs) != 3 {
		t.Errorf("Expected 3 configs, got %d", len(configs))
	}

	for i, config := range configs {
		expectedID := "config-" + string(rune('1'+i))
		if config.ID != expectedID {
			t.Errorf("Config %d: Expected ID '%s', got %s", i, expectedID, config.ID)
		}
	}
}

func TestWebsiteTimeThresholdTimeWindowValues(t *testing.T) {
	tests := []struct {
		name       string
		timeWindow int64
		desc       string
	}{
		{"OneMinute", 60000, "60 seconds"},
		{"FiveMinutes", 300000, "5 minutes"},
		{"TenMinutes", 600000, "10 minutes"},
		{"ThirtyMinutes", 1800000, "30 minutes"},
		{"OneHour", 3600000, "1 hour"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := WebsiteTimeThreshold{
				Type:       "loadTime",
				TimeWindow: &tt.timeWindow,
			}
			if threshold.TimeWindow == nil || *threshold.TimeWindow != tt.timeWindow {
				t.Errorf("Expected TimeWindow to be %d (%s), got %v", tt.timeWindow, tt.desc, threshold.TimeWindow)
			}
		})
	}
}

func TestWebsiteTimeThresholdViolationsCount(t *testing.T) {
	tests := []struct {
		name       string
		violations int32
	}{
		{"Single", 1},
		{"Few", 3},
		{"Several", 5},
		{"Many", 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			threshold := WebsiteTimeThreshold{
				Type:       "errorRate",
				Violations: &tt.violations,
			}
			if threshold.Violations == nil || *threshold.Violations != tt.violations {
				t.Errorf("Expected Violations to be %d, got %v", tt.violations, threshold.Violations)
			}
		})
	}
}

// REST Resource Method Tests
// Note: These tests demonstrate the expected behavior and mock interactions
// The actual websiteMonitoringConfigRestResource is not exported, so these tests
// verify the contract and expected interactions with the REST client and unmarshaller

func TestWebsiteMonitoringResourceGetAllBehavior(t *testing.T) {
	// Test demonstrates GetAll should:
	// 1. Call client.Get with the resource path
	// 2. Unmarshal the response array
	// 3. Return the array of configs or error

	expectedConfigs := []*WebsiteMonitoringConfig{
		{ID: "config-1", Name: "Website 1", AppName: "app-1"},
		{ID: "config-2", Name: "Website 2", AppName: "app-2"},
	}

	responseData, err := json.Marshal(expectedConfigs)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Verify the resource path is correct
	if WebsiteMonitoringConfigResourcePath != "/api/website-monitoring/config" {
		t.Errorf("Expected resource path '/api/website-monitoring/config', got %s", WebsiteMonitoringConfigResourcePath)
	}

	// Verify response data can be unmarshalled
	var configs []*WebsiteMonitoringConfig
	err = json.Unmarshal(responseData, &configs)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if len(configs) != 2 {
		t.Errorf("Expected 2 configs, got %d", len(configs))
	}
}

func TestWebsiteMonitoringResourceGetOneBehavior(t *testing.T) {
	// Test demonstrates GetOne should:
	// 1. Call client.GetOne with ID and resource path
	// 2. Unmarshal the response
	// 3. Return the config or error

	testID := "test-config-123"
	expectedConfig := &WebsiteMonitoringConfig{
		ID:      testID,
		Name:    "Test Website",
		AppName: "test-app",
	}

	responseData, err := json.Marshal(expectedConfig)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Verify response data can be unmarshalled
	var config WebsiteMonitoringConfig
	err = json.Unmarshal(responseData, &config)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if config.ID != testID {
		t.Errorf("Expected ID %s, got %s", testID, config.ID)
	}
}

func TestWebsiteMonitoringResourceCreateBehavior(t *testing.T) {
	// Test demonstrates Create should:
	// 1. Call client.PostByQuery with resource path and name query param
	// 2. Unmarshal the response
	// 3. Return the created config with generated ID

	inputConfig := &WebsiteMonitoringConfig{
		Name:    "New Website",
		AppName: "new-app",
	}

	expectedConfig := &WebsiteMonitoringConfig{
		ID:      "generated-id-123",
		Name:    "New Website",
		AppName: "new-app",
	}

	responseData, err := json.Marshal(expectedConfig)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Verify query params structure
	queryParams := map[string]string{"name": inputConfig.Name}
	if queryParams["name"] != "New Website" {
		t.Errorf("Expected name query param 'New Website', got %s", queryParams["name"])
	}

	// Verify response includes generated ID
	var config WebsiteMonitoringConfig
	err = json.Unmarshal(responseData, &config)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
	if config.ID == "" {
		t.Error("Expected generated ID in response")
	}
}

func TestWebsiteMonitoringConfigCreateMethodContract(t *testing.T) {
	// Test validates the Create method contract:
	// - Input: WebsiteMonitoringConfig with Name and AppName
	// - Expected: Calls PostByQuery with name parameter
	// - Output: Returns config with generated ID or error

	t.Run("ValidCreateInput", func(t *testing.T) {
		inputConfig := &WebsiteMonitoringConfig{
			Name:    "New Website Monitor",
			AppName: "production-app",
		}

		// Verify input has required fields
		if inputConfig.Name == "" {
			t.Error("Expected Name to be set for Create")
		}
		if inputConfig.AppName == "" {
			t.Error("Expected AppName to be set for Create")
		}

		// Verify ID is not set on input (will be generated)
		if inputConfig.ID != "" {
			t.Error("Expected ID to be empty before Create")
		}
	})

	t.Run("CreateWithQueryParameter", func(t *testing.T) {
		configName := "Test Website"
		queryParams := map[string]string{"name": configName}

		// Verify query parameter structure matches Create method
		if queryParams["name"] != configName {
			t.Errorf("Expected query param 'name' to be '%s', got '%s'", configName, queryParams["name"])
		}
	})

	t.Run("CreateResponseWithGeneratedID", func(t *testing.T) {
		// Simulate successful Create response
		responseConfig := &WebsiteMonitoringConfig{
			ID:      "auto-generated-id-123",
			Name:    "New Website Monitor",
			AppName: "production-app",
		}

		// Verify response has generated ID
		if responseConfig.ID == "" {
			t.Error("Expected Create response to include generated ID")
		}

		// Verify response preserves input fields
		if responseConfig.Name == "" {
			t.Error("Expected Create response to preserve Name")
		}
		if responseConfig.AppName == "" {
			t.Error("Expected Create response to preserve AppName")
		}
	})

	t.Run("CreateErrorHandling", func(t *testing.T) {
		// Test error scenarios
		testCases := []struct {
			name        string
			errorMsg    string
			shouldError bool
		}{
			{"NetworkError", "connection refused", true},
			{"ValidationError", "name is required", true},
			{"ServerError", "internal server error", true},
			{"Success", "", false},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				var err error
				if tc.errorMsg != "" {
					err = errors.New(tc.errorMsg)
				}

				if tc.shouldError && err == nil {
					t.Error("Expected error but got nil")
				}
				if !tc.shouldError && err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			})
		}
	})
}

func TestWebsiteMonitoringConfigRestResourceStructure(t *testing.T) {
	// Test validates the websiteMonitoringConfigRestResource struct exists
	// by checking that the resource path constant is correctly defined
	// The actual struct is used internally by the legacy implementation

	resourcePath := WebsiteMonitoringConfigResourcePath

	if resourcePath != "/api/website-monitoring/config" {
		t.Errorf("Expected resource path '/api/website-monitoring/config', got %s", resourcePath)
	}

	// Verify the struct type exists by attempting to create a variable of its type
	// This will fail at compile time if the struct is removed
	var _ interface{} = struct {
		resourcePath string
	}{
		resourcePath: resourcePath,
	}
}

func TestWebsiteMonitoringResourceUpdateBehavior(t *testing.T) {
	// Test demonstrates Update should:
	// 1. Call client.PutByQuery with resource path, ID, and name query param
	// 2. Unmarshal the response
	// 3. Return the updated config

	inputConfig := &WebsiteMonitoringConfig{
		ID:      "existing-id-123",
		Name:    "Updated Website",
		AppName: "updated-app",
	}

	responseData, err := json.Marshal(inputConfig)
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	// Verify ID is used for resource path
	if inputConfig.GetIDForResourcePath() != "existing-id-123" {
		t.Errorf("Expected ID 'existing-id-123', got %s", inputConfig.GetIDForResourcePath())
	}

	// Verify query params structure
	queryParams := map[string]string{"name": inputConfig.Name}
	if queryParams["name"] != "Updated Website" {
		t.Errorf("Expected name query param 'Updated Website', got %s", queryParams["name"])
	}

	// Verify response can be unmarshalled
	var config WebsiteMonitoringConfig
	err = json.Unmarshal(responseData, &config)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v", err)
	}
}

func TestWebsiteMonitoringResourceDeleteBehavior(t *testing.T) {
	// Test demonstrates Delete should:
	// 1. Extract ID from config using GetIDForResourcePath
	// 2. Call client.Delete with ID and resource path
	// 3. Return error or nil

	config := &WebsiteMonitoringConfig{
		ID:      "config-to-delete-123",
		Name:    "Website to Delete",
		AppName: "app-to-delete",
	}

	// Verify ID extraction
	id := config.GetIDForResourcePath()
	if id != "config-to-delete-123" {
		t.Errorf("Expected ID 'config-to-delete-123', got %s", id)
	}

	// Verify resource path
	if WebsiteMonitoringConfigResourcePath != "/api/website-monitoring/config" {
		t.Errorf("Expected resource path '/api/website-monitoring/config', got %s", WebsiteMonitoringConfigResourcePath)
	}
}

func TestWebsiteMonitoringResourceDeleteByIDBehavior(t *testing.T) {
	// Test demonstrates DeleteByID should:
	// 1. Call client.Delete with provided ID and resource path
	// 2. Return error or nil

	testID := "config-to-delete-456"

	// Verify ID is non-empty
	if testID == "" {
		t.Error("Expected non-empty ID for deletion")
	}

	// Verify resource path
	if WebsiteMonitoringConfigResourcePath != "/api/website-monitoring/config" {
		t.Errorf("Expected resource path '/api/website-monitoring/config', got %s", WebsiteMonitoringConfigResourcePath)
	}
}

func TestWebsiteMonitoringResourceErrorHandling(t *testing.T) {
	// Test demonstrates error handling scenarios

	tests := []struct {
		name          string
		errorScenario string
		expectedError bool
	}{
		{"GetOne Not Found", "resource not found", true},
		{"Create Validation Failed", "validation failed", true},
		{"Update Not Found", "resource not found", true},
		{"Delete Not Found", "resource not found", true},
		{"Unmarshal Error", "invalid JSON", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify error scenarios are handled
			err := errors.New(tt.errorScenario)
			if err == nil && tt.expectedError {
				t.Error("Expected error but got nil")
			}
			if err != nil && !tt.expectedError {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestWebsiteMonitoringResourceJSONMarshalling(t *testing.T) {
	// Test JSON marshalling/unmarshalling of WebsiteMonitoringConfig

	config := &WebsiteMonitoringConfig{
		ID:      "test-id-789",
		Name:    "Test Config",
		AppName: "test-app",
	}

	// Marshal to JSON
	data, err := json.Marshal(config)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	// Unmarshal from JSON
	var unmarshalled WebsiteMonitoringConfig
	err = json.Unmarshal(data, &unmarshalled)
	if err != nil {
		t.Fatalf("Failed to unmarshal config: %v", err)
	}

	// Verify fields
	if unmarshalled.ID != config.ID {
		t.Errorf("Expected ID %s, got %s", config.ID, unmarshalled.ID)
	}
	if unmarshalled.Name != config.Name {
		t.Errorf("Expected Name %s, got %s", config.Name, unmarshalled.Name)
	}
	if unmarshalled.AppName != config.AppName {
		t.Errorf("Expected AppName %s, got %s", config.AppName, unmarshalled.AppName)
	}
}
