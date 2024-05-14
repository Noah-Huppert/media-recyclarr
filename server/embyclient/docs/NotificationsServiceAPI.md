# \NotificationsServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsTypes**](NotificationsServiceAPI.md#GetNotificationsTypes) | **Get** /Notifications/Types | Gets notification types
[**PostNotificationsAdmin**](NotificationsServiceAPI.md#PostNotificationsAdmin) | **Post** /Notifications/Admin | Sends a notification to all admin users



## GetNotificationsTypes

> []ModelModelNotificationCategoryInfo GetNotificationsTypes(ctx).Execute()

Gets notification types



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
	resp, r, err := apiClient.NotificationsServiceAPI.GetNotificationsTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.GetNotificationsTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsTypes`: []ModelModelNotificationCategoryInfo
	fmt.Fprintf(os.Stdout, "Response from `NotificationsServiceAPI.GetNotificationsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsTypesRequest struct via the builder pattern


### Return type

[**[]ModelModelNotificationCategoryInfo**](ModelNotificationCategoryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNotificationsAdmin

> PostNotificationsAdmin(ctx).Name(name).Description(description).ImageUrl(imageUrl).Url(url).Level(level).Execute()

Sends a notification to all admin users



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
	name := "name_example" // string | The notification's name
	description := "description_example" // string | The notification's description
	imageUrl := "imageUrl_example" // string | The notification's image url (optional)
	url := "url_example" // string | The notification's info url (optional)
	level := "level_example" // string | The notification level (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsServiceAPI.PostNotificationsAdmin(context.Background()).Name(name).Description(description).ImageUrl(imageUrl).Url(url).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsServiceAPI.PostNotificationsAdmin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostNotificationsAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The notification&#39;s name | 
 **description** | **string** | The notification&#39;s description | 
 **imageUrl** | **string** | The notification&#39;s image url | 
 **url** | **string** | The notification&#39;s info url | 
 **level** | **string** | The notification level | 

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

