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

// checks if the OperationTracing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OperationTracing{}

// OperationTracing struct for OperationTracing
type OperationTracing struct {
	Preprocess  []PipelineTracing `json:"preprocess,omitempty"`
	Postprocess []PipelineTracing `json:"postprocess,omitempty"`
}

// NewOperationTracing instantiates a new OperationTracing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationTracing() *OperationTracing {
	this := OperationTracing{}
	return &this
}

// NewOperationTracingWithDefaults instantiates a new OperationTracing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationTracingWithDefaults() *OperationTracing {
	this := OperationTracing{}
	return &this
}

// GetPreprocess returns the Preprocess field value if set, zero value otherwise.
func (o *OperationTracing) GetPreprocess() []PipelineTracing {
	if o == nil || IsNil(o.Preprocess) {
		var ret []PipelineTracing
		return ret
	}
	return o.Preprocess
}

// GetPreprocessOk returns a tuple with the Preprocess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTracing) GetPreprocessOk() ([]PipelineTracing, bool) {
	if o == nil || IsNil(o.Preprocess) {
		return nil, false
	}
	return o.Preprocess, true
}

// HasPreprocess returns a boolean if a field has been set.
func (o *OperationTracing) HasPreprocess() bool {
	if o != nil && !IsNil(o.Preprocess) {
		return true
	}

	return false
}

// SetPreprocess gets a reference to the given []PipelineTracing and assigns it to the Preprocess field.
func (o *OperationTracing) SetPreprocess(v []PipelineTracing) {
	o.Preprocess = v
}

// GetPostprocess returns the Postprocess field value if set, zero value otherwise.
func (o *OperationTracing) GetPostprocess() []PipelineTracing {
	if o == nil || IsNil(o.Postprocess) {
		var ret []PipelineTracing
		return ret
	}
	return o.Postprocess
}

// GetPostprocessOk returns a tuple with the Postprocess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationTracing) GetPostprocessOk() ([]PipelineTracing, bool) {
	if o == nil || IsNil(o.Postprocess) {
		return nil, false
	}
	return o.Postprocess, true
}

// HasPostprocess returns a boolean if a field has been set.
func (o *OperationTracing) HasPostprocess() bool {
	if o != nil && !IsNil(o.Postprocess) {
		return true
	}

	return false
}

// SetPostprocess gets a reference to the given []PipelineTracing and assigns it to the Postprocess field.
func (o *OperationTracing) SetPostprocess(v []PipelineTracing) {
	o.Postprocess = v
}

func (o OperationTracing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OperationTracing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Preprocess) {
		toSerialize["preprocess"] = o.Preprocess
	}
	if !IsNil(o.Postprocess) {
		toSerialize["postprocess"] = o.Postprocess
	}
	return toSerialize, nil
}

type NullableOperationTracing struct {
	value *OperationTracing
	isSet bool
}

func (v NullableOperationTracing) Get() *OperationTracing {
	return v.value
}

func (v *NullableOperationTracing) Set(val *OperationTracing) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationTracing) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationTracing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationTracing(val *OperationTracing) *NullableOperationTracing {
	return &NullableOperationTracing{value: val, isSet: true}
}

func (v NullableOperationTracing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationTracing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
