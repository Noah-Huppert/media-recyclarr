/*
Overseerr API

Testing SearchAPIService

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

func Test_jellyseerrclient_SearchAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SearchAPIService DiscoverGenresliderMovieGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverGenresliderMovieGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverGenresliderTvGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverGenresliderTvGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverKeywordKeywordIdMoviesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var keywordId float32

		resp, httpRes, err := apiClient.SearchAPI.DiscoverKeywordKeywordIdMoviesGet(context.Background(), keywordId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverMoviesGenreGenreIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var genreId string

		resp, httpRes, err := apiClient.SearchAPI.DiscoverMoviesGenreGenreIdGet(context.Background(), genreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverMoviesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverMoviesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverMoviesLanguageLanguageGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string

		resp, httpRes, err := apiClient.SearchAPI.DiscoverMoviesLanguageLanguageGet(context.Background(), language).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverMoviesStudioStudioIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var studioId string

		resp, httpRes, err := apiClient.SearchAPI.DiscoverMoviesStudioStudioIdGet(context.Background(), studioId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverMoviesUpcomingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverMoviesUpcomingGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverTrendingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverTrendingGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverTvGenreGenreIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var genreId string

		resp, httpRes, err := apiClient.SearchAPI.DiscoverTvGenreGenreIdGet(context.Background(), genreId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverTvGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverTvGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverTvLanguageLanguageGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var language string

		resp, httpRes, err := apiClient.SearchAPI.DiscoverTvLanguageLanguageGet(context.Background(), language).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverTvNetworkNetworkIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var networkId string

		resp, httpRes, err := apiClient.SearchAPI.DiscoverTvNetworkNetworkIdGet(context.Background(), networkId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverTvUpcomingGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverTvUpcomingGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService DiscoverWatchlistGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.DiscoverWatchlistGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService SearchCompanyGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.SearchCompanyGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService SearchGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.SearchGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SearchAPIService SearchKeywordGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.SearchKeywordGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
