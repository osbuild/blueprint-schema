# BlueprintAccountsGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **int32** | The group ID (GID) of the group. | [optional] 
**Name** | **string** | Group name. Accepted characters: lowercase letters, digits, underscores, dollars, and hyphens. Name must not start with a hyphen. Maximum length is 256 characters. The validation pattern is a relaxed version of https://github.com/shadow-maint/shadow/blob/master/lib/chkname.c | 

## Methods

### NewBlueprintAccountsGroupsInner

`func NewBlueprintAccountsGroupsInner(name string, ) *BlueprintAccountsGroupsInner`

NewBlueprintAccountsGroupsInner instantiates a new BlueprintAccountsGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintAccountsGroupsInnerWithDefaults

`func NewBlueprintAccountsGroupsInnerWithDefaults() *BlueprintAccountsGroupsInner`

NewBlueprintAccountsGroupsInnerWithDefaults instantiates a new BlueprintAccountsGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *BlueprintAccountsGroupsInner) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *BlueprintAccountsGroupsInner) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *BlueprintAccountsGroupsInner) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *BlueprintAccountsGroupsInner) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetName

`func (o *BlueprintAccountsGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintAccountsGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintAccountsGroupsInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


