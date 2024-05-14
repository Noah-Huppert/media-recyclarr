# \RemoteImageServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImagesRemote**](RemoteImageServiceAPI.md#GetImagesRemote) | **Get** /Images/Remote | Gets a remote image
[**GetItemsByIdRemoteimages**](RemoteImageServiceAPI.md#GetItemsByIdRemoteimages) | **Get** /Items/{Id}/RemoteImages | Gets available remote images for an item
[**GetItemsByIdRemoteimagesProviders**](RemoteImageServiceAPI.md#GetItemsByIdRemoteimagesProviders) | **Get** /Items/{Id}/RemoteImages/Providers | Gets available remote image providers for an item
[**PostItemsByIdRemoteimagesDownload**](RemoteImageServiceAPI.md#PostItemsByIdRemoteimagesDownload) | **Post** /Items/{Id}/RemoteImages/Download | Downloads a remote image for an item



## GetImagesRemote

> GetImagesRemote(ctx).ImageUrl(imageUrl).Execute()

Gets a remote image



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
	imageUrl := "imageUrl_example" // string | The image url

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteImageServiceAPI.GetImagesRemote(context.Background()).ImageUrl(imageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageServiceAPI.GetImagesRemote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRemoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageUrl** | **string** | The image url | 

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


## GetItemsByIdRemoteimages

> ModelModelRemoteImageResult GetItemsByIdRemoteimages(ctx, id).Type_(type_).StartIndex(startIndex).Limit(limit).ProviderName(providerName).IncludeAllLanguages(includeAllLanguages).Execute()

Gets available remote images for an item



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
	type_ := openapiclient.ImageType("Primary") // ModelImageType | The image type (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	providerName := "providerName_example" // string | Optional. The image provider to use (optional)
	includeAllLanguages := true // bool | Optional. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteImageServiceAPI.GetItemsByIdRemoteimages(context.Background(), id).Type_(type_).StartIndex(startIndex).Limit(limit).ProviderName(providerName).IncludeAllLanguages(includeAllLanguages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageServiceAPI.GetItemsByIdRemoteimages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdRemoteimages`: ModelModelRemoteImageResult
	fmt.Fprintf(os.Stdout, "Response from `RemoteImageServiceAPI.GetItemsByIdRemoteimages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdRemoteimagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ModelImageType**](ModelImageType.md) | The image type | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **providerName** | **string** | Optional. The image provider to use | 
 **includeAllLanguages** | **bool** | Optional. | 

### Return type

[**ModelModelRemoteImageResult**](ModelRemoteImageResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdRemoteimagesProviders

> []ModelModelImageProviderInfo GetItemsByIdRemoteimagesProviders(ctx, id).Execute()

Gets available remote image providers for an item



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteImageServiceAPI.GetItemsByIdRemoteimagesProviders(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageServiceAPI.GetItemsByIdRemoteimagesProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdRemoteimagesProviders`: []ModelModelImageProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `RemoteImageServiceAPI.GetItemsByIdRemoteimagesProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdRemoteimagesProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelModelImageProviderInfo**](ModelImageProviderInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsByIdRemoteimagesDownload

> PostItemsByIdRemoteimagesDownload(ctx, id).Type_(type_).ProviderName(providerName).ImageUrl(imageUrl).Execute()

Downloads a remote image for an item



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
	type_ := openapiclient.ImageType("Primary") // ModelImageType | The image type
	providerName := "providerName_example" // string | The image provider (optional)
	imageUrl := "imageUrl_example" // string | The image url (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteImageServiceAPI.PostItemsByIdRemoteimagesDownload(context.Background(), id).Type_(type_).ProviderName(providerName).ImageUrl(imageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageServiceAPI.PostItemsByIdRemoteimagesDownload``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostItemsByIdRemoteimagesDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ModelImageType**](ModelImageType.md) | The image type | 
 **providerName** | **string** | The image provider | 
 **imageUrl** | **string** | The image url | 

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

