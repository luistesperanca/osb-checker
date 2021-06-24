# ServiceBindingVolumeMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Driver** | **string** |  | 
**ContainerDir** | **string** |  | 
**Mode** | **string** |  | 
**DeviceType** | **string** |  | 
**Device** | [**ServiceBindingVolumeMountDevice**](ServiceBindingVolumeMountDevice.md) |  | 

## Methods

### NewServiceBindingVolumeMount

`func NewServiceBindingVolumeMount(driver string, containerDir string, mode string, deviceType string, device ServiceBindingVolumeMountDevice, ) *ServiceBindingVolumeMount`

NewServiceBindingVolumeMount instantiates a new ServiceBindingVolumeMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceBindingVolumeMountWithDefaults

`func NewServiceBindingVolumeMountWithDefaults() *ServiceBindingVolumeMount`

NewServiceBindingVolumeMountWithDefaults instantiates a new ServiceBindingVolumeMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriver

`func (o *ServiceBindingVolumeMount) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *ServiceBindingVolumeMount) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *ServiceBindingVolumeMount) SetDriver(v string)`

SetDriver sets Driver field to given value.


### GetContainerDir

`func (o *ServiceBindingVolumeMount) GetContainerDir() string`

GetContainerDir returns the ContainerDir field if non-nil, zero value otherwise.

### GetContainerDirOk

`func (o *ServiceBindingVolumeMount) GetContainerDirOk() (*string, bool)`

GetContainerDirOk returns a tuple with the ContainerDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDir

`func (o *ServiceBindingVolumeMount) SetContainerDir(v string)`

SetContainerDir sets ContainerDir field to given value.


### GetMode

`func (o *ServiceBindingVolumeMount) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ServiceBindingVolumeMount) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ServiceBindingVolumeMount) SetMode(v string)`

SetMode sets Mode field to given value.


### GetDeviceType

`func (o *ServiceBindingVolumeMount) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ServiceBindingVolumeMount) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ServiceBindingVolumeMount) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.


### GetDevice

`func (o *ServiceBindingVolumeMount) GetDevice() ServiceBindingVolumeMountDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ServiceBindingVolumeMount) GetDeviceOk() (*ServiceBindingVolumeMountDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ServiceBindingVolumeMount) SetDevice(v ServiceBindingVolumeMountDevice)`

SetDevice sets Device field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


