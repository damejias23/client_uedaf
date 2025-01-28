# UedafEventNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportEvent** | Pointer to [**UedafEventReport**](UedafEventReport.md) |  | [optional] 
**EventSubsSyncInfo** | Pointer to [**UedafEventSubsSyncInfo**](UedafEventSubsSyncInfo.md) |  | [optional] 

## Methods

### NewUedafEventNotification

`func NewUedafEventNotification() *UedafEventNotification`

NewUedafEventNotification instantiates a new UedafEventNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafEventNotificationWithDefaults

`func NewUedafEventNotificationWithDefaults() *UedafEventNotification`

NewUedafEventNotificationWithDefaults instantiates a new UedafEventNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportEvent

`func (o *UedafEventNotification) GetReportEvent() UedafEventReport`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *UedafEventNotification) GetReportEventOk() (*UedafEventReport, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *UedafEventNotification) SetReportEvent(v UedafEventReport)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *UedafEventNotification) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.

### GetEventSubsSyncInfo

`func (o *UedafEventNotification) GetEventSubsSyncInfo() UedafEventSubsSyncInfo`

GetEventSubsSyncInfo returns the EventSubsSyncInfo field if non-nil, zero value otherwise.

### GetEventSubsSyncInfoOk

`func (o *UedafEventNotification) GetEventSubsSyncInfoOk() (*UedafEventSubsSyncInfo, bool)`

GetEventSubsSyncInfoOk returns a tuple with the EventSubsSyncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSubsSyncInfo

`func (o *UedafEventNotification) SetEventSubsSyncInfo(v UedafEventSubsSyncInfo)`

SetEventSubsSyncInfo sets EventSubsSyncInfo field to given value.

### HasEventSubsSyncInfo

`func (o *UedafEventNotification) HasEventSubsSyncInfo() bool`

HasEventSubsSyncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


