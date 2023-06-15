// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// DataFederationLimit - Details of user managed limits.
type DataFederationLimit struct {
	DataFederationQueryLimit *DataFederationQueryLimit
	DefaultLimit             *DefaultLimit
}

// DataFederationQueryLimitAsDataFederationLimit is a convenience function that returns DataFederationQueryLimit wrapped in DataFederationLimit
func DataFederationQueryLimitAsDataFederationLimit(v *DataFederationQueryLimit) DataFederationLimit {
	return DataFederationLimit{
		DataFederationQueryLimit: v,
	}
}

// DefaultLimitAsDataFederationLimit is a convenience function that returns DefaultLimit wrapped in DataFederationLimit
func DefaultLimitAsDataFederationLimit(v *DefaultLimit) DataFederationLimit {
	return DataFederationLimit{
		DefaultLimit: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DataFederationLimit) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'DataFederationQueryLimit'
	if jsonDict["name"] == "DataFederationQueryLimit" {
		// try to unmarshal JSON data into DataFederationQueryLimit
		err = json.Unmarshal(data, &dst.DataFederationQueryLimit)
		if err == nil {
			return nil // data stored in dst.DataFederationQueryLimit, return on the first match
		} else {
			dst.DataFederationQueryLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DataFederationQueryLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'DefaultLimit'
	if jsonDict["name"] == "DefaultLimit" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas.project.deployment.clusters'
	if jsonDict["name"] == "atlas.project.deployment.clusters" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas.project.deployment.nodesPerPrivateLinkRegion'
	if jsonDict["name"] == "atlas.project.deployment.nodesPerPrivateLinkRegion" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas.project.security.databaseAccess.customRoles'
	if jsonDict["name"] == "atlas.project.security.databaseAccess.customRoles" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas.project.security.databaseAccess.users'
	if jsonDict["name"] == "atlas.project.security.databaseAccess.users" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas.project.security.networkAccess.crossRegionEntries'
	if jsonDict["name"] == "atlas.project.security.networkAccess.crossRegionEntries" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'atlas.project.security.networkAccess.entries'
	if jsonDict["name"] == "atlas.project.security.networkAccess.entries" {
		// try to unmarshal JSON data into DefaultLimit
		err = json.Unmarshal(data, &dst.DefaultLimit)
		if err == nil {
			return nil // data stored in dst.DefaultLimit, return on the first match
		} else {
			dst.DefaultLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DefaultLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'dataFederation.bytesProcessed.daily'
	if jsonDict["name"] == "dataFederation.bytesProcessed.daily" {
		// try to unmarshal JSON data into DataFederationQueryLimit
		err = json.Unmarshal(data, &dst.DataFederationQueryLimit)
		if err == nil {
			return nil // data stored in dst.DataFederationQueryLimit, return on the first match
		} else {
			dst.DataFederationQueryLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DataFederationQueryLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'dataFederation.bytesProcessed.monthly'
	if jsonDict["name"] == "dataFederation.bytesProcessed.monthly" {
		// try to unmarshal JSON data into DataFederationQueryLimit
		err = json.Unmarshal(data, &dst.DataFederationQueryLimit)
		if err == nil {
			return nil // data stored in dst.DataFederationQueryLimit, return on the first match
		} else {
			dst.DataFederationQueryLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DataFederationQueryLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'dataFederation.bytesProcessed.query'
	if jsonDict["name"] == "dataFederation.bytesProcessed.query" {
		// try to unmarshal JSON data into DataFederationQueryLimit
		err = json.Unmarshal(data, &dst.DataFederationQueryLimit)
		if err == nil {
			return nil // data stored in dst.DataFederationQueryLimit, return on the first match
		} else {
			dst.DataFederationQueryLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DataFederationQueryLimit: %s", err.Error())
		}
	}

	// check if the discriminator value is 'dataFederation.bytesProcessed.weekly'
	if jsonDict["name"] == "dataFederation.bytesProcessed.weekly" {
		// try to unmarshal JSON data into DataFederationQueryLimit
		err = json.Unmarshal(data, &dst.DataFederationQueryLimit)
		if err == nil {
			return nil // data stored in dst.DataFederationQueryLimit, return on the first match
		} else {
			dst.DataFederationQueryLimit = nil
			return fmt.Errorf("failed to unmarshal DataFederationLimit as DataFederationQueryLimit: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DataFederationLimit) MarshalJSON() ([]byte, error) {
	if src.DataFederationQueryLimit != nil {
		return json.Marshal(&src.DataFederationQueryLimit)
	}

	if src.DefaultLimit != nil {
		return json.Marshal(&src.DefaultLimit)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DataFederationLimit) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.DataFederationQueryLimit != nil {
		return obj.DataFederationQueryLimit
	}

	if obj.DefaultLimit != nil {
		return obj.DefaultLimit
	}

	// all schemas are nil
	return nil
}

type NullableDataFederationLimit struct {
	value *DataFederationLimit
	isSet bool
}

func (v NullableDataFederationLimit) Get() *DataFederationLimit {
	return v.value
}

func (v *NullableDataFederationLimit) Set(val *DataFederationLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFederationLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFederationLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFederationLimit(val *DataFederationLimit) *NullableDataFederationLimit {
	return &NullableDataFederationLimit{value: val, isSet: true}
}

func (v NullableDataFederationLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFederationLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
