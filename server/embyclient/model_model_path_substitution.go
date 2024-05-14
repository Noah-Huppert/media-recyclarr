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

// checks if the ModelPathSubstitution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPathSubstitution{}

// ModelPathSubstitution struct for ModelPathSubstitution
type ModelPathSubstitution struct {
	From *string `json:"From,omitempty"`
	To *string `json:"To,omitempty"`
}

// NewModelPathSubstitution instantiates a new ModelPathSubstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPathSubstitution() *ModelPathSubstitution {
	this := ModelPathSubstitution{}
	return &this
}

// NewModelPathSubstitutionWithDefaults instantiates a new ModelPathSubstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPathSubstitutionWithDefaults() *ModelPathSubstitution {
	this := ModelPathSubstitution{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ModelPathSubstitution) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPathSubstitution) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ModelPathSubstitution) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *ModelPathSubstitution) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ModelPathSubstitution) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPathSubstitution) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ModelPathSubstitution) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *ModelPathSubstitution) SetTo(v string) {
	o.To = &v
}

func (o ModelPathSubstitution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPathSubstitution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["From"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["To"] = o.To
	}
	return toSerialize, nil
}

type NullableModelPathSubstitution struct {
	value *ModelPathSubstitution
	isSet bool
}

func (v NullableModelPathSubstitution) Get() *ModelPathSubstitution {
	return v.value
}

func (v *NullableModelPathSubstitution) Set(val *ModelPathSubstitution) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPathSubstitution) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPathSubstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPathSubstitution(val *ModelPathSubstitution) *NullableModelPathSubstitution {
	return &NullableModelPathSubstitution{value: val, isSet: true}
}

func (v NullableModelPathSubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPathSubstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

