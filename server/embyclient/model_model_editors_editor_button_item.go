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

// checks if the ModelEditorsEditorButtonItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelEditorsEditorButtonItem{}

// ModelEditorsEditorButtonItem struct for ModelEditorsEditorButtonItem
type ModelEditorsEditorButtonItem struct {
	EditorType *ModelCommonEditorTypes `json:"EditorType,omitempty"`
	Name *string `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
	AllowEmpty *bool `json:"AllowEmpty,omitempty"`
	IsReadOnly *bool `json:"IsReadOnly,omitempty"`
	IsAdvanced *bool `json:"IsAdvanced,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty"`
	Description *string `json:"Description,omitempty"`
	ParentId *string `json:"ParentId,omitempty"`
}

// NewModelEditorsEditorButtonItem instantiates a new ModelEditorsEditorButtonItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelEditorsEditorButtonItem() *ModelEditorsEditorButtonItem {
	this := ModelEditorsEditorButtonItem{}
	return &this
}

// NewModelEditorsEditorButtonItemWithDefaults instantiates a new ModelEditorsEditorButtonItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelEditorsEditorButtonItemWithDefaults() *ModelEditorsEditorButtonItem {
	this := ModelEditorsEditorButtonItem{}
	return &this
}

// GetEditorType returns the EditorType field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetEditorType() ModelCommonEditorTypes {
	if o == nil || IsNil(o.EditorType) {
		var ret ModelCommonEditorTypes
		return ret
	}
	return *o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetEditorTypeOk() (*ModelCommonEditorTypes, bool) {
	if o == nil || IsNil(o.EditorType) {
		return nil, false
	}
	return o.EditorType, true
}

// HasEditorType returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasEditorType() bool {
	if o != nil && !IsNil(o.EditorType) {
		return true
	}

	return false
}

// SetEditorType gets a reference to the given ModelCommonEditorTypes and assigns it to the EditorType field.
func (o *ModelEditorsEditorButtonItem) SetEditorType(v ModelCommonEditorTypes) {
	o.EditorType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelEditorsEditorButtonItem) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelEditorsEditorButtonItem) SetId(v string) {
	o.Id = &v
}

// GetAllowEmpty returns the AllowEmpty field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetAllowEmpty() bool {
	if o == nil || IsNil(o.AllowEmpty) {
		var ret bool
		return ret
	}
	return *o.AllowEmpty
}

// GetAllowEmptyOk returns a tuple with the AllowEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetAllowEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowEmpty) {
		return nil, false
	}
	return o.AllowEmpty, true
}

// HasAllowEmpty returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasAllowEmpty() bool {
	if o != nil && !IsNil(o.AllowEmpty) {
		return true
	}

	return false
}

// SetAllowEmpty gets a reference to the given bool and assigns it to the AllowEmpty field.
func (o *ModelEditorsEditorButtonItem) SetAllowEmpty(v bool) {
	o.AllowEmpty = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetIsReadOnly() bool {
	if o == nil || IsNil(o.IsReadOnly) {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReadOnly) {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasIsReadOnly() bool {
	if o != nil && !IsNil(o.IsReadOnly) {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *ModelEditorsEditorButtonItem) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetIsAdvanced returns the IsAdvanced field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetIsAdvanced() bool {
	if o == nil || IsNil(o.IsAdvanced) {
		var ret bool
		return ret
	}
	return *o.IsAdvanced
}

// GetIsAdvancedOk returns a tuple with the IsAdvanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetIsAdvancedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAdvanced) {
		return nil, false
	}
	return o.IsAdvanced, true
}

// HasIsAdvanced returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasIsAdvanced() bool {
	if o != nil && !IsNil(o.IsAdvanced) {
		return true
	}

	return false
}

// SetIsAdvanced gets a reference to the given bool and assigns it to the IsAdvanced field.
func (o *ModelEditorsEditorButtonItem) SetIsAdvanced(v bool) {
	o.IsAdvanced = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ModelEditorsEditorButtonItem) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelEditorsEditorButtonItem) SetDescription(v string) {
	o.Description = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ModelEditorsEditorButtonItem) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEditorsEditorButtonItem) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ModelEditorsEditorButtonItem) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ModelEditorsEditorButtonItem) SetParentId(v string) {
	o.ParentId = &v
}

func (o ModelEditorsEditorButtonItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelEditorsEditorButtonItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EditorType) {
		toSerialize["EditorType"] = o.EditorType
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.AllowEmpty) {
		toSerialize["AllowEmpty"] = o.AllowEmpty
	}
	if !IsNil(o.IsReadOnly) {
		toSerialize["IsReadOnly"] = o.IsReadOnly
	}
	if !IsNil(o.IsAdvanced) {
		toSerialize["IsAdvanced"] = o.IsAdvanced
	}
	if !IsNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.ParentId) {
		toSerialize["ParentId"] = o.ParentId
	}
	return toSerialize, nil
}

type NullableModelEditorsEditorButtonItem struct {
	value *ModelEditorsEditorButtonItem
	isSet bool
}

func (v NullableModelEditorsEditorButtonItem) Get() *ModelEditorsEditorButtonItem {
	return v.value
}

func (v *NullableModelEditorsEditorButtonItem) Set(val *ModelEditorsEditorButtonItem) {
	v.value = val
	v.isSet = true
}

func (v NullableModelEditorsEditorButtonItem) IsSet() bool {
	return v.isSet
}

func (v *NullableModelEditorsEditorButtonItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelEditorsEditorButtonItem(val *ModelEditorsEditorButtonItem) *NullableModelEditorsEditorButtonItem {
	return &NullableModelEditorsEditorButtonItem{value: val, isSet: true}
}

func (v NullableModelEditorsEditorButtonItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelEditorsEditorButtonItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

