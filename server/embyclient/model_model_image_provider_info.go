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

// checks if the ModelImageProviderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelImageProviderInfo{}

// ModelImageProviderInfo struct for ModelImageProviderInfo
type ModelImageProviderInfo struct {
	Name *string `json:"Name,omitempty"`
	SupportedImages []ModelImageType `json:"SupportedImages,omitempty"`
}

// NewModelImageProviderInfo instantiates a new ModelImageProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelImageProviderInfo() *ModelImageProviderInfo {
	this := ModelImageProviderInfo{}
	return &this
}

// NewModelImageProviderInfoWithDefaults instantiates a new ModelImageProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelImageProviderInfoWithDefaults() *ModelImageProviderInfo {
	this := ModelImageProviderInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelImageProviderInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageProviderInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelImageProviderInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelImageProviderInfo) SetName(v string) {
	o.Name = &v
}

// GetSupportedImages returns the SupportedImages field value if set, zero value otherwise.
func (o *ModelImageProviderInfo) GetSupportedImages() []ModelImageType {
	if o == nil || IsNil(o.SupportedImages) {
		var ret []ModelImageType
		return ret
	}
	return o.SupportedImages
}

// GetSupportedImagesOk returns a tuple with the SupportedImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageProviderInfo) GetSupportedImagesOk() ([]ModelImageType, bool) {
	if o == nil || IsNil(o.SupportedImages) {
		return nil, false
	}
	return o.SupportedImages, true
}

// HasSupportedImages returns a boolean if a field has been set.
func (o *ModelImageProviderInfo) HasSupportedImages() bool {
	if o != nil && !IsNil(o.SupportedImages) {
		return true
	}

	return false
}

// SetSupportedImages gets a reference to the given []ModelImageType and assigns it to the SupportedImages field.
func (o *ModelImageProviderInfo) SetSupportedImages(v []ModelImageType) {
	o.SupportedImages = v
}

func (o ModelImageProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelImageProviderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.SupportedImages) {
		toSerialize["SupportedImages"] = o.SupportedImages
	}
	return toSerialize, nil
}

type NullableModelImageProviderInfo struct {
	value *ModelImageProviderInfo
	isSet bool
}

func (v NullableModelImageProviderInfo) Get() *ModelImageProviderInfo {
	return v.value
}

func (v *NullableModelImageProviderInfo) Set(val *ModelImageProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelImageProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelImageProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelImageProviderInfo(val *ModelImageProviderInfo) *NullableModelImageProviderInfo {
	return &NullableModelImageProviderInfo{value: val, isSet: true}
}

func (v NullableModelImageProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelImageProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


