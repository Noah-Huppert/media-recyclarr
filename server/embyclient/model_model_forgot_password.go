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

// checks if the ModelForgotPassword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelForgotPassword{}

// ModelForgotPassword struct for ModelForgotPassword
type ModelForgotPassword struct {
	EnteredUsername *string `json:"EnteredUsername,omitempty"`
}

// NewModelForgotPassword instantiates a new ModelForgotPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelForgotPassword() *ModelForgotPassword {
	this := ModelForgotPassword{}
	return &this
}

// NewModelForgotPasswordWithDefaults instantiates a new ModelForgotPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelForgotPasswordWithDefaults() *ModelForgotPassword {
	this := ModelForgotPassword{}
	return &this
}

// GetEnteredUsername returns the EnteredUsername field value if set, zero value otherwise.
func (o *ModelForgotPassword) GetEnteredUsername() string {
	if o == nil || IsNil(o.EnteredUsername) {
		var ret string
		return ret
	}
	return *o.EnteredUsername
}

// GetEnteredUsernameOk returns a tuple with the EnteredUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelForgotPassword) GetEnteredUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.EnteredUsername) {
		return nil, false
	}
	return o.EnteredUsername, true
}

// HasEnteredUsername returns a boolean if a field has been set.
func (o *ModelForgotPassword) HasEnteredUsername() bool {
	if o != nil && !IsNil(o.EnteredUsername) {
		return true
	}

	return false
}

// SetEnteredUsername gets a reference to the given string and assigns it to the EnteredUsername field.
func (o *ModelForgotPassword) SetEnteredUsername(v string) {
	o.EnteredUsername = &v
}

func (o ModelForgotPassword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelForgotPassword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnteredUsername) {
		toSerialize["EnteredUsername"] = o.EnteredUsername
	}
	return toSerialize, nil
}

type NullableModelForgotPassword struct {
	value *ModelForgotPassword
	isSet bool
}

func (v NullableModelForgotPassword) Get() *ModelForgotPassword {
	return v.value
}

func (v *NullableModelForgotPassword) Set(val *ModelForgotPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableModelForgotPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableModelForgotPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelForgotPassword(val *ModelForgotPassword) *NullableModelForgotPassword {
	return &NullableModelForgotPassword{value: val, isSet: true}
}

func (v NullableModelForgotPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelForgotPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

