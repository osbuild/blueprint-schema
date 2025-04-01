# BlueprintNetworkFirewallServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**From** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] [default to "any"]
**Service** | Pointer to **string** |  | [optional] 
**To** | Pointer to **int32** |  | [optional] 

## Methods

### NewBlueprintNetworkFirewallServicesInner

`func NewBlueprintNetworkFirewallServicesInner() *BlueprintNetworkFirewallServicesInner`

NewBlueprintNetworkFirewallServicesInner instantiates a new BlueprintNetworkFirewallServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintNetworkFirewallServicesInnerWithDefaults

`func NewBlueprintNetworkFirewallServicesInnerWithDefaults() *BlueprintNetworkFirewallServicesInner`

NewBlueprintNetworkFirewallServicesInnerWithDefaults instantiates a new BlueprintNetworkFirewallServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BlueprintNetworkFirewallServicesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintNetworkFirewallServicesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintNetworkFirewallServicesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BlueprintNetworkFirewallServicesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *BlueprintNetworkFirewallServicesInner) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *BlueprintNetworkFirewallServicesInner) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFrom

`func (o *BlueprintNetworkFirewallServicesInner) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *BlueprintNetworkFirewallServicesInner) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *BlueprintNetworkFirewallServicesInner) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *BlueprintNetworkFirewallServicesInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetPort

`func (o *BlueprintNetworkFirewallServicesInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BlueprintNetworkFirewallServicesInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BlueprintNetworkFirewallServicesInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BlueprintNetworkFirewallServicesInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *BlueprintNetworkFirewallServicesInner) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BlueprintNetworkFirewallServicesInner) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BlueprintNetworkFirewallServicesInner) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BlueprintNetworkFirewallServicesInner) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetService

`func (o *BlueprintNetworkFirewallServicesInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BlueprintNetworkFirewallServicesInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BlueprintNetworkFirewallServicesInner) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *BlueprintNetworkFirewallServicesInner) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTo

`func (o *BlueprintNetworkFirewallServicesInner) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *BlueprintNetworkFirewallServicesInner) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *BlueprintNetworkFirewallServicesInner) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *BlueprintNetworkFirewallServicesInner) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


