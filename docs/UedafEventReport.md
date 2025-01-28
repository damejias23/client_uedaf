# UedafEventReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**UedafEventType**](UedafEventType.md) |  | 
**State** | [**UedafEventState**](UedafEventState.md) |  | 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**SubscriptionId** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**PhyUELevelInfos** | Pointer to [**[]PhyUELevelInformation**](PhyUELevelInformation.md) |  | [optional] 
**UeInfos** | Pointer to [**[]UeInformation**](UeInformation.md) |  | [optional] 

## Methods

### NewUedafEventReport

`func NewUedafEventReport(type_ UedafEventType, state UedafEventState, timeStamp time.Time, ) *UedafEventReport`

NewUedafEventReport instantiates a new UedafEventReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafEventReportWithDefaults

`func NewUedafEventReportWithDefaults() *UedafEventReport`

NewUedafEventReportWithDefaults instantiates a new UedafEventReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UedafEventReport) GetType() UedafEventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UedafEventReport) GetTypeOk() (*UedafEventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UedafEventReport) SetType(v UedafEventType)`

SetType sets Type field to given value.


### GetState

`func (o *UedafEventReport) GetState() UedafEventState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UedafEventReport) GetStateOk() (*UedafEventState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UedafEventReport) SetState(v UedafEventState)`

SetState sets State field to given value.


### GetTimeStamp

`func (o *UedafEventReport) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *UedafEventReport) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *UedafEventReport) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetSubscriptionId

`func (o *UedafEventReport) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UedafEventReport) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UedafEventReport) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *UedafEventReport) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetPhyUELevelInfos

`func (o *UedafEventReport) GetPhyUELevelInfos() []PhyUELevelInformation`

GetPhyUELevelInfos returns the PhyUELevelInfos field if non-nil, zero value otherwise.

### GetPhyUELevelInfosOk

`func (o *UedafEventReport) GetPhyUELevelInfosOk() (*[]PhyUELevelInformation, bool)`

GetPhyUELevelInfosOk returns a tuple with the PhyUELevelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhyUELevelInfos

`func (o *UedafEventReport) SetPhyUELevelInfos(v []PhyUELevelInformation)`

SetPhyUELevelInfos sets PhyUELevelInfos field to given value.

### HasPhyUELevelInfos

`func (o *UedafEventReport) HasPhyUELevelInfos() bool`

HasPhyUELevelInfos returns a boolean if a field has been set.

### GetUeInfos

`func (o *UedafEventReport) GetUeInfos() []UeInformation`

GetUeInfos returns the UeInfos field if non-nil, zero value otherwise.

### GetUeInfosOk

`func (o *UedafEventReport) GetUeInfosOk() (*[]UeInformation, bool)`

GetUeInfosOk returns a tuple with the UeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeInfos

`func (o *UedafEventReport) SetUeInfos(v []UeInformation)`

SetUeInfos sets UeInfos field to given value.

### HasUeInfos

`func (o *UedafEventReport) HasUeInfos() bool`

HasUeInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


