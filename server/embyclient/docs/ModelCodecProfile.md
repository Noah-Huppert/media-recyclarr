# ModelCodecProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModelModelCodecType**](ModelCodecType.md) |  | [optional] 
**Conditions** | Pointer to [**[]ModelModelProfileCondition**](ModelModelProfileCondition.md) |  | [optional] 
**ApplyConditions** | Pointer to [**[]ModelModelProfileCondition**](ModelModelProfileCondition.md) |  | [optional] 
**Codec** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCodecProfile

`func NewModelCodecProfile() *ModelCodecProfile`

NewModelCodecProfile instantiates a new ModelCodecProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCodecProfileWithDefaults

`func NewModelCodecProfileWithDefaults() *ModelCodecProfile`

NewModelCodecProfileWithDefaults instantiates a new ModelCodecProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelCodecProfile) GetType() ModelModelCodecType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelCodecProfile) GetTypeOk() (*ModelModelCodecType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelCodecProfile) SetType(v ModelModelCodecType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelCodecProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *ModelCodecProfile) GetConditions() []ModelModelProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ModelCodecProfile) GetConditionsOk() (*[]ModelModelProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ModelCodecProfile) SetConditions(v []ModelModelProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ModelCodecProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetApplyConditions

`func (o *ModelCodecProfile) GetApplyConditions() []ModelModelProfileCondition`

GetApplyConditions returns the ApplyConditions field if non-nil, zero value otherwise.

### GetApplyConditionsOk

`func (o *ModelCodecProfile) GetApplyConditionsOk() (*[]ModelModelProfileCondition, bool)`

GetApplyConditionsOk returns a tuple with the ApplyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyConditions

`func (o *ModelCodecProfile) SetApplyConditions(v []ModelModelProfileCondition)`

SetApplyConditions sets ApplyConditions field to given value.

### HasApplyConditions

`func (o *ModelCodecProfile) HasApplyConditions() bool`

HasApplyConditions returns a boolean if a field has been set.

### GetCodec

`func (o *ModelCodecProfile) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *ModelCodecProfile) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *ModelCodecProfile) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *ModelCodecProfile) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetContainer

`func (o *ModelCodecProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelCodecProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelCodecProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelCodecProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


