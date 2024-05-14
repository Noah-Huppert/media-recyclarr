# ModelIOFileSystemEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelModelIOFileSystemEntryType**](ModelIOFileSystemEntryType.md) |  | [optional] 

## Methods

### NewModelIOFileSystemEntryInfo

`func NewModelIOFileSystemEntryInfo() *ModelIOFileSystemEntryInfo`

NewModelIOFileSystemEntryInfo instantiates a new ModelIOFileSystemEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelIOFileSystemEntryInfoWithDefaults

`func NewModelIOFileSystemEntryInfoWithDefaults() *ModelIOFileSystemEntryInfo`

NewModelIOFileSystemEntryInfoWithDefaults instantiates a new ModelIOFileSystemEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelIOFileSystemEntryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelIOFileSystemEntryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelIOFileSystemEntryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelIOFileSystemEntryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ModelIOFileSystemEntryInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ModelIOFileSystemEntryInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ModelIOFileSystemEntryInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ModelIOFileSystemEntryInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *ModelIOFileSystemEntryInfo) GetType() ModelModelIOFileSystemEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelIOFileSystemEntryInfo) GetTypeOk() (*ModelModelIOFileSystemEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelIOFileSystemEntryInfo) SetType(v ModelModelIOFileSystemEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelIOFileSystemEntryInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


