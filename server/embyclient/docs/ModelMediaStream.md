# ModelMediaStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codec** | Pointer to **string** |  | [optional] 
**CodecTag** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**ColorTransfer** | Pointer to **string** |  | [optional] 
**ColorPrimaries** | Pointer to **string** |  | [optional] 
**ColorSpace** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**StreamStartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**TimeBase** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Extradata** | Pointer to **string** |  | [optional] 
**VideoRange** | Pointer to **string** |  | [optional] 
**DisplayTitle** | Pointer to **string** |  | [optional] 
**DisplayLanguage** | Pointer to **string** |  | [optional] 
**NalLengthSize** | Pointer to **string** |  | [optional] 
**IsInterlaced** | Pointer to **bool** |  | [optional] 
**IsAVC** | Pointer to **NullableBool** |  | [optional] 
**ChannelLayout** | Pointer to **string** |  | [optional] 
**BitRate** | Pointer to **NullableInt32** |  | [optional] 
**BitDepth** | Pointer to **NullableInt32** |  | [optional] 
**RefFrames** | Pointer to **NullableInt32** |  | [optional] 
**Rotation** | Pointer to **NullableInt32** |  | [optional] 
**Channels** | Pointer to **NullableInt32** |  | [optional] 
**SampleRate** | Pointer to **NullableInt32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsForced** | Pointer to **bool** |  | [optional] 
**IsHearingImpaired** | Pointer to **bool** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**AverageFrameRate** | Pointer to **NullableFloat32** |  | [optional] 
**RealFrameRate** | Pointer to **NullableFloat32** |  | [optional] 
**Profile** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelModelMediaStreamType**](ModelMediaStreamType.md) |  | [optional] 
**AspectRatio** | Pointer to **string** |  | [optional] 
**Index** | Pointer to **int32** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**DeliveryMethod** | Pointer to [**ModelModelSubtitleDeliveryMethod**](ModelSubtitleDeliveryMethod.md) |  | [optional] 
**DeliveryUrl** | Pointer to **string** |  | [optional] 
**IsExternalUrl** | Pointer to **NullableBool** |  | [optional] 
**IsTextSubtitleStream** | Pointer to **bool** |  | [optional] 
**SupportsExternalStream** | Pointer to **bool** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to [**ModelModelMediaProtocol**](ModelMediaProtocol.md) |  | [optional] 
**PixelFormat** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **NullableFloat64** |  | [optional] 
**IsAnamorphic** | Pointer to **NullableBool** |  | [optional] 
**ExtendedVideoType** | Pointer to [**ModelModelExtendedVideoTypes**](ModelExtendedVideoTypes.md) |  | [optional] 
**ExtendedVideoSubType** | Pointer to [**ModelModelExtendedVideoSubTypes**](ModelExtendedVideoSubTypes.md) |  | [optional] 
**ExtendedVideoSubTypeDescription** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**AttachmentSize** | Pointer to **NullableInt32** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**SubtitleLocationType** | Pointer to [**ModelModelSubtitleLocationType**](ModelSubtitleLocationType.md) |  | [optional] 

## Methods

### NewModelMediaStream

`func NewModelMediaStream() *ModelMediaStream`

NewModelMediaStream instantiates a new ModelMediaStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMediaStreamWithDefaults

`func NewModelMediaStreamWithDefaults() *ModelMediaStream`

NewModelMediaStreamWithDefaults instantiates a new ModelMediaStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodec

`func (o *ModelMediaStream) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *ModelMediaStream) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *ModelMediaStream) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *ModelMediaStream) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### GetCodecTag

`func (o *ModelMediaStream) GetCodecTag() string`

GetCodecTag returns the CodecTag field if non-nil, zero value otherwise.

### GetCodecTagOk

`func (o *ModelMediaStream) GetCodecTagOk() (*string, bool)`

GetCodecTagOk returns a tuple with the CodecTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTag

`func (o *ModelMediaStream) SetCodecTag(v string)`

SetCodecTag sets CodecTag field to given value.

### HasCodecTag

`func (o *ModelMediaStream) HasCodecTag() bool`

HasCodecTag returns a boolean if a field has been set.

### GetLanguage

`func (o *ModelMediaStream) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ModelMediaStream) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ModelMediaStream) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ModelMediaStream) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetColorTransfer

