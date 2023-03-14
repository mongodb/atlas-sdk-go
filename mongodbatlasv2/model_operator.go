/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
	"fmt"
)

// Operator Comparison operator to apply when checking the current metric value.
type Operator string

// List of Operator
const (
	OPERATOR_LESS_THAN Operator = "<"
	OPERATOR_GREATER_THAN Operator = ">"
)

// All allowed values of Operator enum
var AllowedOperatorEnumValues = []Operator{
	"<",
	">",
}

func (v *Operator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Operator(value)
	for _, existing := range AllowedOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Operator", value)
}

// NewOperatorFromValue returns a pointer to a valid Operator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatorFromValue(v string) (*Operator, error) {
	ev := Operator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Operator: valid values are %v", v, AllowedOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Operator) IsValid() bool {
	for _, existing := range AllowedOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Operator value
func (v Operator) Ptr() *Operator {
	return &v
}


