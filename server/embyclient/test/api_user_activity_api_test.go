/*
Emby Server REST API (BETA)

Testing UserActivityAPIAPIService

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

func Test_embyclient_UserActivityAPIAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsByBreakdowntypeBreakdownreport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var breakdownType string

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsByBreakdowntypeBreakdownreport(context.Background(), breakdownType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsByUseridByDateGetitems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userID string
		var date string

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsByUseridByDateGetitems(context.Background(), userID, date).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsGetItemPath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsGetItemPath(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsGetItemStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsGetItemStats(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsGetItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsGetItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsHourlyreport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsHourlyreport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsLoadBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsLoadBackup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsMoviesreport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsMoviesreport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsPlayactivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsPlayactivity(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsSaveBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsSaveBackup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsSessionList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsSessionList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsTvshowsreport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsTvshowsreport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsTypeFilterList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsTypeFilterList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsUserActivity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserActivity(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsUserList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsUserManageByActionById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var action string
		var id string

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserManageByActionById(context.Background(), action, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService GetUserUsageStatsUserplaylist", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.GetUserUsageStatsUserplaylist(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService PostUserUsageStatsImportBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.UserActivityAPIAPI.PostUserUsageStatsImportBackup(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserActivityAPIAPIService PostUserUsageStatsSubmitCustomQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.UserActivityAPIAPI.PostUserUsageStatsSubmitCustomQuery(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}