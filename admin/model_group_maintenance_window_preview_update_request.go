// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupMaintenanceWindowPreviewUpdateRequest struct for GroupMaintenanceWindowPreviewUpdateRequest
type GroupMaintenanceWindowPreviewUpdateRequest struct {
	// Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them. This setting controls the same underlying auto-deferral feature as the `/maintenanceWindow/autoDefer` endpoint. Use either this field (to set a specific value) or that endpoint (to toggle the current value). For most use cases, this field in the PATCH request is preferred because it allows setting an explicit value rather than toggling.
	AutoDeferOnceEnabled *bool `json:"autoDeferOnceEnabled,omitempty"`
	// One-based integer that represents the day of the week, in the project's configured time zone (see `timeZoneId`), that the maintenance window starts.  - `1`: Sunday. - `2`: Monday. - `3`: Tuesday. - `4`: Wednesday. - `5`: Thursday. - `6`: Friday. - `7`: Saturday.
	DayOfWeek *int `json:"dayOfWeek,omitempty"`
	// Zero-based integer that represents the hour of the day, in the project's configured time zone (see `timeZoneId`), that the maintenance window starts according to a 24-hour clock. Use `0` for midnight and `12` for noon. If you haven't changed your project's time zone, this defaults to UTC.
	HourOfDay      *int            `json:"hourOfDay,omitempty"`
	ProtectedHours *ProtectedHours `json:"protectedHours,omitempty"`
	// Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request. To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to `false` after MongoDB Cloud completes maintenance.
	StartASAP *bool `json:"startASAP,omitempty"`
	// Maintenance wave assigned to this project (1–3). Not used for scheduling when the organization's `effectiveWaveAssignmentMode` is `ENV_TAG_MAPPING`. Pass `null` to clear.
	WaveAssignment *int `json:"waveAssignment,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod GroupMaintenanceWindowPreviewUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupMaintenanceWindowPreviewUpdateRequest instantiates a new GroupMaintenanceWindowPreviewUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMaintenanceWindowPreviewUpdateRequest() *GroupMaintenanceWindowPreviewUpdateRequest {
	this := GroupMaintenanceWindowPreviewUpdateRequest{}
	return &this
}

// NewGroupMaintenanceWindowPreviewUpdateRequestWithDefaults instantiates a new GroupMaintenanceWindowPreviewUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMaintenanceWindowPreviewUpdateRequestWithDefaults() *GroupMaintenanceWindowPreviewUpdateRequest {
	this := GroupMaintenanceWindowPreviewUpdateRequest{}
	return &this
}

// GetAutoDeferOnceEnabled returns the AutoDeferOnceEnabled field value if set, zero value otherwise
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetAutoDeferOnceEnabled() bool {
	if o == nil || IsNil(o.AutoDeferOnceEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoDeferOnceEnabled
}

// GetAutoDeferOnceEnabledOk returns a tuple with the AutoDeferOnceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetAutoDeferOnceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDeferOnceEnabled) {
		return nil, false
	}

	return o.AutoDeferOnceEnabled, true
}

// HasAutoDeferOnceEnabled returns a boolean if a field has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) HasAutoDeferOnceEnabled() bool {
	if o != nil && !IsNil(o.AutoDeferOnceEnabled) {
		return true
	}

	return false
}

// SetAutoDeferOnceEnabled gets a reference to the given bool and assigns it to the AutoDeferOnceEnabled field.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetAutoDeferOnceEnabled(v bool) {
	o.AutoDeferOnceEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "AutoDeferOnceEnabled")
}

// SetAutoDeferOnceEnabledNil sets AutoDeferOnceEnabled to an explicit JSON null when marshaled.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetAutoDeferOnceEnabledNil() {
	o.AutoDeferOnceEnabled = nil
	o.NullFields = addNullField(o.NullFields, "AutoDeferOnceEnabled")
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetDayOfWeek() int {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret int
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetDayOfWeekOk() (*int, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}

	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given int and assigns it to the DayOfWeek field.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetDayOfWeek(v int) {
	o.DayOfWeek = &v
	o.NullFields = removeNullField(o.NullFields, "DayOfWeek")
}

// SetDayOfWeekNil sets DayOfWeek to an explicit JSON null when marshaled.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetDayOfWeekNil() {
	o.DayOfWeek = nil
	o.NullFields = addNullField(o.NullFields, "DayOfWeek")
}

// GetHourOfDay returns the HourOfDay field value if set, zero value otherwise
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetHourOfDay() int {
	if o == nil || IsNil(o.HourOfDay) {
		var ret int
		return ret
	}
	return *o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetHourOfDayOk() (*int, bool) {
	if o == nil || IsNil(o.HourOfDay) {
		return nil, false
	}

	return o.HourOfDay, true
}

// HasHourOfDay returns a boolean if a field has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) HasHourOfDay() bool {
	if o != nil && !IsNil(o.HourOfDay) {
		return true
	}

	return false
}

// SetHourOfDay gets a reference to the given int and assigns it to the HourOfDay field.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetHourOfDay(v int) {
	o.HourOfDay = &v
	o.NullFields = removeNullField(o.NullFields, "HourOfDay")
}

// SetHourOfDayNil sets HourOfDay to an explicit JSON null when marshaled.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetHourOfDayNil() {
	o.HourOfDay = nil
	o.NullFields = addNullField(o.NullFields, "HourOfDay")
}

// GetProtectedHours returns the ProtectedHours field value if set, zero value otherwise
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetProtectedHours() ProtectedHours {
	if o == nil || IsNil(o.ProtectedHours) {
		var ret ProtectedHours
		return ret
	}
	return *o.ProtectedHours
}

// GetProtectedHoursOk returns a tuple with the ProtectedHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetProtectedHoursOk() (*ProtectedHours, bool) {
	if o == nil || IsNil(o.ProtectedHours) {
		return nil, false
	}

	return o.ProtectedHours, true
}

// HasProtectedHours returns a boolean if a field has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) HasProtectedHours() bool {
	if o != nil && !IsNil(o.ProtectedHours) {
		return true
	}

	return false
}

// SetProtectedHours gets a reference to the given ProtectedHours and assigns it to the ProtectedHours field.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetProtectedHours(v ProtectedHours) {
	o.ProtectedHours = &v
	o.NullFields = removeNullField(o.NullFields, "ProtectedHours")
}

// SetProtectedHoursNil sets ProtectedHours to an explicit JSON null when marshaled.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetProtectedHoursNil() {
	o.ProtectedHours = nil
	o.NullFields = addNullField(o.NullFields, "ProtectedHours")
}

// GetStartASAP returns the StartASAP field value if set, zero value otherwise
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetStartASAP() bool {
	if o == nil || IsNil(o.StartASAP) {
		var ret bool
		return ret
	}
	return *o.StartASAP
}

// GetStartASAPOk returns a tuple with the StartASAP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetStartASAPOk() (*bool, bool) {
	if o == nil || IsNil(o.StartASAP) {
		return nil, false
	}

	return o.StartASAP, true
}

// HasStartASAP returns a boolean if a field has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) HasStartASAP() bool {
	if o != nil && !IsNil(o.StartASAP) {
		return true
	}

	return false
}

// SetStartASAP gets a reference to the given bool and assigns it to the StartASAP field.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetStartASAP(v bool) {
	o.StartASAP = &v
	o.NullFields = removeNullField(o.NullFields, "StartASAP")
}

// SetStartASAPNil sets StartASAP to an explicit JSON null when marshaled.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetStartASAPNil() {
	o.StartASAP = nil
	o.NullFields = addNullField(o.NullFields, "StartASAP")
}

// GetWaveAssignment returns the WaveAssignment field value if set, zero value otherwise
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetWaveAssignment() int {
	if o == nil || IsNil(o.WaveAssignment) {
		var ret int
		return ret
	}
	return *o.WaveAssignment
}

// GetWaveAssignmentOk returns a tuple with the WaveAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) GetWaveAssignmentOk() (*int, bool) {
	if o == nil || IsNil(o.WaveAssignment) {
		return nil, false
	}

	return o.WaveAssignment, true
}

// HasWaveAssignment returns a boolean if a field has been set.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) HasWaveAssignment() bool {
	if o != nil && !IsNil(o.WaveAssignment) {
		return true
	}

	return false
}

// SetWaveAssignment gets a reference to the given int and assigns it to the WaveAssignment field.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetWaveAssignment(v int) {
	o.WaveAssignment = &v
	o.NullFields = removeNullField(o.NullFields, "WaveAssignment")
}

// SetWaveAssignmentNil sets WaveAssignment to an explicit JSON null when marshaled.
func (o *GroupMaintenanceWindowPreviewUpdateRequest) SetWaveAssignmentNil() {
	o.WaveAssignment = nil
	o.NullFields = addNullField(o.NullFields, "WaveAssignment")
}
