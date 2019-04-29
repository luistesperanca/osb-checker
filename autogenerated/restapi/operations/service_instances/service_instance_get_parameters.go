// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewServiceInstanceGetParams creates a new ServiceInstanceGetParams object
// no default values defined in spec.
func NewServiceInstanceGetParams() ServiceInstanceGetParams {

	return ServiceInstanceGetParams{}
}

// ServiceInstanceGetParams contains all the bound params for the service instance get operation
// typically these are obtained from a http.Request
//
// swagger:parameters serviceInstance.get
type ServiceInstanceGetParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*identity of the user that initiated the request from the Platform
	  In: header
	*/
	XBrokerAPIOriginatingIdentity *string
	/*idenity of the request from the Platform
	  In: header
	*/
	XBrokerAPIRequestIdentity *string
	/*version number of the Service Broker API that the Platform will use
	  Required: true
	  In: header
	*/
	XBrokerAPIVersion string
	/*instance id of instance to provision
	  Required: true
	  In: path
	*/
	InstanceID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewServiceInstanceGetParams() beforehand.
func (o *ServiceInstanceGetParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXBrokerAPIOriginatingIdentity(r.Header[http.CanonicalHeaderKey("X-Broker-API-Originating-Identity")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXBrokerAPIRequestIdentity(r.Header[http.CanonicalHeaderKey("X-Broker-API-Request-Identity")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindXBrokerAPIVersion(r.Header[http.CanonicalHeaderKey("X-Broker-API-Version")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rInstanceID, rhkInstanceID, _ := route.Params.GetOK("instance_id")
	if err := o.bindInstanceID(rInstanceID, rhkInstanceID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXBrokerAPIOriginatingIdentity binds and validates parameter XBrokerAPIOriginatingIdentity from header.
func (o *ServiceInstanceGetParams) bindXBrokerAPIOriginatingIdentity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XBrokerAPIOriginatingIdentity = &raw

	return nil
}

// bindXBrokerAPIRequestIdentity binds and validates parameter XBrokerAPIRequestIdentity from header.
func (o *ServiceInstanceGetParams) bindXBrokerAPIRequestIdentity(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.XBrokerAPIRequestIdentity = &raw

	return nil
}

// bindXBrokerAPIVersion binds and validates parameter XBrokerAPIVersion from header.
func (o *ServiceInstanceGetParams) bindXBrokerAPIVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-Broker-API-Version", "header")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-Broker-API-Version", "header", raw); err != nil {
		return err
	}

	o.XBrokerAPIVersion = raw

	return nil
}

// bindInstanceID binds and validates parameter InstanceID from path.
func (o *ServiceInstanceGetParams) bindInstanceID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.InstanceID = raw

	return nil
}
