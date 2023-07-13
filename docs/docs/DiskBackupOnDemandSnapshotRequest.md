# DiskBackupOnDemandSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Human-readable phrase or sentence that explains the purpose of the snapshot. The resource returns this parameter when &#x60;\&quot;status\&quot; : \&quot;onDemand\&quot;&#x60;. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RetentionInDays** | Pointer to **int** | Number of days that MongoDB Cloud should retain the on-demand snapshot. Must be at least **1**. | [optional] 

## Methods

### NewDiskBackupOnDemandSnapshotRequest

`func NewDiskBackupOnDemandSnapshotRequest() *DiskBackupOnDemandSnapshotRequest`

NewDiskBackupOnDemandSnapshotRequest instantiates a new DiskBackupOnDemandSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupOnDemandSnapshotRequestWithDefaults

`func NewDiskBackupOnDemandSnapshotRequestWithDefaults() *DiskBackupOnDemandSnapshotRequest`

NewDiskBackupOnDemandSnapshotRequestWithDefaults instantiates a new DiskBackupOnDemandSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DiskBackupOnDemandSnapshotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiskBackupOnDemandSnapshotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiskBackupOnDemandSnapshotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiskBackupOnDemandSnapshotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetLinks

`func (o *DiskBackupOnDemandSnapshotRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupOnDemandSnapshotRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupOnDemandSnapshotRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupOnDemandSnapshotRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRetentionInDays

`func (o *DiskBackupOnDemandSnapshotRequest) GetRetentionInDays() int`

GetRetentionInDays returns the RetentionInDays field if non-nil, zero value otherwise.

### GetRetentionInDaysOk

`func (o *DiskBackupOnDemandSnapshotRequest) GetRetentionInDaysOk() (*int, bool)`

GetRetentionInDaysOk returns a tuple with the RetentionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionInDays

`func (o *DiskBackupOnDemandSnapshotRequest) SetRetentionInDays(v int)`

SetRetentionInDays sets RetentionInDays field to given value.

### HasRetentionInDays

`func (o *DiskBackupOnDemandSnapshotRequest) HasRetentionInDays() bool`

HasRetentionInDays returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


