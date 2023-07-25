# BackupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the cluster with the snapshots you want to return. | [optional] [readonly] 
**Complete** | Pointer to **bool** | Flag that indicates whether the snapshot exists. This flag returns &#x60;false&#x60; while MongoDB Cloud creates the snapshot. | [optional] [readonly] 
**Created** | Pointer to [**ApiBSONTimestamp**](ApiBSONTimestamp.md) |  | [optional] 
**DoNotDelete** | Pointer to **bool** | Flag that indicates whether someone can delete this snapshot. You can&#39;t set &#x60;\&quot;doNotDelete\&quot; : true&#x60; and set a timestamp for **expires** in the same request. | [optional] 
**Expires** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. If &#x60;\&quot;doNotDelete\&quot; : true&#x60;, MongoDB Cloud removes any value set for this parameter. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the project that owns the snapshots. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**Incremental** | Pointer to **bool** | Flag indicating if this is an incremental or a full snapshot. | [optional] [readonly] 
**LastOplogAppliedTimestamp** | Pointer to [**ApiBSONTimestamp**](ApiBSONTimestamp.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Parts** | Pointer to [**[]BackupSnapshotPart**](BackupSnapshotPart.md) | Metadata that describes the complete snapshot.  - For a replica set, this array contains a single document. - For a sharded cluster, this array contains one document for each shard plus one document for the config host. | [optional] [readonly] 

## Methods

### NewBackupSnapshot

`func NewBackupSnapshot() *BackupSnapshot`

NewBackupSnapshot instantiates a new BackupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupSnapshotWithDefaults

`func NewBackupSnapshotWithDefaults() *BackupSnapshot`

NewBackupSnapshotWithDefaults instantiates a new BackupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *BackupSnapshot) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *BackupSnapshot) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *BackupSnapshot) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *BackupSnapshot) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.
### GetComplete

`func (o *BackupSnapshot) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *BackupSnapshot) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *BackupSnapshot) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *BackupSnapshot) HasComplete() bool`

HasComplete returns a boolean if a field has been set.
### GetCreated

`func (o *BackupSnapshot) GetCreated() ApiBSONTimestamp`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BackupSnapshot) GetCreatedOk() (*ApiBSONTimestamp, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BackupSnapshot) SetCreated(v ApiBSONTimestamp)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BackupSnapshot) HasCreated() bool`

HasCreated returns a boolean if a field has been set.
### GetDoNotDelete

`func (o *BackupSnapshot) GetDoNotDelete() bool`

GetDoNotDelete returns the DoNotDelete field if non-nil, zero value otherwise.

### GetDoNotDeleteOk

`func (o *BackupSnapshot) GetDoNotDeleteOk() (*bool, bool)`

GetDoNotDeleteOk returns a tuple with the DoNotDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotDelete

`func (o *BackupSnapshot) SetDoNotDelete(v bool)`

SetDoNotDelete sets DoNotDelete field to given value.

### HasDoNotDelete

`func (o *BackupSnapshot) HasDoNotDelete() bool`

HasDoNotDelete returns a boolean if a field has been set.
### GetExpires

`func (o *BackupSnapshot) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *BackupSnapshot) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *BackupSnapshot) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *BackupSnapshot) HasExpires() bool`

HasExpires returns a boolean if a field has been set.
### GetGroupId

`func (o *BackupSnapshot) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BackupSnapshot) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BackupSnapshot) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BackupSnapshot) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetId

`func (o *BackupSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.
### GetIncremental

`func (o *BackupSnapshot) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *BackupSnapshot) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *BackupSnapshot) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *BackupSnapshot) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.
### GetLastOplogAppliedTimestamp

`func (o *BackupSnapshot) GetLastOplogAppliedTimestamp() ApiBSONTimestamp`

GetLastOplogAppliedTimestamp returns the LastOplogAppliedTimestamp field if non-nil, zero value otherwise.

### GetLastOplogAppliedTimestampOk

`func (o *BackupSnapshot) GetLastOplogAppliedTimestampOk() (*ApiBSONTimestamp, bool)`

GetLastOplogAppliedTimestampOk returns a tuple with the LastOplogAppliedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOplogAppliedTimestamp

`func (o *BackupSnapshot) SetLastOplogAppliedTimestamp(v ApiBSONTimestamp)`

SetLastOplogAppliedTimestamp sets LastOplogAppliedTimestamp field to given value.

### HasLastOplogAppliedTimestamp

`func (o *BackupSnapshot) HasLastOplogAppliedTimestamp() bool`

HasLastOplogAppliedTimestamp returns a boolean if a field has been set.
### GetLinks

`func (o *BackupSnapshot) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupSnapshot) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupSnapshot) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BackupSnapshot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetParts

`func (o *BackupSnapshot) GetParts() []BackupSnapshotPart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *BackupSnapshot) GetPartsOk() (*[]BackupSnapshotPart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *BackupSnapshot) SetParts(v []BackupSnapshotPart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *BackupSnapshot) HasParts() bool`

HasParts returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


