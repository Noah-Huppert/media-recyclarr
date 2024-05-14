# ModelSongInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlbumArtists** | Pointer to **[]string** |  | [optional] 
**Album** | Pointer to **string** |  | [optional] 
**Artists** | Pointer to **[]string** |  | [optional] 
**Composers** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**MetadataLanguage** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**MetadataLanguages** | Pointer to [**[]ModelModelGlobalizationCultureDto**](ModelModelGlobalizationCultureDto.md) |  | [optional] 
**ProviderIds** | Pointer to **map[string]string** |  | [optional] 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**IndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**PremiereDate** | Pointer to **NullableTime** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**EnableAdultMetadata** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelSongInfo

`func NewModelSongInfo() *ModelSongInfo`

NewModelSongInfo instantiates a new ModelSongInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSongInfoWithDefaults

`func NewModelSongInfoWithDefaults() *ModelSongInfo`

NewModelSongInfoWithDefaults instantiates a new ModelSongInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbumArtists

`func (o *ModelSongInfo) GetAlbumArtists() []string`

GetAlbumArtists returns the AlbumArtists field if non-nil, zero value otherwise.

### GetAlbumArtistsOk

`func (o *ModelSongInfo) GetAlbumArtistsOk() (*[]string, bool)`

GetAlbumArtistsOk returns a tuple with the AlbumArtists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtists

`func (o *ModelSongInfo) SetAlbumArtists(v []string)`

SetAlbumArtists sets AlbumArtists field to given value.

### HasAlbumArtists

`func (o *ModelSongInfo) HasAlbumArtists() bool`

HasAlbumArtists returns a boolean if a field has been set.

### GetAlbum

`func (o *ModelSongInfo) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *ModelSongInfo) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *ModelSongInfo) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *ModelSongInfo) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### GetArtists

`func (o *ModelSongInfo) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *ModelSongInfo) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *ModelSongInfo) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *ModelSongInfo) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetComposers

`func (o *ModelSongInfo) GetComposers() []string`

GetComposers returns the Composers field if non-nil, zero value otherwise.

### GetComposersOk

`func (o *ModelSongInfo) GetComposersOk() (*[]string, bool)`

GetComposersOk returns a tuple with the Composers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposers

`func (o *ModelSongInfo) SetComposers(v []string)`

SetComposers sets Composers field to given value.

### HasComposers

`func (o *ModelSongInfo) HasComposers() bool`

HasComposers returns a boolean if a field has been set.

### GetName

`func (o *ModelSongInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSongInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSongInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSongInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadataLanguage

`func (o *ModelSongInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *ModelSongInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *ModelSongInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *ModelSongInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ModelSongInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ModelSongInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ModelSongInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ModelSongInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetMetadataLanguages

`func (o *ModelSongInfo) GetMetadataLanguages() []ModelModelGlobalizationCultureDto`

GetMetadataLanguages returns the MetadataLanguages field if non-nil, zero value otherwise.

### GetMetadataLanguagesOk

`func (o *ModelSongInfo) GetMetadataLanguagesOk() (*[]ModelModelGlobalizationCultureDto, bool)`

GetMetadataLanguagesOk returns a tuple with the MetadataLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguages

`func (o *ModelSongInfo) SetMetadataLanguages(v []ModelModelGlobalizationCultureDto)`

SetMetadataLanguages sets MetadataLanguages field to given value.

### HasMetadataLanguages

`func (o *ModelSongInfo) HasMetadataLanguages() bool`

HasMetadataLanguages returns a boolean if a field has been set.

### GetProviderIds

`func (o *ModelSongInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *ModelSongInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *ModelSongInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *ModelSongInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetYear

`func (o *ModelSongInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ModelSongInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ModelSongInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ModelSongInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *ModelSongInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *ModelSongInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *ModelSongInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *ModelSongInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *ModelSongInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *ModelSongInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *ModelSongInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *ModelSongInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *ModelSongInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *ModelSongInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *ModelSongInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *ModelSongInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *ModelSongInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *ModelSongInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *ModelSongInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *ModelSongInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *ModelSongInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *ModelSongInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *ModelSongInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *ModelSongInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *ModelSongInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *ModelSongInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *ModelSongInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *ModelSongInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetEnableAdultMetadata

`func (o *ModelSongInfo) GetEnableAdultMetadata() bool`

GetEnableAdultMetadata returns the EnableAdultMetadata field if non-nil, zero value otherwise.

### GetEnableAdultMetadataOk

`func (o *ModelSongInfo) GetEnableAdultMetadataOk() (*bool, bool)`

GetEnableAdultMetadataOk returns a tuple with the EnableAdultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdultMetadata

`func (o *ModelSongInfo) SetEnableAdultMetadata(v bool)`

SetEnableAdultMetadata sets EnableAdultMetadata field to given value.

### HasEnableAdultMetadata

`func (o *ModelSongInfo) HasEnableAdultMetadata() bool`

HasEnableAdultMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


