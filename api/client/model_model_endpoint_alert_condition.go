/*
Merlin

API Guide for accessing Merlin's model management, deployment, and serving functionalities

API version: 0.14.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelEndpointAlertCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelEndpointAlertCondition{}

// ModelEndpointAlertCondition struct for ModelEndpointAlertCondition
type ModelEndpointAlertCondition struct {
	Enabled    *bool                     `json:"enabled,omitempty"`
	MetricType *AlertConditionMetricType `json:"metric_type,omitempty"`
	Severity   *AlertConditionSeverity   `json:"severity,omitempty"`
	Target     *float32                  `json:"target,omitempty"`
	Percentile *float32                  `json:"percentile,omitempty"`
	Unit       *string                   `json:"unit,omitempty"`
}

// NewModelEndpointAlertCondition instantiates a new ModelEndpointAlertCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelEndpointAlertCondition() *ModelEndpointAlertCondition {
	this := ModelEndpointAlertCondition{}
	return &this
}

// NewModelEndpointAlertConditionWithDefaults instantiates a new ModelEndpointAlertCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelEndpointAlertConditionWithDefaults() *ModelEndpointAlertCondition {
	this := ModelEndpointAlertCondition{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ModelEndpointAlertCondition) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEndpointAlertCondition) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ModelEndpointAlertCondition) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ModelEndpointAlertCondition) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMetricType returns the MetricType field value if set, zero value otherwise.
func (o *ModelEndpointAlertCondition) GetMetricType() AlertConditionMetricType {
	if o == nil || IsNil(o.MetricType) {
		var ret AlertConditionMetricType
		return ret
	}
	return *o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEndpointAlertCondition) GetMetricTypeOk() (*AlertConditionMetricType, bool) {
	if o == nil || IsNil(o.MetricType) {
		return nil, false
	}
	return o.MetricType, true
}

// HasMetricType returns a boolean if a field has been set.
func (o *ModelEndpointAlertCondition) HasMetricType() bool {
	if o != nil && !IsNil(o.MetricType) {
		return true
	}

	return false
}

// SetMetricType gets a reference to the given AlertConditionMetricType and assigns it to the MetricType field.
func (o *ModelEndpointAlertCondition) SetMetricType(v AlertConditionMetricType) {
	o.MetricType = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ModelEndpointAlertCondition) GetSeverity() AlertConditionSeverity {
	if o == nil || IsNil(o.Severity) {
		var ret AlertConditionSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEndpointAlertCondition) GetSeverityOk() (*AlertConditionSeverity, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ModelEndpointAlertCondition) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given AlertConditionSeverity and assigns it to the Severity field.
func (o *ModelEndpointAlertCondition) SetSeverity(v AlertConditionSeverity) {
	o.Severity = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ModelEndpointAlertCondition) GetTarget() float32 {
	if o == nil || IsNil(o.Target) {
		var ret float32
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEndpointAlertCondition) GetTargetOk() (*float32, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ModelEndpointAlertCondition) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given float32 and assigns it to the Target field.
func (o *ModelEndpointAlertCondition) SetTarget(v float32) {
	o.Target = &v
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *ModelEndpointAlertCondition) GetPercentile() float32 {
	if o == nil || IsNil(o.Percentile) {
		var ret float32
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEndpointAlertCondition) GetPercentileOk() (*float32, bool) {
	if o == nil || IsNil(o.Percentile) {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *ModelEndpointAlertCondition) HasPercentile() bool {
	if o != nil && !IsNil(o.Percentile) {
		return true
	}

	return false
}

// SetPercentile gets a reference to the given float32 and assigns it to the Percentile field.
func (o *ModelEndpointAlertCondition) SetPercentile(v float32) {
	o.Percentile = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ModelEndpointAlertCondition) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelEndpointAlertCondition) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ModelEndpointAlertCondition) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *ModelEndpointAlertCondition) SetUnit(v string) {
	o.Unit = &v
}

func (o ModelEndpointAlertCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelEndpointAlertCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.MetricType) {
		toSerialize["metric_type"] = o.MetricType
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.Percentile) {
		toSerialize["percentile"] = o.Percentile
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableModelEndpointAlertCondition struct {
	value *ModelEndpointAlertCondition
	isSet bool
}

func (v NullableModelEndpointAlertCondition) Get() *ModelEndpointAlertCondition {
	return v.value
}

func (v *NullableModelEndpointAlertCondition) Set(val *ModelEndpointAlertCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableModelEndpointAlertCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableModelEndpointAlertCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelEndpointAlertCondition(val *ModelEndpointAlertCondition) *NullableModelEndpointAlertCondition {
	return &NullableModelEndpointAlertCondition{value: val, isSet: true}
}

func (v NullableModelEndpointAlertCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelEndpointAlertCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
