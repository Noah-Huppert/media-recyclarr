/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyseerrclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SonarrSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SonarrSettings{}

// SonarrSettings struct for SonarrSettings
type SonarrSettings struct {
	Id *float32 `json:"id,omitempty"`
	Name string `json:"name"`
	Hostname string `json:"hostname"`
	Port float32 `json:"port"`
	ApiKey string `json:"apiKey"`
	UseSsl bool `json:"useSsl"`
	BaseUrl *string `json:"baseUrl,omitempty"`
	ActiveProfileId float32 `json:"activeProfileId"`
	ActiveProfileName string `json:"activeProfileName"`
	ActiveDirectory string `json:"activeDirectory"`
	ActiveLanguageProfileId *float32 `json:"activeLanguageProfileId,omitempty"`
	ActiveAnimeProfileId NullableFloat32 `json:"activeAnimeProfileId,omitempty"`
	ActiveAnimeLanguageProfileId NullableFloat32 `json:"activeAnimeLanguageProfileId,omitempty"`
	ActiveAnimeProfileName NullableString `json:"activeAnimeProfileName,omitempty"`
	ActiveAnimeDirectory NullableString `json:"activeAnimeDirectory,omitempty"`
	Is4k bool `json:"is4k"`
	EnableSeasonFolders bool `json:"enableSeasonFolders"`
	IsDefault bool `json:"isDefault"`
	ExternalUrl *string `json:"externalUrl,omitempty"`
	SyncEnabled *bool `json:"syncEnabled,omitempty"`
	PreventSearch *bool `json:"preventSearch,omitempty"`
}

type _SonarrSettings SonarrSettings

// NewSonarrSettings instantiates a new SonarrSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSonarrSettings(name string, hostname string, port float32, apiKey string, useSsl bool, activeProfileId float32, activeProfileName string, activeDirectory string, is4k bool, enableSeasonFolders bool, isDefault bool) *SonarrSettings {
	this := SonarrSettings{}
	this.Name = name
	this.Hostname = hostname
	this.Port = port
	this.ApiKey = apiKey
	this.UseSsl = useSsl
	this.ActiveProfileId = activeProfileId
	this.ActiveProfileName = activeProfileName
	this.ActiveDirectory = activeDirectory
	this.Is4k = is4k
	this.EnableSeasonFolders = enableSeasonFolders
	this.IsDefault = isDefault
	return &this
}

// NewSonarrSettingsWithDefaults instantiates a new SonarrSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSonarrSettingsWithDefaults() *SonarrSettings {
	this := SonarrSettings{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SonarrSettings) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SonarrSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *SonarrSettings) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SonarrSettings) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SonarrSettings) SetName(v string) {
	o.Name = v
}

// GetHostname returns the Hostname field value
func (o *SonarrSettings) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *SonarrSettings) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *SonarrSettings) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *SonarrSettings) SetPort(v float32) {
	o.Port = v
}

// GetApiKey returns the ApiKey field value
func (o *SonarrSettings) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *SonarrSettings) SetApiKey(v string) {
	o.ApiKey = v
}

// GetUseSsl returns the UseSsl field value
func (o *SonarrSettings) GetUseSsl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetUseSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSsl, true
}

// SetUseSsl sets field value
func (o *SonarrSettings) SetUseSsl(v bool) {
	o.UseSsl = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SonarrSettings) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SonarrSettings) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SonarrSettings) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetActiveProfileId returns the ActiveProfileId field value
func (o *SonarrSettings) GetActiveProfileId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ActiveProfileId
}

// GetActiveProfileIdOk returns a tuple with the ActiveProfileId field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetActiveProfileIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveProfileId, true
}

// SetActiveProfileId sets field value
func (o *SonarrSettings) SetActiveProfileId(v float32) {
	o.ActiveProfileId = v
}

// GetActiveProfileName returns the ActiveProfileName field value
func (o *SonarrSettings) GetActiveProfileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveProfileName
}

// GetActiveProfileNameOk returns a tuple with the ActiveProfileName field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetActiveProfileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveProfileName, true
}

// SetActiveProfileName sets field value
func (o *SonarrSettings) SetActiveProfileName(v string) {
	o.ActiveProfileName = v
}

// GetActiveDirectory returns the ActiveDirectory field value
func (o *SonarrSettings) GetActiveDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActiveDirectory
}

// GetActiveDirectoryOk returns a tuple with the ActiveDirectory field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetActiveDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveDirectory, true
}

