# CloudRegionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectableSpecs** | Pointer to [**HardwareSpec**](HardwareSpec.md) |  | [optional] 
**Priority** | Pointer to **int** | Precedence is given to this region when a primary election occurs. If your **regionConfigs** has only **readOnlySpecs**, **analyticsSpecs**, or both, set this value to &#x60;0&#x60;. If you have multiple **regionConfigs** objects (your cluster is multi-region or multi-cloud), they must have priorities in descending order. The highest priority is &#x60;7&#x60;.  **Example:** If you have three regions, their priorities would be &#x60;7&#x60;, &#x60;6&#x60;, and &#x60;5&#x60; respectively. If you added two more regions for supporting electable nodes, the priorities of those regions would be &#x60;4&#x60; and &#x60;3&#x60; respectively. | [optional] 
**ProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to &#x60;AWS&#x60;, &#x60;GCP&#x60;, &#x60;AZURE&#x60; or &#x60;TENANT&#x60;. | [optional] 
**RegionName** | Pointer to **string** | Physical location of your MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. It assigns the VPC a Classless Inter-Domain Routing (CIDR) block. To limit a new VPC peering connection to one Classless Inter-Domain Routing (CIDR) block and region, create the connection first. Deploy the cluster after the connection starts. GCP Clusters and Multi-region clusters require one VPC peering connection for each region. MongoDB nodes can use only the peering connection that resides in the same region as the nodes to communicate with the peered VPC. | [optional] 
**AnalyticsAutoScaling** | Pointer to [**AdvancedAutoScalingSettings**](AdvancedAutoScalingSettings.md) |  | [optional] 
**AnalyticsSpecs** | Pointer to [**DedicatedHardwareSpec**](DedicatedHardwareSpec.md) |  | [optional] 
**AutoScaling** | Pointer to [**AdvancedAutoScalingSettings**](AdvancedAutoScalingSettings.md) |  | [optional] 
**ReadOnlySpecs** | Pointer to [**DedicatedHardwareSpec**](DedicatedHardwareSpec.md) |  | [optional] 
**BackingProviderName** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisioned the multi-tenant cluster. The resource returns this parameter when **providerName** is &#x60;TENANT&#x60; and **electableSpecs.instanceSize** is &#x60;M0&#x60;, &#x60;M2&#x60; or &#x60;M5&#x60;. | [optional] 

## Methods

### NewCloudRegionConfig

`func NewCloudRegionConfig() *CloudRegionConfig`

NewCloudRegionConfig instantiates a new CloudRegionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRegionConfigWithDefaults

`func NewCloudRegionConfigWithDefaults() *CloudRegionConfig`

NewCloudRegionConfigWithDefaults instantiates a new CloudRegionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectableSpecs

`func (o *CloudRegionConfig) GetElectableSpecs() HardwareSpec`

GetElectableSpecs returns the ElectableSpecs field if non-nil, zero value otherwise.

### GetElectableSpecsOk

`func (o *CloudRegionConfig) GetElectableSpecsOk() (*HardwareSpec, bool)`

GetElectableSpecsOk returns a tuple with the ElectableSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableSpecs

`func (o *CloudRegionConfig) SetElectableSpecs(v HardwareSpec)`

SetElectableSpecs sets ElectableSpecs field to given value.

### HasElectableSpecs

`func (o *CloudRegionConfig) HasElectableSpecs() bool`

HasElectableSpecs returns a boolean if a field has been set.
### GetPriority

`func (o *CloudRegionConfig) GetPriority() int`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CloudRegionConfig) GetPriorityOk() (*int, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CloudRegionConfig) SetPriority(v int)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *CloudRegionConfig) HasPriority() bool`

HasPriority returns a boolean if a field has been set.
### GetProviderName

`func (o *CloudRegionConfig) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudRegionConfig) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudRegionConfig) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *CloudRegionConfig) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.
### GetRegionName

