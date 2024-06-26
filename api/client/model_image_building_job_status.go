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

// checks if the ImageBuildingJobStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageBuildingJobStatus{}

// ImageBuildingJobStatus struct for ImageBuildingJobStatus
type ImageBuildingJobStatus struct {
	State   *ImageBuildingJobState `json:"state,omitempty"`
	Message *string                `json:"message,omitempty"`
}

// NewImageBuildingJobStatus instantiates a new ImageBuildingJobStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBuildingJobStatus() *ImageBuildingJobStatus {
	this := ImageBuildingJobStatus{}
	return &this
}

// NewImageBuildingJobStatusWithDefaults instantiates a new ImageBuildingJobStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBuildingJobStatusWithDefaults() *ImageBuildingJobStatus {
	this := ImageBuildingJobStatus{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ImageBuildingJobStatus) GetState() ImageBuildingJobState {
	if o == nil || IsNil(o.State) {
		var ret ImageBuildingJobState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildingJobStatus) GetStateOk() (*ImageBuildingJobState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ImageBuildingJobStatus) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ImageBuildingJobState and assigns it to the State field.
func (o *ImageBuildingJobStatus) SetState(v ImageBuildingJobState) {
	o.State = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ImageBuildingJobStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildingJobStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ImageBuildingJobStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ImageBuildingJobStatus) SetMessage(v string) {
	o.Message = &v
}

func (o ImageBuildingJobStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageBuildingJobStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableImageBuildingJobStatus struct {
	value *ImageBuildingJobStatus
	isSet bool
}

func (v NullableImageBuildingJobStatus) Get() *ImageBuildingJobStatus {
	return v.value
}

func (v *NullableImageBuildingJobStatus) Set(val *ImageBuildingJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuildingJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuildingJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuildingJobStatus(val *ImageBuildingJobStatus) *NullableImageBuildingJobStatus {
	return &NullableImageBuildingJobStatus{value: val, isSet: true}
}

func (v NullableImageBuildingJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuildingJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
