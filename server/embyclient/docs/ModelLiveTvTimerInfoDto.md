# ModelLiveTvTimerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ModelModelLiveTvRecordingStatus**](ModelLiveTvRecordingStatus.md) |  | [optional] 
**SeriesTimerId** | Pointer to **string** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**ProgramInfo** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**TimerType** | Pointer to [**ModelModelLiveTvTimerType**](ModelLiveTvTimerType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**ChannelName** | Pointer to **string** |  | [optional] 
**ChannelNumber** | Pointer to **string** |  | [optional] 
**ChannelPrimaryImageTag** | Pointer to **string** |  | [optional] 
**ProgramId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**ParentFolderId** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**PrePaddingSeconds** | Pointer to **int32** |  | [optional] 
**PostPaddingSeconds** | Pointer to **int32** |  | [optional] 
**IsPrePaddingRequired** | Pointer to **bool** |  | [optional] 
**ParentBackdropItemId** | Pointer to **string** |  | [optional] 
**ParentBackdropImageTags** | Pointer to **[]string** |  | [optional] 
**IsPostPaddingRequired** | Pointer to **bool** |  | [optional] 
**KeepUntil** | Pointer to [**ModelModelLiveTvKeepUntil**](ModelLiveTvKeepUntil.md) |  | [optional] 

## Methods

### NewModelLiveTvTimerInfoDto

`func NewModelLiveTvTimerInfoDto() *ModelLiveTvTimerInfoDto`

NewModelLiveTvTimerInfoDto instantiates a new ModelLiveTvTimerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTvTimerInfoDtoWithDefaults

`func NewModelLiveTvTimerInfoDtoWithDefaults() *ModelLiveTvTimerInfoDto`

NewModelLiveTvTimerInfoDtoWithDefaults instantiates a new ModelLiveTvTimerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ModelLiveTvTimerInfoDto) GetStatus() ModelModelLiveTvRecordingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelLiveTvTimerInfoDto) GetStatusOk() (*ModelModelLiveTvRecordingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelLiveTvTimerInfoDto) SetStatus(v ModelModelLiveTvRecordingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelLiveTvTimerInfoDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeriesTimerId

`func (o *ModelLiveTvTimerInfoDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *ModelLiveTvTimerInfoDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *ModelLiveTvTimerInfoDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *ModelLiveTvTimerInfoDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### GetRunTimeTicks

`func (o *ModelLiveTvTimerInfoDto) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *ModelLiveTvTimerInfoDto) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *ModelLiveTvTimerInfoDto) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *ModelLiveTvTimerInfoDto) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *ModelLiveTvTimerInfoDto) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *ModelLiveTvTimerInfoDto) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetProgramInfo

`func (o *ModelLiveTvTimerInfoDto) GetProgramInfo() ModelModelBaseItemDto`

GetProgramInfo returns the ProgramInfo field if non-nil, zero value otherwise.

### GetProgramInfoOk

`func (o *ModelLiveTvTimerInfoDto) GetProgramInfoOk() (*ModelModelBaseItemDto, bool)`

GetProgramInfoOk returns a tuple with the ProgramInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramInfo

`func (o *ModelLiveTvTimerInfoDto) SetProgramInfo(v ModelModelBaseItemDto)`

SetProgramInfo sets ProgramInfo field to given value.

### HasProgramInfo

`func (o *ModelLiveTvTimerInfoDto) HasProgramInfo() bool`

HasProgramInfo returns a boolean if a field has been set.

### GetTimerType

`func (o *ModelLiveTvTimerInfoDto) GetTimerType() ModelModelLiveTvTimerType`

GetTimerType returns the TimerType field if non-nil, zero value otherwise.

### GetTimerTypeOk

`func (o *ModelLiveTvTimerInfoDto) GetTimerTypeOk() (*ModelModelLiveTvTimerType, bool)`

GetTimerTypeOk returns a tuple with the TimerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerType

`func (o *ModelLiveTvTimerInfoDto) SetTimerType(v ModelModelLiveTvTimerType)`

SetTimerType sets TimerType field to given value.

### HasTimerType

`func (o *ModelLiveTvTimerInfoDto) HasTimerType() bool`

HasTimerType returns a boolean if a field has been set.

### GetId

