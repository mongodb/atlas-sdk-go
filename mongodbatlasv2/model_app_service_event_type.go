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

// AppServiceEventType Unique identifier of event type.
type AppServiceEventType string

// List of AppServiceEventType
const (
	APPSERVICEEVENTTYPE_URL_CONFIRMATION AppServiceEventType = "URL_CONFIRMATION"
	APPSERVICEEVENTTYPE_SUCCESSFUL_DEPLOY AppServiceEventType = "SUCCESSFUL_DEPLOY"
	APPSERVICEEVENTTYPE_DEPLOYMENT_FAILURE AppServiceEventType = "DEPLOYMENT_FAILURE"
	APPSERVICEEVENTTYPE_DEPLOYMENT_MODEL_CHANGE_SUCCESS AppServiceEventType = "DEPLOYMENT_MODEL_CHANGE_SUCCESS"
	APPSERVICEEVENTTYPE_DEPLOYMENT_MODEL_CHANGE_FAILURE AppServiceEventType = "DEPLOYMENT_MODEL_CHANGE_FAILURE"
	APPSERVICEEVENTTYPE_REQUEST_RATE_LIMIT AppServiceEventType = "REQUEST_RATE_LIMIT"
	APPSERVICEEVENTTYPE_LOG_FORWARDER_FAILURE AppServiceEventType = "LOG_FORWARDER_FAILURE"
	APPSERVICEEVENTTYPE_INSIDE_REALM_METRIC_THRESHOLD AppServiceEventType = "INSIDE_REALM_METRIC_THRESHOLD"
	APPSERVICEEVENTTYPE_OUTSIDE_REALM_METRIC_THRESHOLD AppServiceEventType = "OUTSIDE_REALM_METRIC_THRESHOLD"
	APPSERVICEEVENTTYPE_SYNC_FAILURE AppServiceEventType = "SYNC_FAILURE"
	APPSERVICEEVENTTYPE_TRIGGER_FAILURE AppServiceEventType = "TRIGGER_FAILURE"
	APPSERVICEEVENTTYPE_TRIGGER_AUTO_RESUMED AppServiceEventType = "TRIGGER_AUTO_RESUMED"
)

// All allowed values of AppServiceEventType enum
var AllowedAppServiceEventTypeEnumValues = []AppServiceEventType{
	"URL_CONFIRMATION",
	"SUCCESSFUL_DEPLOY",
	"DEPLOYMENT_FAILURE",
	"DEPLOYMENT_MODEL_CHANGE_SUCCESS",
	"DEPLOYMENT_MODEL_CHANGE_FAILURE",
	"REQUEST_RATE_LIMIT",
	"LOG_FORWARDER_FAILURE",
	"INSIDE_REALM_METRIC_THRESHOLD",
	"OUTSIDE_REALM_METRIC_THRESHOLD",
	"SYNC_FAILURE",
	"TRIGGER_FAILURE",
	"TRIGGER_AUTO_RESUMED",
}

func (v *AppServiceEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppServiceEventType(value)
	for _, existing := range AllowedAppServiceEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppServiceEventType", value)
}

// NewAppServiceEventTypeFromValue returns a pointer to a valid AppServiceEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppServiceEventTypeFromValue(v string) (*AppServiceEventType, error) {
	ev := AppServiceEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppServiceEventType: valid values are %v", v, AllowedAppServiceEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppServiceEventType) IsValid() bool {
	for _, existing := range AllowedAppServiceEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppServiceEventType value
func (v AppServiceEventType) Ptr() *AppServiceEventType {
	return &v
}


