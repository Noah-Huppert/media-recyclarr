# ModelLevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Ordinal** | Pointer to **NullableInt32** |  | [optional] 
**MaxBitRate** | Pointer to [**ModelModelBitRate**](ModelBitRate.md) |  | [optional] 
**MaxBitRateDisplay** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ResolutionRates** | Pointer to [**[]ModelModelResolutionWithRate**](ModelModelResolutionWithRate.md) |  | [optional] 
**ResolutionRateStrings** | Pointer to **[]string** |  | [optional] 
**ResolutionRatesDisplay** | Pointer to **string** |  | [optional] 

## Methods

### NewModelLevelInformation

`func NewModelLevelInformation() *ModelLevelInformation`

NewModelLevelInformation instantiates a new ModelLevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLevelInformationWithDefaults

`func NewModelLevelInformationWithDefaults() *ModelLevelInformation`

NewModelLevelInformationWithDefaults instantiates a new ModelLevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortName

`func (o *ModelLevelInformation) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ModelLevelInformation) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ModelLevelInformation) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ModelLevelInformation) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetDescription

`func (o *ModelLevelInformation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelLevelInformation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelLevelInformation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelLevelInformation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrdinal

`func (o *ModelLevelInformation) GetOrdinal() int32`

GetOrdinal returns the Ordinal field if non-nil, zero value otherwise.

### GetOrdinalOk

`func (o *ModelLevelInformation) GetOrdinalOk() (*int32, bool)`

GetOrdinalOk returns a tuple with the Ordinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdinal

`func (o *ModelLevelInformation) SetOrdinal(v int32)`

SetOrdinal sets Ordinal field to given value.

### HasOrdinal

`func (o *ModelLevelInformation) HasOrdinal() bool`

HasOrdinal returns a boolean if a field has been set.

### SetOrdinalNil

`func (o *ModelLevelInformation) SetOrdinalNil(b bool)`

 SetOrdinalNil sets the value for Ordinal to be an explicit nil

### UnsetOrdinal
`func (o *ModelLevelInformation) UnsetOrdinal()`

UnsetOrdinal ensures that no value is present for Ordinal, not even an explicit nil
### GetMaxBitRate

`func (o *ModelLevelInformation) GetMaxBitRate() ModelModelBitRate`

GetMaxBitRate returns the MaxBitRate field if non-nil, zero value otherwise.

### GetMaxBitRateOk

`func (o *ModelLevelInformation) GetMaxBitRateOk() (*ModelModelBitRate, bool)`

GetMaxBitRateOk returns a tuple with the MaxBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRate

`func (o *ModelLevelInformation) SetMaxBitRate(v ModelModelBitRate)`

SetMaxBitRate sets MaxBitRate field to given value.

### HasMaxBitRate

`func (o *ModelLevelInformation) HasMaxBitRate() bool`

HasMaxBitRate returns a boolean if a field has been set.

### GetMaxBitRateDisplay

`func (o *ModelLevelInformation) GetMaxBitRateDisplay() string`

GetMaxBitRateDisplay returns the MaxBitRateDisplay field if non-nil, zero value otherwise.

### GetMaxBitRateDisplayOk

`func (o *ModelLevelInformation) GetMaxBitRateDisplayOk() (*string, bool)`

GetMaxBitRateDisplayOk returns a tuple with the MaxBitRateDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBitRateDisplay

`func (o *ModelLevelInformation) SetMaxBitRateDisplay(v string)`

SetMaxBitRateDisplay sets MaxBitRateDisplay field to given value.

### HasMaxBitRateDisplay

`func (o *ModelLevelInformation) HasMaxBitRateDisplay() bool`

HasMaxBitRateDisplay returns a boolean if a field has been set.

### GetId

`func (o *ModelLevelInformation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelLevelInformation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelLevelInformation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelLevelInformation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResolutionRates

`func (o *ModelLevelInformation) GetResolutionRates() []ModelModelResolutionWithRate`

GetResolutionRates returns the ResolutionRates field if non-nil, zero value otherwise.

### GetResolutionRatesOk

`func (o *ModelLevelInformation) GetResolutionRatesOk() (*[]ModelModelResolutionWithRate, bool)`

GetResolutionRatesOk returns a tuple with the ResolutionRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRates

`func (o *ModelLevelInformation) SetResolutionRates(v []ModelModelResolutionWithRate)`

SetResolutionRates sets ResolutionRates field to given value.

### HasResolutionRates

`func (o *ModelLevelInformation) HasResolutionRates() bool`

HasResolutionRates returns a boolean if a field has been set.

### GetResolutionRateStrings

`func (o *ModelLevelInformation) GetResolutionRateStrings() []string`

GetResolutionRateStrings returns the ResolutionRateStrings field if non-nil, zero value otherwise.

### GetResolutionRateStringsOk

`func (o *ModelLevelInformation) GetResolutionRateStringsOk() (*[]string, bool)`

GetResolutionRateStringsOk returns a tuple with the ResolutionRateStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRateStrings

`func (o *ModelLevelInformation) SetResolutionRateStrings(v []string)`

SetResolutionRateStrings sets ResolutionRateStrings field to given value.

### HasResolutionRateStrings

`func (o *ModelLevelInformation) HasResolutionRateStrings() bool`

HasResolutionRateStrings returns a boolean if a field has been set.

### GetResolutionRatesDisplay

`func (o *ModelLevelInformation) GetResolutionRatesDisplay() string`

GetResolutionRatesDisplay returns the ResolutionRatesDisplay field if non-nil, zero value otherwise.

### GetResolutionRatesDisplayOk

`func (o *ModelLevelInformation) GetResolutionRatesDisplayOk() (*string, bool)`

GetResolutionRatesDisplayOk returns a tuple with the ResolutionRatesDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionRatesDisplay

`func (o *ModelLevelInformation) SetResolutionRatesDisplay(v string)`

SetResolutionRatesDisplay sets ResolutionRatesDisplay field to given value.

### HasResolutionRatesDisplay

`func (o *ModelLevelInformation) HasResolutionRatesDisplay() bool`

HasResolutionRatesDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


