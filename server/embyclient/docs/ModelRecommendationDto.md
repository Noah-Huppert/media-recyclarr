# ModelRecommendationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ModelModelBaseItemDto**](ModelModelBaseItemDto.md) |  | [optional] 
**RecommendationType** | Pointer to [**ModelModelRecommendationType**](ModelRecommendationType.md) |  | [optional] 
**BaselineItemName** | Pointer to **string** |  | [optional] 
**CategoryId** | Pointer to **int64** |  | [optional] 

## Methods

### NewModelRecommendationDto

`func NewModelRecommendationDto() *ModelRecommendationDto`

NewModelRecommendationDto instantiates a new ModelRecommendationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRecommendationDtoWithDefaults

`func NewModelRecommendationDtoWithDefaults() *ModelRecommendationDto`

NewModelRecommendationDtoWithDefaults instantiates a new ModelRecommendationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ModelRecommendationDto) GetItems() []ModelModelBaseItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelRecommendationDto) GetItemsOk() (*[]ModelModelBaseItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelRecommendationDto) SetItems(v []ModelModelBaseItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *ModelRecommendationDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRecommendationType

`func (o *ModelRecommendationDto) GetRecommendationType() ModelModelRecommendationType`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *ModelRecommendationDto) GetRecommendationTypeOk() (*ModelModelRecommendationType, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *ModelRecommendationDto) SetRecommendationType(v ModelModelRecommendationType)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *ModelRecommendationDto) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetBaselineItemName

`func (o *ModelRecommendationDto) GetBaselineItemName() string`

GetBaselineItemName returns the BaselineItemName field if non-nil, zero value otherwise.

### GetBaselineItemNameOk

`func (o *ModelRecommendationDto) GetBaselineItemNameOk() (*string, bool)`

GetBaselineItemNameOk returns a tuple with the BaselineItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineItemName

`func (o *ModelRecommendationDto) SetBaselineItemName(v string)`

SetBaselineItemName sets BaselineItemName field to given value.

### HasBaselineItemName

`func (o *ModelRecommendationDto) HasBaselineItemName() bool`

HasBaselineItemName returns a boolean if a field has been set.

### GetCategoryId

`func (o *ModelRecommendationDto) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *ModelRecommendationDto) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *ModelRecommendationDto) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *ModelRecommendationDto) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


