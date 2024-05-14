# ModelQueryResultSyncJobItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ModelModelSyncJobItem**](ModelModelSyncJobItem.md) |  | [optional] 
**TotalRecordCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelQueryResultSyncJobItem

`func NewModelQueryResultSyncJobItem() *ModelQueryResultSyncJobItem`

NewModelQueryResultSyncJobItem instantiates a new ModelQueryResultSyncJobItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelQueryResultSyncJobItemWithDefaults

`func NewModelQueryResultSyncJobItemWithDefaults() *ModelQueryResultSyncJobItem`

NewModelQueryResultSyncJobItemWithDefaults instantiates a new ModelQueryResultSyncJobItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ModelQueryResultSyncJobItem) GetItems() []ModelModelSyncJobItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ModelQueryResultSyncJobItem) GetItemsOk() (*[]ModelModelSyncJobItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ModelQueryResultSyncJobItem) SetItems(v []ModelModelSyncJobItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ModelQueryResultSyncJobItem) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *ModelQueryResultSyncJobItem) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *ModelQueryResultSyncJobItem) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *ModelQueryResultSyncJobItem) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *ModelQueryResultSyncJobItem) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


