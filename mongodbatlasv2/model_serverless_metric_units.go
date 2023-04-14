/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// ServerlessMetricUnits Element used to express the quantity. This can be an element of time, storage capacity, and the like.
type ServerlessMetricUnits string

// List of ServerlessMetricUnits
const (
	SERVERLESSMETRICUNITS_RPU ServerlessMetricUnits = "RPU"
	SERVERLESSMETRICUNITS_THOUSAND_RPU ServerlessMetricUnits = "THOUSAND_RPU"
	SERVERLESSMETRICUNITS_MILLION_RPU ServerlessMetricUnits = "MILLION_RPU"
	SERVERLESSMETRICUNITS_WPU ServerlessMetricUnits = "WPU"
	SERVERLESSMETRICUNITS_THOUSAND_WPU ServerlessMetricUnits = "THOUSAND_WPU"
	SERVERLESSMETRICUNITS_MILLION_WPU ServerlessMetricUnits = "MILLION_WPU"
)

// All allowed values of ServerlessMetricUnits enum
var AllowedServerlessMetricUnitsEnumValues = []ServerlessMetricUnits{
	"RPU",
	"THOUSAND_RPU",
	"MILLION_RPU",
	"WPU",
	"THOUSAND_WPU",
	"MILLION_WPU",
}

func (v *ServerlessMetricUnits) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServerlessMetricUnits(value)
	for _, existing := range AllowedServerlessMetricUnitsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServerlessMetricUnits", value)
}

// NewServerlessMetricUnitsFromValue returns a pointer to a valid ServerlessMetricUnits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServerlessMetricUnitsFromValue(v string) (*ServerlessMetricUnits, error) {
	ev := ServerlessMetricUnits(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServerlessMetricUnits: valid values are %v", v, AllowedServerlessMetricUnitsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServerlessMetricUnits) IsValid() bool {
	for _, existing := range AllowedServerlessMetricUnitsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServerlessMetricUnits value
func (v ServerlessMetricUnits) Ptr() *ServerlessMetricUnits {
	return &v
}


