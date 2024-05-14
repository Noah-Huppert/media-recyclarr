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

// checks if the ModelLiveTvTimerInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLiveTvTimerInfoDto{}

// ModelLiveTvTimerInfoDto struct for ModelLiveTvTimerInfoDto
type ModelLiveTvTimerInfoDto struct {
	Status *ModelLiveTvRecordingStatus `json:"Status,omitempty"`
	SeriesTimerId *string `json:"SeriesTimerId,omitempty"`
	RunTimeTicks NullableInt64 `json:"RunTimeTicks,omitempty"`
	ProgramInfo *ModelBaseItemDto `json:"ProgramInfo,omitempty"`
	TimerType *ModelLiveTvTimerType `json:"TimerType,omitempty"`
	Id *string `json:"Id,omitempty"`
	Type *string `json:"Type,omitempty"`
	ServerId *string `json:"ServerId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty"`
	ChannelName *string `json:"ChannelName,omitempty"`
	ChannelNumber *string `json:"ChannelNumber,omitempty"`
	ChannelPrimaryImageTag *string `json:"ChannelPrimaryImageTag,omitempty"`
	ProgramId *string `json:"ProgramId,omitempty"`
	Name *string `json:"Name,omitempty"`
	Overview *string `json:"Overview,omitempty"`
	ParentFolderId *string `json:"ParentFolderId,omitempty"`
	StartDate *time.Time `json:"StartDate,omitempty"`
	EndDate *time.Time `json:"EndDate,omitempty"`
	Priority *int32 `json:"Priority,omitempty"`
	PrePaddingSeconds *int32 `json:"PrePaddingSeconds,omitempty"`
	PostPaddingSeconds *int32 `json:"PostPaddingSeconds,omitempty"`
	IsPrePaddingRequired *bool `json:"IsPrePaddingRequired,omitempty"`
	ParentBackdropItemId *string `json:"ParentBackdropItemId,omitempty"`
	ParentBackdropImageTags []string `json:"ParentBackdropImageTags,omitempty"`
	IsPostPaddingRequired *bool `json:"IsPostPaddingRequired,omitempty"`
	KeepUntil *ModelLiveTvKeepUntil `json:"KeepUntil,omitempty"`
}

// NewModelLiveTvTimerInfoDto instantiates a new ModelLiveTvTimerInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLiveTvTimerInfoDto() *ModelLiveTvTimerInfoDto {
	this := ModelLiveTvTimerInfoDto{}
	return &this
}

// NewModelLiveTvTimerInfoDtoWithDefaults instantiates a new ModelLiveTvTimerInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLiveTvTimerInfoDtoWithDefaults() *ModelLiveTvTimerInfoDto {
	this := ModelLiveTvTimerInfoDto{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetStatus() ModelLiveTvRecordingStatus {
	if o == nil || IsNil(o.Status) {
		var ret ModelLiveTvRecordingStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetStatusOk() (*ModelLiveTvRecordingStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ModelLiveTvRecordingStatus and assigns it to the Status field.
func (o *ModelLiveTvTimerInfoDto) SetStatus(v ModelLiveTvRecordingStatus) {
	o.Status = &v
}

// GetSeriesTimerId returns the SeriesTimerId field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetSeriesTimerId() string {
	if o == nil || IsNil(o.SeriesTimerId) {
		var ret string
		return ret
	}
	return *o.SeriesTimerId
}

// GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetSeriesTimerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SeriesTimerId) {
		return nil, false
	}
	return o.SeriesTimerId, true
}

// HasSeriesTimerId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasSeriesTimerId() bool {
	if o != nil && !IsNil(o.SeriesTimerId) {
		return true
	}

	return false
}

// SetSeriesTimerId gets a reference to the given string and assigns it to the SeriesTimerId field.
func (o *ModelLiveTvTimerInfoDto) SetSeriesTimerId(v string) {
	o.SeriesTimerId = &v
}

