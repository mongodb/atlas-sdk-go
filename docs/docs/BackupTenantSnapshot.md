# BackupTenantSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to **time.Time** | Date and time when the download link no longer works. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FinishTime** | Pointer to **time.Time** | Date and time when MongoDB Cloud completed writing this snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the restore job. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongoDBVersion** | Pointer to **string** | MongoDB host version that the snapshot runs. | [optional] [readonly] 
**ScheduledTime** | Pointer to **time.Time** | Date and time when MongoDB Cloud will take the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**StartTime** | Pointer to **time.Time** | Date and time when MongoDB Cloud began taking the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**Status** | Pointer to **string** | Phase of the workflow for this snapshot at the time this resource made this request. | [optional] [readonly] 

## Methods

### NewBackupTenantSnapshot

`func NewBackupTenantSnapshot() *BackupTenantSnapshot`

NewBackupTenantSnapshot instantiates a new BackupTenantSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupTenantSnapshotWithDefaults

`func NewBackupTenantSnapshotWithDefaults() *BackupTenantSnapshot`

NewBackupTenantSnapshotWithDefaults instantiates a new BackupTenantSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *BackupTenantSnapshot) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BackupTenantSnapshot) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BackupTenantSnapshot) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BackupTenantSnapshot) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.
### GetFinishTime

`func (o *BackupTenantSnapshot) GetFinishTime() time.Time`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *BackupTenantSnapshot) GetFinishTimeOk() (*time.Time, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *BackupTenantSnapshot) SetFinishTime(v time.Time)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *BackupTenantSnapshot) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.
### GetId

`func (o *BackupTenantSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupTenantSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupTenantSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupTenantSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.
### GetLinks

`func (o *BackupTenantSnapshot) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupTenantSnapshot) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupTenantSnapshot) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BackupTenantSnapshot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetMongoDBVersion

`func (o *BackupTenantSnapshot) GetMongoDBVersion() string`

GetMongoDBVersion returns the MongoDBVersion field if non-nil, zero value otherwise.

### GetMongoDBVersionOk

`func (o *BackupTenantSnapshot) GetMongoDBVersionOk() (*string, bool)`

GetMongoDBVersionOk returns a tuple with the MongoDBVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoDBVersion

`func (o *BackupTenantSnapshot) SetMongoDBVersion(v string)`

SetMongoDBVersion sets MongoDBVersion field to given value.

### HasMongoDBVersion

`func (o *BackupTenantSnapshot) HasMongoDBVersion() bool`

HasMongoDBVersion returns a boolean if a field has been set.
### GetScheduledTime

`func (o *BackupTenantSnapshot) GetScheduledTime() time.Time`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *BackupTenantSnapshot) GetScheduledTimeOk() (*time.Time, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *BackupTenantSnapshot) SetScheduledTime(v time.Time)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *BackupTenantSnapshot) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.
### GetStartTime

`func (o *BackupTenantSnapshot) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupTenantSnapshot) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupTenantSnapshot) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupTenantSnapshot) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.
### GetStatus

`func (o *BackupTenantSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackupTenantSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackupTenantSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BackupTenantSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


