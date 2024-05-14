# ModelPlaybackStopInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NowPlayingQueue** | Pointer to [**[]ModelModelQueueItem**](ModelModelQueueItem.md) |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**PlaylistIndex** | Pointer to **int32** |  | [optional] 
**PlaylistLength** | Pointer to **int32** |  | [optional] 
**Item** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**PositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**LiveStreamId** | Pointer to **string** |  | [optional] 
**PlaySessionId** | Pointer to **string** |  | [optional] 
**Failed** | Pointer to **bool** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**NextMediaType** | Pointer to **string** |  | [optional] 

## Methods

### NewModelPlaybackStopInfo

`func NewModelPlaybackStopInfo() *ModelPlaybackStopInfo`

NewModelPlaybackStopInfo instantiates a new ModelPlaybackStopInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPlaybackStopInfoWithDefaults

`func NewModelPlaybackStopInfoWithDefaults() *ModelPlaybackStopInfo`

NewModelPlaybackStopInfoWithDefaults instantiates a new ModelPlaybackStopInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNowPlayingQueue

`func (o *ModelPlaybackStopInfo) GetNowPlayingQueue() []ModelModelQueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *ModelPlaybackStopInfo) GetNowPlayingQueueOk() (*[]ModelModelQueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *ModelPlaybackStopInfo) SetNowPlayingQueue(v []ModelModelQueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *ModelPlaybackStopInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *ModelPlaybackStopInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelPlaybackStopInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelPlaybackStopInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelPlaybackStopInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetPlaylistIndex

`func (o *ModelPlaybackStopInfo) GetPlaylistIndex() int32`

GetPlaylistIndex returns the PlaylistIndex field if non-nil, zero value otherwise.

### GetPlaylistIndexOk

`func (o *ModelPlaybackStopInfo) GetPlaylistIndexOk() (*int32, bool)`

GetPlaylistIndexOk returns a tuple with the PlaylistIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistIndex

`func (o *ModelPlaybackStopInfo) SetPlaylistIndex(v int32)`

SetPlaylistIndex sets PlaylistIndex field to given value.

### HasPlaylistIndex

`func (o *ModelPlaybackStopInfo) HasPlaylistIndex() bool`

HasPlaylistIndex returns a boolean if a field has been set.

### GetPlaylistLength

`func (o *ModelPlaybackStopInfo) GetPlaylistLength() int32`

GetPlaylistLength returns the PlaylistLength field if non-nil, zero value otherwise.

### GetPlaylistLengthOk

`func (o *ModelPlaybackStopInfo) GetPlaylistLengthOk() (*int32, bool)`

GetPlaylistLengthOk returns a tuple with the PlaylistLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistLength

`func (o *ModelPlaybackStopInfo) SetPlaylistLength(v int32)`

SetPlaylistLength sets PlaylistLength field to given value.

### HasPlaylistLength

`func (o *ModelPlaybackStopInfo) HasPlaylistLength() bool`

HasPlaylistLength returns a boolean if a field has been set.

### GetItem

`func (o *ModelPlaybackStopInfo) GetItem() ModelModelBaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ModelPlaybackStopInfo) GetItemOk() (*ModelModelBaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ModelPlaybackStopInfo) SetItem(v ModelModelBaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *ModelPlaybackStopInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetItemId

`func (o *ModelPlaybackStopInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelPlaybackStopInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelPlaybackStopInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelPlaybackStopInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *ModelPlaybackStopInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ModelPlaybackStopInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ModelPlaybackStopInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *ModelPlaybackStopInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetMediaSourceId

`func (o *ModelPlaybackStopInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelPlaybackStopInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelPlaybackStopInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelPlaybackStopInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetPositionTicks

`func (o *ModelPlaybackStopInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *ModelPlaybackStopInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *ModelPlaybackStopInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *ModelPlaybackStopInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *ModelPlaybackStopInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *ModelPlaybackStopInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetLiveStreamId

`func (o *ModelPlaybackStopInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *ModelPlaybackStopInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *ModelPlaybackStopInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *ModelPlaybackStopInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### GetPlaySessionId

`func (o *ModelPlaybackStopInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *ModelPlaybackStopInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *ModelPlaybackStopInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *ModelPlaybackStopInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### GetFailed

`func (o *ModelPlaybackStopInfo) GetFailed() bool`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ModelPlaybackStopInfo) GetFailedOk() (*bool, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ModelPlaybackStopInfo) SetFailed(v bool)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ModelPlaybackStopInfo) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetIsAutomated

`func (o *ModelPlaybackStopInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *ModelPlaybackStopInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *ModelPlaybackStopInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *ModelPlaybackStopInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetNextMediaType

`func (o *ModelPlaybackStopInfo) GetNextMediaType() string`

GetNextMediaType returns the NextMediaType field if non-nil, zero value otherwise.

### GetNextMediaTypeOk

`func (o *ModelPlaybackStopInfo) GetNextMediaTypeOk() (*string, bool)`

GetNextMediaTypeOk returns a tuple with the NextMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMediaType

`func (o *ModelPlaybackStopInfo) SetNextMediaType(v string)`

SetNextMediaType sets NextMediaType field to given value.

### HasNextMediaType

`func (o *ModelPlaybackStopInfo) HasNextMediaType() bool`

HasNextMediaType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


