# AutoScalingV15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**ComputeAutoScalingV15**](ComputeAutoScalingV15.md) |  | [optional] 
**DiskGB** | Pointer to [**DiskGBAutoScaling**](DiskGBAutoScaling.md) |  | [optional] 

## Methods

### NewAutoScalingV15

`func NewAutoScalingV15() *AutoScalingV15`

NewAutoScalingV15 instantiates a new AutoScalingV15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingV15WithDefaults

`func NewAutoScalingV15WithDefaults() *AutoScalingV15`

NewAutoScalingV15WithDefaults instantiates a new AutoScalingV15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *AutoScalingV15) GetCompute() ComputeAutoScalingV15`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *AutoScalingV15) GetComputeOk() (*ComputeAutoScalingV15, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *AutoScalingV15) SetCompute(v ComputeAutoScalingV15)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *AutoScalingV15) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### GetDiskGB

`func (o *AutoScalingV15) GetDiskGB() DiskGBAutoScaling`

GetDiskGB returns the DiskGB field if non-nil, zero value otherwise.

### GetDiskGBOk

`func (o *AutoScalingV15) GetDiskGBOk() (*DiskGBAutoScaling, bool)`

GetDiskGBOk returns a tuple with the DiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskGB

`func (o *AutoScalingV15) SetDiskGB(v DiskGBAutoScaling)`

SetDiskGB sets DiskGB field to given value.

### HasDiskGB

`func (o *AutoScalingV15) HasDiskGB() bool`

HasDiskGB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


