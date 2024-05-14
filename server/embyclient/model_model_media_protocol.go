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

// ModelMediaProtocol the model 'ModelMediaProtocol'
type ModelMediaProtocol string

// List of MediaProtocol
const (
	MODELMEDIAPROTOCOL_FILE ModelMediaProtocol = "File"
	MODELMEDIAPROTOCOL_HTTP ModelMediaProtocol = "Http"
	MODELMEDIAPROTOCOL_RTMP ModelMediaProtocol = "Rtmp"
	MODELMEDIAPROTOCOL_RTSP ModelMediaProtocol = "Rtsp"
	MODELMEDIAPROTOCOL_UDP ModelMediaProtocol = "Udp"
	MODELMEDIAPROTOCOL_RTP ModelMediaProtocol = "Rtp"
	MODELMEDIAPROTOCOL_FTP ModelMediaProtocol = "Ftp"
	MODELMEDIAPROTOCOL_MMS ModelMediaProtocol = "Mms"
)

// All allowed values of ModelMediaProtocol enum
var AllowedModelMediaProtocolEnumValues = []ModelMediaProtocol{
	"File",
	"Http",
	"Rtmp",
	"Rtsp",
	"Udp",
	"Rtp",
	"Ftp",
	"Mms",
}

func (v *ModelMediaProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelMediaProtocol(value)
	for _, existing := range AllowedModelMediaProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelMediaProtocol", value)
}

// NewModelMediaProtocolFromValue returns a pointer to a valid ModelMediaProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelMediaProtocolFromValue(v string) (*ModelMediaProtocol, error) {
	ev := ModelMediaProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelMediaProtocol: valid values are %v", v, AllowedModelMediaProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelMediaProtocol) IsValid() bool {
	for _, existing := range AllowedModelMediaProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaProtocol value
func (v ModelMediaProtocol) Ptr() *ModelMediaProtocol {
	return &v
}

type NullableModelMediaProtocol struct {
	value *ModelMediaProtocol
	isSet bool
}

func (v NullableModelMediaProtocol) Get() *ModelMediaProtocol {
	return v.value
}

func (v *NullableModelMediaProtocol) Set(val *ModelMediaProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableModelMediaProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableModelMediaProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelMediaProtocol(val *ModelMediaProtocol) *NullableModelMediaProtocol {
	return &NullableModelMediaProtocol{value: val, isSet: true}
}

func (v NullableModelMediaProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelMediaProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