// GetRunTimeTicks returns the RunTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelLiveTvTimerInfoDto) GetRunTimeTicks() int64 {
	if o == nil || IsNil(o.RunTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.RunTimeTicks.Get()
}

// GetRunTimeTicksOk returns a tuple with the RunTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelLiveTvTimerInfoDto) GetRunTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RunTimeTicks.Get(), o.RunTimeTicks.IsSet()
}

// HasRunTimeTicks returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasRunTimeTicks() bool {
	if o != nil && o.RunTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetRunTimeTicks gets a reference to the given NullableInt64 and assigns it to the RunTimeTicks field.
func (o *ModelLiveTvTimerInfoDto) SetRunTimeTicks(v int64) {
	o.RunTimeTicks.Set(&v)
}
// SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil
func (o *ModelLiveTvTimerInfoDto) SetRunTimeTicksNil() {
	o.RunTimeTicks.Set(nil)
}

// UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
func (o *ModelLiveTvTimerInfoDto) UnsetRunTimeTicks() {
	o.RunTimeTicks.Unset()
}

// GetProgramInfo returns the ProgramInfo field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetProgramInfo() ModelBaseItemDto {
	if o == nil || IsNil(o.ProgramInfo) {
		var ret ModelBaseItemDto
		return ret
	}
	return *o.ProgramInfo
}

// GetProgramInfoOk returns a tuple with the ProgramInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetProgramInfoOk() (*ModelBaseItemDto, bool) {
	if o == nil || IsNil(o.ProgramInfo) {
		return nil, false
	}
	return o.ProgramInfo, true
}

// HasProgramInfo returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasProgramInfo() bool {
	if o != nil && !IsNil(o.ProgramInfo) {
		return true
	}

	return false
}

// SetProgramInfo gets a reference to the given ModelBaseItemDto and assigns it to the ProgramInfo field.
func (o *ModelLiveTvTimerInfoDto) SetProgramInfo(v ModelBaseItemDto) {
	o.ProgramInfo = &v
}

// GetTimerType returns the TimerType field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetTimerType() ModelLiveTvTimerType {
	if o == nil || IsNil(o.TimerType) {
		var ret ModelLiveTvTimerType
		return ret
	}
	return *o.TimerType
}

// GetTimerTypeOk returns a tuple with the TimerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetTimerTypeOk() (*ModelLiveTvTimerType, bool) {
	if o == nil || IsNil(o.TimerType) {
		return nil, false
	}
	return o.TimerType, true
}

// HasTimerType returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasTimerType() bool {
	if o != nil && !IsNil(o.TimerType) {
		return true
	}

	return false
}

// SetTimerType gets a reference to the given ModelLiveTvTimerType and assigns it to the TimerType field.
func (o *ModelLiveTvTimerInfoDto) SetTimerType(v ModelLiveTvTimerType) {
	o.TimerType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelLiveTvTimerInfoDto) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelLiveTvTimerInfoDto) SetType(v string) {
	o.Type = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *ModelLiveTvTimerInfoDto) SetServerId(v string) {
	o.ServerId = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *ModelLiveTvTimerInfoDto) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetChannelName() string {
	if o == nil || IsNil(o.ChannelName) {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetChannelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelName) {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasChannelName() bool {
	if o != nil && !IsNil(o.ChannelName) {
		return true
	}

	return false
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *ModelLiveTvTimerInfoDto) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetChannelNumber returns the ChannelNumber field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetChannelNumber() string {
	if o == nil || IsNil(o.ChannelNumber) {
		var ret string
		return ret
	}
	return *o.ChannelNumber
}

// GetChannelNumberOk returns a tuple with the ChannelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetChannelNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelNumber) {
		return nil, false
	}
	return o.ChannelNumber, true
}

// HasChannelNumber returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasChannelNumber() bool {
	if o != nil && !IsNil(o.ChannelNumber) {
		return true
	}

	return false
}

