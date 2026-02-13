# RegionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsNodes** | Pointer to **int** | Number of analytics nodes in the region. Analytics nodes handle analytic data such as reporting queries from MongoDB Connector for Business Intelligence on MongoDB Cloud. Analytics nodes are read-only, and can never become the primary. Use &#x60;replicationSpecs[n].{region}.analyticsNodes&#x60; instead. | [optional] 
**ElectableNodes** | Pointer to **int** | Number of electable nodes to deploy in the specified region. Electable nodes can become the primary and can facilitate local reads. Use &#x60;replicationSpecs[n].{region}.electableNodes&#x60; instead. | [optional] 
**Priority** | Pointer to **int** | Number that indicates the election priority of the region. To identify the Preferred Region of the cluster, set this parameter to &#x60;7&#x60;. The primary node runs in the **Preferred Region**. To identify a read-only region, set this parameter to &#x60;0&#x60;. | [optional] 
**ReadOnlyNodes** | Pointer to **int** | Number of read-only nodes in the region. Read-only nodes can never become the primary member, but can facilitate local reads. Use &#x60;replicationSpecs[n].{region}.readOnlyNodes&#x60; instead. | [optional] 

## Methods

### NewRegionSpec

`func NewRegionSpec() *RegionSpec`

NewRegionSpec instantiates a new RegionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionSpecWithDefaults

`func NewRegionSpecWithDefaults() *RegionSpec`

NewRegionSpecWithDefaults instantiates a new RegionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsNodes

`func (o *RegionSpec) GetAnalyticsNodes() int`

GetAnalyticsNodes returns the AnalyticsNodes field if non-nil, zero value otherwise.

### GetAnalyticsNodesOk

`func (o *RegionSpec) GetAnalyticsNodesOk() (*int, bool)`

GetAnalyticsNodesOk returns a tuple with the AnalyticsNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsNodes

`func (o *RegionSpec) SetAnalyticsNodes(v int)`

SetAnalyticsNodes sets AnalyticsNodes field to given value.

### HasAnalyticsNodes

`func (o *RegionSpec) HasAnalyticsNodes() bool`

HasAnalyticsNodes returns a boolean if a field has been set.
### GetElectableNodes

`func (o *RegionSpec) GetElectableNodes() int`

GetElectableNodes returns the ElectableNodes field if non-nil, zero value otherwise.

### GetElectableNodesOk

`func (o *RegionSpec) GetElectableNodesOk() (*int, bool)`

GetElectableNodesOk returns a tuple with the ElectableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectableNodes

`func (o *RegionSpec) SetElectableNodes(v int)`

SetElectableNodes sets ElectableNodes field to given value.

### HasElectableNodes

`func (o *RegionSpec) HasElectableNodes() bool`

HasElectableNodes returns a boolean if a field has been set.
### GetPriority

`func (o *RegionSpec) GetPriority() int`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RegionSpec) GetPriorityOk() (*int, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RegionSpec) SetPriority(v int)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RegionSpec) HasPriority() bool`

HasPriority returns a boolean if a field has been set.
### GetReadOnlyNodes

`func (o *RegionSpec) GetReadOnlyNodes() int`

GetReadOnlyNodes returns the ReadOnlyNodes field if non-nil, zero value otherwise.

### GetReadOnlyNodesOk

`func (o *RegionSpec) GetReadOnlyNodesOk() (*int, bool)`

GetReadOnlyNodesOk returns a tuple with the ReadOnlyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyNodes

`func (o *RegionSpec) SetReadOnlyNodes(v int)`

SetReadOnlyNodes sets ReadOnlyNodes field to given value.

### HasReadOnlyNodes

`func (o *RegionSpec) HasReadOnlyNodes() bool`

HasReadOnlyNodes returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


