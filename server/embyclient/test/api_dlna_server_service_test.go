/*
Emby Server REST API (BETA)

Testing DlnaServerServiceAPIService

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

func Test_embyclient_DlnaServerServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidConnectionmanagerConnectionmanager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidConnectionmanagerConnectionmanager(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidConnectionmanagerConnectionmanagerXml", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidConnectionmanagerConnectionmanagerXml(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidContentdirectoryContentdirectory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidContentdirectoryContentdirectory(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidContentdirectoryContentdirectoryXml", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidContentdirectoryContentdirectoryXml(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidDescription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidDescription(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidDescriptionXml", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidDescriptionXml(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaByUuidIconsByFilename", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string
		var filename string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaByUuidIconsByFilename(context.Background(), uuId, filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService GetDlnaIconsByFilename", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filename string

		httpRes, err := apiClient.DlnaServerServiceAPI.GetDlnaIconsByFilename(context.Background(), filename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService HeadDlnaByUuidConnectionmanagerConnectionmanager", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.HeadDlnaByUuidConnectionmanagerConnectionmanager(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService HeadDlnaByUuidConnectionmanagerConnectionmanagerXml", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.HeadDlnaByUuidConnectionmanagerConnectionmanagerXml(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService HeadDlnaByUuidContentdirectoryContentdirectory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.HeadDlnaByUuidContentdirectoryContentdirectory(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService HeadDlnaByUuidContentdirectoryContentdirectoryXml", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.HeadDlnaByUuidContentdirectoryContentdirectoryXml(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService HeadDlnaByUuidDescription", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.HeadDlnaByUuidDescription(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService HeadDlnaByUuidDescriptionXml", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.HeadDlnaByUuidDescriptionXml(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService PostDlnaByUuidConnectionmanagerControl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.PostDlnaByUuidConnectionmanagerControl(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DlnaServerServiceAPIService PostDlnaByUuidContentdirectoryControl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var uuId string

		httpRes, err := apiClient.DlnaServerServiceAPI.PostDlnaByUuidContentdirectoryControl(context.Background(), uuId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
