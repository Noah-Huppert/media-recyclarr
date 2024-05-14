/*
Emby Server REST API (BETA)

Testing LibraryStructureServiceAPIService

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

func Test_embyclient_LibraryStructureServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LibraryStructureServiceAPIService DeleteLibraryVirtualfolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.DeleteLibraryVirtualfolders(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService DeleteLibraryVirtualfoldersPaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.DeleteLibraryVirtualfoldersPaths(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService GetLibraryVirtualfoldersQuery", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfolders(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfoldersDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfoldersLibraryoptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersLibraryoptions(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfoldersName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersName(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfoldersPaths", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfoldersPathsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LibraryStructureServiceAPIService PostLibraryVirtualfoldersPathsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsUpdate(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
