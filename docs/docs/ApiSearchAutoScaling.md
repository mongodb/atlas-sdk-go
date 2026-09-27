# ApiSearchAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compute** | Pointer to [**ApiSearchComputeAutoScaling**](ApiSearchComputeAutoScaling.md) |  | [optional] 

## Methods

### NewApiSearchAutoScaling

`func NewApiSearchAutoScaling() *ApiSearchAutoScaling`

NewApiSearchAutoScaling instantiates a new ApiSearchAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchAutoScalingWithDefaults

`func NewApiSearchAutoScalingWithDefaults() *ApiSearchAutoScaling`

NewApiSearchAutoScalingWithDefaults instantiates a new ApiSearchAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompute

`func (o *ApiSearchAutoScaling) GetCompute() ApiSearchComputeAutoScaling`

GetCompute returns the Compute field if non-nil, zero value otherwise.

### GetComputeOk

`func (o *ApiSearchAutoScaling) GetComputeOk() (*ApiSearchComputeAutoScaling, bool)`

GetComputeOk returns a tuple with the Compute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompute

`func (o *ApiSearchAutoScaling) SetCompute(v ApiSearchComputeAutoScaling)`

SetCompute sets Compute field to given value.

### HasCompute

`func (o *ApiSearchAutoScaling) HasCompute() bool`

HasCompute returns a boolean if a field has been set.

### SetComputeNil

`func (o *ApiSearchAutoScaling) SetComputeNil()`

SetComputeNil sets Compute to an explicit JSON null when marshaled, overriding any value previously set with SetCompute. Calling SetCompute again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


