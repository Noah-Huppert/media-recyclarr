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

// checks if the ModelQueryResultActivityLogEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelQueryResultActivityLogEntry{}

// ModelQueryResultActivityLogEntry struct for ModelQueryResultActivityLogEntry
type ModelQueryResultActivityLogEntry struct {
	Items []ModelActivityLogEntry `json:"Items,omitempty"`
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
}

// NewModelQueryResultActivityLogEntry instantiates a new ModelQueryResultActivityLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelQueryResultActivityLogEntry() *ModelQueryResultActivityLogEntry {
	this := ModelQueryResultActivityLogEntry{}
	return &this
}

// NewModelQueryResultActivityLogEntryWithDefaults instantiates a new ModelQueryResultActivityLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelQueryResultActivityLogEntryWithDefaults() *ModelQueryResultActivityLogEntry {
	this := ModelQueryResultActivityLogEntry{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelQueryResultActivityLogEntry) GetItems() []ModelActivityLogEntry {
	if o == nil || IsNil(o.Items) {
		var ret []ModelActivityLogEntry
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelQueryResultActivityLogEntry) GetItemsOk() ([]ModelActivityLogEntry, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelQueryResultActivityLogEntry) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelActivityLogEntry and assigns it to the Items field.
func (o *ModelQueryResultActivityLogEntry) SetItems(v []ModelActivityLogEntry) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *ModelQueryResultActivityLogEntry) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelQueryResultActivityLogEntry) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *ModelQueryResultActivityLogEntry) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *ModelQueryResultActivityLogEntry) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

func (o ModelQueryResultActivityLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelQueryResultActivityLogEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	return toSerialize, nil
}

type NullableModelQueryResultActivityLogEntry struct {
	value *ModelQueryResultActivityLogEntry
	isSet bool
}

func (v NullableModelQueryResultActivityLogEntry) Get() *ModelQueryResultActivityLogEntry {
	return v.value
}

func (v *NullableModelQueryResultActivityLogEntry) Set(val *ModelQueryResultActivityLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableModelQueryResultActivityLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableModelQueryResultActivityLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelQueryResultActivityLogEntry(val *ModelQueryResultActivityLogEntry) *NullableModelQueryResultActivityLogEntry {
	return &NullableModelQueryResultActivityLogEntry{value: val, isSet: true}
}

func (v NullableModelQueryResultActivityLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelQueryResultActivityLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