`func (o *CloudRegionConfig) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *CloudRegionConfig) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *CloudRegionConfig) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *CloudRegionConfig) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetAnalyticsAutoScaling

`func (o *CloudRegionConfig) GetAnalyticsAutoScaling() AdvancedAutoScalingSettings`

GetAnalyticsAutoScaling returns the AnalyticsAutoScaling field if non-nil, zero value otherwise.

### GetAnalyticsAutoScalingOk

`func (o *CloudRegionConfig) GetAnalyticsAutoScalingOk() (*AdvancedAutoScalingSettings, bool)`

GetAnalyticsAutoScalingOk returns a tuple with the AnalyticsAutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsAutoScaling

`func (o *CloudRegionConfig) SetAnalyticsAutoScaling(v AdvancedAutoScalingSettings)`

SetAnalyticsAutoScaling sets AnalyticsAutoScaling field to given value.

### HasAnalyticsAutoScaling

`func (o *CloudRegionConfig) HasAnalyticsAutoScaling() bool`

HasAnalyticsAutoScaling returns a boolean if a field has been set.
### GetAnalyticsSpecs

`func (o *CloudRegionConfig) GetAnalyticsSpecs() DedicatedHardwareSpec`

GetAnalyticsSpecs returns the AnalyticsSpecs field if non-nil, zero value otherwise.

### GetAnalyticsSpecsOk

`func (o *CloudRegionConfig) GetAnalyticsSpecsOk() (*DedicatedHardwareSpec, bool)`

GetAnalyticsSpecsOk returns a tuple with the AnalyticsSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsSpecs

`func (o *CloudRegionConfig) SetAnalyticsSpecs(v DedicatedHardwareSpec)`

SetAnalyticsSpecs sets AnalyticsSpecs field to given value.

### HasAnalyticsSpecs

`func (o *CloudRegionConfig) HasAnalyticsSpecs() bool`

HasAnalyticsSpecs returns a boolean if a field has been set.
### GetAutoScaling

`func (o *CloudRegionConfig) GetAutoScaling() AdvancedAutoScalingSettings`

GetAutoScaling returns the AutoScaling field if non-nil, zero value otherwise.

### GetAutoScalingOk

`func (o *CloudRegionConfig) GetAutoScalingOk() (*AdvancedAutoScalingSettings, bool)`

GetAutoScalingOk returns a tuple with the AutoScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaling

`func (o *CloudRegionConfig) SetAutoScaling(v AdvancedAutoScalingSettings)`

SetAutoScaling sets AutoScaling field to given value.

### HasAutoScaling

`func (o *CloudRegionConfig) HasAutoScaling() bool`

HasAutoScaling returns a boolean if a field has been set.
### GetReadOnlySpecs

`func (o *CloudRegionConfig) GetReadOnlySpecs() DedicatedHardwareSpec`

GetReadOnlySpecs returns the ReadOnlySpecs field if non-nil, zero value otherwise.

### GetReadOnlySpecsOk

`func (o *CloudRegionConfig) GetReadOnlySpecsOk() (*DedicatedHardwareSpec, bool)`

GetReadOnlySpecsOk returns a tuple with the ReadOnlySpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlySpecs

`func (o *CloudRegionConfig) SetReadOnlySpecs(v DedicatedHardwareSpec)`

SetReadOnlySpecs sets ReadOnlySpecs field to given value.

### HasReadOnlySpecs

`func (o *CloudRegionConfig) HasReadOnlySpecs() bool`

HasReadOnlySpecs returns a boolean if a field has been set.
### GetBackingProviderName

`func (o *CloudRegionConfig) GetBackingProviderName() string`

GetBackingProviderName returns the BackingProviderName field if non-nil, zero value otherwise.

### GetBackingProviderNameOk

`func (o *CloudRegionConfig) GetBackingProviderNameOk() (*string, bool)`

GetBackingProviderNameOk returns a tuple with the BackingProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingProviderName

`func (o *CloudRegionConfig) SetBackingProviderName(v string)`

SetBackingProviderName sets BackingProviderName field to given value.

### HasBackingProviderName

`func (o *CloudRegionConfig) HasBackingProviderName() bool`

HasBackingProviderName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


