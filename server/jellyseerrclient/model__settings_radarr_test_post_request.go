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

// checks if the SettingsRadarrTestPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsRadarrTestPostRequest{}

// SettingsRadarrTestPostRequest struct for SettingsRadarrTestPostRequest
type SettingsRadarrTestPostRequest struct {
	Hostname string `json:"hostname"`
	Port float32 `json:"port"`
	ApiKey string `json:"apiKey"`
	UseSsl bool `json:"useSsl"`
	BaseUrl *string `json:"baseUrl,omitempty"`
}

type _SettingsRadarrTestPostRequest SettingsRadarrTestPostRequest

// NewSettingsRadarrTestPostRequest instantiates a new SettingsRadarrTestPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsRadarrTestPostRequest(hostname string, port float32, apiKey string, useSsl bool) *SettingsRadarrTestPostRequest {
	this := SettingsRadarrTestPostRequest{}
	this.Hostname = hostname
	this.Port = port
	this.ApiKey = apiKey
	this.UseSsl = useSsl
	return &this
}

// NewSettingsRadarrTestPostRequestWithDefaults instantiates a new SettingsRadarrTestPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsRadarrTestPostRequestWithDefaults() *SettingsRadarrTestPostRequest {
	this := SettingsRadarrTestPostRequest{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *SettingsRadarrTestPostRequest) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *SettingsRadarrTestPostRequest) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *SettingsRadarrTestPostRequest) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value
func (o *SettingsRadarrTestPostRequest) GetPort() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *SettingsRadarrTestPostRequest) GetPortOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *SettingsRadarrTestPostRequest) SetPort(v float32) {
	o.Port = v
}

// GetApiKey returns the ApiKey field value
func (o *SettingsRadarrTestPostRequest) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *SettingsRadarrTestPostRequest) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *SettingsRadarrTestPostRequest) SetApiKey(v string) {
	o.ApiKey = v
}

// GetUseSsl returns the UseSsl field value
func (o *SettingsRadarrTestPostRequest) GetUseSsl() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseSsl
}

// GetUseSslOk returns a tuple with the UseSsl field value
// and a boolean to check if the value has been set.
func (o *SettingsRadarrTestPostRequest) GetUseSslOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseSsl, true
}

// SetUseSsl sets field value
func (o *SettingsRadarrTestPostRequest) SetUseSsl(v bool) {
	o.UseSsl = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *SettingsRadarrTestPostRequest) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRadarrTestPostRequest) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SettingsRadarrTestPostRequest) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *SettingsRadarrTestPostRequest) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

func (o SettingsRadarrTestPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsRadarrTestPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hostname"] = o.Hostname
	toSerialize["port"] = o.Port
	toSerialize["apiKey"] = o.ApiKey
	toSerialize["useSsl"] = o.UseSsl
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	return toSerialize, nil
}

func (o *SettingsRadarrTestPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hostname",
		"port",
		"apiKey",
		"useSsl",
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

	varSettingsRadarrTestPostRequest := _SettingsRadarrTestPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSettingsRadarrTestPostRequest)

	if err != nil {
		return err
	}

	*o = SettingsRadarrTestPostRequest(varSettingsRadarrTestPostRequest)

	return err
}

type NullableSettingsRadarrTestPostRequest struct {
	value *SettingsRadarrTestPostRequest
	isSet bool
}

func (v NullableSettingsRadarrTestPostRequest) Get() *SettingsRadarrTestPostRequest {
	return v.value
}

func (v *NullableSettingsRadarrTestPostRequest) Set(val *SettingsRadarrTestPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsRadarrTestPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsRadarrTestPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsRadarrTestPostRequest(val *SettingsRadarrTestPostRequest) *NullableSettingsRadarrTestPostRequest {
	return &NullableSettingsRadarrTestPostRequest{value: val, isSet: true}
}

func (v NullableSettingsRadarrTestPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsRadarrTestPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

