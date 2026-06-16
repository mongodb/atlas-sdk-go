// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrganizationMaintenanceSettingsResponse Maintenance configuration settings for an organization.
type OrganizationMaintenanceSettingsResponse struct {
	// Wave assignment mode that Atlas uses when scheduling maintenance for projects in this organization. This read-only field takes precedence over `waveAssignmentMode` for scheduling. It matches `waveAssignmentMode` except when cross-organization maintenance sequencing is enabled and this organization is a linked non-paying organization, in which case it reflects the paying organization's `waveAssignmentMode`. Possible values are `MANUAL` and `ENV_TAG_MAPPING`. Defaults to `MANUAL` when no mode is configured. Omitted from GET responses when maintenance sequencing is disabled for this organization.
	// Read only field.
	EffectiveWaveAssignmentMode *string `json:"effectiveWaveAssignmentMode,omitempty"`
	// Mode configured for this organization that determines how maintenance waves are assigned to projects. Possible values are `MANUAL` and `ENV_TAG_MAPPING`. Defaults to `MANUAL` when unset. Atlas uses read-only `effectiveWaveAssignmentMode` (not this field) for scheduling when the two values differ.
	WaveAssignmentMode *string `json:"waveAssignmentMode,omitempty"`
}

// NewOrganizationMaintenanceSettingsResponse instantiates a new OrganizationMaintenanceSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationMaintenanceSettingsResponse() *OrganizationMaintenanceSettingsResponse {
	this := OrganizationMaintenanceSettingsResponse{}
	return &this
}

// NewOrganizationMaintenanceSettingsResponseWithDefaults instantiates a new OrganizationMaintenanceSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationMaintenanceSettingsResponseWithDefaults() *OrganizationMaintenanceSettingsResponse {
	this := OrganizationMaintenanceSettingsResponse{}
	return &this
}

// GetEffectiveWaveAssignmentMode returns the EffectiveWaveAssignmentMode field value if set, zero value otherwise
func (o *OrganizationMaintenanceSettingsResponse) GetEffectiveWaveAssignmentMode() string {
	if o == nil || IsNil(o.EffectiveWaveAssignmentMode) {
		var ret string
		return ret
	}
	return *o.EffectiveWaveAssignmentMode
}

// GetEffectiveWaveAssignmentModeOk returns a tuple with the EffectiveWaveAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMaintenanceSettingsResponse) GetEffectiveWaveAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveWaveAssignmentMode) {
		return nil, false
	}

	return o.EffectiveWaveAssignmentMode, true
}

// HasEffectiveWaveAssignmentMode returns a boolean if a field has been set.
func (o *OrganizationMaintenanceSettingsResponse) HasEffectiveWaveAssignmentMode() bool {
	if o != nil && !IsNil(o.EffectiveWaveAssignmentMode) {
		return true
	}

	return false
}

// SetEffectiveWaveAssignmentMode gets a reference to the given string and assigns it to the EffectiveWaveAssignmentMode field.
func (o *OrganizationMaintenanceSettingsResponse) SetEffectiveWaveAssignmentMode(v string) {
	o.EffectiveWaveAssignmentMode = &v
}

// GetWaveAssignmentMode returns the WaveAssignmentMode field value if set, zero value otherwise
func (o *OrganizationMaintenanceSettingsResponse) GetWaveAssignmentMode() string {
	if o == nil || IsNil(o.WaveAssignmentMode) {
		var ret string
		return ret
	}
	return *o.WaveAssignmentMode
}

// GetWaveAssignmentModeOk returns a tuple with the WaveAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMaintenanceSettingsResponse) GetWaveAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.WaveAssignmentMode) {
		return nil, false
	}

	return o.WaveAssignmentMode, true
}

// HasWaveAssignmentMode returns a boolean if a field has been set.
func (o *OrganizationMaintenanceSettingsResponse) HasWaveAssignmentMode() bool {
	if o != nil && !IsNil(o.WaveAssignmentMode) {
		return true
	}

	return false
}

// SetWaveAssignmentMode gets a reference to the given string and assigns it to the WaveAssignmentMode field.
func (o *OrganizationMaintenanceSettingsResponse) SetWaveAssignmentMode(v string) {
	o.WaveAssignmentMode = &v
}
