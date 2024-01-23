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

// checks if the PredictionJobResourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredictionJobResourceRequest{}

// PredictionJobResourceRequest struct for PredictionJobResourceRequest
type PredictionJobResourceRequest struct {
	DriverCpuRequest      *string `json:"driver_cpu_request,omitempty"`
	DriverMemoryRequest   *string `json:"driver_memory_request,omitempty"`
	ExecutorCpuRequest    *string `json:"executor_cpu_request,omitempty"`
	ExecutorMemoryRequest *string `json:"executor_memory_request,omitempty"`
	ExecutorReplica       *int32  `json:"executor_replica,omitempty"`
}

// NewPredictionJobResourceRequest instantiates a new PredictionJobResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionJobResourceRequest() *PredictionJobResourceRequest {
	this := PredictionJobResourceRequest{}
	return &this
}

// NewPredictionJobResourceRequestWithDefaults instantiates a new PredictionJobResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionJobResourceRequestWithDefaults() *PredictionJobResourceRequest {
	this := PredictionJobResourceRequest{}
	return &this
}

// GetDriverCpuRequest returns the DriverCpuRequest field value if set, zero value otherwise.
func (o *PredictionJobResourceRequest) GetDriverCpuRequest() string {
	if o == nil || IsNil(o.DriverCpuRequest) {
		var ret string
		return ret
	}
	return *o.DriverCpuRequest
}

// GetDriverCpuRequestOk returns a tuple with the DriverCpuRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionJobResourceRequest) GetDriverCpuRequestOk() (*string, bool) {
	if o == nil || IsNil(o.DriverCpuRequest) {
		return nil, false
	}
	return o.DriverCpuRequest, true
}

// HasDriverCpuRequest returns a boolean if a field has been set.
func (o *PredictionJobResourceRequest) HasDriverCpuRequest() bool {
	if o != nil && !IsNil(o.DriverCpuRequest) {
		return true
	}

	return false
}

// SetDriverCpuRequest gets a reference to the given string and assigns it to the DriverCpuRequest field.
func (o *PredictionJobResourceRequest) SetDriverCpuRequest(v string) {
	o.DriverCpuRequest = &v
}

// GetDriverMemoryRequest returns the DriverMemoryRequest field value if set, zero value otherwise.
func (o *PredictionJobResourceRequest) GetDriverMemoryRequest() string {
	if o == nil || IsNil(o.DriverMemoryRequest) {
		var ret string
		return ret
	}
	return *o.DriverMemoryRequest
}

// GetDriverMemoryRequestOk returns a tuple with the DriverMemoryRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionJobResourceRequest) GetDriverMemoryRequestOk() (*string, bool) {
	if o == nil || IsNil(o.DriverMemoryRequest) {
		return nil, false
	}
	return o.DriverMemoryRequest, true
}

// HasDriverMemoryRequest returns a boolean if a field has been set.
func (o *PredictionJobResourceRequest) HasDriverMemoryRequest() bool {
	if o != nil && !IsNil(o.DriverMemoryRequest) {
		return true
	}

	return false
}

// SetDriverMemoryRequest gets a reference to the given string and assigns it to the DriverMemoryRequest field.
func (o *PredictionJobResourceRequest) SetDriverMemoryRequest(v string) {
	o.DriverMemoryRequest = &v
}

// GetExecutorCpuRequest returns the ExecutorCpuRequest field value if set, zero value otherwise.
func (o *PredictionJobResourceRequest) GetExecutorCpuRequest() string {
	if o == nil || IsNil(o.ExecutorCpuRequest) {
		var ret string
		return ret
	}
	return *o.ExecutorCpuRequest
}

// GetExecutorCpuRequestOk returns a tuple with the ExecutorCpuRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionJobResourceRequest) GetExecutorCpuRequestOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutorCpuRequest) {
		return nil, false
	}
	return o.ExecutorCpuRequest, true
}

