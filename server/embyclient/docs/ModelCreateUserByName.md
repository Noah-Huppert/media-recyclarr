# ModelCreateUserByName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**CopyFromUserId** | Pointer to **string** |  | [optional] 
**UserCopyOptions** | Pointer to [**[]ModelModelLibraryUserCopyOptions**](ModelModelLibraryUserCopyOptions.md) |  | [optional] 

## Methods

### NewModelCreateUserByName

`func NewModelCreateUserByName() *ModelCreateUserByName`

NewModelCreateUserByName instantiates a new ModelCreateUserByName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCreateUserByNameWithDefaults

`func NewModelCreateUserByNameWithDefaults() *ModelCreateUserByName`

NewModelCreateUserByNameWithDefaults instantiates a new ModelCreateUserByName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelCreateUserByName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCreateUserByName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCreateUserByName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelCreateUserByName) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCopyFromUserId

`func (o *ModelCreateUserByName) GetCopyFromUserId() string`

GetCopyFromUserId returns the CopyFromUserId field if non-nil, zero value otherwise.

### GetCopyFromUserIdOk

`func (o *ModelCreateUserByName) GetCopyFromUserIdOk() (*string, bool)`

GetCopyFromUserIdOk returns a tuple with the CopyFromUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyFromUserId

`func (o *ModelCreateUserByName) SetCopyFromUserId(v string)`

SetCopyFromUserId sets CopyFromUserId field to given value.

### HasCopyFromUserId

`func (o *ModelCreateUserByName) HasCopyFromUserId() bool`

HasCopyFromUserId returns a boolean if a field has been set.

### GetUserCopyOptions

`func (o *ModelCreateUserByName) GetUserCopyOptions() []ModelModelLibraryUserCopyOptions`

GetUserCopyOptions returns the UserCopyOptions field if non-nil, zero value otherwise.

### GetUserCopyOptionsOk

`func (o *ModelCreateUserByName) GetUserCopyOptionsOk() (*[]ModelModelLibraryUserCopyOptions, bool)`

GetUserCopyOptionsOk returns a tuple with the UserCopyOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCopyOptions

`func (o *ModelCreateUserByName) SetUserCopyOptions(v []ModelModelLibraryUserCopyOptions)`

SetUserCopyOptions sets UserCopyOptions field to given value.

### HasUserCopyOptions

`func (o *ModelCreateUserByName) HasUserCopyOptions() bool`

HasUserCopyOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


