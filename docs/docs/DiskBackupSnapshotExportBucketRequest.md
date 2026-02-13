# DiskBackupSnapshotExportBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Human-readable label that identifies the cloud provider. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**BucketName** | Pointer to **string** | Human-readable label that identifies the Google Cloud Storage Bucket that the role is authorized to export to. | [optional] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the Unified AWS Access role ID that MongoDB Cloud uses to access the AWS S3 bucket. | [optional] 
**RequirePrivateNetworking** | Pointer to **bool** | Indicates whether to do exports over PrivateLink as opposed to public IPs. Defaults to False. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the GCP Cloud Provider Access Role that MongoDB Cloud uses to access the Google Cloud Storage Bucket. | [optional] 
**ServiceUrl** | Pointer to **string** | URL of the Azure Storage Account to export to. For example: &#x60;https://examplestorageaccount.blob.core.windows.net/exportcontainer&#x60;. Only standard endpoints (with &#x60;blob.core.windows.net&#x60;) are supported. | [optional] 
**TenantId** | Pointer to **string** | UUID that identifies the Azure Active Directory Tenant ID. Deprecated: this field is ignored; the &#x60;tenantId&#x60; of the Cloud Provider Access role (from &#x60;roleId&#x60;) is used. | [optional] 

## Methods

### NewDiskBackupSnapshotExportBucketRequest

`func NewDiskBackupSnapshotExportBucketRequest(cloudProvider string, ) *DiskBackupSnapshotExportBucketRequest`

NewDiskBackupSnapshotExportBucketRequest instantiates a new DiskBackupSnapshotExportBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotExportBucketRequestWithDefaults

`func NewDiskBackupSnapshotExportBucketRequestWithDefaults() *DiskBackupSnapshotExportBucketRequest`

NewDiskBackupSnapshotExportBucketRequestWithDefaults instantiates a new DiskBackupSnapshotExportBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasBucketName

`func (o *DiskBackupSnapshotExportBucketRequest) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.
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
### GetRequirePrivateNetworking

`func (o *DiskBackupSnapshotExportBucketRequest) GetRequirePrivateNetworking() bool`

GetRequirePrivateNetworking returns the RequirePrivateNetworking field if non-nil, zero value otherwise.

### GetRequirePrivateNetworkingOk

`func (o *DiskBackupSnapshotExportBucketRequest) GetRequirePrivateNetworkingOk() (*bool, bool)`

GetRequirePrivateNetworkingOk returns a tuple with the RequirePrivateNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePrivateNetworking

`func (o *DiskBackupSnapshotExportBucketRequest) SetRequirePrivateNetworking(v bool)`

SetRequirePrivateNetworking sets RequirePrivateNetworking field to given value.

### HasRequirePrivateNetworking

`func (o *DiskBackupSnapshotExportBucketRequest) HasRequirePrivateNetworking() bool`

HasRequirePrivateNetworking returns a boolean if a field has been set.
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


