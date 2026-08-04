# ApiAtlasCollectionRestoreJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionSuffix** | Pointer to **string** | Suffix applied to restored collection names. | [optional] 
**Collections** | Pointer to [**[]ApiAtlasRestoreNamespace**](ApiAtlasRestoreNamespace.md) | List of collections in the restore scope (up to 100 items). | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time when the restore job was created (ISO 8601 format in UTC). | [optional] [readonly] 
**DatabaseSuffix** | Pointer to **string** | Suffix applied to restored database names. | [optional] 
**Databases** | Pointer to [**[]ApiAtlasRestoreNamespace**](ApiAtlasRestoreNamespace.md) | List of databases in the restore scope (up to 100 items). | [optional] 
**ErrorMessage** | Pointer to **string** | Error message when the job has failed or been canceled. | [optional] [readonly] 
**FinishedAt** | Pointer to **time.Time** | Date and time when the restore job finished (ISO 8601 format in UTC). | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the collection restore job. | [optional] [readonly] 
**IndexStatus** | Pointer to [**ApiAtlasCollectionRestoreJobIndexStatus**](ApiAtlasCollectionRestoreJobIndexStatus.md) |  | [optional] 
**IndexStrategy** | Pointer to **string** | Strategy for restoring indexes (all, none, or all except TTL). | [optional] 
**OplogInc** | Pointer to **int** | Oplog increment for point-in-time restore. | [optional] 
**OplogTs** | Pointer to **int** | Oplog timestamp (seconds part) for point-in-time restore. | [optional] 
**PointInTimeUtcSeconds** | Pointer to **int** | Point-in-time restore time in seconds since UNIX epoch. | [optional] 
**RestoredDocuments** | Pointer to **int64** | Number of documents restored so far across all supported collections. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot being restored. | [optional] 
**State** | Pointer to **string** | Current state of the collection restore job. | [optional] [readonly] 
**TargetClusterName** | Pointer to **string** | Human-readable label that identifies the target cluster. | [optional] 
**TargetGroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the target group. | [optional] 
**TotalDocuments** | Pointer to **int64** | Total number of documents across all supported collections in the restore job. This value may initially reflect an estimate based on collection metadata and can change as accurate document counts become available during the restore. | [optional] [readonly] 
**WriteStrategy** | Pointer to **string** | Strategy for writing data on the target (create as new or overwrite existing). With &#x60;OVERWRITE_EXISTING&#x60;, any writes to the affected databases or collections during the restore will be lost when the existing namespaces are dropped and replaced. To avoid data loss, stop writes to the affected namespaces before starting the restore. | [optional] 

## Methods

### NewApiAtlasCollectionRestoreJobResponse

`func NewApiAtlasCollectionRestoreJobResponse() *ApiAtlasCollectionRestoreJobResponse`

NewApiAtlasCollectionRestoreJobResponse instantiates a new ApiAtlasCollectionRestoreJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasCollectionRestoreJobResponseWithDefaults

`func NewApiAtlasCollectionRestoreJobResponseWithDefaults() *ApiAtlasCollectionRestoreJobResponse`

NewApiAtlasCollectionRestoreJobResponseWithDefaults instantiates a new ApiAtlasCollectionRestoreJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionSuffix

`func (o *ApiAtlasCollectionRestoreJobResponse) GetCollectionSuffix() string`

GetCollectionSuffix returns the CollectionSuffix field if non-nil, zero value otherwise.

### GetCollectionSuffixOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetCollectionSuffixOk() (*string, bool)`

GetCollectionSuffixOk returns a tuple with the CollectionSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionSuffix

`func (o *ApiAtlasCollectionRestoreJobResponse) SetCollectionSuffix(v string)`

SetCollectionSuffix sets CollectionSuffix field to given value.

### HasCollectionSuffix

`func (o *ApiAtlasCollectionRestoreJobResponse) HasCollectionSuffix() bool`

HasCollectionSuffix returns a boolean if a field has been set.

### SetCollectionSuffixNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetCollectionSuffixNil()`

SetCollectionSuffixNil sets CollectionSuffix to an explicit JSON null when marshaled, overriding any value previously set with SetCollectionSuffix. Calling SetCollectionSuffix again clears the null override.

### GetCollections

