# ServiceInstanceUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** | See [Context Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#context-object) for more details. | [optional] 
**ServiceId** | **string** |  | 
**PlanId** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**PreviousValues** | Pointer to [**ServiceInstancePreviousValues**](ServiceInstancePreviousValues.md) |  | [optional] 
**MaintenanceInfo** | Pointer to [**MaintenanceInfo**](MaintenanceInfo.md) |  | [optional] 

## Methods

### NewServiceInstanceUpdateRequest

`func NewServiceInstanceUpdateRequest(serviceId string, ) *ServiceInstanceUpdateRequest`

NewServiceInstanceUpdateRequest instantiates a new ServiceInstanceUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceUpdateRequestWithDefaults

`func NewServiceInstanceUpdateRequestWithDefaults() *ServiceInstanceUpdateRequest`

NewServiceInstanceUpdateRequestWithDefaults instantiates a new ServiceInstanceUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ServiceInstanceUpdateRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ServiceInstanceUpdateRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ServiceInstanceUpdateRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ServiceInstanceUpdateRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetServiceId

`func (o *ServiceInstanceUpdateRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceInstanceUpdateRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceInstanceUpdateRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPlanId

`func (o *ServiceInstanceUpdateRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ServiceInstanceUpdateRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ServiceInstanceUpdateRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ServiceInstanceUpdateRequest) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetParameters

`func (o *ServiceInstanceUpdateRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceInstanceUpdateRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceInstanceUpdateRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceInstanceUpdateRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPreviousValues

`func (o *ServiceInstanceUpdateRequest) GetPreviousValues() ServiceInstancePreviousValues`

GetPreviousValues returns the PreviousValues field if non-nil, zero value otherwise.

### GetPreviousValuesOk

`func (o *ServiceInstanceUpdateRequest) GetPreviousValuesOk() (*ServiceInstancePreviousValues, bool)`

GetPreviousValuesOk returns a tuple with the PreviousValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousValues

`func (o *ServiceInstanceUpdateRequest) SetPreviousValues(v ServiceInstancePreviousValues)`

SetPreviousValues sets PreviousValues field to given value.

### HasPreviousValues

`func (o *ServiceInstanceUpdateRequest) HasPreviousValues() bool`

HasPreviousValues returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *ServiceInstanceUpdateRequest) GetMaintenanceInfo() MaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *ServiceInstanceUpdateRequest) GetMaintenanceInfoOk() (*MaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *ServiceInstanceUpdateRequest) SetMaintenanceInfo(v MaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *ServiceInstanceUpdateRequest) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


