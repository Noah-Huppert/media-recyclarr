/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"encoding/json"
	"time"
)

// checks if the ModelUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUserDto{}

// ModelUserDto struct for ModelUserDto
type ModelUserDto struct {
	Name *string `json:"Name,omitempty"`
	ServerId *string `json:"ServerId,omitempty"`
	ServerName *string `json:"ServerName,omitempty"`
	Prefix *string `json:"Prefix,omitempty"`
	ConnectUserName *string `json:"ConnectUserName,omitempty"`
	DateCreated NullableTime `json:"DateCreated,omitempty"`
	ConnectLinkType *ModelConnectUserLinkType `json:"ConnectLinkType,omitempty"`
	Id *string `json:"Id,omitempty"`
	PrimaryImageTag *string `json:"PrimaryImageTag,omitempty"`
	HasPassword *bool `json:"HasPassword,omitempty"`
	HasConfiguredPassword *bool `json:"HasConfiguredPassword,omitempty"`
	EnableAutoLogin NullableBool `json:"EnableAutoLogin,omitempty"`
	LastLoginDate NullableTime `json:"LastLoginDate,omitempty"`
	LastActivityDate NullableTime `json:"LastActivityDate,omitempty"`
	Configuration *ModelUserConfiguration `json:"Configuration,omitempty"`
	Policy *ModelUserPolicy `json:"Policy,omitempty"`
	PrimaryImageAspectRatio NullableFloat64 `json:"PrimaryImageAspectRatio,omitempty"`
	// Deprecated
	HasConfiguredEasyPassword *bool `json:"HasConfiguredEasyPassword,omitempty"`
}

// NewModelUserDto instantiates a new ModelUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUserDto() *ModelUserDto {
	this := ModelUserDto{}
	return &this
}

// NewModelUserDtoWithDefaults instantiates a new ModelUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUserDtoWithDefaults() *ModelUserDto {
	this := ModelUserDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelUserDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelUserDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelUserDto) SetName(v string) {
	o.Name = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *ModelUserDto) GetServerId() string {
	if o == nil || IsNil(o.ServerId) {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetServerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *ModelUserDto) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *ModelUserDto) SetServerId(v string) {
	o.ServerId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *ModelUserDto) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *ModelUserDto) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *ModelUserDto) SetServerName(v string) {
	o.ServerName = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ModelUserDto) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ModelUserDto) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ModelUserDto) SetPrefix(v string) {
	o.Prefix = &v
}

// GetConnectUserName returns the ConnectUserName field value if set, zero value otherwise.
func (o *ModelUserDto) GetConnectUserName() string {
	if o == nil || IsNil(o.ConnectUserName) {
		var ret string
		return ret
	}
	return *o.ConnectUserName
}

// GetConnectUserNameOk returns a tuple with the ConnectUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetConnectUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectUserName) {
		return nil, false
	}
	return o.ConnectUserName, true
}

// HasConnectUserName returns a boolean if a field has been set.
func (o *ModelUserDto) HasConnectUserName() bool {
	if o != nil && !IsNil(o.ConnectUserName) {
		return true
	}

	return false
}

// SetConnectUserName gets a reference to the given string and assigns it to the ConnectUserName field.
func (o *ModelUserDto) SetConnectUserName(v string) {
	o.ConnectUserName = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserDto) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserDto) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ModelUserDto) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableTime and assigns it to the DateCreated field.
func (o *ModelUserDto) SetDateCreated(v time.Time) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *ModelUserDto) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *ModelUserDto) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetConnectLinkType returns the ConnectLinkType field value if set, zero value otherwise.
func (o *ModelUserDto) GetConnectLinkType() ModelConnectUserLinkType {
	if o == nil || IsNil(o.ConnectLinkType) {
		var ret ModelConnectUserLinkType
		return ret
	}
	return *o.ConnectLinkType
}

// GetConnectLinkTypeOk returns a tuple with the ConnectLinkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetConnectLinkTypeOk() (*ModelConnectUserLinkType, bool) {
	if o == nil || IsNil(o.ConnectLinkType) {
		return nil, false
	}
	return o.ConnectLinkType, true
}

// HasConnectLinkType returns a boolean if a field has been set.
func (o *ModelUserDto) HasConnectLinkType() bool {
	if o != nil && !IsNil(o.ConnectLinkType) {
		return true
	}

	return false
}

