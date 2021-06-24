# ServiceBindingEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Ports** | **[]string** |  | 
**Protocol** | Pointer to **string** |  | [optional] [default to "tcp"]

## Methods

### NewServiceBindingEndpoint

`func NewServiceBindingEndpoint(host string, ports []string, ) *ServiceBindingEndpoint`

NewServiceBindingEndpoint instantiates a new ServiceBindingEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBindingEndpointWithDefaults

`func NewServiceBindingEndpointWithDefaults() *ServiceBindingEndpoint`

NewServiceBindingEndpointWithDefaults instantiates a new ServiceBindingEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ServiceBindingEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ServiceBindingEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ServiceBindingEndpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetPorts

`func (o *ServiceBindingEndpoint) GetPorts() []string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ServiceBindingEndpoint) GetPortsOk() (*[]string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ServiceBindingEndpoint) SetPorts(v []string)`

SetPorts sets Ports field to given value.


### GetProtocol

`func (o *ServiceBindingEndpoint) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServiceBindingEndpoint) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServiceBindingEndpoint) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServiceBindingEndpoint) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


