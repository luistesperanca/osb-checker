# ServiceBindingMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** |  | [optional] 
**RenewBefore** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceBindingMetadata

`func NewServiceBindingMetadata() *ServiceBindingMetadata`

NewServiceBindingMetadata instantiates a new ServiceBindingMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBindingMetadataWithDefaults

`func NewServiceBindingMetadataWithDefaults() *ServiceBindingMetadata`

NewServiceBindingMetadataWithDefaults instantiates a new ServiceBindingMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *ServiceBindingMetadata) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ServiceBindingMetadata) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ServiceBindingMetadata) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ServiceBindingMetadata) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRenewBefore

`func (o *ServiceBindingMetadata) GetRenewBefore() string`

GetRenewBefore returns the RenewBefore field if non-nil, zero value otherwise.

### GetRenewBeforeOk

`func (o *ServiceBindingMetadata) GetRenewBeforeOk() (*string, bool)`

GetRenewBeforeOk returns a tuple with the RenewBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewBefore

`func (o *ServiceBindingMetadata) SetRenewBefore(v string)`

SetRenewBefore sets RenewBefore field to given value.

### HasRenewBefore

`func (o *ServiceBindingMetadata) HasRenewBefore() bool`

HasRenewBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


