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

// checks if the ModelPackageVersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPackageVersionInfo{}

// ModelPackageVersionInfo struct for ModelPackageVersionInfo
type ModelPackageVersionInfo struct {
	Name *string `json:"name,omitempty"`
	Guid *string `json:"guid,omitempty"`
	VersionStr *string `json:"versionStr,omitempty"`
	Classification *ModelPackageVersionClass `json:"classification,omitempty"`
	Description *string `json:"description,omitempty"`
	RequiredVersionStr *string `json:"requiredVersionStr,omitempty"`
	SourceUrl *string `json:"sourceUrl,omitempty"`
	Checksum *string `json:"checksum,omitempty"`
	TargetFilename *string `json:"targetFilename,omitempty"`
	InfoUrl *string `json:"infoUrl,omitempty"`
	Runtimes *string `json:"runtimes,omitempty"`
	Timestamp NullableTime `json:"timestamp,omitempty"`
}

// NewModelPackageVersionInfo instantiates a new ModelPackageVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPackageVersionInfo() *ModelPackageVersionInfo {
	this := ModelPackageVersionInfo{}
	return &this
}

// NewModelPackageVersionInfoWithDefaults instantiates a new ModelPackageVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPackageVersionInfoWithDefaults() *ModelPackageVersionInfo {
	this := ModelPackageVersionInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelPackageVersionInfo) SetName(v string) {
	o.Name = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ModelPackageVersionInfo) SetGuid(v string) {
	o.Guid = &v
}

// GetVersionStr returns the VersionStr field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetVersionStr() string {
	if o == nil || IsNil(o.VersionStr) {
		var ret string
		return ret
	}
	return *o.VersionStr
}

// GetVersionStrOk returns a tuple with the VersionStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetVersionStrOk() (*string, bool) {
	if o == nil || IsNil(o.VersionStr) {
		return nil, false
	}
	return o.VersionStr, true
}

// HasVersionStr returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasVersionStr() bool {
	if o != nil && !IsNil(o.VersionStr) {
		return true
	}

	return false
}

// SetVersionStr gets a reference to the given string and assigns it to the VersionStr field.
func (o *ModelPackageVersionInfo) SetVersionStr(v string) {
	o.VersionStr = &v
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetClassification() ModelPackageVersionClass {
	if o == nil || IsNil(o.Classification) {
		var ret ModelPackageVersionClass
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetClassificationOk() (*ModelPackageVersionClass, bool) {
	if o == nil || IsNil(o.Classification) {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasClassification() bool {
	if o != nil && !IsNil(o.Classification) {
		return true
	}

	return false
}

// SetClassification gets a reference to the given ModelPackageVersionClass and assigns it to the Classification field.
func (o *ModelPackageVersionInfo) SetClassification(v ModelPackageVersionClass) {
	o.Classification = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelPackageVersionInfo) SetDescription(v string) {
	o.Description = &v
}

// GetRequiredVersionStr returns the RequiredVersionStr field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetRequiredVersionStr() string {
	if o == nil || IsNil(o.RequiredVersionStr) {
		var ret string
		return ret
	}
	return *o.RequiredVersionStr
}

// GetRequiredVersionStrOk returns a tuple with the RequiredVersionStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetRequiredVersionStrOk() (*string, bool) {
	if o == nil || IsNil(o.RequiredVersionStr) {
		return nil, false
	}
	return o.RequiredVersionStr, true
}

// HasRequiredVersionStr returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasRequiredVersionStr() bool {
	if o != nil && !IsNil(o.RequiredVersionStr) {
		return true
	}

	return false
}

// SetRequiredVersionStr gets a reference to the given string and assigns it to the RequiredVersionStr field.
func (o *ModelPackageVersionInfo) SetRequiredVersionStr(v string) {
	o.RequiredVersionStr = &v
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetSourceUrl() string {
	if o == nil || IsNil(o.SourceUrl) {
		var ret string
		return ret
	}
	return *o.SourceUrl
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetSourceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SourceUrl) {
		return nil, false
	}
	return o.SourceUrl, true
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasSourceUrl() bool {
	if o != nil && !IsNil(o.SourceUrl) {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given string and assigns it to the SourceUrl field.
func (o *ModelPackageVersionInfo) SetSourceUrl(v string) {
	o.SourceUrl = &v
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *ModelPackageVersionInfo) SetChecksum(v string) {
	o.Checksum = &v
}

// GetTargetFilename returns the TargetFilename field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetTargetFilename() string {
	if o == nil || IsNil(o.TargetFilename) {
		var ret string
		return ret
	}
	return *o.TargetFilename
}

// GetTargetFilenameOk returns a tuple with the TargetFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetTargetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetFilename) {
		return nil, false
	}
	return o.TargetFilename, true
}

// HasTargetFilename returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasTargetFilename() bool {
	if o != nil && !IsNil(o.TargetFilename) {
		return true
	}

	return false
}

// SetTargetFilename gets a reference to the given string and assigns it to the TargetFilename field.
func (o *ModelPackageVersionInfo) SetTargetFilename(v string) {
	o.TargetFilename = &v
}

// GetInfoUrl returns the InfoUrl field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetInfoUrl() string {
	if o == nil || IsNil(o.InfoUrl) {
		var ret string
		return ret
	}
	return *o.InfoUrl
}