// SetChannelNumber gets a reference to the given string and assigns it to the ChannelNumber field.
func (o *ModelLiveTvTimerInfoDto) SetChannelNumber(v string) {
	o.ChannelNumber = &v
}

// GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetChannelPrimaryImageTag() string {
	if o == nil || IsNil(o.ChannelPrimaryImageTag) {
		var ret string
		return ret
	}
	return *o.ChannelPrimaryImageTag
}

// GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelPrimaryImageTag) {
		return nil, false
	}
	return o.ChannelPrimaryImageTag, true
}

// HasChannelPrimaryImageTag returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasChannelPrimaryImageTag() bool {
	if o != nil && !IsNil(o.ChannelPrimaryImageTag) {
		return true
	}

	return false
}

// SetChannelPrimaryImageTag gets a reference to the given string and assigns it to the ChannelPrimaryImageTag field.
func (o *ModelLiveTvTimerInfoDto) SetChannelPrimaryImageTag(v string) {
	o.ChannelPrimaryImageTag = &v
}

// GetProgramId returns the ProgramId field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetProgramId() string {
	if o == nil || IsNil(o.ProgramId) {
		var ret string
		return ret
	}
	return *o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetProgramIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramId) {
		return nil, false
	}
	return o.ProgramId, true
}

// HasProgramId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasProgramId() bool {
	if o != nil && !IsNil(o.ProgramId) {
		return true
	}

	return false
}

// SetProgramId gets a reference to the given string and assigns it to the ProgramId field.
func (o *ModelLiveTvTimerInfoDto) SetProgramId(v string) {
	o.ProgramId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelLiveTvTimerInfoDto) SetName(v string) {
	o.Name = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *ModelLiveTvTimerInfoDto) SetOverview(v string) {
	o.Overview = &v
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetParentFolderId() string {
	if o == nil || IsNil(o.ParentFolderId) {
		var ret string
		return ret
	}
	return *o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetParentFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentFolderId) {
		return nil, false
	}
	return o.ParentFolderId, true
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasParentFolderId() bool {
	if o != nil && !IsNil(o.ParentFolderId) {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given string and assigns it to the ParentFolderId field.
func (o *ModelLiveTvTimerInfoDto) SetParentFolderId(v string) {
	o.ParentFolderId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ModelLiveTvTimerInfoDto) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ModelLiveTvTimerInfoDto) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ModelLiveTvTimerInfoDto) SetPriority(v int32) {
	o.Priority = &v
}

// GetPrePaddingSeconds returns the PrePaddingSeconds field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetPrePaddingSeconds() int32 {
	if o == nil || IsNil(o.PrePaddingSeconds) {
		var ret int32
		return ret
	}
	return *o.PrePaddingSeconds
}

// GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PrePaddingSeconds) {
		return nil, false
	}
	return o.PrePaddingSeconds, true
}

// HasPrePaddingSeconds returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasPrePaddingSeconds() bool {
	if o != nil && !IsNil(o.PrePaddingSeconds) {
		return true
	}

	return false
}

// SetPrePaddingSeconds gets a reference to the given int32 and assigns it to the PrePaddingSeconds field.
func (o *ModelLiveTvTimerInfoDto) SetPrePaddingSeconds(v int32) {
	o.PrePaddingSeconds = &v
}

// GetPostPaddingSeconds returns the PostPaddingSeconds field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetPostPaddingSeconds() int32 {
	if o == nil || IsNil(o.PostPaddingSeconds) {
		var ret int32
		return ret
	}
	return *o.PostPaddingSeconds
}

// GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PostPaddingSeconds) {
		return nil, false
	}
	return o.PostPaddingSeconds, true
}

// HasPostPaddingSeconds returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasPostPaddingSeconds() bool {
	if o != nil && !IsNil(o.PostPaddingSeconds) {
		return true
	}

	return false
}

// SetPostPaddingSeconds gets a reference to the given int32 and assigns it to the PostPaddingSeconds field.
func (o *ModelLiveTvTimerInfoDto) SetPostPaddingSeconds(v int32) {
	o.PostPaddingSeconds = &v
}

