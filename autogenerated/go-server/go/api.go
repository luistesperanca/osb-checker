/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
)



// CatalogApiRouter defines the required methods for binding the api requests to a responses for the CatalogApi
// The CatalogApiRouter implementation should parse necessary information from the http request, 
// pass the data to a CatalogApiServicer to perform the required actions, then write the service results to the http response.
type CatalogApiRouter interface { 
	CatalogGet(http.ResponseWriter, *http.Request)
}
// ServiceBindingsApiRouter defines the required methods for binding the api requests to a responses for the ServiceBindingsApi
// The ServiceBindingsApiRouter implementation should parse necessary information from the http request, 
// pass the data to a ServiceBindingsApiServicer to perform the required actions, then write the service results to the http response.
type ServiceBindingsApiRouter interface { 
	ServiceBindingBinding(http.ResponseWriter, *http.Request)
	ServiceBindingGet(http.ResponseWriter, *http.Request)
	ServiceBindingLastOperationGet(http.ResponseWriter, *http.Request)
	ServiceBindingUnbinding(http.ResponseWriter, *http.Request)
}
// ServiceInstancesApiRouter defines the required methods for binding the api requests to a responses for the ServiceInstancesApi
// The ServiceInstancesApiRouter implementation should parse necessary information from the http request, 
// pass the data to a ServiceInstancesApiServicer to perform the required actions, then write the service results to the http response.
type ServiceInstancesApiRouter interface { 
	ServiceInstanceDeprovision(http.ResponseWriter, *http.Request)
	ServiceInstanceGet(http.ResponseWriter, *http.Request)
	ServiceInstanceLastOperationGet(http.ResponseWriter, *http.Request)
	ServiceInstanceProvision(http.ResponseWriter, *http.Request)
	ServiceInstanceUpdate(http.ResponseWriter, *http.Request)
}


// CatalogApiServicer defines the api actions for the CatalogApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type CatalogApiServicer interface { 
	CatalogGet(context.Context, string, string, string) (ImplResponse, error)
}


// ServiceBindingsApiServicer defines the api actions for the ServiceBindingsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type ServiceBindingsApiServicer interface { 
	ServiceBindingBinding(context.Context, string, string, string, ServiceBindingRequest, string, string, bool) (ImplResponse, error)
	ServiceBindingGet(context.Context, string, string, string, string, string, string, string) (ImplResponse, error)
	ServiceBindingLastOperationGet(context.Context, string, string, string, string, string, string, string, string) (ImplResponse, error)
	ServiceBindingUnbinding(context.Context, string, string, string, string, string, string, string, bool) (ImplResponse, error)
}


// ServiceInstancesApiServicer defines the api actions for the ServiceInstancesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it, 
// while the service implementation can ignored with the .openapi-generator-ignore file 
// and updated with the logic required for the API.
type ServiceInstancesApiServicer interface { 
	ServiceInstanceDeprovision(context.Context, string, string, string, string, string, string, bool) (ImplResponse, error)
	ServiceInstanceGet(context.Context, string, string, string, string, string, string) (ImplResponse, error)
	ServiceInstanceLastOperationGet(context.Context, string, string, string, string, string, string, string) (ImplResponse, error)
	ServiceInstanceProvision(context.Context, string, string, ServiceInstanceProvisionRequest, string, string, bool) (ImplResponse, error)
	ServiceInstanceUpdate(context.Context, string, string, ServiceInstanceUpdateRequest, string, string, bool) (ImplResponse, error)
}