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

// checks if the RequestPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestPostRequest{}

// RequestPostRequest struct for RequestPostRequest
type RequestPostRequest struct {
	MediaType string `json:"mediaType"`
	MediaId float32 `json:"mediaId"`
	TvdbId *float32 `json:"tvdbId,omitempty"`
	Seasons *RequestPostRequestSeasons `json:"seasons,omitempty"`
	Is4k *bool `json:"is4k,omitempty"`
	ServerId *float32 `json:"serverId,omitempty"`
	ProfileId *float32 `json:"profileId,omitempty"`
	RootFolder *string `json:"rootFolder,omitempty"`
	LanguageProfileId *float32 `json:"languageProfileId,omitempty"`
	UserId NullableFloat32 `json:"userId,omitempty"`
}

type _RequestPostRequest RequestPostRequest

// NewRequestPostRequest instantiates a new RequestPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestPostRequest(mediaType string, mediaId float32) *RequestPostRequest {
	this := RequestPostRequest{}
	this.MediaType = mediaType
	this.MediaId = mediaId
	return &this
}

// NewRequestPostRequestWithDefaults instantiates a new RequestPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestPostRequestWithDefaults() *RequestPostRequest {
	this := RequestPostRequest{}
	return &this
}

// GetMediaType returns the MediaType field value
func (o *RequestPostRequest) GetMediaType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaType, true
}

// SetMediaType sets field value
func (o *RequestPostRequest) SetMediaType(v string) {
	o.MediaType = v
}

// GetMediaId returns the MediaId field value
func (o *RequestPostRequest) GetMediaId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetMediaIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaId, true
}

// SetMediaId sets field value
func (o *RequestPostRequest) SetMediaId(v float32) {
	o.MediaId = v
}

// GetTvdbId returns the TvdbId field value if set, zero value otherwise.
func (o *RequestPostRequest) GetTvdbId() float32 {
	if o == nil || IsNil(o.TvdbId) {
		var ret float32
		return ret
	}
	return *o.TvdbId
}

// GetTvdbIdOk returns a tuple with the TvdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetTvdbIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TvdbId) {
		return nil, false
	}
	return o.TvdbId, true
}

// HasTvdbId returns a boolean if a field has been set.
func (o *RequestPostRequest) HasTvdbId() bool {
	if o != nil && !IsNil(o.TvdbId) {
		return true
	}

	return false
}

// SetTvdbId gets a reference to the given float32 and assigns it to the TvdbId field.
func (o *RequestPostRequest) SetTvdbId(v float32) {
	o.TvdbId = &v
}

// GetSeasons returns the Seasons field value if set, zero value otherwise.
func (o *RequestPostRequest) GetSeasons() RequestPostRequestSeasons {
	if o == nil || IsNil(o.Seasons) {
		var ret RequestPostRequestSeasons
		return ret
	}
	return *o.Seasons
}

// GetSeasonsOk returns a tuple with the Seasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetSeasonsOk() (*RequestPostRequestSeasons, bool) {
	if o == nil || IsNil(o.Seasons) {
		return nil, false
	}
	return o.Seasons, true
}

// HasSeasons returns a boolean if a field has been set.
func (o *RequestPostRequest) HasSeasons() bool {
	if o != nil && !IsNil(o.Seasons) {
		return true
	}

	return false
}

// SetSeasons gets a reference to the given RequestPostRequestSeasons and assigns it to the Seasons field.
func (o *RequestPostRequest) SetSeasons(v RequestPostRequestSeasons) {
	o.Seasons = &v
}

// GetIs4k returns the Is4k field value if set, zero value otherwise.
func (o *RequestPostRequest) GetIs4k() bool {
	if o == nil || IsNil(o.Is4k) {
		var ret bool
		return ret
	}
	return *o.Is4k
}

// GetIs4kOk returns a tuple with the Is4k field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetIs4kOk() (*bool, bool) {
	if o == nil || IsNil(o.Is4k) {
		return nil, false
	}
	return o.Is4k, true
}

// HasIs4k returns a boolean if a field has been set.
func (o *RequestPostRequest) HasIs4k() bool {
	if o != nil && !IsNil(o.Is4k) {
		return true
	}

	return false
}

