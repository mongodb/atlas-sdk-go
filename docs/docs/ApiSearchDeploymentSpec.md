# ApiSearchDeploymentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSize** | **string** | Hardware specification for the search node instance sizes. | 
**NodeCount** | **int** | Number of search nodes in the cluster. | 

## Methods

### NewApiSearchDeploymentSpec

`func NewApiSearchDeploymentSpec(instanceSize string, nodeCount int, ) *ApiSearchDeploymentSpec`

NewApiSearchDeploymentSpec instantiates a new ApiSearchDeploymentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchDeploymentSpecWithDefaults

`func NewApiSearchDeploymentSpecWithDefaults() *ApiSearchDeploymentSpec`

NewApiSearchDeploymentSpecWithDefaults instantiates a new ApiSearchDeploymentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSize

`func (o *ApiSearchDeploymentSpec) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ApiSearchDeploymentSpec) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ApiSearchDeploymentSpec) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### GetNodeCount

`func (o *ApiSearchDeploymentSpec) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ApiSearchDeploymentSpec) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ApiSearchDeploymentSpec) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


