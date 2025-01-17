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

// ServiceInstancePreviousValues struct for ServiceInstancePreviousValues
type ServiceInstancePreviousValues struct {
	ServiceId *string `json:"service_id,omitempty"`
	PlanId *string `json:"plan_id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	SpaceId *string `json:"space_id,omitempty"`
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`
}

// NewServiceInstancePreviousValues instantiates a new ServiceInstancePreviousValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInstancePreviousValues() *ServiceInstancePreviousValues {
	this := ServiceInstancePreviousValues{}
	return &this
}

// NewServiceInstancePreviousValuesWithDefaults instantiates a new ServiceInstancePreviousValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInstancePreviousValuesWithDefaults() *ServiceInstancePreviousValues {
	this := ServiceInstancePreviousValues{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServiceInstancePreviousValues) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstancePreviousValues) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServiceInstancePreviousValues) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ServiceInstancePreviousValues) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetPlanId returns the PlanId field value if set, zero value otherwise.
func (o *ServiceInstancePreviousValues) GetPlanId() string {
	if o == nil || o.PlanId == nil {
		var ret string
		return ret
	}
	return *o.PlanId
}

// GetPlanIdOk returns a tuple with the PlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstancePreviousValues) GetPlanIdOk() (*string, bool) {
	if o == nil || o.PlanId == nil {
		return nil, false
	}
	return o.PlanId, true
}

// HasPlanId returns a boolean if a field has been set.
func (o *ServiceInstancePreviousValues) HasPlanId() bool {
	if o != nil && o.PlanId != nil {
		return true
	}

	return false
}

// SetPlanId gets a reference to the given string and assigns it to the PlanId field.
func (o *ServiceInstancePreviousValues) SetPlanId(v string) {
	o.PlanId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *ServiceInstancePreviousValues) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstancePreviousValues) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *ServiceInstancePreviousValues) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *ServiceInstancePreviousValues) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *ServiceInstancePreviousValues) GetSpaceId() string {
	if o == nil || o.SpaceId == nil {
		var ret string
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstancePreviousValues) GetSpaceIdOk() (*string, bool) {
	if o == nil || o.SpaceId == nil {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *ServiceInstancePreviousValues) HasSpaceId() bool {
	if o != nil && o.SpaceId != nil {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given string and assigns it to the SpaceId field.
func (o *ServiceInstancePreviousValues) SetSpaceId(v string) {
	o.SpaceId = &v
}

// GetMaintenanceInfo returns the MaintenanceInfo field value if set, zero value otherwise.
func (o *ServiceInstancePreviousValues) GetMaintenanceInfo() MaintenanceInfo {
	if o == nil || o.MaintenanceInfo == nil {
		var ret MaintenanceInfo
		return ret
	}
	return *o.MaintenanceInfo
}

// GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstancePreviousValues) GetMaintenanceInfoOk() (*MaintenanceInfo, bool) {
	if o == nil || o.MaintenanceInfo == nil {
		return nil, false
	}
	return o.MaintenanceInfo, true
}

// HasMaintenanceInfo returns a boolean if a field has been set.
func (o *ServiceInstancePreviousValues) HasMaintenanceInfo() bool {
	if o != nil && o.MaintenanceInfo != nil {
		return true
	}

	return false
}

// SetMaintenanceInfo gets a reference to the given MaintenanceInfo and assigns it to the MaintenanceInfo field.
func (o *ServiceInstancePreviousValues) SetMaintenanceInfo(v MaintenanceInfo) {
	o.MaintenanceInfo = &v
}

func (o ServiceInstancePreviousValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.PlanId != nil {
		toSerialize["plan_id"] = o.PlanId
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.SpaceId != nil {
		toSerialize["space_id"] = o.SpaceId
	}
	if o.MaintenanceInfo != nil {
		toSerialize["maintenance_info"] = o.MaintenanceInfo
	}
	return json.Marshal(toSerialize)
}

type NullableServiceInstancePreviousValues struct {
	value *ServiceInstancePreviousValues
	isSet bool
}

func (v NullableServiceInstancePreviousValues) Get() *ServiceInstancePreviousValues {
	return v.value
}

func (v *NullableServiceInstancePreviousValues) Set(val *ServiceInstancePreviousValues) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInstancePreviousValues) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInstancePreviousValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInstancePreviousValues(val *ServiceInstancePreviousValues) *NullableServiceInstancePreviousValues {
	return &NullableServiceInstancePreviousValues{value: val, isSet: true}
}

func (v NullableServiceInstancePreviousValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInstancePreviousValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


