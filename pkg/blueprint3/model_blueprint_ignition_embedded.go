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

// checks if the BlueprintIgnitionEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintIgnitionEmbedded{}

// BlueprintIgnitionEmbedded struct for BlueprintIgnitionEmbedded
type BlueprintIgnitionEmbedded struct {
	// Ignition data formatted in base64.
	Base64 *string
	// Ignition data formatted in plain text.
	Text *string
}

// NewBlueprintIgnitionEmbedded instantiates a new BlueprintIgnitionEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintIgnitionEmbedded() *BlueprintIgnitionEmbedded {
	this := BlueprintIgnitionEmbedded{}
	return &this
}

// NewBlueprintIgnitionEmbeddedWithDefaults instantiates a new BlueprintIgnitionEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintIgnitionEmbeddedWithDefaults() *BlueprintIgnitionEmbedded {
	this := BlueprintIgnitionEmbedded{}
	return &this
}

// GetBase64 returns the Base64 field value if set, zero value otherwise.
func (o *BlueprintIgnitionEmbedded) GetBase64() string {
	if o == nil || IsNil(o.Base64) {
		var ret string
		return ret
	}
	return *o.Base64
}

// GetBase64Ok returns a tuple with the Base64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintIgnitionEmbedded) GetBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.Base64) {
		return nil, false
	}
	return o.Base64, true
}

// HasBase64 returns a boolean if a field has been set.
func (o *BlueprintIgnitionEmbedded) HasBase64() bool {
	if o != nil && !IsNil(o.Base64) {
		return true
	}

	return false
}

// SetBase64 gets a reference to the given string and assigns it to the Base64 field.
func (o *BlueprintIgnitionEmbedded) SetBase64(v string) {
	o.Base64 = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *BlueprintIgnitionEmbedded) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintIgnitionEmbedded) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *BlueprintIgnitionEmbedded) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *BlueprintIgnitionEmbedded) SetText(v string) {
	o.Text = &v
}

func (o BlueprintIgnitionEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintIgnitionEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Base64) {
		toSerialize["base64"] = o.Base64
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableBlueprintIgnitionEmbedded struct {
	value *BlueprintIgnitionEmbedded
	isSet bool
}

func (v NullableBlueprintIgnitionEmbedded) Get() *BlueprintIgnitionEmbedded {
	return v.value
}

func (v *NullableBlueprintIgnitionEmbedded) Set(val *BlueprintIgnitionEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintIgnitionEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintIgnitionEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintIgnitionEmbedded(val *BlueprintIgnitionEmbedded) *NullableBlueprintIgnitionEmbedded {
	return &NullableBlueprintIgnitionEmbedded{value: val, isSet: true}
}

func (v NullableBlueprintIgnitionEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintIgnitionEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


