// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AtlasGroupBackupPolicyPreview struct for AtlasGroupBackupPolicyPreview
type AtlasGroupBackupPolicyPreview struct {
	// Unique identifier of the project.
	// Read only field.
	GroupId *string `json:"groupId,omitempty"`
	// Flag that indicates whether continuous cloud backup is enabled.
	// Read only field.
	PitEnabled *bool `json:"pitEnabled,omitempty"`
	// Rules that make up the backup policy.
	// Read only field.
	Policies *[]AdvancedDiskBackupSnapshotSchedulePolicy `json:"policies,omitempty"`
	// Source of the project backup policy.
	// Read only field.
	PolicySource *string `json:"policySource,omitempty"`
	// Reference time for scheduled snapshots, in minutes after midnight UTC.
	// Read only field.
	ReferenceTimeInMins *int `json:"referenceTimeInMins,omitempty"`
	// Number of days that continuous cloud backups are retained.
	// Read only field.
	RestoreWindowDays *int `json:"restoreWindowDays,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields. It needs a value receiver: nested struct
// values aren't addressable in marshalWithNullFields.
func (o AtlasGroupBackupPolicyPreview) MarshalJSON() ([]byte, error) {
	type noMethod AtlasGroupBackupPolicyPreview
	return marshalWithNullFields(noMethod(o), o.NullFields)
}

// NewAtlasGroupBackupPolicyPreview instantiates a new AtlasGroupBackupPolicyPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtlasGroupBackupPolicyPreview() *AtlasGroupBackupPolicyPreview {
	this := AtlasGroupBackupPolicyPreview{}
	return &this
}

// NewAtlasGroupBackupPolicyPreviewWithDefaults instantiates a new AtlasGroupBackupPolicyPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtlasGroupBackupPolicyPreviewWithDefaults() *AtlasGroupBackupPolicyPreview {
	this := AtlasGroupBackupPolicyPreview{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise
func (o *AtlasGroupBackupPolicyPreview) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasGroupBackupPolicyPreview) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}

	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AtlasGroupBackupPolicyPreview) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AtlasGroupBackupPolicyPreview) SetGroupId(v string) {
	o.GroupId = &v
	o.NullFields = removeNullField(o.NullFields, "GroupId")
}

// SetGroupIdNil sets GroupId to an explicit JSON null when marshaled.
func (o *AtlasGroupBackupPolicyPreview) SetGroupIdNil() {
	o.GroupId = nil
	o.NullFields = addNullField(o.NullFields, "GroupId")
}

// GetPitEnabled returns the PitEnabled field value if set, zero value otherwise
func (o *AtlasGroupBackupPolicyPreview) GetPitEnabled() bool {
	if o == nil || IsNil(o.PitEnabled) {
		var ret bool
		return ret
	}
	return *o.PitEnabled
}

// GetPitEnabledOk returns a tuple with the PitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasGroupBackupPolicyPreview) GetPitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PitEnabled) {
		return nil, false
	}

	return o.PitEnabled, true
}

// HasPitEnabled returns a boolean if a field has been set.
func (o *AtlasGroupBackupPolicyPreview) HasPitEnabled() bool {
	if o != nil && !IsNil(o.PitEnabled) {
		return true
	}

	return false
}

// SetPitEnabled gets a reference to the given bool and assigns it to the PitEnabled field.
func (o *AtlasGroupBackupPolicyPreview) SetPitEnabled(v bool) {
	o.PitEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "PitEnabled")
}

// SetPitEnabledNil sets PitEnabled to an explicit JSON null when marshaled.
func (o *AtlasGroupBackupPolicyPreview) SetPitEnabledNil() {
	o.PitEnabled = nil
	o.NullFields = addNullField(o.NullFields, "PitEnabled")
}

// GetPolicies returns the Policies field value if set, zero value otherwise
func (o *AtlasGroupBackupPolicyPreview) GetPolicies() []AdvancedDiskBackupSnapshotSchedulePolicy {
	if o == nil || IsNil(o.Policies) {
		var ret []AdvancedDiskBackupSnapshotSchedulePolicy
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasGroupBackupPolicyPreview) GetPoliciesOk() (*[]AdvancedDiskBackupSnapshotSchedulePolicy, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}

	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AtlasGroupBackupPolicyPreview) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []AdvancedDiskBackupSnapshotSchedulePolicy and assigns it to the Policies field.
func (o *AtlasGroupBackupPolicyPreview) SetPolicies(v []AdvancedDiskBackupSnapshotSchedulePolicy) {
	o.Policies = &v
	o.NullFields = removeNullField(o.NullFields, "Policies")
}

// SetPoliciesNil sets Policies to an explicit JSON null when marshaled.
func (o *AtlasGroupBackupPolicyPreview) SetPoliciesNil() {
	o.Policies = nil
	o.NullFields = addNullField(o.NullFields, "Policies")
}

// GetPolicySource returns the PolicySource field value if set, zero value otherwise
func (o *AtlasGroupBackupPolicyPreview) GetPolicySource() string {
	if o == nil || IsNil(o.PolicySource) {
		var ret string
		return ret
	}
	return *o.PolicySource
}

// GetPolicySourceOk returns a tuple with the PolicySource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasGroupBackupPolicyPreview) GetPolicySourceOk() (*string, bool) {
	if o == nil || IsNil(o.PolicySource) {
		return nil, false
	}

	return o.PolicySource, true
}

// HasPolicySource returns a boolean if a field has been set.
func (o *AtlasGroupBackupPolicyPreview) HasPolicySource() bool {
	if o != nil && !IsNil(o.PolicySource) {
		return true
	}

	return false
}

// SetPolicySource gets a reference to the given string and assigns it to the PolicySource field.
func (o *AtlasGroupBackupPolicyPreview) SetPolicySource(v string) {
	o.PolicySource = &v
	o.NullFields = removeNullField(o.NullFields, "PolicySource")
}

// SetPolicySourceNil sets PolicySource to an explicit JSON null when marshaled.
func (o *AtlasGroupBackupPolicyPreview) SetPolicySourceNil() {
	o.PolicySource = nil
	o.NullFields = addNullField(o.NullFields, "PolicySource")
}

// GetReferenceTimeInMins returns the ReferenceTimeInMins field value if set, zero value otherwise
func (o *AtlasGroupBackupPolicyPreview) GetReferenceTimeInMins() int {
	if o == nil || IsNil(o.ReferenceTimeInMins) {
		var ret int
		return ret
	}
	return *o.ReferenceTimeInMins
}

// GetReferenceTimeInMinsOk returns a tuple with the ReferenceTimeInMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasGroupBackupPolicyPreview) GetReferenceTimeInMinsOk() (*int, bool) {
	if o == nil || IsNil(o.ReferenceTimeInMins) {
		return nil, false
	}

	return o.ReferenceTimeInMins, true
}

// HasReferenceTimeInMins returns a boolean if a field has been set.
func (o *AtlasGroupBackupPolicyPreview) HasReferenceTimeInMins() bool {
	if o != nil && !IsNil(o.ReferenceTimeInMins) {
		return true
	}

	return false
}

// SetReferenceTimeInMins gets a reference to the given int and assigns it to the ReferenceTimeInMins field.
func (o *AtlasGroupBackupPolicyPreview) SetReferenceTimeInMins(v int) {
	o.ReferenceTimeInMins = &v
	o.NullFields = removeNullField(o.NullFields, "ReferenceTimeInMins")
}

// SetReferenceTimeInMinsNil sets ReferenceTimeInMins to an explicit JSON null when marshaled.
func (o *AtlasGroupBackupPolicyPreview) SetReferenceTimeInMinsNil() {
	o.ReferenceTimeInMins = nil
	o.NullFields = addNullField(o.NullFields, "ReferenceTimeInMins")
}

// GetRestoreWindowDays returns the RestoreWindowDays field value if set, zero value otherwise
func (o *AtlasGroupBackupPolicyPreview) GetRestoreWindowDays() int {
	if o == nil || IsNil(o.RestoreWindowDays) {
		var ret int
		return ret
	}
	return *o.RestoreWindowDays
}

// GetRestoreWindowDaysOk returns a tuple with the RestoreWindowDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtlasGroupBackupPolicyPreview) GetRestoreWindowDaysOk() (*int, bool) {
	if o == nil || IsNil(o.RestoreWindowDays) {
		return nil, false
	}

	return o.RestoreWindowDays, true
}

// HasRestoreWindowDays returns a boolean if a field has been set.
func (o *AtlasGroupBackupPolicyPreview) HasRestoreWindowDays() bool {
	if o != nil && !IsNil(o.RestoreWindowDays) {
		return true
	}

	return false
}

// SetRestoreWindowDays gets a reference to the given int and assigns it to the RestoreWindowDays field.
func (o *AtlasGroupBackupPolicyPreview) SetRestoreWindowDays(v int) {
	o.RestoreWindowDays = &v
	o.NullFields = removeNullField(o.NullFields, "RestoreWindowDays")
}

// SetRestoreWindowDaysNil sets RestoreWindowDays to an explicit JSON null when marshaled.
func (o *AtlasGroupBackupPolicyPreview) SetRestoreWindowDaysNil() {
	o.RestoreWindowDays = nil
	o.NullFields = addNullField(o.NullFields, "RestoreWindowDays")
}
