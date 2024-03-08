# PaginatedCloudBackupReplicaSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Results** | Pointer to [**[]DiskBackupReplicaSet**](DiskBackupReplicaSet.md) | List of returned documents that MongoDB Cloud providers when completing this request. | [optional] [readonly] 
**TotalCount** | Pointer to **int** | Total number of documents available. MongoDB Cloud omits this value if &#x60;includeCount&#x60; is set to &#x60;false&#x60;. | [optional] [readonly] 

## Methods

### NewPaginatedCloudBackupReplicaSet

`func NewPaginatedCloudBackupReplicaSet() *PaginatedCloudBackupReplicaSet`

NewPaginatedCloudBackupReplicaSet instantiates a new PaginatedCloudBackupReplicaSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedCloudBackupReplicaSetWithDefaults

`func NewPaginatedCloudBackupReplicaSetWithDefaults() *PaginatedCloudBackupReplicaSet`

NewPaginatedCloudBackupReplicaSetWithDefaults instantiates a new PaginatedCloudBackupReplicaSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PaginatedCloudBackupReplicaSet) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PaginatedCloudBackupReplicaSet) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PaginatedCloudBackupReplicaSet) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PaginatedCloudBackupReplicaSet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetResults

`func (o *PaginatedCloudBackupReplicaSet) GetResults() []DiskBackupReplicaSet`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedCloudBackupReplicaSet) GetResultsOk() (*[]DiskBackupReplicaSet, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedCloudBackupReplicaSet) SetResults(v []DiskBackupReplicaSet)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedCloudBackupReplicaSet) HasResults() bool`

HasResults returns a boolean if a field has been set.
### GetTotalCount

`func (o *PaginatedCloudBackupReplicaSet) GetTotalCount() int`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PaginatedCloudBackupReplicaSet) GetTotalCountOk() (*int, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PaginatedCloudBackupReplicaSet) SetTotalCount(v int)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *PaginatedCloudBackupReplicaSet) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


