# BlueprintDnf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to **[]string** | Groups to install, must match exactly. Groups describes groups of packages to be installed into the image. Package groups are defined in the repository metadata. Each group has a descriptive name used primarily for display in user interfaces and an ID more commonly used in kickstart files. Here, the ID is the expected way of listing a group. Groups have three different ways of categorizing their packages: mandatory, default, and optional. For the purposes of blueprints, only mandatory and default packages will be installed. There is no mechanism for selecting optional packages. | [optional] 
**ImportKeys** | Pointer to **[]string** | Additional file paths to the GPG keys to import. The files must be present in the image. Does not support importing from URLs. | [optional] 
**Modules** | Pointer to **[]string** | Modules to enable or disable | [optional] 
**Packages** | Pointer to **[]string** | Packages to install. Package name or NVRA is accepted as long as DNF can resolve it. Examples: vim-enhanced, vim-enhanced-9.1.866-1 or vim-enhanced-9.1.866-1.fc41.x86_64. The packages can also be specified as @group_name to install all packages in the group. | [optional] 
**Repositories** | Pointer to [**[]BlueprintDnfRepositoriesInner**](BlueprintDnfRepositoriesInner.md) | Third-party repositories are supported by the blueprint customizations.  All fields reflect configuration values of dnf, see man dnf.conf(5) for more information. | [optional] 

## Methods

### NewBlueprintDnf

`func NewBlueprintDnf() *BlueprintDnf`

NewBlueprintDnf instantiates a new BlueprintDnf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDnfWithDefaults

`func NewBlueprintDnfWithDefaults() *BlueprintDnf`

NewBlueprintDnfWithDefaults instantiates a new BlueprintDnf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *BlueprintDnf) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BlueprintDnf) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BlueprintDnf) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BlueprintDnf) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *BlueprintDnf) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *BlueprintDnf) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetImportKeys

`func (o *BlueprintDnf) GetImportKeys() []string`

GetImportKeys returns the ImportKeys field if non-nil, zero value otherwise.

### GetImportKeysOk

`func (o *BlueprintDnf) GetImportKeysOk() (*[]string, bool)`

GetImportKeysOk returns a tuple with the ImportKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportKeys

`func (o *BlueprintDnf) SetImportKeys(v []string)`

SetImportKeys sets ImportKeys field to given value.

### HasImportKeys

`func (o *BlueprintDnf) HasImportKeys() bool`

HasImportKeys returns a boolean if a field has been set.

### SetImportKeysNil

`func (o *BlueprintDnf) SetImportKeysNil(b bool)`

 SetImportKeysNil sets the value for ImportKeys to be an explicit nil

### UnsetImportKeys
`func (o *BlueprintDnf) UnsetImportKeys()`

UnsetImportKeys ensures that no value is present for ImportKeys, not even an explicit nil
### GetModules

`func (o *BlueprintDnf) GetModules() []string`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *BlueprintDnf) GetModulesOk() (*[]string, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *BlueprintDnf) SetModules(v []string)`

SetModules sets Modules field to given value.

### HasModules

`func (o *BlueprintDnf) HasModules() bool`

HasModules returns a boolean if a field has been set.

### SetModulesNil

`func (o *BlueprintDnf) SetModulesNil(b bool)`

 SetModulesNil sets the value for Modules to be an explicit nil

### UnsetModules
`func (o *BlueprintDnf) UnsetModules()`

UnsetModules ensures that no value is present for Modules, not even an explicit nil
### GetPackages

`func (o *BlueprintDnf) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *BlueprintDnf) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *BlueprintDnf) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *BlueprintDnf) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### SetPackagesNil

`func (o *BlueprintDnf) SetPackagesNil(b bool)`

 SetPackagesNil sets the value for Packages to be an explicit nil

### UnsetPackages
`func (o *BlueprintDnf) UnsetPackages()`

UnsetPackages ensures that no value is present for Packages, not even an explicit nil
### GetRepositories

`func (o *BlueprintDnf) GetRepositories() []BlueprintDnfRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *BlueprintDnf) GetRepositoriesOk() (*[]BlueprintDnfRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *BlueprintDnf) SetRepositories(v []BlueprintDnfRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *BlueprintDnf) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### SetRepositoriesNil

`func (o *BlueprintDnf) SetRepositoriesNil(b bool)`

 SetRepositoriesNil sets the value for Repositories to be an explicit nil

### UnsetRepositories
`func (o *BlueprintDnf) UnsetRepositories()`

UnsetRepositories ensures that no value is present for Repositories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


