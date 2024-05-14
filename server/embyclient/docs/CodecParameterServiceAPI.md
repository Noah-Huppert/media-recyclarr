# \CodecParameterServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingCodecparameters**](CodecParameterServiceAPI.md#GetEncodingCodecparameters) | **Get** /Encoding/CodecParameters | Gets the parameters for a specified codec.
[**PostEncodingCodecparameters**](CodecParameterServiceAPI.md#PostEncodingCodecparameters) | **Post** /Encoding/CodecParameters | Updates the parameters for a specified codec.



## GetEncodingCodecparameters

> ModelModelEditObjectContainer GetEncodingCodecparameters(ctx).CodecId(codecId).ParameterContext(parameterContext).Execute()

Gets the parameters for a specified codec.



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
	codecId := "codecId_example" // string | Codec Id
	parameterContext := openapiclient.MediaEncoding.CodecParameterContext("Playback") // ModelMediaEncodingCodecParameterContext | Parameter Context

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CodecParameterServiceAPI.GetEncodingCodecparameters(context.Background()).CodecId(codecId).ParameterContext(parameterContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodecParameterServiceAPI.GetEncodingCodecparameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecparameters`: ModelModelEditObjectContainer
	fmt.Fprintf(os.Stdout, "Response from `CodecParameterServiceAPI.GetEncodingCodecparameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecparametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codecId** | **string** | Codec Id | 
 **parameterContext** | [**ModelMediaEncodingCodecParameterContext**](ModelMediaEncodingCodecParameterContext.md) | Parameter Context | 

### Return type

[**ModelModelEditObjectContainer**](ModelEditObjectContainer.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEncodingCodecparameters

> PostEncodingCodecparameters(ctx).CodecId(codecId).ParameterContext(parameterContext).Body(body).Execute()

Updates the parameters for a specified codec.



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
	codecId := "codecId_example" // string | Codec Id
	parameterContext := openapiclient.MediaEncoding.CodecParameterContext("Playback") // ModelMediaEncodingCodecParameterContext | Parameter Context
	body := os.NewFile(1234, "some_file") // *os.File | Binary stream

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CodecParameterServiceAPI.PostEncodingCodecparameters(context.Background()).CodecId(codecId).ParameterContext(parameterContext).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CodecParameterServiceAPI.PostEncodingCodecparameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostEncodingCodecparametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **codecId** | **string** | Codec Id | 
 **parameterContext** | [**ModelMediaEncodingCodecParameterContext**](ModelMediaEncodingCodecParameterContext.md) | Parameter Context | 
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

