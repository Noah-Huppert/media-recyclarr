# \SuggestionsServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsersByUseridSuggestions**](SuggestionsServiceAPI.md#GetUsersByUseridSuggestions) | **Get** /Users/{UserId}/Suggestions | Gets items based on a query.



## GetUsersByUseridSuggestions

> ModelModelQueryResultBaseItemDto GetUsersByUseridSuggestions(ctx, userId).Fields(fields).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()

Gets items based on a query.



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
	userId := "userId_example" // string | 
	fields := "fields_example" // string | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls (optional)
	enableImages := true // bool | Optional, include image information in output (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type (optional)
	enableImageTypes := "enableImageTypes_example" // string | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional, include user data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsServiceAPI.GetUsersByUseridSuggestions(context.Background(), userId).Fields(fields).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsServiceAPI.GetUsersByUseridSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersByUseridSuggestions`: ModelModelQueryResultBaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsServiceAPI.GetUsersByUseridSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimeted. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls | 
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

