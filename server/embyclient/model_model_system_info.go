/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
)

// checks if the ModelSystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSystemInfo{}

// ModelSystemInfo struct for ModelSystemInfo
type ModelSystemInfo struct {
	SystemUpdateLevel *ModelPackageVersionClass `json:"SystemUpdateLevel,omitempty"`
	OperatingSystemDisplayName *string `json:"OperatingSystemDisplayName,omitempty"`
	PackageName *string `json:"PackageName,omitempty"`
	HasPendingRestart *bool `json:"HasPendingRestart,omitempty"`
	IsShuttingDown *bool `json:"IsShuttingDown,omitempty"`
	OperatingSystem *string `json:"OperatingSystem,omitempty"`
	SupportsLibraryMonitor *bool `json:"SupportsLibraryMonitor,omitempty"`
	SupportsLocalPortConfiguration *bool `json:"SupportsLocalPortConfiguration,omitempty"`
	SupportsWakeServer *bool `json:"SupportsWakeServer,omitempty"`
	WebSocketPortNumber *int32 `json:"WebSocketPortNumber,omitempty"`
	CompletedInstallations []ModelInstallationInfo `json:"CompletedInstallations,omitempty"`
	CanSelfRestart *bool `json:"CanSelfRestart,omitempty"`
	CanSelfUpdate *bool `json:"CanSelfUpdate,omitempty"`
	CanLaunchWebBrowser *bool `json:"CanLaunchWebBrowser,omitempty"`
	ProgramDataPath *string `json:"ProgramDataPath,omitempty"`
	ItemsByNamePath *string `json:"ItemsByNamePath,omitempty"`
	CachePath *string `json:"CachePath,omitempty"`
	LogPath *string `json:"LogPath,omitempty"`
	InternalMetadataPath *string `json:"InternalMetadataPath,omitempty"`
	TranscodingTempPath *string `json:"TranscodingTempPath,omitempty"`
	HttpServerPortNumber *int32 `json:"HttpServerPortNumber,omitempty"`
	SupportsHttps *bool `json:"SupportsHttps,omitempty"`
	HttpsPortNumber *int32 `json:"HttpsPortNumber,omitempty"`
	HasUpdateAvailable *bool `json:"HasUpdateAvailable,omitempty"`
	SupportsAutoRunAtStartup *bool `json:"SupportsAutoRunAtStartup,omitempty"`
	HardwareAccelerationRequiresPremiere *bool `json:"HardwareAccelerationRequiresPremiere,omitempty"`
	LocalAddress *string `json:"LocalAddress,omitempty"`
	LocalAddresses []string `json:"LocalAddresses,omitempty"`
	WanAddress *string `json:"WanAddress,omitempty"`
	RemoteAddresses []string `json:"RemoteAddresses,omitempty"`
	ServerName *string `json:"ServerName,omitempty"`
	Version *string `json:"Version,omitempty"`
	Id *string `json:"Id,omitempty"`
}

// NewModelSystemInfo instantiates a new ModelSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSystemInfo() *ModelSystemInfo {
	this := ModelSystemInfo{}
	return &this
}

// NewModelSystemInfoWithDefaults instantiates a new ModelSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSystemInfoWithDefaults() *ModelSystemInfo {
	this := ModelSystemInfo{}
	return &this
}

// GetSystemUpdateLevel returns the SystemUpdateLevel field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetSystemUpdateLevel() ModelPackageVersionClass {
	if o == nil || IsNil(o.SystemUpdateLevel) {
		var ret ModelPackageVersionClass
		return ret
	}
	return *o.SystemUpdateLevel
}

// GetSystemUpdateLevelOk returns a tuple with the SystemUpdateLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetSystemUpdateLevelOk() (*ModelPackageVersionClass, bool) {
	if o == nil || IsNil(o.SystemUpdateLevel) {
		return nil, false
	}
	return o.SystemUpdateLevel, true
}

// HasSystemUpdateLevel returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasSystemUpdateLevel() bool {
	if o != nil && !IsNil(o.SystemUpdateLevel) {
		return true
	}

	return false
}

