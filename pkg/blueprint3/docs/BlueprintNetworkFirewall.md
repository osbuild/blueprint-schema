# BlueprintNetworkFirewall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]BlueprintNetworkFirewallServicesInner**](BlueprintNetworkFirewallServicesInner.md) | Services to enable or disable. The service can be defined via an assigned IANA name, port number or port range.  Services are processed in order, when a service is disabled and then accidentally enabled via a port or a port range, the service will be enabled in the end.  By default the firewall blocks all access, except for services that enable their ports explicitly such as the sshd. | [optional] 

## Methods

### NewBlueprintNetworkFirewall

`func NewBlueprintNetworkFirewall() *BlueprintNetworkFirewall`

NewBlueprintNetworkFirewall instantiates a new BlueprintNetworkFirewall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintNetworkFirewallWithDefaults

`func NewBlueprintNetworkFirewallWithDefaults() *BlueprintNetworkFirewall`

NewBlueprintNetworkFirewallWithDefaults instantiates a new BlueprintNetworkFirewall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *BlueprintNetworkFirewall) GetServices() []BlueprintNetworkFirewallServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *BlueprintNetworkFirewall) GetServicesOk() (*[]BlueprintNetworkFirewallServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *BlueprintNetworkFirewall) SetServices(v []BlueprintNetworkFirewallServicesInner)`

SetServices sets Services field to given value.

### HasServices

`func (o *BlueprintNetworkFirewall) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *BlueprintNetworkFirewall) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *BlueprintNetworkFirewall) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


