# ModelPlayerStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**CanSeek** | Pointer to **bool** |  | [optional] 
**IsPaused** | Pointer to **bool** |  | [optional] 
**IsMuted** | Pointer to **bool** |  | [optional] 
**VolumeLevel** | Pointer to **NullableInt32** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**PlayMethod** | Pointer to [**ModelModelPlayMethod**](ModelPlayMethod.md) |  | [optional] 
**RepeatMode** | Pointer to [**ModelModelRepeatMode**](ModelRepeatMode.md) |  | [optional] 
**SubtitleOffset** | Pointer to **int32** |  | [optional] 
**PlaybackRate** | Pointer to **float64** |  | [optional] 

## Methods

### NewModelPlayerStateInfo

`func NewModelPlayerStateInfo() *ModelPlayerStateInfo`

NewModelPlayerStateInfo instantiates a new ModelPlayerStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlayerStateInfoWithDefaults

`func NewModelPlayerStateInfoWithDefaults() *ModelPlayerStateInfo`

NewModelPlayerStateInfoWithDefaults instantiates a new ModelPlayerStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionTicks

`func (o *ModelPlayerStateInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *ModelPlayerStateInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *ModelPlayerStateInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *ModelPlayerStateInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *ModelPlayerStateInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *ModelPlayerStateInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetCanSeek

`func (o *ModelPlayerStateInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *ModelPlayerStateInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *ModelPlayerStateInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *ModelPlayerStateInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetIsPaused

`func (o *ModelPlayerStateInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ModelPlayerStateInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ModelPlayerStateInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ModelPlayerStateInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *ModelPlayerStateInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *ModelPlayerStateInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *ModelPlayerStateInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *ModelPlayerStateInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetVolumeLevel

`func (o *ModelPlayerStateInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *ModelPlayerStateInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *ModelPlayerStateInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *ModelPlayerStateInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *ModelPlayerStateInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *ModelPlayerStateInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetAudioStreamIndex

`func (o *ModelPlayerStateInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *ModelPlayerStateInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *ModelPlayerStateInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *ModelPlayerStateInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *ModelPlayerStateInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *ModelPlayerStateInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *ModelPlayerStateInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *ModelPlayerStateInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *ModelPlayerStateInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *ModelPlayerStateInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *ModelPlayerStateInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *ModelPlayerStateInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMediaSourceId

`func (o *ModelPlayerStateInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelPlayerStateInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelPlayerStateInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelPlayerStateInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetPlayMethod

`func (o *ModelPlayerStateInfo) GetPlayMethod() ModelModelPlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *ModelPlayerStateInfo) GetPlayMethodOk() (*ModelModelPlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *ModelPlayerStateInfo) SetPlayMethod(v ModelModelPlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *ModelPlayerStateInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetRepeatMode

`func (o *ModelPlayerStateInfo) GetRepeatMode() ModelModelRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *ModelPlayerStateInfo) GetRepeatModeOk() (*ModelModelRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *ModelPlayerStateInfo) SetRepeatMode(v ModelModelRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *ModelPlayerStateInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetSubtitleOffset

`func (o *ModelPlayerStateInfo) GetSubtitleOffset() int32`

GetSubtitleOffset returns the SubtitleOffset field if non-nil, zero value otherwise.

### GetSubtitleOffsetOk

`func (o *ModelPlayerStateInfo) GetSubtitleOffsetOk() (*int32, bool)`

GetSubtitleOffsetOk returns a tuple with the SubtitleOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleOffset

`func (o *ModelPlayerStateInfo) SetSubtitleOffset(v int32)`

SetSubtitleOffset sets SubtitleOffset field to given value.

### HasSubtitleOffset

`func (o *ModelPlayerStateInfo) HasSubtitleOffset() bool`

HasSubtitleOffset returns a boolean if a field has been set.

### GetPlaybackRate

`func (o *ModelPlayerStateInfo) GetPlaybackRate() float64`

GetPlaybackRate returns the PlaybackRate field if non-nil, zero value otherwise.

### GetPlaybackRateOk

`func (o *ModelPlayerStateInfo) GetPlaybackRateOk() (*float64, bool)`

GetPlaybackRateOk returns a tuple with the PlaybackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackRate

`func (o *ModelPlayerStateInfo) SetPlaybackRate(v float64)`

SetPlaybackRate sets PlaybackRate field to given value.

### HasPlaybackRate

`func (o *ModelPlayerStateInfo) HasPlaybackRate() bool`

HasPlaybackRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


