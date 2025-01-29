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

// UedafEventMode Describes how the reports shall be generated by a subscribed event
type UedafEventMode struct {
	Trigger UedafEventTriggerAnyOf `json:"trigger"`
	MaxReports *int32 `json:"maxReports,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	// indicating a time in seconds.
	RepPeriod *int32 `json:"repPeriod,omitempty"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio *int32 `json:"sampRatio,omitempty"`
}

// NewUedafEventMode instantiates a new UedafEventMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUedafEventMode(trigger UedafEventTriggerAnyOf) *UedafEventMode {
	this := UedafEventMode{}
	this.Trigger = trigger
	return &this
}

// NewUedafEventModeWithDefaults instantiates a new UedafEventMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUedafEventModeWithDefaults() *UedafEventMode {
	this := UedafEventMode{}
	return &this
}

// GetTrigger returns the Trigger field value
func (o *UedafEventMode) GetTrigger() UedafEventTriggerAnyOf {
	if o == nil {
		var ret UedafEventTriggerAnyOf
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *UedafEventMode) GetTriggerOk() (*UedafEventTriggerAnyOf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *UedafEventMode) SetTrigger(v UedafEventTriggerAnyOf) {
	o.Trigger = v
}

// GetMaxReports returns the MaxReports field value if set, zero value otherwise.
func (o *UedafEventMode) GetMaxReports() int32 {
	if o == nil || o.MaxReports == nil {
		var ret int32
		return ret
	}
	return *o.MaxReports
}

// GetMaxReportsOk returns a tuple with the MaxReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventMode) GetMaxReportsOk() (*int32, bool) {
	if o == nil || o.MaxReports == nil {
		return nil, false
	}
	return o.MaxReports, true
}

// HasMaxReports returns a boolean if a field has been set.
func (o *UedafEventMode) HasMaxReports() bool {
	if o != nil && o.MaxReports != nil {
		return true
	}

	return false
}

// SetMaxReports gets a reference to the given int32 and assigns it to the MaxReports field.
func (o *UedafEventMode) SetMaxReports(v int32) {
	o.MaxReports = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *UedafEventMode) GetExpiry() time.Time {
	if o == nil || o.Expiry == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventMode) GetExpiryOk() (*time.Time, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *UedafEventMode) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *UedafEventMode) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetRepPeriod returns the RepPeriod field value if set, zero value otherwise.
func (o *UedafEventMode) GetRepPeriod() int32 {
	if o == nil || o.RepPeriod == nil {
		var ret int32
		return ret
	}
	return *o.RepPeriod
}

// GetRepPeriodOk returns a tuple with the RepPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventMode) GetRepPeriodOk() (*int32, bool) {
	if o == nil || o.RepPeriod == nil {
		return nil, false
	}
	return o.RepPeriod, true
}

// HasRepPeriod returns a boolean if a field has been set.
func (o *UedafEventMode) HasRepPeriod() bool {
	if o != nil && o.RepPeriod != nil {
		return true
	}

	return false
}

// SetRepPeriod gets a reference to the given int32 and assigns it to the RepPeriod field.
func (o *UedafEventMode) SetRepPeriod(v int32) {
	o.RepPeriod = &v
}

// GetSampRatio returns the SampRatio field value if set, zero value otherwise.
func (o *UedafEventMode) GetSampRatio() int32 {
	if o == nil || o.SampRatio == nil {
		var ret int32
		return ret
	}
	return *o.SampRatio
}

// GetSampRatioOk returns a tuple with the SampRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UedafEventMode) GetSampRatioOk() (*int32, bool) {
	if o == nil || o.SampRatio == nil {
		return nil, false
	}
	return o.SampRatio, true
}

// HasSampRatio returns a boolean if a field has been set.
func (o *UedafEventMode) HasSampRatio() bool {
	if o != nil && o.SampRatio != nil {
		return true
	}

	return false
}

// SetSampRatio gets a reference to the given int32 and assigns it to the SampRatio field.
func (o *UedafEventMode) SetSampRatio(v int32) {
	o.SampRatio = &v
}

func (o UedafEventMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["trigger"] = o.Trigger
	}
	if o.MaxReports != nil {
		toSerialize["maxReports"] = o.MaxReports
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.RepPeriod != nil {
		toSerialize["repPeriod"] = o.RepPeriod
	}
	if o.SampRatio != nil {
		toSerialize["sampRatio"] = o.SampRatio
	}
	return json.Marshal(toSerialize)
}

type NullableUedafEventMode struct {
	value *UedafEventMode
	isSet bool
}

func (v NullableUedafEventMode) Get() *UedafEventMode {
	return v.value
}

func (v *NullableUedafEventMode) Set(val *UedafEventMode) {
	v.value = val
	v.isSet = true
}

func (v NullableUedafEventMode) IsSet() bool {
	return v.isSet
}

func (v *NullableUedafEventMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUedafEventMode(val *UedafEventMode) *NullableUedafEventMode {
	return &NullableUedafEventMode{value: val, isSet: true}
}

func (v NullableUedafEventMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUedafEventMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


