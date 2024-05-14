# ModelQueryResultBaseItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ModelModelBaseItemDto**](ModelModelBaseItemDto.md) |  | [optional] 
**TotalRecordCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelQueryResultBaseItemDto

`func NewModelQueryResultBaseItemDto() *ModelQueryResultBaseItemDto`

NewModelQueryResultBaseItemDto instantiates a new ModelQueryResultBaseItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelQueryResultBaseItemDtoWithDefaults

`func NewModelQueryResultBaseItemDtoWithDefaults() *ModelQueryResultBaseItemDto`

NewModelQueryResultBaseItemDtoWithDefaults instantiates a new ModelQueryResultBaseItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ModelQueryResultBaseItemDto) GetItems() []ModelModelBaseItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelQueryResultBaseItemDto) GetItemsOk() (*[]ModelModelBaseItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelQueryResultBaseItemDto) SetItems(v []ModelModelBaseItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *ModelQueryResultBaseItemDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *ModelQueryResultBaseItemDto) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *ModelQueryResultBaseItemDto) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *ModelQueryResultBaseItemDto) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *ModelQueryResultBaseItemDto) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


