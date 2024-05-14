/*
Emby Server REST API (BETA)

Testing InstantMixServiceAPIService

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

func Test_embyclient_InstantMixServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InstantMixServiceAPIService GetAlbumsByIdInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetAlbumsByIdInstantmix(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetArtistsInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetArtistsInstantmix(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetAudiobooksNextup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetAudiobooksNextup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetItemsByIdInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetItemsByIdInstantmix(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetMusicgenresByNameInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetMusicgenresByNameInstantmix(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetMusicgenresInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetMusicgenresInstantmix(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetPlaylistsByIdInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetPlaylistsByIdInstantmix(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test InstantMixServiceAPIService GetSongsByIdInstantmix", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.InstantMixServiceAPI.GetSongsByIdInstantmix(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
