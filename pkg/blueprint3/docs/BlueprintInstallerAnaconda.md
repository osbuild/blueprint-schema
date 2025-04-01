# BlueprintInstallerAnaconda

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisabledModules** | Pointer to **[]string** |  | [optional] 
**EnabledModules** | Pointer to **[]string** |  | [optional] 
**Kickstart** | Pointer to [**NullableBlueprintInstallerAnacondaKickstart**](BlueprintInstallerAnacondaKickstart.md) |  | [optional] 
**SudoNopasswd** | Pointer to **[]string** | Sudo users with NOPASSWD option. Adds a snippet to the kickstart file that, after installation, will create drop-in files in /etc/sudoers.d to allow the specified users and groups to run sudo without a password (groups must be prefixed with %). | [optional] 
**Unattended** | Pointer to **bool** | Unattended installation Anaconda flag. When not set, Anaconda installer will ask for user input. | [optional] 

## Methods

### NewBlueprintInstallerAnaconda

`func NewBlueprintInstallerAnaconda() *BlueprintInstallerAnaconda`

NewBlueprintInstallerAnaconda instantiates a new BlueprintInstallerAnaconda object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintInstallerAnacondaWithDefaults

`func NewBlueprintInstallerAnacondaWithDefaults() *BlueprintInstallerAnaconda`

NewBlueprintInstallerAnacondaWithDefaults instantiates a new BlueprintInstallerAnaconda object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabledModules

`func (o *BlueprintInstallerAnaconda) GetDisabledModules() []string`

GetDisabledModules returns the DisabledModules field if non-nil, zero value otherwise.

### GetDisabledModulesOk

`func (o *BlueprintInstallerAnaconda) GetDisabledModulesOk() (*[]string, bool)`

GetDisabledModulesOk returns a tuple with the DisabledModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledModules

`func (o *BlueprintInstallerAnaconda) SetDisabledModules(v []string)`

SetDisabledModules sets DisabledModules field to given value.

### HasDisabledModules

`func (o *BlueprintInstallerAnaconda) HasDisabledModules() bool`

HasDisabledModules returns a boolean if a field has been set.

### SetDisabledModulesNil

`func (o *BlueprintInstallerAnaconda) SetDisabledModulesNil(b bool)`

 SetDisabledModulesNil sets the value for DisabledModules to be an explicit nil

### UnsetDisabledModules
`func (o *BlueprintInstallerAnaconda) UnsetDisabledModules()`

UnsetDisabledModules ensures that no value is present for DisabledModules, not even an explicit nil
### GetEnabledModules

`func (o *BlueprintInstallerAnaconda) GetEnabledModules() []string`

GetEnabledModules returns the EnabledModules field if non-nil, zero value otherwise.

### GetEnabledModulesOk

`func (o *BlueprintInstallerAnaconda) GetEnabledModulesOk() (*[]string, bool)`

GetEnabledModulesOk returns a tuple with the EnabledModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledModules

`func (o *BlueprintInstallerAnaconda) SetEnabledModules(v []string)`

SetEnabledModules sets EnabledModules field to given value.

### HasEnabledModules

`func (o *BlueprintInstallerAnaconda) HasEnabledModules() bool`

HasEnabledModules returns a boolean if a field has been set.

### SetEnabledModulesNil

`func (o *BlueprintInstallerAnaconda) SetEnabledModulesNil(b bool)`

 SetEnabledModulesNil sets the value for EnabledModules to be an explicit nil

### UnsetEnabledModules
`func (o *BlueprintInstallerAnaconda) UnsetEnabledModules()`

UnsetEnabledModules ensures that no value is present for EnabledModules, not even an explicit nil
### GetKickstart

`func (o *BlueprintInstallerAnaconda) GetKickstart() BlueprintInstallerAnacondaKickstart`

GetKickstart returns the Kickstart field if non-nil, zero value otherwise.

### GetKickstartOk

`func (o *BlueprintInstallerAnaconda) GetKickstartOk() (*BlueprintInstallerAnacondaKickstart, bool)`

GetKickstartOk returns a tuple with the Kickstart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKickstart

`func (o *BlueprintInstallerAnaconda) SetKickstart(v BlueprintInstallerAnacondaKickstart)`

SetKickstart sets Kickstart field to given value.

### HasKickstart

`func (o *BlueprintInstallerAnaconda) HasKickstart() bool`

HasKickstart returns a boolean if a field has been set.

### SetKickstartNil

`func (o *BlueprintInstallerAnaconda) SetKickstartNil(b bool)`

 SetKickstartNil sets the value for Kickstart to be an explicit nil

### UnsetKickstart
`func (o *BlueprintInstallerAnaconda) UnsetKickstart()`

UnsetKickstart ensures that no value is present for Kickstart, not even an explicit nil
### GetSudoNopasswd

`func (o *BlueprintInstallerAnaconda) GetSudoNopasswd() []string`

GetSudoNopasswd returns the SudoNopasswd field if non-nil, zero value otherwise.

### GetSudoNopasswdOk

`func (o *BlueprintInstallerAnaconda) GetSudoNopasswdOk() (*[]string, bool)`

GetSudoNopasswdOk returns a tuple with the SudoNopasswd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudoNopasswd

`func (o *BlueprintInstallerAnaconda) SetSudoNopasswd(v []string)`

SetSudoNopasswd sets SudoNopasswd field to given value.

### HasSudoNopasswd

`func (o *BlueprintInstallerAnaconda) HasSudoNopasswd() bool`

HasSudoNopasswd returns a boolean if a field has been set.

### GetUnattended

`func (o *BlueprintInstallerAnaconda) GetUnattended() bool`

GetUnattended returns the Unattended field if non-nil, zero value otherwise.

### GetUnattendedOk

`func (o *BlueprintInstallerAnaconda) GetUnattendedOk() (*bool, bool)`

GetUnattendedOk returns a tuple with the Unattended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnattended

`func (o *BlueprintInstallerAnaconda) SetUnattended(v bool)`

SetUnattended sets Unattended field to given value.

### HasUnattended

`func (o *BlueprintInstallerAnaconda) HasUnattended() bool`

HasUnattended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


