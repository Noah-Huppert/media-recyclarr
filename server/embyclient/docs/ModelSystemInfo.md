# ModelSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SystemUpdateLevel** | Pointer to [**ModelModelPackageVersionClass**](ModelPackageVersionClass.md) |  | [optional] 
**OperatingSystemDisplayName** | Pointer to **string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**HasPendingRestart** | Pointer to **bool** |  | [optional] 
**IsShuttingDown** | Pointer to **bool** |  | [optional] 
**OperatingSystem** | Pointer to **string** |  | [optional] 
**SupportsLibraryMonitor** | Pointer to **bool** |  | [optional] 
**SupportsLocalPortConfiguration** | Pointer to **bool** |  | [optional] 
**SupportsWakeServer** | Pointer to **bool** |  | [optional] 
**WebSocketPortNumber** | Pointer to **int32** |  | [optional] 
**CompletedInstallations** | Pointer to [**[]ModelModelInstallationInfo**](ModelModelInstallationInfo.md) |  | [optional] 
**CanSelfRestart** | Pointer to **bool** |  | [optional] 
**CanSelfUpdate** | Pointer to **bool** |  | [optional] 
**CanLaunchWebBrowser** | Pointer to **bool** |  | [optional] 
**ProgramDataPath** | Pointer to **string** |  | [optional] 
**ItemsByNamePath** | Pointer to **string** |  | [optional] 
**CachePath** | Pointer to **string** |  | [optional] 
**LogPath** | Pointer to **string** |  | [optional] 
**InternalMetadataPath** | Pointer to **string** |  | [optional] 
**TranscodingTempPath** | Pointer to **string** |  | [optional] 
**HttpServerPortNumber** | Pointer to **int32** |  | [optional] 
**SupportsHttps** | Pointer to **bool** |  | [optional] 
**HttpsPortNumber** | Pointer to **int32** |  | [optional] 
**HasUpdateAvailable** | Pointer to **bool** |  | [optional] 
**SupportsAutoRunAtStartup** | Pointer to **bool** |  | [optional] 
**HardwareAccelerationRequiresPremiere** | Pointer to **bool** |  | [optional] 
**LocalAddress** | Pointer to **string** |  | [optional] 
**LocalAddresses** | Pointer to **[]string** |  | [optional] 
**WanAddress** | Pointer to **string** |  | [optional] 
**RemoteAddresses** | Pointer to **[]string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewModelSystemInfo

`func NewModelSystemInfo() *ModelSystemInfo`

NewModelSystemInfo instantiates a new ModelSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSystemInfoWithDefaults

`func NewModelSystemInfoWithDefaults() *ModelSystemInfo`

NewModelSystemInfoWithDefaults instantiates a new ModelSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystemUpdateLevel

`func (o *ModelSystemInfo) GetSystemUpdateLevel() ModelModelPackageVersionClass`

GetSystemUpdateLevel returns the SystemUpdateLevel field if non-nil, zero value otherwise.

### GetSystemUpdateLevelOk

`func (o *ModelSystemInfo) GetSystemUpdateLevelOk() (*ModelModelPackageVersionClass, bool)`

GetSystemUpdateLevelOk returns a tuple with the SystemUpdateLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUpdateLevel

`func (o *ModelSystemInfo) SetSystemUpdateLevel(v ModelModelPackageVersionClass)`

SetSystemUpdateLevel sets SystemUpdateLevel field to given value.

### HasSystemUpdateLevel

`func (o *ModelSystemInfo) HasSystemUpdateLevel() bool`

HasSystemUpdateLevel returns a boolean if a field has been set.

### GetOperatingSystemDisplayName

`func (o *ModelSystemInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *ModelSystemInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *ModelSystemInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *ModelSystemInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### GetPackageName

`func (o *ModelSystemInfo) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *ModelSystemInfo) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *ModelSystemInfo) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *ModelSystemInfo) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetHasPendingRestart

`func (o *ModelSystemInfo) GetHasPendingRestart() bool`

GetHasPendingRestart returns the HasPendingRestart field if non-nil, zero value otherwise.

### GetHasPendingRestartOk

`func (o *ModelSystemInfo) GetHasPendingRestartOk() (*bool, bool)`

GetHasPendingRestartOk returns a tuple with the HasPendingRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingRestart

`func (o *ModelSystemInfo) SetHasPendingRestart(v bool)`

