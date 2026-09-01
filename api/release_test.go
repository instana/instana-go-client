package api_test

import (
	"testing"

	. "github.com/instana/instana-go-client/api"
)

func TestReleaseResourcePath(t *testing.T) {
	expected := "/api/releases"
	if ReleaseResourcePath != expected {
		t.Errorf("Expected ReleaseResourcePath to be %s, got %s", expected, ReleaseResourcePath)
	}
}

func TestReleaseWithMetadataGetIDForResourcePath(t *testing.T) {
	testID := "Tiu16hLCTniHDtHb_uDV1w"
	release := &ReleaseWithMetadata{
		ID:    testID,
		Name:  "demo-app/main-**",
		Start: 1709091782000,
	}

	result := release.GetIDForResourcePath()
	if result != testID {
		t.Errorf("Expected GetIDForResourcePath to return %s, got %s", testID, result)
	}
}

func TestReleaseWithMetadataStructure(t *testing.T) {
	appScope := &ReleaseApplicationScope{Name: "app1"}
	serviceScope := &ReleaseServiceScope{
		Name: "payment",
		ScopedTo: &ReleaseServiceScopedTo{
			Applications: []*ReleaseApplicationScope{{Name: "checkout-app"}},
		},
	}
	release := ReleaseWithMetadata{
		ID:           "XK1e1TF3T9SHKugndn_soQ",
		Name:         "frontend/release-2000",
		Start:        1706674621000,
		LastUpdated:  1706674621604,
		Applications: []*ReleaseApplicationScope{appScope},
		Services:     []*ReleaseServiceScope{serviceScope},
	}

	if release.ID != "XK1e1TF3T9SHKugndn_soQ" {
		t.Errorf("Expected ID 'XK1e1TF3T9SHKugndn_soQ', got %s", release.ID)
	}
	if release.Name != "frontend/release-2000" {
		t.Errorf("Expected Name 'frontend/release-2000', got %s", release.Name)
	}
	if release.Start != 1706674621000 {
		t.Errorf("Expected Start 1706674621000, got %d", release.Start)
	}
	if release.LastUpdated != 1706674621604 {
		t.Errorf("Expected LastUpdated 1706674621604, got %d", release.LastUpdated)
	}
	if len(release.Applications) != 1 || release.Applications[0].Name != "app1" {
		t.Errorf("Expected Applications[0].Name 'app1', got %v", release.Applications)
	}
	if len(release.Services) != 1 || release.Services[0].Name != "payment" {
		t.Errorf("Expected Services[0].Name 'payment', got %v", release.Services)
	}
	if len(release.Services[0].ScopedTo.Applications) != 1 || release.Services[0].ScopedTo.Applications[0].Name != "checkout-app" {
		t.Errorf("Expected Services[0].ScopedTo.Applications[0].Name 'checkout-app', got %v", release.Services[0].ScopedTo.Applications)
	}
}

func TestReleaseStructure(t *testing.T) {
	release := Release{
		Name:  "frontend/release-1000",
		Start: 1742349976000,
	}

	if release.Name != "frontend/release-1000" {
		t.Errorf("Expected Name 'frontend/release-1000', got %s", release.Name)
	}
	if release.Start != 1742349976000 {
		t.Errorf("Expected Start 1742349976000, got %d", release.Start)
	}
}
