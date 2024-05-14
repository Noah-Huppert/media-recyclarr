# \ConnectServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersByIdConnectLink**](ConnectServiceAPI.md#DeleteUsersByIdConnectLink) | **Delete** /Users/{Id}/Connect/Link | Removes a Connect link for a user
[**GetConnectExchange**](ConnectServiceAPI.md#GetConnectExchange) | **Get** /Connect/Exchange | Gets the corresponding local user from a connect user id
[**GetConnectPending**](ConnectServiceAPI.md#GetConnectPending) | **Get** /Connect/Pending | Creates a Connect link for a user
[**PostUsersByIdConnectLink**](ConnectServiceAPI.md#PostUsersByIdConnectLink) | **Post** /Users/{Id}/Connect/Link | Creates a Connect link for a user
[**PostUsersByIdConnectLinkDelete**](ConnectServiceAPI.md#PostUsersByIdConnectLinkDelete) | **Post** /Users/{Id}/Connect/Link/Delete | Removes a Connect link for a user



## DeleteUsersByIdConnectLink

> DeleteUsersByIdConnectLink(ctx, id).Execute()

Removes a Connect link for a user



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
	id := "id_example" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectServiceAPI.DeleteUsersByIdConnectLink(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectServiceAPI.DeleteUsersByIdConnectLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByIdConnectLinkRequest struct via the builder pattern


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


## GetConnectExchange

> ModelModelConnectConnectAuthenticationExchangeResult GetConnectExchange(ctx).ConnectUserId(connectUserId).Execute()

Gets the corresponding local user from a connect user id



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
	connectUserId := "connectUserId_example" // string | ConnectUserId

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectServiceAPI.GetConnectExchange(context.Background()).ConnectUserId(connectUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectServiceAPI.GetConnectExchange``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectExchange`: ModelModelConnectConnectAuthenticationExchangeResult
	fmt.Fprintf(os.Stdout, "Response from `ConnectServiceAPI.GetConnectExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectUserId** | **string** | ConnectUserId | 

### Return type

[**ModelModelConnectConnectAuthenticationExchangeResult**](ModelConnectConnectAuthenticationExchangeResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectPending

> GetConnectPending(ctx).Execute()

Creates a Connect link for a user



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
	r, err := apiClient.ConnectServiceAPI.GetConnectPending(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectServiceAPI.GetConnectPending``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectPendingRequest struct via the builder pattern


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


## PostUsersByIdConnectLink

> ModelModelConnectUserLinkResult PostUsersByIdConnectLink(ctx, id).ConnectUsername(connectUsername).Execute()

Creates a Connect link for a user



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
	id := "id_example" // string | User Id
	connectUsername := "connectUsername_example" // string | Connect username

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectServiceAPI.PostUsersByIdConnectLink(context.Background(), id).ConnectUsername(connectUsername).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectServiceAPI.PostUsersByIdConnectLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByIdConnectLink`: ModelModelConnectUserLinkResult
	fmt.Fprintf(os.Stdout, "Response from `ConnectServiceAPI.PostUsersByIdConnectLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByIdConnectLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectUsername** | **string** | Connect username | 

### Return type

[**ModelModelConnectUserLinkResult**](ModelConnectUserLinkResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByIdConnectLinkDelete

> PostUsersByIdConnectLinkDelete(ctx, id).Execute()

Removes a Connect link for a user



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
	id := "id_example" // string | User Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConnectServiceAPI.PostUsersByIdConnectLinkDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectServiceAPI.PostUsersByIdConnectLinkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | User Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByIdConnectLinkDeleteRequest struct via the builder pattern


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

