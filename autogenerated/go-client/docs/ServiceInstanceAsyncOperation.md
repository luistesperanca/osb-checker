# ServiceInstanceAsyncOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardUrl** | Pointer to **string** |  | [optional] 
**Operation** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ServiceInstanceMetadata**](ServiceInstanceMetadata.md) |  | [optional] 

## Methods

### NewServiceInstanceAsyncOperation

`func NewServiceInstanceAsyncOperation() *ServiceInstanceAsyncOperation`

NewServiceInstanceAsyncOperation instantiates a new ServiceInstanceAsyncOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceAsyncOperationWithDefaults

`func NewServiceInstanceAsyncOperationWithDefaults() *ServiceInstanceAsyncOperation`

NewServiceInstanceAsyncOperationWithDefaults instantiates a new ServiceInstanceAsyncOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardUrl

`func (o *ServiceInstanceAsyncOperation) GetDashboardUrl() string`

GetDashboardUrl returns the DashboardUrl field if non-nil, zero value otherwise.

### GetDashboardUrlOk

`func (o *ServiceInstanceAsyncOperation) GetDashboardUrlOk() (*string, bool)`

GetDashboardUrlOk returns a tuple with the DashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardUrl

`func (o *ServiceInstanceAsyncOperation) SetDashboardUrl(v string)`

SetDashboardUrl sets DashboardUrl field to given value.

### HasDashboardUrl

`func (o *ServiceInstanceAsyncOperation) HasDashboardUrl() bool`

HasDashboardUrl returns a boolean if a field has been set.

### GetOperation

`func (o *ServiceInstanceAsyncOperation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ServiceInstanceAsyncOperation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ServiceInstanceAsyncOperation) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ServiceInstanceAsyncOperation) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceInstanceAsyncOperation) GetMetadata() ServiceInstanceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceInstanceAsyncOperation) GetMetadataOk() (*ServiceInstanceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceInstanceAsyncOperation) SetMetadata(v ServiceInstanceMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceInstanceAsyncOperation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


