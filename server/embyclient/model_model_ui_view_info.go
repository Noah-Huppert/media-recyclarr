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

// checks if the ModelUIViewInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUIViewInfo{}

// ModelUIViewInfo struct for ModelUIViewInfo
type ModelUIViewInfo struct {
	ViewId *string `json:"ViewId,omitempty"`
	PageId *string `json:"PageId,omitempty"`
	Caption *string `json:"Caption,omitempty"`
	SubCaption *string `json:"SubCaption,omitempty"`
	PluginId *string `json:"PluginId,omitempty"`
	ViewType *ModelEnumsUIViewType `json:"ViewType,omitempty"`
	ShowDialogFullScreen *bool `json:"ShowDialogFullScreen,omitempty"`
	IsInSequence *bool `json:"IsInSequence,omitempty"`
	RedirectViewUrl *string `json:"RedirectViewUrl,omitempty"`
	EditObjectContainer *ModelGenericEditIEditObjectContainer `json:"EditObjectContainer,omitempty"`
	Commands []ModelUICommand `json:"Commands,omitempty"`
	TabPageInfos []ModelUITabPageInfo `json:"TabPageInfos,omitempty"`
	IsPageChangeInfo *bool `json:"IsPageChangeInfo,omitempty"`
}

// NewModelUIViewInfo instantiates a new ModelUIViewInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUIViewInfo() *ModelUIViewInfo {
	this := ModelUIViewInfo{}
	return &this
}

// NewModelUIViewInfoWithDefaults instantiates a new ModelUIViewInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUIViewInfoWithDefaults() *ModelUIViewInfo {
	this := ModelUIViewInfo{}
	return &this
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetViewId() string {
	if o == nil || IsNil(o.ViewId) {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.ViewId) {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasViewId() bool {
	if o != nil && !IsNil(o.ViewId) {
		return true
	}

	return false
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *ModelUIViewInfo) SetViewId(v string) {
	o.ViewId = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetPageId() string {
	if o == nil || IsNil(o.PageId) {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetPageIdOk() (*string, bool) {
	if o == nil || IsNil(o.PageId) {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasPageId() bool {
	if o != nil && !IsNil(o.PageId) {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *ModelUIViewInfo) SetPageId(v string) {
	o.PageId = &v
}

// GetCaption returns the Caption field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetCaption() string {
	if o == nil || IsNil(o.Caption) {
		var ret string
		return ret
	}
	return *o.Caption
}

// GetCaptionOk returns a tuple with the Caption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.Caption) {
		return nil, false
	}
	return o.Caption, true
}

// HasCaption returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasCaption() bool {
	if o != nil && !IsNil(o.Caption) {
		return true
	}

	return false
}

// SetCaption gets a reference to the given string and assigns it to the Caption field.
func (o *ModelUIViewInfo) SetCaption(v string) {
	o.Caption = &v
}

// GetSubCaption returns the SubCaption field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetSubCaption() string {
	if o == nil || IsNil(o.SubCaption) {
		var ret string
		return ret
	}
	return *o.SubCaption
}

// GetSubCaptionOk returns a tuple with the SubCaption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetSubCaptionOk() (*string, bool) {
	if o == nil || IsNil(o.SubCaption) {
		return nil, false
	}
	return o.SubCaption, true
}

// HasSubCaption returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasSubCaption() bool {
	if o != nil && !IsNil(o.SubCaption) {
		return true
	}

	return false
}

// SetSubCaption gets a reference to the given string and assigns it to the SubCaption field.
func (o *ModelUIViewInfo) SetSubCaption(v string) {
	o.SubCaption = &v
}

// GetPluginId returns the PluginId field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetPluginId() string {
	if o == nil || IsNil(o.PluginId) {
		var ret string
		return ret
	}
	return *o.PluginId
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetPluginIdOk() (*string, bool) {
	if o == nil || IsNil(o.PluginId) {
		return nil, false
	}
	return o.PluginId, true
}

// HasPluginId returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasPluginId() bool {
	if o != nil && !IsNil(o.PluginId) {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given string and assigns it to the PluginId field.
func (o *ModelUIViewInfo) SetPluginId(v string) {
	o.PluginId = &v
}

// GetViewType returns the ViewType field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetViewType() ModelEnumsUIViewType {
	if o == nil || IsNil(o.ViewType) {
		var ret ModelEnumsUIViewType
		return ret
	}
	return *o.ViewType
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetViewTypeOk() (*ModelEnumsUIViewType, bool) {
	if o == nil || IsNil(o.ViewType) {
		return nil, false
	}
	return o.ViewType, true
}

// HasViewType returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasViewType() bool {
	if o != nil && !IsNil(o.ViewType) {
		return true
	}

	return false
}

// SetViewType gets a reference to the given ModelEnumsUIViewType and assigns it to the ViewType field.
func (o *ModelUIViewInfo) SetViewType(v ModelEnumsUIViewType) {
	o.ViewType = &v
}

// GetShowDialogFullScreen returns the ShowDialogFullScreen field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetShowDialogFullScreen() bool {
	if o == nil || IsNil(o.ShowDialogFullScreen) {
		var ret bool
		return ret
	}
	return *o.ShowDialogFullScreen
}

