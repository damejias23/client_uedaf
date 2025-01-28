# UeInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | Identifies a SUPI for an UE.  | [optional] 
**CellId** | Pointer to **string** | Cell identifier to which the UE is connected  | [optional] 
**Plmn** | Pointer to **string** | PLMN indicator | [optional] 
**Status** | Pointer to **string** | Status indicator | [optional] 
**StatusTimestamp** | Pointer to **string** | Status timestamp indicator | [optional] 

## Methods

### NewUeInformation

`func NewUeInformation() *UeInformation`

NewUeInformation instantiates a new UeInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeInformationWithDefaults

`func NewUeInformationWithDefaults() *UeInformation`

NewUeInformationWithDefaults instantiates a new UeInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeInformation) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeInformation) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeInformation) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UeInformation) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetCellId

`func (o *UeInformation) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *UeInformation) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *UeInformation) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *UeInformation) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetPlmn

`func (o *UeInformation) GetPlmn() string`

GetPlmn returns the Plmn field if non-nil, zero value otherwise.

### GetPlmnOk

`func (o *UeInformation) GetPlmnOk() (*string, bool)`

GetPlmnOk returns a tuple with the Plmn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmn

`func (o *UeInformation) SetPlmn(v string)`

SetPlmn sets Plmn field to given value.

### HasPlmn

`func (o *UeInformation) HasPlmn() bool`

HasPlmn returns a boolean if a field has been set.

### GetStatus

`func (o *UeInformation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UeInformation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UeInformation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UeInformation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusTimestamp

`func (o *UeInformation) GetStatusTimestamp() string`

GetStatusTimestamp returns the StatusTimestamp field if non-nil, zero value otherwise.

### GetStatusTimestampOk

`func (o *UeInformation) GetStatusTimestampOk() (*string, bool)`

GetStatusTimestampOk returns a tuple with the StatusTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTimestamp

`func (o *UeInformation) SetStatusTimestamp(v string)`

SetStatusTimestamp sets StatusTimestamp field to given value.

### HasStatusTimestamp

`func (o *UeInformation) HasStatusTimestamp() bool`

HasStatusTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


