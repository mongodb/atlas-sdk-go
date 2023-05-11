# CloudProviderAccessFeatureUsageExportSnapshotFeatureId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportBucketId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the AWS S3 bucket to which you export your snapshots. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies your project. | [optional] [readonly] 

## Methods

### NewCloudProviderAccessFeatureUsageExportSnapshotFeatureId

`func NewCloudProviderAccessFeatureUsageExportSnapshotFeatureId() *CloudProviderAccessFeatureUsageExportSnapshotFeatureId`

NewCloudProviderAccessFeatureUsageExportSnapshotFeatureId instantiates a new CloudProviderAccessFeatureUsageExportSnapshotFeatureId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderAccessFeatureUsageExportSnapshotFeatureIdWithDefaults

`func NewCloudProviderAccessFeatureUsageExportSnapshotFeatureIdWithDefaults() *CloudProviderAccessFeatureUsageExportSnapshotFeatureId`

NewCloudProviderAccessFeatureUsageExportSnapshotFeatureIdWithDefaults instantiates a new CloudProviderAccessFeatureUsageExportSnapshotFeatureId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportBucketId

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetExportBucketId() string`

GetExportBucketId returns the ExportBucketId field if non-nil, zero value otherwise.

### GetExportBucketIdOk

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetExportBucketIdOk() (*string, bool)`

GetExportBucketIdOk returns a tuple with the ExportBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportBucketId

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) SetExportBucketId(v string)`

SetExportBucketId sets ExportBucketId field to given value.

### HasExportBucketId

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) HasExportBucketId() bool`

HasExportBucketId returns a boolean if a field has been set.

### GetGroupId

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CloudProviderAccessFeatureUsageExportSnapshotFeatureId) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


