# AzureAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**AzureComputeAutoScaling**](AzureComputeAutoScaling.md) |  | [optional] 

## Methods

### NewAzureAutoScaling

`func NewAzureAutoScaling() *AzureAutoScaling`

NewAzureAutoScaling instantiates a new AzureAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureAutoScalingWithDefaults

`func NewAzureAutoScalingWithDefaults() *AzureAutoScaling`

NewAzureAutoScalingWithDefaults instantiates a new AzureAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *AzureAutoScaling) GetCompute() AzureComputeAutoScaling`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *AzureAutoScaling) GetComputeOk() (*AzureComputeAutoScaling, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *AzureAutoScaling) SetCompute(v AzureComputeAutoScaling)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *AzureAutoScaling) HasCompute() bool`

HasCompute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


