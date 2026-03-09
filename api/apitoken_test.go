package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestAPITokensResourcePath(t *testing.T) {
	expected := "/api/settings/api-tokens"
	if APITokensResourcePath != expected {
		t.Errorf("Expected APITokensResourcePath to be %s, got %s", expected, APITokensResourcePath)
	}
}

func TestAPITokenGetIDForResourcePath(t *testing.T) {
	testInternalID := "internal-id-123"
	token := &APIToken{
		ID:         "id-123",
		InternalID: testInternalID,
		Name:       "Test Token",
	}

	result := token.GetIDForResourcePath()
	if result != testInternalID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testInternalID, result)
	}
}

func TestAPITokenStructure(t *testing.T) {
	token := APIToken{
		ID:                                        "token-id-123",
		AccessGrantingToken:                       "access-token-xyz",
		InternalID:                                "internal-123",
		Name:                                      "Test API Token",
		CanConfigureServiceMapping:                true,
		CanConfigureEumApplications:               true,
		CanConfigureMobileAppMonitoring:           false,
		CanConfigureUsers:                         true,
		CanInstallNewAgents:                       false,
		CanConfigureIntegrations:                  true,
		CanConfigureEventsAndAlerts:               true,
		CanConfigureMaintenanceWindows:            false,
		CanConfigureApplicationSmartAlerts:        true,
		CanConfigureWebsiteSmartAlerts:            true,
		CanConfigureMobileAppSmartAlerts:          false,
		CanConfigureAPITokens:                     true,
		CanConfigureAgentRunMode:                  false,
		CanViewAuditLog:                           true,
		CanConfigureAgents:                        true,
		CanConfigureAuthenticationMethods:         false,
		CanConfigureApplications:                  true,
		CanConfigureTeams:                         true,
		CanConfigureReleases:                      false,
		CanConfigureLogManagement:                 true,
		CanCreatePublicCustomDashboards:           true,
		CanViewLogs:                               true,
		CanViewTraceDetails:                       true,
		CanConfigureSessionSettings:               false,
		CanConfigureGlobalAlertPayload:            true,
		CanConfigureGlobalApplicationSmartAlerts:  true,
		CanConfigureGlobalSyntheticSmartAlerts:    false,
		CanConfigureGlobalInfraSmartAlerts:        true,
		CanConfigureGlobalLogSmartAlerts:          true,
		CanViewAccountAndBillingInformation:       false,
		CanEditAllAccessibleCustomDashboards:      true,
		LimitedApplicationsScope:                  false,
		LimitedBizOpsScope:                        false,
		LimitedWebsitesScope:                      false,
		LimitedKubernetesScope:                    false,
		LimitedMobileAppsScope:                    false,
		LimitedInfrastructureScope:                false,
		LimitedSyntheticsScope:                    false,
		LimitedVsphereScope:                       false,
		LimitedPhmcScope:                          false,
		LimitedPvcScope:                           false,
		LimitedZhmcScope:                          false,
		LimitedPcfScope:                           false,
		LimitedOpenstackScope:                     false,
		LimitedAutomationScope:                    false,
		LimitedLogsScope:                          false,
		LimitedNutanixScope:                       false,
		LimitedXenServerScope:                     false,
		LimitedWindowsHypervisorScope:             false,
		LimitedAlertChannelsScope:                 false,
		LimitedLinuxKvmHypervisorScope:            false,
		LimitedAiGatewayScope:                     false,
		LimitedGenAIScope:                         false,
		LimitedServiceLevelScope:                  false,
		CanConfigurePersonalAPITokens:             true,
		CanConfigureDatabaseManagement:            false,
		CanConfigureAutomationActions:             true,
		CanConfigureAutomationPolicies:            true,
		CanRunAutomationActions:                   false,
		CanDeleteAutomationActionHistory:          false,
		CanConfigureSyntheticTests:                true,
		CanConfigureSyntheticLocations:            true,
		CanConfigureSyntheticCredentials:          false,
		CanViewSyntheticTests:                     true,
		CanViewSyntheticLocations:                 true,
		CanViewSyntheticTestResults:               true,
		CanUseSyntheticCredentials:                false,
		CanConfigureBizops:                        true,
		CanViewBusinessProcesses:                  true,
		CanViewBusinessProcessDetails:             true,
		CanViewBusinessActivities:                 true,
		CanViewBizAlerts:                          true,
		CanDeleteLogs:                             false,
		CanCreateHeapDump:                         false,
		CanCreateThreadDump:                       false,
		CanManuallyCloseIssue:                     true,
		CanViewLogVolume:                          true,
		CanConfigureLogRetentionPeriod:            false,
		CanConfigureSubtraces:                     false,
		CanInvokeAlertChannel:                     true,
		CanConfigureLlm:                           false,
		CanConfigureAiAgents:                      false,
		CanConfigureApdex:                         true,
		CanConfigureServiceLevelCorrectionWindows: true,
		CanConfigureServiceLevelSmartAlerts:       true,
		CanConfigureServiceLevels:                 true,
	}

	// Test basic fields
	if token.ID != "token-id-123" {
		t.Errorf("Expected ID 'token-id-123', got %s", token.ID)
	}
	if token.Name != "Test API Token" {
		t.Errorf("Expected Name 'Test API Token', got %s", token.Name)
	}
	if token.InternalID != "internal-123" {
		t.Errorf("Expected InternalID 'internal-123', got %s", token.InternalID)
	}

	// Test some permission fields
	if !token.CanConfigureServiceMapping {
		t.Error("Expected CanConfigureServiceMapping to be true")
	}
	if !token.CanConfigureApplications {
		t.Error("Expected CanConfigureApplications to be true")
	}
	if token.CanConfigureMobileAppMonitoring {
		t.Error("Expected CanConfigureMobileAppMonitoring to be false")
	}
}

func TestAPITokenPermissions(t *testing.T) {
	token := APIToken{
		ID:                         "test-id",
		InternalID:                 "test-internal-id",
		Name:                       "Test Token",
		CanConfigureApplications:   true,
		CanConfigureUsers:          false,
		CanViewLogs:                true,
		LimitedApplicationsScope:   false,
		LimitedKubernetesScope:     true,
		CanConfigureSyntheticTests: true,
	}

	// Verify permission values
	if !token.CanConfigureApplications {
		t.Error("Expected CanConfigureApplications to be true")
	}
	if token.CanConfigureUsers {
		t.Error("Expected CanConfigureUsers to be false")
	}
	if !token.CanViewLogs {
		t.Error("Expected CanViewLogs to be true")
	}
	if token.LimitedApplicationsScope {
		t.Error("Expected LimitedApplicationsScope to be false")
	}
	if !token.LimitedKubernetesScope {
		t.Error("Expected LimitedKubernetesScope to be true")
	}
	if !token.CanConfigureSyntheticTests {
		t.Error("Expected CanConfigureSyntheticTests to be true")
	}
}
