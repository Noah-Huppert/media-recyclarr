# ModelApiOnPlaybackProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistIndex** | Pointer to **int32** |  | [optional] 
**PlaylistLength** | Pointer to **int32** |  | [optional] 
**EventName** | Pointer to [**ModelModelProgressEvent**](ModelProgressEvent.md) |  | [optional] 

## Methods

### NewModelApiOnPlaybackProgress

`func NewModelApiOnPlaybackProgress() *ModelApiOnPlaybackProgress`

NewModelApiOnPlaybackProgress instantiates a new ModelApiOnPlaybackProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelApiOnPlaybackProgressWithDefaults

`func NewModelApiOnPlaybackProgressWithDefaults() *ModelApiOnPlaybackProgress`

NewModelApiOnPlaybackProgressWithDefaults instantiates a new ModelApiOnPlaybackProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistIndex

`func (o *ModelApiOnPlaybackProgress) GetPlaylistIndex() int32`

GetPlaylistIndex returns the PlaylistIndex field if non-nil, zero value otherwise.

### GetPlaylistIndexOk

`func (o *ModelApiOnPlaybackProgress) GetPlaylistIndexOk() (*int32, bool)`

GetPlaylistIndexOk returns a tuple with the PlaylistIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistIndex

`func (o *ModelApiOnPlaybackProgress) SetPlaylistIndex(v int32)`

SetPlaylistIndex sets PlaylistIndex field to given value.

### HasPlaylistIndex

`func (o *ModelApiOnPlaybackProgress) HasPlaylistIndex() bool`

HasPlaylistIndex returns a boolean if a field has been set.

### GetPlaylistLength

`func (o *ModelApiOnPlaybackProgress) GetPlaylistLength() int32`

GetPlaylistLength returns the PlaylistLength field if non-nil, zero value otherwise.

### GetPlaylistLengthOk

`func (o *ModelApiOnPlaybackProgress) GetPlaylistLengthOk() (*int32, bool)`

GetPlaylistLengthOk returns a tuple with the PlaylistLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistLength

`func (o *ModelApiOnPlaybackProgress) SetPlaylistLength(v int32)`

SetPlaylistLength sets PlaylistLength field to given value.

### HasPlaylistLength

`func (o *ModelApiOnPlaybackProgress) HasPlaylistLength() bool`

HasPlaylistLength returns a boolean if a field has been set.

### GetEventName

`func (o *ModelApiOnPlaybackProgress) GetEventName() ModelModelProgressEvent`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ModelApiOnPlaybackProgress) GetEventNameOk() (*ModelModelProgressEvent, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ModelApiOnPlaybackProgress) SetEventName(v ModelModelProgressEvent)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ModelApiOnPlaybackProgress) HasEventName() bool`

HasEventName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


