// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// FlexClusterDescription20250101 Group of settings that configure a MongoDB Flex cluster.
type FlexClusterDescription20250101 struct {
	BackupSettings *FlexBackupSettings20250101 `json:"backupSettings,omitempty"`
	// Flex cluster topology.
	// Read only field.
	ClusterType       *string                        `json:"clusterType,omitempty"`
	ConnectionStrings *FlexConnectionStrings20250101 `json:"connectionStrings,omitempty"`
	// Date and time when MongoDB Cloud created this instance. This parameter expresses its value in ISO 8601 format in UTC.
	// Read only field.
	CreateDate *time.Time `json:"createDate,omitempty"`
	// Unique 24-hexadecimal character string that identifies the project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the instance.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Version of MongoDB that the instance runs.
	// Read only field.
	MongoDBVersion *string `json:"mongoDBVersion,omitempty"`
	// Human-readable label that identifies the instance.
	// Read only field.
	Name             *string                      `json:"name,omitempty"`
	ProviderSettings FlexProviderSettings20250101 `json:"providerSettings"`
	// Human-readable label that indicates the current operating condition of this instance.
	// Read only field.
	StateName *string `json:"stateName,omitempty"`
	// List that contains key-value pairs between 1 to 255 characters in length for tagging and categorizing the instance.
	Tags *[]ResourceTag `json:"tags,omitempty"`
	// Flag that indicates whether termination protection is enabled on the cluster. If set to `true`, MongoDB Cloud won't delete the cluster. If set to `false`, MongoDB Cloud will delete the cluster.
	TerminationProtectionEnabled *bool `json:"terminationProtectionEnabled,omitempty"`
	// Method by which the cluster maintains the MongoDB versions.
	// Read only field.
	VersionReleaseSystem *string `json:"versionReleaseSystem,omitempty"`
}

// NewFlexClusterDescription20250101 instantiates a new FlexClusterDescription20250101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexClusterDescription20250101(providerSettings FlexProviderSettings20250101) *FlexClusterDescription20250101 {
	this := FlexClusterDescription20250101{}
	this.ProviderSettings = providerSettings
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// NewFlexClusterDescription20250101WithDefaults instantiates a new FlexClusterDescription20250101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexClusterDescription20250101WithDefaults() *FlexClusterDescription20250101 {
	this := FlexClusterDescription20250101{}
	var terminationProtectionEnabled bool = false
	this.TerminationProtectionEnabled = &terminationProtectionEnabled
	return &this
}

// GetBackupSettings returns the BackupSettings field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetBackupSettings() FlexBackupSettings20250101 {
	if o == nil || IsNil(o.BackupSettings) {
		var ret FlexBackupSettings20250101
		return ret
	}
	return *o.BackupSettings
}

// GetBackupSettingsOk returns a tuple with the BackupSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetBackupSettingsOk() (*FlexBackupSettings20250101, bool) {
	if o == nil || IsNil(o.BackupSettings) {
		return nil, false
	}

	return o.BackupSettings, true
}

// HasBackupSettings returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasBackupSettings() bool {
	if o != nil && !IsNil(o.BackupSettings) {
		return true
	}

	return false
}

// SetBackupSettings gets a reference to the given FlexBackupSettings20250101 and assigns it to the BackupSettings field.
func (o *FlexClusterDescription20250101) SetBackupSettings(v FlexBackupSettings20250101) {
	o.BackupSettings = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetClusterType() string {
	if o == nil || IsNil(o.ClusterType) {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetClusterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterType) {
		return nil, false
	}

	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasClusterType() bool {
	if o != nil && !IsNil(o.ClusterType) {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *FlexClusterDescription20250101) SetClusterType(v string) {
	o.ClusterType = &v
}

// GetConnectionStrings returns the ConnectionStrings field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetConnectionStrings() FlexConnectionStrings20250101 {
	if o == nil || IsNil(o.ConnectionStrings) {
		var ret FlexConnectionStrings20250101
		return ret
	}
	return *o.ConnectionStrings
}

// GetConnectionStringsOk returns a tuple with the ConnectionStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetConnectionStringsOk() (*FlexConnectionStrings20250101, bool) {
	if o == nil || IsNil(o.ConnectionStrings) {
		return nil, false
	}

	return o.ConnectionStrings, true
}

// HasConnectionStrings returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasConnectionStrings() bool {
	if o != nil && !IsNil(o.ConnectionStrings) {
		return true
	}

	return false
}

