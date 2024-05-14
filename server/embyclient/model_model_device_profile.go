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

// checks if the ModelDeviceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDeviceProfile{}

// ModelDeviceProfile struct for ModelDeviceProfile
type ModelDeviceProfile struct {
	Name *string `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
	SupportedMediaTypes *string `json:"SupportedMediaTypes,omitempty"`
	MaxStreamingBitrate NullableInt64 `json:"MaxStreamingBitrate,omitempty"`
	MusicStreamingTranscodingBitrate NullableInt32 `json:"MusicStreamingTranscodingBitrate,omitempty"`
	MaxStaticMusicBitrate NullableInt32 `json:"MaxStaticMusicBitrate,omitempty"`
	DirectPlayProfiles []ModelDirectPlayProfile `json:"DirectPlayProfiles,omitempty"`
	TranscodingProfiles []ModelTranscodingProfile `json:"TranscodingProfiles,omitempty"`
	ContainerProfiles []ModelContainerProfile `json:"ContainerProfiles,omitempty"`
	CodecProfiles []ModelCodecProfile `json:"CodecProfiles,omitempty"`
	ResponseProfiles []ModelResponseProfile `json:"ResponseProfiles,omitempty"`
	SubtitleProfiles []ModelSubtitleProfile `json:"SubtitleProfiles,omitempty"`
}

// NewModelDeviceProfile instantiates a new ModelDeviceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDeviceProfile() *ModelDeviceProfile {
	this := ModelDeviceProfile{}
	return &this
}

// NewModelDeviceProfileWithDefaults instantiates a new ModelDeviceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDeviceProfileWithDefaults() *ModelDeviceProfile {
	this := ModelDeviceProfile{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelDeviceProfile) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelDeviceProfile) SetId(v string) {
	o.Id = &v
}

// GetSupportedMediaTypes returns the SupportedMediaTypes field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetSupportedMediaTypes() string {
	if o == nil || IsNil(o.SupportedMediaTypes) {
		var ret string
		return ret
	}
	return *o.SupportedMediaTypes
}

// GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetSupportedMediaTypesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedMediaTypes) {
		return nil, false
	}
	return o.SupportedMediaTypes, true
}

// HasSupportedMediaTypes returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasSupportedMediaTypes() bool {
	if o != nil && !IsNil(o.SupportedMediaTypes) {
		return true
	}

	return false
}

// SetSupportedMediaTypes gets a reference to the given string and assigns it to the SupportedMediaTypes field.
func (o *ModelDeviceProfile) SetSupportedMediaTypes(v string) {
	o.SupportedMediaTypes = &v
}

// GetMaxStreamingBitrate returns the MaxStreamingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelDeviceProfile) GetMaxStreamingBitrate() int64 {
	if o == nil || IsNil(o.MaxStreamingBitrate.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxStreamingBitrate.Get()
}

// GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDeviceProfile) GetMaxStreamingBitrateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStreamingBitrate.Get(), o.MaxStreamingBitrate.IsSet()
}

// HasMaxStreamingBitrate returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasMaxStreamingBitrate() bool {
	if o != nil && o.MaxStreamingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStreamingBitrate gets a reference to the given NullableInt64 and assigns it to the MaxStreamingBitrate field.
func (o *ModelDeviceProfile) SetMaxStreamingBitrate(v int64) {
	o.MaxStreamingBitrate.Set(&v)
}
// SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil
func (o *ModelDeviceProfile) SetMaxStreamingBitrateNil() {
	o.MaxStreamingBitrate.Set(nil)
}

// UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
func (o *ModelDeviceProfile) UnsetMaxStreamingBitrate() {
	o.MaxStreamingBitrate.Unset()
}

// GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelDeviceProfile) GetMusicStreamingTranscodingBitrate() int32 {
	if o == nil || IsNil(o.MusicStreamingTranscodingBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MusicStreamingTranscodingBitrate.Get()
}

// GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MusicStreamingTranscodingBitrate.Get(), o.MusicStreamingTranscodingBitrate.IsSet()
}

// HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasMusicStreamingTranscodingBitrate() bool {
	if o != nil && o.MusicStreamingTranscodingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMusicStreamingTranscodingBitrate gets a reference to the given NullableInt32 and assigns it to the MusicStreamingTranscodingBitrate field.
func (o *ModelDeviceProfile) SetMusicStreamingTranscodingBitrate(v int32) {
	o.MusicStreamingTranscodingBitrate.Set(&v)
}
// SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil
func (o *ModelDeviceProfile) SetMusicStreamingTranscodingBitrateNil() {
	o.MusicStreamingTranscodingBitrate.Set(nil)
}

// UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
func (o *ModelDeviceProfile) UnsetMusicStreamingTranscodingBitrate() {
	o.MusicStreamingTranscodingBitrate.Unset()
}

// GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelDeviceProfile) GetMaxStaticMusicBitrate() int32 {
	if o == nil || IsNil(o.MaxStaticMusicBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStaticMusicBitrate.Get()
}

// GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStaticMusicBitrate.Get(), o.MaxStaticMusicBitrate.IsSet()
}

// HasMaxStaticMusicBitrate returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasMaxStaticMusicBitrate() bool {
	if o != nil && o.MaxStaticMusicBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStaticMusicBitrate gets a reference to the given NullableInt32 and assigns it to the MaxStaticMusicBitrate field.
func (o *ModelDeviceProfile) SetMaxStaticMusicBitrate(v int32) {
	o.MaxStaticMusicBitrate.Set(&v)
}
// SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil
func (o *ModelDeviceProfile) SetMaxStaticMusicBitrateNil() {
	o.MaxStaticMusicBitrate.Set(nil)
}

// UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
func (o *ModelDeviceProfile) UnsetMaxStaticMusicBitrate() {
	o.MaxStaticMusicBitrate.Unset()
}

// GetDirectPlayProfiles returns the DirectPlayProfiles field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetDirectPlayProfiles() []ModelDirectPlayProfile {
	if o == nil || IsNil(o.DirectPlayProfiles) {
		var ret []ModelDirectPlayProfile
		return ret
	}
	return o.DirectPlayProfiles
}

// GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetDirectPlayProfilesOk() ([]ModelDirectPlayProfile, bool) {
	if o == nil || IsNil(o.DirectPlayProfiles) {
		return nil, false
	}
	return o.DirectPlayProfiles, true
}

// HasDirectPlayProfiles returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasDirectPlayProfiles() bool {
	if o != nil && !IsNil(o.DirectPlayProfiles) {
		return true
	}

	return false
}

// SetDirectPlayProfiles gets a reference to the given []ModelDirectPlayProfile and assigns it to the DirectPlayProfiles field.
func (o *ModelDeviceProfile) SetDirectPlayProfiles(v []ModelDirectPlayProfile) {
	o.DirectPlayProfiles = v
}

// GetTranscodingProfiles returns the TranscodingProfiles field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetTranscodingProfiles() []ModelTranscodingProfile {
	if o == nil || IsNil(o.TranscodingProfiles) {
		var ret []ModelTranscodingProfile
		return ret
	}
	return o.TranscodingProfiles
}

// GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetTranscodingProfilesOk() ([]ModelTranscodingProfile, bool) {
	if o == nil || IsNil(o.TranscodingProfiles) {
		return nil, false
	}
	return o.TranscodingProfiles, true
}

// HasTranscodingProfiles returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasTranscodingProfiles() bool {
	if o != nil && !IsNil(o.TranscodingProfiles) {
		return true
	}

	return false
}

// SetTranscodingProfiles gets a reference to the given []ModelTranscodingProfile and assigns it to the TranscodingProfiles field.
func (o *ModelDeviceProfile) SetTranscodingProfiles(v []ModelTranscodingProfile) {
	o.TranscodingProfiles = v
}

// GetContainerProfiles returns the ContainerProfiles field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetContainerProfiles() []ModelContainerProfile {
	if o == nil || IsNil(o.ContainerProfiles) {
		var ret []ModelContainerProfile
		return ret
	}
	return o.ContainerProfiles
}

// GetContainerProfilesOk returns a tuple with the ContainerProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetContainerProfilesOk() ([]ModelContainerProfile, bool) {
	if o == nil || IsNil(o.ContainerProfiles) {
		return nil, false
	}
	return o.ContainerProfiles, true
}

// HasContainerProfiles returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasContainerProfiles() bool {
	if o != nil && !IsNil(o.ContainerProfiles) {
		return true
	}

	return false
}

// SetContainerProfiles gets a reference to the given []ModelContainerProfile and assigns it to the ContainerProfiles field.
func (o *ModelDeviceProfile) SetContainerProfiles(v []ModelContainerProfile) {
	o.ContainerProfiles = v
}

// GetCodecProfiles returns the CodecProfiles field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetCodecProfiles() []ModelCodecProfile {
	if o == nil || IsNil(o.CodecProfiles) {
		var ret []ModelCodecProfile
		return ret
	}
	return o.CodecProfiles
}

// GetCodecProfilesOk returns a tuple with the CodecProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetCodecProfilesOk() ([]ModelCodecProfile, bool) {
	if o == nil || IsNil(o.CodecProfiles) {
		return nil, false
	}
	return o.CodecProfiles, true
}

// HasCodecProfiles returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasCodecProfiles() bool {
	if o != nil && !IsNil(o.CodecProfiles) {
		return true
	}

	return false
}

// SetCodecProfiles gets a reference to the given []ModelCodecProfile and assigns it to the CodecProfiles field.
func (o *ModelDeviceProfile) SetCodecProfiles(v []ModelCodecProfile) {
	o.CodecProfiles = v
}

// GetResponseProfiles returns the ResponseProfiles field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetResponseProfiles() []ModelResponseProfile {
	if o == nil || IsNil(o.ResponseProfiles) {
		var ret []ModelResponseProfile
		return ret
	}
	return o.ResponseProfiles
}

// GetResponseProfilesOk returns a tuple with the ResponseProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetResponseProfilesOk() ([]ModelResponseProfile, bool) {
	if o == nil || IsNil(o.ResponseProfiles) {
		return nil, false
	}
	return o.ResponseProfiles, true
}

// HasResponseProfiles returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasResponseProfiles() bool {
	if o != nil && !IsNil(o.ResponseProfiles) {
		return true
	}

	return false
}

// SetResponseProfiles gets a reference to the given []ModelResponseProfile and assigns it to the ResponseProfiles field.
func (o *ModelDeviceProfile) SetResponseProfiles(v []ModelResponseProfile) {
	o.ResponseProfiles = v
}

// GetSubtitleProfiles returns the SubtitleProfiles field value if set, zero value otherwise.
func (o *ModelDeviceProfile) GetSubtitleProfiles() []ModelSubtitleProfile {
	if o == nil || IsNil(o.SubtitleProfiles) {
		var ret []ModelSubtitleProfile
		return ret
	}
	return o.SubtitleProfiles
}

// GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDeviceProfile) GetSubtitleProfilesOk() ([]ModelSubtitleProfile, bool) {
	if o == nil || IsNil(o.SubtitleProfiles) {
		return nil, false
	}
	return o.SubtitleProfiles, true
}

// HasSubtitleProfiles returns a boolean if a field has been set.
func (o *ModelDeviceProfile) HasSubtitleProfiles() bool {
	if o != nil && !IsNil(o.SubtitleProfiles) {
		return true
	}

	return false
}

// SetSubtitleProfiles gets a reference to the given []ModelSubtitleProfile and assigns it to the SubtitleProfiles field.
func (o *ModelDeviceProfile) SetSubtitleProfiles(v []ModelSubtitleProfile) {
	o.SubtitleProfiles = v
}

func (o ModelDeviceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDeviceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.SupportedMediaTypes) {
		toSerialize["SupportedMediaTypes"] = o.SupportedMediaTypes
	}
	if o.MaxStreamingBitrate.IsSet() {
		toSerialize["MaxStreamingBitrate"] = o.MaxStreamingBitrate.Get()
	}
	if o.MusicStreamingTranscodingBitrate.IsSet() {
		toSerialize["MusicStreamingTranscodingBitrate"] = o.MusicStreamingTranscodingBitrate.Get()
	}
	if o.MaxStaticMusicBitrate.IsSet() {
		toSerialize["MaxStaticMusicBitrate"] = o.MaxStaticMusicBitrate.Get()
	}
	if !IsNil(o.DirectPlayProfiles) {
		toSerialize["DirectPlayProfiles"] = o.DirectPlayProfiles
	}
	if !IsNil(o.TranscodingProfiles) {
		toSerialize["TranscodingProfiles"] = o.TranscodingProfiles
	}
	if !IsNil(o.ContainerProfiles) {
		toSerialize["ContainerProfiles"] = o.ContainerProfiles
	}
	if !IsNil(o.CodecProfiles) {
		toSerialize["CodecProfiles"] = o.CodecProfiles
	}
	if !IsNil(o.ResponseProfiles) {
		toSerialize["ResponseProfiles"] = o.ResponseProfiles
	}
	if !IsNil(o.SubtitleProfiles) {
		toSerialize["SubtitleProfiles"] = o.SubtitleProfiles
	}
	return toSerialize, nil
}

type NullableModelDeviceProfile struct {
	value *ModelDeviceProfile
	isSet bool
}

func (v NullableModelDeviceProfile) Get() *ModelDeviceProfile {
	return v.value
}

func (v *NullableModelDeviceProfile) Set(val *ModelDeviceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDeviceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDeviceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDeviceProfile(val *ModelDeviceProfile) *NullableModelDeviceProfile {
	return &NullableModelDeviceProfile{value: val, isSet: true}
}

func (v NullableModelDeviceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDeviceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

