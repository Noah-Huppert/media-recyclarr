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

// checks if the ModelUserLibraryAddTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUserLibraryAddTags{}

// ModelUserLibraryAddTags struct for ModelUserLibraryAddTags
type ModelUserLibraryAddTags struct {
	Tags []ModelNameIdPair `json:"Tags,omitempty"`
}

// NewModelUserLibraryAddTags instantiates a new ModelUserLibraryAddTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUserLibraryAddTags() *ModelUserLibraryAddTags {
	this := ModelUserLibraryAddTags{}
	return &this
}

// NewModelUserLibraryAddTagsWithDefaults instantiates a new ModelUserLibraryAddTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUserLibraryAddTagsWithDefaults() *ModelUserLibraryAddTags {
	this := ModelUserLibraryAddTags{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModelUserLibraryAddTags) GetTags() []ModelNameIdPair {
	if o == nil || IsNil(o.Tags) {
		var ret []ModelNameIdPair
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserLibraryAddTags) GetTagsOk() ([]ModelNameIdPair, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModelUserLibraryAddTags) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ModelNameIdPair and assigns it to the Tags field.
func (o *ModelUserLibraryAddTags) SetTags(v []ModelNameIdPair) {
	o.Tags = v
}

func (o ModelUserLibraryAddTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUserLibraryAddTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["Tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableModelUserLibraryAddTags struct {
	value *ModelUserLibraryAddTags
	isSet bool
}

func (v NullableModelUserLibraryAddTags) Get() *ModelUserLibraryAddTags {
	return v.value
}

func (v *NullableModelUserLibraryAddTags) Set(val *ModelUserLibraryAddTags) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUserLibraryAddTags) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUserLibraryAddTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUserLibraryAddTags(val *ModelUserLibraryAddTags) *NullableModelUserLibraryAddTags {
	return &NullableModelUserLibraryAddTags{value: val, isSet: true}
}

func (v NullableModelUserLibraryAddTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUserLibraryAddTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

