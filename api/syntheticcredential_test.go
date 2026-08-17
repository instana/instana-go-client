package api_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/mocks"
	"github.com/instana/instana-go-client/shared/rest"
	gomock "go.uber.org/mock/gomock"
)

// ── constants ─────────────────────────────────────────────────────────────────

func TestSyntheticCredentialResourcePath(t *testing.T) {
	if api.SyntheticCredentialResourcePath != "/api/synthetics/settings/credentials" {
		t.Errorf("unexpected SyntheticCredentialResourcePath: %s", api.SyntheticCredentialResourcePath)
	}
}

func TestSyntheticCredentialAssociationsResourcePath(t *testing.T) {
	if api.SyntheticCredentialAssociationsResourcePath != "/api/synthetics/settings/credentials/associations" {
		t.Errorf("unexpected SyntheticCredentialAssociationsResourcePath: %s", api.SyntheticCredentialAssociationsResourcePath)
	}
}

// ── SyntheticCredential struct ────────────────────────────────────────────────

func TestSyntheticCredentialGetIDForResourcePath(t *testing.T) {
	cred := &api.SyntheticCredential{CredentialName: "my_token"}
	if got := cred.GetIDForResourcePath(); got != "my_token" {
		t.Errorf("GetIDForResourcePath: want %q, got %q", "my_token", got)
	}
}

func TestSyntheticCredentialGetIDForResourcePath_Empty(t *testing.T) {
	cred := &api.SyntheticCredential{}
	if got := cred.GetIDForResourcePath(); got != "" {
		t.Errorf("GetIDForResourcePath: want empty string, got %q", got)
	}
}

func TestSyntheticCredentialStructure(t *testing.T) {
	cred := api.SyntheticCredential{
		CredentialName:  "db_password",
		CredentialValue: "s3cr3t",
		Applications:    []string{"app-1", "app-2"},
		MobileApps:      []string{"mobile-1"},
		Websites:        []string{"web-1"},
		RbacTags: []api.RbacTag{
			{ID: "team-id", DisplayName: "Platform Team"},
		},
	}

	if cred.CredentialName != "db_password" {
		t.Errorf("CredentialName: want %q, got %q", "db_password", cred.CredentialName)
	}
	if len(cred.Applications) != 2 {
		t.Errorf("Applications: want 2, got %d", len(cred.Applications))
	}
	if len(cred.RbacTags) != 1 || cred.RbacTags[0].ID != "team-id" {
		t.Errorf("RbacTags: unexpected value %+v", cred.RbacTags)
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

func newSyntheticCredentialResource(t *testing.T, mockClient *mocks.MockRestClient) rest.RestResource[*api.SyntheticCredential] {
	t.Helper()
	return api.NewSyntheticCredentialRestResource(
		rest.NewGenericUnmarshaller[*api.SyntheticCredential](),
		mockClient,
	)
}

func marshalCredential(t *testing.T, cred *api.SyntheticCredential) []byte {
	t.Helper()
	data, err := json.Marshal(cred)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	return data
}

// ── GetAll ────────────────────────────────────────────────────────────────────

func TestSyntheticCredentialGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	credentials := []*api.SyntheticCredential{
		{CredentialName: "cred_one"},
		{CredentialName: "cred_two"},
	}
	payload, _ := json.Marshal(credentials)

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Get(api.SyntheticCredentialResourcePath).Return(payload, nil)

	resource := newSyntheticCredentialResource(t, mockClient)
	got, err := resource.GetAll()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || len(*got) != 2 {
		t.Fatalf("want 2 credentials, got %v", got)
	}
	if (*got)[0].CredentialName != "cred_one" {
		t.Errorf("first credential name: want %q, got %q", "cred_one", (*got)[0].CredentialName)
	}
}

func TestSyntheticCredentialGetAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Get(api.SyntheticCredentialResourcePath).Return(nil, errors.New("server error"))

	resource := newSyntheticCredentialResource(t, mockClient)
	_, err := resource.GetAll()

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// ── GetOne ────────────────────────────────────────────────────────────────────

func TestSyntheticCredentialGetOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := &api.SyntheticCredential{
		CredentialName: "api_token",
		Applications:   []string{"app-1"},
	}
	payload := marshalCredential(t, expected)

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().
		GetOne("api_token", api.SyntheticCredentialAssociationsResourcePath).
		Return(payload, nil)

	resource := newSyntheticCredentialResource(t, mockClient)
	got, err := resource.GetOne("api_token")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.CredentialName != expected.CredentialName {
		t.Errorf("CredentialName: want %q, got %q", expected.CredentialName, got.CredentialName)
	}
	if len(got.Applications) != 1 || got.Applications[0] != "app-1" {
		t.Errorf("Applications: want [app-1], got %v", got.Applications)
	}
}

