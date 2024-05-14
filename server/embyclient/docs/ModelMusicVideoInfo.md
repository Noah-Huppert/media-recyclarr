# ModelMusicVideoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artists** | Pointer to **[]string** |  | [optional] 
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

### NewModelMusicVideoInfo

`func NewModelMusicVideoInfo() *ModelMusicVideoInfo`

NewModelMusicVideoInfo instantiates a new ModelMusicVideoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMusicVideoInfoWithDefaults

`func NewModelMusicVideoInfoWithDefaults() *ModelMusicVideoInfo`

NewModelMusicVideoInfoWithDefaults instantiates a new ModelMusicVideoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtists

`func (o *ModelMusicVideoInfo) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *ModelMusicVideoInfo) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *ModelMusicVideoInfo) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *ModelMusicVideoInfo) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetName

`func (o *ModelMusicVideoInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelMusicVideoInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelMusicVideoInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelMusicVideoInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadataLanguage

`func (o *ModelMusicVideoInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *ModelMusicVideoInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *ModelMusicVideoInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *ModelMusicVideoInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ModelMusicVideoInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ModelMusicVideoInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ModelMusicVideoInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ModelMusicVideoInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetMetadataLanguages

`func (o *ModelMusicVideoInfo) GetMetadataLanguages() []ModelModelGlobalizationCultureDto`

GetMetadataLanguages returns the MetadataLanguages field if non-nil, zero value otherwise.

### GetMetadataLanguagesOk

`func (o *ModelMusicVideoInfo) GetMetadataLanguagesOk() (*[]ModelModelGlobalizationCultureDto, bool)`

GetMetadataLanguagesOk returns a tuple with the MetadataLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguages

`func (o *ModelMusicVideoInfo) SetMetadataLanguages(v []ModelModelGlobalizationCultureDto)`

SetMetadataLanguages sets MetadataLanguages field to given value.

### HasMetadataLanguages

`func (o *ModelMusicVideoInfo) HasMetadataLanguages() bool`

HasMetadataLanguages returns a boolean if a field has been set.

### GetProviderIds

`func (o *ModelMusicVideoInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *ModelMusicVideoInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *ModelMusicVideoInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *ModelMusicVideoInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetYear

`func (o *ModelMusicVideoInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ModelMusicVideoInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ModelMusicVideoInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ModelMusicVideoInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *ModelMusicVideoInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *ModelMusicVideoInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *ModelMusicVideoInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *ModelMusicVideoInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *ModelMusicVideoInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *ModelMusicVideoInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *ModelMusicVideoInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *ModelMusicVideoInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *ModelMusicVideoInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *ModelMusicVideoInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *ModelMusicVideoInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *ModelMusicVideoInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *ModelMusicVideoInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *ModelMusicVideoInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *ModelMusicVideoInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *ModelMusicVideoInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *ModelMusicVideoInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *ModelMusicVideoInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *ModelMusicVideoInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *ModelMusicVideoInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *ModelMusicVideoInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *ModelMusicVideoInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *ModelMusicVideoInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *ModelMusicVideoInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetEnableAdultMetadata

`func (o *ModelMusicVideoInfo) GetEnableAdultMetadata() bool`

GetEnableAdultMetadata returns the EnableAdultMetadata field if non-nil, zero value otherwise.

### GetEnableAdultMetadataOk

`func (o *ModelMusicVideoInfo) GetEnableAdultMetadataOk() (*bool, bool)`

GetEnableAdultMetadataOk returns a tuple with the EnableAdultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdultMetadata

`func (o *ModelMusicVideoInfo) SetEnableAdultMetadata(v bool)`

SetEnableAdultMetadata sets EnableAdultMetadata field to given value.

### HasEnableAdultMetadata

`func (o *ModelMusicVideoInfo) HasEnableAdultMetadata() bool`

HasEnableAdultMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