// SetSystemUpdateLevel gets a reference to the given ModelPackageVersionClass and assigns it to the SystemUpdateLevel field.
func (o *ModelSystemInfo) SetSystemUpdateLevel(v ModelPackageVersionClass) {
	o.SystemUpdateLevel = &v
}

// GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetOperatingSystemDisplayName() string {
	if o == nil || IsNil(o.OperatingSystemDisplayName) {
		var ret string
		return ret
	}
	return *o.OperatingSystemDisplayName
}

// GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetOperatingSystemDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystemDisplayName) {
		return nil, false
	}
	return o.OperatingSystemDisplayName, true
}

// HasOperatingSystemDisplayName returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasOperatingSystemDisplayName() bool {
	if o != nil && !IsNil(o.OperatingSystemDisplayName) {
		return true
	}

	return false
}

// SetOperatingSystemDisplayName gets a reference to the given string and assigns it to the OperatingSystemDisplayName field.
func (o *ModelSystemInfo) SetOperatingSystemDisplayName(v string) {
	o.OperatingSystemDisplayName = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetPackageName() string {
	if o == nil || IsNil(o.PackageName) {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetPackageNameOk() (*string, bool) {
	if o == nil || IsNil(o.PackageName) {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasPackageName() bool {
	if o != nil && !IsNil(o.PackageName) {
		return true
	}

	return false
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *ModelSystemInfo) SetPackageName(v string) {
	o.PackageName = &v
}

// GetHasPendingRestart returns the HasPendingRestart field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetHasPendingRestart() bool {
	if o == nil || IsNil(o.HasPendingRestart) {
		var ret bool
		return ret
	}
	return *o.HasPendingRestart
}

// GetHasPendingRestartOk returns a tuple with the HasPendingRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetHasPendingRestartOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPendingRestart) {
		return nil, false
	}
	return o.HasPendingRestart, true
}

// HasHasPendingRestart returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasHasPendingRestart() bool {
	if o != nil && !IsNil(o.HasPendingRestart) {
		return true
	}

	return false
}

// SetHasPendingRestart gets a reference to the given bool and assigns it to the HasPendingRestart field.
func (o *ModelSystemInfo) SetHasPendingRestart(v bool) {
	o.HasPendingRestart = &v
}

// GetIsShuttingDown returns the IsShuttingDown field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetIsShuttingDown() bool {
	if o == nil || IsNil(o.IsShuttingDown) {
		var ret bool
		return ret
	}
	return *o.IsShuttingDown
}

// GetIsShuttingDownOk returns a tuple with the IsShuttingDown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetIsShuttingDownOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShuttingDown) {
		return nil, false
	}
	return o.IsShuttingDown, true
}

// HasIsShuttingDown returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasIsShuttingDown() bool {
	if o != nil && !IsNil(o.IsShuttingDown) {
		return true
	}

	return false
}

// SetIsShuttingDown gets a reference to the given bool and assigns it to the IsShuttingDown field.
func (o *ModelSystemInfo) SetIsShuttingDown(v bool) {
	o.IsShuttingDown = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetOperatingSystem() string {
	if o == nil || IsNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *ModelSystemInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetSupportsLibraryMonitor returns the SupportsLibraryMonitor field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetSupportsLibraryMonitor() bool {
	if o == nil || IsNil(o.SupportsLibraryMonitor) {
		var ret bool
		return ret
	}
	return *o.SupportsLibraryMonitor
}

// GetSupportsLibraryMonitorOk returns a tuple with the SupportsLibraryMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetSupportsLibraryMonitorOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsLibraryMonitor) {
		return nil, false
	}
	return o.SupportsLibraryMonitor, true
}

// HasSupportsLibraryMonitor returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasSupportsLibraryMonitor() bool {
	if o != nil && !IsNil(o.SupportsLibraryMonitor) {
		return true
	}

	return false
}

// SetSupportsLibraryMonitor gets a reference to the given bool and assigns it to the SupportsLibraryMonitor field.
func (o *ModelSystemInfo) SetSupportsLibraryMonitor(v bool) {
	o.SupportsLibraryMonitor = &v
}

