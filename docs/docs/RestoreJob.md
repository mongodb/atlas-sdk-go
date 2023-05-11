# RestoreJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the batch to which this restore job belongs. This parameter exists only for a sharded cluster restore. | [optional] [readonly] 
**CheckpointId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the sharded cluster checkpoint. The checkpoint represents the point in time back to which you want to restore you data. This parameter applies when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;. Use this parameter with sharded clusters only.  - If you set **checkpointId**, you can&#39;t set **oplogInc**, **oplogTs**, **snapshotId**, or **pointInTimeUTCMillis**. - If you provide this parameter, this endpoint restores all data up to this checkpoint to the database you specify in the **delivery** object. | [optional] 
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshot you want to return. This parameter returns for restore clusters. | [optional] [readonly] 
**ClusterName** | Pointer to **string** | Human-readable label that identifies the cluster containing the snapshots you want to retrieve. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | Date and time when someone requested this restore job. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Delivery** | [**RestoreJobDelivery**](RestoreJobDelivery.md) |  | 
**EncryptionEnabled** | Pointer to **bool** | Flag that indicates whether someone encrypted the data in the restored snapshot. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns the snapshots. | [optional] [readonly] 
**Hashes** | Pointer to [**[]RestoreJobFileHash**](RestoreJobFileHash.md) | List that contains documents mapping each restore file to a hashed checksum. This parameter applies after you download the corresponding **delivery.url**. If &#x60;\&quot;methodName\&quot; : \&quot;HTTP\&quot;&#x60;, this list contains one object that represents the hash of the **.tar.gz** file. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the restore job. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MasterKeyUUID** | Pointer to **string** | Universally Unique Identifier (UUID) that identifies the Key Management Interoperability (KMIP) master key used to encrypt the snapshot data. This parameter applies only when &#x60;\&quot;encryptionEnabled\&quot; : \&quot;true\&quot;&#x60;. | [optional] [readonly] 
**OplogInc** | Pointer to **int** | Thirty-two-bit incrementing ordinal that represents operations within a given second. When paired with **oplogTs**, this represents the point in time to which MongoDB Cloud restores your data. This parameter applies when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;.  - If you set **oplogInc**, you must set **oplogTs**, and can&#39;t set **checkpointId**, **snapshotId**, or **pointInTimeUTCMillis**. - If you provide this parameter, this endpoint restores all data up to and including this Oplog timestamp to the database you specified in the **delivery** object. | [optional] 
**OplogTs** | Pointer to **string** | Date and time from which you want to restore this snapshot. This parameter expresses its value in ISO 8601 format in UTC. This represents the first part of an Oplog timestamp. When paired with **oplogInc**, they represent the last database operation to which you want to restore your data. This parameter applies when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;. Run a query against **local.oplog.rs** on your replica set to find the desired timestamp.  - If you set **oplogTs**, you must set **oplogInc**, and you can&#39;t set **checkpointId**, **snapshotId**, or **pointInTimeUTCMillis**. - If you provide this parameter, this endpoint restores all data up to and including this Oplog timestamp to the database you specified in the **delivery** object. | [optional] 
**PointInTimeUTCMillis** | Pointer to **int64** | Timestamp from which you want to restore this snapshot. This parameter expresses its value in the number of milliseconds elapsed since the [UNIX epoch](https://en.wikipedia.org/wiki/Unix_time). This timestamp must fall within the last 24 hours of the current time. This parameter applies when &#x60;\&quot;delivery.methodName\&quot; : \&quot;AUTOMATED_RESTORE\&quot;&#x60;.  - If you provide this parameter, this endpoint restores all data up to this point in time to the database you specified in the **delivery** object. - If you set **pointInTimeUTCMillis**, you can&#39;t set **oplogInc**, **oplogTs**, **snapshotId**, or **checkpointId**. | [optional] 
**SnapshotId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot to restore. If you set **snapshotId**, you can&#39;t set **oplogInc**, **oplogTs**, **pointInTimeUTCMillis**, or **checkpointId**. | [optional] 
**StatusName** | Pointer to **string** | Human-readable label that identifies the status of the downloadable file at the time of the request. | [optional] [readonly] 
**Timestamp** | Pointer to [**BSONTimestamp**](BSONTimestamp.md) |  | [optional] 

## Methods

### NewRestoreJob

`func NewRestoreJob(delivery RestoreJobDelivery, ) *RestoreJob`

NewRestoreJob instantiates a new RestoreJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreJobWithDefaults

`func NewRestoreJobWithDefaults() *RestoreJob`

NewRestoreJobWithDefaults instantiates a new RestoreJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *RestoreJob) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *RestoreJob) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *RestoreJob) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *RestoreJob) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetCheckpointId

`func (o *RestoreJob) GetCheckpointId() string`

GetCheckpointId returns the CheckpointId field if non-nil, zero value otherwise.

### GetCheckpointIdOk

`func (o *RestoreJob) GetCheckpointIdOk() (*string, bool)`

GetCheckpointIdOk returns a tuple with the CheckpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckpointId

`func (o *RestoreJob) SetCheckpointId(v string)`

