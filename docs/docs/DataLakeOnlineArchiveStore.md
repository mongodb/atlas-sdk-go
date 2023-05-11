# DataLakeOnlineArchiveStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | ID of the Cluster the Online Archive belongs to. | 
**ClusterName** | **string** | Name of the Cluster the Online Archive belongs to. | 
**ProjectId** | **string** | ID of the Project the Online Archive belongs to. | 
**Name** | Pointer to **string** | Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an &#x60;M10&#x60; or higher cluster. | [optional] 
**Provider** | **string** |  | 

## Methods

### NewDataLakeOnlineArchiveStore

`func NewDataLakeOnlineArchiveStore(clusterId string, clusterName string, projectId string, provider string, ) *DataLakeOnlineArchiveStore`

NewDataLakeOnlineArchiveStore instantiates a new DataLakeOnlineArchiveStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeOnlineArchiveStoreWithDefaults

`func NewDataLakeOnlineArchiveStoreWithDefaults() *DataLakeOnlineArchiveStore`

NewDataLakeOnlineArchiveStoreWithDefaults instantiates a new DataLakeOnlineArchiveStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *DataLakeOnlineArchiveStore) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DataLakeOnlineArchiveStore) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DataLakeOnlineArchiveStore) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *DataLakeOnlineArchiveStore) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DataLakeOnlineArchiveStore) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DataLakeOnlineArchiveStore) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetProjectId

`func (o *DataLakeOnlineArchiveStore) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DataLakeOnlineArchiveStore) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DataLakeOnlineArchiveStore) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *DataLakeOnlineArchiveStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeOnlineArchiveStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeOnlineArchiveStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeOnlineArchiveStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DataLakeOnlineArchiveStore) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DataLakeOnlineArchiveStore) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DataLakeOnlineArchiveStore) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


