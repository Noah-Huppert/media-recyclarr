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

// checks if the ModelLibraryMediaUpdateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLibraryMediaUpdateInfo{}

// ModelLibraryMediaUpdateInfo struct for ModelLibraryMediaUpdateInfo
type ModelLibraryMediaUpdateInfo struct {
	Path *string `json:"Path,omitempty"`
	UpdateType *string `json:"UpdateType,omitempty"`
}

// NewModelLibraryMediaUpdateInfo instantiates a new ModelLibraryMediaUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLibraryMediaUpdateInfo() *ModelLibraryMediaUpdateInfo {
	this := ModelLibraryMediaUpdateInfo{}
	return &this
}

// NewModelLibraryMediaUpdateInfoWithDefaults instantiates a new ModelLibraryMediaUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLibraryMediaUpdateInfoWithDefaults() *ModelLibraryMediaUpdateInfo {
	this := ModelLibraryMediaUpdateInfo{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ModelLibraryMediaUpdateInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLibraryMediaUpdateInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ModelLibraryMediaUpdateInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ModelLibraryMediaUpdateInfo) SetPath(v string) {
	o.Path = &v
}

// GetUpdateType returns the UpdateType field value if set, zero value otherwise.
func (o *ModelLibraryMediaUpdateInfo) GetUpdateType() string {
	if o == nil || IsNil(o.UpdateType) {
		var ret string
		return ret
	}
	return *o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLibraryMediaUpdateInfo) GetUpdateTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateType) {
		return nil, false
	}
	return o.UpdateType, true
}

// HasUpdateType returns a boolean if a field has been set.
func (o *ModelLibraryMediaUpdateInfo) HasUpdateType() bool {
	if o != nil && !IsNil(o.UpdateType) {
		return true
	}

	return false
}

// SetUpdateType gets a reference to the given string and assigns it to the UpdateType field.
func (o *ModelLibraryMediaUpdateInfo) SetUpdateType(v string) {
	o.UpdateType = &v
}

func (o ModelLibraryMediaUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLibraryMediaUpdateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.UpdateType) {
		toSerialize["UpdateType"] = o.UpdateType
	}
	return toSerialize, nil
}

type NullableModelLibraryMediaUpdateInfo struct {
	value *ModelLibraryMediaUpdateInfo
	isSet bool
}

func (v NullableModelLibraryMediaUpdateInfo) Get() *ModelLibraryMediaUpdateInfo {
	return v.value
}

func (v *NullableModelLibraryMediaUpdateInfo) Set(val *ModelLibraryMediaUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLibraryMediaUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLibraryMediaUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLibraryMediaUpdateInfo(val *ModelLibraryMediaUpdateInfo) *NullableModelLibraryMediaUpdateInfo {
	return &NullableModelLibraryMediaUpdateInfo{value: val, isSet: true}
}

func (v NullableModelLibraryMediaUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLibraryMediaUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


