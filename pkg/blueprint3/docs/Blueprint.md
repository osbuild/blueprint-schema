# Blueprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description attribute is a string that can be a longer description of the blueprint and is only used for display purposes. | [optional] 
**Accounts** | Pointer to [**BlueprintAccounts**](BlueprintAccounts.md) |  | [optional] 
**Cacerts** | Pointer to [**[]BlueprintCacertsInner**](BlueprintCacertsInner.md) | The CA certificates to be added to the image. The certificates are added to the system-wide CA trust store. The certificates are added to the /etc/pki/ca-trust/source/anchors/ directory and the update-ca-trust command is run to update the system-wide CA trust store. | [optional] 
**Containers** | Pointer to [**[]BlueprintContainersInner**](BlueprintContainersInner.md) | Containers to be pulled during the image build and stored in the image at the default local container storage location that is appropriate for the image type, so that all supported container tools like podman and cri-o will be able to work with it. The embedded containers are not started, to do so you can create systemd unit files or quadlets with the files customization. | [optional] 
**Dnf** | Pointer to [**BlueprintDnf**](BlueprintDnf.md) |  | [optional] 
**Fips** | Pointer to [**BlueprintFips**](BlueprintFips.md) |  | [optional] 
**Fsnodes** | Pointer to [**[]BlueprintFsnodesInner**](BlueprintFsnodesInner.md) | File system nodes details.  You can use the customization to create new files or to replace existing ones, if not restricted by the policy specified below. If the target path is an existing symlink to another file, the symlink will be replaced by the custom file.  Please note that the parent directory of a specified file must exist. If it does not exist, the image build will fail. One can ensure that the parent directory exists by specifying \&quot;ensure_parents\&quot;.  In addition, the following files are not allowed to be created or replaced by policy: /etc/fstab, /etc/shadow, /etc/passwd and /etc/group.  Using the files customization comes with a high chance of creating an image that doesn&#39;t boot. Use this feature only if you know what you are doing. Although the files customization can be used to configure parts of the OS which can also be configured by other blueprint customizations, this use is discouraged. If possible, users should always default to using the specialized blueprint customizations. Note that if you combine the files customizations with other customizations, the other customizations may not work as expected or may be overridden by the files customizations.  You can create custom directories as well. The existence of a specified directory is handled gracefully only if no explicit mode, user or group is specified. If any of these customizations are specified and the directory already exists in the image, the image build will fail. The intention is to prevent changing the ownership or permissions of existing directories. | [optional] 
**Hostname** | Pointer to **string** | Hostname is an optional string that can be used to configure the hostname of the final image. | [optional] 
**Ignition** | Pointer to [**NullableBlueprintIgnition**](BlueprintIgnition.md) |  | [optional] 
**Installer** | Pointer to [**BlueprintInstaller**](BlueprintInstaller.md) |  | [optional] 
**Kernel** | Pointer to [**BlueprintKernel**](BlueprintKernel.md) |  | [optional] 
**Locale** | Pointer to [**BlueprintLocale**](BlueprintLocale.md) |  | [optional] 
**Name** | **string** | The name attribute is a string that contains the name of the blueprint. It can contain spaces, but they may be converted to dash characters during build. It should be short and descriptive. | 
**Network** | Pointer to [**BlueprintNetwork**](BlueprintNetwork.md) |  | [optional] 
**Openscap** | Pointer to [**BlueprintOpenscap**](BlueprintOpenscap.md) |  | [optional] 
**Registration** | Pointer to [**BlueprintRegistration**](BlueprintRegistration.md) |  | [optional] 
**Storage** | Pointer to [**BlueprintStorage**](BlueprintStorage.md) |  | [optional] 
**Systemd** | Pointer to [**BlueprintSystemd**](BlueprintSystemd.md) |  | [optional] 
**Timedate** | Pointer to [**BlueprintTimedate**](BlueprintTimedate.md) |  | [optional] 

## Methods

### NewBlueprint

`func NewBlueprint(name string, ) *Blueprint`

NewBlueprint instantiates a new Blueprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintWithDefaults

`func NewBlueprintWithDefaults() *Blueprint`

NewBlueprintWithDefaults instantiates a new Blueprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Blueprint) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Blueprint) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Blueprint) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Blueprint) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccounts

`func (o *Blueprint) GetAccounts() BlueprintAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Blueprint) GetAccountsOk() (*BlueprintAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Blueprint) SetAccounts(v BlueprintAccounts)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Blueprint) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetCacerts

`func (o *Blueprint) GetCacerts() []BlueprintCacertsInner`

GetCacerts returns the Cacerts field if non-nil, zero value otherwise.

### GetCacertsOk

`func (o *Blueprint) GetCacertsOk() (*[]BlueprintCacertsInner, bool)`

GetCacertsOk returns a tuple with the Cacerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacerts

`func (o *Blueprint) SetCacerts(v []BlueprintCacertsInner)`

SetCacerts sets Cacerts field to given value.

### HasCacerts

`func (o *Blueprint) HasCacerts() bool`

HasCacerts returns a boolean if a field has been set.

### SetCacertsNil

`func (o *Blueprint) SetCacertsNil(b bool)`

 SetCacertsNil sets the value for Cacerts to be an explicit nil

### UnsetCacerts
`func (o *Blueprint) UnsetCacerts()`

UnsetCacerts ensures that no value is present for Cacerts, not even an explicit nil
### GetContainers

`func (o *Blueprint) GetContainers() []BlueprintContainersInner`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *Blueprint) GetContainersOk() (*[]BlueprintContainersInner, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *Blueprint) SetContainers(v []BlueprintContainersInner)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *Blueprint) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### SetContainersNil

`func (o *Blueprint) SetContainersNil(b bool)`

 SetContainersNil sets the value for Containers to be an explicit nil

### UnsetContainers
`func (o *Blueprint) UnsetContainers()`

UnsetContainers ensures that no value is present for Containers, not even an explicit nil
### GetDnf

`func (o *Blueprint) GetDnf() BlueprintDnf`

GetDnf returns the Dnf field if non-nil, zero value otherwise.

### GetDnfOk

`func (o *Blueprint) GetDnfOk() (*BlueprintDnf, bool)`

GetDnfOk returns a tuple with the Dnf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnf

`func (o *Blueprint) SetDnf(v BlueprintDnf)`

SetDnf sets Dnf field to given value.

### HasDnf

`func (o *Blueprint) HasDnf() bool`

HasDnf returns a boolean if a field has been set.

### GetFips

`func (o *Blueprint) GetFips() BlueprintFips`

GetFips returns the Fips field if non-nil, zero value otherwise.

### GetFipsOk

`func (o *Blueprint) GetFipsOk() (*BlueprintFips, bool)`

GetFipsOk returns a tuple with the Fips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFips

`func (o *Blueprint) SetFips(v BlueprintFips)`

SetFips sets Fips field to given value.

### HasFips

`func (o *Blueprint) HasFips() bool`

HasFips returns a boolean if a field has been set.

### GetFsnodes

`func (o *Blueprint) GetFsnodes() []BlueprintFsnodesInner`

GetFsnodes returns the Fsnodes field if non-nil, zero value otherwise.

### GetFsnodesOk

`func (o *Blueprint) GetFsnodesOk() (*[]BlueprintFsnodesInner, bool)`

GetFsnodesOk returns a tuple with the Fsnodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsnodes

`func (o *Blueprint) SetFsnodes(v []BlueprintFsnodesInner)`

SetFsnodes sets Fsnodes field to given value.

### HasFsnodes

`func (o *Blueprint) HasFsnodes() bool`

HasFsnodes returns a boolean if a field has been set.

### SetFsnodesNil

`func (o *Blueprint) SetFsnodesNil(b bool)`

 SetFsnodesNil sets the value for Fsnodes to be an explicit nil

### UnsetFsnodes
`func (o *Blueprint) UnsetFsnodes()`

UnsetFsnodes ensures that no value is present for Fsnodes, not even an explicit nil
### GetHostname

`func (o *Blueprint) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Blueprint) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Blueprint) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Blueprint) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIgnition

`func (o *Blueprint) GetIgnition() BlueprintIgnition`

GetIgnition returns the Ignition field if non-nil, zero value otherwise.

### GetIgnitionOk

`func (o *Blueprint) GetIgnitionOk() (*BlueprintIgnition, bool)`

GetIgnitionOk returns a tuple with the Ignition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnition

`func (o *Blueprint) SetIgnition(v BlueprintIgnition)`

SetIgnition sets Ignition field to given value.

