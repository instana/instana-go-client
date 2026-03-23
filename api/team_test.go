package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestTeamsResourcePath(t *testing.T) {
	expected := "/api/settings/rbac/teams"
	if TeamsResourcePath != expected {
		t.Errorf("Expected TeamsResourcePath to be %s, got %s", expected, TeamsResourcePath)
	}
}

func TestTeamGetIDForResourcePath(t *testing.T) {
	testID := "team-id-123"
	team := &Team{
		ID:  testID,
		Tag: "test-team",
	}

	result := team.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestRestrictedApplicationFilterScopeConstants(t *testing.T) {
	tests := []struct {
		name     string
		value    RestrictedApplicationFilterScope
		expected string
	}{
		{"IncludeNoDownstream", RestrictedApplicationFilterScopeIncludeNoDownstream, "INCLUDE_NO_DOWNSTREAM"},
		{"IncludeImmediateDownstream", RestrictedApplicationFilterScopeIncludeImmediateDownstream, "INCLUDE_IMMEDIATE_DOWNSTREAM_DATABASE_AND_MESSAGING"},
		{"IncludeAllDownstream", RestrictedApplicationFilterScopeIncludeAllDownstream, "INCLUDE_ALL_DOWNSTREAM"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.value) != tt.expected {
				t.Errorf("Expected %s to be %s, got %s", tt.name, tt.expected, string(tt.value))
			}
		})
	}
}

func TestTeamStructure(t *testing.T) {
	description := "Test team description"
	roleName := "Admin"
	viaIdP := true
	email := "user@example.com"
	name := "Test User"

	team := Team{
		ID:  "team-123",
		Tag: "test-team-tag",
		Info: &TeamInfo{
			Description: &description,
		},
		Members: []TeamMember{
			{
				UserID: "user-1",
				Email:  &email,
				Name:   &name,
				Roles: []TeamRole{
					{
						RoleID:   "role-1",
						RoleName: &roleName,
						ViaIdP:   &viaIdP,
					},
				},
			},
		},
		Scope: &TeamScope{
			AccessPermissions:    []string{"perm1", "perm2"},
			Applications:         []string{"app1", "app2"},
			KubernetesClusters:   []string{"cluster1"},
			KubernetesNamespaces: []string{"namespace1"},
			MobileApps:           []string{"mobile1"},
			Websites:             []string{"website1"},
		},
	}

	if team.ID != "team-123" {
		t.Errorf("Expected ID 'team-123', got %s", team.ID)
	}
	if team.Tag != "test-team-tag" {
		t.Errorf("Expected Tag 'test-team-tag', got %s", team.Tag)
	}
	if team.Info == nil || team.Info.Description == nil {
		t.Error("TeamInfo or Description not set correctly")
	}
	if len(team.Members) != 1 {
		t.Errorf("Expected 1 member, got %d", len(team.Members))
	}
	if team.Scope == nil {
		t.Error("TeamScope should not be nil")
	}
}

func TestTeamMemberStructure(t *testing.T) {
	email := "member@example.com"
	name := "Team Member"
	roleName := "Developer"
	viaIdP := false

	member := TeamMember{
		UserID: "user-456",
		Email:  &email,
		Name:   &name,
		Roles: []TeamRole{
			{
				RoleID:   "role-dev",
				RoleName: &roleName,
				ViaIdP:   &viaIdP,
			},
		},
	}

	if member.UserID != "user-456" {
		t.Errorf("Expected UserID 'user-456', got %s", member.UserID)
	}
	if member.Email == nil || *member.Email != email {
		t.Error("Email not set correctly")
	}
	if len(member.Roles) != 1 {
		t.Errorf("Expected 1 role, got %d", len(member.Roles))
	}
}

func TestTeamScopeStructure(t *testing.T) {
	infraFilter := "entity.type:host"
	actionFilter := "action.type:restart"
	logFilter := "log.level:error"

	scope := TeamScope{
		AccessPermissions:    []string{"READ", "WRITE"},
		Applications:         []string{"app1", "app2", "app3"},
		KubernetesClusters:   []string{"cluster1", "cluster2"},
		KubernetesNamespaces: []string{"default", "production"},
		MobileApps:           []string{"mobile-app-1"},
		Websites:             []string{"website-1", "website-2"},
		InfraDFQFilter:       &infraFilter,
		ActionFilter:         &actionFilter,
		LogFilter:            &logFilter,
		BusinessPerspectives: []string{"bp1"},
		SloIDs:               []string{"slo1", "slo2"},
		SyntheticTests:       []string{"test1"},
		SyntheticCredentials: []string{"cred1"},
		TagIDs:               []string{"tag1", "tag2"},
	}

	if len(scope.AccessPermissions) != 2 {
		t.Errorf("Expected 2 access permissions, got %d", len(scope.AccessPermissions))
	}
	if len(scope.Applications) != 3 {
		t.Errorf("Expected 3 applications, got %d", len(scope.Applications))
	}
	if scope.InfraDFQFilter == nil || *scope.InfraDFQFilter != infraFilter {
		t.Error("InfraDFQFilter not set correctly")
	}
}

func TestRestrictedApplicationFilterStructure(t *testing.T) {
	label := "Test Filter"
	appID := "app-123"
	scope := RestrictedApplicationFilterScopeIncludeAllDownstream

	filter := RestrictedApplicationFilter{
		Label:                    &label,
		RestrictingApplicationID: &appID,
		Scope:                    &scope,
	}

	if filter.Label == nil || *filter.Label != label {
		t.Error("Label not set correctly")
	}
	if filter.RestrictingApplicationID == nil || *filter.RestrictingApplicationID != appID {
		t.Error("RestrictingApplicationID not set correctly")
	}
	if filter.Scope == nil || *filter.Scope != scope {
		t.Error("Scope not set correctly")
	}
}
