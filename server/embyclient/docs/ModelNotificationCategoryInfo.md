# ModelNotificationCategoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**[]ModelModelNotificationTypeInfo**](ModelModelNotificationTypeInfo.md) |  | [optional] 

## Methods

### NewModelNotificationCategoryInfo

`func NewModelNotificationCategoryInfo() *ModelNotificationCategoryInfo`

NewModelNotificationCategoryInfo instantiates a new ModelNotificationCategoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelNotificationCategoryInfoWithDefaults

`func NewModelNotificationCategoryInfoWithDefaults() *ModelNotificationCategoryInfo`

NewModelNotificationCategoryInfoWithDefaults instantiates a new ModelNotificationCategoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelNotificationCategoryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelNotificationCategoryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelNotificationCategoryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelNotificationCategoryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelNotificationCategoryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelNotificationCategoryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelNotificationCategoryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelNotificationCategoryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvents

`func (o *ModelNotificationCategoryInfo) GetEvents() []ModelModelNotificationTypeInfo`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ModelNotificationCategoryInfo) GetEventsOk() (*[]ModelModelNotificationTypeInfo, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ModelNotificationCategoryInfo) SetEvents(v []ModelModelNotificationTypeInfo)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ModelNotificationCategoryInfo) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


