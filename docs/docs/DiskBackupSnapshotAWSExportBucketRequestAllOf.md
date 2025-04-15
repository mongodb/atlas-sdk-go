# DiskBackupSnapshotAWSExportBucketRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | Human-readable label that identifies the AWS S3 Bucket that the role is authorized to export to. | [optional] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the Unified AWS Access role ID that MongoDB Cloud uses to access the AWS S3 bucket. | [optional] 

## Methods

### NewDiskBackupSnapshotAWSExportBucketRequestAllOf

`func NewDiskBackupSnapshotAWSExportBucketRequestAllOf() *DiskBackupSnapshotAWSExportBucketRequestAllOf`

NewDiskBackupSnapshotAWSExportBucketRequestAllOf instantiates a new DiskBackupSnapshotAWSExportBucketRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskBackupSnapshotAWSExportBucketRequestAllOfWithDefaults

`func NewDiskBackupSnapshotAWSExportBucketRequestAllOfWithDefaults() *DiskBackupSnapshotAWSExportBucketRequestAllOf`

NewDiskBackupSnapshotAWSExportBucketRequestAllOfWithDefaults instantiates a new DiskBackupSnapshotAWSExportBucketRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.
### GetIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


