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

// AppServiceEventTypeViewAlertable Incident that triggered this alert.
type AppServiceEventTypeViewAlertable string

// List of AppServiceEventTypeViewAlertable
const (
	APPSERVICEEVENTTYPEVIEWALERTABLE_URL_CONFIRMATION AppServiceEventTypeViewAlertable = "URL_CONFIRMATION"
	APPSERVICEEVENTTYPEVIEWALERTABLE_SUCCESSFUL_DEPLOY AppServiceEventTypeViewAlertable = "SUCCESSFUL_DEPLOY"
	APPSERVICEEVENTTYPEVIEWALERTABLE_DEPLOYMENT_FAILURE AppServiceEventTypeViewAlertable = "DEPLOYMENT_FAILURE"
	APPSERVICEEVENTTYPEVIEWALERTABLE_DEPLOYMENT_MODEL_CHANGE_SUCCESS AppServiceEventTypeViewAlertable = "DEPLOYMENT_MODEL_CHANGE_SUCCESS"
	APPSERVICEEVENTTYPEVIEWALERTABLE_DEPLOYMENT_MODEL_CHANGE_FAILURE AppServiceEventTypeViewAlertable = "DEPLOYMENT_MODEL_CHANGE_FAILURE"
	APPSERVICEEVENTTYPEVIEWALERTABLE_REQUEST_RATE_LIMIT AppServiceEventTypeViewAlertable = "REQUEST_RATE_LIMIT"
	APPSERVICEEVENTTYPEVIEWALERTABLE_LOG_FORWARDER_FAILURE AppServiceEventTypeViewAlertable = "LOG_FORWARDER_FAILURE"
	APPSERVICEEVENTTYPEVIEWALERTABLE_OUTSIDE_REALM_METRIC_THRESHOLD AppServiceEventTypeViewAlertable = "OUTSIDE_REALM_METRIC_THRESHOLD"
	APPSERVICEEVENTTYPEVIEWALERTABLE_SYNC_FAILURE AppServiceEventTypeViewAlertable = "SYNC_FAILURE"
	APPSERVICEEVENTTYPEVIEWALERTABLE_TRIGGER_FAILURE AppServiceEventTypeViewAlertable = "TRIGGER_FAILURE"
	APPSERVICEEVENTTYPEVIEWALERTABLE_TRIGGER_AUTO_RESUMED AppServiceEventTypeViewAlertable = "TRIGGER_AUTO_RESUMED"
)

// All allowed values of AppServiceEventTypeViewAlertable enum
var AllowedAppServiceEventTypeViewAlertableEnumValues = []AppServiceEventTypeViewAlertable{
	"URL_CONFIRMATION",
	"SUCCESSFUL_DEPLOY",
	"DEPLOYMENT_FAILURE",
	"DEPLOYMENT_MODEL_CHANGE_SUCCESS",
	"DEPLOYMENT_MODEL_CHANGE_FAILURE",
	"REQUEST_RATE_LIMIT",
	"LOG_FORWARDER_FAILURE",
	"OUTSIDE_REALM_METRIC_THRESHOLD",
	"SYNC_FAILURE",
	"TRIGGER_FAILURE",
	"TRIGGER_AUTO_RESUMED",
}

func (v *AppServiceEventTypeViewAlertable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppServiceEventTypeViewAlertable(value)
	for _, existing := range AllowedAppServiceEventTypeViewAlertableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppServiceEventTypeViewAlertable", value)
}

// NewAppServiceEventTypeViewAlertableFromValue returns a pointer to a valid AppServiceEventTypeViewAlertable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppServiceEventTypeViewAlertableFromValue(v string) (*AppServiceEventTypeViewAlertable, error) {
	ev := AppServiceEventTypeViewAlertable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppServiceEventTypeViewAlertable: valid values are %v", v, AllowedAppServiceEventTypeViewAlertableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppServiceEventTypeViewAlertable) IsValid() bool {
	for _, existing := range AllowedAppServiceEventTypeViewAlertableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppServiceEventTypeViewAlertable value
func (v AppServiceEventTypeViewAlertable) Ptr() *AppServiceEventTypeViewAlertable {
	return &v
}


