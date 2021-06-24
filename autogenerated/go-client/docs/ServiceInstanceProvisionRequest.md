# ServiceInstanceProvisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** |  | 
**PlanId** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** | See [Context Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#context-object) for more details. | [optional] 
**OrganizationGuid** | **string** |  | 
**SpaceGuid** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**MaintenanceInfo** | Pointer to [**MaintenanceInfo**](MaintenanceInfo.md) |  | [optional] 

## Methods

### NewServiceInstanceProvisionRequest

`func NewServiceInstanceProvisionRequest(serviceId string, planId string, organizationGuid string, spaceGuid string, ) *ServiceInstanceProvisionRequest`

NewServiceInstanceProvisionRequest instantiates a new ServiceInstanceProvisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceProvisionRequestWithDefaults

`func NewServiceInstanceProvisionRequestWithDefaults() *ServiceInstanceProvisionRequest`

NewServiceInstanceProvisionRequestWithDefaults instantiates a new ServiceInstanceProvisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ServiceInstanceProvisionRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceInstanceProvisionRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceInstanceProvisionRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPlanId

`func (o *ServiceInstanceProvisionRequest) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ServiceInstanceProvisionRequest) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ServiceInstanceProvisionRequest) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetContext

`func (o *ServiceInstanceProvisionRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ServiceInstanceProvisionRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ServiceInstanceProvisionRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ServiceInstanceProvisionRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *ServiceInstanceProvisionRequest) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *ServiceInstanceProvisionRequest) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *ServiceInstanceProvisionRequest) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.


### GetSpaceGuid

`func (o *ServiceInstanceProvisionRequest) GetSpaceGuid() string`

GetSpaceGuid returns the SpaceGuid field if non-nil, zero value otherwise.

### GetSpaceGuidOk

`func (o *ServiceInstanceProvisionRequest) GetSpaceGuidOk() (*string, bool)`

GetSpaceGuidOk returns a tuple with the SpaceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceGuid

`func (o *ServiceInstanceProvisionRequest) SetSpaceGuid(v string)`

SetSpaceGuid sets SpaceGuid field to given value.


### GetParameters

`func (o *ServiceInstanceProvisionRequest) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceInstanceProvisionRequest) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceInstanceProvisionRequest) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceInstanceProvisionRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *ServiceInstanceProvisionRequest) GetMaintenanceInfo() MaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *ServiceInstanceProvisionRequest) GetMaintenanceInfoOk() (*MaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *ServiceInstanceProvisionRequest) SetMaintenanceInfo(v MaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *ServiceInstanceProvisionRequest) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


