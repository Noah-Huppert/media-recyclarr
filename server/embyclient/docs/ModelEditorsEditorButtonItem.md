# ModelEditorsEditorButtonItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditorType** | Pointer to [**ModelModelCommonEditorTypes**](ModelCommonEditorTypes.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**AllowEmpty** | Pointer to **bool** |  | [optional] 
**IsReadOnly** | Pointer to **bool** |  | [optional] 
**IsAdvanced** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelEditorsEditorButtonItem

`func NewModelEditorsEditorButtonItem() *ModelEditorsEditorButtonItem`

NewModelEditorsEditorButtonItem instantiates a new ModelEditorsEditorButtonItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEditorsEditorButtonItemWithDefaults

`func NewModelEditorsEditorButtonItemWithDefaults() *ModelEditorsEditorButtonItem`

NewModelEditorsEditorButtonItemWithDefaults instantiates a new ModelEditorsEditorButtonItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditorType

`func (o *ModelEditorsEditorButtonItem) GetEditorType() ModelModelCommonEditorTypes`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ModelEditorsEditorButtonItem) GetEditorTypeOk() (*ModelModelCommonEditorTypes, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ModelEditorsEditorButtonItem) SetEditorType(v ModelModelCommonEditorTypes)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *ModelEditorsEditorButtonItem) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetName

`func (o *ModelEditorsEditorButtonItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelEditorsEditorButtonItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelEditorsEditorButtonItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelEditorsEditorButtonItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelEditorsEditorButtonItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelEditorsEditorButtonItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelEditorsEditorButtonItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelEditorsEditorButtonItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *ModelEditorsEditorButtonItem) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *ModelEditorsEditorButtonItem) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *ModelEditorsEditorButtonItem) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *ModelEditorsEditorButtonItem) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *ModelEditorsEditorButtonItem) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *ModelEditorsEditorButtonItem) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *ModelEditorsEditorButtonItem) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *ModelEditorsEditorButtonItem) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsAdvanced

`func (o *ModelEditorsEditorButtonItem) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *ModelEditorsEditorButtonItem) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *ModelEditorsEditorButtonItem) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.

### HasIsAdvanced

`func (o *ModelEditorsEditorButtonItem) HasIsAdvanced() bool`

HasIsAdvanced returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelEditorsEditorButtonItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelEditorsEditorButtonItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelEditorsEditorButtonItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelEditorsEditorButtonItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ModelEditorsEditorButtonItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelEditorsEditorButtonItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelEditorsEditorButtonItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelEditorsEditorButtonItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *ModelEditorsEditorButtonItem) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModelEditorsEditorButtonItem) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModelEditorsEditorButtonItem) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModelEditorsEditorButtonItem) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

