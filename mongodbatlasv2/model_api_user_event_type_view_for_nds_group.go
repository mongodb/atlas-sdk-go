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

// ApiUserEventTypeViewForNdsGroup Unique identifier of event type.
type ApiUserEventTypeViewForNdsGroup string

// List of ApiUserEventTypeViewForNdsGroup
const (
	APIUSEREVENTTYPEVIEWFORNDSGROUP_CREATED ApiUserEventTypeViewForNdsGroup = "API_KEY_CREATED"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_DELETED ApiUserEventTypeViewForNdsGroup = "API_KEY_DELETED"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_ACCESS_LIST_ENTRY_ADDED ApiUserEventTypeViewForNdsGroup = "API_KEY_ACCESS_LIST_ENTRY_ADDED"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_ACCESS_LIST_ENTRY_DELETED ApiUserEventTypeViewForNdsGroup = "API_KEY_ACCESS_LIST_ENTRY_DELETED"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_ROLES_CHANGED ApiUserEventTypeViewForNdsGroup = "API_KEY_ROLES_CHANGED"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_DESCRIPTION_CHANGED ApiUserEventTypeViewForNdsGroup = "API_KEY_DESCRIPTION_CHANGED"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_ADDED_TO_GROUP ApiUserEventTypeViewForNdsGroup = "API_KEY_ADDED_TO_GROUP"
	APIUSEREVENTTYPEVIEWFORNDSGROUP_REMOVED_FROM_GROUP ApiUserEventTypeViewForNdsGroup = "API_KEY_REMOVED_FROM_GROUP"
)

// All allowed values of ApiUserEventTypeViewForNdsGroup enum
var AllowedApiUserEventTypeViewForNdsGroupEnumValues = []ApiUserEventTypeViewForNdsGroup{
	"API_KEY_CREATED",
	"API_KEY_DELETED",
	"API_KEY_ACCESS_LIST_ENTRY_ADDED",
	"API_KEY_ACCESS_LIST_ENTRY_DELETED",
	"API_KEY_ROLES_CHANGED",
	"API_KEY_DESCRIPTION_CHANGED",
	"API_KEY_ADDED_TO_GROUP",
	"API_KEY_REMOVED_FROM_GROUP",
}

func (v *ApiUserEventTypeViewForNdsGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiUserEventTypeViewForNdsGroup(value)
	for _, existing := range AllowedApiUserEventTypeViewForNdsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiUserEventTypeViewForNdsGroup", value)
}

// NewApiUserEventTypeViewForNdsGroupFromValue returns a pointer to a valid ApiUserEventTypeViewForNdsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiUserEventTypeViewForNdsGroupFromValue(v string) (*ApiUserEventTypeViewForNdsGroup, error) {
	ev := ApiUserEventTypeViewForNdsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiUserEventTypeViewForNdsGroup: valid values are %v", v, AllowedApiUserEventTypeViewForNdsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiUserEventTypeViewForNdsGroup) IsValid() bool {
	for _, existing := range AllowedApiUserEventTypeViewForNdsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiUserEventTypeViewForNdsGroup value
func (v ApiUserEventTypeViewForNdsGroup) Ptr() *ApiUserEventTypeViewForNdsGroup {
	return &v
}