SetHasPendingRestart sets HasPendingRestart field to given value.

### HasHasPendingRestart

`func (o *ModelSystemInfo) HasHasPendingRestart() bool`

HasHasPendingRestart returns a boolean if a field has been set.

### GetIsShuttingDown

`func (o *ModelSystemInfo) GetIsShuttingDown() bool`

GetIsShuttingDown returns the IsShuttingDown field if non-nil, zero value otherwise.

### GetIsShuttingDownOk

`func (o *ModelSystemInfo) GetIsShuttingDownOk() (*bool, bool)`

GetIsShuttingDownOk returns a tuple with the IsShuttingDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShuttingDown

`func (o *ModelSystemInfo) SetIsShuttingDown(v bool)`

SetIsShuttingDown sets IsShuttingDown field to given value.

### HasIsShuttingDown

`func (o *ModelSystemInfo) HasIsShuttingDown() bool`

HasIsShuttingDown returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *ModelSystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *ModelSystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *ModelSystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *ModelSystemInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetSupportsLibraryMonitor

`func (o *ModelSystemInfo) GetSupportsLibraryMonitor() bool`

GetSupportsLibraryMonitor returns the SupportsLibraryMonitor field if non-nil, zero value otherwise.

### GetSupportsLibraryMonitorOk

`func (o *ModelSystemInfo) GetSupportsLibraryMonitorOk() (*bool, bool)`

GetSupportsLibraryMonitorOk returns a tuple with the SupportsLibraryMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLibraryMonitor

`func (o *ModelSystemInfo) SetSupportsLibraryMonitor(v bool)`

SetSupportsLibraryMonitor sets SupportsLibraryMonitor field to given value.

### HasSupportsLibraryMonitor

`func (o *ModelSystemInfo) HasSupportsLibraryMonitor() bool`

HasSupportsLibraryMonitor returns a boolean if a field has been set.

### GetSupportsLocalPortConfiguration

`func (o *ModelSystemInfo) GetSupportsLocalPortConfiguration() bool`

GetSupportsLocalPortConfiguration returns the SupportsLocalPortConfiguration field if non-nil, zero value otherwise.

### GetSupportsLocalPortConfigurationOk

`func (o *ModelSystemInfo) GetSupportsLocalPortConfigurationOk() (*bool, bool)`

GetSupportsLocalPortConfigurationOk returns a tuple with the SupportsLocalPortConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLocalPortConfiguration

`func (o *ModelSystemInfo) SetSupportsLocalPortConfiguration(v bool)`

SetSupportsLocalPortConfiguration sets SupportsLocalPortConfiguration field to given value.

### HasSupportsLocalPortConfiguration

`func (o *ModelSystemInfo) HasSupportsLocalPortConfiguration() bool`

HasSupportsLocalPortConfiguration returns a boolean if a field has been set.

### GetSupportsWakeServer

`func (o *ModelSystemInfo) GetSupportsWakeServer() bool`

GetSupportsWakeServer returns the SupportsWakeServer field if non-nil, zero value otherwise.

### GetSupportsWakeServerOk

`func (o *ModelSystemInfo) GetSupportsWakeServerOk() (*bool, bool)`

GetSupportsWakeServerOk returns a tuple with the SupportsWakeServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsWakeServer

`func (o *ModelSystemInfo) SetSupportsWakeServer(v bool)`

SetSupportsWakeServer sets SupportsWakeServer field to given value.

### HasSupportsWakeServer

`func (o *ModelSystemInfo) HasSupportsWakeServer() bool`

HasSupportsWakeServer returns a boolean if a field has been set.

### GetWebSocketPortNumber

`func (o *ModelSystemInfo) GetWebSocketPortNumber() int32`

GetWebSocketPortNumber returns the WebSocketPortNumber field if non-nil, zero value otherwise.

### GetWebSocketPortNumberOk

`func (o *ModelSystemInfo) GetWebSocketPortNumberOk() (*int32, bool)`

GetWebSocketPortNumberOk returns a tuple with the WebSocketPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSocketPortNumber

`func (o *ModelSystemInfo) SetWebSocketPortNumber(v int32)`

SetWebSocketPortNumber sets WebSocketPortNumber field to given value.

### HasWebSocketPortNumber

`func (o *ModelSystemInfo) HasWebSocketPortNumber() bool`

