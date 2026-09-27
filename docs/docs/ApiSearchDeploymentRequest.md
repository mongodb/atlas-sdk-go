# ApiSearchDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoScaling** | Pointer to [**ApiSearchAutoScaling**](ApiSearchAutoScaling.md) | Settings that let Atlas change the Search Node tier on its own. Bounds apply to the whole deployment, and Atlas scales each region and shard independently within them. Omit to keep autoscaling off. | [optional] 
**DefaultNodeCount** | Pointer to **int** | Default number of Search Nodes per region. Applied to a region without an explicit override. | [optional] 
**Specs** | [**[]ApiSearchDeploymentRequestSpec**](ApiSearchDeploymentRequestSpec.md) | List of settings that configure the Search Nodes for your cluster. Provide one element per region when configuring asymmetric deployments; a single element applies to all regions. | 

## Methods

### NewApiSearchDeploymentRequest

`func NewApiSearchDeploymentRequest(specs []ApiSearchDeploymentRequestSpec, ) *ApiSearchDeploymentRequest`

NewApiSearchDeploymentRequest instantiates a new ApiSearchDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchDeploymentRequestWithDefaults

`func NewApiSearchDeploymentRequestWithDefaults() *ApiSearchDeploymentRequest`

NewApiSearchDeploymentRequestWithDefaults instantiates a new ApiSearchDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoScaling

`func (o *ApiSearchDeploymentRequest) GetAutoScaling() ApiSearchAutoScaling`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *ApiSearchDeploymentRequest) GetAutoScalingOk() (*ApiSearchAutoScaling, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *ApiSearchDeploymentRequest) SetAutoScaling(v ApiSearchAutoScaling)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *ApiSearchDeploymentRequest) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.

### SetAutoScalingNil

`func (o *ApiSearchDeploymentRequest) SetAutoScalingNil()`

SetAutoScalingNil sets AutoScaling to an explicit JSON null when marshaled, overriding any value previously set with SetAutoScaling. Calling SetAutoScaling again clears the null override.

### GetDefaultNodeCount

`func (o *ApiSearchDeploymentRequest) GetDefaultNodeCount() int`

GetDefaultNodeCount returns the DefaultNodeCount field if non-nil, zero value otherwise.

### GetDefaultNodeCountOk

`func (o *ApiSearchDeploymentRequest) GetDefaultNodeCountOk() (*int, bool)`

GetDefaultNodeCountOk returns a tuple with the DefaultNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNodeCount

`func (o *ApiSearchDeploymentRequest) SetDefaultNodeCount(v int)`

SetDefaultNodeCount sets DefaultNodeCount field to given value.

### HasDefaultNodeCount

`func (o *ApiSearchDeploymentRequest) HasDefaultNodeCount() bool`

HasDefaultNodeCount returns a boolean if a field has been set.

### SetDefaultNodeCountNil

`func (o *ApiSearchDeploymentRequest) SetDefaultNodeCountNil()`

SetDefaultNodeCountNil sets DefaultNodeCount to an explicit JSON null when marshaled, overriding any value previously set with SetDefaultNodeCount. Calling SetDefaultNodeCount again clears the null override.

### GetSpecs

`func (o *ApiSearchDeploymentRequest) GetSpecs() []ApiSearchDeploymentRequestSpec`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *ApiSearchDeploymentRequest) GetSpecsOk() (*[]ApiSearchDeploymentRequestSpec, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *ApiSearchDeploymentRequest) SetSpecs(v []ApiSearchDeploymentRequestSpec)`

SetSpecs sets Specs field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


