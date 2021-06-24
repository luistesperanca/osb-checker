/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ServiceInstanceMetadata struct for ServiceInstanceMetadata
type ServiceInstanceMetadata struct {
	Labels *map[string]interface{} `json:"labels,omitempty"`
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
}

// NewServiceInstanceMetadata instantiates a new ServiceInstanceMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInstanceMetadata() *ServiceInstanceMetadata {
	this := ServiceInstanceMetadata{}
	return &this
}

// NewServiceInstanceMetadataWithDefaults instantiates a new ServiceInstanceMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInstanceMetadataWithDefaults() *ServiceInstanceMetadata {
	this := ServiceInstanceMetadata{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ServiceInstanceMetadata) GetLabels() map[string]interface{} {
	if o == nil || o.Labels == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceMetadata) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ServiceInstanceMetadata) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *ServiceInstanceMetadata) SetLabels(v map[string]interface{}) {
	o.Labels = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ServiceInstanceMetadata) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceMetadata) GetAttributesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ServiceInstanceMetadata) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *ServiceInstanceMetadata) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

func (o ServiceInstanceMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableServiceInstanceMetadata struct {
	value *ServiceInstanceMetadata
	isSet bool
}

func (v NullableServiceInstanceMetadata) Get() *ServiceInstanceMetadata {
	return v.value
}

func (v *NullableServiceInstanceMetadata) Set(val *ServiceInstanceMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInstanceMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInstanceMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInstanceMetadata(val *ServiceInstanceMetadata) *NullableServiceInstanceMetadata {
	return &NullableServiceInstanceMetadata{value: val, isSet: true}
}

func (v NullableServiceInstanceMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInstanceMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


