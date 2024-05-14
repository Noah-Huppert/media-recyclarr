# \EnvironmentServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironmentDefaultdirectorybrowser**](EnvironmentServiceAPI.md#GetEnvironmentDefaultdirectorybrowser) | **Get** /Environment/DefaultDirectoryBrowser | Gets the parent path of a given path
[**GetEnvironmentDirectorycontents**](EnvironmentServiceAPI.md#GetEnvironmentDirectorycontents) | **Get** /Environment/DirectoryContents | Gets the contents of a given directory in the file system
[**GetEnvironmentDrives**](EnvironmentServiceAPI.md#GetEnvironmentDrives) | **Get** /Environment/Drives | Gets available drives from the server&#39;s file system
[**GetEnvironmentNetworkdevices**](EnvironmentServiceAPI.md#GetEnvironmentNetworkdevices) | **Get** /Environment/NetworkDevices | Gets a list of devices on the network
[**GetEnvironmentNetworkshares**](EnvironmentServiceAPI.md#GetEnvironmentNetworkshares) | **Get** /Environment/NetworkShares | Gets shares from a network device
[**GetEnvironmentParentpath**](EnvironmentServiceAPI.md#GetEnvironmentParentpath) | **Get** /Environment/ParentPath | Gets the parent path of a given path
[**PostEnvironmentDirectorycontents**](EnvironmentServiceAPI.md#PostEnvironmentDirectorycontents) | **Post** /Environment/DirectoryContents | Gets the contents of a given directory in the file system
[**PostEnvironmentValidatepath**](EnvironmentServiceAPI.md#PostEnvironmentValidatepath) | **Post** /Environment/ValidatePath | Gets the contents of a given directory in the file system



## GetEnvironmentDefaultdirectorybrowser

> ModelModelDefaultDirectoryBrowserInfo GetEnvironmentDefaultdirectorybrowser(ctx).Execute()

Gets the parent path of a given path



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
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentDefaultdirectorybrowser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentDefaultdirectorybrowser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentDefaultdirectorybrowser`: ModelModelDefaultDirectoryBrowserInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentDefaultdirectorybrowser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDefaultdirectorybrowserRequest struct via the builder pattern


### Return type

[**ModelModelDefaultDirectoryBrowserInfo**](ModelDefaultDirectoryBrowserInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentDirectorycontents

> []ModelModelIOFileSystemEntryInfo GetEnvironmentDirectorycontents(ctx).Path(path).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()

Gets the contents of a given directory in the file system



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
	path := "path_example" // string | 
	includeFiles := true // bool | An optional filter to include or exclude files from the results. true/false (optional)
	includeDirectories := true // bool | An optional filter to include or exclude folders from the results. true/false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentDirectorycontents(context.Background()).Path(path).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentDirectorycontents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentDirectorycontents`: []ModelModelIOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentDirectorycontents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDirectorycontentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **includeFiles** | **bool** | An optional filter to include or exclude files from the results. true/false | 
 **includeDirectories** | **bool** | An optional filter to include or exclude folders from the results. true/false | 

### Return type

[**[]ModelModelIOFileSystemEntryInfo**](ModelIOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentDrives

> []ModelModelIOFileSystemEntryInfo GetEnvironmentDrives(ctx).Execute()

Gets available drives from the server's file system



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
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentDrives(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentDrives`: []ModelModelIOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentDrives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentDrivesRequest struct via the builder pattern


### Return type

[**[]ModelModelIOFileSystemEntryInfo**](ModelIOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentNetworkdevices

> []ModelModelIOFileSystemEntryInfo GetEnvironmentNetworkdevices(ctx).Execute()

Gets a list of devices on the network



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
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentNetworkdevices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentNetworkdevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentNetworkdevices`: []ModelModelIOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentNetworkdevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentNetworkdevicesRequest struct via the builder pattern


### Return type

[**[]ModelModelIOFileSystemEntryInfo**](ModelIOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentNetworkshares

> []ModelModelIOFileSystemEntryInfo GetEnvironmentNetworkshares(ctx).Path(path).Execute()

Gets shares from a network device



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentNetworkshares(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentNetworkshares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentNetworkshares`: []ModelModelIOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentNetworkshares`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentNetworksharesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

### Return type

[**[]ModelModelIOFileSystemEntryInfo**](ModelIOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentParentpath

> string GetEnvironmentParentpath(ctx).Path(path).Execute()

Gets the parent path of a given path



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
	path := "path_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.GetEnvironmentParentpath(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.GetEnvironmentParentpath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentParentpath`: string
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.GetEnvironmentParentpath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentParentpathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 

### Return type

**string**

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnvironmentDirectorycontents

> []ModelModelIOFileSystemEntryInfo PostEnvironmentDirectorycontents(ctx).Path(path).ModelGetDirectoryContents(modelGetDirectoryContents).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()

Gets the contents of a given directory in the file system



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
	path := "path_example" // string | 
	modelGetDirectoryContents := *openapiclient.NewModelGetDirectoryContents() // ModelGetDirectoryContents | GetDirectoryContents
	includeFiles := true // bool | An optional filter to include or exclude files from the results. true/false (optional)
	includeDirectories := true // bool | An optional filter to include or exclude folders from the results. true/false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentServiceAPI.PostEnvironmentDirectorycontents(context.Background()).Path(path).ModelGetDirectoryContents(modelGetDirectoryContents).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.PostEnvironmentDirectorycontents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEnvironmentDirectorycontents`: []ModelModelIOFileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentServiceAPI.PostEnvironmentDirectorycontents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEnvironmentDirectorycontentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **modelGetDirectoryContents** | [**ModelGetDirectoryContents**](ModelGetDirectoryContents.md) | GetDirectoryContents | 
 **includeFiles** | **bool** | An optional filter to include or exclude files from the results. true/false | 
 **includeDirectories** | **bool** | An optional filter to include or exclude folders from the results. true/false | 

### Return type

[**[]ModelModelIOFileSystemEntryInfo**](ModelIOFileSystemEntryInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnvironmentValidatepath

> PostEnvironmentValidatepath(ctx).Path(path).ModelValidatePath(modelValidatePath).Execute()

Gets the contents of a given directory in the file system



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
	path := "path_example" // string | 
	modelValidatePath := *openapiclient.NewModelValidatePath() // ModelValidatePath | ValidatePath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentServiceAPI.PostEnvironmentValidatepath(context.Background()).Path(path).ModelValidatePath(modelValidatePath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentServiceAPI.PostEnvironmentValidatepath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEnvironmentValidatepathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** |  | 
 **modelValidatePath** | [**ModelValidatePath**](ModelValidatePath.md) | ValidatePath | 

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

