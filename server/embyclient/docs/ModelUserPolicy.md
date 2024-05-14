# ModelUserPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdministrator** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsHiddenRemotely** | Pointer to **bool** |  | [optional] 
**IsHiddenFromUnusedDevices** | Pointer to **bool** |  | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**LockedOutDate** | Pointer to **int64** |  | [optional] 
**MaxParentalRating** | Pointer to **NullableInt32** |  | [optional] 
**AllowTagOrRating** | Pointer to **bool** |  | [optional] 
**BlockedTags** | Pointer to **[]string** |  | [optional] 
**IsTagBlockingModeInclusive** | Pointer to **bool** |  | [optional] 
**IncludeTags** | Pointer to **[]string** |  | [optional] 
**EnableUserPreferenceAccess** | Pointer to **bool** |  | [optional] 
**AccessSchedules** | Pointer to [**[]ModelModelAccessSchedule**](ModelModelAccessSchedule.md) |  | [optional] 
**BlockUnratedItems** | Pointer to [**[]ModelModelUnratedItem**](ModelModelUnratedItem.md) |  | [optional] 
**EnableRemoteControlOfOtherUsers** | Pointer to **bool** |  | [optional] 
**EnableSharedDeviceControl** | Pointer to **bool** |  | [optional] 
**EnableRemoteAccess** | Pointer to **bool** |  | [optional] 
**EnableLiveTvManagement** | Pointer to **bool** |  | [optional] 
**EnableLiveTvAccess** | Pointer to **bool** |  | [optional] 
**EnableMediaPlayback** | Pointer to **bool** |  | [optional] 
**EnableAudioPlaybackTranscoding** | Pointer to **bool** |  | [optional] 
**EnableVideoPlaybackTranscoding** | Pointer to **bool** |  | [optional] 
**EnablePlaybackRemuxing** | Pointer to **bool** |  | [optional] 
**EnableContentDeletion** | Pointer to **bool** |  | [optional] 
**RestrictedFeatures** | Pointer to **[]string** |  | [optional] 
**EnableContentDeletionFromFolders** | Pointer to **[]string** |  | [optional] 
**EnableContentDownloading** | Pointer to **bool** |  | [optional] 
**EnableSubtitleDownloading** | Pointer to **bool** |  | [optional] 
**EnableSubtitleManagement** | Pointer to **bool** |  | [optional] 
**EnableSyncTranscoding** | Pointer to **bool** |  | [optional] 
**EnableMediaConversion** | Pointer to **bool** |  | [optional] 
**EnabledChannels** | Pointer to **[]string** |  | [optional] 
**EnableAllChannels** | Pointer to **bool** |  | [optional] 
**EnabledFolders** | Pointer to **[]string** |  | [optional] 
**EnableAllFolders** | Pointer to **bool** |  | [optional] 
**InvalidLoginAttemptCount** | Pointer to **int32** |  | [optional] 
**EnablePublicSharing** | Pointer to **bool** |  | [optional] 
**BlockedMediaFolders** | Pointer to **[]string** |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**AuthenticationProviderId** | Pointer to **string** |  | [optional] 
**ExcludedSubFolders** | Pointer to **[]string** |  | [optional] 
**SimultaneousStreamLimit** | Pointer to **int32** |  | [optional] 
**EnabledDevices** | Pointer to **[]string** |  | [optional] 
**EnableAllDevices** | Pointer to **bool** |  | [optional] 
**AllowCameraUpload** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelUserPolicy

`func NewModelUserPolicy() *ModelUserPolicy`

NewModelUserPolicy instantiates a new ModelUserPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserPolicyWithDefaults

`func NewModelUserPolicyWithDefaults() *ModelUserPolicy`

NewModelUserPolicyWithDefaults instantiates a new ModelUserPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdministrator

`func (o *ModelUserPolicy) GetIsAdministrator() bool`

GetIsAdministrator returns the IsAdministrator field if non-nil, zero value otherwise.

### GetIsAdministratorOk

`func (o *ModelUserPolicy) GetIsAdministratorOk() (*bool, bool)`

GetIsAdministratorOk returns a tuple with the IsAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdministrator

`func (o *ModelUserPolicy) SetIsAdministrator(v bool)`

SetIsAdministrator sets IsAdministrator field to given value.

### HasIsAdministrator

`func (o *ModelUserPolicy) HasIsAdministrator() bool`

HasIsAdministrator returns a boolean if a field has been set.

### GetIsHidden

