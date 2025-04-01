# BlueprintStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Device partitioning type: gpt (default) or mbr. | [default to "gpt"]
**Minsize** | **string** | Minimum size of the storage device. When not set, the image size is acquired from image request.  Size must be formatted as an integer followed by whitespace and then either a decimal unit (B, KB/kB, MB, GB, TB, PB, EB) or binary unit (KiB, MiB, GiB, TiB, PiB, EiB). | 
**Partitions** | [**[]BlueprintStoragePartitionsInner**](BlueprintStoragePartitionsInner.md) | Partitions of the following types: plain (default), lvm, or btrfs. | 

## Methods

### NewBlueprintStorage

`func NewBlueprintStorage(type_ string, minsize string, partitions []BlueprintStoragePartitionsInner, ) *BlueprintStorage`

NewBlueprintStorage instantiates a new BlueprintStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintStorageWithDefaults

`func NewBlueprintStorageWithDefaults() *BlueprintStorage`

NewBlueprintStorageWithDefaults instantiates a new BlueprintStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlueprintStorage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintStorage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintStorage) SetType(v string)`

SetType sets Type field to given value.


### GetMinsize

`func (o *BlueprintStorage) GetMinsize() string`

GetMinsize returns the Minsize field if non-nil, zero value otherwise.

### GetMinsizeOk

`func (o *BlueprintStorage) GetMinsizeOk() (*string, bool)`

GetMinsizeOk returns a tuple with the Minsize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinsize

`func (o *BlueprintStorage) SetMinsize(v string)`

SetMinsize sets Minsize field to given value.


### GetPartitions

`func (o *BlueprintStorage) GetPartitions() []BlueprintStoragePartitionsInner`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *BlueprintStorage) GetPartitionsOk() (*[]BlueprintStoragePartitionsInner, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *BlueprintStorage) SetPartitions(v []BlueprintStoragePartitionsInner)`

SetPartitions sets Partitions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


