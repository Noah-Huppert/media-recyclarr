# \SubtitleServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItemsByIdSubtitlesByIndex**](SubtitleServiceAPI.md#DeleteItemsByIdSubtitlesByIndex) | **Delete** /Items/{Id}/Subtitles/{Index} | Deletes an external subtitle file
[**DeleteVideosByIdSubtitlesByIndex**](SubtitleServiceAPI.md#DeleteVideosByIdSubtitlesByIndex) | **Delete** /Videos/{Id}/Subtitles/{Index} | Deletes an external subtitle file
[**GetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat**](SubtitleServiceAPI.md#GetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat) | **Get** /Items/{Id}/{MediaSourceId}/Subtitles/{Index}/{StartPositionTicks}/Stream.{Format} | Gets subtitles in a specified format.
[**GetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormat**](SubtitleServiceAPI.md#GetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormat) | **Get** /Items/{Id}/{MediaSourceId}/Subtitles/{Index}/Stream.{Format} | Gets subtitles in a specified format.
[**GetItemsByIdRemotesearchSubtitlesByLanguage**](SubtitleServiceAPI.md#GetItemsByIdRemotesearchSubtitlesByLanguage) | **Get** /Items/{Id}/RemoteSearch/Subtitles/{Language} | 
[**GetProvidersSubtitlesSubtitlesById**](SubtitleServiceAPI.md#GetProvidersSubtitlesSubtitlesById) | **Get** /Providers/Subtitles/Subtitles/{Id} | 
[**GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat**](SubtitleServiceAPI.md#GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat) | **Get** /Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/{StartPositionTicks}/Stream.{Format} | Gets subtitles in a specified format.
[**GetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormat**](SubtitleServiceAPI.md#GetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormat) | **Get** /Videos/{Id}/{MediaSourceId}/Subtitles/{Index}/Stream.{Format} | Gets subtitles in a specified format.
[**PostItemsByIdRemotesearchSubtitlesBySubtitleid**](SubtitleServiceAPI.md#PostItemsByIdRemotesearchSubtitlesBySubtitleid) | **Post** /Items/{Id}/RemoteSearch/Subtitles/{SubtitleId} | 
[**PostItemsByIdSubtitlesByIndexDelete**](SubtitleServiceAPI.md#PostItemsByIdSubtitlesByIndexDelete) | **Post** /Items/{Id}/Subtitles/{Index}/Delete | Deletes an external subtitle file
[**PostVideosByIdSubtitlesByIndexDelete**](SubtitleServiceAPI.md#PostVideosByIdSubtitlesByIndexDelete) | **Post** /Videos/{Id}/Subtitles/{Index}/Delete | Deletes an external subtitle file



## DeleteItemsByIdSubtitlesByIndex

> DeleteItemsByIdSubtitlesByIndex(ctx, id, index).MediaSourceId(mediaSourceId).Execute()

Deletes an external subtitle file



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.DeleteItemsByIdSubtitlesByIndex(context.Background(), id, index).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.DeleteItemsByIdSubtitlesByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**index** | **int32** | The subtitle stream index | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemsByIdSubtitlesByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | MediaSourceId | 


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


## DeleteVideosByIdSubtitlesByIndex

> DeleteVideosByIdSubtitlesByIndex(ctx, id, index).MediaSourceId(mediaSourceId).Execute()

Deletes an external subtitle file



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.DeleteVideosByIdSubtitlesByIndex(context.Background(), id, index).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.DeleteVideosByIdSubtitlesByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**index** | **int32** | The subtitle stream index | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVideosByIdSubtitlesByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | MediaSourceId | 


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


## GetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat

> GetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat(ctx, id, mediaSourceId, index, format, startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()

Gets subtitles in a specified format.



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index
	format := "format_example" // string | Format
	startPositionTicks := int64(789) // int64 | StartPositionTicks
	endPositionTicks := int64(789) // int64 | EndPositionTicks (optional)
	copyTimestamps := true // bool | CopyTimestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.GetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat(context.Background(), id, mediaSourceId, index, format, startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**mediaSourceId** | **string** | MediaSourceId | 
**index** | **int32** | The subtitle stream index | 
**format** | **string** | Format | 
**startPositionTicks** | **int64** | StartPositionTicks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **endPositionTicks** | **int64** | EndPositionTicks | 
 **copyTimestamps** | **bool** | CopyTimestamps | 

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


## GetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormat

> GetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormat(ctx, id, mediaSourceId, index, format).StartPositionTicks(startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()

Gets subtitles in a specified format.



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index
	format := "format_example" // string | Format
	startPositionTicks := int64(789) // int64 | StartPositionTicks (optional)
	endPositionTicks := int64(789) // int64 | EndPositionTicks (optional)
	copyTimestamps := true // bool | CopyTimestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.GetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormat(context.Background(), id, mediaSourceId, index, format).StartPositionTicks(startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**mediaSourceId** | **string** | MediaSourceId | 
**index** | **int32** | The subtitle stream index | 
**format** | **string** | Format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdByMediasourceidSubtitlesByIndexStreamByFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **startPositionTicks** | **int64** | StartPositionTicks | 
 **endPositionTicks** | **int64** | EndPositionTicks | 
 **copyTimestamps** | **bool** | CopyTimestamps | 

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


## GetItemsByIdRemotesearchSubtitlesByLanguage

> []ModelModelRemoteSubtitleInfo GetItemsByIdRemotesearchSubtitlesByLanguage(ctx, id, language).MediaSourceId(mediaSourceId).IsPerfectMatch(isPerfectMatch).IsForced(isForced).IsHearingImpaired(isHearingImpaired).Execute()





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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	language := "language_example" // string | Language
	isPerfectMatch := true // bool | IsPerfectMatch (optional)
	isForced := true // bool | IsForced (optional)
	isHearingImpaired := true // bool | IsHearingImpaired (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleServiceAPI.GetItemsByIdRemotesearchSubtitlesByLanguage(context.Background(), id, language).MediaSourceId(mediaSourceId).IsPerfectMatch(isPerfectMatch).IsForced(isForced).IsHearingImpaired(isHearingImpaired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetItemsByIdRemotesearchSubtitlesByLanguage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdRemotesearchSubtitlesByLanguage`: []ModelModelRemoteSubtitleInfo
	fmt.Fprintf(os.Stdout, "Response from `SubtitleServiceAPI.GetItemsByIdRemotesearchSubtitlesByLanguage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**language** | **string** | Language | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdRemotesearchSubtitlesByLanguageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | MediaSourceId | 

 **isPerfectMatch** | **bool** | IsPerfectMatch | 
 **isForced** | **bool** | IsForced | 
 **isHearingImpaired** | **bool** | IsHearingImpaired | 

### Return type

[**[]ModelModelRemoteSubtitleInfo**](ModelRemoteSubtitleInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProvidersSubtitlesSubtitlesById

> GetProvidersSubtitlesSubtitlesById(ctx, id).Execute()





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
	r, err := apiClient.SubtitleServiceAPI.GetProvidersSubtitlesSubtitlesById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetProvidersSubtitlesSubtitlesById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetProvidersSubtitlesSubtitlesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat

> GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat(ctx, id, mediaSourceId, index, format, startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()

Gets subtitles in a specified format.



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index
	format := "format_example" // string | Format
	startPositionTicks := int64(789) // int64 | StartPositionTicks
	endPositionTicks := int64(789) // int64 | EndPositionTicks (optional)
	copyTimestamps := true // bool | CopyTimestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat(context.Background(), id, mediaSourceId, index, format, startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**mediaSourceId** | **string** | MediaSourceId | 
**index** | **int32** | The subtitle stream index | 
**format** | **string** | Format | 
**startPositionTicks** | **int64** | StartPositionTicks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdByMediasourceidSubtitlesByIndexByStartpositionticksStreamByFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **endPositionTicks** | **int64** | EndPositionTicks | 
 **copyTimestamps** | **bool** | CopyTimestamps | 

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


## GetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormat

> GetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormat(ctx, id, mediaSourceId, index, format).StartPositionTicks(startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()

Gets subtitles in a specified format.



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index
	format := "format_example" // string | Format
	startPositionTicks := int64(789) // int64 | StartPositionTicks (optional)
	endPositionTicks := int64(789) // int64 | EndPositionTicks (optional)
	copyTimestamps := true // bool | CopyTimestamps (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormat(context.Background(), id, mediaSourceId, index, format).StartPositionTicks(startPositionTicks).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.GetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**mediaSourceId** | **string** | MediaSourceId | 
**index** | **int32** | The subtitle stream index | 
**format** | **string** | Format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdByMediasourceidSubtitlesByIndexStreamByFormatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **startPositionTicks** | **int64** | StartPositionTicks | 
 **endPositionTicks** | **int64** | EndPositionTicks | 
 **copyTimestamps** | **bool** | CopyTimestamps | 

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


## PostItemsByIdRemotesearchSubtitlesBySubtitleid

> ModelModelSubtitlesSubtitleDownloadResult PostItemsByIdRemotesearchSubtitlesBySubtitleid(ctx, id, subtitleId).MediaSourceId(mediaSourceId).Execute()





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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	subtitleId := "subtitleId_example" // string | SubtitleId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleServiceAPI.PostItemsByIdRemotesearchSubtitlesBySubtitleid(context.Background(), id, subtitleId).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.PostItemsByIdRemotesearchSubtitlesBySubtitleid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsByIdRemotesearchSubtitlesBySubtitleid`: ModelModelSubtitlesSubtitleDownloadResult
	fmt.Fprintf(os.Stdout, "Response from `SubtitleServiceAPI.PostItemsByIdRemotesearchSubtitlesBySubtitleid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**subtitleId** | **string** | SubtitleId | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByIdRemotesearchSubtitlesBySubtitleidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | MediaSourceId | 


### Return type

[**ModelModelSubtitlesSubtitleDownloadResult**](ModelSubtitlesSubtitleDownloadResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsByIdSubtitlesByIndexDelete

> PostItemsByIdSubtitlesByIndexDelete(ctx, id, index).MediaSourceId(mediaSourceId).Execute()

Deletes an external subtitle file



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.PostItemsByIdSubtitlesByIndexDelete(context.Background(), id, index).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.PostItemsByIdSubtitlesByIndexDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**index** | **int32** | The subtitle stream index | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsByIdSubtitlesByIndexDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | MediaSourceId | 


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


## PostVideosByIdSubtitlesByIndexDelete

> PostVideosByIdSubtitlesByIndexDelete(ctx, id, index).MediaSourceId(mediaSourceId).Execute()

Deletes an external subtitle file



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
	mediaSourceId := "mediaSourceId_example" // string | MediaSourceId
	index := int32(56) // int32 | The subtitle stream index

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleServiceAPI.PostVideosByIdSubtitlesByIndexDelete(context.Background(), id, index).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleServiceAPI.PostVideosByIdSubtitlesByIndexDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 
**index** | **int32** | The subtitle stream index | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVideosByIdSubtitlesByIndexDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | MediaSourceId | 


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

