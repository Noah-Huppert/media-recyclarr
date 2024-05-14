# ModelBaseItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**OriginalTitle** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**Etag** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**ExtraType** | Pointer to **string** |  | [optional] 
**SortIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**SortParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**CanDelete** | Pointer to **NullableBool** |  | [optional] 
**CanDownload** | Pointer to **NullableBool** |  | [optional] 
**SupportsResume** | Pointer to **NullableBool** |  | [optional] 
**PresentationUniqueKey** | Pointer to **string** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**PreferredMetadataCountryCode** | Pointer to **string** |  | [optional] 
**SupportsSync** | Pointer to **NullableBool** |  | [optional] 
**SyncStatus** | Pointer to [**ModelModelSyncJobItemStatus**](ModelSyncJobItemStatus.md) |  | [optional] 
**CanManageAccess** | Pointer to **NullableBool** |  | [optional] 
**CanMakePrivate** | Pointer to **NullableBool** |  | [optional] 
**CanMakePublic** | Pointer to **NullableBool** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**SortName** | Pointer to **string** |  | [optional] 
**ForcedSortName** | Pointer to **string** |  | [optional] 
**Video3DFormat** | Pointer to [**ModelModelVideo3DFormat**](ModelVideo3DFormat.md) |  | [optional] 
**PremiereDate** | Pointer to **NullableTime** |  | [optional] 
**ExternalUrls** | Pointer to [**[]ModelModelExternalUrl**](ModelModelExternalUrl.md) |  | [optional] 
**MediaSources** | Pointer to [**[]ModelModelMediaSourceInfo**](ModelModelMediaSourceInfo.md) |  | [optional] 
**CriticRating** | Pointer to **NullableFloat32** |  | [optional] 
**GameSystemId** | Pointer to **NullableInt64** |  | [optional] 
**AsSeries** | Pointer to **NullableBool** |  | [optional] 
**GameSystem** | Pointer to **string** |  | [optional] 
**ProductionLocations** | Pointer to **[]string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**OfficialRating** | Pointer to **string** |  | [optional] 
**CustomRating** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**ChannelName** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**Taglines** | Pointer to **[]string** |  | [optional] 
**Genres** | Pointer to **[]string** |  | [optional] 
**CommunityRating** | Pointer to **NullableFloat32** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**ProductionYear** | Pointer to **NullableInt32** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**ChannelNumber** | Pointer to **string** |  | [optional] 
**IndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**IndexNumberEnd** | Pointer to **NullableInt32** |  | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**RemoteTrailers** | Pointer to [**[]ModelModelMediaUrl**](ModelModelMediaUrl.md) |  | [optional] 
**ProviderIds** | Pointer to **map[string]string** |  | [optional] 
**IsFolder** | Pointer to **NullableBool** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**People** | Pointer to [**[]ModelModelBaseItemPerson**](ModelModelBaseItemPerson.md) |  | [optional] 
**Studios** | Pointer to [**[]ModelModelNameLongIdPair**](ModelModelNameLongIdPair.md) |  | [optional] 
**GenreItems** | Pointer to [**[]ModelModelNameLongIdPair**](ModelModelNameLongIdPair.md) |  | [optional] 
**TagItems** | Pointer to [**[]ModelModelNameLongIdPair**](ModelModelNameLongIdPair.md) |  | [optional] 
**ParentLogoItemId** | Pointer to **string** |  | [optional] 
**ParentBackdropItemId** | Pointer to **string** |  | [optional] 
**ParentBackdropImageTags** | Pointer to **[]string** |  | [optional] 
**LocalTrailerCount** | Pointer to **NullableInt32** |  | [optional] 
**UserData** | Pointer to [**ModelModelUserItemDataDto**](ModelUserItemDataDto.md) |  | [optional] 
**RecursiveItemCount** | Pointer to **NullableInt32** |  | [optional] 
**ChildCount** | Pointer to **NullableInt32** |  | [optional] 
**SeriesName** | Pointer to **string** |  | [optional] 
**SeriesId** | Pointer to **string** |  | [optional] 
**SeasonId** | Pointer to **string** |  | [optional] 
**SpecialFeatureCount** | Pointer to **NullableInt32** |  | [optional] 
**DisplayPreferencesId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AirDays** | Pointer to [**[]ModelModelDayOfWeek**](ModelModelDayOfWeek.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** |  | [optional] 
**Artists** | Pointer to **[]string** |  | [optional] 
**ArtistItems** | Pointer to [**[]ModelModelNameIdPair**](ModelModelNameIdPair.md) |  | [optional] 
**Composers** | Pointer to [**[]ModelModelNameIdPair**](ModelModelNameIdPair.md) |  | [optional] 
**Album** | Pointer to **string** |  | [optional] 
**CollectionType** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **string** |  | [optional] 
**AlbumId** | Pointer to **string** |  | [optional] 
**AlbumPrimaryImageTag** | Pointer to **string** |  | [optional] 
**SeriesPrimaryImageTag** | Pointer to **string** |  | [optional] 
**AlbumArtist** | Pointer to **string** |  | [optional] 
**AlbumArtists** | Pointer to [**[]ModelModelNameIdPair**](ModelModelNameIdPair.md) |  | [optional] 
**SeasonName** | Pointer to **string** |  | [optional] 
**MediaStreams** | Pointer to [**[]ModelModelMediaStream**](ModelModelMediaStream.md) |  | [optional] 
**PartCount** | Pointer to **NullableInt32** |  | [optional] 
**ImageTags** | Pointer to **map[string]string** |  | [optional] 
**BackdropImageTags** | Pointer to **[]string** |  | [optional] 
**ParentLogoImageTag** | Pointer to **string** |  | [optional] 
**SeriesStudio** | Pointer to **string** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**ParentThumbItemId** | Pointer to **string** |  | [optional] 
**ParentThumbImageTag** | Pointer to **string** |  | [optional] 
**Chapters** | Pointer to [**[]ModelModelChapterInfo**](ModelModelChapterInfo.md) |  | [optional] 
**LocationType** | Pointer to [**ModelModelLocationType**](ModelLocationType.md) |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **NullableTime** |  | [optional] 
**LockedFields** | Pointer to [**[]ModelModelMetadataFields**](ModelModelMetadataFields.md) |  | [optional] 
**LockData** | Pointer to **NullableBool** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**CameraMake** | Pointer to **string** |  | [optional] 
**CameraModel** | Pointer to **string** |  | [optional] 
**Software** | Pointer to **string** |  | [optional] 
**ExposureTime** | Pointer to **NullableFloat64** |  | [optional] 
**FocalLength** | Pointer to **NullableFloat64** |  | [optional] 
**ImageOrientation** | Pointer to [**ModelModelDrawingImageOrientation**](ModelDrawingImageOrientation.md) |  | [optional] 
**Aperture** | Pointer to **NullableFloat64** |  | [optional] 
**ShutterSpeed** | Pointer to **NullableFloat64** |  | [optional] 
**Latitude** | Pointer to **NullableFloat64** |  | [optional] 
**Longitude** | Pointer to **NullableFloat64** |  | [optional] 
**Altitude** | Pointer to **NullableFloat64** |  | [optional] 
**IsoSpeedRating** | Pointer to **NullableInt32** |  | [optional] 
**SeriesTimerId** | Pointer to **string** |  | [optional] 
**ChannelPrimaryImageTag** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **NullableTime** |  | [optional] 
**CompletionPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**IsRepeat** | Pointer to **NullableBool** |  | [optional] 
**IsNew** | Pointer to **NullableBool** |  | [optional] 
**EpisodeTitle** | Pointer to **string** |  | [optional] 
**IsMovie** | Pointer to **NullableBool** |  | [optional] 
**IsSports** | Pointer to **NullableBool** |  | [optional] 
**IsSeries** | Pointer to **NullableBool** |  | [optional] 
**IsLive** | Pointer to **NullableBool** |  | [optional] 
**IsNews** | Pointer to **NullableBool** |  | [optional] 
**IsKids** | Pointer to **NullableBool** |  | [optional] 
**IsPremiere** | Pointer to **NullableBool** |  | [optional] 
**TimerType** | Pointer to [**ModelModelLiveTvTimerType**](ModelLiveTvTimerType.md) |  | [optional] 
**Disabled** | Pointer to **NullableBool** |  | [optional] 
**ManagementId** | Pointer to **string** |  | [optional] 
**TimerId** | Pointer to **string** |  | [optional] 
**CurrentProgram** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**MovieCount** | Pointer to **NullableInt32** |  | [optional] 
**SeriesCount** | Pointer to **NullableInt32** |  | [optional] 
**AlbumCount** | Pointer to **NullableInt32** |  | [optional] 
**SongCount** | Pointer to **NullableInt32** |  | [optional] 
**MusicVideoCount** | Pointer to **NullableInt32** |  | [optional] 
**Subviews** | Pointer to **[]string** |  | [optional] 
**ListingsProviderId** | Pointer to **string** |  | [optional] 
**ListingsChannelId** | Pointer to **string** |  | [optional] 
**ListingsPath** | Pointer to **string** |  | [optional] 
**ListingsId** | Pointer to **string** |  | [optional] 
**ListingsChannelName** | Pointer to **string** |  | [optional] 
**ListingsChannelNumber** | Pointer to **string** |  | [optional] 
**AffiliateCallSign** | Pointer to **string** |  | [optional] 

## Methods

### NewModelBaseItemDto

`func NewModelBaseItemDto() *ModelBaseItemDto`

NewModelBaseItemDto instantiates a new ModelBaseItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelBaseItemDtoWithDefaults

`func NewModelBaseItemDtoWithDefaults() *ModelBaseItemDto`

NewModelBaseItemDtoWithDefaults instantiates a new ModelBaseItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelBaseItemDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelBaseItemDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelBaseItemDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelBaseItemDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalTitle

`func (o *ModelBaseItemDto) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *ModelBaseItemDto) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *ModelBaseItemDto) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *ModelBaseItemDto) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### GetServerId

`func (o *ModelBaseItemDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelBaseItemDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelBaseItemDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelBaseItemDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetId

`func (o *ModelBaseItemDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelBaseItemDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelBaseItemDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelBaseItemDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuid

`func (o *ModelBaseItemDto) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ModelBaseItemDto) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ModelBaseItemDto) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ModelBaseItemDto) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetEtag

`func (o *ModelBaseItemDto) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *ModelBaseItemDto) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *ModelBaseItemDto) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *ModelBaseItemDto) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetPrefix

`func (o *ModelBaseItemDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ModelBaseItemDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ModelBaseItemDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ModelBaseItemDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *ModelBaseItemDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelBaseItemDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelBaseItemDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelBaseItemDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ModelBaseItemDto) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelBaseItemDto) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelBaseItemDto) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelBaseItemDto) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ModelBaseItemDto) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ModelBaseItemDto) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetExtraType

`func (o *ModelBaseItemDto) GetExtraType() string`

GetExtraType returns the ExtraType field if non-nil, zero value otherwise.

### GetExtraTypeOk

`func (o *ModelBaseItemDto) GetExtraTypeOk() (*string, bool)`

GetExtraTypeOk returns a tuple with the ExtraType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraType

`func (o *ModelBaseItemDto) SetExtraType(v string)`

SetExtraType sets ExtraType field to given value.

### HasExtraType

`func (o *ModelBaseItemDto) HasExtraType() bool`

HasExtraType returns a boolean if a field has been set.

### GetSortIndexNumber

`func (o *ModelBaseItemDto) GetSortIndexNumber() int32`

GetSortIndexNumber returns the SortIndexNumber field if non-nil, zero value otherwise.

### GetSortIndexNumberOk

`func (o *ModelBaseItemDto) GetSortIndexNumberOk() (*int32, bool)`

