# ModelMediaSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**ModelModelMediaProtocol**](ModelMediaProtocol.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**EncoderPath** | Pointer to **string** |  | [optional] 
**EncoderProtocol** | Pointer to [**ModelModelMediaProtocol**](ModelMediaProtocol.md) |  | [optional] 
**Type** | Pointer to [**ModelModelMediaSourceType**](ModelMediaSourceType.md) |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SortName** | Pointer to **string** |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**ContainerStartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**SupportsTranscoding** | Pointer to **bool** |  | [optional] 
**SupportsDirectStream** | Pointer to **bool** |  | [optional] 
**SupportsDirectPlay** | Pointer to **bool** |  | [optional] 
**IsInfiniteStream** | Pointer to **bool** |  | [optional] 
**RequiresOpening** | Pointer to **bool** |  | [optional] 
**OpenToken** | Pointer to **string** |  | [optional] 
**RequiresClosing** | Pointer to **bool** |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**BufferMs** | Pointer to **NullableInt32** |  | [optional] 
**RequiresLooping** | Pointer to **bool** |  | [optional] 
**SupportsProbing** | Pointer to **bool** |  | [optional] 
**Video3DFormat** | Pointer to [**ModelModelVideo3DFormat**](ModelVideo3DFormat.md) |  | [optional] 
**MediaStreams** | Pointer to [**[]ModelModelMediaStream**](ModelModelMediaStream.md) |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to [**ModelModelTransportStreamTimestamp**](ModelTransportStreamTimestamp.md) |  | [optional] 
**RequiredHttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**DirectStreamUrl** | Pointer to **string** |  | [optional] 
**TranscodingUrl** | Pointer to **string** |  | [optional] 
**TranscodingSubProtocol** | Pointer to **string** |  | [optional] 
**TranscodingContainer** | Pointer to **string** |  | [optional] 
**AnalyzeDurationMs** | Pointer to **NullableInt32** |  | [optional] 
**ReadAtNativeFramerate** | Pointer to **bool** |  | [optional] 
**DefaultAudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelMediaSourceInfo

`func NewModelMediaSourceInfo() *ModelMediaSourceInfo`

NewModelMediaSourceInfo instantiates a new ModelMediaSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMediaSourceInfoWithDefaults

`func NewModelMediaSourceInfoWithDefaults() *ModelMediaSourceInfo`

NewModelMediaSourceInfoWithDefaults instantiates a new ModelMediaSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *ModelMediaSourceInfo) GetProtocol() ModelModelMediaProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelMediaSourceInfo) GetProtocolOk() (*ModelModelMediaProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelMediaSourceInfo) SetProtocol(v ModelModelMediaProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ModelMediaSourceInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetId

`func (o *ModelMediaSourceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelMediaSourceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelMediaSourceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelMediaSourceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *ModelMediaSourceInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelMediaSourceInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelMediaSourceInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelMediaSourceInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetEncoderPath

`func (o *ModelMediaSourceInfo) GetEncoderPath() string`

GetEncoderPath returns the EncoderPath field if non-nil, zero value otherwise.

### GetEncoderPathOk

`func (o *ModelMediaSourceInfo) GetEncoderPathOk() (*string, bool)`

GetEncoderPathOk returns a tuple with the EncoderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPath

`func (o *ModelMediaSourceInfo) SetEncoderPath(v string)`

SetEncoderPath sets EncoderPath field to given value.

### HasEncoderPath

`func (o *ModelMediaSourceInfo) HasEncoderPath() bool`

HasEncoderPath returns a boolean if a field has been set.

### GetEncoderProtocol

`func (o *ModelMediaSourceInfo) GetEncoderProtocol() ModelModelMediaProtocol`

GetEncoderProtocol returns the EncoderProtocol field if non-nil, zero value otherwise.

### GetEncoderProtocolOk

`func (o *ModelMediaSourceInfo) GetEncoderProtocolOk() (*ModelModelMediaProtocol, bool)`

GetEncoderProtocolOk returns a tuple with the EncoderProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderProtocol

`func (o *ModelMediaSourceInfo) SetEncoderProtocol(v ModelModelMediaProtocol)`

SetEncoderProtocol sets EncoderProtocol field to given value.

### HasEncoderProtocol

`func (o *ModelMediaSourceInfo) HasEncoderProtocol() bool`

HasEncoderProtocol returns a boolean if a field has been set.

### GetType

`func (o *ModelMediaSourceInfo) GetType() ModelModelMediaSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelMediaSourceInfo) GetTypeOk() (*ModelModelMediaSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelMediaSourceInfo) SetType(v ModelModelMediaSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelMediaSourceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContainer

`func (o *ModelMediaSourceInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelMediaSourceInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelMediaSourceInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelMediaSourceInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetSize

`func (o *ModelMediaSourceInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ModelMediaSourceInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ModelMediaSourceInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ModelMediaSourceInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ModelMediaSourceInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ModelMediaSourceInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetName

`func (o *ModelMediaSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMediaSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMediaSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelMediaSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSortName

`func (o *ModelMediaSourceInfo) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *ModelMediaSourceInfo) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *ModelMediaSourceInfo) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *ModelMediaSourceInfo) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### GetIsRemote

