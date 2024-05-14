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

// ModelSubtitleDeliveryMethod the model 'ModelSubtitleDeliveryMethod'
type ModelSubtitleDeliveryMethod string

// List of SubtitleDeliveryMethod
const (
	MODELSUBTITLEDELIVERYMETHOD_ENCODE ModelSubtitleDeliveryMethod = "Encode"
	MODELSUBTITLEDELIVERYMETHOD_EMBED ModelSubtitleDeliveryMethod = "Embed"
	MODELSUBTITLEDELIVERYMETHOD_EXTERNAL ModelSubtitleDeliveryMethod = "External"
	MODELSUBTITLEDELIVERYMETHOD_HLS ModelSubtitleDeliveryMethod = "Hls"
	MODELSUBTITLEDELIVERYMETHOD_VIDEO_SIDE_DATA ModelSubtitleDeliveryMethod = "VideoSideData"
)

// All allowed values of ModelSubtitleDeliveryMethod enum
var AllowedModelSubtitleDeliveryMethodEnumValues = []ModelSubtitleDeliveryMethod{
	"Encode",
	"Embed",
	"External",
	"Hls",
	"VideoSideData",
}

func (v *ModelSubtitleDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelSubtitleDeliveryMethod(value)
	for _, existing := range AllowedModelSubtitleDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelSubtitleDeliveryMethod", value)
}

// NewModelSubtitleDeliveryMethodFromValue returns a pointer to a valid ModelSubtitleDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelSubtitleDeliveryMethodFromValue(v string) (*ModelSubtitleDeliveryMethod, error) {
	ev := ModelSubtitleDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelSubtitleDeliveryMethod: valid values are %v", v, AllowedModelSubtitleDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelSubtitleDeliveryMethod) IsValid() bool {
	for _, existing := range AllowedModelSubtitleDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubtitleDeliveryMethod value
func (v ModelSubtitleDeliveryMethod) Ptr() *ModelSubtitleDeliveryMethod {
	return &v
}

type NullableModelSubtitleDeliveryMethod struct {
	value *ModelSubtitleDeliveryMethod
	isSet bool
}

func (v NullableModelSubtitleDeliveryMethod) Get() *ModelSubtitleDeliveryMethod {
	return v.value
}

func (v *NullableModelSubtitleDeliveryMethod) Set(val *ModelSubtitleDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSubtitleDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSubtitleDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSubtitleDeliveryMethod(val *ModelSubtitleDeliveryMethod) *NullableModelSubtitleDeliveryMethod {
	return &NullableModelSubtitleDeliveryMethod{value: val, isSet: true}
}

func (v NullableModelSubtitleDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSubtitleDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
