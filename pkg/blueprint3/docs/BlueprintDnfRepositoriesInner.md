# BlueprintDnfRepositoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Repository ID. Required. | 
**BaseUrls** | Pointer to **[]string** | Base URLs for the repository. | [optional] 
**Filename** | Pointer to **string** | Repository filename to use for the repository configuration file. If not provided, the ID is used. Filename must be provided without the .repo extension. | [optional] 
**GpgCheck** | Pointer to **bool** | Enable GPG check for the repository. | [optional] [default to true]
**GpgCheckRepo** | Pointer to **bool** | Enable GPG check for the repository metadata. | [optional] [default to true]
**GpgKeys** | Pointer to **[]string** | GPG keys for the repository.  The blueprint accepts both inline GPG keys and GPG key urls. If an inline GPG key is provided it will be saved to the /etc/pki/rpm-gpg directory and will be referenced accordingly in the repository configuration. GPG keys are not imported to the RPM database and will only be imported when first installing a package from the third-party repository. | [optional] 
**Metalink** | Pointer to **string** | Metalink for the repository. | [optional] 
**MirrorList** | Pointer to **string** | Mirror list for the repository. | [optional] 
**ModuleHotfixes** | Pointer to **bool** | Enable module hotfixes for the repository.  Adds module_hotfixes flag to all repo types so it can be used during osbuild. This enables users to disable modularity filtering on specific repositories. | [optional] [default to false]
**Name** | Pointer to **string** | Repository name. | [optional] 
**Priority** | Pointer to **int32** | Repository priority. | [optional] [default to 99]
**SslVerify** | Pointer to **bool** | Enable SSL verification for the repository. | [optional] [default to true]
**Usage** | Pointer to [**BlueprintDnfRepositoriesInnerUsage**](BlueprintDnfRepositoriesInnerUsage.md) |  | [optional] 

## Methods

### NewBlueprintDnfRepositoriesInner

`func NewBlueprintDnfRepositoriesInner(id string, ) *BlueprintDnfRepositoriesInner`

NewBlueprintDnfRepositoriesInner instantiates a new BlueprintDnfRepositoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintDnfRepositoriesInnerWithDefaults

`func NewBlueprintDnfRepositoriesInnerWithDefaults() *BlueprintDnfRepositoriesInner`

NewBlueprintDnfRepositoriesInnerWithDefaults instantiates a new BlueprintDnfRepositoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlueprintDnfRepositoriesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlueprintDnfRepositoriesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlueprintDnfRepositoriesInner) SetId(v string)`

SetId sets Id field to given value.


### GetBaseUrls

`func (o *BlueprintDnfRepositoriesInner) GetBaseUrls() []string`

GetBaseUrls returns the BaseUrls field if non-nil, zero value otherwise.

### GetBaseUrlsOk

`func (o *BlueprintDnfRepositoriesInner) GetBaseUrlsOk() (*[]string, bool)`

GetBaseUrlsOk returns a tuple with the BaseUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrls

`func (o *BlueprintDnfRepositoriesInner) SetBaseUrls(v []string)`

SetBaseUrls sets BaseUrls field to given value.

### HasBaseUrls

`func (o *BlueprintDnfRepositoriesInner) HasBaseUrls() bool`

HasBaseUrls returns a boolean if a field has been set.

### GetFilename

`func (o *BlueprintDnfRepositoriesInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *BlueprintDnfRepositoriesInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *BlueprintDnfRepositoriesInner) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *BlueprintDnfRepositoriesInner) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetGpgCheck

`func (o *BlueprintDnfRepositoriesInner) GetGpgCheck() bool`

GetGpgCheck returns the GpgCheck field if non-nil, zero value otherwise.

### GetGpgCheckOk

`func (o *BlueprintDnfRepositoriesInner) GetGpgCheckOk() (*bool, bool)`

GetGpgCheckOk returns a tuple with the GpgCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgCheck

`func (o *BlueprintDnfRepositoriesInner) SetGpgCheck(v bool)`

SetGpgCheck sets GpgCheck field to given value.

### HasGpgCheck

`func (o *BlueprintDnfRepositoriesInner) HasGpgCheck() bool`

HasGpgCheck returns a boolean if a field has been set.

### GetGpgCheckRepo

`func (o *BlueprintDnfRepositoriesInner) GetGpgCheckRepo() bool`

GetGpgCheckRepo returns the GpgCheckRepo field if non-nil, zero value otherwise.

### GetGpgCheckRepoOk

`func (o *BlueprintDnfRepositoriesInner) GetGpgCheckRepoOk() (*bool, bool)`

GetGpgCheckRepoOk returns a tuple with the GpgCheckRepo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgCheckRepo

`func (o *BlueprintDnfRepositoriesInner) SetGpgCheckRepo(v bool)`

SetGpgCheckRepo sets GpgCheckRepo field to given value.

### HasGpgCheckRepo

`func (o *BlueprintDnfRepositoriesInner) HasGpgCheckRepo() bool`

HasGpgCheckRepo returns a boolean if a field has been set.

### GetGpgKeys

`func (o *BlueprintDnfRepositoriesInner) GetGpgKeys() []string`

GetGpgKeys returns the GpgKeys field if non-nil, zero value otherwise.

### GetGpgKeysOk

`func (o *BlueprintDnfRepositoriesInner) GetGpgKeysOk() (*[]string, bool)`

GetGpgKeysOk returns a tuple with the GpgKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpgKeys

`func (o *BlueprintDnfRepositoriesInner) SetGpgKeys(v []string)`

SetGpgKeys sets GpgKeys field to given value.

### HasGpgKeys

`func (o *BlueprintDnfRepositoriesInner) HasGpgKeys() bool`

HasGpgKeys returns a boolean if a field has been set.

### GetMetalink

`func (o *BlueprintDnfRepositoriesInner) GetMetalink() string`

GetMetalink returns the Metalink field if non-nil, zero value otherwise.

### GetMetalinkOk

`func (o *BlueprintDnfRepositoriesInner) GetMetalinkOk() (*string, bool)`

GetMetalinkOk returns a tuple with the Metalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalink

`func (o *BlueprintDnfRepositoriesInner) SetMetalink(v string)`

SetMetalink sets Metalink field to given value.

### HasMetalink

`func (o *BlueprintDnfRepositoriesInner) HasMetalink() bool`

HasMetalink returns a boolean if a field has been set.

### GetMirrorList

`func (o *BlueprintDnfRepositoriesInner) GetMirrorList() string`

GetMirrorList returns the MirrorList field if non-nil, zero value otherwise.

### GetMirrorListOk

`func (o *BlueprintDnfRepositoriesInner) GetMirrorListOk() (*string, bool)`

GetMirrorListOk returns a tuple with the MirrorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirrorList

`func (o *BlueprintDnfRepositoriesInner) SetMirrorList(v string)`

SetMirrorList sets MirrorList field to given value.

### HasMirrorList

`func (o *BlueprintDnfRepositoriesInner) HasMirrorList() bool`

HasMirrorList returns a boolean if a field has been set.

### GetModuleHotfixes

`func (o *BlueprintDnfRepositoriesInner) GetModuleHotfixes() bool`

GetModuleHotfixes returns the ModuleHotfixes field if non-nil, zero value otherwise.

### GetModuleHotfixesOk

`func (o *BlueprintDnfRepositoriesInner) GetModuleHotfixesOk() (*bool, bool)`

GetModuleHotfixesOk returns a tuple with the ModuleHotfixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleHotfixes

`func (o *BlueprintDnfRepositoriesInner) SetModuleHotfixes(v bool)`

SetModuleHotfixes sets ModuleHotfixes field to given value.

### HasModuleHotfixes

`func (o *BlueprintDnfRepositoriesInner) HasModuleHotfixes() bool`

HasModuleHotfixes returns a boolean if a field has been set.

### GetName

`func (o *BlueprintDnfRepositoriesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintDnfRepositoriesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintDnfRepositoriesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintDnfRepositoriesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *BlueprintDnfRepositoriesInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BlueprintDnfRepositoriesInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BlueprintDnfRepositoriesInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BlueprintDnfRepositoriesInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSslVerify

`func (o *BlueprintDnfRepositoriesInner) GetSslVerify() bool`

GetSslVerify returns the SslVerify field if non-nil, zero value otherwise.

### GetSslVerifyOk

`func (o *BlueprintDnfRepositoriesInner) GetSslVerifyOk() (*bool, bool)`

GetSslVerifyOk returns a tuple with the SslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerify

`func (o *BlueprintDnfRepositoriesInner) SetSslVerify(v bool)`

SetSslVerify sets SslVerify field to given value.

### HasSslVerify

`func (o *BlueprintDnfRepositoriesInner) HasSslVerify() bool`

HasSslVerify returns a boolean if a field has been set.

### GetUsage

`func (o *BlueprintDnfRepositoriesInner) GetUsage() BlueprintDnfRepositoriesInnerUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *BlueprintDnfRepositoriesInner) GetUsageOk() (*BlueprintDnfRepositoriesInnerUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *BlueprintDnfRepositoriesInner) SetUsage(v BlueprintDnfRepositoriesInnerUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *BlueprintDnfRepositoriesInner) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


