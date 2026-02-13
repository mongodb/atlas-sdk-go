# BackupOnlineArchive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the online archive. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster that contains the collection for which you want to create an online archive. | [optional] [readonly] 
**CollName** | Pointer to **string** | Human-readable label that identifies the collection for which you created the online archive. | [optional] [readonly] 
**CollectionType** | Pointer to **string** | Classification of MongoDB database collection that you want to return.  If you set this parameter to &#x60;TIMESERIES&#x60;, set &#x60;\&quot;criteria.type\&quot; : \&quot;date\&quot;&#x60; and &#x60;\&quot;criteria.dateFormat\&quot; : \&quot;ISODATE\&quot;&#x60;. | [optional] [readonly] [default to "STANDARD"]
**Criteria** | Pointer to [**Criteria**](Criteria.md) |  | [optional] 
**DataExpirationRule** | Pointer to [**DataExpirationRule**](DataExpirationRule.md) |  | [optional] 
**DataProcessRegion** | Pointer to [**DataProcessRegion**](DataProcessRegion.md) |  | [optional] 
**DataSetName** | Pointer to **string** | Human-readable label that identifies the dataset that Atlas generates for this online archive. | [optional] [readonly] 
**DbName** | Pointer to **string** | Human-readable label of the database that contains the collection that contains the online archive. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that contains the specified cluster. The specified cluster contains the collection for which to create the online archive. | [optional] [readonly] 
**PartitionFields** | Pointer to [**[]PartitionField**](PartitionField.md) | List that contains document parameters to use to logically divide data within a collection. Partitions provide a coarse level of filtering of the underlying collection data. To divide your data, specify parameters that you frequently query. If you specified &#x60;criteria.type&#x60;: &#x60;DATE&#x60; in the Create One Online Archive endpoint, then you can specify up to three parameters by which to query. One of these parameters must be the &#x60;DATE&#x60; value, which is required in this case. If you specified &#x60;criteria.type&#x60;: &#x60;CUSTOM&#x60; in the Create One Online Archive endpoint, then you can specify up to two parameters by which to query. Queries that don&#39;t use &#x60;criteria.type&#x60;: &#x60;DATE&#x60; or &#x60;criteria.type&#x60;: &#x60;CUSTOM&#x60; parameters cause MongoDB to scan a full collection of all archived documents. This takes more time and increases your costs. | [optional] [readonly] 
**Paused** | Pointer to **bool** | Flag that indicates whether this online archive exists in the paused state. A request to resume fails if the collection has another active online archive. To pause an active online archive or resume a paused online archive, you must include this parameter. To pause an active archive, set this to **true**. To resume a paused archive, set this to **false**. | [optional] 
**Schedule** | Pointer to [**OnlineArchiveSchedule**](OnlineArchiveSchedule.md) |  | [optional] 
**State** | Pointer to **string** | Phase of the process to create this online archive when you made this request.  | State       | Indication | |-------------|------------| | &#x60;PENDING&#x60;   | MongoDB Cloud has queued documents for archive. Archiving hasn&#39;t started. | | &#x60;ARCHIVING&#x60; | MongoDB Cloud started archiving documents that meet the archival criteria. | | &#x60;IDLE&#x60;      | MongoDB Cloud waits to start the next archival job. | | &#x60;PAUSING&#x60;   | Someone chose to stop archiving. MongoDB Cloud finishes the running archival job then changes the state to &#x60;PAUSED&#x60; when that job completes. | | &#x60;PAUSED&#x60;    | MongoDB Cloud has stopped archiving. Archived documents can be queried. The specified archiving operation on the active cluster cannot archive additional documents. You can resume archiving for paused archives at any time. | | &#x60;ORPHANED&#x60;  | Someone has deleted the collection associated with an active or paused archive. MongoDB Cloud doesn&#39;t delete the archived data. You must manually delete the online archives associated with the deleted collection. | | &#x60;DELETED&#x60;   | Someone has deleted the archive was deleted. When someone deletes an online archive, MongoDB Cloud removes all associated archived documents from the cloud object storage. | | [optional] [readonly] 

## Methods

### NewBackupOnlineArchive

`func NewBackupOnlineArchive() *BackupOnlineArchive`

NewBackupOnlineArchive instantiates a new BackupOnlineArchive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupOnlineArchiveWithDefaults

`func NewBackupOnlineArchiveWithDefaults() *BackupOnlineArchive`

NewBackupOnlineArchiveWithDefaults instantiates a new BackupOnlineArchive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupOnlineArchive) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupOnlineArchive) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupOnlineArchive) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupOnlineArchive) HasId() bool`

HasId returns a boolean if a field has been set.
### GetClusterName

`func (o *BackupOnlineArchive) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *BackupOnlineArchive) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *BackupOnlineArchive) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *BackupOnlineArchive) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.
### GetCollName

`func (o *BackupOnlineArchive) GetCollName() string`

GetCollName returns the CollName field if non-nil, zero value otherwise.

### GetCollNameOk

`func (o *BackupOnlineArchive) GetCollNameOk() (*string, bool)`

GetCollNameOk returns a tuple with the CollName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollName

