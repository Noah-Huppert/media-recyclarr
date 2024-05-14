/*
Emby Server REST API (BETA)

Testing ConnectServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package embyclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func Test_embyclient_ConnectServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConnectServiceAPIService DeleteUsersByIdConnectLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ConnectServiceAPI.DeleteUsersByIdConnectLink(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectServiceAPIService GetConnectExchange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ConnectServiceAPI.GetConnectExchange(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectServiceAPIService GetConnectPending", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ConnectServiceAPI.GetConnectPending(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectServiceAPIService PostUsersByIdConnectLink", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ConnectServiceAPI.PostUsersByIdConnectLink(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectServiceAPIService PostUsersByIdConnectLinkDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ConnectServiceAPI.PostUsersByIdConnectLinkDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
