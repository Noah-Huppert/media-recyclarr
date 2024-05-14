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

// checks if the ModelRemoteSearchQueryItemLookupInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRemoteSearchQueryItemLookupInfo{}

// ModelRemoteSearchQueryItemLookupInfo struct for ModelRemoteSearchQueryItemLookupInfo
type ModelRemoteSearchQueryItemLookupInfo struct {
	SearchInfo *ModelItemLookupInfo `json:"SearchInfo,omitempty"`
	ItemId *int64 `json:"ItemId,omitempty"`
	SearchProviderName *string `json:"SearchProviderName,omitempty"`
	Providers []string `json:"Providers,omitempty"`
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewModelRemoteSearchQueryItemLookupInfo instantiates a new ModelRemoteSearchQueryItemLookupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRemoteSearchQueryItemLookupInfo() *ModelRemoteSearchQueryItemLookupInfo {
	this := ModelRemoteSearchQueryItemLookupInfo{}
	return &this
}

// NewModelRemoteSearchQueryItemLookupInfoWithDefaults instantiates a new ModelRemoteSearchQueryItemLookupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRemoteSearchQueryItemLookupInfoWithDefaults() *ModelRemoteSearchQueryItemLookupInfo {
	this := ModelRemoteSearchQueryItemLookupInfo{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetSearchInfo() ModelItemLookupInfo {
	if o == nil || IsNil(o.SearchInfo) {
		var ret ModelItemLookupInfo
		return ret
	}
	return *o.SearchInfo
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetSearchInfoOk() (*ModelItemLookupInfo, bool) {
	if o == nil || IsNil(o.SearchInfo) {
		return nil, false
	}
	return o.SearchInfo, true
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) HasSearchInfo() bool {
	if o != nil && !IsNil(o.SearchInfo) {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given ModelItemLookupInfo and assigns it to the SearchInfo field.
func (o *ModelRemoteSearchQueryItemLookupInfo) SetSearchInfo(v ModelItemLookupInfo) {
	o.SearchInfo = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetItemId() int64 {
	if o == nil || IsNil(o.ItemId) {
		var ret int64
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetItemIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int64 and assigns it to the ItemId field.
func (o *ModelRemoteSearchQueryItemLookupInfo) SetItemId(v int64) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName) {
		var ret string
		return ret
	}
	return *o.SearchProviderName
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetSearchProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SearchProviderName) {
		return nil, false
	}
	return o.SearchProviderName, true
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) HasSearchProviderName() bool {
	if o != nil && !IsNil(o.SearchProviderName) {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given string and assigns it to the SearchProviderName field.
func (o *ModelRemoteSearchQueryItemLookupInfo) SetSearchProviderName(v string) {
	o.SearchProviderName = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetProviders() []string {
	if o == nil || IsNil(o.Providers) {
		var ret []string
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *ModelRemoteSearchQueryItemLookupInfo) SetProviders(v []string) {
	o.Providers = v
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *ModelRemoteSearchQueryItemLookupInfo) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *ModelRemoteSearchQueryItemLookupInfo) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o ModelRemoteSearchQueryItemLookupInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRemoteSearchQueryItemLookupInfo) ToMap() (map[string]interface{}, error) {
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

type NullableModelRemoteSearchQueryItemLookupInfo struct {
	value *ModelRemoteSearchQueryItemLookupInfo
	isSet bool
}

func (v NullableModelRemoteSearchQueryItemLookupInfo) Get() *ModelRemoteSearchQueryItemLookupInfo {
	return v.value
}

func (v *NullableModelRemoteSearchQueryItemLookupInfo) Set(val *ModelRemoteSearchQueryItemLookupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRemoteSearchQueryItemLookupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRemoteSearchQueryItemLookupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRemoteSearchQueryItemLookupInfo(val *ModelRemoteSearchQueryItemLookupInfo) *NullableModelRemoteSearchQueryItemLookupInfo {
	return &NullableModelRemoteSearchQueryItemLookupInfo{value: val, isSet: true}
}

func (v NullableModelRemoteSearchQueryItemLookupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRemoteSearchQueryItemLookupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

