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

// checks if the ServiceRadarrRadarrIdGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceRadarrRadarrIdGet200Response{}

// ServiceRadarrRadarrIdGet200Response struct for ServiceRadarrRadarrIdGet200Response
type ServiceRadarrRadarrIdGet200Response struct {
	Server *RadarrSettings `json:"server,omitempty"`
	Profiles *ServiceProfile `json:"profiles,omitempty"`
}

// NewServiceRadarrRadarrIdGet200Response instantiates a new ServiceRadarrRadarrIdGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRadarrRadarrIdGet200Response() *ServiceRadarrRadarrIdGet200Response {
	this := ServiceRadarrRadarrIdGet200Response{}
	return &this
}

// NewServiceRadarrRadarrIdGet200ResponseWithDefaults instantiates a new ServiceRadarrRadarrIdGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRadarrRadarrIdGet200ResponseWithDefaults() *ServiceRadarrRadarrIdGet200Response {
	this := ServiceRadarrRadarrIdGet200Response{}
	return &this
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *ServiceRadarrRadarrIdGet200Response) GetServer() RadarrSettings {
	if o == nil || IsNil(o.Server) {
		var ret RadarrSettings
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRadarrRadarrIdGet200Response) GetServerOk() (*RadarrSettings, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *ServiceRadarrRadarrIdGet200Response) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given RadarrSettings and assigns it to the Server field.
func (o *ServiceRadarrRadarrIdGet200Response) SetServer(v RadarrSettings) {
	o.Server = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise.
func (o *ServiceRadarrRadarrIdGet200Response) GetProfiles() ServiceProfile {
	if o == nil || IsNil(o.Profiles) {
		var ret ServiceProfile
		return ret
	}
	return *o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRadarrRadarrIdGet200Response) GetProfilesOk() (*ServiceProfile, bool) {
	if o == nil || IsNil(o.Profiles) {
		return nil, false
	}
	return o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *ServiceRadarrRadarrIdGet200Response) HasProfiles() bool {
	if o != nil && !IsNil(o.Profiles) {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given ServiceProfile and assigns it to the Profiles field.
func (o *ServiceRadarrRadarrIdGet200Response) SetProfiles(v ServiceProfile) {
	o.Profiles = &v
}

func (o ServiceRadarrRadarrIdGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceRadarrRadarrIdGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !IsNil(o.Profiles) {
		toSerialize["profiles"] = o.Profiles
	}
	return toSerialize, nil
}

type NullableServiceRadarrRadarrIdGet200Response struct {
	value *ServiceRadarrRadarrIdGet200Response
	isSet bool
}

func (v NullableServiceRadarrRadarrIdGet200Response) Get() *ServiceRadarrRadarrIdGet200Response {
	return v.value
}

func (v *NullableServiceRadarrRadarrIdGet200Response) Set(val *ServiceRadarrRadarrIdGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRadarrRadarrIdGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRadarrRadarrIdGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRadarrRadarrIdGet200Response(val *ServiceRadarrRadarrIdGet200Response) *NullableServiceRadarrRadarrIdGet200Response {
	return &NullableServiceRadarrRadarrIdGet200Response{value: val, isSet: true}
}

func (v NullableServiceRadarrRadarrIdGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRadarrRadarrIdGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


