# ModelSyncJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**JobId** | Pointer to **int64** |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**ItemName** | Pointer to **string** |  | [optional] 
**MediaSourceId** | Pointer to **string** |  | [optional] 
**MediaSource** | Pointer to [**ModelModelMediaSourceInfo**](ModelMediaSourceInfo.md) |  | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**InternalTargetId** | Pointer to **int64** |  | [optional] 
**OutputPath** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ModelModelSyncJobItemStatus**](ModelSyncJobItemStatus.md) |  | [optional] 
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**TemporaryPath** | Pointer to **string** |  | [optional] 
**AdditionalFiles** | Pointer to [**[]ModelModelItemFileInfo**](ModelModelItemFileInfo.md) |  | [optional] 

## Methods

### NewModelSyncJobItem

`func NewModelSyncJobItem() *ModelSyncJobItem`

NewModelSyncJobItem instantiates a new ModelSyncJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncJobItemWithDefaults

`func NewModelSyncJobItemWithDefaults() *ModelSyncJobItem`

NewModelSyncJobItemWithDefaults instantiates a new ModelSyncJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelSyncJobItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSyncJobItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSyncJobItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSyncJobItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *ModelSyncJobItem) GetJobId() int64`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ModelSyncJobItem) GetJobIdOk() (*int64, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ModelSyncJobItem) SetJobId(v int64)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ModelSyncJobItem) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetItemId

`func (o *ModelSyncJobItem) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelSyncJobItem) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelSyncJobItem) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelSyncJobItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetItemName

`func (o *ModelSyncJobItem) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *ModelSyncJobItem) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *ModelSyncJobItem) SetItemName(v string)`

SetItemName sets ItemName field to given value.

### HasItemName

`func (o *ModelSyncJobItem) HasItemName() bool`

HasItemName returns a boolean if a field has been set.

### GetMediaSourceId

`func (o *ModelSyncJobItem) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *ModelSyncJobItem) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *ModelSyncJobItem) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *ModelSyncJobItem) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### GetMediaSource

`func (o *ModelSyncJobItem) GetMediaSource() ModelModelMediaSourceInfo`

GetMediaSource returns the MediaSource field if non-nil, zero value otherwise.

### GetMediaSourceOk

`func (o *ModelSyncJobItem) GetMediaSourceOk() (*ModelModelMediaSourceInfo, bool)`

GetMediaSourceOk returns a tuple with the MediaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSource

`func (o *ModelSyncJobItem) SetMediaSource(v ModelModelMediaSourceInfo)`

SetMediaSource sets MediaSource field to given value.

### HasMediaSource

`func (o *ModelSyncJobItem) HasMediaSource() bool`

HasMediaSource returns a boolean if a field has been set.

### GetTargetId

`func (o *ModelSyncJobItem) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ModelSyncJobItem) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ModelSyncJobItem) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ModelSyncJobItem) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetInternalTargetId

`func (o *ModelSyncJobItem) GetInternalTargetId() int64`

GetInternalTargetId returns the InternalTargetId field if non-nil, zero value otherwise.

### GetInternalTargetIdOk

`func (o *ModelSyncJobItem) GetInternalTargetIdOk() (*int64, bool)`

GetInternalTargetIdOk returns a tuple with the InternalTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalTargetId

`func (o *ModelSyncJobItem) SetInternalTargetId(v int64)`

SetInternalTargetId sets InternalTargetId field to given value.

### HasInternalTargetId

`func (o *ModelSyncJobItem) HasInternalTargetId() bool`

HasInternalTargetId returns a boolean if a field has been set.

### GetOutputPath

`func (o *ModelSyncJobItem) GetOutputPath() string`

GetOutputPath returns the OutputPath field if non-nil, zero value otherwise.

### GetOutputPathOk

`func (o *ModelSyncJobItem) GetOutputPathOk() (*string, bool)`

GetOutputPathOk returns a tuple with the OutputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPath

`func (o *ModelSyncJobItem) SetOutputPath(v string)`

SetOutputPath sets OutputPath field to given value.

### HasOutputPath

`func (o *ModelSyncJobItem) HasOutputPath() bool`

HasOutputPath returns a boolean if a field has been set.

### GetStatus

`func (o *ModelSyncJobItem) GetStatus() ModelModelSyncJobItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelSyncJobItem) GetStatusOk() (*ModelModelSyncJobItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelSyncJobItem) SetStatus(v ModelModelSyncJobItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelSyncJobItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProgress

`func (o *ModelSyncJobItem) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ModelSyncJobItem) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ModelSyncJobItem) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ModelSyncJobItem) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *ModelSyncJobItem) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *ModelSyncJobItem) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetDateCreated

`func (o *ModelSyncJobItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelSyncJobItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelSyncJobItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelSyncJobItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *ModelSyncJobItem) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *ModelSyncJobItem) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *ModelSyncJobItem) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *ModelSyncJobItem) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *ModelSyncJobItem) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *ModelSyncJobItem) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *ModelSyncJobItem) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *ModelSyncJobItem) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetTemporaryPath

`func (o *ModelSyncJobItem) GetTemporaryPath() string`

GetTemporaryPath returns the TemporaryPath field if non-nil, zero value otherwise.

### GetTemporaryPathOk

`func (o *ModelSyncJobItem) GetTemporaryPathOk() (*string, bool)`

GetTemporaryPathOk returns a tuple with the TemporaryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryPath

`func (o *ModelSyncJobItem) SetTemporaryPath(v string)`

SetTemporaryPath sets TemporaryPath field to given value.

### HasTemporaryPath

`func (o *ModelSyncJobItem) HasTemporaryPath() bool`

HasTemporaryPath returns a boolean if a field has been set.

### GetAdditionalFiles

`func (o *ModelSyncJobItem) GetAdditionalFiles() []ModelModelItemFileInfo`

GetAdditionalFiles returns the AdditionalFiles field if non-nil, zero value otherwise.

### GetAdditionalFilesOk

`func (o *ModelSyncJobItem) GetAdditionalFilesOk() (*[]ModelModelItemFileInfo, bool)`

GetAdditionalFilesOk returns a tuple with the AdditionalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFiles

`func (o *ModelSyncJobItem) SetAdditionalFiles(v []ModelModelItemFileInfo)`

SetAdditionalFiles sets AdditionalFiles field to given value.

### HasAdditionalFiles

`func (o *ModelSyncJobItem) HasAdditionalFiles() bool`

HasAdditionalFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


