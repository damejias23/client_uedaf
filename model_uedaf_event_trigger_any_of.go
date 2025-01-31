/*
Nuedaf_EventExposure

UEDAF Event Exposure Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package uedaf

import (
	"encoding/json"
	"fmt"
)

// UedafEventTriggerAnyOf the model 'UedafEventTriggerAnyOf'
type UedafEventTriggerAnyOf string

// List of UedafEventTrigger_anyOf
const (
	ONE_TIME UedafEventTriggerAnyOf = "ONE_TIME"
	CONTINUOUS UedafEventTriggerAnyOf = "CONTINUOUS"
	PERIODIC UedafEventTriggerAnyOf = "PERIODIC"
)

// All allowed values of UedafEventTriggerAnyOf enum
var AllowedUedafEventTriggerAnyOfEnumValues = []UedafEventTriggerAnyOf{
	"ONE_TIME",
	"CONTINUOUS",
	"PERIODIC",
}

func (v *UedafEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UedafEventTriggerAnyOf(value)
	for _, existing := range AllowedUedafEventTriggerAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UedafEventTriggerAnyOf", value)
}

// NewUedafEventTriggerAnyOfFromValue returns a pointer to a valid UedafEventTriggerAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUedafEventTriggerAnyOfFromValue(v string) (*UedafEventTriggerAnyOf, error) {
	ev := UedafEventTriggerAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UedafEventTriggerAnyOf: valid values are %v", v, AllowedUedafEventTriggerAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UedafEventTriggerAnyOf) IsValid() bool {
	for _, existing := range AllowedUedafEventTriggerAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UedafEventTrigger_anyOf value
func (v UedafEventTriggerAnyOf) Ptr() *UedafEventTriggerAnyOf {
	return &v
}

type NullableUedafEventTriggerAnyOf struct {
	value *UedafEventTriggerAnyOf
	isSet bool
}

func (v NullableUedafEventTriggerAnyOf) Get() *UedafEventTriggerAnyOf {
	return v.value
}

func (v *NullableUedafEventTriggerAnyOf) Set(val *UedafEventTriggerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUedafEventTriggerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUedafEventTriggerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUedafEventTriggerAnyOf(val *UedafEventTriggerAnyOf) *NullableUedafEventTriggerAnyOf {
	return &NullableUedafEventTriggerAnyOf{value: val, isSet: true}
}

func (v NullableUedafEventTriggerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUedafEventTriggerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

