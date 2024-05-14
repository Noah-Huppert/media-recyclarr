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

// checks if the ModelLibraryDeleteInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLibraryDeleteInfo{}

// ModelLibraryDeleteInfo struct for ModelLibraryDeleteInfo
type ModelLibraryDeleteInfo struct {
	Paths []string `json:"Paths,omitempty"`
}

// NewModelLibraryDeleteInfo instantiates a new ModelLibraryDeleteInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLibraryDeleteInfo() *ModelLibraryDeleteInfo {
	this := ModelLibraryDeleteInfo{}
	return &this
}

// NewModelLibraryDeleteInfoWithDefaults instantiates a new ModelLibraryDeleteInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLibraryDeleteInfoWithDefaults() *ModelLibraryDeleteInfo {
	this := ModelLibraryDeleteInfo{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *ModelLibraryDeleteInfo) GetPaths() []string {
	if o == nil || IsNil(o.Paths) {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLibraryDeleteInfo) GetPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.Paths) {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *ModelLibraryDeleteInfo) HasPaths() bool {
	if o != nil && !IsNil(o.Paths) {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *ModelLibraryDeleteInfo) SetPaths(v []string) {
	o.Paths = v
}

func (o ModelLibraryDeleteInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLibraryDeleteInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paths) {
		toSerialize["Paths"] = o.Paths
	}
	return toSerialize, nil
}

type NullableModelLibraryDeleteInfo struct {
	value *ModelLibraryDeleteInfo
	isSet bool
}

func (v NullableModelLibraryDeleteInfo) Get() *ModelLibraryDeleteInfo {
	return v.value
}

func (v *NullableModelLibraryDeleteInfo) Set(val *ModelLibraryDeleteInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLibraryDeleteInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLibraryDeleteInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLibraryDeleteInfo(val *ModelLibraryDeleteInfo) *NullableModelLibraryDeleteInfo {
	return &NullableModelLibraryDeleteInfo{value: val, isSet: true}
}

func (v NullableModelLibraryDeleteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLibraryDeleteInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

