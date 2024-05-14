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

// checks if the ModelResponseProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelResponseProfile{}

// ModelResponseProfile struct for ModelResponseProfile
type ModelResponseProfile struct {
	Container *string `json:"Container,omitempty"`
	AudioCodec *string `json:"AudioCodec,omitempty"`
	VideoCodec *string `json:"VideoCodec,omitempty"`
	Type *ModelDlnaProfileType `json:"Type,omitempty"`
	OrgPn *string `json:"OrgPn,omitempty"`
	MimeType *string `json:"MimeType,omitempty"`
	Conditions []ModelProfileCondition `json:"Conditions,omitempty"`
}

// NewModelResponseProfile instantiates a new ModelResponseProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelResponseProfile() *ModelResponseProfile {
	this := ModelResponseProfile{}
	return &this
}

// NewModelResponseProfileWithDefaults instantiates a new ModelResponseProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelResponseProfileWithDefaults() *ModelResponseProfile {
	this := ModelResponseProfile{}
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *ModelResponseProfile) SetContainer(v string) {
	o.Container = &v
}

// GetAudioCodec returns the AudioCodec field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetAudioCodec() string {
	if o == nil || IsNil(o.AudioCodec) {
		var ret string
		return ret
	}
	return *o.AudioCodec
}

// GetAudioCodecOk returns a tuple with the AudioCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetAudioCodecOk() (*string, bool) {
	if o == nil || IsNil(o.AudioCodec) {
		return nil, false
	}
	return o.AudioCodec, true
}

// HasAudioCodec returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasAudioCodec() bool {
	if o != nil && !IsNil(o.AudioCodec) {
		return true
	}

	return false
}

// SetAudioCodec gets a reference to the given string and assigns it to the AudioCodec field.
func (o *ModelResponseProfile) SetAudioCodec(v string) {
	o.AudioCodec = &v
}

// GetVideoCodec returns the VideoCodec field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetVideoCodec() string {
	if o == nil || IsNil(o.VideoCodec) {
		var ret string
		return ret
	}
	return *o.VideoCodec
}

// GetVideoCodecOk returns a tuple with the VideoCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetVideoCodecOk() (*string, bool) {
	if o == nil || IsNil(o.VideoCodec) {
		return nil, false
	}
	return o.VideoCodec, true
}

// HasVideoCodec returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasVideoCodec() bool {
	if o != nil && !IsNil(o.VideoCodec) {
		return true
	}

	return false
}

// SetVideoCodec gets a reference to the given string and assigns it to the VideoCodec field.
func (o *ModelResponseProfile) SetVideoCodec(v string) {
	o.VideoCodec = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetType() ModelDlnaProfileType {
	if o == nil || IsNil(o.Type) {
		var ret ModelDlnaProfileType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetTypeOk() (*ModelDlnaProfileType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModelDlnaProfileType and assigns it to the Type field.
func (o *ModelResponseProfile) SetType(v ModelDlnaProfileType) {
	o.Type = &v
}

// GetOrgPn returns the OrgPn field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetOrgPn() string {
	if o == nil || IsNil(o.OrgPn) {
		var ret string
		return ret
	}
	return *o.OrgPn
}

// GetOrgPnOk returns a tuple with the OrgPn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetOrgPnOk() (*string, bool) {
	if o == nil || IsNil(o.OrgPn) {
		return nil, false
	}
	return o.OrgPn, true
}

// HasOrgPn returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasOrgPn() bool {
	if o != nil && !IsNil(o.OrgPn) {
		return true
	}

	return false
}

// SetOrgPn gets a reference to the given string and assigns it to the OrgPn field.
func (o *ModelResponseProfile) SetOrgPn(v string) {
	o.OrgPn = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *ModelResponseProfile) SetMimeType(v string) {
	o.MimeType = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ModelResponseProfile) GetConditions() []ModelProfileCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []ModelProfileCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelResponseProfile) GetConditionsOk() ([]ModelProfileCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ModelResponseProfile) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ModelProfileCondition and assigns it to the Conditions field.
func (o *ModelResponseProfile) SetConditions(v []ModelProfileCondition) {
	o.Conditions = v
}

func (o ModelResponseProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelResponseProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Container) {
		toSerialize["Container"] = o.Container
	}
	if !IsNil(o.AudioCodec) {
		toSerialize["AudioCodec"] = o.AudioCodec
	}
	if !IsNil(o.VideoCodec) {
		toSerialize["VideoCodec"] = o.VideoCodec
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.OrgPn) {
		toSerialize["OrgPn"] = o.OrgPn
	}
	if !IsNil(o.MimeType) {
		toSerialize["MimeType"] = o.MimeType
	}
	if !IsNil(o.Conditions) {
		toSerialize["Conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableModelResponseProfile struct {
	value *ModelResponseProfile
	isSet bool
}

func (v NullableModelResponseProfile) Get() *ModelResponseProfile {
	return v.value
}

func (v *NullableModelResponseProfile) Set(val *ModelResponseProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableModelResponseProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableModelResponseProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelResponseProfile(val *ModelResponseProfile) *NullableModelResponseProfile {
	return &NullableModelResponseProfile{value: val, isSet: true}
}

func (v NullableModelResponseProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelResponseProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