`func (o *ModelUserPolicy) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ModelUserPolicy) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ModelUserPolicy) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *ModelUserPolicy) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsHiddenRemotely

`func (o *ModelUserPolicy) GetIsHiddenRemotely() bool`

GetIsHiddenRemotely returns the IsHiddenRemotely field if non-nil, zero value otherwise.

### GetIsHiddenRemotelyOk

`func (o *ModelUserPolicy) GetIsHiddenRemotelyOk() (*bool, bool)`

GetIsHiddenRemotelyOk returns a tuple with the IsHiddenRemotely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHiddenRemotely

`func (o *ModelUserPolicy) SetIsHiddenRemotely(v bool)`

SetIsHiddenRemotely sets IsHiddenRemotely field to given value.

### HasIsHiddenRemotely

`func (o *ModelUserPolicy) HasIsHiddenRemotely() bool`

HasIsHiddenRemotely returns a boolean if a field has been set.

### GetIsHiddenFromUnusedDevices

`func (o *ModelUserPolicy) GetIsHiddenFromUnusedDevices() bool`

GetIsHiddenFromUnusedDevices returns the IsHiddenFromUnusedDevices field if non-nil, zero value otherwise.

### GetIsHiddenFromUnusedDevicesOk

`func (o *ModelUserPolicy) GetIsHiddenFromUnusedDevicesOk() (*bool, bool)`

GetIsHiddenFromUnusedDevicesOk returns a tuple with the IsHiddenFromUnusedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHiddenFromUnusedDevices

`func (o *ModelUserPolicy) SetIsHiddenFromUnusedDevices(v bool)`

SetIsHiddenFromUnusedDevices sets IsHiddenFromUnusedDevices field to given value.

### HasIsHiddenFromUnusedDevices

`func (o *ModelUserPolicy) HasIsHiddenFromUnusedDevices() bool`

HasIsHiddenFromUnusedDevices returns a boolean if a field has been set.

### GetIsDisabled

`func (o *ModelUserPolicy) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *ModelUserPolicy) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *ModelUserPolicy) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *ModelUserPolicy) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetLockedOutDate

`func (o *ModelUserPolicy) GetLockedOutDate() int64`

GetLockedOutDate returns the LockedOutDate field if non-nil, zero value otherwise.

### GetLockedOutDateOk

`func (o *ModelUserPolicy) GetLockedOutDateOk() (*int64, bool)`

GetLockedOutDateOk returns a tuple with the LockedOutDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedOutDate

`func (o *ModelUserPolicy) SetLockedOutDate(v int64)`

SetLockedOutDate sets LockedOutDate field to given value.

### HasLockedOutDate

`func (o *ModelUserPolicy) HasLockedOutDate() bool`

HasLockedOutDate returns a boolean if a field has been set.

### GetMaxParentalRating

`func (o *ModelUserPolicy) GetMaxParentalRating() int32`

GetMaxParentalRating returns the MaxParentalRating field if non-nil, zero value otherwise.

### GetMaxParentalRatingOk

`func (o *ModelUserPolicy) GetMaxParentalRatingOk() (*int32, bool)`

GetMaxParentalRatingOk returns a tuple with the MaxParentalRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParentalRating

`func (o *ModelUserPolicy) SetMaxParentalRating(v int32)`

SetMaxParentalRating sets MaxParentalRating field to given value.

### HasMaxParentalRating

`func (o *ModelUserPolicy) HasMaxParentalRating() bool`

HasMaxParentalRating returns a boolean if a field has been set.

### SetMaxParentalRatingNil

`func (o *ModelUserPolicy) SetMaxParentalRatingNil(b bool)`

 SetMaxParentalRatingNil sets the value for MaxParentalRating to be an explicit nil

### UnsetMaxParentalRating
`func (o *ModelUserPolicy) UnsetMaxParentalRating()`

UnsetMaxParentalRating ensures that no value is present for MaxParentalRating, not even an explicit nil
### GetAllowTagOrRating

`func (o *ModelUserPolicy) GetAllowTagOrRating() bool`

GetAllowTagOrRating returns the AllowTagOrRating field if non-nil, zero value otherwise.

### GetAllowTagOrRatingOk

`func (o *ModelUserPolicy) GetAllowTagOrRatingOk() (*bool, bool)`

GetAllowTagOrRatingOk returns a tuple with the AllowTagOrRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTagOrRating

`func (o *ModelUserPolicy) SetAllowTagOrRating(v bool)`

SetAllowTagOrRating sets AllowTagOrRating field to given value.

### HasAllowTagOrRating

`func (o *ModelUserPolicy) HasAllowTagOrRating() bool`

HasAllowTagOrRating returns a boolean if a field has been set.

### GetBlockedTags

`func (o *ModelUserPolicy) GetBlockedTags() []string`

GetBlockedTags returns the BlockedTags field if non-nil, zero value otherwise.

