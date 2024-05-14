# ModelLiveTvSeriesTimerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordAnyTime** | Pointer to **bool** |  | [optional] 
**SkipEpisodesInLibrary** | Pointer to **bool** |  | [optional] 
**RecordAnyChannel** | Pointer to **bool** |  | [optional] 
**KeepUpTo** | Pointer to **int32** |  | [optional] 
**MaxRecordingSeconds** | Pointer to **int32** |  | [optional] 
**RecordNewOnly** | Pointer to **bool** |  | [optional] 
**ChannelIds** | Pointer to **[]string** |  | [optional] 
**Days** | Pointer to [**[]ModelModelDayOfWeek**](ModelModelDayOfWeek.md) |  | [optional] 
**ImageTags** | Pointer to **map[string]string** |  | [optional] 
**ParentThumbItemId** | Pointer to **string** |  | [optional] 
**ParentThumbImageTag** | Pointer to **string** |  | [optional] 
**ParentPrimaryImageItemId** | Pointer to **string** |  | [optional] 
**ParentPrimaryImageTag** | Pointer to **string** |  | [optional] 
**SeriesId** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to [**[]ModelModelLiveTvKeywordInfo**](ModelModelLiveTvKeywordInfo.md) |  | [optional] 
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

### NewModelLiveTvSeriesTimerInfoDto

`func NewModelLiveTvSeriesTimerInfoDto() *ModelLiveTvSeriesTimerInfoDto`

NewModelLiveTvSeriesTimerInfoDto instantiates a new ModelLiveTvSeriesTimerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTvSeriesTimerInfoDtoWithDefaults

`func NewModelLiveTvSeriesTimerInfoDtoWithDefaults() *ModelLiveTvSeriesTimerInfoDto`

NewModelLiveTvSeriesTimerInfoDtoWithDefaults instantiates a new ModelLiveTvSeriesTimerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordAnyTime

`func (o *ModelLiveTvSeriesTimerInfoDto) GetRecordAnyTime() bool`

GetRecordAnyTime returns the RecordAnyTime field if non-nil, zero value otherwise.

### GetRecordAnyTimeOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetRecordAnyTimeOk() (*bool, bool)`

GetRecordAnyTimeOk returns a tuple with the RecordAnyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyTime

`func (o *ModelLiveTvSeriesTimerInfoDto) SetRecordAnyTime(v bool)`

SetRecordAnyTime sets RecordAnyTime field to given value.

### HasRecordAnyTime

`func (o *ModelLiveTvSeriesTimerInfoDto) HasRecordAnyTime() bool`

HasRecordAnyTime returns a boolean if a field has been set.

### GetSkipEpisodesInLibrary

`func (o *ModelLiveTvSeriesTimerInfoDto) GetSkipEpisodesInLibrary() bool`

GetSkipEpisodesInLibrary returns the SkipEpisodesInLibrary field if non-nil, zero value otherwise.

### GetSkipEpisodesInLibraryOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetSkipEpisodesInLibraryOk() (*bool, bool)`

GetSkipEpisodesInLibraryOk returns a tuple with the SkipEpisodesInLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEpisodesInLibrary

`func (o *ModelLiveTvSeriesTimerInfoDto) SetSkipEpisodesInLibrary(v bool)`

SetSkipEpisodesInLibrary sets SkipEpisodesInLibrary field to given value.

### HasSkipEpisodesInLibrary

`func (o *ModelLiveTvSeriesTimerInfoDto) HasSkipEpisodesInLibrary() bool`

HasSkipEpisodesInLibrary returns a boolean if a field has been set.

### GetRecordAnyChannel

`func (o *ModelLiveTvSeriesTimerInfoDto) GetRecordAnyChannel() bool`

GetRecordAnyChannel returns the RecordAnyChannel field if non-nil, zero value otherwise.

### GetRecordAnyChannelOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetRecordAnyChannelOk() (*bool, bool)`

GetRecordAnyChannelOk returns a tuple with the RecordAnyChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyChannel

`func (o *ModelLiveTvSeriesTimerInfoDto) SetRecordAnyChannel(v bool)`

SetRecordAnyChannel sets RecordAnyChannel field to given value.

### HasRecordAnyChannel

`func (o *ModelLiveTvSeriesTimerInfoDto) HasRecordAnyChannel() bool`

HasRecordAnyChannel returns a boolean if a field has been set.

### GetKeepUpTo

`func (o *ModelLiveTvSeriesTimerInfoDto) GetKeepUpTo() int32`

GetKeepUpTo returns the KeepUpTo field if non-nil, zero value otherwise.

### GetKeepUpToOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetKeepUpToOk() (*int32, bool)`

GetKeepUpToOk returns a tuple with the KeepUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUpTo

`func (o *ModelLiveTvSeriesTimerInfoDto) SetKeepUpTo(v int32)`

SetKeepUpTo sets KeepUpTo field to given value.

### HasKeepUpTo

`func (o *ModelLiveTvSeriesTimerInfoDto) HasKeepUpTo() bool`

HasKeepUpTo returns a boolean if a field has been set.

### GetMaxRecordingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) GetMaxRecordingSeconds() int32`

GetMaxRecordingSeconds returns the MaxRecordingSeconds field if non-nil, zero value otherwise.

### GetMaxRecordingSecondsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetMaxRecordingSecondsOk() (*int32, bool)`

GetMaxRecordingSecondsOk returns a tuple with the MaxRecordingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) SetMaxRecordingSeconds(v int32)`

SetMaxRecordingSeconds sets MaxRecordingSeconds field to given value.

### HasMaxRecordingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) HasMaxRecordingSeconds() bool`

HasMaxRecordingSeconds returns a boolean if a field has been set.

### GetRecordNewOnly

`func (o *ModelLiveTvSeriesTimerInfoDto) GetRecordNewOnly() bool`

GetRecordNewOnly returns the RecordNewOnly field if non-nil, zero value otherwise.

### GetRecordNewOnlyOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetRecordNewOnlyOk() (*bool, bool)`

GetRecordNewOnlyOk returns a tuple with the RecordNewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNewOnly

`func (o *ModelLiveTvSeriesTimerInfoDto) SetRecordNewOnly(v bool)`

SetRecordNewOnly sets RecordNewOnly field to given value.

### HasRecordNewOnly

`func (o *ModelLiveTvSeriesTimerInfoDto) HasRecordNewOnly() bool`

HasRecordNewOnly returns a boolean if a field has been set.

### GetChannelIds

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *ModelLiveTvSeriesTimerInfoDto) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *ModelLiveTvSeriesTimerInfoDto) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### GetDays

`func (o *ModelLiveTvSeriesTimerInfoDto) GetDays() []ModelModelDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetDaysOk() (*[]ModelModelDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ModelLiveTvSeriesTimerInfoDto) SetDays(v []ModelModelDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *ModelLiveTvSeriesTimerInfoDto) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetImageTags

`func (o *ModelLiveTvSeriesTimerInfoDto) GetImageTags() map[string]string`

GetImageTags returns the ImageTags field if non-nil, zero value otherwise.

### GetImageTagsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetImageTagsOk() (*map[string]string, bool)`

GetImageTagsOk returns a tuple with the ImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTags

`func (o *ModelLiveTvSeriesTimerInfoDto) SetImageTags(v map[string]string)`

SetImageTags sets ImageTags field to given value.

### HasImageTags

`func (o *ModelLiveTvSeriesTimerInfoDto) HasImageTags() bool`

HasImageTags returns a boolean if a field has been set.

### GetParentThumbItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentThumbItemId() string`

GetParentThumbItemId returns the ParentThumbItemId field if non-nil, zero value otherwise.

### GetParentThumbItemIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentThumbItemIdOk() (*string, bool)`

GetParentThumbItemIdOk returns a tuple with the ParentThumbItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentThumbItemId(v string)`

SetParentThumbItemId sets ParentThumbItemId field to given value.

### HasParentThumbItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentThumbItemId() bool`

HasParentThumbItemId returns a boolean if a field has been set.

### GetParentThumbImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentThumbImageTag() string`

GetParentThumbImageTag returns the ParentThumbImageTag field if non-nil, zero value otherwise.

### GetParentThumbImageTagOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentThumbImageTagOk() (*string, bool)`

GetParentThumbImageTagOk returns a tuple with the ParentThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentThumbImageTag(v string)`

SetParentThumbImageTag sets ParentThumbImageTag field to given value.

### HasParentThumbImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentThumbImageTag() bool`

HasParentThumbImageTag returns a boolean if a field has been set.

### GetParentPrimaryImageItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentPrimaryImageItemId() string`

GetParentPrimaryImageItemId returns the ParentPrimaryImageItemId field if non-nil, zero value otherwise.

### GetParentPrimaryImageItemIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentPrimaryImageItemIdOk() (*string, bool)`

GetParentPrimaryImageItemIdOk returns a tuple with the ParentPrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentPrimaryImageItemId(v string)`

SetParentPrimaryImageItemId sets ParentPrimaryImageItemId field to given value.

### HasParentPrimaryImageItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentPrimaryImageItemId() bool`

