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

// checks if the ModelMediaUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelMediaUrl{}

// ModelMediaUrl struct for ModelMediaUrl
type ModelMediaUrl struct {
	Url *string `json:"Url,omitempty"`
	Name *string `json:"Name,omitempty"`
}

// NewModelMediaUrl instantiates a new ModelMediaUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelMediaUrl() *ModelMediaUrl {
	this := ModelMediaUrl{}
	return &this
}

// NewModelMediaUrlWithDefaults instantiates a new ModelMediaUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelMediaUrlWithDefaults() *ModelMediaUrl {
	this := ModelMediaUrl{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ModelMediaUrl) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMediaUrl) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ModelMediaUrl) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ModelMediaUrl) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelMediaUrl) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelMediaUrl) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelMediaUrl) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelMediaUrl) SetName(v string) {
	o.Name = &v
}

func (o ModelMediaUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelMediaUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["Url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableModelMediaUrl struct {
	value *ModelMediaUrl
	isSet bool
}

func (v NullableModelMediaUrl) Get() *ModelMediaUrl {
	return v.value
}

func (v *NullableModelMediaUrl) Set(val *ModelMediaUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMediaUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMediaUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMediaUrl(val *ModelMediaUrl) *NullableModelMediaUrl {
	return &NullableModelMediaUrl{value: val, isSet: true}
}

func (v NullableModelMediaUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMediaUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

