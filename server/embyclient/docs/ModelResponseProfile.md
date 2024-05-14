# ModelResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** |  | [optional] 
**AudioCodec** | Pointer to **string** |  | [optional] 
**VideoCodec** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelModelDlnaProfileType**](ModelDlnaProfileType.md) |  | [optional] 
**OrgPn** | Pointer to **string** |  | [optional] 
**MimeType** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to [**[]ModelModelProfileCondition**](ModelModelProfileCondition.md) |  | [optional] 

## Methods

### NewModelResponseProfile

`func NewModelResponseProfile() *ModelResponseProfile`

NewModelResponseProfile instantiates a new ModelResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelResponseProfileWithDefaults

`func NewModelResponseProfileWithDefaults() *ModelResponseProfile`

NewModelResponseProfileWithDefaults instantiates a new ModelResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ModelResponseProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelResponseProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelResponseProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelResponseProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetAudioCodec

`func (o *ModelResponseProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ModelResponseProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ModelResponseProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ModelResponseProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### GetVideoCodec

`func (o *ModelResponseProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ModelResponseProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ModelResponseProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ModelResponseProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### GetType

`func (o *ModelResponseProfile) GetType() ModelModelDlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelResponseProfile) GetTypeOk() (*ModelModelDlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelResponseProfile) SetType(v ModelModelDlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelResponseProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrgPn

`func (o *ModelResponseProfile) GetOrgPn() string`

GetOrgPn returns the OrgPn field if non-nil, zero value otherwise.

### GetOrgPnOk

`func (o *ModelResponseProfile) GetOrgPnOk() (*string, bool)`

GetOrgPnOk returns a tuple with the OrgPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPn

`func (o *ModelResponseProfile) SetOrgPn(v string)`

SetOrgPn sets OrgPn field to given value.

### HasOrgPn

`func (o *ModelResponseProfile) HasOrgPn() bool`

HasOrgPn returns a boolean if a field has been set.

### GetMimeType

`func (o *ModelResponseProfile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ModelResponseProfile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ModelResponseProfile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ModelResponseProfile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetConditions

`func (o *ModelResponseProfile) GetConditions() []ModelModelProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ModelResponseProfile) GetConditionsOk() (*[]ModelModelProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ModelResponseProfile) SetConditions(v []ModelModelProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ModelResponseProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


