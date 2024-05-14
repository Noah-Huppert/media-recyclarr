# ModelPlaybackStartInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSeek** | Pointer to **bool** |  | [optional] 
**Item** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**NowPlayingQueue** | Pointer to [**[]ModelModelQueueItem**](ModelModelQueueItem.md) |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**IsPaused** | Pointer to **bool** |  | [optional] 
**PlaylistIndex** | Pointer to **int32** |  | [optional] 
**PlaylistLength** | Pointer to **int32** |  | [optional] 
**IsMuted** | Pointer to **bool** |  | [optional] 
**PositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**PlaybackStartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**VolumeLevel** | Pointer to **NullableInt32** |  | [optional] 
**Brightness** | Pointer to **NullableInt32** |  | [optional] 
**AspectRatio** | Pointer to **string** |  | [optional] 
**EventName** | Pointer to [**ModelModelProgressEvent**](ModelProgressEvent.md) |  | [optional] 
**PlayMethod** | Pointer to [**ModelModelPlayMethod**](ModelPlayMethod.md) |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**PlaySessionId** | Pointer to **string** |  | [optional] 
**RepeatMode** | Pointer to [**ModelModelRepeatMode**](ModelRepeatMode.md) |  | [optional] 
**SubtitleOffset** | Pointer to **int32** |  | [optional] 
**PlaybackRate** | Pointer to **float64** |  | [optional] 
**PlaylistItemIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelPlaybackStartInfo

`func NewModelPlaybackStartInfo() *ModelPlaybackStartInfo`

NewModelPlaybackStartInfo instantiates a new ModelPlaybackStartInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaybackStartInfoWithDefaults

`func NewModelPlaybackStartInfoWithDefaults() *ModelPlaybackStartInfo`

NewModelPlaybackStartInfoWithDefaults instantiates a new ModelPlaybackStartInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSeek

`func (o *ModelPlaybackStartInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *ModelPlaybackStartInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *ModelPlaybackStartInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *ModelPlaybackStartInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetItem

`func (o *ModelPlaybackStartInfo) GetItem() ModelModelBaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ModelPlaybackStartInfo) GetItemOk() (*ModelModelBaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ModelPlaybackStartInfo) SetItem(v ModelModelBaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *ModelPlaybackStartInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *ModelPlaybackStartInfo) GetNowPlayingQueue() []ModelModelQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *ModelPlaybackStartInfo) GetNowPlayingQueueOk() (*[]ModelModelQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *ModelPlaybackStartInfo) SetNowPlayingQueue(v []ModelModelQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *ModelPlaybackStartInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *ModelPlaybackStartInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelPlaybackStartInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelPlaybackStartInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelPlaybackStartInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetItemId

`func (o *ModelPlaybackStartInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelPlaybackStartInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelPlaybackStartInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelPlaybackStartInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *ModelPlaybackStartInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ModelPlaybackStartInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ModelPlaybackStartInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ModelPlaybackStartInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetMediaSourceId

`func (o *ModelPlaybackStartInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelPlaybackStartInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelPlaybackStartInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelPlaybackStartInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetAudioStreamIndex

`func (o *ModelPlaybackStartInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *ModelPlaybackStartInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *ModelPlaybackStartInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *ModelPlaybackStartInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *ModelPlaybackStartInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *ModelPlaybackStartInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *ModelPlaybackStartInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *ModelPlaybackStartInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *ModelPlaybackStartInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *ModelPlaybackStartInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *ModelPlaybackStartInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *ModelPlaybackStartInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetIsPaused

`func (o *ModelPlaybackStartInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *ModelPlaybackStartInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *ModelPlaybackStartInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *ModelPlaybackStartInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetPlaylistIndex

`func (o *ModelPlaybackStartInfo) GetPlaylistIndex() int32`

GetPlaylistIndex returns the PlaylistIndex field if non-nil, zero value otherwise.

### GetPlaylistIndexOk

`func (o *ModelPlaybackStartInfo) GetPlaylistIndexOk() (*int32, bool)`

GetPlaylistIndexOk returns a tuple with the PlaylistIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistIndex

`func (o *ModelPlaybackStartInfo) SetPlaylistIndex(v int32)`

SetPlaylistIndex sets PlaylistIndex field to given value.

### HasPlaylistIndex

`func (o *ModelPlaybackStartInfo) HasPlaylistIndex() bool`

HasPlaylistIndex returns a boolean if a field has been set.

### GetPlaylistLength

`func (o *ModelPlaybackStartInfo) GetPlaylistLength() int32`

GetPlaylistLength returns the PlaylistLength field if non-nil, zero value otherwise.

### GetPlaylistLengthOk

`func (o *ModelPlaybackStartInfo) GetPlaylistLengthOk() (*int32, bool)`

GetPlaylistLengthOk returns a tuple with the PlaylistLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistLength

`func (o *ModelPlaybackStartInfo) SetPlaylistLength(v int32)`

SetPlaylistLength sets PlaylistLength field to given value.

### HasPlaylistLength

`func (o *ModelPlaybackStartInfo) HasPlaylistLength() bool`

HasPlaylistLength returns a boolean if a field has been set.

### GetIsMuted

`func (o *ModelPlaybackStartInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *ModelPlaybackStartInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *ModelPlaybackStartInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *ModelPlaybackStartInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetPositionTicks

`func (o *ModelPlaybackStartInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *ModelPlaybackStartInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *ModelPlaybackStartInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *ModelPlaybackStartInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *ModelPlaybackStartInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *ModelPlaybackStartInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetRunTimeTicks

`func (o *ModelPlaybackStartInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *ModelPlaybackStartInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *ModelPlaybackStartInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *ModelPlaybackStartInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *ModelPlaybackStartInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *ModelPlaybackStartInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetPlaybackStartTimeTicks

`func (o *ModelPlaybackStartInfo) GetPlaybackStartTimeTicks() int64`

GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field if non-nil, zero value otherwise.

### GetPlaybackStartTimeTicksOk

`func (o *ModelPlaybackStartInfo) GetPlaybackStartTimeTicksOk() (*int64, bool)`

GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackStartTimeTicks

`func (o *ModelPlaybackStartInfo) SetPlaybackStartTimeTicks(v int64)`

SetPlaybackStartTimeTicks sets PlaybackStartTimeTicks field to given value.

### HasPlaybackStartTimeTicks

`func (o *ModelPlaybackStartInfo) HasPlaybackStartTimeTicks() bool`

HasPlaybackStartTimeTicks returns a boolean if a field has been set.

### SetPlaybackStartTimeTicksNil

`func (o *ModelPlaybackStartInfo) SetPlaybackStartTimeTicksNil(b bool)`

 SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil

### UnsetPlaybackStartTimeTicks
`func (o *ModelPlaybackStartInfo) UnsetPlaybackStartTimeTicks()`

UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
### GetVolumeLevel

`func (o *ModelPlaybackStartInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *ModelPlaybackStartInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *ModelPlaybackStartInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *ModelPlaybackStartInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *ModelPlaybackStartInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *ModelPlaybackStartInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetBrightness

`func (o *ModelPlaybackStartInfo) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *ModelPlaybackStartInfo) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *ModelPlaybackStartInfo) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *ModelPlaybackStartInfo) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### SetBrightnessNil

`func (o *ModelPlaybackStartInfo) SetBrightnessNil(b bool)`

 SetBrightnessNil sets the value for Brightness to be an explicit nil

### UnsetBrightness
`func (o *ModelPlaybackStartInfo) UnsetBrightness()`

UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
### GetAspectRatio

`func (o *ModelPlaybackStartInfo) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ModelPlaybackStartInfo) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ModelPlaybackStartInfo) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *ModelPlaybackStartInfo) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetEventName

`func (o *ModelPlaybackStartInfo) GetEventName() ModelModelProgressEvent`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ModelPlaybackStartInfo) GetEventNameOk() (*ModelModelProgressEvent, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ModelPlaybackStartInfo) SetEventName(v ModelModelProgressEvent)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ModelPlaybackStartInfo) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetPlayMethod

