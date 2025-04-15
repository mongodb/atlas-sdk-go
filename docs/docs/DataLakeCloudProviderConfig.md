# DataLakeCloudProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aws** | Pointer to [**DataLakeAWSCloudProviderConfig**](DataLakeAWSCloudProviderConfig.md) |  | [optional] 
**Azure** | Pointer to [**DataFederationAzureCloudProviderConfig**](DataFederationAzureCloudProviderConfig.md) |  | [optional] 
**Gcp** | Pointer to [**DataFederationGCPCloudProviderConfig**](DataFederationGCPCloudProviderConfig.md) |  | [optional] 

## Methods

### NewDataLakeCloudProviderConfig

`func NewDataLakeCloudProviderConfig() *DataLakeCloudProviderConfig`

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

### HasAws

`func (o *DataLakeCloudProviderConfig) HasAws() bool`

HasAws returns a boolean if a field has been set.
### GetAzure

`func (o *DataLakeCloudProviderConfig) GetAzure() DataFederationAzureCloudProviderConfig`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *DataLakeCloudProviderConfig) GetAzureOk() (*DataFederationAzureCloudProviderConfig, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *DataLakeCloudProviderConfig) SetAzure(v DataFederationAzureCloudProviderConfig)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *DataLakeCloudProviderConfig) HasAzure() bool`

HasAzure returns a boolean if a field has been set.
### GetGcp

`func (o *DataLakeCloudProviderConfig) GetGcp() DataFederationGCPCloudProviderConfig`

GetGcp returns the Gcp field if non-nil, zero value otherwise.

### GetGcpOk

`func (o *DataLakeCloudProviderConfig) GetGcpOk() (*DataFederationGCPCloudProviderConfig, bool)`

GetGcpOk returns a tuple with the Gcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcp

`func (o *DataLakeCloudProviderConfig) SetGcp(v DataFederationGCPCloudProviderConfig)`

SetGcp sets Gcp field to given value.

### HasGcp

`func (o *DataLakeCloudProviderConfig) HasGcp() bool`

HasGcp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