SetCheckpointId sets CheckpointId field to given value.

### HasCheckpointId

`func (o *RestoreJob) HasCheckpointId() bool`

HasCheckpointId returns a boolean if a field has been set.

### GetClusterId

`func (o *RestoreJob) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RestoreJob) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RestoreJob) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *RestoreJob) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *RestoreJob) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RestoreJob) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RestoreJob) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RestoreJob) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCreated

`func (o *RestoreJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestoreJob) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestoreJob) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestoreJob) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDelivery

`func (o *RestoreJob) GetDelivery() RestoreJobDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *RestoreJob) GetDeliveryOk() (*RestoreJobDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *RestoreJob) SetDelivery(v RestoreJobDelivery)`

SetDelivery sets Delivery field to given value.


### GetEncryptionEnabled

`func (o *RestoreJob) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *RestoreJob) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *RestoreJob) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.

### HasEncryptionEnabled

`func (o *RestoreJob) HasEncryptionEnabled() bool`

HasEncryptionEnabled returns a boolean if a field has been set.

### GetGroupId

`func (o *RestoreJob) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *RestoreJob) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *RestoreJob) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *RestoreJob) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetHashes

`func (o *RestoreJob) GetHashes() []RestoreJobFileHash`

GetHashes returns the Hashes field if non-nil, zero value otherwise.

### GetHashesOk

`func (o *RestoreJob) GetHashesOk() (*[]RestoreJobFileHash, bool)`

GetHashesOk returns a tuple with the Hashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashes

`func (o *RestoreJob) SetHashes(v []RestoreJobFileHash)`

SetHashes sets Hashes field to given value.

### HasHashes

`func (o *RestoreJob) HasHashes() bool`

HasHashes returns a boolean if a field has been set.

### GetId

`func (o *RestoreJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RestoreJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *RestoreJob) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RestoreJob) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RestoreJob) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RestoreJob) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMasterKeyUUID

`func (o *RestoreJob) GetMasterKeyUUID() string`

GetMasterKeyUUID returns the MasterKeyUUID field if non-nil, zero value otherwise.

### GetMasterKeyUUIDOk

`func (o *RestoreJob) GetMasterKeyUUIDOk() (*string, bool)`

GetMasterKeyUUIDOk returns a tuple with the MasterKeyUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyUUID

`func (o *RestoreJob) SetMasterKeyUUID(v string)`

SetMasterKeyUUID sets MasterKeyUUID field to given value.

### HasMasterKeyUUID

`func (o *RestoreJob) HasMasterKeyUUID() bool`

HasMasterKeyUUID returns a boolean if a field has been set.

### GetOplogInc

`func (o *RestoreJob) GetOplogInc() int`

GetOplogInc returns the OplogInc field if non-nil, zero value otherwise.

### GetOplogIncOk

`func (o *RestoreJob) GetOplogIncOk() (*int, bool)`

GetOplogIncOk returns a tuple with the OplogInc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogInc

`func (o *RestoreJob) SetOplogInc(v int)`

SetOplogInc sets OplogInc field to given value.

### HasOplogInc

`func (o *RestoreJob) HasOplogInc() bool`

HasOplogInc returns a boolean if a field has been set.

### GetOplogTs

`func (o *RestoreJob) GetOplogTs() string`

GetOplogTs returns the OplogTs field if non-nil, zero value otherwise.

### GetOplogTsOk

`func (o *RestoreJob) GetOplogTsOk() (*string, bool)`

GetOplogTsOk returns a tuple with the OplogTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogTs

`func (o *RestoreJob) SetOplogTs(v string)`

SetOplogTs sets OplogTs field to given value.

### HasOplogTs

`func (o *RestoreJob) HasOplogTs() bool`

HasOplogTs returns a boolean if a field has been set.

### GetPointInTimeUTCMillis

`func (o *RestoreJob) GetPointInTimeUTCMillis() int64`

GetPointInTimeUTCMillis returns the PointInTimeUTCMillis field if non-nil, zero value otherwise.

### GetPointInTimeUTCMillisOk

`func (o *RestoreJob) GetPointInTimeUTCMillisOk() (*int64, bool)`

GetPointInTimeUTCMillisOk returns a tuple with the PointInTimeUTCMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeUTCMillis

`func (o *RestoreJob) SetPointInTimeUTCMillis(v int64)`

SetPointInTimeUTCMillis sets PointInTimeUTCMillis field to given value.

### HasPointInTimeUTCMillis

`func (o *RestoreJob) HasPointInTimeUTCMillis() bool`

HasPointInTimeUTCMillis returns a boolean if a field has been set.

### GetSnapshotId

`func (o *RestoreJob) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RestoreJob) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RestoreJob) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *RestoreJob) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetStatusName

`func (o *RestoreJob) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *RestoreJob) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *RestoreJob) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *RestoreJob) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetTimestamp

`func (o *RestoreJob) GetTimestamp() BSONTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RestoreJob) GetTimestampOk() (*BSONTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RestoreJob) SetTimestamp(v BSONTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RestoreJob) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


