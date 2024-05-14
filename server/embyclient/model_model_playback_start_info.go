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

// checks if the ModelPlaybackStartInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPlaybackStartInfo{}

// ModelPlaybackStartInfo struct for ModelPlaybackStartInfo
type ModelPlaybackStartInfo struct {
	CanSeek *bool `json:"CanSeek,omitempty"`
	Item *ModelBaseItemDto `json:"Item,omitempty"`
	NowPlayingQueue []ModelQueueItem `json:"NowPlayingQueue,omitempty"`
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	SessionId *string `json:"SessionId,omitempty"`
	MediaSourceId *string `json:"MediaSourceId,omitempty"`
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	IsPaused *bool `json:"IsPaused,omitempty"`
	PlaylistIndex *int32 `json:"PlaylistIndex,omitempty"`
	PlaylistLength *int32 `json:"PlaylistLength,omitempty"`
	IsMuted *bool `json:"IsMuted,omitempty"`
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	RunTimeTicks NullableInt64 `json:"RunTimeTicks,omitempty"`
	PlaybackStartTimeTicks NullableInt64 `json:"PlaybackStartTimeTicks,omitempty"`
	VolumeLevel NullableInt32 `json:"VolumeLevel,omitempty"`
	Brightness NullableInt32 `json:"Brightness,omitempty"`
	AspectRatio *string `json:"AspectRatio,omitempty"`
	EventName *ModelProgressEvent `json:"EventName,omitempty"`
	PlayMethod *ModelPlayMethod `json:"PlayMethod,omitempty"`
	LiveStreamId *string `json:"LiveStreamId,omitempty"`
	PlaySessionId *string `json:"PlaySessionId,omitempty"`
	RepeatMode *ModelRepeatMode `json:"RepeatMode,omitempty"`
	SubtitleOffset *int32 `json:"SubtitleOffset,omitempty"`
	PlaybackRate *float64 `json:"PlaybackRate,omitempty"`
	PlaylistItemIds []string `json:"PlaylistItemIds,omitempty"`
}

// NewModelPlaybackStartInfo instantiates a new ModelPlaybackStartInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPlaybackStartInfo() *ModelPlaybackStartInfo {
	this := ModelPlaybackStartInfo{}
	return &this
}

// NewModelPlaybackStartInfoWithDefaults instantiates a new ModelPlaybackStartInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPlaybackStartInfoWithDefaults() *ModelPlaybackStartInfo {
	this := ModelPlaybackStartInfo{}
	return &this
}

// GetCanSeek returns the CanSeek field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetCanSeek() bool {
	if o == nil || IsNil(o.CanSeek) {
		var ret bool
		return ret
	}
	return *o.CanSeek
}

// GetCanSeekOk returns a tuple with the CanSeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetCanSeekOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSeek) {
		return nil, false
	}
	return o.CanSeek, true
}

// HasCanSeek returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasCanSeek() bool {
	if o != nil && !IsNil(o.CanSeek) {
		return true
	}

	return false
}

// SetCanSeek gets a reference to the given bool and assigns it to the CanSeek field.
func (o *ModelPlaybackStartInfo) SetCanSeek(v bool) {
	o.CanSeek = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetItem() ModelBaseItemDto {
	if o == nil || IsNil(o.Item) {
		var ret ModelBaseItemDto
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetItemOk() (*ModelBaseItemDto, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given ModelBaseItemDto and assigns it to the Item field.
func (o *ModelPlaybackStartInfo) SetItem(v ModelBaseItemDto) {
	o.Item = &v
}

// GetNowPlayingQueue returns the NowPlayingQueue field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetNowPlayingQueue() []ModelQueueItem {
	if o == nil || IsNil(o.NowPlayingQueue) {
		var ret []ModelQueueItem
		return ret
	}
	return o.NowPlayingQueue
}

// GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetNowPlayingQueueOk() ([]ModelQueueItem, bool) {
	if o == nil || IsNil(o.NowPlayingQueue) {
		return nil, false
	}
	return o.NowPlayingQueue, true
}

// HasNowPlayingQueue returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasNowPlayingQueue() bool {
	if o != nil && !IsNil(o.NowPlayingQueue) {
		return true
	}

	return false
}

// SetNowPlayingQueue gets a reference to the given []ModelQueueItem and assigns it to the NowPlayingQueue field.
func (o *ModelPlaybackStartInfo) SetNowPlayingQueue(v []ModelQueueItem) {
	o.NowPlayingQueue = v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *ModelPlaybackStartInfo) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *ModelPlaybackStartInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *ModelPlaybackStartInfo) SetSessionId(v string) {
	o.SessionId = &v
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId) {
		var ret string
		return ret
	}
	return *o.MediaSourceId
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetMediaSourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaSourceId) {
		return nil, false
	}
	return o.MediaSourceId, true
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasMediaSourceId() bool {
	if o != nil && !IsNil(o.MediaSourceId) {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given string and assigns it to the MediaSourceId field.
func (o *ModelPlaybackStartInfo) SetMediaSourceId(v string) {
	o.MediaSourceId = &v
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *ModelPlaybackStartInfo) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *ModelPlaybackStartInfo) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *ModelPlaybackStartInfo) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *ModelPlaybackStartInfo) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused) {
		var ret bool
		return ret
	}
	return *o.IsPaused
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetIsPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaused) {
		return nil, false
	}
	return o.IsPaused, true
}

