# ModelFeatureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**FeatureType** | Pointer to [**ModelModelFeatureType**](ModelFeatureType.md) |  | [optional] 

## Methods

### NewModelFeatureInfo

`func NewModelFeatureInfo() *ModelFeatureInfo`

NewModelFeatureInfo instantiates a new ModelFeatureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelFeatureInfoWithDefaults

`func NewModelFeatureInfoWithDefaults() *ModelFeatureInfo`

NewModelFeatureInfoWithDefaults instantiates a new ModelFeatureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelFeatureInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelFeatureInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelFeatureInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelFeatureInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelFeatureInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelFeatureInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelFeatureInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelFeatureInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFeatureType

`func (o *ModelFeatureInfo) GetFeatureType() ModelModelFeatureType`

GetFeatureType returns the FeatureType field if non-nil, zero value otherwise.

### GetFeatureTypeOk

`func (o *ModelFeatureInfo) GetFeatureTypeOk() (*ModelModelFeatureType, bool)`

GetFeatureTypeOk returns a tuple with the FeatureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureType

`func (o *ModelFeatureInfo) SetFeatureType(v ModelModelFeatureType)`

SetFeatureType sets FeatureType field to given value.

### HasFeatureType

`func (o *ModelFeatureInfo) HasFeatureType() bool`

HasFeatureType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


