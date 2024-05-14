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

// ModelCommonEditorTypes the model 'ModelCommonEditorTypes'
type ModelCommonEditorTypes string

// List of Common.EditorTypes
const (
	MODELCOMMONEDITORTYPES_GROUP ModelCommonEditorTypes = "Group"
	MODELCOMMONEDITORTYPES_TEXT ModelCommonEditorTypes = "Text"
	MODELCOMMONEDITORTYPES_NUMERIC ModelCommonEditorTypes = "Numeric"
	MODELCOMMONEDITORTYPES_BOOLEAN ModelCommonEditorTypes = "Boolean"
	MODELCOMMONEDITORTYPES_SELECT_SINGLE ModelCommonEditorTypes = "SelectSingle"
	MODELCOMMONEDITORTYPES_SELECT_MULTIPLE ModelCommonEditorTypes = "SelectMultiple"
	MODELCOMMONEDITORTYPES_DATE ModelCommonEditorTypes = "Date"
	MODELCOMMONEDITORTYPES_FILE_PATH ModelCommonEditorTypes = "FilePath"
	MODELCOMMONEDITORTYPES_FOLDER_PATH ModelCommonEditorTypes = "FolderPath"
	MODELCOMMONEDITORTYPES_STATUS_ITEM ModelCommonEditorTypes = "StatusItem"
	MODELCOMMONEDITORTYPES_PROGRESS_ITEM ModelCommonEditorTypes = "ProgressItem"
	MODELCOMMONEDITORTYPES_BUTTON_ITEM ModelCommonEditorTypes = "ButtonItem"
	MODELCOMMONEDITORTYPES_BUTTON_GROUP ModelCommonEditorTypes = "ButtonGroup"
	MODELCOMMONEDITORTYPES_CAPTION_ITEM ModelCommonEditorTypes = "CaptionItem"
	MODELCOMMONEDITORTYPES_LABEL_ITEM ModelCommonEditorTypes = "LabelItem"
	MODELCOMMONEDITORTYPES_ITEM_LIST ModelCommonEditorTypes = "ItemList"
	MODELCOMMONEDITORTYPES_RADIO_GROUP ModelCommonEditorTypes = "RadioGroup"
	MODELCOMMONEDITORTYPES_DX_DATA_GRID ModelCommonEditorTypes = "DxDataGrid"
	MODELCOMMONEDITORTYPES_DX_PIVOT_GRID ModelCommonEditorTypes = "DxPivotGrid"
	MODELCOMMONEDITORTYPES_SPACER_ITEM ModelCommonEditorTypes = "SpacerItem"
)

// All allowed values of ModelCommonEditorTypes enum
var AllowedModelCommonEditorTypesEnumValues = []ModelCommonEditorTypes{
	"Group",
	"Text",
	"Numeric",
	"Boolean",
	"SelectSingle",
	"SelectMultiple",
	"Date",
	"FilePath",
	"FolderPath",
	"StatusItem",
	"ProgressItem",
	"ButtonItem",
	"ButtonGroup",
	"CaptionItem",
	"LabelItem",
	"ItemList",
	"RadioGroup",
	"DxDataGrid",
	"DxPivotGrid",
	"SpacerItem",
}

func (v *ModelCommonEditorTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelCommonEditorTypes(value)
	for _, existing := range AllowedModelCommonEditorTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelCommonEditorTypes", value)
}

// NewModelCommonEditorTypesFromValue returns a pointer to a valid ModelCommonEditorTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelCommonEditorTypesFromValue(v string) (*ModelCommonEditorTypes, error) {
	ev := ModelCommonEditorTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelCommonEditorTypes: valid values are %v", v, AllowedModelCommonEditorTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelCommonEditorTypes) IsValid() bool {
	for _, existing := range AllowedModelCommonEditorTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Common.EditorTypes value
func (v ModelCommonEditorTypes) Ptr() *ModelCommonEditorTypes {
	return &v
}

type NullableModelCommonEditorTypes struct {
	value *ModelCommonEditorTypes
	isSet bool
}

func (v NullableModelCommonEditorTypes) Get() *ModelCommonEditorTypes {
	return v.value
}

func (v *NullableModelCommonEditorTypes) Set(val *ModelCommonEditorTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCommonEditorTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCommonEditorTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCommonEditorTypes(val *ModelCommonEditorTypes) *NullableModelCommonEditorTypes {
	return &NullableModelCommonEditorTypes{value: val, isSet: true}
}

func (v NullableModelCommonEditorTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCommonEditorTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

