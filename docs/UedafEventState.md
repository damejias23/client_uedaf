# UedafEventState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**RemainReports** | Pointer to **int32** |  | [optional] 
**RemainDuration** | Pointer to **int32** | indicating a time in seconds. | [optional] 

## Methods

### NewUedafEventState

`func NewUedafEventState(active bool, ) *UedafEventState`

NewUedafEventState instantiates a new UedafEventState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUedafEventStateWithDefaults

`func NewUedafEventStateWithDefaults() *UedafEventState`

NewUedafEventStateWithDefaults instantiates a new UedafEventState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UedafEventState) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UedafEventState) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UedafEventState) SetActive(v bool)`

SetActive sets Active field to given value.


### GetRemainReports

`func (o *UedafEventState) GetRemainReports() int32`

GetRemainReports returns the RemainReports field if non-nil, zero value otherwise.

### GetRemainReportsOk

`func (o *UedafEventState) GetRemainReportsOk() (*int32, bool)`

GetRemainReportsOk returns a tuple with the RemainReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainReports

`func (o *UedafEventState) SetRemainReports(v int32)`

SetRemainReports sets RemainReports field to given value.

### HasRemainReports

`func (o *UedafEventState) HasRemainReports() bool`

HasRemainReports returns a boolean if a field has been set.

### GetRemainDuration

`func (o *UedafEventState) GetRemainDuration() int32`

GetRemainDuration returns the RemainDuration field if non-nil, zero value otherwise.

### GetRemainDurationOk

`func (o *UedafEventState) GetRemainDurationOk() (*int32, bool)`

GetRemainDurationOk returns a tuple with the RemainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainDuration

`func (o *UedafEventState) SetRemainDuration(v int32)`

SetRemainDuration sets RemainDuration field to given value.

### HasRemainDuration

`func (o *UedafEventState) HasRemainDuration() bool`

HasRemainDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


