# ModelTranscodingProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelModelDlnaProfileType**](ModelDlnaProfileType.md) |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**EstimateContentLength** | Pointer to **bool** |  | [optional] 
**EnableMpegtsM2TsMode** | Pointer to **bool** |  | [optional] 
**TranscodeSeekInfo** | Pointer to [**ModelModelTranscodeSeekInfo**](ModelTranscodeSeekInfo.md) |  | [optional] 
**CopyTimestamps** | Pointer to **bool** |  | [optional] 
**Context** | Pointer to [**ModelModelEncodingContext**](ModelEncodingContext.md) |  | [optional] 
**MaxAudioChannels** | Pointer to **string** |  | [optional] 
**MinSegments** | Pointer to **int32** |  | [optional] 
**SegmentLength** | Pointer to **int32** |  | [optional] 
**BreakOnNonKeyFrames** | Pointer to **bool** |  | [optional] 
**AllowInterlacedVideoStreamCopy** | Pointer to **bool** |  | [optional] 
**ManifestSubtitles** | Pointer to **string** |  | [optional] 
**MaxManifestSubtitles** | Pointer to **int32** |  | [optional] 
**MaxWidth** | Pointer to **int32** |  | [optional] 
**MaxHeight** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelTranscodingProfile

`func NewModelTranscodingProfile() *ModelTranscodingProfile`

NewModelTranscodingProfile instantiates a new ModelTranscodingProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTranscodingProfileWithDefaults

`func NewModelTranscodingProfileWithDefaults() *ModelTranscodingProfile`

NewModelTranscodingProfileWithDefaults instantiates a new ModelTranscodingProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ModelTranscodingProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelTranscodingProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelTranscodingProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelTranscodingProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetType

`func (o *ModelTranscodingProfile) GetType() ModelModelDlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelTranscodingProfile) GetTypeOk() (*ModelModelDlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelTranscodingProfile) SetType(v ModelModelDlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelTranscodingProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVideoCodec

`func (o *ModelTranscodingProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelTranscodingProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelTranscodingProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelTranscodingProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetAudioCodec

`func (o *ModelTranscodingProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelTranscodingProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelTranscodingProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelTranscodingProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetProtocol

`func (o *ModelTranscodingProfile) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelTranscodingProfile) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelTranscodingProfile) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ModelTranscodingProfile) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetEstimateContentLength

`func (o *ModelTranscodingProfile) GetEstimateContentLength() bool`

GetEstimateContentLength returns the EstimateContentLength field if non-nil, zero value otherwise.

### GetEstimateContentLengthOk

`func (o *ModelTranscodingProfile) GetEstimateContentLengthOk() (*bool, bool)`

GetEstimateContentLengthOk returns a tuple with the EstimateContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateContentLength

`func (o *ModelTranscodingProfile) SetEstimateContentLength(v bool)`

SetEstimateContentLength sets EstimateContentLength field to given value.

### HasEstimateContentLength

`func (o *ModelTranscodingProfile) HasEstimateContentLength() bool`

HasEstimateContentLength returns a boolean if a field has been set.

### GetEnableMpegtsM2TsMode

`func (o *ModelTranscodingProfile) GetEnableMpegtsM2TsMode() bool`

GetEnableMpegtsM2TsMode returns the EnableMpegtsM2TsMode field if non-nil, zero value otherwise.

### GetEnableMpegtsM2TsModeOk

`func (o *ModelTranscodingProfile) GetEnableMpegtsM2TsModeOk() (*bool, bool)`

GetEnableMpegtsM2TsModeOk returns a tuple with the EnableMpegtsM2TsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMpegtsM2TsMode

`func (o *ModelTranscodingProfile) SetEnableMpegtsM2TsMode(v bool)`

SetEnableMpegtsM2TsMode sets EnableMpegtsM2TsMode field to given value.

### HasEnableMpegtsM2TsMode