`func (o *ModelMediaSourceInfo) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *ModelMediaSourceInfo) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *ModelMediaSourceInfo) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *ModelMediaSourceInfo) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetRunTimeTicks

`func (o *ModelMediaSourceInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *ModelMediaSourceInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *ModelMediaSourceInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *ModelMediaSourceInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *ModelMediaSourceInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *ModelMediaSourceInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetContainerStartTimeTicks

`func (o *ModelMediaSourceInfo) GetContainerStartTimeTicks() int64`

GetContainerStartTimeTicks returns the ContainerStartTimeTicks field if non-nil, zero value otherwise.

### GetContainerStartTimeTicksOk

`func (o *ModelMediaSourceInfo) GetContainerStartTimeTicksOk() (*int64, bool)`

GetContainerStartTimeTicksOk returns a tuple with the ContainerStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStartTimeTicks

`func (o *ModelMediaSourceInfo) SetContainerStartTimeTicks(v int64)`

SetContainerStartTimeTicks sets ContainerStartTimeTicks field to given value.

### HasContainerStartTimeTicks

`func (o *ModelMediaSourceInfo) HasContainerStartTimeTicks() bool`

HasContainerStartTimeTicks returns a boolean if a field has been set.

### SetContainerStartTimeTicksNil

`func (o *ModelMediaSourceInfo) SetContainerStartTimeTicksNil(b bool)`

 SetContainerStartTimeTicksNil sets the value for ContainerStartTimeTicks to be an explicit nil

### UnsetContainerStartTimeTicks
`func (o *ModelMediaSourceInfo) UnsetContainerStartTimeTicks()`

UnsetContainerStartTimeTicks ensures that no value is present for ContainerStartTimeTicks, not even an explicit nil
### GetSupportsTranscoding

`func (o *ModelMediaSourceInfo) GetSupportsTranscoding() bool`

GetSupportsTranscoding returns the SupportsTranscoding field if non-nil, zero value otherwise.

### GetSupportsTranscodingOk

`func (o *ModelMediaSourceInfo) GetSupportsTranscodingOk() (*bool, bool)`

GetSupportsTranscodingOk returns a tuple with the SupportsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTranscoding

`func (o *ModelMediaSourceInfo) SetSupportsTranscoding(v bool)`

SetSupportsTranscoding sets SupportsTranscoding field to given value.

### HasSupportsTranscoding

`func (o *ModelMediaSourceInfo) HasSupportsTranscoding() bool`

HasSupportsTranscoding returns a boolean if a field has been set.

### GetSupportsDirectStream

`func (o *ModelMediaSourceInfo) GetSupportsDirectStream() bool`

GetSupportsDirectStream returns the SupportsDirectStream field if non-nil, zero value otherwise.

### GetSupportsDirectStreamOk

`func (o *ModelMediaSourceInfo) GetSupportsDirectStreamOk() (*bool, bool)`

GetSupportsDirectStreamOk returns a tuple with the SupportsDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectStream

`func (o *ModelMediaSourceInfo) SetSupportsDirectStream(v bool)`

SetSupportsDirectStream sets SupportsDirectStream field to given value.

### HasSupportsDirectStream

`func (o *ModelMediaSourceInfo) HasSupportsDirectStream() bool`

HasSupportsDirectStream returns a boolean if a field has been set.

### GetSupportsDirectPlay

`func (o *ModelMediaSourceInfo) GetSupportsDirectPlay() bool`

GetSupportsDirectPlay returns the SupportsDirectPlay field if non-nil, zero value otherwise.

### GetSupportsDirectPlayOk

`func (o *ModelMediaSourceInfo) GetSupportsDirectPlayOk() (*bool, bool)`

GetSupportsDirectPlayOk returns a tuple with the SupportsDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectPlay

`func (o *ModelMediaSourceInfo) SetSupportsDirectPlay(v bool)`

SetSupportsDirectPlay sets SupportsDirectPlay field to given value.

### HasSupportsDirectPlay

`func (o *ModelMediaSourceInfo) HasSupportsDirectPlay() bool`

HasSupportsDirectPlay returns a boolean if a field has been set.

### GetIsInfiniteStream

`func (o *ModelMediaSourceInfo) GetIsInfiniteStream() bool`

GetIsInfiniteStream returns the IsInfiniteStream field if non-nil, zero value otherwise.

### GetIsInfiniteStreamOk

`func (o *ModelMediaSourceInfo) GetIsInfiniteStreamOk() (*bool, bool)`

GetIsInfiniteStreamOk returns a tuple with the IsInfiniteStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfiniteStream

`func (o *ModelMediaSourceInfo) SetIsInfiniteStream(v bool)`

SetIsInfiniteStream sets IsInfiniteStream field to given value.

### HasIsInfiniteStream

`func (o *ModelMediaSourceInfo) HasIsInfiniteStream() bool`

HasIsInfiniteStream returns a boolean if a field has been set.

### GetRequiresOpening

`func (o *ModelMediaSourceInfo) GetRequiresOpening() bool`

GetRequiresOpening returns the RequiresOpening field if non-nil, zero value otherwise.

### GetRequiresOpeningOk

`func (o *ModelMediaSourceInfo) GetRequiresOpeningOk() (*bool, bool)`

GetRequiresOpeningOk returns a tuple with the RequiresOpening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOpening

`func (o *ModelMediaSourceInfo) SetRequiresOpening(v bool)`

SetRequiresOpening sets RequiresOpening field to given value.

### HasRequiresOpening

`func (o *ModelMediaSourceInfo) HasRequiresOpening() bool`

HasRequiresOpening returns a boolean if a field has been set.

### GetOpenToken

`func (o *ModelMediaSourceInfo) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *ModelMediaSourceInfo) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *ModelMediaSourceInfo) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *ModelMediaSourceInfo) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### GetRequiresClosing

`func (o *ModelMediaSourceInfo) GetRequiresClosing() bool`

GetRequiresClosing returns the RequiresClosing field if non-nil, zero value otherwise.

### GetRequiresClosingOk

`func (o *ModelMediaSourceInfo) GetRequiresClosingOk() (*bool, bool)`

GetRequiresClosingOk returns a tuple with the RequiresClosing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresClosing

`func (o *ModelMediaSourceInfo) SetRequiresClosing(v bool)`

SetRequiresClosing sets RequiresClosing field to given value.

### HasRequiresClosing

`func (o *ModelMediaSourceInfo) HasRequiresClosing() bool`

HasRequiresClosing returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *ModelMediaSourceInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *ModelMediaSourceInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *ModelMediaSourceInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *ModelMediaSourceInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetBufferMs

`func (o *ModelMediaSourceInfo) GetBufferMs() int32`

GetBufferMs returns the BufferMs field if non-nil, zero value otherwise.

### GetBufferMsOk

`func (o *ModelMediaSourceInfo) GetBufferMsOk() (*int32, bool)`

GetBufferMsOk returns a tuple with the BufferMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMs

`func (o *ModelMediaSourceInfo) SetBufferMs(v int32)`

SetBufferMs sets BufferMs field to given value.

### HasBufferMs

`func (o *ModelMediaSourceInfo) HasBufferMs() bool`

HasBufferMs returns a boolean if a field has been set.

### SetBufferMsNil

`func (o *ModelMediaSourceInfo) SetBufferMsNil(b bool)`

 SetBufferMsNil sets the value for BufferMs to be an explicit nil

### UnsetBufferMs
`func (o *ModelMediaSourceInfo) UnsetBufferMs()`

UnsetBufferMs ensures that no value is present for BufferMs, not even an explicit nil
### GetRequiresLooping

`func (o *ModelMediaSourceInfo) GetRequiresLooping() bool`

GetRequiresLooping returns the RequiresLooping field if non-nil, zero value otherwise.

### GetRequiresLoopingOk

`func (o *ModelMediaSourceInfo) GetRequiresLoopingOk() (*bool, bool)`

GetRequiresLoopingOk returns a tuple with the RequiresLooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresLooping

`func (o *ModelMediaSourceInfo) SetRequiresLooping(v bool)`

SetRequiresLooping sets RequiresLooping field to given value.

### HasRequiresLooping

`func (o *ModelMediaSourceInfo) HasRequiresLooping() bool`

HasRequiresLooping returns a boolean if a field has been set.

### GetSupportsProbing

`func (o *ModelMediaSourceInfo) GetSupportsProbing() bool`

GetSupportsProbing returns the SupportsProbing field if non-nil, zero value otherwise.

### GetSupportsProbingOk

`func (o *ModelMediaSourceInfo) GetSupportsProbingOk() (*bool, bool)`

GetSupportsProbingOk returns a tuple with the SupportsProbing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsProbing

`func (o *ModelMediaSourceInfo) SetSupportsProbing(v bool)`

SetSupportsProbing sets SupportsProbing field to given value.

### HasSupportsProbing

`func (o *ModelMediaSourceInfo) HasSupportsProbing() bool`

HasSupportsProbing returns a boolean if a field has been set.

### GetVideo3DFormat

`func (o *ModelMediaSourceInfo) GetVideo3DFormat() ModelModelVideo3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *ModelMediaSourceInfo) GetVideo3DFormatOk() (*ModelModelVideo3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *ModelMediaSourceInfo) SetVideo3DFormat(v ModelModelVideo3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *ModelMediaSourceInfo) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### GetMediaStreams

