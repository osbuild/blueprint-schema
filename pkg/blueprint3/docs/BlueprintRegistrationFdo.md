# BlueprintRegistrationFdo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiMfgStringTypeMacIface** | Pointer to **string** | Optional interface name for the MAC address. | [optional] 
**DiunPubKeyHash** | Pointer to **string** | FDO server public key hex-encoded hash. Cannot be used together with insecure option or root certs. | [optional] 
**DiunPubKeyInsecure** | Pointer to **bool** | FDO insecure option. When set, both hash or root certs must not be set. | [optional] [default to false]
**DiunPubKeyRootCerts** | Pointer to **string** | FDO server public key root certificate path. Cannot be used together with insecure option or hash. | [optional] 
**ManufacturingServerUrl** | **string** | FDO manufacturing server URL. | 

## Methods

### NewBlueprintRegistrationFdo

`func NewBlueprintRegistrationFdo(manufacturingServerUrl string, ) *BlueprintRegistrationFdo`

NewBlueprintRegistrationFdo instantiates a new BlueprintRegistrationFdo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintRegistrationFdoWithDefaults

`func NewBlueprintRegistrationFdoWithDefaults() *BlueprintRegistrationFdo`

NewBlueprintRegistrationFdoWithDefaults instantiates a new BlueprintRegistrationFdo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiMfgStringTypeMacIface

`func (o *BlueprintRegistrationFdo) GetDiMfgStringTypeMacIface() string`

GetDiMfgStringTypeMacIface returns the DiMfgStringTypeMacIface field if non-nil, zero value otherwise.

### GetDiMfgStringTypeMacIfaceOk

`func (o *BlueprintRegistrationFdo) GetDiMfgStringTypeMacIfaceOk() (*string, bool)`

GetDiMfgStringTypeMacIfaceOk returns a tuple with the DiMfgStringTypeMacIface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiMfgStringTypeMacIface

`func (o *BlueprintRegistrationFdo) SetDiMfgStringTypeMacIface(v string)`

SetDiMfgStringTypeMacIface sets DiMfgStringTypeMacIface field to given value.

### HasDiMfgStringTypeMacIface

`func (o *BlueprintRegistrationFdo) HasDiMfgStringTypeMacIface() bool`

HasDiMfgStringTypeMacIface returns a boolean if a field has been set.

### GetDiunPubKeyHash

`func (o *BlueprintRegistrationFdo) GetDiunPubKeyHash() string`

GetDiunPubKeyHash returns the DiunPubKeyHash field if non-nil, zero value otherwise.

### GetDiunPubKeyHashOk

`func (o *BlueprintRegistrationFdo) GetDiunPubKeyHashOk() (*string, bool)`

GetDiunPubKeyHashOk returns a tuple with the DiunPubKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiunPubKeyHash

`func (o *BlueprintRegistrationFdo) SetDiunPubKeyHash(v string)`

SetDiunPubKeyHash sets DiunPubKeyHash field to given value.

### HasDiunPubKeyHash

`func (o *BlueprintRegistrationFdo) HasDiunPubKeyHash() bool`

HasDiunPubKeyHash returns a boolean if a field has been set.

### GetDiunPubKeyInsecure

`func (o *BlueprintRegistrationFdo) GetDiunPubKeyInsecure() bool`

GetDiunPubKeyInsecure returns the DiunPubKeyInsecure field if non-nil, zero value otherwise.

### GetDiunPubKeyInsecureOk

`func (o *BlueprintRegistrationFdo) GetDiunPubKeyInsecureOk() (*bool, bool)`

GetDiunPubKeyInsecureOk returns a tuple with the DiunPubKeyInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiunPubKeyInsecure

`func (o *BlueprintRegistrationFdo) SetDiunPubKeyInsecure(v bool)`

SetDiunPubKeyInsecure sets DiunPubKeyInsecure field to given value.

### HasDiunPubKeyInsecure

`func (o *BlueprintRegistrationFdo) HasDiunPubKeyInsecure() bool`

HasDiunPubKeyInsecure returns a boolean if a field has been set.

### GetDiunPubKeyRootCerts

`func (o *BlueprintRegistrationFdo) GetDiunPubKeyRootCerts() string`

GetDiunPubKeyRootCerts returns the DiunPubKeyRootCerts field if non-nil, zero value otherwise.

### GetDiunPubKeyRootCertsOk

`func (o *BlueprintRegistrationFdo) GetDiunPubKeyRootCertsOk() (*string, bool)`

GetDiunPubKeyRootCertsOk returns a tuple with the DiunPubKeyRootCerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiunPubKeyRootCerts

`func (o *BlueprintRegistrationFdo) SetDiunPubKeyRootCerts(v string)`

SetDiunPubKeyRootCerts sets DiunPubKeyRootCerts field to given value.

### HasDiunPubKeyRootCerts

`func (o *BlueprintRegistrationFdo) HasDiunPubKeyRootCerts() bool`

HasDiunPubKeyRootCerts returns a boolean if a field has been set.

### GetManufacturingServerUrl

`func (o *BlueprintRegistrationFdo) GetManufacturingServerUrl() string`

GetManufacturingServerUrl returns the ManufacturingServerUrl field if non-nil, zero value otherwise.

### GetManufacturingServerUrlOk

`func (o *BlueprintRegistrationFdo) GetManufacturingServerUrlOk() (*string, bool)`

GetManufacturingServerUrlOk returns a tuple with the ManufacturingServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturingServerUrl

`func (o *BlueprintRegistrationFdo) SetManufacturingServerUrl(v string)`

SetManufacturingServerUrl sets ManufacturingServerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


