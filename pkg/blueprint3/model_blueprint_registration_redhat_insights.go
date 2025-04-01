/*
Blueprint schema

Image Builder Blueprint  WORK IN PROGRESS 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BlueprintRegistrationRedhatInsights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintRegistrationRedhatInsights{}

// BlueprintRegistrationRedhatInsights struct for BlueprintRegistrationRedhatInsights
type BlueprintRegistrationRedhatInsights struct {
	// Enables insights client during boot.
	Enabled bool `json:"enabled"`
}

type _BlueprintRegistrationRedhatInsights BlueprintRegistrationRedhatInsights

// NewBlueprintRegistrationRedhatInsights instantiates a new BlueprintRegistrationRedhatInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintRegistrationRedhatInsights(enabled bool) *BlueprintRegistrationRedhatInsights {
	this := BlueprintRegistrationRedhatInsights{}
	this.Enabled = enabled
	return &this
}

// NewBlueprintRegistrationRedhatInsightsWithDefaults instantiates a new BlueprintRegistrationRedhatInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintRegistrationRedhatInsightsWithDefaults() *BlueprintRegistrationRedhatInsights {
	this := BlueprintRegistrationRedhatInsights{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *BlueprintRegistrationRedhatInsights) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BlueprintRegistrationRedhatInsights) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BlueprintRegistrationRedhatInsights) SetEnabled(v bool) {
	o.Enabled = v
}

func (o BlueprintRegistrationRedhatInsights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintRegistrationRedhatInsights) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

func (o *BlueprintRegistrationRedhatInsights) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"enabled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBlueprintRegistrationRedhatInsights := _BlueprintRegistrationRedhatInsights{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlueprintRegistrationRedhatInsights)

	if err != nil {
		return err
	}

	*o = BlueprintRegistrationRedhatInsights(varBlueprintRegistrationRedhatInsights)

	return err
}

type NullableBlueprintRegistrationRedhatInsights struct {
	value *BlueprintRegistrationRedhatInsights
	isSet bool
}

func (v NullableBlueprintRegistrationRedhatInsights) Get() *BlueprintRegistrationRedhatInsights {
	return v.value
}

func (v *NullableBlueprintRegistrationRedhatInsights) Set(val *BlueprintRegistrationRedhatInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintRegistrationRedhatInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintRegistrationRedhatInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintRegistrationRedhatInsights(val *BlueprintRegistrationRedhatInsights) *NullableBlueprintRegistrationRedhatInsights {
	return &NullableBlueprintRegistrationRedhatInsights{value: val, isSet: true}
}

func (v NullableBlueprintRegistrationRedhatInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintRegistrationRedhatInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


