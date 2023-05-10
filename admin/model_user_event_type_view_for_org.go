// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// UserEventTypeViewForOrg Unique identifier of event type.
type UserEventTypeViewForOrg string

// List of UserEventTypeViewForOrg
const (
	USEREVENTTYPEVIEWFORORG_JOINED_ORG                          UserEventTypeViewForOrg = "JOINED_ORG"
	USEREVENTTYPEVIEWFORORG_JOINED_TEAM                         UserEventTypeViewForOrg = "JOINED_TEAM"
	USEREVENTTYPEVIEWFORORG_INVITED_TO_ORG                      UserEventTypeViewForOrg = "INVITED_TO_ORG"
	USEREVENTTYPEVIEWFORORG_ORG_INVITATION_DELETED              UserEventTypeViewForOrg = "ORG_INVITATION_DELETED"
	USEREVENTTYPEVIEWFORORG_REMOVED_FROM_ORG                    UserEventTypeViewForOrg = "REMOVED_FROM_ORG"
	USEREVENTTYPEVIEWFORORG_REMOVED_FROM_TEAM                   UserEventTypeViewForOrg = "REMOVED_FROM_TEAM"
	USEREVENTTYPEVIEWFORORG_USER_ROLES_CHANGED_AUDIT            UserEventTypeViewForOrg = "USER_ROLES_CHANGED_AUDIT"
	USEREVENTTYPEVIEWFORORG_ORG_FLEX_CONSULTING_PURCHASED       UserEventTypeViewForOrg = "ORG_FLEX_CONSULTING_PURCHASED"
	USEREVENTTYPEVIEWFORORG_ORG_FLEX_CONSULTING_PURCHASE_FAILED UserEventTypeViewForOrg = "ORG_FLEX_CONSULTING_PURCHASE_FAILED"
)

// All allowed values of UserEventTypeViewForOrg enum
var AllowedUserEventTypeViewForOrgEnumValues = []UserEventTypeViewForOrg{
	"JOINED_ORG",
	"JOINED_TEAM",
	"INVITED_TO_ORG",
	"ORG_INVITATION_DELETED",
	"REMOVED_FROM_ORG",
	"REMOVED_FROM_TEAM",
	"USER_ROLES_CHANGED_AUDIT",
	"ORG_FLEX_CONSULTING_PURCHASED",
	"ORG_FLEX_CONSULTING_PURCHASE_FAILED",
}

func (v *UserEventTypeViewForOrg) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserEventTypeViewForOrg(value)
	for _, existing := range AllowedUserEventTypeViewForOrgEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserEventTypeViewForOrg", value)
}

// NewUserEventTypeViewForOrgFromValue returns a pointer to a valid UserEventTypeViewForOrg
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserEventTypeViewForOrgFromValue(v string) (*UserEventTypeViewForOrg, error) {
	ev := UserEventTypeViewForOrg(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserEventTypeViewForOrg: valid values are %v", v, AllowedUserEventTypeViewForOrgEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserEventTypeViewForOrg) IsValid() bool {
	for _, existing := range AllowedUserEventTypeViewForOrgEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserEventTypeViewForOrg value
func (v UserEventTypeViewForOrg) Ptr() *UserEventTypeViewForOrg {
	return &v
}