// HasIsPaused returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasIsPaused() bool {
	if o != nil && !IsNil(o.IsPaused) {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given bool and assigns it to the IsPaused field.
func (o *ModelPlaybackStartInfo) SetIsPaused(v bool) {
	o.IsPaused = &v
}

// GetPlaylistIndex returns the PlaylistIndex field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlaylistIndex() int32 {
	if o == nil || IsNil(o.PlaylistIndex) {
		var ret int32
		return ret
	}
	return *o.PlaylistIndex
}

// GetPlaylistIndexOk returns a tuple with the PlaylistIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlaylistIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.PlaylistIndex) {
		return nil, false
	}
	return o.PlaylistIndex, true
}

// HasPlaylistIndex returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaylistIndex() bool {
	if o != nil && !IsNil(o.PlaylistIndex) {
		return true
	}

	return false
}

// SetPlaylistIndex gets a reference to the given int32 and assigns it to the PlaylistIndex field.
func (o *ModelPlaybackStartInfo) SetPlaylistIndex(v int32) {
	o.PlaylistIndex = &v
}

// GetPlaylistLength returns the PlaylistLength field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlaylistLength() int32 {
	if o == nil || IsNil(o.PlaylistLength) {
		var ret int32
		return ret
	}
	return *o.PlaylistLength
}

// GetPlaylistLengthOk returns a tuple with the PlaylistLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlaylistLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.PlaylistLength) {
		return nil, false
	}
	return o.PlaylistLength, true
}

// HasPlaylistLength returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaylistLength() bool {
	if o != nil && !IsNil(o.PlaylistLength) {
		return true
	}

	return false
}

// SetPlaylistLength gets a reference to the given int32 and assigns it to the PlaylistLength field.
func (o *ModelPlaybackStartInfo) SetPlaylistLength(v int32) {
	o.PlaylistLength = &v
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *ModelPlaybackStartInfo) SetIsMuted(v bool) {
	o.IsMuted = &v
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *ModelPlaybackStartInfo) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *ModelPlaybackStartInfo) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetRunTimeTicks returns the RunTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetRunTimeTicks() int64 {
	if o == nil || IsNil(o.RunTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.RunTimeTicks.Get()
}

// GetRunTimeTicksOk returns a tuple with the RunTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetRunTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunTimeTicks.Get(), o.RunTimeTicks.IsSet()
}

// HasRunTimeTicks returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasRunTimeTicks() bool {
	if o != nil && o.RunTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetRunTimeTicks gets a reference to the given NullableInt64 and assigns it to the RunTimeTicks field.
func (o *ModelPlaybackStartInfo) SetRunTimeTicks(v int64) {
	o.RunTimeTicks.Set(&v)
}
// SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil
func (o *ModelPlaybackStartInfo) SetRunTimeTicksNil() {
	o.RunTimeTicks.Set(nil)
}

// UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetRunTimeTicks() {
	o.RunTimeTicks.Unset()
}

// GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetPlaybackStartTimeTicks() int64 {
	if o == nil || IsNil(o.PlaybackStartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PlaybackStartTimeTicks.Get()
}

// GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetPlaybackStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaybackStartTimeTicks.Get(), o.PlaybackStartTimeTicks.IsSet()
}

// HasPlaybackStartTimeTicks returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaybackStartTimeTicks() bool {
	if o != nil && o.PlaybackStartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetPlaybackStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the PlaybackStartTimeTicks field.
func (o *ModelPlaybackStartInfo) SetPlaybackStartTimeTicks(v int64) {
	o.PlaybackStartTimeTicks.Set(&v)
}
// SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil
func (o *ModelPlaybackStartInfo) SetPlaybackStartTimeTicksNil() {
	o.PlaybackStartTimeTicks.Set(nil)
}

// UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetPlaybackStartTimeTicks() {
	o.PlaybackStartTimeTicks.Unset()
}

// GetVolumeLevel returns the VolumeLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetVolumeLevel() int32 {
	if o == nil || IsNil(o.VolumeLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.VolumeLevel.Get()
}

// GetVolumeLevelOk returns a tuple with the VolumeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetVolumeLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeLevel.Get(), o.VolumeLevel.IsSet()
}

// HasVolumeLevel returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasVolumeLevel() bool {
	if o != nil && o.VolumeLevel.IsSet() {
		return true
	}

	return false
}

// SetVolumeLevel gets a reference to the given NullableInt32 and assigns it to the VolumeLevel field.
func (o *ModelPlaybackStartInfo) SetVolumeLevel(v int32) {
	o.VolumeLevel.Set(&v)
}
// SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil
func (o *ModelPlaybackStartInfo) SetVolumeLevelNil() {
	o.VolumeLevel.Set(nil)
}

// UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetVolumeLevel() {
	o.VolumeLevel.Unset()
}

// GetBrightness returns the Brightness field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPlaybackStartInfo) GetBrightness() int32 {
	if o == nil || IsNil(o.Brightness.Get()) {
		var ret int32
		return ret
	}
	return *o.Brightness.Get()
}

// GetBrightnessOk returns a tuple with the Brightness field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPlaybackStartInfo) GetBrightnessOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Brightness.Get(), o.Brightness.IsSet()
}

// HasBrightness returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasBrightness() bool {
	if o != nil && o.Brightness.IsSet() {
		return true
	}

	return false
}

// SetBrightness gets a reference to the given NullableInt32 and assigns it to the Brightness field.
func (o *ModelPlaybackStartInfo) SetBrightness(v int32) {
	o.Brightness.Set(&v)
}
// SetBrightnessNil sets the value for Brightness to be an explicit nil
func (o *ModelPlaybackStartInfo) SetBrightnessNil() {
	o.Brightness.Set(nil)
}

// UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
func (o *ModelPlaybackStartInfo) UnsetBrightness() {
	o.Brightness.Unset()
}

// GetAspectRatio returns the AspectRatio field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetAspectRatio() string {
	if o == nil || IsNil(o.AspectRatio) {
		var ret string
		return ret
	}
	return *o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetAspectRatioOk() (*string, bool) {
	if o == nil || IsNil(o.AspectRatio) {
		return nil, false
	}
	return o.AspectRatio, true
}

// HasAspectRatio returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasAspectRatio() bool {
	if o != nil && !IsNil(o.AspectRatio) {
		return true
	}

	return false
}

// SetAspectRatio gets a reference to the given string and assigns it to the AspectRatio field.
func (o *ModelPlaybackStartInfo) SetAspectRatio(v string) {
	o.AspectRatio = &v
}

// GetEventName returns the EventName field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetEventName() ModelProgressEvent {
	if o == nil || IsNil(o.EventName) {
		var ret ModelProgressEvent
		return ret
	}
	return *o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetEventNameOk() (*ModelProgressEvent, bool) {
	if o == nil || IsNil(o.EventName) {
		return nil, false
	}
	return o.EventName, true
}

// HasEventName returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasEventName() bool {
	if o != nil && !IsNil(o.EventName) {
		return true
	}

	return false
}

// SetEventName gets a reference to the given ModelProgressEvent and assigns it to the EventName field.
func (o *ModelPlaybackStartInfo) SetEventName(v ModelProgressEvent) {
	o.EventName = &v
}

