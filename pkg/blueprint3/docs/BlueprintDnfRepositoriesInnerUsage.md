# BlueprintDnfRepositoriesInnerUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configure** | Pointer to **bool** | Configure the repository for dnf.  A repository will be saved to the /etc/yum.repos.d directory in an image. An optional filename argument can be set, otherwise the repository will be saved using the the repository ID, i.e. /etc/yum.repos.d/&lt;repo-id&gt;.repo. | [optional] [default to true]
**Install** | Pointer to **bool** | Use the repository for image build.  When this flag is set, it is possible to install third-party packages during the image build. | [optional] [default to false]

## Methods

### NewBlueprintDnfRepositoriesInnerUsage

`func NewBlueprintDnfRepositoriesInnerUsage() *BlueprintDnfRepositoriesInnerUsage`

NewBlueprintDnfRepositoriesInnerUsage instantiates a new BlueprintDnfRepositoriesInnerUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDnfRepositoriesInnerUsageWithDefaults

`func NewBlueprintDnfRepositoriesInnerUsageWithDefaults() *BlueprintDnfRepositoriesInnerUsage`

NewBlueprintDnfRepositoriesInnerUsageWithDefaults instantiates a new BlueprintDnfRepositoriesInnerUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigure

`func (o *BlueprintDnfRepositoriesInnerUsage) GetConfigure() bool`

GetConfigure returns the Configure field if non-nil, zero value otherwise.

### GetConfigureOk

`func (o *BlueprintDnfRepositoriesInnerUsage) GetConfigureOk() (*bool, bool)`

GetConfigureOk returns a tuple with the Configure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigure

`func (o *BlueprintDnfRepositoriesInnerUsage) SetConfigure(v bool)`

SetConfigure sets Configure field to given value.

### HasConfigure

`func (o *BlueprintDnfRepositoriesInnerUsage) HasConfigure() bool`

HasConfigure returns a boolean if a field has been set.

### GetInstall

`func (o *BlueprintDnfRepositoriesInnerUsage) GetInstall() bool`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *BlueprintDnfRepositoriesInnerUsage) GetInstallOk() (*bool, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *BlueprintDnfRepositoriesInnerUsage) SetInstall(v bool)`

SetInstall sets Install field to given value.

### HasInstall

`func (o *BlueprintDnfRepositoriesInnerUsage) HasInstall() bool`

HasInstall returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


