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

// NDSAutoScalingAuditTypeViewForNdsGroup Unique identifier of event type.
type NDSAutoScalingAuditTypeViewForNdsGroup string

// List of NDSAutoScalingAuditTypeViewForNdsGroup
const (
	NDSAUTOSCALINGAUDITTYPEVIEWFORNDSGROUP_COMPUTE_AUTO_SCALE_INITIATED NDSAutoScalingAuditTypeViewForNdsGroup = "COMPUTE_AUTO_SCALE_INITIATED"
	NDSAUTOSCALINGAUDITTYPEVIEWFORNDSGROUP_DISK_AUTO_SCALE_INITIATED NDSAutoScalingAuditTypeViewForNdsGroup = "DISK_AUTO_SCALE_INITIATED"
	NDSAUTOSCALINGAUDITTYPEVIEWFORNDSGROUP_COMPUTE_AUTO_SCALE_INITIATED_BASE NDSAutoScalingAuditTypeViewForNdsGroup = "COMPUTE_AUTO_SCALE_INITIATED_BASE"
	NDSAUTOSCALINGAUDITTYPEVIEWFORNDSGROUP_COMPUTE_AUTO_SCALE_INITIATED_ANALYTICS NDSAutoScalingAuditTypeViewForNdsGroup = "COMPUTE_AUTO_SCALE_INITIATED_ANALYTICS"
)

// All allowed values of NDSAutoScalingAuditTypeViewForNdsGroup enum
var AllowedNDSAutoScalingAuditTypeViewForNdsGroupEnumValues = []NDSAutoScalingAuditTypeViewForNdsGroup{
	"COMPUTE_AUTO_SCALE_INITIATED",
	"DISK_AUTO_SCALE_INITIATED",
	"COMPUTE_AUTO_SCALE_INITIATED_BASE",
	"COMPUTE_AUTO_SCALE_INITIATED_ANALYTICS",
}

func (v *NDSAutoScalingAuditTypeViewForNdsGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NDSAutoScalingAuditTypeViewForNdsGroup(value)
	for _, existing := range AllowedNDSAutoScalingAuditTypeViewForNdsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NDSAutoScalingAuditTypeViewForNdsGroup", value)
}

// NewNDSAutoScalingAuditTypeViewForNdsGroupFromValue returns a pointer to a valid NDSAutoScalingAuditTypeViewForNdsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNDSAutoScalingAuditTypeViewForNdsGroupFromValue(v string) (*NDSAutoScalingAuditTypeViewForNdsGroup, error) {
	ev := NDSAutoScalingAuditTypeViewForNdsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NDSAutoScalingAuditTypeViewForNdsGroup: valid values are %v", v, AllowedNDSAutoScalingAuditTypeViewForNdsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NDSAutoScalingAuditTypeViewForNdsGroup) IsValid() bool {
	for _, existing := range AllowedNDSAutoScalingAuditTypeViewForNdsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NDSAutoScalingAuditTypeViewForNdsGroup value
func (v NDSAutoScalingAuditTypeViewForNdsGroup) Ptr() *NDSAutoScalingAuditTypeViewForNdsGroup {
	return &v
}


