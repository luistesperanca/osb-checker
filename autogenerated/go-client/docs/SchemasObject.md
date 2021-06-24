# SchemasObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstance** | Pointer to [**ServiceInstanceSchemaObject**](ServiceInstanceSchemaObject.md) |  | [optional] 
**ServiceBinding** | Pointer to [**ServiceBindingSchemaObject**](ServiceBindingSchemaObject.md) |  | [optional] 

## Methods

### NewSchemasObject

`func NewSchemasObject() *SchemasObject`

NewSchemasObject instantiates a new SchemasObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasObjectWithDefaults

`func NewSchemasObjectWithDefaults() *SchemasObject`

NewSchemasObjectWithDefaults instantiates a new SchemasObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstance

`func (o *SchemasObject) GetServiceInstance() ServiceInstanceSchemaObject`

GetServiceInstance returns the ServiceInstance field if non-nil, zero value otherwise.

### GetServiceInstanceOk

`func (o *SchemasObject) GetServiceInstanceOk() (*ServiceInstanceSchemaObject, bool)`

GetServiceInstanceOk returns a tuple with the ServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstance

`func (o *SchemasObject) SetServiceInstance(v ServiceInstanceSchemaObject)`

SetServiceInstance sets ServiceInstance field to given value.

### HasServiceInstance

`func (o *SchemasObject) HasServiceInstance() bool`

HasServiceInstance returns a boolean if a field has been set.

### GetServiceBinding

`func (o *SchemasObject) GetServiceBinding() ServiceBindingSchemaObject`

GetServiceBinding returns the ServiceBinding field if non-nil, zero value otherwise.

### GetServiceBindingOk

`func (o *SchemasObject) GetServiceBindingOk() (*ServiceBindingSchemaObject, bool)`

GetServiceBindingOk returns a tuple with the ServiceBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBinding

`func (o *SchemasObject) SetServiceBinding(v ServiceBindingSchemaObject)`

SetServiceBinding sets ServiceBinding field to given value.

### HasServiceBinding

`func (o *SchemasObject) HasServiceBinding() bool`

HasServiceBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


