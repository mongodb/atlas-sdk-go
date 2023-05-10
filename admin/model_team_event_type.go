// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// TeamEventType Unique identifier of event type.
type TeamEventType string

// List of TeamEventType
const (
	TEAMEVENTTYPE_CREATED            TeamEventType = "TEAM_CREATED"
	TEAMEVENTTYPE_DELETED            TeamEventType = "TEAM_DELETED"
	TEAMEVENTTYPE_UPDATED            TeamEventType = "TEAM_UPDATED"
	TEAMEVENTTYPE_NAME_CHANGED       TeamEventType = "TEAM_NAME_CHANGED"
	TEAMEVENTTYPE_ADDED_TO_GROUP     TeamEventType = "TEAM_ADDED_TO_GROUP"
	TEAMEVENTTYPE_REMOVED_FROM_GROUP TeamEventType = "TEAM_REMOVED_FROM_GROUP"
	TEAMEVENTTYPE_ROLES_MODIFIED     TeamEventType = "TEAM_ROLES_MODIFIED"
)

// All allowed values of TeamEventType enum
var AllowedTeamEventTypeEnumValues = []TeamEventType{
	"TEAM_CREATED",
	"TEAM_DELETED",
	"TEAM_UPDATED",
	"TEAM_NAME_CHANGED",
	"TEAM_ADDED_TO_GROUP",
	"TEAM_REMOVED_FROM_GROUP",
	"TEAM_ROLES_MODIFIED",
}

func (v *TeamEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TeamEventType(value)
	for _, existing := range AllowedTeamEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TeamEventType", value)
}

// NewTeamEventTypeFromValue returns a pointer to a valid TeamEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTeamEventTypeFromValue(v string) (*TeamEventType, error) {
	ev := TeamEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TeamEventType: valid values are %v", v, AllowedTeamEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TeamEventType) IsValid() bool {
	for _, existing := range AllowedTeamEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TeamEventType value
func (v TeamEventType) Ptr() *TeamEventType {
	return &v
}
