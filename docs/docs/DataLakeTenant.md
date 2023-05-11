# DataLakeTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderConfig** | Pointer to [**DataLakeCloudProviderConfig**](DataLakeCloudProviderConfig.md) |  | [optional] 
**DataProcessRegion** | Pointer to [**DataLakeDataProcessRegion**](DataLakeDataProcessRegion.md) |  | [optional] 
**Name** | Pointer to **string** | Human-readable label that identifies the data lake. | [optional] 
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


