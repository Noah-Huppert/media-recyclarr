/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
)

// checks if the ModelExternalIdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelExternalIdInfo{}

// ModelExternalIdInfo struct for ModelExternalIdInfo
type ModelExternalIdInfo struct {
	Name *string `json:"Name,omitempty"`
	Key *string `json:"Key,omitempty"`
	UrlFormatString *string `json:"UrlFormatString,omitempty"`
	IsSupportedAsIdentifier *bool `json:"IsSupportedAsIdentifier,omitempty"`
}

// NewModelExternalIdInfo instantiates a new ModelExternalIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelExternalIdInfo() *ModelExternalIdInfo {
	this := ModelExternalIdInfo{}
	return &this
}

// NewModelExternalIdInfoWithDefaults instantiates a new ModelExternalIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelExternalIdInfoWithDefaults() *ModelExternalIdInfo {
	this := ModelExternalIdInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelExternalIdInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExternalIdInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelExternalIdInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelExternalIdInfo) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ModelExternalIdInfo) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExternalIdInfo) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ModelExternalIdInfo) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ModelExternalIdInfo) SetKey(v string) {
	o.Key = &v
}

// GetUrlFormatString returns the UrlFormatString field value if set, zero value otherwise.
func (o *ModelExternalIdInfo) GetUrlFormatString() string {
	if o == nil || IsNil(o.UrlFormatString) {
		var ret string
		return ret
	}
	return *o.UrlFormatString
}

// GetUrlFormatStringOk returns a tuple with the UrlFormatString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExternalIdInfo) GetUrlFormatStringOk() (*string, bool) {
	if o == nil || IsNil(o.UrlFormatString) {
		return nil, false
	}
	return o.UrlFormatString, true
}

// HasUrlFormatString returns a boolean if a field has been set.
func (o *ModelExternalIdInfo) HasUrlFormatString() bool {
	if o != nil && !IsNil(o.UrlFormatString) {
		return true
	}

	return false
}

// SetUrlFormatString gets a reference to the given string and assigns it to the UrlFormatString field.
func (o *ModelExternalIdInfo) SetUrlFormatString(v string) {
	o.UrlFormatString = &v
}

// GetIsSupportedAsIdentifier returns the IsSupportedAsIdentifier field value if set, zero value otherwise.
func (o *ModelExternalIdInfo) GetIsSupportedAsIdentifier() bool {
	if o == nil || IsNil(o.IsSupportedAsIdentifier) {
		var ret bool
		return ret
	}
	return *o.IsSupportedAsIdentifier
}

// GetIsSupportedAsIdentifierOk returns a tuple with the IsSupportedAsIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExternalIdInfo) GetIsSupportedAsIdentifierOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSupportedAsIdentifier) {
		return nil, false
	}
	return o.IsSupportedAsIdentifier, true
}

// HasIsSupportedAsIdentifier returns a boolean if a field has been set.
func (o *ModelExternalIdInfo) HasIsSupportedAsIdentifier() bool {
	if o != nil && !IsNil(o.IsSupportedAsIdentifier) {
		return true
	}

	return false
}

// SetIsSupportedAsIdentifier gets a reference to the given bool and assigns it to the IsSupportedAsIdentifier field.
func (o *ModelExternalIdInfo) SetIsSupportedAsIdentifier(v bool) {
	o.IsSupportedAsIdentifier = &v
}

func (o ModelExternalIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelExternalIdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.UrlFormatString) {
		toSerialize["UrlFormatString"] = o.UrlFormatString
	}
	if !IsNil(o.IsSupportedAsIdentifier) {
		toSerialize["IsSupportedAsIdentifier"] = o.IsSupportedAsIdentifier
	}
	return toSerialize, nil
}

type NullableModelExternalIdInfo struct {
	value *ModelExternalIdInfo
	isSet bool
}

func (v NullableModelExternalIdInfo) Get() *ModelExternalIdInfo {
	return v.value
}

func (v *NullableModelExternalIdInfo) Set(val *ModelExternalIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelExternalIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelExternalIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelExternalIdInfo(val *ModelExternalIdInfo) *NullableModelExternalIdInfo {
	return &NullableModelExternalIdInfo{value: val, isSet: true}
}

func (v NullableModelExternalIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelExternalIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

