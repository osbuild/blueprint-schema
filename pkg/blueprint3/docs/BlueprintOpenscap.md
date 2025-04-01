# BlueprintOpenscap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastream** | Pointer to **string** | Datastream to use for the scan. The datastream is the path to the SCAP datastream file to use for the scan. If the datastream parameter is not provided, a sensible default based on the selected distro will be chosen. | [optional] 
**ProfileId** | **string** | The desired security profile ID. | 
**Tailoring** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 

## Methods

### NewBlueprintOpenscap

`func NewBlueprintOpenscap(profileId string, ) *BlueprintOpenscap`

NewBlueprintOpenscap instantiates a new BlueprintOpenscap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintOpenscapWithDefaults

`func NewBlueprintOpenscapWithDefaults() *BlueprintOpenscap`

NewBlueprintOpenscapWithDefaults instantiates a new BlueprintOpenscap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastream

`func (o *BlueprintOpenscap) GetDatastream() string`

GetDatastream returns the Datastream field if non-nil, zero value otherwise.

### GetDatastreamOk

`func (o *BlueprintOpenscap) GetDatastreamOk() (*string, bool)`

GetDatastreamOk returns a tuple with the Datastream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastream

`func (o *BlueprintOpenscap) SetDatastream(v string)`

SetDatastream sets Datastream field to given value.

### HasDatastream

`func (o *BlueprintOpenscap) HasDatastream() bool`

HasDatastream returns a boolean if a field has been set.

### GetProfileId

`func (o *BlueprintOpenscap) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *BlueprintOpenscap) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *BlueprintOpenscap) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetTailoring

`func (o *BlueprintOpenscap) GetTailoring() AnyOf`

GetTailoring returns the Tailoring field if non-nil, zero value otherwise.

### GetTailoringOk

`func (o *BlueprintOpenscap) GetTailoringOk() (*AnyOf, bool)`

GetTailoringOk returns a tuple with the Tailoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTailoring

`func (o *BlueprintOpenscap) SetTailoring(v AnyOf)`

SetTailoring sets Tailoring field to given value.

### HasTailoring

`func (o *BlueprintOpenscap) HasTailoring() bool`

HasTailoring returns a boolean if a field has been set.

### SetTailoringNil

`func (o *BlueprintOpenscap) SetTailoringNil(b bool)`

 SetTailoringNil sets the value for Tailoring to be an explicit nil

### UnsetTailoring
`func (o *BlueprintOpenscap) UnsetTailoring()`

UnsetTailoring ensures that no value is present for Tailoring, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


