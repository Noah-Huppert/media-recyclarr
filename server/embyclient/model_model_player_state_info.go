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

// checks if the ModelPlayerStateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPlayerStateInfo{}

// ModelPlayerStateInfo struct for ModelPlayerStateInfo
type ModelPlayerStateInfo struct {
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	CanSeek *bool `json:"CanSeek,omitempty"`
	IsPaused *bool `json:"IsPaused,omitempty"`
	IsMuted *bool `json:"IsMuted,omitempty"`
	VolumeLevel NullableInt32 `json:"VolumeLevel,omitempty"`
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	MediaSourceId *string `json:"MediaSourceId,omitempty"`
	PlayMethod *ModelPlayMethod `json:"PlayMethod,omitempty"`
	RepeatMode *ModelRepeatMode `json:"RepeatMode,omitempty"`
	SubtitleOffset *int32 `json:"SubtitleOffset,omitempty"`
	PlaybackRate *float64 `json:"PlaybackRate,omitempty"`
}

// NewModelPlayerStateInfo instantiates a new ModelPlayerStateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPlayerStateInfo() *ModelPlayerStateInfo {
	this := ModelPlayerStateInfo{}
	return &this
}

// NewModelPlayerStateInfoWithDefaults instantiates a new ModelPlayerStateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPlayerStateInfoWithDefaults() *ModelPlayerStateInfo {
	this := ModelPlayerStateInfo{}
	return &this
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlayerStateInfo) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlayerStateInfo) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *ModelPlayerStateInfo) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *ModelPlayerStateInfo) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *ModelPlayerStateInfo) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetCanSeek returns the CanSeek field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetCanSeek() bool {
	if o == nil || IsNil(o.CanSeek) {
		var ret bool
		return ret
	}
	return *o.CanSeek
}

// GetCanSeekOk returns a tuple with the CanSeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetCanSeekOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSeek) {
		return nil, false
	}
	return o.CanSeek, true
}

// HasCanSeek returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasCanSeek() bool {
	if o != nil && !IsNil(o.CanSeek) {
		return true
	}

	return false
}

// SetCanSeek gets a reference to the given bool and assigns it to the CanSeek field.
func (o *ModelPlayerStateInfo) SetCanSeek(v bool) {
	o.CanSeek = &v
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused) {
		var ret bool
		return ret
	}
	return *o.IsPaused
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetIsPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaused) {
		return nil, false
	}
	return o.IsPaused, true
}

// HasIsPaused returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasIsPaused() bool {
	if o != nil && !IsNil(o.IsPaused) {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given bool and assigns it to the IsPaused field.
func (o *ModelPlayerStateInfo) SetIsPaused(v bool) {
	o.IsPaused = &v
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *ModelPlayerStateInfo) SetIsMuted(v bool) {
	o.IsMuted = &v
}

// GetVolumeLevel returns the VolumeLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlayerStateInfo) GetVolumeLevel() int32 {
	if o == nil || IsNil(o.VolumeLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.VolumeLevel.Get()
}

// GetVolumeLevelOk returns a tuple with the VolumeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlayerStateInfo) GetVolumeLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeLevel.Get(), o.VolumeLevel.IsSet()
}

// HasVolumeLevel returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasVolumeLevel() bool {
	if o != nil && o.VolumeLevel.IsSet() {
		return true
	}

	return false
}

// SetVolumeLevel gets a reference to the given NullableInt32 and assigns it to the VolumeLevel field.
func (o *ModelPlayerStateInfo) SetVolumeLevel(v int32) {
	o.VolumeLevel.Set(&v)
}
// SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil
func (o *ModelPlayerStateInfo) SetVolumeLevelNil() {
	o.VolumeLevel.Set(nil)
}

// UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
func (o *ModelPlayerStateInfo) UnsetVolumeLevel() {
	o.VolumeLevel.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlayerStateInfo) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlayerStateInfo) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *ModelPlayerStateInfo) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *ModelPlayerStateInfo) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *ModelPlayerStateInfo) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlayerStateInfo) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlayerStateInfo) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *ModelPlayerStateInfo) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *ModelPlayerStateInfo) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *ModelPlayerStateInfo) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId) {
		var ret string
		return ret
	}
	return *o.MediaSourceId
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetMediaSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaSourceId) {
		return nil, false
	}
	return o.MediaSourceId, true
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasMediaSourceId() bool {
	if o != nil && !IsNil(o.MediaSourceId) {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given string and assigns it to the MediaSourceId field.
func (o *ModelPlayerStateInfo) SetMediaSourceId(v string) {
	o.MediaSourceId = &v
}

// GetPlayMethod returns the PlayMethod field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetPlayMethod() ModelPlayMethod {
	if o == nil || IsNil(o.PlayMethod) {
		var ret ModelPlayMethod
		return ret
	}
	return *o.PlayMethod
}

// GetPlayMethodOk returns a tuple with the PlayMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetPlayMethodOk() (*ModelPlayMethod, bool) {
	if o == nil || IsNil(o.PlayMethod) {
		return nil, false
	}
	return o.PlayMethod, true
}

