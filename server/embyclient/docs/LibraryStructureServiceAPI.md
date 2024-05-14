# \LibraryStructureServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLibraryVirtualfolders**](LibraryStructureServiceAPI.md#DeleteLibraryVirtualfolders) | **Delete** /Library/VirtualFolders | 
[**DeleteLibraryVirtualfoldersPaths**](LibraryStructureServiceAPI.md#DeleteLibraryVirtualfoldersPaths) | **Delete** /Library/VirtualFolders/Paths | 
[**GetLibraryVirtualfoldersQuery**](LibraryStructureServiceAPI.md#GetLibraryVirtualfoldersQuery) | **Get** /Library/VirtualFolders/Query | 
[**PostLibraryVirtualfolders**](LibraryStructureServiceAPI.md#PostLibraryVirtualfolders) | **Post** /Library/VirtualFolders | 
[**PostLibraryVirtualfoldersDelete**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersDelete) | **Post** /Library/VirtualFolders/Delete | 
[**PostLibraryVirtualfoldersLibraryoptions**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersLibraryoptions) | **Post** /Library/VirtualFolders/LibraryOptions | 
[**PostLibraryVirtualfoldersName**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersName) | **Post** /Library/VirtualFolders/Name | 
[**PostLibraryVirtualfoldersPaths**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersPaths) | **Post** /Library/VirtualFolders/Paths | 
[**PostLibraryVirtualfoldersPathsDelete**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersPathsDelete) | **Post** /Library/VirtualFolders/Paths/Delete | 
[**PostLibraryVirtualfoldersPathsUpdate**](LibraryStructureServiceAPI.md#PostLibraryVirtualfoldersPathsUpdate) | **Post** /Library/VirtualFolders/Paths/Update | 



## DeleteLibraryVirtualfolders

> DeleteLibraryVirtualfolders(ctx).Execute()





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
	r, err := apiClient.LibraryStructureServiceAPI.DeleteLibraryVirtualfolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.DeleteLibraryVirtualfolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLibraryVirtualfoldersRequest struct via the builder pattern


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


## DeleteLibraryVirtualfoldersPaths

> DeleteLibraryVirtualfoldersPaths(ctx).Execute()





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
	r, err := apiClient.LibraryStructureServiceAPI.DeleteLibraryVirtualfoldersPaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.DeleteLibraryVirtualfoldersPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLibraryVirtualfoldersPathsRequest struct via the builder pattern


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


## GetLibraryVirtualfoldersQuery

> ModelModelQueryResultVirtualFolderInfo GetLibraryVirtualfoldersQuery(ctx).StartIndex(startIndex).Limit(limit).Execute()





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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery(context.Background()).StartIndex(startIndex).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryVirtualfoldersQuery`: ModelModelQueryResultVirtualFolderInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryStructureServiceAPI.GetLibraryVirtualfoldersQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryVirtualfoldersQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return | 

### Return type

[**ModelModelQueryResultVirtualFolderInfo**](ModelQueryResultVirtualFolderInfo.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLibraryVirtualfolders

> PostLibraryVirtualfolders(ctx).ModelLibraryAddVirtualFolder(modelLibraryAddVirtualFolder).Execute()





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
	modelLibraryAddVirtualFolder := *openapiclient.NewModelLibraryAddVirtualFolder() // ModelLibraryAddVirtualFolder | AddVirtualFolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfolders(context.Background()).ModelLibraryAddVirtualFolder(modelLibraryAddVirtualFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryAddVirtualFolder** | [**ModelLibraryAddVirtualFolder**](ModelLibraryAddVirtualFolder.md) | AddVirtualFolder | 

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


## PostLibraryVirtualfoldersDelete

> PostLibraryVirtualfoldersDelete(ctx).ModelLibraryRemoveVirtualFolder(modelLibraryRemoveVirtualFolder).Execute()





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
	modelLibraryRemoveVirtualFolder := *openapiclient.NewModelLibraryRemoveVirtualFolder() // ModelLibraryRemoveVirtualFolder | RemoveVirtualFolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete(context.Background()).ModelLibraryRemoveVirtualFolder(modelLibraryRemoveVirtualFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryRemoveVirtualFolder** | [**ModelLibraryRemoveVirtualFolder**](ModelLibraryRemoveVirtualFolder.md) | RemoveVirtualFolder | 

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


## PostLibraryVirtualfoldersLibraryoptions

> PostLibraryVirtualfoldersLibraryoptions(ctx).ModelLibraryUpdateLibraryOptions(modelLibraryUpdateLibraryOptions).Execute()





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
	modelLibraryUpdateLibraryOptions := *openapiclient.NewModelLibraryUpdateLibraryOptions() // ModelLibraryUpdateLibraryOptions | UpdateLibraryOptions

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersLibraryoptions(context.Background()).ModelLibraryUpdateLibraryOptions(modelLibraryUpdateLibraryOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersLibraryoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersLibraryoptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryUpdateLibraryOptions** | [**ModelLibraryUpdateLibraryOptions**](ModelLibraryUpdateLibraryOptions.md) | UpdateLibraryOptions | 

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


## PostLibraryVirtualfoldersName

> PostLibraryVirtualfoldersName(ctx).ModelLibraryRenameVirtualFolder(modelLibraryRenameVirtualFolder).Execute()





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
	modelLibraryRenameVirtualFolder := *openapiclient.NewModelLibraryRenameVirtualFolder() // ModelLibraryRenameVirtualFolder | RenameVirtualFolder

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersName(context.Background()).ModelLibraryRenameVirtualFolder(modelLibraryRenameVirtualFolder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryRenameVirtualFolder** | [**ModelLibraryRenameVirtualFolder**](ModelLibraryRenameVirtualFolder.md) | RenameVirtualFolder | 

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


## PostLibraryVirtualfoldersPaths

> PostLibraryVirtualfoldersPaths(ctx).ModelLibraryAddMediaPath(modelLibraryAddMediaPath).Execute()





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
	modelLibraryAddMediaPath := *openapiclient.NewModelLibraryAddMediaPath() // ModelLibraryAddMediaPath | AddMediaPath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths(context.Background()).ModelLibraryAddMediaPath(modelLibraryAddMediaPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryAddMediaPath** | [**ModelLibraryAddMediaPath**](ModelLibraryAddMediaPath.md) | AddMediaPath | 

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


## PostLibraryVirtualfoldersPathsDelete

> PostLibraryVirtualfoldersPathsDelete(ctx).ModelLibraryRemoveMediaPath(modelLibraryRemoveMediaPath).Execute()





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
	modelLibraryRemoveMediaPath := *openapiclient.NewModelLibraryRemoveMediaPath() // ModelLibraryRemoveMediaPath | RemoveMediaPath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete(context.Background()).ModelLibraryRemoveMediaPath(modelLibraryRemoveMediaPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersPathsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryRemoveMediaPath** | [**ModelLibraryRemoveMediaPath**](ModelLibraryRemoveMediaPath.md) | RemoveMediaPath | 

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


## PostLibraryVirtualfoldersPathsUpdate

> PostLibraryVirtualfoldersPathsUpdate(ctx).ModelLibraryUpdateMediaPath(modelLibraryUpdateMediaPath).Execute()





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
	modelLibraryUpdateMediaPath := *openapiclient.NewModelLibraryUpdateMediaPath() // ModelLibraryUpdateMediaPath | UpdateMediaPath

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsUpdate(context.Background()).ModelLibraryUpdateMediaPath(modelLibraryUpdateMediaPath).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureServiceAPI.PostLibraryVirtualfoldersPathsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLibraryVirtualfoldersPathsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelLibraryUpdateMediaPath** | [**ModelLibraryUpdateMediaPath**](ModelLibraryUpdateMediaPath.md) | UpdateMediaPath | 

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