`func (o *ModelMediaStream) GetColorTransfer() string`

GetColorTransfer returns the ColorTransfer field if non-nil, zero value otherwise.

### GetColorTransferOk

`func (o *ModelMediaStream) GetColorTransferOk() (*string, bool)`

GetColorTransferOk returns a tuple with the ColorTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorTransfer

`func (o *ModelMediaStream) SetColorTransfer(v string)`

SetColorTransfer sets ColorTransfer field to given value.

### HasColorTransfer

`func (o *ModelMediaStream) HasColorTransfer() bool`

HasColorTransfer returns a boolean if a field has been set.

### GetColorPrimaries

`func (o *ModelMediaStream) GetColorPrimaries() string`

GetColorPrimaries returns the ColorPrimaries field if non-nil, zero value otherwise.

### GetColorPrimariesOk

`func (o *ModelMediaStream) GetColorPrimariesOk() (*string, bool)`

GetColorPrimariesOk returns a tuple with the ColorPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorPrimaries

`func (o *ModelMediaStream) SetColorPrimaries(v string)`

SetColorPrimaries sets ColorPrimaries field to given value.

### HasColorPrimaries

`func (o *ModelMediaStream) HasColorPrimaries() bool`

HasColorPrimaries returns a boolean if a field has been set.

### GetColorSpace

`func (o *ModelMediaStream) GetColorSpace() string`

GetColorSpace returns the ColorSpace field if non-nil, zero value otherwise.

### GetColorSpaceOk

`func (o *ModelMediaStream) GetColorSpaceOk() (*string, bool)`

GetColorSpaceOk returns a tuple with the ColorSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorSpace

`func (o *ModelMediaStream) SetColorSpace(v string)`

SetColorSpace sets ColorSpace field to given value.

### HasColorSpace

`func (o *ModelMediaStream) HasColorSpace() bool`

HasColorSpace returns a boolean if a field has been set.

### GetComment

`func (o *ModelMediaStream) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelMediaStream) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelMediaStream) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelMediaStream) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetStreamStartTimeTicks

`func (o *ModelMediaStream) GetStreamStartTimeTicks() int64`

GetStreamStartTimeTicks returns the StreamStartTimeTicks field if non-nil, zero value otherwise.

### GetStreamStartTimeTicksOk

`func (o *ModelMediaStream) GetStreamStartTimeTicksOk() (*int64, bool)`

GetStreamStartTimeTicksOk returns a tuple with the StreamStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamStartTimeTicks

`func (o *ModelMediaStream) SetStreamStartTimeTicks(v int64)`

SetStreamStartTimeTicks sets StreamStartTimeTicks field to given value.

### HasStreamStartTimeTicks

`func (o *ModelMediaStream) HasStreamStartTimeTicks() bool`

HasStreamStartTimeTicks returns a boolean if a field has been set.

### SetStreamStartTimeTicksNil

`func (o *ModelMediaStream) SetStreamStartTimeTicksNil(b bool)`

 SetStreamStartTimeTicksNil sets the value for StreamStartTimeTicks to be an explicit nil

### UnsetStreamStartTimeTicks
`func (o *ModelMediaStream) UnsetStreamStartTimeTicks()`

UnsetStreamStartTimeTicks ensures that no value is present for StreamStartTimeTicks, not even an explicit nil
### GetTimeBase

`func (o *ModelMediaStream) GetTimeBase() string`

GetTimeBase returns the TimeBase field if non-nil, zero value otherwise.

### GetTimeBaseOk

`func (o *ModelMediaStream) GetTimeBaseOk() (*string, bool)`

GetTimeBaseOk returns a tuple with the TimeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBase

`func (o *ModelMediaStream) SetTimeBase(v string)`

SetTimeBase sets TimeBase field to given value.

### HasTimeBase

`func (o *ModelMediaStream) HasTimeBase() bool`

HasTimeBase returns a boolean if a field has been set.

### GetTitle

`func (o *ModelMediaStream) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModelMediaStream) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModelMediaStream) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModelMediaStream) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetExtradata

`func (o *ModelMediaStream) GetExtradata() string`

GetExtradata returns the Extradata field if non-nil, zero value otherwise.

### GetExtradataOk

`func (o *ModelMediaStream) GetExtradataOk() (*string, bool)`

GetExtradataOk returns a tuple with the Extradata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtradata

`func (o *ModelMediaStream) SetExtradata(v string)`

SetExtradata sets Extradata field to given value.

### HasExtradata

`func (o *ModelMediaStream) HasExtradata() bool`

HasExtradata returns a boolean if a field has been set.

### GetVideoRange

`func (o *ModelMediaStream) GetVideoRange() string`

GetVideoRange returns the VideoRange field if non-nil, zero value otherwise.

### GetVideoRangeOk

`func (o *ModelMediaStream) GetVideoRangeOk() (*string, bool)`

GetVideoRangeOk returns a tuple with the VideoRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRange

`func (o *ModelMediaStream) SetVideoRange(v string)`

SetVideoRange sets VideoRange field to given value.

### HasVideoRange

`func (o *ModelMediaStream) HasVideoRange() bool`

HasVideoRange returns a boolean if a field has been set.

### GetDisplayTitle

`func (o *ModelMediaStream) GetDisplayTitle() string`

GetDisplayTitle returns the DisplayTitle field if non-nil, zero value otherwise.

### GetDisplayTitleOk

`func (o *ModelMediaStream) GetDisplayTitleOk() (*string, bool)`

GetDisplayTitleOk returns a tuple with the DisplayTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTitle

`func (o *ModelMediaStream) SetDisplayTitle(v string)`

SetDisplayTitle sets DisplayTitle field to given value.

### HasDisplayTitle

`func (o *ModelMediaStream) HasDisplayTitle() bool`

HasDisplayTitle returns a boolean if a field has been set.

### GetDisplayLanguage

`func (o *ModelMediaStream) GetDisplayLanguage() string`

GetDisplayLanguage returns the DisplayLanguage field if non-nil, zero value otherwise.

### GetDisplayLanguageOk

`func (o *ModelMediaStream) GetDisplayLanguageOk() (*string, bool)`

GetDisplayLanguageOk returns a tuple with the DisplayLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLanguage

`func (o *ModelMediaStream) SetDisplayLanguage(v string)`

SetDisplayLanguage sets DisplayLanguage field to given value.

### HasDisplayLanguage

`func (o *ModelMediaStream) HasDisplayLanguage() bool`

HasDisplayLanguage returns a boolean if a field has been set.

### GetNalLengthSize

`func (o *ModelMediaStream) GetNalLengthSize() string`

GetNalLengthSize returns the NalLengthSize field if non-nil, zero value otherwise.

### GetNalLengthSizeOk

`func (o *ModelMediaStream) GetNalLengthSizeOk() (*string, bool)`

GetNalLengthSizeOk returns a tuple with the NalLengthSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNalLengthSize

`func (o *ModelMediaStream) SetNalLengthSize(v string)`

SetNalLengthSize sets NalLengthSize field to given value.

### HasNalLengthSize

`func (o *ModelMediaStream) HasNalLengthSize() bool`

HasNalLengthSize returns a boolean if a field has been set.

### GetIsInterlaced

`func (o *ModelMediaStream) GetIsInterlaced() bool`

GetIsInterlaced returns the IsInterlaced field if non-nil, zero value otherwise.

### GetIsInterlacedOk

`func (o *ModelMediaStream) GetIsInterlacedOk() (*bool, bool)`

GetIsInterlacedOk returns a tuple with the IsInterlaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterlaced

`func (o *ModelMediaStream) SetIsInterlaced(v bool)`

SetIsInterlaced sets IsInterlaced field to given value.

### HasIsInterlaced

`func (o *ModelMediaStream) HasIsInterlaced() bool`

HasIsInterlaced returns a boolean if a field has been set.

### GetIsAVC

`func (o *ModelMediaStream) GetIsAVC() bool`

GetIsAVC returns the IsAVC field if non-nil, zero value otherwise.

### GetIsAVCOk