HasParentPrimaryImageItemId returns a boolean if a field has been set.

### GetParentPrimaryImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentPrimaryImageTag() string`

GetParentPrimaryImageTag returns the ParentPrimaryImageTag field if non-nil, zero value otherwise.

### GetParentPrimaryImageTagOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentPrimaryImageTagOk() (*string, bool)`

GetParentPrimaryImageTagOk returns a tuple with the ParentPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPrimaryImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentPrimaryImageTag(v string)`

SetParentPrimaryImageTag sets ParentPrimaryImageTag field to given value.

### HasParentPrimaryImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentPrimaryImageTag() bool`

HasParentPrimaryImageTag returns a boolean if a field has been set.

### GetSeriesId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetKeywords

`func (o *ModelLiveTvSeriesTimerInfoDto) GetKeywords() []ModelModelLiveTvKeywordInfo`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetKeywordsOk() (*[]ModelModelLiveTvKeywordInfo, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ModelLiveTvSeriesTimerInfoDto) SetKeywords(v []ModelModelLiveTvKeywordInfo)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ModelLiveTvSeriesTimerInfoDto) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetTimerType

`func (o *ModelLiveTvSeriesTimerInfoDto) GetTimerType() ModelModelLiveTvTimerType`

GetTimerType returns the TimerType field if non-nil, zero value otherwise.

### GetTimerTypeOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetTimerTypeOk() (*ModelModelLiveTvTimerType, bool)`

GetTimerTypeOk returns a tuple with the TimerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerType

`func (o *ModelLiveTvSeriesTimerInfoDto) SetTimerType(v ModelModelLiveTvTimerType)`

SetTimerType sets TimerType field to given value.

### HasTimerType

`func (o *ModelLiveTvSeriesTimerInfoDto) HasTimerType() bool`

HasTimerType returns a boolean if a field has been set.

### GetId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ModelLiveTvSeriesTimerInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelLiveTvSeriesTimerInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelLiveTvSeriesTimerInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetServerId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetChannelId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelName

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ModelLiveTvSeriesTimerInfoDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ModelLiveTvSeriesTimerInfoDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetChannelNumber

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelNumber() string`

GetChannelNumber returns the ChannelNumber field if non-nil, zero value otherwise.

### GetChannelNumberOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelNumberOk() (*string, bool)`

GetChannelNumberOk returns a tuple with the ChannelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNumber

`func (o *ModelLiveTvSeriesTimerInfoDto) SetChannelNumber(v string)`

SetChannelNumber sets ChannelNumber field to given value.

### HasChannelNumber

`func (o *ModelLiveTvSeriesTimerInfoDto) HasChannelNumber() bool`

HasChannelNumber returns a boolean if a field has been set.

### GetChannelPrimaryImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *ModelLiveTvSeriesTimerInfoDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### GetProgramId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetName

`func (o *ModelLiveTvSeriesTimerInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelLiveTvSeriesTimerInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelLiveTvSeriesTimerInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *ModelLiveTvSeriesTimerInfoDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelLiveTvSeriesTimerInfoDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelLiveTvSeriesTimerInfoDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetParentFolderId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetStartDate

`func (o *ModelLiveTvSeriesTimerInfoDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelLiveTvSeriesTimerInfoDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelLiveTvSeriesTimerInfoDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelLiveTvSeriesTimerInfoDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelLiveTvSeriesTimerInfoDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelLiveTvSeriesTimerInfoDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPriority

`func (o *ModelLiveTvSeriesTimerInfoDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelLiveTvSeriesTimerInfoDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelLiveTvSeriesTimerInfoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfoDto) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *ModelLiveTvSeriesTimerInfoDto) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *ModelLiveTvSeriesTimerInfoDto) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *ModelLiveTvSeriesTimerInfoDto) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### GetParentBackdropImageTags

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *ModelLiveTvSeriesTimerInfoDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *ModelLiveTvSeriesTimerInfoDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### GetIsPostPaddingRequired

`func (o *ModelLiveTvSeriesTimerInfoDto) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *ModelLiveTvSeriesTimerInfoDto) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *ModelLiveTvSeriesTimerInfoDto) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetKeepUntil

`func (o *ModelLiveTvSeriesTimerInfoDto) GetKeepUntil() ModelModelLiveTvKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *ModelLiveTvSeriesTimerInfoDto) GetKeepUntilOk() (*ModelModelLiveTvKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *ModelLiveTvSeriesTimerInfoDto) SetKeepUntil(v ModelModelLiveTvKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *ModelLiveTvSeriesTimerInfoDto) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


