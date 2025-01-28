/*
Nuedaf_EventExposure

UEDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uedaf

import (
	"encoding/json"
)

// PhyUELevelInformation Represents load level information of a given NF instance.
type PhyUELevelInformation struct {
	// Identifies a SUPI for an UE. 
	Supi *string `json:"supi,omitempty"`
	// Cell identifier to which the UE is connected 
	CellId *string `json:"cell_id,omitempty"`
	// Channel quality indicator
	ChannelQualityIndicator *int32 `json:"channel_quality_indicator,omitempty"`
	// Signal to noise ratio
	SNR *float32 `json:"SNR,omitempty"`
	// Reference Signal Received Power
	Rsrp *float32 `json:"rsrp,omitempty"`
	// Reference Signal Received Quality 
	Rsrq *float32 `json:"rsrq,omitempty"`
	// Modulation used in uplink
	AvgMCSUplink *float32 `json:"avg_MCS_uplink,omitempty"`
	// Modulation used in downlink
	AvgMCSDownlink *float32 `json:"avg_MCS_downlink,omitempty"`
	// Uplink retransmissions
	RetransmissionUplink *int32 `json:"retransmission_uplink,omitempty"`
	// Downlink retransmissions
	RetransmissionDownlink *int32 `json:"retransmission_downlink,omitempty"`
}

// NewPhyUELevelInformation instantiates a new PhyUELevelInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhyUELevelInformation() *PhyUELevelInformation {
	this := PhyUELevelInformation{}
	return &this
}

// NewPhyUELevelInformationWithDefaults instantiates a new PhyUELevelInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhyUELevelInformationWithDefaults() *PhyUELevelInformation {
	this := PhyUELevelInformation{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetSupi() string {
	if o == nil || o.Supi == nil {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetSupiOk() (*string, bool) {
	if o == nil || o.Supi == nil {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasSupi() bool {
	if o != nil && o.Supi != nil {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *PhyUELevelInformation) SetSupi(v string) {
	o.Supi = &v
}

// GetCellId returns the CellId field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetCellId() string {
	if o == nil || o.CellId == nil {
		var ret string
		return ret
	}
	return *o.CellId
}

// GetCellIdOk returns a tuple with the CellId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetCellIdOk() (*string, bool) {
	if o == nil || o.CellId == nil {
		return nil, false
	}
	return o.CellId, true
}

// HasCellId returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasCellId() bool {
	if o != nil && o.CellId != nil {
		return true
	}

	return false
}

// SetCellId gets a reference to the given string and assigns it to the CellId field.
func (o *PhyUELevelInformation) SetCellId(v string) {
	o.CellId = &v
}

// GetChannelQualityIndicator returns the ChannelQualityIndicator field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetChannelQualityIndicator() int32 {
	if o == nil || o.ChannelQualityIndicator == nil {
		var ret int32
		return ret
	}
	return *o.ChannelQualityIndicator
}

// GetChannelQualityIndicatorOk returns a tuple with the ChannelQualityIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetChannelQualityIndicatorOk() (*int32, bool) {
	if o == nil || o.ChannelQualityIndicator == nil {
		return nil, false
	}
	return o.ChannelQualityIndicator, true
}

// HasChannelQualityIndicator returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasChannelQualityIndicator() bool {
	if o != nil && o.ChannelQualityIndicator != nil {
		return true
	}

	return false
}

// SetChannelQualityIndicator gets a reference to the given int32 and assigns it to the ChannelQualityIndicator field.
func (o *PhyUELevelInformation) SetChannelQualityIndicator(v int32) {
	o.ChannelQualityIndicator = &v
}

// GetSNR returns the SNR field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetSNR() float32 {
	if o == nil || o.SNR == nil {
		var ret float32
		return ret
	}
	return *o.SNR
}

// GetSNROk returns a tuple with the SNR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetSNROk() (*float32, bool) {
	if o == nil || o.SNR == nil {
		return nil, false
	}
	return o.SNR, true
}

// HasSNR returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasSNR() bool {
	if o != nil && o.SNR != nil {
		return true
	}

	return false
}

// SetSNR gets a reference to the given float32 and assigns it to the SNR field.
func (o *PhyUELevelInformation) SetSNR(v float32) {
	o.SNR = &v
}

// GetRsrp returns the Rsrp field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetRsrp() float32 {
	if o == nil || o.Rsrp == nil {
		var ret float32
		return ret
	}
	return *o.Rsrp
}

// GetRsrpOk returns a tuple with the Rsrp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetRsrpOk() (*float32, bool) {
	if o == nil || o.Rsrp == nil {
		return nil, false
	}
	return o.Rsrp, true
}

// HasRsrp returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasRsrp() bool {
	if o != nil && o.Rsrp != nil {
		return true
	}

	return false
}

// SetRsrp gets a reference to the given float32 and assigns it to the Rsrp field.
func (o *PhyUELevelInformation) SetRsrp(v float32) {
	o.Rsrp = &v
}

// GetRsrq returns the Rsrq field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetRsrq() float32 {
	if o == nil || o.Rsrq == nil {
		var ret float32
		return ret
	}
	return *o.Rsrq
}

// GetRsrqOk returns a tuple with the Rsrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetRsrqOk() (*float32, bool) {
	if o == nil || o.Rsrq == nil {
		return nil, false
	}
	return o.Rsrq, true
}

// HasRsrq returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasRsrq() bool {
	if o != nil && o.Rsrq != nil {
		return true
	}

	return false
}

// SetRsrq gets a reference to the given float32 and assigns it to the Rsrq field.
func (o *PhyUELevelInformation) SetRsrq(v float32) {
	o.Rsrq = &v
}

// GetAvgMCSUplink returns the AvgMCSUplink field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetAvgMCSUplink() float32 {
	if o == nil || o.AvgMCSUplink == nil {
		var ret float32
		return ret
	}
	return *o.AvgMCSUplink
}

// GetAvgMCSUplinkOk returns a tuple with the AvgMCSUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetAvgMCSUplinkOk() (*float32, bool) {
	if o == nil || o.AvgMCSUplink == nil {
		return nil, false
	}
	return o.AvgMCSUplink, true
}

// HasAvgMCSUplink returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasAvgMCSUplink() bool {
	if o != nil && o.AvgMCSUplink != nil {
		return true
	}

	return false
}

// SetAvgMCSUplink gets a reference to the given float32 and assigns it to the AvgMCSUplink field.
func (o *PhyUELevelInformation) SetAvgMCSUplink(v float32) {
	o.AvgMCSUplink = &v
}

// GetAvgMCSDownlink returns the AvgMCSDownlink field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetAvgMCSDownlink() float32 {
	if o == nil || o.AvgMCSDownlink == nil {
		var ret float32
		return ret
	}
	return *o.AvgMCSDownlink
}

// GetAvgMCSDownlinkOk returns a tuple with the AvgMCSDownlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetAvgMCSDownlinkOk() (*float32, bool) {
	if o == nil || o.AvgMCSDownlink == nil {
		return nil, false
	}
	return o.AvgMCSDownlink, true
}

// HasAvgMCSDownlink returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasAvgMCSDownlink() bool {
	if o != nil && o.AvgMCSDownlink != nil {
		return true
	}

	return false
}

// SetAvgMCSDownlink gets a reference to the given float32 and assigns it to the AvgMCSDownlink field.
func (o *PhyUELevelInformation) SetAvgMCSDownlink(v float32) {
	o.AvgMCSDownlink = &v
}

// GetRetransmissionUplink returns the RetransmissionUplink field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetRetransmissionUplink() int32 {
	if o == nil || o.RetransmissionUplink == nil {
		var ret int32
		return ret
	}
	return *o.RetransmissionUplink
}

// GetRetransmissionUplinkOk returns a tuple with the RetransmissionUplink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetRetransmissionUplinkOk() (*int32, bool) {
	if o == nil || o.RetransmissionUplink == nil {
		return nil, false
	}
	return o.RetransmissionUplink, true
}

// HasRetransmissionUplink returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasRetransmissionUplink() bool {
	if o != nil && o.RetransmissionUplink != nil {
		return true
	}

	return false
}

// SetRetransmissionUplink gets a reference to the given int32 and assigns it to the RetransmissionUplink field.
func (o *PhyUELevelInformation) SetRetransmissionUplink(v int32) {
	o.RetransmissionUplink = &v
}

// GetRetransmissionDownlink returns the RetransmissionDownlink field value if set, zero value otherwise.
func (o *PhyUELevelInformation) GetRetransmissionDownlink() int32 {
	if o == nil || o.RetransmissionDownlink == nil {
		var ret int32
		return ret
	}
	return *o.RetransmissionDownlink
}

// GetRetransmissionDownlinkOk returns a tuple with the RetransmissionDownlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhyUELevelInformation) GetRetransmissionDownlinkOk() (*int32, bool) {
	if o == nil || o.RetransmissionDownlink == nil {
		return nil, false
	}
	return o.RetransmissionDownlink, true
}

// HasRetransmissionDownlink returns a boolean if a field has been set.
func (o *PhyUELevelInformation) HasRetransmissionDownlink() bool {
	if o != nil && o.RetransmissionDownlink != nil {
		return true
	}

	return false
}

// SetRetransmissionDownlink gets a reference to the given int32 and assigns it to the RetransmissionDownlink field.
func (o *PhyUELevelInformation) SetRetransmissionDownlink(v int32) {
	o.RetransmissionDownlink = &v
}

func (o PhyUELevelInformation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Supi != nil {
		toSerialize["supi"] = o.Supi
	}
	if o.CellId != nil {
		toSerialize["cell_id"] = o.CellId
	}
	if o.ChannelQualityIndicator != nil {
		toSerialize["channel_quality_indicator"] = o.ChannelQualityIndicator
	}
	if o.SNR != nil {
		toSerialize["SNR"] = o.SNR
	}
	if o.Rsrp != nil {
		toSerialize["rsrp"] = o.Rsrp
	}
	if o.Rsrq != nil {
		toSerialize["rsrq"] = o.Rsrq
	}
	if o.AvgMCSUplink != nil {
		toSerialize["avg_MCS_uplink"] = o.AvgMCSUplink
	}
	if o.AvgMCSDownlink != nil {
		toSerialize["avg_MCS_downlink"] = o.AvgMCSDownlink
	}
	if o.RetransmissionUplink != nil {
		toSerialize["retransmission_uplink"] = o.RetransmissionUplink
	}
	if o.RetransmissionDownlink != nil {
		toSerialize["retransmission_downlink"] = o.RetransmissionDownlink
	}
	return json.Marshal(toSerialize)
}

type NullablePhyUELevelInformation struct {
	value *PhyUELevelInformation
	isSet bool
}

func (v NullablePhyUELevelInformation) Get() *PhyUELevelInformation {
	return v.value
}

func (v *NullablePhyUELevelInformation) Set(val *PhyUELevelInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePhyUELevelInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePhyUELevelInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhyUELevelInformation(val *PhyUELevelInformation) *NullablePhyUELevelInformation {
	return &NullablePhyUELevelInformation{value: val, isSet: true}
}

func (v NullablePhyUELevelInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhyUELevelInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


