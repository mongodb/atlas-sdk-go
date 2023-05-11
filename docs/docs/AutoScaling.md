# AutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**ComputeAutoScaling**](ComputeAutoScaling.md) |  | [optional] 
**DiskGBEnabled** | Pointer to **bool** | Flag that indicates whether someone enabled disk auto-scaling for this cluster. | [optional] 

## Methods

### NewAutoScaling

`func NewAutoScaling() *AutoScaling`

NewAutoScaling instantiates a new AutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingWithDefaults

`func NewAutoScalingWithDefaults() *AutoScaling`

NewAutoScalingWithDefaults instantiates a new AutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *AutoScaling) GetCompute() ComputeAutoScaling`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *AutoScaling) GetComputeOk() (*ComputeAutoScaling, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *AutoScaling) SetCompute(v ComputeAutoScaling)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *AutoScaling) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetDiskGBEnabled

`func (o *AutoScaling) GetDiskGBEnabled() bool`

GetDiskGBEnabled returns the DiskGBEnabled field if non-nil, zero value otherwise.

### GetDiskGBEnabledOk

`func (o *AutoScaling) GetDiskGBEnabledOk() (*bool, bool)`

GetDiskGBEnabledOk returns a tuple with the DiskGBEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGBEnabled

`func (o *AutoScaling) SetDiskGBEnabled(v bool)`

SetDiskGBEnabled sets DiskGBEnabled field to given value.

### HasDiskGBEnabled

`func (o *AutoScaling) HasDiskGBEnabled() bool`

HasDiskGBEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


