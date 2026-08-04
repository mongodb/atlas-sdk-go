# ApiAtlasCollectionRestoreJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionSuffix** | Pointer to **string** | Optional suffix applied to restored collection names. | [optional] 
**Collections** | Pointer to [**[]ApiAtlasRestoreNamespace**](ApiAtlasRestoreNamespace.md) | List of collections to restore (up to 100 items). | [optional] 
**DatabaseSuffix** | Pointer to **string** | Optional suffix applied to restored database names. | [optional] 
**Databases** | Pointer to [**[]ApiAtlasRestoreNamespace**](ApiAtlasRestoreNamespace.md) | List of databases to restore (up to 100 items). | [optional] 
**IndexStrategy** | **string** | Strategy for restoring indexes (all, none, or all except TTL). | 
**OplogInc** | Pointer to **int** | Oplog increment for point-in-time restore. | [optional] 
**OplogTs** | Pointer to **int** | Oplog timestamp (seconds part) for point-in-time restore. | [optional] 
**PointInTimeUtcSeconds** | Pointer to **int** | Point-in-time restore time in seconds since UNIX epoch. | [optional] 
**SnapshotId** | Pointer to **string** | ID of the snapshot to restore. | [optional] 
**TargetClusterName** | **string** | Target cluster name. | 
**TargetGroupId** | **string** | Unique 24-hexadecimal digit string that identifies the target group. | 
**WriteStrategy** | **string** | Strategy for writing data on the target (create as new or overwrite existing). With &#x60;OVERWRITE_EXISTING&#x60;, any writes to the affected databases or collections during the restore will be lost when the existing namespaces are dropped and replaced. To avoid data loss, stop writes to the affected namespaces before starting the restore. | 

## Methods

### NewApiAtlasCollectionRestoreJobRequest

`func NewApiAtlasCollectionRestoreJobRequest(indexStrategy string, targetClusterName string, targetGroupId string, writeStrategy string, ) *ApiAtlasCollectionRestoreJobRequest`

NewApiAtlasCollectionRestoreJobRequest instantiates a new ApiAtlasCollectionRestoreJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCollectionRestoreJobRequestWithDefaults

`func NewApiAtlasCollectionRestoreJobRequestWithDefaults() *ApiAtlasCollectionRestoreJobRequest`

NewApiAtlasCollectionRestoreJobRequestWithDefaults instantiates a new ApiAtlasCollectionRestoreJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionSuffix

`func (o *ApiAtlasCollectionRestoreJobRequest) GetCollectionSuffix() string`

GetCollectionSuffix returns the CollectionSuffix field if non-nil, zero value otherwise.

### GetCollectionSuffixOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetCollectionSuffixOk() (*string, bool)`

GetCollectionSuffixOk returns a tuple with the CollectionSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionSuffix

`func (o *ApiAtlasCollectionRestoreJobRequest) SetCollectionSuffix(v string)`

SetCollectionSuffix sets CollectionSuffix field to given value.

### HasCollectionSuffix

`func (o *ApiAtlasCollectionRestoreJobRequest) HasCollectionSuffix() bool`

HasCollectionSuffix returns a boolean if a field has been set.

### SetCollectionSuffixNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetCollectionSuffixNil()`

SetCollectionSuffixNil sets CollectionSuffix to an explicit JSON null when marshaled, overriding any value previously set with SetCollectionSuffix. Calling SetCollectionSuffix again clears the null override.

### GetCollections

`func (o *ApiAtlasCollectionRestoreJobRequest) GetCollections() []ApiAtlasRestoreNamespace`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetCollectionsOk() (*[]ApiAtlasRestoreNamespace, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ApiAtlasCollectionRestoreJobRequest) SetCollections(v []ApiAtlasRestoreNamespace)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *ApiAtlasCollectionRestoreJobRequest) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### SetCollectionsNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetCollectionsNil()`

SetCollectionsNil sets Collections to an explicit JSON null when marshaled, overriding any value previously set with SetCollections. Calling SetCollections again clears the null override.

### GetDatabaseSuffix

`func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabaseSuffix() string`

GetDatabaseSuffix returns the DatabaseSuffix field if non-nil, zero value otherwise.

### GetDatabaseSuffixOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabaseSuffixOk() (*string, bool)`

GetDatabaseSuffixOk returns a tuple with the DatabaseSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSuffix

`func (o *ApiAtlasCollectionRestoreJobRequest) SetDatabaseSuffix(v string)`

SetDatabaseSuffix sets DatabaseSuffix field to given value.

### HasDatabaseSuffix

`func (o *ApiAtlasCollectionRestoreJobRequest) HasDatabaseSuffix() bool`

HasDatabaseSuffix returns a boolean if a field has been set.

### SetDatabaseSuffixNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetDatabaseSuffixNil()`

SetDatabaseSuffixNil sets DatabaseSuffix to an explicit JSON null when marshaled, overriding any value previously set with SetDatabaseSuffix. Calling SetDatabaseSuffix again clears the null override.

### GetDatabases

