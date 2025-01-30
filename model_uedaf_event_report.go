/*
Nuedaf_EventExposure

UEDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uedaf

import (
	"encoding/json"
	"time"
)

// UedafEventReport Represents a report triggered by a subscribed event type
type UedafEventReport struct {
	Type UedafEventTypeAnyOf `json:"type"`
	State UedafEventState `json:"state"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	// String providing an URI formatted according to RFC 3986.
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	PhyUELevelInfos []PhyUELevelInformation `json:"phyUELevelInfos,omitempty"`
	UeInfos []UeInformation `json:"UeInfos,omitempty"`
}

// NewUedafEventReport instantiates a new UedafEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUedafEventReport(type_ UedafEventTypeAnyOf, state UedafEventState, timeStamp time.Time) *UedafEventReport {
	this := UedafEventReport{}
	this.Type = type_
	this.State = state
	this.TimeStamp = timeStamp
	return &this
}

// NewUedafEventReportWithDefaults instantiates a new UedafEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUedafEventReportWithDefaults() *UedafEventReport {
	this := UedafEventReport{}
	return &this
}

// GetType returns the Type field value
func (o *UedafEventReport) GetType() UedafEventTypeAnyOf {
	if o == nil {
		var ret UedafEventTypeAnyOf
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UedafEventReport) GetTypeOk() (*UedafEventTypeAnyOf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UedafEventReport) SetType(v UedafEventTypeAnyOf) {
	o.Type = v
}

// GetState returns the State field value
func (o *UedafEventReport) GetState() UedafEventState {
	if o == nil {
		var ret UedafEventState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UedafEventReport) GetStateOk() (*UedafEventState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UedafEventReport) SetState(v UedafEventState) {
	o.State = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *UedafEventReport) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *UedafEventReport) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *UedafEventReport) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *UedafEventReport) GetSubscriptionId() string {
	if o == nil || o.SubscriptionId == nil {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventReport) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || o.SubscriptionId == nil {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *UedafEventReport) HasSubscriptionId() bool {
	if o != nil && o.SubscriptionId != nil {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *UedafEventReport) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetPhyUELevelInfos returns the PhyUELevelInfos field value if set, zero value otherwise.
func (o *UedafEventReport) GetPhyUELevelInfos() []PhyUELevelInformation {
	if o == nil || o.PhyUELevelInfos == nil {
		var ret []PhyUELevelInformation
		return ret
	}
	return o.PhyUELevelInfos
}

// GetPhyUELevelInfosOk returns a tuple with the PhyUELevelInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventReport) GetPhyUELevelInfosOk() ([]PhyUELevelInformation, bool) {
	// if o == nil || o.PhyUELevelInfos == nil {
	if o == nil {
		return nil, false
	}
	return o.PhyUELevelInfos, true
}

// HasPhyUELevelInfos returns a boolean if a field has been set.
func (o *UedafEventReport) HasPhyUELevelInfos() bool {
	if o != nil && o.PhyUELevelInfos != nil {
		return true
	}

	return false
}

// SetPhyUELevelInfos gets a reference to the given []PhyUELevelInformation and assigns it to the PhyUELevelInfos field.
func (o *UedafEventReport) SetPhyUELevelInfos(v []PhyUELevelInformation) {
	o.PhyUELevelInfos = v
}

// GetUeInfos returns the UeInfos field value if set, zero value otherwise.
func (o *UedafEventReport) GetUeInfos() []UeInformation {
	if o == nil || o.UeInfos == nil {
		var ret []UeInformation
		return ret
	}
	return o.UeInfos
}

// GetUeInfosOk returns a tuple with the UeInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventReport) GetUeInfosOk() ([]UeInformation, bool) {
	if o == nil || o.UeInfos == nil {
		return nil, false
	}
	return o.UeInfos, true
}

// HasUeInfos returns a boolean if a field has been set.
func (o *UedafEventReport) HasUeInfos() bool {
	if o != nil && o.UeInfos != nil {
		return true
	}

	return false
}

// SetUeInfos gets a reference to the given []UeInformation and assigns it to the UeInfos field.
func (o *UedafEventReport) SetUeInfos(v []UeInformation) {
	o.UeInfos = v
}

func (o UedafEventReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if o.SubscriptionId != nil {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	if o.PhyUELevelInfos != nil {
		toSerialize["phyUELevelInfos"] = o.PhyUELevelInfos
	}
	if o.UeInfos != nil {
		toSerialize["UeInfos"] = o.UeInfos
	}
	return json.Marshal(toSerialize)
}

type NullableUedafEventReport struct {
	value *UedafEventReport
	isSet bool
}

func (v NullableUedafEventReport) Get() *UedafEventReport {
	return v.value
}

func (v *NullableUedafEventReport) Set(val *UedafEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableUedafEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableUedafEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUedafEventReport(val *UedafEventReport) *NullableUedafEventReport {
	return &NullableUedafEventReport{value: val, isSet: true}
}

func (v NullableUedafEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUedafEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


