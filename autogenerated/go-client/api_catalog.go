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
)

// Linger please
var (
	_ _context.Context
)

// CatalogApiService CatalogApi service
type CatalogApiService service

type ApiCatalogGetRequest struct {
	ctx _context.Context
	ApiService *CatalogApiService
	xBrokerAPIVersion *string
	xBrokerAPIOriginatingIdentity *string
	xBrokerAPIRequestIdentity *string
}

func (r ApiCatalogGetRequest) XBrokerAPIVersion(xBrokerAPIVersion string) ApiCatalogGetRequest {
	r.xBrokerAPIVersion = &xBrokerAPIVersion
	return r
}
func (r ApiCatalogGetRequest) XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity string) ApiCatalogGetRequest {
	r.xBrokerAPIOriginatingIdentity = &xBrokerAPIOriginatingIdentity
	return r
}
func (r ApiCatalogGetRequest) XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity string) ApiCatalogGetRequest {
	r.xBrokerAPIRequestIdentity = &xBrokerAPIRequestIdentity
	return r
}

func (r ApiCatalogGetRequest) Execute() (Catalog, *_nethttp.Response, error) {
	return r.ApiService.CatalogGetExecute(r)
}

/*
 * CatalogGet get the catalog of services that the service broker offers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCatalogGetRequest
 */
func (a *CatalogApiService) CatalogGet(ctx _context.Context) ApiCatalogGetRequest {
	return ApiCatalogGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Catalog
 */
func (a *CatalogApiService) CatalogGetExecute(r ApiCatalogGetRequest) (Catalog, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Catalog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CatalogApiService.CatalogGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/catalog"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xBrokerAPIVersion == nil {
		return localVarReturnValue, nil, reportError("xBrokerAPIVersion is required and must be specified")
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
