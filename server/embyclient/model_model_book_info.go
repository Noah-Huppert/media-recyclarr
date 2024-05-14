/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
	"time"
)

// checks if the ModelBookInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBookInfo{}

// ModelBookInfo struct for ModelBookInfo
type ModelBookInfo struct {
	SeriesName *string `json:"SeriesName,omitempty"`
	Name *string `json:"Name,omitempty"`
	MetadataLanguage *string `json:"MetadataLanguage,omitempty"`
	MetadataCountryCode *string `json:"MetadataCountryCode,omitempty"`
	MetadataLanguages []ModelGlobalizationCultureDto `json:"MetadataLanguages,omitempty"`
	ProviderIds *map[string]string `json:"ProviderIds,omitempty"`
	Year NullableInt32 `json:"Year,omitempty"`
	IndexNumber NullableInt32 `json:"IndexNumber,omitempty"`
	ParentIndexNumber NullableInt32 `json:"ParentIndexNumber,omitempty"`
	PremiereDate NullableTime `json:"PremiereDate,omitempty"`
	IsAutomated *bool `json:"IsAutomated,omitempty"`
	EnableAdultMetadata *bool `json:"EnableAdultMetadata,omitempty"`
}

// NewModelBookInfo instantiates a new ModelBookInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBookInfo() *ModelBookInfo {
	this := ModelBookInfo{}
	return &this
}

// NewModelBookInfoWithDefaults instantiates a new ModelBookInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBookInfoWithDefaults() *ModelBookInfo {
	this := ModelBookInfo{}
	return &this
}

// GetSeriesName returns the SeriesName field value if set, zero value otherwise.
func (o *ModelBookInfo) GetSeriesName() string {
	if o == nil || IsNil(o.SeriesName) {
		var ret string
		return ret
	}
	return *o.SeriesName
}

// GetSeriesNameOk returns a tuple with the SeriesName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetSeriesNameOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesName) {
		return nil, false
	}
	return o.SeriesName, true
}

// HasSeriesName returns a boolean if a field has been set.
func (o *ModelBookInfo) HasSeriesName() bool {
	if o != nil && !IsNil(o.SeriesName) {
		return true
	}

	return false
}

// SetSeriesName gets a reference to the given string and assigns it to the SeriesName field.
func (o *ModelBookInfo) SetSeriesName(v string) {
	o.SeriesName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelBookInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelBookInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelBookInfo) SetName(v string) {
	o.Name = &v
}

// GetMetadataLanguage returns the MetadataLanguage field value if set, zero value otherwise.
func (o *ModelBookInfo) GetMetadataLanguage() string {
	if o == nil || IsNil(o.MetadataLanguage) {
		var ret string
		return ret
	}
	return *o.MetadataLanguage
}

// GetMetadataLanguageOk returns a tuple with the MetadataLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetMetadataLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataLanguage) {
		return nil, false
	}
	return o.MetadataLanguage, true
}

// HasMetadataLanguage returns a boolean if a field has been set.
func (o *ModelBookInfo) HasMetadataLanguage() bool {
	if o != nil && !IsNil(o.MetadataLanguage) {
		return true
	}

	return false
}

// SetMetadataLanguage gets a reference to the given string and assigns it to the MetadataLanguage field.
func (o *ModelBookInfo) SetMetadataLanguage(v string) {
	o.MetadataLanguage = &v
}

// GetMetadataCountryCode returns the MetadataCountryCode field value if set, zero value otherwise.
func (o *ModelBookInfo) GetMetadataCountryCode() string {
	if o == nil || IsNil(o.MetadataCountryCode) {
		var ret string
		return ret
	}
	return *o.MetadataCountryCode
}

// GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetMetadataCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MetadataCountryCode) {
		return nil, false
	}
	return o.MetadataCountryCode, true
}

// HasMetadataCountryCode returns a boolean if a field has been set.
func (o *ModelBookInfo) HasMetadataCountryCode() bool {
	if o != nil && !IsNil(o.MetadataCountryCode) {
		return true
	}

	return false
}

// SetMetadataCountryCode gets a reference to the given string and assigns it to the MetadataCountryCode field.
func (o *ModelBookInfo) SetMetadataCountryCode(v string) {
	o.MetadataCountryCode = &v
}

