# \ItemRefreshServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostItemsByIdRefresh**](ItemRefreshServiceAPI.md#PostItemsByIdRefresh) | **Post** /Items/{Id}/Refresh | Refreshes metadata for an item



## PostItemsByIdRefresh

> PostItemsByIdRefresh(ctx, id).Recursive(recursive).MetadataRefreshMode(metadataRefreshMode).ImageRefreshMode(imageRefreshMode).ReplaceAllMetadata(replaceAllMetadata).ReplaceAllImages(replaceAllImages).Execute()

Refreshes metadata for an item



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Noah-Huppert/media-recyclarr/embyclient"
)

func main() {
	id := "id_example" // string | Item Id
	recursive := true // bool | Indicates if the refresh should occur recursively. (optional)
	metadataRefreshMode := openapiclient.MetadataRefreshMode("ValidationOnly") // ModelMetadataRefreshMode | Specifies the metadata refresh mode (optional)
	imageRefreshMode := openapiclient.MetadataRefreshMode("ValidationOnly") // ModelMetadataRefreshMode | Specifies the image refresh mode (optional)
	replaceAllMetadata := true // bool | Determines if metadata should be replaced. Only applicable if mode is FullRefresh (optional)
	replaceAllImages := true // bool | Determines if images should be replaced. Only applicable if mode is FullRefresh (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemRefreshServiceAPI.PostItemsByIdRefresh(context.Background(), id).Recursive(recursive).MetadataRefreshMode(metadataRefreshMode).ImageRefreshMode(imageRefreshMode).ReplaceAllMetadata(replaceAllMetadata).ReplaceAllImages(replaceAllImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefreshServiceAPI.PostItemsByIdRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByIdRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | Indicates if the refresh should occur recursively. | 
 **metadataRefreshMode** | [**ModelMetadataRefreshMode**](ModelMetadataRefreshMode.md) | Specifies the metadata refresh mode | 
 **imageRefreshMode** | [**ModelMetadataRefreshMode**](ModelMetadataRefreshMode.md) | Specifies the image refresh mode | 
 **replaceAllMetadata** | **bool** | Determines if metadata should be replaced. Only applicable if mode is FullRefresh | 
 **replaceAllImages** | **bool** | Determines if images should be replaced. Only applicable if mode is FullRefresh | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

