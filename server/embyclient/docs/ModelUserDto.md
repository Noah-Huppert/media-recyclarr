# ModelUserDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**ConnectUserName** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**ConnectLinkType** | Pointer to [**ModelModelConnectUserLinkType**](ModelConnectUserLinkType.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**PrimaryImageTag** | Pointer to **string** |  | [optional] 
**HasPassword** | Pointer to **bool** |  | [optional] 
**HasConfiguredPassword** | Pointer to **bool** |  | [optional] 
**EnableAutoLogin** | Pointer to **NullableBool** |  | [optional] 
**LastLoginDate** | Pointer to **NullableTime** |  | [optional] 
**LastActivityDate** | Pointer to **NullableTime** |  | [optional] 
**Configuration** | Pointer to [**ModelModelUserConfiguration**](ModelUserConfiguration.md) |  | [optional] 
**Policy** | Pointer to [**ModelModelUserPolicy**](ModelUserPolicy.md) |  | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** |  | [optional] 
**HasConfiguredEasyPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelUserDto

`func NewModelUserDto() *ModelUserDto`

NewModelUserDto instantiates a new ModelUserDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserDtoWithDefaults

`func NewModelUserDtoWithDefaults() *ModelUserDto`

NewModelUserDtoWithDefaults instantiates a new ModelUserDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelUserDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelUserDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelUserDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelUserDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerId

`func (o *ModelUserDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelUserDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelUserDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelUserDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *ModelUserDto) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ModelUserDto) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ModelUserDto) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ModelUserDto) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPrefix

`func (o *ModelUserDto) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ModelUserDto) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ModelUserDto) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ModelUserDto) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetConnectUserName

`func (o *ModelUserDto) GetConnectUserName() string`

GetConnectUserName returns the ConnectUserName field if non-nil, zero value otherwise.

### GetConnectUserNameOk

`func (o *ModelUserDto) GetConnectUserNameOk() (*string, bool)`

GetConnectUserNameOk returns a tuple with the ConnectUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectUserName

`func (o *ModelUserDto) SetConnectUserName(v string)`

SetConnectUserName sets ConnectUserName field to given value.

### HasConnectUserName

`func (o *ModelUserDto) HasConnectUserName() bool`

HasConnectUserName returns a boolean if a field has been set.

### GetDateCreated

`func (o *ModelUserDto) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ModelUserDto) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ModelUserDto) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ModelUserDto) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *ModelUserDto) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *ModelUserDto) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetConnectLinkType

`func (o *ModelUserDto) GetConnectLinkType() ModelModelConnectUserLinkType`

GetConnectLinkType returns the ConnectLinkType field if non-nil, zero value otherwise.

### GetConnectLinkTypeOk

`func (o *ModelUserDto) GetConnectLinkTypeOk() (*ModelModelConnectUserLinkType, bool)`

GetConnectLinkTypeOk returns a tuple with the ConnectLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectLinkType

`func (o *ModelUserDto) SetConnectLinkType(v ModelModelConnectUserLinkType)`

SetConnectLinkType sets ConnectLinkType field to given value.

### HasConnectLinkType

`func (o *ModelUserDto) HasConnectLinkType() bool`

HasConnectLinkType returns a boolean if a field has been set.

### GetId

`func (o *ModelUserDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelUserDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelUserDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelUserDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *ModelUserDto) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *ModelUserDto) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *ModelUserDto) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *ModelUserDto) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### GetHasPassword

`func (o *ModelUserDto) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *ModelUserDto) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *ModelUserDto) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *ModelUserDto) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetHasConfiguredPassword

`func (o *ModelUserDto) GetHasConfiguredPassword() bool`

GetHasConfiguredPassword returns the HasConfiguredPassword field if non-nil, zero value otherwise.

### GetHasConfiguredPasswordOk

`func (o *ModelUserDto) GetHasConfiguredPasswordOk() (*bool, bool)`

GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredPassword

`func (o *ModelUserDto) SetHasConfiguredPassword(v bool)`

SetHasConfiguredPassword sets HasConfiguredPassword field to given value.

### HasHasConfiguredPassword

`func (o *ModelUserDto) HasHasConfiguredPassword() bool`

HasHasConfiguredPassword returns a boolean if a field has been set.

### GetEnableAutoLogin

`func (o *ModelUserDto) GetEnableAutoLogin() bool`

GetEnableAutoLogin returns the EnableAutoLogin field if non-nil, zero value otherwise.

### GetEnableAutoLoginOk

`func (o *ModelUserDto) GetEnableAutoLoginOk() (*bool, bool)`

GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoLogin

`func (o *ModelUserDto) SetEnableAutoLogin(v bool)`

SetEnableAutoLogin sets EnableAutoLogin field to given value.

### HasEnableAutoLogin

`func (o *ModelUserDto) HasEnableAutoLogin() bool`

HasEnableAutoLogin returns a boolean if a field has been set.

### SetEnableAutoLoginNil

`func (o *ModelUserDto) SetEnableAutoLoginNil(b bool)`

 SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil

### UnsetEnableAutoLogin
`func (o *ModelUserDto) UnsetEnableAutoLogin()`

UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
### GetLastLoginDate

`func (o *ModelUserDto) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *ModelUserDto) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *ModelUserDto) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *ModelUserDto) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### SetLastLoginDateNil

`func (o *ModelUserDto) SetLastLoginDateNil(b bool)`

 SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil

### UnsetLastLoginDate
`func (o *ModelUserDto) UnsetLastLoginDate()`

UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
### GetLastActivityDate

`func (o *ModelUserDto) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *ModelUserDto) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *ModelUserDto) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *ModelUserDto) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### SetLastActivityDateNil

`func (o *ModelUserDto) SetLastActivityDateNil(b bool)`

 SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil

### UnsetLastActivityDate
`func (o *ModelUserDto) UnsetLastActivityDate()`

UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
### GetConfiguration

`func (o *ModelUserDto) GetConfiguration() ModelModelUserConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ModelUserDto) GetConfigurationOk() (*ModelModelUserConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ModelUserDto) SetConfiguration(v ModelModelUserConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ModelUserDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetPolicy

`func (o *ModelUserDto) GetPolicy() ModelModelUserPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ModelUserDto) GetPolicyOk() (*ModelModelUserPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ModelUserDto) SetPolicy(v ModelModelUserPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ModelUserDto) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetPrimaryImageAspectRatio

`func (o *ModelUserDto) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *ModelUserDto) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *ModelUserDto) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *ModelUserDto) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *ModelUserDto) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *ModelUserDto) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
### GetHasConfiguredEasyPassword

`func (o *ModelUserDto) GetHasConfiguredEasyPassword() bool`

GetHasConfiguredEasyPassword returns the HasConfiguredEasyPassword field if non-nil, zero value otherwise.

### GetHasConfiguredEasyPasswordOk

`func (o *ModelUserDto) GetHasConfiguredEasyPasswordOk() (*bool, bool)`

GetHasConfiguredEasyPasswordOk returns a tuple with the HasConfiguredEasyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredEasyPassword

`func (o *ModelUserDto) SetHasConfiguredEasyPassword(v bool)`

SetHasConfiguredEasyPassword sets HasConfiguredEasyPassword field to given value.

### HasHasConfiguredEasyPassword

`func (o *ModelUserDto) HasHasConfiguredEasyPassword() bool`

HasHasConfiguredEasyPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


