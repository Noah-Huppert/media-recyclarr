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

// ModelExtendedVideoSubTypes the model 'ModelExtendedVideoSubTypes'
type ModelExtendedVideoSubTypes string

// List of ExtendedVideoSubTypes
const (
	MODELEXTENDEDVIDEOSUBTYPES_NONE ModelExtendedVideoSubTypes = "None"
	MODELEXTENDEDVIDEOSUBTYPES_HDR10 ModelExtendedVideoSubTypes = "Hdr10"
	MODELEXTENDEDVIDEOSUBTYPES_HYPER_LOG_GAMMA ModelExtendedVideoSubTypes = "HyperLogGamma"
	MODELEXTENDEDVIDEOSUBTYPES_HDR10_PLUS0 ModelExtendedVideoSubTypes = "Hdr10Plus0"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE02 ModelExtendedVideoSubTypes = "DoviProfile02"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE10 ModelExtendedVideoSubTypes = "DoviProfile10"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE22 ModelExtendedVideoSubTypes = "DoviProfile22"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE30 ModelExtendedVideoSubTypes = "DoviProfile30"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE42 ModelExtendedVideoSubTypes = "DoviProfile42"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE50 ModelExtendedVideoSubTypes = "DoviProfile50"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE61 ModelExtendedVideoSubTypes = "DoviProfile61"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE76 ModelExtendedVideoSubTypes = "DoviProfile76"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE81 ModelExtendedVideoSubTypes = "DoviProfile81"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE82 ModelExtendedVideoSubTypes = "DoviProfile82"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE83 ModelExtendedVideoSubTypes = "DoviProfile83"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE84 ModelExtendedVideoSubTypes = "DoviProfile84"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE85 ModelExtendedVideoSubTypes = "DoviProfile85"
	MODELEXTENDEDVIDEOSUBTYPES_DOVI_PROFILE92 ModelExtendedVideoSubTypes = "DoviProfile92"
)

// All allowed values of ModelExtendedVideoSubTypes enum
var AllowedModelExtendedVideoSubTypesEnumValues = []ModelExtendedVideoSubTypes{
	"None",
	"Hdr10",
	"HyperLogGamma",
	"Hdr10Plus0",
	"DoviProfile02",
	"DoviProfile10",
	"DoviProfile22",
	"DoviProfile30",
	"DoviProfile42",
	"DoviProfile50",
	"DoviProfile61",
	"DoviProfile76",
	"DoviProfile81",
	"DoviProfile82",
	"DoviProfile83",
	"DoviProfile84",
	"DoviProfile85",
	"DoviProfile92",
}

func (v *ModelExtendedVideoSubTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelExtendedVideoSubTypes(value)
	for _, existing := range AllowedModelExtendedVideoSubTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelExtendedVideoSubTypes", value)
}

// NewModelExtendedVideoSubTypesFromValue returns a pointer to a valid ModelExtendedVideoSubTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelExtendedVideoSubTypesFromValue(v string) (*ModelExtendedVideoSubTypes, error) {
	ev := ModelExtendedVideoSubTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelExtendedVideoSubTypes: valid values are %v", v, AllowedModelExtendedVideoSubTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelExtendedVideoSubTypes) IsValid() bool {
	for _, existing := range AllowedModelExtendedVideoSubTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExtendedVideoSubTypes value
func (v ModelExtendedVideoSubTypes) Ptr() *ModelExtendedVideoSubTypes {
	return &v
}

type NullableModelExtendedVideoSubTypes struct {
	value *ModelExtendedVideoSubTypes
	isSet bool
}

func (v NullableModelExtendedVideoSubTypes) Get() *ModelExtendedVideoSubTypes {
	return v.value
}

func (v *NullableModelExtendedVideoSubTypes) Set(val *ModelExtendedVideoSubTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableModelExtendedVideoSubTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableModelExtendedVideoSubTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelExtendedVideoSubTypes(val *ModelExtendedVideoSubTypes) *NullableModelExtendedVideoSubTypes {
	return &NullableModelExtendedVideoSubTypes{value: val, isSet: true}
}

func (v NullableModelExtendedVideoSubTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelExtendedVideoSubTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

