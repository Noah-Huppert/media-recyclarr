# ModelLiveTVApiEpgRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**Programs** | Pointer to [**[]ModelModelBaseItemDto**](ModelModelBaseItemDto.md) |  | [optional] 

## Methods

### NewModelLiveTVApiEpgRow

`func NewModelLiveTVApiEpgRow() *ModelLiveTVApiEpgRow`

NewModelLiveTVApiEpgRow instantiates a new ModelLiveTVApiEpgRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLiveTVApiEpgRowWithDefaults

`func NewModelLiveTVApiEpgRowWithDefaults() *ModelLiveTVApiEpgRow`

NewModelLiveTVApiEpgRowWithDefaults instantiates a new ModelLiveTVApiEpgRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ModelLiveTVApiEpgRow) GetChannel() ModelModelBaseItemDto`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ModelLiveTVApiEpgRow) GetChannelOk() (*ModelModelBaseItemDto, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ModelLiveTVApiEpgRow) SetChannel(v ModelModelBaseItemDto)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ModelLiveTVApiEpgRow) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetPrograms

`func (o *ModelLiveTVApiEpgRow) GetPrograms() []ModelModelBaseItemDto`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *ModelLiveTVApiEpgRow) GetProgramsOk() (*[]ModelModelBaseItemDto, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *ModelLiveTVApiEpgRow) SetPrograms(v []ModelModelBaseItemDto)`

SetPrograms sets Programs field to given value.

### HasPrograms

`func (o *ModelLiveTVApiEpgRow) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


