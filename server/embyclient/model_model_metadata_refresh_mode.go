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

// ModelMetadataRefreshMode the model 'ModelMetadataRefreshMode'
type ModelMetadataRefreshMode string

// List of MetadataRefreshMode
const (
	MODELMETADATAREFRESHMODE_VALIDATION_ONLY ModelMetadataRefreshMode = "ValidationOnly"
	MODELMETADATAREFRESHMODE_DEFAULT ModelMetadataRefreshMode = "Default"
	MODELMETADATAREFRESHMODE_FULL_REFRESH ModelMetadataRefreshMode = "FullRefresh"
)

// All allowed values of ModelMetadataRefreshMode enum
var AllowedModelMetadataRefreshModeEnumValues = []ModelMetadataRefreshMode{
	"ValidationOnly",
	"Default",
	"FullRefresh",
}

func (v *ModelMetadataRefreshMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelMetadataRefreshMode(value)
	for _, existing := range AllowedModelMetadataRefreshModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelMetadataRefreshMode", value)
}

// NewModelMetadataRefreshModeFromValue returns a pointer to a valid ModelMetadataRefreshMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelMetadataRefreshModeFromValue(v string) (*ModelMetadataRefreshMode, error) {
	ev := ModelMetadataRefreshMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelMetadataRefreshMode: valid values are %v", v, AllowedModelMetadataRefreshModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelMetadataRefreshMode) IsValid() bool {
	for _, existing := range AllowedModelMetadataRefreshModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataRefreshMode value
func (v ModelMetadataRefreshMode) Ptr() *ModelMetadataRefreshMode {
	return &v
}

type NullableModelMetadataRefreshMode struct {
	value *ModelMetadataRefreshMode
	isSet bool
}

func (v NullableModelMetadataRefreshMode) Get() *ModelMetadataRefreshMode {
	return v.value
}

func (v *NullableModelMetadataRefreshMode) Set(val *ModelMetadataRefreshMode) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMetadataRefreshMode) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMetadataRefreshMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMetadataRefreshMode(val *ModelMetadataRefreshMode) *NullableModelMetadataRefreshMode {
	return &NullableModelMetadataRefreshMode{value: val, isSet: true}
}

func (v NullableModelMetadataRefreshMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMetadataRefreshMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
