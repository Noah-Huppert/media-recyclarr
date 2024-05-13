/*
Overseerr API

Testing PersonAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package jellyseerrclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/jellyseerrclient"
)

func Test_jellyseerrclient_PersonAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PersonAPIService PersonPersonIdCombinedCreditsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId float32

		resp, httpRes, err := apiClient.PersonAPI.PersonPersonIdCombinedCreditsGet(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PersonAPIService PersonPersonIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var personId float32

		resp, httpRes, err := apiClient.PersonAPI.PersonPersonIdGet(context.Background(), personId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
