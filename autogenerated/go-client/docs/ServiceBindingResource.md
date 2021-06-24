# ServiceBindingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ServiceBindingMetadata**](ServiceBindingMetadata.md) |  | [optional] 
**Credentials** | Pointer to **map[string]interface{}** |  | [optional] 
**SyslogDrainUrl** | Pointer to **string** |  | [optional] 
**RouteServiceUrl** | Pointer to **string** |  | [optional] 
**VolumeMounts** | Pointer to [**[]ServiceBindingVolumeMount**](ServiceBindingVolumeMount.md) |  | [optional] 
**Endpoints** | Pointer to [**[]ServiceBindingEndpoint**](ServiceBindingEndpoint.md) |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewServiceBindingResource

`func NewServiceBindingResource() *ServiceBindingResource`

NewServiceBindingResource instantiates a new ServiceBindingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBindingResourceWithDefaults

`func NewServiceBindingResourceWithDefaults() *ServiceBindingResource`

NewServiceBindingResourceWithDefaults instantiates a new ServiceBindingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ServiceBindingResource) GetMetadata() ServiceBindingMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceBindingResource) GetMetadataOk() (*ServiceBindingMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceBindingResource) SetMetadata(v ServiceBindingMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceBindingResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCredentials

`func (o *ServiceBindingResource) GetCredentials() map[string]interface{}`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ServiceBindingResource) GetCredentialsOk() (*map[string]interface{}, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ServiceBindingResource) SetCredentials(v map[string]interface{})`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ServiceBindingResource) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetSyslogDrainUrl

`func (o *ServiceBindingResource) GetSyslogDrainUrl() string`

GetSyslogDrainUrl returns the SyslogDrainUrl field if non-nil, zero value otherwise.

### GetSyslogDrainUrlOk

`func (o *ServiceBindingResource) GetSyslogDrainUrlOk() (*string, bool)`

GetSyslogDrainUrlOk returns a tuple with the SyslogDrainUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogDrainUrl

`func (o *ServiceBindingResource) SetSyslogDrainUrl(v string)`

SetSyslogDrainUrl sets SyslogDrainUrl field to given value.

### HasSyslogDrainUrl

`func (o *ServiceBindingResource) HasSyslogDrainUrl() bool`

HasSyslogDrainUrl returns a boolean if a field has been set.

### GetRouteServiceUrl

`func (o *ServiceBindingResource) GetRouteServiceUrl() string`

GetRouteServiceUrl returns the RouteServiceUrl field if non-nil, zero value otherwise.

### GetRouteServiceUrlOk

`func (o *ServiceBindingResource) GetRouteServiceUrlOk() (*string, bool)`

GetRouteServiceUrlOk returns a tuple with the RouteServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteServiceUrl

`func (o *ServiceBindingResource) SetRouteServiceUrl(v string)`

SetRouteServiceUrl sets RouteServiceUrl field to given value.

### HasRouteServiceUrl

`func (o *ServiceBindingResource) HasRouteServiceUrl() bool`

HasRouteServiceUrl returns a boolean if a field has been set.

### GetVolumeMounts

`func (o *ServiceBindingResource) GetVolumeMounts() []ServiceBindingVolumeMount`

GetVolumeMounts returns the VolumeMounts field if non-nil, zero value otherwise.

### GetVolumeMountsOk

`func (o *ServiceBindingResource) GetVolumeMountsOk() (*[]ServiceBindingVolumeMount, bool)`

GetVolumeMountsOk returns a tuple with the VolumeMounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeMounts

`func (o *ServiceBindingResource) SetVolumeMounts(v []ServiceBindingVolumeMount)`

SetVolumeMounts sets VolumeMounts field to given value.

### HasVolumeMounts

`func (o *ServiceBindingResource) HasVolumeMounts() bool`

HasVolumeMounts returns a boolean if a field has been set.

### GetEndpoints

`func (o *ServiceBindingResource) GetEndpoints() []ServiceBindingEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServiceBindingResource) GetEndpointsOk() (*[]ServiceBindingEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServiceBindingResource) SetEndpoints(v []ServiceBindingEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ServiceBindingResource) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetParameters

`func (o *ServiceBindingResource) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceBindingResource) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceBindingResource) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceBindingResource) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


