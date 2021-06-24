# ServiceInstanceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**DashboardUrl** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**MaintenanceInfo** | Pointer to [**MaintenanceInfo**](MaintenanceInfo.md) |  | [optional] 
**Metadata** | Pointer to [**ServiceInstanceMetadata**](ServiceInstanceMetadata.md) |  | [optional] 

## Methods

### NewServiceInstanceResource

`func NewServiceInstanceResource() *ServiceInstanceResource`

NewServiceInstanceResource instantiates a new ServiceInstanceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceResourceWithDefaults

`func NewServiceInstanceResourceWithDefaults() *ServiceInstanceResource`

NewServiceInstanceResourceWithDefaults instantiates a new ServiceInstanceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ServiceInstanceResource) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceInstanceResource) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceInstanceResource) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ServiceInstanceResource) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPlanId

`func (o *ServiceInstanceResource) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ServiceInstanceResource) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ServiceInstanceResource) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *ServiceInstanceResource) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetDashboardUrl

`func (o *ServiceInstanceResource) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *ServiceInstanceResource) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *ServiceInstanceResource) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *ServiceInstanceResource) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetParameters

`func (o *ServiceInstanceResource) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceInstanceResource) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceInstanceResource) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceInstanceResource) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *ServiceInstanceResource) GetMaintenanceInfo() MaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *ServiceInstanceResource) GetMaintenanceInfoOk() (*MaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *ServiceInstanceResource) SetMaintenanceInfo(v MaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *ServiceInstanceResource) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceInstanceResource) GetMetadata() ServiceInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceInstanceResource) GetMetadataOk() (*ServiceInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceInstanceResource) SetMetadata(v ServiceInstanceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceInstanceResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


