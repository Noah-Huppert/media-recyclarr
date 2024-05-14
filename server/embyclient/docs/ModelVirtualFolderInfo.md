# ModelVirtualFolderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to **[]string** |  | [optional] 
**CollectionType** | Pointer to **string** |  | [optional] 
**LibraryOptions** | Pointer to [**ModelModelLibraryOptions**](ModelLibraryOptions.md) |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**PrimaryImageItemId** | Pointer to **string** |  | [optional] 
**RefreshProgress** | Pointer to **NullableFloat64** |  | [optional] 
**RefreshStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewModelVirtualFolderInfo

`func NewModelVirtualFolderInfo() *ModelVirtualFolderInfo`

NewModelVirtualFolderInfo instantiates a new ModelVirtualFolderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVirtualFolderInfoWithDefaults

`func NewModelVirtualFolderInfoWithDefaults() *ModelVirtualFolderInfo`

NewModelVirtualFolderInfoWithDefaults instantiates a new ModelVirtualFolderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelVirtualFolderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelVirtualFolderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelVirtualFolderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelVirtualFolderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocations

`func (o *ModelVirtualFolderInfo) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ModelVirtualFolderInfo) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ModelVirtualFolderInfo) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ModelVirtualFolderInfo) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetCollectionType

`func (o *ModelVirtualFolderInfo) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *ModelVirtualFolderInfo) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *ModelVirtualFolderInfo) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *ModelVirtualFolderInfo) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.

### GetLibraryOptions

`func (o *ModelVirtualFolderInfo) GetLibraryOptions() ModelModelLibraryOptions`

GetLibraryOptions returns the LibraryOptions field if non-nil, zero value otherwise.

### GetLibraryOptionsOk

`func (o *ModelVirtualFolderInfo) GetLibraryOptionsOk() (*ModelModelLibraryOptions, bool)`

GetLibraryOptionsOk returns a tuple with the LibraryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryOptions

`func (o *ModelVirtualFolderInfo) SetLibraryOptions(v ModelModelLibraryOptions)`

SetLibraryOptions sets LibraryOptions field to given value.

### HasLibraryOptions

`func (o *ModelVirtualFolderInfo) HasLibraryOptions() bool`

HasLibraryOptions returns a boolean if a field has been set.

### GetItemId

`func (o *ModelVirtualFolderInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelVirtualFolderInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelVirtualFolderInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelVirtualFolderInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetId

`func (o *ModelVirtualFolderInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVirtualFolderInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVirtualFolderInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVirtualFolderInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuid

`func (o *ModelVirtualFolderInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ModelVirtualFolderInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ModelVirtualFolderInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ModelVirtualFolderInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetPrimaryImageItemId

`func (o *ModelVirtualFolderInfo) GetPrimaryImageItemId() string`

GetPrimaryImageItemId returns the PrimaryImageItemId field if non-nil, zero value otherwise.

### GetPrimaryImageItemIdOk

`func (o *ModelVirtualFolderInfo) GetPrimaryImageItemIdOk() (*string, bool)`

GetPrimaryImageItemIdOk returns a tuple with the PrimaryImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageItemId

`func (o *ModelVirtualFolderInfo) SetPrimaryImageItemId(v string)`

SetPrimaryImageItemId sets PrimaryImageItemId field to given value.

### HasPrimaryImageItemId

`func (o *ModelVirtualFolderInfo) HasPrimaryImageItemId() bool`

HasPrimaryImageItemId returns a boolean if a field has been set.

### GetRefreshProgress

`func (o *ModelVirtualFolderInfo) GetRefreshProgress() float64`

GetRefreshProgress returns the RefreshProgress field if non-nil, zero value otherwise.

### GetRefreshProgressOk

`func (o *ModelVirtualFolderInfo) GetRefreshProgressOk() (*float64, bool)`

GetRefreshProgressOk returns a tuple with the RefreshProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshProgress

`func (o *ModelVirtualFolderInfo) SetRefreshProgress(v float64)`

SetRefreshProgress sets RefreshProgress field to given value.

### HasRefreshProgress

`func (o *ModelVirtualFolderInfo) HasRefreshProgress() bool`

HasRefreshProgress returns a boolean if a field has been set.

### SetRefreshProgressNil

`func (o *ModelVirtualFolderInfo) SetRefreshProgressNil(b bool)`

 SetRefreshProgressNil sets the value for RefreshProgress to be an explicit nil

### UnsetRefreshProgress
`func (o *ModelVirtualFolderInfo) UnsetRefreshProgress()`

UnsetRefreshProgress ensures that no value is present for RefreshProgress, not even an explicit nil
### GetRefreshStatus

`func (o *ModelVirtualFolderInfo) GetRefreshStatus() string`

GetRefreshStatus returns the RefreshStatus field if non-nil, zero value otherwise.

### GetRefreshStatusOk

`func (o *ModelVirtualFolderInfo) GetRefreshStatusOk() (*string, bool)`

GetRefreshStatusOk returns a tuple with the RefreshStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshStatus

`func (o *ModelVirtualFolderInfo) SetRefreshStatus(v string)`

SetRefreshStatus sets RefreshStatus field to given value.

### HasRefreshStatus

`func (o *ModelVirtualFolderInfo) HasRefreshStatus() bool`

HasRefreshStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


