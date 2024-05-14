# ModelUserNotificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotifierKey** | Pointer to **string** |  | [optional] 
**SetupModuleUrl** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**PluginId** | Pointer to **string** |  | [optional] 
**FriendlyName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 
**DeviceIds** | Pointer to **[]string** |  | [optional] 
**LibraryIds** | Pointer to **[]string** |  | [optional] 
**EventIds** | Pointer to **[]string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**IsSelfNotification** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelUserNotificationInfo

`func NewModelUserNotificationInfo() *ModelUserNotificationInfo`

NewModelUserNotificationInfo instantiates a new ModelUserNotificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserNotificationInfoWithDefaults

`func NewModelUserNotificationInfoWithDefaults() *ModelUserNotificationInfo`

NewModelUserNotificationInfoWithDefaults instantiates a new ModelUserNotificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifierKey

`func (o *ModelUserNotificationInfo) GetNotifierKey() string`

GetNotifierKey returns the NotifierKey field if non-nil, zero value otherwise.

### GetNotifierKeyOk

`func (o *ModelUserNotificationInfo) GetNotifierKeyOk() (*string, bool)`

GetNotifierKeyOk returns a tuple with the NotifierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierKey

`func (o *ModelUserNotificationInfo) SetNotifierKey(v string)`

SetNotifierKey sets NotifierKey field to given value.

### HasNotifierKey

`func (o *ModelUserNotificationInfo) HasNotifierKey() bool`

HasNotifierKey returns a boolean if a field has been set.

### GetSetupModuleUrl

`func (o *ModelUserNotificationInfo) GetSetupModuleUrl() string`

GetSetupModuleUrl returns the SetupModuleUrl field if non-nil, zero value otherwise.

### GetSetupModuleUrlOk

`func (o *ModelUserNotificationInfo) GetSetupModuleUrlOk() (*string, bool)`

GetSetupModuleUrlOk returns a tuple with the SetupModuleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupModuleUrl

`func (o *ModelUserNotificationInfo) SetSetupModuleUrl(v string)`

SetSetupModuleUrl sets SetupModuleUrl field to given value.

### HasSetupModuleUrl

`func (o *ModelUserNotificationInfo) HasSetupModuleUrl() bool`

HasSetupModuleUrl returns a boolean if a field has been set.

### GetServiceName

`func (o *ModelUserNotificationInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ModelUserNotificationInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ModelUserNotificationInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ModelUserNotificationInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetPluginId

`func (o *ModelUserNotificationInfo) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *ModelUserNotificationInfo) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *ModelUserNotificationInfo) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *ModelUserNotificationInfo) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetFriendlyName

`func (o *ModelUserNotificationInfo) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ModelUserNotificationInfo) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ModelUserNotificationInfo) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ModelUserNotificationInfo) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetId

`func (o *ModelUserNotificationInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelUserNotificationInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelUserNotificationInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelUserNotificationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelUserNotificationInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelUserNotificationInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelUserNotificationInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelUserNotificationInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUserIds

`func (o *ModelUserNotificationInfo) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *ModelUserNotificationInfo) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *ModelUserNotificationInfo) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *ModelUserNotificationInfo) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetDeviceIds

`func (o *ModelUserNotificationInfo) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *ModelUserNotificationInfo) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *ModelUserNotificationInfo) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *ModelUserNotificationInfo) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetLibraryIds

`func (o *ModelUserNotificationInfo) GetLibraryIds() []string`

GetLibraryIds returns the LibraryIds field if non-nil, zero value otherwise.

### GetLibraryIdsOk

`func (o *ModelUserNotificationInfo) GetLibraryIdsOk() (*[]string, bool)`

GetLibraryIdsOk returns a tuple with the LibraryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryIds

`func (o *ModelUserNotificationInfo) SetLibraryIds(v []string)`

SetLibraryIds sets LibraryIds field to given value.

### HasLibraryIds

`func (o *ModelUserNotificationInfo) HasLibraryIds() bool`

HasLibraryIds returns a boolean if a field has been set.

### GetEventIds

`func (o *ModelUserNotificationInfo) GetEventIds() []string`

GetEventIds returns the EventIds field if non-nil, zero value otherwise.

### GetEventIdsOk

`func (o *ModelUserNotificationInfo) GetEventIdsOk() (*[]string, bool)`

GetEventIdsOk returns a tuple with the EventIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIds

`func (o *ModelUserNotificationInfo) SetEventIds(v []string)`

SetEventIds sets EventIds field to given value.

### HasEventIds

`func (o *ModelUserNotificationInfo) HasEventIds() bool`

HasEventIds returns a boolean if a field has been set.

### GetUserId

`func (o *ModelUserNotificationInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelUserNotificationInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelUserNotificationInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelUserNotificationInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsSelfNotification

`func (o *ModelUserNotificationInfo) GetIsSelfNotification() bool`

GetIsSelfNotification returns the IsSelfNotification field if non-nil, zero value otherwise.

### GetIsSelfNotificationOk

`func (o *ModelUserNotificationInfo) GetIsSelfNotificationOk() (*bool, bool)`

GetIsSelfNotificationOk returns a tuple with the IsSelfNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfNotification

`func (o *ModelUserNotificationInfo) SetIsSelfNotification(v bool)`

SetIsSelfNotification sets IsSelfNotification field to given value.

### HasIsSelfNotification

`func (o *ModelUserNotificationInfo) HasIsSelfNotification() bool`

HasIsSelfNotification returns a boolean if a field has been set.

### GetOptions

`func (o *ModelUserNotificationInfo) GetOptions() map[string]string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModelUserNotificationInfo) GetOptionsOk() (*map[string]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModelUserNotificationInfo) SetOptions(v map[string]string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModelUserNotificationInfo) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