`func (o *ModelTranscodingProfile) HasEnableMpegtsM2TsMode() bool`

HasEnableMpegtsM2TsMode returns a boolean if a field has been set.

### GetTranscodeSeekInfo

`func (o *ModelTranscodingProfile) GetTranscodeSeekInfo() ModelModelTranscodeSeekInfo`

GetTranscodeSeekInfo returns the TranscodeSeekInfo field if non-nil, zero value otherwise.

### GetTranscodeSeekInfoOk

`func (o *ModelTranscodingProfile) GetTranscodeSeekInfoOk() (*ModelModelTranscodeSeekInfo, bool)`

GetTranscodeSeekInfoOk returns a tuple with the TranscodeSeekInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeSeekInfo

`func (o *ModelTranscodingProfile) SetTranscodeSeekInfo(v ModelModelTranscodeSeekInfo)`

SetTranscodeSeekInfo sets TranscodeSeekInfo field to given value.

### HasTranscodeSeekInfo

`func (o *ModelTranscodingProfile) HasTranscodeSeekInfo() bool`

HasTranscodeSeekInfo returns a boolean if a field has been set.

### GetCopyTimestamps

`func (o *ModelTranscodingProfile) GetCopyTimestamps() bool`

GetCopyTimestamps returns the CopyTimestamps field if non-nil, zero value otherwise.

### GetCopyTimestampsOk

`func (o *ModelTranscodingProfile) GetCopyTimestampsOk() (*bool, bool)`

GetCopyTimestampsOk returns a tuple with the CopyTimestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTimestamps

`func (o *ModelTranscodingProfile) SetCopyTimestamps(v bool)`

SetCopyTimestamps sets CopyTimestamps field to given value.

### HasCopyTimestamps

`func (o *ModelTranscodingProfile) HasCopyTimestamps() bool`

HasCopyTimestamps returns a boolean if a field has been set.

### GetContext

