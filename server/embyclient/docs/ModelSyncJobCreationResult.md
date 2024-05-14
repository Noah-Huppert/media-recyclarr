# ModelSyncJobCreationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**ModelModelSyncJob**](ModelSyncJob.md) |  | [optional] 
**JobItems** | Pointer to [**[]ModelModelSyncJobItem**](ModelModelSyncJobItem.md) |  | [optional] 

## Methods

### NewModelSyncJobCreationResult

`func NewModelSyncJobCreationResult() *ModelSyncJobCreationResult`

NewModelSyncJobCreationResult instantiates a new ModelSyncJobCreationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncJobCreationResultWithDefaults

`func NewModelSyncJobCreationResultWithDefaults() *ModelSyncJobCreationResult`

NewModelSyncJobCreationResultWithDefaults instantiates a new ModelSyncJobCreationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *ModelSyncJobCreationResult) GetJob() ModelModelSyncJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *ModelSyncJobCreationResult) GetJobOk() (*ModelModelSyncJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *ModelSyncJobCreationResult) SetJob(v ModelModelSyncJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *ModelSyncJobCreationResult) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobItems

`func (o *ModelSyncJobCreationResult) GetJobItems() []ModelModelSyncJobItem`

GetJobItems returns the JobItems field if non-nil, zero value otherwise.

### GetJobItemsOk

`func (o *ModelSyncJobCreationResult) GetJobItemsOk() (*[]ModelModelSyncJobItem, bool)`

GetJobItemsOk returns a tuple with the JobItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobItems

`func (o *ModelSyncJobCreationResult) SetJobItems(v []ModelModelSyncJobItem)`

SetJobItems sets JobItems field to given value.

### HasJobItems

`func (o *ModelSyncJobCreationResult) HasJobItems() bool`

HasJobItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