`func (o *ApiAtlasCollectionRestoreJobResponse) GetCollections() []ApiAtlasRestoreNamespace`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetCollectionsOk() (*[]ApiAtlasRestoreNamespace, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ApiAtlasCollectionRestoreJobResponse) SetCollections(v []ApiAtlasRestoreNamespace)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *ApiAtlasCollectionRestoreJobResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### SetCollectionsNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetCollectionsNil()`

SetCollectionsNil sets Collections to an explicit JSON null when marshaled, overriding any value previously set with SetCollections. Calling SetCollections again clears the null override.

### GetCreatedAt

`func (o *ApiAtlasCollectionRestoreJobResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiAtlasCollectionRestoreJobResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiAtlasCollectionRestoreJobResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetCreatedAtNil()`

SetCreatedAtNil sets CreatedAt to an explicit JSON null when marshaled, overriding any value previously set with SetCreatedAt. Calling SetCreatedAt again clears the null override.

### GetDatabaseSuffix

`func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabaseSuffix() string`

GetDatabaseSuffix returns the DatabaseSuffix field if non-nil, zero value otherwise.

### GetDatabaseSuffixOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabaseSuffixOk() (*string, bool)`

GetDatabaseSuffixOk returns a tuple with the DatabaseSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSuffix

`func (o *ApiAtlasCollectionRestoreJobResponse) SetDatabaseSuffix(v string)`

SetDatabaseSuffix sets DatabaseSuffix field to given value.

### HasDatabaseSuffix

`func (o *ApiAtlasCollectionRestoreJobResponse) HasDatabaseSuffix() bool`

HasDatabaseSuffix returns a boolean if a field has been set.

### SetDatabaseSuffixNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetDatabaseSuffixNil()`

SetDatabaseSuffixNil sets DatabaseSuffix to an explicit JSON null when marshaled, overriding any value previously set with SetDatabaseSuffix. Calling SetDatabaseSuffix again clears the null override.

### GetDatabases

`func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabases() []ApiAtlasRestoreNamespace`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabasesOk() (*[]ApiAtlasRestoreNamespace, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *ApiAtlasCollectionRestoreJobResponse) SetDatabases(v []ApiAtlasRestoreNamespace)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *ApiAtlasCollectionRestoreJobResponse) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### SetDatabasesNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetDatabasesNil()`

SetDatabasesNil sets Databases to an explicit JSON null when marshaled, overriding any value previously set with SetDatabases. Calling SetDatabases again clears the null override.

### GetErrorMessage

`func (o *ApiAtlasCollectionRestoreJobResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiAtlasCollectionRestoreJobResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiAtlasCollectionRestoreJobResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetErrorMessageNil()`

SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled, overriding any value previously set with SetErrorMessage. Calling SetErrorMessage again clears the null override.

### GetFinishedAt

`func (o *ApiAtlasCollectionRestoreJobResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiAtlasCollectionRestoreJobResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiAtlasCollectionRestoreJobResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetFinishedAtNil()`

SetFinishedAtNil sets FinishedAt to an explicit JSON null when marshaled, overriding any value previously set with SetFinishedAt. Calling SetFinishedAt again clears the null override.

### GetId

`func (o *ApiAtlasCollectionRestoreJobResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiAtlasCollectionRestoreJobResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiAtlasCollectionRestoreJobResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetIndexStatus

`func (o *ApiAtlasCollectionRestoreJobResponse) GetIndexStatus() ApiAtlasCollectionRestoreJobIndexStatus`

GetIndexStatus returns the IndexStatus field if non-nil, zero value otherwise.

### GetIndexStatusOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetIndexStatusOk() (*ApiAtlasCollectionRestoreJobIndexStatus, bool)`

GetIndexStatusOk returns a tuple with the IndexStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStatus

`func (o *ApiAtlasCollectionRestoreJobResponse) SetIndexStatus(v ApiAtlasCollectionRestoreJobIndexStatus)`

SetIndexStatus sets IndexStatus field to given value.

### HasIndexStatus

`func (o *ApiAtlasCollectionRestoreJobResponse) HasIndexStatus() bool`

HasIndexStatus returns a boolean if a field has been set.

### SetIndexStatusNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetIndexStatusNil()`

SetIndexStatusNil sets IndexStatus to an explicit JSON null when marshaled, overriding any value previously set with SetIndexStatus. Calling SetIndexStatus again clears the null override.

### GetIndexStrategy

`func (o *ApiAtlasCollectionRestoreJobResponse) GetIndexStrategy() string`

GetIndexStrategy returns the IndexStrategy field if non-nil, zero value otherwise.

### GetIndexStrategyOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetIndexStrategyOk() (*string, bool)`

GetIndexStrategyOk returns a tuple with the IndexStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStrategy

`func (o *ApiAtlasCollectionRestoreJobResponse) SetIndexStrategy(v string)`

SetIndexStrategy sets IndexStrategy field to given value.

### HasIndexStrategy

`func (o *ApiAtlasCollectionRestoreJobResponse) HasIndexStrategy() bool`

HasIndexStrategy returns a boolean if a field has been set.

### SetIndexStrategyNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetIndexStrategyNil()`

SetIndexStrategyNil sets IndexStrategy to an explicit JSON null when marshaled, overriding any value previously set with SetIndexStrategy. Calling SetIndexStrategy again clears the null override.

### GetOplogInc

`func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogInc() int`

GetOplogInc returns the OplogInc field if non-nil, zero value otherwise.

### GetOplogIncOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogIncOk() (*int, bool)`

GetOplogIncOk returns a tuple with the OplogInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogInc

`func (o *ApiAtlasCollectionRestoreJobResponse) SetOplogInc(v int)`

SetOplogInc sets OplogInc field to given value.

### HasOplogInc

`func (o *ApiAtlasCollectionRestoreJobResponse) HasOplogInc() bool`

HasOplogInc returns a boolean if a field has been set.

### SetOplogIncNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetOplogIncNil()`

SetOplogIncNil sets OplogInc to an explicit JSON null when marshaled, overriding any value previously set with SetOplogInc. Calling SetOplogInc again clears the null override.

### GetOplogTs

`func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogTs() int`

GetOplogTs returns the OplogTs field if non-nil, zero value otherwise.

### GetOplogTsOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogTsOk() (*int, bool)`

GetOplogTsOk returns a tuple with the OplogTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogTs

`func (o *ApiAtlasCollectionRestoreJobResponse) SetOplogTs(v int)`

SetOplogTs sets OplogTs field to given value.

### HasOplogTs

`func (o *ApiAtlasCollectionRestoreJobResponse) HasOplogTs() bool`

HasOplogTs returns a boolean if a field has been set.

### SetOplogTsNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetOplogTsNil()`

SetOplogTsNil sets OplogTs to an explicit JSON null when marshaled, overriding any value previously set with SetOplogTs. Calling SetOplogTs again clears the null override.

### GetPointInTimeUtcSeconds

`func (o *ApiAtlasCollectionRestoreJobResponse) GetPointInTimeUtcSeconds() int`

GetPointInTimeUtcSeconds returns the PointInTimeUtcSeconds field if non-nil, zero value otherwise.

### GetPointInTimeUtcSecondsOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetPointInTimeUtcSecondsOk() (*int, bool)`

GetPointInTimeUtcSecondsOk returns a tuple with the PointInTimeUtcSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUtcSeconds

`func (o *ApiAtlasCollectionRestoreJobResponse) SetPointInTimeUtcSeconds(v int)`

SetPointInTimeUtcSeconds sets PointInTimeUtcSeconds field to given value.

### HasPointInTimeUtcSeconds

`func (o *ApiAtlasCollectionRestoreJobResponse) HasPointInTimeUtcSeconds() bool`

HasPointInTimeUtcSeconds returns a boolean if a field has been set.

### SetPointInTimeUtcSecondsNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetPointInTimeUtcSecondsNil()`

SetPointInTimeUtcSecondsNil sets PointInTimeUtcSeconds to an explicit JSON null when marshaled, overriding any value previously set with SetPointInTimeUtcSeconds. Calling SetPointInTimeUtcSeconds again clears the null override.

### GetRestoredDocuments

`func (o *ApiAtlasCollectionRestoreJobResponse) GetRestoredDocuments() int64`

GetRestoredDocuments returns the RestoredDocuments field if non-nil, zero value otherwise.

### GetRestoredDocumentsOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetRestoredDocumentsOk() (*int64, bool)`

GetRestoredDocumentsOk returns a tuple with the RestoredDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredDocuments

`func (o *ApiAtlasCollectionRestoreJobResponse) SetRestoredDocuments(v int64)`

SetRestoredDocuments sets RestoredDocuments field to given value.

### HasRestoredDocuments

`func (o *ApiAtlasCollectionRestoreJobResponse) HasRestoredDocuments() bool`

HasRestoredDocuments returns a boolean if a field has been set.

### SetRestoredDocumentsNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetRestoredDocumentsNil()`

SetRestoredDocumentsNil sets RestoredDocuments to an explicit JSON null when marshaled, overriding any value previously set with SetRestoredDocuments. Calling SetRestoredDocuments again clears the null override.

### GetSnapshotId

`func (o *ApiAtlasCollectionRestoreJobResponse) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ApiAtlasCollectionRestoreJobResponse) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ApiAtlasCollectionRestoreJobResponse) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetSnapshotIdNil()`

SetSnapshotIdNil sets SnapshotId to an explicit JSON null when marshaled, overriding any value previously set with SetSnapshotId. Calling SetSnapshotId again clears the null override.

### GetState

`func (o *ApiAtlasCollectionRestoreJobResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiAtlasCollectionRestoreJobResponse) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ApiAtlasCollectionRestoreJobResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetStateNil()`

SetStateNil sets State to an explicit JSON null when marshaled, overriding any value previously set with SetState. Calling SetState again clears the null override.

### GetTargetClusterName

`func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetClusterName() string`

GetTargetClusterName returns the TargetClusterName field if non-nil, zero value otherwise.

### GetTargetClusterNameOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetClusterNameOk() (*string, bool)`

GetTargetClusterNameOk returns a tuple with the TargetClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetClusterName

`func (o *ApiAtlasCollectionRestoreJobResponse) SetTargetClusterName(v string)`

SetTargetClusterName sets TargetClusterName field to given value.

### HasTargetClusterName

`func (o *ApiAtlasCollectionRestoreJobResponse) HasTargetClusterName() bool`

HasTargetClusterName returns a boolean if a field has been set.

### SetTargetClusterNameNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetTargetClusterNameNil()`

SetTargetClusterNameNil sets TargetClusterName to an explicit JSON null when marshaled, overriding any value previously set with SetTargetClusterName. Calling SetTargetClusterName again clears the null override.

### GetTargetGroupId

`func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *ApiAtlasCollectionRestoreJobResponse) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *ApiAtlasCollectionRestoreJobResponse) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### SetTargetGroupIdNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetTargetGroupIdNil()`

SetTargetGroupIdNil sets TargetGroupId to an explicit JSON null when marshaled, overriding any value previously set with SetTargetGroupId. Calling SetTargetGroupId again clears the null override.

### GetTotalDocuments

`func (o *ApiAtlasCollectionRestoreJobResponse) GetTotalDocuments() int64`

GetTotalDocuments returns the TotalDocuments field if non-nil, zero value otherwise.

### GetTotalDocumentsOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetTotalDocumentsOk() (*int64, bool)`

GetTotalDocumentsOk returns a tuple with the TotalDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDocuments

`func (o *ApiAtlasCollectionRestoreJobResponse) SetTotalDocuments(v int64)`

SetTotalDocuments sets TotalDocuments field to given value.

### HasTotalDocuments

`func (o *ApiAtlasCollectionRestoreJobResponse) HasTotalDocuments() bool`

HasTotalDocuments returns a boolean if a field has been set.

### SetTotalDocumentsNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetTotalDocumentsNil()`

SetTotalDocumentsNil sets TotalDocuments to an explicit JSON null when marshaled, overriding any value previously set with SetTotalDocuments. Calling SetTotalDocuments again clears the null override.

### GetWriteStrategy

`func (o *ApiAtlasCollectionRestoreJobResponse) GetWriteStrategy() string`

GetWriteStrategy returns the WriteStrategy field if non-nil, zero value otherwise.

### GetWriteStrategyOk

`func (o *ApiAtlasCollectionRestoreJobResponse) GetWriteStrategyOk() (*string, bool)`

GetWriteStrategyOk returns a tuple with the WriteStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteStrategy

`func (o *ApiAtlasCollectionRestoreJobResponse) SetWriteStrategy(v string)`

SetWriteStrategy sets WriteStrategy field to given value.

### HasWriteStrategy

`func (o *ApiAtlasCollectionRestoreJobResponse) HasWriteStrategy() bool`

HasWriteStrategy returns a boolean if a field has been set.

### SetWriteStrategyNil

`func (o *ApiAtlasCollectionRestoreJobResponse) SetWriteStrategyNil()`

SetWriteStrategyNil sets WriteStrategy to an explicit JSON null when marshaled, overriding any value previously set with SetWriteStrategy. Calling SetWriteStrategy again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


