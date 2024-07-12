# DataLakeTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderConfig** | Pointer to [**DataLakeCloudProviderConfig**](DataLakeCloudProviderConfig.md) |  | [optional] 
**DataProcessRegion** | Pointer to [**DataLakeDataProcessRegion**](DataLakeDataProcessRegion.md) |  | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Hostnames** | Pointer to **[]string** | List that contains the hostnames assigned to the Federated Database Instance. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the Federated Database Instance. | [optional] 
**PrivateEndpointHostnames** | Pointer to [**[]PrivateEndpointHostname**](PrivateEndpointHostname.md) | List that contains the sets of private endpoints and hostnames. | [optional] [readonly] 
**State** | Pointer to **string** | Label that indicates the status of the Federated Database Instance. | [optional] [readonly] 
**Storage** | Pointer to [**DataLakeStorage**](DataLakeStorage.md) |  | [optional] 

## Methods

### NewDataLakeTenant

`func NewDataLakeTenant() *DataLakeTenant`

NewDataLakeTenant instantiates a new DataLakeTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeTenantWithDefaults

`func NewDataLakeTenantWithDefaults() *DataLakeTenant`

NewDataLakeTenantWithDefaults instantiates a new DataLakeTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderConfig

`func (o *DataLakeTenant) GetCloudProviderConfig() DataLakeCloudProviderConfig`

GetCloudProviderConfig returns the CloudProviderConfig field if non-nil, zero value otherwise.

### GetCloudProviderConfigOk

`func (o *DataLakeTenant) GetCloudProviderConfigOk() (*DataLakeCloudProviderConfig, bool)`

GetCloudProviderConfigOk returns a tuple with the CloudProviderConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderConfig

`func (o *DataLakeTenant) SetCloudProviderConfig(v DataLakeCloudProviderConfig)`

SetCloudProviderConfig sets CloudProviderConfig field to given value.

### HasCloudProviderConfig

`func (o *DataLakeTenant) HasCloudProviderConfig() bool`

HasCloudProviderConfig returns a boolean if a field has been set.
### GetDataProcessRegion

`func (o *DataLakeTenant) GetDataProcessRegion() DataLakeDataProcessRegion`

GetDataProcessRegion returns the DataProcessRegion field if non-nil, zero value otherwise.

### GetDataProcessRegionOk

`func (o *DataLakeTenant) GetDataProcessRegionOk() (*DataLakeDataProcessRegion, bool)`

GetDataProcessRegionOk returns a tuple with the DataProcessRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessRegion

`func (o *DataLakeTenant) SetDataProcessRegion(v DataLakeDataProcessRegion)`

SetDataProcessRegion sets DataProcessRegion field to given value.

### HasDataProcessRegion

`func (o *DataLakeTenant) HasDataProcessRegion() bool`

HasDataProcessRegion returns a boolean if a field has been set.
### GetGroupId

`func (o *DataLakeTenant) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DataLakeTenant) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DataLakeTenant) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DataLakeTenant) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetHostnames

`func (o *DataLakeTenant) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *DataLakeTenant) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *DataLakeTenant) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *DataLakeTenant) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.
### GetName

`func (o *DataLakeTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeTenant) HasName() bool`

HasName returns a boolean if a field has been set.
### GetPrivateEndpointHostnames

`func (o *DataLakeTenant) GetPrivateEndpointHostnames() []PrivateEndpointHostname`

GetPrivateEndpointHostnames returns the PrivateEndpointHostnames field if non-nil, zero value otherwise.

### GetPrivateEndpointHostnamesOk

`func (o *DataLakeTenant) GetPrivateEndpointHostnamesOk() (*[]PrivateEndpointHostname, bool)`

GetPrivateEndpointHostnamesOk returns a tuple with the PrivateEndpointHostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointHostnames

`func (o *DataLakeTenant) SetPrivateEndpointHostnames(v []PrivateEndpointHostname)`

SetPrivateEndpointHostnames sets PrivateEndpointHostnames field to given value.

### HasPrivateEndpointHostnames

`func (o *DataLakeTenant) HasPrivateEndpointHostnames() bool`

HasPrivateEndpointHostnames returns a boolean if a field has been set.
### GetState

`func (o *DataLakeTenant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DataLakeTenant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DataLakeTenant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *DataLakeTenant) HasState() bool`

HasState returns a boolean if a field has been set.
### GetStorage

`func (o *DataLakeTenant) GetStorage() DataLakeStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DataLakeTenant) GetStorageOk() (*DataLakeStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DataLakeTenant) SetStorage(v DataLakeStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DataLakeTenant) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


