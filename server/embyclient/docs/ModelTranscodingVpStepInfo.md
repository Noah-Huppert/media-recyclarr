# ModelTranscodingVpStepInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepType** | Pointer to [**ModelModelTranscodingVpStepTypes**](ModelTranscodingVpStepTypes.md) |  | [optional] 
**StepTypeName** | Pointer to **string** |  | [optional] 
**HardwareContextName** | Pointer to **string** |  | [optional] 
**IsHardwareContext** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Short** | Pointer to **string** |  | [optional] 
**FfmpegName** | Pointer to **string** |  | [optional] 
**FfmpegDescription** | Pointer to **string** |  | [optional] 
**FfmpegOptions** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **string** |  | [optional] 
**ParamShort** | Pointer to **string** |  | [optional] 

## Methods

### NewModelTranscodingVpStepInfo

`func NewModelTranscodingVpStepInfo() *ModelTranscodingVpStepInfo`

NewModelTranscodingVpStepInfo instantiates a new ModelTranscodingVpStepInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTranscodingVpStepInfoWithDefaults

`func NewModelTranscodingVpStepInfoWithDefaults() *ModelTranscodingVpStepInfo`

NewModelTranscodingVpStepInfoWithDefaults instantiates a new ModelTranscodingVpStepInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStepType

`func (o *ModelTranscodingVpStepInfo) GetStepType() ModelModelTranscodingVpStepTypes`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *ModelTranscodingVpStepInfo) GetStepTypeOk() (*ModelModelTranscodingVpStepTypes, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *ModelTranscodingVpStepInfo) SetStepType(v ModelModelTranscodingVpStepTypes)`

SetStepType sets StepType field to given value.

### HasStepType

`func (o *ModelTranscodingVpStepInfo) HasStepType() bool`

HasStepType returns a boolean if a field has been set.

### GetStepTypeName

`func (o *ModelTranscodingVpStepInfo) GetStepTypeName() string`

GetStepTypeName returns the StepTypeName field if non-nil, zero value otherwise.

### GetStepTypeNameOk

`func (o *ModelTranscodingVpStepInfo) GetStepTypeNameOk() (*string, bool)`

GetStepTypeNameOk returns a tuple with the StepTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepTypeName

`func (o *ModelTranscodingVpStepInfo) SetStepTypeName(v string)`

SetStepTypeName sets StepTypeName field to given value.

### HasStepTypeName

`func (o *ModelTranscodingVpStepInfo) HasStepTypeName() bool`

HasStepTypeName returns a boolean if a field has been set.

### GetHardwareContextName

`func (o *ModelTranscodingVpStepInfo) GetHardwareContextName() string`

GetHardwareContextName returns the HardwareContextName field if non-nil, zero value otherwise.

### GetHardwareContextNameOk

`func (o *ModelTranscodingVpStepInfo) GetHardwareContextNameOk() (*string, bool)`

GetHardwareContextNameOk returns a tuple with the HardwareContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareContextName

`func (o *ModelTranscodingVpStepInfo) SetHardwareContextName(v string)`

SetHardwareContextName sets HardwareContextName field to given value.

### HasHardwareContextName

`func (o *ModelTranscodingVpStepInfo) HasHardwareContextName() bool`

HasHardwareContextName returns a boolean if a field has been set.

### GetIsHardwareContext

`func (o *ModelTranscodingVpStepInfo) GetIsHardwareContext() bool`

GetIsHardwareContext returns the IsHardwareContext field if non-nil, zero value otherwise.

### GetIsHardwareContextOk

`func (o *ModelTranscodingVpStepInfo) GetIsHardwareContextOk() (*bool, bool)`

GetIsHardwareContextOk returns a tuple with the IsHardwareContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardwareContext

`func (o *ModelTranscodingVpStepInfo) SetIsHardwareContext(v bool)`

SetIsHardwareContext sets IsHardwareContext field to given value.

### HasIsHardwareContext

`func (o *ModelTranscodingVpStepInfo) HasIsHardwareContext() bool`

HasIsHardwareContext returns a boolean if a field has been set.

### GetName

`func (o *ModelTranscodingVpStepInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTranscodingVpStepInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTranscodingVpStepInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTranscodingVpStepInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShort

`func (o *ModelTranscodingVpStepInfo) GetShort() string`

GetShort returns the Short field if non-nil, zero value otherwise.

### GetShortOk

`func (o *ModelTranscodingVpStepInfo) GetShortOk() (*string, bool)`

GetShortOk returns a tuple with the Short field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShort

`func (o *ModelTranscodingVpStepInfo) SetShort(v string)`

SetShort sets Short field to given value.

### HasShort

`func (o *ModelTranscodingVpStepInfo) HasShort() bool`

HasShort returns a boolean if a field has been set.

### GetFfmpegName

`func (o *ModelTranscodingVpStepInfo) GetFfmpegName() string`

GetFfmpegName returns the FfmpegName field if non-nil, zero value otherwise.

### GetFfmpegNameOk

`func (o *ModelTranscodingVpStepInfo) GetFfmpegNameOk() (*string, bool)`

GetFfmpegNameOk returns a tuple with the FfmpegName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfmpegName

`func (o *ModelTranscodingVpStepInfo) SetFfmpegName(v string)`

SetFfmpegName sets FfmpegName field to given value.

### HasFfmpegName

`func (o *ModelTranscodingVpStepInfo) HasFfmpegName() bool`

HasFfmpegName returns a boolean if a field has been set.

### GetFfmpegDescription

`func (o *ModelTranscodingVpStepInfo) GetFfmpegDescription() string`

GetFfmpegDescription returns the FfmpegDescription field if non-nil, zero value otherwise.

### GetFfmpegDescriptionOk

`func (o *ModelTranscodingVpStepInfo) GetFfmpegDescriptionOk() (*string, bool)`

GetFfmpegDescriptionOk returns a tuple with the FfmpegDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfmpegDescription

`func (o *ModelTranscodingVpStepInfo) SetFfmpegDescription(v string)`

SetFfmpegDescription sets FfmpegDescription field to given value.

### HasFfmpegDescription

`func (o *ModelTranscodingVpStepInfo) HasFfmpegDescription() bool`

HasFfmpegDescription returns a boolean if a field has been set.

### GetFfmpegOptions

`func (o *ModelTranscodingVpStepInfo) GetFfmpegOptions() string`

GetFfmpegOptions returns the FfmpegOptions field if non-nil, zero value otherwise.

### GetFfmpegOptionsOk

`func (o *ModelTranscodingVpStepInfo) GetFfmpegOptionsOk() (*string, bool)`

GetFfmpegOptionsOk returns a tuple with the FfmpegOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFfmpegOptions

`func (o *ModelTranscodingVpStepInfo) SetFfmpegOptions(v string)`

SetFfmpegOptions sets FfmpegOptions field to given value.

### HasFfmpegOptions

`func (o *ModelTranscodingVpStepInfo) HasFfmpegOptions() bool`

HasFfmpegOptions returns a boolean if a field has been set.

### GetParam

`func (o *ModelTranscodingVpStepInfo) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ModelTranscodingVpStepInfo) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ModelTranscodingVpStepInfo) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ModelTranscodingVpStepInfo) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetParamShort

`func (o *ModelTranscodingVpStepInfo) GetParamShort() string`

GetParamShort returns the ParamShort field if non-nil, zero value otherwise.

### GetParamShortOk

`func (o *ModelTranscodingVpStepInfo) GetParamShortOk() (*string, bool)`

GetParamShortOk returns a tuple with the ParamShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamShort

`func (o *ModelTranscodingVpStepInfo) SetParamShort(v string)`

SetParamShort sets ParamShort field to given value.

### HasParamShort

`func (o *ModelTranscodingVpStepInfo) HasParamShort() bool`

HasParamShort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


