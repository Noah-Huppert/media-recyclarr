# ModelDeviceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewModelDeviceProfile

`func NewModelDeviceProfile() *ModelDeviceProfile`

NewModelDeviceProfile instantiates a new ModelDeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDeviceProfileWithDefaults

`func NewModelDeviceProfileWithDefaults() *ModelDeviceProfile`

NewModelDeviceProfileWithDefaults instantiates a new ModelDeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelDeviceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelDeviceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelDeviceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelDeviceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelDeviceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDeviceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDeviceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDeviceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSupportedMediaTypes

`func (o *ModelDeviceProfile) GetSupportedMediaTypes() string`

GetSupportedMediaTypes returns the SupportedMediaTypes field if non-nil, zero value otherwise.

### GetSupportedMediaTypesOk

`func (o *ModelDeviceProfile) GetSupportedMediaTypesOk() (*string, bool)`

GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMediaTypes

`func (o *ModelDeviceProfile) SetSupportedMediaTypes(v string)`

SetSupportedMediaTypes sets SupportedMediaTypes field to given value.

### HasSupportedMediaTypes

`func (o *ModelDeviceProfile) HasSupportedMediaTypes() bool`

HasSupportedMediaTypes returns a boolean if a field has been set.

### GetMaxStreamingBitrate

`func (o *ModelDeviceProfile) GetMaxStreamingBitrate() int64`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *ModelDeviceProfile) GetMaxStreamingBitrateOk() (*int64, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *ModelDeviceProfile) SetMaxStreamingBitrate(v int64)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *ModelDeviceProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *ModelDeviceProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *ModelDeviceProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *ModelDeviceProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *ModelDeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *ModelDeviceProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *ModelDeviceProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *ModelDeviceProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *ModelDeviceProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *ModelDeviceProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *ModelDeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *ModelDeviceProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *ModelDeviceProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *ModelDeviceProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *ModelDeviceProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetDirectPlayProfiles

`func (o *ModelDeviceProfile) GetDirectPlayProfiles() []ModelModelDirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *ModelDeviceProfile) GetDirectPlayProfilesOk() (*[]ModelModelDirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *ModelDeviceProfile) SetDirectPlayProfiles(v []ModelModelDirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *ModelDeviceProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *ModelDeviceProfile) GetTranscodingProfiles() []ModelModelTranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *ModelDeviceProfile) GetTranscodingProfilesOk() (*[]ModelModelTranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *ModelDeviceProfile) SetTranscodingProfiles(v []ModelModelTranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *ModelDeviceProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *ModelDeviceProfile) GetContainerProfiles() []ModelModelContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *ModelDeviceProfile) GetContainerProfilesOk() (*[]ModelModelContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *ModelDeviceProfile) SetContainerProfiles(v []ModelModelContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *ModelDeviceProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *ModelDeviceProfile) GetCodecProfiles() []ModelModelCodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *ModelDeviceProfile) GetCodecProfilesOk() (*[]ModelModelCodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *ModelDeviceProfile) SetCodecProfiles(v []ModelModelCodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *ModelDeviceProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetResponseProfiles

`func (o *ModelDeviceProfile) GetResponseProfiles() []ModelModelResponseProfile`

GetResponseProfiles returns the ResponseProfiles field if non-nil, zero value otherwise.

### GetResponseProfilesOk

`func (o *ModelDeviceProfile) GetResponseProfilesOk() (*[]ModelModelResponseProfile, bool)`

GetResponseProfilesOk returns a tuple with the ResponseProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfiles

`func (o *ModelDeviceProfile) SetResponseProfiles(v []ModelModelResponseProfile)`

SetResponseProfiles sets ResponseProfiles field to given value.

### HasResponseProfiles

`func (o *ModelDeviceProfile) HasResponseProfiles() bool`

HasResponseProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *ModelDeviceProfile) GetSubtitleProfiles() []ModelModelSubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *ModelDeviceProfile) GetSubtitleProfilesOk() (*[]ModelModelSubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *ModelDeviceProfile) SetSubtitleProfiles(v []ModelModelSubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *ModelDeviceProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