func TestSyntheticCredentialGetOne_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().
		GetOne("api_token", api.SyntheticCredentialAssociationsResourcePath).
		Return(nil, errors.New("not found"))

	resource := newSyntheticCredentialResource(t, mockClient)
	_, err := resource.GetOne("api_token")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// ── Create ────────────────────────────────────────────────────────────────────

func TestSyntheticCredentialCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &api.SyntheticCredential{
		CredentialName:  "new_cred",
		CredentialValue: "s3cr3t",
	}
	// POST returns empty body (per API design); GetOne is called afterwards.
	readBack := &api.SyntheticCredential{CredentialName: "new_cred", Applications: []string{"app-1"}}
	readPayload := marshalCredential(t, readBack)

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Post(gomock.Any(), api.SyntheticCredentialResourcePath).Return(nil, nil)
	mockClient.EXPECT().
		GetOne("new_cred", api.SyntheticCredentialAssociationsResourcePath).
		Return(readPayload, nil)

	resource := newSyntheticCredentialResource(t, mockClient)
	got, err := resource.Create(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.CredentialName != "new_cred" {
		t.Errorf("CredentialName: want %q, got %q", "new_cred", got.CredentialName)
	}
}

func TestSyntheticCredentialCreate_PostError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &api.SyntheticCredential{CredentialName: "new_cred", CredentialValue: "s3cr3t"}

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Post(gomock.Any(), api.SyntheticCredentialResourcePath).Return(nil, errors.New("conflict"))

	resource := newSyntheticCredentialResource(t, mockClient)
	_, err := resource.Create(input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// ── Update ────────────────────────────────────────────────────────────────────

func TestSyntheticCredentialUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &api.SyntheticCredential{
		CredentialName:  "existing_cred",
		CredentialValue: "newvalue",
	}
	readBack := &api.SyntheticCredential{CredentialName: "existing_cred"}
	readPayload := marshalCredential(t, readBack)

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Put(gomock.Any(), api.SyntheticCredentialResourcePath).Return(nil, nil)
	mockClient.EXPECT().
		GetOne("existing_cred", api.SyntheticCredentialAssociationsResourcePath).
		Return(readPayload, nil)

	resource := newSyntheticCredentialResource(t, mockClient)
	got, err := resource.Update(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.CredentialName != "existing_cred" {
		t.Errorf("CredentialName: want %q, got %q", "existing_cred", got.CredentialName)
	}
}

func TestSyntheticCredentialUpdate_PutError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &api.SyntheticCredential{CredentialName: "existing_cred", CredentialValue: "v"}

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Put(gomock.Any(), api.SyntheticCredentialResourcePath).Return(nil, errors.New("server error"))

	resource := newSyntheticCredentialResource(t, mockClient)
	_, err := resource.Update(input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

// ── Delete ────────────────────────────────────────────────────────────────────

func TestSyntheticCredentialDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cred := &api.SyntheticCredential{CredentialName: "old_cred"}

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Delete("old_cred", api.SyntheticCredentialResourcePath).Return(nil)

	resource := newSyntheticCredentialResource(t, mockClient)
	if err := resource.Delete(cred); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSyntheticCredentialDelete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cred := &api.SyntheticCredential{CredentialName: "old_cred"}

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Delete("old_cred", api.SyntheticCredentialResourcePath).Return(errors.New("forbidden"))

	resource := newSyntheticCredentialResource(t, mockClient)
	if err := resource.Delete(cred); err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSyntheticCredentialDeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Delete("some_cred", api.SyntheticCredentialResourcePath).Return(nil)

	resource := newSyntheticCredentialResource(t, mockClient)
	if err := resource.DeleteByID("some_cred"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