`func (o *ModelMediaStream) GetIsAVCOk() (*bool, bool)`

GetIsAVCOk returns a tuple with the IsAVC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAVC

`func (o *ModelMediaStream) SetIsAVC(v bool)`

SetIsAVC sets IsAVC field to given value.

### HasIsAVC

`func (o *ModelMediaStream) HasIsAVC() bool`

HasIsAVC returns a boolean if a field has been set.

### SetIsAVCNil

`func (o *ModelMediaStream) SetIsAVCNil(b bool)`

 SetIsAVCNil sets the value for IsAVC to be an explicit nil

### UnsetIsAVC
`func (o *ModelMediaStream) UnsetIsAVC()`

UnsetIsAVC ensures that no value is present for IsAVC, not even an explicit nil
### GetChannelLayout

`func (o *ModelMediaStream) GetChannelLayout() string`

GetChannelLayout returns the ChannelLayout field if non-nil, zero value otherwise.

### GetChannelLayoutOk

`func (o *ModelMediaStream) GetChannelLayoutOk() (*string, bool)`

GetChannelLayoutOk returns a tuple with the ChannelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLayout

`func (o *ModelMediaStream) SetChannelLayout(v string)`

SetChannelLayout sets ChannelLayout field to given value.

### HasChannelLayout

`func (o *ModelMediaStream) HasChannelLayout() bool`

HasChannelLayout returns a boolean if a field has been set.

### GetBitRate

`func (o *ModelMediaStream) GetBitRate() int32`

GetBitRate returns the BitRate field if non-nil, zero value otherwise.

### GetBitRateOk

`func (o *ModelMediaStream) GetBitRateOk() (*int32, bool)`

GetBitRateOk returns a tuple with the BitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitRate

`func (o *ModelMediaStream) SetBitRate(v int32)`

SetBitRate sets BitRate field to given value.

### HasBitRate

`func (o *ModelMediaStream) HasBitRate() bool`

HasBitRate returns a boolean if a field has been set.

### SetBitRateNil

`func (o *ModelMediaStream) SetBitRateNil(b bool)`

 SetBitRateNil sets the value for BitRate to be an explicit nil

### UnsetBitRate
`func (o *ModelMediaStream) UnsetBitRate()`

UnsetBitRate ensures that no value is present for BitRate, not even an explicit nil
### GetBitDepth

`func (o *ModelMediaStream) GetBitDepth() int32`

GetBitDepth returns the BitDepth field if non-nil, zero value otherwise.

### GetBitDepthOk

`func (o *ModelMediaStream) GetBitDepthOk() (*int32, bool)`

GetBitDepthOk returns a tuple with the BitDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitDepth

`func (o *ModelMediaStream) SetBitDepth(v int32)`

SetBitDepth sets BitDepth field to given value.

### HasBitDepth

`func (o *ModelMediaStream) HasBitDepth() bool`

HasBitDepth returns a boolean if a field has been set.

### SetBitDepthNil

`func (o *ModelMediaStream) SetBitDepthNil(b bool)`

 SetBitDepthNil sets the value for BitDepth to be an explicit nil

### UnsetBitDepth
`func (o *ModelMediaStream) UnsetBitDepth()`

UnsetBitDepth ensures that no value is present for BitDepth, not even an explicit nil
### GetRefFrames

`func (o *ModelMediaStream) GetRefFrames() int32`

GetRefFrames returns the RefFrames field if non-nil, zero value otherwise.

### GetRefFramesOk

`func (o *ModelMediaStream) GetRefFramesOk() (*int32, bool)`

GetRefFramesOk returns a tuple with the RefFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefFrames

`func (o *ModelMediaStream) SetRefFrames(v int32)`

SetRefFrames sets RefFrames field to given value.

### HasRefFrames

`func (o *ModelMediaStream) HasRefFrames() bool`

HasRefFrames returns a boolean if a field has been set.

### SetRefFramesNil

`func (o *ModelMediaStream) SetRefFramesNil(b bool)`

 SetRefFramesNil sets the value for RefFrames to be an explicit nil

### UnsetRefFrames
`func (o *ModelMediaStream) UnsetRefFrames()`