// HasExecutorCpuRequest returns a boolean if a field has been set.
func (o *PredictionJobResourceRequest) HasExecutorCpuRequest() bool {
	if o != nil && !IsNil(o.ExecutorCpuRequest) {
		return true
	}

	return false
}

// SetExecutorCpuRequest gets a reference to the given string and assigns it to the ExecutorCpuRequest field.
func (o *PredictionJobResourceRequest) SetExecutorCpuRequest(v string) {
	o.ExecutorCpuRequest = &v
}

// GetExecutorMemoryRequest returns the ExecutorMemoryRequest field value if set, zero value otherwise.
func (o *PredictionJobResourceRequest) GetExecutorMemoryRequest() string {
	if o == nil || IsNil(o.ExecutorMemoryRequest) {
		var ret string
		return ret
	}
	return *o.ExecutorMemoryRequest
}

// GetExecutorMemoryRequestOk returns a tuple with the ExecutorMemoryRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionJobResourceRequest) GetExecutorMemoryRequestOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutorMemoryRequest) {
		return nil, false
	}
	return o.ExecutorMemoryRequest, true
}

// HasExecutorMemoryRequest returns a boolean if a field has been set.
func (o *PredictionJobResourceRequest) HasExecutorMemoryRequest() bool {
	if o != nil && !IsNil(o.ExecutorMemoryRequest) {
		return true
	}

	return false
}

// SetExecutorMemoryRequest gets a reference to the given string and assigns it to the ExecutorMemoryRequest field.
func (o *PredictionJobResourceRequest) SetExecutorMemoryRequest(v string) {
	o.ExecutorMemoryRequest = &v
}

// GetExecutorReplica returns the ExecutorReplica field value if set, zero value otherwise.
func (o *PredictionJobResourceRequest) GetExecutorReplica() int32 {
	if o == nil || IsNil(o.ExecutorReplica) {
		var ret int32
		return ret
	}
	return *o.ExecutorReplica
}

// GetExecutorReplicaOk returns a tuple with the ExecutorReplica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionJobResourceRequest) GetExecutorReplicaOk() (*int32, bool) {
	if o == nil || IsNil(o.ExecutorReplica) {
		return nil, false
	}
	return o.ExecutorReplica, true
}

// HasExecutorReplica returns a boolean if a field has been set.
func (o *PredictionJobResourceRequest) HasExecutorReplica() bool {
	if o != nil && !IsNil(o.ExecutorReplica) {
		return true
	}

	return false
}

// SetExecutorReplica gets a reference to the given int32 and assigns it to the ExecutorReplica field.
func (o *PredictionJobResourceRequest) SetExecutorReplica(v int32) {
	o.ExecutorReplica = &v
}

func (o PredictionJobResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictionJobResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DriverCpuRequest) {
		toSerialize["driver_cpu_request"] = o.DriverCpuRequest
	}
	if !IsNil(o.DriverMemoryRequest) {
		toSerialize["driver_memory_request"] = o.DriverMemoryRequest
	}
	if !IsNil(o.ExecutorCpuRequest) {
		toSerialize["executor_cpu_request"] = o.ExecutorCpuRequest
	}
	if !IsNil(o.ExecutorMemoryRequest) {
		toSerialize["executor_memory_request"] = o.ExecutorMemoryRequest
	}
	if !IsNil(o.ExecutorReplica) {
		toSerialize["executor_replica"] = o.ExecutorReplica
	}
	return toSerialize, nil
}

type NullablePredictionJobResourceRequest struct {
	value *PredictionJobResourceRequest
	isSet bool
}

func (v NullablePredictionJobResourceRequest) Get() *PredictionJobResourceRequest {
	return v.value
}

func (v *NullablePredictionJobResourceRequest) Set(val *PredictionJobResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionJobResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionJobResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionJobResourceRequest(val *PredictionJobResourceRequest) *NullablePredictionJobResourceRequest {
	return &NullablePredictionJobResourceRequest{value: val, isSet: true}
}

func (v NullablePredictionJobResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionJobResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
