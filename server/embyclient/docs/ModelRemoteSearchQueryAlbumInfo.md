# ModelRemoteSearchQueryAlbumInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchInfo** | Pointer to [**ModelModelAlbumInfo**](ModelAlbumInfo.md) |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**SearchProviderName** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to **[]string** |  | [optional] 
**IncludeDisabledProviders** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelRemoteSearchQueryAlbumInfo

`func NewModelRemoteSearchQueryAlbumInfo() *ModelRemoteSearchQueryAlbumInfo`

NewModelRemoteSearchQueryAlbumInfo instantiates a new ModelRemoteSearchQueryAlbumInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRemoteSearchQueryAlbumInfoWithDefaults

`func NewModelRemoteSearchQueryAlbumInfoWithDefaults() *ModelRemoteSearchQueryAlbumInfo`

NewModelRemoteSearchQueryAlbumInfoWithDefaults instantiates a new ModelRemoteSearchQueryAlbumInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchInfo

`func (o *ModelRemoteSearchQueryAlbumInfo) GetSearchInfo() ModelModelAlbumInfo`

GetSearchInfo returns the SearchInfo field if non-nil, zero value otherwise.

### GetSearchInfoOk

`func (o *ModelRemoteSearchQueryAlbumInfo) GetSearchInfoOk() (*ModelModelAlbumInfo, bool)`

GetSearchInfoOk returns a tuple with the SearchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInfo

`func (o *ModelRemoteSearchQueryAlbumInfo) SetSearchInfo(v ModelModelAlbumInfo)`

SetSearchInfo sets SearchInfo field to given value.

### HasSearchInfo

`func (o *ModelRemoteSearchQueryAlbumInfo) HasSearchInfo() bool`

HasSearchInfo returns a boolean if a field has been set.

### GetItemId

`func (o *ModelRemoteSearchQueryAlbumInfo) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelRemoteSearchQueryAlbumInfo) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelRemoteSearchQueryAlbumInfo) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelRemoteSearchQueryAlbumInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSearchProviderName

`func (o *ModelRemoteSearchQueryAlbumInfo) GetSearchProviderName() string`

GetSearchProviderName returns the SearchProviderName field if non-nil, zero value otherwise.

### GetSearchProviderNameOk

`func (o *ModelRemoteSearchQueryAlbumInfo) GetSearchProviderNameOk() (*string, bool)`

GetSearchProviderNameOk returns a tuple with the SearchProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProviderName

`func (o *ModelRemoteSearchQueryAlbumInfo) SetSearchProviderName(v string)`

SetSearchProviderName sets SearchProviderName field to given value.

### HasSearchProviderName

`func (o *ModelRemoteSearchQueryAlbumInfo) HasSearchProviderName() bool`

HasSearchProviderName returns a boolean if a field has been set.

### GetProviders

`func (o *ModelRemoteSearchQueryAlbumInfo) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ModelRemoteSearchQueryAlbumInfo) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ModelRemoteSearchQueryAlbumInfo) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ModelRemoteSearchQueryAlbumInfo) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryAlbumInfo) GetIncludeDisabledProviders() bool`

GetIncludeDisabledProviders returns the IncludeDisabledProviders field if non-nil, zero value otherwise.

### GetIncludeDisabledProvidersOk

`func (o *ModelRemoteSearchQueryAlbumInfo) GetIncludeDisabledProvidersOk() (*bool, bool)`

GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryAlbumInfo) SetIncludeDisabledProviders(v bool)`

SetIncludeDisabledProviders sets IncludeDisabledProviders field to given value.

### HasIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryAlbumInfo) HasIncludeDisabledProviders() bool`

HasIncludeDisabledProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


