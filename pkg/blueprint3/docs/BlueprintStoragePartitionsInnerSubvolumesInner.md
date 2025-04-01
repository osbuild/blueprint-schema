# BlueprintStoragePartitionsInnerSubvolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mountpoint** | Pointer to **string** | Mount point of the subvolume. Required. Swap filesystem type is not supported on BTRFS volumes. | [optional] 
**Name** | Pointer to **string** | Subvolume name, must also define its parent volume. | [optional] 

## Methods

### NewBlueprintStoragePartitionsInnerSubvolumesInner

`func NewBlueprintStoragePartitionsInnerSubvolumesInner() *BlueprintStoragePartitionsInnerSubvolumesInner`

NewBlueprintStoragePartitionsInnerSubvolumesInner instantiates a new BlueprintStoragePartitionsInnerSubvolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintStoragePartitionsInnerSubvolumesInnerWithDefaults

`func NewBlueprintStoragePartitionsInnerSubvolumesInnerWithDefaults() *BlueprintStoragePartitionsInnerSubvolumesInner`

NewBlueprintStoragePartitionsInnerSubvolumesInnerWithDefaults instantiates a new BlueprintStoragePartitionsInnerSubvolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMountpoint

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) GetMountpoint() string`

GetMountpoint returns the Mountpoint field if non-nil, zero value otherwise.

### GetMountpointOk

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) GetMountpointOk() (*string, bool)`

GetMountpointOk returns a tuple with the Mountpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoint

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) SetMountpoint(v string)`

SetMountpoint sets Mountpoint field to given value.

### HasMountpoint

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) HasMountpoint() bool`

HasMountpoint returns a boolean if a field has been set.

### GetName

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintStoragePartitionsInnerSubvolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


