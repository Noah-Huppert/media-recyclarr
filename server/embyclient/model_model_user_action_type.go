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

// ModelUserActionType the model 'ModelUserActionType'
type ModelUserActionType string

// List of UserActionType
const (
	MODELUSERACTIONTYPE_PLAYED_ITEM ModelUserActionType = "PlayedItem"
	MODELUSERACTIONTYPE_MARK_PLAYED ModelUserActionType = "MarkPlayed"
	MODELUSERACTIONTYPE_MARK_FAVORITE ModelUserActionType = "MarkFavorite"
)

// All allowed values of ModelUserActionType enum
var AllowedModelUserActionTypeEnumValues = []ModelUserActionType{
	"PlayedItem",
	"MarkPlayed",
	"MarkFavorite",
}

func (v *ModelUserActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelUserActionType(value)
	for _, existing := range AllowedModelUserActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelUserActionType", value)
}

// NewModelUserActionTypeFromValue returns a pointer to a valid ModelUserActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelUserActionTypeFromValue(v string) (*ModelUserActionType, error) {
	ev := ModelUserActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelUserActionType: valid values are %v", v, AllowedModelUserActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelUserActionType) IsValid() bool {
	for _, existing := range AllowedModelUserActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserActionType value
func (v ModelUserActionType) Ptr() *ModelUserActionType {
	return &v
}

type NullableModelUserActionType struct {
	value *ModelUserActionType
	isSet bool
}

func (v NullableModelUserActionType) Get() *ModelUserActionType {
	return v.value
}

func (v *NullableModelUserActionType) Set(val *ModelUserActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUserActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUserActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUserActionType(val *ModelUserActionType) *NullableModelUserActionType {
	return &NullableModelUserActionType{value: val, isSet: true}
}

func (v NullableModelUserActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUserActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
