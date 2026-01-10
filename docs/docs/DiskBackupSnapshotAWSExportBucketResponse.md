# DiskBackupSnapshotAWSExportBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal character string that identifies the Export Bucket. | 
**BucketName** | **string** | The name of the AWS S3 Bucket, Azure Storage Container, or Google Cloud Storage Bucket that Snapshots are exported to. | 
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider. | 
**IamRoleId** | **string** | Unique 24-hexadecimal character string that identifies the Unified AWS Access role ID that MongoDB Cloud uses to access the AWS S3 bucket. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Region** | Pointer to **string** | AWS region for the export bucket. This is set by Atlas and is never user-supplied. | [optional] [readonly] 
**RequirePrivateNetworking** | Pointer to **bool** | Indicates whether to use privatelink. User supplied. | [optional] 

## Methods

### NewDiskBackupSnapshotAWSExportBucketResponse

`func NewDiskBackupSnapshotAWSExportBucketResponse(id string, bucketName string, cloudProvider string, iamRoleId string, ) *DiskBackupSnapshotAWSExportBucketResponse`

NewDiskBackupSnapshotAWSExportBucketResponse instantiates a new DiskBackupSnapshotAWSExportBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotAWSExportBucketResponseWithDefaults

`func NewDiskBackupSnapshotAWSExportBucketResponseWithDefaults() *DiskBackupSnapshotAWSExportBucketResponse`

NewDiskBackupSnapshotAWSExportBucketResponseWithDefaults instantiates a new DiskBackupSnapshotAWSExportBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetId(v string)`

SetId sets Id field to given value.

### GetBucketName

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### GetCloudProvider

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### GetLinks

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotAWSExportBucketResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRegion

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DiskBackupSnapshotAWSExportBucketResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetRequirePrivateNetworking

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetRequirePrivateNetworking() bool`

GetRequirePrivateNetworking returns the RequirePrivateNetworking field if non-nil, zero value otherwise.

### GetRequirePrivateNetworkingOk

`func (o *DiskBackupSnapshotAWSExportBucketResponse) GetRequirePrivateNetworkingOk() (*bool, bool)`

GetRequirePrivateNetworkingOk returns a tuple with the RequirePrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePrivateNetworking

`func (o *DiskBackupSnapshotAWSExportBucketResponse) SetRequirePrivateNetworking(v bool)`

SetRequirePrivateNetworking sets RequirePrivateNetworking field to given value.

### HasRequirePrivateNetworking

`func (o *DiskBackupSnapshotAWSExportBucketResponse) HasRequirePrivateNetworking() bool`

HasRequirePrivateNetworking returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