`func (o *ModelLiveTvTimerInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelLiveTvTimerInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelLiveTvTimerInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelLiveTvTimerInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ModelLiveTvTimerInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelLiveTvTimerInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelLiveTvTimerInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelLiveTvTimerInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServerId

`func (o *ModelLiveTvTimerInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelLiveTvTimerInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelLiveTvTimerInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelLiveTvTimerInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetChannelId

`func (o *ModelLiveTvTimerInfoDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ModelLiveTvTimerInfoDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ModelLiveTvTimerInfoDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ModelLiveTvTimerInfoDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelName

`func (o *ModelLiveTvTimerInfoDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ModelLiveTvTimerInfoDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ModelLiveTvTimerInfoDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ModelLiveTvTimerInfoDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetChannelNumber

`func (o *ModelLiveTvTimerInfoDto) GetChannelNumber() string`

GetChannelNumber returns the ChannelNumber field if non-nil, zero value otherwise.

### GetChannelNumberOk

`func (o *ModelLiveTvTimerInfoDto) GetChannelNumberOk() (*string, bool)`

GetChannelNumberOk returns a tuple with the ChannelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNumber

`func (o *ModelLiveTvTimerInfoDto) SetChannelNumber(v string)`

SetChannelNumber sets ChannelNumber field to given value.

### HasChannelNumber

`func (o *ModelLiveTvTimerInfoDto) HasChannelNumber() bool`

HasChannelNumber returns a boolean if a field has been set.

### GetChannelPrimaryImageTag

`func (o *ModelLiveTvTimerInfoDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *ModelLiveTvTimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *ModelLiveTvTimerInfoDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *ModelLiveTvTimerInfoDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### GetProgramId

`func (o *ModelLiveTvTimerInfoDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ModelLiveTvTimerInfoDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ModelLiveTvTimerInfoDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ModelLiveTvTimerInfoDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetName

`func (o *ModelLiveTvTimerInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelLiveTvTimerInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelLiveTvTimerInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelLiveTvTimerInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *ModelLiveTvTimerInfoDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelLiveTvTimerInfoDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelLiveTvTimerInfoDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelLiveTvTimerInfoDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetParentFolderId

`func (o *ModelLiveTvTimerInfoDto) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *ModelLiveTvTimerInfoDto) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *ModelLiveTvTimerInfoDto) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *ModelLiveTvTimerInfoDto) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetStartDate

`func (o *ModelLiveTvTimerInfoDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelLiveTvTimerInfoDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelLiveTvTimerInfoDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelLiveTvTimerInfoDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelLiveTvTimerInfoDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelLiveTvTimerInfoDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelLiveTvTimerInfoDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelLiveTvTimerInfoDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriority

`func (o *ModelLiveTvTimerInfoDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelLiveTvTimerInfoDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelLiveTvTimerInfoDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelLiveTvTimerInfoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *ModelLiveTvTimerInfoDto) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *ModelLiveTvTimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *ModelLiveTvTimerInfoDto) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *ModelLiveTvTimerInfoDto) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *ModelLiveTvTimerInfoDto) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *ModelLiveTvTimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *ModelLiveTvTimerInfoDto) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *ModelLiveTvTimerInfoDto) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *ModelLiveTvTimerInfoDto) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *ModelLiveTvTimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *ModelLiveTvTimerInfoDto) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *ModelLiveTvTimerInfoDto) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *ModelLiveTvTimerInfoDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *ModelLiveTvTimerInfoDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *ModelLiveTvTimerInfoDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *ModelLiveTvTimerInfoDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### GetParentBackdropImageTags

`func (o *ModelLiveTvTimerInfoDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *ModelLiveTvTimerInfoDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *ModelLiveTvTimerInfoDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *ModelLiveTvTimerInfoDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### GetIsPostPaddingRequired

`func (o *ModelLiveTvTimerInfoDto) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *ModelLiveTvTimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *ModelLiveTvTimerInfoDto) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *ModelLiveTvTimerInfoDto) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetKeepUntil

`func (o *ModelLiveTvTimerInfoDto) GetKeepUntil() ModelModelLiveTvKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *ModelLiveTvTimerInfoDto) GetKeepUntilOk() (*ModelModelLiveTvKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *ModelLiveTvTimerInfoDto) SetKeepUntil(v ModelModelLiveTvKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *ModelLiveTvTimerInfoDto) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


