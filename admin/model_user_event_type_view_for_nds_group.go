// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// UserEventTypeViewForNdsGroup Unique identifier of event type.
type UserEventTypeViewForNdsGroup string

// List of UserEventTypeViewForNdsGroup
const (
	USEREVENTTYPEVIEWFORNDSGROUP_JOINED_GROUP                      UserEventTypeViewForNdsGroup = "JOINED_GROUP"
	USEREVENTTYPEVIEWFORNDSGROUP_REMOVED_FROM_GROUP                UserEventTypeViewForNdsGroup = "REMOVED_FROM_GROUP"
	USEREVENTTYPEVIEWFORNDSGROUP_INVITED_TO_GROUP                  UserEventTypeViewForNdsGroup = "INVITED_TO_GROUP"
	USEREVENTTYPEVIEWFORNDSGROUP_REQUESTED_TO_JOIN_GROUP           UserEventTypeViewForNdsGroup = "REQUESTED_TO_JOIN_GROUP"
	USEREVENTTYPEVIEWFORNDSGROUP_GROUP_INVITATION_DELETED          UserEventTypeViewForNdsGroup = "GROUP_INVITATION_DELETED"
	USEREVENTTYPEVIEWFORNDSGROUP_USER_ROLES_CHANGED_AUDIT          UserEventTypeViewForNdsGroup = "USER_ROLES_CHANGED_AUDIT"
	USEREVENTTYPEVIEWFORNDSGROUP_JOIN_GROUP_REQUEST_DENIED_AUDIT   UserEventTypeViewForNdsGroup = "JOIN_GROUP_REQUEST_DENIED_AUDIT"
	USEREVENTTYPEVIEWFORNDSGROUP_JOIN_GROUP_REQUEST_APPROVED_AUDIT UserEventTypeViewForNdsGroup = "JOIN_GROUP_REQUEST_APPROVED_AUDIT"
)

// All allowed values of UserEventTypeViewForNdsGroup enum
var AllowedUserEventTypeViewForNdsGroupEnumValues = []UserEventTypeViewForNdsGroup{
	"JOINED_GROUP",
	"REMOVED_FROM_GROUP",
	"INVITED_TO_GROUP",
	"REQUESTED_TO_JOIN_GROUP",
	"GROUP_INVITATION_DELETED",
	"USER_ROLES_CHANGED_AUDIT",
	"JOIN_GROUP_REQUEST_DENIED_AUDIT",
	"JOIN_GROUP_REQUEST_APPROVED_AUDIT",
}

func (v *UserEventTypeViewForNdsGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserEventTypeViewForNdsGroup(value)
	for _, existing := range AllowedUserEventTypeViewForNdsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserEventTypeViewForNdsGroup", value)
}

// NewUserEventTypeViewForNdsGroupFromValue returns a pointer to a valid UserEventTypeViewForNdsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserEventTypeViewForNdsGroupFromValue(v string) (*UserEventTypeViewForNdsGroup, error) {
	ev := UserEventTypeViewForNdsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserEventTypeViewForNdsGroup: valid values are %v", v, AllowedUserEventTypeViewForNdsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserEventTypeViewForNdsGroup) IsValid() bool {
	for _, existing := range AllowedUserEventTypeViewForNdsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserEventTypeViewForNdsGroup value
func (v UserEventTypeViewForNdsGroup) Ptr() *UserEventTypeViewForNdsGroup {
	return &v
}