// GetSupportsLocalPortConfiguration returns the SupportsLocalPortConfiguration field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetSupportsLocalPortConfiguration() bool {
	if o == nil || IsNil(o.SupportsLocalPortConfiguration) {
		var ret bool
		return ret
	}
	return *o.SupportsLocalPortConfiguration
}

// GetSupportsLocalPortConfigurationOk returns a tuple with the SupportsLocalPortConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetSupportsLocalPortConfigurationOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsLocalPortConfiguration) {
		return nil, false
	}
	return o.SupportsLocalPortConfiguration, true
}

// HasSupportsLocalPortConfiguration returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasSupportsLocalPortConfiguration() bool {
	if o != nil && !IsNil(o.SupportsLocalPortConfiguration) {
		return true
	}

	return false
}

// SetSupportsLocalPortConfiguration gets a reference to the given bool and assigns it to the SupportsLocalPortConfiguration field.
func (o *ModelSystemInfo) SetSupportsLocalPortConfiguration(v bool) {
	o.SupportsLocalPortConfiguration = &v
}

// GetSupportsWakeServer returns the SupportsWakeServer field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetSupportsWakeServer() bool {
	if o == nil || IsNil(o.SupportsWakeServer) {
		var ret bool
		return ret
	}
	return *o.SupportsWakeServer
}

// GetSupportsWakeServerOk returns a tuple with the SupportsWakeServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetSupportsWakeServerOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsWakeServer) {
		return nil, false
	}
	return o.SupportsWakeServer, true
}

// HasSupportsWakeServer returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasSupportsWakeServer() bool {
	if o != nil && !IsNil(o.SupportsWakeServer) {
		return true
	}

	return false
}

// SetSupportsWakeServer gets a reference to the given bool and assigns it to the SupportsWakeServer field.
func (o *ModelSystemInfo) SetSupportsWakeServer(v bool) {
	o.SupportsWakeServer = &v
}

// GetWebSocketPortNumber returns the WebSocketPortNumber field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetWebSocketPortNumber() int32 {
	if o == nil || IsNil(o.WebSocketPortNumber) {
		var ret int32
		return ret
	}
	return *o.WebSocketPortNumber
}

// GetWebSocketPortNumberOk returns a tuple with the WebSocketPortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetWebSocketPortNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.WebSocketPortNumber) {
		return nil, false
	}
	return o.WebSocketPortNumber, true
}

// HasWebSocketPortNumber returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasWebSocketPortNumber() bool {
	if o != nil && !IsNil(o.WebSocketPortNumber) {
		return true
	}

	return false
}

// SetWebSocketPortNumber gets a reference to the given int32 and assigns it to the WebSocketPortNumber field.
func (o *ModelSystemInfo) SetWebSocketPortNumber(v int32) {
	o.WebSocketPortNumber = &v
}

// GetCompletedInstallations returns the CompletedInstallations field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetCompletedInstallations() []ModelInstallationInfo {
	if o == nil || IsNil(o.CompletedInstallations) {
		var ret []ModelInstallationInfo
		return ret
	}
	return o.CompletedInstallations
}

// GetCompletedInstallationsOk returns a tuple with the CompletedInstallations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetCompletedInstallationsOk() ([]ModelInstallationInfo, bool) {
	if o == nil || IsNil(o.CompletedInstallations) {
		return nil, false
	}
	return o.CompletedInstallations, true
}

// HasCompletedInstallations returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasCompletedInstallations() bool {
	if o != nil && !IsNil(o.CompletedInstallations) {
		return true
	}

	return false
}

// SetCompletedInstallations gets a reference to the given []ModelInstallationInfo and assigns it to the CompletedInstallations field.
func (o *ModelSystemInfo) SetCompletedInstallations(v []ModelInstallationInfo) {
	o.CompletedInstallations = v
}

// GetCanSelfRestart returns the CanSelfRestart field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetCanSelfRestart() bool {
	if o == nil || IsNil(o.CanSelfRestart) {
		var ret bool
		return ret
	}
	return *o.CanSelfRestart
}

// GetCanSelfRestartOk returns a tuple with the CanSelfRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetCanSelfRestartOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSelfRestart) {
		return nil, false
	}
	return o.CanSelfRestart, true
}

