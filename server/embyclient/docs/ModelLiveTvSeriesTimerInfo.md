# ModelLiveTvSeriesTimerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**ChannelIds** | Pointer to **[]string** |  | [optional] 
**ParentFolderId** | Pointer to **int64** |  | [optional] 
**ProgramId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**RecordAnyTime** | Pointer to **bool** |  | [optional] 
**KeepUpTo** | Pointer to **int32** |  | [optional] 
**KeepUntil** | Pointer to [**ModelModelLiveTvKeepUntil**](ModelLiveTvKeepUntil.md) |  | [optional] 
**SkipEpisodesInLibrary** | Pointer to **bool** |  | [optional] 
**RecordNewOnly** | Pointer to **bool** |  | [optional] 
**Days** | Pointer to [**[]ModelModelDayOfWeek**](ModelModelDayOfWeek.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**PrePaddingSeconds** | Pointer to **int32** |  | [optional] 
**PostPaddingSeconds** | Pointer to **int32** |  | [optional] 
**IsPrePaddingRequired** | Pointer to **bool** |  | [optional] 
**IsPostPaddingRequired** | Pointer to **bool** |  | [optional] 
**SeriesId** | Pointer to **string** |  | [optional] 
**ProviderIds** | Pointer to **map[string]string** |  | [optional] 
**MaxRecordingSeconds** | Pointer to **int32** |  | [optional] 
**Keywords** | Pointer to [**[]ModelModelLiveTvKeywordInfo**](ModelModelLiveTvKeywordInfo.md) |  | [optional] 
**TimerType** | Pointer to [**ModelModelLiveTvTimerType**](ModelLiveTvTimerType.md) |  | [optional] 

## Methods

### NewModelLiveTvSeriesTimerInfo

`func NewModelLiveTvSeriesTimerInfo() *ModelLiveTvSeriesTimerInfo`

NewModelLiveTvSeriesTimerInfo instantiates a new ModelLiveTvSeriesTimerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTvSeriesTimerInfoWithDefaults

`func NewModelLiveTvSeriesTimerInfoWithDefaults() *ModelLiveTvSeriesTimerInfo`

NewModelLiveTvSeriesTimerInfoWithDefaults instantiates a new ModelLiveTvSeriesTimerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelLiveTvSeriesTimerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelLiveTvSeriesTimerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelLiveTvSeriesTimerInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelLiveTvSeriesTimerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannelId

`func (o *ModelLiveTvSeriesTimerInfo) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ModelLiveTvSeriesTimerInfo) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ModelLiveTvSeriesTimerInfo) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ModelLiveTvSeriesTimerInfo) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelIds

`func (o *ModelLiveTvSeriesTimerInfo) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *ModelLiveTvSeriesTimerInfo) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *ModelLiveTvSeriesTimerInfo) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *ModelLiveTvSeriesTimerInfo) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### GetParentFolderId

`func (o *ModelLiveTvSeriesTimerInfo) GetParentFolderId() int64`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *ModelLiveTvSeriesTimerInfo) GetParentFolderIdOk() (*int64, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *ModelLiveTvSeriesTimerInfo) SetParentFolderId(v int64)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *ModelLiveTvSeriesTimerInfo) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### GetProgramId

