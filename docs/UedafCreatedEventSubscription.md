# UedafCreatedEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | [**UedafEventSubscription**](UedafEventSubscription.md) |  | 
**SubscriptionId** | **string** | String providing an URI formatted according to RFC 3986. | 
**ReportEvent** | Pointer to [**UedafEventReport**](UedafEventReport.md) |  | [optional] 

## Methods

### NewUedafCreatedEventSubscription

`func NewUedafCreatedEventSubscription(subscription UedafEventSubscription, subscriptionId string, ) *UedafCreatedEventSubscription`

NewUedafCreatedEventSubscription instantiates a new UedafCreatedEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafCreatedEventSubscriptionWithDefaults

`func NewUedafCreatedEventSubscriptionWithDefaults() *UedafCreatedEventSubscription`

NewUedafCreatedEventSubscriptionWithDefaults instantiates a new UedafCreatedEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *UedafCreatedEventSubscription) GetSubscription() UedafEventSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UedafCreatedEventSubscription) GetSubscriptionOk() (*UedafEventSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UedafCreatedEventSubscription) SetSubscription(v UedafEventSubscription)`

SetSubscription sets Subscription field to given value.


### GetSubscriptionId

`func (o *UedafCreatedEventSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UedafCreatedEventSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UedafCreatedEventSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetReportEvent

`func (o *UedafCreatedEventSubscription) GetReportEvent() UedafEventReport`

GetReportEvent returns the ReportEvent field if non-nil, zero value otherwise.

### GetReportEventOk

`func (o *UedafCreatedEventSubscription) GetReportEventOk() (*UedafEventReport, bool)`

GetReportEventOk returns a tuple with the ReportEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportEvent

`func (o *UedafCreatedEventSubscription) SetReportEvent(v UedafEventReport)`

SetReportEvent sets ReportEvent field to given value.

### HasReportEvent

`func (o *UedafCreatedEventSubscription) HasReportEvent() bool`

HasReportEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


