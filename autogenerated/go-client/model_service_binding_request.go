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

// ServiceBindingRequest struct for ServiceBindingRequest
type ServiceBindingRequest struct {
	// See [Context Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#context-object) for more details.
	Context *map[string]interface{} `json:"context,omitempty"`
	ServiceId string `json:"service_id"`
	PlanId string `json:"plan_id"`
	AppGuid *string `json:"app_guid,omitempty"`
	BindResource *ServiceBindingResourceObject `json:"bind_resource,omitempty"`
	Parameters *map[string]interface{} `json:"parameters,omitempty"`
	PredecessorBindingId *string `json:"predecessor_binding_id,omitempty"`
}

// NewServiceBindingRequest instantiates a new ServiceBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingRequest(serviceId string, planId string) *ServiceBindingRequest {
	this := ServiceBindingRequest{}
	this.ServiceId = serviceId
	this.PlanId = planId
	return &this
}

// NewServiceBindingRequestWithDefaults instantiates a new ServiceBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingRequestWithDefaults() *ServiceBindingRequest {
	this := ServiceBindingRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ServiceBindingRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ServiceBindingRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *ServiceBindingRequest) SetContext(v map[string]interface{}) {
	o.Context = &v
}

// GetServiceId returns the ServiceId field value
func (o *ServiceBindingRequest) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetServiceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ServiceBindingRequest) SetServiceId(v string) {
	o.ServiceId = v
}

// GetPlanId returns the PlanId field value
func (o *ServiceBindingRequest) GetPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetPlanIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PlanId, true
}

// SetPlanId sets field value
func (o *ServiceBindingRequest) SetPlanId(v string) {
	o.PlanId = v
}

// GetAppGuid returns the AppGuid field value if set, zero value otherwise.
func (o *ServiceBindingRequest) GetAppGuid() string {
	if o == nil || o.AppGuid == nil {
		var ret string
		return ret
	}
	return *o.AppGuid
}

// GetAppGuidOk returns a tuple with the AppGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetAppGuidOk() (*string, bool) {
	if o == nil || o.AppGuid == nil {
		return nil, false
	}
	return o.AppGuid, true
}

// HasAppGuid returns a boolean if a field has been set.
func (o *ServiceBindingRequest) HasAppGuid() bool {
	if o != nil && o.AppGuid != nil {
		return true
	}

	return false
}

// SetAppGuid gets a reference to the given string and assigns it to the AppGuid field.
func (o *ServiceBindingRequest) SetAppGuid(v string) {
	o.AppGuid = &v
}

// GetBindResource returns the BindResource field value if set, zero value otherwise.
func (o *ServiceBindingRequest) GetBindResource() ServiceBindingResourceObject {
	if o == nil || o.BindResource == nil {
		var ret ServiceBindingResourceObject
		return ret
	}
	return *o.BindResource
}

// GetBindResourceOk returns a tuple with the BindResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetBindResourceOk() (*ServiceBindingResourceObject, bool) {
	if o == nil || o.BindResource == nil {
		return nil, false
	}
	return o.BindResource, true
}

// HasBindResource returns a boolean if a field has been set.
func (o *ServiceBindingRequest) HasBindResource() bool {
	if o != nil && o.BindResource != nil {
		return true
	}

	return false
}

// SetBindResource gets a reference to the given ServiceBindingResourceObject and assigns it to the BindResource field.
func (o *ServiceBindingRequest) SetBindResource(v ServiceBindingResourceObject) {
	o.BindResource = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ServiceBindingRequest) GetParameters() map[string]interface{} {
	if o == nil || o.Parameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ServiceBindingRequest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *ServiceBindingRequest) SetParameters(v map[string]interface{}) {
	o.Parameters = &v
}

// GetPredecessorBindingId returns the PredecessorBindingId field value if set, zero value otherwise.
func (o *ServiceBindingRequest) GetPredecessorBindingId() string {
	if o == nil || o.PredecessorBindingId == nil {
		var ret string
		return ret
	}
	return *o.PredecessorBindingId
}

// GetPredecessorBindingIdOk returns a tuple with the PredecessorBindingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingRequest) GetPredecessorBindingIdOk() (*string, bool) {
	if o == nil || o.PredecessorBindingId == nil {
		return nil, false
	}
	return o.PredecessorBindingId, true
}

// HasPredecessorBindingId returns a boolean if a field has been set.
func (o *ServiceBindingRequest) HasPredecessorBindingId() bool {
	if o != nil && o.PredecessorBindingId != nil {
		return true
	}

	return false
}

// SetPredecessorBindingId gets a reference to the given string and assigns it to the PredecessorBindingId field.
func (o *ServiceBindingRequest) SetPredecessorBindingId(v string) {
	o.PredecessorBindingId = &v
}

func (o ServiceBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["service_id"] = o.ServiceId
	}
	if true {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.AppGuid != nil {
		toSerialize["app_guid"] = o.AppGuid
	}
	if o.BindResource != nil {
		toSerialize["bind_resource"] = o.BindResource
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.PredecessorBindingId != nil {
		toSerialize["predecessor_binding_id"] = o.PredecessorBindingId
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBindingRequest struct {
	value *ServiceBindingRequest
	isSet bool
}

func (v NullableServiceBindingRequest) Get() *ServiceBindingRequest {
	return v.value
}

func (v *NullableServiceBindingRequest) Set(val *ServiceBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingRequest(val *ServiceBindingRequest) *NullableServiceBindingRequest {
	return &NullableServiceBindingRequest{value: val, isSet: true}
}

func (v NullableServiceBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


