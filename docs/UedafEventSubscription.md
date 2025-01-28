# UedafEventSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventRequest** | [**UedafEvent**](UedafEvent.md) |  | 
**EventNotifyUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**Options** | Pointer to [**UedafEventMode**](UedafEventMode.md) |  | [optional] 
**TgtUe** | Pointer to [**TargetUeInformation**](TargetUeInformation.md) |  | [optional] 

## Methods

### NewUedafEventSubscription

`func NewUedafEventSubscription(eventRequest UedafEvent, eventNotifyUri string, ) *UedafEventSubscription`

NewUedafEventSubscription instantiates a new UedafEventSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafEventSubscriptionWithDefaults

`func NewUedafEventSubscriptionWithDefaults() *UedafEventSubscription`

NewUedafEventSubscriptionWithDefaults instantiates a new UedafEventSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventRequest

`func (o *UedafEventSubscription) GetEventRequest() UedafEvent`

GetEventRequest returns the EventRequest field if non-nil, zero value otherwise.

### GetEventRequestOk

`func (o *UedafEventSubscription) GetEventRequestOk() (*UedafEvent, bool)`

GetEventRequestOk returns a tuple with the EventRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRequest

`func (o *UedafEventSubscription) SetEventRequest(v UedafEvent)`

SetEventRequest sets EventRequest field to given value.


### GetEventNotifyUri

`func (o *UedafEventSubscription) GetEventNotifyUri() string`

GetEventNotifyUri returns the EventNotifyUri field if non-nil, zero value otherwise.

### GetEventNotifyUriOk

`func (o *UedafEventSubscription) GetEventNotifyUriOk() (*string, bool)`

GetEventNotifyUriOk returns a tuple with the EventNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNotifyUri

`func (o *UedafEventSubscription) SetEventNotifyUri(v string)`

SetEventNotifyUri sets EventNotifyUri field to given value.


### GetOptions

`func (o *UedafEventSubscription) GetOptions() UedafEventMode`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UedafEventSubscription) GetOptionsOk() (*UedafEventMode, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UedafEventSubscription) SetOptions(v UedafEventMode)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UedafEventSubscription) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTgtUe

`func (o *UedafEventSubscription) GetTgtUe() TargetUeInformation`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *UedafEventSubscription) GetTgtUeOk() (*TargetUeInformation, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *UedafEventSubscription) SetTgtUe(v TargetUeInformation)`

SetTgtUe sets TgtUe field to given value.

### HasTgtUe

`func (o *UedafEventSubscription) HasTgtUe() bool`

HasTgtUe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


