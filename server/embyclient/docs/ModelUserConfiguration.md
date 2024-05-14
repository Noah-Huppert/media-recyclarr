# ModelUserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioLanguagePreference** | Pointer to **string** |  | [optional] 
**PlayDefaultAudioTrack** | Pointer to **bool** |  | [optional] 
**SubtitleLanguagePreference** | Pointer to **string** |  | [optional] 
**ProfilePin** | Pointer to **string** |  | [optional] 
**DisplayMissingEpisodes** | Pointer to **bool** |  | [optional] 
**SubtitleMode** | Pointer to [**ModelModelSubtitlePlaybackMode**](ModelSubtitlePlaybackMode.md) |  | [optional] 
**OrderedViews** | Pointer to **[]string** |  | [optional] 
**LatestItemsExcludes** | Pointer to **[]string** |  | [optional] 
**MyMediaExcludes** | Pointer to **[]string** |  | [optional] 
**HidePlayedInLatest** | Pointer to **bool** |  | [optional] 
**HidePlayedInMoreLikeThis** | Pointer to **bool** |  | [optional] 
**HidePlayedInSuggestions** | Pointer to **bool** |  | [optional] 
**RememberAudioSelections** | Pointer to **bool** |  | [optional] 
**RememberSubtitleSelections** | Pointer to **bool** |  | [optional] 
**EnableNextEpisodeAutoPlay** | Pointer to **bool** |  | [optional] 
**ResumeRewindSeconds** | Pointer to **int32** |  | [optional] 
**IntroSkipMode** | Pointer to [**ModelModelSegmentSkipMode**](ModelSegmentSkipMode.md) |  | [optional] 
**EnableLocalPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelUserConfiguration

`func NewModelUserConfiguration() *ModelUserConfiguration`

NewModelUserConfiguration instantiates a new ModelUserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserConfigurationWithDefaults

`func NewModelUserConfigurationWithDefaults() *ModelUserConfiguration`

NewModelUserConfigurationWithDefaults instantiates a new ModelUserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioLanguagePreference

`func (o *ModelUserConfiguration) GetAudioLanguagePreference() string`

GetAudioLanguagePreference returns the AudioLanguagePreference field if non-nil, zero value otherwise.

### GetAudioLanguagePreferenceOk

`func (o *ModelUserConfiguration) GetAudioLanguagePreferenceOk() (*string, bool)`

GetAudioLanguagePreferenceOk returns a tuple with the AudioLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguagePreference

`func (o *ModelUserConfiguration) SetAudioLanguagePreference(v string)`

SetAudioLanguagePreference sets AudioLanguagePreference field to given value.

### HasAudioLanguagePreference

`func (o *ModelUserConfiguration) HasAudioLanguagePreference() bool`

HasAudioLanguagePreference returns a boolean if a field has been set.

### GetPlayDefaultAudioTrack

`func (o *ModelUserConfiguration) GetPlayDefaultAudioTrack() bool`

GetPlayDefaultAudioTrack returns the PlayDefaultAudioTrack field if non-nil, zero value otherwise.

### GetPlayDefaultAudioTrackOk

`func (o *ModelUserConfiguration) GetPlayDefaultAudioTrackOk() (*bool, bool)`

GetPlayDefaultAudioTrackOk returns a tuple with the PlayDefaultAudioTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayDefaultAudioTrack

`func (o *ModelUserConfiguration) SetPlayDefaultAudioTrack(v bool)`

SetPlayDefaultAudioTrack sets PlayDefaultAudioTrack field to given value.

### HasPlayDefaultAudioTrack

`func (o *ModelUserConfiguration) HasPlayDefaultAudioTrack() bool`

HasPlayDefaultAudioTrack returns a boolean if a field has been set.

### GetSubtitleLanguagePreference

`func (o *ModelUserConfiguration) GetSubtitleLanguagePreference() string`

GetSubtitleLanguagePreference returns the SubtitleLanguagePreference field if non-nil, zero value otherwise.

### GetSubtitleLanguagePreferenceOk

`func (o *ModelUserConfiguration) GetSubtitleLanguagePreferenceOk() (*string, bool)`

GetSubtitleLanguagePreferenceOk returns a tuple with the SubtitleLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguagePreference

