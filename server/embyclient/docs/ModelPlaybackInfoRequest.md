# ModelPlaybackInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt64** |  | [optional] 
**StartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**MaxAudioChannels** | Pointer to **NullableInt32** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**DeviceProfile** | Pointer to [**ModelModelDeviceProfile**](ModelDeviceProfile.md) |  | [optional] 
**EnableDirectPlay** | Pointer to **bool** |  | [optional] 
**EnableDirectStream** | Pointer to **bool** |  | [optional] 
**EnableTranscoding** | Pointer to **bool** |  | [optional] 
**AllowInterlacedVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**AllowVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**AllowAudioStreamCopy** | Pointer to **bool** |  | [optional] 
**IsPlayback** | Pointer to **bool** |  | [optional] 
**AutoOpenLiveStream** | Pointer to **bool** |  | [optional] 
**DirectPlayProtocols** | Pointer to [**[]ModelModelMediaProtocol**](ModelModelMediaProtocol.md) |  | [optional] 
**CurrentPlaySessionId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelPlaybackInfoRequest

`func NewModelPlaybackInfoRequest() *ModelPlaybackInfoRequest`

NewModelPlaybackInfoRequest instantiates a new ModelPlaybackInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaybackInfoRequestWithDefaults

`func NewModelPlaybackInfoRequestWithDefaults() *ModelPlaybackInfoRequest`

NewModelPlaybackInfoRequestWithDefaults instantiates a new ModelPlaybackInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelPlaybackInfoRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelPlaybackInfoRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelPlaybackInfoRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelPlaybackInfoRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *ModelPlaybackInfoRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelPlaybackInfoRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelPlaybackInfoRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelPlaybackInfoRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetMaxStreamingBitrate

`func (o *ModelPlaybackInfoRequest) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *ModelPlaybackInfoRequest) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *ModelPlaybackInfoRequest) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *ModelPlaybackInfoRequest) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *ModelPlaybackInfoRequest) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *ModelPlaybackInfoRequest) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *ModelPlaybackInfoRequest) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *ModelPlaybackInfoRequest) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *ModelPlaybackInfoRequest) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *ModelPlaybackInfoRequest) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *ModelPlaybackInfoRequest) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *ModelPlaybackInfoRequest) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *ModelPlaybackInfoRequest) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *ModelPlaybackInfoRequest) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *ModelPlaybackInfoRequest) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *ModelPlaybackInfoRequest) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *ModelPlaybackInfoRequest) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *ModelPlaybackInfoRequest) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *ModelPlaybackInfoRequest) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *ModelPlaybackInfoRequest) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *ModelPlaybackInfoRequest) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *ModelPlaybackInfoRequest) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *ModelPlaybackInfoRequest) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *ModelPlaybackInfoRequest) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *ModelPlaybackInfoRequest) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *ModelPlaybackInfoRequest) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *ModelPlaybackInfoRequest) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *ModelPlaybackInfoRequest) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *ModelPlaybackInfoRequest) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *ModelPlaybackInfoRequest) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetMediaSourceId

`func (o *ModelPlaybackInfoRequest) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelPlaybackInfoRequest) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelPlaybackInfoRequest) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelPlaybackInfoRequest) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *ModelPlaybackInfoRequest) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *ModelPlaybackInfoRequest) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *ModelPlaybackInfoRequest) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *ModelPlaybackInfoRequest) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *ModelPlaybackInfoRequest) GetDeviceProfile() ModelModelDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ModelPlaybackInfoRequest) GetDeviceProfileOk() (*ModelModelDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ModelPlaybackInfoRequest) SetDeviceProfile(v ModelModelDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ModelPlaybackInfoRequest) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetEnableDirectPlay

`func (o *ModelPlaybackInfoRequest) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *ModelPlaybackInfoRequest) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *ModelPlaybackInfoRequest) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *ModelPlaybackInfoRequest) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### GetEnableDirectStream

`func (o *ModelPlaybackInfoRequest) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *ModelPlaybackInfoRequest) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *ModelPlaybackInfoRequest) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *ModelPlaybackInfoRequest) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### GetEnableTranscoding

`func (o *ModelPlaybackInfoRequest) GetEnableTranscoding() bool`

GetEnableTranscoding returns the EnableTranscoding field if non-nil, zero value otherwise.

### GetEnableTranscodingOk

`func (o *ModelPlaybackInfoRequest) GetEnableTranscodingOk() (*bool, bool)`

GetEnableTranscodingOk returns a tuple with the EnableTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTranscoding

`func (o *ModelPlaybackInfoRequest) SetEnableTranscoding(v bool)`

SetEnableTranscoding sets EnableTranscoding field to given value.

### HasEnableTranscoding

`func (o *ModelPlaybackInfoRequest) HasEnableTranscoding() bool`

HasEnableTranscoding returns a boolean if a field has been set.

### GetAllowInterlacedVideoStreamCopy

`func (o *ModelPlaybackInfoRequest) GetAllowInterlacedVideoStreamCopy() bool`

GetAllowInterlacedVideoStreamCopy returns the AllowInterlacedVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowInterlacedVideoStreamCopyOk

`func (o *ModelPlaybackInfoRequest) GetAllowInterlacedVideoStreamCopyOk() (*bool, bool)`

