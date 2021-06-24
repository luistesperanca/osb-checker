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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ServiceBindingsApiService ServiceBindingsApi service
type ServiceBindingsApiService service

type ApiServiceBindingBindingRequest struct {
	ctx _context.Context
	ApiService *ServiceBindingsApiService
	xBrokerAPIVersion *string
	instanceId string
	bindingId string
	body *ServiceBindingRequest
	xBrokerAPIOriginatingIdentity *string
	xBrokerAPIRequestIdentity *string
	acceptsIncomplete *bool
}

func (r ApiServiceBindingBindingRequest) XBrokerAPIVersion(xBrokerAPIVersion string) ApiServiceBindingBindingRequest {
	r.xBrokerAPIVersion = &xBrokerAPIVersion
	return r
}
func (r ApiServiceBindingBindingRequest) Body(body ServiceBindingRequest) ApiServiceBindingBindingRequest {
	r.body = &body
	return r
}
func (r ApiServiceBindingBindingRequest) XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity string) ApiServiceBindingBindingRequest {
	r.xBrokerAPIOriginatingIdentity = &xBrokerAPIOriginatingIdentity
	return r
}
func (r ApiServiceBindingBindingRequest) XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity string) ApiServiceBindingBindingRequest {
	r.xBrokerAPIRequestIdentity = &xBrokerAPIRequestIdentity
	return r
}
func (r ApiServiceBindingBindingRequest) AcceptsIncomplete(acceptsIncomplete bool) ApiServiceBindingBindingRequest {
	r.acceptsIncomplete = &acceptsIncomplete
	return r
}

func (r ApiServiceBindingBindingRequest) Execute() (ServiceBindingResponse, *_nethttp.Response, error) {
	return r.ApiService.ServiceBindingBindingExecute(r)
}

/*
 * ServiceBindingBinding generation of a service binding
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param instanceId instance id of instance to provision
 * @param bindingId binding id of binding to create
 * @return ApiServiceBindingBindingRequest
 */
func (a *ServiceBindingsApiService) ServiceBindingBinding(ctx _context.Context, instanceId string, bindingId string) ApiServiceBindingBindingRequest {
	return ApiServiceBindingBindingRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		bindingId: bindingId,
	}
}

/*
 * Execute executes the request
 * @return ServiceBindingResponse
 */
