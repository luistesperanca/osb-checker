# LastOperationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**InstanceUsable** | Pointer to **bool** |  | [optional] 
**UpdateRepeatable** | Pointer to **bool** |  | [optional] 

## Methods

### NewLastOperationResource

`func NewLastOperationResource(state string, ) *LastOperationResource`

NewLastOperationResource instantiates a new LastOperationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastOperationResourceWithDefaults

`func NewLastOperationResourceWithDefaults() *LastOperationResource`

NewLastOperationResourceWithDefaults instantiates a new LastOperationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *LastOperationResource) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LastOperationResource) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LastOperationResource) SetState(v string)`

SetState sets State field to given value.


### GetDescription

`func (o *LastOperationResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LastOperationResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LastOperationResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LastOperationResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstanceUsable

`func (o *LastOperationResource) GetInstanceUsable() bool`

GetInstanceUsable returns the InstanceUsable field if non-nil, zero value otherwise.

### GetInstanceUsableOk

`func (o *LastOperationResource) GetInstanceUsableOk() (*bool, bool)`

GetInstanceUsableOk returns a tuple with the InstanceUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUsable

`func (o *LastOperationResource) SetInstanceUsable(v bool)`

SetInstanceUsable sets InstanceUsable field to given value.

### HasInstanceUsable

`func (o *LastOperationResource) HasInstanceUsable() bool`

HasInstanceUsable returns a boolean if a field has been set.

### GetUpdateRepeatable

`func (o *LastOperationResource) GetUpdateRepeatable() bool`

GetUpdateRepeatable returns the UpdateRepeatable field if non-nil, zero value otherwise.

### GetUpdateRepeatableOk

`func (o *LastOperationResource) GetUpdateRepeatableOk() (*bool, bool)`

GetUpdateRepeatableOk returns a tuple with the UpdateRepeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRepeatable

`func (o *LastOperationResource) SetUpdateRepeatable(v bool)`

SetUpdateRepeatable sets UpdateRepeatable field to given value.

### HasUpdateRepeatable

`func (o *LastOperationResource) HasUpdateRepeatable() bool`

HasUpdateRepeatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


