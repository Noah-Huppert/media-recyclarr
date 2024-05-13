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

// checks if the DiscoverMoviesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoverMoviesGet200Response{}

// DiscoverMoviesGet200Response struct for DiscoverMoviesGet200Response
type DiscoverMoviesGet200Response struct {
	Page *float32 `json:"page,omitempty"`
	TotalPages *float32 `json:"totalPages,omitempty"`
	TotalResults *float32 `json:"totalResults,omitempty"`
	Results []MovieResult `json:"results,omitempty"`
}

// NewDiscoverMoviesGet200Response instantiates a new DiscoverMoviesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoverMoviesGet200Response() *DiscoverMoviesGet200Response {
	this := DiscoverMoviesGet200Response{}
	return &this
}

// NewDiscoverMoviesGet200ResponseWithDefaults instantiates a new DiscoverMoviesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoverMoviesGet200ResponseWithDefaults() *DiscoverMoviesGet200Response {
	this := DiscoverMoviesGet200Response{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DiscoverMoviesGet200Response) GetPage() float32 {
	if o == nil || IsNil(o.Page) {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverMoviesGet200Response) GetPageOk() (*float32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DiscoverMoviesGet200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *DiscoverMoviesGet200Response) SetPage(v float32) {
	o.Page = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *DiscoverMoviesGet200Response) GetTotalPages() float32 {
	if o == nil || IsNil(o.TotalPages) {
		var ret float32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverMoviesGet200Response) GetTotalPagesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalPages) {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *DiscoverMoviesGet200Response) HasTotalPages() bool {
	if o != nil && !IsNil(o.TotalPages) {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given float32 and assigns it to the TotalPages field.
func (o *DiscoverMoviesGet200Response) SetTotalPages(v float32) {
	o.TotalPages = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *DiscoverMoviesGet200Response) GetTotalResults() float32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret float32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverMoviesGet200Response) GetTotalResultsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *DiscoverMoviesGet200Response) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given float32 and assigns it to the TotalResults field.
func (o *DiscoverMoviesGet200Response) SetTotalResults(v float32) {
	o.TotalResults = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *DiscoverMoviesGet200Response) GetResults() []MovieResult {
	if o == nil || IsNil(o.Results) {
		var ret []MovieResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoverMoviesGet200Response) GetResultsOk() ([]MovieResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *DiscoverMoviesGet200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []MovieResult and assigns it to the Results field.
func (o *DiscoverMoviesGet200Response) SetResults(v []MovieResult) {
	o.Results = v
}

func (o DiscoverMoviesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoverMoviesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.TotalPages) {
		toSerialize["totalPages"] = o.TotalPages
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableDiscoverMoviesGet200Response struct {
	value *DiscoverMoviesGet200Response
	isSet bool
}

func (v NullableDiscoverMoviesGet200Response) Get() *DiscoverMoviesGet200Response {
	return v.value
}

func (v *NullableDiscoverMoviesGet200Response) Set(val *DiscoverMoviesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoverMoviesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoverMoviesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoverMoviesGet200Response(val *DiscoverMoviesGet200Response) *NullableDiscoverMoviesGet200Response {
	return &NullableDiscoverMoviesGet200Response{value: val, isSet: true}
}

func (v NullableDiscoverMoviesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoverMoviesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


