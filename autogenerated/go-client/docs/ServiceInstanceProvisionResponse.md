# ServiceInstanceProvisionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardUrl** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ServiceInstanceMetadata**](ServiceInstanceMetadata.md) |  | [optional] 

## Methods

### NewServiceInstanceProvisionResponse

`func NewServiceInstanceProvisionResponse() *ServiceInstanceProvisionResponse`

NewServiceInstanceProvisionResponse instantiates a new ServiceInstanceProvisionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceProvisionResponseWithDefaults

`func NewServiceInstanceProvisionResponseWithDefaults() *ServiceInstanceProvisionResponse`

NewServiceInstanceProvisionResponseWithDefaults instantiates a new ServiceInstanceProvisionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardUrl

`func (o *ServiceInstanceProvisionResponse) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *ServiceInstanceProvisionResponse) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *ServiceInstanceProvisionResponse) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *ServiceInstanceProvisionResponse) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceInstanceProvisionResponse) GetMetadata() ServiceInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceInstanceProvisionResponse) GetMetadataOk() (*ServiceInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceInstanceProvisionResponse) SetMetadata(v ServiceInstanceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceInstanceProvisionResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


