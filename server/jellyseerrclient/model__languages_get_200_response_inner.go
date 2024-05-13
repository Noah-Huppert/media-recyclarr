/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyseerrclient

import (
	"encoding/json"
)

// checks if the LanguagesGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguagesGet200ResponseInner{}

// LanguagesGet200ResponseInner struct for LanguagesGet200ResponseInner
type LanguagesGet200ResponseInner struct {
	Iso6391 *string `json:"iso_639_1,omitempty"`
	EnglishName *string `json:"english_name,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewLanguagesGet200ResponseInner instantiates a new LanguagesGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguagesGet200ResponseInner() *LanguagesGet200ResponseInner {
	this := LanguagesGet200ResponseInner{}
	return &this
}

// NewLanguagesGet200ResponseInnerWithDefaults instantiates a new LanguagesGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguagesGet200ResponseInnerWithDefaults() *LanguagesGet200ResponseInner {
	this := LanguagesGet200ResponseInner{}
	return &this
}

// GetIso6391 returns the Iso6391 field value if set, zero value otherwise.
func (o *LanguagesGet200ResponseInner) GetIso6391() string {
	if o == nil || IsNil(o.Iso6391) {
		var ret string
		return ret
	}
	return *o.Iso6391
}

// GetIso6391Ok returns a tuple with the Iso6391 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguagesGet200ResponseInner) GetIso6391Ok() (*string, bool) {
	if o == nil || IsNil(o.Iso6391) {
		return nil, false
	}
	return o.Iso6391, true
}

// HasIso6391 returns a boolean if a field has been set.
func (o *LanguagesGet200ResponseInner) HasIso6391() bool {
	if o != nil && !IsNil(o.Iso6391) {
		return true
	}

	return false
}

// SetIso6391 gets a reference to the given string and assigns it to the Iso6391 field.
func (o *LanguagesGet200ResponseInner) SetIso6391(v string) {
	o.Iso6391 = &v
}

// GetEnglishName returns the EnglishName field value if set, zero value otherwise.
func (o *LanguagesGet200ResponseInner) GetEnglishName() string {
	if o == nil || IsNil(o.EnglishName) {
		var ret string
		return ret
	}
	return *o.EnglishName
}

// GetEnglishNameOk returns a tuple with the EnglishName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguagesGet200ResponseInner) GetEnglishNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnglishName) {
		return nil, false
	}
	return o.EnglishName, true
}

// HasEnglishName returns a boolean if a field has been set.
func (o *LanguagesGet200ResponseInner) HasEnglishName() bool {
	if o != nil && !IsNil(o.EnglishName) {
		return true
	}

	return false
}

// SetEnglishName gets a reference to the given string and assigns it to the EnglishName field.
func (o *LanguagesGet200ResponseInner) SetEnglishName(v string) {
	o.EnglishName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LanguagesGet200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LanguagesGet200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LanguagesGet200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LanguagesGet200ResponseInner) SetName(v string) {
	o.Name = &v
}

func (o LanguagesGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguagesGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Iso6391) {
		toSerialize["iso_639_1"] = o.Iso6391
	}
	if !IsNil(o.EnglishName) {
		toSerialize["english_name"] = o.EnglishName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLanguagesGet200ResponseInner struct {
	value *LanguagesGet200ResponseInner
	isSet bool
}

func (v NullableLanguagesGet200ResponseInner) Get() *LanguagesGet200ResponseInner {
	return v.value
}

func (v *NullableLanguagesGet200ResponseInner) Set(val *LanguagesGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguagesGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguagesGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguagesGet200ResponseInner(val *LanguagesGet200ResponseInner) *NullableLanguagesGet200ResponseInner {
	return &NullableLanguagesGet200ResponseInner{value: val, isSet: true}
}

func (v NullableLanguagesGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguagesGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


