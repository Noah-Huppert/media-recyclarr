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

// ModelUnratedItem the model 'ModelUnratedItem'
type ModelUnratedItem string

// List of UnratedItem
const (
	MODELUNRATEDITEM_MOVIE ModelUnratedItem = "Movie"
	MODELUNRATEDITEM_TRAILER ModelUnratedItem = "Trailer"
	MODELUNRATEDITEM_SERIES ModelUnratedItem = "Series"
	MODELUNRATEDITEM_MUSIC ModelUnratedItem = "Music"
	MODELUNRATEDITEM_GAME ModelUnratedItem = "Game"
	MODELUNRATEDITEM_BOOK ModelUnratedItem = "Book"
	MODELUNRATEDITEM_LIVE_TV_CHANNEL ModelUnratedItem = "LiveTvChannel"
	MODELUNRATEDITEM_LIVE_TV_PROGRAM ModelUnratedItem = "LiveTvProgram"
	MODELUNRATEDITEM_CHANNEL_CONTENT ModelUnratedItem = "ChannelContent"
	MODELUNRATEDITEM_OTHER ModelUnratedItem = "Other"
)

// All allowed values of ModelUnratedItem enum
var AllowedModelUnratedItemEnumValues = []ModelUnratedItem{
	"Movie",
	"Trailer",
	"Series",
	"Music",
	"Game",
	"Book",
	"LiveTvChannel",
	"LiveTvProgram",
	"ChannelContent",
	"Other",
}

func (v *ModelUnratedItem) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelUnratedItem(value)
	for _, existing := range AllowedModelUnratedItemEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelUnratedItem", value)
}

// NewModelUnratedItemFromValue returns a pointer to a valid ModelUnratedItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelUnratedItemFromValue(v string) (*ModelUnratedItem, error) {
	ev := ModelUnratedItem(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelUnratedItem: valid values are %v", v, AllowedModelUnratedItemEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelUnratedItem) IsValid() bool {
	for _, existing := range AllowedModelUnratedItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnratedItem value
func (v ModelUnratedItem) Ptr() *ModelUnratedItem {
	return &v
}

type NullableModelUnratedItem struct {
	value *ModelUnratedItem
	isSet bool
}

func (v NullableModelUnratedItem) Get() *ModelUnratedItem {
	return v.value
}

func (v *NullableModelUnratedItem) Set(val *ModelUnratedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUnratedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUnratedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUnratedItem(val *ModelUnratedItem) *NullableModelUnratedItem {
	return &NullableModelUnratedItem{value: val, isSet: true}
}

func (v NullableModelUnratedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUnratedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

