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

// checks if the ModelPackageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPackageInfo{}

// ModelPackageInfo struct for ModelPackageInfo
type ModelPackageInfo struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ShortDescription *string `json:"shortDescription,omitempty"`
	Overview *string `json:"overview,omitempty"`
	IsPremium *bool `json:"isPremium,omitempty"`
	Adult *bool `json:"adult,omitempty"`
	RichDescUrl *string `json:"richDescUrl,omitempty"`
	ThumbImage *string `json:"thumbImage,omitempty"`
	PreviewImage *string `json:"previewImage,omitempty"`
	Type *string `json:"type,omitempty"`
	TargetFilename *string `json:"targetFilename,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Category *string `json:"category,omitempty"`
	TileColor *string `json:"tileColor,omitempty"`
	FeatureId *string `json:"featureId,omitempty"`
	Price NullableFloat32 `json:"price,omitempty"`
	TargetSystem *ModelPackageTargetSystem `json:"targetSystem,omitempty"`
	Guid *string `json:"guid,omitempty"`
	IsRegistered *bool `json:"isRegistered,omitempty"`
	ExpDate *time.Time `json:"expDate,omitempty"`
	Versions []ModelPackageVersionInfo `json:"versions,omitempty"`
	EnableInAppStore *bool `json:"enableInAppStore,omitempty"`
	Installs *int32 `json:"installs,omitempty"`
}

// NewModelPackageInfo instantiates a new ModelPackageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPackageInfo() *ModelPackageInfo {
	this := ModelPackageInfo{}
	return &this
}

// NewModelPackageInfoWithDefaults instantiates a new ModelPackageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPackageInfoWithDefaults() *ModelPackageInfo {
	this := ModelPackageInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelPackageInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelPackageInfo) SetName(v string) {
	o.Name = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ModelPackageInfo) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *ModelPackageInfo) SetOverview(v string) {
	o.Overview = &v
}

// GetIsPremium returns the IsPremium field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetIsPremium() bool {
	if o == nil || IsNil(o.IsPremium) {
		var ret bool
		return ret
	}
	return *o.IsPremium
}

// GetIsPremiumOk returns a tuple with the IsPremium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetIsPremiumOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPremium) {
		return nil, false
	}
	return o.IsPremium, true
}

// HasIsPremium returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasIsPremium() bool {
	if o != nil && !IsNil(o.IsPremium) {
		return true
	}

	return false
}

// SetIsPremium gets a reference to the given bool and assigns it to the IsPremium field.
func (o *ModelPackageInfo) SetIsPremium(v bool) {
	o.IsPremium = &v
}

// GetAdult returns the Adult field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetAdult() bool {
	if o == nil || IsNil(o.Adult) {
		var ret bool
		return ret
	}
	return *o.Adult
}

// GetAdultOk returns a tuple with the Adult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetAdultOk() (*bool, bool) {
	if o == nil || IsNil(o.Adult) {
		return nil, false
	}
	return o.Adult, true
}

// HasAdult returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasAdult() bool {
	if o != nil && !IsNil(o.Adult) {
		return true
	}

	return false
}

// SetAdult gets a reference to the given bool and assigns it to the Adult field.
func (o *ModelPackageInfo) SetAdult(v bool) {
	o.Adult = &v
}

// GetRichDescUrl returns the RichDescUrl field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetRichDescUrl() string {
	if o == nil || IsNil(o.RichDescUrl) {
		var ret string
		return ret
	}
	return *o.RichDescUrl
}

// GetRichDescUrlOk returns a tuple with the RichDescUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetRichDescUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RichDescUrl) {
		return nil, false
	}
	return o.RichDescUrl, true
}

// HasRichDescUrl returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasRichDescUrl() bool {
	if o != nil && !IsNil(o.RichDescUrl) {
		return true
	}

	return false
}

// SetRichDescUrl gets a reference to the given string and assigns it to the RichDescUrl field.
func (o *ModelPackageInfo) SetRichDescUrl(v string) {
	o.RichDescUrl = &v
}

// GetThumbImage returns the ThumbImage field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetThumbImage() string {
	if o == nil || IsNil(o.ThumbImage) {
		var ret string
		return ret
	}
	return *o.ThumbImage
}

// GetThumbImageOk returns a tuple with the ThumbImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetThumbImageOk() (*string, bool) {
	if o == nil || IsNil(o.ThumbImage) {
		return nil, false
	}
	return o.ThumbImage, true
}

// HasThumbImage returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasThumbImage() bool {
	if o != nil && !IsNil(o.ThumbImage) {
		return true
	}

	return false
}

// SetThumbImage gets a reference to the given string and assigns it to the ThumbImage field.
func (o *ModelPackageInfo) SetThumbImage(v string) {
	o.ThumbImage = &v
}

// GetPreviewImage returns the PreviewImage field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetPreviewImage() string {
	if o == nil || IsNil(o.PreviewImage) {
		var ret string
		return ret
	}
	return *o.PreviewImage
}

// GetPreviewImageOk returns a tuple with the PreviewImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetPreviewImageOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewImage) {
		return nil, false
	}
	return o.PreviewImage, true
}

// HasPreviewImage returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasPreviewImage() bool {
	if o != nil && !IsNil(o.PreviewImage) {
		return true
	}

	return false
}

// SetPreviewImage gets a reference to the given string and assigns it to the PreviewImage field.
func (o *ModelPackageInfo) SetPreviewImage(v string) {
	o.PreviewImage = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelPackageInfo) SetType(v string) {
	o.Type = &v
}

// GetTargetFilename returns the TargetFilename field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetTargetFilename() string {
	if o == nil || IsNil(o.TargetFilename) {
		var ret string
		return ret
	}
	return *o.TargetFilename
}

// GetTargetFilenameOk returns a tuple with the TargetFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetTargetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetFilename) {
		return nil, false
	}
	return o.TargetFilename, true
}

// HasTargetFilename returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasTargetFilename() bool {
	if o != nil && !IsNil(o.TargetFilename) {
		return true
	}

	return false
}

// SetTargetFilename gets a reference to the given string and assigns it to the TargetFilename field.
func (o *ModelPackageInfo) SetTargetFilename(v string) {
	o.TargetFilename = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ModelPackageInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ModelPackageInfo) SetCategory(v string) {
	o.Category = &v
}

// GetTileColor returns the TileColor field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetTileColor() string {
	if o == nil || IsNil(o.TileColor) {
		var ret string
		return ret
	}
	return *o.TileColor
}

// GetTileColorOk returns a tuple with the TileColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetTileColorOk() (*string, bool) {
	if o == nil || IsNil(o.TileColor) {
		return nil, false
	}
	return o.TileColor, true
}

// HasTileColor returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasTileColor() bool {
	if o != nil && !IsNil(o.TileColor) {
		return true
	}

	return false
}

// SetTileColor gets a reference to the given string and assigns it to the TileColor field.
func (o *ModelPackageInfo) SetTileColor(v string) {
	o.TileColor = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetFeatureId() string {
	if o == nil || IsNil(o.FeatureId) {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetFeatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *ModelPackageInfo) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelPackageInfo) GetPrice() float32 {
	if o == nil || IsNil(o.Price.Get()) {
		var ret float32
		return ret
	}
	return *o.Price.Get()
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPackageInfo) GetPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Price.Get(), o.Price.IsSet()
}

// HasPrice returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasPrice() bool {
	if o != nil && o.Price.IsSet() {
		return true
	}

	return false
}

// SetPrice gets a reference to the given NullableFloat32 and assigns it to the Price field.
func (o *ModelPackageInfo) SetPrice(v float32) {
	o.Price.Set(&v)
}
// SetPriceNil sets the value for Price to be an explicit nil
func (o *ModelPackageInfo) SetPriceNil() {
	o.Price.Set(nil)
}

// UnsetPrice ensures that no value is present for Price, not even an explicit nil
func (o *ModelPackageInfo) UnsetPrice() {
	o.Price.Unset()
}

// GetTargetSystem returns the TargetSystem field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetTargetSystem() ModelPackageTargetSystem {
	if o == nil || IsNil(o.TargetSystem) {
		var ret ModelPackageTargetSystem
		return ret
	}
	return *o.TargetSystem
}

// GetTargetSystemOk returns a tuple with the TargetSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetTargetSystemOk() (*ModelPackageTargetSystem, bool) {
	if o == nil || IsNil(o.TargetSystem) {
		return nil, false
	}
	return o.TargetSystem, true
}

// HasTargetSystem returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasTargetSystem() bool {
	if o != nil && !IsNil(o.TargetSystem) {
		return true
	}

	return false
}

// SetTargetSystem gets a reference to the given ModelPackageTargetSystem and assigns it to the TargetSystem field.
func (o *ModelPackageInfo) SetTargetSystem(v ModelPackageTargetSystem) {
	o.TargetSystem = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *ModelPackageInfo) SetGuid(v string) {
	o.Guid = &v
}

// GetIsRegistered returns the IsRegistered field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetIsRegistered() bool {
	if o == nil || IsNil(o.IsRegistered) {
		var ret bool
		return ret
	}
	return *o.IsRegistered
}

// GetIsRegisteredOk returns a tuple with the IsRegistered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetIsRegisteredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRegistered) {
		return nil, false
	}
	return o.IsRegistered, true
}

// HasIsRegistered returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasIsRegistered() bool {
	if o != nil && !IsNil(o.IsRegistered) {
		return true
	}

	return false
}

// SetIsRegistered gets a reference to the given bool and assigns it to the IsRegistered field.
func (o *ModelPackageInfo) SetIsRegistered(v bool) {
	o.IsRegistered = &v
}

// GetExpDate returns the ExpDate field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetExpDate() time.Time {
	if o == nil || IsNil(o.ExpDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpDate
}

// GetExpDateOk returns a tuple with the ExpDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetExpDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpDate) {
		return nil, false
	}
	return o.ExpDate, true
}

// HasExpDate returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasExpDate() bool {
	if o != nil && !IsNil(o.ExpDate) {
		return true
	}

	return false
}

// SetExpDate gets a reference to the given time.Time and assigns it to the ExpDate field.
func (o *ModelPackageInfo) SetExpDate(v time.Time) {
	o.ExpDate = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetVersions() []ModelPackageVersionInfo {
	if o == nil || IsNil(o.Versions) {
		var ret []ModelPackageVersionInfo
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetVersionsOk() ([]ModelPackageVersionInfo, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []ModelPackageVersionInfo and assigns it to the Versions field.
func (o *ModelPackageInfo) SetVersions(v []ModelPackageVersionInfo) {
	o.Versions = v
}

// GetEnableInAppStore returns the EnableInAppStore field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetEnableInAppStore() bool {
	if o == nil || IsNil(o.EnableInAppStore) {
		var ret bool
		return ret
	}
	return *o.EnableInAppStore
}

// GetEnableInAppStoreOk returns a tuple with the EnableInAppStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetEnableInAppStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableInAppStore) {
		return nil, false
	}
	return o.EnableInAppStore, true
}

// HasEnableInAppStore returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasEnableInAppStore() bool {
	if o != nil && !IsNil(o.EnableInAppStore) {
		return true
	}

	return false
}

// SetEnableInAppStore gets a reference to the given bool and assigns it to the EnableInAppStore field.
func (o *ModelPackageInfo) SetEnableInAppStore(v bool) {
	o.EnableInAppStore = &v
}

// GetInstalls returns the Installs field value if set, zero value otherwise.
func (o *ModelPackageInfo) GetInstalls() int32 {
	if o == nil || IsNil(o.Installs) {
		var ret int32
		return ret
	}
	return *o.Installs
}

// GetInstallsOk returns a tuple with the Installs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelPackageInfo) GetInstallsOk() (*int32, bool) {
	if o == nil || IsNil(o.Installs) {
		return nil, false
	}
	return o.Installs, true
}

// HasInstalls returns a boolean if a field has been set.
func (o *ModelPackageInfo) HasInstalls() bool {
	if o != nil && !IsNil(o.Installs) {
		return true
	}

	return false
}

// SetInstalls gets a reference to the given int32 and assigns it to the Installs field.
func (o *ModelPackageInfo) SetInstalls(v int32) {
	o.Installs = &v
}

func (o ModelPackageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPackageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.IsPremium) {
		toSerialize["isPremium"] = o.IsPremium
	}
	if !IsNil(o.Adult) {
		toSerialize["adult"] = o.Adult
	}
	if !IsNil(o.RichDescUrl) {
		toSerialize["richDescUrl"] = o.RichDescUrl
	}
	if !IsNil(o.ThumbImage) {
		toSerialize["thumbImage"] = o.ThumbImage
	}
	if !IsNil(o.PreviewImage) {
		toSerialize["previewImage"] = o.PreviewImage
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TargetFilename) {
		toSerialize["targetFilename"] = o.TargetFilename
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.TileColor) {
		toSerialize["tileColor"] = o.TileColor
	}
	if !IsNil(o.FeatureId) {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.Price.IsSet() {
		toSerialize["price"] = o.Price.Get()
	}
	if !IsNil(o.TargetSystem) {
		toSerialize["targetSystem"] = o.TargetSystem
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.IsRegistered) {
		toSerialize["isRegistered"] = o.IsRegistered
	}
	if !IsNil(o.ExpDate) {
		toSerialize["expDate"] = o.ExpDate
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.EnableInAppStore) {
		toSerialize["enableInAppStore"] = o.EnableInAppStore
	}
	if !IsNil(o.Installs) {
		toSerialize["installs"] = o.Installs
	}
	return toSerialize, nil
}

type NullableModelPackageInfo struct {
	value *ModelPackageInfo
	isSet bool
}

func (v NullableModelPackageInfo) Get() *ModelPackageInfo {
	return v.value
}

func (v *NullableModelPackageInfo) Set(val *ModelPackageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPackageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPackageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPackageInfo(val *ModelPackageInfo) *NullableModelPackageInfo {
	return &NullableModelPackageInfo{value: val, isSet: true}
}

func (v NullableModelPackageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPackageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