### GetBlockedTagsOk

`func (o *ModelUserPolicy) GetBlockedTagsOk() (*[]string, bool)`

GetBlockedTagsOk returns a tuple with the BlockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTags

`func (o *ModelUserPolicy) SetBlockedTags(v []string)`

SetBlockedTags sets BlockedTags field to given value.

### HasBlockedTags

`func (o *ModelUserPolicy) HasBlockedTags() bool`

HasBlockedTags returns a boolean if a field has been set.

### GetIsTagBlockingModeInclusive

`func (o *ModelUserPolicy) GetIsTagBlockingModeInclusive() bool`

GetIsTagBlockingModeInclusive returns the IsTagBlockingModeInclusive field if non-nil, zero value otherwise.

### GetIsTagBlockingModeInclusiveOk

`func (o *ModelUserPolicy) GetIsTagBlockingModeInclusiveOk() (*bool, bool)`

GetIsTagBlockingModeInclusiveOk returns a tuple with the IsTagBlockingModeInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTagBlockingModeInclusive

`func (o *ModelUserPolicy) SetIsTagBlockingModeInclusive(v bool)`

SetIsTagBlockingModeInclusive sets IsTagBlockingModeInclusive field to given value.

### HasIsTagBlockingModeInclusive

`func (o *ModelUserPolicy) HasIsTagBlockingModeInclusive() bool`

HasIsTagBlockingModeInclusive returns a boolean if a field has been set.

### GetIncludeTags

`func (o *ModelUserPolicy) GetIncludeTags() []string`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ModelUserPolicy) GetIncludeTagsOk() (*[]string, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ModelUserPolicy) SetIncludeTags(v []string)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ModelUserPolicy) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetEnableUserPreferenceAccess

`func (o *ModelUserPolicy) GetEnableUserPreferenceAccess() bool`

GetEnableUserPreferenceAccess returns the EnableUserPreferenceAccess field if non-nil, zero value otherwise.

### GetEnableUserPreferenceAccessOk

`func (o *ModelUserPolicy) GetEnableUserPreferenceAccessOk() (*bool, bool)`

GetEnableUserPreferenceAccessOk returns a tuple with the EnableUserPreferenceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPreferenceAccess

`func (o *ModelUserPolicy) SetEnableUserPreferenceAccess(v bool)`

SetEnableUserPreferenceAccess sets EnableUserPreferenceAccess field to given value.

### HasEnableUserPreferenceAccess

`func (o *ModelUserPolicy) HasEnableUserPreferenceAccess() bool`

HasEnableUserPreferenceAccess returns a boolean if a field has been set.

### GetAccessSchedules

`func (o *ModelUserPolicy) GetAccessSchedules() []ModelModelAccessSchedule`

GetAccessSchedules returns the AccessSchedules field if non-nil, zero value otherwise.

### GetAccessSchedulesOk

`func (o *ModelUserPolicy) GetAccessSchedulesOk() (*[]ModelModelAccessSchedule, bool)`

GetAccessSchedulesOk returns a tuple with the AccessSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSchedules

`func (o *ModelUserPolicy) SetAccessSchedules(v []ModelModelAccessSchedule)`

SetAccessSchedules sets AccessSchedules field to given value.

### HasAccessSchedules

`func (o *ModelUserPolicy) HasAccessSchedules() bool`

HasAccessSchedules returns a boolean if a field has been set.

### GetBlockUnratedItems

`func (o *ModelUserPolicy) GetBlockUnratedItems() []ModelModelUnratedItem`

GetBlockUnratedItems returns the BlockUnratedItems field if non-nil, zero value otherwise.

### GetBlockUnratedItemsOk

`func (o *ModelUserPolicy) GetBlockUnratedItemsOk() (*[]ModelModelUnratedItem, bool)`

GetBlockUnratedItemsOk returns a tuple with the BlockUnratedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUnratedItems

`func (o *ModelUserPolicy) SetBlockUnratedItems(v []ModelModelUnratedItem)`

SetBlockUnratedItems sets BlockUnratedItems field to given value.

### HasBlockUnratedItems

`func (o *ModelUserPolicy) HasBlockUnratedItems() bool`

HasBlockUnratedItems returns a boolean if a field has been set.

### GetEnableRemoteControlOfOtherUsers

`func (o *ModelUserPolicy) GetEnableRemoteControlOfOtherUsers() bool`

GetEnableRemoteControlOfOtherUsers returns the EnableRemoteControlOfOtherUsers field if non-nil, zero value otherwise.

### GetEnableRemoteControlOfOtherUsersOk

`func (o *ModelUserPolicy) GetEnableRemoteControlOfOtherUsersOk() (*bool, bool)`