// HasCanSelfRestart returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasCanSelfRestart() bool {
	if o != nil && !IsNil(o.CanSelfRestart) {
		return true
	}

	return false
}

// SetCanSelfRestart gets a reference to the given bool and assigns it to the CanSelfRestart field.
func (o *ModelSystemInfo) SetCanSelfRestart(v bool) {
	o.CanSelfRestart = &v
}

// GetCanSelfUpdate returns the CanSelfUpdate field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetCanSelfUpdate() bool {
	if o == nil || IsNil(o.CanSelfUpdate) {
		var ret bool
		return ret
	}
	return *o.CanSelfUpdate
}

// GetCanSelfUpdateOk returns a tuple with the CanSelfUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetCanSelfUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSelfUpdate) {
		return nil, false
	}
	return o.CanSelfUpdate, true
}

// HasCanSelfUpdate returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasCanSelfUpdate() bool {
	if o != nil && !IsNil(o.CanSelfUpdate) {
		return true
	}

	return false
}

// SetCanSelfUpdate gets a reference to the given bool and assigns it to the CanSelfUpdate field.
func (o *ModelSystemInfo) SetCanSelfUpdate(v bool) {
	o.CanSelfUpdate = &v
}

// GetCanLaunchWebBrowser returns the CanLaunchWebBrowser field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetCanLaunchWebBrowser() bool {
	if o == nil || IsNil(o.CanLaunchWebBrowser) {
		var ret bool
		return ret
	}
	return *o.CanLaunchWebBrowser
}

// GetCanLaunchWebBrowserOk returns a tuple with the CanLaunchWebBrowser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetCanLaunchWebBrowserOk() (*bool, bool) {
	if o == nil || IsNil(o.CanLaunchWebBrowser) {
		return nil, false
	}
	return o.CanLaunchWebBrowser, true
}

// HasCanLaunchWebBrowser returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasCanLaunchWebBrowser() bool {
	if o != nil && !IsNil(o.CanLaunchWebBrowser) {
		return true
	}

	return false
}

// SetCanLaunchWebBrowser gets a reference to the given bool and assigns it to the CanLaunchWebBrowser field.
func (o *ModelSystemInfo) SetCanLaunchWebBrowser(v bool) {
	o.CanLaunchWebBrowser = &v
}

// GetProgramDataPath returns the ProgramDataPath field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetProgramDataPath() string {
	if o == nil || IsNil(o.ProgramDataPath) {
		var ret string
		return ret
	}
	return *o.ProgramDataPath
}

// GetProgramDataPathOk returns a tuple with the ProgramDataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetProgramDataPathOk() (*string, bool) {
	if o == nil || IsNil(o.ProgramDataPath) {
		return nil, false
	}
	return o.ProgramDataPath, true
}

// HasProgramDataPath returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasProgramDataPath() bool {
	if o != nil && !IsNil(o.ProgramDataPath) {
		return true
	}

	return false
}

// SetProgramDataPath gets a reference to the given string and assigns it to the ProgramDataPath field.
func (o *ModelSystemInfo) SetProgramDataPath(v string) {
	o.ProgramDataPath = &v
}

// GetItemsByNamePath returns the ItemsByNamePath field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetItemsByNamePath() string {
	if o == nil || IsNil(o.ItemsByNamePath) {
		var ret string
		return ret
	}
	return *o.ItemsByNamePath
}

// GetItemsByNamePathOk returns a tuple with the ItemsByNamePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetItemsByNamePathOk() (*string, bool) {
	if o == nil || IsNil(o.ItemsByNamePath) {
		return nil, false
	}
	return o.ItemsByNamePath, true
}

// HasItemsByNamePath returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasItemsByNamePath() bool {
	if o != nil && !IsNil(o.ItemsByNamePath) {
		return true
	}

	return false
}

// SetItemsByNamePath gets a reference to the given string and assigns it to the ItemsByNamePath field.
func (o *ModelSystemInfo) SetItemsByNamePath(v string) {
	o.ItemsByNamePath = &v
}