HasWebSocketPortNumber returns a boolean if a field has been set.

### GetCompletedInstallations

`func (o *ModelSystemInfo) GetCompletedInstallations() []ModelModelInstallationInfo`

GetCompletedInstallations returns the CompletedInstallations field if non-nil, zero value otherwise.

### GetCompletedInstallationsOk

`func (o *ModelSystemInfo) GetCompletedInstallationsOk() (*[]ModelModelInstallationInfo, bool)`

GetCompletedInstallationsOk returns a tuple with the CompletedInstallations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedInstallations

`func (o *ModelSystemInfo) SetCompletedInstallations(v []ModelModelInstallationInfo)`

SetCompletedInstallations sets CompletedInstallations field to given value.

### HasCompletedInstallations

`func (o *ModelSystemInfo) HasCompletedInstallations() bool`

HasCompletedInstallations returns a boolean if a field has been set.

### GetCanSelfRestart

`func (o *ModelSystemInfo) GetCanSelfRestart() bool`

GetCanSelfRestart returns the CanSelfRestart field if non-nil, zero value otherwise.

### GetCanSelfRestartOk

`func (o *ModelSystemInfo) GetCanSelfRestartOk() (*bool, bool)`

GetCanSelfRestartOk returns a tuple with the CanSelfRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelfRestart

`func (o *ModelSystemInfo) SetCanSelfRestart(v bool)`

SetCanSelfRestart sets CanSelfRestart field to given value.

### HasCanSelfRestart

`func (o *ModelSystemInfo) HasCanSelfRestart() bool`

HasCanSelfRestart returns a boolean if a field has been set.

### GetCanSelfUpdate

`func (o *ModelSystemInfo) GetCanSelfUpdate() bool`

GetCanSelfUpdate returns the CanSelfUpdate field if non-nil, zero value otherwise.

### GetCanSelfUpdateOk

`func (o *ModelSystemInfo) GetCanSelfUpdateOk() (*bool, bool)`

GetCanSelfUpdateOk returns a tuple with the CanSelfUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSelfUpdate

`func (o *ModelSystemInfo) SetCanSelfUpdate(v bool)`

SetCanSelfUpdate sets CanSelfUpdate field to given value.

### HasCanSelfUpdate

`func (o *ModelSystemInfo) HasCanSelfUpdate() bool`

HasCanSelfUpdate returns a boolean if a field has been set.

### GetCanLaunchWebBrowser

`func (o *ModelSystemInfo) GetCanLaunchWebBrowser() bool`

GetCanLaunchWebBrowser returns the CanLaunchWebBrowser field if non-nil, zero value otherwise.

### GetCanLaunchWebBrowserOk

`func (o *ModelSystemInfo) GetCanLaunchWebBrowserOk() (*bool, bool)`

GetCanLaunchWebBrowserOk returns a tuple with the CanLaunchWebBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLaunchWebBrowser

`func (o *ModelSystemInfo) SetCanLaunchWebBrowser(v bool)`

SetCanLaunchWebBrowser sets CanLaunchWebBrowser field to given value.

### HasCanLaunchWebBrowser

`func (o *ModelSystemInfo) HasCanLaunchWebBrowser() bool`

HasCanLaunchWebBrowser returns a boolean if a field has been set.

### GetProgramDataPath

`func (o *ModelSystemInfo) GetProgramDataPath() string`

GetProgramDataPath returns the ProgramDataPath field if non-nil, zero value otherwise.

### GetProgramDataPathOk

`func (o *ModelSystemInfo) GetProgramDataPathOk() (*string, bool)`

GetProgramDataPathOk returns a tuple with the ProgramDataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDataPath

`func (o *ModelSystemInfo) SetProgramDataPath(v string)`

SetProgramDataPath sets ProgramDataPath field to given value.

### HasProgramDataPath

`func (o *ModelSystemInfo) HasProgramDataPath() bool`

HasProgramDataPath returns a boolean if a field has been set.

### GetItemsByNamePath

`func (o *ModelSystemInfo) GetItemsByNamePath() string`

GetItemsByNamePath returns the ItemsByNamePath field if non-nil, zero value otherwise.

### GetItemsByNamePathOk

`func (o *ModelSystemInfo) GetItemsByNamePathOk() (*string, bool)`

GetItemsByNamePathOk returns a tuple with the ItemsByNamePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsByNamePath

