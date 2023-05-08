/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// HostMatcherField Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations.
type HostMatcherField string

// List of HostMatcherField
const (
	HOSTMATCHERFIELD_TYPE_NAME HostMatcherField = "TYPE_NAME"
	HOSTMATCHERFIELD_HOSTNAME HostMatcherField = "HOSTNAME"
	HOSTMATCHERFIELD_PORT HostMatcherField = "PORT"
	HOSTMATCHERFIELD_HOSTNAME_AND_PORT HostMatcherField = "HOSTNAME_AND_PORT"
	HOSTMATCHERFIELD_REPLICA_SET_NAME HostMatcherField = "REPLICA_SET_NAME"
)

// All allowed values of HostMatcherField enum
var AllowedHostMatcherFieldEnumValues = []HostMatcherField{
	"TYPE_NAME",
	"HOSTNAME",
	"PORT",
	"HOSTNAME_AND_PORT",
	"REPLICA_SET_NAME",
}

func (v *HostMatcherField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HostMatcherField(value)
	for _, existing := range AllowedHostMatcherFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HostMatcherField", value)
}

// NewHostMatcherFieldFromValue returns a pointer to a valid HostMatcherField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHostMatcherFieldFromValue(v string) (*HostMatcherField, error) {
	ev := HostMatcherField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HostMatcherField: valid values are %v", v, AllowedHostMatcherFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HostMatcherField) IsValid() bool {
	for _, existing := range AllowedHostMatcherFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HostMatcherField value
func (v HostMatcherField) Ptr() *HostMatcherField {
	return &v
}


