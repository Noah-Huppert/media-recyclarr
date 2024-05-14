# ModelDisplayPreferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SortBy** | Pointer to **string** |  | [optional] 
**CustomPrefs** | Pointer to **map[string]string** |  | [optional] 
**SortOrder** | Pointer to [**ModelModelSortOrder**](ModelSortOrder.md) |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 

## Methods

### NewModelDisplayPreferences

`func NewModelDisplayPreferences() *ModelDisplayPreferences`

NewModelDisplayPreferences instantiates a new ModelDisplayPreferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDisplayPreferencesWithDefaults

`func NewModelDisplayPreferencesWithDefaults() *ModelDisplayPreferences`

NewModelDisplayPreferencesWithDefaults instantiates a new ModelDisplayPreferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelDisplayPreferences) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelDisplayPreferences) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelDisplayPreferences) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelDisplayPreferences) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSortBy

`func (o *ModelDisplayPreferences) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *ModelDisplayPreferences) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *ModelDisplayPreferences) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *ModelDisplayPreferences) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetCustomPrefs

`func (o *ModelDisplayPreferences) GetCustomPrefs() map[string]string`

GetCustomPrefs returns the CustomPrefs field if non-nil, zero value otherwise.

### GetCustomPrefsOk

`func (o *ModelDisplayPreferences) GetCustomPrefsOk() (*map[string]string, bool)`

GetCustomPrefsOk returns a tuple with the CustomPrefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrefs

`func (o *ModelDisplayPreferences) SetCustomPrefs(v map[string]string)`

SetCustomPrefs sets CustomPrefs field to given value.

### HasCustomPrefs

`func (o *ModelDisplayPreferences) HasCustomPrefs() bool`

HasCustomPrefs returns a boolean if a field has been set.

### GetSortOrder

`func (o *ModelDisplayPreferences) GetSortOrder() ModelModelSortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ModelDisplayPreferences) GetSortOrderOk() (*ModelModelSortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ModelDisplayPreferences) SetSortOrder(v ModelModelSortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ModelDisplayPreferences) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetClient

`func (o *ModelDisplayPreferences) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ModelDisplayPreferences) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ModelDisplayPreferences) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ModelDisplayPreferences) HasClient() bool`

HasClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


