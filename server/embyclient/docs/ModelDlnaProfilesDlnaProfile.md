# ModelDlnaProfilesDlnaProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ModelModelDlnaProfilesDeviceProfileType**](ModelDlnaProfilesDeviceProfileType.md) |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AlbumArtPn** | Pointer to **string** |  | [optional] 
**MaxAlbumArtWidth** | Pointer to **int32** |  | [optional] 
**MaxAlbumArtHeight** | Pointer to **int32** |  | [optional] 
**MaxIconWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxIconHeight** | Pointer to **NullableInt32** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Manufacturer** | Pointer to **string** |  | [optional] 
**ManufacturerUrl** | Pointer to **string** |  | [optional] 
**ModelName** | Pointer to **string** |  | [optional] 
**ModelDescription** | Pointer to **string** |  | [optional] 
**ModelNumber** | Pointer to **string** |  | [optional] 
**ModelUrl** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**EnableAlbumArtInDidl** | Pointer to **bool** |  | [optional] 
**EnableSingleAlbumArtLimit** | Pointer to **bool** |  | [optional] 
**EnableSingleSubtitleLimit** | Pointer to **bool** |  | [optional] 
**ProtocolInfo** | Pointer to **string** |  | [optional] 
**TimelineOffsetSeconds** | Pointer to **int32** |  | [optional] 
**RequiresPlainVideoItems** | Pointer to **bool** |  | [optional] 
**RequiresPlainFolders** | Pointer to **bool** |  | [optional] 
**IgnoreTranscodeByteRangeRequests** | Pointer to **bool** |  | [optional] 
**SupportsSamsungBookmark** | Pointer to **bool** |  | [optional] 
**Identification** | Pointer to [**ModelModelDlnaProfilesDeviceIdentification**](ModelDlnaProfilesDeviceIdentification.md) |  | [optional] 
**ProtocolInfoDetection** | Pointer to [**ModelModelDlnaProfilesProtocolInfoDetection**](ModelDlnaProfilesProtocolInfoDetection.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SupportedMediaTypes** | Pointer to **string** |  | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt64** |  | [optional] 
**MusicStreamingTranscodingBitrate** | Pointer to **NullableInt32** |  | [optional] 
**MaxStaticMusicBitrate** | Pointer to **NullableInt32** |  | [optional] 
**DirectPlayProfiles** | Pointer to [**[]ModelModelDirectPlayProfile**](ModelModelDirectPlayProfile.md) |  | [optional] 
**TranscodingProfiles** | Pointer to [**[]ModelModelTranscodingProfile**](ModelModelTranscodingProfile.md) |  | [optional] 
**ContainerProfiles** | Pointer to [**[]ModelModelContainerProfile**](ModelModelContainerProfile.md) |  | [optional] 
**CodecProfiles** | Pointer to [**[]ModelModelCodecProfile**](ModelModelCodecProfile.md) |  | [optional] 
**ResponseProfiles** | Pointer to [**[]ModelModelResponseProfile**](ModelModelResponseProfile.md) |  | [optional] 
**SubtitleProfiles** | Pointer to [**[]ModelModelSubtitleProfile**](ModelModelSubtitleProfile.md) |  | [optional] 

## Methods

### NewModelDlnaProfilesDlnaProfile

`func NewModelDlnaProfilesDlnaProfile() *ModelDlnaProfilesDlnaProfile`

NewModelDlnaProfilesDlnaProfile instantiates a new ModelDlnaProfilesDlnaProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDlnaProfilesDlnaProfileWithDefaults

`func NewModelDlnaProfilesDlnaProfileWithDefaults() *ModelDlnaProfilesDlnaProfile`

NewModelDlnaProfilesDlnaProfileWithDefaults instantiates a new ModelDlnaProfilesDlnaProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ModelDlnaProfilesDlnaProfile) GetType() ModelModelDlnaProfilesDeviceProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelDlnaProfilesDlnaProfile) GetTypeOk() (*ModelModelDlnaProfilesDeviceProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelDlnaProfilesDlnaProfile) SetType(v ModelModelDlnaProfilesDeviceProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelDlnaProfilesDlnaProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPath

`func (o *ModelDlnaProfilesDlnaProfile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelDlnaProfilesDlnaProfile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelDlnaProfilesDlnaProfile) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelDlnaProfilesDlnaProfile) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserId

`func (o *ModelDlnaProfilesDlnaProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelDlnaProfilesDlnaProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelDlnaProfilesDlnaProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelDlnaProfilesDlnaProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAlbumArtPn

`func (o *ModelDlnaProfilesDlnaProfile) GetAlbumArtPn() string`

GetAlbumArtPn returns the AlbumArtPn field if non-nil, zero value otherwise.

### GetAlbumArtPnOk

`func (o *ModelDlnaProfilesDlnaProfile) GetAlbumArtPnOk() (*string, bool)`

GetAlbumArtPnOk returns a tuple with the AlbumArtPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtPn

`func (o *ModelDlnaProfilesDlnaProfile) SetAlbumArtPn(v string)`

SetAlbumArtPn sets AlbumArtPn field to given value.

### HasAlbumArtPn

`func (o *ModelDlnaProfilesDlnaProfile) HasAlbumArtPn() bool`

HasAlbumArtPn returns a boolean if a field has been set.

### GetMaxAlbumArtWidth

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxAlbumArtWidth() int32`

GetMaxAlbumArtWidth returns the MaxAlbumArtWidth field if non-nil, zero value otherwise.

### GetMaxAlbumArtWidthOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxAlbumArtWidthOk() (*int32, bool)`

GetMaxAlbumArtWidthOk returns a tuple with the MaxAlbumArtWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtWidth

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxAlbumArtWidth(v int32)`

SetMaxAlbumArtWidth sets MaxAlbumArtWidth field to given value.

### HasMaxAlbumArtWidth

`func (o *ModelDlnaProfilesDlnaProfile) HasMaxAlbumArtWidth() bool`

HasMaxAlbumArtWidth returns a boolean if a field has been set.

### GetMaxAlbumArtHeight

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxAlbumArtHeight() int32`

GetMaxAlbumArtHeight returns the MaxAlbumArtHeight field if non-nil, zero value otherwise.

### GetMaxAlbumArtHeightOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxAlbumArtHeightOk() (*int32, bool)`

GetMaxAlbumArtHeightOk returns a tuple with the MaxAlbumArtHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtHeight

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxAlbumArtHeight(v int32)`

SetMaxAlbumArtHeight sets MaxAlbumArtHeight field to given value.

### HasMaxAlbumArtHeight

`func (o *ModelDlnaProfilesDlnaProfile) HasMaxAlbumArtHeight() bool`

HasMaxAlbumArtHeight returns a boolean if a field has been set.

### GetMaxIconWidth

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxIconWidth() int32`

GetMaxIconWidth returns the MaxIconWidth field if non-nil, zero value otherwise.

### GetMaxIconWidthOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxIconWidthOk() (*int32, bool)`

GetMaxIconWidthOk returns a tuple with the MaxIconWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconWidth

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxIconWidth(v int32)`

SetMaxIconWidth sets MaxIconWidth field to given value.

### HasMaxIconWidth

`func (o *ModelDlnaProfilesDlnaProfile) HasMaxIconWidth() bool`

HasMaxIconWidth returns a boolean if a field has been set.

### SetMaxIconWidthNil

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxIconWidthNil(b bool)`

 SetMaxIconWidthNil sets the value for MaxIconWidth to be an explicit nil

### UnsetMaxIconWidth
`func (o *ModelDlnaProfilesDlnaProfile) UnsetMaxIconWidth()`

UnsetMaxIconWidth ensures that no value is present for MaxIconWidth, not even an explicit nil
### GetMaxIconHeight

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxIconHeight() int32`

GetMaxIconHeight returns the MaxIconHeight field if non-nil, zero value otherwise.

### GetMaxIconHeightOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxIconHeightOk() (*int32, bool)`

GetMaxIconHeightOk returns a tuple with the MaxIconHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconHeight

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxIconHeight(v int32)`

SetMaxIconHeight sets MaxIconHeight field to given value.

### HasMaxIconHeight

`func (o *ModelDlnaProfilesDlnaProfile) HasMaxIconHeight() bool`

HasMaxIconHeight returns a boolean if a field has been set.

### SetMaxIconHeightNil

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxIconHeightNil(b bool)`

 SetMaxIconHeightNil sets the value for MaxIconHeight to be an explicit nil

### UnsetMaxIconHeight
`func (o *ModelDlnaProfilesDlnaProfile) UnsetMaxIconHeight()`

UnsetMaxIconHeight ensures that no value is present for MaxIconHeight, not even an explicit nil
### GetFriendlyName

`func (o *ModelDlnaProfilesDlnaProfile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ModelDlnaProfilesDlnaProfile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ModelDlnaProfilesDlnaProfile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ModelDlnaProfilesDlnaProfile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetManufacturer

`func (o *ModelDlnaProfilesDlnaProfile) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ModelDlnaProfilesDlnaProfile) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ModelDlnaProfilesDlnaProfile) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ModelDlnaProfilesDlnaProfile) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerUrl

