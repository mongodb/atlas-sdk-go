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

// AutomationConfigEventType Unique identifier of event type.
type AutomationConfigEventType string

// List of AutomationConfigEventType
const (
	AUTOMATIONCONFIGEVENTTYPE_AUTOMATION_CONFIG_PUBLISHED_AUDIT AutomationConfigEventType = "AUTOMATION_CONFIG_PUBLISHED_AUDIT"
)

// All allowed values of AutomationConfigEventType enum
var AllowedAutomationConfigEventTypeEnumValues = []AutomationConfigEventType{
	"AUTOMATION_CONFIG_PUBLISHED_AUDIT",
}

func (v *AutomationConfigEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AutomationConfigEventType(value)
	for _, existing := range AllowedAutomationConfigEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AutomationConfigEventType", value)
}

// NewAutomationConfigEventTypeFromValue returns a pointer to a valid AutomationConfigEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAutomationConfigEventTypeFromValue(v string) (*AutomationConfigEventType, error) {
	ev := AutomationConfigEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AutomationConfigEventType: valid values are %v", v, AllowedAutomationConfigEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AutomationConfigEventType) IsValid() bool {
	for _, existing := range AllowedAutomationConfigEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AutomationConfigEventType value
func (v AutomationConfigEventType) Ptr() *AutomationConfigEventType {
	return &v
}