### HasIgnition

`func (o *Blueprint) HasIgnition() bool`

HasIgnition returns a boolean if a field has been set.

### SetIgnitionNil

`func (o *Blueprint) SetIgnitionNil(b bool)`

 SetIgnitionNil sets the value for Ignition to be an explicit nil

### UnsetIgnition
`func (o *Blueprint) UnsetIgnition()`

UnsetIgnition ensures that no value is present for Ignition, not even an explicit nil
### GetInstaller

`func (o *Blueprint) GetInstaller() BlueprintInstaller`

GetInstaller returns the Installer field if non-nil, zero value otherwise.

### GetInstallerOk

`func (o *Blueprint) GetInstallerOk() (*BlueprintInstaller, bool)`

GetInstallerOk returns a tuple with the Installer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstaller

`func (o *Blueprint) SetInstaller(v BlueprintInstaller)`

SetInstaller sets Installer field to given value.

### HasInstaller

`func (o *Blueprint) HasInstaller() bool`

HasInstaller returns a boolean if a field has been set.

### GetKernel

`func (o *Blueprint) GetKernel() BlueprintKernel`

GetKernel returns the Kernel field if non-nil, zero value otherwise.

### GetKernelOk

`func (o *Blueprint) GetKernelOk() (*BlueprintKernel, bool)`

GetKernelOk returns a tuple with the Kernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernel

`func (o *Blueprint) SetKernel(v BlueprintKernel)`

SetKernel sets Kernel field to given value.

### HasKernel

`func (o *Blueprint) HasKernel() bool`

HasKernel returns a boolean if a field has been set.

### GetLocale

`func (o *Blueprint) GetLocale() BlueprintLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Blueprint) GetLocaleOk() (*BlueprintLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Blueprint) SetLocale(v BlueprintLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Blueprint) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *Blueprint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Blueprint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Blueprint) SetName(v string)`

SetName sets Name field to given value.


### GetNetwork

`func (o *Blueprint) GetNetwork() BlueprintNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Blueprint) GetNetworkOk() (*BlueprintNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Blueprint) SetNetwork(v BlueprintNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Blueprint) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOpenscap

`func (o *Blueprint) GetOpenscap() BlueprintOpenscap`

GetOpenscap returns the Openscap field if non-nil, zero value otherwise.

### GetOpenscapOk

`func (o *Blueprint) GetOpenscapOk() (*BlueprintOpenscap, bool)`

GetOpenscapOk returns a tuple with the Openscap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenscap

`func (o *Blueprint) SetOpenscap(v BlueprintOpenscap)`

SetOpenscap sets Openscap field to given value.

### HasOpenscap

`func (o *Blueprint) HasOpenscap() bool`

HasOpenscap returns a boolean if a field has been set.

### GetRegistration

`func (o *Blueprint) GetRegistration() BlueprintRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *Blueprint) GetRegistrationOk() (*BlueprintRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *Blueprint) SetRegistration(v BlueprintRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *Blueprint) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetStorage

`func (o *Blueprint) GetStorage() BlueprintStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Blueprint) GetStorageOk() (*BlueprintStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Blueprint) SetStorage(v BlueprintStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Blueprint) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSystemd

`func (o *Blueprint) GetSystemd() BlueprintSystemd`

GetSystemd returns the Systemd field if non-nil, zero value otherwise.

### GetSystemdOk

`func (o *Blueprint) GetSystemdOk() (*BlueprintSystemd, bool)`

GetSystemdOk returns a tuple with the Systemd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemd

`func (o *Blueprint) SetSystemd(v BlueprintSystemd)`

SetSystemd sets Systemd field to given value.

### HasSystemd

`func (o *Blueprint) HasSystemd() bool`

HasSystemd returns a boolean if a field has been set.

### GetTimedate

`func (o *Blueprint) GetTimedate() BlueprintTimedate`

GetTimedate returns the Timedate field if non-nil, zero value otherwise.

### GetTimedateOk

`func (o *Blueprint) GetTimedateOk() (*BlueprintTimedate, bool)`

GetTimedateOk returns a tuple with the Timedate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedate

`func (o *Blueprint) SetTimedate(v BlueprintTimedate)`

SetTimedate sets Timedate field to given value.

### HasTimedate

`func (o *Blueprint) HasTimedate() bool`

HasTimedate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


