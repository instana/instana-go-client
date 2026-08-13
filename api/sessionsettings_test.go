package api_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/mocks"
	gomock "go.uber.org/mock/gomock"
)

func TestSessionSettingsGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := &api.SessionSettings{
		TokenLifeTimeInMillis: 28800000,
		IdleTimeInMillis:      3600000,
	}
	payload, _ := json.Marshal(expected)

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Get(api.SessionSettingsResourcePath).Return(payload, nil)

	resource := api.NewSessionSettingsRestResource(mockClient)
	got, err := resource.Get()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TokenLifeTimeInMillis != expected.TokenLifeTimeInMillis {
		t.Errorf("TokenLifeTimeInMillis: want %d, got %d", expected.TokenLifeTimeInMillis, got.TokenLifeTimeInMillis)
	}
	if got.IdleTimeInMillis != expected.IdleTimeInMillis {
		t.Errorf("IdleTimeInMillis: want %d, got %d", expected.IdleTimeInMillis, got.IdleTimeInMillis)
	}
}

func TestSessionSettingsGet_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Get(api.SessionSettingsResourcePath).Return(nil, errors.New("not found"))

	resource := api.NewSessionSettingsRestResource(mockClient)
	_, err := resource.Get()

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSessionSettingsUpsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &api.SessionSettings{
		TokenLifeTimeInMillis: 604800000,
		IdleTimeInMillis:      28800000,
	}
	payload, _ := json.Marshal(input)

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().
		Put(gomock.Any(), api.SessionSettingsResourcePath).
		Return(payload, nil)

	resource := api.NewSessionSettingsRestResource(mockClient)
	got, err := resource.Upsert(input)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.TokenLifeTimeInMillis != input.TokenLifeTimeInMillis {
		t.Errorf("TokenLifeTimeInMillis: want %d, got %d", input.TokenLifeTimeInMillis, got.TokenLifeTimeInMillis)
	}
}

func TestSessionSettingsUpsert_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	input := &api.SessionSettings{
		TokenLifeTimeInMillis: 28800000,
		IdleTimeInMillis:      3600000,
	}

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().
		Put(gomock.Any(), api.SessionSettingsResourcePath).
		Return(nil, errors.New("server error"))

	resource := api.NewSessionSettingsRestResource(mockClient)
	_, err := resource.Upsert(input)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSessionSettingsDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	// Delete is called with empty string ID because session settings has no ID
	mockClient.EXPECT().Delete("", api.SessionSettingsResourcePath).Return(nil)

	resource := api.NewSessionSettingsRestResource(mockClient)
	err := resource.Delete()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSessionSettingsDelete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Delete("", api.SessionSettingsResourcePath).Return(errors.New("forbidden"))

	resource := api.NewSessionSettingsRestResource(mockClient)
	err := resource.Delete()

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestSessionSettingsGet_EmptyBody_ReturnsDefaults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Simulate API returning 200 with empty body (no custom settings configured)
	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Get(api.SessionSettingsResourcePath).Return([]byte{}, nil)

	resource := api.NewSessionSettingsRestResource(mockClient)
	got, err := resource.Get()

	if err != nil {
		t.Fatalf("unexpected error for empty body: %v", err)
	}
	if got.TokenLifeTimeInMillis != 604800000 {
		t.Errorf("TokenLifeTimeInMillis: want 604800000 (default), got %d", got.TokenLifeTimeInMillis)
	}
	if got.IdleTimeInMillis != 28800000 {
		t.Errorf("IdleTimeInMillis: want 28800000 (default), got %d", got.IdleTimeInMillis)
	}
}

func TestSessionSettingsGet_NilBody_ReturnsDefaults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Simulate RestClient returning nil bytes (some HTTP client implementations return nil
	// rather than []byte{} on a 200 with no body content).
	mockClient := mocks.NewMockRestClient(ctrl)
	mockClient.EXPECT().Get(api.SessionSettingsResourcePath).Return(nil, nil)

	resource := api.NewSessionSettingsRestResource(mockClient)
	got, err := resource.Get()

	if err != nil {
		t.Fatalf("unexpected error for nil body: %v", err)
	}
	if got.TokenLifeTimeInMillis != 604800000 {
		t.Errorf("TokenLifeTimeInMillis: want 604800000 (default), got %d", got.TokenLifeTimeInMillis)
	}
	if got.IdleTimeInMillis != 28800000 {
		t.Errorf("IdleTimeInMillis: want 28800000 (default), got %d", got.IdleTimeInMillis)
	}
}
