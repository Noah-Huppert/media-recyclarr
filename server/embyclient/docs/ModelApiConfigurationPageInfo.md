# ModelApiConfigurationPageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**EnableInMainMenu** | Pointer to **bool** |  | [optional] 
**EnableInUserMenu** | Pointer to **bool** |  | [optional] 
**FeatureId** | Pointer to **string** |  | [optional] 
**MenuSection** | Pointer to **string** |  | [optional] 
**MenuIcon** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ConfigurationPageType** | Pointer to [**ModelModelPluginsConfigurationPageType**](ModelPluginsConfigurationPageType.md) |  | [optional] 
**PluginId** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**NavMenuId** | Pointer to **string** |  | [optional] 
**Plugin** | Pointer to [**ModelModelCommonPluginsIPlugin**](ModelCommonPluginsIPlugin.md) |  | [optional] 
**Translations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelApiConfigurationPageInfo

`func NewModelApiConfigurationPageInfo() *ModelApiConfigurationPageInfo`

NewModelApiConfigurationPageInfo instantiates a new ModelApiConfigurationPageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelApiConfigurationPageInfoWithDefaults

`func NewModelApiConfigurationPageInfoWithDefaults() *ModelApiConfigurationPageInfo`

NewModelApiConfigurationPageInfoWithDefaults instantiates a new ModelApiConfigurationPageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelApiConfigurationPageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelApiConfigurationPageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelApiConfigurationPageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelApiConfigurationPageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnableInMainMenu

`func (o *ModelApiConfigurationPageInfo) GetEnableInMainMenu() bool`

GetEnableInMainMenu returns the EnableInMainMenu field if non-nil, zero value otherwise.

### GetEnableInMainMenuOk

`func (o *ModelApiConfigurationPageInfo) GetEnableInMainMenuOk() (*bool, bool)`

GetEnableInMainMenuOk returns a tuple with the EnableInMainMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInMainMenu

`func (o *ModelApiConfigurationPageInfo) SetEnableInMainMenu(v bool)`

SetEnableInMainMenu sets EnableInMainMenu field to given value.

### HasEnableInMainMenu

`func (o *ModelApiConfigurationPageInfo) HasEnableInMainMenu() bool`

HasEnableInMainMenu returns a boolean if a field has been set.

### GetEnableInUserMenu

`func (o *ModelApiConfigurationPageInfo) GetEnableInUserMenu() bool`

GetEnableInUserMenu returns the EnableInUserMenu field if non-nil, zero value otherwise.

### GetEnableInUserMenuOk

`func (o *ModelApiConfigurationPageInfo) GetEnableInUserMenuOk() (*bool, bool)`

GetEnableInUserMenuOk returns a tuple with the EnableInUserMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInUserMenu

`func (o *ModelApiConfigurationPageInfo) SetEnableInUserMenu(v bool)`

SetEnableInUserMenu sets EnableInUserMenu field to given value.

### HasEnableInUserMenu

`func (o *ModelApiConfigurationPageInfo) HasEnableInUserMenu() bool`

HasEnableInUserMenu returns a boolean if a field has been set.

### GetFeatureId

`func (o *ModelApiConfigurationPageInfo) GetFeatureId() string`

GetFeatureId returns the FeatureId field if non-nil, zero value otherwise.

### GetFeatureIdOk

`func (o *ModelApiConfigurationPageInfo) GetFeatureIdOk() (*string, bool)`

GetFeatureIdOk returns a tuple with the FeatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureId

`func (o *ModelApiConfigurationPageInfo) SetFeatureId(v string)`

SetFeatureId sets FeatureId field to given value.

### HasFeatureId

`func (o *ModelApiConfigurationPageInfo) HasFeatureId() bool`

HasFeatureId returns a boolean if a field has been set.

### GetMenuSection

`func (o *ModelApiConfigurationPageInfo) GetMenuSection() string`

GetMenuSection returns the MenuSection field if non-nil, zero value otherwise.

### GetMenuSectionOk

`func (o *ModelApiConfigurationPageInfo) GetMenuSectionOk() (*string, bool)`

GetMenuSectionOk returns a tuple with the MenuSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuSection

