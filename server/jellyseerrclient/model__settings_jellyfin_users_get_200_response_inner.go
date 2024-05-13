/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyseerrclient

import (
	"encoding/json"
)

// checks if the SettingsJellyfinUsersGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsJellyfinUsersGet200ResponseInner{}

// SettingsJellyfinUsersGet200ResponseInner struct for SettingsJellyfinUsersGet200ResponseInner
type SettingsJellyfinUsersGet200ResponseInner struct {
	Username *string `json:"username,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
}

// NewSettingsJellyfinUsersGet200ResponseInner instantiates a new SettingsJellyfinUsersGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsJellyfinUsersGet200ResponseInner() *SettingsJellyfinUsersGet200ResponseInner {
	this := SettingsJellyfinUsersGet200ResponseInner{}
	return &this
}

// NewSettingsJellyfinUsersGet200ResponseInnerWithDefaults instantiates a new SettingsJellyfinUsersGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsJellyfinUsersGet200ResponseInnerWithDefaults() *SettingsJellyfinUsersGet200ResponseInner {
	this := SettingsJellyfinUsersGet200ResponseInner{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SettingsJellyfinUsersGet200ResponseInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsJellyfinUsersGet200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SettingsJellyfinUsersGet200ResponseInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SettingsJellyfinUsersGet200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SettingsJellyfinUsersGet200ResponseInner) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsJellyfinUsersGet200ResponseInner) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SettingsJellyfinUsersGet200ResponseInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *SettingsJellyfinUsersGet200ResponseInner) SetUserId(v int32) {
	o.UserId = &v
}

func (o SettingsJellyfinUsersGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsJellyfinUsersGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableSettingsJellyfinUsersGet200ResponseInner struct {
	value *SettingsJellyfinUsersGet200ResponseInner
	isSet bool
}

func (v NullableSettingsJellyfinUsersGet200ResponseInner) Get() *SettingsJellyfinUsersGet200ResponseInner {
	return v.value
}

func (v *NullableSettingsJellyfinUsersGet200ResponseInner) Set(val *SettingsJellyfinUsersGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsJellyfinUsersGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsJellyfinUsersGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsJellyfinUsersGet200ResponseInner(val *SettingsJellyfinUsersGet200ResponseInner) *NullableSettingsJellyfinUsersGet200ResponseInner {
	return &NullableSettingsJellyfinUsersGet200ResponseInner{value: val, isSet: true}
}

func (v NullableSettingsJellyfinUsersGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsJellyfinUsersGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