GetEnableRemoteControlOfOtherUsersOk returns a tuple with the EnableRemoteControlOfOtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteControlOfOtherUsers

`func (o *ModelUserPolicy) SetEnableRemoteControlOfOtherUsers(v bool)`

SetEnableRemoteControlOfOtherUsers sets EnableRemoteControlOfOtherUsers field to given value.

### HasEnableRemoteControlOfOtherUsers

`func (o *ModelUserPolicy) HasEnableRemoteControlOfOtherUsers() bool`

HasEnableRemoteControlOfOtherUsers returns a boolean if a field has been set.

### GetEnableSharedDeviceControl

`func (o *ModelUserPolicy) GetEnableSharedDeviceControl() bool`

GetEnableSharedDeviceControl returns the EnableSharedDeviceControl field if non-nil, zero value otherwise.

### GetEnableSharedDeviceControlOk

`func (o *ModelUserPolicy) GetEnableSharedDeviceControlOk() (*bool, bool)`

GetEnableSharedDeviceControlOk returns a tuple with the EnableSharedDeviceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSharedDeviceControl

`func (o *ModelUserPolicy) SetEnableSharedDeviceControl(v bool)`

SetEnableSharedDeviceControl sets EnableSharedDeviceControl field to given value.

### HasEnableSharedDeviceControl

`func (o *ModelUserPolicy) HasEnableSharedDeviceControl() bool`

HasEnableSharedDeviceControl returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *ModelUserPolicy) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *ModelUserPolicy) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *ModelUserPolicy) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *ModelUserPolicy) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetEnableLiveTvManagement

`func (o *ModelUserPolicy) GetEnableLiveTvManagement() bool`

GetEnableLiveTvManagement returns the EnableLiveTvManagement field if non-nil, zero value otherwise.

### GetEnableLiveTvManagementOk

`func (o *ModelUserPolicy) GetEnableLiveTvManagementOk() (*bool, bool)`

GetEnableLiveTvManagementOk returns a tuple with the EnableLiveTvManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvManagement

`func (o *ModelUserPolicy) SetEnableLiveTvManagement(v bool)`

SetEnableLiveTvManagement sets EnableLiveTvManagement field to given value.

### HasEnableLiveTvManagement

`func (o *ModelUserPolicy) HasEnableLiveTvManagement() bool`

HasEnableLiveTvManagement returns a boolean if a field has been set.

### GetEnableLiveTvAccess

`func (o *ModelUserPolicy) GetEnableLiveTvAccess() bool`

GetEnableLiveTvAccess returns the EnableLiveTvAccess field if non-nil, zero value otherwise.

### GetEnableLiveTvAccessOk

`func (o *ModelUserPolicy) GetEnableLiveTvAccessOk() (*bool, bool)`

GetEnableLiveTvAccessOk returns a tuple with the EnableLiveTvAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvAccess

`func (o *ModelUserPolicy) SetEnableLiveTvAccess(v bool)`

SetEnableLiveTvAccess sets EnableLiveTvAccess field to given value.

### HasEnableLiveTvAccess

`func (o *ModelUserPolicy) HasEnableLiveTvAccess() bool`

HasEnableLiveTvAccess returns a boolean if a field has been set.

### GetEnableMediaPlayback

`func (o *ModelUserPolicy) GetEnableMediaPlayback() bool`

GetEnableMediaPlayback returns the EnableMediaPlayback field if non-nil, zero value otherwise.

### GetEnableMediaPlaybackOk

`func (o *ModelUserPolicy) GetEnableMediaPlaybackOk() (*bool, bool)`

GetEnableMediaPlaybackOk returns a tuple with the EnableMediaPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaPlayback

`func (o *ModelUserPolicy) SetEnableMediaPlayback(v bool)`

SetEnableMediaPlayback sets EnableMediaPlayback field to given value.

### HasEnableMediaPlayback

`func (o *ModelUserPolicy) HasEnableMediaPlayback() bool`

HasEnableMediaPlayback returns a boolean if a field has been set.

### GetEnableAudioPlaybackTranscoding

`func (o *ModelUserPolicy) GetEnableAudioPlaybackTranscoding() bool`

GetEnableAudioPlaybackTranscoding returns the EnableAudioPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableAudioPlaybackTranscodingOk

`func (o *ModelUserPolicy) GetEnableAudioPlaybackTranscodingOk() (*bool, bool)`

GetEnableAudioPlaybackTranscodingOk returns a tuple with the EnableAudioPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioPlaybackTranscoding

`func (o *ModelUserPolicy) SetEnableAudioPlaybackTranscoding(v bool)`