func (a *ServiceBindingsApiService) ServiceBindingBindingExecute(r ApiServiceBindingBindingRequest) (ServiceBindingResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceBindingResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceBindingsApiService.ServiceBindingBinding")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"instance_id"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"binding_id"+"}", _neturl.PathEscape(parameterToString(r.bindingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xBrokerAPIVersion == nil {
		return localVarReturnValue, nil, reportError("xBrokerAPIVersion is required and must be specified")
	}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.acceptsIncomplete != nil {
		localVarQueryParams.Add("accepts_incomplete", parameterToString(*r.acceptsIncomplete, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Broker-API-Version"] = parameterToString(*r.xBrokerAPIVersion, "")
	if r.xBrokerAPIOriginatingIdentity != nil {
		localVarHeaderParams["X-Broker-API-Originating-Identity"] = parameterToString(*r.xBrokerAPIOriginatingIdentity, "")
	}
	if r.xBrokerAPIRequestIdentity != nil {
		localVarHeaderParams["X-Broker-API-Request-Identity"] = parameterToString(*r.xBrokerAPIRequestIdentity, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServiceBindingGetRequest struct {
	ctx _context.Context
	ApiService *ServiceBindingsApiService
	xBrokerAPIVersion *string
	instanceId string
	bindingId string
	xBrokerAPIOriginatingIdentity *string
	xBrokerAPIRequestIdentity *string
	serviceId *string
	planId *string
}

func (r ApiServiceBindingGetRequest) XBrokerAPIVersion(xBrokerAPIVersion string) ApiServiceBindingGetRequest {
	r.xBrokerAPIVersion = &xBrokerAPIVersion
	return r
}
func (r ApiServiceBindingGetRequest) XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity string) ApiServiceBindingGetRequest {
	r.xBrokerAPIOriginatingIdentity = &xBrokerAPIOriginatingIdentity
	return r
}
func (r ApiServiceBindingGetRequest) XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity string) ApiServiceBindingGetRequest {
	r.xBrokerAPIRequestIdentity = &xBrokerAPIRequestIdentity
	return r
}
func (r ApiServiceBindingGetRequest) ServiceId(serviceId string) ApiServiceBindingGetRequest {
	r.serviceId = &serviceId
	return r
}
func (r ApiServiceBindingGetRequest) PlanId(planId string) ApiServiceBindingGetRequest {
	r.planId = &planId
	return r
}

func (r ApiServiceBindingGetRequest) Execute() (ServiceBindingResource, *_nethttp.Response, error) {
	return r.ApiService.ServiceBindingGetExecute(r)
}

/*
 * ServiceBindingGet gets a service binding
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param instanceId instance id of instance to provision
 * @param bindingId binding id of binding to create
 * @return ApiServiceBindingGetRequest
 */
func (a *ServiceBindingsApiService) ServiceBindingGet(ctx _context.Context, instanceId string, bindingId string) ApiServiceBindingGetRequest {
	return ApiServiceBindingGetRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		bindingId: bindingId,
	}
}

/*
 * Execute executes the request
 * @return ServiceBindingResource
 */
func (a *ServiceBindingsApiService) ServiceBindingGetExecute(r ApiServiceBindingGetRequest) (ServiceBindingResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ServiceBindingResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceBindingsApiService.ServiceBindingGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"instance_id"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"binding_id"+"}", _neturl.PathEscape(parameterToString(r.bindingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xBrokerAPIVersion == nil {
		return localVarReturnValue, nil, reportError("xBrokerAPIVersion is required and must be specified")
	}

	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.planId != nil {
		localVarQueryParams.Add("plan_id", parameterToString(*r.planId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Broker-API-Version"] = parameterToString(*r.xBrokerAPIVersion, "")
	if r.xBrokerAPIOriginatingIdentity != nil {
		localVarHeaderParams["X-Broker-API-Originating-Identity"] = parameterToString(*r.xBrokerAPIOriginatingIdentity, "")
	}
	if r.xBrokerAPIRequestIdentity != nil {
		localVarHeaderParams["X-Broker-API-Request-Identity"] = parameterToString(*r.xBrokerAPIRequestIdentity, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServiceBindingLastOperationGetRequest struct {
	ctx _context.Context
	ApiService *ServiceBindingsApiService
	xBrokerAPIVersion *string
	instanceId string
	bindingId string
	xBrokerAPIOriginatingIdentity *string
	xBrokerAPIRequestIdentity *string
	serviceId *string
	planId *string
	operation *string
}

func (r ApiServiceBindingLastOperationGetRequest) XBrokerAPIVersion(xBrokerAPIVersion string) ApiServiceBindingLastOperationGetRequest {
	r.xBrokerAPIVersion = &xBrokerAPIVersion
	return r
}
func (r ApiServiceBindingLastOperationGetRequest) XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity string) ApiServiceBindingLastOperationGetRequest {
	r.xBrokerAPIOriginatingIdentity = &xBrokerAPIOriginatingIdentity
	return r
}
func (r ApiServiceBindingLastOperationGetRequest) XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity string) ApiServiceBindingLastOperationGetRequest {
	r.xBrokerAPIRequestIdentity = &xBrokerAPIRequestIdentity
	return r
}
func (r ApiServiceBindingLastOperationGetRequest) ServiceId(serviceId string) ApiServiceBindingLastOperationGetRequest {
	r.serviceId = &serviceId
	return r
}
func (r ApiServiceBindingLastOperationGetRequest) PlanId(planId string) ApiServiceBindingLastOperationGetRequest {
	r.planId = &planId
	return r
}
func (r ApiServiceBindingLastOperationGetRequest) Operation(operation string) ApiServiceBindingLastOperationGetRequest {
	r.operation = &operation
	return r
}

func (r ApiServiceBindingLastOperationGetRequest) Execute() (LastOperationResource, *_nethttp.Response, error) {
	return r.ApiService.ServiceBindingLastOperationGetExecute(r)
}

/*
 * ServiceBindingLastOperationGet last requested operation state for service binding
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param instanceId instance id of instance to provision
 * @param bindingId binding id of binding to create
 * @return ApiServiceBindingLastOperationGetRequest
 */
func (a *ServiceBindingsApiService) ServiceBindingLastOperationGet(ctx _context.Context, instanceId string, bindingId string) ApiServiceBindingLastOperationGetRequest {
	return ApiServiceBindingLastOperationGetRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		bindingId: bindingId,
	}
}

/*
 * Execute executes the request
 * @return LastOperationResource
 */
func (a *ServiceBindingsApiService) ServiceBindingLastOperationGetExecute(r ApiServiceBindingLastOperationGetRequest) (LastOperationResource, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  LastOperationResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceBindingsApiService.ServiceBindingLastOperationGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation"
	localVarPath = strings.Replace(localVarPath, "{"+"instance_id"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"binding_id"+"}", _neturl.PathEscape(parameterToString(r.bindingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xBrokerAPIVersion == nil {
		return localVarReturnValue, nil, reportError("xBrokerAPIVersion is required and must be specified")
	}

	if r.serviceId != nil {
		localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	}
	if r.planId != nil {
		localVarQueryParams.Add("plan_id", parameterToString(*r.planId, ""))
	}
	if r.operation != nil {
		localVarQueryParams.Add("operation", parameterToString(*r.operation, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Broker-API-Version"] = parameterToString(*r.xBrokerAPIVersion, "")
	if r.xBrokerAPIOriginatingIdentity != nil {
		localVarHeaderParams["X-Broker-API-Originating-Identity"] = parameterToString(*r.xBrokerAPIOriginatingIdentity, "")
	}
	if r.xBrokerAPIRequestIdentity != nil {
		localVarHeaderParams["X-Broker-API-Request-Identity"] = parameterToString(*r.xBrokerAPIRequestIdentity, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServiceBindingUnbindingRequest struct {
	ctx _context.Context
	ApiService *ServiceBindingsApiService
	xBrokerAPIVersion *string
	instanceId string
	bindingId string
	serviceId *string
	planId *string
	xBrokerAPIOriginatingIdentity *string
	xBrokerAPIRequestIdentity *string
	acceptsIncomplete *bool
}

func (r ApiServiceBindingUnbindingRequest) XBrokerAPIVersion(xBrokerAPIVersion string) ApiServiceBindingUnbindingRequest {
	r.xBrokerAPIVersion = &xBrokerAPIVersion
	return r
}
func (r ApiServiceBindingUnbindingRequest) ServiceId(serviceId string) ApiServiceBindingUnbindingRequest {
	r.serviceId = &serviceId
	return r
}
func (r ApiServiceBindingUnbindingRequest) PlanId(planId string) ApiServiceBindingUnbindingRequest {
	r.planId = &planId
	return r
}
func (r ApiServiceBindingUnbindingRequest) XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity string) ApiServiceBindingUnbindingRequest {
	r.xBrokerAPIOriginatingIdentity = &xBrokerAPIOriginatingIdentity
	return r
}
func (r ApiServiceBindingUnbindingRequest) XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity string) ApiServiceBindingUnbindingRequest {
	r.xBrokerAPIRequestIdentity = &xBrokerAPIRequestIdentity
	return r
}
func (r ApiServiceBindingUnbindingRequest) AcceptsIncomplete(acceptsIncomplete bool) ApiServiceBindingUnbindingRequest {
	r.acceptsIncomplete = &acceptsIncomplete
	return r
}

func (r ApiServiceBindingUnbindingRequest) Execute() (map[string]interface{}, *_nethttp.Response, error) {
	return r.ApiService.ServiceBindingUnbindingExecute(r)
}

/*
 * ServiceBindingUnbinding deprovision of a service binding
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param instanceId instance id of instance to provision
 * @param bindingId binding id of binding to create
 * @return ApiServiceBindingUnbindingRequest
 */
func (a *ServiceBindingsApiService) ServiceBindingUnbinding(ctx _context.Context, instanceId string, bindingId string) ApiServiceBindingUnbindingRequest {
	return ApiServiceBindingUnbindingRequest{
		ApiService: a,
		ctx: ctx,
		instanceId: instanceId,
		bindingId: bindingId,
	}
}

/*
 * Execute executes the request
 * @return map[string]interface{}
 */
func (a *ServiceBindingsApiService) ServiceBindingUnbindingExecute(r ApiServiceBindingUnbindingRequest) (map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceBindingsApiService.ServiceBindingUnbinding")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/service_instances/{instance_id}/service_bindings/{binding_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"instance_id"+"}", _neturl.PathEscape(parameterToString(r.instanceId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"binding_id"+"}", _neturl.PathEscape(parameterToString(r.bindingId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xBrokerAPIVersion == nil {
		return localVarReturnValue, nil, reportError("xBrokerAPIVersion is required and must be specified")
	}
	if r.serviceId == nil {
		return localVarReturnValue, nil, reportError("serviceId is required and must be specified")
	}
	if r.planId == nil {
		return localVarReturnValue, nil, reportError("planId is required and must be specified")
	}

	localVarQueryParams.Add("service_id", parameterToString(*r.serviceId, ""))
	localVarQueryParams.Add("plan_id", parameterToString(*r.planId, ""))
	if r.acceptsIncomplete != nil {
		localVarQueryParams.Add("accepts_incomplete", parameterToString(*r.acceptsIncomplete, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarHeaderParams["X-Broker-API-Version"] = parameterToString(*r.xBrokerAPIVersion, "")
	if r.xBrokerAPIOriginatingIdentity != nil {
		localVarHeaderParams["X-Broker-API-Originating-Identity"] = parameterToString(*r.xBrokerAPIOriginatingIdentity, "")
	}
	if r.xBrokerAPIRequestIdentity != nil {
		localVarHeaderParams["X-Broker-API-Request-Identity"] = parameterToString(*r.xBrokerAPIRequestIdentity, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 410 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 412 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
