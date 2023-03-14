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

// DataLakeRegion Atlas Data Lake Regions.
type DataLakeRegion string

// List of DataLakeRegion
const (
	DATALAKEREGION_SYDNEY_AUS DataLakeRegion = "SYDNEY_AUS"
	DATALAKEREGION_MUMBAI_IND DataLakeRegion = "MUMBAI_IND"
	DATALAKEREGION_FRANKFURT_DEU DataLakeRegion = "FRANKFURT_DEU"
	DATALAKEREGION_DUBLIN_IRL DataLakeRegion = "DUBLIN_IRL"
	DATALAKEREGION_LONDON_GBR DataLakeRegion = "LONDON_GBR"
	DATALAKEREGION_VIRGINIA_USA DataLakeRegion = "VIRGINIA_USA"
	DATALAKEREGION_OREGON_USA DataLakeRegion = "OREGON_USA"
	DATALAKEREGION_SAOPAULO_BRA DataLakeRegion = "SAOPAULO_BRA"
	DATALAKEREGION_SINGAPORE_SGP DataLakeRegion = "SINGAPORE_SGP"
)

// All allowed values of DataLakeRegion enum
var AllowedDataLakeRegionEnumValues = []DataLakeRegion{
	"SYDNEY_AUS",
	"MUMBAI_IND",
	"FRANKFURT_DEU",
	"DUBLIN_IRL",
	"LONDON_GBR",
	"VIRGINIA_USA",
	"OREGON_USA",
	"SAOPAULO_BRA",
	"SINGAPORE_SGP",
}

func (v *DataLakeRegion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataLakeRegion(value)
	for _, existing := range AllowedDataLakeRegionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataLakeRegion", value)
}

// NewDataLakeRegionFromValue returns a pointer to a valid DataLakeRegion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataLakeRegionFromValue(v string) (*DataLakeRegion, error) {
	ev := DataLakeRegion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataLakeRegion: valid values are %v", v, AllowedDataLakeRegionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataLakeRegion) IsValid() bool {
	for _, existing := range AllowedDataLakeRegionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataLakeRegion value
func (v DataLakeRegion) Ptr() *DataLakeRegion {
	return &v
}


