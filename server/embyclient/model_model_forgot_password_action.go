/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
	"fmt"
)

// ModelForgotPasswordAction the model 'ModelForgotPasswordAction'
type ModelForgotPasswordAction string

// List of ForgotPasswordAction
const (
	MODELFORGOTPASSWORDACTION_CONTACT_ADMIN ModelForgotPasswordAction = "ContactAdmin"
	MODELFORGOTPASSWORDACTION_PIN_CODE ModelForgotPasswordAction = "PinCode"
)

// All allowed values of ModelForgotPasswordAction enum
var AllowedModelForgotPasswordActionEnumValues = []ModelForgotPasswordAction{
	"ContactAdmin",
	"PinCode",
}

func (v *ModelForgotPasswordAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelForgotPasswordAction(value)
	for _, existing := range AllowedModelForgotPasswordActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelForgotPasswordAction", value)
}

// NewModelForgotPasswordActionFromValue returns a pointer to a valid ModelForgotPasswordAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelForgotPasswordActionFromValue(v string) (*ModelForgotPasswordAction, error) {
	ev := ModelForgotPasswordAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelForgotPasswordAction: valid values are %v", v, AllowedModelForgotPasswordActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelForgotPasswordAction) IsValid() bool {
	for _, existing := range AllowedModelForgotPasswordActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ForgotPasswordAction value
func (v ModelForgotPasswordAction) Ptr() *ModelForgotPasswordAction {
	return &v
}

type NullableModelForgotPasswordAction struct {
	value *ModelForgotPasswordAction
	isSet bool
}

func (v NullableModelForgotPasswordAction) Get() *ModelForgotPasswordAction {
	return v.value
}

func (v *NullableModelForgotPasswordAction) Set(val *ModelForgotPasswordAction) {
	v.value = val
	v.isSet = true
}

func (v NullableModelForgotPasswordAction) IsSet() bool {
	return v.isSet
}

func (v *NullableModelForgotPasswordAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelForgotPasswordAction(val *ModelForgotPasswordAction) *NullableModelForgotPasswordAction {
	return &NullableModelForgotPasswordAction{value: val, isSet: true}
}

func (v NullableModelForgotPasswordAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelForgotPasswordAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