// SetActiveDirectory sets field value
func (o *SonarrSettings) SetActiveDirectory(v string) {
	o.ActiveDirectory = v
}

// GetActiveLanguageProfileId returns the ActiveLanguageProfileId field value if set, zero value otherwise.
func (o *SonarrSettings) GetActiveLanguageProfileId() float32 {
	if o == nil || IsNil(o.ActiveLanguageProfileId) {
		var ret float32
		return ret
	}
	return *o.ActiveLanguageProfileId
}

// GetActiveLanguageProfileIdOk returns a tuple with the ActiveLanguageProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetActiveLanguageProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ActiveLanguageProfileId) {
		return nil, false
	}
	return o.ActiveLanguageProfileId, true
}

// HasActiveLanguageProfileId returns a boolean if a field has been set.
func (o *SonarrSettings) HasActiveLanguageProfileId() bool {
	if o != nil && !IsNil(o.ActiveLanguageProfileId) {
		return true
	}

	return false
}

// SetActiveLanguageProfileId gets a reference to the given float32 and assigns it to the ActiveLanguageProfileId field.
func (o *SonarrSettings) SetActiveLanguageProfileId(v float32) {
	o.ActiveLanguageProfileId = &v
}

// GetActiveAnimeProfileId returns the ActiveAnimeProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSettings) GetActiveAnimeProfileId() float32 {
	if o == nil || IsNil(o.ActiveAnimeProfileId.Get()) {
		var ret float32
		return ret
	}
	return *o.ActiveAnimeProfileId.Get()
}

// GetActiveAnimeProfileIdOk returns a tuple with the ActiveAnimeProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSettings) GetActiveAnimeProfileIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAnimeProfileId.Get(), o.ActiveAnimeProfileId.IsSet()
}

// HasActiveAnimeProfileId returns a boolean if a field has been set.
func (o *SonarrSettings) HasActiveAnimeProfileId() bool {
	if o != nil && o.ActiveAnimeProfileId.IsSet() {
		return true
	}

	return false
}

// SetActiveAnimeProfileId gets a reference to the given NullableFloat32 and assigns it to the ActiveAnimeProfileId field.
func (o *SonarrSettings) SetActiveAnimeProfileId(v float32) {
	o.ActiveAnimeProfileId.Set(&v)
}
// SetActiveAnimeProfileIdNil sets the value for ActiveAnimeProfileId to be an explicit nil
func (o *SonarrSettings) SetActiveAnimeProfileIdNil() {
	o.ActiveAnimeProfileId.Set(nil)
}

// UnsetActiveAnimeProfileId ensures that no value is present for ActiveAnimeProfileId, not even an explicit nil
func (o *SonarrSettings) UnsetActiveAnimeProfileId() {
	o.ActiveAnimeProfileId.Unset()
}

// GetActiveAnimeLanguageProfileId returns the ActiveAnimeLanguageProfileId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSettings) GetActiveAnimeLanguageProfileId() float32 {
	if o == nil || IsNil(o.ActiveAnimeLanguageProfileId.Get()) {
		var ret float32
		return ret
	}
	return *o.ActiveAnimeLanguageProfileId.Get()
}

// GetActiveAnimeLanguageProfileIdOk returns a tuple with the ActiveAnimeLanguageProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSettings) GetActiveAnimeLanguageProfileIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAnimeLanguageProfileId.Get(), o.ActiveAnimeLanguageProfileId.IsSet()
}

// HasActiveAnimeLanguageProfileId returns a boolean if a field has been set.
func (o *SonarrSettings) HasActiveAnimeLanguageProfileId() bool {
	if o != nil && o.ActiveAnimeLanguageProfileId.IsSet() {
		return true
	}

	return false
}

// SetActiveAnimeLanguageProfileId gets a reference to the given NullableFloat32 and assigns it to the ActiveAnimeLanguageProfileId field.
func (o *SonarrSettings) SetActiveAnimeLanguageProfileId(v float32) {
	o.ActiveAnimeLanguageProfileId.Set(&v)
}
// SetActiveAnimeLanguageProfileIdNil sets the value for ActiveAnimeLanguageProfileId to be an explicit nil
func (o *SonarrSettings) SetActiveAnimeLanguageProfileIdNil() {
	o.ActiveAnimeLanguageProfileId.Set(nil)
}

