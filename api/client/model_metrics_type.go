/*
Merlin

API Guide for accessing Merlin's model management, deployment, and serving functionalities

API version: 0.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// MetricsType the model 'MetricsType'
type MetricsType string

// List of MetricsType
const (
	METRICSTYPE_CONCURRENCY        MetricsType = "concurrency"
	METRICSTYPE_CPU_UTILIZATION    MetricsType = "cpu_utilization"
	METRICSTYPE_MEMORY_UTILIZATION MetricsType = "memory_utilization"
	METRICSTYPE_RPS                MetricsType = "rps"
)

// All allowed values of MetricsType enum
var AllowedMetricsTypeEnumValues = []MetricsType{
	"concurrency",
	"cpu_utilization",
	"memory_utilization",
	"rps",
}

func (v *MetricsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricsType(value)
	for _, existing := range AllowedMetricsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricsType", value)
}

// NewMetricsTypeFromValue returns a pointer to a valid MetricsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricsTypeFromValue(v string) (*MetricsType, error) {
	ev := MetricsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricsType: valid values are %v", v, AllowedMetricsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricsType) IsValid() bool {
	for _, existing := range AllowedMetricsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricsType value
func (v MetricsType) Ptr() *MetricsType {
	return &v
}

type NullableMetricsType struct {
	value *MetricsType
	isSet bool
}

func (v NullableMetricsType) Get() *MetricsType {
	return v.value
}

func (v *NullableMetricsType) Set(val *MetricsType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsType(val *MetricsType) *NullableMetricsType {
	return &NullableMetricsType{value: val, isSet: true}
}

func (v NullableMetricsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
