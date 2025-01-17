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

// ServiceInstanceSchemaObject struct for ServiceInstanceSchemaObject
type ServiceInstanceSchemaObject struct {
	Create *SchemaParameters `json:"create,omitempty"`
	Update *SchemaParameters `json:"update,omitempty"`
}

// NewServiceInstanceSchemaObject instantiates a new ServiceInstanceSchemaObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInstanceSchemaObject() *ServiceInstanceSchemaObject {
	this := ServiceInstanceSchemaObject{}
	return &this
}

// NewServiceInstanceSchemaObjectWithDefaults instantiates a new ServiceInstanceSchemaObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInstanceSchemaObjectWithDefaults() *ServiceInstanceSchemaObject {
	this := ServiceInstanceSchemaObject{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ServiceInstanceSchemaObject) GetCreate() SchemaParameters {
	if o == nil || o.Create == nil {
		var ret SchemaParameters
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceSchemaObject) GetCreateOk() (*SchemaParameters, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ServiceInstanceSchemaObject) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given SchemaParameters and assigns it to the Create field.
func (o *ServiceInstanceSchemaObject) SetCreate(v SchemaParameters) {
	o.Create = &v
}

// GetUpdate returns the Update field value if set, zero value otherwise.
func (o *ServiceInstanceSchemaObject) GetUpdate() SchemaParameters {
	if o == nil || o.Update == nil {
		var ret SchemaParameters
		return ret
	}
	return *o.Update
}

// GetUpdateOk returns a tuple with the Update field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceSchemaObject) GetUpdateOk() (*SchemaParameters, bool) {
	if o == nil || o.Update == nil {
		return nil, false
	}
	return o.Update, true
}

// HasUpdate returns a boolean if a field has been set.
func (o *ServiceInstanceSchemaObject) HasUpdate() bool {
	if o != nil && o.Update != nil {
		return true
	}

	return false
}

// SetUpdate gets a reference to the given SchemaParameters and assigns it to the Update field.
func (o *ServiceInstanceSchemaObject) SetUpdate(v SchemaParameters) {
	o.Update = &v
}

func (o ServiceInstanceSchemaObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	if o.Update != nil {
		toSerialize["update"] = o.Update
	}
	return json.Marshal(toSerialize)
}

type NullableServiceInstanceSchemaObject struct {
	value *ServiceInstanceSchemaObject
	isSet bool
}

func (v NullableServiceInstanceSchemaObject) Get() *ServiceInstanceSchemaObject {
	return v.value
}

func (v *NullableServiceInstanceSchemaObject) Set(val *ServiceInstanceSchemaObject) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInstanceSchemaObject) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInstanceSchemaObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInstanceSchemaObject(val *ServiceInstanceSchemaObject) *NullableServiceInstanceSchemaObject {
	return &NullableServiceInstanceSchemaObject{value: val, isSet: true}
}

func (v NullableServiceInstanceSchemaObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInstanceSchemaObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


