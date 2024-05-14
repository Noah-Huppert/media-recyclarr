# ModelApiBaseItemsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Is4K** | Pointer to **NullableBool** |  | [optional] 
**EnableTotalRecordCount** | Pointer to **bool** |  | [optional] 
**RecordingKeyword** | Pointer to **string** |  | [optional] 
**RecordingKeywordType** | Pointer to [**ModelModelLiveTvKeywordType**](ModelLiveTvKeywordType.md) |  | [optional] 
**RandomSeed** | Pointer to **int32** |  | [optional] 
**GenreIds** | Pointer to **string** |  | [optional] 
**CollectionIds** | Pointer to **string** |  | [optional] 
**TagIds** | Pointer to **string** |  | [optional] 
**ExcludeTagIds** | Pointer to **string** |  | [optional] 
**ExcludeArtistIds** | Pointer to **string** |  | [optional] 
**AlbumArtistIds** | Pointer to **string** |  | [optional] 
**ContributingArtistIds** | Pointer to **string** |  | [optional] 
**AlbumIds** | Pointer to **string** |  | [optional] 
**OuterIds** | Pointer to **string** |  | [optional] 
**ListItemIds** | Pointer to **string** |  | [optional] 
**AudioLanguages** | Pointer to **string** |  | [optional] 
**SubtitleLanguages** | Pointer to **string** |  | [optional] 
**GroupItemsInto** | Pointer to [**ModelModelLibraryItemLinkType**](ModelLibraryItemLinkType.md) |  | [optional] 
**MinWidth** | Pointer to **NullableInt32** |  | [optional] 
**MinHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**GroupProgramsBySeries** | Pointer to **bool** |  | [optional] 
**AirDays** | Pointer to [**[]ModelModelDayOfWeek**](ModelModelDayOfWeek.md) |  | [optional] 
**IsAiring** | Pointer to **NullableBool** |  | [optional] 
**HasAired** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewModelApiBaseItemsRequest

`func NewModelApiBaseItemsRequest() *ModelApiBaseItemsRequest`

NewModelApiBaseItemsRequest instantiates a new ModelApiBaseItemsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelApiBaseItemsRequestWithDefaults

`func NewModelApiBaseItemsRequestWithDefaults() *ModelApiBaseItemsRequest`

NewModelApiBaseItemsRequestWithDefaults instantiates a new ModelApiBaseItemsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIs4K

`func (o *ModelApiBaseItemsRequest) GetIs4K() bool`

GetIs4K returns the Is4K field if non-nil, zero value otherwise.

### GetIs4KOk

`func (o *ModelApiBaseItemsRequest) GetIs4KOk() (*bool, bool)`

GetIs4KOk returns a tuple with the Is4K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIs4K

`func (o *ModelApiBaseItemsRequest) SetIs4K(v bool)`

SetIs4K sets Is4K field to given value.

### HasIs4K

`func (o *ModelApiBaseItemsRequest) HasIs4K() bool`

HasIs4K returns a boolean if a field has been set.

### SetIs4KNil

`func (o *ModelApiBaseItemsRequest) SetIs4KNil(b bool)`

 SetIs4KNil sets the value for Is4K to be an explicit nil

### UnsetIs4K
`func (o *ModelApiBaseItemsRequest) UnsetIs4K()`

UnsetIs4K ensures that no value is present for Is4K, not even an explicit nil
### GetEnableTotalRecordCount

`func (o *ModelApiBaseItemsRequest) GetEnableTotalRecordCount() bool`

GetEnableTotalRecordCount returns the EnableTotalRecordCount field if non-nil, zero value otherwise.

### GetEnableTotalRecordCountOk

`func (o *ModelApiBaseItemsRequest) GetEnableTotalRecordCountOk() (*bool, bool)`

GetEnableTotalRecordCountOk returns a tuple with the EnableTotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTotalRecordCount

`func (o *ModelApiBaseItemsRequest) SetEnableTotalRecordCount(v bool)`

SetEnableTotalRecordCount sets EnableTotalRecordCount field to given value.

