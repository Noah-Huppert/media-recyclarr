# ModelSessionSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**ModelModelPlayerStateInfo**](ModelPlayerStateInfo.md) |  | [optional] 
**AdditionalUsers** | Pointer to [**[]ModelModelSessionUserInfo**](ModelModelSessionUserInfo.md) |  | [optional] 
**RemoteEndPoint** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**PlayableMediaTypes** | Pointer to **[]string** |  | [optional] 
**PlaylistItemId** | Pointer to **string** |  | [optional] 
**PlaylistIndex** | Pointer to **int32** |  | [optional] 
**PlaylistLength** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **string** |  | [optional] 
**Client** | Pointer to **string** |  | [optional] 
**LastActivityDate** | Pointer to **time.Time** |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**DeviceType** | Pointer to **string** |  | [optional] 
**NowPlayingItem** | Pointer to [**ModelModelBaseItemDto**](ModelBaseItemDto.md) |  | [optional] 
**InternalDeviceId** | Pointer to **int64** |  | [optional] 
**DeviceId** | Pointer to **string** |  | [optional] 
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**AppIconUrl** | Pointer to **string** |  | [optional] 
**SupportedCommands** | Pointer to **[]string** |  | [optional] 
**TranscodingInfo** | Pointer to [**ModelModelTranscodingInfo**](ModelTranscodingInfo.md) |  | [optional] 
**SupportsRemoteControl** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelSessionSessionInfo

`func NewModelSessionSessionInfo() *ModelSessionSessionInfo`

NewModelSessionSessionInfo instantiates a new ModelSessionSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelSessionSessionInfoWithDefaults

`func NewModelSessionSessionInfoWithDefaults() *ModelSessionSessionInfo`

NewModelSessionSessionInfoWithDefaults instantiates a new ModelSessionSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *ModelSessionSessionInfo) GetPlayState() ModelModelPlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *ModelSessionSessionInfo) GetPlayStateOk() (*ModelModelPlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *ModelSessionSessionInfo) SetPlayState(v ModelModelPlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *ModelSessionSessionInfo) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### GetAdditionalUsers

`func (o *ModelSessionSessionInfo) GetAdditionalUsers() []ModelModelSessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *ModelSessionSessionInfo) GetAdditionalUsersOk() (*[]ModelModelSessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *ModelSessionSessionInfo) SetAdditionalUsers(v []ModelModelSessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *ModelSessionSessionInfo) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### GetRemoteEndPoint

`func (o *ModelSessionSessionInfo) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *ModelSessionSessionInfo) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *ModelSessionSessionInfo) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *ModelSessionSessionInfo) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### GetProtocol

`func (o *ModelSessionSessionInfo) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ModelSessionSessionInfo) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ModelSessionSessionInfo) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ModelSessionSessionInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPlayableMediaTypes

`func (o *ModelSessionSessionInfo) GetPlayableMediaTypes() []string`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ModelSessionSessionInfo) GetPlayableMediaTypesOk() (*[]string, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ModelSessionSessionInfo) SetPlayableMediaTypes(v []string)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ModelSessionSessionInfo) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *ModelSessionSessionInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *ModelSessionSessionInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *ModelSessionSessionInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *ModelSessionSessionInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetPlaylistIndex

`func (o *ModelSessionSessionInfo) GetPlaylistIndex() int32`

GetPlaylistIndex returns the PlaylistIndex field if non-nil, zero value otherwise.

### GetPlaylistIndexOk

`func (o *ModelSessionSessionInfo) GetPlaylistIndexOk() (*int32, bool)`

GetPlaylistIndexOk returns a tuple with the PlaylistIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistIndex

`func (o *ModelSessionSessionInfo) SetPlaylistIndex(v int32)`

SetPlaylistIndex sets PlaylistIndex field to given value.

### HasPlaylistIndex

`func (o *ModelSessionSessionInfo) HasPlaylistIndex() bool`

HasPlaylistIndex returns a boolean if a field has been set.

### GetPlaylistLength

`func (o *ModelSessionSessionInfo) GetPlaylistLength() int32`

GetPlaylistLength returns the PlaylistLength field if non-nil, zero value otherwise.

### GetPlaylistLengthOk

`func (o *ModelSessionSessionInfo) GetPlaylistLengthOk() (*int32, bool)`

GetPlaylistLengthOk returns a tuple with the PlaylistLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistLength

`func (o *ModelSessionSessionInfo) SetPlaylistLength(v int32)`

SetPlaylistLength sets PlaylistLength field to given value.

### HasPlaylistLength

`func (o *ModelSessionSessionInfo) HasPlaylistLength() bool`

HasPlaylistLength returns a boolean if a field has been set.

### GetId

`func (o *ModelSessionSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelSessionSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelSessionSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelSessionSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerId

`func (o *ModelSessionSessionInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelSessionSessionInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelSessionSessionInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelSessionSessionInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetUserId

`func (o *ModelSessionSessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelSessionSessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelSessionSessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelSessionSessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *ModelSessionSessionInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ModelSessionSessionInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ModelSessionSessionInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ModelSessionSessionInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *ModelSessionSessionInfo) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *ModelSessionSessionInfo) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *ModelSessionSessionInfo) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *ModelSessionSessionInfo) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### GetClient

