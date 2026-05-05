// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiLifecycleManagementSchedule Schedule that controls when the lifecycle management policy runs.
type ApiLifecycleManagementSchedule struct {
	// Type of schedule.
	Type *string `json:"type,omitempty"`
	// Hour of the day when the scheduled window ends.
	EndHour *int `json:"endHour,omitempty"`
	// Minute of the hour when the scheduled window ends.
	EndMinute *int `json:"endMinute,omitempty"`
	// Hour of the day when the scheduled window starts.
	StartHour *int `json:"startHour,omitempty"`
	// Minute of the hour when the scheduled window starts.
	StartMinute *int `json:"startMinute,omitempty"`
	// Day of the week when the scheduled window starts. The week starts with Monday (`1`) and ends with Sunday (`7`).
	DayOfWeek *int `json:"dayOfWeek,omitempty"`
	// Day of the month when the scheduled window starts.
	DayOfMonth *int `json:"dayOfMonth,omitempty"`
}

// NewApiLifecycleManagementSchedule instantiates a new ApiLifecycleManagementSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLifecycleManagementSchedule() *ApiLifecycleManagementSchedule {
	this := ApiLifecycleManagementSchedule{}
	return &this
}

// NewApiLifecycleManagementScheduleWithDefaults instantiates a new ApiLifecycleManagementSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLifecycleManagementScheduleWithDefaults() *ApiLifecycleManagementSchedule {
	this := ApiLifecycleManagementSchedule{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}

	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiLifecycleManagementSchedule) SetType(v string) {
	o.Type = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetEndHour() int {
	if o == nil || IsNil(o.EndHour) {
		var ret int
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetEndHourOk() (*int, bool) {
	if o == nil || IsNil(o.EndHour) {
		return nil, false
	}

	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasEndHour() bool {
	if o != nil && !IsNil(o.EndHour) {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int and assigns it to the EndHour field.
func (o *ApiLifecycleManagementSchedule) SetEndHour(v int) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetEndMinute() int {
	if o == nil || IsNil(o.EndMinute) {
		var ret int
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetEndMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.EndMinute) {
		return nil, false
	}

	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasEndMinute() bool {
	if o != nil && !IsNil(o.EndMinute) {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int and assigns it to the EndMinute field.
func (o *ApiLifecycleManagementSchedule) SetEndMinute(v int) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetStartHour() int {
	if o == nil || IsNil(o.StartHour) {
		var ret int
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetStartHourOk() (*int, bool) {
	if o == nil || IsNil(o.StartHour) {
		return nil, false
	}

	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasStartHour() bool {
	if o != nil && !IsNil(o.StartHour) {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int and assigns it to the StartHour field.
func (o *ApiLifecycleManagementSchedule) SetStartHour(v int) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetStartMinute() int {
	if o == nil || IsNil(o.StartMinute) {
		var ret int
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetStartMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.StartMinute) {
		return nil, false
	}

	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasStartMinute() bool {
	if o != nil && !IsNil(o.StartMinute) {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int and assigns it to the StartMinute field.
func (o *ApiLifecycleManagementSchedule) SetStartMinute(v int) {
	o.StartMinute = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetDayOfWeek() int {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret int
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetDayOfWeekOk() (*int, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}

	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given int and assigns it to the DayOfWeek field.
func (o *ApiLifecycleManagementSchedule) SetDayOfWeek(v int) {
	o.DayOfWeek = &v
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise
func (o *ApiLifecycleManagementSchedule) GetDayOfMonth() int {
	if o == nil || IsNil(o.DayOfMonth) {
		var ret int
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLifecycleManagementSchedule) GetDayOfMonthOk() (*int, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}

	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *ApiLifecycleManagementSchedule) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int and assigns it to the DayOfMonth field.
func (o *ApiLifecycleManagementSchedule) SetDayOfMonth(v int) {
	o.DayOfMonth = &v
}
