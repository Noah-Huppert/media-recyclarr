# \LiveTvServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLivetvChannelmappingoptions**](LiveTvServiceAPI.md#DeleteLivetvChannelmappingoptions) | **Delete** /LiveTv/ChannelMappingOptions | 
[**DeleteLivetvChannelmappings**](LiveTvServiceAPI.md#DeleteLivetvChannelmappings) | **Delete** /LiveTv/ChannelMappings | 
[**DeleteLivetvListingproviders**](LiveTvServiceAPI.md#DeleteLivetvListingproviders) | **Delete** /LiveTv/ListingProviders | Deletes a listing provider
[**DeleteLivetvRecordingsById**](LiveTvServiceAPI.md#DeleteLivetvRecordingsById) | **Delete** /LiveTv/Recordings/{Id} | Deletes a live tv recording
[**DeleteLivetvSeriestimersById**](LiveTvServiceAPI.md#DeleteLivetvSeriestimersById) | **Delete** /LiveTv/SeriesTimers/{Id} | Cancels a live tv series timer
[**DeleteLivetvTimersById**](LiveTvServiceAPI.md#DeleteLivetvTimersById) | **Delete** /LiveTv/Timers/{Id} | Cancels a live tv timer
[**DeleteLivetvTunerhosts**](LiveTvServiceAPI.md#DeleteLivetvTunerhosts) | **Delete** /LiveTv/TunerHosts | Deletes a tuner host
[**GetLivetvAvailablerecordingoptions**](LiveTvServiceAPI.md#GetLivetvAvailablerecordingoptions) | **Get** /LiveTv/AvailableRecordingOptions | Gets available recording options
[**GetLivetvChannelmappingoptions**](LiveTvServiceAPI.md#GetLivetvChannelmappingoptions) | **Get** /LiveTv/ChannelMappingOptions | 
[**GetLivetvChannelmappings**](LiveTvServiceAPI.md#GetLivetvChannelmappings) | **Get** /LiveTv/ChannelMappings | 
[**GetLivetvChannels**](LiveTvServiceAPI.md#GetLivetvChannels) | **Get** /LiveTv/Channels | Gets available live tv channels.
[**GetLivetvChannelsById**](LiveTvServiceAPI.md#GetLivetvChannelsById) | **Get** /LiveTv/Channels/{Id} | Gets a live tv channel
[**GetLivetvChanneltags**](LiveTvServiceAPI.md#GetLivetvChanneltags) | **Get** /LiveTv/ChannelTags | Gets live tv channel tags
[**GetLivetvChanneltagsPrefixes**](LiveTvServiceAPI.md#GetLivetvChanneltagsPrefixes) | **Get** /LiveTv/ChannelTags/Prefixes | Gets live tv channel tag prefixes
[**GetLivetvEPG**](LiveTvServiceAPI.md#GetLivetvEPG) | **Get** /LiveTv/EPG | Gets the epg.
[**GetLivetvFolder**](LiveTvServiceAPI.md#GetLivetvFolder) | **Get** /LiveTv/Folder | Gets the top level live tv folder
[**GetLivetvGuideinfo**](LiveTvServiceAPI.md#GetLivetvGuideinfo) | **Get** /LiveTv/GuideInfo | Gets guide info
[**GetLivetvInfo**](LiveTvServiceAPI.md#GetLivetvInfo) | **Get** /LiveTv/Info | Gets available live tv services.
[**GetLivetvListingproviders**](LiveTvServiceAPI.md#GetLivetvListingproviders) | **Get** /LiveTv/ListingProviders | Gets current listing providers
[**GetLivetvListingprovidersAvailable**](LiveTvServiceAPI.md#GetLivetvListingprovidersAvailable) | **Get** /LiveTv/ListingProviders/Available | Gets listing provider
[**GetLivetvListingprovidersDefault**](LiveTvServiceAPI.md#GetLivetvListingprovidersDefault) | **Get** /LiveTv/ListingProviders/Default | 
[**GetLivetvListingprovidersLineups**](LiveTvServiceAPI.md#GetLivetvListingprovidersLineups) | **Get** /LiveTv/ListingProviders/Lineups | Gets available lineups
[**GetLivetvListingprovidersSchedulesdirectCountries**](LiveTvServiceAPI.md#GetLivetvListingprovidersSchedulesdirectCountries) | **Get** /LiveTv/ListingProviders/SchedulesDirect/Countries | Gets available lineups
[**GetLivetvLiverecordingsByIdStream**](LiveTvServiceAPI.md#GetLivetvLiverecordingsByIdStream) | **Get** /LiveTv/LiveRecordings/{Id}/stream | Gets a live tv channel
[**GetLivetvLivestreamfilesByIdStreamByContainer**](LiveTvServiceAPI.md#GetLivetvLivestreamfilesByIdStreamByContainer) | **Get** /LiveTv/LiveStreamFiles/{Id}/stream.{Container} | Gets a live tv channel
[**GetLivetvManageChannels**](LiveTvServiceAPI.md#GetLivetvManageChannels) | **Get** /LiveTv/Manage/Channels | Gets the channel management list
[**GetLivetvPrograms**](LiveTvServiceAPI.md#GetLivetvPrograms) | **Get** /LiveTv/Programs | Gets available live tv epgs..
[**GetLivetvProgramsRecommended**](LiveTvServiceAPI.md#GetLivetvProgramsRecommended) | **Get** /LiveTv/Programs/Recommended | Gets available live tv epgs..
[**GetLivetvRecordings**](LiveTvServiceAPI.md#GetLivetvRecordings) | **Get** /LiveTv/Recordings | Gets live tv recordings
[**GetLivetvRecordingsById**](LiveTvServiceAPI.md#GetLivetvRecordingsById) | **Get** /LiveTv/Recordings/{Id} | Gets a live tv recording
[**GetLivetvRecordingsFolders**](LiveTvServiceAPI.md#GetLivetvRecordingsFolders) | **Get** /LiveTv/Recordings/Folders | Gets recording folders
[**GetLivetvRecordingsGroups**](LiveTvServiceAPI.md#GetLivetvRecordingsGroups) | **Get** /LiveTv/Recordings/Groups | Gets live tv recording groups
[**GetLivetvRecordingsSeries**](LiveTvServiceAPI.md#GetLivetvRecordingsSeries) | **Get** /LiveTv/Recordings/Series | Gets live tv recordings
[**GetLivetvSeriestimers**](LiveTvServiceAPI.md#GetLivetvSeriestimers) | **Get** /LiveTv/SeriesTimers | Gets live tv series timers
[**GetLivetvSeriestimersById**](LiveTvServiceAPI.md#GetLivetvSeriestimersById) | **Get** /LiveTv/SeriesTimers/{Id} | Gets a live tv series timer
[**GetLivetvTimers**](LiveTvServiceAPI.md#GetLivetvTimers) | **Get** /LiveTv/Timers | Gets live tv timers
[**GetLivetvTimersById**](LiveTvServiceAPI.md#GetLivetvTimersById) | **Get** /LiveTv/Timers/{Id} | Gets a live tv timer
[**GetLivetvTimersDefaults**](LiveTvServiceAPI.md#GetLivetvTimersDefaults) | **Get** /LiveTv/Timers/Defaults | Gets default values for a new timer
[**GetLivetvTunerhosts**](LiveTvServiceAPI.md#GetLivetvTunerhosts) | **Get** /LiveTv/TunerHosts | Gets tuner hosts
[**GetLivetvTunerhostsDefaultByType**](LiveTvServiceAPI.md#GetLivetvTunerhostsDefaultByType) | **Get** /LiveTv/TunerHosts/Default/{Type} | Gets tuner hosts
[**GetLivetvTunerhostsTypes**](LiveTvServiceAPI.md#GetLivetvTunerhostsTypes) | **Get** /LiveTv/TunerHosts/Types | 
[**GetLivetvTunersDiscvover**](LiveTvServiceAPI.md#GetLivetvTunersDiscvover) | **Get** /LiveTv/Tuners/Discvover | 
[**HeadLivetvChannelmappingoptions**](LiveTvServiceAPI.md#HeadLivetvChannelmappingoptions) | **Head** /LiveTv/ChannelMappingOptions | 
[**HeadLivetvChannelmappings**](LiveTvServiceAPI.md#HeadLivetvChannelmappings) | **Head** /LiveTv/ChannelMappings | 
[**PostLivetvChannelmappingoptions**](LiveTvServiceAPI.md#PostLivetvChannelmappingoptions) | **Post** /LiveTv/ChannelMappingOptions | 
[**PostLivetvChannelmappings**](LiveTvServiceAPI.md#PostLivetvChannelmappings) | **Post** /LiveTv/ChannelMappings | 
[**PostLivetvListingproviders**](LiveTvServiceAPI.md#PostLivetvListingproviders) | **Post** /LiveTv/ListingProviders | Adds a listing provider
[**PostLivetvListingprovidersDelete**](LiveTvServiceAPI.md#PostLivetvListingprovidersDelete) | **Post** /LiveTv/ListingProviders/Delete | Deletes a listing provider
[**PostLivetvManageChannelsByIdDisabled**](LiveTvServiceAPI.md#PostLivetvManageChannelsByIdDisabled) | **Post** /LiveTv/Manage/Channels/{Id}/Disabled | Sets a channel disabled or not
[**PostLivetvManageChannelsByIdSortindex**](LiveTvServiceAPI.md#PostLivetvManageChannelsByIdSortindex) | **Post** /LiveTv/Manage/Channels/{Id}/SortIndex | Sets a channel sort index
[**PostLivetvPrograms**](LiveTvServiceAPI.md#PostLivetvPrograms) | **Post** /LiveTv/Programs | Gets available live tv epgs..
[**PostLivetvRecordingsByIdDelete**](LiveTvServiceAPI.md#PostLivetvRecordingsByIdDelete) | **Post** /LiveTv/Recordings/{Id}/Delete | Deletes a live tv recording
[**PostLivetvSeriestimers**](LiveTvServiceAPI.md#PostLivetvSeriestimers) | **Post** /LiveTv/SeriesTimers | Creates a live tv series timer
[**PostLivetvSeriestimersById**](LiveTvServiceAPI.md#PostLivetvSeriestimersById) | **Post** /LiveTv/SeriesTimers/{Id} | Updates a live tv series timer
[**PostLivetvSeriestimersByIdDelete**](LiveTvServiceAPI.md#PostLivetvSeriestimersByIdDelete) | **Post** /LiveTv/SeriesTimers/{Id}/Delete | Cancels a live tv series timer
[**PostLivetvTimers**](LiveTvServiceAPI.md#PostLivetvTimers) | **Post** /LiveTv/Timers | Creates a live tv timer
[**PostLivetvTimersById**](LiveTvServiceAPI.md#PostLivetvTimersById) | **Post** /LiveTv/Timers/{Id} | Updates a live tv timer
[**PostLivetvTimersByIdDelete**](LiveTvServiceAPI.md#PostLivetvTimersByIdDelete) | **Post** /LiveTv/Timers/{Id}/Delete | Cancels a live tv timer
[**PostLivetvTunerhosts**](LiveTvServiceAPI.md#PostLivetvTunerhosts) | **Post** /LiveTv/TunerHosts | Adds a tuner host
[**PostLivetvTunerhostsDelete**](LiveTvServiceAPI.md#PostLivetvTunerhostsDelete) | **Post** /LiveTv/TunerHosts/Delete | Deletes a tuner host
[**PostLivetvTunersByIdReset**](LiveTvServiceAPI.md#PostLivetvTunersByIdReset) | **Post** /LiveTv/Tuners/{Id}/Reset | Resets a tv tuner
[**PutLivetvChannelmappingoptions**](LiveTvServiceAPI.md#PutLivetvChannelmappingoptions) | **Put** /LiveTv/ChannelMappingOptions | 
[**PutLivetvChannelmappings**](LiveTvServiceAPI.md#PutLivetvChannelmappings) | **Put** /LiveTv/ChannelMappings | 



## DeleteLivetvChannelmappingoptions

> DeleteLivetvChannelmappingoptions(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvChannelmappingoptions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvChannelmappingoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvChannelmappingoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## DeleteLivetvChannelmappings

> DeleteLivetvChannelmappings(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvChannelmappings(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvChannelmappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvChannelmappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## DeleteLivetvListingproviders

> DeleteLivetvListingproviders(ctx).Id(id).Execute()

Deletes a listing provider



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
	id := "id_example" // string | Provider id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvListingproviders(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvListingproviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvListingprovidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Provider id | 

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


## DeleteLivetvRecordingsById

> DeleteLivetvRecordingsById(ctx, id).Execute()

Deletes a live tv recording



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
	id := "id_example" // string | Recording Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvRecordingsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvRecordingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Recording Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvRecordingsByIdRequest struct via the builder pattern


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


## DeleteLivetvSeriestimersById

> DeleteLivetvSeriestimersById(ctx, id).Execute()

Cancels a live tv series timer



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
	id := "id_example" // string | Timer Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvSeriestimersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvSeriestimersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Timer Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvSeriestimersByIdRequest struct via the builder pattern


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


## DeleteLivetvTimersById

> DeleteLivetvTimersById(ctx, id).Execute()

Cancels a live tv timer



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
	id := "id_example" // string | Timer Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvTimersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvTimersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Timer Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvTimersByIdRequest struct via the builder pattern


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


## DeleteLivetvTunerhosts

> DeleteLivetvTunerhosts(ctx).Id(id).Execute()

Deletes a tuner host



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
	id := "id_example" // string | Tuner host id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.DeleteLivetvTunerhosts(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.DeleteLivetvTunerhosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLivetvTunerhostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Tuner host id | 

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


## GetLivetvAvailablerecordingoptions

> ModelModelLiveTVApiAvailableRecordingOptions GetLivetvAvailablerecordingoptions(ctx).Execute()

Gets available recording options



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvAvailablerecordingoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvAvailablerecordingoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvAvailablerecordingoptions`: ModelModelLiveTVApiAvailableRecordingOptions
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvAvailablerecordingoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvAvailablerecordingoptionsRequest struct via the builder pattern


### Return type

[**ModelModelLiveTVApiAvailableRecordingOptions**](ModelLiveTVApiAvailableRecordingOptions.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvChannelmappingoptions

> GetLivetvChannelmappingoptions(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvChannelmappingoptions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvChannelmappingoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvChannelmappingoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## GetLivetvChannelmappings

> GetLivetvChannelmappings(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvChannelmappings(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvChannelmappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvChannelmappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## GetLivetvChannels

> ModelModelQueryResultBaseItemDto GetLivetvChannels(ctx).Type_(type_).IsLiked(isLiked).IsDisliked(isDisliked).EnableFavoriteSorting(enableFavoriteSorting).AddCurrentProgram(addCurrentProgram).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets available live tv channels.



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
	type_ := openapiclient.LiveTv.ChannelType("TV") // ModelLiveTvChannelType | Optional filter by channel type. (optional)
	isLiked := true // bool | Filter by channels that are liked, or not. (optional)
	isDisliked := true // bool | Filter by channels that are disliked, or not. (optional)
	enableFavoriteSorting := true // bool | Incorporate favorite and like status into channel sorting. (optional)
	addCurrentProgram := true // bool | Optional. Adds current program info to each channel (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvChannels(context.Background()).Type_(type_).IsLiked(isLiked).IsDisliked(isDisliked).EnableFavoriteSorting(enableFavoriteSorting).AddCurrentProgram(addCurrentProgram).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvChannels`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ModelLiveTvChannelType**](ModelLiveTvChannelType.md) | Optional filter by channel type. | 
 **isLiked** | **bool** | Filter by channels that are liked, or not. | 
 **isDisliked** | **bool** | Filter by channels that are disliked, or not. | 
 **enableFavoriteSorting** | **bool** | Incorporate favorite and like status into channel sorting. | 
 **addCurrentProgram** | **bool** | Optional. Adds current program info to each channel | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**ModelModelQueryResultBaseItemDto**](ModelQueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvChannelsById

> ModelModelBaseItemDto GetLivetvChannelsById(ctx, id).UserId(userId).Execute()

Gets a live tv channel



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
	id := "id_example" // string | Channel Id
	userId := "userId_example" // string | Optional attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvChannelsById(context.Background(), id).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvChannelsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvChannelsById`: ModelModelBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvChannelsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Channel Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvChannelsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional attach user data. | 

### Return type

[**ModelModelBaseItemDto**](ModelBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvChanneltags

> ModelModelQueryResultBaseItemDto GetLivetvChanneltags(ctx).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets live tv channel tags



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
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvChanneltags(context.Background()).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvChanneltags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvChanneltags`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvChanneltags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvChanneltagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**ModelModelQueryResultBaseItemDto**](ModelQueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvChanneltagsPrefixes

> []ModelModelLiveTVApiTagItem GetLivetvChanneltagsPrefixes(ctx).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets live tv channel tag prefixes



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
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvChanneltagsPrefixes(context.Background()).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvChanneltagsPrefixes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvChanneltagsPrefixes`: []ModelModelLiveTVApiTagItem
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvChanneltagsPrefixes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvChanneltagsPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**[]ModelModelLiveTVApiTagItem**](ModelLiveTVApiTagItem.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvEPG

> ModelModelQueryResultLiveTVApiEpgRow GetLivetvEPG(ctx).Type_(type_).IsLiked(isLiked).IsDisliked(isDisliked).EnableFavoriteSorting(enableFavoriteSorting).AddCurrentProgram(addCurrentProgram).ChannelIds(channelIds).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets the epg.



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
	type_ := openapiclient.LiveTv.ChannelType("TV") // ModelLiveTvChannelType | Optional filter by channel type. (optional)
	isLiked := true // bool | Filter by channels that are liked, or not. (optional)
	isDisliked := true // bool | Filter by channels that are disliked, or not. (optional)
	enableFavoriteSorting := true // bool | Incorporate favorite and like status into channel sorting. (optional)
	addCurrentProgram := true // bool | Optional. Adds current program info to each channel (optional)
	channelIds := "channelIds_example" // string | The channels to return guide information for. (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvEPG(context.Background()).Type_(type_).IsLiked(isLiked).IsDisliked(isDisliked).EnableFavoriteSorting(enableFavoriteSorting).AddCurrentProgram(addCurrentProgram).ChannelIds(channelIds).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvEPG``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvEPG`: ModelModelQueryResultLiveTVApiEpgRow
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvEPG`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvEPGRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**ModelLiveTvChannelType**](ModelLiveTvChannelType.md) | Optional filter by channel type. | 
 **isLiked** | **bool** | Filter by channels that are liked, or not. | 
 **isDisliked** | **bool** | Filter by channels that are disliked, or not. | 
 **enableFavoriteSorting** | **bool** | Incorporate favorite and like status into channel sorting. | 
 **addCurrentProgram** | **bool** | Optional. Adds current program info to each channel | 
 **channelIds** | **string** | The channels to return guide information for. | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

[**ModelModelQueryResultLiveTVApiEpgRow**](ModelQueryResultLiveTVApiEpgRow.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvFolder

> ModelModelBaseItemDto GetLivetvFolder(ctx).Execute()

Gets the top level live tv folder



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvFolder(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvFolder`: ModelModelBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvFolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvFolderRequest struct via the builder pattern


### Return type

[**ModelModelBaseItemDto**](ModelBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvGuideinfo

> ModelModelLiveTvGuideInfo GetLivetvGuideinfo(ctx).Execute()

Gets guide info



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvGuideinfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvGuideinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvGuideinfo`: ModelModelLiveTvGuideInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvGuideinfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvGuideinfoRequest struct via the builder pattern


### Return type

[**ModelModelLiveTvGuideInfo**](ModelLiveTvGuideInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvInfo

> ModelModelLiveTvLiveTvInfo GetLivetvInfo(ctx).Execute()

Gets available live tv services.



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvInfo`: ModelModelLiveTvLiveTvInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvInfoRequest struct via the builder pattern


### Return type

[**ModelModelLiveTvLiveTvInfo**](ModelLiveTvLiveTvInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvListingproviders

> []ModelModelLiveTvListingsProviderInfo GetLivetvListingproviders(ctx).ChannelId(channelId).Execute()

Gets current listing providers



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
	channelId := "channelId_example" // string | Channel id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvListingproviders(context.Background()).ChannelId(channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvListingproviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvListingproviders`: []ModelModelLiveTvListingsProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvListingproviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvListingprovidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Channel id | 

### Return type

[**[]ModelModelLiveTvListingsProviderInfo**](ModelLiveTvListingsProviderInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvListingprovidersAvailable

> []ModelModelLiveTVApiListingProviderTypeInfo GetLivetvListingprovidersAvailable(ctx).Execute()

Gets listing provider



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersAvailable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvListingprovidersAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvListingprovidersAvailable`: []ModelModelLiveTVApiListingProviderTypeInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvListingprovidersAvailable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvListingprovidersAvailableRequest struct via the builder pattern


### Return type

[**[]ModelModelLiveTVApiListingProviderTypeInfo**](ModelLiveTVApiListingProviderTypeInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvListingprovidersDefault

> ModelModelLiveTvListingsProviderInfo GetLivetvListingprovidersDefault(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersDefault(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvListingprovidersDefault``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvListingprovidersDefault`: ModelModelLiveTvListingsProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvListingprovidersDefault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvListingprovidersDefaultRequest struct via the builder pattern


### Return type

[**ModelModelLiveTvListingsProviderInfo**](ModelLiveTvListingsProviderInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvListingprovidersLineups

> []ModelModelNameIdPair GetLivetvListingprovidersLineups(ctx).Id(id).Type_(type_).Location(location).Country(country).Execute()

Gets available lineups



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
	id := "id_example" // string | Provider id (optional)
	type_ := "type__example" // string | Provider Type (optional)
	location := "location_example" // string | Location (optional)
	country := "country_example" // string | Country (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersLineups(context.Background()).Id(id).Type_(type_).Location(location).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvListingprovidersLineups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvListingprovidersLineups`: []ModelModelNameIdPair
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvListingprovidersLineups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvListingprovidersLineupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Provider id | 
 **type_** | **string** | Provider Type | 
 **location** | **string** | Location | 
 **country** | **string** | Country | 

### Return type

[**[]ModelModelNameIdPair**](ModelNameIdPair.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvListingprovidersSchedulesdirectCountries

> GetLivetvListingprovidersSchedulesdirectCountries(ctx).Execute()

Gets available lineups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvListingprovidersSchedulesdirectCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvListingprovidersSchedulesdirectCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvListingprovidersSchedulesdirectCountriesRequest struct via the builder pattern


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


## GetLivetvLiverecordingsByIdStream

> GetLivetvLiverecordingsByIdStream(ctx, id).Execute()

Gets a live tv channel



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvLiverecordingsByIdStream(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvLiverecordingsByIdStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLiverecordingsByIdStreamRequest struct via the builder pattern


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


## GetLivetvLivestreamfilesByIdStreamByContainer

> GetLivetvLivestreamfilesByIdStreamByContainer(ctx, id, container).Execute()

Gets a live tv channel



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
	container := "container_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvLivestreamfilesByIdStreamByContainer(context.Background(), id, container).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvLivestreamfilesByIdStreamByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**container** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvLivestreamfilesByIdStreamByContainerRequest struct via the builder pattern


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


## GetLivetvManageChannels

> ModelModelQueryResultChannelManagementInfo GetLivetvManageChannels(ctx).StartIndex(startIndex).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()

Gets the channel management list



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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Name, StartDate (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvManageChannels(context.Background()).StartIndex(startIndex).Limit(limit).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvManageChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvManageChannels`: ModelModelQueryResultChannelManagementInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvManageChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvManageChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Name, StartDate | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

### Return type

[**ModelModelQueryResultChannelManagementInfo**](ModelQueryResultChannelManagementInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvPrograms

> GetLivetvPrograms(ctx).ChannelIds(channelIds).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets available live tv epgs..



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
	channelIds := "channelIds_example" // string | The channels to return guide information for. (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvPrograms(context.Background()).ChannelIds(channelIds).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelIds** | **string** | The channels to return guide information for. | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

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


## GetLivetvProgramsRecommended

> ModelModelQueryResultBaseItemDto GetLivetvProgramsRecommended(ctx).UserId(userId).Limit(limit).IsAiring(isAiring).HasAired(hasAired).IsSeries(isSeries).IsMovie(isMovie).IsNews(isNews).IsKids(isKids).IsSports(isSports).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).GenreIds(genreIds).Fields(fields).EnableUserData(enableUserData).Execute()

Gets available live tv epgs..



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
	userId := "userId_example" // string | Optional filter by user id. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	isAiring := true // bool | Optional. Filter by programs that are currently airing, or not. (optional)
	hasAired := true // bool | Optional. Filter by programs that have completed airing, or not. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	genreIds := "genreIds_example" // string | The genres to return guide information for. (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	enableUserData := true // bool | Optional, include user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvProgramsRecommended(context.Background()).UserId(userId).Limit(limit).IsAiring(isAiring).HasAired(hasAired).IsSeries(isSeries).IsMovie(isMovie).IsNews(isNews).IsKids(isKids).IsSports(isSports).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).GenreIds(genreIds).Fields(fields).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvProgramsRecommended``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvProgramsRecommended`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvProgramsRecommended`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvProgramsRecommendedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional filter by user id. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **isAiring** | **bool** | Optional. Filter by programs that are currently airing, or not. | 
 **hasAired** | **bool** | Optional. Filter by programs that have completed airing, or not. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **genreIds** | **string** | The genres to return guide information for. | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **enableUserData** | **bool** | Optional, include user data | 

### Return type

[**ModelModelQueryResultBaseItemDto**](ModelQueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvRecordings

> GetLivetvRecordings(ctx).ChannelId(channelId).Status(status).IsInProgress(isInProgress).SeriesTimerId(seriesTimerId).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets live tv recordings



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
	channelId := "channelId_example" // string | Optional filter by channel id. (optional)
	status := openapiclient.LiveTv.RecordingStatus("New") // ModelLiveTvRecordingStatus | Optional filter by recording status. (optional)
	isInProgress := true // bool | Optional filter by recordings that are in progress, or not. (optional)
	seriesTimerId := "seriesTimerId_example" // string | Optional filter by recordings belonging to a series timer (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.GetLivetvRecordings(context.Background()).ChannelId(channelId).Status(status).IsInProgress(isInProgress).SeriesTimerId(seriesTimerId).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvRecordings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvRecordingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Optional filter by channel id. | 
 **status** | [**ModelLiveTvRecordingStatus**](ModelLiveTvRecordingStatus.md) | Optional filter by recording status. | 
 **isInProgress** | **bool** | Optional filter by recordings that are in progress, or not. | 
 **seriesTimerId** | **string** | Optional filter by recordings belonging to a series timer | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

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


## GetLivetvRecordingsById

> ModelModelBaseItemDto GetLivetvRecordingsById(ctx, id).UserId(userId).Execute()

Gets a live tv recording



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
	id := "id_example" // string | Recording Id
	userId := "userId_example" // string | Optional attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsById(context.Background(), id).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvRecordingsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvRecordingsById`: ModelModelBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvRecordingsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Recording Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvRecordingsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional attach user data. | 

### Return type

[**ModelModelBaseItemDto**](ModelBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvRecordingsFolders

> []ModelModelBaseItemDto GetLivetvRecordingsFolders(ctx).UserId(userId).Fields(fields).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()

Gets recording folders



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
	userId := "userId_example" // string | Optional filter by user and attach user data. (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsFolders(context.Background()).UserId(userId).Fields(fields).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvRecordingsFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvRecordingsFolders`: []ModelModelBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvRecordingsFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvRecordingsFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional filter by user and attach user data. | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional, include user data | 

### Return type

[**[]ModelModelBaseItemDto**](ModelBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvRecordingsGroups

> ModelModelQueryResultBaseItemDto GetLivetvRecordingsGroups(ctx).Execute()

Gets live tv recording groups



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvRecordingsGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvRecordingsGroups`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvRecordingsGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvRecordingsGroupsRequest struct via the builder pattern


### Return type

[**ModelModelQueryResultBaseItemDto**](ModelQueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvRecordingsSeries

> ModelModelQueryResultBaseItemDto GetLivetvRecordingsSeries(ctx).Execute()

Gets live tv recordings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvRecordingsSeries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvRecordingsSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvRecordingsSeries`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvRecordingsSeries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvRecordingsSeriesRequest struct via the builder pattern


### Return type

[**ModelModelQueryResultBaseItemDto**](ModelQueryResultBaseItemDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvSeriestimers

> ModelModelQueryResultLiveTvSeriesTimerInfoDto GetLivetvSeriestimers(ctx).SortBy(sortBy).SortOrder(sortOrder).StartIndex(startIndex).Limit(limit).Execute()

Gets live tv series timers



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
	sortBy := "sortBy_example" // string | Optional. Sort by SortName or Priority (optional)
	sortOrder := openapiclient.SortOrder("Ascending") // ModelSortOrder | Optional. Sort in Ascending or Descending order (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvSeriestimers(context.Background()).SortBy(sortBy).SortOrder(sortOrder).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvSeriestimers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvSeriestimers`: ModelModelQueryResultLiveTvSeriesTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvSeriestimers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvSeriestimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortBy** | **string** | Optional. Sort by SortName or Priority | 
 **sortOrder** | [**ModelSortOrder**](ModelSortOrder.md) | Optional. Sort in Ascending or Descending order | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

### Return type

[**ModelModelQueryResultLiveTvSeriesTimerInfoDto**](ModelQueryResultLiveTvSeriesTimerInfoDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvSeriestimersById

> ModelModelLiveTvTimerInfoDto GetLivetvSeriestimersById(ctx, id).Execute()

Gets a live tv series timer



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
	id := "id_example" // string | Timer Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvSeriestimersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvSeriestimersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvSeriestimersById`: ModelModelLiveTvTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvSeriestimersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Timer Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvSeriestimersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelModelLiveTvTimerInfoDto**](ModelLiveTvTimerInfoDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTimers

> ModelModelQueryResultLiveTvTimerInfoDto GetLivetvTimers(ctx).ChannelId(channelId).SeriesTimerId(seriesTimerId).Execute()

Gets live tv timers



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
	channelId := "channelId_example" // string | Optional filter by channel id. (optional)
	seriesTimerId := "seriesTimerId_example" // string | Optional filter by timers belonging to a series timer (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTimers(context.Background()).ChannelId(channelId).SeriesTimerId(seriesTimerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTimers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTimers`: ModelModelQueryResultLiveTvTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTimers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Optional filter by channel id. | 
 **seriesTimerId** | **string** | Optional filter by timers belonging to a series timer | 

### Return type

[**ModelModelQueryResultLiveTvTimerInfoDto**](ModelQueryResultLiveTvTimerInfoDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTimersById

> ModelModelLiveTvTimerInfoDto GetLivetvTimersById(ctx, id).Execute()

Gets a live tv timer



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
	id := "id_example" // string | Timer Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTimersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTimersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTimersById`: ModelModelLiveTvTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTimersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Timer Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTimersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelModelLiveTvTimerInfoDto**](ModelLiveTvTimerInfoDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTimersDefaults

> ModelModelLiveTvSeriesTimerInfoDto GetLivetvTimersDefaults(ctx).ProgramId(programId).Execute()

Gets default values for a new timer



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
	programId := "programId_example" // string | Optional, to attach default values based on a program. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTimersDefaults(context.Background()).ProgramId(programId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTimersDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTimersDefaults`: ModelModelLiveTvSeriesTimerInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTimersDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTimersDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **programId** | **string** | Optional, to attach default values based on a program. | 

### Return type

[**ModelModelLiveTvSeriesTimerInfoDto**](ModelLiveTvSeriesTimerInfoDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTunerhosts

> []ModelModelLiveTvTunerHostInfo GetLivetvTunerhosts(ctx).Execute()

Gets tuner hosts



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTunerhosts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTunerhosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTunerhosts`: []ModelModelLiveTvTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTunerhosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTunerhostsRequest struct via the builder pattern


### Return type

[**[]ModelModelLiveTvTunerHostInfo**](ModelLiveTvTunerHostInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTunerhostsDefaultByType

> ModelModelLiveTvTunerHostInfo GetLivetvTunerhostsDefaultByType(ctx, type_).Execute()

Gets tuner hosts



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
	type_ := "type__example" // string | Type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTunerhostsDefaultByType(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTunerhostsDefaultByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTunerhostsDefaultByType`: ModelModelLiveTvTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTunerhostsDefaultByType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTunerhostsDefaultByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelModelLiveTvTunerHostInfo**](ModelLiveTvTunerHostInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTunerhostsTypes

> []ModelModelNameIdPair GetLivetvTunerhostsTypes(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTunerhostsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTunerhostsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTunerhostsTypes`: []ModelModelNameIdPair
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTunerhostsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTunerhostsTypesRequest struct via the builder pattern


### Return type

[**[]ModelModelNameIdPair**](ModelNameIdPair.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLivetvTunersDiscvover

> []ModelModelLiveTvTunerHostInfo GetLivetvTunersDiscvover(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.GetLivetvTunersDiscvover(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.GetLivetvTunersDiscvover``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLivetvTunersDiscvover`: []ModelModelLiveTvTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.GetLivetvTunersDiscvover`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLivetvTunersDiscvoverRequest struct via the builder pattern


### Return type

[**[]ModelModelLiveTvTunerHostInfo**](ModelLiveTvTunerHostInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadLivetvChannelmappingoptions

> HeadLivetvChannelmappingoptions(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.HeadLivetvChannelmappingoptions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.HeadLivetvChannelmappingoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvChannelmappingoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## HeadLivetvChannelmappings

> HeadLivetvChannelmappings(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.HeadLivetvChannelmappings(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.HeadLivetvChannelmappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeadLivetvChannelmappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## PostLivetvChannelmappingoptions

> PostLivetvChannelmappingoptions(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvChannelmappingoptions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvChannelmappingoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvChannelmappingoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## PostLivetvChannelmappings

> PostLivetvChannelmappings(ctx).ProviderId(providerId).ModelLiveTVApiSetChannelMapping(modelLiveTVApiSetChannelMapping).Execute()





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
	providerId := "providerId_example" // string | Provider id
	modelLiveTVApiSetChannelMapping := *openapiclient.NewModelLiveTVApiSetChannelMapping() // ModelLiveTVApiSetChannelMapping | SetChannelMapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvChannelmappings(context.Background()).ProviderId(providerId).ModelLiveTVApiSetChannelMapping(modelLiveTVApiSetChannelMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvChannelmappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvChannelmappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 
 **modelLiveTVApiSetChannelMapping** | [**ModelLiveTVApiSetChannelMapping**](ModelLiveTVApiSetChannelMapping.md) | SetChannelMapping | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvListingproviders

> ModelModelLiveTvListingsProviderInfo PostLivetvListingproviders(ctx).ModelLiveTvListingsProviderInfo(modelLiveTvListingsProviderInfo).Execute()

Adds a listing provider



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
	modelLiveTvListingsProviderInfo := *openapiclient.NewModelLiveTvListingsProviderInfo() // ModelLiveTvListingsProviderInfo | ListingsProviderInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.PostLivetvListingproviders(context.Background()).ModelLiveTvListingsProviderInfo(modelLiveTvListingsProviderInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvListingproviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLivetvListingproviders`: ModelModelLiveTvListingsProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.PostLivetvListingproviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvListingprovidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLiveTvListingsProviderInfo** | [**ModelLiveTvListingsProviderInfo**](ModelLiveTvListingsProviderInfo.md) | ListingsProviderInfo:  | 

### Return type

[**ModelModelLiveTvListingsProviderInfo**](ModelLiveTvListingsProviderInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvListingprovidersDelete

> PostLivetvListingprovidersDelete(ctx).Id(id).Execute()

Deletes a listing provider



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
	id := "id_example" // string | Provider id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvListingprovidersDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvListingprovidersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvListingprovidersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Provider id | 

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


## PostLivetvManageChannelsByIdDisabled

> ModelModelQueryResultChannelManagementInfo PostLivetvManageChannelsByIdDisabled(ctx, id).ModelLiveTVApiSetChannelDisabled(modelLiveTVApiSetChannelDisabled).Execute()

Sets a channel disabled or not



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
	modelLiveTVApiSetChannelDisabled := *openapiclient.NewModelLiveTVApiSetChannelDisabled() // ModelLiveTVApiSetChannelDisabled | SetChannelDisabled

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.PostLivetvManageChannelsByIdDisabled(context.Background(), id).ModelLiveTVApiSetChannelDisabled(modelLiveTVApiSetChannelDisabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvManageChannelsByIdDisabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLivetvManageChannelsByIdDisabled`: ModelModelQueryResultChannelManagementInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.PostLivetvManageChannelsByIdDisabled`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvManageChannelsByIdDisabledRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelLiveTVApiSetChannelDisabled** | [**ModelLiveTVApiSetChannelDisabled**](ModelLiveTVApiSetChannelDisabled.md) | SetChannelDisabled | 

### Return type

[**ModelModelQueryResultChannelManagementInfo**](ModelQueryResultChannelManagementInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvManageChannelsByIdSortindex

> ModelModelQueryResultChannelManagementInfo PostLivetvManageChannelsByIdSortindex(ctx, id).ModelLiveTVApiSetChannelSortIndex(modelLiveTVApiSetChannelSortIndex).Execute()

Sets a channel sort index



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
	modelLiveTVApiSetChannelSortIndex := *openapiclient.NewModelLiveTVApiSetChannelSortIndex() // ModelLiveTVApiSetChannelSortIndex | SetChannelSortIndex

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.PostLivetvManageChannelsByIdSortindex(context.Background(), id).ModelLiveTVApiSetChannelSortIndex(modelLiveTVApiSetChannelSortIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvManageChannelsByIdSortindex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLivetvManageChannelsByIdSortindex`: ModelModelQueryResultChannelManagementInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.PostLivetvManageChannelsByIdSortindex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvManageChannelsByIdSortindexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelLiveTVApiSetChannelSortIndex** | [**ModelLiveTVApiSetChannelSortIndex**](ModelLiveTVApiSetChannelSortIndex.md) | SetChannelSortIndex | 

### Return type

[**ModelModelQueryResultChannelManagementInfo**](ModelQueryResultChannelManagementInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvPrograms

> PostLivetvPrograms(ctx).ModelApiBaseItemsRequest(modelApiBaseItemsRequest).ChannelIds(channelIds).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()

Gets available live tv epgs..



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
	modelApiBaseItemsRequest := *openapiclient.NewModelApiBaseItemsRequest() // ModelApiBaseItemsRequest | BaseItemsRequest: 
	channelIds := "channelIds_example" // string | The channels to return guide information for. (optional)
	artistType := "artistType_example" // string | Artist or AlbumArtist (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "adjacentTo_example" // string | Optional. Return items that are siblings of a supplied item. (optional)
	minIndexNumber := int32(56) // int32 | Optional filter by minimum index number. (optional)
	minStartDate := "minStartDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxStartDate := "maxStartDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minEndDate := "minEndDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxEndDate := "maxEndDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	minPlayers := int32(56) // int32 | Optional filter by minimum number of game players. (optional)
	maxPlayers := int32(56) // int32 | Optional filter by maximum number of game players. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating (optional)
	isHD := true // bool | Optional filter by items that are HD or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	airedDuringSeason := int32(56) // int32 | Gets all episodes that aired during a season, including specials. (optional)
	minPremiereDate := "minPremiereDate_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSaved := "minDateLastSaved_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	minDateLastSavedForUser := "minDateLastSavedForUser_example" // string | Optional. The minimum premiere date. Format = ISO (optional)
	maxPremiereDate := "maxPremiereDate_example" // string | Optional. The maximum premiere date. Format = ISO (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an imdb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a tmdb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a tvdb id or not. (optional)
	excludeItemIds := "excludeItemIds_example" // string | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false (optional)
	searchTerm := "searchTerm_example" // string | Enter a search term to perform a search request (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	excludeItemTypes := "excludeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	anyProviderIdEquals := "anyProviderIdEquals_example" // string | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form 'prov.id', e.g. 'imdb.tt123456'. This allows multiple, comma delimeted value pairs. (optional)
	filters := "filters_example" // string | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isFolder := true // bool | Optional filter for folders. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	isNew := true // bool | Optional filter for IsNew. (optional)
	isPremiere := true // bool | Optional filter for IsPremiere. (optional)
	isNewOrPremiere := true // bool | Optional filter for IsNewOrPremiere. (optional)
	isRepeat := true // bool | Optional filter for IsRepeat. (optional)
	projectToMedia := true // bool | ProjectToMedia (optional)
	mediaTypes := "mediaTypes_example" // string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := "imageTypes_example" // string | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := "sortBy_example" // string | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := "genres_example" // string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. (optional)
	officialRatings := "officialRatings_example" // string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. (optional)
	tags := "tags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	excludeTags := "excludeTags_example" // string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. (optional)
	years := "years_example" // string | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := "personIds_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personTypes := "personTypes_example" // string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited (optional)
	studios := "studios_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	studioIds := "studioIds_example" // string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. (optional)
	artists := "artists_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	artistIds := "artistIds_example" // string | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. (optional)
	albums := "albums_example" // string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. (optional)
	ids := "ids_example" // string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := "videoTypes_example" // string | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. (optional)
	containers := "containers_example" // string | Optional filter by Container. Allows multiple, comma delimeted. (optional)
	audioCodecs := "audioCodecs_example" // string | Optional filter by AudioCodec. Allows multiple, comma delimeted. (optional)
	audioLayouts := "audioLayouts_example" // string | Optional filter by AudioLayout. Allows multiple, comma delimeted. (optional)
	videoCodecs := "videoCodecs_example" // string | Optional filter by VideoCodec. Allows multiple, comma delimeted. (optional)
	extendedVideoTypes := "extendedVideoTypes_example" // string | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. (optional)
	subtitleCodecs := "subtitleCodecs_example" // string | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. (optional)
	path := "path_example" // string | Optional filter by Path. (optional)
	userId := "userId_example" // string | User Id (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings (optional)
	groupItemsIntoCollections := true // bool | Whether or not to hide items behind their boxsets. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := "seriesStatus_example" // string | Optional filter by Series Status. Allows multiple, comma delimeted. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	artistStartsWithOrGreater := "artistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	albumArtistStartsWithOrGreater := "albumArtistStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvPrograms(context.Background()).ModelApiBaseItemsRequest(modelApiBaseItemsRequest).ChannelIds(channelIds).ArtistType(artistType).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).MinIndexNumber(minIndexNumber).MinStartDate(minStartDate).MaxStartDate(maxStartDate).MinEndDate(minEndDate).MaxEndDate(maxEndDate).MinPlayers(minPlayers).MaxPlayers(maxPlayers).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHD(isHD).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).AiredDuringSeason(airedDuringSeason).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).AnyProviderIdEquals(anyProviderIdEquals).Filters(filters).IsFavorite(isFavorite).IsMovie(isMovie).IsSeries(isSeries).IsFolder(isFolder).IsNews(isNews).IsKids(isKids).IsSports(isSports).IsNew(isNew).IsPremiere(isPremiere).IsNewOrPremiere(isNewOrPremiere).IsRepeat(isRepeat).ProjectToMedia(projectToMedia).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).ExcludeTags(excludeTags).Years(years).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).Artists(artists).ArtistIds(artistIds).Albums(albums).Ids(ids).VideoTypes(videoTypes).Containers(containers).AudioCodecs(audioCodecs).AudioLayouts(audioLayouts).VideoCodecs(videoCodecs).ExtendedVideoTypes(extendedVideoTypes).SubtitleCodecs(subtitleCodecs).Path(path).UserId(userId).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).GroupItemsIntoCollections(groupItemsIntoCollections).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).ArtistStartsWithOrGreater(artistStartsWithOrGreater).AlbumArtistStartsWithOrGreater(albumArtistStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvProgramsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelApiBaseItemsRequest** | [**ModelApiBaseItemsRequest**](ModelApiBaseItemsRequest.md) | BaseItemsRequest:  | 
 **channelIds** | **string** | The channels to return guide information for. | 
 **artistType** | **string** | Artist or AlbumArtist | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **minIndexNumber** | **int32** | Optional filter by minimum index number. | 
 **minStartDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxStartDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minEndDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxEndDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **minPlayers** | **int32** | Optional filter by minimum number of game players. | 
 **maxPlayers** | **int32** | Optional filter by maximum number of game players. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating | 
 **isHD** | **bool** | Optional filter by items that are HD or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **airedDuringSeason** | **int32** | Gets all episodes that aired during a season, including specials. | 
 **minPremiereDate** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSaved** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **minDateLastSavedForUser** | **string** | Optional. The minimum premiere date. Format &#x3D; ISO | 
 **maxPremiereDate** | **string** | Optional. The maximum premiere date. Format &#x3D; ISO | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an imdb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a tmdb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a tvdb id or not. | 
 **excludeItemIds** | **string** | Optional. If specified, results will be filtered by exxcluding item ids. This allows multiple, comma delimeted. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false | 
 **searchTerm** | **string** | Enter a search term to perform a search request | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **excludeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **anyProviderIdEquals** | **string** | Optional. If specified, result will be filtered to contain only items which match at least one of the specified IDs. Each provider ID must be in the form &#39;prov.id&#39;, e.g. &#39;imdb.tt123456&#39;. This allows multiple, comma delimeted value pairs. | 
 **filters** | **string** | Optional. Specify additional filters to apply. This allows multiple, comma delimeted. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isFolder** | **bool** | Optional filter for folders. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **isNew** | **bool** | Optional filter for IsNew. | 
 **isPremiere** | **bool** | Optional filter for IsPremiere. | 
 **isNewOrPremiere** | **bool** | Optional filter for IsNewOrPremiere. | 
 **isRepeat** | **bool** | Optional filter for IsRepeat. | 
 **projectToMedia** | **bool** | ProjectToMedia | 
 **mediaTypes** | **string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | **string** | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | **string** | Optional. Specify one or more sort orders, comma delimeted. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimeted. | 
 **officialRatings** | **string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimeted. | 
 **tags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **excludeTags** | **string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimeted. | 
 **years** | **string** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personTypes** | **string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited | 
 **studios** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **studioIds** | **string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimeted. | 
 **artists** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **artistIds** | **string** | Optional. If specified, results will be filtered based on artist. This allows multiple, pipe delimeted. | 
 **albums** | **string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimeted. | 
 **ids** | **string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | **string** | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimeted. | 
 **containers** | **string** | Optional filter by Container. Allows multiple, comma delimeted. | 
 **audioCodecs** | **string** | Optional filter by AudioCodec. Allows multiple, comma delimeted. | 
 **audioLayouts** | **string** | Optional filter by AudioLayout. Allows multiple, comma delimeted. | 
 **videoCodecs** | **string** | Optional filter by VideoCodec. Allows multiple, comma delimeted. | 
 **extendedVideoTypes** | **string** | Optional filter by ExtendedVideoType. Allows multiple, comma delimeted. | 
 **subtitleCodecs** | **string** | Optional filter by SubtitleCodec. Allows multiple, comma delimeted. | 
 **path** | **string** | Optional filter by Path. | 
 **userId** | **string** | User Id | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings | 
 **groupItemsIntoCollections** | **bool** | Whether or not to hide items behind their boxsets. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | **string** | Optional filter by Series Status. Allows multiple, comma delimeted. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **artistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **albumArtistStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvRecordingsByIdDelete

> PostLivetvRecordingsByIdDelete(ctx, id).Execute()

Deletes a live tv recording



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
	id := "id_example" // string | Recording Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvRecordingsByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvRecordingsByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Recording Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvRecordingsByIdDeleteRequest struct via the builder pattern


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


## PostLivetvSeriestimers

> PostLivetvSeriestimers(ctx).ModelLiveTvSeriesTimerInfo(modelLiveTvSeriesTimerInfo).Execute()

Creates a live tv series timer



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
	modelLiveTvSeriesTimerInfo := *openapiclient.NewModelLiveTvSeriesTimerInfo() // ModelLiveTvSeriesTimerInfo | SeriesTimerInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvSeriestimers(context.Background()).ModelLiveTvSeriesTimerInfo(modelLiveTvSeriesTimerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvSeriestimers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvSeriestimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLiveTvSeriesTimerInfo** | [**ModelLiveTvSeriesTimerInfo**](ModelLiveTvSeriesTimerInfo.md) | SeriesTimerInfo:  | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvSeriestimersById

> PostLivetvSeriestimersById(ctx, id).ModelLiveTvSeriesTimerInfo(modelLiveTvSeriesTimerInfo).Execute()

Updates a live tv series timer



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
	modelLiveTvSeriesTimerInfo := *openapiclient.NewModelLiveTvSeriesTimerInfo() // ModelLiveTvSeriesTimerInfo | SeriesTimerInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvSeriestimersById(context.Background(), id).ModelLiveTvSeriesTimerInfo(modelLiveTvSeriesTimerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvSeriestimersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvSeriestimersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelLiveTvSeriesTimerInfo** | [**ModelLiveTvSeriesTimerInfo**](ModelLiveTvSeriesTimerInfo.md) | SeriesTimerInfo:  | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvSeriestimersByIdDelete

> PostLivetvSeriestimersByIdDelete(ctx, id).Execute()

Cancels a live tv series timer



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
	id := "id_example" // string | Timer Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvSeriestimersByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvSeriestimersByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Timer Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvSeriestimersByIdDeleteRequest struct via the builder pattern


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


## PostLivetvTimers

> PostLivetvTimers(ctx).ModelLiveTvTimerInfoDto(modelLiveTvTimerInfoDto).Execute()

Creates a live tv timer



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
	modelLiveTvTimerInfoDto := *openapiclient.NewModelLiveTvTimerInfoDto() // ModelLiveTvTimerInfoDto | TimerInfoDto: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvTimers(context.Background()).ModelLiveTvTimerInfoDto(modelLiveTvTimerInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvTimers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvTimersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLiveTvTimerInfoDto** | [**ModelLiveTvTimerInfoDto**](ModelLiveTvTimerInfoDto.md) | TimerInfoDto:  | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvTimersById

> PostLivetvTimersById(ctx, id).ModelLiveTvTimerInfoDto(modelLiveTvTimerInfoDto).Execute()

Updates a live tv timer



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
	modelLiveTvTimerInfoDto := *openapiclient.NewModelLiveTvTimerInfoDto() // ModelLiveTvTimerInfoDto | TimerInfoDto: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvTimersById(context.Background(), id).ModelLiveTvTimerInfoDto(modelLiveTvTimerInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvTimersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvTimersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelLiveTvTimerInfoDto** | [**ModelLiveTvTimerInfoDto**](ModelLiveTvTimerInfoDto.md) | TimerInfoDto:  | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvTimersByIdDelete

> PostLivetvTimersByIdDelete(ctx, id).Execute()

Cancels a live tv timer



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
	id := "id_example" // string | Timer Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvTimersByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvTimersByIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Timer Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvTimersByIdDeleteRequest struct via the builder pattern


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


## PostLivetvTunerhosts

> ModelModelLiveTvTunerHostInfo PostLivetvTunerhosts(ctx).ModelLiveTvTunerHostInfo(modelLiveTvTunerHostInfo).Execute()

Adds a tuner host



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
	modelLiveTvTunerHostInfo := *openapiclient.NewModelLiveTvTunerHostInfo() // ModelLiveTvTunerHostInfo | TunerHostInfo: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiveTvServiceAPI.PostLivetvTunerhosts(context.Background()).ModelLiveTvTunerHostInfo(modelLiveTvTunerHostInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvTunerhosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLivetvTunerhosts`: ModelModelLiveTvTunerHostInfo
	fmt.Fprintf(os.Stdout, "Response from `LiveTvServiceAPI.PostLivetvTunerhosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvTunerhostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLiveTvTunerHostInfo** | [**ModelLiveTvTunerHostInfo**](ModelLiveTvTunerHostInfo.md) | TunerHostInfo:  | 

### Return type

[**ModelModelLiveTvTunerHostInfo**](ModelLiveTvTunerHostInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLivetvTunerhostsDelete

> PostLivetvTunerhostsDelete(ctx).Id(id).Execute()

Deletes a tuner host



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
	id := "id_example" // string | Tuner host id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvTunerhostsDelete(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvTunerhostsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvTunerhostsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Tuner host id | 

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


## PostLivetvTunersByIdReset

> PostLivetvTunersByIdReset(ctx, id).Execute()

Resets a tv tuner



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
	id := "id_example" // string | Tuner Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PostLivetvTunersByIdReset(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PostLivetvTunersByIdReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Tuner Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLivetvTunersByIdResetRequest struct via the builder pattern


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


## PutLivetvChannelmappingoptions

> PutLivetvChannelmappingoptions(ctx).ProviderId(providerId).Execute()





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
	providerId := "providerId_example" // string | Provider id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PutLivetvChannelmappingoptions(context.Background()).ProviderId(providerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PutLivetvChannelmappingoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLivetvChannelmappingoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 

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


## PutLivetvChannelmappings

> PutLivetvChannelmappings(ctx).ProviderId(providerId).ModelLiveTVApiSetChannelMapping(modelLiveTVApiSetChannelMapping).Execute()





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
	providerId := "providerId_example" // string | Provider id
	modelLiveTVApiSetChannelMapping := *openapiclient.NewModelLiveTVApiSetChannelMapping() // ModelLiveTVApiSetChannelMapping | SetChannelMapping

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LiveTvServiceAPI.PutLivetvChannelmappings(context.Background()).ProviderId(providerId).ModelLiveTVApiSetChannelMapping(modelLiveTVApiSetChannelMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiveTvServiceAPI.PutLivetvChannelmappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutLivetvChannelmappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **providerId** | **string** | Provider id | 
 **modelLiveTVApiSetChannelMapping** | [**ModelLiveTVApiSetChannelMapping**](ModelLiveTVApiSetChannelMapping.md) | SetChannelMapping | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

