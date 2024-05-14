# \BifServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsByIdThumbnailset**](BifServiceAPI.md#GetItemsByIdThumbnailset) | **Get** /Items/{Id}/ThumbnailSet | 
[**GetVideosByIdIndexBif**](BifServiceAPI.md#GetVideosByIdIndexBif) | **Get** /Videos/{Id}/index.bif | 



## GetItemsByIdThumbnailset

> ModelModelRokuMetadataApiThumbnailSetInfo GetItemsByIdThumbnailset(ctx, id).Width(width).Execute()





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
	width := int32(56) // int32 | 
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BifServiceAPI.GetItemsByIdThumbnailset(context.Background(), id).Width(width).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BifServiceAPI.GetItemsByIdThumbnailset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThumbnailset`: ModelModelRokuMetadataApiThumbnailSetInfo
	fmt.Fprintf(os.Stdout, "Response from `BifServiceAPI.GetItemsByIdThumbnailset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThumbnailsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **int32** |  | 


### Return type

[**ModelModelRokuMetadataApiThumbnailSetInfo**](ModelRokuMetadataApiThumbnailSetInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideosByIdIndexBif

> GetVideosByIdIndexBif(ctx, id).Width(width).Execute()





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
	width := int32(56) // int32 | 
	id := "id_example" // string | Item Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BifServiceAPI.GetVideosByIdIndexBif(context.Background(), id).Width(width).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BifServiceAPI.GetVideosByIdIndexBif``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetVideosByIdIndexBifRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **width** | **int32** |  | 


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