// SetConnectLinkType gets a reference to the given ModelConnectUserLinkType and assigns it to the ConnectLinkType field.
func (o *ModelUserDto) SetConnectLinkType(v ModelConnectUserLinkType) {
	o.ConnectLinkType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelUserDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelUserDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelUserDto) SetId(v string) {
	o.Id = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise.
func (o *ModelUserDto) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryImageTag) {
		return nil, false
	}
	return o.PrimaryImageTag, true
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *ModelUserDto) HasPrimaryImageTag() bool {
	if o != nil && !IsNil(o.PrimaryImageTag) {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given string and assigns it to the PrimaryImageTag field.
func (o *ModelUserDto) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *ModelUserDto) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *ModelUserDto) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *ModelUserDto) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetHasConfiguredPassword returns the HasConfiguredPassword field value if set, zero value otherwise.
func (o *ModelUserDto) GetHasConfiguredPassword() bool {
	if o == nil || IsNil(o.HasConfiguredPassword) {
		var ret bool
		return ret
	}
	return *o.HasConfiguredPassword
}

// GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetHasConfiguredPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfiguredPassword) {
		return nil, false
	}
	return o.HasConfiguredPassword, true
}

// HasHasConfiguredPassword returns a boolean if a field has been set.
func (o *ModelUserDto) HasHasConfiguredPassword() bool {
	if o != nil && !IsNil(o.HasConfiguredPassword) {
		return true
	}

	return false
}

// SetHasConfiguredPassword gets a reference to the given bool and assigns it to the HasConfiguredPassword field.
func (o *ModelUserDto) SetHasConfiguredPassword(v bool) {
	o.HasConfiguredPassword = &v
}

// GetEnableAutoLogin returns the EnableAutoLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserDto) GetEnableAutoLogin() bool {
	if o == nil || IsNil(o.EnableAutoLogin.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAutoLogin.Get()
}

// GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserDto) GetEnableAutoLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableAutoLogin.Get(), o.EnableAutoLogin.IsSet()
}

// HasEnableAutoLogin returns a boolean if a field has been set.
func (o *ModelUserDto) HasEnableAutoLogin() bool {
	if o != nil && o.EnableAutoLogin.IsSet() {
		return true
	}

	return false
}

// SetEnableAutoLogin gets a reference to the given NullableBool and assigns it to the EnableAutoLogin field.
func (o *ModelUserDto) SetEnableAutoLogin(v bool) {
	o.EnableAutoLogin.Set(&v)
}
// SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil
func (o *ModelUserDto) SetEnableAutoLoginNil() {
	o.EnableAutoLogin.Set(nil)
}

// UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
func (o *ModelUserDto) UnsetEnableAutoLogin() {
	o.EnableAutoLogin.Unset()
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserDto) GetLastLoginDate() time.Time {
	if o == nil || IsNil(o.LastLoginDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate.Get()
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserDto) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoginDate.Get(), o.LastLoginDate.IsSet()
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *ModelUserDto) HasLastLoginDate() bool {
	if o != nil && o.LastLoginDate.IsSet() {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given NullableTime and assigns it to the LastLoginDate field.
func (o *ModelUserDto) SetLastLoginDate(v time.Time) {
	o.LastLoginDate.Set(&v)
}
// SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil
func (o *ModelUserDto) SetLastLoginDateNil() {
	o.LastLoginDate.Set(nil)
}

// UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
func (o *ModelUserDto) UnsetLastLoginDate() {
	o.LastLoginDate.Unset()
}

// GetLastActivityDate returns the LastActivityDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserDto) GetLastActivityDate() time.Time {
	if o == nil || IsNil(o.LastActivityDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastActivityDate.Get()
}

// GetLastActivityDateOk returns a tuple with the LastActivityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserDto) GetLastActivityDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastActivityDate.Get(), o.LastActivityDate.IsSet()
}

// HasLastActivityDate returns a boolean if a field has been set.
func (o *ModelUserDto) HasLastActivityDate() bool {
	if o != nil && o.LastActivityDate.IsSet() {
		return true
	}

	return false
}

// SetLastActivityDate gets a reference to the given NullableTime and assigns it to the LastActivityDate field.
func (o *ModelUserDto) SetLastActivityDate(v time.Time) {
	o.LastActivityDate.Set(&v)
}
// SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil
func (o *ModelUserDto) SetLastActivityDateNil() {
	o.LastActivityDate.Set(nil)
}

// UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
func (o *ModelUserDto) UnsetLastActivityDate() {
	o.LastActivityDate.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ModelUserDto) GetConfiguration() ModelUserConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret ModelUserConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetConfigurationOk() (*ModelUserConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ModelUserDto) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given ModelUserConfiguration and assigns it to the Configuration field.
func (o *ModelUserDto) SetConfiguration(v ModelUserConfiguration) {
	o.Configuration = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *ModelUserDto) GetPolicy() ModelUserPolicy {
	if o == nil || IsNil(o.Policy) {
		var ret ModelUserPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserDto) GetPolicyOk() (*ModelUserPolicy, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *ModelUserDto) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given ModelUserPolicy and assigns it to the Policy field.
func (o *ModelUserDto) SetPolicy(v ModelUserPolicy) {
	o.Policy = &v
}

// GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelUserDto) GetPrimaryImageAspectRatio() float64 {
	if o == nil || IsNil(o.PrimaryImageAspectRatio.Get()) {
		var ret float64
		return ret
	}
	return *o.PrimaryImageAspectRatio.Get()
}

// GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelUserDto) GetPrimaryImageAspectRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageAspectRatio.Get(), o.PrimaryImageAspectRatio.IsSet()
}