GetSortIndexNumberOk returns a tuple with the SortIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortIndexNumber

`func (o *ModelBaseItemDto) SetSortIndexNumber(v int32)`

SetSortIndexNumber sets SortIndexNumber field to given value.

### HasSortIndexNumber

`func (o *ModelBaseItemDto) HasSortIndexNumber() bool`

HasSortIndexNumber returns a boolean if a field has been set.

### SetSortIndexNumberNil

`func (o *ModelBaseItemDto) SetSortIndexNumberNil(b bool)`

 SetSortIndexNumberNil sets the value for SortIndexNumber to be an explicit nil

### UnsetSortIndexNumber
`func (o *ModelBaseItemDto) UnsetSortIndexNumber()`

UnsetSortIndexNumber ensures that no value is present for SortIndexNumber, not even an explicit nil
### GetSortParentIndexNumber

`func (o *ModelBaseItemDto) GetSortParentIndexNumber() int32`

GetSortParentIndexNumber returns the SortParentIndexNumber field if non-nil, zero value otherwise.

### GetSortParentIndexNumberOk

`func (o *ModelBaseItemDto) GetSortParentIndexNumberOk() (*int32, bool)`

GetSortParentIndexNumberOk returns a tuple with the SortParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortParentIndexNumber

`func (o *ModelBaseItemDto) SetSortParentIndexNumber(v int32)`

SetSortParentIndexNumber sets SortParentIndexNumber field to given value.

### HasSortParentIndexNumber

`func (o *ModelBaseItemDto) HasSortParentIndexNumber() bool`

HasSortParentIndexNumber returns a boolean if a field has been set.

### SetSortParentIndexNumberNil

`func (o *ModelBaseItemDto) SetSortParentIndexNumberNil(b bool)`

 SetSortParentIndexNumberNil sets the value for SortParentIndexNumber to be an explicit nil

### UnsetSortParentIndexNumber
`func (o *ModelBaseItemDto) UnsetSortParentIndexNumber()`

UnsetSortParentIndexNumber ensures that no value is present for SortParentIndexNumber, not even an explicit nil
### GetCanDelete

`func (o *ModelBaseItemDto) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *ModelBaseItemDto) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *ModelBaseItemDto) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *ModelBaseItemDto) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### SetCanDeleteNil

`func (o *ModelBaseItemDto) SetCanDeleteNil(b bool)`

 SetCanDeleteNil sets the value for CanDelete to be an explicit nil

### UnsetCanDelete
`func (o *ModelBaseItemDto) UnsetCanDelete()`

UnsetCanDelete ensures that no value is present for CanDelete, not even an explicit nil
### GetCanDownload

`func (o *ModelBaseItemDto) GetCanDownload() bool`

GetCanDownload returns the CanDownload field if non-nil, zero value otherwise.

### GetCanDownloadOk

`func (o *ModelBaseItemDto) GetCanDownloadOk() (*bool, bool)`

GetCanDownloadOk returns a tuple with the CanDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDownload

`func (o *ModelBaseItemDto) SetCanDownload(v bool)`

SetCanDownload sets CanDownload field to given value.

### HasCanDownload

`func (o *ModelBaseItemDto) HasCanDownload() bool`

HasCanDownload returns a boolean if a field has been set.

### SetCanDownloadNil

`func (o *ModelBaseItemDto) SetCanDownloadNil(b bool)`

 SetCanDownloadNil sets the value for CanDownload to be an explicit nil

### UnsetCanDownload
`func (o *ModelBaseItemDto) UnsetCanDownload()`

UnsetCanDownload ensures that no value is present for CanDownload, not even an explicit nil
### GetSupportsResume

`func (o *ModelBaseItemDto) GetSupportsResume() bool`

GetSupportsResume returns the SupportsResume field if non-nil, zero value otherwise.

### GetSupportsResumeOk

`func (o *ModelBaseItemDto) GetSupportsResumeOk() (*bool, bool)`

GetSupportsResumeOk returns a tuple with the SupportsResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsResume

`func (o *ModelBaseItemDto) SetSupportsResume(v bool)`

SetSupportsResume sets SupportsResume field to given value.

### HasSupportsResume

`func (o *ModelBaseItemDto) HasSupportsResume() bool`

HasSupportsResume returns a boolean if a field has been set.

### SetSupportsResumeNil

`func (o *ModelBaseItemDto) SetSupportsResumeNil(b bool)`

 SetSupportsResumeNil sets the value for SupportsResume to be an explicit nil

### UnsetSupportsResume
`func (o *ModelBaseItemDto) UnsetSupportsResume()`

UnsetSupportsResume ensures that no value is present for SupportsResume, not even an explicit nil
### GetPresentationUniqueKey

`func (o *ModelBaseItemDto) GetPresentationUniqueKey() string`

GetPresentationUniqueKey returns the PresentationUniqueKey field if non-nil, zero value otherwise.

### GetPresentationUniqueKeyOk

`func (o *ModelBaseItemDto) GetPresentationUniqueKeyOk() (*string, bool)`

GetPresentationUniqueKeyOk returns a tuple with the PresentationUniqueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentationUniqueKey

`func (o *ModelBaseItemDto) SetPresentationUniqueKey(v string)`

SetPresentationUniqueKey sets PresentationUniqueKey field to given value.

### HasPresentationUniqueKey

`func (o *ModelBaseItemDto) HasPresentationUniqueKey() bool`

HasPresentationUniqueKey returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ModelBaseItemDto) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ModelBaseItemDto) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ModelBaseItemDto) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ModelBaseItemDto) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetPreferredMetadataCountryCode

`func (o *ModelBaseItemDto) GetPreferredMetadataCountryCode() string`

GetPreferredMetadataCountryCode returns the PreferredMetadataCountryCode field if non-nil, zero value otherwise.

### GetPreferredMetadataCountryCodeOk

`func (o *ModelBaseItemDto) GetPreferredMetadataCountryCodeOk() (*string, bool)`

GetPreferredMetadataCountryCodeOk returns a tuple with the PreferredMetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataCountryCode

`func (o *ModelBaseItemDto) SetPreferredMetadataCountryCode(v string)`

SetPreferredMetadataCountryCode sets PreferredMetadataCountryCode field to given value.

### HasPreferredMetadataCountryCode

`func (o *ModelBaseItemDto) HasPreferredMetadataCountryCode() bool`

HasPreferredMetadataCountryCode returns a boolean if a field has been set.

### GetSupportsSync

`func (o *ModelBaseItemDto) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *ModelBaseItemDto) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *ModelBaseItemDto) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *ModelBaseItemDto) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### SetSupportsSyncNil

`func (o *ModelBaseItemDto) SetSupportsSyncNil(b bool)`

 SetSupportsSyncNil sets the value for SupportsSync to be an explicit nil

### UnsetSupportsSync
`func (o *ModelBaseItemDto) UnsetSupportsSync()`

UnsetSupportsSync ensures that no value is present for SupportsSync, not even an explicit nil
### GetSyncStatus

`func (o *ModelBaseItemDto) GetSyncStatus() ModelModelSyncJobItemStatus`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *ModelBaseItemDto) GetSyncStatusOk() (*ModelModelSyncJobItemStatus, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *ModelBaseItemDto) SetSyncStatus(v ModelModelSyncJobItemStatus)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *ModelBaseItemDto) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetCanManageAccess

`func (o *ModelBaseItemDto) GetCanManageAccess() bool`

GetCanManageAccess returns the CanManageAccess field if non-nil, zero value otherwise.

### GetCanManageAccessOk

`func (o *ModelBaseItemDto) GetCanManageAccessOk() (*bool, bool)`

GetCanManageAccessOk returns a tuple with the CanManageAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageAccess

`func (o *ModelBaseItemDto) SetCanManageAccess(v bool)`

SetCanManageAccess sets CanManageAccess field to given value.

### HasCanManageAccess

`func (o *ModelBaseItemDto) HasCanManageAccess() bool`

HasCanManageAccess returns a boolean if a field has been set.

### SetCanManageAccessNil

`func (o *ModelBaseItemDto) SetCanManageAccessNil(b bool)`

 SetCanManageAccessNil sets the value for CanManageAccess to be an explicit nil

### UnsetCanManageAccess
`func (o *ModelBaseItemDto) UnsetCanManageAccess()`

UnsetCanManageAccess ensures that no value is present for CanManageAccess, not even an explicit nil
### GetCanMakePrivate

`func (o *ModelBaseItemDto) GetCanMakePrivate() bool`

GetCanMakePrivate returns the CanMakePrivate field if non-nil, zero value otherwise.

### GetCanMakePrivateOk

`func (o *ModelBaseItemDto) GetCanMakePrivateOk() (*bool, bool)`

GetCanMakePrivateOk returns a tuple with the CanMakePrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMakePrivate

`func (o *ModelBaseItemDto) SetCanMakePrivate(v bool)`

SetCanMakePrivate sets CanMakePrivate field to given value.

### HasCanMakePrivate

`func (o *ModelBaseItemDto) HasCanMakePrivate() bool`

HasCanMakePrivate returns a boolean if a field has been set.

### SetCanMakePrivateNil

`func (o *ModelBaseItemDto) SetCanMakePrivateNil(b bool)`

 SetCanMakePrivateNil sets the value for CanMakePrivate to be an explicit nil

### UnsetCanMakePrivate
`func (o *ModelBaseItemDto) UnsetCanMakePrivate()`

UnsetCanMakePrivate ensures that no value is present for CanMakePrivate, not even an explicit nil
### GetCanMakePublic

`func (o *ModelBaseItemDto) GetCanMakePublic() bool`

GetCanMakePublic returns the CanMakePublic field if non-nil, zero value otherwise.

### GetCanMakePublicOk

`func (o *ModelBaseItemDto) GetCanMakePublicOk() (*bool, bool)`

GetCanMakePublicOk returns a tuple with the CanMakePublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMakePublic

`func (o *ModelBaseItemDto) SetCanMakePublic(v bool)`

SetCanMakePublic sets CanMakePublic field to given value.

### HasCanMakePublic

`func (o *ModelBaseItemDto) HasCanMakePublic() bool`

HasCanMakePublic returns a boolean if a field has been set.

### SetCanMakePublicNil

`func (o *ModelBaseItemDto) SetCanMakePublicNil(b bool)`

 SetCanMakePublicNil sets the value for CanMakePublic to be an explicit nil

### UnsetCanMakePublic
`func (o *ModelBaseItemDto) UnsetCanMakePublic()`

UnsetCanMakePublic ensures that no value is present for CanMakePublic, not even an explicit nil
### GetContainer

`func (o *ModelBaseItemDto) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelBaseItemDto) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelBaseItemDto) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelBaseItemDto) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetSortName

`func (o *ModelBaseItemDto) GetSortName() string`

GetSortName returns the SortName field if non-nil, zero value otherwise.

### GetSortNameOk

`func (o *ModelBaseItemDto) GetSortNameOk() (*string, bool)`

GetSortNameOk returns a tuple with the SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortName

`func (o *ModelBaseItemDto) SetSortName(v string)`

SetSortName sets SortName field to given value.

### HasSortName

`func (o *ModelBaseItemDto) HasSortName() bool`

HasSortName returns a boolean if a field has been set.

### GetForcedSortName

`func (o *ModelBaseItemDto) GetForcedSortName() string`

GetForcedSortName returns the ForcedSortName field if non-nil, zero value otherwise.

### GetForcedSortNameOk

`func (o *ModelBaseItemDto) GetForcedSortNameOk() (*string, bool)`

GetForcedSortNameOk returns a tuple with the ForcedSortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForcedSortName

`func (o *ModelBaseItemDto) SetForcedSortName(v string)`

SetForcedSortName sets ForcedSortName field to given value.

### HasForcedSortName

`func (o *ModelBaseItemDto) HasForcedSortName() bool`

HasForcedSortName returns a boolean if a field has been set.

### GetVideo3DFormat

`func (o *ModelBaseItemDto) GetVideo3DFormat() ModelModelVideo3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *ModelBaseItemDto) GetVideo3DFormatOk() (*ModelModelVideo3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *ModelBaseItemDto) SetVideo3DFormat(v ModelModelVideo3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *ModelBaseItemDto) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### GetPremiereDate

