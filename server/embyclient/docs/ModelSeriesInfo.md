# ModelSeriesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EpisodeAirDate** | Pointer to **NullableTime** |  | [optional] 
**DisplayOrder** | Pointer to [**ModelModelSeriesDisplayOrder**](ModelSeriesDisplayOrder.md) |  | [optional] 
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

### NewModelSeriesInfo

`func NewModelSeriesInfo() *ModelSeriesInfo`

NewModelSeriesInfo instantiates a new ModelSeriesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSeriesInfoWithDefaults

`func NewModelSeriesInfoWithDefaults() *ModelSeriesInfo`

NewModelSeriesInfoWithDefaults instantiates a new ModelSeriesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpisodeAirDate

`func (o *ModelSeriesInfo) GetEpisodeAirDate() time.Time`

GetEpisodeAirDate returns the EpisodeAirDate field if non-nil, zero value otherwise.

### GetEpisodeAirDateOk

`func (o *ModelSeriesInfo) GetEpisodeAirDateOk() (*time.Time, bool)`

GetEpisodeAirDateOk returns a tuple with the EpisodeAirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeAirDate

`func (o *ModelSeriesInfo) SetEpisodeAirDate(v time.Time)`

SetEpisodeAirDate sets EpisodeAirDate field to given value.

### HasEpisodeAirDate

`func (o *ModelSeriesInfo) HasEpisodeAirDate() bool`

HasEpisodeAirDate returns a boolean if a field has been set.

### SetEpisodeAirDateNil

`func (o *ModelSeriesInfo) SetEpisodeAirDateNil(b bool)`

 SetEpisodeAirDateNil sets the value for EpisodeAirDate to be an explicit nil

### UnsetEpisodeAirDate
`func (o *ModelSeriesInfo) UnsetEpisodeAirDate()`

UnsetEpisodeAirDate ensures that no value is present for EpisodeAirDate, not even an explicit nil
### GetDisplayOrder

`func (o *ModelSeriesInfo) GetDisplayOrder() ModelModelSeriesDisplayOrder`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ModelSeriesInfo) GetDisplayOrderOk() (*ModelModelSeriesDisplayOrder, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ModelSeriesInfo) SetDisplayOrder(v ModelModelSeriesDisplayOrder)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ModelSeriesInfo) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetName

`func (o *ModelSeriesInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelSeriesInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelSeriesInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelSeriesInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadataLanguage

`func (o *ModelSeriesInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *ModelSeriesInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *ModelSeriesInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *ModelSeriesInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ModelSeriesInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ModelSeriesInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ModelSeriesInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ModelSeriesInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetMetadataLanguages

`func (o *ModelSeriesInfo) GetMetadataLanguages() []ModelModelGlobalizationCultureDto`

GetMetadataLanguages returns the MetadataLanguages field if non-nil, zero value otherwise.

### GetMetadataLanguagesOk

`func (o *ModelSeriesInfo) GetMetadataLanguagesOk() (*[]ModelModelGlobalizationCultureDto, bool)`

GetMetadataLanguagesOk returns a tuple with the MetadataLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguages

`func (o *ModelSeriesInfo) SetMetadataLanguages(v []ModelModelGlobalizationCultureDto)`

SetMetadataLanguages sets MetadataLanguages field to given value.

### HasMetadataLanguages

`func (o *ModelSeriesInfo) HasMetadataLanguages() bool`

HasMetadataLanguages returns a boolean if a field has been set.

### GetProviderIds

`func (o *ModelSeriesInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *ModelSeriesInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *ModelSeriesInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *ModelSeriesInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### GetYear

`func (o *ModelSeriesInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ModelSeriesInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ModelSeriesInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ModelSeriesInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *ModelSeriesInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *ModelSeriesInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *ModelSeriesInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *ModelSeriesInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *ModelSeriesInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *ModelSeriesInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *ModelSeriesInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *ModelSeriesInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *ModelSeriesInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *ModelSeriesInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *ModelSeriesInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *ModelSeriesInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *ModelSeriesInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *ModelSeriesInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *ModelSeriesInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *ModelSeriesInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *ModelSeriesInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *ModelSeriesInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *ModelSeriesInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *ModelSeriesInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *ModelSeriesInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *ModelSeriesInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *ModelSeriesInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *ModelSeriesInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetEnableAdultMetadata

`func (o *ModelSeriesInfo) GetEnableAdultMetadata() bool`

GetEnableAdultMetadata returns the EnableAdultMetadata field if non-nil, zero value otherwise.

### GetEnableAdultMetadataOk

`func (o *ModelSeriesInfo) GetEnableAdultMetadataOk() (*bool, bool)`

GetEnableAdultMetadataOk returns a tuple with the EnableAdultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdultMetadata

`func (o *ModelSeriesInfo) SetEnableAdultMetadata(v bool)`

SetEnableAdultMetadata sets EnableAdultMetadata field to given value.

### HasEnableAdultMetadata

`func (o *ModelSeriesInfo) HasEnableAdultMetadata() bool`

HasEnableAdultMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


