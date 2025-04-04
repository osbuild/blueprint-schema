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

// checks if the BlueprintStoragePartitionsInnerLogicalVolumesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlueprintStoragePartitionsInnerLogicalVolumesInner{}

// BlueprintStoragePartitionsInnerLogicalVolumesInner struct for BlueprintStoragePartitionsInnerLogicalVolumesInner
type BlueprintStoragePartitionsInnerLogicalVolumesInner struct {
	// File system type: ext4 (default), xfs, swap, or vfat.
	FsType *string `json:"fs_type,omitempty"`
	// Label of the logical volume.
	Label *string `json:"label,omitempty"`
	// Minimum size of the logical volume.  Size must be formatted as an integer followed by whitespace and then either a decimal unit (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB).
	Minsize *string `json:"minsize,omitempty" validate:"regexp=^\\\\d+\\\\s*[BKkMGTPE]i?[BKMGTPE]?$"`
	// Mount point of the logical volume. Required except for swap fs_type.
	Mountpoint *string `json:"mountpoint,omitempty" validate:"regexp=^\\/"`
	// Logical volume name. When not set, will be generated automatically.
	Name *string `json:"name,omitempty"`
}

// NewBlueprintStoragePartitionsInnerLogicalVolumesInner instantiates a new BlueprintStoragePartitionsInnerLogicalVolumesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlueprintStoragePartitionsInnerLogicalVolumesInner() *BlueprintStoragePartitionsInnerLogicalVolumesInner {
	this := BlueprintStoragePartitionsInnerLogicalVolumesInner{}
	var fsType string = "ext4"
	this.FsType = &fsType
	return &this
}

// NewBlueprintStoragePartitionsInnerLogicalVolumesInnerWithDefaults instantiates a new BlueprintStoragePartitionsInnerLogicalVolumesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlueprintStoragePartitionsInnerLogicalVolumesInnerWithDefaults() *BlueprintStoragePartitionsInnerLogicalVolumesInner {
	this := BlueprintStoragePartitionsInnerLogicalVolumesInner{}
	var fsType string = "ext4"
	this.FsType = &fsType
	return &this
}

// GetFsType returns the FsType field value if set, zero value otherwise.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetFsType() string {
	if o == nil || IsNil(o.FsType) {
		var ret string
		return ret
	}
	return *o.FsType
}

// GetFsTypeOk returns a tuple with the FsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetFsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FsType) {
		return nil, false
	}
	return o.FsType, true
}

// HasFsType returns a boolean if a field has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasFsType() bool {
	if o != nil && !IsNil(o.FsType) {
		return true
	}

	return false
}

// SetFsType gets a reference to the given string and assigns it to the FsType field.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetFsType(v string) {
	o.FsType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetLabel(v string) {
	o.Label = &v
}

// GetMinsize returns the Minsize field value if set, zero value otherwise.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMinsize() string {
	if o == nil || IsNil(o.Minsize) {
		var ret string
		return ret
	}
	return *o.Minsize
}

// GetMinsizeOk returns a tuple with the Minsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMinsizeOk() (*string, bool) {
	if o == nil || IsNil(o.Minsize) {
		return nil, false
	}
	return o.Minsize, true
}

// HasMinsize returns a boolean if a field has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasMinsize() bool {
	if o != nil && !IsNil(o.Minsize) {
		return true
	}

	return false
}

// SetMinsize gets a reference to the given string and assigns it to the Minsize field.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetMinsize(v string) {
	o.Minsize = &v
}

// GetMountpoint returns the Mountpoint field value if set, zero value otherwise.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMountpoint() string {
	if o == nil || IsNil(o.Mountpoint) {
		var ret string
		return ret
	}
	return *o.Mountpoint
}

// GetMountpointOk returns a tuple with the Mountpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMountpointOk() (*string, bool) {
	if o == nil || IsNil(o.Mountpoint) {
		return nil, false
	}
	return o.Mountpoint, true
}

// HasMountpoint returns a boolean if a field has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasMountpoint() bool {
	if o != nil && !IsNil(o.Mountpoint) {
		return true
	}

	return false
}

// SetMountpoint gets a reference to the given string and assigns it to the Mountpoint field.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetMountpoint(v string) {
	o.Mountpoint = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetName(v string) {
	o.Name = &v
}

func (o BlueprintStoragePartitionsInnerLogicalVolumesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlueprintStoragePartitionsInnerLogicalVolumesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FsType) {
		toSerialize["fs_type"] = o.FsType
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Minsize) {
		toSerialize["minsize"] = o.Minsize
	}
	if !IsNil(o.Mountpoint) {
		toSerialize["mountpoint"] = o.Mountpoint
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableBlueprintStoragePartitionsInnerLogicalVolumesInner struct {
	value *BlueprintStoragePartitionsInnerLogicalVolumesInner
	isSet bool
}

func (v NullableBlueprintStoragePartitionsInnerLogicalVolumesInner) Get() *BlueprintStoragePartitionsInnerLogicalVolumesInner {
	return v.value
}

func (v *NullableBlueprintStoragePartitionsInnerLogicalVolumesInner) Set(val *BlueprintStoragePartitionsInnerLogicalVolumesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlueprintStoragePartitionsInnerLogicalVolumesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlueprintStoragePartitionsInnerLogicalVolumesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlueprintStoragePartitionsInnerLogicalVolumesInner(val *BlueprintStoragePartitionsInnerLogicalVolumesInner) *NullableBlueprintStoragePartitionsInnerLogicalVolumesInner {
	return &NullableBlueprintStoragePartitionsInnerLogicalVolumesInner{value: val, isSet: true}
}

func (v NullableBlueprintStoragePartitionsInnerLogicalVolumesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlueprintStoragePartitionsInnerLogicalVolumesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