`func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabases() []ApiAtlasRestoreNamespace`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabasesOk() (*[]ApiAtlasRestoreNamespace, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *ApiAtlasCollectionRestoreJobRequest) SetDatabases(v []ApiAtlasRestoreNamespace)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *ApiAtlasCollectionRestoreJobRequest) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### SetDatabasesNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetDatabasesNil()`

SetDatabasesNil sets Databases to an explicit JSON null when marshaled, overriding any value previously set with SetDatabases. Calling SetDatabases again clears the null override.

### GetIndexStrategy

`func (o *ApiAtlasCollectionRestoreJobRequest) GetIndexStrategy() string`

GetIndexStrategy returns the IndexStrategy field if non-nil, zero value otherwise.

### GetIndexStrategyOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetIndexStrategyOk() (*string, bool)`

GetIndexStrategyOk returns a tuple with the IndexStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStrategy

`func (o *ApiAtlasCollectionRestoreJobRequest) SetIndexStrategy(v string)`

SetIndexStrategy sets IndexStrategy field to given value.

### GetOplogInc

`func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogInc() int`

GetOplogInc returns the OplogInc field if non-nil, zero value otherwise.

### GetOplogIncOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogIncOk() (*int, bool)`

GetOplogIncOk returns a tuple with the OplogInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogInc

`func (o *ApiAtlasCollectionRestoreJobRequest) SetOplogInc(v int)`

SetOplogInc sets OplogInc field to given value.

### HasOplogInc

`func (o *ApiAtlasCollectionRestoreJobRequest) HasOplogInc() bool`

HasOplogInc returns a boolean if a field has been set.

### SetOplogIncNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetOplogIncNil()`

SetOplogIncNil sets OplogInc to an explicit JSON null when marshaled, overriding any value previously set with SetOplogInc. Calling SetOplogInc again clears the null override.

### GetOplogTs

`func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogTs() int`

GetOplogTs returns the OplogTs field if non-nil, zero value otherwise.

### GetOplogTsOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogTsOk() (*int, bool)`

GetOplogTsOk returns a tuple with the OplogTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogTs

`func (o *ApiAtlasCollectionRestoreJobRequest) SetOplogTs(v int)`

SetOplogTs sets OplogTs field to given value.

### HasOplogTs

`func (o *ApiAtlasCollectionRestoreJobRequest) HasOplogTs() bool`

HasOplogTs returns a boolean if a field has been set.

### SetOplogTsNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetOplogTsNil()`

SetOplogTsNil sets OplogTs to an explicit JSON null when marshaled, overriding any value previously set with SetOplogTs. Calling SetOplogTs again clears the null override.

### GetPointInTimeUtcSeconds

`func (o *ApiAtlasCollectionRestoreJobRequest) GetPointInTimeUtcSeconds() int`

GetPointInTimeUtcSeconds returns the PointInTimeUtcSeconds field if non-nil, zero value otherwise.

### GetPointInTimeUtcSecondsOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetPointInTimeUtcSecondsOk() (*int, bool)`

GetPointInTimeUtcSecondsOk returns a tuple with the PointInTimeUtcSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUtcSeconds

`func (o *ApiAtlasCollectionRestoreJobRequest) SetPointInTimeUtcSeconds(v int)`

SetPointInTimeUtcSeconds sets PointInTimeUtcSeconds field to given value.

### HasPointInTimeUtcSeconds

`func (o *ApiAtlasCollectionRestoreJobRequest) HasPointInTimeUtcSeconds() bool`

HasPointInTimeUtcSeconds returns a boolean if a field has been set.

### SetPointInTimeUtcSecondsNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetPointInTimeUtcSecondsNil()`

SetPointInTimeUtcSecondsNil sets PointInTimeUtcSeconds to an explicit JSON null when marshaled, overriding any value previously set with SetPointInTimeUtcSeconds. Calling SetPointInTimeUtcSeconds again clears the null override.

### GetSnapshotId

`func (o *ApiAtlasCollectionRestoreJobRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasCollectionRestoreJobRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ApiAtlasCollectionRestoreJobRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ApiAtlasCollectionRestoreJobRequest) SetSnapshotIdNil()`

SetSnapshotIdNil sets SnapshotId to an explicit JSON null when marshaled, overriding any value previously set with SetSnapshotId. Calling SetSnapshotId again clears the null override.

### GetTargetClusterName

`func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *ApiAtlasCollectionRestoreJobRequest) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### GetTargetGroupId

`func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ApiAtlasCollectionRestoreJobRequest) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### GetWriteStrategy

`func (o *ApiAtlasCollectionRestoreJobRequest) GetWriteStrategy() string`

GetWriteStrategy returns the WriteStrategy field if non-nil, zero value otherwise.

### GetWriteStrategyOk

`func (o *ApiAtlasCollectionRestoreJobRequest) GetWriteStrategyOk() (*string, bool)`

GetWriteStrategyOk returns a tuple with the WriteStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteStrategy

`func (o *ApiAtlasCollectionRestoreJobRequest) SetWriteStrategy(v string)`

SetWriteStrategy sets WriteStrategy field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


