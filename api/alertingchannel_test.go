package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestAlertingChannelResourcePath(t *testing.T) {
	expected := "/api/events/settings/alertingChannels"
	if AlertingchannelResourcePath != expected {
		t.Errorf("Expected AlertingchannelResourcePath to be %s, got %s", expected, AlertingchannelResourcePath)
	}
}

func TestAlertingChannelTypeConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    AlertingChannelType
		expected string
	}{
		{"EmailChannelType", EmailChannelType, "EMAIL"},
		{"GoogleChatChannelType", GoogleChatChannelType, "GOOGLE_CHAT"},
		{"Office365ChannelType", Office365ChannelType, "OFFICE_365"},
		{"OpsGenieChannelType", OpsGenieChannelType, "OPS_GENIE"},
		{"PagerDutyChannelType", PagerDutyChannelType, "PAGER_DUTY"},
		{"SlackChannelType", SlackChannelType, "SLACK"},
		{"SplunkChannelType", SplunkChannelType, "SPLUNK"},
		{"VictorOpsChannelType", VictorOpsChannelType, "VICTOR_OPS"},
		{"WebhookChannelType", WebhookChannelType, "WEB_HOOK"},
		{"ServiceNowChannelType", ServiceNowChannelType, "SERVICE_NOW_WEBHOOK"},
		{"ServiceNowApplicationChannelType", ServiceNowApplicationChannelType, "SERVICE_NOW_APPLICATION"},
		{"PrometheusWebhookChannelType", PrometheusWebhookChannelType, "PROMETHEUS_WEBHOOK"},
		{"WebexTeamsWebhookChannelType", WebexTeamsWebhookChannelType, "WEBEX_TEAMS_WEBHOOK"},
		{"WatsonAIOpsWebhookChannelType", WatsonAIOpsWebhookChannelType, "WATSON_AIOPS_WEBHOOK"},
		{"SlackAppChannelType", SlackAppChannelType, "BIDIRECTIONAL_SLACK"},
		{"MsTeamsAppChannelType", MsTeamsAppChannelType, "BIDIRECTIONAL_MS_TEAMS"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, string(tt.value))
			}
		})
	}
}

func TestAlertingChannelGetIDForResourcePath(t *testing.T) {
	testID := "test-channel-id-123"
	channel := &AlertingChannel{
		ID:   testID,
		Name: "Test Channel",
		Kind: EmailChannelType,
	}

	result := channel.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestAlertingChannelStructure(t *testing.T) {
	webhookURL := "https://example.com/webhook"
	apiKey := "test-api-key"
	tags := "tag1,tag2"
	region := "us-east-1"
	routingKey := "routing-key-123"
	serviceIntegrationKey := "service-integration-key"
	iconURL := "https://example.com/icon.png"
	channel := "test-channel"
	url := "https://example.com"
	token := "test-token"
	serviceNowURL := "https://servicenow.example.com"
	username := "testuser"
	password := "testpass"
	autoCloseIncidents := true
	tenant := "test-tenant"
	unit := "test-unit"
	instanaURL := "https://instana.example.com"
	enableSendInstanaNotes := true
	enableSendServiceNowActivities := true
	enableSendServiceNowWorkNotes := true
	manuallyClosedIncidents := true
	resolutionOfIncident := true
	snowStatusOnCloseEvent := 6
	receiver := "test-receiver"
	appID := "app-123"
	teamID := "team-123"
	teamName := "Test Team"
	channelID := "channel-123"
	channelName := "Test Channel Name"
	emojiRendering := true
	apiTokenID := "api-token-123"
	serviceURL := "https://service.example.com"
	tenantID := "tenant-123"
	tenantName := "Test Tenant"

	alertingChannel := AlertingChannel{
		ID:                             "test-id",
		Name:                           "Test Alerting Channel",
		Kind:                           EmailChannelType,
		Emails:                         []string{"test@example.com", "admin@example.com"},
		WebhookURL:                     &webhookURL,
		APIKey:                         &apiKey,
		Tags:                           &tags,
		Region:                         &region,
		RoutingKey:                     &routingKey,
		ServiceIntegrationKey:          &serviceIntegrationKey,
		IconURL:                        &iconURL,
		Channel:                        &channel,
		URL:                            &url,
		Token:                          &token,
		WebhookURLs:                    []string{"https://webhook1.com", "https://webhook2.com"},
		Headers:                        []string{"Authorization: Bearer token", "Content-Type: application/json"},
		ServiceNowURL:                  &serviceNowURL,
		Username:                       &username,
		Password:                       &password,
		AutoCloseIncidents:             &autoCloseIncidents,
		Tenant:                         &tenant,
		Unit:                           &unit,
		InstanaURL:                     &instanaURL,
		EnableSendInstanaNotes:         &enableSendInstanaNotes,
		EnableSendServiceNowActivities: &enableSendServiceNowActivities,
		EnableSendServiceNowWorkNotes:  &enableSendServiceNowWorkNotes,
		ManuallyClosedIncidents:        &manuallyClosedIncidents,
		ResolutionOfIncident:           &resolutionOfIncident,
		SnowStatusOnCloseEvent:         &snowStatusOnCloseEvent,
		Receiver:                       &receiver,
		AppID:                          &appID,
		TeamID:                         &teamID,
		TeamName:                       &teamName,
		ChannelID:                      &channelID,
		ChannelName:                    &channelName,
		EmojiRendering:                 &emojiRendering,
		APITokenID:                     &apiTokenID,
		ServiceURL:                     &serviceURL,
		TenantID:                       &tenantID,
		TenantName:                     &tenantName,
	}

	// Test basic fields
	if alertingChannel.ID != "test-id" {
		t.Errorf("Expected ID to be 'test-id', got %s", alertingChannel.ID)
	}
	if alertingChannel.Name != "Test Alerting Channel" {
		t.Errorf("Expected Name to be 'Test Alerting Channel', got %s", alertingChannel.Name)
	}
	if alertingChannel.Kind != EmailChannelType {
		t.Errorf("Expected Kind to be EmailChannelType, got %s", alertingChannel.Kind)
	}
	if len(alertingChannel.Emails) != 2 {
		t.Errorf("Expected 2 emails, got %d", len(alertingChannel.Emails))
	}

	// Test pointer fields
	if alertingChannel.WebhookURL == nil || *alertingChannel.WebhookURL != webhookURL {
		t.Error("WebhookURL not set correctly")
	}
	if alertingChannel.APIKey == nil || *alertingChannel.APIKey != apiKey {
		t.Error("APIKey not set correctly")
	}
	if alertingChannel.AutoCloseIncidents == nil || *alertingChannel.AutoCloseIncidents != autoCloseIncidents {
		t.Error("AutoCloseIncidents not set correctly")
	}
	if alertingChannel.SnowStatusOnCloseEvent == nil || *alertingChannel.SnowStatusOnCloseEvent != snowStatusOnCloseEvent {
		t.Error("SnowStatusOnCloseEvent not set correctly")
	}
	if alertingChannel.EmojiRendering == nil || *alertingChannel.EmojiRendering != emojiRendering {
		t.Error("EmojiRendering not set correctly")
	}
}
