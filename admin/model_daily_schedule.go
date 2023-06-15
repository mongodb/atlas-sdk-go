// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the DailySchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DailySchedule{}

// DailySchedule struct for DailySchedule
type DailySchedule struct {
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

// NewDailySchedule instantiates a new DailySchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDailySchedule(type_ string) *DailySchedule {
	this := DailySchedule{}
	this.Type = type_
	return &this
}

// NewDailyScheduleWithDefaults instantiates a new DailySchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDailyScheduleWithDefaults() *DailySchedule {
	this := DailySchedule{}
	return &this
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *DailySchedule) GetEndHour() int {
	if o == nil || IsNil(o.EndHour) {
		var ret int
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySchedule) GetEndHourOk() (*int, bool) {
	if o == nil || IsNil(o.EndHour) {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *DailySchedule) HasEndHour() bool {
	if o != nil && !IsNil(o.EndHour) {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given int and assigns it to the EndHour field.
func (o *DailySchedule) SetEndHour(v int) {
	o.EndHour = &v
}

// GetEndMinute returns the EndMinute field value if set, zero value otherwise.
func (o *DailySchedule) GetEndMinute() int {
	if o == nil || IsNil(o.EndMinute) {
		var ret int
		return ret
	}
	return *o.EndMinute
}

// GetEndMinuteOk returns a tuple with the EndMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySchedule) GetEndMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.EndMinute) {
		return nil, false
	}
	return o.EndMinute, true
}

// HasEndMinute returns a boolean if a field has been set.
func (o *DailySchedule) HasEndMinute() bool {
	if o != nil && !IsNil(o.EndMinute) {
		return true
	}

	return false
}

// SetEndMinute gets a reference to the given int and assigns it to the EndMinute field.
func (o *DailySchedule) SetEndMinute(v int) {
	o.EndMinute = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *DailySchedule) GetStartHour() int {
	if o == nil || IsNil(o.StartHour) {
		var ret int
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySchedule) GetStartHourOk() (*int, bool) {
	if o == nil || IsNil(o.StartHour) {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *DailySchedule) HasStartHour() bool {
	if o != nil && !IsNil(o.StartHour) {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given int and assigns it to the StartHour field.
func (o *DailySchedule) SetStartHour(v int) {
	o.StartHour = &v
}

// GetStartMinute returns the StartMinute field value if set, zero value otherwise.
func (o *DailySchedule) GetStartMinute() int {
	if o == nil || IsNil(o.StartMinute) {
		var ret int
		return ret
	}
	return *o.StartMinute
}

// GetStartMinuteOk returns a tuple with the StartMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DailySchedule) GetStartMinuteOk() (*int, bool) {
	if o == nil || IsNil(o.StartMinute) {
		return nil, false
	}
	return o.StartMinute, true
}

// HasStartMinute returns a boolean if a field has been set.
func (o *DailySchedule) HasStartMinute() bool {
	if o != nil && !IsNil(o.StartMinute) {
		return true
	}

	return false
}

// SetStartMinute gets a reference to the given int and assigns it to the StartMinute field.
func (o *DailySchedule) SetStartMinute(v int) {
	o.StartMinute = &v
}

// GetType returns the Type field value
func (o *DailySchedule) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DailySchedule) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DailySchedule) SetType(v string) {
	o.Type = v
}

func (o DailySchedule) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o DailySchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableDailySchedule struct {
	value *DailySchedule
	isSet bool
}

func (v NullableDailySchedule) Get() *DailySchedule {
	return v.value
}

func (v *NullableDailySchedule) Set(val *DailySchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableDailySchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableDailySchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDailySchedule(val *DailySchedule) *NullableDailySchedule {
	return &NullableDailySchedule{value: val, isSet: true}
}

func (v NullableDailySchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDailySchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
