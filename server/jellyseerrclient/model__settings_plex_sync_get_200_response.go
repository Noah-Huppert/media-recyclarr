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

// checks if the SettingsPlexSyncGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SettingsPlexSyncGet200Response{}

// SettingsPlexSyncGet200Response struct for SettingsPlexSyncGet200Response
type SettingsPlexSyncGet200Response struct {
	Running *bool `json:"running,omitempty"`
	Progress *float32 `json:"progress,omitempty"`
	Total *float32 `json:"total,omitempty"`
	CurrentLibrary *PlexLibrary `json:"currentLibrary,omitempty"`
	Libraries []PlexLibrary `json:"libraries,omitempty"`
}

// NewSettingsPlexSyncGet200Response instantiates a new SettingsPlexSyncGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsPlexSyncGet200Response() *SettingsPlexSyncGet200Response {
	this := SettingsPlexSyncGet200Response{}
	return &this
}

// NewSettingsPlexSyncGet200ResponseWithDefaults instantiates a new SettingsPlexSyncGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsPlexSyncGet200ResponseWithDefaults() *SettingsPlexSyncGet200Response {
	this := SettingsPlexSyncGet200Response{}
	return &this
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *SettingsPlexSyncGet200Response) GetRunning() bool {
	if o == nil || IsNil(o.Running) {
		var ret bool
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncGet200Response) GetRunningOk() (*bool, bool) {
	if o == nil || IsNil(o.Running) {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *SettingsPlexSyncGet200Response) HasRunning() bool {
	if o != nil && !IsNil(o.Running) {
		return true
	}

	return false
}

// SetRunning gets a reference to the given bool and assigns it to the Running field.
func (o *SettingsPlexSyncGet200Response) SetRunning(v bool) {
	o.Running = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *SettingsPlexSyncGet200Response) GetProgress() float32 {
	if o == nil || IsNil(o.Progress) {
		var ret float32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncGet200Response) GetProgressOk() (*float32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *SettingsPlexSyncGet200Response) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given float32 and assigns it to the Progress field.
func (o *SettingsPlexSyncGet200Response) SetProgress(v float32) {
	o.Progress = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SettingsPlexSyncGet200Response) GetTotal() float32 {
	if o == nil || IsNil(o.Total) {
		var ret float32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncGet200Response) GetTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SettingsPlexSyncGet200Response) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given float32 and assigns it to the Total field.
func (o *SettingsPlexSyncGet200Response) SetTotal(v float32) {
	o.Total = &v
}

// GetCurrentLibrary returns the CurrentLibrary field value if set, zero value otherwise.
func (o *SettingsPlexSyncGet200Response) GetCurrentLibrary() PlexLibrary {
	if o == nil || IsNil(o.CurrentLibrary) {
		var ret PlexLibrary
		return ret
	}
	return *o.CurrentLibrary
}

// GetCurrentLibraryOk returns a tuple with the CurrentLibrary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncGet200Response) GetCurrentLibraryOk() (*PlexLibrary, bool) {
	if o == nil || IsNil(o.CurrentLibrary) {
		return nil, false
	}
	return o.CurrentLibrary, true
}

// HasCurrentLibrary returns a boolean if a field has been set.
func (o *SettingsPlexSyncGet200Response) HasCurrentLibrary() bool {
	if o != nil && !IsNil(o.CurrentLibrary) {
		return true
	}

	return false
}

// SetCurrentLibrary gets a reference to the given PlexLibrary and assigns it to the CurrentLibrary field.
func (o *SettingsPlexSyncGet200Response) SetCurrentLibrary(v PlexLibrary) {
	o.CurrentLibrary = &v
}

// GetLibraries returns the Libraries field value if set, zero value otherwise.
func (o *SettingsPlexSyncGet200Response) GetLibraries() []PlexLibrary {
	if o == nil || IsNil(o.Libraries) {
		var ret []PlexLibrary
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsPlexSyncGet200Response) GetLibrariesOk() ([]PlexLibrary, bool) {
	if o == nil || IsNil(o.Libraries) {
		return nil, false
	}
	return o.Libraries, true
}

// HasLibraries returns a boolean if a field has been set.
func (o *SettingsPlexSyncGet200Response) HasLibraries() bool {
	if o != nil && !IsNil(o.Libraries) {
		return true
	}

	return false
}

// SetLibraries gets a reference to the given []PlexLibrary and assigns it to the Libraries field.
func (o *SettingsPlexSyncGet200Response) SetLibraries(v []PlexLibrary) {
	o.Libraries = v
}

func (o SettingsPlexSyncGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SettingsPlexSyncGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Running) {
		toSerialize["running"] = o.Running
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.CurrentLibrary) {
		toSerialize["currentLibrary"] = o.CurrentLibrary
	}
	if !IsNil(o.Libraries) {
		toSerialize["libraries"] = o.Libraries
	}
	return toSerialize, nil
}

type NullableSettingsPlexSyncGet200Response struct {
	value *SettingsPlexSyncGet200Response
	isSet bool
}

func (v NullableSettingsPlexSyncGet200Response) Get() *SettingsPlexSyncGet200Response {
	return v.value
}

func (v *NullableSettingsPlexSyncGet200Response) Set(val *SettingsPlexSyncGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsPlexSyncGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsPlexSyncGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsPlexSyncGet200Response(val *SettingsPlexSyncGet200Response) *NullableSettingsPlexSyncGet200Response {
	return &NullableSettingsPlexSyncGet200Response{value: val, isSet: true}
}

func (v NullableSettingsPlexSyncGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsPlexSyncGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


