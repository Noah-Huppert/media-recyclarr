# ModelAuthenticationAuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**ModelModelUserDto**](ModelUserDto.md) |  | [optional] 
**SessionInfo** | Pointer to [**ModelModelSessionSessionInfo**](ModelSessionSessionInfo.md) |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAuthenticationAuthenticationResult

`func NewModelAuthenticationAuthenticationResult() *ModelAuthenticationAuthenticationResult`

NewModelAuthenticationAuthenticationResult instantiates a new ModelAuthenticationAuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAuthenticationAuthenticationResultWithDefaults

`func NewModelAuthenticationAuthenticationResultWithDefaults() *ModelAuthenticationAuthenticationResult`

NewModelAuthenticationAuthenticationResultWithDefaults instantiates a new ModelAuthenticationAuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ModelAuthenticationAuthenticationResult) GetUser() ModelModelUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelAuthenticationAuthenticationResult) GetUserOk() (*ModelModelUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelAuthenticationAuthenticationResult) SetUser(v ModelModelUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelAuthenticationAuthenticationResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSessionInfo

`func (o *ModelAuthenticationAuthenticationResult) GetSessionInfo() ModelModelSessionSessionInfo`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *ModelAuthenticationAuthenticationResult) GetSessionInfoOk() (*ModelModelSessionSessionInfo, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *ModelAuthenticationAuthenticationResult) SetSessionInfo(v ModelModelSessionSessionInfo)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *ModelAuthenticationAuthenticationResult) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### GetAccessToken

`func (o *ModelAuthenticationAuthenticationResult) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ModelAuthenticationAuthenticationResult) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ModelAuthenticationAuthenticationResult) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ModelAuthenticationAuthenticationResult) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetServerId

`func (o *ModelAuthenticationAuthenticationResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelAuthenticationAuthenticationResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelAuthenticationAuthenticationResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelAuthenticationAuthenticationResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


