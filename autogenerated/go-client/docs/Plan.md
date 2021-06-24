# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** | See [Service Metadata Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#service-metadata) for more details. | [optional] 
**Free** | Pointer to **bool** |  | [optional] [default to true]
**Bindable** | Pointer to **bool** |  | [optional] 
**PlanUpdateable** | Pointer to **bool** |  | [optional] 
**Schemas** | Pointer to [**SchemasObject**](SchemasObject.md) |  | [optional] 
**MaximumPollingDuration** | Pointer to **int32** |  | [optional] 
**MaintenanceInfo** | Pointer to [**MaintenanceInfo**](MaintenanceInfo.md) |  | [optional] 
**BindingRotatable** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPlan

`func NewPlan(id string, name string, description string, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Plan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plan) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetadata

`func (o *Plan) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Plan) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Plan) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Plan) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFree

`func (o *Plan) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *Plan) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *Plan) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *Plan) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetBindable

`func (o *Plan) GetBindable() bool`

GetBindable returns the Bindable field if non-nil, zero value otherwise.

### GetBindableOk

`func (o *Plan) GetBindableOk() (*bool, bool)`

GetBindableOk returns a tuple with the Bindable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindable

`func (o *Plan) SetBindable(v bool)`

SetBindable sets Bindable field to given value.

### HasBindable

`func (o *Plan) HasBindable() bool`

HasBindable returns a boolean if a field has been set.

### GetPlanUpdateable

`func (o *Plan) GetPlanUpdateable() bool`

GetPlanUpdateable returns the PlanUpdateable field if non-nil, zero value otherwise.

### GetPlanUpdateableOk

`func (o *Plan) GetPlanUpdateableOk() (*bool, bool)`

GetPlanUpdateableOk returns a tuple with the PlanUpdateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanUpdateable

`func (o *Plan) SetPlanUpdateable(v bool)`

SetPlanUpdateable sets PlanUpdateable field to given value.

### HasPlanUpdateable

`func (o *Plan) HasPlanUpdateable() bool`

HasPlanUpdateable returns a boolean if a field has been set.

### GetSchemas

`func (o *Plan) GetSchemas() SchemasObject`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Plan) GetSchemasOk() (*SchemasObject, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Plan) SetSchemas(v SchemasObject)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *Plan) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetMaximumPollingDuration

`func (o *Plan) GetMaximumPollingDuration() int32`

GetMaximumPollingDuration returns the MaximumPollingDuration field if non-nil, zero value otherwise.

### GetMaximumPollingDurationOk

`func (o *Plan) GetMaximumPollingDurationOk() (*int32, bool)`

GetMaximumPollingDurationOk returns a tuple with the MaximumPollingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPollingDuration

`func (o *Plan) SetMaximumPollingDuration(v int32)`

SetMaximumPollingDuration sets MaximumPollingDuration field to given value.

### HasMaximumPollingDuration

`func (o *Plan) HasMaximumPollingDuration() bool`

HasMaximumPollingDuration returns a boolean if a field has been set.

### GetMaintenanceInfo

`func (o *Plan) GetMaintenanceInfo() MaintenanceInfo`

GetMaintenanceInfo returns the MaintenanceInfo field if non-nil, zero value otherwise.

### GetMaintenanceInfoOk

`func (o *Plan) GetMaintenanceInfoOk() (*MaintenanceInfo, bool)`

GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceInfo

`func (o *Plan) SetMaintenanceInfo(v MaintenanceInfo)`

SetMaintenanceInfo sets MaintenanceInfo field to given value.

### HasMaintenanceInfo

`func (o *Plan) HasMaintenanceInfo() bool`

HasMaintenanceInfo returns a boolean if a field has been set.

### GetBindingRotatable

`func (o *Plan) GetBindingRotatable() bool`

GetBindingRotatable returns the BindingRotatable field if non-nil, zero value otherwise.

### GetBindingRotatableOk

`func (o *Plan) GetBindingRotatableOk() (*bool, bool)`

GetBindingRotatableOk returns a tuple with the BindingRotatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingRotatable

`func (o *Plan) SetBindingRotatable(v bool)`

SetBindingRotatable sets BindingRotatable field to given value.

### HasBindingRotatable

`func (o *Plan) HasBindingRotatable() bool`

HasBindingRotatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