`func (o *ModelBaseItemDto) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *ModelBaseItemDto) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *ModelBaseItemDto) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *ModelBaseItemDto) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *ModelBaseItemDto) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *ModelBaseItemDto) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetExternalUrls

`func (o *ModelBaseItemDto) GetExternalUrls() []ModelModelExternalUrl`

GetExternalUrls returns the ExternalUrls field if non-nil, zero value otherwise.

### GetExternalUrlsOk

`func (o *ModelBaseItemDto) GetExternalUrlsOk() (*[]ModelModelExternalUrl, bool)`

GetExternalUrlsOk returns a tuple with the ExternalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUrls

`func (o *ModelBaseItemDto) SetExternalUrls(v []ModelModelExternalUrl)`

SetExternalUrls sets ExternalUrls field to given value.

### HasExternalUrls

`func (o *ModelBaseItemDto) HasExternalUrls() bool`

HasExternalUrls returns a boolean if a field has been set.

### GetMediaSources

`func (o *ModelBaseItemDto) GetMediaSources() []ModelModelMediaSourceInfo`

GetMediaSources returns the MediaSources field if non-nil, zero value otherwise.

### GetMediaSourcesOk

`func (o *ModelBaseItemDto) GetMediaSourcesOk() (*[]ModelModelMediaSourceInfo, bool)`

GetMediaSourcesOk returns a tuple with the MediaSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSources

`func (o *ModelBaseItemDto) SetMediaSources(v []ModelModelMediaSourceInfo)`

SetMediaSources sets MediaSources field to given value.

### HasMediaSources

`func (o *ModelBaseItemDto) HasMediaSources() bool`

HasMediaSources returns a boolean if a field has been set.

### GetCriticRating

`func (o *ModelBaseItemDto) GetCriticRating() float32`

GetCriticRating returns the CriticRating field if non-nil, zero value otherwise.

### GetCriticRatingOk

`func (o *ModelBaseItemDto) GetCriticRatingOk() (*float32, bool)`

GetCriticRatingOk returns a tuple with the CriticRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticRating

`func (o *ModelBaseItemDto) SetCriticRating(v float32)`

SetCriticRating sets CriticRating field to given value.

### HasCriticRating

`func (o *ModelBaseItemDto) HasCriticRating() bool`

HasCriticRating returns a boolean if a field has been set.

### SetCriticRatingNil

`func (o *ModelBaseItemDto) SetCriticRatingNil(b bool)`

 SetCriticRatingNil sets the value for CriticRating to be an explicit nil

### UnsetCriticRating
`func (o *ModelBaseItemDto) UnsetCriticRating()`

UnsetCriticRating ensures that no value is present for CriticRating, not even an explicit nil
### GetGameSystemId

`func (o *ModelBaseItemDto) GetGameSystemId() int64`

GetGameSystemId returns the GameSystemId field if non-nil, zero value otherwise.

### GetGameSystemIdOk

`func (o *ModelBaseItemDto) GetGameSystemIdOk() (*int64, bool)`

GetGameSystemIdOk returns a tuple with the GameSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameSystemId

`func (o *ModelBaseItemDto) SetGameSystemId(v int64)`

SetGameSystemId sets GameSystemId field to given value.

### HasGameSystemId

`func (o *ModelBaseItemDto) HasGameSystemId() bool`

HasGameSystemId returns a boolean if a field has been set.

### SetGameSystemIdNil

`func (o *ModelBaseItemDto) SetGameSystemIdNil(b bool)`

 SetGameSystemIdNil sets the value for GameSystemId to be an explicit nil

### UnsetGameSystemId
`func (o *ModelBaseItemDto) UnsetGameSystemId()`

UnsetGameSystemId ensures that no value is present for GameSystemId, not even an explicit nil
### GetAsSeries

`func (o *ModelBaseItemDto) GetAsSeries() bool`

GetAsSeries returns the AsSeries field if non-nil, zero value otherwise.

### GetAsSeriesOk

`func (o *ModelBaseItemDto) GetAsSeriesOk() (*bool, bool)`

GetAsSeriesOk returns a tuple with the AsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSeries

`func (o *ModelBaseItemDto) SetAsSeries(v bool)`

SetAsSeries sets AsSeries field to given value.

### HasAsSeries

`func (o *ModelBaseItemDto) HasAsSeries() bool`

HasAsSeries returns a boolean if a field has been set.

### SetAsSeriesNil

`func (o *ModelBaseItemDto) SetAsSeriesNil(b bool)`

 SetAsSeriesNil sets the value for AsSeries to be an explicit nil

### UnsetAsSeries
`func (o *ModelBaseItemDto) UnsetAsSeries()`

UnsetAsSeries ensures that no value is present for AsSeries, not even an explicit nil
### GetGameSystem

`func (o *ModelBaseItemDto) GetGameSystem() string`

GetGameSystem returns the GameSystem field if non-nil, zero value otherwise.

### GetGameSystemOk

`func (o *ModelBaseItemDto) GetGameSystemOk() (*string, bool)`

GetGameSystemOk returns a tuple with the GameSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameSystem

`func (o *ModelBaseItemDto) SetGameSystem(v string)`

SetGameSystem sets GameSystem field to given value.

### HasGameSystem

`func (o *ModelBaseItemDto) HasGameSystem() bool`

HasGameSystem returns a boolean if a field has been set.

### GetProductionLocations

`func (o *ModelBaseItemDto) GetProductionLocations() []string`

GetProductionLocations returns the ProductionLocations field if non-nil, zero value otherwise.

### GetProductionLocationsOk

`func (o *ModelBaseItemDto) GetProductionLocationsOk() (*[]string, bool)`

GetProductionLocationsOk returns a tuple with the ProductionLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionLocations

`func (o *ModelBaseItemDto) SetProductionLocations(v []string)`

SetProductionLocations sets ProductionLocations field to given value.

### HasProductionLocations

`func (o *ModelBaseItemDto) HasProductionLocations() bool`

HasProductionLocations returns a boolean if a field has been set.

### GetPath

`func (o *ModelBaseItemDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelBaseItemDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelBaseItemDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelBaseItemDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetOfficialRating

`func (o *ModelBaseItemDto) GetOfficialRating() string`

GetOfficialRating returns the OfficialRating field if non-nil, zero value otherwise.

### GetOfficialRatingOk

`func (o *ModelBaseItemDto) GetOfficialRatingOk() (*string, bool)`

GetOfficialRatingOk returns a tuple with the OfficialRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialRating

`func (o *ModelBaseItemDto) SetOfficialRating(v string)`

SetOfficialRating sets OfficialRating field to given value.

### HasOfficialRating

`func (o *ModelBaseItemDto) HasOfficialRating() bool`

HasOfficialRating returns a boolean if a field has been set.

### GetCustomRating

`func (o *ModelBaseItemDto) GetCustomRating() string`

GetCustomRating returns the CustomRating field if non-nil, zero value otherwise.

### GetCustomRatingOk

`func (o *ModelBaseItemDto) GetCustomRatingOk() (*string, bool)`

GetCustomRatingOk returns a tuple with the CustomRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRating

`func (o *ModelBaseItemDto) SetCustomRating(v string)`

SetCustomRating sets CustomRating field to given value.

### HasCustomRating

`func (o *ModelBaseItemDto) HasCustomRating() bool`

HasCustomRating returns a boolean if a field has been set.

### GetChannelId

`func (o *ModelBaseItemDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ModelBaseItemDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ModelBaseItemDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ModelBaseItemDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelName

`func (o *ModelBaseItemDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *ModelBaseItemDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *ModelBaseItemDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *ModelBaseItemDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetOverview

`func (o *ModelBaseItemDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelBaseItemDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelBaseItemDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelBaseItemDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetTaglines

`func (o *ModelBaseItemDto) GetTaglines() []string`

GetTaglines returns the Taglines field if non-nil, zero value otherwise.

### GetTaglinesOk

`func (o *ModelBaseItemDto) GetTaglinesOk() (*[]string, bool)`

GetTaglinesOk returns a tuple with the Taglines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaglines

`func (o *ModelBaseItemDto) SetTaglines(v []string)`

SetTaglines sets Taglines field to given value.

### HasTaglines

`func (o *ModelBaseItemDto) HasTaglines() bool`

HasTaglines returns a boolean if a field has been set.

### GetGenres

`func (o *ModelBaseItemDto) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *ModelBaseItemDto) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *ModelBaseItemDto) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *ModelBaseItemDto) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCommunityRating

`func (o *ModelBaseItemDto) GetCommunityRating() float32`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *ModelBaseItemDto) GetCommunityRatingOk() (*float32, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *ModelBaseItemDto) SetCommunityRating(v float32)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *ModelBaseItemDto) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *ModelBaseItemDto) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *ModelBaseItemDto) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetRunTimeTicks

`func (o *ModelBaseItemDto) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *ModelBaseItemDto) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *ModelBaseItemDto) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *ModelBaseItemDto) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *ModelBaseItemDto) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *ModelBaseItemDto) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetSize

`func (o *ModelBaseItemDto) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ModelBaseItemDto) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ModelBaseItemDto) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ModelBaseItemDto) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ModelBaseItemDto) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ModelBaseItemDto) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetFileName

`func (o *ModelBaseItemDto) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *ModelBaseItemDto) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *ModelBaseItemDto) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *ModelBaseItemDto) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetBitrate

`func (o *ModelBaseItemDto) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *ModelBaseItemDto) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *ModelBaseItemDto) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *ModelBaseItemDto) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *ModelBaseItemDto) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *ModelBaseItemDto) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetProductionYear

`func (o *ModelBaseItemDto) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *ModelBaseItemDto) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *ModelBaseItemDto) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *ModelBaseItemDto) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *ModelBaseItemDto) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *ModelBaseItemDto) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetNumber

`func (o *ModelBaseItemDto) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ModelBaseItemDto) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ModelBaseItemDto) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *ModelBaseItemDto) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetChannelNumber

`func (o *ModelBaseItemDto) GetChannelNumber() string`

GetChannelNumber returns the ChannelNumber field if non-nil, zero value otherwise.

### GetChannelNumberOk

`func (o *ModelBaseItemDto) GetChannelNumberOk() (*string, bool)`

GetChannelNumberOk returns a tuple with the ChannelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelNumber

`func (o *ModelBaseItemDto) SetChannelNumber(v string)`

SetChannelNumber sets ChannelNumber field to given value.

### HasChannelNumber

`func (o *ModelBaseItemDto) HasChannelNumber() bool`

HasChannelNumber returns a boolean if a field has been set.

### GetIndexNumber