// GetPlayMethod returns the PlayMethod field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlayMethod() ModelPlayMethod {
	if o == nil || IsNil(o.PlayMethod) {
		var ret ModelPlayMethod
		return ret
	}
	return *o.PlayMethod
}

// GetPlayMethodOk returns a tuple with the PlayMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlayMethodOk() (*ModelPlayMethod, bool) {
	if o == nil || IsNil(o.PlayMethod) {
		return nil, false
	}
	return o.PlayMethod, true
}

// HasPlayMethod returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlayMethod() bool {
	if o != nil && !IsNil(o.PlayMethod) {
		return true
	}

	return false
}

// SetPlayMethod gets a reference to the given ModelPlayMethod and assigns it to the PlayMethod field.
func (o *ModelPlaybackStartInfo) SetPlayMethod(v ModelPlayMethod) {
	o.PlayMethod = &v
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId) {
		var ret string
		return ret
	}
	return *o.LiveStreamId
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetLiveStreamIdOk() (*string, bool) {
	if o == nil || IsNil(o.LiveStreamId) {
		return nil, false
	}
	return o.LiveStreamId, true
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasLiveStreamId() bool {
	if o != nil && !IsNil(o.LiveStreamId) {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given string and assigns it to the LiveStreamId field.
func (o *ModelPlaybackStartInfo) SetLiveStreamId(v string) {
	o.LiveStreamId = &v
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId) {
		var ret string
		return ret
	}
	return *o.PlaySessionId
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlaySessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaySessionId) {
		return nil, false
	}
	return o.PlaySessionId, true
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaySessionId() bool {
	if o != nil && !IsNil(o.PlaySessionId) {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given string and assigns it to the PlaySessionId field.
func (o *ModelPlaybackStartInfo) SetPlaySessionId(v string) {
	o.PlaySessionId = &v
}

// GetRepeatMode returns the RepeatMode field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetRepeatMode() ModelRepeatMode {
	if o == nil || IsNil(o.RepeatMode) {
		var ret ModelRepeatMode
		return ret
	}
	return *o.RepeatMode
}

// GetRepeatModeOk returns a tuple with the RepeatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetRepeatModeOk() (*ModelRepeatMode, bool) {
	if o == nil || IsNil(o.RepeatMode) {
		return nil, false
	}
	return o.RepeatMode, true
}

// HasRepeatMode returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasRepeatMode() bool {
	if o != nil && !IsNil(o.RepeatMode) {
		return true
	}

	return false
}

// SetRepeatMode gets a reference to the given ModelRepeatMode and assigns it to the RepeatMode field.
func (o *ModelPlaybackStartInfo) SetRepeatMode(v ModelRepeatMode) {
	o.RepeatMode = &v
}

// GetSubtitleOffset returns the SubtitleOffset field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetSubtitleOffset() int32 {
	if o == nil || IsNil(o.SubtitleOffset) {
		var ret int32
		return ret
	}
	return *o.SubtitleOffset
}

// GetSubtitleOffsetOk returns a tuple with the SubtitleOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetSubtitleOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.SubtitleOffset) {
		return nil, false
	}
	return o.SubtitleOffset, true
}

// HasSubtitleOffset returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasSubtitleOffset() bool {
	if o != nil && !IsNil(o.SubtitleOffset) {
		return true
	}

	return false
}

// SetSubtitleOffset gets a reference to the given int32 and assigns it to the SubtitleOffset field.
func (o *ModelPlaybackStartInfo) SetSubtitleOffset(v int32) {
	o.SubtitleOffset = &v
}

// GetPlaybackRate returns the PlaybackRate field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlaybackRate() float64 {
	if o == nil || IsNil(o.PlaybackRate) {
		var ret float64
		return ret
	}
	return *o.PlaybackRate
}

// GetPlaybackRateOk returns a tuple with the PlaybackRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlaybackRateOk() (*float64, bool) {
	if o == nil || IsNil(o.PlaybackRate) {
		return nil, false
	}
	return o.PlaybackRate, true
}

// HasPlaybackRate returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaybackRate() bool {
	if o != nil && !IsNil(o.PlaybackRate) {
		return true
	}

	return false
}

