# ModelSyncDialogOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]ModelModelSyncTarget**](ModelModelSyncTarget.md) |  | [optional] 
**Options** | Pointer to [**[]ModelModelSyncJobOption**](ModelModelSyncJobOption.md) |  | [optional] 
**QualityOptions** | Pointer to [**[]ModelModelSyncQualityOption**](ModelModelSyncQualityOption.md) |  | [optional] 
**ProfileOptions** | Pointer to [**[]ModelModelSyncProfileOption**](ModelModelSyncProfileOption.md) |  | [optional] 

## Methods

### NewModelSyncDialogOptions

`func NewModelSyncDialogOptions() *ModelSyncDialogOptions`

NewModelSyncDialogOptions instantiates a new ModelSyncDialogOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncDialogOptionsWithDefaults

`func NewModelSyncDialogOptionsWithDefaults() *ModelSyncDialogOptions`

NewModelSyncDialogOptionsWithDefaults instantiates a new ModelSyncDialogOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ModelSyncDialogOptions) GetTargets() []ModelModelSyncTarget`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ModelSyncDialogOptions) GetTargetsOk() (*[]ModelModelSyncTarget, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ModelSyncDialogOptions) SetTargets(v []ModelModelSyncTarget)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ModelSyncDialogOptions) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### GetOptions

`func (o *ModelSyncDialogOptions) GetOptions() []ModelModelSyncJobOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ModelSyncDialogOptions) GetOptionsOk() (*[]ModelModelSyncJobOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ModelSyncDialogOptions) SetOptions(v []ModelModelSyncJobOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ModelSyncDialogOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetQualityOptions

`func (o *ModelSyncDialogOptions) GetQualityOptions() []ModelModelSyncQualityOption`

GetQualityOptions returns the QualityOptions field if non-nil, zero value otherwise.

### GetQualityOptionsOk

`func (o *ModelSyncDialogOptions) GetQualityOptionsOk() (*[]ModelModelSyncQualityOption, bool)`

GetQualityOptionsOk returns a tuple with the QualityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOptions

`func (o *ModelSyncDialogOptions) SetQualityOptions(v []ModelModelSyncQualityOption)`

SetQualityOptions sets QualityOptions field to given value.

### HasQualityOptions

`func (o *ModelSyncDialogOptions) HasQualityOptions() bool`

HasQualityOptions returns a boolean if a field has been set.

### GetProfileOptions

`func (o *ModelSyncDialogOptions) GetProfileOptions() []ModelModelSyncProfileOption`

GetProfileOptions returns the ProfileOptions field if non-nil, zero value otherwise.

### GetProfileOptionsOk

`func (o *ModelSyncDialogOptions) GetProfileOptionsOk() (*[]ModelModelSyncProfileOption, bool)`

GetProfileOptionsOk returns a tuple with the ProfileOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOptions

`func (o *ModelSyncDialogOptions) SetProfileOptions(v []ModelModelSyncProfileOption)`

SetProfileOptions sets ProfileOptions field to given value.

### HasProfileOptions

`func (o *ModelSyncDialogOptions) HasProfileOptions() bool`

HasProfileOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


