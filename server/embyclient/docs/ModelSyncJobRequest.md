# ModelSyncJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | Pointer to **string** |  | [optional] 
**ItemIds** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to [**ModelModelSyncCategory**](ModelSyncCategory.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UnwatchedOnly** | Pointer to **bool** |  | [optional] 
**SyncNewContent** | Pointer to **bool** |  | [optional] 
**ItemLimit** | Pointer to **NullableInt32** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Downloaded** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelSyncJobRequest

`func NewModelSyncJobRequest() *ModelSyncJobRequest`

NewModelSyncJobRequest instantiates a new ModelSyncJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncJobRequestWithDefaults

`func NewModelSyncJobRequestWithDefaults() *ModelSyncJobRequest`

NewModelSyncJobRequestWithDefaults instantiates a new ModelSyncJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetId

`func (o *ModelSyncJobRequest) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ModelSyncJobRequest) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ModelSyncJobRequest) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ModelSyncJobRequest) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetItemIds

`func (o *ModelSyncJobRequest) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *ModelSyncJobRequest) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *ModelSyncJobRequest) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *ModelSyncJobRequest) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetCategory

`func (o *ModelSyncJobRequest) GetCategory() ModelModelSyncCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelSyncJobRequest) GetCategoryOk() (*ModelModelSyncCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelSyncJobRequest) SetCategory(v ModelModelSyncCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelSyncJobRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetParentId

`func (o *ModelSyncJobRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModelSyncJobRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModelSyncJobRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModelSyncJobRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetQuality

`func (o *ModelSyncJobRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ModelSyncJobRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ModelSyncJobRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ModelSyncJobRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetProfile

`func (o *ModelSyncJobRequest) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ModelSyncJobRequest) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ModelSyncJobRequest) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ModelSyncJobRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetContainer

`func (o *ModelSyncJobRequest) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelSyncJobRequest) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelSyncJobRequest) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelSyncJobRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetVideoCodec

`func (o *ModelSyncJobRequest) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelSyncJobRequest) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelSyncJobRequest) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelSyncJobRequest) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *ModelSyncJobRequest) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelSyncJobRequest) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelSyncJobRequest) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelSyncJobRequest) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetName

`func (o *ModelSyncJobRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSyncJobRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSyncJobRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSyncJobRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *ModelSyncJobRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelSyncJobRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelSyncJobRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelSyncJobRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUnwatchedOnly

`func (o *ModelSyncJobRequest) GetUnwatchedOnly() bool`

GetUnwatchedOnly returns the UnwatchedOnly field if non-nil, zero value otherwise.

### GetUnwatchedOnlyOk

`func (o *ModelSyncJobRequest) GetUnwatchedOnlyOk() (*bool, bool)`

GetUnwatchedOnlyOk returns a tuple with the UnwatchedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwatchedOnly

`func (o *ModelSyncJobRequest) SetUnwatchedOnly(v bool)`

SetUnwatchedOnly sets UnwatchedOnly field to given value.

### HasUnwatchedOnly

`func (o *ModelSyncJobRequest) HasUnwatchedOnly() bool`

HasUnwatchedOnly returns a boolean if a field has been set.

### GetSyncNewContent

`func (o *ModelSyncJobRequest) GetSyncNewContent() bool`

GetSyncNewContent returns the SyncNewContent field if non-nil, zero value otherwise.

### GetSyncNewContentOk

`func (o *ModelSyncJobRequest) GetSyncNewContentOk() (*bool, bool)`

GetSyncNewContentOk returns a tuple with the SyncNewContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncNewContent

`func (o *ModelSyncJobRequest) SetSyncNewContent(v bool)`

SetSyncNewContent sets SyncNewContent field to given value.

### HasSyncNewContent

`func (o *ModelSyncJobRequest) HasSyncNewContent() bool`

HasSyncNewContent returns a boolean if a field has been set.

### GetItemLimit

`func (o *ModelSyncJobRequest) GetItemLimit() int32`

GetItemLimit returns the ItemLimit field if non-nil, zero value otherwise.

### GetItemLimitOk

`func (o *ModelSyncJobRequest) GetItemLimitOk() (*int32, bool)`

GetItemLimitOk returns a tuple with the ItemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLimit

`func (o *ModelSyncJobRequest) SetItemLimit(v int32)`

SetItemLimit sets ItemLimit field to given value.

### HasItemLimit

`func (o *ModelSyncJobRequest) HasItemLimit() bool`

HasItemLimit returns a boolean if a field has been set.

### SetItemLimitNil

`func (o *ModelSyncJobRequest) SetItemLimitNil(b bool)`

 SetItemLimitNil sets the value for ItemLimit to be an explicit nil

### UnsetItemLimit
`func (o *ModelSyncJobRequest) UnsetItemLimit()`

UnsetItemLimit ensures that no value is present for ItemLimit, not even an explicit nil
### GetBitrate

`func (o *ModelSyncJobRequest) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelSyncJobRequest) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelSyncJobRequest) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelSyncJobRequest) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelSyncJobRequest) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelSyncJobRequest) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetDownloaded

`func (o *ModelSyncJobRequest) GetDownloaded() bool`

GetDownloaded returns the Downloaded field if non-nil, zero value otherwise.

### GetDownloadedOk

`func (o *ModelSyncJobRequest) GetDownloadedOk() (*bool, bool)`

GetDownloadedOk returns a tuple with the Downloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloaded

`func (o *ModelSyncJobRequest) SetDownloaded(v bool)`

SetDownloaded sets Downloaded field to given value.

### HasDownloaded

`func (o *ModelSyncJobRequest) HasDownloaded() bool`

HasDownloaded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