SetEnableAudioPlaybackTranscoding sets EnableAudioPlaybackTranscoding field to given value.

### HasEnableAudioPlaybackTranscoding

`func (o *ModelUserPolicy) HasEnableAudioPlaybackTranscoding() bool`

HasEnableAudioPlaybackTranscoding returns a boolean if a field has been set.

### GetEnableVideoPlaybackTranscoding

`func (o *ModelUserPolicy) GetEnableVideoPlaybackTranscoding() bool`

GetEnableVideoPlaybackTranscoding returns the EnableVideoPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableVideoPlaybackTranscodingOk

`func (o *ModelUserPolicy) GetEnableVideoPlaybackTranscodingOk() (*bool, bool)`

GetEnableVideoPlaybackTranscodingOk returns a tuple with the EnableVideoPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoPlaybackTranscoding

`func (o *ModelUserPolicy) SetEnableVideoPlaybackTranscoding(v bool)`

SetEnableVideoPlaybackTranscoding sets EnableVideoPlaybackTranscoding field to given value.

### HasEnableVideoPlaybackTranscoding

`func (o *ModelUserPolicy) HasEnableVideoPlaybackTranscoding() bool`

HasEnableVideoPlaybackTranscoding returns a boolean if a field has been set.

### GetEnablePlaybackRemuxing

`func (o *ModelUserPolicy) GetEnablePlaybackRemuxing() bool`

GetEnablePlaybackRemuxing returns the EnablePlaybackRemuxing field if non-nil, zero value otherwise.

### GetEnablePlaybackRemuxingOk

`func (o *ModelUserPolicy) GetEnablePlaybackRemuxingOk() (*bool, bool)`

GetEnablePlaybackRemuxingOk returns a tuple with the EnablePlaybackRemuxing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlaybackRemuxing

`func (o *ModelUserPolicy) SetEnablePlaybackRemuxing(v bool)`

SetEnablePlaybackRemuxing sets EnablePlaybackRemuxing field to given value.

### HasEnablePlaybackRemuxing

`func (o *ModelUserPolicy) HasEnablePlaybackRemuxing() bool`

HasEnablePlaybackRemuxing returns a boolean if a field has been set.

### GetEnableContentDeletion

`func (o *ModelUserPolicy) GetEnableContentDeletion() bool`

GetEnableContentDeletion returns the EnableContentDeletion field if non-nil, zero value otherwise.

### GetEnableContentDeletionOk

`func (o *ModelUserPolicy) GetEnableContentDeletionOk() (*bool, bool)`

GetEnableContentDeletionOk returns a tuple with the EnableContentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletion

`func (o *ModelUserPolicy) SetEnableContentDeletion(v bool)`

SetEnableContentDeletion sets EnableContentDeletion field to given value.

### HasEnableContentDeletion

`func (o *ModelUserPolicy) HasEnableContentDeletion() bool`

HasEnableContentDeletion returns a boolean if a field has been set.

### GetRestrictedFeatures

`func (o *ModelUserPolicy) GetRestrictedFeatures() []string`

GetRestrictedFeatures returns the RestrictedFeatures field if non-nil, zero value otherwise.

### GetRestrictedFeaturesOk

`func (o *ModelUserPolicy) GetRestrictedFeaturesOk() (*[]string, bool)`

GetRestrictedFeaturesOk returns a tuple with the RestrictedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedFeatures

`func (o *ModelUserPolicy) SetRestrictedFeatures(v []string)`

SetRestrictedFeatures sets RestrictedFeatures field to given value.

### HasRestrictedFeatures

`func (o *ModelUserPolicy) HasRestrictedFeatures() bool`

HasRestrictedFeatures returns a boolean if a field has been set.

### GetEnableContentDeletionFromFolders

`func (o *ModelUserPolicy) GetEnableContentDeletionFromFolders() []string`

GetEnableContentDeletionFromFolders returns the EnableContentDeletionFromFolders field if non-nil, zero value otherwise.

### GetEnableContentDeletionFromFoldersOk

`func (o *ModelUserPolicy) GetEnableContentDeletionFromFoldersOk() (*[]string, bool)`

GetEnableContentDeletionFromFoldersOk returns a tuple with the EnableContentDeletionFromFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletionFromFolders

`func (o *ModelUserPolicy) SetEnableContentDeletionFromFolders(v []string)`

SetEnableContentDeletionFromFolders sets EnableContentDeletionFromFolders field to given value.

### HasEnableContentDeletionFromFolders

`func (o *ModelUserPolicy) HasEnableContentDeletionFromFolders() bool`

HasEnableContentDeletionFromFolders returns a boolean if a field has been set.

### GetEnableContentDownloading

