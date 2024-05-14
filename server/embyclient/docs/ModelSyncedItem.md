# ModelSyncedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**SyncJobId** | Pointer to **int64** |  | [optional] 
**SyncJobName** | Pointer to **string** |  | [optional] 
**SyncJobDateCreated** | Pointer to **time.Time** |  | [optional] 
**SyncJobItemId** | Pointer to **int64** |  | [optional] 
**OriginalFileName** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AdditionalFiles** | Pointer to [**[]ModelModelItemFileInfo**](ModelModelItemFileInfo.md) |  | [optional] 

## Methods

### NewModelSyncedItem

`func NewModelSyncedItem() *ModelSyncedItem`

NewModelSyncedItem instantiates a new ModelSyncedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSyncedItemWithDefaults

`func NewModelSyncedItemWithDefaults() *ModelSyncedItem`

NewModelSyncedItemWithDefaults instantiates a new ModelSyncedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *ModelSyncedItem) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelSyncedItem) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelSyncedItem) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelSyncedItem) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetSyncJobId

`func (o *ModelSyncedItem) GetSyncJobId() int64`

GetSyncJobId returns the SyncJobId field if non-nil, zero value otherwise.

### GetSyncJobIdOk

`func (o *ModelSyncedItem) GetSyncJobIdOk() (*int64, bool)`

GetSyncJobIdOk returns a tuple with the SyncJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobId

`func (o *ModelSyncedItem) SetSyncJobId(v int64)`

SetSyncJobId sets SyncJobId field to given value.

### HasSyncJobId

`func (o *ModelSyncedItem) HasSyncJobId() bool`

HasSyncJobId returns a boolean if a field has been set.

### GetSyncJobName

`func (o *ModelSyncedItem) GetSyncJobName() string`

GetSyncJobName returns the SyncJobName field if non-nil, zero value otherwise.

### GetSyncJobNameOk

`func (o *ModelSyncedItem) GetSyncJobNameOk() (*string, bool)`

GetSyncJobNameOk returns a tuple with the SyncJobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobName

`func (o *ModelSyncedItem) SetSyncJobName(v string)`

SetSyncJobName sets SyncJobName field to given value.

### HasSyncJobName

`func (o *ModelSyncedItem) HasSyncJobName() bool`

HasSyncJobName returns a boolean if a field has been set.

### GetSyncJobDateCreated

`func (o *ModelSyncedItem) GetSyncJobDateCreated() time.Time`

GetSyncJobDateCreated returns the SyncJobDateCreated field if non-nil, zero value otherwise.

### GetSyncJobDateCreatedOk

`func (o *ModelSyncedItem) GetSyncJobDateCreatedOk() (*time.Time, bool)`

GetSyncJobDateCreatedOk returns a tuple with the SyncJobDateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobDateCreated

`func (o *ModelSyncedItem) SetSyncJobDateCreated(v time.Time)`

SetSyncJobDateCreated sets SyncJobDateCreated field to given value.

### HasSyncJobDateCreated

`func (o *ModelSyncedItem) HasSyncJobDateCreated() bool`

HasSyncJobDateCreated returns a boolean if a field has been set.

### GetSyncJobItemId

`func (o *ModelSyncedItem) GetSyncJobItemId() int64`

GetSyncJobItemId returns the SyncJobItemId field if non-nil, zero value otherwise.

### GetSyncJobItemIdOk

`func (o *ModelSyncedItem) GetSyncJobItemIdOk() (*int64, bool)`

GetSyncJobItemIdOk returns a tuple with the SyncJobItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncJobItemId

`func (o *ModelSyncedItem) SetSyncJobItemId(v int64)`

SetSyncJobItemId sets SyncJobItemId field to given value.

### HasSyncJobItemId

`func (o *ModelSyncedItem) HasSyncJobItemId() bool`

HasSyncJobItemId returns a boolean if a field has been set.

### GetOriginalFileName

`func (o *ModelSyncedItem) GetOriginalFileName() string`

GetOriginalFileName returns the OriginalFileName field if non-nil, zero value otherwise.

### GetOriginalFileNameOk

`func (o *ModelSyncedItem) GetOriginalFileNameOk() (*string, bool)`

GetOriginalFileNameOk returns a tuple with the OriginalFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFileName

`func (o *ModelSyncedItem) SetOriginalFileName(v string)`

SetOriginalFileName sets OriginalFileName field to given value.

### HasOriginalFileName

`func (o *ModelSyncedItem) HasOriginalFileName() bool`

HasOriginalFileName returns a boolean if a field has been set.

### GetItem

`func (o *ModelSyncedItem) GetItem() ModelModelBaseItemDto`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ModelSyncedItem) GetItemOk() (*ModelModelBaseItemDto, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ModelSyncedItem) SetItem(v ModelModelBaseItemDto)`

SetItem sets Item field to given value.

### HasItem

`func (o *ModelSyncedItem) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetUserId

`func (o *ModelSyncedItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelSyncedItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelSyncedItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelSyncedItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAdditionalFiles

`func (o *ModelSyncedItem) GetAdditionalFiles() []ModelModelItemFileInfo`

GetAdditionalFiles returns the AdditionalFiles field if non-nil, zero value otherwise.

### GetAdditionalFilesOk

`func (o *ModelSyncedItem) GetAdditionalFilesOk() (*[]ModelModelItemFileInfo, bool)`

GetAdditionalFilesOk returns a tuple with the AdditionalFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalFiles

`func (o *ModelSyncedItem) SetAdditionalFiles(v []ModelModelItemFileInfo)`

SetAdditionalFiles sets AdditionalFiles field to given value.

### HasAdditionalFiles

`func (o *ModelSyncedItem) HasAdditionalFiles() bool`

HasAdditionalFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


