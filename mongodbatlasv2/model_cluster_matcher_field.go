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

// ClusterMatcherField Name of the parameter in the target object that MongoDB Cloud checks. The parameter must match all rules for MongoDB Cloud to check for alert configurations.
type ClusterMatcherField string

// List of ClusterMatcherField
const (
	CLUSTERMATCHERFIELD_CLUSTER_NAME ClusterMatcherField = "CLUSTER_NAME"
)

// All allowed values of ClusterMatcherField enum
var AllowedClusterMatcherFieldEnumValues = []ClusterMatcherField{
	"CLUSTER_NAME",
}

func (v *ClusterMatcherField) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterMatcherField(value)
	for _, existing := range AllowedClusterMatcherFieldEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterMatcherField", value)
}

// NewClusterMatcherFieldFromValue returns a pointer to a valid ClusterMatcherField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterMatcherFieldFromValue(v string) (*ClusterMatcherField, error) {
	ev := ClusterMatcherField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterMatcherField: valid values are %v", v, AllowedClusterMatcherFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterMatcherField) IsValid() bool {
	for _, existing := range AllowedClusterMatcherFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterMatcherField value
func (v ClusterMatcherField) Ptr() *ClusterMatcherField {
	return &v
}


