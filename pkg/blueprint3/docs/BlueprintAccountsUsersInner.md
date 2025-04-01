# BlueprintAccountsUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A longer description of the account. | [optional] 
**Expires** | Pointer to **string** | Date type which accepts date (YYYY-MM-DD) or date-time (RFC3339) format and only marshals into date format. This is needed for JSON/YAML compatibility since YAML automatically converts strings which look like dates into time.Time. | [optional] 
**Gid** | Pointer to **int32** | The primary group ID (GID) of the user. Value of zero (or ommited value) means that the next available UID will be assigned. | [optional] 
**Groups** | Pointer to **[]string** | Additional groups that the user should be a member of. | [optional] 
**Home** | Pointer to **string** | The home directory of the user. | [optional] 
**Name** | **string** | Account name. Accepted characters: lowercase letters, digits, underscores, dollars, and hyphens. Name must not start with a hyphen. Maximum length is 256 characters. The validation pattern is a relaxed version of https://github.com/shadow-maint/shadow/blob/master/lib/chkname.c | 
**Password** | Pointer to **string** | Password either in plain text or encrypted form. If the password is not provided, the account will be locked and the user will not be able to log in with a password. The password can be encrypted using the crypt(3) function. The format of the encrypted password is $id$salt$hashed, where $id is the algorithm used (1, 5, 6, or 2a). | [optional] 
**Shell** | Pointer to **string** | The shell of the user. | [optional] 
**SshKeys** | Pointer to **[]string** | SSH keys to be added to the account&#39;s authorized_keys file. | [optional] 
**Uid** | Pointer to **int32** | The user ID (UID) of the user. Value of zero (or ommited value) means that the next available UID will be assigned. | [optional] 

## Methods

### NewBlueprintAccountsUsersInner

`func NewBlueprintAccountsUsersInner(name string, ) *BlueprintAccountsUsersInner`

NewBlueprintAccountsUsersInner instantiates a new BlueprintAccountsUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintAccountsUsersInnerWithDefaults

`func NewBlueprintAccountsUsersInnerWithDefaults() *BlueprintAccountsUsersInner`

NewBlueprintAccountsUsersInnerWithDefaults instantiates a new BlueprintAccountsUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BlueprintAccountsUsersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlueprintAccountsUsersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlueprintAccountsUsersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlueprintAccountsUsersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpires

`func (o *BlueprintAccountsUsersInner) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *BlueprintAccountsUsersInner) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *BlueprintAccountsUsersInner) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *BlueprintAccountsUsersInner) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetGid

`func (o *BlueprintAccountsUsersInner) GetGid() int32`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *BlueprintAccountsUsersInner) GetGidOk() (*int32, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *BlueprintAccountsUsersInner) SetGid(v int32)`

SetGid sets Gid field to given value.

### HasGid

`func (o *BlueprintAccountsUsersInner) HasGid() bool`

HasGid returns a boolean if a field has been set.

### GetGroups

`func (o *BlueprintAccountsUsersInner) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BlueprintAccountsUsersInner) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BlueprintAccountsUsersInner) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BlueprintAccountsUsersInner) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHome

`func (o *BlueprintAccountsUsersInner) GetHome() string`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *BlueprintAccountsUsersInner) GetHomeOk() (*string, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *BlueprintAccountsUsersInner) SetHome(v string)`

SetHome sets Home field to given value.

### HasHome

`func (o *BlueprintAccountsUsersInner) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetName

`func (o *BlueprintAccountsUsersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintAccountsUsersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintAccountsUsersInner) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *BlueprintAccountsUsersInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BlueprintAccountsUsersInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BlueprintAccountsUsersInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BlueprintAccountsUsersInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetShell

`func (o *BlueprintAccountsUsersInner) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *BlueprintAccountsUsersInner) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *BlueprintAccountsUsersInner) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *BlueprintAccountsUsersInner) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetSshKeys

`func (o *BlueprintAccountsUsersInner) GetSshKeys() []string`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *BlueprintAccountsUsersInner) GetSshKeysOk() (*[]string, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *BlueprintAccountsUsersInner) SetSshKeys(v []string)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *BlueprintAccountsUsersInner) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetUid

`func (o *BlueprintAccountsUsersInner) GetUid() int32`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *BlueprintAccountsUsersInner) GetUidOk() (*int32, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *BlueprintAccountsUsersInner) SetUid(v int32)`

SetUid sets Uid field to given value.

### HasUid

`func (o *BlueprintAccountsUsersInner) HasUid() bool`

HasUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