`func (o *ModelSystemInfo) SetItemsByNamePath(v string)`

SetItemsByNamePath sets ItemsByNamePath field to given value.

### HasItemsByNamePath

`func (o *ModelSystemInfo) HasItemsByNamePath() bool`

HasItemsByNamePath returns a boolean if a field has been set.

### GetCachePath

`func (o *ModelSystemInfo) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *ModelSystemInfo) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *ModelSystemInfo) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *ModelSystemInfo) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.

### GetLogPath

`func (o *ModelSystemInfo) GetLogPath() string`

GetLogPath returns the LogPath field if non-nil, zero value otherwise.

### GetLogPathOk

`func (o *ModelSystemInfo) GetLogPathOk() (*string, bool)`

GetLogPathOk returns a tuple with the LogPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogPath

`func (o *ModelSystemInfo) SetLogPath(v string)`

SetLogPath sets LogPath field to given value.

### HasLogPath

`func (o *ModelSystemInfo) HasLogPath() bool`

HasLogPath returns a boolean if a field has been set.

### GetInternalMetadataPath

`func (o *ModelSystemInfo) GetInternalMetadataPath() string`

GetInternalMetadataPath returns the InternalMetadataPath field if non-nil, zero value otherwise.

### GetInternalMetadataPathOk

`func (o *ModelSystemInfo) GetInternalMetadataPathOk() (*string, bool)`

GetInternalMetadataPathOk returns a tuple with the InternalMetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadataPath

`func (o *ModelSystemInfo) SetInternalMetadataPath(v string)`

SetInternalMetadataPath sets InternalMetadataPath field to given value.

### HasInternalMetadataPath

`func (o *ModelSystemInfo) HasInternalMetadataPath() bool`

HasInternalMetadataPath returns a boolean if a field has been set.

### GetTranscodingTempPath

`func (o *ModelSystemInfo) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *ModelSystemInfo) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *ModelSystemInfo) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *ModelSystemInfo) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### GetHttpServerPortNumber

`func (o *ModelSystemInfo) GetHttpServerPortNumber() int32`

GetHttpServerPortNumber returns the HttpServerPortNumber field if non-nil, zero value otherwise.

### GetHttpServerPortNumberOk

`func (o *ModelSystemInfo) GetHttpServerPortNumberOk() (*int32, bool)`

GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerPortNumber

`func (o *ModelSystemInfo) SetHttpServerPortNumber(v int32)`

SetHttpServerPortNumber sets HttpServerPortNumber field to given value.

### HasHttpServerPortNumber

`func (o *ModelSystemInfo) HasHttpServerPortNumber() bool`

HasHttpServerPortNumber returns a boolean if a field has been set.

### GetSupportsHttps

`func (o *ModelSystemInfo) GetSupportsHttps() bool`

GetSupportsHttps returns the SupportsHttps field if non-nil, zero value otherwise.

### GetSupportsHttpsOk

`func (o *ModelSystemInfo) GetSupportsHttpsOk() (*bool, bool)`

GetSupportsHttpsOk returns a tuple with the SupportsHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsHttps

`func (o *ModelSystemInfo) SetSupportsHttps(v bool)`

SetSupportsHttps sets SupportsHttps field to given value.

### HasSupportsHttps

`func (o *ModelSystemInfo) HasSupportsHttps() bool`

HasSupportsHttps returns a boolean if a field has been set.

### GetHttpsPortNumber

`func (o *ModelSystemInfo) GetHttpsPortNumber() int32`

GetHttpsPortNumber returns the HttpsPortNumber field if non-nil, zero value otherwise.

### GetHttpsPortNumberOk

`func (o *ModelSystemInfo) GetHttpsPortNumberOk() (*int32, bool)`

GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPortNumber

`func (o *ModelSystemInfo) SetHttpsPortNumber(v int32)`

SetHttpsPortNumber sets HttpsPortNumber field to given value.

### HasHttpsPortNumber

`func (o *ModelSystemInfo) HasHttpsPortNumber() bool`

HasHttpsPortNumber returns a boolean if a field has been set.

### GetHasUpdateAvailable

