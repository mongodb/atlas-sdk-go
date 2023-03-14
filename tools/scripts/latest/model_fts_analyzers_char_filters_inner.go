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

// FTSAnalyzersCharFiltersInner - struct for FTSAnalyzersCharFiltersInner
type FTSAnalyzersCharFiltersInner struct {
	CharFilterhtmlStrip *CharFilterhtmlStrip
	CharFiltericuNormalize *CharFiltericuNormalize
	CharFiltermapping *CharFiltermapping
	CharFilterpersian *CharFilterpersian
}

// CharFilterhtmlStripAsFTSAnalyzersCharFiltersInner is a convenience function that returns CharFilterhtmlStrip wrapped in FTSAnalyzersCharFiltersInner
func CharFilterhtmlStripAsFTSAnalyzersCharFiltersInner(v *CharFilterhtmlStrip) FTSAnalyzersCharFiltersInner {
	return FTSAnalyzersCharFiltersInner{
		CharFilterhtmlStrip: v,
	}
}

// CharFiltericuNormalizeAsFTSAnalyzersCharFiltersInner is a convenience function that returns CharFiltericuNormalize wrapped in FTSAnalyzersCharFiltersInner
func CharFiltericuNormalizeAsFTSAnalyzersCharFiltersInner(v *CharFiltericuNormalize) FTSAnalyzersCharFiltersInner {
	return FTSAnalyzersCharFiltersInner{
		CharFiltericuNormalize: v,
	}
}

// CharFiltermappingAsFTSAnalyzersCharFiltersInner is a convenience function that returns CharFiltermapping wrapped in FTSAnalyzersCharFiltersInner
func CharFiltermappingAsFTSAnalyzersCharFiltersInner(v *CharFiltermapping) FTSAnalyzersCharFiltersInner {
	return FTSAnalyzersCharFiltersInner{
		CharFiltermapping: v,
	}
}

// CharFilterpersianAsFTSAnalyzersCharFiltersInner is a convenience function that returns CharFilterpersian wrapped in FTSAnalyzersCharFiltersInner
func CharFilterpersianAsFTSAnalyzersCharFiltersInner(v *CharFilterpersian) FTSAnalyzersCharFiltersInner {
	return FTSAnalyzersCharFiltersInner{
		CharFilterpersian: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *FTSAnalyzersCharFiltersInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CharFilterhtmlStrip
	err = json.Unmarshal(data, &dst.CharFilterhtmlStrip)
	if err == nil {
		jsonCharFilterhtmlStrip, _ := json.Marshal(dst.CharFilterhtmlStrip)
		if string(jsonCharFilterhtmlStrip) == "{}" { // empty struct
			dst.CharFilterhtmlStrip = nil
		} else {
			match++
		}
	} else {
		dst.CharFilterhtmlStrip = nil
	}

	// try to unmarshal data into CharFiltericuNormalize
	err = json.Unmarshal(data, &dst.CharFiltericuNormalize)
	if err == nil {
		jsonCharFiltericuNormalize, _ := json.Marshal(dst.CharFiltericuNormalize)
		if string(jsonCharFiltericuNormalize) == "{}" { // empty struct
			dst.CharFiltericuNormalize = nil
		} else {
			match++
		}
	} else {
		dst.CharFiltericuNormalize = nil
	}

	// try to unmarshal data into CharFiltermapping
	err = json.Unmarshal(data, &dst.CharFiltermapping)
	if err == nil {
		jsonCharFiltermapping, _ := json.Marshal(dst.CharFiltermapping)
		if string(jsonCharFiltermapping) == "{}" { // empty struct
			dst.CharFiltermapping = nil
		} else {
			match++
		}
	} else {
		dst.CharFiltermapping = nil
	}

	// try to unmarshal data into CharFilterpersian
	err = json.Unmarshal(data, &dst.CharFilterpersian)
	if err == nil {
		jsonCharFilterpersian, _ := json.Marshal(dst.CharFilterpersian)
		if string(jsonCharFilterpersian) == "{}" { // empty struct
			dst.CharFilterpersian = nil
		} else {
			match++
		}
	} else {
		dst.CharFilterpersian = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CharFilterhtmlStrip = nil
		dst.CharFiltericuNormalize = nil
		dst.CharFiltermapping = nil
		dst.CharFilterpersian = nil

		return fmt.Errorf("data matches more than one schema in oneOf(FTSAnalyzersCharFiltersInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FTSAnalyzersCharFiltersInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FTSAnalyzersCharFiltersInner) MarshalJSON() ([]byte, error) {
	if src.CharFilterhtmlStrip != nil {
		return json.Marshal(&src.CharFilterhtmlStrip)
	}

	if src.CharFiltericuNormalize != nil {
		return json.Marshal(&src.CharFiltericuNormalize)
	}

	if src.CharFiltermapping != nil {
		return json.Marshal(&src.CharFiltermapping)
	}

	if src.CharFilterpersian != nil {
		return json.Marshal(&src.CharFilterpersian)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FTSAnalyzersCharFiltersInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CharFilterhtmlStrip != nil {
		return obj.CharFilterhtmlStrip
	}

	if obj.CharFiltericuNormalize != nil {
		return obj.CharFiltericuNormalize
	}

	if obj.CharFiltermapping != nil {
		return obj.CharFiltermapping
	}

	if obj.CharFilterpersian != nil {
		return obj.CharFilterpersian
	}

	// all schemas are nil
	return nil
}

type NullableFTSAnalyzersCharFiltersInner struct {
	value *FTSAnalyzersCharFiltersInner
	isSet bool
}

func (v NullableFTSAnalyzersCharFiltersInner) Get() *FTSAnalyzersCharFiltersInner {
	return v.value
}

func (v *NullableFTSAnalyzersCharFiltersInner) Set(val *FTSAnalyzersCharFiltersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFTSAnalyzersCharFiltersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFTSAnalyzersCharFiltersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFTSAnalyzersCharFiltersInner(val *FTSAnalyzersCharFiltersInner) *NullableFTSAnalyzersCharFiltersInner {
	return &NullableFTSAnalyzersCharFiltersInner{value: val, isSet: true}
}

func (v NullableFTSAnalyzersCharFiltersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFTSAnalyzersCharFiltersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


