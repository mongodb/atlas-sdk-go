// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// HostEventTypeViewForNdsGroupAlertable Event type that triggers an alert.
type HostEventTypeViewForNdsGroupAlertable string

// List of HostEventTypeViewForNdsGroupAlertable
const (
	HOSTEVENTTYPEVIEWFORNDSGROUPALERTABLE_DOWN                                             HostEventTypeViewForNdsGroupAlertable = "HOST_DOWN"
	HOSTEVENTTYPEVIEWFORNDSGROUPALERTABLE_HAS_INDEX_SUGGESTIONS                            HostEventTypeViewForNdsGroupAlertable = "HOST_HAS_INDEX_SUGGESTIONS"
	HOSTEVENTTYPEVIEWFORNDSGROUPALERTABLE_MONGOT_CRASHING_OOM                              HostEventTypeViewForNdsGroupAlertable = "HOST_MONGOT_CRASHING_OOM"
	HOSTEVENTTYPEVIEWFORNDSGROUPALERTABLE_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD HostEventTypeViewForNdsGroupAlertable = "HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD"
	HOSTEVENTTYPEVIEWFORNDSGROUPALERTABLE_NOT_ENOUGH_DISK_SPACE                            HostEventTypeViewForNdsGroupAlertable = "HOST_NOT_ENOUGH_DISK_SPACE"
)

// All allowed values of HostEventTypeViewForNdsGroupAlertable enum
var AllowedHostEventTypeViewForNdsGroupAlertableEnumValues = []HostEventTypeViewForNdsGroupAlertable{
	"HOST_DOWN",
	"HOST_HAS_INDEX_SUGGESTIONS",
	"HOST_MONGOT_CRASHING_OOM",
	"HOST_DISK_SPACE_INSUFFICIENT_FOR_SEARCH_INDEX_REBUILD",
	"HOST_NOT_ENOUGH_DISK_SPACE",
}

func (v *HostEventTypeViewForNdsGroupAlertable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HostEventTypeViewForNdsGroupAlertable(value)
	for _, existing := range AllowedHostEventTypeViewForNdsGroupAlertableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HostEventTypeViewForNdsGroupAlertable", value)
}

// NewHostEventTypeViewForNdsGroupAlertableFromValue returns a pointer to a valid HostEventTypeViewForNdsGroupAlertable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostEventTypeViewForNdsGroupAlertableFromValue(v string) (*HostEventTypeViewForNdsGroupAlertable, error) {
	ev := HostEventTypeViewForNdsGroupAlertable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostEventTypeViewForNdsGroupAlertable: valid values are %v", v, AllowedHostEventTypeViewForNdsGroupAlertableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostEventTypeViewForNdsGroupAlertable) IsValid() bool {
	for _, existing := range AllowedHostEventTypeViewForNdsGroupAlertableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostEventTypeViewForNdsGroupAlertable value
func (v HostEventTypeViewForNdsGroupAlertable) Ptr() *HostEventTypeViewForNdsGroupAlertable {
	return &v
}