`func (o *ModelUserPolicy) GetEnableContentDownloading() bool`

GetEnableContentDownloading returns the EnableContentDownloading field if non-nil, zero value otherwise.

### GetEnableContentDownloadingOk

`func (o *ModelUserPolicy) GetEnableContentDownloadingOk() (*bool, bool)`

GetEnableContentDownloadingOk returns a tuple with the EnableContentDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDownloading

`func (o *ModelUserPolicy) SetEnableContentDownloading(v bool)`

SetEnableContentDownloading sets EnableContentDownloading field to given value.

### HasEnableContentDownloading

`func (o *ModelUserPolicy) HasEnableContentDownloading() bool`

HasEnableContentDownloading returns a boolean if a field has been set.

### GetEnableSubtitleDownloading

`func (o *ModelUserPolicy) GetEnableSubtitleDownloading() bool`

GetEnableSubtitleDownloading returns the EnableSubtitleDownloading field if non-nil, zero value otherwise.

### GetEnableSubtitleDownloadingOk

`func (o *ModelUserPolicy) GetEnableSubtitleDownloadingOk() (*bool, bool)`

GetEnableSubtitleDownloadingOk returns a tuple with the EnableSubtitleDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleDownloading

`func (o *ModelUserPolicy) SetEnableSubtitleDownloading(v bool)`

SetEnableSubtitleDownloading sets EnableSubtitleDownloading field to given value.

### HasEnableSubtitleDownloading

`func (o *ModelUserPolicy) HasEnableSubtitleDownloading() bool`

HasEnableSubtitleDownloading returns a boolean if a field has been set.

### GetEnableSubtitleManagement

`func (o *ModelUserPolicy) GetEnableSubtitleManagement() bool`

GetEnableSubtitleManagement returns the EnableSubtitleManagement field if non-nil, zero value otherwise.

### GetEnableSubtitleManagementOk

`func (o *ModelUserPolicy) GetEnableSubtitleManagementOk() (*bool, bool)`

GetEnableSubtitleManagementOk returns a tuple with the EnableSubtitleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleManagement

`func (o *ModelUserPolicy) SetEnableSubtitleManagement(v bool)`

SetEnableSubtitleManagement sets EnableSubtitleManagement field to given value.

### HasEnableSubtitleManagement

`func (o *ModelUserPolicy) HasEnableSubtitleManagement() bool`

HasEnableSubtitleManagement returns a boolean if a field has been set.

### GetEnableSyncTranscoding

`func (o *ModelUserPolicy) GetEnableSyncTranscoding() bool`

GetEnableSyncTranscoding returns the EnableSyncTranscoding field if non-nil, zero value otherwise.

### GetEnableSyncTranscodingOk

`func (o *ModelUserPolicy) GetEnableSyncTranscodingOk() (*bool, bool)`

GetEnableSyncTranscodingOk returns a tuple with the EnableSyncTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyncTranscoding

`func (o *ModelUserPolicy) SetEnableSyncTranscoding(v bool)`

SetEnableSyncTranscoding sets EnableSyncTranscoding field to given value.

### HasEnableSyncTranscoding

`func (o *ModelUserPolicy) HasEnableSyncTranscoding() bool`

HasEnableSyncTranscoding returns a boolean if a field has been set.

### GetEnableMediaConversion

`func (o *ModelUserPolicy) GetEnableMediaConversion() bool`

GetEnableMediaConversion returns the EnableMediaConversion field if non-nil, zero value otherwise.

### GetEnableMediaConversionOk

`func (o *ModelUserPolicy) GetEnableMediaConversionOk() (*bool, bool)`

GetEnableMediaConversionOk returns a tuple with the EnableMediaConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaConversion

`func (o *ModelUserPolicy) SetEnableMediaConversion(v bool)`

SetEnableMediaConversion sets EnableMediaConversion field to given value.

### HasEnableMediaConversion

`func (o *ModelUserPolicy) HasEnableMediaConversion() bool`

HasEnableMediaConversion returns a boolean if a field has been set.

### GetEnabledChannels

`func (o *ModelUserPolicy) GetEnabledChannels() []string`

GetEnabledChannels returns the EnabledChannels field if non-nil, zero value otherwise.

### GetEnabledChannelsOk

`func (o *ModelUserPolicy) GetEnabledChannelsOk() (*[]string, bool)`

GetEnabledChannelsOk returns a tuple with the EnabledChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledChannels

`func (o *ModelUserPolicy) SetEnabledChannels(v []string)`

SetEnabledChannels sets EnabledChannels field to given value.

### HasEnabledChannels

`func (o *ModelUserPolicy) HasEnabledChannels() bool`

HasEnabledChannels returns a boolean if a field has been set.

