# ModelTranscodingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioCodec** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**SubProtocol** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**IsVideoDirect** | Pointer to **bool** |  | [optional] 
**IsAudioDirect** | Pointer to **bool** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**AudioBitrate** | Pointer to **NullableInt32** |  | [optional] 
**VideoBitrate** | Pointer to **NullableInt32** |  | [optional] 
**Framerate** | Pointer to **NullableFloat32** |  | [optional] 
**CompletionPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**TranscodingPositionTicks** | Pointer to **NullableFloat64** |  | [optional] 
**TranscodingStartPositionTicks** | Pointer to **NullableFloat64** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**AudioChannels** | Pointer to **NullableInt32** |  | [optional] 
**TranscodeReasons** | Pointer to [**[]ModelModelTranscodeReason**](ModelModelTranscodeReason.md) |  | [optional] 
**CurrentCpuUsage** | Pointer to **NullableFloat64** | Deprecated, please use ProcessStatistics instead | [optional] 
**AverageCpuUsage** | Pointer to **NullableFloat64** | Deprecated, please use ProcessStatistics instead | [optional] 
**CpuHistory** | Pointer to [**[]ModelModelTupleDoubleDouble**](ModelModelTupleDoubleDouble.md) | Deprecated, please use ProcessStatistics instead | [optional] 
**ProcessStatistics** | Pointer to [**ModelModelProcessRunMetricsProcessStatistics**](ModelProcessRunMetricsProcessStatistics.md) |  | [optional] 
**CurrentThrottle** | Pointer to **NullableInt32** |  | [optional] 
**VideoDecoder** | Pointer to **string** |  | [optional] 
**VideoDecoderIsHardware** | Pointer to **bool** |  | [optional] 
**VideoDecoderMediaType** | Pointer to **string** |  | [optional] 
**VideoDecoderHwAccel** | Pointer to **string** |  | [optional] 
**VideoEncoder** | Pointer to **string** |  | [optional] 
**VideoEncoderIsHardware** | Pointer to **bool** |  | [optional] 
**VideoEncoderMediaType** | Pointer to **string** |  | [optional] 
**VideoEncoderHwAccel** | Pointer to **string** |  | [optional] 
**VideoPipelineInfo** | Pointer to [**[]ModelModelTranscodingVpStepInfo**](ModelModelTranscodingVpStepInfo.md) |  | [optional] 
**SubtitlePipelineInfos** | Pointer to [**[][]ModelModelTranscodingVpStepInfo**]([]ModelModelTranscodingVpStepInfo.md) |  | [optional] 

## Methods

### NewModelTranscodingInfo

`func NewModelTranscodingInfo() *ModelTranscodingInfo`

NewModelTranscodingInfo instantiates a new ModelTranscodingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTranscodingInfoWithDefaults

`func NewModelTranscodingInfoWithDefaults() *ModelTranscodingInfo`

NewModelTranscodingInfoWithDefaults instantiates a new ModelTranscodingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioCodec

`func (o *ModelTranscodingInfo) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelTranscodingInfo) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelTranscodingInfo) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelTranscodingInfo) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoCodec

`func (o *ModelTranscodingInfo) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelTranscodingInfo) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelTranscodingInfo) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelTranscodingInfo) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetSubProtocol

`func (o *ModelTranscodingInfo) GetSubProtocol() string`

GetSubProtocol returns the SubProtocol field if non-nil, zero value otherwise.

### GetSubProtocolOk

`func (o *ModelTranscodingInfo) GetSubProtocolOk() (*string, bool)`

GetSubProtocolOk returns a tuple with the SubProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubProtocol

`func (o *ModelTranscodingInfo) SetSubProtocol(v string)`

SetSubProtocol sets SubProtocol field to given value.

### HasSubProtocol

`func (o *ModelTranscodingInfo) HasSubProtocol() bool`

HasSubProtocol returns a boolean if a field has been set.

