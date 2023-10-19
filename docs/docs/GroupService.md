# GroupService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]ClusterIPAddresses**](ClusterIPAddresses.md) | IP addresses of clusters. | [optional] [readonly] 

## Methods

### NewGroupService

`func NewGroupService() *GroupService`

NewGroupService instantiates a new GroupService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupServiceWithDefaults

`func NewGroupServiceWithDefaults() *GroupService`

NewGroupServiceWithDefaults instantiates a new GroupService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *GroupService) GetClusters() []ClusterIPAddresses`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *GroupService) GetClustersOk() (*[]ClusterIPAddresses, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *GroupService) SetClusters(v []ClusterIPAddresses)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *GroupService) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


