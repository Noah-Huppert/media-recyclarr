# ModelEditorsEditorRoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyConditions** | Pointer to [**[]ModelModelConditionsPropertyCondition**](ModelModelConditionsPropertyCondition.md) |  | [optional] 
**PostbackActions** | Pointer to [**[]ModelModelActionsPostbackAction**](ModelModelActionsPostbackAction.md) |  | [optional] 
**TitleButton** | Pointer to [**ModelModelEditorsEditorButtonItem**](ModelEditorsEditorButtonItem.md) |  | [optional] 
**EditorItems** | Pointer to [**[]ModelModelEditorsEditorBase**](ModelModelEditorsEditorBase.md) |  | [optional] 
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

### NewModelEditorsEditorRoot

`func NewModelEditorsEditorRoot() *ModelEditorsEditorRoot`

NewModelEditorsEditorRoot instantiates a new ModelEditorsEditorRoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelEditorsEditorRootWithDefaults

`func NewModelEditorsEditorRootWithDefaults() *ModelEditorsEditorRoot`

NewModelEditorsEditorRootWithDefaults instantiates a new ModelEditorsEditorRoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyConditions

`func (o *ModelEditorsEditorRoot) GetPropertyConditions() []ModelModelConditionsPropertyCondition`

GetPropertyConditions returns the PropertyConditions field if non-nil, zero value otherwise.

### GetPropertyConditionsOk

`func (o *ModelEditorsEditorRoot) GetPropertyConditionsOk() (*[]ModelModelConditionsPropertyCondition, bool)`

GetPropertyConditionsOk returns a tuple with the PropertyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyConditions

`func (o *ModelEditorsEditorRoot) SetPropertyConditions(v []ModelModelConditionsPropertyCondition)`

SetPropertyConditions sets PropertyConditions field to given value.

### HasPropertyConditions

`func (o *ModelEditorsEditorRoot) HasPropertyConditions() bool`

HasPropertyConditions returns a boolean if a field has been set.

### GetPostbackActions

`func (o *ModelEditorsEditorRoot) GetPostbackActions() []ModelModelActionsPostbackAction`

GetPostbackActions returns the PostbackActions field if non-nil, zero value otherwise.

### GetPostbackActionsOk

`func (o *ModelEditorsEditorRoot) GetPostbackActionsOk() (*[]ModelModelActionsPostbackAction, bool)`

GetPostbackActionsOk returns a tuple with the PostbackActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackActions

`func (o *ModelEditorsEditorRoot) SetPostbackActions(v []ModelModelActionsPostbackAction)`

SetPostbackActions sets PostbackActions field to given value.

### HasPostbackActions

`func (o *ModelEditorsEditorRoot) HasPostbackActions() bool`

HasPostbackActions returns a boolean if a field has been set.

### GetTitleButton

`func (o *ModelEditorsEditorRoot) GetTitleButton() ModelModelEditorsEditorButtonItem`

GetTitleButton returns the TitleButton field if non-nil, zero value otherwise.

### GetTitleButtonOk

`func (o *ModelEditorsEditorRoot) GetTitleButtonOk() (*ModelModelEditorsEditorButtonItem, bool)`

GetTitleButtonOk returns a tuple with the TitleButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleButton

`func (o *ModelEditorsEditorRoot) SetTitleButton(v ModelModelEditorsEditorButtonItem)`

SetTitleButton sets TitleButton field to given value.

### HasTitleButton

`func (o *ModelEditorsEditorRoot) HasTitleButton() bool`

HasTitleButton returns a boolean if a field has been set.

### GetEditorItems

`func (o *ModelEditorsEditorRoot) GetEditorItems() []ModelModelEditorsEditorBase`

GetEditorItems returns the EditorItems field if non-nil, zero value otherwise.

### GetEditorItemsOk

`func (o *ModelEditorsEditorRoot) GetEditorItemsOk() (*[]ModelModelEditorsEditorBase, bool)`

GetEditorItemsOk returns a tuple with the EditorItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorItems

`func (o *ModelEditorsEditorRoot) SetEditorItems(v []ModelModelEditorsEditorBase)`

SetEditorItems sets EditorItems field to given value.

### HasEditorItems

`func (o *ModelEditorsEditorRoot) HasEditorItems() bool`

HasEditorItems returns a boolean if a field has been set.

### GetEditorType

`func (o *ModelEditorsEditorRoot) GetEditorType() ModelModelCommonEditorTypes`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ModelEditorsEditorRoot) GetEditorTypeOk() (*ModelModelCommonEditorTypes, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ModelEditorsEditorRoot) SetEditorType(v ModelModelCommonEditorTypes)`

SetEditorType sets EditorType field to given value.

### HasEditorType

`func (o *ModelEditorsEditorRoot) HasEditorType() bool`

HasEditorType returns a boolean if a field has been set.

### GetName

`func (o *ModelEditorsEditorRoot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelEditorsEditorRoot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelEditorsEditorRoot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelEditorsEditorRoot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelEditorsEditorRoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelEditorsEditorRoot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelEditorsEditorRoot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelEditorsEditorRoot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAllowEmpty

`func (o *ModelEditorsEditorRoot) GetAllowEmpty() bool`

GetAllowEmpty returns the AllowEmpty field if non-nil, zero value otherwise.

### GetAllowEmptyOk

`func (o *ModelEditorsEditorRoot) GetAllowEmptyOk() (*bool, bool)`

GetAllowEmptyOk returns a tuple with the AllowEmpty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmpty

`func (o *ModelEditorsEditorRoot) SetAllowEmpty(v bool)`

SetAllowEmpty sets AllowEmpty field to given value.

### HasAllowEmpty

`func (o *ModelEditorsEditorRoot) HasAllowEmpty() bool`

HasAllowEmpty returns a boolean if a field has been set.

### GetIsReadOnly

`func (o *ModelEditorsEditorRoot) GetIsReadOnly() bool`

GetIsReadOnly returns the IsReadOnly field if non-nil, zero value otherwise.

### GetIsReadOnlyOk

`func (o *ModelEditorsEditorRoot) GetIsReadOnlyOk() (*bool, bool)`

GetIsReadOnlyOk returns a tuple with the IsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadOnly

`func (o *ModelEditorsEditorRoot) SetIsReadOnly(v bool)`

SetIsReadOnly sets IsReadOnly field to given value.

### HasIsReadOnly

`func (o *ModelEditorsEditorRoot) HasIsReadOnly() bool`

HasIsReadOnly returns a boolean if a field has been set.

### GetIsAdvanced

`func (o *ModelEditorsEditorRoot) GetIsAdvanced() bool`

GetIsAdvanced returns the IsAdvanced field if non-nil, zero value otherwise.

### GetIsAdvancedOk

`func (o *ModelEditorsEditorRoot) GetIsAdvancedOk() (*bool, bool)`

GetIsAdvancedOk returns a tuple with the IsAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdvanced

`func (o *ModelEditorsEditorRoot) SetIsAdvanced(v bool)`

SetIsAdvanced sets IsAdvanced field to given value.

### HasIsAdvanced

`func (o *ModelEditorsEditorRoot) HasIsAdvanced() bool`

HasIsAdvanced returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelEditorsEditorRoot) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelEditorsEditorRoot) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelEditorsEditorRoot) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelEditorsEditorRoot) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ModelEditorsEditorRoot) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelEditorsEditorRoot) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelEditorsEditorRoot) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelEditorsEditorRoot) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *ModelEditorsEditorRoot) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModelEditorsEditorRoot) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModelEditorsEditorRoot) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModelEditorsEditorRoot) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


