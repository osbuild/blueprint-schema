# BlueprintFsnodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type is the type of the file system node, one of: file, dir. | [optional] [default to "file"]
**Contents** | Pointer to [**BlueprintFsnodesInnerContents**](BlueprintFsnodesInnerContents.md) |  | [optional] 
**EnsureParents** | Pointer to **bool** | EnsureParents is a boolean that determines if the parent directories should be created if they do not exist. | [optional] [default to false]
**Group** | Pointer to **string** | Group is the file system node group. Defaults to root. | [optional] [default to "root"]
**Mode** | Pointer to **int32** | Mode is the file system node permissions. Defaults to 0644 for files and 0755 for directories. | [optional] 
**Path** | **string** | Path is the absolute path to the file or directory. | 
**State** | Pointer to **string** | State is the state of the file system node, one of: present, absent. | [optional] [default to "present"]
**User** | Pointer to **string** | User is the file system node owner. Defaults to root. | [optional] [default to "root"]

## Methods

### NewBlueprintFsnodesInner

`func NewBlueprintFsnodesInner(path string, ) *BlueprintFsnodesInner`

NewBlueprintFsnodesInner instantiates a new BlueprintFsnodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintFsnodesInnerWithDefaults

`func NewBlueprintFsnodesInnerWithDefaults() *BlueprintFsnodesInner`

NewBlueprintFsnodesInnerWithDefaults instantiates a new BlueprintFsnodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintFsnodesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintFsnodesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintFsnodesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintFsnodesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContents

`func (o *BlueprintFsnodesInner) GetContents() BlueprintFsnodesInnerContents`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *BlueprintFsnodesInner) GetContentsOk() (*BlueprintFsnodesInnerContents, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *BlueprintFsnodesInner) SetContents(v BlueprintFsnodesInnerContents)`

SetContents sets Contents field to given value.

### HasContents

`func (o *BlueprintFsnodesInner) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetEnsureParents

`func (o *BlueprintFsnodesInner) GetEnsureParents() bool`

GetEnsureParents returns the EnsureParents field if non-nil, zero value otherwise.

### GetEnsureParentsOk

`func (o *BlueprintFsnodesInner) GetEnsureParentsOk() (*bool, bool)`

GetEnsureParentsOk returns a tuple with the EnsureParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsureParents

`func (o *BlueprintFsnodesInner) SetEnsureParents(v bool)`

SetEnsureParents sets EnsureParents field to given value.

### HasEnsureParents

`func (o *BlueprintFsnodesInner) HasEnsureParents() bool`

HasEnsureParents returns a boolean if a field has been set.

### GetGroup

`func (o *BlueprintFsnodesInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *BlueprintFsnodesInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *BlueprintFsnodesInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *BlueprintFsnodesInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMode

`func (o *BlueprintFsnodesInner) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BlueprintFsnodesInner) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BlueprintFsnodesInner) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BlueprintFsnodesInner) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPath

`func (o *BlueprintFsnodesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BlueprintFsnodesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BlueprintFsnodesInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetState

`func (o *BlueprintFsnodesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BlueprintFsnodesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BlueprintFsnodesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BlueprintFsnodesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUser

`func (o *BlueprintFsnodesInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BlueprintFsnodesInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BlueprintFsnodesInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *BlueprintFsnodesInner) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


