# CloudProviderAWSAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**AWSComputeAutoScaling**](AWSComputeAutoScaling.md) |  | [optional] 

## Methods

### NewCloudProviderAWSAutoScaling

`func NewCloudProviderAWSAutoScaling() *CloudProviderAWSAutoScaling`

NewCloudProviderAWSAutoScaling instantiates a new CloudProviderAWSAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAWSAutoScalingWithDefaults

`func NewCloudProviderAWSAutoScalingWithDefaults() *CloudProviderAWSAutoScaling`

NewCloudProviderAWSAutoScalingWithDefaults instantiates a new CloudProviderAWSAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *CloudProviderAWSAutoScaling) GetCompute() AWSComputeAutoScaling`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *CloudProviderAWSAutoScaling) GetComputeOk() (*AWSComputeAutoScaling, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *CloudProviderAWSAutoScaling) SetCompute(v AWSComputeAutoScaling)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *CloudProviderAWSAutoScaling) HasCompute() bool`

HasCompute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


