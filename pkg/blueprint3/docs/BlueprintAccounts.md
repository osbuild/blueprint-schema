# BlueprintAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]BlueprintAccountsGroupsInner**](BlueprintAccountsGroupsInner.md) | Operating system group accounts to be created on the image. | [optional] 
**Users** | Pointer to [**[]BlueprintAccountsUsersInner**](BlueprintAccountsUsersInner.md) | Operating system user accounts to be created on the image. | [optional] 

## Methods

### NewBlueprintAccounts

`func NewBlueprintAccounts() *BlueprintAccounts`

NewBlueprintAccounts instantiates a new BlueprintAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintAccountsWithDefaults

`func NewBlueprintAccountsWithDefaults() *BlueprintAccounts`

NewBlueprintAccountsWithDefaults instantiates a new BlueprintAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *BlueprintAccounts) GetGroups() []BlueprintAccountsGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BlueprintAccounts) GetGroupsOk() (*[]BlueprintAccountsGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BlueprintAccounts) SetGroups(v []BlueprintAccountsGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BlueprintAccounts) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *BlueprintAccounts) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *BlueprintAccounts) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetUsers

`func (o *BlueprintAccounts) GetUsers() []BlueprintAccountsUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BlueprintAccounts) GetUsersOk() (*[]BlueprintAccountsUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BlueprintAccounts) SetUsers(v []BlueprintAccountsUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BlueprintAccounts) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *BlueprintAccounts) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *BlueprintAccounts) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


