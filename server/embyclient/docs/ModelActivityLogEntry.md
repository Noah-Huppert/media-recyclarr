# ModelActivityLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**ShortOverview** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ItemId** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to [**ModelModelLoggingLogSeverity**](ModelLoggingLogSeverity.md) |  | [optional] 

## Methods

### NewModelActivityLogEntry

`func NewModelActivityLogEntry() *ModelActivityLogEntry`

NewModelActivityLogEntry instantiates a new ModelActivityLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelActivityLogEntryWithDefaults

`func NewModelActivityLogEntryWithDefaults() *ModelActivityLogEntry`

NewModelActivityLogEntryWithDefaults instantiates a new ModelActivityLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelActivityLogEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelActivityLogEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelActivityLogEntry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelActivityLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelActivityLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelActivityLogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelActivityLogEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelActivityLogEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *ModelActivityLogEntry) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ModelActivityLogEntry) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ModelActivityLogEntry) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ModelActivityLogEntry) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetShortOverview

`func (o *ModelActivityLogEntry) GetShortOverview() string`

GetShortOverview returns the ShortOverview field if non-nil, zero value otherwise.

### GetShortOverviewOk

`func (o *ModelActivityLogEntry) GetShortOverviewOk() (*string, bool)`

GetShortOverviewOk returns a tuple with the ShortOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOverview

`func (o *ModelActivityLogEntry) SetShortOverview(v string)`

SetShortOverview sets ShortOverview field to given value.

### HasShortOverview

`func (o *ModelActivityLogEntry) HasShortOverview() bool`

HasShortOverview returns a boolean if a field has been set.

### GetType

`func (o *ModelActivityLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelActivityLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelActivityLogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelActivityLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemId

`func (o *ModelActivityLogEntry) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ModelActivityLogEntry) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ModelActivityLogEntry) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ModelActivityLogEntry) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetDate

`func (o *ModelActivityLogEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ModelActivityLogEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ModelActivityLogEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ModelActivityLogEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUserId

`func (o *ModelActivityLogEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelActivityLogEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelActivityLogEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelActivityLogEntry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *ModelActivityLogEntry) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *ModelActivityLogEntry) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *ModelActivityLogEntry) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *ModelActivityLogEntry) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### GetSeverity

`func (o *ModelActivityLogEntry) GetSeverity() ModelModelLoggingLogSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ModelActivityLogEntry) GetSeverityOk() (*ModelModelLoggingLogSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ModelActivityLogEntry) SetSeverity(v ModelModelLoggingLogSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ModelActivityLogEntry) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


