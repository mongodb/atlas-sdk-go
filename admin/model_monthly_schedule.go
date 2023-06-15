// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the MonthlySchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonthlySchedule{}

// MonthlySchedule struct for MonthlySchedule
type MonthlySchedule struct {
	// Day of the month when the scheduled archive starts.
	DayOfMonth *int `json:"dayOfMonth,omitempty"`
	// Hour of the day when the scheduled window to run one online archive ends.
	EndHour *int `json:"endHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive ends.
	EndMinute *int `json:"endMinute,omitempty"`
	// Hour of the day when the when the scheduled window to run one online archive starts.
	StartHour *int `json:"startHour,omitempty"`
	// Minute of the hour when the scheduled window to run one online archive starts.
	StartMinute *int   `json:"startMinute,omitempty"`
	Type        string `json:"type"`
}

// NewMonthlySchedule instantiates a new MonthlySchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonthlySchedule(type_ string) *MonthlySchedule {
	this := MonthlySchedule{}
	this.Type = type_
	return &this
}

// NewMonthlyScheduleWithDefaults instantiates a new MonthlySchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonthlyScheduleWithDefaults() *MonthlySchedule {
	this := MonthlySchedule{}
	return &this
}

// GetDayOfMonth returns the DayOfMonth field value if set, zero value otherwise.
func (o *MonthlySchedule) GetDayOfMonth() int {
	if o == nil || IsNil(o.DayOfMonth) {
		var ret int
		return ret
	}
	return *o.DayOfMonth
}

// GetDayOfMonthOk returns a tuple with the DayOfMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlySchedule) GetDayOfMonthOk() (*int, bool) {
	if o == nil || IsNil(o.DayOfMonth) {
		return nil, false
	}
	return o.DayOfMonth, true
}

// HasDayOfMonth returns a boolean if a field has been set.
func (o *MonthlySchedule) HasDayOfMonth() bool {
	if o != nil && !IsNil(o.DayOfMonth) {
		return true
	}

	return false
}

// SetDayOfMonth gets a reference to the given int and assigns it to the DayOfMonth field.
func (o *MonthlySchedule) SetDayOfMonth(v int) {
	o.DayOfMonth = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *MonthlySchedule) GetEndHour() int {
	if o == nil || IsNil(o.EndHour) {
		var ret int
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlySchedule) GetEndHourOk() (*int, bool) {
	if o == nil || IsNil(o.EndHour) {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *MonthlySchedule) HasEndHour() bool {
	if o != nil && !IsNil(o.EndHour) {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int and assigns it to the EndHour field.
func (o *MonthlySchedule) SetEndHour(v int) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *MonthlySchedule) GetEndMinute() int {
	if o == nil || IsNil(o.EndMinute) {
		var ret int
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlySchedule) GetEndMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.EndMinute) {
		return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *MonthlySchedule) HasEndMinute() bool {
	if o != nil && !IsNil(o.EndMinute) {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int and assigns it to the EndMinute field.
func (o *MonthlySchedule) SetEndMinute(v int) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *MonthlySchedule) GetStartHour() int {
	if o == nil || IsNil(o.StartHour) {
		var ret int
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlySchedule) GetStartHourOk() (*int, bool) {
	if o == nil || IsNil(o.StartHour) {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *MonthlySchedule) HasStartHour() bool {
	if o != nil && !IsNil(o.StartHour) {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int and assigns it to the StartHour field.
func (o *MonthlySchedule) SetStartHour(v int) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *MonthlySchedule) GetStartMinute() int {
	if o == nil || IsNil(o.StartMinute) {
		var ret int
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlySchedule) GetStartMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.StartMinute) {
		return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *MonthlySchedule) HasStartMinute() bool {
	if o != nil && !IsNil(o.StartMinute) {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int and assigns it to the StartMinute field.
func (o *MonthlySchedule) SetStartMinute(v int) {
	o.StartMinute = &v
}

// GetType returns the Type field value
func (o *MonthlySchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MonthlySchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MonthlySchedule) SetType(v string) {
	o.Type = v
}

func (o MonthlySchedule) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o MonthlySchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DayOfMonth) {
		toSerialize["dayOfMonth"] = o.DayOfMonth
	}
	if !IsNil(o.EndHour) {
		toSerialize["endHour"] = o.EndHour
	}
	if !IsNil(o.EndMinute) {
		toSerialize["endMinute"] = o.EndMinute
	}
	if !IsNil(o.StartHour) {
		toSerialize["startHour"] = o.StartHour
	}
	if !IsNil(o.StartMinute) {
		toSerialize["startMinute"] = o.StartMinute
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMonthlySchedule struct {
	value *MonthlySchedule
	isSet bool
}

func (v NullableMonthlySchedule) Get() *MonthlySchedule {
	return v.value
}

func (v *NullableMonthlySchedule) Set(val *MonthlySchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableMonthlySchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableMonthlySchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonthlySchedule(val *MonthlySchedule) *NullableMonthlySchedule {
	return &NullableMonthlySchedule{value: val, isSet: true}
}

func (v NullableMonthlySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonthlySchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
