# \PlaylistServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlaylistsByIdItems**](PlaylistServiceAPI.md#DeletePlaylistsByIdItems) | **Delete** /Playlists/{Id}/Items | Removes items from a playlist
[**GetPlaylistsByIdAddtoplaylistinfo**](PlaylistServiceAPI.md#GetPlaylistsByIdAddtoplaylistinfo) | **Get** /Playlists/{Id}/AddToPlaylistInfo | Gets add to playlist info
[**GetPlaylistsByIdItems**](PlaylistServiceAPI.md#GetPlaylistsByIdItems) | **Get** /Playlists/{Id}/Items | Gets the original items of a playlist
[**PostPlaylists**](PlaylistServiceAPI.md#PostPlaylists) | **Post** /Playlists | Creates a new playlist
[**PostPlaylistsByIdItems**](PlaylistServiceAPI.md#PostPlaylistsByIdItems) | **Post** /Playlists/{Id}/Items | Adds items to a playlist
[**PostPlaylistsByIdItemsByItemidMoveByNewindex**](PlaylistServiceAPI.md#PostPlaylistsByIdItemsByItemidMoveByNewindex) | **Post** /Playlists/{Id}/Items/{ItemId}/Move/{NewIndex} | Moves a playlist item
[**PostPlaylistsByIdItemsDelete**](PlaylistServiceAPI.md#PostPlaylistsByIdItemsDelete) | **Post** /Playlists/{Id}/Items/Delete | Removes items from a playlist



## DeletePlaylistsByIdItems

> DeletePlaylistsByIdItems(ctx, id).EntryIds(entryIds).Execute()

Removes items from a playlist



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
	entryIds := "entryIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistServiceAPI.DeletePlaylistsByIdItems(context.Background(), id).EntryIds(entryIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.DeletePlaylistsByIdItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeletePlaylistsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryIds** | **string** |  | 

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


## GetPlaylistsByIdAddtoplaylistinfo

> interface{} GetPlaylistsByIdAddtoplaylistinfo(ctx, id).Ids(ids).UserId(userId).Execute()

Gets add to playlist info



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
	ids := "ids_example" // string | Item id, comma delimited
	id := "id_example" // string | 
	userId := "userId_example" // string | User Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistServiceAPI.GetPlaylistsByIdAddtoplaylistinfo(context.Background(), id).Ids(ids).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.GetPlaylistsByIdAddtoplaylistinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistsByIdAddtoplaylistinfo`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaylistServiceAPI.GetPlaylistsByIdAddtoplaylistinfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsByIdAddtoplaylistinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Item id, comma delimited | 

 **userId** | **string** | User Id | 

### Return type

**interface{}**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistsByIdItems

> ModelModelQueryResultBaseItemDto GetPlaylistsByIdItems(ctx, id).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()

Gets the original items of a playlist



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
	userId := "userId_example" // string | User Id (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistServiceAPI.GetPlaylistsByIdItems(context.Background(), id).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.GetPlaylistsByIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistsByIdItems`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `PlaylistServiceAPI.GetPlaylistsByIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 

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


## PostPlaylists

> ModelModelPlaylistsPlaylistCreationResult PostPlaylists(ctx).Name(name).Ids(ids).MediaType(mediaType).Execute()

Creates a new playlist



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
	name := "name_example" // string | The name of the new playlist. (optional)
	ids := "ids_example" // string | Item Ids to add to the playlist (optional)
	mediaType := "mediaType_example" // string | The playlist media type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistServiceAPI.PostPlaylists(context.Background()).Name(name).Ids(ids).MediaType(mediaType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.PostPlaylists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlaylists`: ModelModelPlaylistsPlaylistCreationResult
	fmt.Fprintf(os.Stdout, "Response from `PlaylistServiceAPI.PostPlaylists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPlaylistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the new playlist. | 
 **ids** | **string** | Item Ids to add to the playlist | 
 **mediaType** | **string** | The playlist media type | 

### Return type

[**ModelModelPlaylistsPlaylistCreationResult**](ModelPlaylistsPlaylistCreationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlaylistsByIdItems

> ModelModelPlaylistsAddToPlaylistResult PostPlaylistsByIdItems(ctx, id).Ids(ids).UserId(userId).Execute()

Adds items to a playlist



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
	ids := "ids_example" // string | Item id, comma delimited
	id := "id_example" // string | 
	userId := "userId_example" // string | User Id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistServiceAPI.PostPlaylistsByIdItems(context.Background(), id).Ids(ids).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.PostPlaylistsByIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlaylistsByIdItems`: ModelModelPlaylistsAddToPlaylistResult
	fmt.Fprintf(os.Stdout, "Response from `PlaylistServiceAPI.PostPlaylistsByIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlaylistsByIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Item id, comma delimited | 

 **userId** | **string** | User Id | 

### Return type

[**ModelModelPlaylistsAddToPlaylistResult**](ModelPlaylistsAddToPlaylistResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlaylistsByIdItemsByItemidMoveByNewindex

> PostPlaylistsByIdItemsByItemidMoveByNewindex(ctx, itemId, id, newIndex).Execute()

Moves a playlist item



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
	itemId := int64(789) // int64 | ItemId
	id := "id_example" // string | 
	newIndex := int32(56) // int32 | NewIndex

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistServiceAPI.PostPlaylistsByIdItemsByItemidMoveByNewindex(context.Background(), itemId, id, newIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.PostPlaylistsByIdItemsByItemidMoveByNewindex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **int64** | ItemId | 
**id** | **string** |  | 
**newIndex** | **int32** | NewIndex | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlaylistsByIdItemsByItemidMoveByNewindexRequest struct via the builder pattern


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


## PostPlaylistsByIdItemsDelete

> PostPlaylistsByIdItemsDelete(ctx, id).EntryIds(entryIds).Execute()

Removes items from a playlist



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
	entryIds := "entryIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistServiceAPI.PostPlaylistsByIdItemsDelete(context.Background(), id).EntryIds(entryIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistServiceAPI.PostPlaylistsByIdItemsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostPlaylistsByIdItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryIds** | **string** |  | 

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

