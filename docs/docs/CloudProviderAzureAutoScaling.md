# CloudProviderAzureAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**AzureComputeAutoScalingRules**](AzureComputeAutoScalingRules.md) |  | [optional] 

## Methods

### NewCloudProviderAzureAutoScaling

`func NewCloudProviderAzureAutoScaling() *CloudProviderAzureAutoScaling`

NewCloudProviderAzureAutoScaling instantiates a new CloudProviderAzureAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAzureAutoScalingWithDefaults

`func NewCloudProviderAzureAutoScalingWithDefaults() *CloudProviderAzureAutoScaling`

NewCloudProviderAzureAutoScalingWithDefaults instantiates a new CloudProviderAzureAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *CloudProviderAzureAutoScaling) GetCompute() AzureComputeAutoScalingRules`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *CloudProviderAzureAutoScaling) GetComputeOk() (*AzureComputeAutoScalingRules, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *CloudProviderAzureAutoScaling) SetCompute(v AzureComputeAutoScalingRules)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *CloudProviderAzureAutoScaling) HasCompute() bool`

HasCompute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