### GetEnableAllChannels

`func (o *ModelUserPolicy) GetEnableAllChannels() bool`

GetEnableAllChannels returns the EnableAllChannels field if non-nil, zero value otherwise.

### GetEnableAllChannelsOk

`func (o *ModelUserPolicy) GetEnableAllChannelsOk() (*bool, bool)`

GetEnableAllChannelsOk returns a tuple with the EnableAllChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllChannels

`func (o *ModelUserPolicy) SetEnableAllChannels(v bool)`

SetEnableAllChannels sets EnableAllChannels field to given value.

### HasEnableAllChannels

`func (o *ModelUserPolicy) HasEnableAllChannels() bool`

HasEnableAllChannels returns a boolean if a field has been set.

### GetEnabledFolders

`func (o *ModelUserPolicy) GetEnabledFolders() []string`

GetEnabledFolders returns the EnabledFolders field if non-nil, zero value otherwise.

### GetEnabledFoldersOk

`func (o *ModelUserPolicy) GetEnabledFoldersOk() (*[]string, bool)`

GetEnabledFoldersOk returns a tuple with the EnabledFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFolders

`func (o *ModelUserPolicy) SetEnabledFolders(v []string)`

SetEnabledFolders sets EnabledFolders field to given value.

### HasEnabledFolders

`func (o *ModelUserPolicy) HasEnabledFolders() bool`

HasEnabledFolders returns a boolean if a field has been set.

### GetEnableAllFolders

`func (o *ModelUserPolicy) GetEnableAllFolders() bool`

GetEnableAllFolders returns the EnableAllFolders field if non-nil, zero value otherwise.

### GetEnableAllFoldersOk

`func (o *ModelUserPolicy) GetEnableAllFoldersOk() (*bool, bool)`

GetEnableAllFoldersOk returns a tuple with the EnableAllFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllFolders

`func (o *ModelUserPolicy) SetEnableAllFolders(v bool)`

SetEnableAllFolders sets EnableAllFolders field to given value.

### HasEnableAllFolders

`func (o *ModelUserPolicy) HasEnableAllFolders() bool`

HasEnableAllFolders returns a boolean if a field has been set.

### GetInvalidLoginAttemptCount

`func (o *ModelUserPolicy) GetInvalidLoginAttemptCount() int32`

GetInvalidLoginAttemptCount returns the InvalidLoginAttemptCount field if non-nil, zero value otherwise.

### GetInvalidLoginAttemptCountOk

`func (o *ModelUserPolicy) GetInvalidLoginAttemptCountOk() (*int32, bool)`

GetInvalidLoginAttemptCountOk returns a tuple with the InvalidLoginAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLoginAttemptCount

`func (o *ModelUserPolicy) SetInvalidLoginAttemptCount(v int32)`

SetInvalidLoginAttemptCount sets InvalidLoginAttemptCount field to given value.

### HasInvalidLoginAttemptCount

`func (o *ModelUserPolicy) HasInvalidLoginAttemptCount() bool`

HasInvalidLoginAttemptCount returns a boolean if a field has been set.

### GetEnablePublicSharing

`func (o *ModelUserPolicy) GetEnablePublicSharing() bool`

GetEnablePublicSharing returns the EnablePublicSharing field if non-nil, zero value otherwise.

### GetEnablePublicSharingOk

`func (o *ModelUserPolicy) GetEnablePublicSharingOk() (*bool, bool)`

GetEnablePublicSharingOk returns a tuple with the EnablePublicSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicSharing

`func (o *ModelUserPolicy) SetEnablePublicSharing(v bool)`

SetEnablePublicSharing sets EnablePublicSharing field to given value.

### HasEnablePublicSharing

`func (o *ModelUserPolicy) HasEnablePublicSharing() bool`

HasEnablePublicSharing returns a boolean if a field has been set.

### GetBlockedMediaFolders

`func (o *ModelUserPolicy) GetBlockedMediaFolders() []string`

GetBlockedMediaFolders returns the BlockedMediaFolders field if non-nil, zero value otherwise.

### GetBlockedMediaFoldersOk

`func (o *ModelUserPolicy) GetBlockedMediaFoldersOk() (*[]string, bool)`

GetBlockedMediaFoldersOk returns a tuple with the BlockedMediaFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMediaFolders

`func (o *ModelUserPolicy) SetBlockedMediaFolders(v []string)`

SetBlockedMediaFolders sets BlockedMediaFolders field to given value.

### HasBlockedMediaFolders

`func (o *ModelUserPolicy) HasBlockedMediaFolders() bool`

