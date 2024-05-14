# ModelConditionsPropertyCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPropertyId** | Pointer to **string** |  | [optional] 
**ConditionType** | Pointer to [**ModelModelConditionsPropertyConditionType**](ModelConditionsPropertyConditionType.md) |  | [optional] 
**TargetPropertyId** | Pointer to **string** |  | [optional] 
**SimpleCondition** | Pointer to [**ModelModelAttributesSimpleCondition**](ModelAttributesSimpleCondition.md) |  | [optional] 
**ValueCondition** | Pointer to [**ModelModelAttributesValueCondition**](ModelAttributesValueCondition.md) |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelConditionsPropertyCondition

`func NewModelConditionsPropertyCondition() *ModelConditionsPropertyCondition`

NewModelConditionsPropertyCondition instantiates a new ModelConditionsPropertyCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelConditionsPropertyConditionWithDefaults

`func NewModelConditionsPropertyConditionWithDefaults() *ModelConditionsPropertyCondition`

NewModelConditionsPropertyConditionWithDefaults instantiates a new ModelConditionsPropertyCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPropertyId

`func (o *ModelConditionsPropertyCondition) GetAffectedPropertyId() string`

GetAffectedPropertyId returns the AffectedPropertyId field if non-nil, zero value otherwise.

### GetAffectedPropertyIdOk

`func (o *ModelConditionsPropertyCondition) GetAffectedPropertyIdOk() (*string, bool)`

GetAffectedPropertyIdOk returns a tuple with the AffectedPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPropertyId

`func (o *ModelConditionsPropertyCondition) SetAffectedPropertyId(v string)`

SetAffectedPropertyId sets AffectedPropertyId field to given value.

### HasAffectedPropertyId

`func (o *ModelConditionsPropertyCondition) HasAffectedPropertyId() bool`

HasAffectedPropertyId returns a boolean if a field has been set.

### GetConditionType

`func (o *ModelConditionsPropertyCondition) GetConditionType() ModelModelConditionsPropertyConditionType`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ModelConditionsPropertyCondition) GetConditionTypeOk() (*ModelModelConditionsPropertyConditionType, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ModelConditionsPropertyCondition) SetConditionType(v ModelModelConditionsPropertyConditionType)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *ModelConditionsPropertyCondition) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetTargetPropertyId

`func (o *ModelConditionsPropertyCondition) GetTargetPropertyId() string`

GetTargetPropertyId returns the TargetPropertyId field if non-nil, zero value otherwise.

### GetTargetPropertyIdOk

`func (o *ModelConditionsPropertyCondition) GetTargetPropertyIdOk() (*string, bool)`

GetTargetPropertyIdOk returns a tuple with the TargetPropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPropertyId

`func (o *ModelConditionsPropertyCondition) SetTargetPropertyId(v string)`

SetTargetPropertyId sets TargetPropertyId field to given value.

### HasTargetPropertyId

`func (o *ModelConditionsPropertyCondition) HasTargetPropertyId() bool`

HasTargetPropertyId returns a boolean if a field has been set.

### GetSimpleCondition

`func (o *ModelConditionsPropertyCondition) GetSimpleCondition() ModelModelAttributesSimpleCondition`

GetSimpleCondition returns the SimpleCondition field if non-nil, zero value otherwise.

### GetSimpleConditionOk

`func (o *ModelConditionsPropertyCondition) GetSimpleConditionOk() (*ModelModelAttributesSimpleCondition, bool)`

GetSimpleConditionOk returns a tuple with the SimpleCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleCondition

`func (o *ModelConditionsPropertyCondition) SetSimpleCondition(v ModelModelAttributesSimpleCondition)`

SetSimpleCondition sets SimpleCondition field to given value.

### HasSimpleCondition

`func (o *ModelConditionsPropertyCondition) HasSimpleCondition() bool`

HasSimpleCondition returns a boolean if a field has been set.

### GetValueCondition

`func (o *ModelConditionsPropertyCondition) GetValueCondition() ModelModelAttributesValueCondition`

GetValueCondition returns the ValueCondition field if non-nil, zero value otherwise.

### GetValueConditionOk

`func (o *ModelConditionsPropertyCondition) GetValueConditionOk() (*ModelModelAttributesValueCondition, bool)`

GetValueConditionOk returns a tuple with the ValueCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCondition

`func (o *ModelConditionsPropertyCondition) SetValueCondition(v ModelModelAttributesValueCondition)`

SetValueCondition sets ValueCondition field to given value.

### HasValueCondition

`func (o *ModelConditionsPropertyCondition) HasValueCondition() bool`

HasValueCondition returns a boolean if a field has been set.

### GetValue

`func (o *ModelConditionsPropertyCondition) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelConditionsPropertyCondition) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelConditionsPropertyCondition) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelConditionsPropertyCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


