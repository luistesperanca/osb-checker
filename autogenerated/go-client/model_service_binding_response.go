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

// ServiceBindingResponse struct for ServiceBindingResponse
type ServiceBindingResponse struct {
	Metadata *ServiceBindingMetadata `json:"metadata,omitempty"`
	Credentials *map[string]interface{} `json:"credentials,omitempty"`
	SyslogDrainUrl *string `json:"syslog_drain_url,omitempty"`
	RouteServiceUrl *string `json:"route_service_url,omitempty"`
	VolumeMounts *[]ServiceBindingVolumeMount `json:"volume_mounts,omitempty"`
	Endpoints *[]ServiceBindingEndpoint `json:"endpoints,omitempty"`
}

// NewServiceBindingResponse instantiates a new ServiceBindingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingResponse() *ServiceBindingResponse {
	this := ServiceBindingResponse{}
	return &this
}

// NewServiceBindingResponseWithDefaults instantiates a new ServiceBindingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingResponseWithDefaults() *ServiceBindingResponse {
	this := ServiceBindingResponse{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceBindingResponse) GetMetadata() ServiceBindingMetadata {
	if o == nil || o.Metadata == nil {
		var ret ServiceBindingMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResponse) GetMetadataOk() (*ServiceBindingMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceBindingResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ServiceBindingMetadata and assigns it to the Metadata field.
func (o *ServiceBindingResponse) SetMetadata(v ServiceBindingMetadata) {
	o.Metadata = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ServiceBindingResponse) GetCredentials() map[string]interface{} {
	if o == nil || o.Credentials == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResponse) GetCredentialsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ServiceBindingResponse) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *ServiceBindingResponse) SetCredentials(v map[string]interface{}) {
	o.Credentials = &v
}

// GetSyslogDrainUrl returns the SyslogDrainUrl field value if set, zero value otherwise.
func (o *ServiceBindingResponse) GetSyslogDrainUrl() string {
	if o == nil || o.SyslogDrainUrl == nil {
		var ret string
		return ret
	}
	return *o.SyslogDrainUrl
}

// GetSyslogDrainUrlOk returns a tuple with the SyslogDrainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResponse) GetSyslogDrainUrlOk() (*string, bool) {
	if o == nil || o.SyslogDrainUrl == nil {
		return nil, false
	}
	return o.SyslogDrainUrl, true
}

// HasSyslogDrainUrl returns a boolean if a field has been set.
func (o *ServiceBindingResponse) HasSyslogDrainUrl() bool {
	if o != nil && o.SyslogDrainUrl != nil {
		return true
	}

	return false
}

// SetSyslogDrainUrl gets a reference to the given string and assigns it to the SyslogDrainUrl field.
func (o *ServiceBindingResponse) SetSyslogDrainUrl(v string) {
	o.SyslogDrainUrl = &v
}

// GetRouteServiceUrl returns the RouteServiceUrl field value if set, zero value otherwise.
func (o *ServiceBindingResponse) GetRouteServiceUrl() string {
	if o == nil || o.RouteServiceUrl == nil {
		var ret string
		return ret
	}
	return *o.RouteServiceUrl
}

// GetRouteServiceUrlOk returns a tuple with the RouteServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResponse) GetRouteServiceUrlOk() (*string, bool) {
	if o == nil || o.RouteServiceUrl == nil {
		return nil, false
	}
	return o.RouteServiceUrl, true
}

// HasRouteServiceUrl returns a boolean if a field has been set.
func (o *ServiceBindingResponse) HasRouteServiceUrl() bool {
	if o != nil && o.RouteServiceUrl != nil {
		return true
	}

	return false
}

// SetRouteServiceUrl gets a reference to the given string and assigns it to the RouteServiceUrl field.
func (o *ServiceBindingResponse) SetRouteServiceUrl(v string) {
	o.RouteServiceUrl = &v
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *ServiceBindingResponse) GetVolumeMounts() []ServiceBindingVolumeMount {
	if o == nil || o.VolumeMounts == nil {
		var ret []ServiceBindingVolumeMount
		return ret
	}
	return *o.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResponse) GetVolumeMountsOk() (*[]ServiceBindingVolumeMount, bool) {
	if o == nil || o.VolumeMounts == nil {
		return nil, false
	}
	return o.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *ServiceBindingResponse) HasVolumeMounts() bool {
	if o != nil && o.VolumeMounts != nil {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given []ServiceBindingVolumeMount and assigns it to the VolumeMounts field.
func (o *ServiceBindingResponse) SetVolumeMounts(v []ServiceBindingVolumeMount) {
	o.VolumeMounts = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ServiceBindingResponse) GetEndpoints() []ServiceBindingEndpoint {
	if o == nil || o.Endpoints == nil {
		var ret []ServiceBindingEndpoint
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResponse) GetEndpointsOk() (*[]ServiceBindingEndpoint, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ServiceBindingResponse) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ServiceBindingEndpoint and assigns it to the Endpoints field.
func (o *ServiceBindingResponse) SetEndpoints(v []ServiceBindingEndpoint) {
	o.Endpoints = &v
}

func (o ServiceBindingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.SyslogDrainUrl != nil {
		toSerialize["syslog_drain_url"] = o.SyslogDrainUrl
	}
	if o.RouteServiceUrl != nil {
		toSerialize["route_service_url"] = o.RouteServiceUrl
	}
	if o.VolumeMounts != nil {
		toSerialize["volume_mounts"] = o.VolumeMounts
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBindingResponse struct {
	value *ServiceBindingResponse
	isSet bool
}

func (v NullableServiceBindingResponse) Get() *ServiceBindingResponse {
	return v.value
}

func (v *NullableServiceBindingResponse) Set(val *ServiceBindingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingResponse(val *ServiceBindingResponse) *NullableServiceBindingResponse {
	return &NullableServiceBindingResponse{value: val, isSet: true}
}

func (v NullableServiceBindingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


