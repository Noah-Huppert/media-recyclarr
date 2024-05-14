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

// checks if the ModelQueryResultLiveTvSeriesTimerInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelQueryResultLiveTvSeriesTimerInfoDto{}

// ModelQueryResultLiveTvSeriesTimerInfoDto struct for ModelQueryResultLiveTvSeriesTimerInfoDto
type ModelQueryResultLiveTvSeriesTimerInfoDto struct {
	Items []ModelLiveTvSeriesTimerInfoDto `json:"Items,omitempty"`
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
}

// NewModelQueryResultLiveTvSeriesTimerInfoDto instantiates a new ModelQueryResultLiveTvSeriesTimerInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelQueryResultLiveTvSeriesTimerInfoDto() *ModelQueryResultLiveTvSeriesTimerInfoDto {
	this := ModelQueryResultLiveTvSeriesTimerInfoDto{}
	return &this
}

// NewModelQueryResultLiveTvSeriesTimerInfoDtoWithDefaults instantiates a new ModelQueryResultLiveTvSeriesTimerInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelQueryResultLiveTvSeriesTimerInfoDtoWithDefaults() *ModelQueryResultLiveTvSeriesTimerInfoDto {
	this := ModelQueryResultLiveTvSeriesTimerInfoDto{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) GetItems() []ModelLiveTvSeriesTimerInfoDto {
	if o == nil || IsNil(o.Items) {
		var ret []ModelLiveTvSeriesTimerInfoDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) GetItemsOk() ([]ModelLiveTvSeriesTimerInfoDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ModelLiveTvSeriesTimerInfoDto and assigns it to the Items field.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) SetItems(v []ModelLiveTvSeriesTimerInfoDto) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *ModelQueryResultLiveTvSeriesTimerInfoDto) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

func (o ModelQueryResultLiveTvSeriesTimerInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelQueryResultLiveTvSeriesTimerInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	return toSerialize, nil
}

type NullableModelQueryResultLiveTvSeriesTimerInfoDto struct {
	value *ModelQueryResultLiveTvSeriesTimerInfoDto
	isSet bool
}

func (v NullableModelQueryResultLiveTvSeriesTimerInfoDto) Get() *ModelQueryResultLiveTvSeriesTimerInfoDto {
	return v.value
}

func (v *NullableModelQueryResultLiveTvSeriesTimerInfoDto) Set(val *ModelQueryResultLiveTvSeriesTimerInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModelQueryResultLiveTvSeriesTimerInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModelQueryResultLiveTvSeriesTimerInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelQueryResultLiveTvSeriesTimerInfoDto(val *ModelQueryResultLiveTvSeriesTimerInfoDto) *NullableModelQueryResultLiveTvSeriesTimerInfoDto {
	return &NullableModelQueryResultLiveTvSeriesTimerInfoDto{value: val, isSet: true}
}

func (v NullableModelQueryResultLiveTvSeriesTimerInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelQueryResultLiveTvSeriesTimerInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