// GetShowDialogFullScreenOk returns a tuple with the ShowDialogFullScreen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetShowDialogFullScreenOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowDialogFullScreen) {
		return nil, false
	}
	return o.ShowDialogFullScreen, true
}

// HasShowDialogFullScreen returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasShowDialogFullScreen() bool {
	if o != nil && !IsNil(o.ShowDialogFullScreen) {
		return true
	}

	return false
}

// SetShowDialogFullScreen gets a reference to the given bool and assigns it to the ShowDialogFullScreen field.
func (o *ModelUIViewInfo) SetShowDialogFullScreen(v bool) {
	o.ShowDialogFullScreen = &v
}

// GetIsInSequence returns the IsInSequence field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetIsInSequence() bool {
	if o == nil || IsNil(o.IsInSequence) {
		var ret bool
		return ret
	}
	return *o.IsInSequence
}

// GetIsInSequenceOk returns a tuple with the IsInSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetIsInSequenceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInSequence) {
		return nil, false
	}
	return o.IsInSequence, true
}

// HasIsInSequence returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasIsInSequence() bool {
	if o != nil && !IsNil(o.IsInSequence) {
		return true
	}

	return false
}

// SetIsInSequence gets a reference to the given bool and assigns it to the IsInSequence field.
func (o *ModelUIViewInfo) SetIsInSequence(v bool) {
	o.IsInSequence = &v
}

// GetRedirectViewUrl returns the RedirectViewUrl field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetRedirectViewUrl() string {
	if o == nil || IsNil(o.RedirectViewUrl) {
		var ret string
		return ret
	}
	return *o.RedirectViewUrl
}

// GetRedirectViewUrlOk returns a tuple with the RedirectViewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetRedirectViewUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectViewUrl) {
		return nil, false
	}
	return o.RedirectViewUrl, true
}

// HasRedirectViewUrl returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasRedirectViewUrl() bool {
	if o != nil && !IsNil(o.RedirectViewUrl) {
		return true
	}

	return false
}

// SetRedirectViewUrl gets a reference to the given string and assigns it to the RedirectViewUrl field.
func (o *ModelUIViewInfo) SetRedirectViewUrl(v string) {
	o.RedirectViewUrl = &v
}

// GetEditObjectContainer returns the EditObjectContainer field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetEditObjectContainer() ModelGenericEditIEditObjectContainer {
	if o == nil || IsNil(o.EditObjectContainer) {
		var ret ModelGenericEditIEditObjectContainer
		return ret
	}
	return *o.EditObjectContainer
}

// GetEditObjectContainerOk returns a tuple with the EditObjectContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetEditObjectContainerOk() (*ModelGenericEditIEditObjectContainer, bool) {
	if o == nil || IsNil(o.EditObjectContainer) {
		return nil, false
	}
	return o.EditObjectContainer, true
}

// HasEditObjectContainer returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasEditObjectContainer() bool {
	if o != nil && !IsNil(o.EditObjectContainer) {
		return true
	}

	return false
}

// SetEditObjectContainer gets a reference to the given ModelGenericEditIEditObjectContainer and assigns it to the EditObjectContainer field.
func (o *ModelUIViewInfo) SetEditObjectContainer(v ModelGenericEditIEditObjectContainer) {
	o.EditObjectContainer = &v
}

// GetCommands returns the Commands field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetCommands() []ModelUICommand {
	if o == nil || IsNil(o.Commands) {
		var ret []ModelUICommand
		return ret
	}
	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetCommandsOk() ([]ModelUICommand, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// HasCommands returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasCommands() bool {
	if o != nil && !IsNil(o.Commands) {
		return true
	}

	return false
}

