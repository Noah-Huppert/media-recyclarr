# \MediaInfoServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsByIdPlaybackinfo**](MediaInfoServiceAPI.md#GetItemsByIdPlaybackinfo) | **Get** /Items/{Id}/PlaybackInfo | Gets live playback media info for an item
[**GetPlaybackBitratetest**](MediaInfoServiceAPI.md#GetPlaybackBitratetest) | **Get** /Playback/BitrateTest | 
[**PostItemsByIdPlaybackinfo**](MediaInfoServiceAPI.md#PostItemsByIdPlaybackinfo) | **Post** /Items/{Id}/PlaybackInfo | Gets live playback media info for an item
[**PostLivestreamsClose**](MediaInfoServiceAPI.md#PostLivestreamsClose) | **Post** /LiveStreams/Close | Closes a media source
[**PostLivestreamsMediainfo**](MediaInfoServiceAPI.md#PostLivestreamsMediainfo) | **Post** /LiveStreams/MediaInfo | Closes a media source
[**PostLivestreamsOpen**](MediaInfoServiceAPI.md#PostLivestreamsOpen) | **Post** /LiveStreams/Open | Opens a media source



## GetItemsByIdPlaybackinfo

> ModelModelPlaybackInfoResponse GetItemsByIdPlaybackinfo(ctx, id).UserId(userId).Execute()

Gets live playback media info for an item



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
	userId := "userId_example" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoServiceAPI.GetItemsByIdPlaybackinfo(context.Background(), id).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoServiceAPI.GetItemsByIdPlaybackinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdPlaybackinfo`: ModelModelPlaybackInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoServiceAPI.GetItemsByIdPlaybackinfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdPlaybackinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 

### Return type

[**ModelModelPlaybackInfoResponse**](ModelPlaybackInfoResponse.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaybackBitratetest

> GetPlaybackBitratetest(ctx).Size(size).Execute()





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
	size := int64(789) // int64 | Size

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaInfoServiceAPI.GetPlaybackBitratetest(context.Background()).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoServiceAPI.GetPlaybackBitratetest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaybackBitratetestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int64** | Size | 

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


## PostItemsByIdPlaybackinfo

> ModelModelPlaybackInfoResponse PostItemsByIdPlaybackinfo(ctx, id).ModelPlaybackInfoRequest(modelPlaybackInfoRequest).Execute()

Gets live playback media info for an item



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
	id := "id_example" // string | 
	modelPlaybackInfoRequest := *openapiclient.NewModelPlaybackInfoRequest() // ModelPlaybackInfoRequest | PlaybackInfoRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoServiceAPI.PostItemsByIdPlaybackinfo(context.Background(), id).ModelPlaybackInfoRequest(modelPlaybackInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoServiceAPI.PostItemsByIdPlaybackinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsByIdPlaybackinfo`: ModelModelPlaybackInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoServiceAPI.PostItemsByIdPlaybackinfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByIdPlaybackinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelPlaybackInfoRequest** | [**ModelPlaybackInfoRequest**](ModelPlaybackInfoRequest.md) | PlaybackInfoRequest:  | 

### Return type

[**ModelModelPlaybackInfoResponse**](ModelPlaybackInfoResponse.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivestreamsClose

> PostLivestreamsClose(ctx).LiveStreamId(liveStreamId).Execute()

Closes a media source



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
	liveStreamId := "liveStreamId_example" // string | LiveStreamId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaInfoServiceAPI.PostLivestreamsClose(context.Background()).LiveStreamId(liveStreamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoServiceAPI.PostLivestreamsClose``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivestreamsCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveStreamId** | **string** | LiveStreamId | 

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


## PostLivestreamsMediainfo

> PostLivestreamsMediainfo(ctx).LiveStreamId(liveStreamId).Execute()

Closes a media source



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
	liveStreamId := "liveStreamId_example" // string | LiveStreamId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaInfoServiceAPI.PostLivestreamsMediainfo(context.Background()).LiveStreamId(liveStreamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoServiceAPI.PostLivestreamsMediainfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivestreamsMediainfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveStreamId** | **string** | LiveStreamId | 

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


## PostLivestreamsOpen

> ModelModelLiveStreamResponse PostLivestreamsOpen(ctx).ModelLiveStreamRequest(modelLiveStreamRequest).Execute()

Opens a media source



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
	modelLiveStreamRequest := *openapiclient.NewModelLiveStreamRequest() // ModelLiveStreamRequest | LiveStreamRequest: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoServiceAPI.PostLivestreamsOpen(context.Background()).ModelLiveStreamRequest(modelLiveStreamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoServiceAPI.PostLivestreamsOpen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLivestreamsOpen`: ModelModelLiveStreamResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoServiceAPI.PostLivestreamsOpen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivestreamsOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLiveStreamRequest** | [**ModelLiveStreamRequest**](ModelLiveStreamRequest.md) | LiveStreamRequest:  | 

### Return type

[**ModelModelLiveStreamResponse**](ModelLiveStreamResponse.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

