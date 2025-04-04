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

// checks if the BlueprintFips type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintFips{}

// BlueprintFips struct for BlueprintFips
type BlueprintFips struct {
	// Enables the system FIPS mode (disabled by default). Currently only edge-raw-image, edge-installer, edge-simplified-installer, edge-ami and edge-vsphere images support this customization.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewBlueprintFips instantiates a new BlueprintFips object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintFips() *BlueprintFips {
	this := BlueprintFips{}
	return &this
}

// NewBlueprintFipsWithDefaults instantiates a new BlueprintFips object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintFipsWithDefaults() *BlueprintFips {
	this := BlueprintFips{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BlueprintFips) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintFips) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BlueprintFips) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BlueprintFips) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o BlueprintFips) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintFips) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableBlueprintFips struct {
	value *BlueprintFips
	isSet bool
}

func (v NullableBlueprintFips) Get() *BlueprintFips {
	return v.value
}

func (v *NullableBlueprintFips) Set(val *BlueprintFips) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintFips) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintFips) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintFips(val *BlueprintFips) *NullableBlueprintFips {
	return &NullableBlueprintFips{value: val, isSet: true}
}

func (v NullableBlueprintFips) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintFips) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


