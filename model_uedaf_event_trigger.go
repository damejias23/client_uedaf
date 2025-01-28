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

// UedafEventTrigger Describes how UEDAF should generate the report for the event
type UedafEventTrigger struct {
	UedafEventTriggerAnyOf *UedafEventTriggerAnyOf
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UedafEventTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UedafEventTriggerAnyOf
	err = json.Unmarshal(data, &dst.UedafEventTriggerAnyOf);
	if err == nil {
		jsonUedafEventTriggerAnyOf, _ := json.Marshal(dst.UedafEventTriggerAnyOf)
		if string(jsonUedafEventTriggerAnyOf) == "{}" { // empty struct
			dst.UedafEventTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.UedafEventTriggerAnyOf, return on the first match
		}
	} else {
		dst.UedafEventTriggerAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(UedafEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UedafEventTrigger) MarshalJSON() ([]byte, error) {
	if src.UedafEventTriggerAnyOf != nil {
		return json.Marshal(&src.UedafEventTriggerAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUedafEventTrigger struct {
	value *UedafEventTrigger
	isSet bool
}

func (v NullableUedafEventTrigger) Get() *UedafEventTrigger {
	return v.value
}

func (v *NullableUedafEventTrigger) Set(val *UedafEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableUedafEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableUedafEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUedafEventTrigger(val *UedafEventTrigger) *NullableUedafEventTrigger {
	return &NullableUedafEventTrigger{value: val, isSet: true}
}

func (v NullableUedafEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUedafEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


