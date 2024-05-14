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

// checks if the ModelRemoteSearchQuerySeriesInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRemoteSearchQuerySeriesInfo{}

// ModelRemoteSearchQuerySeriesInfo struct for ModelRemoteSearchQuerySeriesInfo
type ModelRemoteSearchQuerySeriesInfo struct {
	SearchInfo *ModelSeriesInfo `json:"SearchInfo,omitempty"`
	ItemId *int64 `json:"ItemId,omitempty"`
	SearchProviderName *string `json:"SearchProviderName,omitempty"`
	Providers []string `json:"Providers,omitempty"`
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewModelRemoteSearchQuerySeriesInfo instantiates a new ModelRemoteSearchQuerySeriesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRemoteSearchQuerySeriesInfo() *ModelRemoteSearchQuerySeriesInfo {
	this := ModelRemoteSearchQuerySeriesInfo{}
	return &this
}

// NewModelRemoteSearchQuerySeriesInfoWithDefaults instantiates a new ModelRemoteSearchQuerySeriesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRemoteSearchQuerySeriesInfoWithDefaults() *ModelRemoteSearchQuerySeriesInfo {
	this := ModelRemoteSearchQuerySeriesInfo{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise.
func (o *ModelRemoteSearchQuerySeriesInfo) GetSearchInfo() ModelSeriesInfo {
	if o == nil || IsNil(o.SearchInfo) {
		var ret ModelSeriesInfo
		return ret
	}
	return *o.SearchInfo
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) GetSearchInfoOk() (*ModelSeriesInfo, bool) {
	if o == nil || IsNil(o.SearchInfo) {
		return nil, false
	}
	return o.SearchInfo, true
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) HasSearchInfo() bool {
	if o != nil && !IsNil(o.SearchInfo) {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given ModelSeriesInfo and assigns it to the SearchInfo field.
func (o *ModelRemoteSearchQuerySeriesInfo) SetSearchInfo(v ModelSeriesInfo) {
	o.SearchInfo = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ModelRemoteSearchQuerySeriesInfo) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *ModelRemoteSearchQuerySeriesInfo) SetItemId(v int64) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise.
func (o *ModelRemoteSearchQuerySeriesInfo) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName) {
		var ret string
		return ret
	}
	return *o.SearchProviderName
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) GetSearchProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SearchProviderName) {
		return nil, false
	}
	return o.SearchProviderName, true
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) HasSearchProviderName() bool {
	if o != nil && !IsNil(o.SearchProviderName) {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given string and assigns it to the SearchProviderName field.
func (o *ModelRemoteSearchQuerySeriesInfo) SetSearchProviderName(v string) {
	o.SearchProviderName = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *ModelRemoteSearchQuerySeriesInfo) GetProviders() []string {
	if o == nil || IsNil(o.Providers) {
		var ret []string
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *ModelRemoteSearchQuerySeriesInfo) SetProviders(v []string) {
	o.Providers = v
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *ModelRemoteSearchQuerySeriesInfo) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *ModelRemoteSearchQuerySeriesInfo) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *ModelRemoteSearchQuerySeriesInfo) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o ModelRemoteSearchQuerySeriesInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRemoteSearchQuerySeriesInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchInfo) {
		toSerialize["SearchInfo"] = o.SearchInfo
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.SearchProviderName) {
		toSerialize["SearchProviderName"] = o.SearchProviderName
	}
	if !IsNil(o.Providers) {
		toSerialize["Providers"] = o.Providers
	}
	if !IsNil(o.IncludeDisabledProviders) {
		toSerialize["IncludeDisabledProviders"] = o.IncludeDisabledProviders
	}
	return toSerialize, nil
}

type NullableModelRemoteSearchQuerySeriesInfo struct {
	value *ModelRemoteSearchQuerySeriesInfo
	isSet bool
}

func (v NullableModelRemoteSearchQuerySeriesInfo) Get() *ModelRemoteSearchQuerySeriesInfo {
	return v.value
}

func (v *NullableModelRemoteSearchQuerySeriesInfo) Set(val *ModelRemoteSearchQuerySeriesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRemoteSearchQuerySeriesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRemoteSearchQuerySeriesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRemoteSearchQuerySeriesInfo(val *ModelRemoteSearchQuerySeriesInfo) *NullableModelRemoteSearchQuerySeriesInfo {
	return &NullableModelRemoteSearchQuerySeriesInfo{value: val, isSet: true}
}

func (v NullableModelRemoteSearchQuerySeriesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRemoteSearchQuerySeriesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

