# ModelServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableUPnP** | Pointer to **bool** |  | [optional] 
**PublicPort** | Pointer to **int32** |  | [optional] 
**PublicHttpsPort** | Pointer to **int32** |  | [optional] 
**HttpServerPortNumber** | Pointer to **int32** |  | [optional] 
**HttpsPortNumber** | Pointer to **int32** |  | [optional] 
**EnableHttps** | Pointer to **bool** |  | [optional] 
**CertificatePath** | Pointer to **string** |  | [optional] 
**CertificatePassword** | Pointer to **string** |  | [optional] 
**IsPortAuthorized** | Pointer to **bool** |  | [optional] 
**AutoRunWebApp** | Pointer to **bool** |  | [optional] 
**EnableRemoteAccess** | Pointer to **bool** |  | [optional] 
**LogAllQueryTimes** | Pointer to **bool** |  | [optional] 
**EnableCaseSensitiveItemIds** | Pointer to **bool** |  | [optional] 
**MetadataPath** | Pointer to **string** |  | [optional] 
**MetadataNetworkPath** | Pointer to **string** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** |  | [optional] 
**MetadataCountryCode** | Pointer to **string** |  | [optional] 
**SortRemoveWords** | Pointer to **[]string** |  | [optional] 
**LibraryMonitorDelay** | Pointer to **int32** |  | [optional] 
**EnableDashboardResponseCaching** | Pointer to **bool** |  | [optional] 
**DashboardSourcePath** | Pointer to **string** |  | [optional] 
**ImageSavingConvention** | Pointer to [**ModelModelImageSavingConvention**](ModelImageSavingConvention.md) |  | [optional] 
**EnableAutomaticRestart** | Pointer to **bool** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**PreferredDetectedRemoteAddressFamily** | Pointer to [**ModelModelNetSocketsAddressFamily**](ModelNetSocketsAddressFamily.md) |  | [optional] 
**WanDdns** | Pointer to **string** |  | [optional] 
**UICulture** | Pointer to **string** |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**LocalNetworkSubnets** | Pointer to **[]string** |  | [optional] 
**LocalNetworkAddresses** | Pointer to **[]string** |  | [optional] 
**EnableExternalContentInSuggestions** | Pointer to **bool** |  | [optional] 
**RequireHttps** | Pointer to **bool** |  | [optional] 
**IsBehindProxy** | Pointer to **bool** |  | [optional] 
**RemoteIPFilter** | Pointer to **[]string** |  | [optional] 
**IsRemoteIPFilterBlacklist** | Pointer to **bool** |  | [optional] 
**ImageExtractionTimeoutMs** | Pointer to **int32** |  | [optional] 
**PathSubstitutions** | Pointer to [**[]ModelModelPathSubstitution**](ModelModelPathSubstitution.md) |  | [optional] 
**UninstalledPlugins** | Pointer to **[]string** |  | [optional] 
**CollapseVideoFolders** | Pointer to **bool** |  | [optional] 
**EnableOriginalTrackTitles** | Pointer to **bool** |  | [optional] 
**VacuumDatabaseOnStartup** | Pointer to **bool** |  | [optional] 
**SimultaneousStreamLimit** | Pointer to **int32** |  | [optional] 
**DatabaseCacheSizeMB** | Pointer to **int32** |  | [optional] 
**EnableSqLiteMmio** | Pointer to **bool** |  | [optional] 
**PlaylistsUpgradedToM3U** | Pointer to **bool** |  | [optional] 
**ImageExtractorUpgraded** | Pointer to **bool** |  | [optional] 
**EnablePeopleLetterSubFolders** | Pointer to **bool** |  | [optional] 
**OptimizeDatabaseOnShutdown** | Pointer to **bool** |  | [optional] 
**DatabaseAnalysisLimit** | Pointer to **int32** |  | [optional] 
**DisableAsyncIO** | Pointer to **bool** |  | [optional] 
**MigratedToUserItemShares** | Pointer to **bool** |  | [optional] 
**MigratedLibraryOptionsToDb** | Pointer to **bool** |  | [optional] 
**AllowLegacyLocalNetworkPassword** | Pointer to **bool** |  | [optional] 
**EnableSavedMetadataForPeople** | Pointer to **bool** |  | [optional] 
**ProxyHeaderMode** | Pointer to [**ModelModelProxyHeaderMode**](ModelProxyHeaderMode.md) |  | [optional] 
**EnableDebugLevelLogging** | Pointer to **bool** |  | [optional] 
**RevertDebugLogging** | Pointer to **string** |  | [optional] 
**EnableAutoUpdate** | Pointer to **bool** |  | [optional] 
**LogFileRetentionDays** | Pointer to **int32** |  | [optional] 
**RunAtStartup** | Pointer to **bool** |  | [optional] 
**IsStartupWizardCompleted** | Pointer to **bool** |  | [optional] 
**CachePath** | Pointer to **string** |  | [optional] 

## Methods

### NewModelServerConfiguration

`func NewModelServerConfiguration() *ModelServerConfiguration`

NewModelServerConfiguration instantiates a new ModelServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelServerConfigurationWithDefaults

`func NewModelServerConfigurationWithDefaults() *ModelServerConfiguration`

NewModelServerConfigurationWithDefaults instantiates a new ModelServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableUPnP

`func (o *ModelServerConfiguration) GetEnableUPnP() bool`

GetEnableUPnP returns the EnableUPnP field if non-nil, zero value otherwise.

### GetEnableUPnPOk

`func (o *ModelServerConfiguration) GetEnableUPnPOk() (*bool, bool)`

GetEnableUPnPOk returns a tuple with the EnableUPnP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUPnP

`func (o *ModelServerConfiguration) SetEnableUPnP(v bool)`

SetEnableUPnP sets EnableUPnP field to given value.

### HasEnableUPnP

`func (o *ModelServerConfiguration) HasEnableUPnP() bool`

HasEnableUPnP returns a boolean if a field has been set.

### GetPublicPort

`func (o *ModelServerConfiguration) GetPublicPort() int32`

GetPublicPort returns the PublicPort field if non-nil, zero value otherwise.

### GetPublicPortOk

`func (o *ModelServerConfiguration) GetPublicPortOk() (*int32, bool)`

GetPublicPortOk returns a tuple with the PublicPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicPort

`func (o *ModelServerConfiguration) SetPublicPort(v int32)`

SetPublicPort sets PublicPort field to given value.

### HasPublicPort

`func (o *ModelServerConfiguration) HasPublicPort() bool`

HasPublicPort returns a boolean if a field has been set.