// GetCachePath returns the CachePath field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetCachePath() string {
	if o == nil || IsNil(o.CachePath) {
		var ret string
		return ret
	}
	return *o.CachePath
}

// GetCachePathOk returns a tuple with the CachePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetCachePathOk() (*string, bool) {
	if o == nil || IsNil(o.CachePath) {
		return nil, false
	}
	return o.CachePath, true
}

// HasCachePath returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasCachePath() bool {
	if o != nil && !IsNil(o.CachePath) {
		return true
	}

	return false
}

// SetCachePath gets a reference to the given string and assigns it to the CachePath field.
func (o *ModelSystemInfo) SetCachePath(v string) {
	o.CachePath = &v
}

// GetLogPath returns the LogPath field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetLogPath() string {
	if o == nil || IsNil(o.LogPath) {
		var ret string
		return ret
	}
	return *o.LogPath
}

// GetLogPathOk returns a tuple with the LogPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetLogPathOk() (*string, bool) {
	if o == nil || IsNil(o.LogPath) {
		return nil, false
	}
	return o.LogPath, true
}

// HasLogPath returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasLogPath() bool {
	if o != nil && !IsNil(o.LogPath) {
		return true
	}

	return false
}

// SetLogPath gets a reference to the given string and assigns it to the LogPath field.
func (o *ModelSystemInfo) SetLogPath(v string) {
	o.LogPath = &v
}

// GetInternalMetadataPath returns the InternalMetadataPath field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetInternalMetadataPath() string {
	if o == nil || IsNil(o.InternalMetadataPath) {
		var ret string
		return ret
	}
	return *o.InternalMetadataPath
}

// GetInternalMetadataPathOk returns a tuple with the InternalMetadataPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetInternalMetadataPathOk() (*string, bool) {
	if o == nil || IsNil(o.InternalMetadataPath) {
		return nil, false
	}
	return o.InternalMetadataPath, true
}

// HasInternalMetadataPath returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasInternalMetadataPath() bool {
	if o != nil && !IsNil(o.InternalMetadataPath) {
		return true
	}

	return false
}

// SetInternalMetadataPath gets a reference to the given string and assigns it to the InternalMetadataPath field.
func (o *ModelSystemInfo) SetInternalMetadataPath(v string) {
	o.InternalMetadataPath = &v
}

// GetTranscodingTempPath returns the TranscodingTempPath field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetTranscodingTempPath() string {
	if o == nil || IsNil(o.TranscodingTempPath) {
		var ret string
		return ret
	}
	return *o.TranscodingTempPath
}

// GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetTranscodingTempPathOk() (*string, bool) {
	if o == nil || IsNil(o.TranscodingTempPath) {
		return nil, false
	}
	return o.TranscodingTempPath, true
}

// HasTranscodingTempPath returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasTranscodingTempPath() bool {
	if o != nil && !IsNil(o.TranscodingTempPath) {
		return true
	}

	return false
}

// SetTranscodingTempPath gets a reference to the given string and assigns it to the TranscodingTempPath field.
func (o *ModelSystemInfo) SetTranscodingTempPath(v string) {
	o.TranscodingTempPath = &v
}

// GetHttpServerPortNumber returns the HttpServerPortNumber field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetHttpServerPortNumber() int32 {
	if o == nil || IsNil(o.HttpServerPortNumber) {
		var ret int32
		return ret
	}
	return *o.HttpServerPortNumber
}

// GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetHttpServerPortNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpServerPortNumber) {
		return nil, false
	}
	return o.HttpServerPortNumber, true
}

// HasHttpServerPortNumber returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasHttpServerPortNumber() bool {
	if o != nil && !IsNil(o.HttpServerPortNumber) {
		return true
	}

	return false
}

// SetHttpServerPortNumber gets a reference to the given int32 and assigns it to the HttpServerPortNumber field.
func (o *ModelSystemInfo) SetHttpServerPortNumber(v int32) {
	o.HttpServerPortNumber = &v
}

// GetSupportsHttps returns the SupportsHttps field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetSupportsHttps() bool {
	if o == nil || IsNil(o.SupportsHttps) {
		var ret bool
		return ret
	}
	return *o.SupportsHttps
}