### HasEnableTotalRecordCount

`func (o *ModelApiBaseItemsRequest) HasEnableTotalRecordCount() bool`

HasEnableTotalRecordCount returns a boolean if a field has been set.

### GetRecordingKeyword

`func (o *ModelApiBaseItemsRequest) GetRecordingKeyword() string`

GetRecordingKeyword returns the RecordingKeyword field if non-nil, zero value otherwise.

### GetRecordingKeywordOk

`func (o *ModelApiBaseItemsRequest) GetRecordingKeywordOk() (*string, bool)`

GetRecordingKeywordOk returns a tuple with the RecordingKeyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingKeyword

`func (o *ModelApiBaseItemsRequest) SetRecordingKeyword(v string)`

SetRecordingKeyword sets RecordingKeyword field to given value.

### HasRecordingKeyword

`func (o *ModelApiBaseItemsRequest) HasRecordingKeyword() bool`

HasRecordingKeyword returns a boolean if a field has been set.

### GetRecordingKeywordType

`func (o *ModelApiBaseItemsRequest) GetRecordingKeywordType() ModelModelLiveTvKeywordType`

GetRecordingKeywordType returns the RecordingKeywordType field if non-nil, zero value otherwise.

### GetRecordingKeywordTypeOk

`func (o *ModelApiBaseItemsRequest) GetRecordingKeywordTypeOk() (*ModelModelLiveTvKeywordType, bool)`

GetRecordingKeywordTypeOk returns a tuple with the RecordingKeywordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingKeywordType

`func (o *ModelApiBaseItemsRequest) SetRecordingKeywordType(v ModelModelLiveTvKeywordType)`

SetRecordingKeywordType sets RecordingKeywordType field to given value.

### HasRecordingKeywordType

`func (o *ModelApiBaseItemsRequest) HasRecordingKeywordType() bool`

HasRecordingKeywordType returns a boolean if a field has been set.

### GetRandomSeed

`func (o *ModelApiBaseItemsRequest) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *ModelApiBaseItemsRequest) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *ModelApiBaseItemsRequest) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *ModelApiBaseItemsRequest) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### GetGenreIds

`func (o *ModelApiBaseItemsRequest) GetGenreIds() string`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *ModelApiBaseItemsRequest) GetGenreIdsOk() (*string, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *ModelApiBaseItemsRequest) SetGenreIds(v string)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *ModelApiBaseItemsRequest) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetCollectionIds

`func (o *ModelApiBaseItemsRequest) GetCollectionIds() string`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *ModelApiBaseItemsRequest) GetCollectionIdsOk() (*string, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *ModelApiBaseItemsRequest) SetCollectionIds(v string)`

SetCollectionIds sets CollectionIds field to given value.

### HasCollectionIds

`func (o *ModelApiBaseItemsRequest) HasCollectionIds() bool`

HasCollectionIds returns a boolean if a field has been set.

### GetTagIds

