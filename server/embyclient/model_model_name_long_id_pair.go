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

// checks if the ModelNameLongIdPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelNameLongIdPair{}

// ModelNameLongIdPair struct for ModelNameLongIdPair
type ModelNameLongIdPair struct {
	Name *string `json:"Name,omitempty"`
	Id *int64 `json:"Id,omitempty"`
}

// NewModelNameLongIdPair instantiates a new ModelNameLongIdPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelNameLongIdPair() *ModelNameLongIdPair {
	this := ModelNameLongIdPair{}
	return &this
}

// NewModelNameLongIdPairWithDefaults instantiates a new ModelNameLongIdPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelNameLongIdPairWithDefaults() *ModelNameLongIdPair {
	this := ModelNameLongIdPair{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelNameLongIdPair) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelNameLongIdPair) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelNameLongIdPair) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelNameLongIdPair) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelNameLongIdPair) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelNameLongIdPair) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelNameLongIdPair) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ModelNameLongIdPair) SetId(v int64) {
	o.Id = &v
}

func (o ModelNameLongIdPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelNameLongIdPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	return toSerialize, nil
}

type NullableModelNameLongIdPair struct {
	value *ModelNameLongIdPair
	isSet bool
}

func (v NullableModelNameLongIdPair) Get() *ModelNameLongIdPair {
	return v.value
}

func (v *NullableModelNameLongIdPair) Set(val *ModelNameLongIdPair) {
	v.value = val
	v.isSet = true
}

func (v NullableModelNameLongIdPair) IsSet() bool {
	return v.isSet
}

func (v *NullableModelNameLongIdPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelNameLongIdPair(val *ModelNameLongIdPair) *NullableModelNameLongIdPair {
	return &NullableModelNameLongIdPair{value: val, isSet: true}
}

func (v NullableModelNameLongIdPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelNameLongIdPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


