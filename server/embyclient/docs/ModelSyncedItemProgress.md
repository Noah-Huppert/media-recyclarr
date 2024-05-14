# ModelSyncedItemProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to **NullableFloat64** |  | [optional] 
**Status** | Pointer to [**ModelModelSyncJobItemStatus**](ModelSyncJobItemStatus.md) |  | [optional] 

## Methods

### NewModelSyncedItemProgress

`func NewModelSyncedItemProgress() *ModelSyncedItemProgress`

NewModelSyncedItemProgress instantiates a new ModelSyncedItemProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncedItemProgressWithDefaults

`func NewModelSyncedItemProgressWithDefaults() *ModelSyncedItemProgress`

NewModelSyncedItemProgressWithDefaults instantiates a new ModelSyncedItemProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *ModelSyncedItemProgress) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ModelSyncedItemProgress) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ModelSyncedItemProgress) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ModelSyncedItemProgress) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### SetProgressNil

`func (o *ModelSyncedItemProgress) SetProgressNil(b bool)`

 SetProgressNil sets the value for Progress to be an explicit nil

### UnsetProgress
`func (o *ModelSyncedItemProgress) UnsetProgress()`

UnsetProgress ensures that no value is present for Progress, not even an explicit nil
### GetStatus

`func (o *ModelSyncedItemProgress) GetStatus() ModelModelSyncJobItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelSyncedItemProgress) GetStatusOk() (*ModelModelSyncJobItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelSyncedItemProgress) SetStatus(v ModelModelSyncJobItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelSyncedItemProgress) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


