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

// checks if the ModelTupleDoubleDouble type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelTupleDoubleDouble{}

// ModelTupleDoubleDouble struct for ModelTupleDoubleDouble
type ModelTupleDoubleDouble struct {
	Item1 *float64 `json:"Item1,omitempty"`
	Item2 *float64 `json:"Item2,omitempty"`
}

// NewModelTupleDoubleDouble instantiates a new ModelTupleDoubleDouble object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelTupleDoubleDouble() *ModelTupleDoubleDouble {
	this := ModelTupleDoubleDouble{}
	return &this
}

// NewModelTupleDoubleDoubleWithDefaults instantiates a new ModelTupleDoubleDouble object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelTupleDoubleDoubleWithDefaults() *ModelTupleDoubleDouble {
	this := ModelTupleDoubleDouble{}
	return &this
}

// GetItem1 returns the Item1 field value if set, zero value otherwise.
func (o *ModelTupleDoubleDouble) GetItem1() float64 {
	if o == nil || IsNil(o.Item1) {
		var ret float64
		return ret
	}
	return *o.Item1
}

// GetItem1Ok returns a tuple with the Item1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTupleDoubleDouble) GetItem1Ok() (*float64, bool) {
	if o == nil || IsNil(o.Item1) {
		return nil, false
	}
	return o.Item1, true
}

// HasItem1 returns a boolean if a field has been set.
func (o *ModelTupleDoubleDouble) HasItem1() bool {
	if o != nil && !IsNil(o.Item1) {
		return true
	}

	return false
}

// SetItem1 gets a reference to the given float64 and assigns it to the Item1 field.
func (o *ModelTupleDoubleDouble) SetItem1(v float64) {
	o.Item1 = &v
}

// GetItem2 returns the Item2 field value if set, zero value otherwise.
func (o *ModelTupleDoubleDouble) GetItem2() float64 {
	if o == nil || IsNil(o.Item2) {
		var ret float64
		return ret
	}
	return *o.Item2
}

// GetItem2Ok returns a tuple with the Item2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTupleDoubleDouble) GetItem2Ok() (*float64, bool) {
	if o == nil || IsNil(o.Item2) {
		return nil, false
	}
	return o.Item2, true
}

// HasItem2 returns a boolean if a field has been set.
func (o *ModelTupleDoubleDouble) HasItem2() bool {
	if o != nil && !IsNil(o.Item2) {
		return true
	}

	return false
}

// SetItem2 gets a reference to the given float64 and assigns it to the Item2 field.
func (o *ModelTupleDoubleDouble) SetItem2(v float64) {
	o.Item2 = &v
}

func (o ModelTupleDoubleDouble) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelTupleDoubleDouble) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item1) {
		toSerialize["Item1"] = o.Item1
	}
	if !IsNil(o.Item2) {
		toSerialize["Item2"] = o.Item2
	}
	return toSerialize, nil
}

type NullableModelTupleDoubleDouble struct {
	value *ModelTupleDoubleDouble
	isSet bool
}

func (v NullableModelTupleDoubleDouble) Get() *ModelTupleDoubleDouble {
	return v.value
}

func (v *NullableModelTupleDoubleDouble) Set(val *ModelTupleDoubleDouble) {
	v.value = val
	v.isSet = true
}

func (v NullableModelTupleDoubleDouble) IsSet() bool {
	return v.isSet
}

func (v *NullableModelTupleDoubleDouble) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelTupleDoubleDouble(val *ModelTupleDoubleDouble) *NullableModelTupleDoubleDouble {
	return &NullableModelTupleDoubleDouble{value: val, isSet: true}
}

func (v NullableModelTupleDoubleDouble) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelTupleDoubleDouble) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