`func (o *ModelTranscodingProfile) GetContext() ModelModelEncodingContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ModelTranscodingProfile) GetContextOk() (*ModelModelEncodingContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ModelTranscodingProfile) SetContext(v ModelModelEncodingContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ModelTranscodingProfile) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMaxAudioChannels

`func (o *ModelTranscodingProfile) GetMaxAudioChannels() string`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *ModelTranscodingProfile) GetMaxAudioChannelsOk() (*string, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *ModelTranscodingProfile) SetMaxAudioChannels(v string)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *ModelTranscodingProfile) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### GetMinSegments

`func (o *ModelTranscodingProfile) GetMinSegments() int32`

GetMinSegments returns the MinSegments field if non-nil, zero value otherwise.

### GetMinSegmentsOk

`func (o *ModelTranscodingProfile) GetMinSegmentsOk() (*int32, bool)`

GetMinSegmentsOk returns a tuple with the MinSegments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSegments

`func (o *ModelTranscodingProfile) SetMinSegments(v int32)`

SetMinSegments sets MinSegments field to given value.

### HasMinSegments

`func (o *ModelTranscodingProfile) HasMinSegments() bool`

HasMinSegments returns a boolean if a field has been set.

### GetSegmentLength

`func (o *ModelTranscodingProfile) GetSegmentLength() int32`

GetSegmentLength returns the SegmentLength field if non-nil, zero value otherwise.

### GetSegmentLengthOk

`func (o *ModelTranscodingProfile) GetSegmentLengthOk() (*int32, bool)`

GetSegmentLengthOk returns a tuple with the SegmentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentLength

`func (o *ModelTranscodingProfile) SetSegmentLength(v int32)`

SetSegmentLength sets SegmentLength field to given value.

### HasSegmentLength

`func (o *ModelTranscodingProfile) HasSegmentLength() bool`

HasSegmentLength returns a boolean if a field has been set.

### GetBreakOnNonKeyFrames

`func (o *ModelTranscodingProfile) GetBreakOnNonKeyFrames() bool`

GetBreakOnNonKeyFrames returns the BreakOnNonKeyFrames field if non-nil, zero value otherwise.

### GetBreakOnNonKeyFramesOk

`func (o *ModelTranscodingProfile) GetBreakOnNonKeyFramesOk() (*bool, bool)`

GetBreakOnNonKeyFramesOk returns a tuple with the BreakOnNonKeyFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakOnNonKeyFrames

`func (o *ModelTranscodingProfile) SetBreakOnNonKeyFrames(v bool)`

SetBreakOnNonKeyFrames sets BreakOnNonKeyFrames field to given value.

### HasBreakOnNonKeyFrames

`func (o *ModelTranscodingProfile) HasBreakOnNonKeyFrames() bool`

HasBreakOnNonKeyFrames returns a boolean if a field has been set.

### GetAllowInterlacedVideoStreamCopy

`func (o *ModelTranscodingProfile) GetAllowInterlacedVideoStreamCopy() bool`

GetAllowInterlacedVideoStreamCopy returns the AllowInterlacedVideoStreamCopy field if non-nil, zero value otherwise.

### GetAllowInterlacedVideoStreamCopyOk

`func (o *ModelTranscodingProfile) GetAllowInterlacedVideoStreamCopyOk() (*bool, bool)`

GetAllowInterlacedVideoStreamCopyOk returns a tuple with the AllowInterlacedVideoStreamCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInterlacedVideoStreamCopy

`func (o *ModelTranscodingProfile) SetAllowInterlacedVideoStreamCopy(v bool)`

SetAllowInterlacedVideoStreamCopy sets AllowInterlacedVideoStreamCopy field to given value.

### HasAllowInterlacedVideoStreamCopy

`func (o *ModelTranscodingProfile) HasAllowInterlacedVideoStreamCopy() bool`

HasAllowInterlacedVideoStreamCopy returns a boolean if a field has been set.

### GetManifestSubtitles

`func (o *ModelTranscodingProfile) GetManifestSubtitles() string`

GetManifestSubtitles returns the ManifestSubtitles field if non-nil, zero value otherwise.

### GetManifestSubtitlesOk

`func (o *ModelTranscodingProfile) GetManifestSubtitlesOk() (*string, bool)`

GetManifestSubtitlesOk returns a tuple with the ManifestSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSubtitles

`func (o *ModelTranscodingProfile) SetManifestSubtitles(v string)`

SetManifestSubtitles sets ManifestSubtitles field to given value.

### HasManifestSubtitles

`func (o *ModelTranscodingProfile) HasManifestSubtitles() bool`

HasManifestSubtitles returns a boolean if a field has been set.

### GetMaxManifestSubtitles

`func (o *ModelTranscodingProfile) GetMaxManifestSubtitles() int32`

GetMaxManifestSubtitles returns the MaxManifestSubtitles field if non-nil, zero value otherwise.

### GetMaxManifestSubtitlesOk

`func (o *ModelTranscodingProfile) GetMaxManifestSubtitlesOk() (*int32, bool)`

GetMaxManifestSubtitlesOk returns a tuple with the MaxManifestSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxManifestSubtitles

`func (o *ModelTranscodingProfile) SetMaxManifestSubtitles(v int32)`

SetMaxManifestSubtitles sets MaxManifestSubtitles field to given value.

### HasMaxManifestSubtitles

`func (o *ModelTranscodingProfile) HasMaxManifestSubtitles() bool`

HasMaxManifestSubtitles returns a boolean if a field has been set.

### GetMaxWidth

`func (o *ModelTranscodingProfile) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *ModelTranscodingProfile) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *ModelTranscodingProfile) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *ModelTranscodingProfile) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### GetMaxHeight

`func (o *ModelTranscodingProfile) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *ModelTranscodingProfile) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *ModelTranscodingProfile) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *ModelTranscodingProfile) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