`func (o *ModelBaseItemDto) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *ModelBaseItemDto) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *ModelBaseItemDto) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *ModelBaseItemDto) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *ModelBaseItemDto) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *ModelBaseItemDto) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetIndexNumberEnd

`func (o *ModelBaseItemDto) GetIndexNumberEnd() int32`

GetIndexNumberEnd returns the IndexNumberEnd field if non-nil, zero value otherwise.

### GetIndexNumberEndOk

`func (o *ModelBaseItemDto) GetIndexNumberEndOk() (*int32, bool)`

GetIndexNumberEndOk returns a tuple with the IndexNumberEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumberEnd

`func (o *ModelBaseItemDto) SetIndexNumberEnd(v int32)`

SetIndexNumberEnd sets IndexNumberEnd field to given value.

### HasIndexNumberEnd

`func (o *ModelBaseItemDto) HasIndexNumberEnd() bool`

HasIndexNumberEnd returns a boolean if a field has been set.

### SetIndexNumberEndNil

`func (o *ModelBaseItemDto) SetIndexNumberEndNil(b bool)`

 SetIndexNumberEndNil sets the value for IndexNumberEnd to be an explicit nil

### UnsetIndexNumberEnd
`func (o *ModelBaseItemDto) UnsetIndexNumberEnd()`

UnsetIndexNumberEnd ensures that no value is present for IndexNumberEnd, not even an explicit nil
### GetParentIndexNumber

`func (o *ModelBaseItemDto) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *ModelBaseItemDto) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *ModelBaseItemDto) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *ModelBaseItemDto) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *ModelBaseItemDto) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *ModelBaseItemDto) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetRemoteTrailers

`func (o *ModelBaseItemDto) GetRemoteTrailers() []ModelModelMediaUrl`

GetRemoteTrailers returns the RemoteTrailers field if non-nil, zero value otherwise.

### GetRemoteTrailersOk

`func (o *ModelBaseItemDto) GetRemoteTrailersOk() (*[]ModelModelMediaUrl, bool)`

GetRemoteTrailersOk returns a tuple with the RemoteTrailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTrailers

`func (o *ModelBaseItemDto) SetRemoteTrailers(v []ModelModelMediaUrl)`

SetRemoteTrailers sets RemoteTrailers field to given value.

### HasRemoteTrailers

`func (o *ModelBaseItemDto) HasRemoteTrailers() bool`

HasRemoteTrailers returns a boolean if a field has been set.

### GetProviderIds

`func (o *ModelBaseItemDto) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *ModelBaseItemDto) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *ModelBaseItemDto) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *ModelBaseItemDto) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetIsFolder

`func (o *ModelBaseItemDto) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *ModelBaseItemDto) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *ModelBaseItemDto) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *ModelBaseItemDto) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *ModelBaseItemDto) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *ModelBaseItemDto) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil
### GetParentId

`func (o *ModelBaseItemDto) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModelBaseItemDto) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModelBaseItemDto) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModelBaseItemDto) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetType

`func (o *ModelBaseItemDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelBaseItemDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelBaseItemDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelBaseItemDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPeople

`func (o *ModelBaseItemDto) GetPeople() []ModelModelBaseItemPerson`

GetPeople returns the People field if non-nil, zero value otherwise.

### GetPeopleOk

`func (o *ModelBaseItemDto) GetPeopleOk() (*[]ModelModelBaseItemPerson, bool)`

GetPeopleOk returns a tuple with the People field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeople

`func (o *ModelBaseItemDto) SetPeople(v []ModelModelBaseItemPerson)`

SetPeople sets People field to given value.

### HasPeople

`func (o *ModelBaseItemDto) HasPeople() bool`

HasPeople returns a boolean if a field has been set.

### GetStudios

`func (o *ModelBaseItemDto) GetStudios() []ModelModelNameLongIdPair`

GetStudios returns the Studios field if non-nil, zero value otherwise.

### GetStudiosOk

`func (o *ModelBaseItemDto) GetStudiosOk() (*[]ModelModelNameLongIdPair, bool)`

GetStudiosOk returns a tuple with the Studios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudios

`func (o *ModelBaseItemDto) SetStudios(v []ModelModelNameLongIdPair)`

SetStudios sets Studios field to given value.

### HasStudios

`func (o *ModelBaseItemDto) HasStudios() bool`

HasStudios returns a boolean if a field has been set.

### GetGenreItems

`func (o *ModelBaseItemDto) GetGenreItems() []ModelModelNameLongIdPair`

GetGenreItems returns the GenreItems field if non-nil, zero value otherwise.

### GetGenreItemsOk

`func (o *ModelBaseItemDto) GetGenreItemsOk() (*[]ModelModelNameLongIdPair, bool)`

GetGenreItemsOk returns a tuple with the GenreItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreItems

`func (o *ModelBaseItemDto) SetGenreItems(v []ModelModelNameLongIdPair)`

SetGenreItems sets GenreItems field to given value.

### HasGenreItems

`func (o *ModelBaseItemDto) HasGenreItems() bool`

HasGenreItems returns a boolean if a field has been set.

### GetTagItems

`func (o *ModelBaseItemDto) GetTagItems() []ModelModelNameLongIdPair`

GetTagItems returns the TagItems field if non-nil, zero value otherwise.

### GetTagItemsOk

`func (o *ModelBaseItemDto) GetTagItemsOk() (*[]ModelModelNameLongIdPair, bool)`

GetTagItemsOk returns a tuple with the TagItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagItems

`func (o *ModelBaseItemDto) SetTagItems(v []ModelModelNameLongIdPair)`

SetTagItems sets TagItems field to given value.

### HasTagItems

`func (o *ModelBaseItemDto) HasTagItems() bool`

HasTagItems returns a boolean if a field has been set.

### GetParentLogoItemId

`func (o *ModelBaseItemDto) GetParentLogoItemId() string`

GetParentLogoItemId returns the ParentLogoItemId field if non-nil, zero value otherwise.

### GetParentLogoItemIdOk

`func (o *ModelBaseItemDto) GetParentLogoItemIdOk() (*string, bool)`

GetParentLogoItemIdOk returns a tuple with the ParentLogoItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogoItemId

`func (o *ModelBaseItemDto) SetParentLogoItemId(v string)`

SetParentLogoItemId sets ParentLogoItemId field to given value.

### HasParentLogoItemId

`func (o *ModelBaseItemDto) HasParentLogoItemId() bool`

HasParentLogoItemId returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *ModelBaseItemDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *ModelBaseItemDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *ModelBaseItemDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *ModelBaseItemDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### GetParentBackdropImageTags

`func (o *ModelBaseItemDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *ModelBaseItemDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *ModelBaseItemDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *ModelBaseItemDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### GetLocalTrailerCount

`func (o *ModelBaseItemDto) GetLocalTrailerCount() int32`

GetLocalTrailerCount returns the LocalTrailerCount field if non-nil, zero value otherwise.

### GetLocalTrailerCountOk

`func (o *ModelBaseItemDto) GetLocalTrailerCountOk() (*int32, bool)`

GetLocalTrailerCountOk returns a tuple with the LocalTrailerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTrailerCount

`func (o *ModelBaseItemDto) SetLocalTrailerCount(v int32)`

SetLocalTrailerCount sets LocalTrailerCount field to given value.

### HasLocalTrailerCount

`func (o *ModelBaseItemDto) HasLocalTrailerCount() bool`

HasLocalTrailerCount returns a boolean if a field has been set.

### SetLocalTrailerCountNil

`func (o *ModelBaseItemDto) SetLocalTrailerCountNil(b bool)`

 SetLocalTrailerCountNil sets the value for LocalTrailerCount to be an explicit nil

### UnsetLocalTrailerCount
`func (o *ModelBaseItemDto) UnsetLocalTrailerCount()`

UnsetLocalTrailerCount ensures that no value is present for LocalTrailerCount, not even an explicit nil
### GetUserData

`func (o *ModelBaseItemDto) GetUserData() ModelModelUserItemDataDto`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *ModelBaseItemDto) GetUserDataOk() (*ModelModelUserItemDataDto, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *ModelBaseItemDto) SetUserData(v ModelModelUserItemDataDto)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *ModelBaseItemDto) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetRecursiveItemCount

`func (o *ModelBaseItemDto) GetRecursiveItemCount() int32`

GetRecursiveItemCount returns the RecursiveItemCount field if non-nil, zero value otherwise.

### GetRecursiveItemCountOk

`func (o *ModelBaseItemDto) GetRecursiveItemCountOk() (*int32, bool)`

GetRecursiveItemCountOk returns a tuple with the RecursiveItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecursiveItemCount

`func (o *ModelBaseItemDto) SetRecursiveItemCount(v int32)`

SetRecursiveItemCount sets RecursiveItemCount field to given value.

### HasRecursiveItemCount

`func (o *ModelBaseItemDto) HasRecursiveItemCount() bool`

HasRecursiveItemCount returns a boolean if a field has been set.

### SetRecursiveItemCountNil

`func (o *ModelBaseItemDto) SetRecursiveItemCountNil(b bool)`

 SetRecursiveItemCountNil sets the value for RecursiveItemCount to be an explicit nil

### UnsetRecursiveItemCount
`func (o *ModelBaseItemDto) UnsetRecursiveItemCount()`

UnsetRecursiveItemCount ensures that no value is present for RecursiveItemCount, not even an explicit nil
### GetChildCount

`func (o *ModelBaseItemDto) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *ModelBaseItemDto) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *ModelBaseItemDto) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *ModelBaseItemDto) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCountNil

`func (o *ModelBaseItemDto) SetChildCountNil(b bool)`

 SetChildCountNil sets the value for ChildCount to be an explicit nil

### UnsetChildCount
`func (o *ModelBaseItemDto) UnsetChildCount()`

UnsetChildCount ensures that no value is present for ChildCount, not even an explicit nil
### GetSeriesName

`func (o *ModelBaseItemDto) GetSeriesName() string`

GetSeriesName returns the SeriesName field if non-nil, zero value otherwise.

### GetSeriesNameOk

`func (o *ModelBaseItemDto) GetSeriesNameOk() (*string, bool)`

GetSeriesNameOk returns a tuple with the SeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesName

`func (o *ModelBaseItemDto) SetSeriesName(v string)`

SetSeriesName sets SeriesName field to given value.

### HasSeriesName

`func (o *ModelBaseItemDto) HasSeriesName() bool`

HasSeriesName returns a boolean if a field has been set.

### GetSeriesId

`func (o *ModelBaseItemDto) GetSeriesId() string`

GetSeriesId returns the SeriesId field if non-nil, zero value otherwise.

### GetSeriesIdOk

`func (o *ModelBaseItemDto) GetSeriesIdOk() (*string, bool)`

GetSeriesIdOk returns a tuple with the SeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesId

`func (o *ModelBaseItemDto) SetSeriesId(v string)`

SetSeriesId sets SeriesId field to given value.

### HasSeriesId

`func (o *ModelBaseItemDto) HasSeriesId() bool`

HasSeriesId returns a boolean if a field has been set.

### GetSeasonId

`func (o *ModelBaseItemDto) GetSeasonId() string`

GetSeasonId returns the SeasonId field if non-nil, zero value otherwise.

### GetSeasonIdOk

`func (o *ModelBaseItemDto) GetSeasonIdOk() (*string, bool)`

GetSeasonIdOk returns a tuple with the SeasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonId

`func (o *ModelBaseItemDto) SetSeasonId(v string)`

SetSeasonId sets SeasonId field to given value.

### HasSeasonId

`func (o *ModelBaseItemDto) HasSeasonId() bool`

HasSeasonId returns a boolean if a field has been set.

### GetSpecialFeatureCount

`func (o *ModelBaseItemDto) GetSpecialFeatureCount() int32`

GetSpecialFeatureCount returns the SpecialFeatureCount field if non-nil, zero value otherwise.

### GetSpecialFeatureCountOk

`func (o *ModelBaseItemDto) GetSpecialFeatureCountOk() (*int32, bool)`

GetSpecialFeatureCountOk returns a tuple with the SpecialFeatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialFeatureCount

`func (o *ModelBaseItemDto) SetSpecialFeatureCount(v int32)`

SetSpecialFeatureCount sets SpecialFeatureCount field to given value.

### HasSpecialFeatureCount

`func (o *ModelBaseItemDto) HasSpecialFeatureCount() bool`