// GetIsPrePaddingRequired returns the IsPrePaddingRequired field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetIsPrePaddingRequired() bool {
	if o == nil || IsNil(o.IsPrePaddingRequired) {
		var ret bool
		return ret
	}
	return *o.IsPrePaddingRequired
}

// GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrePaddingRequired) {
		return nil, false
	}
	return o.IsPrePaddingRequired, true
}

// HasIsPrePaddingRequired returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasIsPrePaddingRequired() bool {
	if o != nil && !IsNil(o.IsPrePaddingRequired) {
		return true
	}

	return false
}

// SetIsPrePaddingRequired gets a reference to the given bool and assigns it to the IsPrePaddingRequired field.
func (o *ModelLiveTvTimerInfoDto) SetIsPrePaddingRequired(v bool) {
	o.IsPrePaddingRequired = &v
}

// GetParentBackdropItemId returns the ParentBackdropItemId field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetParentBackdropItemId() string {
	if o == nil || IsNil(o.ParentBackdropItemId) {
		var ret string
		return ret
	}
	return *o.ParentBackdropItemId
}

// GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetParentBackdropItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentBackdropItemId) {
		return nil, false
	}
	return o.ParentBackdropItemId, true
}

// HasParentBackdropItemId returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasParentBackdropItemId() bool {
	if o != nil && !IsNil(o.ParentBackdropItemId) {
		return true
	}

	return false
}

// SetParentBackdropItemId gets a reference to the given string and assigns it to the ParentBackdropItemId field.
func (o *ModelLiveTvTimerInfoDto) SetParentBackdropItemId(v string) {
	o.ParentBackdropItemId = &v
}

// GetParentBackdropImageTags returns the ParentBackdropImageTags field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetParentBackdropImageTags() []string {
	if o == nil || IsNil(o.ParentBackdropImageTags) {
		var ret []string
		return ret
	}
	return o.ParentBackdropImageTags
}

// GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetParentBackdropImageTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentBackdropImageTags) {
		return nil, false
	}
	return o.ParentBackdropImageTags, true
}

// HasParentBackdropImageTags returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasParentBackdropImageTags() bool {
	if o != nil && !IsNil(o.ParentBackdropImageTags) {
		return true
	}

	return false
}

// SetParentBackdropImageTags gets a reference to the given []string and assigns it to the ParentBackdropImageTags field.
func (o *ModelLiveTvTimerInfoDto) SetParentBackdropImageTags(v []string) {
	o.ParentBackdropImageTags = v
}

// GetIsPostPaddingRequired returns the IsPostPaddingRequired field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetIsPostPaddingRequired() bool {
	if o == nil || IsNil(o.IsPostPaddingRequired) {
		var ret bool
		return ret
	}
	return *o.IsPostPaddingRequired
}

// GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPostPaddingRequired) {
		return nil, false
	}
	return o.IsPostPaddingRequired, true
}

// HasIsPostPaddingRequired returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasIsPostPaddingRequired() bool {
	if o != nil && !IsNil(o.IsPostPaddingRequired) {
		return true
	}

	return false
}

// SetIsPostPaddingRequired gets a reference to the given bool and assigns it to the IsPostPaddingRequired field.
func (o *ModelLiveTvTimerInfoDto) SetIsPostPaddingRequired(v bool) {
	o.IsPostPaddingRequired = &v
}

// GetKeepUntil returns the KeepUntil field value if set, zero value otherwise.
func (o *ModelLiveTvTimerInfoDto) GetKeepUntil() ModelLiveTvKeepUntil {
	if o == nil || IsNil(o.KeepUntil) {
		var ret ModelLiveTvKeepUntil
		return ret
	}
	return *o.KeepUntil
}