// GetInfoUrlOk returns a tuple with the InfoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetInfoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InfoUrl) {
		return nil, false
	}
	return o.InfoUrl, true
}

// HasInfoUrl returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasInfoUrl() bool {
	if o != nil && !IsNil(o.InfoUrl) {
		return true
	}

	return false
}

// SetInfoUrl gets a reference to the given string and assigns it to the InfoUrl field.
func (o *ModelPackageVersionInfo) SetInfoUrl(v string) {
	o.InfoUrl = &v
}

// GetRuntimes returns the Runtimes field value if set, zero value otherwise.
func (o *ModelPackageVersionInfo) GetRuntimes() string {
	if o == nil || IsNil(o.Runtimes) {
		var ret string
		return ret
	}
	return *o.Runtimes
}

// GetRuntimesOk returns a tuple with the Runtimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageVersionInfo) GetRuntimesOk() (*string, bool) {
	if o == nil || IsNil(o.Runtimes) {
		return nil, false
	}
	return o.Runtimes, true
}

// HasRuntimes returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasRuntimes() bool {
	if o != nil && !IsNil(o.Runtimes) {
		return true
	}

	return false
}

// SetRuntimes gets a reference to the given string and assigns it to the Runtimes field.
func (o *ModelPackageVersionInfo) SetRuntimes(v string) {
	o.Runtimes = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPackageVersionInfo) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPackageVersionInfo) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ModelPackageVersionInfo) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableTime and assigns it to the Timestamp field.
func (o *ModelPackageVersionInfo) SetTimestamp(v time.Time) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *ModelPackageVersionInfo) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *ModelPackageVersionInfo) UnsetTimestamp() {
	o.Timestamp.Unset()
}

func (o ModelPackageVersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPackageVersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.VersionStr) {
		toSerialize["versionStr"] = o.VersionStr
	}
	if !IsNil(o.Classification) {
		toSerialize["classification"] = o.Classification
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RequiredVersionStr) {
		toSerialize["requiredVersionStr"] = o.RequiredVersionStr
	}
	if !IsNil(o.SourceUrl) {
		toSerialize["sourceUrl"] = o.SourceUrl
	}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.TargetFilename) {
		toSerialize["targetFilename"] = o.TargetFilename
	}
	if !IsNil(o.InfoUrl) {
		toSerialize["infoUrl"] = o.InfoUrl
	}
	if !IsNil(o.Runtimes) {
		toSerialize["runtimes"] = o.Runtimes
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	return toSerialize, nil
}

type NullableModelPackageVersionInfo struct {
	value *ModelPackageVersionInfo
	isSet bool
}

func (v NullableModelPackageVersionInfo) Get() *ModelPackageVersionInfo {
	return v.value
}

func (v *NullableModelPackageVersionInfo) Set(val *ModelPackageVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPackageVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPackageVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPackageVersionInfo(val *ModelPackageVersionInfo) *NullableModelPackageVersionInfo {
	return &NullableModelPackageVersionInfo{value: val, isSet: true}
}

func (v NullableModelPackageVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPackageVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