HasSpecialFeatureCount returns a boolean if a field has been set.

### SetSpecialFeatureCountNil

`func (o *ModelBaseItemDto) SetSpecialFeatureCountNil(b bool)`

 SetSpecialFeatureCountNil sets the value for SpecialFeatureCount to be an explicit nil

### UnsetSpecialFeatureCount
`func (o *ModelBaseItemDto) UnsetSpecialFeatureCount()`

UnsetSpecialFeatureCount ensures that no value is present for SpecialFeatureCount, not even an explicit nil
### GetDisplayPreferencesId

`func (o *ModelBaseItemDto) GetDisplayPreferencesId() string`

GetDisplayPreferencesId returns the DisplayPreferencesId field if non-nil, zero value otherwise.

### GetDisplayPreferencesIdOk

`func (o *ModelBaseItemDto) GetDisplayPreferencesIdOk() (*string, bool)`

GetDisplayPreferencesIdOk returns a tuple with the DisplayPreferencesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPreferencesId

`func (o *ModelBaseItemDto) SetDisplayPreferencesId(v string)`

SetDisplayPreferencesId sets DisplayPreferencesId field to given value.

### HasDisplayPreferencesId

`func (o *ModelBaseItemDto) HasDisplayPreferencesId() bool`

HasDisplayPreferencesId returns a boolean if a field has been set.

### GetStatus

`func (o *ModelBaseItemDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelBaseItemDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelBaseItemDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelBaseItemDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAirDays

`func (o *ModelBaseItemDto) GetAirDays() []ModelModelDayOfWeek`

GetAirDays returns the AirDays field if non-nil, zero value otherwise.

### GetAirDaysOk

`func (o *ModelBaseItemDto) GetAirDaysOk() (*[]ModelModelDayOfWeek, bool)`

GetAirDaysOk returns a tuple with the AirDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDays

`func (o *ModelBaseItemDto) SetAirDays(v []ModelModelDayOfWeek)`

SetAirDays sets AirDays field to given value.

### HasAirDays

`func (o *ModelBaseItemDto) HasAirDays() bool`

HasAirDays returns a boolean if a field has been set.

### GetTags

`func (o *ModelBaseItemDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelBaseItemDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelBaseItemDto) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelBaseItemDto) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetPrimaryImageAspectRatio

`func (o *ModelBaseItemDto) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *ModelBaseItemDto) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *ModelBaseItemDto) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *ModelBaseItemDto) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *ModelBaseItemDto) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *ModelBaseItemDto) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
### GetArtists

`func (o *ModelBaseItemDto) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *ModelBaseItemDto) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *ModelBaseItemDto) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *ModelBaseItemDto) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetArtistItems

`func (o *ModelBaseItemDto) GetArtistItems() []ModelModelNameIdPair`

GetArtistItems returns the ArtistItems field if non-nil, zero value otherwise.

### GetArtistItemsOk

`func (o *ModelBaseItemDto) GetArtistItemsOk() (*[]ModelModelNameIdPair, bool)`

GetArtistItemsOk returns a tuple with the ArtistItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistItems

`func (o *ModelBaseItemDto) SetArtistItems(v []ModelModelNameIdPair)`

SetArtistItems sets ArtistItems field to given value.

### HasArtistItems

`func (o *ModelBaseItemDto) HasArtistItems() bool`

HasArtistItems returns a boolean if a field has been set.

### GetComposers

`func (o *ModelBaseItemDto) GetComposers() []ModelModelNameIdPair`

GetComposers returns the Composers field if non-nil, zero value otherwise.

### GetComposersOk

`func (o *ModelBaseItemDto) GetComposersOk() (*[]ModelModelNameIdPair, bool)`

GetComposersOk returns a tuple with the Composers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposers

`func (o *ModelBaseItemDto) SetComposers(v []ModelModelNameIdPair)`

SetComposers sets Composers field to given value.

### HasComposers

`func (o *ModelBaseItemDto) HasComposers() bool`

HasComposers returns a boolean if a field has been set.

### GetAlbum

`func (o *ModelBaseItemDto) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *ModelBaseItemDto) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *ModelBaseItemDto) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *ModelBaseItemDto) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetCollectionType

`func (o *ModelBaseItemDto) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *ModelBaseItemDto) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *ModelBaseItemDto) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *ModelBaseItemDto) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ModelBaseItemDto) GetDisplayOrder() string`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ModelBaseItemDto) GetDisplayOrderOk() (*string, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ModelBaseItemDto) SetDisplayOrder(v string)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ModelBaseItemDto) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetAlbumId

`func (o *ModelBaseItemDto) GetAlbumId() string`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *ModelBaseItemDto) GetAlbumIdOk() (*string, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *ModelBaseItemDto) SetAlbumId(v string)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *ModelBaseItemDto) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### GetAlbumPrimaryImageTag

`func (o *ModelBaseItemDto) GetAlbumPrimaryImageTag() string`

GetAlbumPrimaryImageTag returns the AlbumPrimaryImageTag field if non-nil, zero value otherwise.

### GetAlbumPrimaryImageTagOk

`func (o *ModelBaseItemDto) GetAlbumPrimaryImageTagOk() (*string, bool)`

GetAlbumPrimaryImageTagOk returns a tuple with the AlbumPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumPrimaryImageTag

`func (o *ModelBaseItemDto) SetAlbumPrimaryImageTag(v string)`

SetAlbumPrimaryImageTag sets AlbumPrimaryImageTag field to given value.

### HasAlbumPrimaryImageTag

`func (o *ModelBaseItemDto) HasAlbumPrimaryImageTag() bool`

HasAlbumPrimaryImageTag returns a boolean if a field has been set.

### GetSeriesPrimaryImageTag

`func (o *ModelBaseItemDto) GetSeriesPrimaryImageTag() string`

GetSeriesPrimaryImageTag returns the SeriesPrimaryImageTag field if non-nil, zero value otherwise.

### GetSeriesPrimaryImageTagOk

`func (o *ModelBaseItemDto) GetSeriesPrimaryImageTagOk() (*string, bool)`

GetSeriesPrimaryImageTagOk returns a tuple with the SeriesPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesPrimaryImageTag

`func (o *ModelBaseItemDto) SetSeriesPrimaryImageTag(v string)`

SetSeriesPrimaryImageTag sets SeriesPrimaryImageTag field to given value.

### HasSeriesPrimaryImageTag

`func (o *ModelBaseItemDto) HasSeriesPrimaryImageTag() bool`

HasSeriesPrimaryImageTag returns a boolean if a field has been set.

### GetAlbumArtist

`func (o *ModelBaseItemDto) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *ModelBaseItemDto) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *ModelBaseItemDto) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *ModelBaseItemDto) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### GetAlbumArtists

`func (o *ModelBaseItemDto) GetAlbumArtists() []ModelModelNameIdPair`

GetAlbumArtists returns the AlbumArtists field if non-nil, zero value otherwise.

### GetAlbumArtistsOk

`func (o *ModelBaseItemDto) GetAlbumArtistsOk() (*[]ModelModelNameIdPair, bool)`

GetAlbumArtistsOk returns a tuple with the AlbumArtists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtists

`func (o *ModelBaseItemDto) SetAlbumArtists(v []ModelModelNameIdPair)`

SetAlbumArtists sets AlbumArtists field to given value.

### HasAlbumArtists

`func (o *ModelBaseItemDto) HasAlbumArtists() bool`

HasAlbumArtists returns a boolean if a field has been set.

### GetSeasonName

`func (o *ModelBaseItemDto) GetSeasonName() string`

GetSeasonName returns the SeasonName field if non-nil, zero value otherwise.

### GetSeasonNameOk

`func (o *ModelBaseItemDto) GetSeasonNameOk() (*string, bool)`

GetSeasonNameOk returns a tuple with the SeasonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonName

`func (o *ModelBaseItemDto) SetSeasonName(v string)`

SetSeasonName sets SeasonName field to given value.

### HasSeasonName

`func (o *ModelBaseItemDto) HasSeasonName() bool`

HasSeasonName returns a boolean if a field has been set.

### GetMediaStreams

`func (o *ModelBaseItemDto) GetMediaStreams() []ModelModelMediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *ModelBaseItemDto) GetMediaStreamsOk() (*[]ModelModelMediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *ModelBaseItemDto) SetMediaStreams(v []ModelModelMediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *ModelBaseItemDto) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### GetPartCount

`func (o *ModelBaseItemDto) GetPartCount() int32`

GetPartCount returns the PartCount field if non-nil, zero value otherwise.

### GetPartCountOk

`func (o *ModelBaseItemDto) GetPartCountOk() (*int32, bool)`

GetPartCountOk returns a tuple with the PartCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartCount

`func (o *ModelBaseItemDto) SetPartCount(v int32)`

SetPartCount sets PartCount field to given value.

### HasPartCount

`func (o *ModelBaseItemDto) HasPartCount() bool`

HasPartCount returns a boolean if a field has been set.

### SetPartCountNil

`func (o *ModelBaseItemDto) SetPartCountNil(b bool)`

 SetPartCountNil sets the value for PartCount to be an explicit nil

### UnsetPartCount
`func (o *ModelBaseItemDto) UnsetPartCount()`

UnsetPartCount ensures that no value is present for PartCount, not even an explicit nil
### GetImageTags

`func (o *ModelBaseItemDto) GetImageTags() map[string]string`

GetImageTags returns the ImageTags field if non-nil, zero value otherwise.

### GetImageTagsOk

`func (o *ModelBaseItemDto) GetImageTagsOk() (*map[string]string, bool)`

GetImageTagsOk returns a tuple with the ImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTags

`func (o *ModelBaseItemDto) SetImageTags(v map[string]string)`

SetImageTags sets ImageTags field to given value.

### HasImageTags

`func (o *ModelBaseItemDto) HasImageTags() bool`

HasImageTags returns a boolean if a field has been set.

### GetBackdropImageTags

`func (o *ModelBaseItemDto) GetBackdropImageTags() []string`

GetBackdropImageTags returns the BackdropImageTags field if non-nil, zero value otherwise.

### GetBackdropImageTagsOk

`func (o *ModelBaseItemDto) GetBackdropImageTagsOk() (*[]string, bool)`

GetBackdropImageTagsOk returns a tuple with the BackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageTags

`func (o *ModelBaseItemDto) SetBackdropImageTags(v []string)`

SetBackdropImageTags sets BackdropImageTags field to given value.

### HasBackdropImageTags

`func (o *ModelBaseItemDto) HasBackdropImageTags() bool`

HasBackdropImageTags returns a boolean if a field has been set.

### GetParentLogoImageTag

`func (o *ModelBaseItemDto) GetParentLogoImageTag() string`

GetParentLogoImageTag returns the ParentLogoImageTag field if non-nil, zero value otherwise.

### GetParentLogoImageTagOk

`func (o *ModelBaseItemDto) GetParentLogoImageTagOk() (*string, bool)`

GetParentLogoImageTagOk returns a tuple with the ParentLogoImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentLogoImageTag

`func (o *ModelBaseItemDto) SetParentLogoImageTag(v string)`

SetParentLogoImageTag sets ParentLogoImageTag field to given value.

### HasParentLogoImageTag

`func (o *ModelBaseItemDto) HasParentLogoImageTag() bool`

HasParentLogoImageTag returns a boolean if a field has been set.

### GetSeriesStudio

`func (o *ModelBaseItemDto) GetSeriesStudio() string`

GetSeriesStudio returns the SeriesStudio field if non-nil, zero value otherwise.

### GetSeriesStudioOk

`func (o *ModelBaseItemDto) GetSeriesStudioOk() (*string, bool)`

GetSeriesStudioOk returns a tuple with the SeriesStudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesStudio

`func (o *ModelBaseItemDto) SetSeriesStudio(v string)`

SetSeriesStudio sets SeriesStudio field to given value.

### HasSeriesStudio

`func (o *ModelBaseItemDto) HasSeriesStudio() bool`

HasSeriesStudio returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *ModelBaseItemDto) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *ModelBaseItemDto) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *ModelBaseItemDto) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *ModelBaseItemDto) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *ModelBaseItemDto) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *ModelBaseItemDto) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *ModelBaseItemDto) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *ModelBaseItemDto) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetParentThumbItemId

`func (o *ModelBaseItemDto) GetParentThumbItemId() string`

GetParentThumbItemId returns the ParentThumbItemId field if non-nil, zero value otherwise.

### GetParentThumbItemIdOk

`func (o *ModelBaseItemDto) GetParentThumbItemIdOk() (*string, bool)`

GetParentThumbItemIdOk returns a tuple with the ParentThumbItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbItemId

`func (o *ModelBaseItemDto) SetParentThumbItemId(v string)`

SetParentThumbItemId sets ParentThumbItemId field to given value.

### HasParentThumbItemId

`func (o *ModelBaseItemDto) HasParentThumbItemId() bool`

HasParentThumbItemId returns a boolean if a field has been set.

### GetParentThumbImageTag

`func (o *ModelBaseItemDto) GetParentThumbImageTag() string`

GetParentThumbImageTag returns the ParentThumbImageTag field if non-nil, zero value otherwise.

### GetParentThumbImageTagOk

`func (o *ModelBaseItemDto) GetParentThumbImageTagOk() (*string, bool)`

GetParentThumbImageTagOk returns a tuple with the ParentThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentThumbImageTag

`func (o *ModelBaseItemDto) SetParentThumbImageTag(v string)`

SetParentThumbImageTag sets ParentThumbImageTag field to given value.

### HasParentThumbImageTag

`func (o *ModelBaseItemDto) HasParentThumbImageTag() bool`

HasParentThumbImageTag returns a boolean if a field has been set.

### GetChapters

`func (o *ModelBaseItemDto) GetChapters() []ModelModelChapterInfo`

GetChapters returns the Chapters field if non-nil, zero value otherwise.

### GetChaptersOk

`func (o *ModelBaseItemDto) GetChaptersOk() (*[]ModelModelChapterInfo, bool)`

GetChaptersOk returns a tuple with the Chapters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChapters

`func (o *ModelBaseItemDto) SetChapters(v []ModelModelChapterInfo)`

SetChapters sets Chapters field to given value.

### HasChapters

`func (o *ModelBaseItemDto) HasChapters() bool`

HasChapters returns a boolean if a field has been set.

### GetLocationType

`func (o *ModelBaseItemDto) GetLocationType() ModelModelLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *ModelBaseItemDto) GetLocationTypeOk() (*ModelModelLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *ModelBaseItemDto) SetLocationType(v ModelModelLocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *ModelBaseItemDto) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetMediaType

`func (o *ModelBaseItemDto) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ModelBaseItemDto) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ModelBaseItemDto) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *ModelBaseItemDto) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetEndDate

`func (o *ModelBaseItemDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModelBaseItemDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModelBaseItemDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModelBaseItemDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *ModelBaseItemDto) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *ModelBaseItemDto) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetLockedFields

`func (o *ModelBaseItemDto) GetLockedFields() []ModelModelMetadataFields`

GetLockedFields returns the LockedFields field if non-nil, zero value otherwise.

### GetLockedFieldsOk

`func (o *ModelBaseItemDto) GetLockedFieldsOk() (*[]ModelModelMetadataFields, bool)`

GetLockedFieldsOk returns a tuple with the LockedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedFields

`func (o *ModelBaseItemDto) SetLockedFields(v []ModelModelMetadataFields)`

SetLockedFields sets LockedFields field to given value.

### HasLockedFields

`func (o *ModelBaseItemDto) HasLockedFields() bool`

HasLockedFields returns a boolean if a field has been set.

### GetLockData

`func (o *ModelBaseItemDto) GetLockData() bool`

GetLockData returns the LockData field if non-nil, zero value otherwise.

### GetLockDataOk

`func (o *ModelBaseItemDto) GetLockDataOk() (*bool, bool)`

GetLockDataOk returns a tuple with the LockData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockData

`func (o *ModelBaseItemDto) SetLockData(v bool)`

SetLockData sets LockData field to given value.

### HasLockData

`func (o *ModelBaseItemDto) HasLockData() bool`

HasLockData returns a boolean if a field has been set.

### SetLockDataNil

`func (o *ModelBaseItemDto) SetLockDataNil(b bool)`

 SetLockDataNil sets the value for LockData to be an explicit nil

### UnsetLockData
`func (o *ModelBaseItemDto) UnsetLockData()`

UnsetLockData ensures that no value is present for LockData, not even an explicit nil
### GetWidth

`func (o *ModelBaseItemDto) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelBaseItemDto) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelBaseItemDto) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelBaseItemDto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ModelBaseItemDto) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ModelBaseItemDto) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *ModelBaseItemDto) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelBaseItemDto) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelBaseItemDto) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelBaseItemDto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ModelBaseItemDto) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ModelBaseItemDto) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetCameraMake

`func (o *ModelBaseItemDto) GetCameraMake() string`

GetCameraMake returns the CameraMake field if non-nil, zero value otherwise.

### GetCameraMakeOk

`func (o *ModelBaseItemDto) GetCameraMakeOk() (*string, bool)`

GetCameraMakeOk returns a tuple with the CameraMake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraMake

`func (o *ModelBaseItemDto) SetCameraMake(v string)`

SetCameraMake sets CameraMake field to given value.

### HasCameraMake

`func (o *ModelBaseItemDto) HasCameraMake() bool`

HasCameraMake returns a boolean if a field has been set.

### GetCameraModel

`func (o *ModelBaseItemDto) GetCameraModel() string`

GetCameraModel returns the CameraModel field if non-nil, zero value otherwise.

### GetCameraModelOk

`func (o *ModelBaseItemDto) GetCameraModelOk() (*string, bool)`

GetCameraModelOk returns a tuple with the CameraModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraModel

`func (o *ModelBaseItemDto) SetCameraModel(v string)`

SetCameraModel sets CameraModel field to given value.

### HasCameraModel

`func (o *ModelBaseItemDto) HasCameraModel() bool`

HasCameraModel returns a boolean if a field has been set.

### GetSoftware

`func (o *ModelBaseItemDto) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ModelBaseItemDto) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ModelBaseItemDto) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ModelBaseItemDto) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetExposureTime

`func (o *ModelBaseItemDto) GetExposureTime() float64`

GetExposureTime returns the ExposureTime field if non-nil, zero value otherwise.

### GetExposureTimeOk

`func (o *ModelBaseItemDto) GetExposureTimeOk() (*float64, bool)`

GetExposureTimeOk returns a tuple with the ExposureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposureTime

`func (o *ModelBaseItemDto) SetExposureTime(v float64)`

SetExposureTime sets ExposureTime field to given value.

### HasExposureTime

`func (o *ModelBaseItemDto) HasExposureTime() bool`

HasExposureTime returns a boolean if a field has been set.

### SetExposureTimeNil

`func (o *ModelBaseItemDto) SetExposureTimeNil(b bool)`

 SetExposureTimeNil sets the value for ExposureTime to be an explicit nil

### UnsetExposureTime
`func (o *ModelBaseItemDto) UnsetExposureTime()`

UnsetExposureTime ensures that no value is present for ExposureTime, not even an explicit nil
### GetFocalLength

`func (o *ModelBaseItemDto) GetFocalLength() float64`

GetFocalLength returns the FocalLength field if non-nil, zero value otherwise.

### GetFocalLengthOk

`func (o *ModelBaseItemDto) GetFocalLengthOk() (*float64, bool)`

GetFocalLengthOk returns a tuple with the FocalLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocalLength

`func (o *ModelBaseItemDto) SetFocalLength(v float64)`

SetFocalLength sets FocalLength field to given value.

### HasFocalLength

`func (o *ModelBaseItemDto) HasFocalLength() bool`

HasFocalLength returns a boolean if a field has been set.

### SetFocalLengthNil

`func (o *ModelBaseItemDto) SetFocalLengthNil(b bool)`

 SetFocalLengthNil sets the value for FocalLength to be an explicit nil

### UnsetFocalLength
`func (o *ModelBaseItemDto) UnsetFocalLength()`

UnsetFocalLength ensures that no value is present for FocalLength, not even an explicit nil
### GetImageOrientation

`func (o *ModelBaseItemDto) GetImageOrientation() ModelModelDrawingImageOrientation`

GetImageOrientation returns the ImageOrientation field if non-nil, zero value otherwise.

### GetImageOrientationOk

`func (o *ModelBaseItemDto) GetImageOrientationOk() (*ModelModelDrawingImageOrientation, bool)`

GetImageOrientationOk returns a tuple with the ImageOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOrientation

`func (o *ModelBaseItemDto) SetImageOrientation(v ModelModelDrawingImageOrientation)`

SetImageOrientation sets ImageOrientation field to given value.

### HasImageOrientation

`func (o *ModelBaseItemDto) HasImageOrientation() bool`

HasImageOrientation returns a boolean if a field has been set.

### GetAperture

`func (o *ModelBaseItemDto) GetAperture() float64`

GetAperture returns the Aperture field if non-nil, zero value otherwise.

### GetApertureOk

`func (o *ModelBaseItemDto) GetApertureOk() (*float64, bool)`

GetApertureOk returns a tuple with the Aperture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAperture

`func (o *ModelBaseItemDto) SetAperture(v float64)`

SetAperture sets Aperture field to given value.

### HasAperture

`func (o *ModelBaseItemDto) HasAperture() bool`

HasAperture returns a boolean if a field has been set.

### SetApertureNil

`func (o *ModelBaseItemDto) SetApertureNil(b bool)`

 SetApertureNil sets the value for Aperture to be an explicit nil

### UnsetAperture
`func (o *ModelBaseItemDto) UnsetAperture()`

UnsetAperture ensures that no value is present for Aperture, not even an explicit nil
### GetShutterSpeed

`func (o *ModelBaseItemDto) GetShutterSpeed() float64`

GetShutterSpeed returns the ShutterSpeed field if non-nil, zero value otherwise.

### GetShutterSpeedOk

`func (o *ModelBaseItemDto) GetShutterSpeedOk() (*float64, bool)`

GetShutterSpeedOk returns a tuple with the ShutterSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShutterSpeed

`func (o *ModelBaseItemDto) SetShutterSpeed(v float64)`

SetShutterSpeed sets ShutterSpeed field to given value.

### HasShutterSpeed

`func (o *ModelBaseItemDto) HasShutterSpeed() bool`

HasShutterSpeed returns a boolean if a field has been set.

### SetShutterSpeedNil

`func (o *ModelBaseItemDto) SetShutterSpeedNil(b bool)`

 SetShutterSpeedNil sets the value for ShutterSpeed to be an explicit nil

### UnsetShutterSpeed
`func (o *ModelBaseItemDto) UnsetShutterSpeed()`

UnsetShutterSpeed ensures that no value is present for ShutterSpeed, not even an explicit nil
### GetLatitude

