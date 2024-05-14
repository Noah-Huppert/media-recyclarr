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

// checks if the ModelDlnaProfilesHttpHeaderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDlnaProfilesHttpHeaderInfo{}

// ModelDlnaProfilesHttpHeaderInfo struct for ModelDlnaProfilesHttpHeaderInfo
type ModelDlnaProfilesHttpHeaderInfo struct {
	Name *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
	Match *ModelDlnaProfilesHeaderMatchType `json:"Match,omitempty"`
}

// NewModelDlnaProfilesHttpHeaderInfo instantiates a new ModelDlnaProfilesHttpHeaderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDlnaProfilesHttpHeaderInfo() *ModelDlnaProfilesHttpHeaderInfo {
	this := ModelDlnaProfilesHttpHeaderInfo{}
	return &this
}

// NewModelDlnaProfilesHttpHeaderInfoWithDefaults instantiates a new ModelDlnaProfilesHttpHeaderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDlnaProfilesHttpHeaderInfoWithDefaults() *ModelDlnaProfilesHttpHeaderInfo {
	this := ModelDlnaProfilesHttpHeaderInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelDlnaProfilesHttpHeaderInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDlnaProfilesHttpHeaderInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelDlnaProfilesHttpHeaderInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelDlnaProfilesHttpHeaderInfo) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ModelDlnaProfilesHttpHeaderInfo) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDlnaProfilesHttpHeaderInfo) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ModelDlnaProfilesHttpHeaderInfo) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ModelDlnaProfilesHttpHeaderInfo) SetValue(v string) {
	o.Value = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *ModelDlnaProfilesHttpHeaderInfo) GetMatch() ModelDlnaProfilesHeaderMatchType {
	if o == nil || IsNil(o.Match) {
		var ret ModelDlnaProfilesHeaderMatchType
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDlnaProfilesHttpHeaderInfo) GetMatchOk() (*ModelDlnaProfilesHeaderMatchType, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *ModelDlnaProfilesHttpHeaderInfo) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given ModelDlnaProfilesHeaderMatchType and assigns it to the Match field.
func (o *ModelDlnaProfilesHttpHeaderInfo) SetMatch(v ModelDlnaProfilesHeaderMatchType) {
	o.Match = &v
}

func (o ModelDlnaProfilesHttpHeaderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDlnaProfilesHttpHeaderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	if !IsNil(o.Match) {
		toSerialize["Match"] = o.Match
	}
	return toSerialize, nil
}

type NullableModelDlnaProfilesHttpHeaderInfo struct {
	value *ModelDlnaProfilesHttpHeaderInfo
	isSet bool
}

func (v NullableModelDlnaProfilesHttpHeaderInfo) Get() *ModelDlnaProfilesHttpHeaderInfo {
	return v.value
}

func (v *NullableModelDlnaProfilesHttpHeaderInfo) Set(val *ModelDlnaProfilesHttpHeaderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDlnaProfilesHttpHeaderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDlnaProfilesHttpHeaderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDlnaProfilesHttpHeaderInfo(val *ModelDlnaProfilesHttpHeaderInfo) *NullableModelDlnaProfilesHttpHeaderInfo {
	return &NullableModelDlnaProfilesHttpHeaderInfo{value: val, isSet: true}
}

func (v NullableModelDlnaProfilesHttpHeaderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDlnaProfilesHttpHeaderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

