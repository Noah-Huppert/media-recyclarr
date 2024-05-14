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

// checks if the ModelImageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelImageInfo{}

// ModelImageInfo struct for ModelImageInfo
type ModelImageInfo struct {
	ImageType *ModelImageType `json:"ImageType,omitempty"`
	ImageIndex NullableInt32 `json:"ImageIndex,omitempty"`
	Path *string `json:"Path,omitempty"`
	Filename *string `json:"Filename,omitempty"`
	Height NullableInt32 `json:"Height,omitempty"`
	Width NullableInt32 `json:"Width,omitempty"`
	Size *int64 `json:"Size,omitempty"`
}

// NewModelImageInfo instantiates a new ModelImageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelImageInfo() *ModelImageInfo {
	this := ModelImageInfo{}
	return &this
}

// NewModelImageInfoWithDefaults instantiates a new ModelImageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelImageInfoWithDefaults() *ModelImageInfo {
	this := ModelImageInfo{}
	return &this
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *ModelImageInfo) GetImageType() ModelImageType {
	if o == nil || IsNil(o.ImageType) {
		var ret ModelImageType
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageInfo) GetImageTypeOk() (*ModelImageType, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *ModelImageInfo) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given ModelImageType and assigns it to the ImageType field.
func (o *ModelImageInfo) SetImageType(v ModelImageType) {
	o.ImageType = &v
}

// GetImageIndex returns the ImageIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelImageInfo) GetImageIndex() int32 {
	if o == nil || IsNil(o.ImageIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.ImageIndex.Get()
}

// GetImageIndexOk returns a tuple with the ImageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelImageInfo) GetImageIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageIndex.Get(), o.ImageIndex.IsSet()
}

// HasImageIndex returns a boolean if a field has been set.
func (o *ModelImageInfo) HasImageIndex() bool {
	if o != nil && o.ImageIndex.IsSet() {
		return true
	}

	return false
}

// SetImageIndex gets a reference to the given NullableInt32 and assigns it to the ImageIndex field.
func (o *ModelImageInfo) SetImageIndex(v int32) {
	o.ImageIndex.Set(&v)
}
// SetImageIndexNil sets the value for ImageIndex to be an explicit nil
func (o *ModelImageInfo) SetImageIndexNil() {
	o.ImageIndex.Set(nil)
}

// UnsetImageIndex ensures that no value is present for ImageIndex, not even an explicit nil
func (o *ModelImageInfo) UnsetImageIndex() {
	o.ImageIndex.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ModelImageInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ModelImageInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ModelImageInfo) SetPath(v string) {
	o.Path = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *ModelImageInfo) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageInfo) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *ModelImageInfo) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *ModelImageInfo) SetFilename(v string) {
	o.Filename = &v
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelImageInfo) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelImageInfo) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *ModelImageInfo) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *ModelImageInfo) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *ModelImageInfo) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *ModelImageInfo) UnsetHeight() {
	o.Height.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelImageInfo) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelImageInfo) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *ModelImageInfo) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *ModelImageInfo) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *ModelImageInfo) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *ModelImageInfo) UnsetWidth() {
	o.Width.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ModelImageInfo) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageInfo) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ModelImageInfo) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ModelImageInfo) SetSize(v int64) {
	o.Size = &v
}

func (o ModelImageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelImageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageType) {
		toSerialize["ImageType"] = o.ImageType
	}
	if o.ImageIndex.IsSet() {
		toSerialize["ImageIndex"] = o.ImageIndex.Get()
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.Filename) {
		toSerialize["Filename"] = o.Filename
	}
	if o.Height.IsSet() {
		toSerialize["Height"] = o.Height.Get()
	}
	if o.Width.IsSet() {
		toSerialize["Width"] = o.Width.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	return toSerialize, nil
}

type NullableModelImageInfo struct {
	value *ModelImageInfo
	isSet bool
}

func (v NullableModelImageInfo) Get() *ModelImageInfo {
	return v.value
}

func (v *NullableModelImageInfo) Set(val *ModelImageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelImageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelImageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelImageInfo(val *ModelImageInfo) *NullableModelImageInfo {
	return &NullableModelImageInfo{value: val, isSet: true}
}

func (v NullableModelImageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelImageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