// SetConnectionStrings gets a reference to the given FlexConnectionStrings20250101 and assigns it to the ConnectionStrings field.
func (o *FlexClusterDescription20250101) SetConnectionStrings(v FlexConnectionStrings20250101) {
	o.ConnectionStrings = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetCreateDate() time.Time {
	if o == nil || IsNil(o.CreateDate) {
		var ret time.Time
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetCreateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateDate) {
		return nil, false
	}

	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasCreateDate() bool {
	if o != nil && !IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given time.Time and assigns it to the CreateDate field.
func (o *FlexClusterDescription20250101) SetCreateDate(v time.Time) {
	o.CreateDate = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *FlexClusterDescription20250101) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FlexClusterDescription20250101) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *FlexClusterDescription20250101) SetLinks(v []Link) {
	o.Links = &v
}

// GetMongoDBVersion returns the MongoDBVersion field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetMongoDBVersion() string {
	if o == nil || IsNil(o.MongoDBVersion) {
		var ret string
		return ret
	}
	return *o.MongoDBVersion
}

// GetMongoDBVersionOk returns a tuple with the MongoDBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetMongoDBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MongoDBVersion) {
		return nil, false
	}

	return o.MongoDBVersion, true
}

// HasMongoDBVersion returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasMongoDBVersion() bool {
	if o != nil && !IsNil(o.MongoDBVersion) {
		return true
	}

	return false
}

// SetMongoDBVersion gets a reference to the given string and assigns it to the MongoDBVersion field.
func (o *FlexClusterDescription20250101) SetMongoDBVersion(v string) {
	o.MongoDBVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FlexClusterDescription20250101) SetName(v string) {
	o.Name = &v
}

// GetProviderSettings returns the ProviderSettings field value
func (o *FlexClusterDescription20250101) GetProviderSettings() FlexProviderSettings20250101 {
	if o == nil {
		var ret FlexProviderSettings20250101
		return ret
	}

	return o.ProviderSettings
}

// GetProviderSettingsOk returns a tuple with the ProviderSettings field value
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetProviderSettingsOk() (*FlexProviderSettings20250101, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderSettings, true
}

// SetProviderSettings sets field value
func (o *FlexClusterDescription20250101) SetProviderSettings(v FlexProviderSettings20250101) {
	o.ProviderSettings = v
}

// GetStateName returns the StateName field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetStateName() string {
	if o == nil || IsNil(o.StateName) {
		var ret string
		return ret
	}
	return *o.StateName
}

// GetStateNameOk returns a tuple with the StateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetStateNameOk() (*string, bool) {
	if o == nil || IsNil(o.StateName) {
		return nil, false
	}

	return o.StateName, true
}

// HasStateName returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasStateName() bool {
	if o != nil && !IsNil(o.StateName) {
		return true
	}

	return false
}

// SetStateName gets a reference to the given string and assigns it to the StateName field.
func (o *FlexClusterDescription20250101) SetStateName(v string) {
	o.StateName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetTags() []ResourceTag {
	if o == nil || IsNil(o.Tags) {
		var ret []ResourceTag
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetTagsOk() (*[]ResourceTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}

	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []ResourceTag and assigns it to the Tags field.
func (o *FlexClusterDescription20250101) SetTags(v []ResourceTag) {
	o.Tags = &v
}

// GetTerminationProtectionEnabled returns the TerminationProtectionEnabled field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetTerminationProtectionEnabled() bool {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		var ret bool
		return ret
	}
	return *o.TerminationProtectionEnabled
}

// GetTerminationProtectionEnabledOk returns a tuple with the TerminationProtectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetTerminationProtectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TerminationProtectionEnabled) {
		return nil, false
	}

	return o.TerminationProtectionEnabled, true
}

// HasTerminationProtectionEnabled returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasTerminationProtectionEnabled() bool {
	if o != nil && !IsNil(o.TerminationProtectionEnabled) {
		return true
	}

	return false
}

// SetTerminationProtectionEnabled gets a reference to the given bool and assigns it to the TerminationProtectionEnabled field.
func (o *FlexClusterDescription20250101) SetTerminationProtectionEnabled(v bool) {
	o.TerminationProtectionEnabled = &v
}

// GetVersionReleaseSystem returns the VersionReleaseSystem field value if set, zero value otherwise
func (o *FlexClusterDescription20250101) GetVersionReleaseSystem() string {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		var ret string
		return ret
	}
	return *o.VersionReleaseSystem
}

// GetVersionReleaseSystemOk returns a tuple with the VersionReleaseSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexClusterDescription20250101) GetVersionReleaseSystemOk() (*string, bool) {
	if o == nil || IsNil(o.VersionReleaseSystem) {
		return nil, false
	}

	return o.VersionReleaseSystem, true
}

// HasVersionReleaseSystem returns a boolean if a field has been set.
func (o *FlexClusterDescription20250101) HasVersionReleaseSystem() bool {
	if o != nil && !IsNil(o.VersionReleaseSystem) {
		return true
	}

	return false
}

// SetVersionReleaseSystem gets a reference to the given string and assigns it to the VersionReleaseSystem field.
func (o *FlexClusterDescription20250101) SetVersionReleaseSystem(v string) {
	o.VersionReleaseSystem = &v
}