// GetSupportsHttpsOk returns a tuple with the SupportsHttps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetSupportsHttpsOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsHttps) {
		return nil, false
	}
	return o.SupportsHttps, true
}

// HasSupportsHttps returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasSupportsHttps() bool {
	if o != nil && !IsNil(o.SupportsHttps) {
		return true
	}

	return false
}

// SetSupportsHttps gets a reference to the given bool and assigns it to the SupportsHttps field.
func (o *ModelSystemInfo) SetSupportsHttps(v bool) {
	o.SupportsHttps = &v
}

// GetHttpsPortNumber returns the HttpsPortNumber field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetHttpsPortNumber() int32 {
	if o == nil || IsNil(o.HttpsPortNumber) {
		var ret int32
		return ret
	}
	return *o.HttpsPortNumber
}

// GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetHttpsPortNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.HttpsPortNumber) {
		return nil, false
	}
	return o.HttpsPortNumber, true
}

// HasHttpsPortNumber returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasHttpsPortNumber() bool {
	if o != nil && !IsNil(o.HttpsPortNumber) {
		return true
	}

	return false
}

// SetHttpsPortNumber gets a reference to the given int32 and assigns it to the HttpsPortNumber field.
func (o *ModelSystemInfo) SetHttpsPortNumber(v int32) {
	o.HttpsPortNumber = &v
}

// GetHasUpdateAvailable returns the HasUpdateAvailable field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetHasUpdateAvailable() bool {
	if o == nil || IsNil(o.HasUpdateAvailable) {
		var ret bool
		return ret
	}
	return *o.HasUpdateAvailable
}

// GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetHasUpdateAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUpdateAvailable) {
		return nil, false
	}
	return o.HasUpdateAvailable, true
}

// HasHasUpdateAvailable returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasHasUpdateAvailable() bool {
	if o != nil && !IsNil(o.HasUpdateAvailable) {
		return true
	}

	return false
}

// SetHasUpdateAvailable gets a reference to the given bool and assigns it to the HasUpdateAvailable field.
func (o *ModelSystemInfo) SetHasUpdateAvailable(v bool) {
	o.HasUpdateAvailable = &v
}

// GetSupportsAutoRunAtStartup returns the SupportsAutoRunAtStartup field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetSupportsAutoRunAtStartup() bool {
	if o == nil || IsNil(o.SupportsAutoRunAtStartup) {
		var ret bool
		return ret
	}
	return *o.SupportsAutoRunAtStartup
}

// GetSupportsAutoRunAtStartupOk returns a tuple with the SupportsAutoRunAtStartup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetSupportsAutoRunAtStartupOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsAutoRunAtStartup) {
		return nil, false
	}
	return o.SupportsAutoRunAtStartup, true
}

// HasSupportsAutoRunAtStartup returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasSupportsAutoRunAtStartup() bool {
	if o != nil && !IsNil(o.SupportsAutoRunAtStartup) {
		return true
	}

	return false
}

// SetSupportsAutoRunAtStartup gets a reference to the given bool and assigns it to the SupportsAutoRunAtStartup field.
func (o *ModelSystemInfo) SetSupportsAutoRunAtStartup(v bool) {
	o.SupportsAutoRunAtStartup = &v
}

// GetHardwareAccelerationRequiresPremiere returns the HardwareAccelerationRequiresPremiere field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetHardwareAccelerationRequiresPremiere() bool {
	if o == nil || IsNil(o.HardwareAccelerationRequiresPremiere) {
		var ret bool
		return ret
	}
	return *o.HardwareAccelerationRequiresPremiere
}

// GetHardwareAccelerationRequiresPremiereOk returns a tuple with the HardwareAccelerationRequiresPremiere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetHardwareAccelerationRequiresPremiereOk() (*bool, bool) {
	if o == nil || IsNil(o.HardwareAccelerationRequiresPremiere) {
		return nil, false
	}
	return o.HardwareAccelerationRequiresPremiere, true
}

// HasHardwareAccelerationRequiresPremiere returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasHardwareAccelerationRequiresPremiere() bool {
	if o != nil && !IsNil(o.HardwareAccelerationRequiresPremiere) {
		return true
	}

	return false
}

