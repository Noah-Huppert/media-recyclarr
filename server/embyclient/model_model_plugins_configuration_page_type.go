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

// ModelPluginsConfigurationPageType the model 'ModelPluginsConfigurationPageType'
type ModelPluginsConfigurationPageType string

// List of Plugins.ConfigurationPageType
const (
	MODELPLUGINSCONFIGURATIONPAGETYPE_PLUGIN_CONFIGURATION ModelPluginsConfigurationPageType = "PluginConfiguration"
	MODELPLUGINSCONFIGURATIONPAGETYPE_NONE ModelPluginsConfigurationPageType = "None"
)

// All allowed values of ModelPluginsConfigurationPageType enum
var AllowedModelPluginsConfigurationPageTypeEnumValues = []ModelPluginsConfigurationPageType{
	"PluginConfiguration",
	"None",
}

func (v *ModelPluginsConfigurationPageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModelPluginsConfigurationPageType(value)
	for _, existing := range AllowedModelPluginsConfigurationPageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModelPluginsConfigurationPageType", value)
}

// NewModelPluginsConfigurationPageTypeFromValue returns a pointer to a valid ModelPluginsConfigurationPageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModelPluginsConfigurationPageTypeFromValue(v string) (*ModelPluginsConfigurationPageType, error) {
	ev := ModelPluginsConfigurationPageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModelPluginsConfigurationPageType: valid values are %v", v, AllowedModelPluginsConfigurationPageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModelPluginsConfigurationPageType) IsValid() bool {
	for _, existing := range AllowedModelPluginsConfigurationPageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Plugins.ConfigurationPageType value
func (v ModelPluginsConfigurationPageType) Ptr() *ModelPluginsConfigurationPageType {
	return &v
}

type NullableModelPluginsConfigurationPageType struct {
	value *ModelPluginsConfigurationPageType
	isSet bool
}

func (v NullableModelPluginsConfigurationPageType) Get() *ModelPluginsConfigurationPageType {
	return v.value
}

func (v *NullableModelPluginsConfigurationPageType) Set(val *ModelPluginsConfigurationPageType) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPluginsConfigurationPageType) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPluginsConfigurationPageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPluginsConfigurationPageType(val *ModelPluginsConfigurationPageType) *NullableModelPluginsConfigurationPageType {
	return &NullableModelPluginsConfigurationPageType{value: val, isSet: true}
}

func (v NullableModelPluginsConfigurationPageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPluginsConfigurationPageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
