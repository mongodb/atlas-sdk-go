// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// IngestionSink - Ingestion destination of a Data Lake Pipeline.
type IngestionSink struct {
	DLSIngestionSink *DLSIngestionSink
}

// DLSIngestionSinkAsIngestionSink is a convenience function that returns DLSIngestionSink wrapped in IngestionSink
func DLSIngestionSinkAsIngestionSink(v *DLSIngestionSink) IngestionSink {
	return IngestionSink{
		DLSIngestionSink: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IngestionSink) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DLS'
	if jsonDict["type"] == "DLS" {
		// try to unmarshal JSON data into DLSIngestionSink
		err = json.Unmarshal(data, &dst.DLSIngestionSink)
		if err == nil {
			return nil // data stored in dst.DLSIngestionSink, return on the first match
		} else {
			dst.DLSIngestionSink = nil
			return fmt.Errorf("failed to unmarshal IngestionSink as DLSIngestionSink: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DLSIngestionSink'
	if jsonDict["type"] == "DLSIngestionSink" {
		// try to unmarshal JSON data into DLSIngestionSink
		err = json.Unmarshal(data, &dst.DLSIngestionSink)
		if err == nil {
			return nil // data stored in dst.DLSIngestionSink, return on the first match
		} else {
			dst.DLSIngestionSink = nil
			return fmt.Errorf("failed to unmarshal IngestionSink as DLSIngestionSink: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IngestionSink) MarshalJSON() ([]byte, error) {
	if src.DLSIngestionSink != nil {
		return json.Marshal(&src.DLSIngestionSink)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IngestionSink) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DLSIngestionSink != nil {
		return obj.DLSIngestionSink
	}

	// all schemas are nil
	return nil
}

type NullableIngestionSink struct {
	value *IngestionSink
	isSet bool
}

func (v NullableIngestionSink) Get() *IngestionSink {
	return v.value
}

func (v *NullableIngestionSink) Set(val *IngestionSink) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestionSink) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestionSink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestionSink(val *IngestionSink) *NullableIngestionSink {
	return &NullableIngestionSink{value: val, isSet: true}
}

func (v NullableIngestionSink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestionSink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