`func (o *ModelSystemInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *ModelSystemInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *ModelSystemInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *ModelSystemInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetSupportsAutoRunAtStartup

`func (o *ModelSystemInfo) GetSupportsAutoRunAtStartup() bool`

GetSupportsAutoRunAtStartup returns the SupportsAutoRunAtStartup field if non-nil, zero value otherwise.

### GetSupportsAutoRunAtStartupOk

`func (o *ModelSystemInfo) GetSupportsAutoRunAtStartupOk() (*bool, bool)`

GetSupportsAutoRunAtStartupOk returns a tuple with the SupportsAutoRunAtStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAutoRunAtStartup

`func (o *ModelSystemInfo) SetSupportsAutoRunAtStartup(v bool)`

SetSupportsAutoRunAtStartup sets SupportsAutoRunAtStartup field to given value.

### HasSupportsAutoRunAtStartup

`func (o *ModelSystemInfo) HasSupportsAutoRunAtStartup() bool`

HasSupportsAutoRunAtStartup returns a boolean if a field has been set.

### GetHardwareAccelerationRequiresPremiere

`func (o *ModelSystemInfo) GetHardwareAccelerationRequiresPremiere() bool`

GetHardwareAccelerationRequiresPremiere returns the HardwareAccelerationRequiresPremiere field if non-nil, zero value otherwise.

### GetHardwareAccelerationRequiresPremiereOk

`func (o *ModelSystemInfo) GetHardwareAccelerationRequiresPremiereOk() (*bool, bool)`

GetHardwareAccelerationRequiresPremiereOk returns a tuple with the HardwareAccelerationRequiresPremiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationRequiresPremiere

`func (o *ModelSystemInfo) SetHardwareAccelerationRequiresPremiere(v bool)`

SetHardwareAccelerationRequiresPremiere sets HardwareAccelerationRequiresPremiere field to given value.

### HasHardwareAccelerationRequiresPremiere

`func (o *ModelSystemInfo) HasHardwareAccelerationRequiresPremiere() bool`

HasHardwareAccelerationRequiresPremiere returns a boolean if a field has been set.

### GetLocalAddress

`func (o *ModelSystemInfo) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *ModelSystemInfo) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *ModelSystemInfo) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *ModelSystemInfo) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### GetLocalAddresses

`func (o *ModelSystemInfo) GetLocalAddresses() []string`

GetLocalAddresses returns the LocalAddresses field if non-nil, zero value otherwise.

### GetLocalAddressesOk

`func (o *ModelSystemInfo) GetLocalAddressesOk() (*[]string, bool)`

GetLocalAddressesOk returns a tuple with the LocalAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddresses

`func (o *ModelSystemInfo) SetLocalAddresses(v []string)`

SetLocalAddresses sets LocalAddresses field to given value.

### HasLocalAddresses

`func (o *ModelSystemInfo) HasLocalAddresses() bool`

HasLocalAddresses returns a boolean if a field has been set.

### GetWanAddress

`func (o *ModelSystemInfo) GetWanAddress() string`

GetWanAddress returns the WanAddress field if non-nil, zero value otherwise.

### GetWanAddressOk

`func (o *ModelSystemInfo) GetWanAddressOk() (*string, bool)`

GetWanAddressOk returns a tuple with the WanAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAddress

`func (o *ModelSystemInfo) SetWanAddress(v string)`

SetWanAddress sets WanAddress field to given value.

### HasWanAddress

`func (o *ModelSystemInfo) HasWanAddress() bool`

HasWanAddress returns a boolean if a field has been set.

### GetRemoteAddresses

`func (o *ModelSystemInfo) GetRemoteAddresses() []string`

GetRemoteAddresses returns the RemoteAddresses field if non-nil, zero value otherwise.

### GetRemoteAddressesOk

`func (o *ModelSystemInfo) GetRemoteAddressesOk() (*[]string, bool)`

GetRemoteAddressesOk returns a tuple with the RemoteAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddresses

`func (o *ModelSystemInfo) SetRemoteAddresses(v []string)`

SetRemoteAddresses sets RemoteAddresses field to given value.

### HasRemoteAddresses

`func (o *ModelSystemInfo) HasRemoteAddresses() bool`

HasRemoteAddresses returns a boolean if a field has been set.

### GetServerName

`func (o *ModelSystemInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ModelSystemInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ModelSystemInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ModelSystemInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetVersion

`func (o *ModelSystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelSystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelSystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelSystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetId

`func (o *ModelSystemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSystemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSystemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSystemInfo) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


