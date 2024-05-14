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

// checks if the ModelLibraryPostUpdatedMedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLibraryPostUpdatedMedia{}

// ModelLibraryPostUpdatedMedia struct for ModelLibraryPostUpdatedMedia
type ModelLibraryPostUpdatedMedia struct {
	Updates []ModelLibraryMediaUpdateInfo `json:"Updates,omitempty"`
}

// NewModelLibraryPostUpdatedMedia instantiates a new ModelLibraryPostUpdatedMedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLibraryPostUpdatedMedia() *ModelLibraryPostUpdatedMedia {
	this := ModelLibraryPostUpdatedMedia{}
	return &this
}

// NewModelLibraryPostUpdatedMediaWithDefaults instantiates a new ModelLibraryPostUpdatedMedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLibraryPostUpdatedMediaWithDefaults() *ModelLibraryPostUpdatedMedia {
	this := ModelLibraryPostUpdatedMedia{}
	return &this
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *ModelLibraryPostUpdatedMedia) GetUpdates() []ModelLibraryMediaUpdateInfo {
	if o == nil || IsNil(o.Updates) {
		var ret []ModelLibraryMediaUpdateInfo
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLibraryPostUpdatedMedia) GetUpdatesOk() ([]ModelLibraryMediaUpdateInfo, bool) {
	if o == nil || IsNil(o.Updates) {
		return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *ModelLibraryPostUpdatedMedia) HasUpdates() bool {
	if o != nil && !IsNil(o.Updates) {
		return true
	}

	return false
}

// SetUpdates gets a reference to the given []ModelLibraryMediaUpdateInfo and assigns it to the Updates field.
func (o *ModelLibraryPostUpdatedMedia) SetUpdates(v []ModelLibraryMediaUpdateInfo) {
	o.Updates = v
}

func (o ModelLibraryPostUpdatedMedia) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLibraryPostUpdatedMedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Updates) {
		toSerialize["Updates"] = o.Updates
	}
	return toSerialize, nil
}

type NullableModelLibraryPostUpdatedMedia struct {
	value *ModelLibraryPostUpdatedMedia
	isSet bool
}

func (v NullableModelLibraryPostUpdatedMedia) Get() *ModelLibraryPostUpdatedMedia {
	return v.value
}

func (v *NullableModelLibraryPostUpdatedMedia) Set(val *ModelLibraryPostUpdatedMedia) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLibraryPostUpdatedMedia) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLibraryPostUpdatedMedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLibraryPostUpdatedMedia(val *ModelLibraryPostUpdatedMedia) *NullableModelLibraryPostUpdatedMedia {
	return &NullableModelLibraryPostUpdatedMedia{value: val, isSet: true}
}

func (v NullableModelLibraryPostUpdatedMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLibraryPostUpdatedMedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

