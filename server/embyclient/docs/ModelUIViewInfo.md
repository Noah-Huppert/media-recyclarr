# ModelUIViewInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewId** | Pointer to **string** |  | [optional] 
**PageId** | Pointer to **string** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**SubCaption** | Pointer to **string** |  | [optional] 
**PluginId** | Pointer to **string** |  | [optional] 
**ViewType** | Pointer to [**ModelModelEnumsUIViewType**](ModelEnumsUIViewType.md) |  | [optional] 
**ShowDialogFullScreen** | Pointer to **bool** |  | [optional] 
**IsInSequence** | Pointer to **bool** |  | [optional] 
**RedirectViewUrl** | Pointer to **string** |  | [optional] 
**EditObjectContainer** | Pointer to [**ModelModelGenericEditIEditObjectContainer**](ModelGenericEditIEditObjectContainer.md) |  | [optional] 
**Commands** | Pointer to [**[]ModelModelUICommand**](ModelModelUICommand.md) |  | [optional] 
**TabPageInfos** | Pointer to [**[]ModelModelUITabPageInfo**](ModelModelUITabPageInfo.md) |  | [optional] 
**IsPageChangeInfo** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelUIViewInfo

`func NewModelUIViewInfo() *ModelUIViewInfo`

NewModelUIViewInfo instantiates a new ModelUIViewInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUIViewInfoWithDefaults

`func NewModelUIViewInfoWithDefaults() *ModelUIViewInfo`

NewModelUIViewInfoWithDefaults instantiates a new ModelUIViewInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewId

`func (o *ModelUIViewInfo) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *ModelUIViewInfo) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *ModelUIViewInfo) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *ModelUIViewInfo) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetPageId

`func (o *ModelUIViewInfo) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *ModelUIViewInfo) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *ModelUIViewInfo) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *ModelUIViewInfo) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetCaption

`func (o *ModelUIViewInfo) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *ModelUIViewInfo) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *ModelUIViewInfo) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *ModelUIViewInfo) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetSubCaption

`func (o *ModelUIViewInfo) GetSubCaption() string`

GetSubCaption returns the SubCaption field if non-nil, zero value otherwise.

### GetSubCaptionOk

`func (o *ModelUIViewInfo) GetSubCaptionOk() (*string, bool)`

GetSubCaptionOk returns a tuple with the SubCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCaption

`func (o *ModelUIViewInfo) SetSubCaption(v string)`

SetSubCaption sets SubCaption field to given value.

### HasSubCaption

`func (o *ModelUIViewInfo) HasSubCaption() bool`

HasSubCaption returns a boolean if a field has been set.

### GetPluginId

`func (o *ModelUIViewInfo) GetPluginId() string`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *ModelUIViewInfo) GetPluginIdOk() (*string, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *ModelUIViewInfo) SetPluginId(v string)`

SetPluginId sets PluginId field to given value.

### HasPluginId

`func (o *ModelUIViewInfo) HasPluginId() bool`

HasPluginId returns a boolean if a field has been set.

### GetViewType

`func (o *ModelUIViewInfo) GetViewType() ModelModelEnumsUIViewType`

GetViewType returns the ViewType field if non-nil, zero value otherwise.

### GetViewTypeOk

`func (o *ModelUIViewInfo) GetViewTypeOk() (*ModelModelEnumsUIViewType, bool)`

GetViewTypeOk returns a tuple with the ViewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewType

`func (o *ModelUIViewInfo) SetViewType(v ModelModelEnumsUIViewType)`

SetViewType sets ViewType field to given value.

### HasViewType

`func (o *ModelUIViewInfo) HasViewType() bool`

HasViewType returns a boolean if a field has been set.

### GetShowDialogFullScreen

`func (o *ModelUIViewInfo) GetShowDialogFullScreen() bool`

GetShowDialogFullScreen returns the ShowDialogFullScreen field if non-nil, zero value otherwise.

### GetShowDialogFullScreenOk

`func (o *ModelUIViewInfo) GetShowDialogFullScreenOk() (*bool, bool)`

GetShowDialogFullScreenOk returns a tuple with the ShowDialogFullScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDialogFullScreen

`func (o *ModelUIViewInfo) SetShowDialogFullScreen(v bool)`

SetShowDialogFullScreen sets ShowDialogFullScreen field to given value.

