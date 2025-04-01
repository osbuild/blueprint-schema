# BlueprintLocale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyboards** | Pointer to **[]string** | The keyboards attribute is a list of strings that contains the keyboards to be installed on the image. To list available keyboards, run: localectl list-keymaps | [optional] [default to ["us"]]
**Languages** | Pointer to **[]string** | The languages attribute is a list of strings that contains the languages to be installed on the image. To list available languages, run: localectl list-locales | [optional] [default to ["en_US.UTF-8"]]

## Methods

### NewBlueprintLocale

`func NewBlueprintLocale() *BlueprintLocale`

NewBlueprintLocale instantiates a new BlueprintLocale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintLocaleWithDefaults

`func NewBlueprintLocaleWithDefaults() *BlueprintLocale`

NewBlueprintLocaleWithDefaults instantiates a new BlueprintLocale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyboards

`func (o *BlueprintLocale) GetKeyboards() []string`

GetKeyboards returns the Keyboards field if non-nil, zero value otherwise.

### GetKeyboardsOk

`func (o *BlueprintLocale) GetKeyboardsOk() (*[]string, bool)`

GetKeyboardsOk returns a tuple with the Keyboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboards

`func (o *BlueprintLocale) SetKeyboards(v []string)`

SetKeyboards sets Keyboards field to given value.

### HasKeyboards

`func (o *BlueprintLocale) HasKeyboards() bool`

HasKeyboards returns a boolean if a field has been set.

### SetKeyboardsNil

`func (o *BlueprintLocale) SetKeyboardsNil(b bool)`

 SetKeyboardsNil sets the value for Keyboards to be an explicit nil

### UnsetKeyboards
`func (o *BlueprintLocale) UnsetKeyboards()`

UnsetKeyboards ensures that no value is present for Keyboards, not even an explicit nil
### GetLanguages

`func (o *BlueprintLocale) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *BlueprintLocale) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *BlueprintLocale) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *BlueprintLocale) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *BlueprintLocale) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *BlueprintLocale) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