// SetCommands gets a reference to the given []ModelUICommand and assigns it to the Commands field.
func (o *ModelUIViewInfo) SetCommands(v []ModelUICommand) {
	o.Commands = v
}

// GetTabPageInfos returns the TabPageInfos field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetTabPageInfos() []ModelUITabPageInfo {
	if o == nil || IsNil(o.TabPageInfos) {
		var ret []ModelUITabPageInfo
		return ret
	}
	return o.TabPageInfos
}

// GetTabPageInfosOk returns a tuple with the TabPageInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetTabPageInfosOk() ([]ModelUITabPageInfo, bool) {
	if o == nil || IsNil(o.TabPageInfos) {
		return nil, false
	}
	return o.TabPageInfos, true
}

// HasTabPageInfos returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasTabPageInfos() bool {
	if o != nil && !IsNil(o.TabPageInfos) {
		return true
	}

	return false
}

// SetTabPageInfos gets a reference to the given []ModelUITabPageInfo and assigns it to the TabPageInfos field.
func (o *ModelUIViewInfo) SetTabPageInfos(v []ModelUITabPageInfo) {
	o.TabPageInfos = v
}

// GetIsPageChangeInfo returns the IsPageChangeInfo field value if set, zero value otherwise.
func (o *ModelUIViewInfo) GetIsPageChangeInfo() bool {
	if o == nil || IsNil(o.IsPageChangeInfo) {
		var ret bool
		return ret
	}
	return *o.IsPageChangeInfo
}

// GetIsPageChangeInfoOk returns a tuple with the IsPageChangeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUIViewInfo) GetIsPageChangeInfoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPageChangeInfo) {
		return nil, false
	}
	return o.IsPageChangeInfo, true
}

// HasIsPageChangeInfo returns a boolean if a field has been set.
func (o *ModelUIViewInfo) HasIsPageChangeInfo() bool {
	if o != nil && !IsNil(o.IsPageChangeInfo) {
		return true
	}

	return false
}

// SetIsPageChangeInfo gets a reference to the given bool and assigns it to the IsPageChangeInfo field.
func (o *ModelUIViewInfo) SetIsPageChangeInfo(v bool) {
	o.IsPageChangeInfo = &v
}

func (o ModelUIViewInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUIViewInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViewId) {
		toSerialize["ViewId"] = o.ViewId
	}
	if !IsNil(o.PageId) {
		toSerialize["PageId"] = o.PageId
	}
	if !IsNil(o.Caption) {
		toSerialize["Caption"] = o.Caption
	}
	if !IsNil(o.SubCaption) {
		toSerialize["SubCaption"] = o.SubCaption
	}
	if !IsNil(o.PluginId) {
		toSerialize["PluginId"] = o.PluginId
	}
	if !IsNil(o.ViewType) {
		toSerialize["ViewType"] = o.ViewType
	}
	if !IsNil(o.ShowDialogFullScreen) {
		toSerialize["ShowDialogFullScreen"] = o.ShowDialogFullScreen
	}
	if !IsNil(o.IsInSequence) {
		toSerialize["IsInSequence"] = o.IsInSequence
	}
	if !IsNil(o.RedirectViewUrl) {
		toSerialize["RedirectViewUrl"] = o.RedirectViewUrl
	}
	if !IsNil(o.EditObjectContainer) {
		toSerialize["EditObjectContainer"] = o.EditObjectContainer
	}
	if !IsNil(o.Commands) {
		toSerialize["Commands"] = o.Commands
	}
	if !IsNil(o.TabPageInfos) {
		toSerialize["TabPageInfos"] = o.TabPageInfos
	}
	if !IsNil(o.IsPageChangeInfo) {
		toSerialize["IsPageChangeInfo"] = o.IsPageChangeInfo
	}
	return toSerialize, nil
}

type NullableModelUIViewInfo struct {
	value *ModelUIViewInfo
	isSet bool
}

func (v NullableModelUIViewInfo) Get() *ModelUIViewInfo {
	return v.value
}

func (v *NullableModelUIViewInfo) Set(val *ModelUIViewInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUIViewInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUIViewInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUIViewInfo(val *ModelUIViewInfo) *NullableModelUIViewInfo {
	return &NullableModelUIViewInfo{value: val, isSet: true}
}

func (v NullableModelUIViewInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUIViewInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


