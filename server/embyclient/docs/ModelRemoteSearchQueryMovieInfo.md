# ModelRemoteSearchQueryMovieInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchInfo** | Pointer to [**ModelModelMovieInfo**](ModelMovieInfo.md) |  | [optional] 
**ItemId** | Pointer to **int64** |  | [optional] 
**SearchProviderName** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to **[]string** |  | [optional] 
**IncludeDisabledProviders** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelRemoteSearchQueryMovieInfo

`func NewModelRemoteSearchQueryMovieInfo() *ModelRemoteSearchQueryMovieInfo`

NewModelRemoteSearchQueryMovieInfo instantiates a new ModelRemoteSearchQueryMovieInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelRemoteSearchQueryMovieInfoWithDefaults

`func NewModelRemoteSearchQueryMovieInfoWithDefaults() *ModelRemoteSearchQueryMovieInfo`

NewModelRemoteSearchQueryMovieInfoWithDefaults instantiates a new ModelRemoteSearchQueryMovieInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchInfo

`func (o *ModelRemoteSearchQueryMovieInfo) GetSearchInfo() ModelModelMovieInfo`

GetSearchInfo returns the SearchInfo field if non-nil, zero value otherwise.

### GetSearchInfoOk

`func (o *ModelRemoteSearchQueryMovieInfo) GetSearchInfoOk() (*ModelModelMovieInfo, bool)`

GetSearchInfoOk returns a tuple with the SearchInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInfo

`func (o *ModelRemoteSearchQueryMovieInfo) SetSearchInfo(v ModelModelMovieInfo)`

SetSearchInfo sets SearchInfo field to given value.

### HasSearchInfo

`func (o *ModelRemoteSearchQueryMovieInfo) HasSearchInfo() bool`

HasSearchInfo returns a boolean if a field has been set.

### GetItemId

`func (o *ModelRemoteSearchQueryMovieInfo) GetItemId() int64`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelRemoteSearchQueryMovieInfo) GetItemIdOk() (*int64, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelRemoteSearchQueryMovieInfo) SetItemId(v int64)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelRemoteSearchQueryMovieInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSearchProviderName

`func (o *ModelRemoteSearchQueryMovieInfo) GetSearchProviderName() string`

GetSearchProviderName returns the SearchProviderName field if non-nil, zero value otherwise.

### GetSearchProviderNameOk

`func (o *ModelRemoteSearchQueryMovieInfo) GetSearchProviderNameOk() (*string, bool)`

GetSearchProviderNameOk returns a tuple with the SearchProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProviderName

`func (o *ModelRemoteSearchQueryMovieInfo) SetSearchProviderName(v string)`

SetSearchProviderName sets SearchProviderName field to given value.

### HasSearchProviderName

`func (o *ModelRemoteSearchQueryMovieInfo) HasSearchProviderName() bool`

HasSearchProviderName returns a boolean if a field has been set.

### GetProviders

`func (o *ModelRemoteSearchQueryMovieInfo) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ModelRemoteSearchQueryMovieInfo) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ModelRemoteSearchQueryMovieInfo) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ModelRemoteSearchQueryMovieInfo) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryMovieInfo) GetIncludeDisabledProviders() bool`

GetIncludeDisabledProviders returns the IncludeDisabledProviders field if non-nil, zero value otherwise.

### GetIncludeDisabledProvidersOk

`func (o *ModelRemoteSearchQueryMovieInfo) GetIncludeDisabledProvidersOk() (*bool, bool)`

GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryMovieInfo) SetIncludeDisabledProviders(v bool)`

SetIncludeDisabledProviders sets IncludeDisabledProviders field to given value.

### HasIncludeDisabledProviders

`func (o *ModelRemoteSearchQueryMovieInfo) HasIncludeDisabledProviders() bool`

HasIncludeDisabledProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