// HasPlayMethod returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasPlayMethod() bool {
	if o != nil && !IsNil(o.PlayMethod) {
		return true
	}

	return false
}

// SetPlayMethod gets a reference to the given ModelPlayMethod and assigns it to the PlayMethod field.
func (o *ModelPlayerStateInfo) SetPlayMethod(v ModelPlayMethod) {
	o.PlayMethod = &v
}

// GetRepeatMode returns the RepeatMode field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetRepeatMode() ModelRepeatMode {
	if o == nil || IsNil(o.RepeatMode) {
		var ret ModelRepeatMode
		return ret
	}
	return *o.RepeatMode
}

// GetRepeatModeOk returns a tuple with the RepeatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetRepeatModeOk() (*ModelRepeatMode, bool) {
	if o == nil || IsNil(o.RepeatMode) {
		return nil, false
	}
	return o.RepeatMode, true
}

// HasRepeatMode returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasRepeatMode() bool {
	if o != nil && !IsNil(o.RepeatMode) {
		return true
	}

	return false
}

// SetRepeatMode gets a reference to the given ModelRepeatMode and assigns it to the RepeatMode field.
func (o *ModelPlayerStateInfo) SetRepeatMode(v ModelRepeatMode) {
	o.RepeatMode = &v
}

// GetSubtitleOffset returns the SubtitleOffset field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetSubtitleOffset() int32 {
	if o == nil || IsNil(o.SubtitleOffset) {
		var ret int32
		return ret
	}
	return *o.SubtitleOffset
}

// GetSubtitleOffsetOk returns a tuple with the SubtitleOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetSubtitleOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.SubtitleOffset) {
		return nil, false
	}
	return o.SubtitleOffset, true
}

// HasSubtitleOffset returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasSubtitleOffset() bool {
	if o != nil && !IsNil(o.SubtitleOffset) {
		return true
	}

	return false
}

// SetSubtitleOffset gets a reference to the given int32 and assigns it to the SubtitleOffset field.
func (o *ModelPlayerStateInfo) SetSubtitleOffset(v int32) {
	o.SubtitleOffset = &v
}

// GetPlaybackRate returns the PlaybackRate field value if set, zero value otherwise.
func (o *ModelPlayerStateInfo) GetPlaybackRate() float64 {
	if o == nil || IsNil(o.PlaybackRate) {
		var ret float64
		return ret
	}
	return *o.PlaybackRate
}

// GetPlaybackRateOk returns a tuple with the PlaybackRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlayerStateInfo) GetPlaybackRateOk() (*float64, bool) {
	if o == nil || IsNil(o.PlaybackRate) {
		return nil, false
	}
	return o.PlaybackRate, true
}

// HasPlaybackRate returns a boolean if a field has been set.
func (o *ModelPlayerStateInfo) HasPlaybackRate() bool {
	if o != nil && !IsNil(o.PlaybackRate) {
		return true
	}

	return false
}

// SetPlaybackRate gets a reference to the given float64 and assigns it to the PlaybackRate field.
func (o *ModelPlayerStateInfo) SetPlaybackRate(v float64) {
	o.PlaybackRate = &v
}

func (o ModelPlayerStateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPlayerStateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if !IsNil(o.CanSeek) {
		toSerialize["CanSeek"] = o.CanSeek
	}
	if !IsNil(o.IsPaused) {
		toSerialize["IsPaused"] = o.IsPaused
	}
	if !IsNil(o.IsMuted) {
		toSerialize["IsMuted"] = o.IsMuted
	}
	if o.VolumeLevel.IsSet() {
		toSerialize["VolumeLevel"] = o.VolumeLevel.Get()
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if !IsNil(o.MediaSourceId) {
		toSerialize["MediaSourceId"] = o.MediaSourceId
	}
	if !IsNil(o.PlayMethod) {
		toSerialize["PlayMethod"] = o.PlayMethod
	}
	if !IsNil(o.RepeatMode) {
		toSerialize["RepeatMode"] = o.RepeatMode
	}
	if !IsNil(o.SubtitleOffset) {
		toSerialize["SubtitleOffset"] = o.SubtitleOffset
	}
	if !IsNil(o.PlaybackRate) {
		toSerialize["PlaybackRate"] = o.PlaybackRate
	}
	return toSerialize, nil
}

type NullableModelPlayerStateInfo struct {
	value *ModelPlayerStateInfo
	isSet bool
}

func (v NullableModelPlayerStateInfo) Get() *ModelPlayerStateInfo {
	return v.value
}

func (v *NullableModelPlayerStateInfo) Set(val *ModelPlayerStateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPlayerStateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPlayerStateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPlayerStateInfo(val *ModelPlayerStateInfo) *NullableModelPlayerStateInfo {
	return &NullableModelPlayerStateInfo{value: val, isSet: true}
}

func (v NullableModelPlayerStateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPlayerStateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

