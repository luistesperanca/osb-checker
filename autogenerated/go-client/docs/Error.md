# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstanceUsable** | Pointer to **bool** |  | [optional] 
**UpdateRepeatable** | Pointer to **bool** |  | [optional] 

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *Error) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Error) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Error) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *Error) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDescription

`func (o *Error) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Error) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Error) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Error) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstanceUsable

`func (o *Error) GetInstanceUsable() bool`

GetInstanceUsable returns the InstanceUsable field if non-nil, zero value otherwise.

### GetInstanceUsableOk

`func (o *Error) GetInstanceUsableOk() (*bool, bool)`

GetInstanceUsableOk returns a tuple with the InstanceUsable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUsable

`func (o *Error) SetInstanceUsable(v bool)`

SetInstanceUsable sets InstanceUsable field to given value.

### HasInstanceUsable

`func (o *Error) HasInstanceUsable() bool`

HasInstanceUsable returns a boolean if a field has been set.

### GetUpdateRepeatable

`func (o *Error) GetUpdateRepeatable() bool`

GetUpdateRepeatable returns the UpdateRepeatable field if non-nil, zero value otherwise.

### GetUpdateRepeatableOk

`func (o *Error) GetUpdateRepeatableOk() (*bool, bool)`

GetUpdateRepeatableOk returns a tuple with the UpdateRepeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRepeatable

`func (o *Error) SetUpdateRepeatable(v bool)`

SetUpdateRepeatable sets UpdateRepeatable field to given value.

### HasUpdateRepeatable

`func (o *Error) HasUpdateRepeatable() bool`

HasUpdateRepeatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


