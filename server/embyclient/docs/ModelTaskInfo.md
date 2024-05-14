# ModelTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**ModelModelTaskState**](ModelTaskState.md) |  | [optional] 
**CurrentProgressPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastExecutionResult** | Pointer to [**ModelModelTaskResult**](ModelTaskResult.md) |  | [optional] 
**Triggers** | Pointer to [**[]ModelModelTaskTriggerInfo**](ModelModelTaskTriggerInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewModelTaskInfo

`func NewModelTaskInfo() *ModelTaskInfo`

NewModelTaskInfo instantiates a new ModelTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTaskInfoWithDefaults

`func NewModelTaskInfoWithDefaults() *ModelTaskInfo`

NewModelTaskInfoWithDefaults instantiates a new ModelTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ModelTaskInfo) GetState() ModelModelTaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelTaskInfo) GetStateOk() (*ModelModelTaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelTaskInfo) SetState(v ModelModelTaskState)`

SetState sets State field to given value.

### HasState

`func (o *ModelTaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCurrentProgressPercentage

`func (o *ModelTaskInfo) GetCurrentProgressPercentage() float64`

GetCurrentProgressPercentage returns the CurrentProgressPercentage field if non-nil, zero value otherwise.

### GetCurrentProgressPercentageOk

`func (o *ModelTaskInfo) GetCurrentProgressPercentageOk() (*float64, bool)`

GetCurrentProgressPercentageOk returns a tuple with the CurrentProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgressPercentage

`func (o *ModelTaskInfo) SetCurrentProgressPercentage(v float64)`

SetCurrentProgressPercentage sets CurrentProgressPercentage field to given value.

### HasCurrentProgressPercentage

`func (o *ModelTaskInfo) HasCurrentProgressPercentage() bool`

HasCurrentProgressPercentage returns a boolean if a field has been set.

### SetCurrentProgressPercentageNil

`func (o *ModelTaskInfo) SetCurrentProgressPercentageNil(b bool)`

 SetCurrentProgressPercentageNil sets the value for CurrentProgressPercentage to be an explicit nil

### UnsetCurrentProgressPercentage
`func (o *ModelTaskInfo) UnsetCurrentProgressPercentage()`

UnsetCurrentProgressPercentage ensures that no value is present for CurrentProgressPercentage, not even an explicit nil
### GetId

`func (o *ModelTaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelTaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelTaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelTaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastExecutionResult

`func (o *ModelTaskInfo) GetLastExecutionResult() ModelModelTaskResult`

GetLastExecutionResult returns the LastExecutionResult field if non-nil, zero value otherwise.

### GetLastExecutionResultOk

`func (o *ModelTaskInfo) GetLastExecutionResultOk() (*ModelModelTaskResult, bool)`

GetLastExecutionResultOk returns a tuple with the LastExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionResult

`func (o *ModelTaskInfo) SetLastExecutionResult(v ModelModelTaskResult)`

SetLastExecutionResult sets LastExecutionResult field to given value.

### HasLastExecutionResult

`func (o *ModelTaskInfo) HasLastExecutionResult() bool`

HasLastExecutionResult returns a boolean if a field has been set.

### GetTriggers

`func (o *ModelTaskInfo) GetTriggers() []ModelModelTaskTriggerInfo`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ModelTaskInfo) GetTriggersOk() (*[]ModelModelTaskTriggerInfo, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ModelTaskInfo) SetTriggers(v []ModelModelTaskTriggerInfo)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ModelTaskInfo) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetDescription

`func (o *ModelTaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelTaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelTaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelTaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ModelTaskInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelTaskInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelTaskInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelTaskInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetIsHidden

`func (o *ModelTaskInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ModelTaskInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ModelTaskInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *ModelTaskInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetKey

`func (o *ModelTaskInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelTaskInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelTaskInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelTaskInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