// GetKeepUntilOk returns a tuple with the KeepUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLiveTvTimerInfoDto) GetKeepUntilOk() (*ModelLiveTvKeepUntil, bool) {
	if o == nil || IsNil(o.KeepUntil) {
		return nil, false
	}
	return o.KeepUntil, true
}

// HasKeepUntil returns a boolean if a field has been set.
func (o *ModelLiveTvTimerInfoDto) HasKeepUntil() bool {
	if o != nil && !IsNil(o.KeepUntil) {
		return true
	}

	return false
}

// SetKeepUntil gets a reference to the given ModelLiveTvKeepUntil and assigns it to the KeepUntil field.
func (o *ModelLiveTvTimerInfoDto) SetKeepUntil(v ModelLiveTvKeepUntil) {
	o.KeepUntil = &v
}

func (o ModelLiveTvTimerInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLiveTvTimerInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.SeriesTimerId) {
		toSerialize["SeriesTimerId"] = o.SeriesTimerId
	}
	if o.RunTimeTicks.IsSet() {
		toSerialize["RunTimeTicks"] = o.RunTimeTicks.Get()
	}
	if !IsNil(o.ProgramInfo) {
		toSerialize["ProgramInfo"] = o.ProgramInfo
	}
	if !IsNil(o.TimerType) {
		toSerialize["TimerType"] = o.TimerType
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.ServerId) {
		toSerialize["ServerId"] = o.ServerId
	}
	if !IsNil(o.ChannelId) {
		toSerialize["ChannelId"] = o.ChannelId
	}
	if !IsNil(o.ChannelName) {
		toSerialize["ChannelName"] = o.ChannelName
	}
	if !IsNil(o.ChannelNumber) {
		toSerialize["ChannelNumber"] = o.ChannelNumber
	}
	if !IsNil(o.ChannelPrimaryImageTag) {
		toSerialize["ChannelPrimaryImageTag"] = o.ChannelPrimaryImageTag
	}
	if !IsNil(o.ProgramId) {
		toSerialize["ProgramId"] = o.ProgramId
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Overview) {
		toSerialize["Overview"] = o.Overview
	}
	if !IsNil(o.ParentFolderId) {
		toSerialize["ParentFolderId"] = o.ParentFolderId
	}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["EndDate"] = o.EndDate
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.PrePaddingSeconds) {
		toSerialize["PrePaddingSeconds"] = o.PrePaddingSeconds
	}
	if !IsNil(o.PostPaddingSeconds) {
		toSerialize["PostPaddingSeconds"] = o.PostPaddingSeconds
	}
	if !IsNil(o.IsPrePaddingRequired) {
		toSerialize["IsPrePaddingRequired"] = o.IsPrePaddingRequired
	}
	if !IsNil(o.ParentBackdropItemId) {
		toSerialize["ParentBackdropItemId"] = o.ParentBackdropItemId
	}
	if !IsNil(o.ParentBackdropImageTags) {
		toSerialize["ParentBackdropImageTags"] = o.ParentBackdropImageTags
	}
	if !IsNil(o.IsPostPaddingRequired) {
		toSerialize["IsPostPaddingRequired"] = o.IsPostPaddingRequired
	}
	if !IsNil(o.KeepUntil) {
		toSerialize["KeepUntil"] = o.KeepUntil
	}
	return toSerialize, nil
}

type NullableModelLiveTvTimerInfoDto struct {
	value *ModelLiveTvTimerInfoDto
	isSet bool
}

func (v NullableModelLiveTvTimerInfoDto) Get() *ModelLiveTvTimerInfoDto {
	return v.value
}

func (v *NullableModelLiveTvTimerInfoDto) Set(val *ModelLiveTvTimerInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLiveTvTimerInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLiveTvTimerInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLiveTvTimerInfoDto(val *ModelLiveTvTimerInfoDto) *NullableModelLiveTvTimerInfoDto {
	return &NullableModelLiveTvTimerInfoDto{value: val, isSet: true}
}

func (v NullableModelLiveTvTimerInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLiveTvTimerInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