`func (o *ModelPlaybackStartInfo) GetPlayMethod() ModelModelPlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *ModelPlaybackStartInfo) GetPlayMethodOk() (*ModelModelPlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *ModelPlaybackStartInfo) SetPlayMethod(v ModelModelPlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *ModelPlaybackStartInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *ModelPlaybackStartInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *ModelPlaybackStartInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *ModelPlaybackStartInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *ModelPlaybackStartInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetPlaySessionId

`func (o *ModelPlaybackStartInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *ModelPlaybackStartInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *ModelPlaybackStartInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *ModelPlaybackStartInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### GetRepeatMode

`func (o *ModelPlaybackStartInfo) GetRepeatMode() ModelModelRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *ModelPlaybackStartInfo) GetRepeatModeOk() (*ModelModelRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *ModelPlaybackStartInfo) SetRepeatMode(v ModelModelRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *ModelPlaybackStartInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetSubtitleOffset

`func (o *ModelPlaybackStartInfo) GetSubtitleOffset() int32`

GetSubtitleOffset returns the SubtitleOffset field if non-nil, zero value otherwise.

### GetSubtitleOffsetOk

`func (o *ModelPlaybackStartInfo) GetSubtitleOffsetOk() (*int32, bool)`

GetSubtitleOffsetOk returns a tuple with the SubtitleOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleOffset

`func (o *ModelPlaybackStartInfo) SetSubtitleOffset(v int32)`

SetSubtitleOffset sets SubtitleOffset field to given value.

### HasSubtitleOffset

`func (o *ModelPlaybackStartInfo) HasSubtitleOffset() bool`

HasSubtitleOffset returns a boolean if a field has been set.

### GetPlaybackRate

`func (o *ModelPlaybackStartInfo) GetPlaybackRate() float64`

GetPlaybackRate returns the PlaybackRate field if non-nil, zero value otherwise.

### GetPlaybackRateOk

`func (o *ModelPlaybackStartInfo) GetPlaybackRateOk() (*float64, bool)`

GetPlaybackRateOk returns a tuple with the PlaybackRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackRate

`func (o *ModelPlaybackStartInfo) SetPlaybackRate(v float64)`

SetPlaybackRate sets PlaybackRate field to given value.

### HasPlaybackRate

`func (o *ModelPlaybackStartInfo) HasPlaybackRate() bool`

HasPlaybackRate returns a boolean if a field has been set.

### GetPlaylistItemIds

`func (o *ModelPlaybackStartInfo) GetPlaylistItemIds() []string`

GetPlaylistItemIds returns the PlaylistItemIds field if non-nil, zero value otherwise.

### GetPlaylistItemIdsOk

`func (o *ModelPlaybackStartInfo) GetPlaylistItemIdsOk() (*[]string, bool)`

GetPlaylistItemIdsOk returns a tuple with the PlaylistItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemIds

`func (o *ModelPlaybackStartInfo) SetPlaylistItemIds(v []string)`

SetPlaylistItemIds sets PlaylistItemIds field to given value.

### HasPlaylistItemIds

`func (o *ModelPlaybackStartInfo) HasPlaylistItemIds() bool`

HasPlaylistItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


