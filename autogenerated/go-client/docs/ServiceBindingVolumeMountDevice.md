# ServiceBindingVolumeMountDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeId** | **string** |  | 
**MountConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewServiceBindingVolumeMountDevice

`func NewServiceBindingVolumeMountDevice(volumeId string, ) *ServiceBindingVolumeMountDevice`

NewServiceBindingVolumeMountDevice instantiates a new ServiceBindingVolumeMountDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBindingVolumeMountDeviceWithDefaults

`func NewServiceBindingVolumeMountDeviceWithDefaults() *ServiceBindingVolumeMountDevice`

NewServiceBindingVolumeMountDeviceWithDefaults instantiates a new ServiceBindingVolumeMountDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeId

`func (o *ServiceBindingVolumeMountDevice) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ServiceBindingVolumeMountDevice) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ServiceBindingVolumeMountDevice) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### GetMountConfig

`func (o *ServiceBindingVolumeMountDevice) GetMountConfig() map[string]interface{}`

GetMountConfig returns the MountConfig field if non-nil, zero value otherwise.

### GetMountConfigOk

`func (o *ServiceBindingVolumeMountDevice) GetMountConfigOk() (*map[string]interface{}, bool)`

GetMountConfigOk returns a tuple with the MountConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountConfig

`func (o *ServiceBindingVolumeMountDevice) SetMountConfig(v map[string]interface{})`

SetMountConfig sets MountConfig field to given value.

### HasMountConfig

`func (o *ServiceBindingVolumeMountDevice) HasMountConfig() bool`

HasMountConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


