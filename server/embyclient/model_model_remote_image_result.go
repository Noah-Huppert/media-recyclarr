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

// checks if the ModelRemoteImageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRemoteImageResult{}

// ModelRemoteImageResult struct for ModelRemoteImageResult
type ModelRemoteImageResult struct {
	Images []ModelRemoteImageInfo `json:"Images,omitempty"`
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	Providers []string `json:"Providers,omitempty"`
}

// NewModelRemoteImageResult instantiates a new ModelRemoteImageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRemoteImageResult() *ModelRemoteImageResult {
	this := ModelRemoteImageResult{}
	return &this
}

// NewModelRemoteImageResultWithDefaults instantiates a new ModelRemoteImageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRemoteImageResultWithDefaults() *ModelRemoteImageResult {
	this := ModelRemoteImageResult{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ModelRemoteImageResult) GetImages() []ModelRemoteImageInfo {
	if o == nil || IsNil(o.Images) {
		var ret []ModelRemoteImageInfo
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteImageResult) GetImagesOk() ([]ModelRemoteImageInfo, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ModelRemoteImageResult) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []ModelRemoteImageInfo and assigns it to the Images field.
func (o *ModelRemoteImageResult) SetImages(v []ModelRemoteImageInfo) {
	o.Images = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *ModelRemoteImageResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteImageResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *ModelRemoteImageResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *ModelRemoteImageResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *ModelRemoteImageResult) GetProviders() []string {
	if o == nil || IsNil(o.Providers) {
		var ret []string
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteImageResult) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *ModelRemoteImageResult) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *ModelRemoteImageResult) SetProviders(v []string) {
	o.Providers = v
}

func (o ModelRemoteImageResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRemoteImageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Images) {
		toSerialize["Images"] = o.Images
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	if !IsNil(o.Providers) {
		toSerialize["Providers"] = o.Providers
	}
	return toSerialize, nil
}

type NullableModelRemoteImageResult struct {
	value *ModelRemoteImageResult
	isSet bool
}

func (v NullableModelRemoteImageResult) Get() *ModelRemoteImageResult {
	return v.value
}

func (v *NullableModelRemoteImageResult) Set(val *ModelRemoteImageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRemoteImageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRemoteImageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRemoteImageResult(val *ModelRemoteImageResult) *NullableModelRemoteImageResult {
	return &NullableModelRemoteImageResult{value: val, isSet: true}
}

func (v NullableModelRemoteImageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRemoteImageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

