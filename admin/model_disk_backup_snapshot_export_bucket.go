// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// DiskBackupSnapshotExportBucket Disk backup snapshot Export Bucket.
type DiskBackupSnapshotExportBucket struct {
	// Unique 24-hexadecimal character string that identifies the Export Bucket.
	// Read only field.
	Id *string `json:"_id,omitempty"`
	// Human-readable label that identifies the AWS S3 Bucket or Azure Storage Container that the role is authorized to export to.
	BucketName string `json:"bucketName"`
	// Human-readable label that identifies the cloud provider that Snapshots will be exported to.
	CloudProvider string `json:"cloudProvider"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Unique 24-hexadecimal character string that identifies the <a href='https://www.mongodb.com/docs/atlas/security/set-up-unified-aws-access/' target='_blank'>Unified AWS Access role ID</a>  that MongoDB Cloud uses to access the AWS S3 bucket.
	IamRoleId *string `json:"iamRoleId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Azure Service Principal that MongoDB Cloud uses to access the Azure Blob Storage Container.
	RoleId *string `json:"roleId,omitempty"`
	// URL that identifies the blob Endpoint of the Azure Blob Storage Account.
	ServiceUrl *string `json:"serviceUrl,omitempty"`
	// UUID that identifies the Azure Active Directory Tenant ID.
	TenantId *string `json:"tenantId,omitempty"`
}

// NewDiskBackupSnapshotExportBucket instantiates a new DiskBackupSnapshotExportBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotExportBucket(bucketName string, cloudProvider string) *DiskBackupSnapshotExportBucket {
	this := DiskBackupSnapshotExportBucket{}
	this.BucketName = bucketName
	this.CloudProvider = cloudProvider
	return &this
}

// NewDiskBackupSnapshotExportBucketWithDefaults instantiates a new DiskBackupSnapshotExportBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotExportBucketWithDefaults() *DiskBackupSnapshotExportBucket {
	this := DiskBackupSnapshotExportBucket{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *DiskBackupSnapshotExportBucket) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotExportBucket) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DiskBackupSnapshotExportBucket) SetId(v string) {
	o.Id = &v
}

// GetBucketName returns the BucketName field value
func (o *DiskBackupSnapshotExportBucket) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *DiskBackupSnapshotExportBucket) SetBucketName(v string) {
	o.BucketName = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *DiskBackupSnapshotExportBucket) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DiskBackupSnapshotExportBucket) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *DiskBackupSnapshotExportBucket) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupSnapshotExportBucket) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupSnapshotExportBucket) SetLinks(v []Link) {
	o.Links = &v
}

// GetIamRoleId returns the IamRoleId field value if set, zero value otherwise
func (o *DiskBackupSnapshotExportBucket) GetIamRoleId() string {
	if o == nil || IsNil(o.IamRoleId) {
		var ret string
		return ret
	}
	return *o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetIamRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.IamRoleId) {
		return nil, false
	}

	return o.IamRoleId, true
}

// HasIamRoleId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotExportBucket) HasIamRoleId() bool {
	if o != nil && !IsNil(o.IamRoleId) {
		return true
	}

	return false
}

// SetIamRoleId gets a reference to the given string and assigns it to the IamRoleId field.
func (o *DiskBackupSnapshotExportBucket) SetIamRoleId(v string) {
	o.IamRoleId = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise
func (o *DiskBackupSnapshotExportBucket) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}

	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotExportBucket) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *DiskBackupSnapshotExportBucket) SetRoleId(v string) {
	o.RoleId = &v
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise
func (o *DiskBackupSnapshotExportBucket) GetServiceUrl() string {
	if o == nil || IsNil(o.ServiceUrl) {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUrl) {
		return nil, false
	}

	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *DiskBackupSnapshotExportBucket) HasServiceUrl() bool {
	if o != nil && !IsNil(o.ServiceUrl) {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *DiskBackupSnapshotExportBucket) SetServiceUrl(v string) {
	o.ServiceUrl = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise
func (o *DiskBackupSnapshotExportBucket) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotExportBucket) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}

	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotExportBucket) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *DiskBackupSnapshotExportBucket) SetTenantId(v string) {
	o.TenantId = &v
}

func (o DiskBackupSnapshotExportBucket) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupSnapshotExportBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucketName"] = o.BucketName
	toSerialize["cloudProvider"] = o.CloudProvider
	if !IsNil(o.IamRoleId) {
		toSerialize["iamRoleId"] = o.IamRoleId
	}
	if !IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !IsNil(o.ServiceUrl) {
		toSerialize["serviceUrl"] = o.ServiceUrl
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	return toSerialize, nil
}
