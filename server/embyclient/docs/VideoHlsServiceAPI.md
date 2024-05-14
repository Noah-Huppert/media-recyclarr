# \VideoHlsServiceAPI

All URIs are relative to *http://view.media.k8s.funkyboy.zone/emby*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer**](VideoHlsServiceAPI.md#GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer) | **Get** /Audio/{Id}/hls/{PlaylistId}/{SegmentId}.{SegmentContainer} | 
[**GetAudioByIdLiveM3u8**](VideoHlsServiceAPI.md#GetAudioByIdLiveM3u8) | **Get** /Audio/{Id}/live.m3u8 | 
[**GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer**](VideoHlsServiceAPI.md#GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer) | **Get** /Videos/{Id}/hls/{PlaylistId}/{SegmentId}.{SegmentContainer} | 
[**GetVideosByIdLiveM3u8**](VideoHlsServiceAPI.md#GetVideosByIdLiveM3u8) | **Get** /Videos/{Id}/live.m3u8 | 



## GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer

> GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer(ctx, segmentContainer, segmentId, id, playlistId).Execute()





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
	segmentContainer := "segmentContainer_example" // string | 
	segmentId := "segmentId_example" // string | 
	id := "id_example" // string | 
	playlistId := "playlistId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoHlsServiceAPI.GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer(context.Background(), segmentContainer, segmentId, id, playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoHlsServiceAPI.GetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentContainer** | **string** |  | 
**segmentId** | **string** |  | 
**id** | **string** |  | 
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioByIdHlsByPlaylistidBySegmentidBySegmentcontainerRequest struct via the builder pattern


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


## GetAudioByIdLiveM3u8

> GetAudioByIdLiveM3u8(ctx, id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()





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
	container := "container_example" // string | Container
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100 (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2 (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false (optional)
	profile := "profile_example" // string | Optional. Specify a specific h264 profile, e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1ms = 10000 ticks. (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := openapiclient.SubtitleDeliveryMethod("Encode") // ModelSubtitleDeliveryMethod | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h264, mpeg4, theora, vpx, wmv. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoHlsServiceAPI.GetAudioByIdLiveM3u8(context.Background(), id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoHlsServiceAPI.GetAudioByIdLiveM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioByIdLiveM3u8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | Container | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100 | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2 | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false | 
 **profile** | **string** | Optional. Specify a specific h264 profile, e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1ms &#x3D; 10000 ticks. | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | [**ModelSubtitleDeliveryMethod**](ModelSubtitleDeliveryMethod.md) | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h264, mpeg4, theora, vpx, wmv. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 

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


## GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer

> GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer(ctx, segmentContainer, segmentId, id, playlistId).Execute()





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
	segmentContainer := "segmentContainer_example" // string | 
	segmentId := "segmentId_example" // string | 
	id := "id_example" // string | 
	playlistId := "playlistId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoHlsServiceAPI.GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer(context.Background(), segmentContainer, segmentId, id, playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoHlsServiceAPI.GetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**segmentContainer** | **string** |  | 
**segmentId** | **string** |  | 
**id** | **string** |  | 
**playlistId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdHlsByPlaylistidBySegmentidBySegmentcontainerRequest struct via the builder pattern


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


## GetVideosByIdLiveM3u8

> GetVideosByIdLiveM3u8(ctx, id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()





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
	container := "container_example" // string | Container
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100 (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2 (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false (optional)
	profile := "profile_example" // string | Optional. Specify a specific h264 profile, e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1ms = 10000 ticks. (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := openapiclient.SubtitleDeliveryMethod("Encode") // ModelSubtitleDeliveryMethod | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h264, mpeg4, theora, vpx, wmv. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideoHlsServiceAPI.GetVideosByIdLiveM3u8(context.Background(), id).Container(container).DeviceProfileId(deviceProfileId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AudioSampleRate(audioSampleRate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Static(static).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).VideoCodec(videoCodec).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoHlsServiceAPI.GetVideosByIdLiveM3u8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Item Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosByIdLiveM3u8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | Container | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100 | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2 | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2 | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false | 
 **profile** | **string** | Optional. Specify a specific h264 profile, e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the h264 profile, e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1ms &#x3D; 10000 ticks. | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | [**ModelSubtitleDeliveryMethod**](ModelSubtitleDeliveryMethod.md) | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h264, mpeg4, theora, vpx, wmv. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 

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