GetAllowInterlacedVideoStreamCopyOk returns a tuple with the AllowInterlacedVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInterlacedVideoStreamCopy

`func (o *ModelPlaybackInfoRequest) SetAllowInterlacedVideoStreamCopy(v bool)`

SetAllowInterlacedVideoStreamCopy sets AllowInterlacedVideoStreamCopy field to given value.

### HasAllowInterlacedVideoStreamCopy

`func (o *ModelPlaybackInfoRequest) HasAllowInterlacedVideoStreamCopy() bool`

HasAllowInterlacedVideoStreamCopy returns a boolean if a field has been set.

### GetAllowVideoStreamCopy

`func (o *ModelPlaybackInfoRequest) GetAllowVideoStreamCopy() bool`

GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowVideoStreamCopyOk

`func (o *ModelPlaybackInfoRequest) GetAllowVideoStreamCopyOk() (*bool, bool)`

GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVideoStreamCopy

`func (o *ModelPlaybackInfoRequest) SetAllowVideoStreamCopy(v bool)`

SetAllowVideoStreamCopy sets AllowVideoStreamCopy field to given value.

### HasAllowVideoStreamCopy

`func (o *ModelPlaybackInfoRequest) HasAllowVideoStreamCopy() bool`

HasAllowVideoStreamCopy returns a boolean if a field has been set.

### GetAllowAudioStreamCopy

`func (o *ModelPlaybackInfoRequest) GetAllowAudioStreamCopy() bool`

GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field if non-nil, zero value otherwise.

### GetAllowAudioStreamCopyOk

`func (o *ModelPlaybackInfoRequest) GetAllowAudioStreamCopyOk() (*bool, bool)`

GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAudioStreamCopy

`func (o *ModelPlaybackInfoRequest) SetAllowAudioStreamCopy(v bool)`

SetAllowAudioStreamCopy sets AllowAudioStreamCopy field to given value.

### HasAllowAudioStreamCopy

`func (o *ModelPlaybackInfoRequest) HasAllowAudioStreamCopy() bool`

HasAllowAudioStreamCopy returns a boolean if a field has been set.

### GetIsPlayback

`func (o *ModelPlaybackInfoRequest) GetIsPlayback() bool`

GetIsPlayback returns the IsPlayback field if non-nil, zero value otherwise.

### GetIsPlaybackOk

`func (o *ModelPlaybackInfoRequest) GetIsPlaybackOk() (*bool, bool)`

GetIsPlaybackOk returns a tuple with the IsPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlayback

`func (o *ModelPlaybackInfoRequest) SetIsPlayback(v bool)`

SetIsPlayback sets IsPlayback field to given value.

### HasIsPlayback

`func (o *ModelPlaybackInfoRequest) HasIsPlayback() bool`

HasIsPlayback returns a boolean if a field has been set.

### GetAutoOpenLiveStream

`func (o *ModelPlaybackInfoRequest) GetAutoOpenLiveStream() bool`

GetAutoOpenLiveStream returns the AutoOpenLiveStream field if non-nil, zero value otherwise.

### GetAutoOpenLiveStreamOk

`func (o *ModelPlaybackInfoRequest) GetAutoOpenLiveStreamOk() (*bool, bool)`

GetAutoOpenLiveStreamOk returns a tuple with the AutoOpenLiveStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoOpenLiveStream

`func (o *ModelPlaybackInfoRequest) SetAutoOpenLiveStream(v bool)`

SetAutoOpenLiveStream sets AutoOpenLiveStream field to given value.

### HasAutoOpenLiveStream

`func (o *ModelPlaybackInfoRequest) HasAutoOpenLiveStream() bool`

HasAutoOpenLiveStream returns a boolean if a field has been set.

### GetDirectPlayProtocols

`func (o *ModelPlaybackInfoRequest) GetDirectPlayProtocols() []ModelModelMediaProtocol`

GetDirectPlayProtocols returns the DirectPlayProtocols field if non-nil, zero value otherwise.

### GetDirectPlayProtocolsOk

`func (o *ModelPlaybackInfoRequest) GetDirectPlayProtocolsOk() (*[]ModelModelMediaProtocol, bool)`

GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProtocols

`func (o *ModelPlaybackInfoRequest) SetDirectPlayProtocols(v []ModelModelMediaProtocol)`

SetDirectPlayProtocols sets DirectPlayProtocols field to given value.

### HasDirectPlayProtocols

`func (o *ModelPlaybackInfoRequest) HasDirectPlayProtocols() bool`

HasDirectPlayProtocols returns a boolean if a field has been set.

### GetCurrentPlaySessionId

`func (o *ModelPlaybackInfoRequest) GetCurrentPlaySessionId() string`

GetCurrentPlaySessionId returns the CurrentPlaySessionId field if non-nil, zero value otherwise.

### GetCurrentPlaySessionIdOk

`func (o *ModelPlaybackInfoRequest) GetCurrentPlaySessionIdOk() (*string, bool)`

GetCurrentPlaySessionIdOk returns a tuple with the CurrentPlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPlaySessionId

`func (o *ModelPlaybackInfoRequest) SetCurrentPlaySessionId(v string)`

SetCurrentPlaySessionId sets CurrentPlaySessionId field to given value.

### HasCurrentPlaySessionId

`func (o *ModelPlaybackInfoRequest) HasCurrentPlaySessionId() bool`

HasCurrentPlaySessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