`func (o *ModelLiveTvSeriesTimerInfo) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *ModelLiveTvSeriesTimerInfo) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *ModelLiveTvSeriesTimerInfo) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *ModelLiveTvSeriesTimerInfo) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### GetName

`func (o *ModelLiveTvSeriesTimerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelLiveTvSeriesTimerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelLiveTvSeriesTimerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelLiveTvSeriesTimerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceName

`func (o *ModelLiveTvSeriesTimerInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ModelLiveTvSeriesTimerInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ModelLiveTvSeriesTimerInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ModelLiveTvSeriesTimerInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetOverview

`func (o *ModelLiveTvSeriesTimerInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelLiveTvSeriesTimerInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelLiveTvSeriesTimerInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelLiveTvSeriesTimerInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetStartDate

`func (o *ModelLiveTvSeriesTimerInfo) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelLiveTvSeriesTimerInfo) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelLiveTvSeriesTimerInfo) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelLiveTvSeriesTimerInfo) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelLiveTvSeriesTimerInfo) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelLiveTvSeriesTimerInfo) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelLiveTvSeriesTimerInfo) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelLiveTvSeriesTimerInfo) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetRecordAnyTime

`func (o *ModelLiveTvSeriesTimerInfo) GetRecordAnyTime() bool`

GetRecordAnyTime returns the RecordAnyTime field if non-nil, zero value otherwise.

### GetRecordAnyTimeOk

`func (o *ModelLiveTvSeriesTimerInfo) GetRecordAnyTimeOk() (*bool, bool)`

GetRecordAnyTimeOk returns a tuple with the RecordAnyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordAnyTime

`func (o *ModelLiveTvSeriesTimerInfo) SetRecordAnyTime(v bool)`

SetRecordAnyTime sets RecordAnyTime field to given value.

### HasRecordAnyTime

`func (o *ModelLiveTvSeriesTimerInfo) HasRecordAnyTime() bool`

HasRecordAnyTime returns a boolean if a field has been set.

### GetKeepUpTo

`func (o *ModelLiveTvSeriesTimerInfo) GetKeepUpTo() int32`

GetKeepUpTo returns the KeepUpTo field if non-nil, zero value otherwise.

### GetKeepUpToOk

`func (o *ModelLiveTvSeriesTimerInfo) GetKeepUpToOk() (*int32, bool)`

GetKeepUpToOk returns a tuple with the KeepUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUpTo

`func (o *ModelLiveTvSeriesTimerInfo) SetKeepUpTo(v int32)`

SetKeepUpTo sets KeepUpTo field to given value.

### HasKeepUpTo

`func (o *ModelLiveTvSeriesTimerInfo) HasKeepUpTo() bool`

HasKeepUpTo returns a boolean if a field has been set.

### GetKeepUntil

`func (o *ModelLiveTvSeriesTimerInfo) GetKeepUntil() ModelModelLiveTvKeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *ModelLiveTvSeriesTimerInfo) GetKeepUntilOk() (*ModelModelLiveTvKeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *ModelLiveTvSeriesTimerInfo) SetKeepUntil(v ModelModelLiveTvKeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *ModelLiveTvSeriesTimerInfo) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.

### GetSkipEpisodesInLibrary

`func (o *ModelLiveTvSeriesTimerInfo) GetSkipEpisodesInLibrary() bool`

GetSkipEpisodesInLibrary returns the SkipEpisodesInLibrary field if non-nil, zero value otherwise.

### GetSkipEpisodesInLibraryOk

`func (o *ModelLiveTvSeriesTimerInfo) GetSkipEpisodesInLibraryOk() (*bool, bool)`

GetSkipEpisodesInLibraryOk returns a tuple with the SkipEpisodesInLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipEpisodesInLibrary

`func (o *ModelLiveTvSeriesTimerInfo) SetSkipEpisodesInLibrary(v bool)`

SetSkipEpisodesInLibrary sets SkipEpisodesInLibrary field to given value.

### HasSkipEpisodesInLibrary

`func (o *ModelLiveTvSeriesTimerInfo) HasSkipEpisodesInLibrary() bool`

HasSkipEpisodesInLibrary returns a boolean if a field has been set.

### GetRecordNewOnly

`func (o *ModelLiveTvSeriesTimerInfo) GetRecordNewOnly() bool`

GetRecordNewOnly returns the RecordNewOnly field if non-nil, zero value otherwise.

### GetRecordNewOnlyOk

`func (o *ModelLiveTvSeriesTimerInfo) GetRecordNewOnlyOk() (*bool, bool)`

GetRecordNewOnlyOk returns a tuple with the RecordNewOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNewOnly

`func (o *ModelLiveTvSeriesTimerInfo) SetRecordNewOnly(v bool)`

SetRecordNewOnly sets RecordNewOnly field to given value.

### HasRecordNewOnly

`func (o *ModelLiveTvSeriesTimerInfo) HasRecordNewOnly() bool`

HasRecordNewOnly returns a boolean if a field has been set.

### GetDays

`func (o *ModelLiveTvSeriesTimerInfo) GetDays() []ModelModelDayOfWeek`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ModelLiveTvSeriesTimerInfo) GetDaysOk() (*[]ModelModelDayOfWeek, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ModelLiveTvSeriesTimerInfo) SetDays(v []ModelModelDayOfWeek)`

SetDays sets Days field to given value.

### HasDays

`func (o *ModelLiveTvSeriesTimerInfo) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetPriority

`func (o *ModelLiveTvSeriesTimerInfo) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelLiveTvSeriesTimerInfo) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelLiveTvSeriesTimerInfo) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelLiveTvSeriesTimerInfo) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *ModelLiveTvSeriesTimerInfo) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *ModelLiveTvSeriesTimerInfo) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *ModelLiveTvSeriesTimerInfo) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *ModelLiveTvSeriesTimerInfo) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *ModelLiveTvSeriesTimerInfo) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *ModelLiveTvSeriesTimerInfo) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetIsPostPaddingRequired

`func (o *ModelLiveTvSeriesTimerInfo) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *ModelLiveTvSeriesTimerInfo) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *ModelLiveTvSeriesTimerInfo) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *ModelLiveTvSeriesTimerInfo) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetSeriesId

`func (o *ModelLiveTvSeriesTimerInfo) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *ModelLiveTvSeriesTimerInfo) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *ModelLiveTvSeriesTimerInfo) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *ModelLiveTvSeriesTimerInfo) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetProviderIds

`func (o *ModelLiveTvSeriesTimerInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *ModelLiveTvSeriesTimerInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *ModelLiveTvSeriesTimerInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *ModelLiveTvSeriesTimerInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetMaxRecordingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) GetMaxRecordingSeconds() int32`

GetMaxRecordingSeconds returns the MaxRecordingSeconds field if non-nil, zero value otherwise.

### GetMaxRecordingSecondsOk

`func (o *ModelLiveTvSeriesTimerInfo) GetMaxRecordingSecondsOk() (*int32, bool)`

GetMaxRecordingSecondsOk returns a tuple with the MaxRecordingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) SetMaxRecordingSeconds(v int32)`

SetMaxRecordingSeconds sets MaxRecordingSeconds field to given value.

### HasMaxRecordingSeconds

`func (o *ModelLiveTvSeriesTimerInfo) HasMaxRecordingSeconds() bool`

HasMaxRecordingSeconds returns a boolean if a field has been set.

### GetKeywords

`func (o *ModelLiveTvSeriesTimerInfo) GetKeywords() []ModelModelLiveTvKeywordInfo`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ModelLiveTvSeriesTimerInfo) GetKeywordsOk() (*[]ModelModelLiveTvKeywordInfo, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ModelLiveTvSeriesTimerInfo) SetKeywords(v []ModelModelLiveTvKeywordInfo)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ModelLiveTvSeriesTimerInfo) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetTimerType

`func (o *ModelLiveTvSeriesTimerInfo) GetTimerType() ModelModelLiveTvTimerType`

GetTimerType returns the TimerType field if non-nil, zero value otherwise.

### GetTimerTypeOk

`func (o *ModelLiveTvSeriesTimerInfo) GetTimerTypeOk() (*ModelModelLiveTvTimerType, bool)`

GetTimerTypeOk returns a tuple with the TimerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerType

`func (o *ModelLiveTvSeriesTimerInfo) SetTimerType(v ModelModelLiveTvTimerType)`

SetTimerType sets TimerType field to given value.

### HasTimerType

`func (o *ModelLiveTvSeriesTimerInfo) HasTimerType() bool`

HasTimerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


