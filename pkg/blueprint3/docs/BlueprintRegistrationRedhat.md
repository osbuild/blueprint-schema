# BlueprintRegistrationRedhat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationKey** | Pointer to **string** | Subscription manager activation key to use during registration. A list of keys to use to redeem or apply specific subscriptions to the system. | [optional] 
**Connector** | Pointer to [**BlueprintRegistrationRedhatConnector**](BlueprintRegistrationRedhatConnector.md) |  | [optional] 
**Insights** | Pointer to [**BlueprintRegistrationRedhatInsights**](BlueprintRegistrationRedhatInsights.md) |  | [optional] 
**Organization** | Pointer to **string** | Subscription manager organization name to use during registration. | [optional] 
**SubscriptionManager** | Pointer to [**BlueprintRegistrationRedhatSubscriptionManager**](BlueprintRegistrationRedhatSubscriptionManager.md) |  | [optional] 

## Methods

### NewBlueprintRegistrationRedhat

`func NewBlueprintRegistrationRedhat() *BlueprintRegistrationRedhat`

NewBlueprintRegistrationRedhat instantiates a new BlueprintRegistrationRedhat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintRegistrationRedhatWithDefaults

`func NewBlueprintRegistrationRedhatWithDefaults() *BlueprintRegistrationRedhat`

NewBlueprintRegistrationRedhatWithDefaults instantiates a new BlueprintRegistrationRedhat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationKey

`func (o *BlueprintRegistrationRedhat) GetActivationKey() string`

GetActivationKey returns the ActivationKey field if non-nil, zero value otherwise.

### GetActivationKeyOk

`func (o *BlueprintRegistrationRedhat) GetActivationKeyOk() (*string, bool)`

GetActivationKeyOk returns a tuple with the ActivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationKey

`func (o *BlueprintRegistrationRedhat) SetActivationKey(v string)`

SetActivationKey sets ActivationKey field to given value.

### HasActivationKey

`func (o *BlueprintRegistrationRedhat) HasActivationKey() bool`

HasActivationKey returns a boolean if a field has been set.

### GetConnector

`func (o *BlueprintRegistrationRedhat) GetConnector() BlueprintRegistrationRedhatConnector`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *BlueprintRegistrationRedhat) GetConnectorOk() (*BlueprintRegistrationRedhatConnector, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *BlueprintRegistrationRedhat) SetConnector(v BlueprintRegistrationRedhatConnector)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *BlueprintRegistrationRedhat) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetInsights

`func (o *BlueprintRegistrationRedhat) GetInsights() BlueprintRegistrationRedhatInsights`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *BlueprintRegistrationRedhat) GetInsightsOk() (*BlueprintRegistrationRedhatInsights, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *BlueprintRegistrationRedhat) SetInsights(v BlueprintRegistrationRedhatInsights)`

SetInsights sets Insights field to given value.

### HasInsights

`func (o *BlueprintRegistrationRedhat) HasInsights() bool`

HasInsights returns a boolean if a field has been set.

### GetOrganization

`func (o *BlueprintRegistrationRedhat) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *BlueprintRegistrationRedhat) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *BlueprintRegistrationRedhat) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *BlueprintRegistrationRedhat) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSubscriptionManager

`func (o *BlueprintRegistrationRedhat) GetSubscriptionManager() BlueprintRegistrationRedhatSubscriptionManager`

GetSubscriptionManager returns the SubscriptionManager field if non-nil, zero value otherwise.

### GetSubscriptionManagerOk

`func (o *BlueprintRegistrationRedhat) GetSubscriptionManagerOk() (*BlueprintRegistrationRedhatSubscriptionManager, bool)`

GetSubscriptionManagerOk returns a tuple with the SubscriptionManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionManager

`func (o *BlueprintRegistrationRedhat) SetSubscriptionManager(v BlueprintRegistrationRedhatSubscriptionManager)`

SetSubscriptionManager sets SubscriptionManager field to given value.

### HasSubscriptionManager

`func (o *BlueprintRegistrationRedhat) HasSubscriptionManager() bool`

HasSubscriptionManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


