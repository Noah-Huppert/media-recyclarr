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

// checks if the ModelNotificationTypeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelNotificationTypeInfo{}

// ModelNotificationTypeInfo struct for ModelNotificationTypeInfo
type ModelNotificationTypeInfo struct {
	Name *string `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty"`
	CategoryId *string `json:"CategoryId,omitempty"`
}

// NewModelNotificationTypeInfo instantiates a new ModelNotificationTypeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelNotificationTypeInfo() *ModelNotificationTypeInfo {
	this := ModelNotificationTypeInfo{}
	return &this
}

// NewModelNotificationTypeInfoWithDefaults instantiates a new ModelNotificationTypeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelNotificationTypeInfoWithDefaults() *ModelNotificationTypeInfo {
	this := ModelNotificationTypeInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelNotificationTypeInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelNotificationTypeInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelNotificationTypeInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelNotificationTypeInfo) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelNotificationTypeInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelNotificationTypeInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelNotificationTypeInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelNotificationTypeInfo) SetId(v string) {
	o.Id = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *ModelNotificationTypeInfo) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName) {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelNotificationTypeInfo) GetCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryName) {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *ModelNotificationTypeInfo) HasCategoryName() bool {
	if o != nil && !IsNil(o.CategoryName) {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *ModelNotificationTypeInfo) SetCategoryName(v string) {
	o.CategoryName = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ModelNotificationTypeInfo) GetCategoryId() string {
	if o == nil || IsNil(o.CategoryId) {
		var ret string
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelNotificationTypeInfo) GetCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ModelNotificationTypeInfo) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given string and assigns it to the CategoryId field.
func (o *ModelNotificationTypeInfo) SetCategoryId(v string) {
	o.CategoryId = &v
}

func (o ModelNotificationTypeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelNotificationTypeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.CategoryName) {
		toSerialize["CategoryName"] = o.CategoryName
	}
	if !IsNil(o.CategoryId) {
		toSerialize["CategoryId"] = o.CategoryId
	}
	return toSerialize, nil
}

type NullableModelNotificationTypeInfo struct {
	value *ModelNotificationTypeInfo
	isSet bool
}

func (v NullableModelNotificationTypeInfo) Get() *ModelNotificationTypeInfo {
	return v.value
}

func (v *NullableModelNotificationTypeInfo) Set(val *ModelNotificationTypeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelNotificationTypeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelNotificationTypeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelNotificationTypeInfo(val *ModelNotificationTypeInfo) *NullableModelNotificationTypeInfo {
	return &NullableModelNotificationTypeInfo{value: val, isSet: true}
}

func (v NullableModelNotificationTypeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelNotificationTypeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

