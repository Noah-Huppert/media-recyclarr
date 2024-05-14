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

// checks if the ModelCodecProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCodecProfile{}

// ModelCodecProfile struct for ModelCodecProfile
type ModelCodecProfile struct {
	Type *ModelCodecType `json:"Type,omitempty"`
	Conditions []ModelProfileCondition `json:"Conditions,omitempty"`
	ApplyConditions []ModelProfileCondition `json:"ApplyConditions,omitempty"`
	Codec *string `json:"Codec,omitempty"`
	Container *string `json:"Container,omitempty"`
}

// NewModelCodecProfile instantiates a new ModelCodecProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCodecProfile() *ModelCodecProfile {
	this := ModelCodecProfile{}
	return &this
}

// NewModelCodecProfileWithDefaults instantiates a new ModelCodecProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCodecProfileWithDefaults() *ModelCodecProfile {
	this := ModelCodecProfile{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelCodecProfile) GetType() ModelCodecType {
	if o == nil || IsNil(o.Type) {
		var ret ModelCodecType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCodecProfile) GetTypeOk() (*ModelCodecType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelCodecProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModelCodecType and assigns it to the Type field.
func (o *ModelCodecProfile) SetType(v ModelCodecType) {
	o.Type = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ModelCodecProfile) GetConditions() []ModelProfileCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []ModelProfileCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCodecProfile) GetConditionsOk() ([]ModelProfileCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ModelCodecProfile) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ModelProfileCondition and assigns it to the Conditions field.
func (o *ModelCodecProfile) SetConditions(v []ModelProfileCondition) {
	o.Conditions = v
}

// GetApplyConditions returns the ApplyConditions field value if set, zero value otherwise.
func (o *ModelCodecProfile) GetApplyConditions() []ModelProfileCondition {
	if o == nil || IsNil(o.ApplyConditions) {
		var ret []ModelProfileCondition
		return ret
	}
	return o.ApplyConditions
}

// GetApplyConditionsOk returns a tuple with the ApplyConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCodecProfile) GetApplyConditionsOk() ([]ModelProfileCondition, bool) {
	if o == nil || IsNil(o.ApplyConditions) {
		return nil, false
	}
	return o.ApplyConditions, true
}

// HasApplyConditions returns a boolean if a field has been set.
func (o *ModelCodecProfile) HasApplyConditions() bool {
	if o != nil && !IsNil(o.ApplyConditions) {
		return true
	}

	return false
}

// SetApplyConditions gets a reference to the given []ModelProfileCondition and assigns it to the ApplyConditions field.
func (o *ModelCodecProfile) SetApplyConditions(v []ModelProfileCondition) {
	o.ApplyConditions = v
}

// GetCodec returns the Codec field value if set, zero value otherwise.
func (o *ModelCodecProfile) GetCodec() string {
	if o == nil || IsNil(o.Codec) {
		var ret string
		return ret
	}
	return *o.Codec
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCodecProfile) GetCodecOk() (*string, bool) {
	if o == nil || IsNil(o.Codec) {
		return nil, false
	}
	return o.Codec, true
}

// HasCodec returns a boolean if a field has been set.
func (o *ModelCodecProfile) HasCodec() bool {
	if o != nil && !IsNil(o.Codec) {
		return true
	}

	return false
}

// SetCodec gets a reference to the given string and assigns it to the Codec field.
func (o *ModelCodecProfile) SetCodec(v string) {
	o.Codec = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *ModelCodecProfile) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCodecProfile) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *ModelCodecProfile) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *ModelCodecProfile) SetContainer(v string) {
	o.Container = &v
}

func (o ModelCodecProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCodecProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Conditions) {
		toSerialize["Conditions"] = o.Conditions
	}
	if !IsNil(o.ApplyConditions) {
		toSerialize["ApplyConditions"] = o.ApplyConditions
	}
	if !IsNil(o.Codec) {
		toSerialize["Codec"] = o.Codec
	}
	if !IsNil(o.Container) {
		toSerialize["Container"] = o.Container
	}
	return toSerialize, nil
}

type NullableModelCodecProfile struct {
	value *ModelCodecProfile
	isSet bool
}

func (v NullableModelCodecProfile) Get() *ModelCodecProfile {
	return v.value
}

func (v *NullableModelCodecProfile) Set(val *ModelCodecProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCodecProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCodecProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCodecProfile(val *ModelCodecProfile) *NullableModelCodecProfile {
	return &NullableModelCodecProfile{value: val, isSet: true}
}

func (v NullableModelCodecProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCodecProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