// SetHardwareAccelerationRequiresPremiere gets a reference to the given bool and assigns it to the HardwareAccelerationRequiresPremiere field.
func (o *ModelSystemInfo) SetHardwareAccelerationRequiresPremiere(v bool) {
	o.HardwareAccelerationRequiresPremiere = &v
}

// GetLocalAddress returns the LocalAddress field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetLocalAddress() string {
	if o == nil || IsNil(o.LocalAddress) {
		var ret string
		return ret
	}
	return *o.LocalAddress
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetLocalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LocalAddress) {
		return nil, false
	}
	return o.LocalAddress, true
}

// HasLocalAddress returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasLocalAddress() bool {
	if o != nil && !IsNil(o.LocalAddress) {
		return true
	}

	return false
}

// SetLocalAddress gets a reference to the given string and assigns it to the LocalAddress field.
func (o *ModelSystemInfo) SetLocalAddress(v string) {
	o.LocalAddress = &v
}

// GetLocalAddresses returns the LocalAddresses field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetLocalAddresses() []string {
	if o == nil || IsNil(o.LocalAddresses) {
		var ret []string
		return ret
	}
	return o.LocalAddresses
}

// GetLocalAddressesOk returns a tuple with the LocalAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetLocalAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalAddresses) {
		return nil, false
	}
	return o.LocalAddresses, true
}

// HasLocalAddresses returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasLocalAddresses() bool {
	if o != nil && !IsNil(o.LocalAddresses) {
		return true
	}

	return false
}

// SetLocalAddresses gets a reference to the given []string and assigns it to the LocalAddresses field.
func (o *ModelSystemInfo) SetLocalAddresses(v []string) {
	o.LocalAddresses = v
}

// GetWanAddress returns the WanAddress field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetWanAddress() string {
	if o == nil || IsNil(o.WanAddress) {
		var ret string
		return ret
	}
	return *o.WanAddress
}

// GetWanAddressOk returns a tuple with the WanAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetWanAddressOk() (*string, bool) {
	if o == nil || IsNil(o.WanAddress) {
		return nil, false
	}
	return o.WanAddress, true
}

// HasWanAddress returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasWanAddress() bool {
	if o != nil && !IsNil(o.WanAddress) {
		return true
	}

	return false
}

// SetWanAddress gets a reference to the given string and assigns it to the WanAddress field.
func (o *ModelSystemInfo) SetWanAddress(v string) {
	o.WanAddress = &v
}

// GetRemoteAddresses returns the RemoteAddresses field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetRemoteAddresses() []string {
	if o == nil || IsNil(o.RemoteAddresses) {
		var ret []string
		return ret
	}
	return o.RemoteAddresses
}

// GetRemoteAddressesOk returns a tuple with the RemoteAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetRemoteAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoteAddresses) {
		return nil, false
	}
	return o.RemoteAddresses, true
}

// HasRemoteAddresses returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasRemoteAddresses() bool {
	if o != nil && !IsNil(o.RemoteAddresses) {
		return true
	}

	return false
}

// SetRemoteAddresses gets a reference to the given []string and assigns it to the RemoteAddresses field.
func (o *ModelSystemInfo) SetRemoteAddresses(v []string) {
	o.RemoteAddresses = v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *ModelSystemInfo) SetServerName(v string) {
	o.ServerName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ModelSystemInfo) SetVersion(v string) {
	o.Version = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelSystemInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSystemInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelSystemInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelSystemInfo) SetId(v string) {
	o.Id = &v
}

