# \ItemLookupServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemsByIdExternalidinfos**](ItemLookupServiceAPI.md#GetItemsByIdExternalidinfos) | **Get** /Items/{Id}/ExternalIdInfos | Gets external id infos for an item
[**GetItemsRemotesearchImage**](ItemLookupServiceAPI.md#GetItemsRemotesearchImage) | **Get** /Items/RemoteSearch/Image | Gets a remote image
[**PostItemsMetadataReset**](ItemLookupServiceAPI.md#PostItemsMetadataReset) | **Post** /Items/Metadata/Reset | Resets metadata for one or more items
[**PostItemsRemotesearchApplyById**](ItemLookupServiceAPI.md#PostItemsRemotesearchApplyById) | **Post** /Items/RemoteSearch/Apply/{Id} | Applies search criteria to an item and refreshes metadata
[**PostItemsRemotesearchBook**](ItemLookupServiceAPI.md#PostItemsRemotesearchBook) | **Post** /Items/RemoteSearch/Book | 
[**PostItemsRemotesearchBoxset**](ItemLookupServiceAPI.md#PostItemsRemotesearchBoxset) | **Post** /Items/RemoteSearch/BoxSet | 
[**PostItemsRemotesearchGame**](ItemLookupServiceAPI.md#PostItemsRemotesearchGame) | **Post** /Items/RemoteSearch/Game | 
[**PostItemsRemotesearchMovie**](ItemLookupServiceAPI.md#PostItemsRemotesearchMovie) | **Post** /Items/RemoteSearch/Movie | 
[**PostItemsRemotesearchMusicalbum**](ItemLookupServiceAPI.md#PostItemsRemotesearchMusicalbum) | **Post** /Items/RemoteSearch/MusicAlbum | 
[**PostItemsRemotesearchMusicartist**](ItemLookupServiceAPI.md#PostItemsRemotesearchMusicartist) | **Post** /Items/RemoteSearch/MusicArtist | 
[**PostItemsRemotesearchMusicvideo**](ItemLookupServiceAPI.md#PostItemsRemotesearchMusicvideo) | **Post** /Items/RemoteSearch/MusicVideo | 
[**PostItemsRemotesearchPerson**](ItemLookupServiceAPI.md#PostItemsRemotesearchPerson) | **Post** /Items/RemoteSearch/Person | 
[**PostItemsRemotesearchSeries**](ItemLookupServiceAPI.md#PostItemsRemotesearchSeries) | **Post** /Items/RemoteSearch/Series | 
[**PostItemsRemotesearchTrailer**](ItemLookupServiceAPI.md#PostItemsRemotesearchTrailer) | **Post** /Items/RemoteSearch/Trailer | 



## GetItemsByIdExternalidinfos

> []ModelModelExternalIdInfo GetItemsByIdExternalidinfos(ctx, id).Execute()

Gets external id infos for an item



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
	resp, r, err := apiClient.ItemLookupServiceAPI.GetItemsByIdExternalidinfos(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.GetItemsByIdExternalidinfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsByIdExternalidinfos`: []ModelModelExternalIdInfo
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.GetItemsByIdExternalidinfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsByIdExternalidinfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelModelExternalIdInfo**](ModelExternalIdInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemsRemotesearchImage

> GetItemsRemotesearchImage(ctx).ImageUrl(imageUrl).ProviderName(providerName).Execute()

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
	providerName := "providerName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupServiceAPI.GetItemsRemotesearchImage(context.Background()).ImageUrl(imageUrl).ProviderName(providerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.GetItemsRemotesearchImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRemotesearchImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imageUrl** | **string** | The image url | 
 **providerName** | **string** |  | 

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


## PostItemsMetadataReset

> PostItemsMetadataReset(ctx).ItemIds(itemIds).Execute()

Resets metadata for one or more items



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
	itemIds := "itemIds_example" // string | The item ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupServiceAPI.PostItemsMetadataReset(context.Background()).ItemIds(itemIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsMetadataReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsMetadataResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemIds** | **string** | The item ids | 

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


## PostItemsRemotesearchApplyById

> PostItemsRemotesearchApplyById(ctx, id).ModelRemoteSearchResult(modelRemoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()

Applies search criteria to an item and refreshes metadata



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
	id := "id_example" // string | The item id
	modelRemoteSearchResult := *openapiclient.NewModelRemoteSearchResult() // ModelRemoteSearchResult | RemoteSearchResult: 
	replaceAllImages := true // bool | Whether or not to replace all images (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchApplyById(context.Background(), id).ModelRemoteSearchResult(modelRemoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchApplyById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The item id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchApplyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelRemoteSearchResult** | [**ModelRemoteSearchResult**](ModelRemoteSearchResult.md) | RemoteSearchResult:  | 
 **replaceAllImages** | **bool** | Whether or not to replace all images | 

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


## PostItemsRemotesearchBook

> []ModelModelRemoteSearchResult PostItemsRemotesearchBook(ctx).ModelRemoteSearchQueryBookInfo(modelRemoteSearchQueryBookInfo).Execute()





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
	modelRemoteSearchQueryBookInfo := *openapiclient.NewModelRemoteSearchQueryBookInfo() // ModelRemoteSearchQueryBookInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchBook(context.Background()).ModelRemoteSearchQueryBookInfo(modelRemoteSearchQueryBookInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchBook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchBook`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryBookInfo** | [**ModelRemoteSearchQueryBookInfo**](ModelRemoteSearchQueryBookInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchBoxset

> []ModelModelRemoteSearchResult PostItemsRemotesearchBoxset(ctx).ModelRemoteSearchQueryItemLookupInfo(modelRemoteSearchQueryItemLookupInfo).Execute()





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
	modelRemoteSearchQueryItemLookupInfo := *openapiclient.NewModelRemoteSearchQueryItemLookupInfo() // ModelRemoteSearchQueryItemLookupInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchBoxset(context.Background()).ModelRemoteSearchQueryItemLookupInfo(modelRemoteSearchQueryItemLookupInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchBoxset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchBoxset`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchBoxset`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchBoxsetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryItemLookupInfo** | [**ModelRemoteSearchQueryItemLookupInfo**](ModelRemoteSearchQueryItemLookupInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchGame

> []ModelModelRemoteSearchResult PostItemsRemotesearchGame(ctx).ModelRemoteSearchQueryGameInfo(modelRemoteSearchQueryGameInfo).Execute()





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
	modelRemoteSearchQueryGameInfo := *openapiclient.NewModelRemoteSearchQueryGameInfo() // ModelRemoteSearchQueryGameInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchGame(context.Background()).ModelRemoteSearchQueryGameInfo(modelRemoteSearchQueryGameInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchGame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchGame`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchGame`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryGameInfo** | [**ModelRemoteSearchQueryGameInfo**](ModelRemoteSearchQueryGameInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMovie

> []ModelModelRemoteSearchResult PostItemsRemotesearchMovie(ctx).ModelRemoteSearchQueryMovieInfo(modelRemoteSearchQueryMovieInfo).Execute()





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
	modelRemoteSearchQueryMovieInfo := *openapiclient.NewModelRemoteSearchQueryMovieInfo() // ModelRemoteSearchQueryMovieInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMovie(context.Background()).ModelRemoteSearchQueryMovieInfo(modelRemoteSearchQueryMovieInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMovie``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMovie`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMovie`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMovieRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryMovieInfo** | [**ModelRemoteSearchQueryMovieInfo**](ModelRemoteSearchQueryMovieInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMusicalbum

> []ModelModelRemoteSearchResult PostItemsRemotesearchMusicalbum(ctx).ModelRemoteSearchQueryAlbumInfo(modelRemoteSearchQueryAlbumInfo).Execute()





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
	modelRemoteSearchQueryAlbumInfo := *openapiclient.NewModelRemoteSearchQueryAlbumInfo() // ModelRemoteSearchQueryAlbumInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMusicalbum(context.Background()).ModelRemoteSearchQueryAlbumInfo(modelRemoteSearchQueryAlbumInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMusicalbum``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMusicalbum`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMusicalbum`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMusicalbumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryAlbumInfo** | [**ModelRemoteSearchQueryAlbumInfo**](ModelRemoteSearchQueryAlbumInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMusicartist

> []ModelModelRemoteSearchResult PostItemsRemotesearchMusicartist(ctx).ModelRemoteSearchQueryArtistInfo(modelRemoteSearchQueryArtistInfo).Execute()





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
	modelRemoteSearchQueryArtistInfo := *openapiclient.NewModelRemoteSearchQueryArtistInfo() // ModelRemoteSearchQueryArtistInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMusicartist(context.Background()).ModelRemoteSearchQueryArtistInfo(modelRemoteSearchQueryArtistInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMusicartist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMusicartist`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMusicartist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMusicartistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryArtistInfo** | [**ModelRemoteSearchQueryArtistInfo**](ModelRemoteSearchQueryArtistInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchMusicvideo

> []ModelModelRemoteSearchResult PostItemsRemotesearchMusicvideo(ctx).ModelRemoteSearchQueryMusicVideoInfo(modelRemoteSearchQueryMusicVideoInfo).Execute()





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
	modelRemoteSearchQueryMusicVideoInfo := *openapiclient.NewModelRemoteSearchQueryMusicVideoInfo() // ModelRemoteSearchQueryMusicVideoInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchMusicvideo(context.Background()).ModelRemoteSearchQueryMusicVideoInfo(modelRemoteSearchQueryMusicVideoInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchMusicvideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchMusicvideo`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchMusicvideo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchMusicvideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryMusicVideoInfo** | [**ModelRemoteSearchQueryMusicVideoInfo**](ModelRemoteSearchQueryMusicVideoInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchPerson

> []ModelModelRemoteSearchResult PostItemsRemotesearchPerson(ctx).ModelRemoteSearchQueryPersonLookupInfo(modelRemoteSearchQueryPersonLookupInfo).Execute()





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
	modelRemoteSearchQueryPersonLookupInfo := *openapiclient.NewModelRemoteSearchQueryPersonLookupInfo() // ModelRemoteSearchQueryPersonLookupInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchPerson(context.Background()).ModelRemoteSearchQueryPersonLookupInfo(modelRemoteSearchQueryPersonLookupInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchPerson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchPerson`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchPerson`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryPersonLookupInfo** | [**ModelRemoteSearchQueryPersonLookupInfo**](ModelRemoteSearchQueryPersonLookupInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchSeries

> []ModelModelRemoteSearchResult PostItemsRemotesearchSeries(ctx).ModelRemoteSearchQuerySeriesInfo(modelRemoteSearchQuerySeriesInfo).Execute()





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
	modelRemoteSearchQuerySeriesInfo := *openapiclient.NewModelRemoteSearchQuerySeriesInfo() // ModelRemoteSearchQuerySeriesInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchSeries(context.Background()).ModelRemoteSearchQuerySeriesInfo(modelRemoteSearchQuerySeriesInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchSeries`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchSeries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQuerySeriesInfo** | [**ModelRemoteSearchQuerySeriesInfo**](ModelRemoteSearchQuerySeriesInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostItemsRemotesearchTrailer

> []ModelModelRemoteSearchResult PostItemsRemotesearchTrailer(ctx).ModelRemoteSearchQueryTrailerInfo(modelRemoteSearchQueryTrailerInfo).Execute()





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
	modelRemoteSearchQueryTrailerInfo := *openapiclient.NewModelRemoteSearchQueryTrailerInfo() // ModelRemoteSearchQueryTrailerInfo | RemoteSearchQuery`1: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupServiceAPI.PostItemsRemotesearchTrailer(context.Background()).ModelRemoteSearchQueryTrailerInfo(modelRemoteSearchQueryTrailerInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupServiceAPI.PostItemsRemotesearchTrailer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostItemsRemotesearchTrailer`: []ModelModelRemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupServiceAPI.PostItemsRemotesearchTrailer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostItemsRemotesearchTrailerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelRemoteSearchQueryTrailerInfo** | [**ModelRemoteSearchQueryTrailerInfo**](ModelRemoteSearchQueryTrailerInfo.md) | RemoteSearchQuery&#x60;1:  | 

### Return type

[**[]ModelModelRemoteSearchResult**](ModelRemoteSearchResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

