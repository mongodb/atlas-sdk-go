/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
	"fmt"
)

// AppServiceEventTypeViewAlertableWithThreshold Event type that triggers an alert.
type AppServiceEventTypeViewAlertableWithThreshold string

// List of AppServiceEventTypeViewAlertableWithThreshold
const (
	APPSERVICEEVENTTYPEVIEWALERTABLEWITHTHRESHOLD_OUTSIDE_REALM_METRIC_THRESHOLD AppServiceEventTypeViewAlertableWithThreshold = "OUTSIDE_REALM_METRIC_THRESHOLD"
)

// All allowed values of AppServiceEventTypeViewAlertableWithThreshold enum
var AllowedAppServiceEventTypeViewAlertableWithThresholdEnumValues = []AppServiceEventTypeViewAlertableWithThreshold{
	"OUTSIDE_REALM_METRIC_THRESHOLD",
}

func (v *AppServiceEventTypeViewAlertableWithThreshold) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppServiceEventTypeViewAlertableWithThreshold(value)
	for _, existing := range AllowedAppServiceEventTypeViewAlertableWithThresholdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppServiceEventTypeViewAlertableWithThreshold", value)
}

// NewAppServiceEventTypeViewAlertableWithThresholdFromValue returns a pointer to a valid AppServiceEventTypeViewAlertableWithThreshold
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppServiceEventTypeViewAlertableWithThresholdFromValue(v string) (*AppServiceEventTypeViewAlertableWithThreshold, error) {
	ev := AppServiceEventTypeViewAlertableWithThreshold(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppServiceEventTypeViewAlertableWithThreshold: valid values are %v", v, AllowedAppServiceEventTypeViewAlertableWithThresholdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppServiceEventTypeViewAlertableWithThreshold) IsValid() bool {
	for _, existing := range AllowedAppServiceEventTypeViewAlertableWithThresholdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppServiceEventTypeViewAlertableWithThreshold value
func (v AppServiceEventTypeViewAlertableWithThreshold) Ptr() *AppServiceEventTypeViewAlertableWithThreshold {
	return &v
}


