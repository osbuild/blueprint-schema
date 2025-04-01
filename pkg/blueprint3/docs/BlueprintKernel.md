# BlueprintKernel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmdlineAppend** | Pointer to **[]string** | An optional string to append arguments to the bootloader kernel command line. The list will be concatenated with spaces. | [optional] 
**Package** | Pointer to **string** | Kernel DNF package name to replace the standard kernel with. | [optional] 

## Methods

### NewBlueprintKernel

`func NewBlueprintKernel() *BlueprintKernel`

NewBlueprintKernel instantiates a new BlueprintKernel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintKernelWithDefaults

`func NewBlueprintKernelWithDefaults() *BlueprintKernel`

NewBlueprintKernelWithDefaults instantiates a new BlueprintKernel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCmdlineAppend

`func (o *BlueprintKernel) GetCmdlineAppend() []string`

GetCmdlineAppend returns the CmdlineAppend field if non-nil, zero value otherwise.

### GetCmdlineAppendOk

`func (o *BlueprintKernel) GetCmdlineAppendOk() (*[]string, bool)`

GetCmdlineAppendOk returns a tuple with the CmdlineAppend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdlineAppend

`func (o *BlueprintKernel) SetCmdlineAppend(v []string)`

SetCmdlineAppend sets CmdlineAppend field to given value.

### HasCmdlineAppend

`func (o *BlueprintKernel) HasCmdlineAppend() bool`

HasCmdlineAppend returns a boolean if a field has been set.

### SetCmdlineAppendNil

`func (o *BlueprintKernel) SetCmdlineAppendNil(b bool)`

 SetCmdlineAppendNil sets the value for CmdlineAppend to be an explicit nil

### UnsetCmdlineAppend
`func (o *BlueprintKernel) UnsetCmdlineAppend()`

UnsetCmdlineAppend ensures that no value is present for CmdlineAppend, not even an explicit nil
### GetPackage

`func (o *BlueprintKernel) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *BlueprintKernel) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *BlueprintKernel) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *BlueprintKernel) HasPackage() bool`

HasPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


