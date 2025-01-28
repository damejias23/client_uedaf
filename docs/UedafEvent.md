# UedafEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**UedafEventType**](UedafEventType.md) |  | 
**MaxReports** | Pointer to **int32** |  | [optional] 
**MaxResponseTime** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**MinInterval** | Pointer to **int32** | indicating a time in seconds. | [optional] 
**NextReport** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NextPeriodicReportTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewUedafEvent

`func NewUedafEvent(type_ UedafEventType, ) *UedafEvent`

NewUedafEvent instantiates a new UedafEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafEventWithDefaults

`func NewUedafEventWithDefaults() *UedafEvent`

NewUedafEventWithDefaults instantiates a new UedafEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UedafEvent) GetType() UedafEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UedafEvent) GetTypeOk() (*UedafEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UedafEvent) SetType(v UedafEventType)`

SetType sets Type field to given value.


### GetMaxReports

`func (o *UedafEvent) GetMaxReports() int32`

GetMaxReports returns the MaxReports field if non-nil, zero value otherwise.

### GetMaxReportsOk

`func (o *UedafEvent) GetMaxReportsOk() (*int32, bool)`

GetMaxReportsOk returns a tuple with the MaxReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReports

`func (o *UedafEvent) SetMaxReports(v int32)`

SetMaxReports sets MaxReports field to given value.

### HasMaxReports

`func (o *UedafEvent) HasMaxReports() bool`

HasMaxReports returns a boolean if a field has been set.

### GetMaxResponseTime

`func (o *UedafEvent) GetMaxResponseTime() int32`

GetMaxResponseTime returns the MaxResponseTime field if non-nil, zero value otherwise.

### GetMaxResponseTimeOk

`func (o *UedafEvent) GetMaxResponseTimeOk() (*int32, bool)`

GetMaxResponseTimeOk returns a tuple with the MaxResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseTime

`func (o *UedafEvent) SetMaxResponseTime(v int32)`

SetMaxResponseTime sets MaxResponseTime field to given value.

### HasMaxResponseTime

`func (o *UedafEvent) HasMaxResponseTime() bool`

HasMaxResponseTime returns a boolean if a field has been set.

### GetMinInterval

`func (o *UedafEvent) GetMinInterval() int32`

GetMinInterval returns the MinInterval field if non-nil, zero value otherwise.

### GetMinIntervalOk

`func (o *UedafEvent) GetMinIntervalOk() (*int32, bool)`

GetMinIntervalOk returns a tuple with the MinInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInterval

`func (o *UedafEvent) SetMinInterval(v int32)`

SetMinInterval sets MinInterval field to given value.

### HasMinInterval

`func (o *UedafEvent) HasMinInterval() bool`

HasMinInterval returns a boolean if a field has been set.

### GetNextReport

`func (o *UedafEvent) GetNextReport() time.Time`

GetNextReport returns the NextReport field if non-nil, zero value otherwise.

### GetNextReportOk

`func (o *UedafEvent) GetNextReportOk() (*time.Time, bool)`

GetNextReportOk returns a tuple with the NextReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextReport

`func (o *UedafEvent) SetNextReport(v time.Time)`

SetNextReport sets NextReport field to given value.

### HasNextReport

`func (o *UedafEvent) HasNextReport() bool`

HasNextReport returns a boolean if a field has been set.

### GetNextPeriodicReportTime

`func (o *UedafEvent) GetNextPeriodicReportTime() time.Time`

GetNextPeriodicReportTime returns the NextPeriodicReportTime field if non-nil, zero value otherwise.

### GetNextPeriodicReportTimeOk

`func (o *UedafEvent) GetNextPeriodicReportTimeOk() (*time.Time, bool)`

GetNextPeriodicReportTimeOk returns a tuple with the NextPeriodicReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPeriodicReportTime

`func (o *UedafEvent) SetNextPeriodicReportTime(v time.Time)`

SetNextPeriodicReportTime sets NextPeriodicReportTime field to given value.

### HasNextPeriodicReportTime

`func (o *UedafEvent) HasNextPeriodicReportTime() bool`

HasNextPeriodicReportTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


