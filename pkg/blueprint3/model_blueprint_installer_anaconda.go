/*
Blueprint schema

Image Builder Blueprint  WORK IN PROGRESS 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BlueprintInstallerAnaconda type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintInstallerAnaconda{}

// BlueprintInstallerAnaconda struct for BlueprintInstallerAnaconda
type BlueprintInstallerAnaconda struct {
	DisabledModules []string `json:"disabled_modules,omitempty"`
	EnabledModules []string `json:"enabled_modules,omitempty"`
	Kickstart NullableBlueprintInstallerAnacondaKickstart `json:"kickstart,omitempty"`
	// Sudo users with NOPASSWD option. Adds a snippet to the kickstart file that, after installation, will create drop-in files in /etc/sudoers.d to allow the specified users and groups to run sudo without a password (groups must be prefixed with %).
	SudoNopasswd []string `json:"sudo_nopasswd,omitempty"`
	// Unattended installation Anaconda flag. When not set, Anaconda installer will ask for user input.
	Unattended *bool `json:"unattended,omitempty"`
}

// NewBlueprintInstallerAnaconda instantiates a new BlueprintInstallerAnaconda object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintInstallerAnaconda() *BlueprintInstallerAnaconda {
	this := BlueprintInstallerAnaconda{}
	return &this
}

// NewBlueprintInstallerAnacondaWithDefaults instantiates a new BlueprintInstallerAnaconda object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintInstallerAnacondaWithDefaults() *BlueprintInstallerAnaconda {
	this := BlueprintInstallerAnaconda{}
	return &this
}

// GetDisabledModules returns the DisabledModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintInstallerAnaconda) GetDisabledModules() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisabledModules
}

// GetDisabledModulesOk returns a tuple with the DisabledModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintInstallerAnaconda) GetDisabledModulesOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledModules) {
		return nil, false
	}
	return o.DisabledModules, true
}

// HasDisabledModules returns a boolean if a field has been set.
func (o *BlueprintInstallerAnaconda) HasDisabledModules() bool {
	if o != nil && !IsNil(o.DisabledModules) {
		return true
	}

	return false
}

// SetDisabledModules gets a reference to the given []string and assigns it to the DisabledModules field.
func (o *BlueprintInstallerAnaconda) SetDisabledModules(v []string) {
	o.DisabledModules = v
}

// GetEnabledModules returns the EnabledModules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintInstallerAnaconda) GetEnabledModules() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnabledModules
}

// GetEnabledModulesOk returns a tuple with the EnabledModules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintInstallerAnaconda) GetEnabledModulesOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledModules) {
		return nil, false
	}
	return o.EnabledModules, true
}

// HasEnabledModules returns a boolean if a field has been set.
func (o *BlueprintInstallerAnaconda) HasEnabledModules() bool {
	if o != nil && !IsNil(o.EnabledModules) {
		return true
	}

	return false
}

// SetEnabledModules gets a reference to the given []string and assigns it to the EnabledModules field.
func (o *BlueprintInstallerAnaconda) SetEnabledModules(v []string) {
	o.EnabledModules = v
}

// GetKickstart returns the Kickstart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlueprintInstallerAnaconda) GetKickstart() BlueprintInstallerAnacondaKickstart {
	if o == nil || IsNil(o.Kickstart.Get()) {
		var ret BlueprintInstallerAnacondaKickstart
		return ret
	}
	return *o.Kickstart.Get()
}

// GetKickstartOk returns a tuple with the Kickstart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlueprintInstallerAnaconda) GetKickstartOk() (*BlueprintInstallerAnacondaKickstart, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kickstart.Get(), o.Kickstart.IsSet()
}

// HasKickstart returns a boolean if a field has been set.
func (o *BlueprintInstallerAnaconda) HasKickstart() bool {
	if o != nil && o.Kickstart.IsSet() {
		return true
	}

	return false
}

// SetKickstart gets a reference to the given NullableBlueprintInstallerAnacondaKickstart and assigns it to the Kickstart field.
func (o *BlueprintInstallerAnaconda) SetKickstart(v BlueprintInstallerAnacondaKickstart) {
	o.Kickstart.Set(&v)
}
// SetKickstartNil sets the value for Kickstart to be an explicit nil
func (o *BlueprintInstallerAnaconda) SetKickstartNil() {
	o.Kickstart.Set(nil)
}

// UnsetKickstart ensures that no value is present for Kickstart, not even an explicit nil
func (o *BlueprintInstallerAnaconda) UnsetKickstart() {
	o.Kickstart.Unset()
}

// GetSudoNopasswd returns the SudoNopasswd field value if set, zero value otherwise.
func (o *BlueprintInstallerAnaconda) GetSudoNopasswd() []string {
	if o == nil || IsNil(o.SudoNopasswd) {
		var ret []string
		return ret
	}
	return o.SudoNopasswd
}

// GetSudoNopasswdOk returns a tuple with the SudoNopasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstallerAnaconda) GetSudoNopasswdOk() ([]string, bool) {
	if o == nil || IsNil(o.SudoNopasswd) {
		return nil, false
	}
	return o.SudoNopasswd, true
}

// HasSudoNopasswd returns a boolean if a field has been set.
func (o *BlueprintInstallerAnaconda) HasSudoNopasswd() bool {
	if o != nil && !IsNil(o.SudoNopasswd) {
		return true
	}

	return false
}

// SetSudoNopasswd gets a reference to the given []string and assigns it to the SudoNopasswd field.
func (o *BlueprintInstallerAnaconda) SetSudoNopasswd(v []string) {
	o.SudoNopasswd = v
}

// GetUnattended returns the Unattended field value if set, zero value otherwise.
func (o *BlueprintInstallerAnaconda) GetUnattended() bool {
	if o == nil || IsNil(o.Unattended) {
		var ret bool
		return ret
	}
	return *o.Unattended
}

// GetUnattendedOk returns a tuple with the Unattended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintInstallerAnaconda) GetUnattendedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unattended) {
		return nil, false
	}
	return o.Unattended, true
}

// HasUnattended returns a boolean if a field has been set.
func (o *BlueprintInstallerAnaconda) HasUnattended() bool {
	if o != nil && !IsNil(o.Unattended) {
		return true
	}

	return false
}

// SetUnattended gets a reference to the given bool and assigns it to the Unattended field.
func (o *BlueprintInstallerAnaconda) SetUnattended(v bool) {
	o.Unattended = &v
}

func (o BlueprintInstallerAnaconda) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintInstallerAnaconda) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DisabledModules != nil {
		toSerialize["disabled_modules"] = o.DisabledModules
	}
	if o.EnabledModules != nil {
		toSerialize["enabled_modules"] = o.EnabledModules
	}
	if o.Kickstart.IsSet() {
		toSerialize["kickstart"] = o.Kickstart.Get()
	}
	if !IsNil(o.SudoNopasswd) {
		toSerialize["sudo_nopasswd"] = o.SudoNopasswd
	}
	if !IsNil(o.Unattended) {
		toSerialize["unattended"] = o.Unattended
	}
	return toSerialize, nil
}

type NullableBlueprintInstallerAnaconda struct {
	value *BlueprintInstallerAnaconda
	isSet bool
}

func (v NullableBlueprintInstallerAnaconda) Get() *BlueprintInstallerAnaconda {
	return v.value
}

func (v *NullableBlueprintInstallerAnaconda) Set(val *BlueprintInstallerAnaconda) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintInstallerAnaconda) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintInstallerAnaconda) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintInstallerAnaconda(val *BlueprintInstallerAnaconda) *NullableBlueprintInstallerAnaconda {
	return &NullableBlueprintInstallerAnaconda{value: val, isSet: true}
}

func (v NullableBlueprintInstallerAnaconda) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintInstallerAnaconda) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


