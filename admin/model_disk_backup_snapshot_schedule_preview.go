// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// DiskBackupSnapshotSchedulePreview struct for DiskBackupSnapshotSchedulePreview
type DiskBackupSnapshotSchedulePreview struct {
	// Flag that indicates whether the copy settings are automatically managed by MongoDB Cloud and sync to the cluster topology.
	AutoCopySettingsEnabled *bool `json:"autoCopySettingsEnabled,omitempty"`
	// Flag that indicates whether MongoDB Cloud automatically exports Cloud Backup Snapshots to the Export Bucket.
	AutoExportEnabled *bool `json:"autoExportEnabled,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the cluster with the Snapshot you want to return.
	// Read only field.
	ClusterId *string `json:"clusterId,omitempty"`
	// Human-readable label that identifies the cluster with the Snapshot you want to return.
	// Read only field.
	ClusterName *string `json:"clusterName,omitempty"`
	// Flag that indicates whether copy settings use `copyPolicyItems` instead of `frequencies`. When true, requests must supply `copyPolicyItems` and responses return `copyPolicyItems` only. When false or omitted, requests must supply `frequencies` and responses return `frequencies` only.
	CopyPolicyItemsEnabled *bool `json:"copyPolicyItemsEnabled,omitempty"`
	// List that contains a document for each copy setting item in the desired backup policy.
	CopySettings *[]DiskBackupCopySetting20240805 `json:"copySettings,omitempty"`
	// List that contains a document for each deleted copy setting whose backup copies you want to delete.
	// Write only field.
	DeleteCopiedBackups *[]DeleteCopiedBackups20240805 `json:"deleteCopiedBackups,omitempty"`
	// Flag that indicates whether to delete Snapshot copies that MongoDB Cloud took previously when their associated `copyPolicyItems` are removed from a `copySetting`. This option requires `copyPolicyItemsEnabled` to be true.
	// Write only field.
	DeleteCopySnapshots *bool `json:"deleteCopySnapshots,omitempty"`
	// Flag that indicates whether to delete Snapshots that MongoDB Cloud took previously when deleting the associated backup policy.
	// Write only field.
	DeleteSnapshots *bool             `json:"deleteSnapshots,omitempty"`
	Export          *AutoExportPolicy `json:"export,omitempty"`
	// List that contains a document for each extra retention setting item in the desired backup policy.
	ExtraRetentionSettings *[]ExtraRetentionSetting `json:"extraRetentionSettings,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Date and time when MongoDB Cloud takes the next Snapshot. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	NextSnapshot *time.Time `json:"nextSnapshot,omitempty"`
	// Flag that indicates whether continuous cloud backup is enabled.
	// Read only field.
	PitEnabled *bool `json:"pitEnabled,omitempty"`
	// Rules set for this backup schedule.
	Policies []AdvancedDiskBackupSnapshotSchedulePolicy `json:"policies"`
	// Source of the cluster backup policy.
	// Read only field.
	PolicySource *string `json:"policySource,omitempty"`
	// Hour of day in Coordinated Universal Time (UTC) that represents when MongoDB Cloud takes the Snapshot.
	ReferenceHourOfDay *int `json:"referenceHourOfDay,omitempty"`
	// Minute of the `referenceHourOfDay` that represents when MongoDB Cloud takes the Snapshot.
	ReferenceMinuteOfHour *int `json:"referenceMinuteOfHour,omitempty"`
	// Number of previous days that you can restore back to with Continuous Cloud Backup accuracy. You must specify a positive, non-zero integer. This parameter applies to continuous Cloud Backups only.
	RestoreWindowDays *int `json:"restoreWindowDays,omitempty"`
	// Flag that indicates whether to apply the retention changes for updated copy policy items to Snapshot copies that MongoDB Cloud took previously.
	// Write only field.
	UpdateCopySnapshots *bool `json:"updateCopySnapshots,omitempty"`
	// Flag that indicates whether to apply the retention changes in the updated backup policy to Snapshots that MongoDB Cloud took previously.
	// Write only field.
	UpdateSnapshots *bool `json:"updateSnapshots,omitempty"`
	// Flag that indicates whether to use organization and project names instead of organization and project UUIDs in the path to the metadata files that MongoDB Cloud uploads to your Export Bucket.
	UseOrgAndGroupNamesInExportPrefix *bool `json:"useOrgAndGroupNamesInExportPrefix,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o DiskBackupSnapshotSchedulePreview) MarshalJSON() ([]byte, error) {
	type noMethod DiskBackupSnapshotSchedulePreview
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewDiskBackupSnapshotSchedulePreview instantiates a new DiskBackupSnapshotSchedulePreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupSnapshotSchedulePreview(policies []AdvancedDiskBackupSnapshotSchedulePolicy) *DiskBackupSnapshotSchedulePreview {
	this := DiskBackupSnapshotSchedulePreview{}
	this.Policies = policies
	return &this
}

// NewDiskBackupSnapshotSchedulePreviewWithDefaults instantiates a new DiskBackupSnapshotSchedulePreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupSnapshotSchedulePreviewWithDefaults() *DiskBackupSnapshotSchedulePreview {
	this := DiskBackupSnapshotSchedulePreview{}
	return &this
}

// GetAutoCopySettingsEnabled returns the AutoCopySettingsEnabled field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetAutoCopySettingsEnabled() bool {
	if o == nil || IsNil(o.AutoCopySettingsEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoCopySettingsEnabled
}

// GetAutoCopySettingsEnabledOk returns a tuple with the AutoCopySettingsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetAutoCopySettingsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCopySettingsEnabled) {
		return nil, false
	}

	return o.AutoCopySettingsEnabled, true
}

// HasAutoCopySettingsEnabled returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasAutoCopySettingsEnabled() bool {
	if o != nil && !IsNil(o.AutoCopySettingsEnabled) {
		return true
	}

	return false
}

// SetAutoCopySettingsEnabled gets a reference to the given bool and assigns it to the AutoCopySettingsEnabled field.
func (o *DiskBackupSnapshotSchedulePreview) SetAutoCopySettingsEnabled(v bool) {
	o.AutoCopySettingsEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "AutoCopySettingsEnabled")
}

// SetAutoCopySettingsEnabledNil sets AutoCopySettingsEnabled to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetAutoCopySettingsEnabledNil() {
	o.AutoCopySettingsEnabled = nil
	o.NullFields = addNullField(o.NullFields, "AutoCopySettingsEnabled")
}

// GetAutoExportEnabled returns the AutoExportEnabled field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetAutoExportEnabled() bool {
	if o == nil || IsNil(o.AutoExportEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoExportEnabled
}

// GetAutoExportEnabledOk returns a tuple with the AutoExportEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetAutoExportEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoExportEnabled) {
		return nil, false
	}

	return o.AutoExportEnabled, true
}

// HasAutoExportEnabled returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasAutoExportEnabled() bool {
	if o != nil && !IsNil(o.AutoExportEnabled) {
		return true
	}

	return false
}

// SetAutoExportEnabled gets a reference to the given bool and assigns it to the AutoExportEnabled field.
func (o *DiskBackupSnapshotSchedulePreview) SetAutoExportEnabled(v bool) {
	o.AutoExportEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "AutoExportEnabled")
}

// SetAutoExportEnabledNil sets AutoExportEnabled to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetAutoExportEnabledNil() {
	o.AutoExportEnabled = nil
	o.NullFields = addNullField(o.NullFields, "AutoExportEnabled")
}

// GetClusterId returns the ClusterId field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetClusterId() string {
	if o == nil || IsNil(o.ClusterId) {
		var ret string
		return ret
	}
	return *o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetClusterIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterId) {
		return nil, false
	}

	return o.ClusterId, true
}

// HasClusterId returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasClusterId() bool {
	if o != nil && !IsNil(o.ClusterId) {
		return true
	}

	return false
}

// SetClusterId gets a reference to the given string and assigns it to the ClusterId field.
func (o *DiskBackupSnapshotSchedulePreview) SetClusterId(v string) {
	o.ClusterId = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterId")
}

// SetClusterIdNil sets ClusterId to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetClusterIdNil() {
	o.ClusterId = nil
	o.NullFields = addNullField(o.NullFields, "ClusterId")
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetClusterName() string {
	if o == nil || IsNil(o.ClusterName) {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterName) {
		return nil, false
	}

	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasClusterName() bool {
	if o != nil && !IsNil(o.ClusterName) {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *DiskBackupSnapshotSchedulePreview) SetClusterName(v string) {
	o.ClusterName = &v
	o.NullFields = removeNullField(o.NullFields, "ClusterName")
}

// SetClusterNameNil sets ClusterName to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetClusterNameNil() {
	o.ClusterName = nil
	o.NullFields = addNullField(o.NullFields, "ClusterName")
}

// GetCopyPolicyItemsEnabled returns the CopyPolicyItemsEnabled field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetCopyPolicyItemsEnabled() bool {
	if o == nil || IsNil(o.CopyPolicyItemsEnabled) {
		var ret bool
		return ret
	}
	return *o.CopyPolicyItemsEnabled
}

// GetCopyPolicyItemsEnabledOk returns a tuple with the CopyPolicyItemsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetCopyPolicyItemsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyPolicyItemsEnabled) {
		return nil, false
	}

	return o.CopyPolicyItemsEnabled, true
}

// HasCopyPolicyItemsEnabled returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasCopyPolicyItemsEnabled() bool {
	if o != nil && !IsNil(o.CopyPolicyItemsEnabled) {
		return true
	}

	return false
}

// SetCopyPolicyItemsEnabled gets a reference to the given bool and assigns it to the CopyPolicyItemsEnabled field.
func (o *DiskBackupSnapshotSchedulePreview) SetCopyPolicyItemsEnabled(v bool) {
	o.CopyPolicyItemsEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "CopyPolicyItemsEnabled")
}

// SetCopyPolicyItemsEnabledNil sets CopyPolicyItemsEnabled to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetCopyPolicyItemsEnabledNil() {
	o.CopyPolicyItemsEnabled = nil
	o.NullFields = addNullField(o.NullFields, "CopyPolicyItemsEnabled")
}

// GetCopySettings returns the CopySettings field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetCopySettings() []DiskBackupCopySetting20240805 {
	if o == nil || IsNil(o.CopySettings) {
		var ret []DiskBackupCopySetting20240805
		return ret
	}
	return *o.CopySettings
}

// GetCopySettingsOk returns a tuple with the CopySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetCopySettingsOk() (*[]DiskBackupCopySetting20240805, bool) {
	if o == nil || IsNil(o.CopySettings) {
		return nil, false
	}

	return o.CopySettings, true
}

// HasCopySettings returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasCopySettings() bool {
	if o != nil && !IsNil(o.CopySettings) {
		return true
	}

	return false
}

// SetCopySettings gets a reference to the given []DiskBackupCopySetting20240805 and assigns it to the CopySettings field.
func (o *DiskBackupSnapshotSchedulePreview) SetCopySettings(v []DiskBackupCopySetting20240805) {
	o.CopySettings = &v
	o.NullFields = removeNullField(o.NullFields, "CopySettings")
}

// SetCopySettingsNil sets CopySettings to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetCopySettingsNil() {
	o.CopySettings = nil
	o.NullFields = addNullField(o.NullFields, "CopySettings")
}

// GetDeleteCopiedBackups returns the DeleteCopiedBackups field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetDeleteCopiedBackups() []DeleteCopiedBackups20240805 {
	if o == nil || IsNil(o.DeleteCopiedBackups) {
		var ret []DeleteCopiedBackups20240805
		return ret
	}
	return *o.DeleteCopiedBackups
}

// GetDeleteCopiedBackupsOk returns a tuple with the DeleteCopiedBackups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetDeleteCopiedBackupsOk() (*[]DeleteCopiedBackups20240805, bool) {
	if o == nil || IsNil(o.DeleteCopiedBackups) {
		return nil, false
	}

	return o.DeleteCopiedBackups, true
}

// HasDeleteCopiedBackups returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasDeleteCopiedBackups() bool {
	if o != nil && !IsNil(o.DeleteCopiedBackups) {
		return true
	}

	return false
}

// SetDeleteCopiedBackups gets a reference to the given []DeleteCopiedBackups20240805 and assigns it to the DeleteCopiedBackups field.
func (o *DiskBackupSnapshotSchedulePreview) SetDeleteCopiedBackups(v []DeleteCopiedBackups20240805) {
	o.DeleteCopiedBackups = &v
	o.NullFields = removeNullField(o.NullFields, "DeleteCopiedBackups")
}

// SetDeleteCopiedBackupsNil sets DeleteCopiedBackups to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetDeleteCopiedBackupsNil() {
	o.DeleteCopiedBackups = nil
	o.NullFields = addNullField(o.NullFields, "DeleteCopiedBackups")
}

// GetDeleteCopySnapshots returns the DeleteCopySnapshots field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetDeleteCopySnapshots() bool {
	if o == nil || IsNil(o.DeleteCopySnapshots) {
		var ret bool
		return ret
	}
	return *o.DeleteCopySnapshots
}

// GetDeleteCopySnapshotsOk returns a tuple with the DeleteCopySnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetDeleteCopySnapshotsOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteCopySnapshots) {
		return nil, false
	}

	return o.DeleteCopySnapshots, true
}

// HasDeleteCopySnapshots returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasDeleteCopySnapshots() bool {
	if o != nil && !IsNil(o.DeleteCopySnapshots) {
		return true
	}

	return false
}

// SetDeleteCopySnapshots gets a reference to the given bool and assigns it to the DeleteCopySnapshots field.
func (o *DiskBackupSnapshotSchedulePreview) SetDeleteCopySnapshots(v bool) {
	o.DeleteCopySnapshots = &v
	o.NullFields = removeNullField(o.NullFields, "DeleteCopySnapshots")
}

// SetDeleteCopySnapshotsNil sets DeleteCopySnapshots to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetDeleteCopySnapshotsNil() {
	o.DeleteCopySnapshots = nil
	o.NullFields = addNullField(o.NullFields, "DeleteCopySnapshots")
}

// GetDeleteSnapshots returns the DeleteSnapshots field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetDeleteSnapshots() bool {
	if o == nil || IsNil(o.DeleteSnapshots) {
		var ret bool
		return ret
	}
	return *o.DeleteSnapshots
}

// GetDeleteSnapshotsOk returns a tuple with the DeleteSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetDeleteSnapshotsOk() (*bool, bool) {
	if o == nil || IsNil(o.DeleteSnapshots) {
		return nil, false
	}

	return o.DeleteSnapshots, true
}

// HasDeleteSnapshots returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasDeleteSnapshots() bool {
	if o != nil && !IsNil(o.DeleteSnapshots) {
		return true
	}

	return false
}

// SetDeleteSnapshots gets a reference to the given bool and assigns it to the DeleteSnapshots field.
func (o *DiskBackupSnapshotSchedulePreview) SetDeleteSnapshots(v bool) {
	o.DeleteSnapshots = &v
	o.NullFields = removeNullField(o.NullFields, "DeleteSnapshots")
}

// SetDeleteSnapshotsNil sets DeleteSnapshots to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetDeleteSnapshotsNil() {
	o.DeleteSnapshots = nil
	o.NullFields = addNullField(o.NullFields, "DeleteSnapshots")
}

// GetExport returns the Export field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetExport() AutoExportPolicy {
	if o == nil || IsNil(o.Export) {
		var ret AutoExportPolicy
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetExportOk() (*AutoExportPolicy, bool) {
	if o == nil || IsNil(o.Export) {
		return nil, false
	}

	return o.Export, true
}

// HasExport returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasExport() bool {
	if o != nil && !IsNil(o.Export) {
		return true
	}

	return false
}

// SetExport gets a reference to the given AutoExportPolicy and assigns it to the Export field.
func (o *DiskBackupSnapshotSchedulePreview) SetExport(v AutoExportPolicy) {
	o.Export = &v
	o.NullFields = removeNullField(o.NullFields, "Export")
}

// SetExportNil sets Export to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetExportNil() {
	o.Export = nil
	o.NullFields = addNullField(o.NullFields, "Export")
}

// GetExtraRetentionSettings returns the ExtraRetentionSettings field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetExtraRetentionSettings() []ExtraRetentionSetting {
	if o == nil || IsNil(o.ExtraRetentionSettings) {
		var ret []ExtraRetentionSetting
		return ret
	}
	return *o.ExtraRetentionSettings
}

// GetExtraRetentionSettingsOk returns a tuple with the ExtraRetentionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetExtraRetentionSettingsOk() (*[]ExtraRetentionSetting, bool) {
	if o == nil || IsNil(o.ExtraRetentionSettings) {
		return nil, false
	}

	return o.ExtraRetentionSettings, true
}

// HasExtraRetentionSettings returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasExtraRetentionSettings() bool {
	if o != nil && !IsNil(o.ExtraRetentionSettings) {
		return true
	}

	return false
}

// SetExtraRetentionSettings gets a reference to the given []ExtraRetentionSetting and assigns it to the ExtraRetentionSettings field.
func (o *DiskBackupSnapshotSchedulePreview) SetExtraRetentionSettings(v []ExtraRetentionSetting) {
	o.ExtraRetentionSettings = &v
	o.NullFields = removeNullField(o.NullFields, "ExtraRetentionSettings")
}

// SetExtraRetentionSettingsNil sets ExtraRetentionSettings to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetExtraRetentionSettingsNil() {
	o.ExtraRetentionSettings = nil
	o.NullFields = addNullField(o.NullFields, "ExtraRetentionSettings")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *DiskBackupSnapshotSchedulePreview) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetNextSnapshot returns the NextSnapshot field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetNextSnapshot() time.Time {
	if o == nil || IsNil(o.NextSnapshot) {
		var ret time.Time
		return ret
	}
	return *o.NextSnapshot
}

// GetNextSnapshotOk returns a tuple with the NextSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetNextSnapshotOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextSnapshot) {
		return nil, false
	}

	return o.NextSnapshot, true
}

// HasNextSnapshot returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasNextSnapshot() bool {
	if o != nil && !IsNil(o.NextSnapshot) {
		return true
	}

	return false
}

// SetNextSnapshot gets a reference to the given time.Time and assigns it to the NextSnapshot field.
func (o *DiskBackupSnapshotSchedulePreview) SetNextSnapshot(v time.Time) {
	o.NextSnapshot = &v
	o.NullFields = removeNullField(o.NullFields, "NextSnapshot")
}

// SetNextSnapshotNil sets NextSnapshot to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetNextSnapshotNil() {
	o.NextSnapshot = nil
	o.NullFields = addNullField(o.NullFields, "NextSnapshot")
}

// GetPitEnabled returns the PitEnabled field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetPitEnabled() bool {
	if o == nil || IsNil(o.PitEnabled) {
		var ret bool
		return ret
	}
	return *o.PitEnabled
}

// GetPitEnabledOk returns a tuple with the PitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetPitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PitEnabled) {
		return nil, false
	}

	return o.PitEnabled, true
}

// HasPitEnabled returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasPitEnabled() bool {
	if o != nil && !IsNil(o.PitEnabled) {
		return true
	}

	return false
}

// SetPitEnabled gets a reference to the given bool and assigns it to the PitEnabled field.
func (o *DiskBackupSnapshotSchedulePreview) SetPitEnabled(v bool) {
	o.PitEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "PitEnabled")
}

// SetPitEnabledNil sets PitEnabled to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetPitEnabledNil() {
	o.PitEnabled = nil
	o.NullFields = addNullField(o.NullFields, "PitEnabled")
}

// GetPolicies returns the Policies field value
func (o *DiskBackupSnapshotSchedulePreview) GetPolicies() []AdvancedDiskBackupSnapshotSchedulePolicy {
	if o == nil {
		var ret []AdvancedDiskBackupSnapshotSchedulePolicy
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetPoliciesOk() (*[]AdvancedDiskBackupSnapshotSchedulePolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Policies, true
}

// SetPolicies sets field value
func (o *DiskBackupSnapshotSchedulePreview) SetPolicies(v []AdvancedDiskBackupSnapshotSchedulePolicy) {
	o.Policies = v
}

// GetPolicySource returns the PolicySource field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetPolicySource() string {
	if o == nil || IsNil(o.PolicySource) {
		var ret string
		return ret
	}
	return *o.PolicySource
}

// GetPolicySourceOk returns a tuple with the PolicySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetPolicySourceOk() (*string, bool) {
	if o == nil || IsNil(o.PolicySource) {
		return nil, false
	}

	return o.PolicySource, true
}

// HasPolicySource returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasPolicySource() bool {
	if o != nil && !IsNil(o.PolicySource) {
		return true
	}

	return false
}

// SetPolicySource gets a reference to the given string and assigns it to the PolicySource field.
func (o *DiskBackupSnapshotSchedulePreview) SetPolicySource(v string) {
	o.PolicySource = &v
	o.NullFields = removeNullField(o.NullFields, "PolicySource")
}

// SetPolicySourceNil sets PolicySource to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetPolicySourceNil() {
	o.PolicySource = nil
	o.NullFields = addNullField(o.NullFields, "PolicySource")
}

// GetReferenceHourOfDay returns the ReferenceHourOfDay field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetReferenceHourOfDay() int {
	if o == nil || IsNil(o.ReferenceHourOfDay) {
		var ret int
		return ret
	}
	return *o.ReferenceHourOfDay
}

// GetReferenceHourOfDayOk returns a tuple with the ReferenceHourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetReferenceHourOfDayOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceHourOfDay) {
		return nil, false
	}

	return o.ReferenceHourOfDay, true
}

// HasReferenceHourOfDay returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasReferenceHourOfDay() bool {
	if o != nil && !IsNil(o.ReferenceHourOfDay) {
		return true
	}

	return false
}

// SetReferenceHourOfDay gets a reference to the given int and assigns it to the ReferenceHourOfDay field.
func (o *DiskBackupSnapshotSchedulePreview) SetReferenceHourOfDay(v int) {
	o.ReferenceHourOfDay = &v
	o.NullFields = removeNullField(o.NullFields, "ReferenceHourOfDay")
}

// SetReferenceHourOfDayNil sets ReferenceHourOfDay to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetReferenceHourOfDayNil() {
	o.ReferenceHourOfDay = nil
	o.NullFields = addNullField(o.NullFields, "ReferenceHourOfDay")
}

// GetReferenceMinuteOfHour returns the ReferenceMinuteOfHour field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetReferenceMinuteOfHour() int {
	if o == nil || IsNil(o.ReferenceMinuteOfHour) {
		var ret int
		return ret
	}
	return *o.ReferenceMinuteOfHour
}

// GetReferenceMinuteOfHourOk returns a tuple with the ReferenceMinuteOfHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetReferenceMinuteOfHourOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceMinuteOfHour) {
		return nil, false
	}

	return o.ReferenceMinuteOfHour, true
}

// HasReferenceMinuteOfHour returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasReferenceMinuteOfHour() bool {
	if o != nil && !IsNil(o.ReferenceMinuteOfHour) {
		return true
	}

	return false
}

// SetReferenceMinuteOfHour gets a reference to the given int and assigns it to the ReferenceMinuteOfHour field.
func (o *DiskBackupSnapshotSchedulePreview) SetReferenceMinuteOfHour(v int) {
	o.ReferenceMinuteOfHour = &v
	o.NullFields = removeNullField(o.NullFields, "ReferenceMinuteOfHour")
}

// SetReferenceMinuteOfHourNil sets ReferenceMinuteOfHour to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetReferenceMinuteOfHourNil() {
	o.ReferenceMinuteOfHour = nil
	o.NullFields = addNullField(o.NullFields, "ReferenceMinuteOfHour")
}

// GetRestoreWindowDays returns the RestoreWindowDays field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetRestoreWindowDays() int {
	if o == nil || IsNil(o.RestoreWindowDays) {
		var ret int
		return ret
	}
	return *o.RestoreWindowDays
}

// GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetRestoreWindowDaysOk() (*int, bool) {
	if o == nil || IsNil(o.RestoreWindowDays) {
		return nil, false
	}

	return o.RestoreWindowDays, true
}

// HasRestoreWindowDays returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasRestoreWindowDays() bool {
	if o != nil && !IsNil(o.RestoreWindowDays) {
		return true
	}

	return false
}

// SetRestoreWindowDays gets a reference to the given int and assigns it to the RestoreWindowDays field.
func (o *DiskBackupSnapshotSchedulePreview) SetRestoreWindowDays(v int) {
	o.RestoreWindowDays = &v
	o.NullFields = removeNullField(o.NullFields, "RestoreWindowDays")
}

// SetRestoreWindowDaysNil sets RestoreWindowDays to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetRestoreWindowDaysNil() {
	o.RestoreWindowDays = nil
	o.NullFields = addNullField(o.NullFields, "RestoreWindowDays")
}

// GetUpdateCopySnapshots returns the UpdateCopySnapshots field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetUpdateCopySnapshots() bool {
	if o == nil || IsNil(o.UpdateCopySnapshots) {
		var ret bool
		return ret
	}
	return *o.UpdateCopySnapshots
}

// GetUpdateCopySnapshotsOk returns a tuple with the UpdateCopySnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetUpdateCopySnapshotsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateCopySnapshots) {
		return nil, false
	}

	return o.UpdateCopySnapshots, true
}

// HasUpdateCopySnapshots returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasUpdateCopySnapshots() bool {
	if o != nil && !IsNil(o.UpdateCopySnapshots) {
		return true
	}

	return false
}

// SetUpdateCopySnapshots gets a reference to the given bool and assigns it to the UpdateCopySnapshots field.
func (o *DiskBackupSnapshotSchedulePreview) SetUpdateCopySnapshots(v bool) {
	o.UpdateCopySnapshots = &v
	o.NullFields = removeNullField(o.NullFields, "UpdateCopySnapshots")
}

// SetUpdateCopySnapshotsNil sets UpdateCopySnapshots to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetUpdateCopySnapshotsNil() {
	o.UpdateCopySnapshots = nil
	o.NullFields = addNullField(o.NullFields, "UpdateCopySnapshots")
}

// GetUpdateSnapshots returns the UpdateSnapshots field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetUpdateSnapshots() bool {
	if o == nil || IsNil(o.UpdateSnapshots) {
		var ret bool
		return ret
	}
	return *o.UpdateSnapshots
}

// GetUpdateSnapshotsOk returns a tuple with the UpdateSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetUpdateSnapshotsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateSnapshots) {
		return nil, false
	}

	return o.UpdateSnapshots, true
}

// HasUpdateSnapshots returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasUpdateSnapshots() bool {
	if o != nil && !IsNil(o.UpdateSnapshots) {
		return true
	}

	return false
}

// SetUpdateSnapshots gets a reference to the given bool and assigns it to the UpdateSnapshots field.
func (o *DiskBackupSnapshotSchedulePreview) SetUpdateSnapshots(v bool) {
	o.UpdateSnapshots = &v
	o.NullFields = removeNullField(o.NullFields, "UpdateSnapshots")
}

// SetUpdateSnapshotsNil sets UpdateSnapshots to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetUpdateSnapshotsNil() {
	o.UpdateSnapshots = nil
	o.NullFields = addNullField(o.NullFields, "UpdateSnapshots")
}

// GetUseOrgAndGroupNamesInExportPrefix returns the UseOrgAndGroupNamesInExportPrefix field value if set, zero value otherwise
func (o *DiskBackupSnapshotSchedulePreview) GetUseOrgAndGroupNamesInExportPrefix() bool {
	if o == nil || IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		var ret bool
		return ret
	}
	return *o.UseOrgAndGroupNamesInExportPrefix
}

// GetUseOrgAndGroupNamesInExportPrefixOk returns a tuple with the UseOrgAndGroupNamesInExportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupSnapshotSchedulePreview) GetUseOrgAndGroupNamesInExportPrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		return nil, false
	}

	return o.UseOrgAndGroupNamesInExportPrefix, true
}

// HasUseOrgAndGroupNamesInExportPrefix returns a boolean if a field has been set.
func (o *DiskBackupSnapshotSchedulePreview) HasUseOrgAndGroupNamesInExportPrefix() bool {
	if o != nil && !IsNil(o.UseOrgAndGroupNamesInExportPrefix) {
		return true
	}

	return false
}

// SetUseOrgAndGroupNamesInExportPrefix gets a reference to the given bool and assigns it to the UseOrgAndGroupNamesInExportPrefix field.
func (o *DiskBackupSnapshotSchedulePreview) SetUseOrgAndGroupNamesInExportPrefix(v bool) {
	o.UseOrgAndGroupNamesInExportPrefix = &v
	o.NullFields = removeNullField(o.NullFields, "UseOrgAndGroupNamesInExportPrefix")
}

// SetUseOrgAndGroupNamesInExportPrefixNil sets UseOrgAndGroupNamesInExportPrefix to an explicit JSON null when marshaled.
func (o *DiskBackupSnapshotSchedulePreview) SetUseOrgAndGroupNamesInExportPrefixNil() {
	o.UseOrgAndGroupNamesInExportPrefix = nil
	o.NullFields = addNullField(o.NullFields, "UseOrgAndGroupNamesInExportPrefix")
}