`func (o *ModelApiConfigurationPageInfo) SetMenuSection(v string)`

SetMenuSection sets MenuSection field to given value.

### HasMenuSection

`func (o *ModelApiConfigurationPageInfo) HasMenuSection() bool`

HasMenuSection returns a boolean if a field has been set.

### GetMenuIcon

`func (o *ModelApiConfigurationPageInfo) GetMenuIcon() string`

GetMenuIcon returns the MenuIcon field if non-nil, zero value otherwise.

### GetMenuIconOk

`func (o *ModelApiConfigurationPageInfo) GetMenuIconOk() (*string, bool)`

GetMenuIconOk returns a tuple with the MenuIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenuIcon

`func (o *ModelApiConfigurationPageInfo) SetMenuIcon(v string)`

SetMenuIcon sets MenuIcon field to given value.

### HasMenuIcon

`func (o *ModelApiConfigurationPageInfo) HasMenuIcon() bool`

HasMenuIcon returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelApiConfigurationPageInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelApiConfigurationPageInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelApiConfigurationPageInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelApiConfigurationPageInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetConfigurationPageType

`func (o *ModelApiConfigurationPageInfo) GetConfigurationPageType() ModelModelPluginsConfigurationPageType`

GetConfigurationPageType returns the ConfigurationPageType field if non-nil, zero value otherwise.

### GetConfigurationPageTypeOk

`func (o *ModelApiConfigurationPageInfo) GetConfigurationPageTypeOk() (*ModelModelPluginsConfigurationPageType, bool)`

GetConfigurationPageTypeOk returns a tuple with the ConfigurationPageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationPageType

`func (o *ModelApiConfigurationPageInfo) SetConfigurationPageType(v ModelModelPluginsConfigurationPageType)`

SetConfigurationPageType sets ConfigurationPageType field to given value.

### HasConfigurationPageType

`func (o *ModelApiConfigurationPageInfo) HasConfigurationPageType() bool`

HasConfigurationPageType returns a boolean if a field has been set.

### GetPluginId

`func (o *ModelApiConfigurationPageInfo) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *ModelApiConfigurationPageInfo) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *ModelApiConfigurationPageInfo) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *ModelApiConfigurationPageInfo) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetHref

`func (o *ModelApiConfigurationPageInfo) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ModelApiConfigurationPageInfo) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ModelApiConfigurationPageInfo) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ModelApiConfigurationPageInfo) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetNavMenuId

`func (o *ModelApiConfigurationPageInfo) GetNavMenuId() string`

GetNavMenuId returns the NavMenuId field if non-nil, zero value otherwise.

### GetNavMenuIdOk

`func (o *ModelApiConfigurationPageInfo) GetNavMenuIdOk() (*string, bool)`

GetNavMenuIdOk returns a tuple with the NavMenuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavMenuId

`func (o *ModelApiConfigurationPageInfo) SetNavMenuId(v string)`

SetNavMenuId sets NavMenuId field to given value.

### HasNavMenuId

`func (o *ModelApiConfigurationPageInfo) HasNavMenuId() bool`

HasNavMenuId returns a boolean if a field has been set.

### GetPlugin

`func (o *ModelApiConfigurationPageInfo) GetPlugin() ModelModelCommonPluginsIPlugin`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *ModelApiConfigurationPageInfo) GetPluginOk() (*ModelModelCommonPluginsIPlugin, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *ModelApiConfigurationPageInfo) SetPlugin(v ModelModelCommonPluginsIPlugin)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *ModelApiConfigurationPageInfo) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetTranslations

`func (o *ModelApiConfigurationPageInfo) GetTranslations() []string`

GetTranslations returns the Translations field if non-nil, zero value otherwise.

### GetTranslationsOk

`func (o *ModelApiConfigurationPageInfo) GetTranslationsOk() (*[]string, bool)`

GetTranslationsOk returns a tuple with the Translations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslations

`func (o *ModelApiConfigurationPageInfo) SetTranslations(v []string)`

SetTranslations sets Translations field to given value.

### HasTranslations

`func (o *ModelApiConfigurationPageInfo) HasTranslations() bool`

HasTranslations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