// UnsetActiveAnimeLanguageProfileId ensures that no value is present for ActiveAnimeLanguageProfileId, not even an explicit nil
func (o *SonarrSettings) UnsetActiveAnimeLanguageProfileId() {
	o.ActiveAnimeLanguageProfileId.Unset()
}

// GetActiveAnimeProfileName returns the ActiveAnimeProfileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSettings) GetActiveAnimeProfileName() string {
	if o == nil || IsNil(o.ActiveAnimeProfileName.Get()) {
		var ret string
		return ret
	}
	return *o.ActiveAnimeProfileName.Get()
}

// GetActiveAnimeProfileNameOk returns a tuple with the ActiveAnimeProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSettings) GetActiveAnimeProfileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAnimeProfileName.Get(), o.ActiveAnimeProfileName.IsSet()
}

// HasActiveAnimeProfileName returns a boolean if a field has been set.
func (o *SonarrSettings) HasActiveAnimeProfileName() bool {
	if o != nil && o.ActiveAnimeProfileName.IsSet() {
		return true
	}

	return false
}

// SetActiveAnimeProfileName gets a reference to the given NullableString and assigns it to the ActiveAnimeProfileName field.
func (o *SonarrSettings) SetActiveAnimeProfileName(v string) {
	o.ActiveAnimeProfileName.Set(&v)
}
// SetActiveAnimeProfileNameNil sets the value for ActiveAnimeProfileName to be an explicit nil
func (o *SonarrSettings) SetActiveAnimeProfileNameNil() {
	o.ActiveAnimeProfileName.Set(nil)
}

// UnsetActiveAnimeProfileName ensures that no value is present for ActiveAnimeProfileName, not even an explicit nil
func (o *SonarrSettings) UnsetActiveAnimeProfileName() {
	o.ActiveAnimeProfileName.Unset()
}

// GetActiveAnimeDirectory returns the ActiveAnimeDirectory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SonarrSettings) GetActiveAnimeDirectory() string {
	if o == nil || IsNil(o.ActiveAnimeDirectory.Get()) {
		var ret string
		return ret
	}
	return *o.ActiveAnimeDirectory.Get()
}

// GetActiveAnimeDirectoryOk returns a tuple with the ActiveAnimeDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SonarrSettings) GetActiveAnimeDirectoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveAnimeDirectory.Get(), o.ActiveAnimeDirectory.IsSet()
}

// HasActiveAnimeDirectory returns a boolean if a field has been set.
func (o *SonarrSettings) HasActiveAnimeDirectory() bool {
	if o != nil && o.ActiveAnimeDirectory.IsSet() {
		return true
	}

	return false
}

// SetActiveAnimeDirectory gets a reference to the given NullableString and assigns it to the ActiveAnimeDirectory field.
func (o *SonarrSettings) SetActiveAnimeDirectory(v string) {
	o.ActiveAnimeDirectory.Set(&v)
}
// SetActiveAnimeDirectoryNil sets the value for ActiveAnimeDirectory to be an explicit nil
func (o *SonarrSettings) SetActiveAnimeDirectoryNil() {
	o.ActiveAnimeDirectory.Set(nil)
}

// UnsetActiveAnimeDirectory ensures that no value is present for ActiveAnimeDirectory, not even an explicit nil
func (o *SonarrSettings) UnsetActiveAnimeDirectory() {
	o.ActiveAnimeDirectory.Unset()
}

// GetIs4k returns the Is4k field value
func (o *SonarrSettings) GetIs4k() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Is4k
}

// GetIs4kOk returns a tuple with the Is4k field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetIs4kOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Is4k, true
}

// SetIs4k sets field value
func (o *SonarrSettings) SetIs4k(v bool) {
	o.Is4k = v
}

// GetEnableSeasonFolders returns the EnableSeasonFolders field value
func (o *SonarrSettings) GetEnableSeasonFolders() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableSeasonFolders
}

// GetEnableSeasonFoldersOk returns a tuple with the EnableSeasonFolders field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetEnableSeasonFoldersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableSeasonFolders, true
}

// SetEnableSeasonFolders sets field value
func (o *SonarrSettings) SetEnableSeasonFolders(v bool) {
	o.EnableSeasonFolders = v
}

// GetIsDefault returns the IsDefault field value
func (o *SonarrSettings) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetIsDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *SonarrSettings) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *SonarrSettings) GetExternalUrl() string {
	if o == nil || IsNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetExternalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUrl) {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *SonarrSettings) HasExternalUrl() bool {
	if o != nil && !IsNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *SonarrSettings) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