`func (o *ModelSessionSessionInfo) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ModelSessionSessionInfo) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ModelSessionSessionInfo) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *ModelSessionSessionInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetLastActivityDate

`func (o *ModelSessionSessionInfo) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *ModelSessionSessionInfo) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *ModelSessionSessionInfo) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *ModelSessionSessionInfo) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetDeviceName

`func (o *ModelSessionSessionInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ModelSessionSessionInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ModelSessionSessionInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ModelSessionSessionInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *ModelSessionSessionInfo) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModelSessionSessionInfo) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModelSessionSessionInfo) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ModelSessionSessionInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetNowPlayingItem

`func (o *ModelSessionSessionInfo) GetNowPlayingItem() ModelModelBaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *ModelSessionSessionInfo) GetNowPlayingItemOk() (*ModelModelBaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *ModelSessionSessionInfo) SetNowPlayingItem(v ModelModelBaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *ModelSessionSessionInfo) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### GetInternalDeviceId

`func (o *ModelSessionSessionInfo) GetInternalDeviceId() int64`

GetInternalDeviceId returns the InternalDeviceId field if non-nil, zero value otherwise.

### GetInternalDeviceIdOk

`func (o *ModelSessionSessionInfo) GetInternalDeviceIdOk() (*int64, bool)`

GetInternalDeviceIdOk returns a tuple with the InternalDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDeviceId

`func (o *ModelSessionSessionInfo) SetInternalDeviceId(v int64)`

SetInternalDeviceId sets InternalDeviceId field to given value.

### HasInternalDeviceId

`func (o *ModelSessionSessionInfo) HasInternalDeviceId() bool`

HasInternalDeviceId returns a boolean if a field has been set.

### GetDeviceId

`func (o *ModelSessionSessionInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ModelSessionSessionInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ModelSessionSessionInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ModelSessionSessionInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *ModelSessionSessionInfo) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *ModelSessionSessionInfo) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *ModelSessionSessionInfo) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *ModelSessionSessionInfo) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetAppIconUrl

`func (o *ModelSessionSessionInfo) GetAppIconUrl() string`

GetAppIconUrl returns the AppIconUrl field if non-nil, zero value otherwise.

### GetAppIconUrlOk

`func (o *ModelSessionSessionInfo) GetAppIconUrlOk() (*string, bool)`

GetAppIconUrlOk returns a tuple with the AppIconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIconUrl

`func (o *ModelSessionSessionInfo) SetAppIconUrl(v string)`

SetAppIconUrl sets AppIconUrl field to given value.

### HasAppIconUrl

`func (o *ModelSessionSessionInfo) HasAppIconUrl() bool`

HasAppIconUrl returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *ModelSessionSessionInfo) GetSupportedCommands() []string`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ModelSessionSessionInfo) GetSupportedCommandsOk() (*[]string, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ModelSessionSessionInfo) SetSupportedCommands(v []string)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ModelSessionSessionInfo) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetTranscodingInfo

`func (o *ModelSessionSessionInfo) GetTranscodingInfo() ModelModelTranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *ModelSessionSessionInfo) GetTranscodingInfoOk() (*ModelModelTranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *ModelSessionSessionInfo) SetTranscodingInfo(v ModelModelTranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *ModelSessionSessionInfo) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *ModelSessionSessionInfo) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *ModelSessionSessionInfo) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *ModelSessionSessionInfo) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *ModelSessionSessionInfo) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


