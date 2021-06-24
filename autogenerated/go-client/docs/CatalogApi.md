# \CatalogApi

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogGet**](CatalogApi.md#CatalogGet) | **Get** /v2/catalog | get the catalog of services that the service broker offers



## CatalogGet

> Catalog CatalogGet(ctx).XBrokerAPIVersion(xBrokerAPIVersion).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity).Execute()

get the catalog of services that the service broker offers

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xBrokerAPIVersion := "xBrokerAPIVersion_example" // string | version number of the Service Broker API that the Platform will use
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    xBrokerAPIRequestIdentity := "xBrokerAPIRequestIdentity_example" // string | idenity of the request from the Platform (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CatalogApi.CatalogGet(context.Background()).XBrokerAPIVersion(xBrokerAPIVersion).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).XBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.CatalogGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CatalogGet`: Catalog
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.CatalogGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCatalogGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | 
 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **xBrokerAPIRequestIdentity** | **string** | idenity of the request from the Platform | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