### HasShowDialogFullScreen

`func (o *ModelUIViewInfo) HasShowDialogFullScreen() bool`

HasShowDialogFullScreen returns a boolean if a field has been set.

### GetIsInSequence

`func (o *ModelUIViewInfo) GetIsInSequence() bool`

GetIsInSequence returns the IsInSequence field if non-nil, zero value otherwise.

### GetIsInSequenceOk

`func (o *ModelUIViewInfo) GetIsInSequenceOk() (*bool, bool)`

GetIsInSequenceOk returns a tuple with the IsInSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSequence

`func (o *ModelUIViewInfo) SetIsInSequence(v bool)`

SetIsInSequence sets IsInSequence field to given value.

### HasIsInSequence

`func (o *ModelUIViewInfo) HasIsInSequence() bool`

HasIsInSequence returns a boolean if a field has been set.

### GetRedirectViewUrl

`func (o *ModelUIViewInfo) GetRedirectViewUrl() string`

GetRedirectViewUrl returns the RedirectViewUrl field if non-nil, zero value otherwise.

### GetRedirectViewUrlOk

`func (o *ModelUIViewInfo) GetRedirectViewUrlOk() (*string, bool)`

GetRedirectViewUrlOk returns a tuple with the RedirectViewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectViewUrl

`func (o *ModelUIViewInfo) SetRedirectViewUrl(v string)`

SetRedirectViewUrl sets RedirectViewUrl field to given value.

### HasRedirectViewUrl

`func (o *ModelUIViewInfo) HasRedirectViewUrl() bool`

HasRedirectViewUrl returns a boolean if a field has been set.

### GetEditObjectContainer

`func (o *ModelUIViewInfo) GetEditObjectContainer() ModelModelGenericEditIEditObjectContainer`

GetEditObjectContainer returns the EditObjectContainer field if non-nil, zero value otherwise.

### GetEditObjectContainerOk

`func (o *ModelUIViewInfo) GetEditObjectContainerOk() (*ModelModelGenericEditIEditObjectContainer, bool)`

GetEditObjectContainerOk returns a tuple with the EditObjectContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditObjectContainer

`func (o *ModelUIViewInfo) SetEditObjectContainer(v ModelModelGenericEditIEditObjectContainer)`

SetEditObjectContainer sets EditObjectContainer field to given value.

### HasEditObjectContainer

`func (o *ModelUIViewInfo) HasEditObjectContainer() bool`

HasEditObjectContainer returns a boolean if a field has been set.

### GetCommands

`func (o *ModelUIViewInfo) GetCommands() []ModelModelUICommand`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ModelUIViewInfo) GetCommandsOk() (*[]ModelModelUICommand, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ModelUIViewInfo) SetCommands(v []ModelModelUICommand)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ModelUIViewInfo) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetTabPageInfos

`func (o *ModelUIViewInfo) GetTabPageInfos() []ModelModelUITabPageInfo`

GetTabPageInfos returns the TabPageInfos field if non-nil, zero value otherwise.

### GetTabPageInfosOk

`func (o *ModelUIViewInfo) GetTabPageInfosOk() (*[]ModelModelUITabPageInfo, bool)`

GetTabPageInfosOk returns a tuple with the TabPageInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabPageInfos

`func (o *ModelUIViewInfo) SetTabPageInfos(v []ModelModelUITabPageInfo)`

SetTabPageInfos sets TabPageInfos field to given value.

### HasTabPageInfos

`func (o *ModelUIViewInfo) HasTabPageInfos() bool`

HasTabPageInfos returns a boolean if a field has been set.

### GetIsPageChangeInfo

`func (o *ModelUIViewInfo) GetIsPageChangeInfo() bool`

GetIsPageChangeInfo returns the IsPageChangeInfo field if non-nil, zero value otherwise.

### GetIsPageChangeInfoOk

`func (o *ModelUIViewInfo) GetIsPageChangeInfoOk() (*bool, bool)`

GetIsPageChangeInfoOk returns a tuple with the IsPageChangeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPageChangeInfo

`func (o *ModelUIViewInfo) SetIsPageChangeInfo(v bool)`

SetIsPageChangeInfo sets IsPageChangeInfo field to given value.

### HasIsPageChangeInfo

`func (o *ModelUIViewInfo) HasIsPageChangeInfo() bool`

HasIsPageChangeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


