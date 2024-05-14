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

// checks if the ModelPublicSystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPublicSystemInfo{}

// ModelPublicSystemInfo struct for ModelPublicSystemInfo
type ModelPublicSystemInfo struct {
	LocalAddress *string `json:"LocalAddress,omitempty"`
	LocalAddresses []string `json:"LocalAddresses,omitempty"`
	WanAddress *string `json:"WanAddress,omitempty"`
	RemoteAddresses []string `json:"RemoteAddresses,omitempty"`
	ServerName *string `json:"ServerName,omitempty"`
	Version *string `json:"Version,omitempty"`
	Id *string `json:"Id,omitempty"`
}

// NewModelPublicSystemInfo instantiates a new ModelPublicSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPublicSystemInfo() *ModelPublicSystemInfo {
	this := ModelPublicSystemInfo{}
	return &this
}

// NewModelPublicSystemInfoWithDefaults instantiates a new ModelPublicSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPublicSystemInfoWithDefaults() *ModelPublicSystemInfo {
	this := ModelPublicSystemInfo{}
	return &this
}

// GetLocalAddress returns the LocalAddress field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetLocalAddress() string {
	if o == nil || IsNil(o.LocalAddress) {
		var ret string
		return ret
	}
	return *o.LocalAddress
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetLocalAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LocalAddress) {
		return nil, false
	}
	return o.LocalAddress, true
}

// HasLocalAddress returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasLocalAddress() bool {
	if o != nil && !IsNil(o.LocalAddress) {
		return true
	}

	return false
}

// SetLocalAddress gets a reference to the given string and assigns it to the LocalAddress field.
func (o *ModelPublicSystemInfo) SetLocalAddress(v string) {
	o.LocalAddress = &v
}

// GetLocalAddresses returns the LocalAddresses field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetLocalAddresses() []string {
	if o == nil || IsNil(o.LocalAddresses) {
		var ret []string
		return ret
	}
	return o.LocalAddresses
}

// GetLocalAddressesOk returns a tuple with the LocalAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetLocalAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalAddresses) {
		return nil, false
	}
	return o.LocalAddresses, true
}

// HasLocalAddresses returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasLocalAddresses() bool {
	if o != nil && !IsNil(o.LocalAddresses) {
		return true
	}

	return false
}

// SetLocalAddresses gets a reference to the given []string and assigns it to the LocalAddresses field.
func (o *ModelPublicSystemInfo) SetLocalAddresses(v []string) {
	o.LocalAddresses = v
}

// GetWanAddress returns the WanAddress field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetWanAddress() string {
	if o == nil || IsNil(o.WanAddress) {
		var ret string
		return ret
	}
	return *o.WanAddress
}

// GetWanAddressOk returns a tuple with the WanAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetWanAddressOk() (*string, bool) {
	if o == nil || IsNil(o.WanAddress) {
		return nil, false
	}
	return o.WanAddress, true
}

// HasWanAddress returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasWanAddress() bool {
	if o != nil && !IsNil(o.WanAddress) {
		return true
	}

	return false
}

// SetWanAddress gets a reference to the given string and assigns it to the WanAddress field.
func (o *ModelPublicSystemInfo) SetWanAddress(v string) {
	o.WanAddress = &v
}

// GetRemoteAddresses returns the RemoteAddresses field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetRemoteAddresses() []string {
	if o == nil || IsNil(o.RemoteAddresses) {
		var ret []string
		return ret
	}
	return o.RemoteAddresses
}

// GetRemoteAddressesOk returns a tuple with the RemoteAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetRemoteAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.RemoteAddresses) {
		return nil, false
	}
	return o.RemoteAddresses, true
}

// HasRemoteAddresses returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasRemoteAddresses() bool {
	if o != nil && !IsNil(o.RemoteAddresses) {
		return true
	}

	return false
}

// SetRemoteAddresses gets a reference to the given []string and assigns it to the RemoteAddresses field.
func (o *ModelPublicSystemInfo) SetRemoteAddresses(v []string) {
	o.RemoteAddresses = v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *ModelPublicSystemInfo) SetServerName(v string) {
	o.ServerName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ModelPublicSystemInfo) SetVersion(v string) {
	o.Version = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelPublicSystemInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPublicSystemInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelPublicSystemInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelPublicSystemInfo) SetId(v string) {
	o.Id = &v
}

func (o ModelPublicSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPublicSystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableModelPublicSystemInfo struct {
	value *ModelPublicSystemInfo
	isSet bool
}

func (v NullableModelPublicSystemInfo) Get() *ModelPublicSystemInfo {
	return v.value
}

func (v *NullableModelPublicSystemInfo) Set(val *ModelPublicSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPublicSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPublicSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPublicSystemInfo(val *ModelPublicSystemInfo) *NullableModelPublicSystemInfo {
	return &NullableModelPublicSystemInfo{value: val, isSet: true}
}

func (v NullableModelPublicSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPublicSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

