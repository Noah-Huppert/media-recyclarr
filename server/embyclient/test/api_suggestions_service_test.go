/*
Emby Server REST API (BETA)

Testing SuggestionsServiceAPIService

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

func Test_embyclient_SuggestionsServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SuggestionsServiceAPIService GetUsersByUseridSuggestions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string

		resp, httpRes, err := apiClient.SuggestionsServiceAPI.GetUsersByUseridSuggestions(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
