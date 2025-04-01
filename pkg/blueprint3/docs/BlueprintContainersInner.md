# BlueprintContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalStorage** | Pointer to **bool** | Whether to pull the container image from the host&#39;s local-storage. | [optional] [default to false]
**Name** | **string** | Container name is an optional string to set the name under which the container image will be saved in the image. If not specified name falls back to the same value as source. | 
**Source** | **string** | Container image URL is a reference to a container image at a registry. | 
**TlsVerify** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBlueprintContainersInner

`func NewBlueprintContainersInner(name string, source string, ) *BlueprintContainersInner`

NewBlueprintContainersInner instantiates a new BlueprintContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintContainersInnerWithDefaults

`func NewBlueprintContainersInnerWithDefaults() *BlueprintContainersInner`

NewBlueprintContainersInnerWithDefaults instantiates a new BlueprintContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalStorage

`func (o *BlueprintContainersInner) GetLocalStorage() bool`

GetLocalStorage returns the LocalStorage field if non-nil, zero value otherwise.

### GetLocalStorageOk

`func (o *BlueprintContainersInner) GetLocalStorageOk() (*bool, bool)`

GetLocalStorageOk returns a tuple with the LocalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStorage

`func (o *BlueprintContainersInner) SetLocalStorage(v bool)`

SetLocalStorage sets LocalStorage field to given value.

### HasLocalStorage

`func (o *BlueprintContainersInner) HasLocalStorage() bool`

HasLocalStorage returns a boolean if a field has been set.

### GetName

`func (o *BlueprintContainersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintContainersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintContainersInner) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *BlueprintContainersInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BlueprintContainersInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BlueprintContainersInner) SetSource(v string)`

SetSource sets Source field to given value.


### GetTlsVerify

`func (o *BlueprintContainersInner) GetTlsVerify() bool`

GetTlsVerify returns the TlsVerify field if non-nil, zero value otherwise.

### GetTlsVerifyOk

`func (o *BlueprintContainersInner) GetTlsVerifyOk() (*bool, bool)`

GetTlsVerifyOk returns a tuple with the TlsVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsVerify

`func (o *BlueprintContainersInner) SetTlsVerify(v bool)`

SetTlsVerify sets TlsVerify field to given value.

### HasTlsVerify

`func (o *BlueprintContainersInner) HasTlsVerify() bool`

HasTlsVerify returns a boolean if a field has been set.

### SetTlsVerifyNil

`func (o *BlueprintContainersInner) SetTlsVerifyNil(b bool)`

 SetTlsVerifyNil sets the value for TlsVerify to be an explicit nil

### UnsetTlsVerify
`func (o *BlueprintContainersInner) UnsetTlsVerify()`

UnsetTlsVerify ensures that no value is present for TlsVerify, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


