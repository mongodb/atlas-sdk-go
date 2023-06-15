// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// IngestionSource - Ingestion Source of a Data Lake Pipeline.
type IngestionSource struct {
	OnDemandCpsSnapshotSource *OnDemandCpsSnapshotSource
	PeriodicCpsSnapshotSource *PeriodicCpsSnapshotSource
}

// OnDemandCpsSnapshotSourceAsIngestionSource is a convenience function that returns OnDemandCpsSnapshotSource wrapped in IngestionSource
func OnDemandCpsSnapshotSourceAsIngestionSource(v *OnDemandCpsSnapshotSource) IngestionSource {
	return IngestionSource{
		OnDemandCpsSnapshotSource: v,
	}
}

// PeriodicCpsSnapshotSourceAsIngestionSource is a convenience function that returns PeriodicCpsSnapshotSource wrapped in IngestionSource
func PeriodicCpsSnapshotSourceAsIngestionSource(v *PeriodicCpsSnapshotSource) IngestionSource {
	return IngestionSource{
		PeriodicCpsSnapshotSource: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IngestionSource) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'ON_DEMAND_CPS'
	if jsonDict["type"] == "ON_DEMAND_CPS" {
		// try to unmarshal JSON data into OnDemandCpsSnapshotSource
		err = json.Unmarshal(data, &dst.OnDemandCpsSnapshotSource)
		if err == nil {
			return nil // data stored in dst.OnDemandCpsSnapshotSource, return on the first match
		} else {
			dst.OnDemandCpsSnapshotSource = nil
			return fmt.Errorf("failed to unmarshal IngestionSource as OnDemandCpsSnapshotSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OnDemandCpsSnapshotSource'
	if jsonDict["type"] == "OnDemandCpsSnapshotSource" {
		// try to unmarshal JSON data into OnDemandCpsSnapshotSource
		err = json.Unmarshal(data, &dst.OnDemandCpsSnapshotSource)
		if err == nil {
			return nil // data stored in dst.OnDemandCpsSnapshotSource, return on the first match
		} else {
			dst.OnDemandCpsSnapshotSource = nil
			return fmt.Errorf("failed to unmarshal IngestionSource as OnDemandCpsSnapshotSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PERIODIC_CPS'
	if jsonDict["type"] == "PERIODIC_CPS" {
		// try to unmarshal JSON data into PeriodicCpsSnapshotSource
		err = json.Unmarshal(data, &dst.PeriodicCpsSnapshotSource)
		if err == nil {
			return nil // data stored in dst.PeriodicCpsSnapshotSource, return on the first match
		} else {
			dst.PeriodicCpsSnapshotSource = nil
			return fmt.Errorf("failed to unmarshal IngestionSource as PeriodicCpsSnapshotSource: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PeriodicCpsSnapshotSource'
	if jsonDict["type"] == "PeriodicCpsSnapshotSource" {
		// try to unmarshal JSON data into PeriodicCpsSnapshotSource
		err = json.Unmarshal(data, &dst.PeriodicCpsSnapshotSource)
		if err == nil {
			return nil // data stored in dst.PeriodicCpsSnapshotSource, return on the first match
		} else {
			dst.PeriodicCpsSnapshotSource = nil
			return fmt.Errorf("failed to unmarshal IngestionSource as PeriodicCpsSnapshotSource: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IngestionSource) MarshalJSON() ([]byte, error) {
	if src.OnDemandCpsSnapshotSource != nil {
		return json.Marshal(&src.OnDemandCpsSnapshotSource)
	}

	if src.PeriodicCpsSnapshotSource != nil {
		return json.Marshal(&src.PeriodicCpsSnapshotSource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IngestionSource) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.OnDemandCpsSnapshotSource != nil {
		return obj.OnDemandCpsSnapshotSource
	}

	if obj.PeriodicCpsSnapshotSource != nil {
		return obj.PeriodicCpsSnapshotSource
	}

	// all schemas are nil
	return nil
}

type NullableIngestionSource struct {
	value *IngestionSource
	isSet bool
}

func (v NullableIngestionSource) Get() *IngestionSource {
	return v.value
}

func (v *NullableIngestionSource) Set(val *IngestionSource) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionSource) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionSource(val *IngestionSource) *NullableIngestionSource {
	return &NullableIngestionSource{value: val, isSet: true}
}

func (v NullableIngestionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
