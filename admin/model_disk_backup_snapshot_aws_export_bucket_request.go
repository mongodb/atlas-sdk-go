// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// DiskBackupSnapshotAWSExportBucketRequest struct for DiskBackupSnapshotAWSExportBucketRequest
type DiskBackupSnapshotAWSExportBucketRequest struct {
	// Human-readable label that identifies the AWS S3 Bucket that the role is authorized to export to.
	BucketName string `json:"bucketName"`
	// Unique 24-hexadecimal character string that identifies the Unified AWS Access role ID that MongoDB Cloud uses to access the AWS S3 bucket.
	IamRoleId string `json:"iamRoleId"`
	// Human-readable label that identifies the cloud provider that Snapshots are exported to.
	CloudProvider string `json:"cloudProvider"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the Azure Cloud Provider Access Role that MongoDB Cloud uses to access the Azure Blob Storage Container.
	RoleId *string `json:"roleId,omitempty"`
	// URL of the Azure Storage Account to export to. For example: \"https://examplestorageaccount.blob.core.windows.net/exportcontainer\". Only standard endpoints (with \"blob.core.windows.net\") are supported.
	ServiceUrl *string `json:"serviceUrl,omitempty"`
	// UUID that identifies the Azure Active Directory Tenant ID. Deprecated: this field is ignored; the tenantId of the Cloud Provider Access role (from roleId) is used.
	// Deprecated
	TenantId *string `json:"tenantId,omitempty"`
}

// NewDiskBackupSnapshotAWSExportBucketRequest instantiates a new DiskBackupSnapshotAWSExportBucketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotAWSExportBucketRequest(bucketName string, iamRoleId string, cloudProvider string) *DiskBackupSnapshotAWSExportBucketRequest {
	this := DiskBackupSnapshotAWSExportBucketRequest{}
	this.CloudProvider = cloudProvider
	this.BucketName = bucketName
	this.IamRoleId = iamRoleId
	return &this
}

// NewDiskBackupSnapshotAWSExportBucketRequestWithDefaults instantiates a new DiskBackupSnapshotAWSExportBucketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotAWSExportBucketRequestWithDefaults() *DiskBackupSnapshotAWSExportBucketRequest {
	this := DiskBackupSnapshotAWSExportBucketRequest{}
	return &this
}

// GetBucketName returns the BucketName field value
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetBucketName(v string) {
	o.BucketName = v
}

// GetIamRoleId returns the IamRoleId field value
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetIamRoleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IamRoleId
}

// GetIamRoleIdOk returns a tuple with the IamRoleId field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetIamRoleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IamRoleId, true
}

// SetIamRoleId sets field value
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetIamRoleId(v string) {
	o.IamRoleId = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetLinks(v []Link) {
	o.Links = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetRoleId() string {
	if o == nil || IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetRoleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoleId) {
		return nil, false
	}

	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) HasRoleId() bool {
	if o != nil && !IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetRoleId(v string) {
	o.RoleId = &v
}

// GetServiceUrl returns the ServiceUrl field value if set, zero value otherwise
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetServiceUrl() string {
	if o == nil || IsNil(o.ServiceUrl) {
		var ret string
		return ret
	}
	return *o.ServiceUrl
}

// GetServiceUrlOk returns a tuple with the ServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetServiceUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceUrl) {
		return nil, false
	}

	return o.ServiceUrl, true
}

// HasServiceUrl returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) HasServiceUrl() bool {
	if o != nil && !IsNil(o.ServiceUrl) {
		return true
	}

	return false
}

// SetServiceUrl gets a reference to the given string and assigns it to the ServiceUrl field.
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetServiceUrl(v string) {
	o.ServiceUrl = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise
// Deprecated
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DiskBackupSnapshotAWSExportBucketRequest) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}

	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotAWSExportBucketRequest) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
// Deprecated
func (o *DiskBackupSnapshotAWSExportBucketRequest) SetTenantId(v string) {
	o.TenantId = &v
}

func (o DiskBackupSnapshotAWSExportBucketRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DiskBackupSnapshotAWSExportBucketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucketName"] = o.BucketName
	toSerialize["iamRoleId"] = o.IamRoleId
	toSerialize["cloudProvider"] = o.CloudProvider
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
