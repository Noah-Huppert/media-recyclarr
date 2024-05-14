# ModelVideoCodecBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodecDeviceInfo** | Pointer to [**ModelModelCommonInterfacesICodecDeviceInfo**](ModelCommonInterfacesICodecDeviceInfo.md) |  | [optional] 
**CodecKind** | Pointer to [**ModelModelCodecKinds**](ModelCodecKinds.md) |  | [optional] 
**MediaTypeName** | Pointer to **string** |  | [optional] 
**VideoMediaType** | Pointer to [**ModelModelVideoMediaTypes**](ModelVideoMediaTypes.md) |  | [optional] 
**MinWidth** | Pointer to **NullableInt32** |  | [optional] 
**MaxWidth** | Pointer to **NullableInt32** |  | [optional] 
**MinHeight** | Pointer to **NullableInt32** |  | [optional] 
**MaxHeight** | Pointer to **NullableInt32** |  | [optional] 
**WidthAlignment** | Pointer to **NullableInt32** |  | [optional] 
**HeightAlignment** | Pointer to **NullableInt32** |  | [optional] 
**MaxBitRate** | Pointer to [**ModelModelBitRate**](ModelBitRate.md) |  | [optional] 
**SupportedColorFormats** | Pointer to [**[]ModelModelColorFormats**](ModelModelColorFormats.md) |  | [optional] 
**SupportedColorFormatStrings** | Pointer to **[]string** |  | [optional] 
**ProfileAndLevelInformation** | Pointer to [**[]ModelModelProfileLevelInformation**](ModelModelProfileLevelInformation.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**ModelModelCodecDirections**](ModelCodecDirections.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FrameworkCodec** | Pointer to **string** |  | [optional] 
**IsHardwareCodec** | Pointer to **bool** |  | [optional] 
**SecondaryFramework** | Pointer to [**ModelModelSecondaryFrameworks**](ModelSecondaryFrameworks.md) |  | [optional] 
**SecondaryFrameworkCodec** | Pointer to **string** |  | [optional] 
**MaxInstanceCount** | Pointer to **NullableInt32** |  | [optional] 
**IsEnabledByDefault** | Pointer to **bool** |  | [optional] 
**DefaultPriority** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelVideoCodecBase

`func NewModelVideoCodecBase() *ModelVideoCodecBase`

NewModelVideoCodecBase instantiates a new ModelVideoCodecBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVideoCodecBaseWithDefaults

`func NewModelVideoCodecBaseWithDefaults() *ModelVideoCodecBase`

NewModelVideoCodecBaseWithDefaults instantiates a new ModelVideoCodecBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodecDeviceInfo

`func (o *ModelVideoCodecBase) GetCodecDeviceInfo() ModelModelCommonInterfacesICodecDeviceInfo`

GetCodecDeviceInfo returns the CodecDeviceInfo field if non-nil, zero value otherwise.

### GetCodecDeviceInfoOk

`func (o *ModelVideoCodecBase) GetCodecDeviceInfoOk() (*ModelModelCommonInterfacesICodecDeviceInfo, bool)`

GetCodecDeviceInfoOk returns a tuple with the CodecDeviceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecDeviceInfo

`func (o *ModelVideoCodecBase) SetCodecDeviceInfo(v ModelModelCommonInterfacesICodecDeviceInfo)`

SetCodecDeviceInfo sets CodecDeviceInfo field to given value.

### HasCodecDeviceInfo

`func (o *ModelVideoCodecBase) HasCodecDeviceInfo() bool`

HasCodecDeviceInfo returns a boolean if a field has been set.

### GetCodecKind

`func (o *ModelVideoCodecBase) GetCodecKind() ModelModelCodecKinds`

GetCodecKind returns the CodecKind field if non-nil, zero value otherwise.

### GetCodecKindOk

`func (o *ModelVideoCodecBase) GetCodecKindOk() (*ModelModelCodecKinds, bool)`

GetCodecKindOk returns a tuple with the CodecKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecKind

`func (o *ModelVideoCodecBase) SetCodecKind(v ModelModelCodecKinds)`

SetCodecKind sets CodecKind field to given value.

### HasCodecKind

`func (o *ModelVideoCodecBase) HasCodecKind() bool`

HasCodecKind returns a boolean if a field has been set.

### GetMediaTypeName

`func (o *ModelVideoCodecBase) GetMediaTypeName() string`

GetMediaTypeName returns the MediaTypeName field if non-nil, zero value otherwise.

### GetMediaTypeNameOk

`func (o *ModelVideoCodecBase) GetMediaTypeNameOk() (*string, bool)`

GetMediaTypeNameOk returns a tuple with the MediaTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypeName

`func (o *ModelVideoCodecBase) SetMediaTypeName(v string)`

SetMediaTypeName sets MediaTypeName field to given value.

### HasMediaTypeName

`func (o *ModelVideoCodecBase) HasMediaTypeName() bool`

HasMediaTypeName returns a boolean if a field has been set.

### GetVideoMediaType

`func (o *ModelVideoCodecBase) GetVideoMediaType() ModelModelVideoMediaTypes`

GetVideoMediaType returns the VideoMediaType field if non-nil, zero value otherwise.

### GetVideoMediaTypeOk

`func (o *ModelVideoCodecBase) GetVideoMediaTypeOk() (*ModelModelVideoMediaTypes, bool)`

GetVideoMediaTypeOk returns a tuple with the VideoMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoMediaType

`func (o *ModelVideoCodecBase) SetVideoMediaType(v ModelModelVideoMediaTypes)`

SetVideoMediaType sets VideoMediaType field to given value.

### HasVideoMediaType

`func (o *ModelVideoCodecBase) HasVideoMediaType() bool`

HasVideoMediaType returns a boolean if a field has been set.

### GetMinWidth

`func (o *ModelVideoCodecBase) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *ModelVideoCodecBase) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *ModelVideoCodecBase) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *ModelVideoCodecBase) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### SetMinWidthNil

`func (o *ModelVideoCodecBase) SetMinWidthNil(b bool)`

 SetMinWidthNil sets the value for MinWidth to be an explicit nil

### UnsetMinWidth
`func (o *ModelVideoCodecBase) UnsetMinWidth()`

UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
### GetMaxWidth

`func (o *ModelVideoCodecBase) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *ModelVideoCodecBase) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *ModelVideoCodecBase) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *ModelVideoCodecBase) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *ModelVideoCodecBase) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *ModelVideoCodecBase) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMinHeight

`func (o *ModelVideoCodecBase) GetMinHeight() int32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *ModelVideoCodecBase) GetMinHeightOk() (*int32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *ModelVideoCodecBase) SetMinHeight(v int32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *ModelVideoCodecBase) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### SetMinHeightNil

`func (o *ModelVideoCodecBase) SetMinHeightNil(b bool)`

 SetMinHeightNil sets the value for MinHeight to be an explicit nil

### UnsetMinHeight
`func (o *ModelVideoCodecBase) UnsetMinHeight()`

UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
### GetMaxHeight

`func (o *ModelVideoCodecBase) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ModelVideoCodecBase) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ModelVideoCodecBase) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ModelVideoCodecBase) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *ModelVideoCodecBase) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *ModelVideoCodecBase) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetWidthAlignment

`func (o *ModelVideoCodecBase) GetWidthAlignment() int32`

GetWidthAlignment returns the WidthAlignment field if non-nil, zero value otherwise.

### GetWidthAlignmentOk

`func (o *ModelVideoCodecBase) GetWidthAlignmentOk() (*int32, bool)`

GetWidthAlignmentOk returns a tuple with the WidthAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthAlignment

`func (o *ModelVideoCodecBase) SetWidthAlignment(v int32)`

SetWidthAlignment sets WidthAlignment field to given value.

### HasWidthAlignment

`func (o *ModelVideoCodecBase) HasWidthAlignment() bool`

HasWidthAlignment returns a boolean if a field has been set.

### SetWidthAlignmentNil

`func (o *ModelVideoCodecBase) SetWidthAlignmentNil(b bool)`

 SetWidthAlignmentNil sets the value for WidthAlignment to be an explicit nil

### UnsetWidthAlignment
`func (o *ModelVideoCodecBase) UnsetWidthAlignment()`

UnsetWidthAlignment ensures that no value is present for WidthAlignment, not even an explicit nil
### GetHeightAlignment

`func (o *ModelVideoCodecBase) GetHeightAlignment() int32`

GetHeightAlignment returns the HeightAlignment field if non-nil, zero value otherwise.

### GetHeightAlignmentOk

`func (o *ModelVideoCodecBase) GetHeightAlignmentOk() (*int32, bool)`

GetHeightAlignmentOk returns a tuple with the HeightAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightAlignment

`func (o *ModelVideoCodecBase) SetHeightAlignment(v int32)`

SetHeightAlignment sets HeightAlignment field to given value.

### HasHeightAlignment

`func (o *ModelVideoCodecBase) HasHeightAlignment() bool`

HasHeightAlignment returns a boolean if a field has been set.

### SetHeightAlignmentNil

`func (o *ModelVideoCodecBase) SetHeightAlignmentNil(b bool)`

 SetHeightAlignmentNil sets the value for HeightAlignment to be an explicit nil

### UnsetHeightAlignment
`func (o *ModelVideoCodecBase) UnsetHeightAlignment()`

UnsetHeightAlignment ensures that no value is present for HeightAlignment, not even an explicit nil
### GetMaxBitRate

