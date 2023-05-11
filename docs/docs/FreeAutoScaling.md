# FreeAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to **string** | Collection of settings that configures how a cluster might scale its cluster tier and whether the cluster can scale down. | [optional] 

## Methods

### NewFreeAutoScaling

`func NewFreeAutoScaling() *FreeAutoScaling`

NewFreeAutoScaling instantiates a new FreeAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeAutoScalingWithDefaults

`func NewFreeAutoScalingWithDefaults() *FreeAutoScaling`

NewFreeAutoScalingWithDefaults instantiates a new FreeAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *FreeAutoScaling) GetCompute() string`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *FreeAutoScaling) GetComputeOk() (*string, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *FreeAutoScaling) SetCompute(v string)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *FreeAutoScaling) HasCompute() bool`

HasCompute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


