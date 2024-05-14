# ModelSyncJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**InternalTargetId** | Pointer to **int64** |  | [optional] 
**TargetName** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Category** | Pointer to [**ModelModelSyncCategory**](ModelSyncCategory.md) |  | [optional] 
**ParentId** | Pointer to **int64** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelModelSyncJobStatus**](ModelSyncJobStatus.md) |  | [optional] 
**UserId** | Pointer to **int64** |  | [optional] 
**UnwatchedOnly** | Pointer to **bool** |  | [optional] 
**SyncNewContent** | Pointer to **bool** |  | [optional] 
**ItemLimit** | Pointer to **NullableInt32** |  | [optional] 
**RequestedItemIds** | Pointer to **[]int64** |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**DateLastModified** | Pointer to **time.Time** |  | [optional] 
**ItemCount** | Pointer to **int32** |  | [optional] 
**ParentName** | Pointer to **string** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 

## Methods

### NewModelSyncJob

`func NewModelSyncJob() *ModelSyncJob`

NewModelSyncJob instantiates a new ModelSyncJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncJobWithDefaults

`func NewModelSyncJobWithDefaults() *ModelSyncJob`

NewModelSyncJobWithDefaults instantiates a new ModelSyncJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelSyncJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSyncJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSyncJob) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSyncJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetId

`func (o *ModelSyncJob) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ModelSyncJob) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ModelSyncJob) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ModelSyncJob) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetInternalTargetId

`func (o *ModelSyncJob) GetInternalTargetId() int64`

GetInternalTargetId returns the InternalTargetId field if non-nil, zero value otherwise.

### GetInternalTargetIdOk

`func (o *ModelSyncJob) GetInternalTargetIdOk() (*int64, bool)`

GetInternalTargetIdOk returns a tuple with the InternalTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTargetId

`func (o *ModelSyncJob) SetInternalTargetId(v int64)`

SetInternalTargetId sets InternalTargetId field to given value.

### HasInternalTargetId

`func (o *ModelSyncJob) HasInternalTargetId() bool`

HasInternalTargetId returns a boolean if a field has been set.

### GetTargetName

`func (o *ModelSyncJob) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ModelSyncJob) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ModelSyncJob) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ModelSyncJob) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetQuality

`func (o *ModelSyncJob) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *ModelSyncJob) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *ModelSyncJob) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *ModelSyncJob) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetBitrate

`func (o *ModelSyncJob) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelSyncJob) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelSyncJob) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelSyncJob) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelSyncJob) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelSyncJob) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetContainer

`func (o *ModelSyncJob) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelSyncJob) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelSyncJob) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelSyncJob) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetVideoCodec

`func (o *ModelSyncJob) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelSyncJob) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelSyncJob) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelSyncJob) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *ModelSyncJob) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelSyncJob) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelSyncJob) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelSyncJob) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetProfile

`func (o *ModelSyncJob) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ModelSyncJob) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ModelSyncJob) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ModelSyncJob) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetCategory

`func (o *ModelSyncJob) GetCategory() ModelModelSyncCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelSyncJob) GetCategoryOk() (*ModelModelSyncCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelSyncJob) SetCategory(v ModelModelSyncCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelSyncJob) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetParentId

`func (o *ModelSyncJob) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModelSyncJob) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModelSyncJob) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModelSyncJob) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProgress

`func (o *ModelSyncJob) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ModelSyncJob) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ModelSyncJob) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ModelSyncJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetName

`func (o *ModelSyncJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSyncJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSyncJob) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSyncJob) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ModelSyncJob) GetStatus() ModelModelSyncJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelSyncJob) GetStatusOk() (*ModelModelSyncJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelSyncJob) SetStatus(v ModelModelSyncJobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelSyncJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserId

`func (o *ModelSyncJob) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelSyncJob) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelSyncJob) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelSyncJob) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUnwatchedOnly

`func (o *ModelSyncJob) GetUnwatchedOnly() bool`

GetUnwatchedOnly returns the UnwatchedOnly field if non-nil, zero value otherwise.

### GetUnwatchedOnlyOk

`func (o *ModelSyncJob) GetUnwatchedOnlyOk() (*bool, bool)`

