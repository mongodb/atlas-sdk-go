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

// DataExplorerAccessedEventType Unique identifier of event type.
type DataExplorerAccessedEventType string

// List of DataExplorerAccessedEventType
const (
	DATAEXPLORERACCESSEDEVENTTYPE_EXPLORER DataExplorerAccessedEventType = "DATA_EXPLORER"
	DATAEXPLORERACCESSEDEVENTTYPE_EXPLORER_CRUD_ATTEMPT DataExplorerAccessedEventType = "DATA_EXPLORER_CRUD_ATTEMPT"
	DATAEXPLORERACCESSEDEVENTTYPE_EXPLORER_CRUD_ERROR DataExplorerAccessedEventType = "DATA_EXPLORER_CRUD_ERROR"
	DATAEXPLORERACCESSEDEVENTTYPE_EXPLORER_CRUD DataExplorerAccessedEventType = "DATA_EXPLORER_CRUD"
)

// All allowed values of DataExplorerAccessedEventType enum
var AllowedDataExplorerAccessedEventTypeEnumValues = []DataExplorerAccessedEventType{
	"DATA_EXPLORER",
	"DATA_EXPLORER_CRUD_ATTEMPT",
	"DATA_EXPLORER_CRUD_ERROR",
	"DATA_EXPLORER_CRUD",
}

func (v *DataExplorerAccessedEventType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataExplorerAccessedEventType(value)
	for _, existing := range AllowedDataExplorerAccessedEventTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataExplorerAccessedEventType", value)
}

// NewDataExplorerAccessedEventTypeFromValue returns a pointer to a valid DataExplorerAccessedEventType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataExplorerAccessedEventTypeFromValue(v string) (*DataExplorerAccessedEventType, error) {
	ev := DataExplorerAccessedEventType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataExplorerAccessedEventType: valid values are %v", v, AllowedDataExplorerAccessedEventTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataExplorerAccessedEventType) IsValid() bool {
	for _, existing := range AllowedDataExplorerAccessedEventTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataExplorerAccessedEventType value
func (v DataExplorerAccessedEventType) Ptr() *DataExplorerAccessedEventType {
	return &v
}


