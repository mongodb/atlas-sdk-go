# ApiSearchDeploymentRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSize** | **string** | Hardware specification for the Search Node instance sizes. | 
**NodeCount** | Pointer to **int** | Number of Search Nodes in this region. Optional; falls back to the request-level default when omitted. | [optional] 

## Methods

### NewApiSearchDeploymentRequestSpec

`func NewApiSearchDeploymentRequestSpec(instanceSize string, ) *ApiSearchDeploymentRequestSpec`

NewApiSearchDeploymentRequestSpec instantiates a new ApiSearchDeploymentRequestSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchDeploymentRequestSpecWithDefaults

`func NewApiSearchDeploymentRequestSpecWithDefaults() *ApiSearchDeploymentRequestSpec`

NewApiSearchDeploymentRequestSpecWithDefaults instantiates a new ApiSearchDeploymentRequestSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSize

`func (o *ApiSearchDeploymentRequestSpec) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ApiSearchDeploymentRequestSpec) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ApiSearchDeploymentRequestSpec) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### GetNodeCount

`func (o *ApiSearchDeploymentRequestSpec) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ApiSearchDeploymentRequestSpec) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ApiSearchDeploymentRequestSpec) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ApiSearchDeploymentRequestSpec) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


