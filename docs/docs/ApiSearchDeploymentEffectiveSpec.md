# ApiSearchDeploymentEffectiveSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Cloud service provider on which Search Nodes are provisioned. | [optional] [readonly] 
**InstanceSize** | Pointer to **string** | Hardware specification for the Search Node instance sizes. | [optional] [readonly] 
**NodeCount** | Pointer to **int** | Number of Search Nodes in this region. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region where Search Nodes are provisioned. | [optional] [readonly] 

## Methods

### NewApiSearchDeploymentEffectiveSpec

`func NewApiSearchDeploymentEffectiveSpec() *ApiSearchDeploymentEffectiveSpec`

NewApiSearchDeploymentEffectiveSpec instantiates a new ApiSearchDeploymentEffectiveSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchDeploymentEffectiveSpecWithDefaults

`func NewApiSearchDeploymentEffectiveSpecWithDefaults() *ApiSearchDeploymentEffectiveSpec`

NewApiSearchDeploymentEffectiveSpecWithDefaults instantiates a new ApiSearchDeploymentEffectiveSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ApiSearchDeploymentEffectiveSpec) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiSearchDeploymentEffectiveSpec) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiSearchDeploymentEffectiveSpec) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ApiSearchDeploymentEffectiveSpec) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### SetCloudProviderNil

`func (o *ApiSearchDeploymentEffectiveSpec) SetCloudProviderNil()`

SetCloudProviderNil sets CloudProvider to an explicit JSON null when marshaled, overriding any value previously set with SetCloudProvider. Calling SetCloudProvider again clears the null override.

### GetInstanceSize

`func (o *ApiSearchDeploymentEffectiveSpec) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *ApiSearchDeploymentEffectiveSpec) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *ApiSearchDeploymentEffectiveSpec) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *ApiSearchDeploymentEffectiveSpec) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.

### SetInstanceSizeNil

`func (o *ApiSearchDeploymentEffectiveSpec) SetInstanceSizeNil()`

SetInstanceSizeNil sets InstanceSize to an explicit JSON null when marshaled, overriding any value previously set with SetInstanceSize. Calling SetInstanceSize again clears the null override.

### GetNodeCount

`func (o *ApiSearchDeploymentEffectiveSpec) GetNodeCount() int`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ApiSearchDeploymentEffectiveSpec) GetNodeCountOk() (*int, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ApiSearchDeploymentEffectiveSpec) SetNodeCount(v int)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ApiSearchDeploymentEffectiveSpec) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### SetNodeCountNil

`func (o *ApiSearchDeploymentEffectiveSpec) SetNodeCountNil()`

SetNodeCountNil sets NodeCount to an explicit JSON null when marshaled, overriding any value previously set with SetNodeCount. Calling SetNodeCount again clears the null override.

### GetRegionName

`func (o *ApiSearchDeploymentEffectiveSpec) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ApiSearchDeploymentEffectiveSpec) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ApiSearchDeploymentEffectiveSpec) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ApiSearchDeploymentEffectiveSpec) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### SetRegionNameNil

`func (o *ApiSearchDeploymentEffectiveSpec) SetRegionNameNil()`

SetRegionNameNil sets RegionName to an explicit JSON null when marshaled, overriding any value previously set with SetRegionName. Calling SetRegionName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