// GetSyncEnabled returns the SyncEnabled field value if set, zero value otherwise.
func (o *SonarrSettings) GetSyncEnabled() bool {
	if o == nil || IsNil(o.SyncEnabled) {
		var ret bool
		return ret
	}
	return *o.SyncEnabled
}

// GetSyncEnabledOk returns a tuple with the SyncEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetSyncEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncEnabled) {
		return nil, false
	}
	return o.SyncEnabled, true
}

// HasSyncEnabled returns a boolean if a field has been set.
func (o *SonarrSettings) HasSyncEnabled() bool {
	if o != nil && !IsNil(o.SyncEnabled) {
		return true
	}

	return false
}

// SetSyncEnabled gets a reference to the given bool and assigns it to the SyncEnabled field.
func (o *SonarrSettings) SetSyncEnabled(v bool) {
	o.SyncEnabled = &v
}

// GetPreventSearch returns the PreventSearch field value if set, zero value otherwise.
func (o *SonarrSettings) GetPreventSearch() bool {
	if o == nil || IsNil(o.PreventSearch) {
		var ret bool
		return ret
	}
	return *o.PreventSearch
}

// GetPreventSearchOk returns a tuple with the PreventSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SonarrSettings) GetPreventSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.PreventSearch) {
		return nil, false
	}
	return o.PreventSearch, true
}

// HasPreventSearch returns a boolean if a field has been set.
func (o *SonarrSettings) HasPreventSearch() bool {
	if o != nil && !IsNil(o.PreventSearch) {
		return true
	}

	return false
}

// SetPreventSearch gets a reference to the given bool and assigns it to the PreventSearch field.
func (o *SonarrSettings) SetPreventSearch(v bool) {
	o.PreventSearch = &v
}

func (o SonarrSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SonarrSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	toSerialize["apiKey"] = o.ApiKey
	toSerialize["useSsl"] = o.UseSsl
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	toSerialize["activeProfileId"] = o.ActiveProfileId
	toSerialize["activeProfileName"] = o.ActiveProfileName
	toSerialize["activeDirectory"] = o.ActiveDirectory
	if !IsNil(o.ActiveLanguageProfileId) {
		toSerialize["activeLanguageProfileId"] = o.ActiveLanguageProfileId
	}
	if o.ActiveAnimeProfileId.IsSet() {
		toSerialize["activeAnimeProfileId"] = o.ActiveAnimeProfileId.Get()
	}
	if o.ActiveAnimeLanguageProfileId.IsSet() {
		toSerialize["activeAnimeLanguageProfileId"] = o.ActiveAnimeLanguageProfileId.Get()
	}
	if o.ActiveAnimeProfileName.IsSet() {
		toSerialize["activeAnimeProfileName"] = o.ActiveAnimeProfileName.Get()
	}
	if o.ActiveAnimeDirectory.IsSet() {
		toSerialize["activeAnimeDirectory"] = o.ActiveAnimeDirectory.Get()
	}
	toSerialize["is4k"] = o.Is4k
	toSerialize["enableSeasonFolders"] = o.EnableSeasonFolders
	toSerialize["isDefault"] = o.IsDefault
	if !IsNil(o.ExternalUrl) {
		toSerialize["externalUrl"] = o.ExternalUrl
	}
	if !IsNil(o.SyncEnabled) {
		toSerialize["syncEnabled"] = o.SyncEnabled
	}
	if !IsNil(o.PreventSearch) {
		toSerialize["preventSearch"] = o.PreventSearch
	}
	return toSerialize, nil
}

func (o *SonarrSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"hostname",
		"port",
		"apiKey",
		"useSsl",
		"activeProfileId",
		"activeProfileName",
		"activeDirectory",
		"is4k",
		"enableSeasonFolders",
		"isDefault",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSonarrSettings := _SonarrSettings{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSonarrSettings)

	if err != nil {
		return err
	}

	*o = SonarrSettings(varSonarrSettings)

	return err
}

type NullableSonarrSettings struct {
	value *SonarrSettings
	isSet bool
}

func (v NullableSonarrSettings) Get() *SonarrSettings {
	return v.value
}

func (v *NullableSonarrSettings) Set(val *SonarrSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSonarrSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSonarrSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSonarrSettings(val *SonarrSettings) *NullableSonarrSettings {
	return &NullableSonarrSettings{value: val, isSet: true}
}

func (v NullableSonarrSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSonarrSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