`func (o *ModelApiBaseItemsRequest) GetTagIds() string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ModelApiBaseItemsRequest) GetTagIdsOk() (*string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ModelApiBaseItemsRequest) SetTagIds(v string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ModelApiBaseItemsRequest) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetExcludeTagIds

`func (o *ModelApiBaseItemsRequest) GetExcludeTagIds() string`

GetExcludeTagIds returns the ExcludeTagIds field if non-nil, zero value otherwise.

### GetExcludeTagIdsOk

`func (o *ModelApiBaseItemsRequest) GetExcludeTagIdsOk() (*string, bool)`

GetExcludeTagIdsOk returns a tuple with the ExcludeTagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeTagIds

`func (o *ModelApiBaseItemsRequest) SetExcludeTagIds(v string)`

SetExcludeTagIds sets ExcludeTagIds field to given value.

### HasExcludeTagIds

`func (o *ModelApiBaseItemsRequest) HasExcludeTagIds() bool`

HasExcludeTagIds returns a boolean if a field has been set.

### GetExcludeArtistIds

`func (o *ModelApiBaseItemsRequest) GetExcludeArtistIds() string`

GetExcludeArtistIds returns the ExcludeArtistIds field if non-nil, zero value otherwise.

### GetExcludeArtistIdsOk

`func (o *ModelApiBaseItemsRequest) GetExcludeArtistIdsOk() (*string, bool)`

GetExcludeArtistIdsOk returns a tuple with the ExcludeArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeArtistIds

`func (o *ModelApiBaseItemsRequest) SetExcludeArtistIds(v string)`

SetExcludeArtistIds sets ExcludeArtistIds field to given value.

### HasExcludeArtistIds

`func (o *ModelApiBaseItemsRequest) HasExcludeArtistIds() bool`

HasExcludeArtistIds returns a boolean if a field has been set.

### GetAlbumArtistIds

`func (o *ModelApiBaseItemsRequest) GetAlbumArtistIds() string`

GetAlbumArtistIds returns the AlbumArtistIds field if non-nil, zero value otherwise.

### GetAlbumArtistIdsOk

`func (o *ModelApiBaseItemsRequest) GetAlbumArtistIdsOk() (*string, bool)`

GetAlbumArtistIdsOk returns a tuple with the AlbumArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtistIds

`func (o *ModelApiBaseItemsRequest) SetAlbumArtistIds(v string)`

SetAlbumArtistIds sets AlbumArtistIds field to given value.

### HasAlbumArtistIds

`func (o *ModelApiBaseItemsRequest) HasAlbumArtistIds() bool`

HasAlbumArtistIds returns a boolean if a field has been set.

### GetContributingArtistIds

`func (o *ModelApiBaseItemsRequest) GetContributingArtistIds() string`

GetContributingArtistIds returns the ContributingArtistIds field if non-nil, zero value otherwise.

### GetContributingArtistIdsOk

`func (o *ModelApiBaseItemsRequest) GetContributingArtistIdsOk() (*string, bool)`

GetContributingArtistIdsOk returns a tuple with the ContributingArtistIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributingArtistIds

`func (o *ModelApiBaseItemsRequest) SetContributingArtistIds(v string)`

SetContributingArtistIds sets ContributingArtistIds field to given value.

### HasContributingArtistIds

`func (o *ModelApiBaseItemsRequest) HasContributingArtistIds() bool`

HasContributingArtistIds returns a boolean if a field has been set.

### GetAlbumIds

`func (o *ModelApiBaseItemsRequest) GetAlbumIds() string`

GetAlbumIds returns the AlbumIds field if non-nil, zero value otherwise.

### GetAlbumIdsOk

`func (o *ModelApiBaseItemsRequest) GetAlbumIdsOk() (*string, bool)`

GetAlbumIdsOk returns a tuple with the AlbumIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumIds

`func (o *ModelApiBaseItemsRequest) SetAlbumIds(v string)`

SetAlbumIds sets AlbumIds field to given value.

### HasAlbumIds

`func (o *ModelApiBaseItemsRequest) HasAlbumIds() bool`

HasAlbumIds returns a boolean if a field has been set.

### GetOuterIds

`func (o *ModelApiBaseItemsRequest) GetOuterIds() string`

GetOuterIds returns the OuterIds field if non-nil, zero value otherwise.

### GetOuterIdsOk

`func (o *ModelApiBaseItemsRequest) GetOuterIdsOk() (*string, bool)`

GetOuterIdsOk returns a tuple with the OuterIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuterIds

`func (o *ModelApiBaseItemsRequest) SetOuterIds(v string)`

SetOuterIds sets OuterIds field to given value.

### HasOuterIds

`func (o *ModelApiBaseItemsRequest) HasOuterIds() bool`

HasOuterIds returns a boolean if a field has been set.

### GetListItemIds

`func (o *ModelApiBaseItemsRequest) GetListItemIds() string`

GetListItemIds returns the ListItemIds field if non-nil, zero value otherwise.

### GetListItemIdsOk

`func (o *ModelApiBaseItemsRequest) GetListItemIdsOk() (*string, bool)`

GetListItemIdsOk returns a tuple with the ListItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListItemIds

`func (o *ModelApiBaseItemsRequest) SetListItemIds(v string)`

SetListItemIds sets ListItemIds field to given value.

### HasListItemIds

`func (o *ModelApiBaseItemsRequest) HasListItemIds() bool`

HasListItemIds returns a boolean if a field has been set.

### GetAudioLanguages

`func (o *ModelApiBaseItemsRequest) GetAudioLanguages() string`

GetAudioLanguages returns the AudioLanguages field if non-nil, zero value otherwise.

### GetAudioLanguagesOk

`func (o *ModelApiBaseItemsRequest) GetAudioLanguagesOk() (*string, bool)`

GetAudioLanguagesOk returns a tuple with the AudioLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguages

`func (o *ModelApiBaseItemsRequest) SetAudioLanguages(v string)`

SetAudioLanguages sets AudioLanguages field to given value.

### HasAudioLanguages

`func (o *ModelApiBaseItemsRequest) HasAudioLanguages() bool`

HasAudioLanguages returns a boolean if a field has been set.

### GetSubtitleLanguages

`func (o *ModelApiBaseItemsRequest) GetSubtitleLanguages() string`

GetSubtitleLanguages returns the SubtitleLanguages field if non-nil, zero value otherwise.

### GetSubtitleLanguagesOk

`func (o *ModelApiBaseItemsRequest) GetSubtitleLanguagesOk() (*string, bool)`

GetSubtitleLanguagesOk returns a tuple with the SubtitleLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguages

`func (o *ModelApiBaseItemsRequest) SetSubtitleLanguages(v string)`

SetSubtitleLanguages sets SubtitleLanguages field to given value.

### HasSubtitleLanguages

`func (o *ModelApiBaseItemsRequest) HasSubtitleLanguages() bool`

HasSubtitleLanguages returns a boolean if a field has been set.

### GetGroupItemsInto

`func (o *ModelApiBaseItemsRequest) GetGroupItemsInto() ModelModelLibraryItemLinkType`

GetGroupItemsInto returns the GroupItemsInto field if non-nil, zero value otherwise.

### GetGroupItemsIntoOk

`func (o *ModelApiBaseItemsRequest) GetGroupItemsIntoOk() (*ModelModelLibraryItemLinkType, bool)`

GetGroupItemsIntoOk returns a tuple with the GroupItemsInto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupItemsInto

`func (o *ModelApiBaseItemsRequest) SetGroupItemsInto(v ModelModelLibraryItemLinkType)`

SetGroupItemsInto sets GroupItemsInto field to given value.

### HasGroupItemsInto

`func (o *ModelApiBaseItemsRequest) HasGroupItemsInto() bool`

HasGroupItemsInto returns a boolean if a field has been set.

### GetMinWidth

`func (o *ModelApiBaseItemsRequest) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *ModelApiBaseItemsRequest) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *ModelApiBaseItemsRequest) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *ModelApiBaseItemsRequest) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### SetMinWidthNil

