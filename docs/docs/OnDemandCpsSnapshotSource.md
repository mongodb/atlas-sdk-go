# OnDemandCpsSnapshotSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | Pointer to **string** | Human-readable name that identifies the cluster. | [optional] 
**CollectionName** | Pointer to **string** | Human-readable name that identifies the collection. | [optional] 
**DatabaseName** | Pointer to **string** | Human-readable name that identifies the database. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of ingestion source of this Data Lake Pipeline. | [optional] 

## Methods

### NewOnDemandCpsSnapshotSource

`func NewOnDemandCpsSnapshotSource() *OnDemandCpsSnapshotSource`

NewOnDemandCpsSnapshotSource instantiates a new OnDemandCpsSnapshotSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnDemandCpsSnapshotSourceWithDefaults

`func NewOnDemandCpsSnapshotSourceWithDefaults() *OnDemandCpsSnapshotSource`

NewOnDemandCpsSnapshotSourceWithDefaults instantiates a new OnDemandCpsSnapshotSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *OnDemandCpsSnapshotSource) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *OnDemandCpsSnapshotSource) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *OnDemandCpsSnapshotSource) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *OnDemandCpsSnapshotSource) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCollectionName

`func (o *OnDemandCpsSnapshotSource) GetCollectionName() string`

GetCollectionName returns the CollectionName field if non-nil, zero value otherwise.

### GetCollectionNameOk

`func (o *OnDemandCpsSnapshotSource) GetCollectionNameOk() (*string, bool)`

GetCollectionNameOk returns a tuple with the CollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionName

`func (o *OnDemandCpsSnapshotSource) SetCollectionName(v string)`

SetCollectionName sets CollectionName field to given value.

### HasCollectionName

`func (o *OnDemandCpsSnapshotSource) HasCollectionName() bool`

HasCollectionName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *OnDemandCpsSnapshotSource) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *OnDemandCpsSnapshotSource) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *OnDemandCpsSnapshotSource) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *OnDemandCpsSnapshotSource) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetGroupId

`func (o *OnDemandCpsSnapshotSource) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OnDemandCpsSnapshotSource) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OnDemandCpsSnapshotSource) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OnDemandCpsSnapshotSource) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *OnDemandCpsSnapshotSource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnDemandCpsSnapshotSource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnDemandCpsSnapshotSource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OnDemandCpsSnapshotSource) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


