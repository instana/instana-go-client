package api_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/instana/instana-go-client/api"
	"github.com/instana/instana-go-client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestIPFilteringGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	lastChanged := int64(1762356274224)
	expected := &api.IPFiltering{
		Active:               true,
		DenyAll:              false,
		Enabled:              true,
		SupportAccessEnabled: true,
		Rules: []api.IPFilteringRule{
			{Target: "127.0.0.1", Block: true},
			{Target: "0:0:0:0:0:0:0:1", Block: false},
		},
		LastChangedAt:  &lastChanged,
		LastVerifiedAt: nil,
	}

	payload, err := json.Marshal(expected)
	require.NoError(t, err)

	mockClient.EXPECT().Get(api.IPFilteringResourcePath).Return(payload, nil)

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Get()

	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestIPFilteringGet_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	mockClient.EXPECT().Get(api.IPFilteringResourcePath).Return(nil, errors.New("not found"))

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Get()

	assert.Nil(t, result)
	assert.EqualError(t, err, "not found")
}

func TestIPFilteringGet_EmptyBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	mockClient.EXPECT().Get(api.IPFilteringResourcePath).Return([]byte{}, nil)

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Get()

	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestIPFilteringUpsert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	input := &api.IPFiltering{
		DenyAll:              false,
		Enabled:              true,
		SupportAccessEnabled: true,
		Rules: []api.IPFilteringRule{
			{Target: "10.0.0.0/8", Block: false},
		},
	}

	responseObj := &api.IPFiltering{
		Active:               true,
		DenyAll:              false,
		Enabled:              true,
		SupportAccessEnabled: true,
		Rules: []api.IPFilteringRule{
			{Target: "10.0.0.0/8", Block: false},
		},
	}

	responsePayload, err := json.Marshal(responseObj)
	require.NoError(t, err)

	mockClient.EXPECT().
		Put(gomock.Any(), api.IPFilteringResourcePath).
		Return(responsePayload, nil)

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Upsert(input)

	require.NoError(t, err)
	assert.Equal(t, responseObj, result)
}

func TestIPFilteringUpsert_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	input := &api.IPFiltering{
		DenyAll: true,
		Enabled: true,
		Rules: []api.IPFilteringRule{
			{Target: "127.0.0.1", Block: true},
		},
	}

	mockClient.EXPECT().
		Put(gomock.Any(), api.IPFilteringResourcePath).
		Return(nil, errors.New("bad request"))

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Upsert(input)

	assert.Nil(t, result)
	assert.EqualError(t, err, "bad request")
}

func TestIPFilteringDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	mockClient.EXPECT().Delete("", api.IPFilteringResourcePath).Return(nil)

	resource := api.NewIPFilteringRestResource(mockClient)
	err := resource.Delete()

	require.NoError(t, err)
}

func TestIPFilteringVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)
	lastChanged := int64(1762356274224)
	lastVerified := int64(1762356274300)
	expected := &api.IPFiltering{
		Active:               true,
		DenyAll:              false,
		Enabled:              true,
		SupportAccessEnabled: true,
		Rules: []api.IPFilteringRule{
			{Target: "127.0.0.1", Block: true},
		},
		LastChangedAt:  &lastChanged,
		LastVerifiedAt: &lastVerified,
	}

	payload, err := json.Marshal(expected)
	require.NoError(t, err)

	mockClient.EXPECT().
		Patch(gomock.Any(), api.IPFilteringVerifyResourcePath).
		Return(payload, nil)

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Verify()

	require.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestIPFilteringVerify_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	mockClient.EXPECT().
		Patch(gomock.Any(), api.IPFilteringVerifyResourcePath).
		Return(nil, errors.New("failed to verify"))

	resource := api.NewIPFilteringRestResource(mockClient)
	result, err := resource.Verify()

	assert.Nil(t, result)
	assert.EqualError(t, err, "failed to verify")
}

func TestIPFilteringDelete_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRestClient(ctrl)

	mockClient.EXPECT().Delete("", api.IPFilteringResourcePath).Return(errors.New("forbidden"))

	resource := api.NewIPFilteringRestResource(mockClient)
	err := resource.Delete()

	assert.EqualError(t, err, "forbidden")
}