### GetPublicHttpsPort

`func (o *ModelServerConfiguration) GetPublicHttpsPort() int32`

GetPublicHttpsPort returns the PublicHttpsPort field if non-nil, zero value otherwise.

### GetPublicHttpsPortOk

`func (o *ModelServerConfiguration) GetPublicHttpsPortOk() (*int32, bool)`

GetPublicHttpsPortOk returns a tuple with the PublicHttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHttpsPort

`func (o *ModelServerConfiguration) SetPublicHttpsPort(v int32)`

SetPublicHttpsPort sets PublicHttpsPort field to given value.

### HasPublicHttpsPort

`func (o *ModelServerConfiguration) HasPublicHttpsPort() bool`

HasPublicHttpsPort returns a boolean if a field has been set.

### GetHttpServerPortNumber

`func (o *ModelServerConfiguration) GetHttpServerPortNumber() int32`

GetHttpServerPortNumber returns the HttpServerPortNumber field if non-nil, zero value otherwise.

### GetHttpServerPortNumberOk

`func (o *ModelServerConfiguration) GetHttpServerPortNumberOk() (*int32, bool)`

GetHttpServerPortNumberOk returns a tuple with the HttpServerPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpServerPortNumber

`func (o *ModelServerConfiguration) SetHttpServerPortNumber(v int32)`

SetHttpServerPortNumber sets HttpServerPortNumber field to given value.

### HasHttpServerPortNumber

`func (o *ModelServerConfiguration) HasHttpServerPortNumber() bool`

HasHttpServerPortNumber returns a boolean if a field has been set.

### GetHttpsPortNumber

`func (o *ModelServerConfiguration) GetHttpsPortNumber() int32`

GetHttpsPortNumber returns the HttpsPortNumber field if non-nil, zero value otherwise.

### GetHttpsPortNumberOk

`func (o *ModelServerConfiguration) GetHttpsPortNumberOk() (*int32, bool)`

GetHttpsPortNumberOk returns a tuple with the HttpsPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPortNumber

`func (o *ModelServerConfiguration) SetHttpsPortNumber(v int32)`

SetHttpsPortNumber sets HttpsPortNumber field to given value.

### HasHttpsPortNumber

`func (o *ModelServerConfiguration) HasHttpsPortNumber() bool`

HasHttpsPortNumber returns a boolean if a field has been set.

### GetEnableHttps

`func (o *ModelServerConfiguration) GetEnableHttps() bool`

GetEnableHttps returns the EnableHttps field if non-nil, zero value otherwise.

### GetEnableHttpsOk

`func (o *ModelServerConfiguration) GetEnableHttpsOk() (*bool, bool)`

GetEnableHttpsOk returns a tuple with the EnableHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttps

`func (o *ModelServerConfiguration) SetEnableHttps(v bool)`

SetEnableHttps sets EnableHttps field to given value.

### HasEnableHttps

`func (o *ModelServerConfiguration) HasEnableHttps() bool`

HasEnableHttps returns a boolean if a field has been set.

### GetCertificatePath

`func (o *ModelServerConfiguration) GetCertificatePath() string`

GetCertificatePath returns the CertificatePath field if non-nil, zero value otherwise.

### GetCertificatePathOk

`func (o *ModelServerConfiguration) GetCertificatePathOk() (*string, bool)`

GetCertificatePathOk returns a tuple with the CertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePath

`func (o *ModelServerConfiguration) SetCertificatePath(v string)`

SetCertificatePath sets CertificatePath field to given value.

### HasCertificatePath

`func (o *ModelServerConfiguration) HasCertificatePath() bool`

HasCertificatePath returns a boolean if a field has been set.

### GetCertificatePassword

`func (o *ModelServerConfiguration) GetCertificatePassword() string`

GetCertificatePassword returns the CertificatePassword field if non-nil, zero value otherwise.

### GetCertificatePasswordOk

`func (o *ModelServerConfiguration) GetCertificatePasswordOk() (*string, bool)`

GetCertificatePasswordOk returns a tuple with the CertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatePassword

`func (o *ModelServerConfiguration) SetCertificatePassword(v string)`

SetCertificatePassword sets CertificatePassword field to given value.

### HasCertificatePassword

`func (o *ModelServerConfiguration) HasCertificatePassword() bool`

HasCertificatePassword returns a boolean if a field has been set.

### GetIsPortAuthorized

`func (o *ModelServerConfiguration) GetIsPortAuthorized() bool`

GetIsPortAuthorized returns the IsPortAuthorized field if non-nil, zero value otherwise.

### GetIsPortAuthorizedOk

`func (o *ModelServerConfiguration) GetIsPortAuthorizedOk() (*bool, bool)`

GetIsPortAuthorizedOk returns a tuple with the IsPortAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPortAuthorized

`func (o *ModelServerConfiguration) SetIsPortAuthorized(v bool)`

SetIsPortAuthorized sets IsPortAuthorized field to given value.

### HasIsPortAuthorized

`func (o *ModelServerConfiguration) HasIsPortAuthorized() bool`

HasIsPortAuthorized returns a boolean if a field has been set.

### GetAutoRunWebApp

`func (o *ModelServerConfiguration) GetAutoRunWebApp() bool`

GetAutoRunWebApp returns the AutoRunWebApp field if non-nil, zero value otherwise.

### GetAutoRunWebAppOk

`func (o *ModelServerConfiguration) GetAutoRunWebAppOk() (*bool, bool)`

GetAutoRunWebAppOk returns a tuple with the AutoRunWebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRunWebApp

`func (o *ModelServerConfiguration) SetAutoRunWebApp(v bool)`

SetAutoRunWebApp sets AutoRunWebApp field to given value.

### HasAutoRunWebApp

`func (o *ModelServerConfiguration) HasAutoRunWebApp() bool`

HasAutoRunWebApp returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *ModelServerConfiguration) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *ModelServerConfiguration) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *ModelServerConfiguration) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *ModelServerConfiguration) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetLogAllQueryTimes

`func (o *ModelServerConfiguration) GetLogAllQueryTimes() bool`

GetLogAllQueryTimes returns the LogAllQueryTimes field if non-nil, zero value otherwise.

### GetLogAllQueryTimesOk

`func (o *ModelServerConfiguration) GetLogAllQueryTimesOk() (*bool, bool)`

GetLogAllQueryTimesOk returns a tuple with the LogAllQueryTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAllQueryTimes

`func (o *ModelServerConfiguration) SetLogAllQueryTimes(v bool)`

SetLogAllQueryTimes sets LogAllQueryTimes field to given value.

