# DiskBackupSnapshotAWSExportBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Human-readable label that identifies the AWS S3 Bucket that the role is authorized to export to. | 
**IamRoleId** | **string** | Unique 24-hexadecimal character string that identifies the Unified AWS Access role ID that MongoDB Cloud uses to access the AWS S3 bucket. | 
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider that Snapshots are exported to. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Azure Cloud Provider Access Role that MongoDB Cloud uses to access the Azure Blob Storage Container. | [optional] 
**ServiceUrl** | Pointer to **string** | URL of the Azure Storage Account to export to. For example: \&quot;https://examplestorageaccount.blob.core.windows.net/exportcontainer\&quot;. Only standard endpoints (with \&quot;blob.core.windows.net\&quot;) are supported. | [optional] 
**TenantId** | Pointer to **string** | UUID that identifies the Azure Active Directory Tenant ID. Deprecated: this field is ignored; the tenantId of the Cloud Provider Access role (from roleId) is used. | [optional] 

## Methods

### NewDiskBackupSnapshotAWSExportBucketRequest

`func NewDiskBackupSnapshotAWSExportBucketRequest(bucketName string, iamRoleId string, cloudProvider string, ) *DiskBackupSnapshotAWSExportBucketRequest`

NewDiskBackupSnapshotAWSExportBucketRequest instantiates a new DiskBackupSnapshotAWSExportBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotAWSExportBucketRequestWithDefaults

`func NewDiskBackupSnapshotAWSExportBucketRequestWithDefaults() *DiskBackupSnapshotAWSExportBucketRequest`

NewDiskBackupSnapshotAWSExportBucketRequestWithDefaults instantiates a new DiskBackupSnapshotAWSExportBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### GetIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### GetCloudProvider

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetLinks

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotAWSExportBucketRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetServiceUrl

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *DiskBackupSnapshotAWSExportBucketRequest) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.
### GetTenantId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiskBackupSnapshotAWSExportBucketRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiskBackupSnapshotAWSExportBucketRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