`func (o *ModelUserConfiguration) SetSubtitleLanguagePreference(v string)`

SetSubtitleLanguagePreference sets SubtitleLanguagePreference field to given value.

### HasSubtitleLanguagePreference

`func (o *ModelUserConfiguration) HasSubtitleLanguagePreference() bool`

HasSubtitleLanguagePreference returns a boolean if a field has been set.

### GetProfilePin

`func (o *ModelUserConfiguration) GetProfilePin() string`

GetProfilePin returns the ProfilePin field if non-nil, zero value otherwise.

### GetProfilePinOk

`func (o *ModelUserConfiguration) GetProfilePinOk() (*string, bool)`

GetProfilePinOk returns a tuple with the ProfilePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePin

`func (o *ModelUserConfiguration) SetProfilePin(v string)`

SetProfilePin sets ProfilePin field to given value.

### HasProfilePin

`func (o *ModelUserConfiguration) HasProfilePin() bool`

HasProfilePin returns a boolean if a field has been set.

### GetDisplayMissingEpisodes

`func (o *ModelUserConfiguration) GetDisplayMissingEpisodes() bool`

GetDisplayMissingEpisodes returns the DisplayMissingEpisodes field if non-nil, zero value otherwise.

### GetDisplayMissingEpisodesOk

`func (o *ModelUserConfiguration) GetDisplayMissingEpisodesOk() (*bool, bool)`

GetDisplayMissingEpisodesOk returns a tuple with the DisplayMissingEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMissingEpisodes

`func (o *ModelUserConfiguration) SetDisplayMissingEpisodes(v bool)`

SetDisplayMissingEpisodes sets DisplayMissingEpisodes field to given value.

### HasDisplayMissingEpisodes

`func (o *ModelUserConfiguration) HasDisplayMissingEpisodes() bool`

HasDisplayMissingEpisodes returns a boolean if a field has been set.

### GetSubtitleMode

`func (o *ModelUserConfiguration) GetSubtitleMode() ModelModelSubtitlePlaybackMode`

GetSubtitleMode returns the SubtitleMode field if non-nil, zero value otherwise.

### GetSubtitleModeOk

`func (o *ModelUserConfiguration) GetSubtitleModeOk() (*ModelModelSubtitlePlaybackMode, bool)`

GetSubtitleModeOk returns a tuple with the SubtitleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleMode

`func (o *ModelUserConfiguration) SetSubtitleMode(v ModelModelSubtitlePlaybackMode)`

SetSubtitleMode sets SubtitleMode field to given value.

### HasSubtitleMode

`func (o *ModelUserConfiguration) HasSubtitleMode() bool`

HasSubtitleMode returns a boolean if a field has been set.

### GetOrderedViews

`func (o *ModelUserConfiguration) GetOrderedViews() []string`

GetOrderedViews returns the OrderedViews field if non-nil, zero value otherwise.

### GetOrderedViewsOk

`func (o *ModelUserConfiguration) GetOrderedViewsOk() (*[]string, bool)`

GetOrderedViewsOk returns a tuple with the OrderedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedViews

`func (o *ModelUserConfiguration) SetOrderedViews(v []string)`

SetOrderedViews sets OrderedViews field to given value.

### HasOrderedViews

`func (o *ModelUserConfiguration) HasOrderedViews() bool`

HasOrderedViews returns a boolean if a field has been set.

### GetLatestItemsExcludes

`func (o *ModelUserConfiguration) GetLatestItemsExcludes() []string`

GetLatestItemsExcludes returns the LatestItemsExcludes field if non-nil, zero value otherwise.

### GetLatestItemsExcludesOk

`func (o *ModelUserConfiguration) GetLatestItemsExcludesOk() (*[]string, bool)`

GetLatestItemsExcludesOk returns a tuple with the LatestItemsExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestItemsExcludes

`func (o *ModelUserConfiguration) SetLatestItemsExcludes(v []string)`

SetLatestItemsExcludes sets LatestItemsExcludes field to given value.

### HasLatestItemsExcludes

`func (o *ModelUserConfiguration) HasLatestItemsExcludes() bool`

HasLatestItemsExcludes returns a boolean if a field has been set.

### GetMyMediaExcludes

`func (o *ModelUserConfiguration) GetMyMediaExcludes() []string`