`func (o *ModelApiBaseItemsRequest) SetMinWidthNil(b bool)`

 SetMinWidthNil sets the value for MinWidth to be an explicit nil

### UnsetMinWidth
`func (o *ModelApiBaseItemsRequest) UnsetMinWidth()`

UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
### GetMinHeight

`func (o *ModelApiBaseItemsRequest) GetMinHeight() int32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *ModelApiBaseItemsRequest) GetMinHeightOk() (*int32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *ModelApiBaseItemsRequest) SetMinHeight(v int32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *ModelApiBaseItemsRequest) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### SetMinHeightNil

`func (o *ModelApiBaseItemsRequest) SetMinHeightNil(b bool)`

 SetMinHeightNil sets the value for MinHeight to be an explicit nil

### UnsetMinHeight
`func (o *ModelApiBaseItemsRequest) UnsetMinHeight()`

UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
### GetMaxWidth

`func (o *ModelApiBaseItemsRequest) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *ModelApiBaseItemsRequest) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *ModelApiBaseItemsRequest) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *ModelApiBaseItemsRequest) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *ModelApiBaseItemsRequest) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *ModelApiBaseItemsRequest) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMaxHeight

`func (o *ModelApiBaseItemsRequest) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ModelApiBaseItemsRequest) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ModelApiBaseItemsRequest) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ModelApiBaseItemsRequest) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *ModelApiBaseItemsRequest) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *ModelApiBaseItemsRequest) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetGroupProgramsBySeries

`func (o *ModelApiBaseItemsRequest) GetGroupProgramsBySeries() bool`

GetGroupProgramsBySeries returns the GroupProgramsBySeries field if non-nil, zero value otherwise.

### GetGroupProgramsBySeriesOk

`func (o *ModelApiBaseItemsRequest) GetGroupProgramsBySeriesOk() (*bool, bool)`

GetGroupProgramsBySeriesOk returns a tuple with the GroupProgramsBySeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupProgramsBySeries

`func (o *ModelApiBaseItemsRequest) SetGroupProgramsBySeries(v bool)`

SetGroupProgramsBySeries sets GroupProgramsBySeries field to given value.

### HasGroupProgramsBySeries

`func (o *ModelApiBaseItemsRequest) HasGroupProgramsBySeries() bool`

HasGroupProgramsBySeries returns a boolean if a field has been set.

### GetAirDays

`func (o *ModelApiBaseItemsRequest) GetAirDays() []ModelModelDayOfWeek`

GetAirDays returns the AirDays field if non-nil, zero value otherwise.

### GetAirDaysOk

`func (o *ModelApiBaseItemsRequest) GetAirDaysOk() (*[]ModelModelDayOfWeek, bool)`

GetAirDaysOk returns a tuple with the AirDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDays

`func (o *ModelApiBaseItemsRequest) SetAirDays(v []ModelModelDayOfWeek)`

SetAirDays sets AirDays field to given value.

### HasAirDays

`func (o *ModelApiBaseItemsRequest) HasAirDays() bool`

HasAirDays returns a boolean if a field has been set.

### GetIsAiring

`func (o *ModelApiBaseItemsRequest) GetIsAiring() bool`

GetIsAiring returns the IsAiring field if non-nil, zero value otherwise.

### GetIsAiringOk

`func (o *ModelApiBaseItemsRequest) GetIsAiringOk() (*bool, bool)`

GetIsAiringOk returns a tuple with the IsAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiring

`func (o *ModelApiBaseItemsRequest) SetIsAiring(v bool)`

SetIsAiring sets IsAiring field to given value.

### HasIsAiring

`func (o *ModelApiBaseItemsRequest) HasIsAiring() bool`

HasIsAiring returns a boolean if a field has been set.

### SetIsAiringNil

`func (o *ModelApiBaseItemsRequest) SetIsAiringNil(b bool)`

 SetIsAiringNil sets the value for IsAiring to be an explicit nil

### UnsetIsAiring
`func (o *ModelApiBaseItemsRequest) UnsetIsAiring()`

UnsetIsAiring ensures that no value is present for IsAiring, not even an explicit nil
### GetHasAired

`func (o *ModelApiBaseItemsRequest) GetHasAired() bool`

GetHasAired returns the HasAired field if non-nil, zero value otherwise.

### GetHasAiredOk

`func (o *ModelApiBaseItemsRequest) GetHasAiredOk() (*bool, bool)`

GetHasAiredOk returns a tuple with the HasAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAired

`func (o *ModelApiBaseItemsRequest) SetHasAired(v bool)`

SetHasAired sets HasAired field to given value.

### HasHasAired

`func (o *ModelApiBaseItemsRequest) HasHasAired() bool`

HasHasAired returns a boolean if a field has been set.

### SetHasAiredNil

`func (o *ModelApiBaseItemsRequest) SetHasAiredNil(b bool)`

 SetHasAiredNil sets the value for HasAired to be an explicit nil

### UnsetHasAired
`func (o *ModelApiBaseItemsRequest) UnsetHasAired()`

UnsetHasAired ensures that no value is present for HasAired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