`func (o *ModelVideoCodecBase) GetMaxBitRate() ModelModelBitRate`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *ModelVideoCodecBase) GetMaxBitRateOk() (*ModelModelBitRate, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *ModelVideoCodecBase) SetMaxBitRate(v ModelModelBitRate)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *ModelVideoCodecBase) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetSupportedColorFormats

`func (o *ModelVideoCodecBase) GetSupportedColorFormats() []ModelModelColorFormats`

GetSupportedColorFormats returns the SupportedColorFormats field if non-nil, zero value otherwise.

### GetSupportedColorFormatsOk

`func (o *ModelVideoCodecBase) GetSupportedColorFormatsOk() (*[]ModelModelColorFormats, bool)`

GetSupportedColorFormatsOk returns a tuple with the SupportedColorFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedColorFormats

`func (o *ModelVideoCodecBase) SetSupportedColorFormats(v []ModelModelColorFormats)`

SetSupportedColorFormats sets SupportedColorFormats field to given value.

### HasSupportedColorFormats

`func (o *ModelVideoCodecBase) HasSupportedColorFormats() bool`

HasSupportedColorFormats returns a boolean if a field has been set.

### GetSupportedColorFormatStrings

`func (o *ModelVideoCodecBase) GetSupportedColorFormatStrings() []string`

GetSupportedColorFormatStrings returns the SupportedColorFormatStrings field if non-nil, zero value otherwise.

### GetSupportedColorFormatStringsOk

`func (o *ModelVideoCodecBase) GetSupportedColorFormatStringsOk() (*[]string, bool)`

GetSupportedColorFormatStringsOk returns a tuple with the SupportedColorFormatStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedColorFormatStrings

`func (o *ModelVideoCodecBase) SetSupportedColorFormatStrings(v []string)`

SetSupportedColorFormatStrings sets SupportedColorFormatStrings field to given value.

### HasSupportedColorFormatStrings

`func (o *ModelVideoCodecBase) HasSupportedColorFormatStrings() bool`

HasSupportedColorFormatStrings returns a boolean if a field has been set.

### GetProfileAndLevelInformation

`func (o *ModelVideoCodecBase) GetProfileAndLevelInformation() []ModelModelProfileLevelInformation`

GetProfileAndLevelInformation returns the ProfileAndLevelInformation field if non-nil, zero value otherwise.

### GetProfileAndLevelInformationOk

`func (o *ModelVideoCodecBase) GetProfileAndLevelInformationOk() (*[]ModelModelProfileLevelInformation, bool)`

GetProfileAndLevelInformationOk returns a tuple with the ProfileAndLevelInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAndLevelInformation

`func (o *ModelVideoCodecBase) SetProfileAndLevelInformation(v []ModelModelProfileLevelInformation)`

SetProfileAndLevelInformation sets ProfileAndLevelInformation field to given value.

### HasProfileAndLevelInformation

`func (o *ModelVideoCodecBase) HasProfileAndLevelInformation() bool`

HasProfileAndLevelInformation returns a boolean if a field has been set.

### GetId

`func (o *ModelVideoCodecBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVideoCodecBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVideoCodecBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVideoCodecBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *ModelVideoCodecBase) GetDirection() ModelModelCodecDirections`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ModelVideoCodecBase) GetDirectionOk() (*ModelModelCodecDirections, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ModelVideoCodecBase) SetDirection(v ModelModelCodecDirections)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ModelVideoCodecBase) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetName

`func (o *ModelVideoCodecBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelVideoCodecBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelVideoCodecBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelVideoCodecBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ModelVideoCodecBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelVideoCodecBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelVideoCodecBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelVideoCodecBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrameworkCodec

`func (o *ModelVideoCodecBase) GetFrameworkCodec() string`

GetFrameworkCodec returns the FrameworkCodec field if non-nil, zero value otherwise.

### GetFrameworkCodecOk

`func (o *ModelVideoCodecBase) GetFrameworkCodecOk() (*string, bool)`

GetFrameworkCodecOk returns a tuple with the FrameworkCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameworkCodec

`func (o *ModelVideoCodecBase) SetFrameworkCodec(v string)`

SetFrameworkCodec sets FrameworkCodec field to given value.

### HasFrameworkCodec

`func (o *ModelVideoCodecBase) HasFrameworkCodec() bool`

HasFrameworkCodec returns a boolean if a field has been set.

### GetIsHardwareCodec

`func (o *ModelVideoCodecBase) GetIsHardwareCodec() bool`

GetIsHardwareCodec returns the IsHardwareCodec field if non-nil, zero value otherwise.

### GetIsHardwareCodecOk

`func (o *ModelVideoCodecBase) GetIsHardwareCodecOk() (*bool, bool)`

GetIsHardwareCodecOk returns a tuple with the IsHardwareCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardwareCodec

`func (o *ModelVideoCodecBase) SetIsHardwareCodec(v bool)`

SetIsHardwareCodec sets IsHardwareCodec field to given value.

### HasIsHardwareCodec

`func (o *ModelVideoCodecBase) HasIsHardwareCodec() bool`

HasIsHardwareCodec returns a boolean if a field has been set.

### GetSecondaryFramework

`func (o *ModelVideoCodecBase) GetSecondaryFramework() ModelModelSecondaryFrameworks`

GetSecondaryFramework returns the SecondaryFramework field if non-nil, zero value otherwise.

### GetSecondaryFrameworkOk

`func (o *ModelVideoCodecBase) GetSecondaryFrameworkOk() (*ModelModelSecondaryFrameworks, bool)`

GetSecondaryFrameworkOk returns a tuple with the SecondaryFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFramework

`func (o *ModelVideoCodecBase) SetSecondaryFramework(v ModelModelSecondaryFrameworks)`

SetSecondaryFramework sets SecondaryFramework field to given value.

### HasSecondaryFramework

`func (o *ModelVideoCodecBase) HasSecondaryFramework() bool`

HasSecondaryFramework returns a boolean if a field has been set.

### GetSecondaryFrameworkCodec

`func (o *ModelVideoCodecBase) GetSecondaryFrameworkCodec() string`

GetSecondaryFrameworkCodec returns the SecondaryFrameworkCodec field if non-nil, zero value otherwise.

### GetSecondaryFrameworkCodecOk

`func (o *ModelVideoCodecBase) GetSecondaryFrameworkCodecOk() (*string, bool)`

GetSecondaryFrameworkCodecOk returns a tuple with the SecondaryFrameworkCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryFrameworkCodec

`func (o *ModelVideoCodecBase) SetSecondaryFrameworkCodec(v string)`

SetSecondaryFrameworkCodec sets SecondaryFrameworkCodec field to given value.

### HasSecondaryFrameworkCodec

`func (o *ModelVideoCodecBase) HasSecondaryFrameworkCodec() bool`

HasSecondaryFrameworkCodec returns a boolean if a field has been set.

### GetMaxInstanceCount

`func (o *ModelVideoCodecBase) GetMaxInstanceCount() int32`

GetMaxInstanceCount returns the MaxInstanceCount field if non-nil, zero value otherwise.

### GetMaxInstanceCountOk

`func (o *ModelVideoCodecBase) GetMaxInstanceCountOk() (*int32, bool)`

GetMaxInstanceCountOk returns a tuple with the MaxInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInstanceCount

`func (o *ModelVideoCodecBase) SetMaxInstanceCount(v int32)`

SetMaxInstanceCount sets MaxInstanceCount field to given value.

### HasMaxInstanceCount

`func (o *ModelVideoCodecBase) HasMaxInstanceCount() bool`

HasMaxInstanceCount returns a boolean if a field has been set.

### SetMaxInstanceCountNil

`func (o *ModelVideoCodecBase) SetMaxInstanceCountNil(b bool)`

 SetMaxInstanceCountNil sets the value for MaxInstanceCount to be an explicit nil

### UnsetMaxInstanceCount
`func (o *ModelVideoCodecBase) UnsetMaxInstanceCount()`

UnsetMaxInstanceCount ensures that no value is present for MaxInstanceCount, not even an explicit nil
### GetIsEnabledByDefault

`func (o *ModelVideoCodecBase) GetIsEnabledByDefault() bool`

GetIsEnabledByDefault returns the IsEnabledByDefault field if non-nil, zero value otherwise.

### GetIsEnabledByDefaultOk

`func (o *ModelVideoCodecBase) GetIsEnabledByDefaultOk() (*bool, bool)`

GetIsEnabledByDefaultOk returns a tuple with the IsEnabledByDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabledByDefault

`func (o *ModelVideoCodecBase) SetIsEnabledByDefault(v bool)`

SetIsEnabledByDefault sets IsEnabledByDefault field to given value.

### HasIsEnabledByDefault

`func (o *ModelVideoCodecBase) HasIsEnabledByDefault() bool`

HasIsEnabledByDefault returns a boolean if a field has been set.

### GetDefaultPriority

`func (o *ModelVideoCodecBase) GetDefaultPriority() int32`

GetDefaultPriority returns the DefaultPriority field if non-nil, zero value otherwise.

### GetDefaultPriorityOk

`func (o *ModelVideoCodecBase) GetDefaultPriorityOk() (*int32, bool)`

GetDefaultPriorityOk returns a tuple with the DefaultPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPriority

`func (o *ModelVideoCodecBase) SetDefaultPriority(v int32)`

SetDefaultPriority sets DefaultPriority field to given value.

### HasDefaultPriority

`func (o *ModelVideoCodecBase) HasDefaultPriority() bool`

HasDefaultPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


