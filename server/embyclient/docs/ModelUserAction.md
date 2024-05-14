# ModelUserAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ModelModelUserActionType**](ModelUserActionType.md) |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**PositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**Played** | Pointer to **NullableBool** |  | [optional] 
**IsFavorite** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewModelUserAction

`func NewModelUserAction() *ModelUserAction`

NewModelUserAction instantiates a new ModelUserAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUserActionWithDefaults

`func NewModelUserActionWithDefaults() *ModelUserAction`

NewModelUserActionWithDefaults instantiates a new ModelUserAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelUserAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelUserAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelUserAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelUserAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerId

`func (o *ModelUserAction) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelUserAction) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelUserAction) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelUserAction) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetUserId

`func (o *ModelUserAction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelUserAction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelUserAction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelUserAction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetItemId

`func (o *ModelUserAction) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelUserAction) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelUserAction) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelUserAction) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetType

`func (o *ModelUserAction) GetType() ModelModelUserActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelUserAction) GetTypeOk() (*ModelModelUserActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelUserAction) SetType(v ModelModelUserActionType)`

SetType sets Type field to given value.

### HasType

`func (o *ModelUserAction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDate

`func (o *ModelUserAction) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ModelUserAction) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ModelUserAction) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ModelUserAction) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPositionTicks

`func (o *ModelUserAction) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *ModelUserAction) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *ModelUserAction) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *ModelUserAction) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *ModelUserAction) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *ModelUserAction) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetPlayed

`func (o *ModelUserAction) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *ModelUserAction) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *ModelUserAction) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *ModelUserAction) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### SetPlayedNil

`func (o *ModelUserAction) SetPlayedNil(b bool)`

 SetPlayedNil sets the value for Played to be an explicit nil

### UnsetPlayed
`func (o *ModelUserAction) UnsetPlayed()`

UnsetPlayed ensures that no value is present for Played, not even an explicit nil
### GetIsFavorite

`func (o *ModelUserAction) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *ModelUserAction) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *ModelUserAction) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *ModelUserAction) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *ModelUserAction) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *ModelUserAction) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