### GetContainer

`func (o *ModelTranscodingInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelTranscodingInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelTranscodingInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelTranscodingInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetIsVideoDirect

`func (o *ModelTranscodingInfo) GetIsVideoDirect() bool`

GetIsVideoDirect returns the IsVideoDirect field if non-nil, zero value otherwise.

### GetIsVideoDirectOk

`func (o *ModelTranscodingInfo) GetIsVideoDirectOk() (*bool, bool)`

GetIsVideoDirectOk returns a tuple with the IsVideoDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideoDirect

`func (o *ModelTranscodingInfo) SetIsVideoDirect(v bool)`

SetIsVideoDirect sets IsVideoDirect field to given value.

### HasIsVideoDirect

`func (o *ModelTranscodingInfo) HasIsVideoDirect() bool`

HasIsVideoDirect returns a boolean if a field has been set.

### GetIsAudioDirect

`func (o *ModelTranscodingInfo) GetIsAudioDirect() bool`

GetIsAudioDirect returns the IsAudioDirect field if non-nil, zero value otherwise.

### GetIsAudioDirectOk

`func (o *ModelTranscodingInfo) GetIsAudioDirectOk() (*bool, bool)`

GetIsAudioDirectOk returns a tuple with the IsAudioDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAudioDirect

`func (o *ModelTranscodingInfo) SetIsAudioDirect(v bool)`

SetIsAudioDirect sets IsAudioDirect field to given value.

### HasIsAudioDirect

`func (o *ModelTranscodingInfo) HasIsAudioDirect() bool`

HasIsAudioDirect returns a boolean if a field has been set.

### GetBitrate

`func (o *ModelTranscodingInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelTranscodingInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelTranscodingInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelTranscodingInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelTranscodingInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelTranscodingInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetAudioBitrate

`func (o *ModelTranscodingInfo) GetAudioBitrate() int32`

GetAudioBitrate returns the AudioBitrate field if non-nil, zero value otherwise.

### GetAudioBitrateOk

`func (o *ModelTranscodingInfo) GetAudioBitrateOk() (*int32, bool)`

GetAudioBitrateOk returns a tuple with the AudioBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBitrate

`func (o *ModelTranscodingInfo) SetAudioBitrate(v int32)`

SetAudioBitrate sets AudioBitrate field to given value.

### HasAudioBitrate

`func (o *ModelTranscodingInfo) HasAudioBitrate() bool`

HasAudioBitrate returns a boolean if a field has been set.

### SetAudioBitrateNil

`func (o *ModelTranscodingInfo) SetAudioBitrateNil(b bool)`

 SetAudioBitrateNil sets the value for AudioBitrate to be an explicit nil

### UnsetAudioBitrate
`func (o *ModelTranscodingInfo) UnsetAudioBitrate()`

UnsetAudioBitrate ensures that no value is present for AudioBitrate, not even an explicit nil
### GetVideoBitrate

`func (o *ModelTranscodingInfo) GetVideoBitrate() int32`

GetVideoBitrate returns the VideoBitrate field if non-nil, zero value otherwise.

### GetVideoBitrateOk

`func (o *ModelTranscodingInfo) GetVideoBitrateOk() (*int32, bool)`

GetVideoBitrateOk returns a tuple with the VideoBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoBitrate

`func (o *ModelTranscodingInfo) SetVideoBitrate(v int32)`

SetVideoBitrate sets VideoBitrate field to given value.

### HasVideoBitrate

`func (o *ModelTranscodingInfo) HasVideoBitrate() bool`

HasVideoBitrate returns a boolean if a field has been set.

### SetVideoBitrateNil

`func (o *ModelTranscodingInfo) SetVideoBitrateNil(b bool)`

 SetVideoBitrateNil sets the value for VideoBitrate to be an explicit nil

### UnsetVideoBitrate
`func (o *ModelTranscodingInfo) UnsetVideoBitrate()`

UnsetVideoBitrate ensures that no value is present for VideoBitrate, not even an explicit nil
### GetFramerate

`func (o *ModelTranscodingInfo) GetFramerate() float32`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *ModelTranscodingInfo) GetFramerateOk() (*float32, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *ModelTranscodingInfo) SetFramerate(v float32)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *ModelTranscodingInfo) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### SetFramerateNil

