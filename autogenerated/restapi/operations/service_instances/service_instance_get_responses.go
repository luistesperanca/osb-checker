// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
)

// ServiceInstanceGetOKCode is the HTTP code returned for type ServiceInstanceGetOK
const ServiceInstanceGetOKCode int = 200

/*ServiceInstanceGetOK OK

swagger:response serviceInstanceGetOK
*/
type ServiceInstanceGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceInstanceResource `json:"body,omitempty"`
}

// NewServiceInstanceGetOK creates ServiceInstanceGetOK with default headers values
func NewServiceInstanceGetOK() *ServiceInstanceGetOK {

	return &ServiceInstanceGetOK{}
}

// WithPayload adds the payload to the service instance get o k response
func (o *ServiceInstanceGetOK) WithPayload(payload *models.ServiceInstanceResource) *ServiceInstanceGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance get o k response
func (o *ServiceInstanceGetOK) SetPayload(payload *models.ServiceInstanceResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceGetUnauthorizedCode is the HTTP code returned for type ServiceInstanceGetUnauthorized
const ServiceInstanceGetUnauthorizedCode int = 401

/*ServiceInstanceGetUnauthorized Unauthorized

swagger:response serviceInstanceGetUnauthorized
*/
type ServiceInstanceGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceGetUnauthorized creates ServiceInstanceGetUnauthorized with default headers values
func NewServiceInstanceGetUnauthorized() *ServiceInstanceGetUnauthorized {

	return &ServiceInstanceGetUnauthorized{}
}

// WithPayload adds the payload to the service instance get unauthorized response
func (o *ServiceInstanceGetUnauthorized) WithPayload(payload *models.Error) *ServiceInstanceGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance get unauthorized response
func (o *ServiceInstanceGetUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceGetNotFoundCode is the HTTP code returned for type ServiceInstanceGetNotFound
const ServiceInstanceGetNotFoundCode int = 404

/*ServiceInstanceGetNotFound Not Found

swagger:response serviceInstanceGetNotFound
*/
type ServiceInstanceGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceGetNotFound creates ServiceInstanceGetNotFound with default headers values
func NewServiceInstanceGetNotFound() *ServiceInstanceGetNotFound {

	return &ServiceInstanceGetNotFound{}
}

// WithPayload adds the payload to the service instance get not found response
func (o *ServiceInstanceGetNotFound) WithPayload(payload *models.Error) *ServiceInstanceGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance get not found response
func (o *ServiceInstanceGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceInstanceGetPreconditionFailedCode is the HTTP code returned for type ServiceInstanceGetPreconditionFailed
const ServiceInstanceGetPreconditionFailedCode int = 412

/*ServiceInstanceGetPreconditionFailed Precondition Failed

swagger:response serviceInstanceGetPreconditionFailed
*/
type ServiceInstanceGetPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceGetPreconditionFailed creates ServiceInstanceGetPreconditionFailed with default headers values
func NewServiceInstanceGetPreconditionFailed() *ServiceInstanceGetPreconditionFailed {

	return &ServiceInstanceGetPreconditionFailed{}
}

// WithPayload adds the payload to the service instance get precondition failed response
func (o *ServiceInstanceGetPreconditionFailed) WithPayload(payload *models.Error) *ServiceInstanceGetPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance get precondition failed response
func (o *ServiceInstanceGetPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceGetPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ServiceInstanceGetDefault Unexpected

swagger:response serviceInstanceGetDefault
*/
type ServiceInstanceGetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceInstanceGetDefault creates ServiceInstanceGetDefault with default headers values
func NewServiceInstanceGetDefault(code int) *ServiceInstanceGetDefault {
	if code <= 0 {
		code = 500
	}

	return &ServiceInstanceGetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the service instance get default response
func (o *ServiceInstanceGetDefault) WithStatusCode(code int) *ServiceInstanceGetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the service instance get default response
func (o *ServiceInstanceGetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the service instance get default response
func (o *ServiceInstanceGetDefault) WithPayload(payload *models.Error) *ServiceInstanceGetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service instance get default response
func (o *ServiceInstanceGetDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceInstanceGetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