// GetMetadataLanguages returns the MetadataLanguages field value if set, zero value otherwise.
func (o *ModelBookInfo) GetMetadataLanguages() []ModelGlobalizationCultureDto {
	if o == nil || IsNil(o.MetadataLanguages) {
		var ret []ModelGlobalizationCultureDto
		return ret
	}
	return o.MetadataLanguages
}

// GetMetadataLanguagesOk returns a tuple with the MetadataLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetMetadataLanguagesOk() ([]ModelGlobalizationCultureDto, bool) {
	if o == nil || IsNil(o.MetadataLanguages) {
		return nil, false
	}
	return o.MetadataLanguages, true
}

// HasMetadataLanguages returns a boolean if a field has been set.
func (o *ModelBookInfo) HasMetadataLanguages() bool {
	if o != nil && !IsNil(o.MetadataLanguages) {
		return true
	}

	return false
}

// SetMetadataLanguages gets a reference to the given []ModelGlobalizationCultureDto and assigns it to the MetadataLanguages field.
func (o *ModelBookInfo) SetMetadataLanguages(v []ModelGlobalizationCultureDto) {
	o.MetadataLanguages = v
}

// GetProviderIds returns the ProviderIds field value if set, zero value otherwise.
func (o *ModelBookInfo) GetProviderIds() map[string]string {
	if o == nil || IsNil(o.ProviderIds) {
		var ret map[string]string
		return ret
	}
	return *o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetProviderIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ProviderIds) {
		return nil, false
	}
	return o.ProviderIds, true
}

// HasProviderIds returns a boolean if a field has been set.
func (o *ModelBookInfo) HasProviderIds() bool {
	if o != nil && !IsNil(o.ProviderIds) {
		return true
	}

	return false
}

// SetProviderIds gets a reference to the given map[string]string and assigns it to the ProviderIds field.
func (o *ModelBookInfo) SetProviderIds(v map[string]string) {
	o.ProviderIds = &v
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelBookInfo) GetYear() int32 {
	if o == nil || IsNil(o.Year.Get()) {
		var ret int32
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelBookInfo) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *ModelBookInfo) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableInt32 and assigns it to the Year field.
func (o *ModelBookInfo) SetYear(v int32) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *ModelBookInfo) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *ModelBookInfo) UnsetYear() {
	o.Year.Unset()
}

// GetIndexNumber returns the IndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelBookInfo) GetIndexNumber() int32 {
	if o == nil || IsNil(o.IndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexNumber.Get()
}

// GetIndexNumberOk returns a tuple with the IndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelBookInfo) GetIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexNumber.Get(), o.IndexNumber.IsSet()
}

// HasIndexNumber returns a boolean if a field has been set.
func (o *ModelBookInfo) HasIndexNumber() bool {
	if o != nil && o.IndexNumber.IsSet() {
		return true
	}

	return false
}

// SetIndexNumber gets a reference to the given NullableInt32 and assigns it to the IndexNumber field.
func (o *ModelBookInfo) SetIndexNumber(v int32) {
	o.IndexNumber.Set(&v)
}
// SetIndexNumberNil sets the value for IndexNumber to be an explicit nil
func (o *ModelBookInfo) SetIndexNumberNil() {
	o.IndexNumber.Set(nil)
}

// UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
func (o *ModelBookInfo) UnsetIndexNumber() {
	o.IndexNumber.Unset()
}

// GetParentIndexNumber returns the ParentIndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelBookInfo) GetParentIndexNumber() int32 {
	if o == nil || IsNil(o.ParentIndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentIndexNumber.Get()
}

// GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelBookInfo) GetParentIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentIndexNumber.Get(), o.ParentIndexNumber.IsSet()
}

// HasParentIndexNumber returns a boolean if a field has been set.
func (o *ModelBookInfo) HasParentIndexNumber() bool {
	if o != nil && o.ParentIndexNumber.IsSet() {
		return true
	}

	return false
}

// SetParentIndexNumber gets a reference to the given NullableInt32 and assigns it to the ParentIndexNumber field.
func (o *ModelBookInfo) SetParentIndexNumber(v int32) {
	o.ParentIndexNumber.Set(&v)
}
// SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil
func (o *ModelBookInfo) SetParentIndexNumberNil() {
	o.ParentIndexNumber.Set(nil)
}

// UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
func (o *ModelBookInfo) UnsetParentIndexNumber() {
	o.ParentIndexNumber.Unset()
}