HasBlockedMediaFolders returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *ModelUserPolicy) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *ModelUserPolicy) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *ModelUserPolicy) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *ModelUserPolicy) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetAuthenticationProviderId

`func (o *ModelUserPolicy) GetAuthenticationProviderId() string`

GetAuthenticationProviderId returns the AuthenticationProviderId field if non-nil, zero value otherwise.

### GetAuthenticationProviderIdOk

`func (o *ModelUserPolicy) GetAuthenticationProviderIdOk() (*string, bool)`

GetAuthenticationProviderIdOk returns a tuple with the AuthenticationProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderId

`func (o *ModelUserPolicy) SetAuthenticationProviderId(v string)`

SetAuthenticationProviderId sets AuthenticationProviderId field to given value.

### HasAuthenticationProviderId

`func (o *ModelUserPolicy) HasAuthenticationProviderId() bool`

HasAuthenticationProviderId returns a boolean if a field has been set.

### GetExcludedSubFolders

`func (o *ModelUserPolicy) GetExcludedSubFolders() []string`

GetExcludedSubFolders returns the ExcludedSubFolders field if non-nil, zero value otherwise.

### GetExcludedSubFoldersOk

`func (o *ModelUserPolicy) GetExcludedSubFoldersOk() (*[]string, bool)`

GetExcludedSubFoldersOk returns a tuple with the ExcludedSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedSubFolders

`func (o *ModelUserPolicy) SetExcludedSubFolders(v []string)`

SetExcludedSubFolders sets ExcludedSubFolders field to given value.

### HasExcludedSubFolders

`func (o *ModelUserPolicy) HasExcludedSubFolders() bool`

HasExcludedSubFolders returns a boolean if a field has been set.

### GetSimultaneousStreamLimit

`func (o *ModelUserPolicy) GetSimultaneousStreamLimit() int32`

GetSimultaneousStreamLimit returns the SimultaneousStreamLimit field if non-nil, zero value otherwise.

### GetSimultaneousStreamLimitOk

`func (o *ModelUserPolicy) GetSimultaneousStreamLimitOk() (*int32, bool)`

GetSimultaneousStreamLimitOk returns a tuple with the SimultaneousStreamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousStreamLimit

`func (o *ModelUserPolicy) SetSimultaneousStreamLimit(v int32)`

SetSimultaneousStreamLimit sets SimultaneousStreamLimit field to given value.

### HasSimultaneousStreamLimit

`func (o *ModelUserPolicy) HasSimultaneousStreamLimit() bool`

HasSimultaneousStreamLimit returns a boolean if a field has been set.

### GetEnabledDevices

`func (o *ModelUserPolicy) GetEnabledDevices() []string`

GetEnabledDevices returns the EnabledDevices field if non-nil, zero value otherwise.

### GetEnabledDevicesOk

`func (o *ModelUserPolicy) GetEnabledDevicesOk() (*[]string, bool)`

GetEnabledDevicesOk returns a tuple with the EnabledDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledDevices

`func (o *ModelUserPolicy) SetEnabledDevices(v []string)`

SetEnabledDevices sets EnabledDevices field to given value.

### HasEnabledDevices

`func (o *ModelUserPolicy) HasEnabledDevices() bool`

HasEnabledDevices returns a boolean if a field has been set.

### GetEnableAllDevices

`func (o *ModelUserPolicy) GetEnableAllDevices() bool`

GetEnableAllDevices returns the EnableAllDevices field if non-nil, zero value otherwise.

### GetEnableAllDevicesOk

`func (o *ModelUserPolicy) GetEnableAllDevicesOk() (*bool, bool)`

GetEnableAllDevicesOk returns a tuple with the EnableAllDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllDevices

`func (o *ModelUserPolicy) SetEnableAllDevices(v bool)`

SetEnableAllDevices sets EnableAllDevices field to given value.

### HasEnableAllDevices

`func (o *ModelUserPolicy) HasEnableAllDevices() bool`

HasEnableAllDevices returns a boolean if a field has been set.

### GetAllowCameraUpload

`func (o *ModelUserPolicy) GetAllowCameraUpload() bool`

GetAllowCameraUpload returns the AllowCameraUpload field if non-nil, zero value otherwise.

### GetAllowCameraUploadOk

`func (o *ModelUserPolicy) GetAllowCameraUploadOk() (*bool, bool)`

GetAllowCameraUploadOk returns a tuple with the AllowCameraUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCameraUpload

`func (o *ModelUserPolicy) SetAllowCameraUpload(v bool)`

SetAllowCameraUpload sets AllowCameraUpload field to given value.

### HasAllowCameraUpload

`func (o *ModelUserPolicy) HasAllowCameraUpload() bool`

HasAllowCameraUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


