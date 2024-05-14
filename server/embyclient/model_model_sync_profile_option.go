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

// checks if the ModelSyncProfileOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSyncProfileOption{}

// ModelSyncProfileOption struct for ModelSyncProfileOption
type ModelSyncProfileOption struct {
	Name *string `json:"Name,omitempty"`
	Description *string `json:"Description,omitempty"`
	Id *string `json:"Id,omitempty"`
	IsDefault *bool `json:"IsDefault,omitempty"`
	EnableQualityOptions *bool `json:"EnableQualityOptions,omitempty"`
}

// NewModelSyncProfileOption instantiates a new ModelSyncProfileOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSyncProfileOption() *ModelSyncProfileOption {
	this := ModelSyncProfileOption{}
	return &this
}

// NewModelSyncProfileOptionWithDefaults instantiates a new ModelSyncProfileOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSyncProfileOptionWithDefaults() *ModelSyncProfileOption {
	this := ModelSyncProfileOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelSyncProfileOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncProfileOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelSyncProfileOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelSyncProfileOption) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelSyncProfileOption) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncProfileOption) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelSyncProfileOption) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelSyncProfileOption) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelSyncProfileOption) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncProfileOption) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelSyncProfileOption) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelSyncProfileOption) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ModelSyncProfileOption) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncProfileOption) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ModelSyncProfileOption) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ModelSyncProfileOption) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetEnableQualityOptions returns the EnableQualityOptions field value if set, zero value otherwise.
func (o *ModelSyncProfileOption) GetEnableQualityOptions() bool {
	if o == nil || IsNil(o.EnableQualityOptions) {
		var ret bool
		return ret
	}
	return *o.EnableQualityOptions
}

// GetEnableQualityOptionsOk returns a tuple with the EnableQualityOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSyncProfileOption) GetEnableQualityOptionsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableQualityOptions) {
		return nil, false
	}
	return o.EnableQualityOptions, true
}

// HasEnableQualityOptions returns a boolean if a field has been set.
func (o *ModelSyncProfileOption) HasEnableQualityOptions() bool {
	if o != nil && !IsNil(o.EnableQualityOptions) {
		return true
	}

	return false
}

// SetEnableQualityOptions gets a reference to the given bool and assigns it to the EnableQualityOptions field.
func (o *ModelSyncProfileOption) SetEnableQualityOptions(v bool) {
	o.EnableQualityOptions = &v
}

func (o ModelSyncProfileOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSyncProfileOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.IsDefault) {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if !IsNil(o.EnableQualityOptions) {
		toSerialize["EnableQualityOptions"] = o.EnableQualityOptions
	}
	return toSerialize, nil
}

type NullableModelSyncProfileOption struct {
	value *ModelSyncProfileOption
	isSet bool
}

func (v NullableModelSyncProfileOption) Get() *ModelSyncProfileOption {
	return v.value
}

func (v *NullableModelSyncProfileOption) Set(val *ModelSyncProfileOption) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSyncProfileOption) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSyncProfileOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSyncProfileOption(val *ModelSyncProfileOption) *NullableModelSyncProfileOption {
	return &NullableModelSyncProfileOption{value: val, isSet: true}
}

func (v NullableModelSyncProfileOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSyncProfileOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


