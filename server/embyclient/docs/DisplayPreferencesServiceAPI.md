# \DisplayPreferencesServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisplaypreferencesById**](DisplayPreferencesServiceAPI.md#GetDisplaypreferencesById) | **Get** /DisplayPreferences/{Id} | Gets a user&#39;s display preferences for an item
[**PostDisplaypreferencesByDisplaypreferencesid**](DisplayPreferencesServiceAPI.md#PostDisplaypreferencesByDisplaypreferencesid) | **Post** /DisplayPreferences/{DisplayPreferencesId} | Updates a user&#39;s display preferences for an item



## GetDisplaypreferencesById

> ModelModelDisplayPreferences GetDisplaypreferencesById(ctx, id).UserId(userId).Client(client).Execute()

Gets a user's display preferences for an item



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
	client := "client_example" // string | Client

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayPreferencesServiceAPI.GetDisplaypreferencesById(context.Background(), id).UserId(userId).Client(client).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.GetDisplaypreferencesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisplaypreferencesById`: ModelModelDisplayPreferences
	fmt.Fprintf(os.Stdout, "Response from `DisplayPreferencesServiceAPI.GetDisplaypreferencesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisplaypreferencesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 
 **client** | **string** | Client | 

### Return type

[**ModelModelDisplayPreferences**](ModelDisplayPreferences.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDisplaypreferencesByDisplaypreferencesid

> PostDisplaypreferencesByDisplaypreferencesid(ctx, displayPreferencesId).UserId(userId).ModelDisplayPreferences(modelDisplayPreferences).Execute()

Updates a user's display preferences for an item



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
	displayPreferencesId := "displayPreferencesId_example" // string | DisplayPreferences Id
	userId := "userId_example" // string | User Id
	modelDisplayPreferences := *openapiclient.NewModelDisplayPreferences() // ModelDisplayPreferences | DisplayPreferences: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DisplayPreferencesServiceAPI.PostDisplaypreferencesByDisplaypreferencesid(context.Background(), displayPreferencesId).UserId(userId).ModelDisplayPreferences(modelDisplayPreferences).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesServiceAPI.PostDisplaypreferencesByDisplaypreferencesid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayPreferencesId** | **string** | DisplayPreferences Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDisplaypreferencesByDisplaypreferencesidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User Id | 
 **modelDisplayPreferences** | [**ModelDisplayPreferences**](ModelDisplayPreferences.md) | DisplayPreferences:  | 

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

