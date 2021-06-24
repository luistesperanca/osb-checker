# Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Id** | **string** |  | 
**Description** | **string** |  | 
**Tags** | Pointer to **[]string** |  | [optional] 
**Requires** | Pointer to **[]string** |  | [optional] 
**Bindable** | **bool** |  | 
**InstancesRetrievable** | Pointer to **bool** |  | [optional] 
**BindingsRetrievable** | Pointer to **bool** |  | [optional] 
**AllowContextUpdates** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | See [Service Metadata Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#service-metadata) for more details. | [optional] 
**DashboardClient** | Pointer to [**DashboardClient**](DashboardClient.md) |  | [optional] 
**BindingRotatable** | Pointer to **bool** |  | [optional] 
**PlanUpdateable** | Pointer to **bool** |  | [optional] 
**Plans** | [**[]Plan**](Plan.md) |  | 

## Methods

### NewService

`func NewService(name string, id string, description string, bindable bool, plans []Plan, ) *Service`

NewService instantiates a new Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWithDefaults

`func NewServiceWithDefaults() *Service`

NewServiceWithDefaults instantiates a new Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Service) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Service) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *Service) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Service) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Service) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *Service) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Service) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Service) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Service) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRequires

`func (o *Service) GetRequires() []string`

GetRequires returns the Requires field if non-nil, zero value otherwise.

### GetRequiresOk

`func (o *Service) GetRequiresOk() (*[]string, bool)`

GetRequiresOk returns a tuple with the Requires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequires

`func (o *Service) SetRequires(v []string)`

SetRequires sets Requires field to given value.

### HasRequires

`func (o *Service) HasRequires() bool`

HasRequires returns a boolean if a field has been set.

### GetBindable

`func (o *Service) GetBindable() bool`

GetBindable returns the Bindable field if non-nil, zero value otherwise.

### GetBindableOk

`func (o *Service) GetBindableOk() (*bool, bool)`

GetBindableOk returns a tuple with the Bindable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindable

`func (o *Service) SetBindable(v bool)`

SetBindable sets Bindable field to given value.


### GetInstancesRetrievable

`func (o *Service) GetInstancesRetrievable() bool`

GetInstancesRetrievable returns the InstancesRetrievable field if non-nil, zero value otherwise.

### GetInstancesRetrievableOk

`func (o *Service) GetInstancesRetrievableOk() (*bool, bool)`

GetInstancesRetrievableOk returns a tuple with the InstancesRetrievable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesRetrievable

`func (o *Service) SetInstancesRetrievable(v bool)`

SetInstancesRetrievable sets InstancesRetrievable field to given value.

### HasInstancesRetrievable

`func (o *Service) HasInstancesRetrievable() bool`

HasInstancesRetrievable returns a boolean if a field has been set.

### GetBindingsRetrievable

`func (o *Service) GetBindingsRetrievable() bool`

GetBindingsRetrievable returns the BindingsRetrievable field if non-nil, zero value otherwise.

### GetBindingsRetrievableOk

`func (o *Service) GetBindingsRetrievableOk() (*bool, bool)`

GetBindingsRetrievableOk returns a tuple with the BindingsRetrievable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingsRetrievable

`func (o *Service) SetBindingsRetrievable(v bool)`

SetBindingsRetrievable sets BindingsRetrievable field to given value.

### HasBindingsRetrievable

`func (o *Service) HasBindingsRetrievable() bool`

HasBindingsRetrievable returns a boolean if a field has been set.

### GetAllowContextUpdates

`func (o *Service) GetAllowContextUpdates() bool`

GetAllowContextUpdates returns the AllowContextUpdates field if non-nil, zero value otherwise.

### GetAllowContextUpdatesOk

`func (o *Service) GetAllowContextUpdatesOk() (*bool, bool)`

GetAllowContextUpdatesOk returns a tuple with the AllowContextUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowContextUpdates

`func (o *Service) SetAllowContextUpdates(v bool)`

SetAllowContextUpdates sets AllowContextUpdates field to given value.

### HasAllowContextUpdates

`func (o *Service) HasAllowContextUpdates() bool`

HasAllowContextUpdates returns a boolean if a field has been set.

### GetMetadata

`func (o *Service) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Service) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Service) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Service) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDashboardClient

`func (o *Service) GetDashboardClient() DashboardClient`

GetDashboardClient returns the DashboardClient field if non-nil, zero value otherwise.

### GetDashboardClientOk

`func (o *Service) GetDashboardClientOk() (*DashboardClient, bool)`

GetDashboardClientOk returns a tuple with the DashboardClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardClient

`func (o *Service) SetDashboardClient(v DashboardClient)`

SetDashboardClient sets DashboardClient field to given value.

### HasDashboardClient

`func (o *Service) HasDashboardClient() bool`

HasDashboardClient returns a boolean if a field has been set.

### GetBindingRotatable

`func (o *Service) GetBindingRotatable() bool`

GetBindingRotatable returns the BindingRotatable field if non-nil, zero value otherwise.

### GetBindingRotatableOk

`func (o *Service) GetBindingRotatableOk() (*bool, bool)`

GetBindingRotatableOk returns a tuple with the BindingRotatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingRotatable

`func (o *Service) SetBindingRotatable(v bool)`

SetBindingRotatable sets BindingRotatable field to given value.

### HasBindingRotatable

`func (o *Service) HasBindingRotatable() bool`

HasBindingRotatable returns a boolean if a field has been set.

### GetPlanUpdateable

`func (o *Service) GetPlanUpdateable() bool`

GetPlanUpdateable returns the PlanUpdateable field if non-nil, zero value otherwise.

### GetPlanUpdateableOk

`func (o *Service) GetPlanUpdateableOk() (*bool, bool)`

GetPlanUpdateableOk returns a tuple with the PlanUpdateable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanUpdateable

`func (o *Service) SetPlanUpdateable(v bool)`

SetPlanUpdateable sets PlanUpdateable field to given value.

### HasPlanUpdateable

`func (o *Service) HasPlanUpdateable() bool`

HasPlanUpdateable returns a boolean if a field has been set.

### GetPlans

`func (o *Service) GetPlans() []Plan`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *Service) GetPlansOk() (*[]Plan, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *Service) SetPlans(v []Plan)`

SetPlans sets Plans field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


