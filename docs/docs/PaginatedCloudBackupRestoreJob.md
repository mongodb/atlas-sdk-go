# PaginatedCloudBackupRestoreJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]DiskBackupSnapshotRestoreJob**](DiskBackupSnapshotRestoreJob.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Number of documents returned in this response. | [optional] [readonly] 

## Methods

### NewPaginatedCloudBackupRestoreJob

`func NewPaginatedCloudBackupRestoreJob() *PaginatedCloudBackupRestoreJob`

NewPaginatedCloudBackupRestoreJob instantiates a new PaginatedCloudBackupRestoreJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCloudBackupRestoreJobWithDefaults

`func NewPaginatedCloudBackupRestoreJobWithDefaults() *PaginatedCloudBackupRestoreJob`

NewPaginatedCloudBackupRestoreJobWithDefaults instantiates a new PaginatedCloudBackupRestoreJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedCloudBackupRestoreJob) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedCloudBackupRestoreJob) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedCloudBackupRestoreJob) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedCloudBackupRestoreJob) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResults

`func (o *PaginatedCloudBackupRestoreJob) GetResults() []DiskBackupSnapshotRestoreJob`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCloudBackupRestoreJob) GetResultsOk() (*[]DiskBackupSnapshotRestoreJob, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCloudBackupRestoreJob) SetResults(v []DiskBackupSnapshotRestoreJob)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedCloudBackupRestoreJob) HasResults() bool`

HasResults returns a boolean if a field has been set.
### GetTotalCount

`func (o *PaginatedCloudBackupRestoreJob) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedCloudBackupRestoreJob) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedCloudBackupRestoreJob) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedCloudBackupRestoreJob) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