UnsetRefFrames ensures that no value is present for RefFrames, not even an explicit nil
### GetRotation

`func (o *ModelMediaStream) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *ModelMediaStream) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *ModelMediaStream) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *ModelMediaStream) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### SetRotationNil

`func (o *ModelMediaStream) SetRotationNil(b bool)`

 SetRotationNil sets the value for Rotation to be an explicit nil

### UnsetRotation
`func (o *ModelMediaStream) UnsetRotation()`

UnsetRotation ensures that no value is present for Rotation, not even an explicit nil
### GetChannels

`func (o *ModelMediaStream) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ModelMediaStream) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ModelMediaStream) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ModelMediaStream) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *ModelMediaStream) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *ModelMediaStream) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetSampleRate

`func (o *ModelMediaStream) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *ModelMediaStream) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *ModelMediaStream) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *ModelMediaStream) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### SetSampleRateNil

`func (o *ModelMediaStream) SetSampleRateNil(b bool)`

 SetSampleRateNil sets the value for SampleRate to be an explicit nil

### UnsetSampleRate
`func (o *ModelMediaStream) UnsetSampleRate()`

UnsetSampleRate ensures that no value is present for SampleRate, not even an explicit nil
### GetIsDefault

`func (o *ModelMediaStream) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ModelMediaStream) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ModelMediaStream) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ModelMediaStream) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsForced

`func (o *ModelMediaStream) GetIsForced() bool`

GetIsForced returns the IsForced field if non-nil, zero value otherwise.

### GetIsForcedOk

`func (o *ModelMediaStream) GetIsForcedOk() (*bool, bool)`

GetIsForcedOk returns a tuple with the IsForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForced

`func (o *ModelMediaStream) SetIsForced(v bool)`

SetIsForced sets IsForced field to given value.

### HasIsForced

`func (o *ModelMediaStream) HasIsForced() bool`

HasIsForced returns a boolean if a field has been set.

### GetIsHearingImpaired

`func (o *ModelMediaStream) GetIsHearingImpaired() bool`

GetIsHearingImpaired returns the IsHearingImpaired field if non-nil, zero value otherwise.

### GetIsHearingImpairedOk

`func (o *ModelMediaStream) GetIsHearingImpairedOk() (*bool, bool)`

GetIsHearingImpairedOk returns a tuple with the IsHearingImpaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHearingImpaired

`func (o *ModelMediaStream) SetIsHearingImpaired(v bool)`

SetIsHearingImpaired sets IsHearingImpaired field to given value.

### HasIsHearingImpaired

`func (o *ModelMediaStream) HasIsHearingImpaired() bool`

HasIsHearingImpaired returns a boolean if a field has been set.

### GetHeight

`func (o *ModelMediaStream) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ModelMediaStream) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ModelMediaStream) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ModelMediaStream) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ModelMediaStream) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ModelMediaStream) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *ModelMediaStream) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ModelMediaStream) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ModelMediaStream) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ModelMediaStream) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ModelMediaStream) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ModelMediaStream) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetAverageFrameRate

`func (o *ModelMediaStream) GetAverageFrameRate() float32`

GetAverageFrameRate returns the AverageFrameRate field if non-nil, zero value otherwise.

### GetAverageFrameRateOk

`func (o *ModelMediaStream) GetAverageFrameRateOk() (*float32, bool)`

GetAverageFrameRateOk returns a tuple with the AverageFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFrameRate

`func (o *ModelMediaStream) SetAverageFrameRate(v float32)`

SetAverageFrameRate sets AverageFrameRate field to given value.

### HasAverageFrameRate

`func (o *ModelMediaStream) HasAverageFrameRate() bool`

HasAverageFrameRate returns a boolean if a field has been set.

### SetAverageFrameRateNil

`func (o *ModelMediaStream) SetAverageFrameRateNil(b bool)`

 SetAverageFrameRateNil sets the value for AverageFrameRate to be an explicit nil

### UnsetAverageFrameRate
`func (o *ModelMediaStream) UnsetAverageFrameRate()`

UnsetAverageFrameRate ensures that no value is present for AverageFrameRate, not even an explicit nil
### GetRealFrameRate

`func (o *ModelMediaStream) GetRealFrameRate() float32`

GetRealFrameRate returns the RealFrameRate field if non-nil, zero value otherwise.

### GetRealFrameRateOk

`func (o *ModelMediaStream) GetRealFrameRateOk() (*float32, bool)`

GetRealFrameRateOk returns a tuple with the RealFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealFrameRate

`func (o *ModelMediaStream) SetRealFrameRate(v float32)`

SetRealFrameRate sets RealFrameRate field to given value.

### HasRealFrameRate

`func (o *ModelMediaStream) HasRealFrameRate() bool`

HasRealFrameRate returns a boolean if a field has been set.

### SetRealFrameRateNil

`func (o *ModelMediaStream) SetRealFrameRateNil(b bool)`

 SetRealFrameRateNil sets the value for RealFrameRate to be an explicit nil

### UnsetRealFrameRate
`func (o *ModelMediaStream) UnsetRealFrameRate()`

UnsetRealFrameRate ensures that no value is present for RealFrameRate, not even an explicit nil
### GetProfile

`func (o *ModelMediaStream) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ModelMediaStream) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ModelMediaStream) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ModelMediaStream) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetType

`func (o *ModelMediaStream) GetType() ModelModelMediaStreamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelMediaStream) GetTypeOk() (*ModelModelMediaStreamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelMediaStream) SetType(v ModelModelMediaStreamType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelMediaStream) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAspectRatio

`func (o *ModelMediaStream) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ModelMediaStream) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ModelMediaStream) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *ModelMediaStream) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetIndex

`func (o *ModelMediaStream) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ModelMediaStream) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ModelMediaStream) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ModelMediaStream) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIsExternal

`func (o *ModelMediaStream) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *ModelMediaStream) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *ModelMediaStream) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *ModelMediaStream) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *ModelMediaStream) GetDeliveryMethod() ModelModelSubtitleDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *ModelMediaStream) GetDeliveryMethodOk() (*ModelModelSubtitleDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *ModelMediaStream) SetDeliveryMethod(v ModelModelSubtitleDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *ModelMediaStream) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### GetDeliveryUrl

`func (o *ModelMediaStream) GetDeliveryUrl() string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ModelMediaStream) GetDeliveryUrlOk() (*string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ModelMediaStream) SetDeliveryUrl(v string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ModelMediaStream) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### GetIsExternalUrl

`func (o *ModelMediaStream) GetIsExternalUrl() bool`

GetIsExternalUrl returns the IsExternalUrl field if non-nil, zero value otherwise.

### GetIsExternalUrlOk

`func (o *ModelMediaStream) GetIsExternalUrlOk() (*bool, bool)`

GetIsExternalUrlOk returns a tuple with the IsExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalUrl

`func (o *ModelMediaStream) SetIsExternalUrl(v bool)`

SetIsExternalUrl sets IsExternalUrl field to given value.

### HasIsExternalUrl

`func (o *ModelMediaStream) HasIsExternalUrl() bool`

HasIsExternalUrl returns a boolean if a field has been set.

### SetIsExternalUrlNil

`func (o *ModelMediaStream) SetIsExternalUrlNil(b bool)`

 SetIsExternalUrlNil sets the value for IsExternalUrl to be an explicit nil

### UnsetIsExternalUrl
`func (o *ModelMediaStream) UnsetIsExternalUrl()`

UnsetIsExternalUrl ensures that no value is present for IsExternalUrl, not even an explicit nil
### GetIsTextSubtitleStream

`func (o *ModelMediaStream) GetIsTextSubtitleStream() bool`

GetIsTextSubtitleStream returns the IsTextSubtitleStream field if non-nil, zero value otherwise.

### GetIsTextSubtitleStreamOk

`func (o *ModelMediaStream) GetIsTextSubtitleStreamOk() (*bool, bool)`

GetIsTextSubtitleStreamOk returns a tuple with the IsTextSubtitleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTextSubtitleStream

`func (o *ModelMediaStream) SetIsTextSubtitleStream(v bool)`

SetIsTextSubtitleStream sets IsTextSubtitleStream field to given value.

### HasIsTextSubtitleStream

`func (o *ModelMediaStream) HasIsTextSubtitleStream() bool`

HasIsTextSubtitleStream returns a boolean if a field has been set.

### GetSupportsExternalStream

`func (o *ModelMediaStream) GetSupportsExternalStream() bool`

GetSupportsExternalStream returns the SupportsExternalStream field if non-nil, zero value otherwise.

### GetSupportsExternalStreamOk

`func (o *ModelMediaStream) GetSupportsExternalStreamOk() (*bool, bool)`

GetSupportsExternalStreamOk returns a tuple with the SupportsExternalStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExternalStream

`func (o *ModelMediaStream) SetSupportsExternalStream(v bool)`

SetSupportsExternalStream sets SupportsExternalStream field to given value.

### HasSupportsExternalStream

`func (o *ModelMediaStream) HasSupportsExternalStream() bool`

HasSupportsExternalStream returns a boolean if a field has been set.

### GetPath

`func (o *ModelMediaStream) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelMediaStream) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelMediaStream) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelMediaStream) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *ModelMediaStream) GetProtocol() ModelModelMediaProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelMediaStream) GetProtocolOk() (*ModelModelMediaProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelMediaStream) SetProtocol(v ModelModelMediaProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ModelMediaStream) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPixelFormat

`func (o *ModelMediaStream) GetPixelFormat() string`

GetPixelFormat returns the PixelFormat field if non-nil, zero value otherwise.

### GetPixelFormatOk

`func (o *ModelMediaStream) GetPixelFormatOk() (*string, bool)`

GetPixelFormatOk returns a tuple with the PixelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelFormat

`func (o *ModelMediaStream) SetPixelFormat(v string)`

SetPixelFormat sets PixelFormat field to given value.

### HasPixelFormat

`func (o *ModelMediaStream) HasPixelFormat() bool`

HasPixelFormat returns a boolean if a field has been set.

### GetLevel

`func (o *ModelMediaStream) GetLevel() float64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ModelMediaStream) GetLevelOk() (*float64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ModelMediaStream) SetLevel(v float64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ModelMediaStream) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *ModelMediaStream) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *ModelMediaStream) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetIsAnamorphic

`func (o *ModelMediaStream) GetIsAnamorphic() bool`

GetIsAnamorphic returns the IsAnamorphic field if non-nil, zero value otherwise.

### GetIsAnamorphicOk

`func (o *ModelMediaStream) GetIsAnamorphicOk() (*bool, bool)`

GetIsAnamorphicOk returns a tuple with the IsAnamorphic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnamorphic

`func (o *ModelMediaStream) SetIsAnamorphic(v bool)`

SetIsAnamorphic sets IsAnamorphic field to given value.

### HasIsAnamorphic

`func (o *ModelMediaStream) HasIsAnamorphic() bool`

HasIsAnamorphic returns a boolean if a field has been set.

### SetIsAnamorphicNil

`func (o *ModelMediaStream) SetIsAnamorphicNil(b bool)`

 SetIsAnamorphicNil sets the value for IsAnamorphic to be an explicit nil

### UnsetIsAnamorphic
`func (o *ModelMediaStream) UnsetIsAnamorphic()`

UnsetIsAnamorphic ensures that no value is present for IsAnamorphic, not even an explicit nil
### GetExtendedVideoType

`func (o *ModelMediaStream) GetExtendedVideoType() ModelModelExtendedVideoTypes`

GetExtendedVideoType returns the ExtendedVideoType field if non-nil, zero value otherwise.

### GetExtendedVideoTypeOk

`func (o *ModelMediaStream) GetExtendedVideoTypeOk() (*ModelModelExtendedVideoTypes, bool)`

GetExtendedVideoTypeOk returns a tuple with the ExtendedVideoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVideoType

`func (o *ModelMediaStream) SetExtendedVideoType(v ModelModelExtendedVideoTypes)`

SetExtendedVideoType sets ExtendedVideoType field to given value.

### HasExtendedVideoType

`func (o *ModelMediaStream) HasExtendedVideoType() bool`

HasExtendedVideoType returns a boolean if a field has been set.

### GetExtendedVideoSubType

`func (o *ModelMediaStream) GetExtendedVideoSubType() ModelModelExtendedVideoSubTypes`

GetExtendedVideoSubType returns the ExtendedVideoSubType field if non-nil, zero value otherwise.

### GetExtendedVideoSubTypeOk

`func (o *ModelMediaStream) GetExtendedVideoSubTypeOk() (*ModelModelExtendedVideoSubTypes, bool)`

GetExtendedVideoSubTypeOk returns a tuple with the ExtendedVideoSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVideoSubType

`func (o *ModelMediaStream) SetExtendedVideoSubType(v ModelModelExtendedVideoSubTypes)`

SetExtendedVideoSubType sets ExtendedVideoSubType field to given value.

### HasExtendedVideoSubType

`func (o *ModelMediaStream) HasExtendedVideoSubType() bool`

HasExtendedVideoSubType returns a boolean if a field has been set.

### GetExtendedVideoSubTypeDescription

`func (o *ModelMediaStream) GetExtendedVideoSubTypeDescription() string`

GetExtendedVideoSubTypeDescription returns the ExtendedVideoSubTypeDescription field if non-nil, zero value otherwise.

### GetExtendedVideoSubTypeDescriptionOk

`func (o *ModelMediaStream) GetExtendedVideoSubTypeDescriptionOk() (*string, bool)`

GetExtendedVideoSubTypeDescriptionOk returns a tuple with the ExtendedVideoSubTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedVideoSubTypeDescription

`func (o *ModelMediaStream) SetExtendedVideoSubTypeDescription(v string)`

SetExtendedVideoSubTypeDescription sets ExtendedVideoSubTypeDescription field to given value.

### HasExtendedVideoSubTypeDescription

`func (o *ModelMediaStream) HasExtendedVideoSubTypeDescription() bool`

HasExtendedVideoSubTypeDescription returns a boolean if a field has been set.

### GetItemId

`func (o *ModelMediaStream) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelMediaStream) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelMediaStream) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelMediaStream) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetServerId

`func (o *ModelMediaStream) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelMediaStream) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelMediaStream) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelMediaStream) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetAttachmentSize

`func (o *ModelMediaStream) GetAttachmentSize() int32`

GetAttachmentSize returns the AttachmentSize field if non-nil, zero value otherwise.

### GetAttachmentSizeOk

`func (o *ModelMediaStream) GetAttachmentSizeOk() (*int32, bool)`

GetAttachmentSizeOk returns a tuple with the AttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentSize

`func (o *ModelMediaStream) SetAttachmentSize(v int32)`

SetAttachmentSize sets AttachmentSize field to given value.

### HasAttachmentSize

`func (o *ModelMediaStream) HasAttachmentSize() bool`

HasAttachmentSize returns a boolean if a field has been set.

### SetAttachmentSizeNil

`func (o *ModelMediaStream) SetAttachmentSizeNil(b bool)`

 SetAttachmentSizeNil sets the value for AttachmentSize to be an explicit nil

### UnsetAttachmentSize
`func (o *ModelMediaStream) UnsetAttachmentSize()`

UnsetAttachmentSize ensures that no value is present for AttachmentSize, not even an explicit nil
### GetMimeType

`func (o *ModelMediaStream) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ModelMediaStream) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ModelMediaStream) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ModelMediaStream) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetSubtitleLocationType

`func (o *ModelMediaStream) GetSubtitleLocationType() ModelModelSubtitleLocationType`

GetSubtitleLocationType returns the SubtitleLocationType field if non-nil, zero value otherwise.

### GetSubtitleLocationTypeOk

`func (o *ModelMediaStream) GetSubtitleLocationTypeOk() (*ModelModelSubtitleLocationType, bool)`

GetSubtitleLocationTypeOk returns a tuple with the SubtitleLocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLocationType

`func (o *ModelMediaStream) SetSubtitleLocationType(v ModelModelSubtitleLocationType)`

SetSubtitleLocationType sets SubtitleLocationType field to given value.

### HasSubtitleLocationType

`func (o *ModelMediaStream) HasSubtitleLocationType() bool`

HasSubtitleLocationType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


