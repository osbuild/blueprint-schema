# BlueprintTimedate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NtpServers** | Pointer to **[]string** | An optional list of strings containing NTP servers to use. If not provided the distribution defaults are used | [optional] 
**Timezone** | **string** | System time zone. Defaults to UTC. To list available time zones run: timedatectl list-timezones | [default to "UTC"]

## Methods

### NewBlueprintTimedate

`func NewBlueprintTimedate(timezone string, ) *BlueprintTimedate`

NewBlueprintTimedate instantiates a new BlueprintTimedate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintTimedateWithDefaults

`func NewBlueprintTimedateWithDefaults() *BlueprintTimedate`

NewBlueprintTimedateWithDefaults instantiates a new BlueprintTimedate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNtpServers

`func (o *BlueprintTimedate) GetNtpServers() []string`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *BlueprintTimedate) GetNtpServersOk() (*[]string, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *BlueprintTimedate) SetNtpServers(v []string)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *BlueprintTimedate) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetTimezone

`func (o *BlueprintTimedate) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BlueprintTimedate) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BlueprintTimedate) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


