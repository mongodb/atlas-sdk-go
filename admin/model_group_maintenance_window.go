// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupMaintenanceWindow struct for GroupMaintenanceWindow
type GroupMaintenanceWindow struct {
	// Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them.
	AutoDeferOnceEnabled *bool `json:"autoDeferOnceEnabled,omitempty"`
	// One-based integer that represents the day of the week that the maintenance window starts.  | Value | Day of Week | |---|---| | `1` | Sunday | | `2` | Monday | | `3` | Tuesday | | `4` | Wednesday | | `5` | Thursday | | `6` | Friday | | `7` | Saturday |
	DayOfWeek int `json:"dayOfWeek"`
	// Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock. Use `0` for midnight and `12` for noon.
	HourOfDay int `json:"hourOfDay"`
	// Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request. To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to `false` after MongoDB Cloud completes maintenance.
	StartASAP *bool `json:"startASAP,omitempty"`
}

// NewGroupMaintenanceWindow instantiates a new GroupMaintenanceWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMaintenanceWindow(dayOfWeek int, hourOfDay int) *GroupMaintenanceWindow {
	this := GroupMaintenanceWindow{}
	this.DayOfWeek = dayOfWeek
	this.HourOfDay = hourOfDay
	return &this
}

// NewGroupMaintenanceWindowWithDefaults instantiates a new GroupMaintenanceWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMaintenanceWindowWithDefaults() *GroupMaintenanceWindow {
	this := GroupMaintenanceWindow{}
	return &this
}

// GetAutoDeferOnceEnabled returns the AutoDeferOnceEnabled field value if set, zero value otherwise
func (o *GroupMaintenanceWindow) GetAutoDeferOnceEnabled() bool {
	if o == nil || IsNil(o.AutoDeferOnceEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoDeferOnceEnabled
}

// GetAutoDeferOnceEnabledOk returns a tuple with the AutoDeferOnceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindow) GetAutoDeferOnceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDeferOnceEnabled) {
		return nil, false
	}

	return o.AutoDeferOnceEnabled, true
}

// HasAutoDeferOnceEnabled returns a boolean if a field has been set.
func (o *GroupMaintenanceWindow) HasAutoDeferOnceEnabled() bool {
	if o != nil && !IsNil(o.AutoDeferOnceEnabled) {
		return true
	}

	return false
}

// SetAutoDeferOnceEnabled gets a reference to the given bool and assigns it to the AutoDeferOnceEnabled field.
func (o *GroupMaintenanceWindow) SetAutoDeferOnceEnabled(v bool) {
	o.AutoDeferOnceEnabled = &v
}

// GetDayOfWeek returns the DayOfWeek field value
func (o *GroupMaintenanceWindow) GetDayOfWeek() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindow) GetDayOfWeekOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayOfWeek, true
}

// SetDayOfWeek sets field value
func (o *GroupMaintenanceWindow) SetDayOfWeek(v int) {
	o.DayOfWeek = v
}

// GetHourOfDay returns the HourOfDay field value
func (o *GroupMaintenanceWindow) GetHourOfDay() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.HourOfDay
}

// GetHourOfDayOk returns a tuple with the HourOfDay field value
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindow) GetHourOfDayOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourOfDay, true
}

// SetHourOfDay sets field value
func (o *GroupMaintenanceWindow) SetHourOfDay(v int) {
	o.HourOfDay = v
}

// GetStartASAP returns the StartASAP field value if set, zero value otherwise
func (o *GroupMaintenanceWindow) GetStartASAP() bool {
	if o == nil || IsNil(o.StartASAP) {
		var ret bool
		return ret
	}
	return *o.StartASAP
}

// GetStartASAPOk returns a tuple with the StartASAP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMaintenanceWindow) GetStartASAPOk() (*bool, bool) {
	if o == nil || IsNil(o.StartASAP) {
		return nil, false
	}

	return o.StartASAP, true
}

// HasStartASAP returns a boolean if a field has been set.
func (o *GroupMaintenanceWindow) HasStartASAP() bool {
	if o != nil && !IsNil(o.StartASAP) {
		return true
	}

	return false
}

// SetStartASAP gets a reference to the given bool and assigns it to the StartASAP field.
func (o *GroupMaintenanceWindow) SetStartASAP(v bool) {
	o.StartASAP = &v
}

func (o GroupMaintenanceWindow) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupMaintenanceWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoDeferOnceEnabled) {
		toSerialize["autoDeferOnceEnabled"] = o.AutoDeferOnceEnabled
	}
	toSerialize["dayOfWeek"] = o.DayOfWeek
	toSerialize["hourOfDay"] = o.HourOfDay
	if !IsNil(o.StartASAP) {
		toSerialize["startASAP"] = o.StartASAP
	}
	return toSerialize, nil
}
