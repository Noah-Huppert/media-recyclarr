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

// ModelPlaybackErrorCode the model 'ModelPlaybackErrorCode'
type ModelPlaybackErrorCode string

// List of PlaybackErrorCode
const (
	MODELPLAYBACKERRORCODE_NOT_ALLOWED ModelPlaybackErrorCode = "NotAllowed"
	MODELPLAYBACKERRORCODE_NO_COMPATIBLE_STREAM ModelPlaybackErrorCode = "NoCompatibleStream"
	MODELPLAYBACKERRORCODE_RATE_LIMIT_EXCEEDED ModelPlaybackErrorCode = "RateLimitExceeded"
)

// All allowed values of ModelPlaybackErrorCode enum
var AllowedModelPlaybackErrorCodeEnumValues = []ModelPlaybackErrorCode{
	"NotAllowed",
	"NoCompatibleStream",
	"RateLimitExceeded",
}

func (v *ModelPlaybackErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelPlaybackErrorCode(value)
	for _, existing := range AllowedModelPlaybackErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelPlaybackErrorCode", value)
}

// NewModelPlaybackErrorCodeFromValue returns a pointer to a valid ModelPlaybackErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelPlaybackErrorCodeFromValue(v string) (*ModelPlaybackErrorCode, error) {
	ev := ModelPlaybackErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelPlaybackErrorCode: valid values are %v", v, AllowedModelPlaybackErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelPlaybackErrorCode) IsValid() bool {
	for _, existing := range AllowedModelPlaybackErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaybackErrorCode value
func (v ModelPlaybackErrorCode) Ptr() *ModelPlaybackErrorCode {
	return &v
}

type NullableModelPlaybackErrorCode struct {
	value *ModelPlaybackErrorCode
	isSet bool
}

func (v NullableModelPlaybackErrorCode) Get() *ModelPlaybackErrorCode {
	return v.value
}

func (v *NullableModelPlaybackErrorCode) Set(val *ModelPlaybackErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPlaybackErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPlaybackErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPlaybackErrorCode(val *ModelPlaybackErrorCode) *NullableModelPlaybackErrorCode {
	return &NullableModelPlaybackErrorCode{value: val, isSet: true}
}

func (v NullableModelPlaybackErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPlaybackErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

