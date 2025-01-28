// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DiskBackupSnapshotAWSExportBucketRequestAllOf struct for DiskBackupSnapshotAWSExportBucketRequestAllOf
type DiskBackupSnapshotAWSExportBucketRequestAllOf struct {
	// Unique 24-hexadecimal character string that identifies the <a href='https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/' target='_blank'>Unified AWS Access role ID</a>  that MongoDB Cloud uses to access the AWS S3 bucket.
	IamRoleId *string `json:"iamRoleId,omitempty"`
}

// NewDiskBackupSnapshotAWSExportBucketRequestAllOf instantiates a new DiskBackupSnapshotAWSExportBucketRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotAWSExportBucketRequestAllOf() *DiskBackupSnapshotAWSExportBucketRequestAllOf {
	this := DiskBackupSnapshotAWSExportBucketRequestAllOf{}
	return &this
}

// NewDiskBackupSnapshotAWSExportBucketRequestAllOfWithDefaults instantiates a new DiskBackupSnapshotAWSExportBucketRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotAWSExportBucketRequestAllOfWithDefaults() *DiskBackupSnapshotAWSExportBucketRequestAllOf {
	this := DiskBackupSnapshotAWSExportBucketRequestAllOf{}
	return &this
}

// GetIamRoleId returns the IamRoleId field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) GetIamRoleId() string {
	if o == nil || IsNil(o.IamRoleId) {
		var ret string
		return ret
	}
	return *o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) GetIamRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleId) {
		return nil, false
	}

	return o.IamRoleId, true
}

// HasIamRoleId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) HasIamRoleId() bool {
	if o != nil && !IsNil(o.IamRoleId) {
		return true
	}

	return false
}

// SetIamRoleId gets a reference to the given string and assigns it to the IamRoleId field.
func (o *DiskBackupSnapshotAWSExportBucketRequestAllOf) SetIamRoleId(v string) {
	o.IamRoleId = &v
}
