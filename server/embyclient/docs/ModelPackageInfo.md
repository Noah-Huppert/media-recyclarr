# ModelPackageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**IsPremium** | Pointer to **bool** |  | [optional] 
**Adult** | Pointer to **bool** |  | [optional] 
**RichDescUrl** | Pointer to **string** |  | [optional] 
**ThumbImage** | Pointer to **string** |  | [optional] 
**PreviewImage** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TargetFilename** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**TileColor** | Pointer to **string** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**TargetSystem** | Pointer to [**ModelModelPackageTargetSystem**](ModelPackageTargetSystem.md) |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**IsRegistered** | Pointer to **bool** |  | [optional] 
**ExpDate** | Pointer to **time.Time** |  | [optional] 
**Versions** | Pointer to [**[]ModelModelPackageVersionInfo**](ModelModelPackageVersionInfo.md) |  | [optional] 
**EnableInAppStore** | Pointer to **bool** |  | [optional] 
**Installs** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelPackageInfo

`func NewModelPackageInfo() *ModelPackageInfo`

NewModelPackageInfo instantiates a new ModelPackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPackageInfoWithDefaults

`func NewModelPackageInfoWithDefaults() *ModelPackageInfo`

NewModelPackageInfoWithDefaults instantiates a new ModelPackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelPackageInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelPackageInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelPackageInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelPackageInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelPackageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelPackageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelPackageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelPackageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortDescription

`func (o *ModelPackageInfo) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *ModelPackageInfo) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *ModelPackageInfo) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *ModelPackageInfo) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetOverview

`func (o *ModelPackageInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelPackageInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelPackageInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelPackageInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetIsPremium

`func (o *ModelPackageInfo) GetIsPremium() bool`

GetIsPremium returns the IsPremium field if non-nil, zero value otherwise.

### GetIsPremiumOk

`func (o *ModelPackageInfo) GetIsPremiumOk() (*bool, bool)`

GetIsPremiumOk returns a tuple with the IsPremium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPremium

`func (o *ModelPackageInfo) SetIsPremium(v bool)`

SetIsPremium sets IsPremium field to given value.

### HasIsPremium

`func (o *ModelPackageInfo) HasIsPremium() bool`

HasIsPremium returns a boolean if a field has been set.

### GetAdult

`func (o *ModelPackageInfo) GetAdult() bool`

GetAdult returns the Adult field if non-nil, zero value otherwise.

### GetAdultOk

`func (o *ModelPackageInfo) GetAdultOk() (*bool, bool)`

GetAdultOk returns a tuple with the Adult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdult

`func (o *ModelPackageInfo) SetAdult(v bool)`

SetAdult sets Adult field to given value.

### HasAdult

`func (o *ModelPackageInfo) HasAdult() bool`

HasAdult returns a boolean if a field has been set.

### GetRichDescUrl

`func (o *ModelPackageInfo) GetRichDescUrl() string`

GetRichDescUrl returns the RichDescUrl field if non-nil, zero value otherwise.

### GetRichDescUrlOk

`func (o *ModelPackageInfo) GetRichDescUrlOk() (*string, bool)`

GetRichDescUrlOk returns a tuple with the RichDescUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichDescUrl

`func (o *ModelPackageInfo) SetRichDescUrl(v string)`

SetRichDescUrl sets RichDescUrl field to given value.

### HasRichDescUrl

`func (o *ModelPackageInfo) HasRichDescUrl() bool`

HasRichDescUrl returns a boolean if a field has been set.

### GetThumbImage

`func (o *ModelPackageInfo) GetThumbImage() string`

GetThumbImage returns the ThumbImage field if non-nil, zero value otherwise.

### GetThumbImageOk

`func (o *ModelPackageInfo) GetThumbImageOk() (*string, bool)`

GetThumbImageOk returns a tuple with the ThumbImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImage

`func (o *ModelPackageInfo) SetThumbImage(v string)`

SetThumbImage sets ThumbImage field to given value.

### HasThumbImage

`func (o *ModelPackageInfo) HasThumbImage() bool`

HasThumbImage returns a boolean if a field has been set.

### GetPreviewImage

`func (o *ModelPackageInfo) GetPreviewImage() string`

GetPreviewImage returns the PreviewImage field if non-nil, zero value otherwise.

### GetPreviewImageOk

`func (o *ModelPackageInfo) GetPreviewImageOk() (*string, bool)`

GetPreviewImageOk returns a tuple with the PreviewImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImage

`func (o *ModelPackageInfo) SetPreviewImage(v string)`

SetPreviewImage sets PreviewImage field to given value.

### HasPreviewImage

`func (o *ModelPackageInfo) HasPreviewImage() bool`

HasPreviewImage returns a boolean if a field has been set.

### GetType

`func (o *ModelPackageInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelPackageInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelPackageInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelPackageInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTargetFilename

`func (o *ModelPackageInfo) GetTargetFilename() string`

GetTargetFilename returns the TargetFilename field if non-nil, zero value otherwise.

### GetTargetFilenameOk

`func (o *ModelPackageInfo) GetTargetFilenameOk() (*string, bool)`

GetTargetFilenameOk returns a tuple with the TargetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilename

`func (o *ModelPackageInfo) SetTargetFilename(v string)`

SetTargetFilename sets TargetFilename field to given value.

### HasTargetFilename

`func (o *ModelPackageInfo) HasTargetFilename() bool`

HasTargetFilename returns a boolean if a field has been set.

### GetOwner

`func (o *ModelPackageInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelPackageInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelPackageInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelPackageInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCategory

`func (o *ModelPackageInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ModelPackageInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ModelPackageInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ModelPackageInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetTileColor

`func (o *ModelPackageInfo) GetTileColor() string`

GetTileColor returns the TileColor field if non-nil, zero value otherwise.

### GetTileColorOk

`func (o *ModelPackageInfo) GetTileColorOk() (*string, bool)`

GetTileColorOk returns a tuple with the TileColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileColor

`func (o *ModelPackageInfo) SetTileColor(v string)`

SetTileColor sets TileColor field to given value.

### HasTileColor

`func (o *ModelPackageInfo) HasTileColor() bool`

HasTileColor returns a boolean if a field has been set.

### GetFeatureId

`func (o *ModelPackageInfo) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *ModelPackageInfo) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *ModelPackageInfo) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *ModelPackageInfo) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetPrice

`func (o *ModelPackageInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ModelPackageInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ModelPackageInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ModelPackageInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ModelPackageInfo) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ModelPackageInfo) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetTargetSystem

`func (o *ModelPackageInfo) GetTargetSystem() ModelModelPackageTargetSystem`

GetTargetSystem returns the TargetSystem field if non-nil, zero value otherwise.

### GetTargetSystemOk

`func (o *ModelPackageInfo) GetTargetSystemOk() (*ModelModelPackageTargetSystem, bool)`

GetTargetSystemOk returns a tuple with the TargetSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystem

`func (o *ModelPackageInfo) SetTargetSystem(v ModelModelPackageTargetSystem)`

SetTargetSystem sets TargetSystem field to given value.

### HasTargetSystem

`func (o *ModelPackageInfo) HasTargetSystem() bool`

HasTargetSystem returns a boolean if a field has been set.

### GetGuid

`func (o *ModelPackageInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ModelPackageInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ModelPackageInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ModelPackageInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetIsRegistered

`func (o *ModelPackageInfo) GetIsRegistered() bool`

GetIsRegistered returns the IsRegistered field if non-nil, zero value otherwise.

### GetIsRegisteredOk

`func (o *ModelPackageInfo) GetIsRegisteredOk() (*bool, bool)`

GetIsRegisteredOk returns a tuple with the IsRegistered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegistered

`func (o *ModelPackageInfo) SetIsRegistered(v bool)`

SetIsRegistered sets IsRegistered field to given value.

### HasIsRegistered

`func (o *ModelPackageInfo) HasIsRegistered() bool`

HasIsRegistered returns a boolean if a field has been set.

### GetExpDate

`func (o *ModelPackageInfo) GetExpDate() time.Time`

GetExpDate returns the ExpDate field if non-nil, zero value otherwise.

### GetExpDateOk

`func (o *ModelPackageInfo) GetExpDateOk() (*time.Time, bool)`

GetExpDateOk returns a tuple with the ExpDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpDate

`func (o *ModelPackageInfo) SetExpDate(v time.Time)`

SetExpDate sets ExpDate field to given value.

### HasExpDate

`func (o *ModelPackageInfo) HasExpDate() bool`

HasExpDate returns a boolean if a field has been set.

### GetVersions

`func (o *ModelPackageInfo) GetVersions() []ModelModelPackageVersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ModelPackageInfo) GetVersionsOk() (*[]ModelModelPackageVersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ModelPackageInfo) SetVersions(v []ModelModelPackageVersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ModelPackageInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetEnableInAppStore

`func (o *ModelPackageInfo) GetEnableInAppStore() bool`

GetEnableInAppStore returns the EnableInAppStore field if non-nil, zero value otherwise.

### GetEnableInAppStoreOk

`func (o *ModelPackageInfo) GetEnableInAppStoreOk() (*bool, bool)`

GetEnableInAppStoreOk returns a tuple with the EnableInAppStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInAppStore

`func (o *ModelPackageInfo) SetEnableInAppStore(v bool)`

SetEnableInAppStore sets EnableInAppStore field to given value.

### HasEnableInAppStore

`func (o *ModelPackageInfo) HasEnableInAppStore() bool`

HasEnableInAppStore returns a boolean if a field has been set.

### GetInstalls

`func (o *ModelPackageInfo) GetInstalls() int32`

GetInstalls returns the Installs field if non-nil, zero value otherwise.

### GetInstallsOk

`func (o *ModelPackageInfo) GetInstallsOk() (*int32, bool)`

GetInstallsOk returns a tuple with the Installs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalls

`func (o *ModelPackageInfo) SetInstalls(v int32)`

SetInstalls sets Installs field to given value.

### HasInstalls

`func (o *ModelPackageInfo) HasInstalls() bool`

HasInstalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


