// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// DataMetricUnits Element used to express the quantity. This can be an element of time, storage capacity, and the like.
type DataMetricUnits string

// List of DataMetricUnits
const (
	DATAMETRICUNITS_BITS      DataMetricUnits = "BITS"
	DATAMETRICUNITS_KILOBITS  DataMetricUnits = "KILOBITS"
	DATAMETRICUNITS_MEGABITS  DataMetricUnits = "MEGABITS"
	DATAMETRICUNITS_GIGABITS  DataMetricUnits = "GIGABITS"
	DATAMETRICUNITS_BYTES     DataMetricUnits = "BYTES"
	DATAMETRICUNITS_KILOBYTES DataMetricUnits = "KILOBYTES"
	DATAMETRICUNITS_MEGABYTES DataMetricUnits = "MEGABYTES"
	DATAMETRICUNITS_GIGABYTES DataMetricUnits = "GIGABYTES"
	DATAMETRICUNITS_TERABYTES DataMetricUnits = "TERABYTES"
	DATAMETRICUNITS_PETABYTES DataMetricUnits = "PETABYTES"
)

// All allowed values of DataMetricUnits enum
var AllowedDataMetricUnitsEnumValues = []DataMetricUnits{
	"BITS",
	"KILOBITS",
	"MEGABITS",
	"GIGABITS",
	"BYTES",
	"KILOBYTES",
	"MEGABYTES",
	"GIGABYTES",
	"TERABYTES",
	"PETABYTES",
}

func (v *DataMetricUnits) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataMetricUnits(value)
	for _, existing := range AllowedDataMetricUnitsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataMetricUnits", value)
}

// NewDataMetricUnitsFromValue returns a pointer to a valid DataMetricUnits
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataMetricUnitsFromValue(v string) (*DataMetricUnits, error) {
	ev := DataMetricUnits(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataMetricUnits: valid values are %v", v, AllowedDataMetricUnitsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataMetricUnits) IsValid() bool {
	for _, existing := range AllowedDataMetricUnitsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataMetricUnits value
func (v DataMetricUnits) Ptr() *DataMetricUnits {
	return &v
}
