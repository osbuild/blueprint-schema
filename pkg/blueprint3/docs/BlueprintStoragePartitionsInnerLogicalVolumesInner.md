# BlueprintStoragePartitionsInnerLogicalVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FsType** | Pointer to **string** | File system type: ext4 (default), xfs, swap, or vfat. | [optional] [default to "ext4"]
**Label** | Pointer to **string** | Label of the logical volume. | [optional] 
**Minsize** | Pointer to **string** | Minimum size of the logical volume.  Size must be formatted as an integer followed by whitespace and then either a decimal unit (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB). | [optional] 
**Mountpoint** | Pointer to **string** | Mount point of the logical volume. Required except for swap fs_type. | [optional] 
**Name** | Pointer to **string** | Logical volume name. When not set, will be generated automatically. | [optional] 

## Methods

### NewBlueprintStoragePartitionsInnerLogicalVolumesInner

`func NewBlueprintStoragePartitionsInnerLogicalVolumesInner() *BlueprintStoragePartitionsInnerLogicalVolumesInner`

NewBlueprintStoragePartitionsInnerLogicalVolumesInner instantiates a new BlueprintStoragePartitionsInnerLogicalVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintStoragePartitionsInnerLogicalVolumesInnerWithDefaults

`func NewBlueprintStoragePartitionsInnerLogicalVolumesInnerWithDefaults() *BlueprintStoragePartitionsInnerLogicalVolumesInner`

NewBlueprintStoragePartitionsInnerLogicalVolumesInnerWithDefaults instantiates a new BlueprintStoragePartitionsInnerLogicalVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFsType

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetFsType() string`

GetFsType returns the FsType field if non-nil, zero value otherwise.

### GetFsTypeOk

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetFsTypeOk() (*string, bool)`

GetFsTypeOk returns a tuple with the FsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFsType

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetFsType(v string)`

SetFsType sets FsType field to given value.

### HasFsType

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasFsType() bool`

HasFsType returns a boolean if a field has been set.

### GetLabel

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMinsize

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMinsize() string`

GetMinsize returns the Minsize field if non-nil, zero value otherwise.

### GetMinsizeOk

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMinsizeOk() (*string, bool)`

GetMinsizeOk returns a tuple with the Minsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinsize

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetMinsize(v string)`

SetMinsize sets Minsize field to given value.

### HasMinsize

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasMinsize() bool`

HasMinsize returns a boolean if a field has been set.

### GetMountpoint

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetName

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintStoragePartitionsInnerLogicalVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


