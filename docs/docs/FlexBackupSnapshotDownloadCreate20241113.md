# FlexBackupSnapshotDownloadCreate20241113

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotId** | **string** | Unique 24-hexadecimal digit string that identifies the snapshot to download. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewFlexBackupSnapshotDownloadCreate20241113

`func NewFlexBackupSnapshotDownloadCreate20241113(snapshotId string, ) *FlexBackupSnapshotDownloadCreate20241113`

NewFlexBackupSnapshotDownloadCreate20241113 instantiates a new FlexBackupSnapshotDownloadCreate20241113 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlexBackupSnapshotDownloadCreate20241113WithDefaults

`func NewFlexBackupSnapshotDownloadCreate20241113WithDefaults() *FlexBackupSnapshotDownloadCreate20241113`

NewFlexBackupSnapshotDownloadCreate20241113WithDefaults instantiates a new FlexBackupSnapshotDownloadCreate20241113 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotId

`func (o *FlexBackupSnapshotDownloadCreate20241113) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *FlexBackupSnapshotDownloadCreate20241113) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *FlexBackupSnapshotDownloadCreate20241113) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### GetLinks

`func (o *FlexBackupSnapshotDownloadCreate20241113) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FlexBackupSnapshotDownloadCreate20241113) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FlexBackupSnapshotDownloadCreate20241113) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FlexBackupSnapshotDownloadCreate20241113) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


