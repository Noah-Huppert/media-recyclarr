# \EncodingInfoServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEncodingCodecconfigurationDefaults**](EncodingInfoServiceAPI.md#GetEncodingCodecconfigurationDefaults) | **Get** /Encoding/CodecConfiguration/Defaults | Gets default codec configurations
[**GetEncodingCodecinformationVideo**](EncodingInfoServiceAPI.md#GetEncodingCodecinformationVideo) | **Get** /Encoding/CodecInformation/Video | Gets details about available video encoders and decoders
[**GetEncodingTonemapoptions**](EncodingInfoServiceAPI.md#GetEncodingTonemapoptions) | **Get** /Encoding/ToneMapOptions | Gets available tone mapping options



## GetEncodingCodecconfigurationDefaults

> []ModelModelCodecConfiguration GetEncodingCodecconfigurationDefaults(ctx).Execute()

Gets default codec configurations



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecconfigurationDefaults`: []ModelModelCodecConfiguration
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingCodecconfigurationDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecconfigurationDefaultsRequest struct via the builder pattern


### Return type

[**[]ModelModelCodecConfiguration**](ModelCodecConfiguration.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncodingCodecinformationVideo

> []ModelModelVideoCodecBase GetEncodingCodecinformationVideo(ctx).Execute()

Gets details about available video encoders and decoders



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingCodecinformationVideo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingCodecinformationVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingCodecinformationVideo`: []ModelModelVideoCodecBase
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingCodecinformationVideo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingCodecinformationVideoRequest struct via the builder pattern


### Return type

[**[]ModelModelVideoCodecBase**](ModelVideoCodecBase.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEncodingTonemapoptions

> ModelModelConfigurationToneMappingToneMapOptionsVisibility GetEncodingTonemapoptions(ctx).Execute()

Gets available tone mapping options



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
	resp, r, err := apiClient.EncodingInfoServiceAPI.GetEncodingTonemapoptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EncodingInfoServiceAPI.GetEncodingTonemapoptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEncodingTonemapoptions`: ModelModelConfigurationToneMappingToneMapOptionsVisibility
	fmt.Fprintf(os.Stdout, "Response from `EncodingInfoServiceAPI.GetEncodingTonemapoptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEncodingTonemapoptionsRequest struct via the builder pattern


### Return type

[**ModelModelConfigurationToneMappingToneMapOptionsVisibility**](ModelConfigurationToneMappingToneMapOptionsVisibility.md)

### Authorization

[apikeyauth](../README.md#apikeyauth), [embyauth](../README.md#embyauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

