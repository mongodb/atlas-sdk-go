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

// ReplicaSetEventTypeViewForNdsGroupAlertable Incident that triggered this alert.
type ReplicaSetEventTypeViewForNdsGroupAlertable string

// List of ReplicaSetEventTypeViewForNdsGroupAlertable
const (
	REPLICASETEVENTTYPEVIEWFORNDSGROUPALERTABLE_REPLICATION_OPLOG_WINDOW_RUNNING_OUT ReplicaSetEventTypeViewForNdsGroupAlertable = "REPLICATION_OPLOG_WINDOW_RUNNING_OUT"
	REPLICASETEVENTTYPEVIEWFORNDSGROUPALERTABLE_NO_PRIMARY ReplicaSetEventTypeViewForNdsGroupAlertable = "NO_PRIMARY"
	REPLICASETEVENTTYPEVIEWFORNDSGROUPALERTABLE_PRIMARY_ELECTED ReplicaSetEventTypeViewForNdsGroupAlertable = "PRIMARY_ELECTED"
	REPLICASETEVENTTYPEVIEWFORNDSGROUPALERTABLE_TOO_MANY_ELECTIONS ReplicaSetEventTypeViewForNdsGroupAlertable = "TOO_MANY_ELECTIONS"
)

// All allowed values of ReplicaSetEventTypeViewForNdsGroupAlertable enum
var AllowedReplicaSetEventTypeViewForNdsGroupAlertableEnumValues = []ReplicaSetEventTypeViewForNdsGroupAlertable{
	"REPLICATION_OPLOG_WINDOW_RUNNING_OUT",
	"NO_PRIMARY",
	"PRIMARY_ELECTED",
	"TOO_MANY_ELECTIONS",
}

func (v *ReplicaSetEventTypeViewForNdsGroupAlertable) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReplicaSetEventTypeViewForNdsGroupAlertable(value)
	for _, existing := range AllowedReplicaSetEventTypeViewForNdsGroupAlertableEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReplicaSetEventTypeViewForNdsGroupAlertable", value)
}

// NewReplicaSetEventTypeViewForNdsGroupAlertableFromValue returns a pointer to a valid ReplicaSetEventTypeViewForNdsGroupAlertable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReplicaSetEventTypeViewForNdsGroupAlertableFromValue(v string) (*ReplicaSetEventTypeViewForNdsGroupAlertable, error) {
	ev := ReplicaSetEventTypeViewForNdsGroupAlertable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReplicaSetEventTypeViewForNdsGroupAlertable: valid values are %v", v, AllowedReplicaSetEventTypeViewForNdsGroupAlertableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReplicaSetEventTypeViewForNdsGroupAlertable) IsValid() bool {
	for _, existing := range AllowedReplicaSetEventTypeViewForNdsGroupAlertableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReplicaSetEventTypeViewForNdsGroupAlertable value
func (v ReplicaSetEventTypeViewForNdsGroupAlertable) Ptr() *ReplicaSetEventTypeViewForNdsGroupAlertable {
	return &v
}


