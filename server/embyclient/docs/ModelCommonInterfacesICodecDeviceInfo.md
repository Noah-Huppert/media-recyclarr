# ModelCommonInterfacesICodecDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**ModelModelCommonInterfacesICodecDeviceCapabilities**](ModelCommonInterfacesICodecDeviceCapabilities.md) |  | [optional] 
**Adapter** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Desription** | Pointer to **string** |  | [optional] 
**Driver** | Pointer to **string** |  | [optional] 
**DriverVersion** | Pointer to [**ModelModelVersion**](ModelVersion.md) |  | [optional] 
**ApiVersion** | Pointer to [**ModelModelVersion**](ModelVersion.md) |  | [optional] 
**VendorId** | Pointer to **int32** |  | [optional] 
**DeviceId** | Pointer to **int32** |  | [optional] 
**DeviceIdentifier** | Pointer to **string** |  | [optional] 
**HardwareContextFramework** | Pointer to [**ModelModelSecondaryFrameworks**](ModelSecondaryFrameworks.md) |  | [optional] 
**DevPath** | Pointer to **string** |  | [optional] 
**DrmNode** | Pointer to **string** |  | [optional] 
**VendorName** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 

## Methods

### NewModelCommonInterfacesICodecDeviceInfo

`func NewModelCommonInterfacesICodecDeviceInfo() *ModelCommonInterfacesICodecDeviceInfo`

NewModelCommonInterfacesICodecDeviceInfo instantiates a new ModelCommonInterfacesICodecDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCommonInterfacesICodecDeviceInfoWithDefaults

`func NewModelCommonInterfacesICodecDeviceInfoWithDefaults() *ModelCommonInterfacesICodecDeviceInfo`

NewModelCommonInterfacesICodecDeviceInfoWithDefaults instantiates a new ModelCommonInterfacesICodecDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetCapabilities() ModelModelCommonInterfacesICodecDeviceCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetCapabilitiesOk() (*ModelModelCommonInterfacesICodecDeviceCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetCapabilities(v ModelModelCommonInterfacesICodecDeviceCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetAdapter

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetAdapter() int32`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetAdapterOk() (*int32, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetAdapter(v int32)`

SetAdapter sets Adapter field to given value.

### HasAdapter

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasAdapter() bool`

HasAdapter returns a boolean if a field has been set.

### GetName

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDesription

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDesription() string`

GetDesription returns the Desription field if non-nil, zero value otherwise.

### GetDesriptionOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDesriptionOk() (*string, bool)`

GetDesriptionOk returns a tuple with the Desription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesription

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDesription(v string)`

SetDesription sets Desription field to given value.

### HasDesription

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDesription() bool`

HasDesription returns a boolean if a field has been set.

### GetDriver

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetDriverVersion

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDriverVersion() ModelModelVersion`

GetDriverVersion returns the DriverVersion field if non-nil, zero value otherwise.

### GetDriverVersionOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDriverVersionOk() (*ModelModelVersion, bool)`

GetDriverVersionOk returns a tuple with the DriverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverVersion

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDriverVersion(v ModelModelVersion)`

SetDriverVersion sets DriverVersion field to given value.

### HasDriverVersion

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDriverVersion() bool`

HasDriverVersion returns a boolean if a field has been set.

### GetApiVersion

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetApiVersion() ModelModelVersion`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetApiVersionOk() (*ModelModelVersion, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetApiVersion(v ModelModelVersion)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetVendorId

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetVendorId() int32`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetVendorIdOk() (*int32, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetVendorId(v int32)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetDeviceId

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceIdentifier

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetHardwareContextFramework

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetHardwareContextFramework() ModelModelSecondaryFrameworks`

GetHardwareContextFramework returns the HardwareContextFramework field if non-nil, zero value otherwise.

### GetHardwareContextFrameworkOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetHardwareContextFrameworkOk() (*ModelModelSecondaryFrameworks, bool)`

GetHardwareContextFrameworkOk returns a tuple with the HardwareContextFramework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareContextFramework

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetHardwareContextFramework(v ModelModelSecondaryFrameworks)`

SetHardwareContextFramework sets HardwareContextFramework field to given value.

### HasHardwareContextFramework

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasHardwareContextFramework() bool`

HasHardwareContextFramework returns a boolean if a field has been set.

### GetDevPath

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDevPath() string`

GetDevPath returns the DevPath field if non-nil, zero value otherwise.

### GetDevPathOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDevPathOk() (*string, bool)`

GetDevPathOk returns a tuple with the DevPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevPath

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDevPath(v string)`

SetDevPath sets DevPath field to given value.

### HasDevPath

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDevPath() bool`

HasDevPath returns a boolean if a field has been set.

### GetDrmNode

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDrmNode() string`

GetDrmNode returns the DrmNode field if non-nil, zero value otherwise.

### GetDrmNodeOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDrmNodeOk() (*string, bool)`

GetDrmNodeOk returns a tuple with the DrmNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrmNode

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDrmNode(v string)`

SetDrmNode sets DrmNode field to given value.

### HasDrmNode

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDrmNode() bool`

HasDrmNode returns a boolean if a field has been set.

### GetVendorName

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetVendorName() string`

GetVendorName returns the VendorName field if non-nil, zero value otherwise.

### GetVendorNameOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetVendorNameOk() (*string, bool)`

GetVendorNameOk returns a tuple with the VendorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorName

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetVendorName(v string)`

SetVendorName sets VendorName field to given value.

### HasVendorName

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasVendorName() bool`

HasVendorName returns a boolean if a field has been set.

### GetDeviceName

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ModelCommonInterfacesICodecDeviceInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ModelCommonInterfacesICodecDeviceInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ModelCommonInterfacesICodecDeviceInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


