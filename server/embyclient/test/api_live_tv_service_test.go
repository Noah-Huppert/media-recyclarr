/*
Emby Server REST API (BETA)

Testing LiveTvServiceAPIService

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

func Test_embyclient_LiveTvServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LiveTvServiceAPIService DeleteLivetvChannelmappingoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvChannelmappingoptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService DeleteLivetvChannelmappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvChannelmappings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService DeleteLivetvListingproviders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvListingproviders(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService DeleteLivetvRecordingsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvRecordingsById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService DeleteLivetvSeriestimersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvSeriestimersById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService DeleteLivetvTimersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvTimersById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService DeleteLivetvTunerhosts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.DeleteLivetvTunerhosts(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvAvailablerecordingoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvAvailablerecordingoptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvChannelmappingoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvChannelmappingoptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvChannelmappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvChannelmappings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvChannels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvChannels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvChannelsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvChannelsById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvChanneltags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvChanneltags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvChanneltagsPrefixes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvChanneltagsPrefixes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvEPG", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvEPG(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvFolder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvGuideinfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvGuideinfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvListingproviders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvListingproviders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvListingprovidersAvailable", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersAvailable(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvListingprovidersDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersDefault(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvListingprovidersLineups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersLineups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvListingprovidersSchedulesdirectCountries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersSchedulesdirectCountries(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvLiverecordingsByIdStream", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvLiverecordingsByIdStream(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvLivestreamfilesByIdStreamByContainer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var container string

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvLivestreamfilesByIdStreamByContainer(context.Background(), id, container).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvManageChannels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvManageChannels(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvPrograms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvPrograms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvProgramsRecommended", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvProgramsRecommended(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvRecordings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvRecordings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvRecordingsById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvRecordingsFolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsFolders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvRecordingsGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvRecordingsSeries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsSeries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvSeriestimers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvSeriestimers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvSeriestimersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvSeriestimersById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTimers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTimers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTimersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTimersById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTimersDefaults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTimersDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTunerhosts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTunerhosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTunerhostsDefaultByType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTunerhostsDefaultByType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTunerhostsTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTunerhostsTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService GetLivetvTunersDiscvover", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.GetLivetvTunersDiscvover(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService HeadLivetvChannelmappingoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.HeadLivetvChannelmappingoptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService HeadLivetvChannelmappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.HeadLivetvChannelmappings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvChannelmappingoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvChannelmappingoptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvChannelmappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvChannelmappings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvListingproviders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvListingproviders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvListingprovidersDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvListingprovidersDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvManageChannelsByIdDisabled", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvManageChannelsByIdDisabled(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvManageChannelsByIdSortindex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvManageChannelsByIdSortindex(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvPrograms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvPrograms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvRecordingsByIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvRecordingsByIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvSeriestimers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvSeriestimers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvSeriestimersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvSeriestimersById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvSeriestimersByIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvSeriestimersByIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvTimers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvTimers(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvTimersById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvTimersById(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvTimersByIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvTimersByIdDelete(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvTunerhosts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvTunerhosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvTunerhostsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvTunerhostsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PostLivetvTunersByIdReset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.LiveTvServiceAPI.PostLivetvTunersByIdReset(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PutLivetvChannelmappingoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PutLivetvChannelmappingoptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LiveTvServiceAPIService PutLivetvChannelmappings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LiveTvServiceAPI.PutLivetvChannelmappings(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