### HasLogAllQueryTimes

`func (o *ModelServerConfiguration) HasLogAllQueryTimes() bool`

HasLogAllQueryTimes returns a boolean if a field has been set.

### GetEnableCaseSensitiveItemIds

`func (o *ModelServerConfiguration) GetEnableCaseSensitiveItemIds() bool`

GetEnableCaseSensitiveItemIds returns the EnableCaseSensitiveItemIds field if non-nil, zero value otherwise.

### GetEnableCaseSensitiveItemIdsOk

`func (o *ModelServerConfiguration) GetEnableCaseSensitiveItemIdsOk() (*bool, bool)`

GetEnableCaseSensitiveItemIdsOk returns a tuple with the EnableCaseSensitiveItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaseSensitiveItemIds

`func (o *ModelServerConfiguration) SetEnableCaseSensitiveItemIds(v bool)`

SetEnableCaseSensitiveItemIds sets EnableCaseSensitiveItemIds field to given value.

### HasEnableCaseSensitiveItemIds

`func (o *ModelServerConfiguration) HasEnableCaseSensitiveItemIds() bool`

HasEnableCaseSensitiveItemIds returns a boolean if a field has been set.

### GetMetadataPath

`func (o *ModelServerConfiguration) GetMetadataPath() string`

GetMetadataPath returns the MetadataPath field if non-nil, zero value otherwise.

### GetMetadataPathOk

`func (o *ModelServerConfiguration) GetMetadataPathOk() (*string, bool)`

GetMetadataPathOk returns a tuple with the MetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPath

`func (o *ModelServerConfiguration) SetMetadataPath(v string)`

SetMetadataPath sets MetadataPath field to given value.

### HasMetadataPath

`func (o *ModelServerConfiguration) HasMetadataPath() bool`

HasMetadataPath returns a boolean if a field has been set.

### GetMetadataNetworkPath

`func (o *ModelServerConfiguration) GetMetadataNetworkPath() string`

GetMetadataNetworkPath returns the MetadataNetworkPath field if non-nil, zero value otherwise.

### GetMetadataNetworkPathOk

`func (o *ModelServerConfiguration) GetMetadataNetworkPathOk() (*string, bool)`

GetMetadataNetworkPathOk returns a tuple with the MetadataNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataNetworkPath

`func (o *ModelServerConfiguration) SetMetadataNetworkPath(v string)`

SetMetadataNetworkPath sets MetadataNetworkPath field to given value.

### HasMetadataNetworkPath

`func (o *ModelServerConfiguration) HasMetadataNetworkPath() bool`

HasMetadataNetworkPath returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ModelServerConfiguration) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ModelServerConfiguration) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ModelServerConfiguration) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ModelServerConfiguration) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ModelServerConfiguration) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ModelServerConfiguration) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ModelServerConfiguration) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ModelServerConfiguration) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetSortRemoveWords

`func (o *ModelServerConfiguration) GetSortRemoveWords() []string`

GetSortRemoveWords returns the SortRemoveWords field if non-nil, zero value otherwise.

### GetSortRemoveWordsOk

`func (o *ModelServerConfiguration) GetSortRemoveWordsOk() (*[]string, bool)`

GetSortRemoveWordsOk returns a tuple with the SortRemoveWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveWords

`func (o *ModelServerConfiguration) SetSortRemoveWords(v []string)`

SetSortRemoveWords sets SortRemoveWords field to given value.

### HasSortRemoveWords

`func (o *ModelServerConfiguration) HasSortRemoveWords() bool`

HasSortRemoveWords returns a boolean if a field has been set.

### GetLibraryMonitorDelay

`func (o *ModelServerConfiguration) GetLibraryMonitorDelay() int32`

GetLibraryMonitorDelay returns the LibraryMonitorDelay field if non-nil, zero value otherwise.

### GetLibraryMonitorDelayOk

`func (o *ModelServerConfiguration) GetLibraryMonitorDelayOk() (*int32, bool)`

GetLibraryMonitorDelayOk returns a tuple with the LibraryMonitorDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMonitorDelay

`func (o *ModelServerConfiguration) SetLibraryMonitorDelay(v int32)`

SetLibraryMonitorDelay sets LibraryMonitorDelay field to given value.

### HasLibraryMonitorDelay

`func (o *ModelServerConfiguration) HasLibraryMonitorDelay() bool`

HasLibraryMonitorDelay returns a boolean if a field has been set.

### GetEnableDashboardResponseCaching

`func (o *ModelServerConfiguration) GetEnableDashboardResponseCaching() bool`

GetEnableDashboardResponseCaching returns the EnableDashboardResponseCaching field if non-nil, zero value otherwise.

### GetEnableDashboardResponseCachingOk

`func (o *ModelServerConfiguration) GetEnableDashboardResponseCachingOk() (*bool, bool)`

GetEnableDashboardResponseCachingOk returns a tuple with the EnableDashboardResponseCaching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDashboardResponseCaching

`func (o *ModelServerConfiguration) SetEnableDashboardResponseCaching(v bool)`

SetEnableDashboardResponseCaching sets EnableDashboardResponseCaching field to given value.

### HasEnableDashboardResponseCaching

`func (o *ModelServerConfiguration) HasEnableDashboardResponseCaching() bool`

HasEnableDashboardResponseCaching returns a boolean if a field has been set.

### GetDashboardSourcePath

`func (o *ModelServerConfiguration) GetDashboardSourcePath() string`

GetDashboardSourcePath returns the DashboardSourcePath field if non-nil, zero value otherwise.

### GetDashboardSourcePathOk

`func (o *ModelServerConfiguration) GetDashboardSourcePathOk() (*string, bool)`

GetDashboardSourcePathOk returns a tuple with the DashboardSourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardSourcePath

`func (o *ModelServerConfiguration) SetDashboardSourcePath(v string)`

SetDashboardSourcePath sets DashboardSourcePath field to given value.

### HasDashboardSourcePath

`func (o *ModelServerConfiguration) HasDashboardSourcePath() bool`

HasDashboardSourcePath returns a boolean if a field has been set.

### GetImageSavingConvention

`func (o *ModelServerConfiguration) GetImageSavingConvention() ModelModelImageSavingConvention`

GetImageSavingConvention returns the ImageSavingConvention field if non-nil, zero value otherwise.

### GetImageSavingConventionOk

`func (o *ModelServerConfiguration) GetImageSavingConventionOk() (*ModelModelImageSavingConvention, bool)`

GetImageSavingConventionOk returns a tuple with the ImageSavingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSavingConvention

`func (o *ModelServerConfiguration) SetImageSavingConvention(v ModelModelImageSavingConvention)`

SetImageSavingConvention sets ImageSavingConvention field to given value.

### HasImageSavingConvention

`func (o *ModelServerConfiguration) HasImageSavingConvention() bool`

HasImageSavingConvention returns a boolean if a field has been set.

### GetEnableAutomaticRestart

`func (o *ModelServerConfiguration) GetEnableAutomaticRestart() bool`

GetEnableAutomaticRestart returns the EnableAutomaticRestart field if non-nil, zero value otherwise.

### GetEnableAutomaticRestartOk

`func (o *ModelServerConfiguration) GetEnableAutomaticRestartOk() (*bool, bool)`

GetEnableAutomaticRestartOk returns a tuple with the EnableAutomaticRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticRestart

`func (o *ModelServerConfiguration) SetEnableAutomaticRestart(v bool)`

SetEnableAutomaticRestart sets EnableAutomaticRestart field to given value.

### HasEnableAutomaticRestart

`func (o *ModelServerConfiguration) HasEnableAutomaticRestart() bool`

HasEnableAutomaticRestart returns a boolean if a field has been set.

### GetServerName

`func (o *ModelServerConfiguration) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ModelServerConfiguration) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ModelServerConfiguration) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ModelServerConfiguration) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPreferredDetectedRemoteAddressFamily

`func (o *ModelServerConfiguration) GetPreferredDetectedRemoteAddressFamily() ModelModelNetSocketsAddressFamily`

GetPreferredDetectedRemoteAddressFamily returns the PreferredDetectedRemoteAddressFamily field if non-nil, zero value otherwise.

### GetPreferredDetectedRemoteAddressFamilyOk

`func (o *ModelServerConfiguration) GetPreferredDetectedRemoteAddressFamilyOk() (*ModelModelNetSocketsAddressFamily, bool)`

GetPreferredDetectedRemoteAddressFamilyOk returns a tuple with the PreferredDetectedRemoteAddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredDetectedRemoteAddressFamily

`func (o *ModelServerConfiguration) SetPreferredDetectedRemoteAddressFamily(v ModelModelNetSocketsAddressFamily)`

SetPreferredDetectedRemoteAddressFamily sets PreferredDetectedRemoteAddressFamily field to given value.

### HasPreferredDetectedRemoteAddressFamily

`func (o *ModelServerConfiguration) HasPreferredDetectedRemoteAddressFamily() bool`

HasPreferredDetectedRemoteAddressFamily returns a boolean if a field has been set.

### GetWanDdns

`func (o *ModelServerConfiguration) GetWanDdns() string`

GetWanDdns returns the WanDdns field if non-nil, zero value otherwise.

### GetWanDdnsOk

`func (o *ModelServerConfiguration) GetWanDdnsOk() (*string, bool)`

GetWanDdnsOk returns a tuple with the WanDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanDdns

`func (o *ModelServerConfiguration) SetWanDdns(v string)`

SetWanDdns sets WanDdns field to given value.

### HasWanDdns

`func (o *ModelServerConfiguration) HasWanDdns() bool`

HasWanDdns returns a boolean if a field has been set.

### GetUICulture

`func (o *ModelServerConfiguration) GetUICulture() string`

GetUICulture returns the UICulture field if non-nil, zero value otherwise.

### GetUICultureOk

`func (o *ModelServerConfiguration) GetUICultureOk() (*string, bool)`

GetUICultureOk returns a tuple with the UICulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUICulture

`func (o *ModelServerConfiguration) SetUICulture(v string)`

SetUICulture sets UICulture field to given value.

### HasUICulture

`func (o *ModelServerConfiguration) HasUICulture() bool`

HasUICulture returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *ModelServerConfiguration) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *ModelServerConfiguration) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *ModelServerConfiguration) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *ModelServerConfiguration) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetLocalNetworkSubnets

`func (o *ModelServerConfiguration) GetLocalNetworkSubnets() []string`

GetLocalNetworkSubnets returns the LocalNetworkSubnets field if non-nil, zero value otherwise.

### GetLocalNetworkSubnetsOk

`func (o *ModelServerConfiguration) GetLocalNetworkSubnetsOk() (*[]string, bool)`

GetLocalNetworkSubnetsOk returns a tuple with the LocalNetworkSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkSubnets

`func (o *ModelServerConfiguration) SetLocalNetworkSubnets(v []string)`

SetLocalNetworkSubnets sets LocalNetworkSubnets field to given value.

### HasLocalNetworkSubnets

`func (o *ModelServerConfiguration) HasLocalNetworkSubnets() bool`

HasLocalNetworkSubnets returns a boolean if a field has been set.

### GetLocalNetworkAddresses

`func (o *ModelServerConfiguration) GetLocalNetworkAddresses() []string`

GetLocalNetworkAddresses returns the LocalNetworkAddresses field if non-nil, zero value otherwise.

### GetLocalNetworkAddressesOk

`func (o *ModelServerConfiguration) GetLocalNetworkAddressesOk() (*[]string, bool)`

GetLocalNetworkAddressesOk returns a tuple with the LocalNetworkAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNetworkAddresses

`func (o *ModelServerConfiguration) SetLocalNetworkAddresses(v []string)`

SetLocalNetworkAddresses sets LocalNetworkAddresses field to given value.

### HasLocalNetworkAddresses

`func (o *ModelServerConfiguration) HasLocalNetworkAddresses() bool`

HasLocalNetworkAddresses returns a boolean if a field has been set.

### GetEnableExternalContentInSuggestions

`func (o *ModelServerConfiguration) GetEnableExternalContentInSuggestions() bool`

GetEnableExternalContentInSuggestions returns the EnableExternalContentInSuggestions field if non-nil, zero value otherwise.

### GetEnableExternalContentInSuggestionsOk

`func (o *ModelServerConfiguration) GetEnableExternalContentInSuggestionsOk() (*bool, bool)`

GetEnableExternalContentInSuggestionsOk returns a tuple with the EnableExternalContentInSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalContentInSuggestions

`func (o *ModelServerConfiguration) SetEnableExternalContentInSuggestions(v bool)`

SetEnableExternalContentInSuggestions sets EnableExternalContentInSuggestions field to given value.

### HasEnableExternalContentInSuggestions

`func (o *ModelServerConfiguration) HasEnableExternalContentInSuggestions() bool`

HasEnableExternalContentInSuggestions returns a boolean if a field has been set.

### GetRequireHttps

`func (o *ModelServerConfiguration) GetRequireHttps() bool`

GetRequireHttps returns the RequireHttps field if non-nil, zero value otherwise.

### GetRequireHttpsOk

`func (o *ModelServerConfiguration) GetRequireHttpsOk() (*bool, bool)`

GetRequireHttpsOk returns a tuple with the RequireHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHttps

`func (o *ModelServerConfiguration) SetRequireHttps(v bool)`

SetRequireHttps sets RequireHttps field to given value.

### HasRequireHttps

`func (o *ModelServerConfiguration) HasRequireHttps() bool`

HasRequireHttps returns a boolean if a field has been set.

### GetIsBehindProxy

`func (o *ModelServerConfiguration) GetIsBehindProxy() bool`

GetIsBehindProxy returns the IsBehindProxy field if non-nil, zero value otherwise.

### GetIsBehindProxyOk

`func (o *ModelServerConfiguration) GetIsBehindProxyOk() (*bool, bool)`

GetIsBehindProxyOk returns a tuple with the IsBehindProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBehindProxy

`func (o *ModelServerConfiguration) SetIsBehindProxy(v bool)`

SetIsBehindProxy sets IsBehindProxy field to given value.

### HasIsBehindProxy

`func (o *ModelServerConfiguration) HasIsBehindProxy() bool`

HasIsBehindProxy returns a boolean if a field has been set.

### GetRemoteIPFilter

`func (o *ModelServerConfiguration) GetRemoteIPFilter() []string`

GetRemoteIPFilter returns the RemoteIPFilter field if non-nil, zero value otherwise.

### GetRemoteIPFilterOk

`func (o *ModelServerConfiguration) GetRemoteIPFilterOk() (*[]string, bool)`

GetRemoteIPFilterOk returns a tuple with the RemoteIPFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIPFilter

`func (o *ModelServerConfiguration) SetRemoteIPFilter(v []string)`

SetRemoteIPFilter sets RemoteIPFilter field to given value.

### HasRemoteIPFilter

`func (o *ModelServerConfiguration) HasRemoteIPFilter() bool`

HasRemoteIPFilter returns a boolean if a field has been set.

### GetIsRemoteIPFilterBlacklist

`func (o *ModelServerConfiguration) GetIsRemoteIPFilterBlacklist() bool`

GetIsRemoteIPFilterBlacklist returns the IsRemoteIPFilterBlacklist field if non-nil, zero value otherwise.

### GetIsRemoteIPFilterBlacklistOk

`func (o *ModelServerConfiguration) GetIsRemoteIPFilterBlacklistOk() (*bool, bool)`

GetIsRemoteIPFilterBlacklistOk returns a tuple with the IsRemoteIPFilterBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemoteIPFilterBlacklist

`func (o *ModelServerConfiguration) SetIsRemoteIPFilterBlacklist(v bool)`

SetIsRemoteIPFilterBlacklist sets IsRemoteIPFilterBlacklist field to given value.

### HasIsRemoteIPFilterBlacklist

`func (o *ModelServerConfiguration) HasIsRemoteIPFilterBlacklist() bool`

HasIsRemoteIPFilterBlacklist returns a boolean if a field has been set.

### GetImageExtractionTimeoutMs

`func (o *ModelServerConfiguration) GetImageExtractionTimeoutMs() int32`

GetImageExtractionTimeoutMs returns the ImageExtractionTimeoutMs field if non-nil, zero value otherwise.

### GetImageExtractionTimeoutMsOk

`func (o *ModelServerConfiguration) GetImageExtractionTimeoutMsOk() (*int32, bool)`

GetImageExtractionTimeoutMsOk returns a tuple with the ImageExtractionTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractionTimeoutMs

`func (o *ModelServerConfiguration) SetImageExtractionTimeoutMs(v int32)`

SetImageExtractionTimeoutMs sets ImageExtractionTimeoutMs field to given value.

### HasImageExtractionTimeoutMs

`func (o *ModelServerConfiguration) HasImageExtractionTimeoutMs() bool`

HasImageExtractionTimeoutMs returns a boolean if a field has been set.

### GetPathSubstitutions

`func (o *ModelServerConfiguration) GetPathSubstitutions() []ModelModelPathSubstitution`

GetPathSubstitutions returns the PathSubstitutions field if non-nil, zero value otherwise.

### GetPathSubstitutionsOk

`func (o *ModelServerConfiguration) GetPathSubstitutionsOk() (*[]ModelModelPathSubstitution, bool)`

GetPathSubstitutionsOk returns a tuple with the PathSubstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSubstitutions

`func (o *ModelServerConfiguration) SetPathSubstitutions(v []ModelModelPathSubstitution)`

SetPathSubstitutions sets PathSubstitutions field to given value.

### HasPathSubstitutions

`func (o *ModelServerConfiguration) HasPathSubstitutions() bool`

HasPathSubstitutions returns a boolean if a field has been set.

### GetUninstalledPlugins

`func (o *ModelServerConfiguration) GetUninstalledPlugins() []string`

GetUninstalledPlugins returns the UninstalledPlugins field if non-nil, zero value otherwise.

### GetUninstalledPluginsOk

`func (o *ModelServerConfiguration) GetUninstalledPluginsOk() (*[]string, bool)`

GetUninstalledPluginsOk returns a tuple with the UninstalledPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalledPlugins

`func (o *ModelServerConfiguration) SetUninstalledPlugins(v []string)`

SetUninstalledPlugins sets UninstalledPlugins field to given value.

### HasUninstalledPlugins

`func (o *ModelServerConfiguration) HasUninstalledPlugins() bool`

HasUninstalledPlugins returns a boolean if a field has been set.

### GetCollapseVideoFolders

`func (o *ModelServerConfiguration) GetCollapseVideoFolders() bool`

GetCollapseVideoFolders returns the CollapseVideoFolders field if non-nil, zero value otherwise.

### GetCollapseVideoFoldersOk

`func (o *ModelServerConfiguration) GetCollapseVideoFoldersOk() (*bool, bool)`

GetCollapseVideoFoldersOk returns a tuple with the CollapseVideoFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapseVideoFolders

`func (o *ModelServerConfiguration) SetCollapseVideoFolders(v bool)`

SetCollapseVideoFolders sets CollapseVideoFolders field to given value.

### HasCollapseVideoFolders

`func (o *ModelServerConfiguration) HasCollapseVideoFolders() bool`

HasCollapseVideoFolders returns a boolean if a field has been set.

### GetEnableOriginalTrackTitles

`func (o *ModelServerConfiguration) GetEnableOriginalTrackTitles() bool`

GetEnableOriginalTrackTitles returns the EnableOriginalTrackTitles field if non-nil, zero value otherwise.

### GetEnableOriginalTrackTitlesOk

`func (o *ModelServerConfiguration) GetEnableOriginalTrackTitlesOk() (*bool, bool)`

GetEnableOriginalTrackTitlesOk returns a tuple with the EnableOriginalTrackTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOriginalTrackTitles

`func (o *ModelServerConfiguration) SetEnableOriginalTrackTitles(v bool)`

SetEnableOriginalTrackTitles sets EnableOriginalTrackTitles field to given value.

### HasEnableOriginalTrackTitles

`func (o *ModelServerConfiguration) HasEnableOriginalTrackTitles() bool`

HasEnableOriginalTrackTitles returns a boolean if a field has been set.

### GetVacuumDatabaseOnStartup

`func (o *ModelServerConfiguration) GetVacuumDatabaseOnStartup() bool`

GetVacuumDatabaseOnStartup returns the VacuumDatabaseOnStartup field if non-nil, zero value otherwise.

### GetVacuumDatabaseOnStartupOk

`func (o *ModelServerConfiguration) GetVacuumDatabaseOnStartupOk() (*bool, bool)`

GetVacuumDatabaseOnStartupOk returns a tuple with the VacuumDatabaseOnStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacuumDatabaseOnStartup

`func (o *ModelServerConfiguration) SetVacuumDatabaseOnStartup(v bool)`

SetVacuumDatabaseOnStartup sets VacuumDatabaseOnStartup field to given value.

### HasVacuumDatabaseOnStartup

`func (o *ModelServerConfiguration) HasVacuumDatabaseOnStartup() bool`

HasVacuumDatabaseOnStartup returns a boolean if a field has been set.

### GetSimultaneousStreamLimit

`func (o *ModelServerConfiguration) GetSimultaneousStreamLimit() int32`

GetSimultaneousStreamLimit returns the SimultaneousStreamLimit field if non-nil, zero value otherwise.

### GetSimultaneousStreamLimitOk

`func (o *ModelServerConfiguration) GetSimultaneousStreamLimitOk() (*int32, bool)`

GetSimultaneousStreamLimitOk returns a tuple with the SimultaneousStreamLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousStreamLimit

`func (o *ModelServerConfiguration) SetSimultaneousStreamLimit(v int32)`

SetSimultaneousStreamLimit sets SimultaneousStreamLimit field to given value.

### HasSimultaneousStreamLimit

`func (o *ModelServerConfiguration) HasSimultaneousStreamLimit() bool`

HasSimultaneousStreamLimit returns a boolean if a field has been set.

### GetDatabaseCacheSizeMB

`func (o *ModelServerConfiguration) GetDatabaseCacheSizeMB() int32`

GetDatabaseCacheSizeMB returns the DatabaseCacheSizeMB field if non-nil, zero value otherwise.

### GetDatabaseCacheSizeMBOk

`func (o *ModelServerConfiguration) GetDatabaseCacheSizeMBOk() (*int32, bool)`

GetDatabaseCacheSizeMBOk returns a tuple with the DatabaseCacheSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCacheSizeMB

`func (o *ModelServerConfiguration) SetDatabaseCacheSizeMB(v int32)`

SetDatabaseCacheSizeMB sets DatabaseCacheSizeMB field to given value.

### HasDatabaseCacheSizeMB

`func (o *ModelServerConfiguration) HasDatabaseCacheSizeMB() bool`

HasDatabaseCacheSizeMB returns a boolean if a field has been set.

### GetEnableSqLiteMmio

`func (o *ModelServerConfiguration) GetEnableSqLiteMmio() bool`

GetEnableSqLiteMmio returns the EnableSqLiteMmio field if non-nil, zero value otherwise.

### GetEnableSqLiteMmioOk

`func (o *ModelServerConfiguration) GetEnableSqLiteMmioOk() (*bool, bool)`

GetEnableSqLiteMmioOk returns a tuple with the EnableSqLiteMmio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSqLiteMmio

`func (o *ModelServerConfiguration) SetEnableSqLiteMmio(v bool)`

SetEnableSqLiteMmio sets EnableSqLiteMmio field to given value.

### HasEnableSqLiteMmio

`func (o *ModelServerConfiguration) HasEnableSqLiteMmio() bool`

HasEnableSqLiteMmio returns a boolean if a field has been set.

### GetPlaylistsUpgradedToM3U

`func (o *ModelServerConfiguration) GetPlaylistsUpgradedToM3U() bool`

GetPlaylistsUpgradedToM3U returns the PlaylistsUpgradedToM3U field if non-nil, zero value otherwise.

### GetPlaylistsUpgradedToM3UOk

`func (o *ModelServerConfiguration) GetPlaylistsUpgradedToM3UOk() (*bool, bool)`

GetPlaylistsUpgradedToM3UOk returns a tuple with the PlaylistsUpgradedToM3U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistsUpgradedToM3U

`func (o *ModelServerConfiguration) SetPlaylistsUpgradedToM3U(v bool)`

SetPlaylistsUpgradedToM3U sets PlaylistsUpgradedToM3U field to given value.

### HasPlaylistsUpgradedToM3U

`func (o *ModelServerConfiguration) HasPlaylistsUpgradedToM3U() bool`

HasPlaylistsUpgradedToM3U returns a boolean if a field has been set.

### GetImageExtractorUpgraded

`func (o *ModelServerConfiguration) GetImageExtractorUpgraded() bool`

GetImageExtractorUpgraded returns the ImageExtractorUpgraded field if non-nil, zero value otherwise.

### GetImageExtractorUpgradedOk

`func (o *ModelServerConfiguration) GetImageExtractorUpgradedOk() (*bool, bool)`

GetImageExtractorUpgradedOk returns a tuple with the ImageExtractorUpgraded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractorUpgraded

`func (o *ModelServerConfiguration) SetImageExtractorUpgraded(v bool)`

SetImageExtractorUpgraded sets ImageExtractorUpgraded field to given value.

### HasImageExtractorUpgraded

`func (o *ModelServerConfiguration) HasImageExtractorUpgraded() bool`

HasImageExtractorUpgraded returns a boolean if a field has been set.

### GetEnablePeopleLetterSubFolders

`func (o *ModelServerConfiguration) GetEnablePeopleLetterSubFolders() bool`

GetEnablePeopleLetterSubFolders returns the EnablePeopleLetterSubFolders field if non-nil, zero value otherwise.

### GetEnablePeopleLetterSubFoldersOk

`func (o *ModelServerConfiguration) GetEnablePeopleLetterSubFoldersOk() (*bool, bool)`

GetEnablePeopleLetterSubFoldersOk returns a tuple with the EnablePeopleLetterSubFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePeopleLetterSubFolders

`func (o *ModelServerConfiguration) SetEnablePeopleLetterSubFolders(v bool)`

SetEnablePeopleLetterSubFolders sets EnablePeopleLetterSubFolders field to given value.

### HasEnablePeopleLetterSubFolders

`func (o *ModelServerConfiguration) HasEnablePeopleLetterSubFolders() bool`

HasEnablePeopleLetterSubFolders returns a boolean if a field has been set.

### GetOptimizeDatabaseOnShutdown

`func (o *ModelServerConfiguration) GetOptimizeDatabaseOnShutdown() bool`

GetOptimizeDatabaseOnShutdown returns the OptimizeDatabaseOnShutdown field if non-nil, zero value otherwise.

### GetOptimizeDatabaseOnShutdownOk

`func (o *ModelServerConfiguration) GetOptimizeDatabaseOnShutdownOk() (*bool, bool)`

GetOptimizeDatabaseOnShutdownOk returns a tuple with the OptimizeDatabaseOnShutdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeDatabaseOnShutdown

`func (o *ModelServerConfiguration) SetOptimizeDatabaseOnShutdown(v bool)`

SetOptimizeDatabaseOnShutdown sets OptimizeDatabaseOnShutdown field to given value.

### HasOptimizeDatabaseOnShutdown

`func (o *ModelServerConfiguration) HasOptimizeDatabaseOnShutdown() bool`

HasOptimizeDatabaseOnShutdown returns a boolean if a field has been set.

### GetDatabaseAnalysisLimit

`func (o *ModelServerConfiguration) GetDatabaseAnalysisLimit() int32`

GetDatabaseAnalysisLimit returns the DatabaseAnalysisLimit field if non-nil, zero value otherwise.

### GetDatabaseAnalysisLimitOk

`func (o *ModelServerConfiguration) GetDatabaseAnalysisLimitOk() (*int32, bool)`

GetDatabaseAnalysisLimitOk returns a tuple with the DatabaseAnalysisLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseAnalysisLimit

`func (o *ModelServerConfiguration) SetDatabaseAnalysisLimit(v int32)`

SetDatabaseAnalysisLimit sets DatabaseAnalysisLimit field to given value.

### HasDatabaseAnalysisLimit

`func (o *ModelServerConfiguration) HasDatabaseAnalysisLimit() bool`

HasDatabaseAnalysisLimit returns a boolean if a field has been set.

### GetDisableAsyncIO

`func (o *ModelServerConfiguration) GetDisableAsyncIO() bool`

GetDisableAsyncIO returns the DisableAsyncIO field if non-nil, zero value otherwise.

### GetDisableAsyncIOOk

`func (o *ModelServerConfiguration) GetDisableAsyncIOOk() (*bool, bool)`

GetDisableAsyncIOOk returns a tuple with the DisableAsyncIO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAsyncIO

`func (o *ModelServerConfiguration) SetDisableAsyncIO(v bool)`

SetDisableAsyncIO sets DisableAsyncIO field to given value.

### HasDisableAsyncIO

`func (o *ModelServerConfiguration) HasDisableAsyncIO() bool`

HasDisableAsyncIO returns a boolean if a field has been set.

### GetMigratedToUserItemShares

`func (o *ModelServerConfiguration) GetMigratedToUserItemShares() bool`

GetMigratedToUserItemShares returns the MigratedToUserItemShares field if non-nil, zero value otherwise.

### GetMigratedToUserItemSharesOk

`func (o *ModelServerConfiguration) GetMigratedToUserItemSharesOk() (*bool, bool)`

GetMigratedToUserItemSharesOk returns a tuple with the MigratedToUserItemShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedToUserItemShares

`func (o *ModelServerConfiguration) SetMigratedToUserItemShares(v bool)`

SetMigratedToUserItemShares sets MigratedToUserItemShares field to given value.

### HasMigratedToUserItemShares

`func (o *ModelServerConfiguration) HasMigratedToUserItemShares() bool`

HasMigratedToUserItemShares returns a boolean if a field has been set.

### GetMigratedLibraryOptionsToDb

`func (o *ModelServerConfiguration) GetMigratedLibraryOptionsToDb() bool`

GetMigratedLibraryOptionsToDb returns the MigratedLibraryOptionsToDb field if non-nil, zero value otherwise.

### GetMigratedLibraryOptionsToDbOk

`func (o *ModelServerConfiguration) GetMigratedLibraryOptionsToDbOk() (*bool, bool)`

GetMigratedLibraryOptionsToDbOk returns a tuple with the MigratedLibraryOptionsToDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedLibraryOptionsToDb

`func (o *ModelServerConfiguration) SetMigratedLibraryOptionsToDb(v bool)`

SetMigratedLibraryOptionsToDb sets MigratedLibraryOptionsToDb field to given value.

### HasMigratedLibraryOptionsToDb

`func (o *ModelServerConfiguration) HasMigratedLibraryOptionsToDb() bool`

HasMigratedLibraryOptionsToDb returns a boolean if a field has been set.

### GetAllowLegacyLocalNetworkPassword

`func (o *ModelServerConfiguration) GetAllowLegacyLocalNetworkPassword() bool`

GetAllowLegacyLocalNetworkPassword returns the AllowLegacyLocalNetworkPassword field if non-nil, zero value otherwise.

### GetAllowLegacyLocalNetworkPasswordOk

`func (o *ModelServerConfiguration) GetAllowLegacyLocalNetworkPasswordOk() (*bool, bool)`

GetAllowLegacyLocalNetworkPasswordOk returns a tuple with the AllowLegacyLocalNetworkPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLegacyLocalNetworkPassword

`func (o *ModelServerConfiguration) SetAllowLegacyLocalNetworkPassword(v bool)`

SetAllowLegacyLocalNetworkPassword sets AllowLegacyLocalNetworkPassword field to given value.

### HasAllowLegacyLocalNetworkPassword

`func (o *ModelServerConfiguration) HasAllowLegacyLocalNetworkPassword() bool`

HasAllowLegacyLocalNetworkPassword returns a boolean if a field has been set.

### GetEnableSavedMetadataForPeople

`func (o *ModelServerConfiguration) GetEnableSavedMetadataForPeople() bool`

GetEnableSavedMetadataForPeople returns the EnableSavedMetadataForPeople field if non-nil, zero value otherwise.

### GetEnableSavedMetadataForPeopleOk

`func (o *ModelServerConfiguration) GetEnableSavedMetadataForPeopleOk() (*bool, bool)`

GetEnableSavedMetadataForPeopleOk returns a tuple with the EnableSavedMetadataForPeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSavedMetadataForPeople

`func (o *ModelServerConfiguration) SetEnableSavedMetadataForPeople(v bool)`

SetEnableSavedMetadataForPeople sets EnableSavedMetadataForPeople field to given value.

### HasEnableSavedMetadataForPeople

`func (o *ModelServerConfiguration) HasEnableSavedMetadataForPeople() bool`

HasEnableSavedMetadataForPeople returns a boolean if a field has been set.

### GetProxyHeaderMode

`func (o *ModelServerConfiguration) GetProxyHeaderMode() ModelModelProxyHeaderMode`

GetProxyHeaderMode returns the ProxyHeaderMode field if non-nil, zero value otherwise.

### GetProxyHeaderModeOk

`func (o *ModelServerConfiguration) GetProxyHeaderModeOk() (*ModelModelProxyHeaderMode, bool)`

GetProxyHeaderModeOk returns a tuple with the ProxyHeaderMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyHeaderMode

`func (o *ModelServerConfiguration) SetProxyHeaderMode(v ModelModelProxyHeaderMode)`

SetProxyHeaderMode sets ProxyHeaderMode field to given value.

### HasProxyHeaderMode

`func (o *ModelServerConfiguration) HasProxyHeaderMode() bool`

HasProxyHeaderMode returns a boolean if a field has been set.

### GetEnableDebugLevelLogging

`func (o *ModelServerConfiguration) GetEnableDebugLevelLogging() bool`

GetEnableDebugLevelLogging returns the EnableDebugLevelLogging field if non-nil, zero value otherwise.

### GetEnableDebugLevelLoggingOk

`func (o *ModelServerConfiguration) GetEnableDebugLevelLoggingOk() (*bool, bool)`

GetEnableDebugLevelLoggingOk returns a tuple with the EnableDebugLevelLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugLevelLogging

`func (o *ModelServerConfiguration) SetEnableDebugLevelLogging(v bool)`

SetEnableDebugLevelLogging sets EnableDebugLevelLogging field to given value.

### HasEnableDebugLevelLogging

`func (o *ModelServerConfiguration) HasEnableDebugLevelLogging() bool`

HasEnableDebugLevelLogging returns a boolean if a field has been set.

### GetRevertDebugLogging

`func (o *ModelServerConfiguration) GetRevertDebugLogging() string`

GetRevertDebugLogging returns the RevertDebugLogging field if non-nil, zero value otherwise.

### GetRevertDebugLoggingOk

`func (o *ModelServerConfiguration) GetRevertDebugLoggingOk() (*string, bool)`

GetRevertDebugLoggingOk returns a tuple with the RevertDebugLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertDebugLogging

`func (o *ModelServerConfiguration) SetRevertDebugLogging(v string)`

SetRevertDebugLogging sets RevertDebugLogging field to given value.

### HasRevertDebugLogging

`func (o *ModelServerConfiguration) HasRevertDebugLogging() bool`

HasRevertDebugLogging returns a boolean if a field has been set.

### GetEnableAutoUpdate

`func (o *ModelServerConfiguration) GetEnableAutoUpdate() bool`

GetEnableAutoUpdate returns the EnableAutoUpdate field if non-nil, zero value otherwise.

### GetEnableAutoUpdateOk

`func (o *ModelServerConfiguration) GetEnableAutoUpdateOk() (*bool, bool)`

GetEnableAutoUpdateOk returns a tuple with the EnableAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoUpdate

`func (o *ModelServerConfiguration) SetEnableAutoUpdate(v bool)`

SetEnableAutoUpdate sets EnableAutoUpdate field to given value.

### HasEnableAutoUpdate

`func (o *ModelServerConfiguration) HasEnableAutoUpdate() bool`

HasEnableAutoUpdate returns a boolean if a field has been set.

### GetLogFileRetentionDays

`func (o *ModelServerConfiguration) GetLogFileRetentionDays() int32`

GetLogFileRetentionDays returns the LogFileRetentionDays field if non-nil, zero value otherwise.

### GetLogFileRetentionDaysOk

`func (o *ModelServerConfiguration) GetLogFileRetentionDaysOk() (*int32, bool)`

GetLogFileRetentionDaysOk returns a tuple with the LogFileRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileRetentionDays

`func (o *ModelServerConfiguration) SetLogFileRetentionDays(v int32)`

SetLogFileRetentionDays sets LogFileRetentionDays field to given value.

### HasLogFileRetentionDays

`func (o *ModelServerConfiguration) HasLogFileRetentionDays() bool`

HasLogFileRetentionDays returns a boolean if a field has been set.

### GetRunAtStartup

`func (o *ModelServerConfiguration) GetRunAtStartup() bool`

GetRunAtStartup returns the RunAtStartup field if non-nil, zero value otherwise.

### GetRunAtStartupOk

`func (o *ModelServerConfiguration) GetRunAtStartupOk() (*bool, bool)`

GetRunAtStartupOk returns a tuple with the RunAtStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAtStartup

`func (o *ModelServerConfiguration) SetRunAtStartup(v bool)`

SetRunAtStartup sets RunAtStartup field to given value.

### HasRunAtStartup

`func (o *ModelServerConfiguration) HasRunAtStartup() bool`

HasRunAtStartup returns a boolean if a field has been set.

### GetIsStartupWizardCompleted

`func (o *ModelServerConfiguration) GetIsStartupWizardCompleted() bool`

GetIsStartupWizardCompleted returns the IsStartupWizardCompleted field if non-nil, zero value otherwise.

### GetIsStartupWizardCompletedOk

`func (o *ModelServerConfiguration) GetIsStartupWizardCompletedOk() (*bool, bool)`

GetIsStartupWizardCompletedOk returns a tuple with the IsStartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartupWizardCompleted

`func (o *ModelServerConfiguration) SetIsStartupWizardCompleted(v bool)`

SetIsStartupWizardCompleted sets IsStartupWizardCompleted field to given value.

### HasIsStartupWizardCompleted

`func (o *ModelServerConfiguration) HasIsStartupWizardCompleted() bool`

HasIsStartupWizardCompleted returns a boolean if a field has been set.

### GetCachePath

`func (o *ModelServerConfiguration) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *ModelServerConfiguration) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *ModelServerConfiguration) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *ModelServerConfiguration) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