// GetPremiereDate returns the PremiereDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelBookInfo) GetPremiereDate() time.Time {
	if o == nil || IsNil(o.PremiereDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PremiereDate.Get()
}

// GetPremiereDateOk returns a tuple with the PremiereDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelBookInfo) GetPremiereDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiereDate.Get(), o.PremiereDate.IsSet()
}

// HasPremiereDate returns a boolean if a field has been set.
func (o *ModelBookInfo) HasPremiereDate() bool {
	if o != nil && o.PremiereDate.IsSet() {
		return true
	}

	return false
}

// SetPremiereDate gets a reference to the given NullableTime and assigns it to the PremiereDate field.
func (o *ModelBookInfo) SetPremiereDate(v time.Time) {
	o.PremiereDate.Set(&v)
}
// SetPremiereDateNil sets the value for PremiereDate to be an explicit nil
func (o *ModelBookInfo) SetPremiereDateNil() {
	o.PremiereDate.Set(nil)
}

// UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
func (o *ModelBookInfo) UnsetPremiereDate() {
	o.PremiereDate.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise.
func (o *ModelBookInfo) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated) {
		var ret bool
		return ret
	}
	return *o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetIsAutomatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomated) {
		return nil, false
	}
	return o.IsAutomated, true
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *ModelBookInfo) HasIsAutomated() bool {
	if o != nil && !IsNil(o.IsAutomated) {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given bool and assigns it to the IsAutomated field.
func (o *ModelBookInfo) SetIsAutomated(v bool) {
	o.IsAutomated = &v
}

// GetEnableAdultMetadata returns the EnableAdultMetadata field value if set, zero value otherwise.
func (o *ModelBookInfo) GetEnableAdultMetadata() bool {
	if o == nil || IsNil(o.EnableAdultMetadata) {
		var ret bool
		return ret
	}
	return *o.EnableAdultMetadata
}

// GetEnableAdultMetadataOk returns a tuple with the EnableAdultMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBookInfo) GetEnableAdultMetadataOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAdultMetadata) {
		return nil, false
	}
	return o.EnableAdultMetadata, true
}

// HasEnableAdultMetadata returns a boolean if a field has been set.
func (o *ModelBookInfo) HasEnableAdultMetadata() bool {
	if o != nil && !IsNil(o.EnableAdultMetadata) {
		return true
	}

	return false
}

// SetEnableAdultMetadata gets a reference to the given bool and assigns it to the EnableAdultMetadata field.
func (o *ModelBookInfo) SetEnableAdultMetadata(v bool) {
	o.EnableAdultMetadata = &v
}

func (o ModelBookInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBookInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SeriesName) {
		toSerialize["SeriesName"] = o.SeriesName
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.MetadataLanguage) {
		toSerialize["MetadataLanguage"] = o.MetadataLanguage
	}
	if !IsNil(o.MetadataCountryCode) {
		toSerialize["MetadataCountryCode"] = o.MetadataCountryCode
	}
	if !IsNil(o.MetadataLanguages) {
		toSerialize["MetadataLanguages"] = o.MetadataLanguages
	}
	if !IsNil(o.ProviderIds) {
		toSerialize["ProviderIds"] = o.ProviderIds
	}
	if o.Year.IsSet() {
		toSerialize["Year"] = o.Year.Get()
	}
	if o.IndexNumber.IsSet() {
		toSerialize["IndexNumber"] = o.IndexNumber.Get()
	}
	if o.ParentIndexNumber.IsSet() {
		toSerialize["ParentIndexNumber"] = o.ParentIndexNumber.Get()
	}
	if o.PremiereDate.IsSet() {
		toSerialize["PremiereDate"] = o.PremiereDate.Get()
	}
	if !IsNil(o.IsAutomated) {
		toSerialize["IsAutomated"] = o.IsAutomated
	}
	if !IsNil(o.EnableAdultMetadata) {
		toSerialize["EnableAdultMetadata"] = o.EnableAdultMetadata
	}
	return toSerialize, nil
}

type NullableModelBookInfo struct {
	value *ModelBookInfo
	isSet bool
}

func (v NullableModelBookInfo) Get() *ModelBookInfo {
	return v.value
}

func (v *NullableModelBookInfo) Set(val *ModelBookInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBookInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBookInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBookInfo(val *ModelBookInfo) *NullableModelBookInfo {
	return &NullableModelBookInfo{value: val, isSet: true}
}

func (v NullableModelBookInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBookInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

