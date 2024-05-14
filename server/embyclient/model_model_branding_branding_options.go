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

// checks if the ModelBrandingBrandingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBrandingBrandingOptions{}

// ModelBrandingBrandingOptions struct for ModelBrandingBrandingOptions
type ModelBrandingBrandingOptions struct {
	LoginDisclaimer *string `json:"LoginDisclaimer,omitempty"`
	CustomCss *string `json:"CustomCss,omitempty"`
}

// NewModelBrandingBrandingOptions instantiates a new ModelBrandingBrandingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBrandingBrandingOptions() *ModelBrandingBrandingOptions {
	this := ModelBrandingBrandingOptions{}
	return &this
}

// NewModelBrandingBrandingOptionsWithDefaults instantiates a new ModelBrandingBrandingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBrandingBrandingOptionsWithDefaults() *ModelBrandingBrandingOptions {
	this := ModelBrandingBrandingOptions{}
	return &this
}

// GetLoginDisclaimer returns the LoginDisclaimer field value if set, zero value otherwise.
func (o *ModelBrandingBrandingOptions) GetLoginDisclaimer() string {
	if o == nil || IsNil(o.LoginDisclaimer) {
		var ret string
		return ret
	}
	return *o.LoginDisclaimer
}

// GetLoginDisclaimerOk returns a tuple with the LoginDisclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBrandingBrandingOptions) GetLoginDisclaimerOk() (*string, bool) {
	if o == nil || IsNil(o.LoginDisclaimer) {
		return nil, false
	}
	return o.LoginDisclaimer, true
}

// HasLoginDisclaimer returns a boolean if a field has been set.
func (o *ModelBrandingBrandingOptions) HasLoginDisclaimer() bool {
	if o != nil && !IsNil(o.LoginDisclaimer) {
		return true
	}

	return false
}

// SetLoginDisclaimer gets a reference to the given string and assigns it to the LoginDisclaimer field.
func (o *ModelBrandingBrandingOptions) SetLoginDisclaimer(v string) {
	o.LoginDisclaimer = &v
}

// GetCustomCss returns the CustomCss field value if set, zero value otherwise.
func (o *ModelBrandingBrandingOptions) GetCustomCss() string {
	if o == nil || IsNil(o.CustomCss) {
		var ret string
		return ret
	}
	return *o.CustomCss
}

// GetCustomCssOk returns a tuple with the CustomCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBrandingBrandingOptions) GetCustomCssOk() (*string, bool) {
	if o == nil || IsNil(o.CustomCss) {
		return nil, false
	}
	return o.CustomCss, true
}

// HasCustomCss returns a boolean if a field has been set.
func (o *ModelBrandingBrandingOptions) HasCustomCss() bool {
	if o != nil && !IsNil(o.CustomCss) {
		return true
	}

	return false
}

// SetCustomCss gets a reference to the given string and assigns it to the CustomCss field.
func (o *ModelBrandingBrandingOptions) SetCustomCss(v string) {
	o.CustomCss = &v
}

func (o ModelBrandingBrandingOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBrandingBrandingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LoginDisclaimer) {
		toSerialize["LoginDisclaimer"] = o.LoginDisclaimer
	}
	if !IsNil(o.CustomCss) {
		toSerialize["CustomCss"] = o.CustomCss
	}
	return toSerialize, nil
}

type NullableModelBrandingBrandingOptions struct {
	value *ModelBrandingBrandingOptions
	isSet bool
}

func (v NullableModelBrandingBrandingOptions) Get() *ModelBrandingBrandingOptions {
	return v.value
}

func (v *NullableModelBrandingBrandingOptions) Set(val *ModelBrandingBrandingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBrandingBrandingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBrandingBrandingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBrandingBrandingOptions(val *ModelBrandingBrandingOptions) *NullableModelBrandingBrandingOptions {
	return &NullableModelBrandingBrandingOptions{value: val, isSet: true}
}

func (v NullableModelBrandingBrandingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBrandingBrandingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


