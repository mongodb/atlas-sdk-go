/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// HostMetricEventTypeViewAlertable Event type that triggers an alert.
type HostMetricEventTypeViewAlertable string

// List of HostMetricEventTypeViewAlertable
const (
	HOSTMETRICEVENTTYPEVIEWALERTABLE_OUTSIDE_METRIC_THRESHOLD HostMetricEventTypeViewAlertable = "OUTSIDE_METRIC_THRESHOLD"
)

// All allowed values of HostMetricEventTypeViewAlertable enum
var AllowedHostMetricEventTypeViewAlertableEnumValues = []HostMetricEventTypeViewAlertable{
	"OUTSIDE_METRIC_THRESHOLD",
}

func (v *HostMetricEventTypeViewAlertable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HostMetricEventTypeViewAlertable(value)
	for _, existing := range AllowedHostMetricEventTypeViewAlertableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HostMetricEventTypeViewAlertable", value)
}

// NewHostMetricEventTypeViewAlertableFromValue returns a pointer to a valid HostMetricEventTypeViewAlertable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostMetricEventTypeViewAlertableFromValue(v string) (*HostMetricEventTypeViewAlertable, error) {
	ev := HostMetricEventTypeViewAlertable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostMetricEventTypeViewAlertable: valid values are %v", v, AllowedHostMetricEventTypeViewAlertableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostMetricEventTypeViewAlertable) IsValid() bool {
	for _, existing := range AllowedHostMetricEventTypeViewAlertableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMetricEventTypeViewAlertable value
func (v HostMetricEventTypeViewAlertable) Ptr() *HostMetricEventTypeViewAlertable {
	return &v
}


