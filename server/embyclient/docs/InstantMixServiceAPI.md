# \InstantMixServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlbumsByIdInstantmix**](InstantMixServiceAPI.md#GetAlbumsByIdInstantmix) | **Get** /Albums/{Id}/InstantMix | Creates an instant playlist based on a given album
[**GetArtistsInstantmix**](InstantMixServiceAPI.md#GetArtistsInstantmix) | **Get** /Artists/InstantMix | Creates an instant playlist based on a given artist
[**GetAudiobooksNextup**](InstantMixServiceAPI.md#GetAudiobooksNextup) | **Get** /AudioBooks/NextUp | Gets a list of next up episodes
[**GetItemsByIdInstantmix**](InstantMixServiceAPI.md#GetItemsByIdInstantmix) | **Get** /Items/{Id}/InstantMix | Creates an instant playlist based on a given item
[**GetMusicgenresByNameInstantmix**](InstantMixServiceAPI.md#GetMusicgenresByNameInstantmix) | **Get** /MusicGenres/{Name}/InstantMix | Creates an instant playlist based on a music genre
[**GetMusicgenresInstantmix**](InstantMixServiceAPI.md#GetMusicgenresInstantmix) | **Get** /MusicGenres/InstantMix | Creates an instant playlist based on a music genre
[**GetPlaylistsByIdInstantmix**](InstantMixServiceAPI.md#GetPlaylistsByIdInstantmix) | **Get** /Playlists/{Id}/InstantMix | Creates an instant playlist based on a given playlist
[**GetSongsByIdInstantmix**](InstantMixServiceAPI.md#GetSongsByIdInstantmix) | **Get** /Songs/{Id}/InstantMix | Creates an instant playlist based on a given song



## GetAlbumsByIdInstantmix

> ModelModelQueryResultBaseItemDto GetAlbumsByIdInstantmix(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a given album



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
	resp, r, err := apiClient.InstantMixServiceAPI.GetAlbumsByIdInstantmix(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetAlbumsByIdInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlbumsByIdInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetAlbumsByIdInstantmix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlbumsByIdInstantmixRequest struct via the builder pattern


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


## GetArtistsInstantmix

> ModelModelQueryResultBaseItemDto GetArtistsInstantmix(ctx).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a given artist



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
	resp, r, err := apiClient.InstantMixServiceAPI.GetArtistsInstantmix(context.Background()).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetArtistsInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtistsInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetArtistsInstantmix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistsInstantmixRequest struct via the builder pattern


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


## GetAudiobooksNextup

> ModelModelQueryResultBaseItemDto GetAudiobooksNextup(ctx).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).AlbumId(albumId).ParentId(parentId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()

Gets a list of next up episodes



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
	userId := "userId_example" // string | User Id
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)
	albumId := "albumId_example" // string | Optional. Filter by series id (optional)
	parentId := "parentId_example" // string | Specify this to localize the search to a specific item or folder. Omit to use the root (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstantMixServiceAPI.GetAudiobooksNextup(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).AlbumId(albumId).ParentId(parentId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetAudiobooksNextup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudiobooksNextup`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetAudiobooksNextup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiobooksNextupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 
 **albumId** | **string** | Optional. Filter by series id | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root | 
 **enableImages** | **bool** | Optional, include image information in output | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type | 
 **enableImageTypes** | **string** | Optional. The image types to include in the output. | 
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


## GetItemsByIdInstantmix

> ModelModelQueryResultBaseItemDto GetItemsByIdInstantmix(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a given item



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
	resp, r, err := apiClient.InstantMixServiceAPI.GetItemsByIdInstantmix(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetItemsByIdInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetItemsByIdInstantmix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdInstantmixRequest struct via the builder pattern


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


## GetMusicgenresByNameInstantmix

> ModelModelQueryResultBaseItemDto GetMusicgenresByNameInstantmix(ctx, name).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a music genre



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
	name := "name_example" // string | The genre name
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
	resp, r, err := apiClient.InstantMixServiceAPI.GetMusicgenresByNameInstantmix(context.Background(), name).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetMusicgenresByNameInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicgenresByNameInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetMusicgenresByNameInstantmix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The genre name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicgenresByNameInstantmixRequest struct via the builder pattern


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


## GetMusicgenresInstantmix

> ModelModelQueryResultBaseItemDto GetMusicgenresInstantmix(ctx).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a music genre



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
	resp, r, err := apiClient.InstantMixServiceAPI.GetMusicgenresInstantmix(context.Background()).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetMusicgenresInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicgenresInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetMusicgenresInstantmix`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicgenresInstantmixRequest struct via the builder pattern


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


## GetPlaylistsByIdInstantmix

> ModelModelQueryResultBaseItemDto GetPlaylistsByIdInstantmix(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a given playlist



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
	resp, r, err := apiClient.InstantMixServiceAPI.GetPlaylistsByIdInstantmix(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetPlaylistsByIdInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistsByIdInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetPlaylistsByIdInstantmix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistsByIdInstantmixRequest struct via the builder pattern


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


## GetSongsByIdInstantmix

> ModelModelQueryResultBaseItemDto GetSongsByIdInstantmix(ctx, id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()

Creates an instant playlist based on a given song



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
	resp, r, err := apiClient.InstantMixServiceAPI.GetSongsByIdInstantmix(context.Background(), id).IncludeItemTypes(includeItemTypes).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstantMixServiceAPI.GetSongsByIdInstantmix``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSongsByIdInstantmix`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `InstantMixServiceAPI.GetSongsByIdInstantmix`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSongsByIdInstantmixRequest struct via the builder pattern


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

