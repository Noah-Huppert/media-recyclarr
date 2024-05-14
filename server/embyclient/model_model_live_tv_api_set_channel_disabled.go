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

// checks if the ModelLiveTVApiSetChannelDisabled type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLiveTVApiSetChannelDisabled{}

// ModelLiveTVApiSetChannelDisabled struct for ModelLiveTVApiSetChannelDisabled
type ModelLiveTVApiSetChannelDisabled struct {
	Id *string `json:"Id,omitempty"`
	ManagementId *string `json:"ManagementId,omitempty"`
	Disabled *bool `json:"Disabled,omitempty"`
}

// NewModelLiveTVApiSetChannelDisabled instantiates a new ModelLiveTVApiSetChannelDisabled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLiveTVApiSetChannelDisabled() *ModelLiveTVApiSetChannelDisabled {
	this := ModelLiveTVApiSetChannelDisabled{}
	return &this
}

// NewModelLiveTVApiSetChannelDisabledWithDefaults instantiates a new ModelLiveTVApiSetChannelDisabled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLiveTVApiSetChannelDisabledWithDefaults() *ModelLiveTVApiSetChannelDisabled {
	this := ModelLiveTVApiSetChannelDisabled{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelLiveTVApiSetChannelDisabled) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTVApiSetChannelDisabled) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelLiveTVApiSetChannelDisabled) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelLiveTVApiSetChannelDisabled) SetId(v string) {
	o.Id = &v
}

// GetManagementId returns the ManagementId field value if set, zero value otherwise.
func (o *ModelLiveTVApiSetChannelDisabled) GetManagementId() string {
	if o == nil || IsNil(o.ManagementId) {
		var ret string
		return ret
	}
	return *o.ManagementId
}

// GetManagementIdOk returns a tuple with the ManagementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTVApiSetChannelDisabled) GetManagementIdOk() (*string, bool) {
	if o == nil || IsNil(o.ManagementId) {
		return nil, false
	}
	return o.ManagementId, true
}

// HasManagementId returns a boolean if a field has been set.
func (o *ModelLiveTVApiSetChannelDisabled) HasManagementId() bool {
	if o != nil && !IsNil(o.ManagementId) {
		return true
	}

	return false
}

// SetManagementId gets a reference to the given string and assigns it to the ManagementId field.
func (o *ModelLiveTVApiSetChannelDisabled) SetManagementId(v string) {
	o.ManagementId = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ModelLiveTVApiSetChannelDisabled) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTVApiSetChannelDisabled) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ModelLiveTVApiSetChannelDisabled) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ModelLiveTVApiSetChannelDisabled) SetDisabled(v bool) {
	o.Disabled = &v
}

func (o ModelLiveTVApiSetChannelDisabled) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLiveTVApiSetChannelDisabled) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.ManagementId) {
		toSerialize["ManagementId"] = o.ManagementId
	}
	if !IsNil(o.Disabled) {
		toSerialize["Disabled"] = o.Disabled
	}
	return toSerialize, nil
}

type NullableModelLiveTVApiSetChannelDisabled struct {
	value *ModelLiveTVApiSetChannelDisabled
	isSet bool
}

func (v NullableModelLiveTVApiSetChannelDisabled) Get() *ModelLiveTVApiSetChannelDisabled {
	return v.value
}

func (v *NullableModelLiveTVApiSetChannelDisabled) Set(val *ModelLiveTVApiSetChannelDisabled) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLiveTVApiSetChannelDisabled) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLiveTVApiSetChannelDisabled) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLiveTVApiSetChannelDisabled(val *ModelLiveTVApiSetChannelDisabled) *NullableModelLiveTVApiSetChannelDisabled {
	return &NullableModelLiveTVApiSetChannelDisabled{value: val, isSet: true}
}

func (v NullableModelLiveTVApiSetChannelDisabled) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLiveTVApiSetChannelDisabled) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