`func (o *ModelTranscodingInfo) SetFramerateNil(b bool)`

 SetFramerateNil sets the value for Framerate to be an explicit nil

### UnsetFramerate
`func (o *ModelTranscodingInfo) UnsetFramerate()`

UnsetFramerate ensures that no value is present for Framerate, not even an explicit nil
### GetCompletionPercentage

`func (o *ModelTranscodingInfo) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *ModelTranscodingInfo) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *ModelTranscodingInfo) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *ModelTranscodingInfo) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *ModelTranscodingInfo) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *ModelTranscodingInfo) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetTranscodingPositionTicks

`func (o *ModelTranscodingInfo) GetTranscodingPositionTicks() float64`

GetTranscodingPositionTicks returns the TranscodingPositionTicks field if non-nil, zero value otherwise.

### GetTranscodingPositionTicksOk

`func (o *ModelTranscodingInfo) GetTranscodingPositionTicksOk() (*float64, bool)`

GetTranscodingPositionTicksOk returns a tuple with the TranscodingPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingPositionTicks

`func (o *ModelTranscodingInfo) SetTranscodingPositionTicks(v float64)`

SetTranscodingPositionTicks sets TranscodingPositionTicks field to given value.

### HasTranscodingPositionTicks

`func (o *ModelTranscodingInfo) HasTranscodingPositionTicks() bool`

HasTranscodingPositionTicks returns a boolean if a field has been set.

### SetTranscodingPositionTicksNil

`func (o *ModelTranscodingInfo) SetTranscodingPositionTicksNil(b bool)`

 SetTranscodingPositionTicksNil sets the value for TranscodingPositionTicks to be an explicit nil

### UnsetTranscodingPositionTicks
`func (o *ModelTranscodingInfo) UnsetTranscodingPositionTicks()`

UnsetTranscodingPositionTicks ensures that no value is present for TranscodingPositionTicks, not even an explicit nil
### GetTranscodingStartPositionTicks

`func (o *ModelTranscodingInfo) GetTranscodingStartPositionTicks() float64`

GetTranscodingStartPositionTicks returns the TranscodingStartPositionTicks field if non-nil, zero value otherwise.

### GetTranscodingStartPositionTicksOk

`func (o *ModelTranscodingInfo) GetTranscodingStartPositionTicksOk() (*float64, bool)`

GetTranscodingStartPositionTicksOk returns a tuple with the TranscodingStartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingStartPositionTicks

`func (o *ModelTranscodingInfo) SetTranscodingStartPositionTicks(v float64)`

SetTranscodingStartPositionTicks sets TranscodingStartPositionTicks field to given value.

### HasTranscodingStartPositionTicks

`func (o *ModelTranscodingInfo) HasTranscodingStartPositionTicks() bool`

HasTranscodingStartPositionTicks returns a boolean if a field has been set.

### SetTranscodingStartPositionTicksNil

`func (o *ModelTranscodingInfo) SetTranscodingStartPositionTicksNil(b bool)`

 SetTranscodingStartPositionTicksNil sets the value for TranscodingStartPositionTicks to be an explicit nil

### UnsetTranscodingStartPositionTicks
`func (o *ModelTranscodingInfo) UnsetTranscodingStartPositionTicks()`

UnsetTranscodingStartPositionTicks ensures that no value is present for TranscodingStartPositionTicks, not even an explicit nil
### GetWidth

`func (o *ModelTranscodingInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelTranscodingInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelTranscodingInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelTranscodingInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ModelTranscodingInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ModelTranscodingInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *ModelTranscodingInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelTranscodingInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelTranscodingInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelTranscodingInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ModelTranscodingInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ModelTranscodingInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetAudioChannels

`func (o *ModelTranscodingInfo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *ModelTranscodingInfo) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *ModelTranscodingInfo) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *ModelTranscodingInfo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### SetAudioChannelsNil

`func (o *ModelTranscodingInfo) SetAudioChannelsNil(b bool)`

 SetAudioChannelsNil sets the value for AudioChannels to be an explicit nil

### UnsetAudioChannels
`func (o *ModelTranscodingInfo) UnsetAudioChannels()`

UnsetAudioChannels ensures that no value is present for AudioChannels, not even an explicit nil
### GetTranscodeReasons

`func (o *ModelTranscodingInfo) GetTranscodeReasons() []ModelModelTranscodeReason`

GetTranscodeReasons returns the TranscodeReasons field if non-nil, zero value otherwise.

### GetTranscodeReasonsOk

`func (o *ModelTranscodingInfo) GetTranscodeReasonsOk() (*[]ModelModelTranscodeReason, bool)`

GetTranscodeReasonsOk returns a tuple with the TranscodeReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeReasons

`func (o *ModelTranscodingInfo) SetTranscodeReasons(v []ModelModelTranscodeReason)`

SetTranscodeReasons sets TranscodeReasons field to given value.

### HasTranscodeReasons

`func (o *ModelTranscodingInfo) HasTranscodeReasons() bool`

HasTranscodeReasons returns a boolean if a field has been set.

### GetCurrentCpuUsage

`func (o *ModelTranscodingInfo) GetCurrentCpuUsage() float64`

GetCurrentCpuUsage returns the CurrentCpuUsage field if non-nil, zero value otherwise.

### GetCurrentCpuUsageOk

`func (o *ModelTranscodingInfo) GetCurrentCpuUsageOk() (*float64, bool)`

GetCurrentCpuUsageOk returns a tuple with the CurrentCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCpuUsage

`func (o *ModelTranscodingInfo) SetCurrentCpuUsage(v float64)`

SetCurrentCpuUsage sets CurrentCpuUsage field to given value.

### HasCurrentCpuUsage

`func (o *ModelTranscodingInfo) HasCurrentCpuUsage() bool`

HasCurrentCpuUsage returns a boolean if a field has been set.

### SetCurrentCpuUsageNil

`func (o *ModelTranscodingInfo) SetCurrentCpuUsageNil(b bool)`

 SetCurrentCpuUsageNil sets the value for CurrentCpuUsage to be an explicit nil

### UnsetCurrentCpuUsage
`func (o *ModelTranscodingInfo) UnsetCurrentCpuUsage()`

UnsetCurrentCpuUsage ensures that no value is present for CurrentCpuUsage, not even an explicit nil
### GetAverageCpuUsage

`func (o *ModelTranscodingInfo) GetAverageCpuUsage() float64`

GetAverageCpuUsage returns the AverageCpuUsage field if non-nil, zero value otherwise.

### GetAverageCpuUsageOk

`func (o *ModelTranscodingInfo) GetAverageCpuUsageOk() (*float64, bool)`

GetAverageCpuUsageOk returns a tuple with the AverageCpuUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageCpuUsage

`func (o *ModelTranscodingInfo) SetAverageCpuUsage(v float64)`

SetAverageCpuUsage sets AverageCpuUsage field to given value.

### HasAverageCpuUsage

`func (o *ModelTranscodingInfo) HasAverageCpuUsage() bool`

HasAverageCpuUsage returns a boolean if a field has been set.

### SetAverageCpuUsageNil

`func (o *ModelTranscodingInfo) SetAverageCpuUsageNil(b bool)`

 SetAverageCpuUsageNil sets the value for AverageCpuUsage to be an explicit nil

### UnsetAverageCpuUsage
`func (o *ModelTranscodingInfo) UnsetAverageCpuUsage()`

UnsetAverageCpuUsage ensures that no value is present for AverageCpuUsage, not even an explicit nil
### GetCpuHistory

`func (o *ModelTranscodingInfo) GetCpuHistory() []ModelModelTupleDoubleDouble`

GetCpuHistory returns the CpuHistory field if non-nil, zero value otherwise.

### GetCpuHistoryOk

`func (o *ModelTranscodingInfo) GetCpuHistoryOk() (*[]ModelModelTupleDoubleDouble, bool)`

GetCpuHistoryOk returns a tuple with the CpuHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuHistory

`func (o *ModelTranscodingInfo) SetCpuHistory(v []ModelModelTupleDoubleDouble)`

SetCpuHistory sets CpuHistory field to given value.

### HasCpuHistory

`func (o *ModelTranscodingInfo) HasCpuHistory() bool`

HasCpuHistory returns a boolean if a field has been set.

### GetProcessStatistics

`func (o *ModelTranscodingInfo) GetProcessStatistics() ModelModelProcessRunMetricsProcessStatistics`

GetProcessStatistics returns the ProcessStatistics field if non-nil, zero value otherwise.

### GetProcessStatisticsOk

`func (o *ModelTranscodingInfo) GetProcessStatisticsOk() (*ModelModelProcessRunMetricsProcessStatistics, bool)`

GetProcessStatisticsOk returns a tuple with the ProcessStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStatistics

`func (o *ModelTranscodingInfo) SetProcessStatistics(v ModelModelProcessRunMetricsProcessStatistics)`

SetProcessStatistics sets ProcessStatistics field to given value.

### HasProcessStatistics

`func (o *ModelTranscodingInfo) HasProcessStatistics() bool`

HasProcessStatistics returns a boolean if a field has been set.

### GetCurrentThrottle

`func (o *ModelTranscodingInfo) GetCurrentThrottle() int32`

GetCurrentThrottle returns the CurrentThrottle field if non-nil, zero value otherwise.

### GetCurrentThrottleOk

`func (o *ModelTranscodingInfo) GetCurrentThrottleOk() (*int32, bool)`

GetCurrentThrottleOk returns a tuple with the CurrentThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentThrottle

`func (o *ModelTranscodingInfo) SetCurrentThrottle(v int32)`

SetCurrentThrottle sets CurrentThrottle field to given value.

### HasCurrentThrottle

`func (o *ModelTranscodingInfo) HasCurrentThrottle() bool`

HasCurrentThrottle returns a boolean if a field has been set.

### SetCurrentThrottleNil

`func (o *ModelTranscodingInfo) SetCurrentThrottleNil(b bool)`

 SetCurrentThrottleNil sets the value for CurrentThrottle to be an explicit nil

### UnsetCurrentThrottle
`func (o *ModelTranscodingInfo) UnsetCurrentThrottle()`

UnsetCurrentThrottle ensures that no value is present for CurrentThrottle, not even an explicit nil
### GetVideoDecoder

`func (o *ModelTranscodingInfo) GetVideoDecoder() string`

GetVideoDecoder returns the VideoDecoder field if non-nil, zero value otherwise.

### GetVideoDecoderOk

`func (o *ModelTranscodingInfo) GetVideoDecoderOk() (*string, bool)`

GetVideoDecoderOk returns a tuple with the VideoDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoder

`func (o *ModelTranscodingInfo) SetVideoDecoder(v string)`

SetVideoDecoder sets VideoDecoder field to given value.

### HasVideoDecoder

`func (o *ModelTranscodingInfo) HasVideoDecoder() bool`

HasVideoDecoder returns a boolean if a field has been set.

### GetVideoDecoderIsHardware

`func (o *ModelTranscodingInfo) GetVideoDecoderIsHardware() bool`

GetVideoDecoderIsHardware returns the VideoDecoderIsHardware field if non-nil, zero value otherwise.

### GetVideoDecoderIsHardwareOk

`func (o *ModelTranscodingInfo) GetVideoDecoderIsHardwareOk() (*bool, bool)`

GetVideoDecoderIsHardwareOk returns a tuple with the VideoDecoderIsHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoderIsHardware

`func (o *ModelTranscodingInfo) SetVideoDecoderIsHardware(v bool)`

SetVideoDecoderIsHardware sets VideoDecoderIsHardware field to given value.

### HasVideoDecoderIsHardware

`func (o *ModelTranscodingInfo) HasVideoDecoderIsHardware() bool`

HasVideoDecoderIsHardware returns a boolean if a field has been set.

### GetVideoDecoderMediaType

`func (o *ModelTranscodingInfo) GetVideoDecoderMediaType() string`

GetVideoDecoderMediaType returns the VideoDecoderMediaType field if non-nil, zero value otherwise.

### GetVideoDecoderMediaTypeOk

`func (o *ModelTranscodingInfo) GetVideoDecoderMediaTypeOk() (*string, bool)`

GetVideoDecoderMediaTypeOk returns a tuple with the VideoDecoderMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoderMediaType

`func (o *ModelTranscodingInfo) SetVideoDecoderMediaType(v string)`

SetVideoDecoderMediaType sets VideoDecoderMediaType field to given value.

### HasVideoDecoderMediaType

`func (o *ModelTranscodingInfo) HasVideoDecoderMediaType() bool`

HasVideoDecoderMediaType returns a boolean if a field has been set.

### GetVideoDecoderHwAccel

`func (o *ModelTranscodingInfo) GetVideoDecoderHwAccel() string`

GetVideoDecoderHwAccel returns the VideoDecoderHwAccel field if non-nil, zero value otherwise.

### GetVideoDecoderHwAccelOk

`func (o *ModelTranscodingInfo) GetVideoDecoderHwAccelOk() (*string, bool)`

GetVideoDecoderHwAccelOk returns a tuple with the VideoDecoderHwAccel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDecoderHwAccel

`func (o *ModelTranscodingInfo) SetVideoDecoderHwAccel(v string)`

SetVideoDecoderHwAccel sets VideoDecoderHwAccel field to given value.

### HasVideoDecoderHwAccel

`func (o *ModelTranscodingInfo) HasVideoDecoderHwAccel() bool`

HasVideoDecoderHwAccel returns a boolean if a field has been set.

### GetVideoEncoder

`func (o *ModelTranscodingInfo) GetVideoEncoder() string`

GetVideoEncoder returns the VideoEncoder field if non-nil, zero value otherwise.

### GetVideoEncoderOk

`func (o *ModelTranscodingInfo) GetVideoEncoderOk() (*string, bool)`

GetVideoEncoderOk returns a tuple with the VideoEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoder

`func (o *ModelTranscodingInfo) SetVideoEncoder(v string)`

SetVideoEncoder sets VideoEncoder field to given value.

### HasVideoEncoder

`func (o *ModelTranscodingInfo) HasVideoEncoder() bool`

HasVideoEncoder returns a boolean if a field has been set.

### GetVideoEncoderIsHardware

`func (o *ModelTranscodingInfo) GetVideoEncoderIsHardware() bool`

GetVideoEncoderIsHardware returns the VideoEncoderIsHardware field if non-nil, zero value otherwise.

### GetVideoEncoderIsHardwareOk

`func (o *ModelTranscodingInfo) GetVideoEncoderIsHardwareOk() (*bool, bool)`

GetVideoEncoderIsHardwareOk returns a tuple with the VideoEncoderIsHardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoderIsHardware

`func (o *ModelTranscodingInfo) SetVideoEncoderIsHardware(v bool)`

SetVideoEncoderIsHardware sets VideoEncoderIsHardware field to given value.

### HasVideoEncoderIsHardware

`func (o *ModelTranscodingInfo) HasVideoEncoderIsHardware() bool`

HasVideoEncoderIsHardware returns a boolean if a field has been set.

### GetVideoEncoderMediaType

`func (o *ModelTranscodingInfo) GetVideoEncoderMediaType() string`

GetVideoEncoderMediaType returns the VideoEncoderMediaType field if non-nil, zero value otherwise.

### GetVideoEncoderMediaTypeOk

`func (o *ModelTranscodingInfo) GetVideoEncoderMediaTypeOk() (*string, bool)`

GetVideoEncoderMediaTypeOk returns a tuple with the VideoEncoderMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoderMediaType

`func (o *ModelTranscodingInfo) SetVideoEncoderMediaType(v string)`

SetVideoEncoderMediaType sets VideoEncoderMediaType field to given value.

### HasVideoEncoderMediaType

`func (o *ModelTranscodingInfo) HasVideoEncoderMediaType() bool`

HasVideoEncoderMediaType returns a boolean if a field has been set.

### GetVideoEncoderHwAccel

`func (o *ModelTranscodingInfo) GetVideoEncoderHwAccel() string`

GetVideoEncoderHwAccel returns the VideoEncoderHwAccel field if non-nil, zero value otherwise.

### GetVideoEncoderHwAccelOk

`func (o *ModelTranscodingInfo) GetVideoEncoderHwAccelOk() (*string, bool)`

GetVideoEncoderHwAccelOk returns a tuple with the VideoEncoderHwAccel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoEncoderHwAccel

`func (o *ModelTranscodingInfo) SetVideoEncoderHwAccel(v string)`

SetVideoEncoderHwAccel sets VideoEncoderHwAccel field to given value.

### HasVideoEncoderHwAccel

`func (o *ModelTranscodingInfo) HasVideoEncoderHwAccel() bool`

HasVideoEncoderHwAccel returns a boolean if a field has been set.

### GetVideoPipelineInfo

`func (o *ModelTranscodingInfo) GetVideoPipelineInfo() []ModelModelTranscodingVpStepInfo`

GetVideoPipelineInfo returns the VideoPipelineInfo field if non-nil, zero value otherwise.

### GetVideoPipelineInfoOk

`func (o *ModelTranscodingInfo) GetVideoPipelineInfoOk() (*[]ModelModelTranscodingVpStepInfo, bool)`

GetVideoPipelineInfoOk returns a tuple with the VideoPipelineInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoPipelineInfo

`func (o *ModelTranscodingInfo) SetVideoPipelineInfo(v []ModelModelTranscodingVpStepInfo)`

SetVideoPipelineInfo sets VideoPipelineInfo field to given value.

### HasVideoPipelineInfo

`func (o *ModelTranscodingInfo) HasVideoPipelineInfo() bool`

HasVideoPipelineInfo returns a boolean if a field has been set.

### GetSubtitlePipelineInfos

`func (o *ModelTranscodingInfo) GetSubtitlePipelineInfos() [][]ModelModelTranscodingVpStepInfo`

GetSubtitlePipelineInfos returns the SubtitlePipelineInfos field if non-nil, zero value otherwise.

### GetSubtitlePipelineInfosOk

`func (o *ModelTranscodingInfo) GetSubtitlePipelineInfosOk() (*[][]ModelModelTranscodingVpStepInfo, bool)`

GetSubtitlePipelineInfosOk returns a tuple with the SubtitlePipelineInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitlePipelineInfos

`func (o *ModelTranscodingInfo) SetSubtitlePipelineInfos(v [][]ModelModelTranscodingVpStepInfo)`

SetSubtitlePipelineInfos sets SubtitlePipelineInfos field to given value.

### HasSubtitlePipelineInfos

`func (o *ModelTranscodingInfo) HasSubtitlePipelineInfos() bool`

HasSubtitlePipelineInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


