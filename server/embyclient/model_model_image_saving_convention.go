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

// ModelImageSavingConvention the model 'ModelImageSavingConvention'
type ModelImageSavingConvention string

// List of ImageSavingConvention
const (
	MODELIMAGESAVINGCONVENTION_LEGACY ModelImageSavingConvention = "Legacy"
	MODELIMAGESAVINGCONVENTION_COMPATIBLE ModelImageSavingConvention = "Compatible"
)

// All allowed values of ModelImageSavingConvention enum
var AllowedModelImageSavingConventionEnumValues = []ModelImageSavingConvention{
	"Legacy",
	"Compatible",
}

func (v *ModelImageSavingConvention) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelImageSavingConvention(value)
	for _, existing := range AllowedModelImageSavingConventionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelImageSavingConvention", value)
}

// NewModelImageSavingConventionFromValue returns a pointer to a valid ModelImageSavingConvention
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelImageSavingConventionFromValue(v string) (*ModelImageSavingConvention, error) {
	ev := ModelImageSavingConvention(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelImageSavingConvention: valid values are %v", v, AllowedModelImageSavingConventionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelImageSavingConvention) IsValid() bool {
	for _, existing := range AllowedModelImageSavingConventionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageSavingConvention value
func (v ModelImageSavingConvention) Ptr() *ModelImageSavingConvention {
	return &v
}

type NullableModelImageSavingConvention struct {
	value *ModelImageSavingConvention
	isSet bool
}

func (v NullableModelImageSavingConvention) Get() *ModelImageSavingConvention {
	return v.value
}

func (v *NullableModelImageSavingConvention) Set(val *ModelImageSavingConvention) {
	v.value = val
	v.isSet = true
}

func (v NullableModelImageSavingConvention) IsSet() bool {
	return v.isSet
}

func (v *NullableModelImageSavingConvention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelImageSavingConvention(val *ModelImageSavingConvention) *NullableModelImageSavingConvention {
	return &NullableModelImageSavingConvention{value: val, isSet: true}
}

func (v NullableModelImageSavingConvention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelImageSavingConvention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
