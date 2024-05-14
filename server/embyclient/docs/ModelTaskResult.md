# ModelTaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeUtc** | Pointer to **time.Time** |  | [optional] 
**EndTimeUtc** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**ModelModelTaskCompletionStatus**](ModelTaskCompletionStatus.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**LongErrorMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewModelTaskResult

`func NewModelTaskResult() *ModelTaskResult`

NewModelTaskResult instantiates a new ModelTaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTaskResultWithDefaults

`func NewModelTaskResultWithDefaults() *ModelTaskResult`

NewModelTaskResultWithDefaults instantiates a new ModelTaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeUtc

`func (o *ModelTaskResult) GetStartTimeUtc() time.Time`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *ModelTaskResult) GetStartTimeUtcOk() (*time.Time, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *ModelTaskResult) SetStartTimeUtc(v time.Time)`

SetStartTimeUtc sets StartTimeUtc field to given value.

### HasStartTimeUtc

`func (o *ModelTaskResult) HasStartTimeUtc() bool`

HasStartTimeUtc returns a boolean if a field has been set.

### GetEndTimeUtc

`func (o *ModelTaskResult) GetEndTimeUtc() time.Time`

GetEndTimeUtc returns the EndTimeUtc field if non-nil, zero value otherwise.

### GetEndTimeUtcOk

`func (o *ModelTaskResult) GetEndTimeUtcOk() (*time.Time, bool)`

GetEndTimeUtcOk returns a tuple with the EndTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUtc

`func (o *ModelTaskResult) SetEndTimeUtc(v time.Time)`

SetEndTimeUtc sets EndTimeUtc field to given value.

### HasEndTimeUtc

`func (o *ModelTaskResult) HasEndTimeUtc() bool`

HasEndTimeUtc returns a boolean if a field has been set.

### GetStatus

`func (o *ModelTaskResult) GetStatus() ModelModelTaskCompletionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelTaskResult) GetStatusOk() (*ModelModelTaskCompletionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelTaskResult) SetStatus(v ModelModelTaskCompletionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelTaskResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *ModelTaskResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTaskResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTaskResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTaskResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKey

`func (o *ModelTaskResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelTaskResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelTaskResult) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelTaskResult) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *ModelTaskResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelTaskResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelTaskResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelTaskResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ModelTaskResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ModelTaskResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ModelTaskResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ModelTaskResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetLongErrorMessage

`func (o *ModelTaskResult) GetLongErrorMessage() string`

GetLongErrorMessage returns the LongErrorMessage field if non-nil, zero value otherwise.

### GetLongErrorMessageOk

`func (o *ModelTaskResult) GetLongErrorMessageOk() (*string, bool)`

GetLongErrorMessageOk returns a tuple with the LongErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongErrorMessage

`func (o *ModelTaskResult) SetLongErrorMessage(v string)`

SetLongErrorMessage sets LongErrorMessage field to given value.

### HasLongErrorMessage

`func (o *ModelTaskResult) HasLongErrorMessage() bool`

HasLongErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


