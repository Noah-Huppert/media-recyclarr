# ModelUICommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandType** | Pointer to [**ModelModelEnumsUICommandType**](ModelEnumsUICommandType.md) |  | [optional] 
**CommandId** | Pointer to **string** |  | [optional] 
**IsVisible** | Pointer to **bool** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**SetFocus** | Pointer to **bool** |  | [optional] 
**ConfirmationPrompt** | Pointer to **string** |  | [optional] 

## Methods

### NewModelUICommand

`func NewModelUICommand() *ModelUICommand`

NewModelUICommand instantiates a new ModelUICommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUICommandWithDefaults

`func NewModelUICommandWithDefaults() *ModelUICommand`

NewModelUICommandWithDefaults instantiates a new ModelUICommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandType

`func (o *ModelUICommand) GetCommandType() ModelModelEnumsUICommandType`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *ModelUICommand) GetCommandTypeOk() (*ModelModelEnumsUICommandType, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *ModelUICommand) SetCommandType(v ModelModelEnumsUICommandType)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *ModelUICommand) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetCommandId

`func (o *ModelUICommand) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *ModelUICommand) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *ModelUICommand) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *ModelUICommand) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetIsVisible

`func (o *ModelUICommand) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *ModelUICommand) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *ModelUICommand) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *ModelUICommand) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetIsEnabled

`func (o *ModelUICommand) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *ModelUICommand) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *ModelUICommand) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *ModelUICommand) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetCaption

`func (o *ModelUICommand) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *ModelUICommand) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *ModelUICommand) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *ModelUICommand) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetSetFocus

`func (o *ModelUICommand) GetSetFocus() bool`

GetSetFocus returns the SetFocus field if non-nil, zero value otherwise.

### GetSetFocusOk

`func (o *ModelUICommand) GetSetFocusOk() (*bool, bool)`

GetSetFocusOk returns a tuple with the SetFocus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetFocus

`func (o *ModelUICommand) SetSetFocus(v bool)`

SetSetFocus sets SetFocus field to given value.

### HasSetFocus

`func (o *ModelUICommand) HasSetFocus() bool`

HasSetFocus returns a boolean if a field has been set.

### GetConfirmationPrompt

`func (o *ModelUICommand) GetConfirmationPrompt() string`

GetConfirmationPrompt returns the ConfirmationPrompt field if non-nil, zero value otherwise.

### GetConfirmationPromptOk

`func (o *ModelUICommand) GetConfirmationPromptOk() (*string, bool)`

GetConfirmationPromptOk returns a tuple with the ConfirmationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationPrompt

`func (o *ModelUICommand) SetConfirmationPrompt(v string)`

SetConfirmationPrompt sets ConfirmationPrompt field to given value.

### HasConfirmationPrompt

`func (o *ModelUICommand) HasConfirmationPrompt() bool`

HasConfirmationPrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


