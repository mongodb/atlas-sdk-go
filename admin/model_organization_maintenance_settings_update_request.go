// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrganizationMaintenanceSettingsUpdateRequest struct for OrganizationMaintenanceSettingsUpdateRequest
type OrganizationMaintenanceSettingsUpdateRequest struct {
	// Mode configured for this organization that determines how maintenance waves are assigned to projects. Possible values are `MANUAL` and `ENV_TAG_MAPPING`. Defaults to `MANUAL` when unset. Only this field can be updated; Atlas derives read-only `effectiveWaveAssignmentMode` on GET responses and uses that value for scheduling when it differs from `waveAssignmentMode`. Omit this field to leave the current value unchanged. Specify null to reset to the default value (`MANUAL`).
	WaveAssignmentMode *string `json:"waveAssignmentMode,omitempty"`
}

// NewOrganizationMaintenanceSettingsUpdateRequest instantiates a new OrganizationMaintenanceSettingsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationMaintenanceSettingsUpdateRequest() *OrganizationMaintenanceSettingsUpdateRequest {
	this := OrganizationMaintenanceSettingsUpdateRequest{}
	return &this
}

// NewOrganizationMaintenanceSettingsUpdateRequestWithDefaults instantiates a new OrganizationMaintenanceSettingsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationMaintenanceSettingsUpdateRequestWithDefaults() *OrganizationMaintenanceSettingsUpdateRequest {
	this := OrganizationMaintenanceSettingsUpdateRequest{}
	return &this
}

// GetWaveAssignmentMode returns the WaveAssignmentMode field value if set, zero value otherwise
func (o *OrganizationMaintenanceSettingsUpdateRequest) GetWaveAssignmentMode() string {
	if o == nil || IsNil(o.WaveAssignmentMode) {
		var ret string
		return ret
	}
	return *o.WaveAssignmentMode
}

// GetWaveAssignmentModeOk returns a tuple with the WaveAssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationMaintenanceSettingsUpdateRequest) GetWaveAssignmentModeOk() (*string, bool) {
	if o == nil || IsNil(o.WaveAssignmentMode) {
		return nil, false
	}

	return o.WaveAssignmentMode, true
}

// HasWaveAssignmentMode returns a boolean if a field has been set.
func (o *OrganizationMaintenanceSettingsUpdateRequest) HasWaveAssignmentMode() bool {
	if o != nil && !IsNil(o.WaveAssignmentMode) {
		return true
	}

	return false
}

// SetWaveAssignmentMode gets a reference to the given string and assigns it to the WaveAssignmentMode field.
func (o *OrganizationMaintenanceSettingsUpdateRequest) SetWaveAssignmentMode(v string) {
	o.WaveAssignmentMode = &v
}
