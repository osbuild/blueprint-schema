# BlueprintStoragePartitionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Partition type: plain (default), lvm, or btrfs. | [default to "plain"]
**FsType** | Pointer to **string** | File system type: ext4 (default), xfs, swap, or vfat.  Relevant for partition types: plain. | [optional] [default to "ext4"]
**Label** | Pointer to **string** | Label of the partition.  Relevant for partition types: plain. | [optional] 
**LogicalVolumes** | Pointer to [**[]BlueprintStoragePartitionsInnerLogicalVolumesInner**](BlueprintStoragePartitionsInnerLogicalVolumesInner.md) | LVM logical volumes to create within the volume group.  Relevant for partition types: lvm. | [optional] 
**Minsize** | Pointer to **string** | Minimum size of the volume.  Size must be formatted as an integer followed by whitespace and then either a decimal unit (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB).  Relevant for partition types: plain, lvm, btrfs. | [optional] 
**Mountpoint** | Pointer to **string** | Mount point of the partition. Required except for swap fs_type.  Relevant for partition types: plain. | [optional] 
**Name** | Pointer to **string** | LVM volume group name. When not set, will be generated automatically.  Relevant for partition types: lvm. | [optional] 
**Subvolumes** | Pointer to [**[]BlueprintStoragePartitionsInnerSubvolumesInner**](BlueprintStoragePartitionsInnerSubvolumesInner.md) | BTRFS subvolumes to create.  Relevant for partition types: btrfs. | [optional] 

## Methods

### NewBlueprintStoragePartitionsInner

`func NewBlueprintStoragePartitionsInner(type_ string, ) *BlueprintStoragePartitionsInner`

NewBlueprintStoragePartitionsInner instantiates a new BlueprintStoragePartitionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintStoragePartitionsInnerWithDefaults

`func NewBlueprintStoragePartitionsInnerWithDefaults() *BlueprintStoragePartitionsInner`

NewBlueprintStoragePartitionsInnerWithDefaults instantiates a new BlueprintStoragePartitionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintStoragePartitionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintStoragePartitionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintStoragePartitionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetFsType

`func (o *BlueprintStoragePartitionsInner) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *BlueprintStoragePartitionsInner) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *BlueprintStoragePartitionsInner) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *BlueprintStoragePartitionsInner) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetLabel

`func (o *BlueprintStoragePartitionsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BlueprintStoragePartitionsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BlueprintStoragePartitionsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BlueprintStoragePartitionsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLogicalVolumes

`func (o *BlueprintStoragePartitionsInner) GetLogicalVolumes() []BlueprintStoragePartitionsInnerLogicalVolumesInner`

GetLogicalVolumes returns the LogicalVolumes field if non-nil, zero value otherwise.

### GetLogicalVolumesOk

`func (o *BlueprintStoragePartitionsInner) GetLogicalVolumesOk() (*[]BlueprintStoragePartitionsInnerLogicalVolumesInner, bool)`

GetLogicalVolumesOk returns a tuple with the LogicalVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalVolumes

`func (o *BlueprintStoragePartitionsInner) SetLogicalVolumes(v []BlueprintStoragePartitionsInnerLogicalVolumesInner)`

SetLogicalVolumes sets LogicalVolumes field to given value.

### HasLogicalVolumes

`func (o *BlueprintStoragePartitionsInner) HasLogicalVolumes() bool`

HasLogicalVolumes returns a boolean if a field has been set.

### GetMinsize

`func (o *BlueprintStoragePartitionsInner) GetMinsize() string`

GetMinsize returns the Minsize field if non-nil, zero value otherwise.

### GetMinsizeOk

`func (o *BlueprintStoragePartitionsInner) GetMinsizeOk() (*string, bool)`

GetMinsizeOk returns a tuple with the Minsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinsize

`func (o *BlueprintStoragePartitionsInner) SetMinsize(v string)`

SetMinsize sets Minsize field to given value.

### HasMinsize

`func (o *BlueprintStoragePartitionsInner) HasMinsize() bool`

HasMinsize returns a boolean if a field has been set.

### GetMountpoint

`func (o *BlueprintStoragePartitionsInner) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *BlueprintStoragePartitionsInner) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *BlueprintStoragePartitionsInner) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *BlueprintStoragePartitionsInner) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetName

`func (o *BlueprintStoragePartitionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintStoragePartitionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintStoragePartitionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintStoragePartitionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubvolumes

`func (o *BlueprintStoragePartitionsInner) GetSubvolumes() []BlueprintStoragePartitionsInnerSubvolumesInner`

GetSubvolumes returns the Subvolumes field if non-nil, zero value otherwise.

### GetSubvolumesOk

`func (o *BlueprintStoragePartitionsInner) GetSubvolumesOk() (*[]BlueprintStoragePartitionsInnerSubvolumesInner, bool)`

GetSubvolumesOk returns a tuple with the Subvolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubvolumes

`func (o *BlueprintStoragePartitionsInner) SetSubvolumes(v []BlueprintStoragePartitionsInnerSubvolumesInner)`

SetSubvolumes sets Subvolumes field to given value.

### HasSubvolumes

`func (o *BlueprintStoragePartitionsInner) HasSubvolumes() bool`

HasSubvolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


