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

// checks if the ModelLogFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLogFile{}

// ModelLogFile struct for ModelLogFile
type ModelLogFile struct {
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	DateModified *time.Time `json:"DateModified,omitempty"`
	Size *int64 `json:"Size,omitempty"`
	Name *string `json:"Name,omitempty"`
}

// NewModelLogFile instantiates a new ModelLogFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLogFile() *ModelLogFile {
	this := ModelLogFile{}
	return &this
}

// NewModelLogFileWithDefaults instantiates a new ModelLogFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLogFileWithDefaults() *ModelLogFile {
	this := ModelLogFile{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *ModelLogFile) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLogFile) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *ModelLogFile) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *ModelLogFile) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateModified returns the DateModified field value if set, zero value otherwise.
func (o *ModelLogFile) GetDateModified() time.Time {
	if o == nil || IsNil(o.DateModified) {
		var ret time.Time
		return ret
	}
	return *o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLogFile) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateModified) {
		return nil, false
	}
	return o.DateModified, true
}

// HasDateModified returns a boolean if a field has been set.
func (o *ModelLogFile) HasDateModified() bool {
	if o != nil && !IsNil(o.DateModified) {
		return true
	}

	return false
}

// SetDateModified gets a reference to the given time.Time and assigns it to the DateModified field.
func (o *ModelLogFile) SetDateModified(v time.Time) {
	o.DateModified = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ModelLogFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLogFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ModelLogFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ModelLogFile) SetSize(v int64) {
	o.Size = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelLogFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLogFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelLogFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelLogFile) SetName(v string) {
	o.Name = &v
}

func (o ModelLogFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLogFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if !IsNil(o.DateModified) {
		toSerialize["DateModified"] = o.DateModified
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableModelLogFile struct {
	value *ModelLogFile
	isSet bool
}

func (v NullableModelLogFile) Get() *ModelLogFile {
	return v.value
}

func (v *NullableModelLogFile) Set(val *ModelLogFile) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLogFile) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLogFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLogFile(val *ModelLogFile) *NullableModelLogFile {
	return &NullableModelLogFile{value: val, isSet: true}
}

func (v NullableModelLogFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLogFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


