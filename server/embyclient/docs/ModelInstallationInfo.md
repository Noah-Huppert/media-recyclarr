# ModelInstallationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AssemblyGuid** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**UpdateClass** | Pointer to [**ModelModelPackageVersionClass**](ModelPackageVersionClass.md) |  | [optional] 
**PercentComplete** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewModelInstallationInfo

`func NewModelInstallationInfo() *ModelInstallationInfo`

NewModelInstallationInfo instantiates a new ModelInstallationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInstallationInfoWithDefaults

`func NewModelInstallationInfoWithDefaults() *ModelInstallationInfo`

NewModelInstallationInfoWithDefaults instantiates a new ModelInstallationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelInstallationInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelInstallationInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelInstallationInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelInstallationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelInstallationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelInstallationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelInstallationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelInstallationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssemblyGuid

`func (o *ModelInstallationInfo) GetAssemblyGuid() string`

GetAssemblyGuid returns the AssemblyGuid field if non-nil, zero value otherwise.

### GetAssemblyGuidOk

`func (o *ModelInstallationInfo) GetAssemblyGuidOk() (*string, bool)`

GetAssemblyGuidOk returns a tuple with the AssemblyGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyGuid

`func (o *ModelInstallationInfo) SetAssemblyGuid(v string)`

SetAssemblyGuid sets AssemblyGuid field to given value.

### HasAssemblyGuid

`func (o *ModelInstallationInfo) HasAssemblyGuid() bool`

HasAssemblyGuid returns a boolean if a field has been set.

### GetVersion

`func (o *ModelInstallationInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelInstallationInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelInstallationInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelInstallationInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUpdateClass

`func (o *ModelInstallationInfo) GetUpdateClass() ModelModelPackageVersionClass`

GetUpdateClass returns the UpdateClass field if non-nil, zero value otherwise.

### GetUpdateClassOk

`func (o *ModelInstallationInfo) GetUpdateClassOk() (*ModelModelPackageVersionClass, bool)`

GetUpdateClassOk returns a tuple with the UpdateClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateClass

`func (o *ModelInstallationInfo) SetUpdateClass(v ModelModelPackageVersionClass)`

SetUpdateClass sets UpdateClass field to given value.

### HasUpdateClass

`func (o *ModelInstallationInfo) HasUpdateClass() bool`

HasUpdateClass returns a boolean if a field has been set.

### GetPercentComplete

`func (o *ModelInstallationInfo) GetPercentComplete() float64`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *ModelInstallationInfo) GetPercentCompleteOk() (*float64, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *ModelInstallationInfo) SetPercentComplete(v float64)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *ModelInstallationInfo) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *ModelInstallationInfo) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *ModelInstallationInfo) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


