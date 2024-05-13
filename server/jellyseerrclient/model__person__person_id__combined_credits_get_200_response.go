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

// checks if the PersonPersonIdCombinedCreditsGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonPersonIdCombinedCreditsGet200Response{}

// PersonPersonIdCombinedCreditsGet200Response struct for PersonPersonIdCombinedCreditsGet200Response
type PersonPersonIdCombinedCreditsGet200Response struct {
	Cast []CreditCast `json:"cast,omitempty"`
	Crew []CreditCrew `json:"crew,omitempty"`
	Id *float32 `json:"id,omitempty"`
}

// NewPersonPersonIdCombinedCreditsGet200Response instantiates a new PersonPersonIdCombinedCreditsGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonPersonIdCombinedCreditsGet200Response() *PersonPersonIdCombinedCreditsGet200Response {
	this := PersonPersonIdCombinedCreditsGet200Response{}
	return &this
}

// NewPersonPersonIdCombinedCreditsGet200ResponseWithDefaults instantiates a new PersonPersonIdCombinedCreditsGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonPersonIdCombinedCreditsGet200ResponseWithDefaults() *PersonPersonIdCombinedCreditsGet200Response {
	this := PersonPersonIdCombinedCreditsGet200Response{}
	return &this
}

// GetCast returns the Cast field value if set, zero value otherwise.
func (o *PersonPersonIdCombinedCreditsGet200Response) GetCast() []CreditCast {
	if o == nil || IsNil(o.Cast) {
		var ret []CreditCast
		return ret
	}
	return o.Cast
}

// GetCastOk returns a tuple with the Cast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonPersonIdCombinedCreditsGet200Response) GetCastOk() ([]CreditCast, bool) {
	if o == nil || IsNil(o.Cast) {
		return nil, false
	}
	return o.Cast, true
}

// HasCast returns a boolean if a field has been set.
func (o *PersonPersonIdCombinedCreditsGet200Response) HasCast() bool {
	if o != nil && !IsNil(o.Cast) {
		return true
	}

	return false
}

// SetCast gets a reference to the given []CreditCast and assigns it to the Cast field.
func (o *PersonPersonIdCombinedCreditsGet200Response) SetCast(v []CreditCast) {
	o.Cast = v
}

// GetCrew returns the Crew field value if set, zero value otherwise.
func (o *PersonPersonIdCombinedCreditsGet200Response) GetCrew() []CreditCrew {
	if o == nil || IsNil(o.Crew) {
		var ret []CreditCrew
		return ret
	}
	return o.Crew
}

// GetCrewOk returns a tuple with the Crew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonPersonIdCombinedCreditsGet200Response) GetCrewOk() ([]CreditCrew, bool) {
	if o == nil || IsNil(o.Crew) {
		return nil, false
	}
	return o.Crew, true
}

// HasCrew returns a boolean if a field has been set.
func (o *PersonPersonIdCombinedCreditsGet200Response) HasCrew() bool {
	if o != nil && !IsNil(o.Crew) {
		return true
	}

	return false
}

// SetCrew gets a reference to the given []CreditCrew and assigns it to the Crew field.
func (o *PersonPersonIdCombinedCreditsGet200Response) SetCrew(v []CreditCrew) {
	o.Crew = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PersonPersonIdCombinedCreditsGet200Response) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonPersonIdCombinedCreditsGet200Response) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PersonPersonIdCombinedCreditsGet200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *PersonPersonIdCombinedCreditsGet200Response) SetId(v float32) {
	o.Id = &v
}

func (o PersonPersonIdCombinedCreditsGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonPersonIdCombinedCreditsGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cast) {
		toSerialize["cast"] = o.Cast
	}
	if !IsNil(o.Crew) {
		toSerialize["crew"] = o.Crew
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePersonPersonIdCombinedCreditsGet200Response struct {
	value *PersonPersonIdCombinedCreditsGet200Response
	isSet bool
}

func (v NullablePersonPersonIdCombinedCreditsGet200Response) Get() *PersonPersonIdCombinedCreditsGet200Response {
	return v.value
}

func (v *NullablePersonPersonIdCombinedCreditsGet200Response) Set(val *PersonPersonIdCombinedCreditsGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonPersonIdCombinedCreditsGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonPersonIdCombinedCreditsGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonPersonIdCombinedCreditsGet200Response(val *PersonPersonIdCombinedCreditsGet200Response) *NullablePersonPersonIdCombinedCreditsGet200Response {
	return &NullablePersonPersonIdCombinedCreditsGet200Response{value: val, isSet: true}
}

func (v NullablePersonPersonIdCombinedCreditsGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonPersonIdCombinedCreditsGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


