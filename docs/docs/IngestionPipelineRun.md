# IngestionPipelineRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline run. | [optional] [readonly] 
**BackupFrequencyType** | Pointer to **string** | Backup schedule interval of the Data Lake Pipeline. | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** | Timestamp that indicates when the pipeline run was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**DatasetName** | Pointer to **string** | Human-readable label that identifies the dataset that Atlas generates during this pipeline run. You can use this dataset as a &#x60;dataSource&#x60; in a Federated Database collection. | [optional] [readonly] 
**DatasetRetentionPolicy** | Pointer to [**DatasetRetentionPolicy**](DatasetRetentionPolicy.md) |  | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**LastUpdatedDate** | Pointer to **time.Time** | Timestamp that indicates the last time that the pipeline run was updated. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Phase** | Pointer to **string** | Processing phase of the Data Lake Pipeline. | [optional] [readonly] 
**PipelineId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies a Data Lake Pipeline. | [optional] [readonly] 
**ScheduledDeletionDate** | Pointer to **time.Time** | Timestamp that indicates when the pipeline run will expire and its dataset will be deleted. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the snapshot of a cluster. | [optional] [readonly] 
**State** | Pointer to **string** | State of the pipeline run. | [optional] [readonly] 
**Stats** | Pointer to [**PipelineRunStats**](PipelineRunStats.md) |  | [optional] 

## Methods

### NewIngestionPipelineRun

`func NewIngestionPipelineRun() *IngestionPipelineRun`

NewIngestionPipelineRun instantiates a new IngestionPipelineRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestionPipelineRunWithDefaults

`func NewIngestionPipelineRunWithDefaults() *IngestionPipelineRun`

NewIngestionPipelineRunWithDefaults instantiates a new IngestionPipelineRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestionPipelineRun) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestionPipelineRun) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestionPipelineRun) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IngestionPipelineRun) HasId() bool`

HasId returns a boolean if a field has been set.
### GetBackupFrequencyType

`func (o *IngestionPipelineRun) GetBackupFrequencyType() string`

GetBackupFrequencyType returns the BackupFrequencyType field if non-nil, zero value otherwise.

### GetBackupFrequencyTypeOk

`func (o *IngestionPipelineRun) GetBackupFrequencyTypeOk() (*string, bool)`

GetBackupFrequencyTypeOk returns a tuple with the BackupFrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFrequencyType

`func (o *IngestionPipelineRun) SetBackupFrequencyType(v string)`

SetBackupFrequencyType sets BackupFrequencyType field to given value.

### HasBackupFrequencyType

`func (o *IngestionPipelineRun) HasBackupFrequencyType() bool`

HasBackupFrequencyType returns a boolean if a field has been set.
### GetCreatedDate

`func (o *IngestionPipelineRun) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *IngestionPipelineRun) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *IngestionPipelineRun) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *IngestionPipelineRun) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.
### GetDatasetName

`func (o *IngestionPipelineRun) GetDatasetName() string`

GetDatasetName returns the DatasetName field if non-nil, zero value otherwise.

### GetDatasetNameOk

`func (o *IngestionPipelineRun) GetDatasetNameOk() (*string, bool)`

GetDatasetNameOk returns a tuple with the DatasetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetName

`func (o *IngestionPipelineRun) SetDatasetName(v string)`

SetDatasetName sets DatasetName field to given value.

### HasDatasetName

`func (o *IngestionPipelineRun) HasDatasetName() bool`

HasDatasetName returns a boolean if a field has been set.
### GetDatasetRetentionPolicy

`func (o *IngestionPipelineRun) GetDatasetRetentionPolicy() DatasetRetentionPolicy`

GetDatasetRetentionPolicy returns the DatasetRetentionPolicy field if non-nil, zero value otherwise.

### GetDatasetRetentionPolicyOk

`func (o *IngestionPipelineRun) GetDatasetRetentionPolicyOk() (*DatasetRetentionPolicy, bool)`

GetDatasetRetentionPolicyOk returns a tuple with the DatasetRetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetRetentionPolicy

`func (o *IngestionPipelineRun) SetDatasetRetentionPolicy(v DatasetRetentionPolicy)`

SetDatasetRetentionPolicy sets DatasetRetentionPolicy field to given value.

### HasDatasetRetentionPolicy

`func (o *IngestionPipelineRun) HasDatasetRetentionPolicy() bool`

HasDatasetRetentionPolicy returns a boolean if a field has been set.
### GetGroupId

`func (o *IngestionPipelineRun) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *IngestionPipelineRun) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *IngestionPipelineRun) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *IngestionPipelineRun) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetLastUpdatedDate

`func (o *IngestionPipelineRun) GetLastUpdatedDate() time.Time`

GetLastUpdatedDate returns the LastUpdatedDate field if non-nil, zero value otherwise.

### GetLastUpdatedDateOk

`func (o *IngestionPipelineRun) GetLastUpdatedDateOk() (*time.Time, bool)`

GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedDate

`func (o *IngestionPipelineRun) SetLastUpdatedDate(v time.Time)`

SetLastUpdatedDate sets LastUpdatedDate field to given value.

### HasLastUpdatedDate

`func (o *IngestionPipelineRun) HasLastUpdatedDate() bool`

HasLastUpdatedDate returns a boolean if a field has been set.
### GetPhase

`func (o *IngestionPipelineRun) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *IngestionPipelineRun) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *IngestionPipelineRun) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *IngestionPipelineRun) HasPhase() bool`

HasPhase returns a boolean if a field has been set.
### GetPipelineId

`func (o *IngestionPipelineRun) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *IngestionPipelineRun) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *IngestionPipelineRun) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *IngestionPipelineRun) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.
### GetScheduledDeletionDate

`func (o *IngestionPipelineRun) GetScheduledDeletionDate() time.Time`

GetScheduledDeletionDate returns the ScheduledDeletionDate field if non-nil, zero value otherwise.

### GetScheduledDeletionDateOk

`func (o *IngestionPipelineRun) GetScheduledDeletionDateOk() (*time.Time, bool)`

GetScheduledDeletionDateOk returns a tuple with the ScheduledDeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeletionDate

`func (o *IngestionPipelineRun) SetScheduledDeletionDate(v time.Time)`

SetScheduledDeletionDate sets ScheduledDeletionDate field to given value.

### HasScheduledDeletionDate

`func (o *IngestionPipelineRun) HasScheduledDeletionDate() bool`

HasScheduledDeletionDate returns a boolean if a field has been set.
### GetSnapshotId

`func (o *IngestionPipelineRun) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *IngestionPipelineRun) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *IngestionPipelineRun) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *IngestionPipelineRun) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.
### GetState

`func (o *IngestionPipelineRun) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IngestionPipelineRun) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IngestionPipelineRun) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IngestionPipelineRun) HasState() bool`

HasState returns a boolean if a field has been set.
### GetStats

`func (o *IngestionPipelineRun) GetStats() PipelineRunStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *IngestionPipelineRun) GetStatsOk() (*PipelineRunStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *IngestionPipelineRun) SetStats(v PipelineRunStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *IngestionPipelineRun) HasStats() bool`

HasStats returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


