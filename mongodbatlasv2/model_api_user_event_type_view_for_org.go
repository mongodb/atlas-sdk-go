/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// ApiUserEventTypeViewForOrg Unique identifier of event type.
type ApiUserEventTypeViewForOrg string

// List of ApiUserEventTypeViewForOrg
const (
	APIUSEREVENTTYPEVIEWFORORG_CREATED ApiUserEventTypeViewForOrg = "API_KEY_CREATED"
	APIUSEREVENTTYPEVIEWFORORG_DELETED ApiUserEventTypeViewForOrg = "API_KEY_DELETED"
	APIUSEREVENTTYPEVIEWFORORG_ACCESS_LIST_ENTRY_ADDED ApiUserEventTypeViewForOrg = "API_KEY_ACCESS_LIST_ENTRY_ADDED"
	APIUSEREVENTTYPEVIEWFORORG_ACCESS_LIST_ENTRY_DELETED ApiUserEventTypeViewForOrg = "API_KEY_ACCESS_LIST_ENTRY_DELETED"
	APIUSEREVENTTYPEVIEWFORORG_ROLES_CHANGED ApiUserEventTypeViewForOrg = "API_KEY_ROLES_CHANGED"
	APIUSEREVENTTYPEVIEWFORORG_DESCRIPTION_CHANGED ApiUserEventTypeViewForOrg = "API_KEY_DESCRIPTION_CHANGED"
	APIUSEREVENTTYPEVIEWFORORG_ADDED_TO_GROUP ApiUserEventTypeViewForOrg = "API_KEY_ADDED_TO_GROUP"
	APIUSEREVENTTYPEVIEWFORORG_REMOVED_FROM_GROUP ApiUserEventTypeViewForOrg = "API_KEY_REMOVED_FROM_GROUP"
	APIUSEREVENTTYPEVIEWFORORG_UI_IP_ACCESS_LIST_INHERITANCE_ENABLED ApiUserEventTypeViewForOrg = "API_KEY_UI_IP_ACCESS_LIST_INHERITANCE_ENABLED"
	APIUSEREVENTTYPEVIEWFORORG_UI_IP_ACCESS_LIST_INHERITANCE_DISABLED ApiUserEventTypeViewForOrg = "API_KEY_UI_IP_ACCESS_LIST_INHERITANCE_DISABLED"
)

// All allowed values of ApiUserEventTypeViewForOrg enum
var AllowedApiUserEventTypeViewForOrgEnumValues = []ApiUserEventTypeViewForOrg{
	"API_KEY_CREATED",
	"API_KEY_DELETED",
	"API_KEY_ACCESS_LIST_ENTRY_ADDED",
	"API_KEY_ACCESS_LIST_ENTRY_DELETED",
	"API_KEY_ROLES_CHANGED",
	"API_KEY_DESCRIPTION_CHANGED",
	"API_KEY_ADDED_TO_GROUP",
	"API_KEY_REMOVED_FROM_GROUP",
	"API_KEY_UI_IP_ACCESS_LIST_INHERITANCE_ENABLED",
	"API_KEY_UI_IP_ACCESS_LIST_INHERITANCE_DISABLED",
}

func (v *ApiUserEventTypeViewForOrg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiUserEventTypeViewForOrg(value)
	for _, existing := range AllowedApiUserEventTypeViewForOrgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiUserEventTypeViewForOrg", value)
}

// NewApiUserEventTypeViewForOrgFromValue returns a pointer to a valid ApiUserEventTypeViewForOrg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiUserEventTypeViewForOrgFromValue(v string) (*ApiUserEventTypeViewForOrg, error) {
	ev := ApiUserEventTypeViewForOrg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiUserEventTypeViewForOrg: valid values are %v", v, AllowedApiUserEventTypeViewForOrgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiUserEventTypeViewForOrg) IsValid() bool {
	for _, existing := range AllowedApiUserEventTypeViewForOrgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiUserEventTypeViewForOrg value
func (v ApiUserEventTypeViewForOrg) Ptr() *ApiUserEventTypeViewForOrg {
	return &v
}


