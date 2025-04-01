# BlueprintInstaller

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anaconda** | Pointer to [**BlueprintInstallerAnaconda**](BlueprintInstallerAnaconda.md) |  | [optional] 
**Coreos** | Pointer to [**BlueprintInstallerCoreos**](BlueprintInstallerCoreos.md) |  | [optional] 

## Methods

### NewBlueprintInstaller

`func NewBlueprintInstaller() *BlueprintInstaller`

NewBlueprintInstaller instantiates a new BlueprintInstaller object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintInstallerWithDefaults

`func NewBlueprintInstallerWithDefaults() *BlueprintInstaller`

NewBlueprintInstallerWithDefaults instantiates a new BlueprintInstaller object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnaconda

`func (o *BlueprintInstaller) GetAnaconda() BlueprintInstallerAnaconda`

GetAnaconda returns the Anaconda field if non-nil, zero value otherwise.

### GetAnacondaOk

`func (o *BlueprintInstaller) GetAnacondaOk() (*BlueprintInstallerAnaconda, bool)`

GetAnacondaOk returns a tuple with the Anaconda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaconda

`func (o *BlueprintInstaller) SetAnaconda(v BlueprintInstallerAnaconda)`

SetAnaconda sets Anaconda field to given value.

### HasAnaconda

`func (o *BlueprintInstaller) HasAnaconda() bool`

HasAnaconda returns a boolean if a field has been set.

### GetCoreos

`func (o *BlueprintInstaller) GetCoreos() BlueprintInstallerCoreos`

GetCoreos returns the Coreos field if non-nil, zero value otherwise.

### GetCoreosOk

`func (o *BlueprintInstaller) GetCoreosOk() (*BlueprintInstallerCoreos, bool)`

GetCoreosOk returns a tuple with the Coreos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoreos

`func (o *BlueprintInstaller) SetCoreos(v BlueprintInstallerCoreos)`

SetCoreos sets Coreos field to given value.

### HasCoreos

`func (o *BlueprintInstaller) HasCoreos() bool`

HasCoreos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


