# ApiSearchDeploymentRequestSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Cloud service provider that hosts the Search Nodes in this region. Required when a region is specified. | [optional] 
**InstanceSize** | **string** | Hardware specification for the Search Node instance sizes. | 
**NodeCount** | Pointer to **int** | Number of Search Nodes in this region. Optional; falls back to the request-level default when omitted. | [optional] 
**RegionName** | Pointer to **string** | Cloud provider region where Search Nodes are provisioned. Required when the request configures more than one region; optional for single-region requests. | [optional] 

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

### GetCloudProvider

`func (o *ApiSearchDeploymentRequestSpec) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiSearchDeploymentRequestSpec) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiSearchDeploymentRequestSpec) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ApiSearchDeploymentRequestSpec) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### SetCloudProviderNil

`func (o *ApiSearchDeploymentRequestSpec) SetCloudProviderNil()`

SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled, overriding any value previously set with SetCloudProvider. Calling SetCloudProvider again clears the null override.

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

### SetNodeCountNil

`func (o *ApiSearchDeploymentRequestSpec) SetNodeCountNil()`

SetNodeCountNil sets NodeCount to an explicit JSON null when marshaled, overriding any value previously set with SetNodeCount. Calling SetNodeCount again clears the null override.

### GetRegionName

`func (o *ApiSearchDeploymentRequestSpec) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiSearchDeploymentRequestSpec) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiSearchDeploymentRequestSpec) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiSearchDeploymentRequestSpec) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### SetRegionNameNil

`func (o *ApiSearchDeploymentRequestSpec) SetRegionNameNil()`

SetRegionNameNil sets RegionName to an explicit JSON null when marshaled, overriding any value previously set with SetRegionName. Calling SetRegionName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


