# DiskBackupSnapshotExportBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | The name of the Azure Storage Container to export to. Deprecated: provide the Container&#39;s URL in serviceUrl instead. | 
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider that Snapshots are exported to. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the &lt;a href&#x3D;&#39;https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/&#39; target&#x3D;&#39;_blank&#39;&gt;Unified AWS Access role ID&lt;/a&gt;  that MongoDB Cloud uses to access the AWS S3 bucket. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Azure Cloud Provider Access Role that MongoDB Cloud uses to access the Azure Blob Storage Container. | [optional] 
**ServiceUrl** | Pointer to **string** | URL of the Azure Storage Account to export to. For example: \&quot;https://examplestorageaccount.blob.core.windows.net\&quot;. Only standard endpoints (with \&quot;blob.core.windows.net\&quot;) are supported. | [optional] 
**TenantId** | Pointer to **string** | UUID that identifies the Azure Active Directory Tenant ID. Deprecated: this field is ignored; the tenantId of the Cloud Provider Access role (from roleId) is used. | [optional] 

## Methods

### NewDiskBackupSnapshotExportBucketRequest

`func NewDiskBackupSnapshotExportBucketRequest(bucketName string, cloudProvider string, ) *DiskBackupSnapshotExportBucketRequest`

NewDiskBackupSnapshotExportBucketRequest instantiates a new DiskBackupSnapshotExportBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotExportBucketRequestWithDefaults

`func NewDiskBackupSnapshotExportBucketRequestWithDefaults() *DiskBackupSnapshotExportBucketRequest`

NewDiskBackupSnapshotExportBucketRequestWithDefaults instantiates a new DiskBackupSnapshotExportBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *DiskBackupSnapshotExportBucketRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DiskBackupSnapshotExportBucketRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### GetCloudProvider

`func (o *DiskBackupSnapshotExportBucketRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupSnapshotExportBucketRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetLinks

`func (o *DiskBackupSnapshotExportBucketRequest) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotExportBucketRequest) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotExportBucketRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetIamRoleId

`func (o *DiskBackupSnapshotExportBucketRequest) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *DiskBackupSnapshotExportBucketRequest) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *DiskBackupSnapshotExportBucketRequest) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.
### GetRoleId

`func (o *DiskBackupSnapshotExportBucketRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DiskBackupSnapshotExportBucketRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *DiskBackupSnapshotExportBucketRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetServiceUrl

`func (o *DiskBackupSnapshotExportBucketRequest) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *DiskBackupSnapshotExportBucketRequest) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *DiskBackupSnapshotExportBucketRequest) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.
### GetTenantId

`func (o *DiskBackupSnapshotExportBucketRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiskBackupSnapshotExportBucketRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiskBackupSnapshotExportBucketRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


