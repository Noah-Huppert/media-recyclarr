# ModelClientCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to **[]string** |  | [optional] 
**SupportedCommands** | Pointer to **[]string** |  | [optional] 
**SupportsMediaControl** | Pointer to **bool** |  | [optional] 
**PushToken** | Pointer to **string** |  | [optional] 
**PushTokenType** | Pointer to **string** |  | [optional] 
**SupportsSync** | Pointer to **bool** |  | [optional] 
**DeviceProfile** | Pointer to [**ModelModelDeviceProfile**](ModelDeviceProfile.md) |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelClientCapabilities

`func NewModelClientCapabilities() *ModelClientCapabilities`

NewModelClientCapabilities instantiates a new ModelClientCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelClientCapabilitiesWithDefaults

`func NewModelClientCapabilitiesWithDefaults() *ModelClientCapabilities`

NewModelClientCapabilitiesWithDefaults instantiates a new ModelClientCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *ModelClientCapabilities) GetPlayableMediaTypes() []string`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ModelClientCapabilities) GetPlayableMediaTypesOk() (*[]string, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ModelClientCapabilities) SetPlayableMediaTypes(v []string)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ModelClientCapabilities) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *ModelClientCapabilities) GetSupportedCommands() []string`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ModelClientCapabilities) GetSupportedCommandsOk() (*[]string, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ModelClientCapabilities) SetSupportedCommands(v []string)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ModelClientCapabilities) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *ModelClientCapabilities) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *ModelClientCapabilities) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *ModelClientCapabilities) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *ModelClientCapabilities) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetPushToken

`func (o *ModelClientCapabilities) GetPushToken() string`

GetPushToken returns the PushToken field if non-nil, zero value otherwise.

### GetPushTokenOk

`func (o *ModelClientCapabilities) GetPushTokenOk() (*string, bool)`

GetPushTokenOk returns a tuple with the PushToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushToken

`func (o *ModelClientCapabilities) SetPushToken(v string)`

SetPushToken sets PushToken field to given value.

### HasPushToken

`func (o *ModelClientCapabilities) HasPushToken() bool`

HasPushToken returns a boolean if a field has been set.

### GetPushTokenType

`func (o *ModelClientCapabilities) GetPushTokenType() string`

GetPushTokenType returns the PushTokenType field if non-nil, zero value otherwise.

### GetPushTokenTypeOk

`func (o *ModelClientCapabilities) GetPushTokenTypeOk() (*string, bool)`

GetPushTokenTypeOk returns a tuple with the PushTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTokenType

`func (o *ModelClientCapabilities) SetPushTokenType(v string)`

SetPushTokenType sets PushTokenType field to given value.

### HasPushTokenType

`func (o *ModelClientCapabilities) HasPushTokenType() bool`

HasPushTokenType returns a boolean if a field has been set.

### GetSupportsSync

`func (o *ModelClientCapabilities) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *ModelClientCapabilities) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *ModelClientCapabilities) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *ModelClientCapabilities) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *ModelClientCapabilities) GetDeviceProfile() ModelModelDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ModelClientCapabilities) GetDeviceProfileOk() (*ModelModelDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ModelClientCapabilities) SetDeviceProfile(v ModelModelDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ModelClientCapabilities) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### GetIconUrl

`func (o *ModelClientCapabilities) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ModelClientCapabilities) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ModelClientCapabilities) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ModelClientCapabilities) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAppId

`func (o *ModelClientCapabilities) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ModelClientCapabilities) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ModelClientCapabilities) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ModelClientCapabilities) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


