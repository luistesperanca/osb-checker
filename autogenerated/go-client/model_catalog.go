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

// Catalog struct for Catalog
type Catalog struct {
	Services *[]Service `json:"services,omitempty"`
}

// NewCatalog instantiates a new Catalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalog() *Catalog {
	this := Catalog{}
	return &this
}

// NewCatalogWithDefaults instantiates a new Catalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogWithDefaults() *Catalog {
	this := Catalog{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Catalog) GetServices() []Service {
	if o == nil || o.Services == nil {
		var ret []Service
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Catalog) GetServicesOk() (*[]Service, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Catalog) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []Service and assigns it to the Services field.
func (o *Catalog) SetServices(v []Service) {
	o.Services = &v
}

func (o Catalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableCatalog struct {
	value *Catalog
	isSet bool
}

func (v NullableCatalog) Get() *Catalog {
	return v.value
}

func (v *NullableCatalog) Set(val *Catalog) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalog(val *Catalog) *NullableCatalog {
	return &NullableCatalog{value: val, isSet: true}
}

func (v NullableCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


