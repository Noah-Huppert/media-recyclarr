/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
)

// checks if the ModelTranscodingVpStepInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelTranscodingVpStepInfo{}

// ModelTranscodingVpStepInfo struct for ModelTranscodingVpStepInfo
type ModelTranscodingVpStepInfo struct {
	StepType *ModelTranscodingVpStepTypes `json:"StepType,omitempty"`
	StepTypeName *string `json:"StepTypeName,omitempty"`
	HardwareContextName *string `json:"HardwareContextName,omitempty"`
	IsHardwareContext *bool `json:"IsHardwareContext,omitempty"`
	Name *string `json:"Name,omitempty"`
	Short *string `json:"Short,omitempty"`
	FfmpegName *string `json:"FfmpegName,omitempty"`
	FfmpegDescription *string `json:"FfmpegDescription,omitempty"`
	FfmpegOptions *string `json:"FfmpegOptions,omitempty"`
	Param *string `json:"Param,omitempty"`
	ParamShort *string `json:"ParamShort,omitempty"`
}

// NewModelTranscodingVpStepInfo instantiates a new ModelTranscodingVpStepInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelTranscodingVpStepInfo() *ModelTranscodingVpStepInfo {
	this := ModelTranscodingVpStepInfo{}
	return &this
}

// NewModelTranscodingVpStepInfoWithDefaults instantiates a new ModelTranscodingVpStepInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelTranscodingVpStepInfoWithDefaults() *ModelTranscodingVpStepInfo {
	this := ModelTranscodingVpStepInfo{}
	return &this
}

// GetStepType returns the StepType field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetStepType() ModelTranscodingVpStepTypes {
	if o == nil || IsNil(o.StepType) {
		var ret ModelTranscodingVpStepTypes
		return ret
	}
	return *o.StepType
}

// GetStepTypeOk returns a tuple with the StepType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetStepTypeOk() (*ModelTranscodingVpStepTypes, bool) {
	if o == nil || IsNil(o.StepType) {
		return nil, false
	}
	return o.StepType, true
}

// HasStepType returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasStepType() bool {
	if o != nil && !IsNil(o.StepType) {
		return true
	}

	return false
}

// SetStepType gets a reference to the given ModelTranscodingVpStepTypes and assigns it to the StepType field.
func (o *ModelTranscodingVpStepInfo) SetStepType(v ModelTranscodingVpStepTypes) {
	o.StepType = &v
}

// GetStepTypeName returns the StepTypeName field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetStepTypeName() string {
	if o == nil || IsNil(o.StepTypeName) {
		var ret string
		return ret
	}
	return *o.StepTypeName
}

// GetStepTypeNameOk returns a tuple with the StepTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetStepTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.StepTypeName) {
		return nil, false
	}
	return o.StepTypeName, true
}

// HasStepTypeName returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasStepTypeName() bool {
	if o != nil && !IsNil(o.StepTypeName) {
		return true
	}

	return false
}

// SetStepTypeName gets a reference to the given string and assigns it to the StepTypeName field.
func (o *ModelTranscodingVpStepInfo) SetStepTypeName(v string) {
	o.StepTypeName = &v
}

// GetHardwareContextName returns the HardwareContextName field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetHardwareContextName() string {
	if o == nil || IsNil(o.HardwareContextName) {
		var ret string
		return ret
	}
	return *o.HardwareContextName
}

// GetHardwareContextNameOk returns a tuple with the HardwareContextName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetHardwareContextNameOk() (*string, bool) {
	if o == nil || IsNil(o.HardwareContextName) {
		return nil, false
	}
	return o.HardwareContextName, true
}

// HasHardwareContextName returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasHardwareContextName() bool {
	if o != nil && !IsNil(o.HardwareContextName) {
		return true
	}

	return false
}

// SetHardwareContextName gets a reference to the given string and assigns it to the HardwareContextName field.
func (o *ModelTranscodingVpStepInfo) SetHardwareContextName(v string) {
	o.HardwareContextName = &v
}

// GetIsHardwareContext returns the IsHardwareContext field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetIsHardwareContext() bool {
	if o == nil || IsNil(o.IsHardwareContext) {
		var ret bool
		return ret
	}
	return *o.IsHardwareContext
}

// GetIsHardwareContextOk returns a tuple with the IsHardwareContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetIsHardwareContextOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHardwareContext) {
		return nil, false
	}
	return o.IsHardwareContext, true
}

// HasIsHardwareContext returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasIsHardwareContext() bool {
	if o != nil && !IsNil(o.IsHardwareContext) {
		return true
	}

	return false
}

