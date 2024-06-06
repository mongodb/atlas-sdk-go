# DiskBackupSnapshotExportBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the Amazon Web Services (AWS) Simple Storage Service (S3) export bucket. | [optional] [readonly] 
**BucketName** | Pointer to **string** | Human-readable label that identifies the AWS bucket that the role is authorized to access. | [optional] 
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider that stores this snapshot. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the &lt;a href&#x3D;&#39;https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/&#39; target&#x3D;&#39;_blank&#39;&gt;Unified AWS Access role ID&lt;/a&gt;  that MongoDB Cloud uses to access the AWS S3 bucket. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Azure Service Principal that MongoDB Cloud uses to access the Azure Blob Storage. | [optional] 
**ServiceUrl** | Pointer to **string** | Url that identifies the Azure Blob Storage Account. | [optional] 
**TenantId** | Pointer to **string** | UUID String that identifies the Azure Active Directory Tenant ID. | [optional] 

## Methods

### NewDiskBackupSnapshotExportBucket

`func NewDiskBackupSnapshotExportBucket() *DiskBackupSnapshotExportBucket`

NewDiskBackupSnapshotExportBucket instantiates a new DiskBackupSnapshotExportBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotExportBucketWithDefaults

`func NewDiskBackupSnapshotExportBucketWithDefaults() *DiskBackupSnapshotExportBucket`

NewDiskBackupSnapshotExportBucketWithDefaults instantiates a new DiskBackupSnapshotExportBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiskBackupSnapshotExportBucket) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskBackupSnapshotExportBucket) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskBackupSnapshotExportBucket) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiskBackupSnapshotExportBucket) HasId() bool`

HasId returns a boolean if a field has been set.
### GetBucketName

`func (o *DiskBackupSnapshotExportBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DiskBackupSnapshotExportBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DiskBackupSnapshotExportBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *DiskBackupSnapshotExportBucket) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.
### GetCloudProvider

`func (o *DiskBackupSnapshotExportBucket) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DiskBackupSnapshotExportBucket) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DiskBackupSnapshotExportBucket) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DiskBackupSnapshotExportBucket) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetLinks

`func (o *DiskBackupSnapshotExportBucket) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DiskBackupSnapshotExportBucket) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DiskBackupSnapshotExportBucket) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DiskBackupSnapshotExportBucket) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetIamRoleId

`func (o *DiskBackupSnapshotExportBucket) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *DiskBackupSnapshotExportBucket) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *DiskBackupSnapshotExportBucket) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *DiskBackupSnapshotExportBucket) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.
### GetRoleId

`func (o *DiskBackupSnapshotExportBucket) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DiskBackupSnapshotExportBucket) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DiskBackupSnapshotExportBucket) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *DiskBackupSnapshotExportBucket) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetServiceUrl

`func (o *DiskBackupSnapshotExportBucket) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *DiskBackupSnapshotExportBucket) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *DiskBackupSnapshotExportBucket) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *DiskBackupSnapshotExportBucket) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.
### GetTenantId

`func (o *DiskBackupSnapshotExportBucket) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *DiskBackupSnapshotExportBucket) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *DiskBackupSnapshotExportBucket) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *DiskBackupSnapshotExportBucket) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