`func (o *BackupOnlineArchive) SetCollName(v string)`

SetCollName sets CollName field to given value.

### HasCollName

`func (o *BackupOnlineArchive) HasCollName() bool`

HasCollName returns a boolean if a field has been set.
### GetCollectionType

`func (o *BackupOnlineArchive) GetCollectionType() string`

GetCollectionType returns the CollectionType field if non-nil, zero value otherwise.

### GetCollectionTypeOk

`func (o *BackupOnlineArchive) GetCollectionTypeOk() (*string, bool)`

GetCollectionTypeOk returns a tuple with the CollectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionType

`func (o *BackupOnlineArchive) SetCollectionType(v string)`

SetCollectionType sets CollectionType field to given value.

### HasCollectionType

`func (o *BackupOnlineArchive) HasCollectionType() bool`

HasCollectionType returns a boolean if a field has been set.
### GetCriteria

`func (o *BackupOnlineArchive) GetCriteria() Criteria`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *BackupOnlineArchive) GetCriteriaOk() (*Criteria, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *BackupOnlineArchive) SetCriteria(v Criteria)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *BackupOnlineArchive) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.
### GetDataExpirationRule

`func (o *BackupOnlineArchive) GetDataExpirationRule() DataExpirationRule`

GetDataExpirationRule returns the DataExpirationRule field if non-nil, zero value otherwise.

### GetDataExpirationRuleOk

`func (o *BackupOnlineArchive) GetDataExpirationRuleOk() (*DataExpirationRule, bool)`

GetDataExpirationRuleOk returns a tuple with the DataExpirationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataExpirationRule

`func (o *BackupOnlineArchive) SetDataExpirationRule(v DataExpirationRule)`

SetDataExpirationRule sets DataExpirationRule field to given value.

### HasDataExpirationRule

`func (o *BackupOnlineArchive) HasDataExpirationRule() bool`

HasDataExpirationRule returns a boolean if a field has been set.
### GetDataProcessRegion

`func (o *BackupOnlineArchive) GetDataProcessRegion() DataProcessRegion`

GetDataProcessRegion returns the DataProcessRegion field if non-nil, zero value otherwise.

### GetDataProcessRegionOk

`func (o *BackupOnlineArchive) GetDataProcessRegionOk() (*DataProcessRegion, bool)`

GetDataProcessRegionOk returns a tuple with the DataProcessRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessRegion

`func (o *BackupOnlineArchive) SetDataProcessRegion(v DataProcessRegion)`

SetDataProcessRegion sets DataProcessRegion field to given value.

### HasDataProcessRegion

`func (o *BackupOnlineArchive) HasDataProcessRegion() bool`

HasDataProcessRegion returns a boolean if a field has been set.
### GetDataSetName

`func (o *BackupOnlineArchive) GetDataSetName() string`

GetDataSetName returns the DataSetName field if non-nil, zero value otherwise.

### GetDataSetNameOk

`func (o *BackupOnlineArchive) GetDataSetNameOk() (*string, bool)`

GetDataSetNameOk returns a tuple with the DataSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSetName

`func (o *BackupOnlineArchive) SetDataSetName(v string)`

SetDataSetName sets DataSetName field to given value.

### HasDataSetName

`func (o *BackupOnlineArchive) HasDataSetName() bool`

HasDataSetName returns a boolean if a field has been set.
### GetDbName

`func (o *BackupOnlineArchive) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *BackupOnlineArchive) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *BackupOnlineArchive) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *BackupOnlineArchive) HasDbName() bool`

HasDbName returns a boolean if a field has been set.
### GetGroupId

`func (o *BackupOnlineArchive) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BackupOnlineArchive) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BackupOnlineArchive) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BackupOnlineArchive) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetPartitionFields

`func (o *BackupOnlineArchive) GetPartitionFields() []PartitionField`

GetPartitionFields returns the PartitionFields field if non-nil, zero value otherwise.

### GetPartitionFieldsOk

`func (o *BackupOnlineArchive) GetPartitionFieldsOk() (*[]PartitionField, bool)`

GetPartitionFieldsOk returns a tuple with the PartitionFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionFields

`func (o *BackupOnlineArchive) SetPartitionFields(v []PartitionField)`

SetPartitionFields sets PartitionFields field to given value.

### HasPartitionFields

`func (o *BackupOnlineArchive) HasPartitionFields() bool`

HasPartitionFields returns a boolean if a field has been set.
### GetPaused

`func (o *BackupOnlineArchive) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *BackupOnlineArchive) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *BackupOnlineArchive) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *BackupOnlineArchive) HasPaused() bool`

HasPaused returns a boolean if a field has been set.
### GetSchedule

`func (o *BackupOnlineArchive) GetSchedule() OnlineArchiveSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BackupOnlineArchive) GetScheduleOk() (*OnlineArchiveSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BackupOnlineArchive) SetSchedule(v OnlineArchiveSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *BackupOnlineArchive) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.
### GetState

`func (o *BackupOnlineArchive) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupOnlineArchive) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupOnlineArchive) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *BackupOnlineArchive) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


