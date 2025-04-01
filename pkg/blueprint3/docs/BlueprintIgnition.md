# BlueprintIgnition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**NullableBlueprintIgnitionEmbedded**](BlueprintIgnitionEmbedded.md) |  | [optional] 
**FirstbootUrl** | Pointer to **string** | The URL to the Ignition configuration to be used by Ignition. This configuration is a URL to a remote Ignition configuration. The firstboot_url is used if the embedded configuration is not specified.  Cannot be used with embedded_base64 or embedded_text. | [optional] 

## Methods

### NewBlueprintIgnition

`func NewBlueprintIgnition() *BlueprintIgnition`

NewBlueprintIgnition instantiates a new BlueprintIgnition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintIgnitionWithDefaults

`func NewBlueprintIgnitionWithDefaults() *BlueprintIgnition`

NewBlueprintIgnitionWithDefaults instantiates a new BlueprintIgnition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *BlueprintIgnition) GetEmbedded() BlueprintIgnitionEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *BlueprintIgnition) GetEmbeddedOk() (*BlueprintIgnitionEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *BlueprintIgnition) SetEmbedded(v BlueprintIgnitionEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *BlueprintIgnition) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### SetEmbeddedNil

`func (o *BlueprintIgnition) SetEmbeddedNil(b bool)`

 SetEmbeddedNil sets the value for Embedded to be an explicit nil

### UnsetEmbedded
`func (o *BlueprintIgnition) UnsetEmbedded()`

UnsetEmbedded ensures that no value is present for Embedded, not even an explicit nil
### GetFirstbootUrl

`func (o *BlueprintIgnition) GetFirstbootUrl() string`

GetFirstbootUrl returns the FirstbootUrl field if non-nil, zero value otherwise.

### GetFirstbootUrlOk

`func (o *BlueprintIgnition) GetFirstbootUrlOk() (*string, bool)`

GetFirstbootUrlOk returns a tuple with the FirstbootUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstbootUrl

`func (o *BlueprintIgnition) SetFirstbootUrl(v string)`

SetFirstbootUrl sets FirstbootUrl field to given value.

### HasFirstbootUrl

`func (o *BlueprintIgnition) HasFirstbootUrl() bool`

HasFirstbootUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