GetUnwatchedOnlyOk returns a tuple with the UnwatchedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnwatchedOnly

`func (o *ModelSyncJob) SetUnwatchedOnly(v bool)`

SetUnwatchedOnly sets UnwatchedOnly field to given value.

### HasUnwatchedOnly

`func (o *ModelSyncJob) HasUnwatchedOnly() bool`

HasUnwatchedOnly returns a boolean if a field has been set.

### GetSyncNewContent

`func (o *ModelSyncJob) GetSyncNewContent() bool`

GetSyncNewContent returns the SyncNewContent field if non-nil, zero value otherwise.

### GetSyncNewContentOk

`func (o *ModelSyncJob) GetSyncNewContentOk() (*bool, bool)`

GetSyncNewContentOk returns a tuple with the SyncNewContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncNewContent

`func (o *ModelSyncJob) SetSyncNewContent(v bool)`

SetSyncNewContent sets SyncNewContent field to given value.

### HasSyncNewContent

`func (o *ModelSyncJob) HasSyncNewContent() bool`

HasSyncNewContent returns a boolean if a field has been set.

### GetItemLimit

`func (o *ModelSyncJob) GetItemLimit() int32`

GetItemLimit returns the ItemLimit field if non-nil, zero value otherwise.

### GetItemLimitOk

`func (o *ModelSyncJob) GetItemLimitOk() (*int32, bool)`

GetItemLimitOk returns a tuple with the ItemLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLimit

`func (o *ModelSyncJob) SetItemLimit(v int32)`

SetItemLimit sets ItemLimit field to given value.

### HasItemLimit

`func (o *ModelSyncJob) HasItemLimit() bool`

HasItemLimit returns a boolean if a field has been set.

### SetItemLimitNil

`func (o *ModelSyncJob) SetItemLimitNil(b bool)`

 SetItemLimitNil sets the value for ItemLimit to be an explicit nil

### UnsetItemLimit
`func (o *ModelSyncJob) UnsetItemLimit()`

UnsetItemLimit ensures that no value is present for ItemLimit, not even an explicit nil
### GetRequestedItemIds

`func (o *ModelSyncJob) GetRequestedItemIds() []int64`

GetRequestedItemIds returns the RequestedItemIds field if non-nil, zero value otherwise.

### GetRequestedItemIdsOk

`func (o *ModelSyncJob) GetRequestedItemIdsOk() (*[]int64, bool)`

GetRequestedItemIdsOk returns a tuple with the RequestedItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedItemIds

`func (o *ModelSyncJob) SetRequestedItemIds(v []int64)`

SetRequestedItemIds sets RequestedItemIds field to given value.

### HasRequestedItemIds

`func (o *ModelSyncJob) HasRequestedItemIds() bool`

HasRequestedItemIds returns a boolean if a field has been set.

### GetItemId

`func (o *ModelSyncJob) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelSyncJob) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelSyncJob) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelSyncJob) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ModelSyncJob) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelSyncJob) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelSyncJob) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelSyncJob) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateLastModified

`func (o *ModelSyncJob) GetDateLastModified() time.Time`

GetDateLastModified returns the DateLastModified field if non-nil, zero value otherwise.

### GetDateLastModifiedOk

`func (o *ModelSyncJob) GetDateLastModifiedOk() (*time.Time, bool)`

GetDateLastModifiedOk returns a tuple with the DateLastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastModified

`func (o *ModelSyncJob) SetDateLastModified(v time.Time)`

SetDateLastModified sets DateLastModified field to given value.

### HasDateLastModified

`func (o *ModelSyncJob) HasDateLastModified() bool`

HasDateLastModified returns a boolean if a field has been set.

### GetItemCount

`func (o *ModelSyncJob) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ModelSyncJob) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ModelSyncJob) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *ModelSyncJob) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetParentName

`func (o *ModelSyncJob) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *ModelSyncJob) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *ModelSyncJob) SetParentName(v string)`

SetParentName sets ParentName field to given value.

### HasParentName

`func (o *ModelSyncJob) HasParentName() bool`

HasParentName returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *ModelSyncJob) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *ModelSyncJob) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *ModelSyncJob) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *ModelSyncJob) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *ModelSyncJob) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *ModelSyncJob) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *ModelSyncJob) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *ModelSyncJob) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