// HasPrimaryImageAspectRatio returns a boolean if a field has been set.
func (o *ModelUserDto) HasPrimaryImageAspectRatio() bool {
	if o != nil && o.PrimaryImageAspectRatio.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageAspectRatio gets a reference to the given NullableFloat64 and assigns it to the PrimaryImageAspectRatio field.
func (o *ModelUserDto) SetPrimaryImageAspectRatio(v float64) {
	o.PrimaryImageAspectRatio.Set(&v)
}
// SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil
func (o *ModelUserDto) SetPrimaryImageAspectRatioNil() {
	o.PrimaryImageAspectRatio.Set(nil)
}

// UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
func (o *ModelUserDto) UnsetPrimaryImageAspectRatio() {
	o.PrimaryImageAspectRatio.Unset()
}

// GetHasConfiguredEasyPassword returns the HasConfiguredEasyPassword field value if set, zero value otherwise.
// Deprecated
func (o *ModelUserDto) GetHasConfiguredEasyPassword() bool {
	if o == nil || IsNil(o.HasConfiguredEasyPassword) {
		var ret bool
		return ret
	}
	return *o.HasConfiguredEasyPassword
}

// GetHasConfiguredEasyPasswordOk returns a tuple with the HasConfiguredEasyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ModelUserDto) GetHasConfiguredEasyPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfiguredEasyPassword) {
		return nil, false
	}
	return o.HasConfiguredEasyPassword, true
}

// HasHasConfiguredEasyPassword returns a boolean if a field has been set.
func (o *ModelUserDto) HasHasConfiguredEasyPassword() bool {
	if o != nil && !IsNil(o.HasConfiguredEasyPassword) {
		return true
	}

	return false
}

// SetHasConfiguredEasyPassword gets a reference to the given bool and assigns it to the HasConfiguredEasyPassword field.
// Deprecated
func (o *ModelUserDto) SetHasConfiguredEasyPassword(v bool) {
	o.HasConfiguredEasyPassword = &v
}

func (o ModelUserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.ServerId) {
		toSerialize["ServerId"] = o.ServerId
	}
	if !IsNil(o.ServerName) {
		toSerialize["ServerName"] = o.ServerName
	}
	if !IsNil(o.Prefix) {
		toSerialize["Prefix"] = o.Prefix
	}
	if !IsNil(o.ConnectUserName) {
		toSerialize["ConnectUserName"] = o.ConnectUserName
	}
	if o.DateCreated.IsSet() {
		toSerialize["DateCreated"] = o.DateCreated.Get()
	}
	if !IsNil(o.ConnectLinkType) {
		toSerialize["ConnectLinkType"] = o.ConnectLinkType
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.PrimaryImageTag) {
		toSerialize["PrimaryImageTag"] = o.PrimaryImageTag
	}
	if !IsNil(o.HasPassword) {
		toSerialize["HasPassword"] = o.HasPassword
	}
	if !IsNil(o.HasConfiguredPassword) {
		toSerialize["HasConfiguredPassword"] = o.HasConfiguredPassword
	}
	if o.EnableAutoLogin.IsSet() {
		toSerialize["EnableAutoLogin"] = o.EnableAutoLogin.Get()
	}
	if o.LastLoginDate.IsSet() {
		toSerialize["LastLoginDate"] = o.LastLoginDate.Get()
	}
	if o.LastActivityDate.IsSet() {
		toSerialize["LastActivityDate"] = o.LastActivityDate.Get()
	}
	if !IsNil(o.Configuration) {
		toSerialize["Configuration"] = o.Configuration
	}
	if !IsNil(o.Policy) {
		toSerialize["Policy"] = o.Policy
	}
	if o.PrimaryImageAspectRatio.IsSet() {
		toSerialize["PrimaryImageAspectRatio"] = o.PrimaryImageAspectRatio.Get()
	}
	if !IsNil(o.HasConfiguredEasyPassword) {
		toSerialize["HasConfiguredEasyPassword"] = o.HasConfiguredEasyPassword
	}
	return toSerialize, nil
}

type NullableModelUserDto struct {
	value *ModelUserDto
	isSet bool
}

func (v NullableModelUserDto) Get() *ModelUserDto {
	return v.value
}

func (v *NullableModelUserDto) Set(val *ModelUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUserDto(val *ModelUserDto) *NullableModelUserDto {
	return &NullableModelUserDto{value: val, isSet: true}
}

func (v NullableModelUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

