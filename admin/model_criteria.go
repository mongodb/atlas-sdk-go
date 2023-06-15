// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// Criteria - Rules by which MongoDB MongoDB Cloud archives data.  Use the **criteria.type** field to choose how MongoDB Cloud selects data to archive. Choose data using the age of the data or a MongoDB query. **\"criteria.type\": \"DATE\"** selects documents to archive based on a date. **\"criteria.type\": \"CUSTOM\"** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **\"criteria.type\": \"CUSTOM\"** when **\"collectionType\": \"TIMESERIES\"**.
type Criteria struct {
	CustomCriteria *CustomCriteria
	DateCriteria   *DateCriteria
}

// CustomCriteriaAsCriteria is a convenience function that returns CustomCriteria wrapped in Criteria
func CustomCriteriaAsCriteria(v *CustomCriteria) Criteria {
	return Criteria{
		CustomCriteria: v,
	}
}

// DateCriteriaAsCriteria is a convenience function that returns DateCriteria wrapped in Criteria
func DateCriteriaAsCriteria(v *DateCriteria) Criteria {
	return Criteria{
		DateCriteria: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Criteria) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'CUSTOM'
	if jsonDict["type"] == "CUSTOM" {
		// try to unmarshal JSON data into CustomCriteria
		err = json.Unmarshal(data, &dst.CustomCriteria)
		if err == nil {
			return nil // data stored in dst.CustomCriteria, return on the first match
		} else {
			dst.CustomCriteria = nil
			return fmt.Errorf("failed to unmarshal Criteria as CustomCriteria: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CustomCriteria'
	if jsonDict["type"] == "CustomCriteria" {
		// try to unmarshal JSON data into CustomCriteria
		err = json.Unmarshal(data, &dst.CustomCriteria)
		if err == nil {
			return nil // data stored in dst.CustomCriteria, return on the first match
		} else {
			dst.CustomCriteria = nil
			return fmt.Errorf("failed to unmarshal Criteria as CustomCriteria: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DATE'
	if jsonDict["type"] == "DATE" {
		// try to unmarshal JSON data into DateCriteria
		err = json.Unmarshal(data, &dst.DateCriteria)
		if err == nil {
			return nil // data stored in dst.DateCriteria, return on the first match
		} else {
			dst.DateCriteria = nil
			return fmt.Errorf("failed to unmarshal Criteria as DateCriteria: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DateCriteria'
	if jsonDict["type"] == "DateCriteria" {
		// try to unmarshal JSON data into DateCriteria
		err = json.Unmarshal(data, &dst.DateCriteria)
		if err == nil {
			return nil // data stored in dst.DateCriteria, return on the first match
		} else {
			dst.DateCriteria = nil
			return fmt.Errorf("failed to unmarshal Criteria as DateCriteria: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Criteria) MarshalJSON() ([]byte, error) {
	if src.CustomCriteria != nil {
		return json.Marshal(&src.CustomCriteria)
	}

	if src.DateCriteria != nil {
		return json.Marshal(&src.DateCriteria)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Criteria) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CustomCriteria != nil {
		return obj.CustomCriteria
	}

	if obj.DateCriteria != nil {
		return obj.DateCriteria
	}

	// all schemas are nil
	return nil
}

type NullableCriteria struct {
	value *Criteria
	isSet bool
}

func (v NullableCriteria) Get() *Criteria {
	return v.value
}

func (v *NullableCriteria) Set(val *Criteria) {
	v.value = val
	v.isSet = true
}

func (v NullableCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCriteria(val *Criteria) *NullableCriteria {
	return &NullableCriteria{value: val, isSet: true}
}

func (v NullableCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
