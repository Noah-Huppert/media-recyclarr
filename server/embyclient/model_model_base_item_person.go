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

// checks if the ModelBaseItemPerson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBaseItemPerson{}

// ModelBaseItemPerson struct for ModelBaseItemPerson
type ModelBaseItemPerson struct {
	Name *string `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
	Role *string `json:"Role,omitempty"`
	Type *ModelPersonType `json:"Type,omitempty"`
	PrimaryImageTag *string `json:"PrimaryImageTag,omitempty"`
}

// NewModelBaseItemPerson instantiates a new ModelBaseItemPerson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBaseItemPerson() *ModelBaseItemPerson {
	this := ModelBaseItemPerson{}
	return &this
}

// NewModelBaseItemPersonWithDefaults instantiates a new ModelBaseItemPerson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBaseItemPersonWithDefaults() *ModelBaseItemPerson {
	this := ModelBaseItemPerson{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelBaseItemPerson) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBaseItemPerson) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelBaseItemPerson) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelBaseItemPerson) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelBaseItemPerson) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBaseItemPerson) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelBaseItemPerson) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelBaseItemPerson) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ModelBaseItemPerson) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBaseItemPerson) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ModelBaseItemPerson) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ModelBaseItemPerson) SetRole(v string) {
	o.Role = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelBaseItemPerson) GetType() ModelPersonType {
	if o == nil || IsNil(o.Type) {
		var ret ModelPersonType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBaseItemPerson) GetTypeOk() (*ModelPersonType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelBaseItemPerson) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ModelPersonType and assigns it to the Type field.
func (o *ModelBaseItemPerson) SetType(v ModelPersonType) {
	o.Type = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise.
func (o *ModelBaseItemPerson) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelBaseItemPerson) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryImageTag) {
		return nil, false
	}
	return o.PrimaryImageTag, true
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *ModelBaseItemPerson) HasPrimaryImageTag() bool {
	if o != nil && !IsNil(o.PrimaryImageTag) {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given string and assigns it to the PrimaryImageTag field.
func (o *ModelBaseItemPerson) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag = &v
}

func (o ModelBaseItemPerson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBaseItemPerson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Role) {
		toSerialize["Role"] = o.Role
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.PrimaryImageTag) {
		toSerialize["PrimaryImageTag"] = o.PrimaryImageTag
	}
	return toSerialize, nil
}

type NullableModelBaseItemPerson struct {
	value *ModelBaseItemPerson
	isSet bool
}

func (v NullableModelBaseItemPerson) Get() *ModelBaseItemPerson {
	return v.value
}

func (v *NullableModelBaseItemPerson) Set(val *ModelBaseItemPerson) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBaseItemPerson) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBaseItemPerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBaseItemPerson(val *ModelBaseItemPerson) *NullableModelBaseItemPerson {
	return &NullableModelBaseItemPerson{value: val, isSet: true}
}

func (v NullableModelBaseItemPerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBaseItemPerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

