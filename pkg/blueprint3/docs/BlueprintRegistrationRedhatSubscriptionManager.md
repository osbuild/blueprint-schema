# BlueprintRegistrationRedhatSubscriptionManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRegistration** | **bool** | Enabled auto_registration plugin configuration. | [default to true]
**Enabled** | **bool** | Enables the subscription-manager DNF plugin. | [default to true]
**ProductPluginEnabled** | **bool** | Enables the product-id DNF plugin. | [default to true]
**RepositoryManagement** | **bool** | Enabled repository_management plugin configuration. | [default to true]

## Methods

### NewBlueprintRegistrationRedhatSubscriptionManager

`func NewBlueprintRegistrationRedhatSubscriptionManager(autoRegistration bool, enabled bool, productPluginEnabled bool, repositoryManagement bool, ) *BlueprintRegistrationRedhatSubscriptionManager`

NewBlueprintRegistrationRedhatSubscriptionManager instantiates a new BlueprintRegistrationRedhatSubscriptionManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintRegistrationRedhatSubscriptionManagerWithDefaults

`func NewBlueprintRegistrationRedhatSubscriptionManagerWithDefaults() *BlueprintRegistrationRedhatSubscriptionManager`

NewBlueprintRegistrationRedhatSubscriptionManagerWithDefaults instantiates a new BlueprintRegistrationRedhatSubscriptionManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRegistration

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetAutoRegistration() bool`

GetAutoRegistration returns the AutoRegistration field if non-nil, zero value otherwise.

### GetAutoRegistrationOk

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetAutoRegistrationOk() (*bool, bool)`

GetAutoRegistrationOk returns a tuple with the AutoRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRegistration

`func (o *BlueprintRegistrationRedhatSubscriptionManager) SetAutoRegistration(v bool)`

SetAutoRegistration sets AutoRegistration field to given value.


### GetEnabled

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlueprintRegistrationRedhatSubscriptionManager) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetProductPluginEnabled

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetProductPluginEnabled() bool`

GetProductPluginEnabled returns the ProductPluginEnabled field if non-nil, zero value otherwise.

### GetProductPluginEnabledOk

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetProductPluginEnabledOk() (*bool, bool)`

GetProductPluginEnabledOk returns a tuple with the ProductPluginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPluginEnabled

`func (o *BlueprintRegistrationRedhatSubscriptionManager) SetProductPluginEnabled(v bool)`

SetProductPluginEnabled sets ProductPluginEnabled field to given value.


### GetRepositoryManagement

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetRepositoryManagement() bool`

GetRepositoryManagement returns the RepositoryManagement field if non-nil, zero value otherwise.

### GetRepositoryManagementOk

`func (o *BlueprintRegistrationRedhatSubscriptionManager) GetRepositoryManagementOk() (*bool, bool)`

GetRepositoryManagementOk returns a tuple with the RepositoryManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryManagement

`func (o *BlueprintRegistrationRedhatSubscriptionManager) SetRepositoryManagement(v bool)`

SetRepositoryManagement sets RepositoryManagement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


