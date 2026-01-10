# DiskBackupSnapshotExportBucketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-hexadecimal character string that identifies the Export Bucket. | 
**BucketName** | **string** | The name of the AWS S3 Bucket, Azure Storage Container, or Google Cloud Storage Bucket that Snapshots are exported to. | 
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the Unified AWS Access role ID that MongoDB Cloud uses to access the AWS S3 bucket. | [optional] 
**Region** | Pointer to **string** | AWS region for the export bucket. This is set by Atlas and is never user-supplied. | [optional] [readonly] 
**RequirePrivateNetworking** | Pointer to **bool** | Indicates whether to use privatelink. User supplied. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the GCP Cloud Provider Access Role that MongoDB Cloud uses to access the Google Cloud Storage Bucket. | [optional] 
**ServiceUrl** | Pointer to **string** | URL of the Azure Storage Account to export to. Only standard endpoints (with \&quot;blob.core.windows.net\&quot;) are supported. | [optional] 
**TenantId** | Pointer to **string** | UUID that identifies the Azure Active Directory Tenant ID used during exports. | [optional] 

## Methods

### NewDiskBackupSnapshotExportBucketResponse

`func NewDiskBackupSnapshotExportBucketResponse(id string, bucketName string, cloudProvider string, ) *DiskBackupSnapshotExportBucketResponse`

NewDiskBackupSnapshotExportBucketResponse instantiates a new DiskBackupSnapshotExportBucketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotExportBucketResponseWithDefaults

`func NewDiskBackupSnapshotExportBucketResponseWithDefaults() *DiskBackupSnapshotExportBucketResponse`

NewDiskBackupSnapshotExportBucketResponseWithDefaults instantiates a new DiskBackupSnapshotExportBucketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiskBackupSnapshotExportBucketResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupSnapshotExportBucketResponse) SetId(v string)`

SetId sets Id field to given value.

### GetBucketName

`func (o *DiskBackupSnapshotExportBucketResponse) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DiskBackupSnapshotExportBucketResponse) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### GetCloudProvider

`func (o *DiskBackupSnapshotExportBucketResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupSnapshotExportBucketResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetLinks

`func (o *DiskBackupSnapshotExportBucketResponse) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotExportBucketResponse) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotExportBucketResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetIamRoleId

`func (o *DiskBackupSnapshotExportBucketResponse) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *DiskBackupSnapshotExportBucketResponse) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *DiskBackupSnapshotExportBucketResponse) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.
### GetRegion

`func (o *DiskBackupSnapshotExportBucketResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DiskBackupSnapshotExportBucketResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DiskBackupSnapshotExportBucketResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetRequirePrivateNetworking

`func (o *DiskBackupSnapshotExportBucketResponse) GetRequirePrivateNetworking() bool`

GetRequirePrivateNetworking returns the RequirePrivateNetworking field if non-nil, zero value otherwise.

### GetRequirePrivateNetworkingOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetRequirePrivateNetworkingOk() (*bool, bool)`

GetRequirePrivateNetworkingOk returns a tuple with the RequirePrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePrivateNetworking

`func (o *DiskBackupSnapshotExportBucketResponse) SetRequirePrivateNetworking(v bool)`

SetRequirePrivateNetworking sets RequirePrivateNetworking field to given value.

### HasRequirePrivateNetworking

`func (o *DiskBackupSnapshotExportBucketResponse) HasRequirePrivateNetworking() bool`

HasRequirePrivateNetworking returns a boolean if a field has been set.
### GetRoleId

`func (o *DiskBackupSnapshotExportBucketResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DiskBackupSnapshotExportBucketResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *DiskBackupSnapshotExportBucketResponse) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetServiceUrl

`func (o *DiskBackupSnapshotExportBucketResponse) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *DiskBackupSnapshotExportBucketResponse) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *DiskBackupSnapshotExportBucketResponse) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.
### GetTenantId

`func (o *DiskBackupSnapshotExportBucketResponse) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiskBackupSnapshotExportBucketResponse) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiskBackupSnapshotExportBucketResponse) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiskBackupSnapshotExportBucketResponse) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