`func (o *ModelMediaSourceInfo) GetMediaStreams() []ModelModelMediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *ModelMediaSourceInfo) GetMediaStreamsOk() (*[]ModelModelMediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *ModelMediaSourceInfo) SetMediaStreams(v []ModelModelMediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *ModelMediaSourceInfo) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### GetFormats

`func (o *ModelMediaSourceInfo) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *ModelMediaSourceInfo) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *ModelMediaSourceInfo) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *ModelMediaSourceInfo) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### GetBitrate

`func (o *ModelMediaSourceInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelMediaSourceInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelMediaSourceInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelMediaSourceInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelMediaSourceInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelMediaSourceInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetTimestamp

`func (o *ModelMediaSourceInfo) GetTimestamp() ModelModelTransportStreamTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelMediaSourceInfo) GetTimestampOk() (*ModelModelTransportStreamTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelMediaSourceInfo) SetTimestamp(v ModelModelTransportStreamTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelMediaSourceInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRequiredHttpHeaders

`func (o *ModelMediaSourceInfo) GetRequiredHttpHeaders() map[string]string`

GetRequiredHttpHeaders returns the RequiredHttpHeaders field if non-nil, zero value otherwise.

### GetRequiredHttpHeadersOk

`func (o *ModelMediaSourceInfo) GetRequiredHttpHeadersOk() (*map[string]string, bool)`

GetRequiredHttpHeadersOk returns a tuple with the RequiredHttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredHttpHeaders

`func (o *ModelMediaSourceInfo) SetRequiredHttpHeaders(v map[string]string)`

SetRequiredHttpHeaders sets RequiredHttpHeaders field to given value.

### HasRequiredHttpHeaders

`func (o *ModelMediaSourceInfo) HasRequiredHttpHeaders() bool`

HasRequiredHttpHeaders returns a boolean if a field has been set.

### GetDirectStreamUrl

`func (o *ModelMediaSourceInfo) GetDirectStreamUrl() string`

GetDirectStreamUrl returns the DirectStreamUrl field if non-nil, zero value otherwise.

### GetDirectStreamUrlOk

`func (o *ModelMediaSourceInfo) GetDirectStreamUrlOk() (*string, bool)`

GetDirectStreamUrlOk returns a tuple with the DirectStreamUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectStreamUrl

`func (o *ModelMediaSourceInfo) SetDirectStreamUrl(v string)`

SetDirectStreamUrl sets DirectStreamUrl field to given value.

### HasDirectStreamUrl

`func (o *ModelMediaSourceInfo) HasDirectStreamUrl() bool`

HasDirectStreamUrl returns a boolean if a field has been set.

### GetTranscodingUrl

`func (o *ModelMediaSourceInfo) GetTranscodingUrl() string`

GetTranscodingUrl returns the TranscodingUrl field if non-nil, zero value otherwise.

### GetTranscodingUrlOk

`func (o *ModelMediaSourceInfo) GetTranscodingUrlOk() (*string, bool)`

GetTranscodingUrlOk returns a tuple with the TranscodingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingUrl

`func (o *ModelMediaSourceInfo) SetTranscodingUrl(v string)`

SetTranscodingUrl sets TranscodingUrl field to given value.

### HasTranscodingUrl

`func (o *ModelMediaSourceInfo) HasTranscodingUrl() bool`

HasTranscodingUrl returns a boolean if a field has been set.

### GetTranscodingSubProtocol

`func (o *ModelMediaSourceInfo) GetTranscodingSubProtocol() string`

GetTranscodingSubProtocol returns the TranscodingSubProtocol field if non-nil, zero value otherwise.

### GetTranscodingSubProtocolOk

`func (o *ModelMediaSourceInfo) GetTranscodingSubProtocolOk() (*string, bool)`

GetTranscodingSubProtocolOk returns a tuple with the TranscodingSubProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingSubProtocol

`func (o *ModelMediaSourceInfo) SetTranscodingSubProtocol(v string)`

SetTranscodingSubProtocol sets TranscodingSubProtocol field to given value.

### HasTranscodingSubProtocol

`func (o *ModelMediaSourceInfo) HasTranscodingSubProtocol() bool`

HasTranscodingSubProtocol returns a boolean if a field has been set.

### GetTranscodingContainer

`func (o *ModelMediaSourceInfo) GetTranscodingContainer() string`

GetTranscodingContainer returns the TranscodingContainer field if non-nil, zero value otherwise.

### GetTranscodingContainerOk

`func (o *ModelMediaSourceInfo) GetTranscodingContainerOk() (*string, bool)`

GetTranscodingContainerOk returns a tuple with the TranscodingContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingContainer

`func (o *ModelMediaSourceInfo) SetTranscodingContainer(v string)`

SetTranscodingContainer sets TranscodingContainer field to given value.

### HasTranscodingContainer

`func (o *ModelMediaSourceInfo) HasTranscodingContainer() bool`

HasTranscodingContainer returns a boolean if a field has been set.

### GetAnalyzeDurationMs

`func (o *ModelMediaSourceInfo) GetAnalyzeDurationMs() int32`

GetAnalyzeDurationMs returns the AnalyzeDurationMs field if non-nil, zero value otherwise.

### GetAnalyzeDurationMsOk

`func (o *ModelMediaSourceInfo) GetAnalyzeDurationMsOk() (*int32, bool)`

GetAnalyzeDurationMsOk returns a tuple with the AnalyzeDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeDurationMs

`func (o *ModelMediaSourceInfo) SetAnalyzeDurationMs(v int32)`

SetAnalyzeDurationMs sets AnalyzeDurationMs field to given value.

### HasAnalyzeDurationMs

`func (o *ModelMediaSourceInfo) HasAnalyzeDurationMs() bool`

HasAnalyzeDurationMs returns a boolean if a field has been set.

### SetAnalyzeDurationMsNil

`func (o *ModelMediaSourceInfo) SetAnalyzeDurationMsNil(b bool)`

 SetAnalyzeDurationMsNil sets the value for AnalyzeDurationMs to be an explicit nil

### UnsetAnalyzeDurationMs
`func (o *ModelMediaSourceInfo) UnsetAnalyzeDurationMs()`

UnsetAnalyzeDurationMs ensures that no value is present for AnalyzeDurationMs, not even an explicit nil
### GetReadAtNativeFramerate

`func (o *ModelMediaSourceInfo) GetReadAtNativeFramerate() bool`

GetReadAtNativeFramerate returns the ReadAtNativeFramerate field if non-nil, zero value otherwise.

### GetReadAtNativeFramerateOk

`func (o *ModelMediaSourceInfo) GetReadAtNativeFramerateOk() (*bool, bool)`

GetReadAtNativeFramerateOk returns a tuple with the ReadAtNativeFramerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAtNativeFramerate

`func (o *ModelMediaSourceInfo) SetReadAtNativeFramerate(v bool)`

SetReadAtNativeFramerate sets ReadAtNativeFramerate field to given value.

### HasReadAtNativeFramerate

`func (o *ModelMediaSourceInfo) HasReadAtNativeFramerate() bool`

HasReadAtNativeFramerate returns a boolean if a field has been set.

### GetDefaultAudioStreamIndex

`func (o *ModelMediaSourceInfo) GetDefaultAudioStreamIndex() int32`

GetDefaultAudioStreamIndex returns the DefaultAudioStreamIndex field if non-nil, zero value otherwise.

### GetDefaultAudioStreamIndexOk

`func (o *ModelMediaSourceInfo) GetDefaultAudioStreamIndexOk() (*int32, bool)`

GetDefaultAudioStreamIndexOk returns a tuple with the DefaultAudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAudioStreamIndex

`func (o *ModelMediaSourceInfo) SetDefaultAudioStreamIndex(v int32)`

SetDefaultAudioStreamIndex sets DefaultAudioStreamIndex field to given value.

### HasDefaultAudioStreamIndex

`func (o *ModelMediaSourceInfo) HasDefaultAudioStreamIndex() bool`

HasDefaultAudioStreamIndex returns a boolean if a field has been set.

### SetDefaultAudioStreamIndexNil

`func (o *ModelMediaSourceInfo) SetDefaultAudioStreamIndexNil(b bool)`

 SetDefaultAudioStreamIndexNil sets the value for DefaultAudioStreamIndex to be an explicit nil

### UnsetDefaultAudioStreamIndex
`func (o *ModelMediaSourceInfo) UnsetDefaultAudioStreamIndex()`

UnsetDefaultAudioStreamIndex ensures that no value is present for DefaultAudioStreamIndex, not even an explicit nil
### GetDefaultSubtitleStreamIndex

`func (o *ModelMediaSourceInfo) GetDefaultSubtitleStreamIndex() int32`

GetDefaultSubtitleStreamIndex returns the DefaultSubtitleStreamIndex field if non-nil, zero value otherwise.

### GetDefaultSubtitleStreamIndexOk

`func (o *ModelMediaSourceInfo) GetDefaultSubtitleStreamIndexOk() (*int32, bool)`

GetDefaultSubtitleStreamIndexOk returns a tuple with the DefaultSubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubtitleStreamIndex

`func (o *ModelMediaSourceInfo) SetDefaultSubtitleStreamIndex(v int32)`

SetDefaultSubtitleStreamIndex sets DefaultSubtitleStreamIndex field to given value.

### HasDefaultSubtitleStreamIndex

`func (o *ModelMediaSourceInfo) HasDefaultSubtitleStreamIndex() bool`

HasDefaultSubtitleStreamIndex returns a boolean if a field has been set.

### SetDefaultSubtitleStreamIndexNil

`func (o *ModelMediaSourceInfo) SetDefaultSubtitleStreamIndexNil(b bool)`

 SetDefaultSubtitleStreamIndexNil sets the value for DefaultSubtitleStreamIndex to be an explicit nil

### UnsetDefaultSubtitleStreamIndex
`func (o *ModelMediaSourceInfo) UnsetDefaultSubtitleStreamIndex()`

UnsetDefaultSubtitleStreamIndex ensures that no value is present for DefaultSubtitleStreamIndex, not even an explicit nil
### GetItemId

`func (o *ModelMediaSourceInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelMediaSourceInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelMediaSourceInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelMediaSourceInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetServerId

`func (o *ModelMediaSourceInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelMediaSourceInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelMediaSourceInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelMediaSourceInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