func (o ModelSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SystemUpdateLevel) {
		toSerialize["SystemUpdateLevel"] = o.SystemUpdateLevel
	}
	if !IsNil(o.OperatingSystemDisplayName) {
		toSerialize["OperatingSystemDisplayName"] = o.OperatingSystemDisplayName
	}
	if !IsNil(o.PackageName) {
		toSerialize["PackageName"] = o.PackageName
	}
	if !IsNil(o.HasPendingRestart) {
		toSerialize["HasPendingRestart"] = o.HasPendingRestart
	}
	if !IsNil(o.IsShuttingDown) {
		toSerialize["IsShuttingDown"] = o.IsShuttingDown
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["OperatingSystem"] = o.OperatingSystem
	}
	if !IsNil(o.SupportsLibraryMonitor) {
		toSerialize["SupportsLibraryMonitor"] = o.SupportsLibraryMonitor
	}
	if !IsNil(o.SupportsLocalPortConfiguration) {
		toSerialize["SupportsLocalPortConfiguration"] = o.SupportsLocalPortConfiguration
	}
	if !IsNil(o.SupportsWakeServer) {
		toSerialize["SupportsWakeServer"] = o.SupportsWakeServer
	}
	if !IsNil(o.WebSocketPortNumber) {
		toSerialize["WebSocketPortNumber"] = o.WebSocketPortNumber
	}
	if !IsNil(o.CompletedInstallations) {
		toSerialize["CompletedInstallations"] = o.CompletedInstallations
	}
	if !IsNil(o.CanSelfRestart) {
		toSerialize["CanSelfRestart"] = o.CanSelfRestart
	}
	if !IsNil(o.CanSelfUpdate) {
		toSerialize["CanSelfUpdate"] = o.CanSelfUpdate
	}
	if !IsNil(o.CanLaunchWebBrowser) {
		toSerialize["CanLaunchWebBrowser"] = o.CanLaunchWebBrowser
	}
	if !IsNil(o.ProgramDataPath) {
		toSerialize["ProgramDataPath"] = o.ProgramDataPath
	}
	if !IsNil(o.ItemsByNamePath) {
		toSerialize["ItemsByNamePath"] = o.ItemsByNamePath
	}
	if !IsNil(o.CachePath) {
		toSerialize["CachePath"] = o.CachePath
	}
	if !IsNil(o.LogPath) {
		toSerialize["LogPath"] = o.LogPath
	}
	if !IsNil(o.InternalMetadataPath) {
		toSerialize["InternalMetadataPath"] = o.InternalMetadataPath
	}
	if !IsNil(o.TranscodingTempPath) {
		toSerialize["TranscodingTempPath"] = o.TranscodingTempPath
	}
	if !IsNil(o.HttpServerPortNumber) {
		toSerialize["HttpServerPortNumber"] = o.HttpServerPortNumber
	}
	if !IsNil(o.SupportsHttps) {
		toSerialize["SupportsHttps"] = o.SupportsHttps
	}
	if !IsNil(o.HttpsPortNumber) {
		toSerialize["HttpsPortNumber"] = o.HttpsPortNumber
	}
	if !IsNil(o.HasUpdateAvailable) {
		toSerialize["HasUpdateAvailable"] = o.HasUpdateAvailable
	}
	if !IsNil(o.SupportsAutoRunAtStartup) {
		toSerialize["SupportsAutoRunAtStartup"] = o.SupportsAutoRunAtStartup
	}
	if !IsNil(o.HardwareAccelerationRequiresPremiere) {
		toSerialize["HardwareAccelerationRequiresPremiere"] = o.HardwareAccelerationRequiresPremiere
	}
	if !IsNil(o.LocalAddress) {
		toSerialize["LocalAddress"] = o.LocalAddress
	}
	if !IsNil(o.LocalAddresses) {
		toSerialize["LocalAddresses"] = o.LocalAddresses
	}
	if !IsNil(o.WanAddress) {
		toSerialize["WanAddress"] = o.WanAddress
	}
	if !IsNil(o.RemoteAddresses) {
		toSerialize["RemoteAddresses"] = o.RemoteAddresses
	}
	if !IsNil(o.ServerName) {
		toSerialize["ServerName"] = o.ServerName
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	return toSerialize, nil
}

type NullableModelSystemInfo struct {
	value *ModelSystemInfo
	isSet bool
}

func (v NullableModelSystemInfo) Get() *ModelSystemInfo {
	return v.value
}

func (v *NullableModelSystemInfo) Set(val *ModelSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSystemInfo(val *ModelSystemInfo) *NullableModelSystemInfo {
	return &NullableModelSystemInfo{value: val, isSet: true}
}

func (v NullableModelSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