// SetIs4k gets a reference to the given bool and assigns it to the Is4k field.
func (o *RequestPostRequest) SetIs4k(v bool) {
	o.Is4k = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *RequestPostRequest) GetServerId() float32 {
	if o == nil || IsNil(o.ServerId) {
		var ret float32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetServerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *RequestPostRequest) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given float32 and assigns it to the ServerId field.
func (o *RequestPostRequest) SetServerId(v float32) {
	o.ServerId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *RequestPostRequest) GetProfileId() float32 {
	if o == nil || IsNil(o.ProfileId) {
		var ret float32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *RequestPostRequest) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given float32 and assigns it to the ProfileId field.
func (o *RequestPostRequest) SetProfileId(v float32) {
	o.ProfileId = &v
}

// GetRootFolder returns the RootFolder field value if set, zero value otherwise.
func (o *RequestPostRequest) GetRootFolder() string {
	if o == nil || IsNil(o.RootFolder) {
		var ret string
		return ret
	}
	return *o.RootFolder
}

// GetRootFolderOk returns a tuple with the RootFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetRootFolderOk() (*string, bool) {
	if o == nil || IsNil(o.RootFolder) {
		return nil, false
	}
	return o.RootFolder, true
}

// HasRootFolder returns a boolean if a field has been set.
func (o *RequestPostRequest) HasRootFolder() bool {
	if o != nil && !IsNil(o.RootFolder) {
		return true
	}

	return false
}

// SetRootFolder gets a reference to the given string and assigns it to the RootFolder field.
func (o *RequestPostRequest) SetRootFolder(v string) {
	o.RootFolder = &v
}

// GetLanguageProfileId returns the LanguageProfileId field value if set, zero value otherwise.
func (o *RequestPostRequest) GetLanguageProfileId() float32 {
	if o == nil || IsNil(o.LanguageProfileId) {
		var ret float32
		return ret
	}
	return *o.LanguageProfileId
}

// GetLanguageProfileIdOk returns a tuple with the LanguageProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPostRequest) GetLanguageProfileIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LanguageProfileId) {
		return nil, false
	}
	return o.LanguageProfileId, true
}

// HasLanguageProfileId returns a boolean if a field has been set.
func (o *RequestPostRequest) HasLanguageProfileId() bool {
	if o != nil && !IsNil(o.LanguageProfileId) {
		return true
	}

	return false
}

// SetLanguageProfileId gets a reference to the given float32 and assigns it to the LanguageProfileId field.
func (o *RequestPostRequest) SetLanguageProfileId(v float32) {
	o.LanguageProfileId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestPostRequest) GetUserId() float32 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret float32
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestPostRequest) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *RequestPostRequest) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableFloat32 and assigns it to the UserId field.
func (o *RequestPostRequest) SetUserId(v float32) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *RequestPostRequest) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *RequestPostRequest) UnsetUserId() {
	o.UserId.Unset()
}

func (o RequestPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mediaType"] = o.MediaType
	toSerialize["mediaId"] = o.MediaId
	if !IsNil(o.TvdbId) {
		toSerialize["tvdbId"] = o.TvdbId
	}
	if !IsNil(o.Seasons) {
		toSerialize["seasons"] = o.Seasons
	}
	if !IsNil(o.Is4k) {
		toSerialize["is4k"] = o.Is4k
	}
	if !IsNil(o.ServerId) {
		toSerialize["serverId"] = o.ServerId
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.RootFolder) {
		toSerialize["rootFolder"] = o.RootFolder
	}
	if !IsNil(o.LanguageProfileId) {
		toSerialize["languageProfileId"] = o.LanguageProfileId
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	return toSerialize, nil
}

func (o *RequestPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mediaType",
		"mediaId",
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

	varRequestPostRequest := _RequestPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRequestPostRequest)

	if err != nil {
		return err
	}

	*o = RequestPostRequest(varRequestPostRequest)

	return err
}

type NullableRequestPostRequest struct {
	value *RequestPostRequest
	isSet bool
}

func (v NullableRequestPostRequest) Get() *RequestPostRequest {
	return v.value
}

func (v *NullableRequestPostRequest) Set(val *RequestPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPostRequest(val *RequestPostRequest) *NullableRequestPostRequest {
	return &NullableRequestPostRequest{value: val, isSet: true}
}

func (v NullableRequestPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


