# ServiceBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** | See [Context Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#context-object) for more details. | [optional] 
**ServiceId** | **string** |  | 
**PlanId** | **string** |  | 
**AppGuid** | Pointer to **string** |  | [optional] 
**BindResource** | Pointer to [**ServiceBindingResourceObject**](ServiceBindingResourceObject.md) |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**PredecessorBindingId** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceBindingRequest

`func NewServiceBindingRequest(serviceId string, planId string, ) *ServiceBindingRequest`

NewServiceBindingRequest instantiates a new ServiceBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBindingRequestWithDefaults

`func NewServiceBindingRequestWithDefaults() *ServiceBindingRequest`

NewServiceBindingRequestWithDefaults instantiates a new ServiceBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ServiceBindingRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ServiceBindingRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ServiceBindingRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ServiceBindingRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceBindingRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceBindingRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceBindingRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPlanId

`func (o *ServiceBindingRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ServiceBindingRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ServiceBindingRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetAppGuid

`func (o *ServiceBindingRequest) GetAppGuid() string`

GetAppGuid returns the AppGuid field if non-nil, zero value otherwise.

### GetAppGuidOk

`func (o *ServiceBindingRequest) GetAppGuidOk() (*string, bool)`

GetAppGuidOk returns a tuple with the AppGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppGuid

`func (o *ServiceBindingRequest) SetAppGuid(v string)`

SetAppGuid sets AppGuid field to given value.

### HasAppGuid

`func (o *ServiceBindingRequest) HasAppGuid() bool`

HasAppGuid returns a boolean if a field has been set.

### GetBindResource

`func (o *ServiceBindingRequest) GetBindResource() ServiceBindingResourceObject`

GetBindResource returns the BindResource field if non-nil, zero value otherwise.

### GetBindResourceOk

`func (o *ServiceBindingRequest) GetBindResourceOk() (*ServiceBindingResourceObject, bool)`

GetBindResourceOk returns a tuple with the BindResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindResource

`func (o *ServiceBindingRequest) SetBindResource(v ServiceBindingResourceObject)`

SetBindResource sets BindResource field to given value.

### HasBindResource

`func (o *ServiceBindingRequest) HasBindResource() bool`

HasBindResource returns a boolean if a field has been set.

### GetParameters

`func (o *ServiceBindingRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceBindingRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceBindingRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceBindingRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPredecessorBindingId

`func (o *ServiceBindingRequest) GetPredecessorBindingId() string`

GetPredecessorBindingId returns the PredecessorBindingId field if non-nil, zero value otherwise.

### GetPredecessorBindingIdOk

`func (o *ServiceBindingRequest) GetPredecessorBindingIdOk() (*string, bool)`

GetPredecessorBindingIdOk returns a tuple with the PredecessorBindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredecessorBindingId

`func (o *ServiceBindingRequest) SetPredecessorBindingId(v string)`

SetPredecessorBindingId sets PredecessorBindingId field to given value.

### HasPredecessorBindingId

`func (o *ServiceBindingRequest) HasPredecessorBindingId() bool`

HasPredecessorBindingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


