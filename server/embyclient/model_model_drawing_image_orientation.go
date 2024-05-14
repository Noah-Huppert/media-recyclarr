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

// ModelDrawingImageOrientation the model 'ModelDrawingImageOrientation'
type ModelDrawingImageOrientation string

// List of Drawing.ImageOrientation
const (
	MODELDRAWINGIMAGEORIENTATION_TOP_LEFT ModelDrawingImageOrientation = "TopLeft"
	MODELDRAWINGIMAGEORIENTATION_TOP_RIGHT ModelDrawingImageOrientation = "TopRight"
	MODELDRAWINGIMAGEORIENTATION_BOTTOM_RIGHT ModelDrawingImageOrientation = "BottomRight"
	MODELDRAWINGIMAGEORIENTATION_BOTTOM_LEFT ModelDrawingImageOrientation = "BottomLeft"
	MODELDRAWINGIMAGEORIENTATION_LEFT_TOP ModelDrawingImageOrientation = "LeftTop"
	MODELDRAWINGIMAGEORIENTATION_RIGHT_TOP ModelDrawingImageOrientation = "RightTop"
	MODELDRAWINGIMAGEORIENTATION_RIGHT_BOTTOM ModelDrawingImageOrientation = "RightBottom"
	MODELDRAWINGIMAGEORIENTATION_LEFT_BOTTOM ModelDrawingImageOrientation = "LeftBottom"
)

// All allowed values of ModelDrawingImageOrientation enum
var AllowedModelDrawingImageOrientationEnumValues = []ModelDrawingImageOrientation{
	"TopLeft",
	"TopRight",
	"BottomRight",
	"BottomLeft",
	"LeftTop",
	"RightTop",
	"RightBottom",
	"LeftBottom",
}

func (v *ModelDrawingImageOrientation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelDrawingImageOrientation(value)
	for _, existing := range AllowedModelDrawingImageOrientationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelDrawingImageOrientation", value)
}

// NewModelDrawingImageOrientationFromValue returns a pointer to a valid ModelDrawingImageOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelDrawingImageOrientationFromValue(v string) (*ModelDrawingImageOrientation, error) {
	ev := ModelDrawingImageOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelDrawingImageOrientation: valid values are %v", v, AllowedModelDrawingImageOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelDrawingImageOrientation) IsValid() bool {
	for _, existing := range AllowedModelDrawingImageOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Drawing.ImageOrientation value
func (v ModelDrawingImageOrientation) Ptr() *ModelDrawingImageOrientation {
	return &v
}

type NullableModelDrawingImageOrientation struct {
	value *ModelDrawingImageOrientation
	isSet bool
}

func (v NullableModelDrawingImageOrientation) Get() *ModelDrawingImageOrientation {
	return v.value
}

func (v *NullableModelDrawingImageOrientation) Set(val *ModelDrawingImageOrientation) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDrawingImageOrientation) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDrawingImageOrientation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDrawingImageOrientation(val *ModelDrawingImageOrientation) *NullableModelDrawingImageOrientation {
	return &NullableModelDrawingImageOrientation{value: val, isSet: true}
}

func (v NullableModelDrawingImageOrientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDrawingImageOrientation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
