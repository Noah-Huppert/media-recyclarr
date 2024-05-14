# ModelRemoteSubtitleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeLetterISOLanguageName** | Pointer to **string** | Use language instead to return the language specified by the subtitle provider | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**CommunityRating** | Pointer to **NullableFloat32** |  | [optional] 
**DownloadCount** | Pointer to **NullableInt32** |  | [optional] 
**IsHashMatch** | Pointer to **NullableBool** |  | [optional] 
**IsForced** | Pointer to **NullableBool** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewModelRemoteSubtitleInfo

`func NewModelRemoteSubtitleInfo() *ModelRemoteSubtitleInfo`

NewModelRemoteSubtitleInfo instantiates a new ModelRemoteSubtitleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRemoteSubtitleInfoWithDefaults

`func NewModelRemoteSubtitleInfoWithDefaults() *ModelRemoteSubtitleInfo`

NewModelRemoteSubtitleInfoWithDefaults instantiates a new ModelRemoteSubtitleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeLetterISOLanguageName

`func (o *ModelRemoteSubtitleInfo) GetThreeLetterISOLanguageName() string`

GetThreeLetterISOLanguageName returns the ThreeLetterISOLanguageName field if non-nil, zero value otherwise.

### GetThreeLetterISOLanguageNameOk

`func (o *ModelRemoteSubtitleInfo) GetThreeLetterISOLanguageNameOk() (*string, bool)`

GetThreeLetterISOLanguageNameOk returns a tuple with the ThreeLetterISOLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeLetterISOLanguageName

`func (o *ModelRemoteSubtitleInfo) SetThreeLetterISOLanguageName(v string)`

SetThreeLetterISOLanguageName sets ThreeLetterISOLanguageName field to given value.

### HasThreeLetterISOLanguageName

`func (o *ModelRemoteSubtitleInfo) HasThreeLetterISOLanguageName() bool`

HasThreeLetterISOLanguageName returns a boolean if a field has been set.

### GetId

`func (o *ModelRemoteSubtitleInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelRemoteSubtitleInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelRemoteSubtitleInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelRemoteSubtitleInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderName

`func (o *ModelRemoteSubtitleInfo) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ModelRemoteSubtitleInfo) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ModelRemoteSubtitleInfo) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ModelRemoteSubtitleInfo) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetName

`func (o *ModelRemoteSubtitleInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelRemoteSubtitleInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelRemoteSubtitleInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelRemoteSubtitleInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFormat

`func (o *ModelRemoteSubtitleInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ModelRemoteSubtitleInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ModelRemoteSubtitleInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ModelRemoteSubtitleInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetAuthor

`func (o *ModelRemoteSubtitleInfo) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelRemoteSubtitleInfo) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelRemoteSubtitleInfo) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelRemoteSubtitleInfo) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetComment

`func (o *ModelRemoteSubtitleInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelRemoteSubtitleInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelRemoteSubtitleInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelRemoteSubtitleInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDateCreated

`func (o *ModelRemoteSubtitleInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelRemoteSubtitleInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelRemoteSubtitleInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelRemoteSubtitleInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ModelRemoteSubtitleInfo) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ModelRemoteSubtitleInfo) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCommunityRating

`func (o *ModelRemoteSubtitleInfo) GetCommunityRating() float32`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *ModelRemoteSubtitleInfo) GetCommunityRatingOk() (*float32, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *ModelRemoteSubtitleInfo) SetCommunityRating(v float32)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *ModelRemoteSubtitleInfo) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *ModelRemoteSubtitleInfo) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *ModelRemoteSubtitleInfo) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetDownloadCount

`func (o *ModelRemoteSubtitleInfo) GetDownloadCount() int32`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ModelRemoteSubtitleInfo) GetDownloadCountOk() (*int32, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ModelRemoteSubtitleInfo) SetDownloadCount(v int32)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *ModelRemoteSubtitleInfo) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### SetDownloadCountNil

`func (o *ModelRemoteSubtitleInfo) SetDownloadCountNil(b bool)`

 SetDownloadCountNil sets the value for DownloadCount to be an explicit nil

### UnsetDownloadCount
`func (o *ModelRemoteSubtitleInfo) UnsetDownloadCount()`

UnsetDownloadCount ensures that no value is present for DownloadCount, not even an explicit nil
### GetIsHashMatch

`func (o *ModelRemoteSubtitleInfo) GetIsHashMatch() bool`

GetIsHashMatch returns the IsHashMatch field if non-nil, zero value otherwise.

### GetIsHashMatchOk

`func (o *ModelRemoteSubtitleInfo) GetIsHashMatchOk() (*bool, bool)`

GetIsHashMatchOk returns a tuple with the IsHashMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHashMatch

`func (o *ModelRemoteSubtitleInfo) SetIsHashMatch(v bool)`

SetIsHashMatch sets IsHashMatch field to given value.

### HasIsHashMatch

`func (o *ModelRemoteSubtitleInfo) HasIsHashMatch() bool`

HasIsHashMatch returns a boolean if a field has been set.

### SetIsHashMatchNil

`func (o *ModelRemoteSubtitleInfo) SetIsHashMatchNil(b bool)`

 SetIsHashMatchNil sets the value for IsHashMatch to be an explicit nil

### UnsetIsHashMatch
`func (o *ModelRemoteSubtitleInfo) UnsetIsHashMatch()`

UnsetIsHashMatch ensures that no value is present for IsHashMatch, not even an explicit nil
### GetIsForced

`func (o *ModelRemoteSubtitleInfo) GetIsForced() bool`

GetIsForced returns the IsForced field if non-nil, zero value otherwise.

### GetIsForcedOk

`func (o *ModelRemoteSubtitleInfo) GetIsForcedOk() (*bool, bool)`

GetIsForcedOk returns a tuple with the IsForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForced

`func (o *ModelRemoteSubtitleInfo) SetIsForced(v bool)`

SetIsForced sets IsForced field to given value.

### HasIsForced

`func (o *ModelRemoteSubtitleInfo) HasIsForced() bool`

HasIsForced returns a boolean if a field has been set.

### SetIsForcedNil

`func (o *ModelRemoteSubtitleInfo) SetIsForcedNil(b bool)`

 SetIsForcedNil sets the value for IsForced to be an explicit nil

### UnsetIsForced
`func (o *ModelRemoteSubtitleInfo) UnsetIsForced()`

UnsetIsForced ensures that no value is present for IsForced, not even an explicit nil
### GetLanguage

`func (o *ModelRemoteSubtitleInfo) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ModelRemoteSubtitleInfo) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ModelRemoteSubtitleInfo) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ModelRemoteSubtitleInfo) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