GetMyMediaExcludes returns the MyMediaExcludes field if non-nil, zero value otherwise.

### GetMyMediaExcludesOk

`func (o *ModelUserConfiguration) GetMyMediaExcludesOk() (*[]string, bool)`

GetMyMediaExcludesOk returns a tuple with the MyMediaExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyMediaExcludes

`func (o *ModelUserConfiguration) SetMyMediaExcludes(v []string)`

SetMyMediaExcludes sets MyMediaExcludes field to given value.

### HasMyMediaExcludes

`func (o *ModelUserConfiguration) HasMyMediaExcludes() bool`

HasMyMediaExcludes returns a boolean if a field has been set.

### GetHidePlayedInLatest

`func (o *ModelUserConfiguration) GetHidePlayedInLatest() bool`

GetHidePlayedInLatest returns the HidePlayedInLatest field if non-nil, zero value otherwise.

### GetHidePlayedInLatestOk

`func (o *ModelUserConfiguration) GetHidePlayedInLatestOk() (*bool, bool)`

GetHidePlayedInLatestOk returns a tuple with the HidePlayedInLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePlayedInLatest

`func (o *ModelUserConfiguration) SetHidePlayedInLatest(v bool)`

SetHidePlayedInLatest sets HidePlayedInLatest field to given value.

### HasHidePlayedInLatest

`func (o *ModelUserConfiguration) HasHidePlayedInLatest() bool`

HasHidePlayedInLatest returns a boolean if a field has been set.

### GetHidePlayedInMoreLikeThis

`func (o *ModelUserConfiguration) GetHidePlayedInMoreLikeThis() bool`

GetHidePlayedInMoreLikeThis returns the HidePlayedInMoreLikeThis field if non-nil, zero value otherwise.

### GetHidePlayedInMoreLikeThisOk

`func (o *ModelUserConfiguration) GetHidePlayedInMoreLikeThisOk() (*bool, bool)`

GetHidePlayedInMoreLikeThisOk returns a tuple with the HidePlayedInMoreLikeThis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePlayedInMoreLikeThis

`func (o *ModelUserConfiguration) SetHidePlayedInMoreLikeThis(v bool)`

SetHidePlayedInMoreLikeThis sets HidePlayedInMoreLikeThis field to given value.

### HasHidePlayedInMoreLikeThis

`func (o *ModelUserConfiguration) HasHidePlayedInMoreLikeThis() bool`

HasHidePlayedInMoreLikeThis returns a boolean if a field has been set.

### GetHidePlayedInSuggestions

`func (o *ModelUserConfiguration) GetHidePlayedInSuggestions() bool`

GetHidePlayedInSuggestions returns the HidePlayedInSuggestions field if non-nil, zero value otherwise.

### GetHidePlayedInSuggestionsOk

`func (o *ModelUserConfiguration) GetHidePlayedInSuggestionsOk() (*bool, bool)`

GetHidePlayedInSuggestionsOk returns a tuple with the HidePlayedInSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePlayedInSuggestions

`func (o *ModelUserConfiguration) SetHidePlayedInSuggestions(v bool)`

SetHidePlayedInSuggestions sets HidePlayedInSuggestions field to given value.

### HasHidePlayedInSuggestions

`func (o *ModelUserConfiguration) HasHidePlayedInSuggestions() bool`

HasHidePlayedInSuggestions returns a boolean if a field has been set.

### GetRememberAudioSelections

`func (o *ModelUserConfiguration) GetRememberAudioSelections() bool`

GetRememberAudioSelections returns the RememberAudioSelections field if non-nil, zero value otherwise.

### GetRememberAudioSelectionsOk

`func (o *ModelUserConfiguration) GetRememberAudioSelectionsOk() (*bool, bool)`

GetRememberAudioSelectionsOk returns a tuple with the RememberAudioSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberAudioSelections

`func (o *ModelUserConfiguration) SetRememberAudioSelections(v bool)`

SetRememberAudioSelections sets RememberAudioSelections field to given value.

### HasRememberAudioSelections

`func (o *ModelUserConfiguration) HasRememberAudioSelections() bool`

HasRememberAudioSelections returns a boolean if a field has been set.

### GetRememberSubtitleSelections