`func (o *ModelBaseItemDto) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ModelBaseItemDto) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ModelBaseItemDto) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *ModelBaseItemDto) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *ModelBaseItemDto) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *ModelBaseItemDto) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *ModelBaseItemDto) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ModelBaseItemDto) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ModelBaseItemDto) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *ModelBaseItemDto) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *ModelBaseItemDto) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *ModelBaseItemDto) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetAltitude

`func (o *ModelBaseItemDto) GetAltitude() float64`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *ModelBaseItemDto) GetAltitudeOk() (*float64, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *ModelBaseItemDto) SetAltitude(v float64)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *ModelBaseItemDto) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitudeNil

`func (o *ModelBaseItemDto) SetAltitudeNil(b bool)`

 SetAltitudeNil sets the value for Altitude to be an explicit nil

### UnsetAltitude
`func (o *ModelBaseItemDto) UnsetAltitude()`

UnsetAltitude ensures that no value is present for Altitude, not even an explicit nil
### GetIsoSpeedRating

`func (o *ModelBaseItemDto) GetIsoSpeedRating() int32`

GetIsoSpeedRating returns the IsoSpeedRating field if non-nil, zero value otherwise.

### GetIsoSpeedRatingOk

`func (o *ModelBaseItemDto) GetIsoSpeedRatingOk() (*int32, bool)`

GetIsoSpeedRatingOk returns a tuple with the IsoSpeedRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoSpeedRating

`func (o *ModelBaseItemDto) SetIsoSpeedRating(v int32)`

SetIsoSpeedRating sets IsoSpeedRating field to given value.

### HasIsoSpeedRating

`func (o *ModelBaseItemDto) HasIsoSpeedRating() bool`

HasIsoSpeedRating returns a boolean if a field has been set.

### SetIsoSpeedRatingNil

`func (o *ModelBaseItemDto) SetIsoSpeedRatingNil(b bool)`

 SetIsoSpeedRatingNil sets the value for IsoSpeedRating to be an explicit nil

### UnsetIsoSpeedRating
`func (o *ModelBaseItemDto) UnsetIsoSpeedRating()`

UnsetIsoSpeedRating ensures that no value is present for IsoSpeedRating, not even an explicit nil
### GetSeriesTimerId

`func (o *ModelBaseItemDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *ModelBaseItemDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *ModelBaseItemDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *ModelBaseItemDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### GetChannelPrimaryImageTag

`func (o *ModelBaseItemDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *ModelBaseItemDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *ModelBaseItemDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *ModelBaseItemDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### GetStartDate

`func (o *ModelBaseItemDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModelBaseItemDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModelBaseItemDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModelBaseItemDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *ModelBaseItemDto) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *ModelBaseItemDto) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetCompletionPercentage

`func (o *ModelBaseItemDto) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *ModelBaseItemDto) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *ModelBaseItemDto) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *ModelBaseItemDto) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *ModelBaseItemDto) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *ModelBaseItemDto) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetIsRepeat

`func (o *ModelBaseItemDto) GetIsRepeat() bool`

GetIsRepeat returns the IsRepeat field if non-nil, zero value otherwise.

### GetIsRepeatOk

`func (o *ModelBaseItemDto) GetIsRepeatOk() (*bool, bool)`

GetIsRepeatOk returns a tuple with the IsRepeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRepeat

`func (o *ModelBaseItemDto) SetIsRepeat(v bool)`

SetIsRepeat sets IsRepeat field to given value.

### HasIsRepeat

`func (o *ModelBaseItemDto) HasIsRepeat() bool`

HasIsRepeat returns a boolean if a field has been set.

### SetIsRepeatNil

`func (o *ModelBaseItemDto) SetIsRepeatNil(b bool)`

 SetIsRepeatNil sets the value for IsRepeat to be an explicit nil

### UnsetIsRepeat
`func (o *ModelBaseItemDto) UnsetIsRepeat()`

UnsetIsRepeat ensures that no value is present for IsRepeat, not even an explicit nil
### GetIsNew

`func (o *ModelBaseItemDto) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *ModelBaseItemDto) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *ModelBaseItemDto) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *ModelBaseItemDto) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### SetIsNewNil

`func (o *ModelBaseItemDto) SetIsNewNil(b bool)`

 SetIsNewNil sets the value for IsNew to be an explicit nil

### UnsetIsNew
`func (o *ModelBaseItemDto) UnsetIsNew()`

UnsetIsNew ensures that no value is present for IsNew, not even an explicit nil
### GetEpisodeTitle

`func (o *ModelBaseItemDto) GetEpisodeTitle() string`

GetEpisodeTitle returns the EpisodeTitle field if non-nil, zero value otherwise.

### GetEpisodeTitleOk

`func (o *ModelBaseItemDto) GetEpisodeTitleOk() (*string, bool)`

GetEpisodeTitleOk returns a tuple with the EpisodeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeTitle

`func (o *ModelBaseItemDto) SetEpisodeTitle(v string)`

SetEpisodeTitle sets EpisodeTitle field to given value.

### HasEpisodeTitle

`func (o *ModelBaseItemDto) HasEpisodeTitle() bool`

HasEpisodeTitle returns a boolean if a field has been set.

### GetIsMovie

`func (o *ModelBaseItemDto) GetIsMovie() bool`

GetIsMovie returns the IsMovie field if non-nil, zero value otherwise.

### GetIsMovieOk

`func (o *ModelBaseItemDto) GetIsMovieOk() (*bool, bool)`

GetIsMovieOk returns a tuple with the IsMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovie

`func (o *ModelBaseItemDto) SetIsMovie(v bool)`

SetIsMovie sets IsMovie field to given value.

### HasIsMovie

`func (o *ModelBaseItemDto) HasIsMovie() bool`

HasIsMovie returns a boolean if a field has been set.

### SetIsMovieNil

`func (o *ModelBaseItemDto) SetIsMovieNil(b bool)`

 SetIsMovieNil sets the value for IsMovie to be an explicit nil

### UnsetIsMovie
`func (o *ModelBaseItemDto) UnsetIsMovie()`

UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
### GetIsSports

`func (o *ModelBaseItemDto) GetIsSports() bool`

GetIsSports returns the IsSports field if non-nil, zero value otherwise.

### GetIsSportsOk

`func (o *ModelBaseItemDto) GetIsSportsOk() (*bool, bool)`

GetIsSportsOk returns a tuple with the IsSports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSports

`func (o *ModelBaseItemDto) SetIsSports(v bool)`

SetIsSports sets IsSports field to given value.

### HasIsSports

`func (o *ModelBaseItemDto) HasIsSports() bool`

HasIsSports returns a boolean if a field has been set.

### SetIsSportsNil

`func (o *ModelBaseItemDto) SetIsSportsNil(b bool)`

 SetIsSportsNil sets the value for IsSports to be an explicit nil

### UnsetIsSports
`func (o *ModelBaseItemDto) UnsetIsSports()`

UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
### GetIsSeries

`func (o *ModelBaseItemDto) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *ModelBaseItemDto) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *ModelBaseItemDto) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.

### HasIsSeries

`func (o *ModelBaseItemDto) HasIsSeries() bool`

HasIsSeries returns a boolean if a field has been set.

### SetIsSeriesNil

`func (o *ModelBaseItemDto) SetIsSeriesNil(b bool)`

 SetIsSeriesNil sets the value for IsSeries to be an explicit nil

### UnsetIsSeries
`func (o *ModelBaseItemDto) UnsetIsSeries()`

UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
### GetIsLive

`func (o *ModelBaseItemDto) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *ModelBaseItemDto) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *ModelBaseItemDto) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *ModelBaseItemDto) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### SetIsLiveNil

`func (o *ModelBaseItemDto) SetIsLiveNil(b bool)`

 SetIsLiveNil sets the value for IsLive to be an explicit nil

### UnsetIsLive
`func (o *ModelBaseItemDto) UnsetIsLive()`

UnsetIsLive ensures that no value is present for IsLive, not even an explicit nil
### GetIsNews

`func (o *ModelBaseItemDto) GetIsNews() bool`

GetIsNews returns the IsNews field if non-nil, zero value otherwise.

### GetIsNewsOk

`func (o *ModelBaseItemDto) GetIsNewsOk() (*bool, bool)`

GetIsNewsOk returns a tuple with the IsNews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNews

`func (o *ModelBaseItemDto) SetIsNews(v bool)`

SetIsNews sets IsNews field to given value.

### HasIsNews

`func (o *ModelBaseItemDto) HasIsNews() bool`

HasIsNews returns a boolean if a field has been set.

### SetIsNewsNil

`func (o *ModelBaseItemDto) SetIsNewsNil(b bool)`

 SetIsNewsNil sets the value for IsNews to be an explicit nil

### UnsetIsNews
`func (o *ModelBaseItemDto) UnsetIsNews()`

UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
### GetIsKids

`func (o *ModelBaseItemDto) GetIsKids() bool`

GetIsKids returns the IsKids field if non-nil, zero value otherwise.

### GetIsKidsOk

`func (o *ModelBaseItemDto) GetIsKidsOk() (*bool, bool)`

GetIsKidsOk returns a tuple with the IsKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKids

`func (o *ModelBaseItemDto) SetIsKids(v bool)`

SetIsKids sets IsKids field to given value.

### HasIsKids

`func (o *ModelBaseItemDto) HasIsKids() bool`

HasIsKids returns a boolean if a field has been set.

### SetIsKidsNil

`func (o *ModelBaseItemDto) SetIsKidsNil(b bool)`

 SetIsKidsNil sets the value for IsKids to be an explicit nil

### UnsetIsKids
`func (o *ModelBaseItemDto) UnsetIsKids()`

UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
### GetIsPremiere

`func (o *ModelBaseItemDto) GetIsPremiere() bool`

GetIsPremiere returns the IsPremiere field if non-nil, zero value otherwise.

### GetIsPremiereOk

`func (o *ModelBaseItemDto) GetIsPremiereOk() (*bool, bool)`

GetIsPremiereOk returns a tuple with the IsPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremiere

`func (o *ModelBaseItemDto) SetIsPremiere(v bool)`

SetIsPremiere sets IsPremiere field to given value.

### HasIsPremiere

`func (o *ModelBaseItemDto) HasIsPremiere() bool`

HasIsPremiere returns a boolean if a field has been set.

### SetIsPremiereNil

`func (o *ModelBaseItemDto) SetIsPremiereNil(b bool)`

 SetIsPremiereNil sets the value for IsPremiere to be an explicit nil

### UnsetIsPremiere
`func (o *ModelBaseItemDto) UnsetIsPremiere()`

UnsetIsPremiere ensures that no value is present for IsPremiere, not even an explicit nil
### GetTimerType

`func (o *ModelBaseItemDto) GetTimerType() ModelModelLiveTvTimerType`

GetTimerType returns the TimerType field if non-nil, zero value otherwise.

### GetTimerTypeOk

`func (o *ModelBaseItemDto) GetTimerTypeOk() (*ModelModelLiveTvTimerType, bool)`

GetTimerTypeOk returns a tuple with the TimerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerType

`func (o *ModelBaseItemDto) SetTimerType(v ModelModelLiveTvTimerType)`

SetTimerType sets TimerType field to given value.

### HasTimerType

`func (o *ModelBaseItemDto) HasTimerType() bool`

HasTimerType returns a boolean if a field has been set.

### GetDisabled

