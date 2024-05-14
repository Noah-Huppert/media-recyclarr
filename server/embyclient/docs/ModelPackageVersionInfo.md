# ModelPackageVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Guid** | Pointer to **string** |  | [optional] 
**VersionStr** | Pointer to **string** |  | [optional] 
**Classification** | Pointer to [**ModelModelPackageVersionClass**](ModelPackageVersionClass.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequiredVersionStr** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**Checksum** | Pointer to **string** |  | [optional] 
**TargetFilename** | Pointer to **string** |  | [optional] 
**InfoUrl** | Pointer to **string** |  | [optional] 
**Runtimes** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewModelPackageVersionInfo

`func NewModelPackageVersionInfo() *ModelPackageVersionInfo`

NewModelPackageVersionInfo instantiates a new ModelPackageVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelPackageVersionInfoWithDefaults

`func NewModelPackageVersionInfoWithDefaults() *ModelPackageVersionInfo`

NewModelPackageVersionInfoWithDefaults instantiates a new ModelPackageVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelPackageVersionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelPackageVersionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelPackageVersionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelPackageVersionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGuid

`func (o *ModelPackageVersionInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *ModelPackageVersionInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *ModelPackageVersionInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *ModelPackageVersionInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetVersionStr

`func (o *ModelPackageVersionInfo) GetVersionStr() string`

GetVersionStr returns the VersionStr field if non-nil, zero value otherwise.

### GetVersionStrOk

`func (o *ModelPackageVersionInfo) GetVersionStrOk() (*string, bool)`

GetVersionStrOk returns a tuple with the VersionStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionStr

`func (o *ModelPackageVersionInfo) SetVersionStr(v string)`

SetVersionStr sets VersionStr field to given value.

### HasVersionStr

`func (o *ModelPackageVersionInfo) HasVersionStr() bool`

HasVersionStr returns a boolean if a field has been set.

### GetClassification

`func (o *ModelPackageVersionInfo) GetClassification() ModelModelPackageVersionClass`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ModelPackageVersionInfo) GetClassificationOk() (*ModelModelPackageVersionClass, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ModelPackageVersionInfo) SetClassification(v ModelModelPackageVersionClass)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ModelPackageVersionInfo) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetDescription

`func (o *ModelPackageVersionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelPackageVersionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelPackageVersionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelPackageVersionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiredVersionStr

`func (o *ModelPackageVersionInfo) GetRequiredVersionStr() string`

GetRequiredVersionStr returns the RequiredVersionStr field if non-nil, zero value otherwise.

### GetRequiredVersionStrOk

`func (o *ModelPackageVersionInfo) GetRequiredVersionStrOk() (*string, bool)`

GetRequiredVersionStrOk returns a tuple with the RequiredVersionStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredVersionStr

`func (o *ModelPackageVersionInfo) SetRequiredVersionStr(v string)`

SetRequiredVersionStr sets RequiredVersionStr field to given value.

### HasRequiredVersionStr

`func (o *ModelPackageVersionInfo) HasRequiredVersionStr() bool`

HasRequiredVersionStr returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ModelPackageVersionInfo) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ModelPackageVersionInfo) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ModelPackageVersionInfo) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ModelPackageVersionInfo) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetChecksum

`func (o *ModelPackageVersionInfo) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ModelPackageVersionInfo) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ModelPackageVersionInfo) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ModelPackageVersionInfo) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### GetTargetFilename

`func (o *ModelPackageVersionInfo) GetTargetFilename() string`

GetTargetFilename returns the TargetFilename field if non-nil, zero value otherwise.

### GetTargetFilenameOk

`func (o *ModelPackageVersionInfo) GetTargetFilenameOk() (*string, bool)`

GetTargetFilenameOk returns a tuple with the TargetFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilename

`func (o *ModelPackageVersionInfo) SetTargetFilename(v string)`

SetTargetFilename sets TargetFilename field to given value.

### HasTargetFilename

`func (o *ModelPackageVersionInfo) HasTargetFilename() bool`

HasTargetFilename returns a boolean if a field has been set.

### GetInfoUrl

`func (o *ModelPackageVersionInfo) GetInfoUrl() string`

GetInfoUrl returns the InfoUrl field if non-nil, zero value otherwise.

### GetInfoUrlOk

`func (o *ModelPackageVersionInfo) GetInfoUrlOk() (*string, bool)`

GetInfoUrlOk returns a tuple with the InfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoUrl

`func (o *ModelPackageVersionInfo) SetInfoUrl(v string)`

SetInfoUrl sets InfoUrl field to given value.

### HasInfoUrl

`func (o *ModelPackageVersionInfo) HasInfoUrl() bool`

HasInfoUrl returns a boolean if a field has been set.

### GetRuntimes

`func (o *ModelPackageVersionInfo) GetRuntimes() string`

GetRuntimes returns the Runtimes field if non-nil, zero value otherwise.

### GetRuntimesOk

`func (o *ModelPackageVersionInfo) GetRuntimesOk() (*string, bool)`

GetRuntimesOk returns a tuple with the Runtimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimes

`func (o *ModelPackageVersionInfo) SetRuntimes(v string)`

SetRuntimes sets Runtimes field to given value.

### HasRuntimes

`func (o *ModelPackageVersionInfo) HasRuntimes() bool`

HasRuntimes returns a boolean if a field has been set.

### GetTimestamp

`func (o *ModelPackageVersionInfo) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelPackageVersionInfo) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelPackageVersionInfo) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelPackageVersionInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ModelPackageVersionInfo) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ModelPackageVersionInfo) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


