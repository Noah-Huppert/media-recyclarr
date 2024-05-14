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

// ModelTransportStreamTimestamp the model 'ModelTransportStreamTimestamp'
type ModelTransportStreamTimestamp string

// List of TransportStreamTimestamp
const (
	MODELTRANSPORTSTREAMTIMESTAMP_NONE ModelTransportStreamTimestamp = "None"
	MODELTRANSPORTSTREAMTIMESTAMP_ZERO ModelTransportStreamTimestamp = "Zero"
	MODELTRANSPORTSTREAMTIMESTAMP_VALID ModelTransportStreamTimestamp = "Valid"
)

// All allowed values of ModelTransportStreamTimestamp enum
var AllowedModelTransportStreamTimestampEnumValues = []ModelTransportStreamTimestamp{
	"None",
	"Zero",
	"Valid",
}

func (v *ModelTransportStreamTimestamp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelTransportStreamTimestamp(value)
	for _, existing := range AllowedModelTransportStreamTimestampEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelTransportStreamTimestamp", value)
}

// NewModelTransportStreamTimestampFromValue returns a pointer to a valid ModelTransportStreamTimestamp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelTransportStreamTimestampFromValue(v string) (*ModelTransportStreamTimestamp, error) {
	ev := ModelTransportStreamTimestamp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelTransportStreamTimestamp: valid values are %v", v, AllowedModelTransportStreamTimestampEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelTransportStreamTimestamp) IsValid() bool {
	for _, existing := range AllowedModelTransportStreamTimestampEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransportStreamTimestamp value
func (v ModelTransportStreamTimestamp) Ptr() *ModelTransportStreamTimestamp {
	return &v
}

type NullableModelTransportStreamTimestamp struct {
	value *ModelTransportStreamTimestamp
	isSet bool
}

func (v NullableModelTransportStreamTimestamp) Get() *ModelTransportStreamTimestamp {
	return v.value
}

func (v *NullableModelTransportStreamTimestamp) Set(val *ModelTransportStreamTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelTransportStreamTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelTransportStreamTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelTransportStreamTimestamp(val *ModelTransportStreamTimestamp) *NullableModelTransportStreamTimestamp {
	return &NullableModelTransportStreamTimestamp{value: val, isSet: true}
}

func (v NullableModelTransportStreamTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelTransportStreamTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