`func (o *ModelBaseItemDto) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ModelBaseItemDto) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ModelBaseItemDto) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ModelBaseItemDto) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *ModelBaseItemDto) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *ModelBaseItemDto) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetManagementId

`func (o *ModelBaseItemDto) GetManagementId() string`

GetManagementId returns the ManagementId field if non-nil, zero value otherwise.

### GetManagementIdOk

`func (o *ModelBaseItemDto) GetManagementIdOk() (*string, bool)`

GetManagementIdOk returns a tuple with the ManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementId

`func (o *ModelBaseItemDto) SetManagementId(v string)`

SetManagementId sets ManagementId field to given value.

### HasManagementId

`func (o *ModelBaseItemDto) HasManagementId() bool`

HasManagementId returns a boolean if a field has been set.

### GetTimerId

`func (o *ModelBaseItemDto) GetTimerId() string`

GetTimerId returns the TimerId field if non-nil, zero value otherwise.

### GetTimerIdOk

`func (o *ModelBaseItemDto) GetTimerIdOk() (*string, bool)`

GetTimerIdOk returns a tuple with the TimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerId

`func (o *ModelBaseItemDto) SetTimerId(v string)`

SetTimerId sets TimerId field to given value.

### HasTimerId

`func (o *ModelBaseItemDto) HasTimerId() bool`

HasTimerId returns a boolean if a field has been set.

### GetCurrentProgram

`func (o *ModelBaseItemDto) GetCurrentProgram() ModelModelBaseItemDto`

GetCurrentProgram returns the CurrentProgram field if non-nil, zero value otherwise.

### GetCurrentProgramOk

`func (o *ModelBaseItemDto) GetCurrentProgramOk() (*ModelModelBaseItemDto, bool)`

GetCurrentProgramOk returns a tuple with the CurrentProgram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgram

`func (o *ModelBaseItemDto) SetCurrentProgram(v ModelModelBaseItemDto)`

SetCurrentProgram sets CurrentProgram field to given value.

### HasCurrentProgram

`func (o *ModelBaseItemDto) HasCurrentProgram() bool`

HasCurrentProgram returns a boolean if a field has been set.

### GetMovieCount

`func (o *ModelBaseItemDto) GetMovieCount() int32`

GetMovieCount returns the MovieCount field if non-nil, zero value otherwise.

### GetMovieCountOk

`func (o *ModelBaseItemDto) GetMovieCountOk() (*int32, bool)`

GetMovieCountOk returns a tuple with the MovieCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieCount

`func (o *ModelBaseItemDto) SetMovieCount(v int32)`

SetMovieCount sets MovieCount field to given value.

### HasMovieCount

`func (o *ModelBaseItemDto) HasMovieCount() bool`

HasMovieCount returns a boolean if a field has been set.

### SetMovieCountNil

`func (o *ModelBaseItemDto) SetMovieCountNil(b bool)`

 SetMovieCountNil sets the value for MovieCount to be an explicit nil

### UnsetMovieCount
`func (o *ModelBaseItemDto) UnsetMovieCount()`

UnsetMovieCount ensures that no value is present for MovieCount, not even an explicit nil
### GetSeriesCount

`func (o *ModelBaseItemDto) GetSeriesCount() int32`

GetSeriesCount returns the SeriesCount field if non-nil, zero value otherwise.

### GetSeriesCountOk

`func (o *ModelBaseItemDto) GetSeriesCountOk() (*int32, bool)`

GetSeriesCountOk returns a tuple with the SeriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesCount

`func (o *ModelBaseItemDto) SetSeriesCount(v int32)`

SetSeriesCount sets SeriesCount field to given value.

### HasSeriesCount

`func (o *ModelBaseItemDto) HasSeriesCount() bool`

HasSeriesCount returns a boolean if a field has been set.

### SetSeriesCountNil

`func (o *ModelBaseItemDto) SetSeriesCountNil(b bool)`

 SetSeriesCountNil sets the value for SeriesCount to be an explicit nil

### UnsetSeriesCount
`func (o *ModelBaseItemDto) UnsetSeriesCount()`

UnsetSeriesCount ensures that no value is present for SeriesCount, not even an explicit nil
### GetAlbumCount

`func (o *ModelBaseItemDto) GetAlbumCount() int32`

GetAlbumCount returns the AlbumCount field if non-nil, zero value otherwise.

### GetAlbumCountOk

`func (o *ModelBaseItemDto) GetAlbumCountOk() (*int32, bool)`

GetAlbumCountOk returns a tuple with the AlbumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumCount

`func (o *ModelBaseItemDto) SetAlbumCount(v int32)`

SetAlbumCount sets AlbumCount field to given value.

### HasAlbumCount

`func (o *ModelBaseItemDto) HasAlbumCount() bool`

HasAlbumCount returns a boolean if a field has been set.

### SetAlbumCountNil

`func (o *ModelBaseItemDto) SetAlbumCountNil(b bool)`

 SetAlbumCountNil sets the value for AlbumCount to be an explicit nil

### UnsetAlbumCount
`func (o *ModelBaseItemDto) UnsetAlbumCount()`

UnsetAlbumCount ensures that no value is present for AlbumCount, not even an explicit nil
### GetSongCount

`func (o *ModelBaseItemDto) GetSongCount() int32`

GetSongCount returns the SongCount field if non-nil, zero value otherwise.

### GetSongCountOk

`func (o *ModelBaseItemDto) GetSongCountOk() (*int32, bool)`

GetSongCountOk returns a tuple with the SongCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongCount

`func (o *ModelBaseItemDto) SetSongCount(v int32)`

SetSongCount sets SongCount field to given value.

### HasSongCount

`func (o *ModelBaseItemDto) HasSongCount() bool`

HasSongCount returns a boolean if a field has been set.

### SetSongCountNil

`func (o *ModelBaseItemDto) SetSongCountNil(b bool)`

 SetSongCountNil sets the value for SongCount to be an explicit nil

### UnsetSongCount
`func (o *ModelBaseItemDto) UnsetSongCount()`

UnsetSongCount ensures that no value is present for SongCount, not even an explicit nil
### GetMusicVideoCount

`func (o *ModelBaseItemDto) GetMusicVideoCount() int32`

GetMusicVideoCount returns the MusicVideoCount field if non-nil, zero value otherwise.

### GetMusicVideoCountOk

`func (o *ModelBaseItemDto) GetMusicVideoCountOk() (*int32, bool)`

GetMusicVideoCountOk returns a tuple with the MusicVideoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicVideoCount

`func (o *ModelBaseItemDto) SetMusicVideoCount(v int32)`

SetMusicVideoCount sets MusicVideoCount field to given value.

### HasMusicVideoCount

`func (o *ModelBaseItemDto) HasMusicVideoCount() bool`

HasMusicVideoCount returns a boolean if a field has been set.

### SetMusicVideoCountNil

`func (o *ModelBaseItemDto) SetMusicVideoCountNil(b bool)`

 SetMusicVideoCountNil sets the value for MusicVideoCount to be an explicit nil

### UnsetMusicVideoCount
`func (o *ModelBaseItemDto) UnsetMusicVideoCount()`

UnsetMusicVideoCount ensures that no value is present for MusicVideoCount, not even an explicit nil
### GetSubviews

`func (o *ModelBaseItemDto) GetSubviews() []string`

GetSubviews returns the Subviews field if non-nil, zero value otherwise.

### GetSubviewsOk

`func (o *ModelBaseItemDto) GetSubviewsOk() (*[]string, bool)`

GetSubviewsOk returns a tuple with the Subviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubviews

`func (o *ModelBaseItemDto) SetSubviews(v []string)`

SetSubviews sets Subviews field to given value.

### HasSubviews

`func (o *ModelBaseItemDto) HasSubviews() bool`

HasSubviews returns a boolean if a field has been set.

### GetListingsProviderId

`func (o *ModelBaseItemDto) GetListingsProviderId() string`

GetListingsProviderId returns the ListingsProviderId field if non-nil, zero value otherwise.

### GetListingsProviderIdOk

`func (o *ModelBaseItemDto) GetListingsProviderIdOk() (*string, bool)`

GetListingsProviderIdOk returns a tuple with the ListingsProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsProviderId

`func (o *ModelBaseItemDto) SetListingsProviderId(v string)`

SetListingsProviderId sets ListingsProviderId field to given value.

### HasListingsProviderId

`func (o *ModelBaseItemDto) HasListingsProviderId() bool`

HasListingsProviderId returns a boolean if a field has been set.

### GetListingsChannelId

`func (o *ModelBaseItemDto) GetListingsChannelId() string`

GetListingsChannelId returns the ListingsChannelId field if non-nil, zero value otherwise.

### GetListingsChannelIdOk

`func (o *ModelBaseItemDto) GetListingsChannelIdOk() (*string, bool)`

GetListingsChannelIdOk returns a tuple with the ListingsChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsChannelId

`func (o *ModelBaseItemDto) SetListingsChannelId(v string)`

SetListingsChannelId sets ListingsChannelId field to given value.

### HasListingsChannelId

`func (o *ModelBaseItemDto) HasListingsChannelId() bool`

HasListingsChannelId returns a boolean if a field has been set.

### GetListingsPath

`func (o *ModelBaseItemDto) GetListingsPath() string`

GetListingsPath returns the ListingsPath field if non-nil, zero value otherwise.

### GetListingsPathOk

`func (o *ModelBaseItemDto) GetListingsPathOk() (*string, bool)`

GetListingsPathOk returns a tuple with the ListingsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsPath

`func (o *ModelBaseItemDto) SetListingsPath(v string)`

SetListingsPath sets ListingsPath field to given value.

### HasListingsPath

`func (o *ModelBaseItemDto) HasListingsPath() bool`

HasListingsPath returns a boolean if a field has been set.

### GetListingsId

`func (o *ModelBaseItemDto) GetListingsId() string`

GetListingsId returns the ListingsId field if non-nil, zero value otherwise.

### GetListingsIdOk

`func (o *ModelBaseItemDto) GetListingsIdOk() (*string, bool)`

GetListingsIdOk returns a tuple with the ListingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsId

`func (o *ModelBaseItemDto) SetListingsId(v string)`

SetListingsId sets ListingsId field to given value.

### HasListingsId

`func (o *ModelBaseItemDto) HasListingsId() bool`

HasListingsId returns a boolean if a field has been set.

### GetListingsChannelName

`func (o *ModelBaseItemDto) GetListingsChannelName() string`

GetListingsChannelName returns the ListingsChannelName field if non-nil, zero value otherwise.

### GetListingsChannelNameOk

`func (o *ModelBaseItemDto) GetListingsChannelNameOk() (*string, bool)`

GetListingsChannelNameOk returns a tuple with the ListingsChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsChannelName

`func (o *ModelBaseItemDto) SetListingsChannelName(v string)`

SetListingsChannelName sets ListingsChannelName field to given value.

### HasListingsChannelName

`func (o *ModelBaseItemDto) HasListingsChannelName() bool`

HasListingsChannelName returns a boolean if a field has been set.

### GetListingsChannelNumber

`func (o *ModelBaseItemDto) GetListingsChannelNumber() string`

GetListingsChannelNumber returns the ListingsChannelNumber field if non-nil, zero value otherwise.

### GetListingsChannelNumberOk

`func (o *ModelBaseItemDto) GetListingsChannelNumberOk() (*string, bool)`

GetListingsChannelNumberOk returns a tuple with the ListingsChannelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsChannelNumber

`func (o *ModelBaseItemDto) SetListingsChannelNumber(v string)`

SetListingsChannelNumber sets ListingsChannelNumber field to given value.

### HasListingsChannelNumber

`func (o *ModelBaseItemDto) HasListingsChannelNumber() bool`

HasListingsChannelNumber returns a boolean if a field has been set.

### GetAffiliateCallSign

`func (o *ModelBaseItemDto) GetAffiliateCallSign() string`

GetAffiliateCallSign returns the AffiliateCallSign field if non-nil, zero value otherwise.

### GetAffiliateCallSignOk

`func (o *ModelBaseItemDto) GetAffiliateCallSignOk() (*string, bool)`

GetAffiliateCallSignOk returns a tuple with the AffiliateCallSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateCallSign

`func (o *ModelBaseItemDto) SetAffiliateCallSign(v string)`

SetAffiliateCallSign sets AffiliateCallSign field to given value.

### HasAffiliateCallSign

`func (o *ModelBaseItemDto) HasAffiliateCallSign() bool`

HasAffiliateCallSign returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


