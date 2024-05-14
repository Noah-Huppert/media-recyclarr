# \UserServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsersById**](UserServiceAPI.md#DeleteUsersById) | **Delete** /Users/{Id} | Deletes a user
[**DeleteUsersByIdTrackselectionsByTracktype**](UserServiceAPI.md#DeleteUsersByIdTrackselectionsByTracktype) | **Delete** /Users/{Id}/TrackSelections/{TrackType} | Clears audio or subtitle track selections for a user
[**GetUsersById**](UserServiceAPI.md#GetUsersById) | **Get** /Users/{Id} | Gets a user by Id
[**GetUsersByUseridTypedsettingsByKey**](UserServiceAPI.md#GetUsersByUseridTypedsettingsByKey) | **Get** /Users/{UserId}/TypedSettings/{Key} | Gets a typed user setting
[**GetUsersPrefixes**](UserServiceAPI.md#GetUsersPrefixes) | **Get** /Users/Prefixes | Gets a list of users
[**GetUsersPublic**](UserServiceAPI.md#GetUsersPublic) | **Get** /Users/Public | Gets a list of publicly visible users for display on a login screen.
[**GetUsersQuery**](UserServiceAPI.md#GetUsersQuery) | **Get** /Users/Query | Gets a list of users
[**PostUsersAuthenticatebyname**](UserServiceAPI.md#PostUsersAuthenticatebyname) | **Post** /Users/AuthenticateByName | Authenticates a user
[**PostUsersById**](UserServiceAPI.md#PostUsersById) | **Post** /Users/{Id} | Updates a user
[**PostUsersByIdAuthenticate**](UserServiceAPI.md#PostUsersByIdAuthenticate) | **Post** /Users/{Id}/Authenticate | Authenticates a user
[**PostUsersByIdConfiguration**](UserServiceAPI.md#PostUsersByIdConfiguration) | **Post** /Users/{Id}/Configuration | Updates a user configuration
[**PostUsersByIdConfigurationPartial**](UserServiceAPI.md#PostUsersByIdConfigurationPartial) | **Post** /Users/{Id}/Configuration/Partial | Updates a user configuration
[**PostUsersByIdDelete**](UserServiceAPI.md#PostUsersByIdDelete) | **Post** /Users/{Id}/Delete | Deletes a user
[**PostUsersByIdPassword**](UserServiceAPI.md#PostUsersByIdPassword) | **Post** /Users/{Id}/Password | Updates a user&#39;s password
[**PostUsersByIdPolicy**](UserServiceAPI.md#PostUsersByIdPolicy) | **Post** /Users/{Id}/Policy | Updates a user policy
[**PostUsersByIdTrackselectionsByTracktypeDelete**](UserServiceAPI.md#PostUsersByIdTrackselectionsByTracktypeDelete) | **Post** /Users/{Id}/TrackSelections/{TrackType}/Delete | Clears audio or subtitle track selections for a user
[**PostUsersByUseridTypedsettingsByKey**](UserServiceAPI.md#PostUsersByUseridTypedsettingsByKey) | **Post** /Users/{UserId}/TypedSettings/{Key} | Updates a typed user setting
[**PostUsersForgotpassword**](UserServiceAPI.md#PostUsersForgotpassword) | **Post** /Users/ForgotPassword | Initiates the forgot password process for a local user
[**PostUsersForgotpasswordPin**](UserServiceAPI.md#PostUsersForgotpasswordPin) | **Post** /Users/ForgotPassword/Pin | Redeems a forgot password pin
[**PostUsersNew**](UserServiceAPI.md#PostUsersNew) | **Post** /Users/New | Creates a user



## DeleteUsersById

> DeleteUsersById(ctx, id).Execute()

Deletes a user



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
	r, err := apiClient.UserServiceAPI.DeleteUsersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.DeleteUsersById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteUsersByIdRequest struct via the builder pattern


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


## DeleteUsersByIdTrackselectionsByTracktype

> DeleteUsersByIdTrackselectionsByTracktype(ctx, id, trackType).Execute()

Clears audio or subtitle track selections for a user



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
	trackType := "trackType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.DeleteUsersByIdTrackselectionsByTracktype(context.Background(), id, trackType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.DeleteUsersByIdTrackselectionsByTracktype``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**trackType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersByIdTrackselectionsByTracktypeRequest struct via the builder pattern


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


## GetUsersById

> ModelModelUserDto GetUsersById(ctx, id).Execute()

Gets a user by Id



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
	resp, r, err := apiClient.UserServiceAPI.GetUsersById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersById`: ModelModelUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelModelUserDto**](ModelUserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersByUseridTypedsettingsByKey

> GetUsersByUseridTypedsettingsByKey(ctx, key, userId).Execute()

Gets a typed user setting



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
	key := "key_example" // string | Key
	userId := "userId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.GetUsersByUseridTypedsettingsByKey(context.Background(), key, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersByUseridTypedsettingsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersByUseridTypedsettingsByKeyRequest struct via the builder pattern


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


## GetUsersPrefixes

> []ModelModelNameIdPair GetUsersPrefixes(ctx).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()

Gets a list of users



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
	isHidden := true // bool | Optional filter by IsHidden=true or false (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.GetUsersPrefixes(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersPrefixes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersPrefixes`: []ModelModelNameIdPair
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersPrefixes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersPrefixesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

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


## GetUsersPublic

> []ModelModelUserDto GetUsersPublic(ctx).Execute()

Gets a list of publicly visible users for display on a login screen.



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
	resp, r, err := apiClient.UserServiceAPI.GetUsersPublic(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersPublic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersPublic`: []ModelModelUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersPublicRequest struct via the builder pattern


### Return type

[**[]ModelModelUserDto**](ModelUserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersQuery

> ModelModelQueryResultUserDto GetUsersQuery(ctx).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()

Gets a list of users



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
	isHidden := true // bool | Optional filter by IsHidden=true or false (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	sortOrder := "sortOrder_example" // string | Sort Order - Ascending,Descending (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.GetUsersQuery(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).StartIndex(startIndex).Limit(limit).NameStartsWithOrGreater(nameStartsWithOrGreater).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.GetUsersQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersQuery`: ModelModelQueryResultUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.GetUsersQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **sortOrder** | **string** | Sort Order - Ascending,Descending | 

### Return type

[**ModelModelQueryResultUserDto**](ModelQueryResultUserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersAuthenticatebyname

> ModelModelAuthenticationAuthenticationResult PostUsersAuthenticatebyname(ctx).XEmbyAuthorization(xEmbyAuthorization).ModelAuthenticateUserByName(modelAuthenticateUserByName).Execute()

Authenticates a user



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
	xEmbyAuthorization := "xEmbyAuthorization_example" // string | The authorization header can be either named 'Authorization' or 'X-Emby-Authorization'.    It must be of the following schema:     Emby UserId=\"(guid)\", Client=\"(string)\", Device=\"(string)\", DeviceId=\"(string)\", Version=\"string\", Token=\"(string)\"     Please consult the documentation for further details.
	modelAuthenticateUserByName := *openapiclient.NewModelAuthenticateUserByName() // ModelAuthenticateUserByName | AuthenticateUserByName

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersAuthenticatebyname(context.Background()).XEmbyAuthorization(xEmbyAuthorization).ModelAuthenticateUserByName(modelAuthenticateUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersAuthenticatebyname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersAuthenticatebyname`: ModelModelAuthenticationAuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersAuthenticatebyname`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersAuthenticatebynameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEmbyAuthorization** | **string** | The authorization header can be either named &#39;Authorization&#39; or &#39;X-Emby-Authorization&#39;.    It must be of the following schema:     Emby UserId&#x3D;\&quot;(guid)\&quot;, Client&#x3D;\&quot;(string)\&quot;, Device&#x3D;\&quot;(string)\&quot;, DeviceId&#x3D;\&quot;(string)\&quot;, Version&#x3D;\&quot;string\&quot;, Token&#x3D;\&quot;(string)\&quot;     Please consult the documentation for further details. | 
 **modelAuthenticateUserByName** | [**ModelAuthenticateUserByName**](ModelAuthenticateUserByName.md) | AuthenticateUserByName | 

### Return type

[**ModelModelAuthenticationAuthenticationResult**](ModelAuthenticationAuthenticationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersById

> PostUsersById(ctx, id).ModelUserDto(modelUserDto).Execute()

Updates a user



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
	modelUserDto := *openapiclient.NewModelUserDto() // ModelUserDto | UserDto: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersById(context.Background(), id).ModelUserDto(modelUserDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersById``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUserDto** | [**ModelUserDto**](ModelUserDto.md) | UserDto:  | 

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


## PostUsersByIdAuthenticate

> ModelModelAuthenticationAuthenticationResult PostUsersByIdAuthenticate(ctx, id).ModelAuthenticateUser(modelAuthenticateUser).Execute()

Authenticates a user



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
	modelAuthenticateUser := *openapiclient.NewModelAuthenticateUser() // ModelAuthenticateUser | AuthenticateUser

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersByIdAuthenticate(context.Background(), id).ModelAuthenticateUser(modelAuthenticateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersByIdAuthenticate`: ModelModelAuthenticationAuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersByIdAuthenticate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByIdAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelAuthenticateUser** | [**ModelAuthenticateUser**](ModelAuthenticateUser.md) | AuthenticateUser | 

### Return type

[**ModelModelAuthenticationAuthenticationResult**](ModelAuthenticationAuthenticationResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByIdConfiguration

> PostUsersByIdConfiguration(ctx, id).ModelUserConfiguration(modelUserConfiguration).Execute()

Updates a user configuration



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
	modelUserConfiguration := *openapiclient.NewModelUserConfiguration() // ModelUserConfiguration | UserConfiguration: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdConfiguration(context.Background(), id).ModelUserConfiguration(modelUserConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUserConfiguration** | [**ModelUserConfiguration**](ModelUserConfiguration.md) | UserConfiguration:  | 

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


## PostUsersByIdConfigurationPartial

> PostUsersByIdConfigurationPartial(ctx, id).Body(body).Execute()

Updates a user configuration



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
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdConfigurationPartial(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdConfigurationPartial``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdConfigurationPartialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersByIdDelete

> PostUsersByIdDelete(ctx, id).Execute()

Deletes a user



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
	r, err := apiClient.UserServiceAPI.PostUsersByIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdDeleteRequest struct via the builder pattern


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


## PostUsersByIdPassword

> PostUsersByIdPassword(ctx, id).ModelUpdateUserPassword(modelUpdateUserPassword).Execute()

Updates a user's password



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
	modelUpdateUserPassword := *openapiclient.NewModelUpdateUserPassword() // ModelUpdateUserPassword | UpdateUserPassword

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdPassword(context.Background(), id).ModelUpdateUserPassword(modelUpdateUserPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdPassword``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUpdateUserPassword** | [**ModelUpdateUserPassword**](ModelUpdateUserPassword.md) | UpdateUserPassword | 

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


## PostUsersByIdPolicy

> PostUsersByIdPolicy(ctx, id).ModelUserPolicy(modelUserPolicy).Execute()

Updates a user policy



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
	modelUserPolicy := *openapiclient.NewModelUserPolicy() // ModelUserPolicy | UserPolicy: 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdPolicy(context.Background(), id).ModelUserPolicy(modelUserPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPostUsersByIdPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelUserPolicy** | [**ModelUserPolicy**](ModelUserPolicy.md) | UserPolicy:  | 

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


## PostUsersByIdTrackselectionsByTracktypeDelete

> PostUsersByIdTrackselectionsByTracktypeDelete(ctx, id, trackType).Execute()

Clears audio or subtitle track selections for a user



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
	trackType := "trackType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByIdTrackselectionsByTracktypeDelete(context.Background(), id, trackType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByIdTrackselectionsByTracktypeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**trackType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByIdTrackselectionsByTracktypeDeleteRequest struct via the builder pattern


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


## PostUsersByUseridTypedsettingsByKey

> PostUsersByUseridTypedsettingsByKey(ctx, userId, key).Body(body).Execute()

Updates a typed user setting



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
	key := "key_example" // string | Key
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserServiceAPI.PostUsersByUseridTypedsettingsByKey(context.Background(), userId, key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersByUseridTypedsettingsByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**key** | **string** | Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersByUseridTypedsettingsByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | Binary stream | 

### Return type

 (empty response body)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersForgotpassword

> ModelModelForgotPasswordResult PostUsersForgotpassword(ctx).ModelForgotPassword(modelForgotPassword).Execute()

Initiates the forgot password process for a local user



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
	modelForgotPassword := *openapiclient.NewModelForgotPassword() // ModelForgotPassword | ForgotPassword

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersForgotpassword(context.Background()).ModelForgotPassword(modelForgotPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersForgotpassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersForgotpassword`: ModelModelForgotPasswordResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersForgotpassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersForgotpasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelForgotPassword** | [**ModelForgotPassword**](ModelForgotPassword.md) | ForgotPassword | 

### Return type

[**ModelModelForgotPasswordResult**](ModelForgotPasswordResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersForgotpasswordPin

> ModelModelPinRedeemResult PostUsersForgotpasswordPin(ctx).ModelForgotPasswordPin(modelForgotPasswordPin).Execute()

Redeems a forgot password pin



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
	modelForgotPasswordPin := *openapiclient.NewModelForgotPasswordPin() // ModelForgotPasswordPin | ForgotPasswordPin

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersForgotpasswordPin(context.Background()).ModelForgotPasswordPin(modelForgotPasswordPin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersForgotpasswordPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersForgotpasswordPin`: ModelModelPinRedeemResult
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersForgotpasswordPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersForgotpasswordPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelForgotPasswordPin** | [**ModelForgotPasswordPin**](ModelForgotPasswordPin.md) | ForgotPasswordPin | 

### Return type

[**ModelModelPinRedeemResult**](ModelPinRedeemResult.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsersNew

> ModelModelUserDto PostUsersNew(ctx).ModelCreateUserByName(modelCreateUserByName).Execute()

Creates a user



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
	modelCreateUserByName := *openapiclient.NewModelCreateUserByName() // ModelCreateUserByName | CreateUserByName

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserServiceAPI.PostUsersNew(context.Background()).ModelCreateUserByName(modelCreateUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserServiceAPI.PostUsersNew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUsersNew`: ModelModelUserDto
	fmt.Fprintf(os.Stdout, "Response from `UserServiceAPI.PostUsersNew`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsersNewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelCreateUserByName** | [**ModelCreateUserByName**](ModelCreateUserByName.md) | CreateUserByName | 

### Return type

[**ModelModelUserDto**](ModelUserDto.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