`func (o *ModelUserConfiguration) GetRememberSubtitleSelections() bool`

GetRememberSubtitleSelections returns the RememberSubtitleSelections field if non-nil, zero value otherwise.

### GetRememberSubtitleSelectionsOk

`func (o *ModelUserConfiguration) GetRememberSubtitleSelectionsOk() (*bool, bool)`

GetRememberSubtitleSelectionsOk returns a tuple with the RememberSubtitleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberSubtitleSelections

`func (o *ModelUserConfiguration) SetRememberSubtitleSelections(v bool)`

SetRememberSubtitleSelections sets RememberSubtitleSelections field to given value.

### HasRememberSubtitleSelections

`func (o *ModelUserConfiguration) HasRememberSubtitleSelections() bool`

HasRememberSubtitleSelections returns a boolean if a field has been set.

### GetEnableNextEpisodeAutoPlay

`func (o *ModelUserConfiguration) GetEnableNextEpisodeAutoPlay() bool`

GetEnableNextEpisodeAutoPlay returns the EnableNextEpisodeAutoPlay field if non-nil, zero value otherwise.

### GetEnableNextEpisodeAutoPlayOk

`func (o *ModelUserConfiguration) GetEnableNextEpisodeAutoPlayOk() (*bool, bool)`

GetEnableNextEpisodeAutoPlayOk returns a tuple with the EnableNextEpisodeAutoPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNextEpisodeAutoPlay

`func (o *ModelUserConfiguration) SetEnableNextEpisodeAutoPlay(v bool)`

SetEnableNextEpisodeAutoPlay sets EnableNextEpisodeAutoPlay field to given value.

### HasEnableNextEpisodeAutoPlay

`func (o *ModelUserConfiguration) HasEnableNextEpisodeAutoPlay() bool`

HasEnableNextEpisodeAutoPlay returns a boolean if a field has been set.

### GetResumeRewindSeconds

`func (o *ModelUserConfiguration) GetResumeRewindSeconds() int32`

GetResumeRewindSeconds returns the ResumeRewindSeconds field if non-nil, zero value otherwise.

### GetResumeRewindSecondsOk

`func (o *ModelUserConfiguration) GetResumeRewindSecondsOk() (*int32, bool)`

GetResumeRewindSecondsOk returns a tuple with the ResumeRewindSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeRewindSeconds

`func (o *ModelUserConfiguration) SetResumeRewindSeconds(v int32)`

SetResumeRewindSeconds sets ResumeRewindSeconds field to given value.

### HasResumeRewindSeconds

`func (o *ModelUserConfiguration) HasResumeRewindSeconds() bool`

HasResumeRewindSeconds returns a boolean if a field has been set.

### GetIntroSkipMode

`func (o *ModelUserConfiguration) GetIntroSkipMode() ModelModelSegmentSkipMode`

GetIntroSkipMode returns the IntroSkipMode field if non-nil, zero value otherwise.

### GetIntroSkipModeOk

`func (o *ModelUserConfiguration) GetIntroSkipModeOk() (*ModelModelSegmentSkipMode, bool)`

GetIntroSkipModeOk returns a tuple with the IntroSkipMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroSkipMode

`func (o *ModelUserConfiguration) SetIntroSkipMode(v ModelModelSegmentSkipMode)`

SetIntroSkipMode sets IntroSkipMode field to given value.

### HasIntroSkipMode

`func (o *ModelUserConfiguration) HasIntroSkipMode() bool`

HasIntroSkipMode returns a boolean if a field has been set.

### GetEnableLocalPassword

`func (o *ModelUserConfiguration) GetEnableLocalPassword() bool`

GetEnableLocalPassword returns the EnableLocalPassword field if non-nil, zero value otherwise.

### GetEnableLocalPasswordOk

`func (o *ModelUserConfiguration) GetEnableLocalPasswordOk() (*bool, bool)`

GetEnableLocalPasswordOk returns a tuple with the EnableLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalPassword

`func (o *ModelUserConfiguration) SetEnableLocalPassword(v bool)`

SetEnableLocalPassword sets EnableLocalPassword field to given value.

### HasEnableLocalPassword

`func (o *ModelUserConfiguration) HasEnableLocalPassword() bool`

HasEnableLocalPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