// SetIsHardwareContext gets a reference to the given bool and assigns it to the IsHardwareContext field.
func (o *ModelTranscodingVpStepInfo) SetIsHardwareContext(v bool) {
	o.IsHardwareContext = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelTranscodingVpStepInfo) SetName(v string) {
	o.Name = &v
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetShort() string {
	if o == nil || IsNil(o.Short) {
		var ret string
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetShortOk() (*string, bool) {
	if o == nil || IsNil(o.Short) {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasShort() bool {
	if o != nil && !IsNil(o.Short) {
		return true
	}

	return false
}

// SetShort gets a reference to the given string and assigns it to the Short field.
func (o *ModelTranscodingVpStepInfo) SetShort(v string) {
	o.Short = &v
}

// GetFfmpegName returns the FfmpegName field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetFfmpegName() string {
	if o == nil || IsNil(o.FfmpegName) {
		var ret string
		return ret
	}
	return *o.FfmpegName
}

// GetFfmpegNameOk returns a tuple with the FfmpegName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetFfmpegNameOk() (*string, bool) {
	if o == nil || IsNil(o.FfmpegName) {
		return nil, false
	}
	return o.FfmpegName, true
}

// HasFfmpegName returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasFfmpegName() bool {
	if o != nil && !IsNil(o.FfmpegName) {
		return true
	}

	return false
}

// SetFfmpegName gets a reference to the given string and assigns it to the FfmpegName field.
func (o *ModelTranscodingVpStepInfo) SetFfmpegName(v string) {
	o.FfmpegName = &v
}

// GetFfmpegDescription returns the FfmpegDescription field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetFfmpegDescription() string {
	if o == nil || IsNil(o.FfmpegDescription) {
		var ret string
		return ret
	}
	return *o.FfmpegDescription
}

// GetFfmpegDescriptionOk returns a tuple with the FfmpegDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetFfmpegDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.FfmpegDescription) {
		return nil, false
	}
	return o.FfmpegDescription, true
}

// HasFfmpegDescription returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasFfmpegDescription() bool {
	if o != nil && !IsNil(o.FfmpegDescription) {
		return true
	}

	return false
}

// SetFfmpegDescription gets a reference to the given string and assigns it to the FfmpegDescription field.
func (o *ModelTranscodingVpStepInfo) SetFfmpegDescription(v string) {
	o.FfmpegDescription = &v
}

// GetFfmpegOptions returns the FfmpegOptions field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetFfmpegOptions() string {
	if o == nil || IsNil(o.FfmpegOptions) {
		var ret string
		return ret
	}
	return *o.FfmpegOptions
}

// GetFfmpegOptionsOk returns a tuple with the FfmpegOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetFfmpegOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.FfmpegOptions) {
		return nil, false
	}
	return o.FfmpegOptions, true
}

// HasFfmpegOptions returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasFfmpegOptions() bool {
	if o != nil && !IsNil(o.FfmpegOptions) {
		return true
	}

	return false
}

// SetFfmpegOptions gets a reference to the given string and assigns it to the FfmpegOptions field.
func (o *ModelTranscodingVpStepInfo) SetFfmpegOptions(v string) {
	o.FfmpegOptions = &v
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *ModelTranscodingVpStepInfo) SetParam(v string) {
	o.Param = &v
}

// GetParamShort returns the ParamShort field value if set, zero value otherwise.
func (o *ModelTranscodingVpStepInfo) GetParamShort() string {
	if o == nil || IsNil(o.ParamShort) {
		var ret string
		return ret
	}
	return *o.ParamShort
}

// GetParamShortOk returns a tuple with the ParamShort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTranscodingVpStepInfo) GetParamShortOk() (*string, bool) {
	if o == nil || IsNil(o.ParamShort) {
		return nil, false
	}
	return o.ParamShort, true
}

// HasParamShort returns a boolean if a field has been set.
func (o *ModelTranscodingVpStepInfo) HasParamShort() bool {
	if o != nil && !IsNil(o.ParamShort) {
		return true
	}

	return false
}

// SetParamShort gets a reference to the given string and assigns it to the ParamShort field.
func (o *ModelTranscodingVpStepInfo) SetParamShort(v string) {
	o.ParamShort = &v
}

func (o ModelTranscodingVpStepInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelTranscodingVpStepInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StepType) {
		toSerialize["StepType"] = o.StepType
	}
	if !IsNil(o.StepTypeName) {
		toSerialize["StepTypeName"] = o.StepTypeName
	}
	if !IsNil(o.HardwareContextName) {
		toSerialize["HardwareContextName"] = o.HardwareContextName
	}
	if !IsNil(o.IsHardwareContext) {
		toSerialize["IsHardwareContext"] = o.IsHardwareContext
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Short) {
		toSerialize["Short"] = o.Short
	}
	if !IsNil(o.FfmpegName) {
		toSerialize["FfmpegName"] = o.FfmpegName
	}
	if !IsNil(o.FfmpegDescription) {
		toSerialize["FfmpegDescription"] = o.FfmpegDescription
	}
	if !IsNil(o.FfmpegOptions) {
		toSerialize["FfmpegOptions"] = o.FfmpegOptions
	}
	if !IsNil(o.Param) {
		toSerialize["Param"] = o.Param
	}
	if !IsNil(o.ParamShort) {
		toSerialize["ParamShort"] = o.ParamShort
	}
	return toSerialize, nil
}

type NullableModelTranscodingVpStepInfo struct {
	value *ModelTranscodingVpStepInfo
	isSet bool
}

func (v NullableModelTranscodingVpStepInfo) Get() *ModelTranscodingVpStepInfo {
	return v.value
}

func (v *NullableModelTranscodingVpStepInfo) Set(val *ModelTranscodingVpStepInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelTranscodingVpStepInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelTranscodingVpStepInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelTranscodingVpStepInfo(val *ModelTranscodingVpStepInfo) *NullableModelTranscodingVpStepInfo {
	return &NullableModelTranscodingVpStepInfo{value: val, isSet: true}
}

func (v NullableModelTranscodingVpStepInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelTranscodingVpStepInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


