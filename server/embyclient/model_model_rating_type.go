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

// ModelRatingType the model 'ModelRatingType'
type ModelRatingType string

// List of RatingType
const (
	MODELRATINGTYPE_SCORE ModelRatingType = "Score"
	MODELRATINGTYPE_LIKES ModelRatingType = "Likes"
)

// All allowed values of ModelRatingType enum
var AllowedModelRatingTypeEnumValues = []ModelRatingType{
	"Score",
	"Likes",
}

func (v *ModelRatingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelRatingType(value)
	for _, existing := range AllowedModelRatingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelRatingType", value)
}

// NewModelRatingTypeFromValue returns a pointer to a valid ModelRatingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelRatingTypeFromValue(v string) (*ModelRatingType, error) {
	ev := ModelRatingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelRatingType: valid values are %v", v, AllowedModelRatingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelRatingType) IsValid() bool {
	for _, existing := range AllowedModelRatingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RatingType value
func (v ModelRatingType) Ptr() *ModelRatingType {
	return &v
}

type NullableModelRatingType struct {
	value *ModelRatingType
	isSet bool
}

func (v NullableModelRatingType) Get() *ModelRatingType {
	return v.value
}

func (v *NullableModelRatingType) Set(val *ModelRatingType) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRatingType) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRatingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRatingType(val *ModelRatingType) *NullableModelRatingType {
	return &NullableModelRatingType{value: val, isSet: true}
}

func (v NullableModelRatingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRatingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
