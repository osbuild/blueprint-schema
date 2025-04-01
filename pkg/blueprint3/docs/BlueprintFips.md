# BlueprintFips

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enables the system FIPS mode (disabled by default). Currently only edge-raw-image, edge-installer, edge-simplified-installer, edge-ami and edge-vsphere images support this customization. | [optional] 

## Methods

### NewBlueprintFips

`func NewBlueprintFips() *BlueprintFips`

NewBlueprintFips instantiates a new BlueprintFips object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintFipsWithDefaults

`func NewBlueprintFipsWithDefaults() *BlueprintFips`

NewBlueprintFipsWithDefaults instantiates a new BlueprintFips object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BlueprintFips) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintFips) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintFips) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BlueprintFips) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


