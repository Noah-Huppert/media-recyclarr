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

// checks if the ModelSyncDialogOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSyncDialogOptions{}

// ModelSyncDialogOptions struct for ModelSyncDialogOptions
type ModelSyncDialogOptions struct {
	Targets []ModelSyncTarget `json:"Targets,omitempty"`
	Options []ModelSyncJobOption `json:"Options,omitempty"`
	QualityOptions []ModelSyncQualityOption `json:"QualityOptions,omitempty"`
	ProfileOptions []ModelSyncProfileOption `json:"ProfileOptions,omitempty"`
}

// NewModelSyncDialogOptions instantiates a new ModelSyncDialogOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSyncDialogOptions() *ModelSyncDialogOptions {
	this := ModelSyncDialogOptions{}
	return &this
}

// NewModelSyncDialogOptionsWithDefaults instantiates a new ModelSyncDialogOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSyncDialogOptionsWithDefaults() *ModelSyncDialogOptions {
	this := ModelSyncDialogOptions{}
	return &this
}

// GetTargets returns the Targets field value if set, zero value otherwise.
func (o *ModelSyncDialogOptions) GetTargets() []ModelSyncTarget {
	if o == nil || IsNil(o.Targets) {
		var ret []ModelSyncTarget
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncDialogOptions) GetTargetsOk() ([]ModelSyncTarget, bool) {
	if o == nil || IsNil(o.Targets) {
		return nil, false
	}
	return o.Targets, true
}

// HasTargets returns a boolean if a field has been set.
func (o *ModelSyncDialogOptions) HasTargets() bool {
	if o != nil && !IsNil(o.Targets) {
		return true
	}

	return false
}

// SetTargets gets a reference to the given []ModelSyncTarget and assigns it to the Targets field.
func (o *ModelSyncDialogOptions) SetTargets(v []ModelSyncTarget) {
	o.Targets = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ModelSyncDialogOptions) GetOptions() []ModelSyncJobOption {
	if o == nil || IsNil(o.Options) {
		var ret []ModelSyncJobOption
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncDialogOptions) GetOptionsOk() ([]ModelSyncJobOption, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ModelSyncDialogOptions) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []ModelSyncJobOption and assigns it to the Options field.
func (o *ModelSyncDialogOptions) SetOptions(v []ModelSyncJobOption) {
	o.Options = v
}

// GetQualityOptions returns the QualityOptions field value if set, zero value otherwise.
func (o *ModelSyncDialogOptions) GetQualityOptions() []ModelSyncQualityOption {
	if o == nil || IsNil(o.QualityOptions) {
		var ret []ModelSyncQualityOption
		return ret
	}
	return o.QualityOptions
}

// GetQualityOptionsOk returns a tuple with the QualityOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncDialogOptions) GetQualityOptionsOk() ([]ModelSyncQualityOption, bool) {
	if o == nil || IsNil(o.QualityOptions) {
		return nil, false
	}
	return o.QualityOptions, true
}

// HasQualityOptions returns a boolean if a field has been set.
func (o *ModelSyncDialogOptions) HasQualityOptions() bool {
	if o != nil && !IsNil(o.QualityOptions) {
		return true
	}

	return false
}

// SetQualityOptions gets a reference to the given []ModelSyncQualityOption and assigns it to the QualityOptions field.
func (o *ModelSyncDialogOptions) SetQualityOptions(v []ModelSyncQualityOption) {
	o.QualityOptions = v
}

// GetProfileOptions returns the ProfileOptions field value if set, zero value otherwise.
func (o *ModelSyncDialogOptions) GetProfileOptions() []ModelSyncProfileOption {
	if o == nil || IsNil(o.ProfileOptions) {
		var ret []ModelSyncProfileOption
		return ret
	}
	return o.ProfileOptions
}

// GetProfileOptionsOk returns a tuple with the ProfileOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncDialogOptions) GetProfileOptionsOk() ([]ModelSyncProfileOption, bool) {
	if o == nil || IsNil(o.ProfileOptions) {
		return nil, false
	}
	return o.ProfileOptions, true
}

// HasProfileOptions returns a boolean if a field has been set.
func (o *ModelSyncDialogOptions) HasProfileOptions() bool {
	if o != nil && !IsNil(o.ProfileOptions) {
		return true
	}

	return false
}

// SetProfileOptions gets a reference to the given []ModelSyncProfileOption and assigns it to the ProfileOptions field.
func (o *ModelSyncDialogOptions) SetProfileOptions(v []ModelSyncProfileOption) {
	o.ProfileOptions = v
}

func (o ModelSyncDialogOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSyncDialogOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Targets) {
		toSerialize["Targets"] = o.Targets
	}
	if !IsNil(o.Options) {
		toSerialize["Options"] = o.Options
	}
	if !IsNil(o.QualityOptions) {
		toSerialize["QualityOptions"] = o.QualityOptions
	}
	if !IsNil(o.ProfileOptions) {
		toSerialize["ProfileOptions"] = o.ProfileOptions
	}
	return toSerialize, nil
}

type NullableModelSyncDialogOptions struct {
	value *ModelSyncDialogOptions
	isSet bool
}

func (v NullableModelSyncDialogOptions) Get() *ModelSyncDialogOptions {
	return v.value
}

func (v *NullableModelSyncDialogOptions) Set(val *ModelSyncDialogOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSyncDialogOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSyncDialogOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSyncDialogOptions(val *ModelSyncDialogOptions) *NullableModelSyncDialogOptions {
	return &NullableModelSyncDialogOptions{value: val, isSet: true}
}

func (v NullableModelSyncDialogOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSyncDialogOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

