# DataLakeAtlasStoreInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable label of the MongoDB Cloud cluster on which the store is based. | [optional] 
**ProjectId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project. | [optional] [readonly] 
**ReadPreference** | Pointer to [**DataLakeAtlasStoreReadPreference**](DataLakeAtlasStoreReadPreference.md) |  | [optional] 
**Name** | Pointer to **string** | Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an &#x60;M10&#x60; or higher cluster. | [optional] 
**Provider** | **string** |  | 

## Methods

### NewDataLakeAtlasStoreInstance

`func NewDataLakeAtlasStoreInstance(provider string, ) *DataLakeAtlasStoreInstance`

NewDataLakeAtlasStoreInstance instantiates a new DataLakeAtlasStoreInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeAtlasStoreInstanceWithDefaults

`func NewDataLakeAtlasStoreInstanceWithDefaults() *DataLakeAtlasStoreInstance`

NewDataLakeAtlasStoreInstanceWithDefaults instantiates a new DataLakeAtlasStoreInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *DataLakeAtlasStoreInstance) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DataLakeAtlasStoreInstance) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DataLakeAtlasStoreInstance) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *DataLakeAtlasStoreInstance) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetProjectId

`func (o *DataLakeAtlasStoreInstance) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DataLakeAtlasStoreInstance) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DataLakeAtlasStoreInstance) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DataLakeAtlasStoreInstance) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReadPreference

`func (o *DataLakeAtlasStoreInstance) GetReadPreference() DataLakeAtlasStoreReadPreference`

GetReadPreference returns the ReadPreference field if non-nil, zero value otherwise.

### GetReadPreferenceOk

`func (o *DataLakeAtlasStoreInstance) GetReadPreferenceOk() (*DataLakeAtlasStoreReadPreference, bool)`

GetReadPreferenceOk returns a tuple with the ReadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPreference

`func (o *DataLakeAtlasStoreInstance) SetReadPreference(v DataLakeAtlasStoreReadPreference)`

SetReadPreference sets ReadPreference field to given value.

### HasReadPreference

`func (o *DataLakeAtlasStoreInstance) HasReadPreference() bool`

HasReadPreference returns a boolean if a field has been set.

### GetName

`func (o *DataLakeAtlasStoreInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeAtlasStoreInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeAtlasStoreInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeAtlasStoreInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DataLakeAtlasStoreInstance) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DataLakeAtlasStoreInstance) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DataLakeAtlasStoreInstance) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


