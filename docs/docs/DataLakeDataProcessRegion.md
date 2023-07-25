# DataLakeDataProcessRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Name of the cloud service that hosts the data lake&#39;s data stores. | 
**Region** | **string** | Name of the region to which the data lake routes client connections. | 

## Methods

### NewDataLakeDataProcessRegion

`func NewDataLakeDataProcessRegion(cloudProvider string, region string, ) *DataLakeDataProcessRegion`

NewDataLakeDataProcessRegion instantiates a new DataLakeDataProcessRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeDataProcessRegionWithDefaults

`func NewDataLakeDataProcessRegionWithDefaults() *DataLakeDataProcessRegion`

NewDataLakeDataProcessRegionWithDefaults instantiates a new DataLakeDataProcessRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DataLakeDataProcessRegion) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DataLakeDataProcessRegion) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DataLakeDataProcessRegion) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetRegion

`func (o *DataLakeDataProcessRegion) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DataLakeDataProcessRegion) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DataLakeDataProcessRegion) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


