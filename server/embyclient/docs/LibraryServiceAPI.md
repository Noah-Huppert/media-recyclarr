# \LibraryServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItems**](LibraryServiceAPI.md#DeleteItems) | **Delete** /Items | Deletes an item from the library and file system
[**DeleteItemsById**](LibraryServiceAPI.md#DeleteItemsById) | **Delete** /Items/{Id} | Deletes an item from the library and file system
[**GetAlbumsByIdSimilar**](LibraryServiceAPI.md#GetAlbumsByIdSimilar) | **Get** /Albums/{Id}/Similar | Finds albums similar to a given album.
[**GetArtistsByIdSimilar**](LibraryServiceAPI.md#GetArtistsByIdSimilar) | **Get** /Artists/{Id}/Similar | Finds albums similar to a given album.
[**GetGamesByIdSimilar**](LibraryServiceAPI.md#GetGamesByIdSimilar) | **Get** /Games/{Id}/Similar | Finds games similar to a given game.
[**GetItemsByIdAncestors**](LibraryServiceAPI.md#GetItemsByIdAncestors) | **Get** /Items/{Id}/Ancestors | Gets all parents of an item
[**GetItemsByIdCriticreviews**](LibraryServiceAPI.md#GetItemsByIdCriticreviews) | **Get** /Items/{Id}/CriticReviews | Gets critic reviews for an item
[**GetItemsByIdDeleteinfo**](LibraryServiceAPI.md#GetItemsByIdDeleteinfo) | **Get** /Items/{Id}/DeleteInfo | Gets delete info for an item
[**GetItemsByIdDownload**](LibraryServiceAPI.md#GetItemsByIdDownload) | **Get** /Items/{Id}/Download | Downloads item media
[**GetItemsByIdFile**](LibraryServiceAPI.md#GetItemsByIdFile) | **Get** /Items/{Id}/File | Gets the original file of an item
[**GetItemsByIdSimilar**](LibraryServiceAPI.md#GetItemsByIdSimilar) | **Get** /Items/{Id}/Similar | Gets similar items
[**GetItemsByIdThememedia**](LibraryServiceAPI.md#GetItemsByIdThememedia) | **Get** /Items/{Id}/ThemeMedia | Gets theme videos and songs for an item
[**GetItemsByIdThemesongs**](LibraryServiceAPI.md#GetItemsByIdThemesongs) | **Get** /Items/{Id}/ThemeSongs | Gets theme songs for an item
[**GetItemsByIdThemevideos**](LibraryServiceAPI.md#GetItemsByIdThemevideos) | **Get** /Items/{Id}/ThemeVideos | Gets theme videos for an item
[**GetItemsCounts**](LibraryServiceAPI.md#GetItemsCounts) | **Get** /Items/Counts | 
[**GetItemsIntros**](LibraryServiceAPI.md#GetItemsIntros) | **Get** /Items/Intros | Gets info to debug intros
[**GetLibrariesAvailableoptions**](LibraryServiceAPI.md#GetLibrariesAvailableoptions) | **Get** /Libraries/AvailableOptions | 
[**GetLibraryMediafolders**](LibraryServiceAPI.md#GetLibraryMediafolders) | **Get** /Library/MediaFolders | Gets all user media folders.
[**GetLibraryPhysicalpaths**](LibraryServiceAPI.md#GetLibraryPhysicalpaths) | **Get** /Library/PhysicalPaths | Gets a list of physical paths from virtual folders
[**GetLibrarySelectablemediafolders**](LibraryServiceAPI.md#GetLibrarySelectablemediafolders) | **Get** /Library/SelectableMediaFolders | Gets all user media folders.
[**GetMoviesByIdSimilar**](LibraryServiceAPI.md#GetMoviesByIdSimilar) | **Get** /Movies/{Id}/Similar | Finds movies and trailers similar to a given movie.
[**GetShowsByIdSimilar**](LibraryServiceAPI.md#GetShowsByIdSimilar) | **Get** /Shows/{Id}/Similar | Finds tv shows similar to a given one.
[**GetTrailersByIdSimilar**](LibraryServiceAPI.md#GetTrailersByIdSimilar) | **Get** /Trailers/{Id}/Similar | Finds movies and trailers similar to a given trailer.
[**PostItemsByIdDelete**](LibraryServiceAPI.md#PostItemsByIdDelete) | **Post** /Items/{Id}/Delete | Deletes an item from the library and file system
[**PostItemsDelete**](LibraryServiceAPI.md#PostItemsDelete) | **Post** /Items/Delete | Deletes an item from the library and file system
[**PostLibraryMediaUpdated**](LibraryServiceAPI.md#PostLibraryMediaUpdated) | **Post** /Library/Media/Updated | Reports that new movies have been added by an external source
[**PostLibraryMoviesAdded**](LibraryServiceAPI.md#PostLibraryMoviesAdded) | **Post** /Library/Movies/Added | Deprecated. Use /Library/Media/Updated
[**PostLibraryMoviesUpdated**](LibraryServiceAPI.md#PostLibraryMoviesUpdated) | **Post** /Library/Movies/Updated | Deprecated. Use /Library/Media/Updated
[**PostLibraryRefresh**](LibraryServiceAPI.md#PostLibraryRefresh) | **Post** /Library/Refresh | Starts a library scan
[**PostLibrarySeriesAdded**](LibraryServiceAPI.md#PostLibrarySeriesAdded) | **Post** /Library/Series/Added | Deprecated. Use /Library/Media/Updated
[**PostLibrarySeriesUpdated**](LibraryServiceAPI.md#PostLibrarySeriesUpdated) | **Post** /Library/Series/Updated | Deprecated. Use /Library/Media/Updated



## DeleteItems

> DeleteItems(ctx).Ids(ids).Execute()

Deletes an item from the library and file system



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
	ids := "ids_example" // string | Ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.DeleteItems(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.DeleteItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Ids | 

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


## DeleteItemsById

> DeleteItemsById(ctx, id).Execute()

Deletes an item from the library and file system



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
	r, err := apiClient.LibraryServiceAPI.DeleteItemsById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.DeleteItemsById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteItemsByIdRequest struct via the builder pattern


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


## GetAlbumsByIdSimilar

> ModelModelQueryResultBaseItemDto GetAlbumsByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Finds albums similar to a given album.



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetAlbumsByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetAlbumsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlbumsByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetAlbumsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlbumsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## GetArtistsByIdSimilar

> ModelModelQueryResultBaseItemDto GetArtistsByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Finds albums similar to a given album.



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetArtistsByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetArtistsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtistsByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetArtistsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## GetGamesByIdSimilar

> ModelModelQueryResultBaseItemDto GetGamesByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Finds games similar to a given game.



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetGamesByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetGamesByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGamesByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetGamesByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGamesByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## GetItemsByIdAncestors

> []ModelModelBaseItemDto GetItemsByIdAncestors(ctx, id).UserId(userId).Execute()

Gets all parents of an item



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
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdAncestors(context.Background(), id).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdAncestors`: []ModelModelBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data | 

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


## GetItemsByIdCriticreviews

> ModelModelQueryResultBaseItemDto GetItemsByIdCriticreviews(ctx, id).StartIndex(startIndex).Limit(limit).Execute()

Gets critic reviews for an item



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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdCriticreviews(context.Background(), id).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdCriticreviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdCriticreviews`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdCriticreviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdCriticreviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

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


## GetItemsByIdDeleteinfo

> ModelModelLibraryDeleteInfo GetItemsByIdDeleteinfo(ctx, id).Execute()

Gets delete info for an item



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
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdDeleteinfo(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdDeleteinfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdDeleteinfo`: ModelModelLibraryDeleteInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdDeleteinfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdDeleteinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelModelLibraryDeleteInfo**](ModelLibraryDeleteInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdDownload

> GetItemsByIdDownload(ctx, id).Execute()

Downloads item media



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
	r, err := apiClient.LibraryServiceAPI.GetItemsByIdDownload(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdDownload``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetItemsByIdDownloadRequest struct via the builder pattern


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


## GetItemsByIdFile

> GetItemsByIdFile(ctx, id).Execute()

Gets the original file of an item



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
	r, err := apiClient.LibraryServiceAPI.GetItemsByIdFile(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdFile``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetItemsByIdFileRequest struct via the builder pattern


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


## GetItemsByIdSimilar

> ModelModelQueryResultBaseItemDto GetItemsByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## GetItemsByIdThememedia

> ModelModelAllThemeMediaResult GetItemsByIdThememedia(ctx, id).UserId(userId).InheritFromParent(inheritFromParent).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Fields(fields).Execute()

Gets theme videos and songs for an item



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
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	inheritFromParent := true // bool | Determines whether or not parent items should be searched for theme media. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdThememedia(context.Background(), id).UserId(userId).InheritFromParent(inheritFromParent).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdThememedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThememedia`: ModelModelAllThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdThememedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThememediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **inheritFromParent** | **bool** | Determines whether or not parent items should be searched for theme media. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional, include user data | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

### Return type

[**ModelModelAllThemeMediaResult**](ModelAllThemeMediaResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdThemesongs

> ModelModelThemeMediaResult GetItemsByIdThemesongs(ctx, id).UserId(userId).InheritFromParent(inheritFromParent).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Fields(fields).Execute()

Gets theme songs for an item



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
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	inheritFromParent := true // bool | Determines whether or not parent items should be searched for theme media. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdThemesongs(context.Background(), id).UserId(userId).InheritFromParent(inheritFromParent).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdThemesongs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThemesongs`: ModelModelThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdThemesongs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThemesongsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **inheritFromParent** | **bool** | Determines whether or not parent items should be searched for theme media. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional, include user data | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

### Return type

[**ModelModelThemeMediaResult**](ModelThemeMediaResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsByIdThemevideos

> ModelModelThemeMediaResult GetItemsByIdThemevideos(ctx, id).UserId(userId).InheritFromParent(inheritFromParent).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Fields(fields).Execute()

Gets theme videos for an item



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
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	inheritFromParent := true // bool | Determines whether or not parent items should be searched for theme media. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsByIdThemevideos(context.Background(), id).UserId(userId).InheritFromParent(inheritFromParent).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsByIdThemevideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdThemevideos`: ModelModelThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsByIdThemevideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdThemevideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **inheritFromParent** | **bool** | Determines whether or not parent items should be searched for theme media. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional, include user data | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

### Return type

[**ModelModelThemeMediaResult**](ModelThemeMediaResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsCounts

> ModelModelItemCounts GetItemsCounts(ctx).UserId(userId).IsFavorite(isFavorite).Execute()





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
	userId := "userId_example" // string | Optional. Get counts from a specific user's library. (optional)
	isFavorite := true // bool | Optional. Get counts of favorite items (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsCounts(context.Background()).UserId(userId).IsFavorite(isFavorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsCounts`: ModelModelItemCounts
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Get counts from a specific user&#39;s library. | 
 **isFavorite** | **bool** | Optional. Get counts of favorite items | 

### Return type

[**ModelModelItemCounts**](ModelItemCounts.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsIntros

> []ModelModelPersistenceIntroDebugInfo GetItemsIntros(ctx).Execute()

Gets info to debug intros



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
	resp, r, err := apiClient.LibraryServiceAPI.GetItemsIntros(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetItemsIntros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsIntros`: []ModelModelPersistenceIntroDebugInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetItemsIntros`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsIntrosRequest struct via the builder pattern


### Return type

[**[]ModelModelPersistenceIntroDebugInfo**](ModelPersistenceIntroDebugInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibrariesAvailableoptions

> ModelModelLibraryLibraryOptionsResult GetLibrariesAvailableoptions(ctx).Execute()





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
	resp, r, err := apiClient.LibraryServiceAPI.GetLibrariesAvailableoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibrariesAvailableoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibrariesAvailableoptions`: ModelModelLibraryLibraryOptionsResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibrariesAvailableoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibrariesAvailableoptionsRequest struct via the builder pattern


### Return type

[**ModelModelLibraryLibraryOptionsResult**](ModelLibraryLibraryOptionsResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryMediafolders

> ModelModelQueryResultBaseItemDto GetLibraryMediafolders(ctx).IsHidden(isHidden).Execute()

Gets all user media folders.



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
	isHidden := true // bool | Optional. Filter by folders that are marked hidden, or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetLibraryMediafolders(context.Background()).IsHidden(isHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibraryMediafolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryMediafolders`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibraryMediafolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryMediafoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional. Filter by folders that are marked hidden, or not. | 

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


## GetLibraryPhysicalpaths

> []string GetLibraryPhysicalpaths(ctx).Execute()

Gets a list of physical paths from virtual folders



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
	resp, r, err := apiClient.LibraryServiceAPI.GetLibraryPhysicalpaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibraryPhysicalpaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryPhysicalpaths`: []string
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibraryPhysicalpaths`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryPhysicalpathsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibrarySelectablemediafolders

> []ModelModelLibraryMediaFolder GetLibrarySelectablemediafolders(ctx).Execute()

Gets all user media folders.



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
	resp, r, err := apiClient.LibraryServiceAPI.GetLibrarySelectablemediafolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetLibrarySelectablemediafolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibrarySelectablemediafolders`: []ModelModelLibraryMediaFolder
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetLibrarySelectablemediafolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibrarySelectablemediafoldersRequest struct via the builder pattern


### Return type

[**[]ModelModelLibraryMediaFolder**](ModelLibraryMediaFolder.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoviesByIdSimilar

> ModelModelQueryResultBaseItemDto GetMoviesByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Finds movies and trailers similar to a given movie.



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetMoviesByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetMoviesByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoviesByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetMoviesByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoviesByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## GetShowsByIdSimilar

> ModelModelQueryResultBaseItemDto GetShowsByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Finds tv shows similar to a given one.



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetShowsByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetShowsByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShowsByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetShowsByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShowsByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## GetTrailersByIdSimilar

> ModelModelQueryResultBaseItemDto GetTrailersByIdSimilar(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Finds movies and trailers similar to a given trailer.



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
	includeItemTypes := "includeItemTypes_example" // string | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	enableUserData := true // bool | Optional, include user data (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	userId := "userId_example" // string | Optional. Filter by user id, and attach user data (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryServiceAPI.GetTrailersByIdSimilar(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.GetTrailersByIdSimilar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrailersByIdSimilar`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryServiceAPI.GetTrailersByIdSimilar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrailersByIdSimilarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeItemTypes** | **string** | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimeted. | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **enableUserData** | **bool** | Optional, include user data | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 

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


## PostItemsByIdDelete

> PostItemsByIdDelete(ctx, id).Execute()

Deletes an item from the library and file system



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
	r, err := apiClient.LibraryServiceAPI.PostItemsByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostItemsByIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostItemsByIdDeleteRequest struct via the builder pattern


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


## PostItemsDelete

> PostItemsDelete(ctx).Ids(ids).Execute()

Deletes an item from the library and file system



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
	ids := "ids_example" // string | Ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostItemsDelete(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostItemsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Ids | 

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


## PostLibraryMediaUpdated

> PostLibraryMediaUpdated(ctx).ModelLibraryPostUpdatedMedia(modelLibraryPostUpdatedMedia).Execute()

Reports that new movies have been added by an external source



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
	modelLibraryPostUpdatedMedia := *openapiclient.NewModelLibraryPostUpdatedMedia() // ModelLibraryPostUpdatedMedia | PostUpdatedMedia

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryServiceAPI.PostLibraryMediaUpdated(context.Background()).ModelLibraryPostUpdatedMedia(modelLibraryPostUpdatedMedia).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryMediaUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryMediaUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryPostUpdatedMedia** | [**ModelLibraryPostUpdatedMedia**](ModelLibraryPostUpdatedMedia.md) | PostUpdatedMedia | 

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


## PostLibraryMoviesAdded

> PostLibraryMoviesAdded(ctx).Execute()

Deprecated. Use /Library/Media/Updated



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
	r, err := apiClient.LibraryServiceAPI.PostLibraryMoviesAdded(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryMoviesAdded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryMoviesAddedRequest struct via the builder pattern


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


## PostLibraryMoviesUpdated

> PostLibraryMoviesUpdated(ctx).Execute()

Deprecated. Use /Library/Media/Updated



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
	r, err := apiClient.LibraryServiceAPI.PostLibraryMoviesUpdated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryMoviesUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryMoviesUpdatedRequest struct via the builder pattern


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


## PostLibraryRefresh

> PostLibraryRefresh(ctx).Execute()

Starts a library scan



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
	r, err := apiClient.LibraryServiceAPI.PostLibraryRefresh(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibraryRefresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryRefreshRequest struct via the builder pattern


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


## PostLibrarySeriesAdded

> PostLibrarySeriesAdded(ctx).Execute()

Deprecated. Use /Library/Media/Updated



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
	r, err := apiClient.LibraryServiceAPI.PostLibrarySeriesAdded(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibrarySeriesAdded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibrarySeriesAddedRequest struct via the builder pattern


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


## PostLibrarySeriesUpdated

> PostLibrarySeriesUpdated(ctx).Execute()

Deprecated. Use /Library/Media/Updated



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
	r, err := apiClient.LibraryServiceAPI.PostLibrarySeriesUpdated(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryServiceAPI.PostLibrarySeriesUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostLibrarySeriesUpdatedRequest struct via the builder pattern


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

