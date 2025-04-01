# BlueprintSystemd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **[]string** | The disabled attribute is a list of strings that contains the systemd units to be disabled. | [optional] 
**Enabled** | Pointer to **[]string** | The enabled attribute is a list of strings that contains the systemd units to be enabled. | [optional] 
**Masked** | Pointer to **[]string** | The masked attribute is a list of strings that contains the systemd units to be masked. | [optional] 

## Methods

### NewBlueprintSystemd

`func NewBlueprintSystemd() *BlueprintSystemd`

NewBlueprintSystemd instantiates a new BlueprintSystemd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintSystemdWithDefaults

`func NewBlueprintSystemdWithDefaults() *BlueprintSystemd`

NewBlueprintSystemdWithDefaults instantiates a new BlueprintSystemd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *BlueprintSystemd) GetDisabled() []string`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *BlueprintSystemd) GetDisabledOk() (*[]string, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *BlueprintSystemd) SetDisabled(v []string)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *BlueprintSystemd) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### SetDisabledNil

`func (o *BlueprintSystemd) SetDisabledNil(b bool)`

 SetDisabledNil sets the value for Disabled to be an explicit nil

### UnsetDisabled
`func (o *BlueprintSystemd) UnsetDisabled()`

UnsetDisabled ensures that no value is present for Disabled, not even an explicit nil
### GetEnabled

`func (o *BlueprintSystemd) GetEnabled() []string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintSystemd) GetEnabledOk() (*[]string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintSystemd) SetEnabled(v []string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BlueprintSystemd) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *BlueprintSystemd) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *BlueprintSystemd) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetMasked

`func (o *BlueprintSystemd) GetMasked() []string`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *BlueprintSystemd) GetMaskedOk() (*[]string, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *BlueprintSystemd) SetMasked(v []string)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *BlueprintSystemd) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### SetMaskedNil

`func (o *BlueprintSystemd) SetMaskedNil(b bool)`

 SetMaskedNil sets the value for Masked to be an explicit nil

### UnsetMasked
`func (o *BlueprintSystemd) UnsetMasked()`

UnsetMasked ensures that no value is present for Masked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


