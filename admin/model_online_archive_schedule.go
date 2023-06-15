// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// OnlineArchiveSchedule - Regular frequency and duration when archiving process occurs.
type OnlineArchiveSchedule struct {
	DailySchedule   *DailySchedule
	DefaultSchedule *DefaultSchedule
	MonthlySchedule *MonthlySchedule
	WeeklySchedule  *WeeklySchedule
}

// DailyScheduleAsOnlineArchiveSchedule is a convenience function that returns DailySchedule wrapped in OnlineArchiveSchedule
func DailyScheduleAsOnlineArchiveSchedule(v *DailySchedule) OnlineArchiveSchedule {
	return OnlineArchiveSchedule{
		DailySchedule: v,
	}
}

// DefaultScheduleAsOnlineArchiveSchedule is a convenience function that returns DefaultSchedule wrapped in OnlineArchiveSchedule
func DefaultScheduleAsOnlineArchiveSchedule(v *DefaultSchedule) OnlineArchiveSchedule {
	return OnlineArchiveSchedule{
		DefaultSchedule: v,
	}
}

// MonthlyScheduleAsOnlineArchiveSchedule is a convenience function that returns MonthlySchedule wrapped in OnlineArchiveSchedule
func MonthlyScheduleAsOnlineArchiveSchedule(v *MonthlySchedule) OnlineArchiveSchedule {
	return OnlineArchiveSchedule{
		MonthlySchedule: v,
	}
}

// WeeklyScheduleAsOnlineArchiveSchedule is a convenience function that returns WeeklySchedule wrapped in OnlineArchiveSchedule
func WeeklyScheduleAsOnlineArchiveSchedule(v *WeeklySchedule) OnlineArchiveSchedule {
	return OnlineArchiveSchedule{
		WeeklySchedule: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *OnlineArchiveSchedule) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DAILY'
	if jsonDict["type"] == "DAILY" {
		// try to unmarshal JSON data into DailySchedule
		err = json.Unmarshal(data, &dst.DailySchedule)
		if err == nil {
			return nil // data stored in dst.DailySchedule, return on the first match
		} else {
			dst.DailySchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as DailySchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DEFAULT'
	if jsonDict["type"] == "DEFAULT" {
		// try to unmarshal JSON data into DefaultSchedule
		err = json.Unmarshal(data, &dst.DefaultSchedule)
		if err == nil {
			return nil // data stored in dst.DefaultSchedule, return on the first match
		} else {
			dst.DefaultSchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as DefaultSchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DailySchedule'
	if jsonDict["type"] == "DailySchedule" {
		// try to unmarshal JSON data into DailySchedule
		err = json.Unmarshal(data, &dst.DailySchedule)
		if err == nil {
			return nil // data stored in dst.DailySchedule, return on the first match
		} else {
			dst.DailySchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as DailySchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DefaultSchedule'
	if jsonDict["type"] == "DefaultSchedule" {
		// try to unmarshal JSON data into DefaultSchedule
		err = json.Unmarshal(data, &dst.DefaultSchedule)
		if err == nil {
			return nil // data stored in dst.DefaultSchedule, return on the first match
		} else {
			dst.DefaultSchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as DefaultSchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MONTHLY'
	if jsonDict["type"] == "MONTHLY" {
		// try to unmarshal JSON data into MonthlySchedule
		err = json.Unmarshal(data, &dst.MonthlySchedule)
		if err == nil {
			return nil // data stored in dst.MonthlySchedule, return on the first match
		} else {
			dst.MonthlySchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as MonthlySchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'MonthlySchedule'
	if jsonDict["type"] == "MonthlySchedule" {
		// try to unmarshal JSON data into MonthlySchedule
		err = json.Unmarshal(data, &dst.MonthlySchedule)
		if err == nil {
			return nil // data stored in dst.MonthlySchedule, return on the first match
		} else {
			dst.MonthlySchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as MonthlySchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WEEKLY'
	if jsonDict["type"] == "WEEKLY" {
		// try to unmarshal JSON data into WeeklySchedule
		err = json.Unmarshal(data, &dst.WeeklySchedule)
		if err == nil {
			return nil // data stored in dst.WeeklySchedule, return on the first match
		} else {
			dst.WeeklySchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as WeeklySchedule: %s", err.Error())
		}
	}

	// check if the discriminator value is 'WeeklySchedule'
	if jsonDict["type"] == "WeeklySchedule" {
		// try to unmarshal JSON data into WeeklySchedule
		err = json.Unmarshal(data, &dst.WeeklySchedule)
		if err == nil {
			return nil // data stored in dst.WeeklySchedule, return on the first match
		} else {
			dst.WeeklySchedule = nil
			return fmt.Errorf("failed to unmarshal OnlineArchiveSchedule as WeeklySchedule: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OnlineArchiveSchedule) MarshalJSON() ([]byte, error) {
	if src.DailySchedule != nil {
		return json.Marshal(&src.DailySchedule)
	}

	if src.DefaultSchedule != nil {
		return json.Marshal(&src.DefaultSchedule)
	}

	if src.MonthlySchedule != nil {
		return json.Marshal(&src.MonthlySchedule)
	}

	if src.WeeklySchedule != nil {
		return json.Marshal(&src.WeeklySchedule)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OnlineArchiveSchedule) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DailySchedule != nil {
		return obj.DailySchedule
	}

	if obj.DefaultSchedule != nil {
		return obj.DefaultSchedule
	}

	if obj.MonthlySchedule != nil {
		return obj.MonthlySchedule
	}

	if obj.WeeklySchedule != nil {
		return obj.WeeklySchedule
	}

	// all schemas are nil
	return nil
}

type NullableOnlineArchiveSchedule struct {
	value *OnlineArchiveSchedule
	isSet bool
}

func (v NullableOnlineArchiveSchedule) Get() *OnlineArchiveSchedule {
	return v.value
}

func (v *NullableOnlineArchiveSchedule) Set(val *OnlineArchiveSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableOnlineArchiveSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableOnlineArchiveSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnlineArchiveSchedule(val *OnlineArchiveSchedule) *NullableOnlineArchiveSchedule {
	return &NullableOnlineArchiveSchedule{value: val, isSet: true}
}

func (v NullableOnlineArchiveSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnlineArchiveSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
