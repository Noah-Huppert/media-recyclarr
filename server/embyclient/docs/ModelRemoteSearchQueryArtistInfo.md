# ModelRemoteSearchQueryArtistInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchInfo** | Pointer to [**ModelModelArtistInfo**](ModelArtistInfo.md) |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**SearchProviderName** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to **[]string** |  | [optional] 
**IncludeDisabledProviders** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelRemoteSearchQueryArtistInfo

`func NewModelRemoteSearchQueryArtistInfo() *ModelRemoteSearchQueryArtistInfo`

NewModelRemoteSearchQueryArtistInfo instantiates a new ModelRemoteSearchQueryArtistInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRemoteSearchQueryArtistInfoWithDefaults

`func NewModelRemoteSearchQueryArtistInfoWithDefaults() *ModelRemoteSearchQueryArtistInfo`

NewModelRemoteSearchQueryArtistInfoWithDefaults instantiates a new ModelRemoteSearchQueryArtistInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchInfo

`func (o *ModelRemoteSearchQueryArtistInfo) GetSearchInfo() ModelModelArtistInfo`

GetSearchInfo returns the SearchInfo field if non-nil, zero value otherwise.

### GetSearchInfoOk

`func (o *ModelRemoteSearchQueryArtistInfo) GetSearchInfoOk() (*ModelModelArtistInfo, bool)`

GetSearchInfoOk returns a tuple with the SearchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInfo

`func (o *ModelRemoteSearchQueryArtistInfo) SetSearchInfo(v ModelModelArtistInfo)`

SetSearchInfo sets SearchInfo field to given value.

### HasSearchInfo

`func (o *ModelRemoteSearchQueryArtistInfo) HasSearchInfo() bool`

HasSearchInfo returns a boolean if a field has been set.

### GetItemId

`func (o *ModelRemoteSearchQueryArtistInfo) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelRemoteSearchQueryArtistInfo) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelRemoteSearchQueryArtistInfo) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelRemoteSearchQueryArtistInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSearchProviderName

`func (o *ModelRemoteSearchQueryArtistInfo) GetSearchProviderName() string`

GetSearchProviderName returns the SearchProviderName field if non-nil, zero value otherwise.

### GetSearchProviderNameOk

`func (o *ModelRemoteSearchQueryArtistInfo) GetSearchProviderNameOk() (*string, bool)`

GetSearchProviderNameOk returns a tuple with the SearchProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProviderName

`func (o *ModelRemoteSearchQueryArtistInfo) SetSearchProviderName(v string)`

SetSearchProviderName sets SearchProviderName field to given value.

### HasSearchProviderName

`func (o *ModelRemoteSearchQueryArtistInfo) HasSearchProviderName() bool`

HasSearchProviderName returns a boolean if a field has been set.

### GetProviders

`func (o *ModelRemoteSearchQueryArtistInfo) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ModelRemoteSearchQueryArtistInfo) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ModelRemoteSearchQueryArtistInfo) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ModelRemoteSearchQueryArtistInfo) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryArtistInfo) GetIncludeDisabledProviders() bool`

GetIncludeDisabledProviders returns the IncludeDisabledProviders field if non-nil, zero value otherwise.

### GetIncludeDisabledProvidersOk

`func (o *ModelRemoteSearchQueryArtistInfo) GetIncludeDisabledProvidersOk() (*bool, bool)`

GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryArtistInfo) SetIncludeDisabledProviders(v bool)`

SetIncludeDisabledProviders sets IncludeDisabledProviders field to given value.

### HasIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryArtistInfo) HasIncludeDisabledProviders() bool`

HasIncludeDisabledProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


