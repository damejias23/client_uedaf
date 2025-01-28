# UedafEventMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**UedafEventTrigger**](UedafEventTrigger.md) |  | 
**MaxReports** | Pointer to **int32** |  | [optional] 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RepPeriod** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**SampRatio** | Pointer to **int32** | Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.   | [optional] 

## Methods

### NewUedafEventMode

`func NewUedafEventMode(trigger UedafEventTrigger, ) *UedafEventMode`

NewUedafEventMode instantiates a new UedafEventMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafEventModeWithDefaults

`func NewUedafEventModeWithDefaults() *UedafEventMode`

NewUedafEventModeWithDefaults instantiates a new UedafEventMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *UedafEventMode) GetTrigger() UedafEventTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *UedafEventMode) GetTriggerOk() (*UedafEventTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *UedafEventMode) SetTrigger(v UedafEventTrigger)`

SetTrigger sets Trigger field to given value.


### GetMaxReports

`func (o *UedafEventMode) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *UedafEventMode) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *UedafEventMode) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *UedafEventMode) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetExpiry

`func (o *UedafEventMode) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *UedafEventMode) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *UedafEventMode) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *UedafEventMode) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetRepPeriod

`func (o *UedafEventMode) GetRepPeriod() int32`

GetRepPeriod returns the RepPeriod field if non-nil, zero value otherwise.

### GetRepPeriodOk

`func (o *UedafEventMode) GetRepPeriodOk() (*int32, bool)`

GetRepPeriodOk returns a tuple with the RepPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepPeriod

`func (o *UedafEventMode) SetRepPeriod(v int32)`

SetRepPeriod sets RepPeriod field to given value.

### HasRepPeriod

`func (o *UedafEventMode) HasRepPeriod() bool`

HasRepPeriod returns a boolean if a field has been set.

### GetSampRatio

`func (o *UedafEventMode) GetSampRatio() int32`

GetSampRatio returns the SampRatio field if non-nil, zero value otherwise.

### GetSampRatioOk

`func (o *UedafEventMode) GetSampRatioOk() (*int32, bool)`

GetSampRatioOk returns a tuple with the SampRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampRatio

`func (o *UedafEventMode) SetSampRatio(v int32)`

SetSampRatio sets SampRatio field to given value.

### HasSampRatio

`func (o *UedafEventMode) HasSampRatio() bool`

HasSampRatio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


