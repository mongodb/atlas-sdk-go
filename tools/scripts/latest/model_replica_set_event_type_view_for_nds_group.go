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

// ReplicaSetEventTypeViewForNdsGroup Unique identifier of event type.
type ReplicaSetEventTypeViewForNdsGroup string

// List of ReplicaSetEventTypeViewForNdsGroup
const (
	REPLICASETEVENTTYPEVIEWFORNDSGROUP_PRIMARY_ELECTED ReplicaSetEventTypeViewForNdsGroup = "PRIMARY_ELECTED"
	REPLICASETEVENTTYPEVIEWFORNDSGROUP_REPLICATION_OPLOG_WINDOW_HEALTHY ReplicaSetEventTypeViewForNdsGroup = "REPLICATION_OPLOG_WINDOW_HEALTHY"
	REPLICASETEVENTTYPEVIEWFORNDSGROUP_REPLICATION_OPLOG_WINDOW_RUNNING_OUT ReplicaSetEventTypeViewForNdsGroup = "REPLICATION_OPLOG_WINDOW_RUNNING_OUT"
	REPLICASETEVENTTYPEVIEWFORNDSGROUP_ONE_PRIMARY ReplicaSetEventTypeViewForNdsGroup = "ONE_PRIMARY"
	REPLICASETEVENTTYPEVIEWFORNDSGROUP_NO_PRIMARY ReplicaSetEventTypeViewForNdsGroup = "NO_PRIMARY"
	REPLICASETEVENTTYPEVIEWFORNDSGROUP_TOO_MANY_ELECTIONS ReplicaSetEventTypeViewForNdsGroup = "TOO_MANY_ELECTIONS"
)

// All allowed values of ReplicaSetEventTypeViewForNdsGroup enum
var AllowedReplicaSetEventTypeViewForNdsGroupEnumValues = []ReplicaSetEventTypeViewForNdsGroup{
	"PRIMARY_ELECTED",
	"REPLICATION_OPLOG_WINDOW_HEALTHY",
	"REPLICATION_OPLOG_WINDOW_RUNNING_OUT",
	"ONE_PRIMARY",
	"NO_PRIMARY",
	"TOO_MANY_ELECTIONS",
}

func (v *ReplicaSetEventTypeViewForNdsGroup) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReplicaSetEventTypeViewForNdsGroup(value)
	for _, existing := range AllowedReplicaSetEventTypeViewForNdsGroupEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReplicaSetEventTypeViewForNdsGroup", value)
}

// NewReplicaSetEventTypeViewForNdsGroupFromValue returns a pointer to a valid ReplicaSetEventTypeViewForNdsGroup
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReplicaSetEventTypeViewForNdsGroupFromValue(v string) (*ReplicaSetEventTypeViewForNdsGroup, error) {
	ev := ReplicaSetEventTypeViewForNdsGroup(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReplicaSetEventTypeViewForNdsGroup: valid values are %v", v, AllowedReplicaSetEventTypeViewForNdsGroupEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReplicaSetEventTypeViewForNdsGroup) IsValid() bool {
	for _, existing := range AllowedReplicaSetEventTypeViewForNdsGroupEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReplicaSetEventTypeViewForNdsGroup value
func (v ReplicaSetEventTypeViewForNdsGroup) Ptr() *ReplicaSetEventTypeViewForNdsGroup {
	return &v
}


