# ApiSearchDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