`func (o *ModelDlnaProfilesDlnaProfile) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *ModelDlnaProfilesDlnaProfile) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *ModelDlnaProfilesDlnaProfile) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *ModelDlnaProfilesDlnaProfile) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### GetModelName

`func (o *ModelDlnaProfilesDlnaProfile) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModelDlnaProfilesDlnaProfile) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModelDlnaProfilesDlnaProfile) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ModelDlnaProfilesDlnaProfile) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelDescription

`func (o *ModelDlnaProfilesDlnaProfile) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *ModelDlnaProfilesDlnaProfile) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *ModelDlnaProfilesDlnaProfile) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *ModelDlnaProfilesDlnaProfile) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### GetModelNumber

`func (o *ModelDlnaProfilesDlnaProfile) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *ModelDlnaProfilesDlnaProfile) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *ModelDlnaProfilesDlnaProfile) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *ModelDlnaProfilesDlnaProfile) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetModelUrl

`func (o *ModelDlnaProfilesDlnaProfile) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *ModelDlnaProfilesDlnaProfile) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *ModelDlnaProfilesDlnaProfile) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *ModelDlnaProfilesDlnaProfile) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ModelDlnaProfilesDlnaProfile) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelDlnaProfilesDlnaProfile) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelDlnaProfilesDlnaProfile) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelDlnaProfilesDlnaProfile) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetEnableAlbumArtInDidl

`func (o *ModelDlnaProfilesDlnaProfile) GetEnableAlbumArtInDidl() bool`

GetEnableAlbumArtInDidl returns the EnableAlbumArtInDidl field if non-nil, zero value otherwise.

### GetEnableAlbumArtInDidlOk

`func (o *ModelDlnaProfilesDlnaProfile) GetEnableAlbumArtInDidlOk() (*bool, bool)`

GetEnableAlbumArtInDidlOk returns a tuple with the EnableAlbumArtInDidl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAlbumArtInDidl

`func (o *ModelDlnaProfilesDlnaProfile) SetEnableAlbumArtInDidl(v bool)`

SetEnableAlbumArtInDidl sets EnableAlbumArtInDidl field to given value.

### HasEnableAlbumArtInDidl

`func (o *ModelDlnaProfilesDlnaProfile) HasEnableAlbumArtInDidl() bool`

HasEnableAlbumArtInDidl returns a boolean if a field has been set.

### GetEnableSingleAlbumArtLimit

`func (o *ModelDlnaProfilesDlnaProfile) GetEnableSingleAlbumArtLimit() bool`

GetEnableSingleAlbumArtLimit returns the EnableSingleAlbumArtLimit field if non-nil, zero value otherwise.

### GetEnableSingleAlbumArtLimitOk

`func (o *ModelDlnaProfilesDlnaProfile) GetEnableSingleAlbumArtLimitOk() (*bool, bool)`

GetEnableSingleAlbumArtLimitOk returns a tuple with the EnableSingleAlbumArtLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleAlbumArtLimit

`func (o *ModelDlnaProfilesDlnaProfile) SetEnableSingleAlbumArtLimit(v bool)`

SetEnableSingleAlbumArtLimit sets EnableSingleAlbumArtLimit field to given value.

### HasEnableSingleAlbumArtLimit

`func (o *ModelDlnaProfilesDlnaProfile) HasEnableSingleAlbumArtLimit() bool`

HasEnableSingleAlbumArtLimit returns a boolean if a field has been set.

### GetEnableSingleSubtitleLimit

`func (o *ModelDlnaProfilesDlnaProfile) GetEnableSingleSubtitleLimit() bool`

GetEnableSingleSubtitleLimit returns the EnableSingleSubtitleLimit field if non-nil, zero value otherwise.

### GetEnableSingleSubtitleLimitOk

`func (o *ModelDlnaProfilesDlnaProfile) GetEnableSingleSubtitleLimitOk() (*bool, bool)`

GetEnableSingleSubtitleLimitOk returns a tuple with the EnableSingleSubtitleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSubtitleLimit

`func (o *ModelDlnaProfilesDlnaProfile) SetEnableSingleSubtitleLimit(v bool)`

SetEnableSingleSubtitleLimit sets EnableSingleSubtitleLimit field to given value.

### HasEnableSingleSubtitleLimit

`func (o *ModelDlnaProfilesDlnaProfile) HasEnableSingleSubtitleLimit() bool`

HasEnableSingleSubtitleLimit returns a boolean if a field has been set.

### GetProtocolInfo

`func (o *ModelDlnaProfilesDlnaProfile) GetProtocolInfo() string`

GetProtocolInfo returns the ProtocolInfo field if non-nil, zero value otherwise.

### GetProtocolInfoOk

`func (o *ModelDlnaProfilesDlnaProfile) GetProtocolInfoOk() (*string, bool)`

GetProtocolInfoOk returns a tuple with the ProtocolInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfo

`func (o *ModelDlnaProfilesDlnaProfile) SetProtocolInfo(v string)`

SetProtocolInfo sets ProtocolInfo field to given value.

### HasProtocolInfo

`func (o *ModelDlnaProfilesDlnaProfile) HasProtocolInfo() bool`

HasProtocolInfo returns a boolean if a field has been set.

### GetTimelineOffsetSeconds

`func (o *ModelDlnaProfilesDlnaProfile) GetTimelineOffsetSeconds() int32`

GetTimelineOffsetSeconds returns the TimelineOffsetSeconds field if non-nil, zero value otherwise.

### GetTimelineOffsetSecondsOk

`func (o *ModelDlnaProfilesDlnaProfile) GetTimelineOffsetSecondsOk() (*int32, bool)`

GetTimelineOffsetSecondsOk returns a tuple with the TimelineOffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineOffsetSeconds

`func (o *ModelDlnaProfilesDlnaProfile) SetTimelineOffsetSeconds(v int32)`

SetTimelineOffsetSeconds sets TimelineOffsetSeconds field to given value.

### HasTimelineOffsetSeconds

`func (o *ModelDlnaProfilesDlnaProfile) HasTimelineOffsetSeconds() bool`

HasTimelineOffsetSeconds returns a boolean if a field has been set.

### GetRequiresPlainVideoItems

`func (o *ModelDlnaProfilesDlnaProfile) GetRequiresPlainVideoItems() bool`

GetRequiresPlainVideoItems returns the RequiresPlainVideoItems field if non-nil, zero value otherwise.

### GetRequiresPlainVideoItemsOk

`func (o *ModelDlnaProfilesDlnaProfile) GetRequiresPlainVideoItemsOk() (*bool, bool)`

GetRequiresPlainVideoItemsOk returns a tuple with the RequiresPlainVideoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainVideoItems

`func (o *ModelDlnaProfilesDlnaProfile) SetRequiresPlainVideoItems(v bool)`

SetRequiresPlainVideoItems sets RequiresPlainVideoItems field to given value.

### HasRequiresPlainVideoItems

`func (o *ModelDlnaProfilesDlnaProfile) HasRequiresPlainVideoItems() bool`

HasRequiresPlainVideoItems returns a boolean if a field has been set.

### GetRequiresPlainFolders

`func (o *ModelDlnaProfilesDlnaProfile) GetRequiresPlainFolders() bool`

GetRequiresPlainFolders returns the RequiresPlainFolders field if non-nil, zero value otherwise.

### GetRequiresPlainFoldersOk

`func (o *ModelDlnaProfilesDlnaProfile) GetRequiresPlainFoldersOk() (*bool, bool)`

GetRequiresPlainFoldersOk returns a tuple with the RequiresPlainFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainFolders

`func (o *ModelDlnaProfilesDlnaProfile) SetRequiresPlainFolders(v bool)`

SetRequiresPlainFolders sets RequiresPlainFolders field to given value.

### HasRequiresPlainFolders

`func (o *ModelDlnaProfilesDlnaProfile) HasRequiresPlainFolders() bool`

HasRequiresPlainFolders returns a boolean if a field has been set.

### GetIgnoreTranscodeByteRangeRequests

`func (o *ModelDlnaProfilesDlnaProfile) GetIgnoreTranscodeByteRangeRequests() bool`

GetIgnoreTranscodeByteRangeRequests returns the IgnoreTranscodeByteRangeRequests field if non-nil, zero value otherwise.

### GetIgnoreTranscodeByteRangeRequestsOk

`func (o *ModelDlnaProfilesDlnaProfile) GetIgnoreTranscodeByteRangeRequestsOk() (*bool, bool)`

GetIgnoreTranscodeByteRangeRequestsOk returns a tuple with the IgnoreTranscodeByteRangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTranscodeByteRangeRequests

`func (o *ModelDlnaProfilesDlnaProfile) SetIgnoreTranscodeByteRangeRequests(v bool)`

SetIgnoreTranscodeByteRangeRequests sets IgnoreTranscodeByteRangeRequests field to given value.

### HasIgnoreTranscodeByteRangeRequests

`func (o *ModelDlnaProfilesDlnaProfile) HasIgnoreTranscodeByteRangeRequests() bool`

HasIgnoreTranscodeByteRangeRequests returns a boolean if a field has been set.

### GetSupportsSamsungBookmark

`func (o *ModelDlnaProfilesDlnaProfile) GetSupportsSamsungBookmark() bool`

GetSupportsSamsungBookmark returns the SupportsSamsungBookmark field if non-nil, zero value otherwise.

### GetSupportsSamsungBookmarkOk

`func (o *ModelDlnaProfilesDlnaProfile) GetSupportsSamsungBookmarkOk() (*bool, bool)`

GetSupportsSamsungBookmarkOk returns a tuple with the SupportsSamsungBookmark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSamsungBookmark

`func (o *ModelDlnaProfilesDlnaProfile) SetSupportsSamsungBookmark(v bool)`

SetSupportsSamsungBookmark sets SupportsSamsungBookmark field to given value.

### HasSupportsSamsungBookmark

`func (o *ModelDlnaProfilesDlnaProfile) HasSupportsSamsungBookmark() bool`

HasSupportsSamsungBookmark returns a boolean if a field has been set.

### GetIdentification

`func (o *ModelDlnaProfilesDlnaProfile) GetIdentification() ModelModelDlnaProfilesDeviceIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *ModelDlnaProfilesDlnaProfile) GetIdentificationOk() (*ModelModelDlnaProfilesDeviceIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *ModelDlnaProfilesDlnaProfile) SetIdentification(v ModelModelDlnaProfilesDeviceIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *ModelDlnaProfilesDlnaProfile) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetProtocolInfoDetection

`func (o *ModelDlnaProfilesDlnaProfile) GetProtocolInfoDetection() ModelModelDlnaProfilesProtocolInfoDetection`

GetProtocolInfoDetection returns the ProtocolInfoDetection field if non-nil, zero value otherwise.

### GetProtocolInfoDetectionOk

`func (o *ModelDlnaProfilesDlnaProfile) GetProtocolInfoDetectionOk() (*ModelModelDlnaProfilesProtocolInfoDetection, bool)`

GetProtocolInfoDetectionOk returns a tuple with the ProtocolInfoDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfoDetection

`func (o *ModelDlnaProfilesDlnaProfile) SetProtocolInfoDetection(v ModelModelDlnaProfilesProtocolInfoDetection)`

SetProtocolInfoDetection sets ProtocolInfoDetection field to given value.

### HasProtocolInfoDetection

`func (o *ModelDlnaProfilesDlnaProfile) HasProtocolInfoDetection() bool`

HasProtocolInfoDetection returns a boolean if a field has been set.

### GetName

`func (o *ModelDlnaProfilesDlnaProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelDlnaProfilesDlnaProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelDlnaProfilesDlnaProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelDlnaProfilesDlnaProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelDlnaProfilesDlnaProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDlnaProfilesDlnaProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDlnaProfilesDlnaProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDlnaProfilesDlnaProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSupportedMediaTypes

`func (o *ModelDlnaProfilesDlnaProfile) GetSupportedMediaTypes() string`

GetSupportedMediaTypes returns the SupportedMediaTypes field if non-nil, zero value otherwise.

### GetSupportedMediaTypesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetSupportedMediaTypesOk() (*string, bool)`

GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMediaTypes

`func (o *ModelDlnaProfilesDlnaProfile) SetSupportedMediaTypes(v string)`

SetSupportedMediaTypes sets SupportedMediaTypes field to given value.

### HasSupportedMediaTypes

`func (o *ModelDlnaProfilesDlnaProfile) HasSupportedMediaTypes() bool`

HasSupportedMediaTypes returns a boolean if a field has been set.

### GetMaxStreamingBitrate

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *ModelDlnaProfilesDlnaProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *ModelDlnaProfilesDlnaProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *ModelDlnaProfilesDlnaProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *ModelDlnaProfilesDlnaProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *ModelDlnaProfilesDlnaProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *ModelDlnaProfilesDlnaProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *ModelDlnaProfilesDlnaProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *ModelDlnaProfilesDlnaProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *ModelDlnaProfilesDlnaProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *ModelDlnaProfilesDlnaProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *ModelDlnaProfilesDlnaProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetDirectPlayProfiles

`func (o *ModelDlnaProfilesDlnaProfile) GetDirectPlayProfiles() []ModelModelDirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetDirectPlayProfilesOk() (*[]ModelModelDirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *ModelDlnaProfilesDlnaProfile) SetDirectPlayProfiles(v []ModelModelDirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *ModelDlnaProfilesDlnaProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *ModelDlnaProfilesDlnaProfile) GetTranscodingProfiles() []ModelModelTranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetTranscodingProfilesOk() (*[]ModelModelTranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *ModelDlnaProfilesDlnaProfile) SetTranscodingProfiles(v []ModelModelTranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *ModelDlnaProfilesDlnaProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *ModelDlnaProfilesDlnaProfile) GetContainerProfiles() []ModelModelContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetContainerProfilesOk() (*[]ModelModelContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *ModelDlnaProfilesDlnaProfile) SetContainerProfiles(v []ModelModelContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *ModelDlnaProfilesDlnaProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *ModelDlnaProfilesDlnaProfile) GetCodecProfiles() []ModelModelCodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetCodecProfilesOk() (*[]ModelModelCodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *ModelDlnaProfilesDlnaProfile) SetCodecProfiles(v []ModelModelCodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *ModelDlnaProfilesDlnaProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetResponseProfiles

`func (o *ModelDlnaProfilesDlnaProfile) GetResponseProfiles() []ModelModelResponseProfile`

GetResponseProfiles returns the ResponseProfiles field if non-nil, zero value otherwise.

### GetResponseProfilesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetResponseProfilesOk() (*[]ModelModelResponseProfile, bool)`

GetResponseProfilesOk returns a tuple with the ResponseProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfiles

`func (o *ModelDlnaProfilesDlnaProfile) SetResponseProfiles(v []ModelModelResponseProfile)`

SetResponseProfiles sets ResponseProfiles field to given value.

### HasResponseProfiles

`func (o *ModelDlnaProfilesDlnaProfile) HasResponseProfiles() bool`

HasResponseProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *ModelDlnaProfilesDlnaProfile) GetSubtitleProfiles() []ModelModelSubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *ModelDlnaProfilesDlnaProfile) GetSubtitleProfilesOk() (*[]ModelModelSubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *ModelDlnaProfilesDlnaProfile) SetSubtitleProfiles(v []ModelModelSubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *ModelDlnaProfilesDlnaProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


