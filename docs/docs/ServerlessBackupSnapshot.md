# ServerlessBackupSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud took the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC. | [optional] [readonly] 
**FrequencyType** | Pointer to **string** | Human-readable label that identifies how often this snapshot triggers. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the snapshot. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**MongodVersion** | Pointer to **string** | Version of the MongoDB host that this snapshot backs up. | [optional] [readonly] 
**ServerlessInstanceName** | Pointer to **string** | Human-readable label given to the serverless instance from which MongoDB Cloud took this snapshot. | [optional] [readonly] 
**SnapshotType** | Pointer to **string** | Human-readable label that identifies when this snapshot triggers. | [optional] [readonly] 
**Status** | Pointer to **string** | Human-readable label that indicates the stage of the backup process for this snapshot. | [optional] [readonly] 
**StorageSizeBytes** | Pointer to **int64** | Number of bytes taken to store the backup snapshot. | [optional] [readonly] 

## Methods

### NewServerlessBackupSnapshot

`func NewServerlessBackupSnapshot() *ServerlessBackupSnapshot`

NewServerlessBackupSnapshot instantiates a new ServerlessBackupSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessBackupSnapshotWithDefaults

`func NewServerlessBackupSnapshotWithDefaults() *ServerlessBackupSnapshot`

NewServerlessBackupSnapshotWithDefaults instantiates a new ServerlessBackupSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ServerlessBackupSnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ServerlessBackupSnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ServerlessBackupSnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ServerlessBackupSnapshot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ServerlessBackupSnapshot) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ServerlessBackupSnapshot) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ServerlessBackupSnapshot) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ServerlessBackupSnapshot) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFrequencyType

`func (o *ServerlessBackupSnapshot) GetFrequencyType() string`

GetFrequencyType returns the FrequencyType field if non-nil, zero value otherwise.

### GetFrequencyTypeOk

`func (o *ServerlessBackupSnapshot) GetFrequencyTypeOk() (*string, bool)`

GetFrequencyTypeOk returns a tuple with the FrequencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyType

`func (o *ServerlessBackupSnapshot) SetFrequencyType(v string)`

SetFrequencyType sets FrequencyType field to given value.

### HasFrequencyType

`func (o *ServerlessBackupSnapshot) HasFrequencyType() bool`

HasFrequencyType returns a boolean if a field has been set.

### GetId

`func (o *ServerlessBackupSnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerlessBackupSnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerlessBackupSnapshot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerlessBackupSnapshot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ServerlessBackupSnapshot) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServerlessBackupSnapshot) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServerlessBackupSnapshot) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServerlessBackupSnapshot) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMongodVersion

`func (o *ServerlessBackupSnapshot) GetMongodVersion() string`

GetMongodVersion returns the MongodVersion field if non-nil, zero value otherwise.

### GetMongodVersionOk

`func (o *ServerlessBackupSnapshot) GetMongodVersionOk() (*string, bool)`

GetMongodVersionOk returns a tuple with the MongodVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodVersion

`func (o *ServerlessBackupSnapshot) SetMongodVersion(v string)`

SetMongodVersion sets MongodVersion field to given value.

### HasMongodVersion

`func (o *ServerlessBackupSnapshot) HasMongodVersion() bool`

HasMongodVersion returns a boolean if a field has been set.

### GetServerlessInstanceName

`func (o *ServerlessBackupSnapshot) GetServerlessInstanceName() string`

GetServerlessInstanceName returns the ServerlessInstanceName field if non-nil, zero value otherwise.

### GetServerlessInstanceNameOk

`func (o *ServerlessBackupSnapshot) GetServerlessInstanceNameOk() (*string, bool)`

GetServerlessInstanceNameOk returns a tuple with the ServerlessInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerlessInstanceName

`func (o *ServerlessBackupSnapshot) SetServerlessInstanceName(v string)`

SetServerlessInstanceName sets ServerlessInstanceName field to given value.

### HasServerlessInstanceName

`func (o *ServerlessBackupSnapshot) HasServerlessInstanceName() bool`

HasServerlessInstanceName returns a boolean if a field has been set.

### GetSnapshotType

`func (o *ServerlessBackupSnapshot) GetSnapshotType() string`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *ServerlessBackupSnapshot) GetSnapshotTypeOk() (*string, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *ServerlessBackupSnapshot) SetSnapshotType(v string)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *ServerlessBackupSnapshot) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetStatus

`func (o *ServerlessBackupSnapshot) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServerlessBackupSnapshot) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServerlessBackupSnapshot) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServerlessBackupSnapshot) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageSizeBytes

`func (o *ServerlessBackupSnapshot) GetStorageSizeBytes() int64`

GetStorageSizeBytes returns the StorageSizeBytes field if non-nil, zero value otherwise.

### GetStorageSizeBytesOk

`func (o *ServerlessBackupSnapshot) GetStorageSizeBytesOk() (*int64, bool)`

GetStorageSizeBytesOk returns a tuple with the StorageSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeBytes

`func (o *ServerlessBackupSnapshot) SetStorageSizeBytes(v int64)`

SetStorageSizeBytes sets StorageSizeBytes field to given value.

### HasStorageSizeBytes

`func (o *ServerlessBackupSnapshot) HasStorageSizeBytes() bool`

HasStorageSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


