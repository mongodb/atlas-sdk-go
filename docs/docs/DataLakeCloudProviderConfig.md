# DataLakeCloudProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | [**DataLakeAWSCloudProviderConfig**](DataLakeAWSCloudProviderConfig.md) |  | 

## Methods

### NewDataLakeCloudProviderConfig

`func NewDataLakeCloudProviderConfig(aws DataLakeAWSCloudProviderConfig, ) *DataLakeCloudProviderConfig`

NewDataLakeCloudProviderConfig instantiates a new DataLakeCloudProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeCloudProviderConfigWithDefaults

`func NewDataLakeCloudProviderConfigWithDefaults() *DataLakeCloudProviderConfig`

NewDataLakeCloudProviderConfigWithDefaults instantiates a new DataLakeCloudProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAws

`func (o *DataLakeCloudProviderConfig) GetAws() DataLakeAWSCloudProviderConfig`

GetAws returns the Aws field if non-nil, zero value otherwise.

### GetAwsOk

`func (o *DataLakeCloudProviderConfig) GetAwsOk() (*DataLakeAWSCloudProviderConfig, bool)`

GetAwsOk returns a tuple with the Aws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAws

`func (o *DataLakeCloudProviderConfig) SetAws(v DataLakeAWSCloudProviderConfig)`

SetAws sets Aws field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