// SetPlaybackRate gets a reference to the given float64 and assigns it to the PlaybackRate field.
func (o *ModelPlaybackStartInfo) SetPlaybackRate(v float64) {
	o.PlaybackRate = &v
}

// GetPlaylistItemIds returns the PlaylistItemIds field value if set, zero value otherwise.
func (o *ModelPlaybackStartInfo) GetPlaylistItemIds() []string {
	if o == nil || IsNil(o.PlaylistItemIds) {
		var ret []string
		return ret
	}
	return o.PlaylistItemIds
}

// GetPlaylistItemIdsOk returns a tuple with the PlaylistItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPlaybackStartInfo) GetPlaylistItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PlaylistItemIds) {
		return nil, false
	}
	return o.PlaylistItemIds, true
}

// HasPlaylistItemIds returns a boolean if a field has been set.
func (o *ModelPlaybackStartInfo) HasPlaylistItemIds() bool {
	if o != nil && !IsNil(o.PlaylistItemIds) {
		return true
	}

	return false
}

// SetPlaylistItemIds gets a reference to the given []string and assigns it to the PlaylistItemIds field.
func (o *ModelPlaybackStartInfo) SetPlaylistItemIds(v []string) {
	o.PlaylistItemIds = v
}

func (o ModelPlaybackStartInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPlaybackStartInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanSeek) {
		toSerialize["CanSeek"] = o.CanSeek
	}
	if !IsNil(o.Item) {
		toSerialize["Item"] = o.Item
	}
	if !IsNil(o.NowPlayingQueue) {
		toSerialize["NowPlayingQueue"] = o.NowPlayingQueue
	}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.SessionId) {
		toSerialize["SessionId"] = o.SessionId
	}
	if !IsNil(o.MediaSourceId) {
		toSerialize["MediaSourceId"] = o.MediaSourceId
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if !IsNil(o.IsPaused) {
		toSerialize["IsPaused"] = o.IsPaused
	}
	if !IsNil(o.PlaylistIndex) {
		toSerialize["PlaylistIndex"] = o.PlaylistIndex
	}
	if !IsNil(o.PlaylistLength) {
		toSerialize["PlaylistLength"] = o.PlaylistLength
	}
	if !IsNil(o.IsMuted) {
		toSerialize["IsMuted"] = o.IsMuted
	}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if o.RunTimeTicks.IsSet() {
		toSerialize["RunTimeTicks"] = o.RunTimeTicks.Get()
	}
	if o.PlaybackStartTimeTicks.IsSet() {
		toSerialize["PlaybackStartTimeTicks"] = o.PlaybackStartTimeTicks.Get()
	}
	if o.VolumeLevel.IsSet() {
		toSerialize["VolumeLevel"] = o.VolumeLevel.Get()
	}
	if o.Brightness.IsSet() {
		toSerialize["Brightness"] = o.Brightness.Get()
	}
	if !IsNil(o.AspectRatio) {
		toSerialize["AspectRatio"] = o.AspectRatio
	}
	if !IsNil(o.EventName) {
		toSerialize["EventName"] = o.EventName
	}
	if !IsNil(o.PlayMethod) {
		toSerialize["PlayMethod"] = o.PlayMethod
	}
	if !IsNil(o.LiveStreamId) {
		toSerialize["LiveStreamId"] = o.LiveStreamId
	}
	if !IsNil(o.PlaySessionId) {
		toSerialize["PlaySessionId"] = o.PlaySessionId
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
	if !IsNil(o.PlaylistItemIds) {
		toSerialize["PlaylistItemIds"] = o.PlaylistItemIds
	}
	return toSerialize, nil
}

type NullableModelPlaybackStartInfo struct {
	value *ModelPlaybackStartInfo
	isSet bool
}

func (v NullableModelPlaybackStartInfo) Get() *ModelPlaybackStartInfo {
	return v.value
}

func (v *NullableModelPlaybackStartInfo) Set(val *ModelPlaybackStartInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPlaybackStartInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPlaybackStartInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPlaybackStartInfo(val *ModelPlaybackStartInfo) *NullableModelPlaybackStartInfo {
	return &NullableModelPlaybackStartInfo{value: val, isSet: true}
}

func (v NullableModelPlaybackStartInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPlaybackStartInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


