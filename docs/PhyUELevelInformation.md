# PhyUELevelInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | Identifies a SUPI for an UE.  | [optional] 
**CellId** | Pointer to **string** | Cell identifier to which the UE is connected  | [optional] 
**ChannelQualityIndicator** | Pointer to **int32** | Channel quality indicator | [optional] 
**SNR** | Pointer to **float32** | Signal to noise ratio | [optional] 
**Rsrp** | Pointer to **float32** | Reference Signal Received Power | [optional] 
**Rsrq** | Pointer to **float32** | Reference Signal Received Quality  | [optional] 
**AvgMCSUplink** | Pointer to **float32** | Modulation used in uplink | [optional] 
**AvgMCSDownlink** | Pointer to **float32** | Modulation used in downlink | [optional] 
**RetransmissionUplink** | Pointer to **int32** | Uplink retransmissions | [optional] 
**RetransmissionDownlink** | Pointer to **int32** | Downlink retransmissions | [optional] 

## Methods

### NewPhyUELevelInformation

`func NewPhyUELevelInformation() *PhyUELevelInformation`

NewPhyUELevelInformation instantiates a new PhyUELevelInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhyUELevelInformationWithDefaults

`func NewPhyUELevelInformationWithDefaults() *PhyUELevelInformation`

NewPhyUELevelInformationWithDefaults instantiates a new PhyUELevelInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *PhyUELevelInformation) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PhyUELevelInformation) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PhyUELevelInformation) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *PhyUELevelInformation) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetCellId

`func (o *PhyUELevelInformation) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *PhyUELevelInformation) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *PhyUELevelInformation) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *PhyUELevelInformation) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetChannelQualityIndicator

`func (o *PhyUELevelInformation) GetChannelQualityIndicator() int32`

GetChannelQualityIndicator returns the ChannelQualityIndicator field if non-nil, zero value otherwise.

### GetChannelQualityIndicatorOk

`func (o *PhyUELevelInformation) GetChannelQualityIndicatorOk() (*int32, bool)`

GetChannelQualityIndicatorOk returns a tuple with the ChannelQualityIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelQualityIndicator

`func (o *PhyUELevelInformation) SetChannelQualityIndicator(v int32)`

SetChannelQualityIndicator sets ChannelQualityIndicator field to given value.

### HasChannelQualityIndicator

`func (o *PhyUELevelInformation) HasChannelQualityIndicator() bool`

HasChannelQualityIndicator returns a boolean if a field has been set.

### GetSNR

`func (o *PhyUELevelInformation) GetSNR() float32`

GetSNR returns the SNR field if non-nil, zero value otherwise.

### GetSNROk

`func (o *PhyUELevelInformation) GetSNROk() (*float32, bool)`

GetSNROk returns a tuple with the SNR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNR

`func (o *PhyUELevelInformation) SetSNR(v float32)`

SetSNR sets SNR field to given value.

### HasSNR

`func (o *PhyUELevelInformation) HasSNR() bool`

HasSNR returns a boolean if a field has been set.

### GetRsrp

`func (o *PhyUELevelInformation) GetRsrp() float32`

GetRsrp returns the Rsrp field if non-nil, zero value otherwise.

### GetRsrpOk

`func (o *PhyUELevelInformation) GetRsrpOk() (*float32, bool)`

GetRsrpOk returns a tuple with the Rsrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsrp

`func (o *PhyUELevelInformation) SetRsrp(v float32)`

SetRsrp sets Rsrp field to given value.

### HasRsrp

`func (o *PhyUELevelInformation) HasRsrp() bool`

HasRsrp returns a boolean if a field has been set.

### GetRsrq

`func (o *PhyUELevelInformation) GetRsrq() float32`

GetRsrq returns the Rsrq field if non-nil, zero value otherwise.

### GetRsrqOk

`func (o *PhyUELevelInformation) GetRsrqOk() (*float32, bool)`

GetRsrqOk returns a tuple with the Rsrq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsrq

`func (o *PhyUELevelInformation) SetRsrq(v float32)`

SetRsrq sets Rsrq field to given value.

### HasRsrq

`func (o *PhyUELevelInformation) HasRsrq() bool`

HasRsrq returns a boolean if a field has been set.

### GetAvgMCSUplink

`func (o *PhyUELevelInformation) GetAvgMCSUplink() float32`

GetAvgMCSUplink returns the AvgMCSUplink field if non-nil, zero value otherwise.

### GetAvgMCSUplinkOk

`func (o *PhyUELevelInformation) GetAvgMCSUplinkOk() (*float32, bool)`

GetAvgMCSUplinkOk returns a tuple with the AvgMCSUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMCSUplink

`func (o *PhyUELevelInformation) SetAvgMCSUplink(v float32)`

SetAvgMCSUplink sets AvgMCSUplink field to given value.

### HasAvgMCSUplink

`func (o *PhyUELevelInformation) HasAvgMCSUplink() bool`

HasAvgMCSUplink returns a boolean if a field has been set.

### GetAvgMCSDownlink

`func (o *PhyUELevelInformation) GetAvgMCSDownlink() float32`

GetAvgMCSDownlink returns the AvgMCSDownlink field if non-nil, zero value otherwise.

### GetAvgMCSDownlinkOk

`func (o *PhyUELevelInformation) GetAvgMCSDownlinkOk() (*float32, bool)`

GetAvgMCSDownlinkOk returns a tuple with the AvgMCSDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMCSDownlink

`func (o *PhyUELevelInformation) SetAvgMCSDownlink(v float32)`

SetAvgMCSDownlink sets AvgMCSDownlink field to given value.

### HasAvgMCSDownlink

`func (o *PhyUELevelInformation) HasAvgMCSDownlink() bool`

HasAvgMCSDownlink returns a boolean if a field has been set.

### GetRetransmissionUplink

`func (o *PhyUELevelInformation) GetRetransmissionUplink() int32`

GetRetransmissionUplink returns the RetransmissionUplink field if non-nil, zero value otherwise.

### GetRetransmissionUplinkOk

`func (o *PhyUELevelInformation) GetRetransmissionUplinkOk() (*int32, bool)`

GetRetransmissionUplinkOk returns a tuple with the RetransmissionUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmissionUplink

`func (o *PhyUELevelInformation) SetRetransmissionUplink(v int32)`

SetRetransmissionUplink sets RetransmissionUplink field to given value.

### HasRetransmissionUplink

`func (o *PhyUELevelInformation) HasRetransmissionUplink() bool`

HasRetransmissionUplink returns a boolean if a field has been set.

### GetRetransmissionDownlink

`func (o *PhyUELevelInformation) GetRetransmissionDownlink() int32`

GetRetransmissionDownlink returns the RetransmissionDownlink field if non-nil, zero value otherwise.

### GetRetransmissionDownlinkOk

`func (o *PhyUELevelInformation) GetRetransmissionDownlinkOk() (*int32, bool)`

GetRetransmissionDownlinkOk returns a tuple with the RetransmissionDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmissionDownlink

`func (o *PhyUELevelInformation) SetRetransmissionDownlink(v int32)`

SetRetransmissionDownlink sets RetransmissionDownlink field to given value.

### HasRetransmissionDownlink

`func (o *PhyUELevelInformation) HasRetransmissionDownlink() bool`

HasRetransmissionDownlink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


