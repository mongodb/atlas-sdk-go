// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// THIS TEMPLATE IS NOT BEING TRIGGERED BY THE CURRENT MODEL

// GroupAlertEventTypeName - Incident that triggered this alert.
type GroupAlertEventTypeName struct {
	String *string
}

// stringAsGroupAlertEventTypeName is a convenience function that returns string wrapped in GroupAlertEventTypeName
func StringAsGroupAlertEventTypeName(v *string) GroupAlertEventTypeName {
	return GroupAlertEventTypeName{
		String: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GroupAlertEventTypeName) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.String)
		if string(jsonstring) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GroupAlertEventTypeName)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GroupAlertEventTypeName)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GroupAlertEventTypeName) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GroupAlertEventTypeName) GetActualInstance() any {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

// THIS TEMPLATE IS NOT BEING TRIGGERED BY THE CURRENT MODEL
