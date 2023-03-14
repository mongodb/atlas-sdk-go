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

// Endpoint - struct for Endpoint
type Endpoint struct {
	AWSInterfaceEndpoint *AWSInterfaceEndpoint
	AzurePrivateEndpoint *AzurePrivateEndpoint
	GCPEndpointGroup *GCPEndpointGroup
}

// AWSInterfaceEndpointAsEndpoint is a convenience function that returns AWSInterfaceEndpoint wrapped in Endpoint
func AWSInterfaceEndpointAsEndpoint(v *AWSInterfaceEndpoint) Endpoint {
	return Endpoint{
		AWSInterfaceEndpoint: v,
	}
}

// AzurePrivateEndpointAsEndpoint is a convenience function that returns AzurePrivateEndpoint wrapped in Endpoint
func AzurePrivateEndpointAsEndpoint(v *AzurePrivateEndpoint) Endpoint {
	return Endpoint{
		AzurePrivateEndpoint: v,
	}
}

// GCPEndpointGroupAsEndpoint is a convenience function that returns GCPEndpointGroup wrapped in Endpoint
func GCPEndpointGroupAsEndpoint(v *GCPEndpointGroup) Endpoint {
	return Endpoint{
		GCPEndpointGroup: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Endpoint) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSInterfaceEndpoint
	err = json.Unmarshal(data, &dst.AWSInterfaceEndpoint)
	if err == nil {
		jsonAWSInterfaceEndpoint, _ := json.Marshal(dst.AWSInterfaceEndpoint)
		if string(jsonAWSInterfaceEndpoint) == "{}" { // empty struct
			dst.AWSInterfaceEndpoint = nil
		} else {
			match++
		}
	} else {
		dst.AWSInterfaceEndpoint = nil
	}

	// try to unmarshal data into AzurePrivateEndpoint
	err = json.Unmarshal(data, &dst.AzurePrivateEndpoint)
	if err == nil {
		jsonAzurePrivateEndpoint, _ := json.Marshal(dst.AzurePrivateEndpoint)
		if string(jsonAzurePrivateEndpoint) == "{}" { // empty struct
			dst.AzurePrivateEndpoint = nil
		} else {
			match++
		}
	} else {
		dst.AzurePrivateEndpoint = nil
	}

	// try to unmarshal data into GCPEndpointGroup
	err = json.Unmarshal(data, &dst.GCPEndpointGroup)
	if err == nil {
		jsonGCPEndpointGroup, _ := json.Marshal(dst.GCPEndpointGroup)
		if string(jsonGCPEndpointGroup) == "{}" { // empty struct
			dst.GCPEndpointGroup = nil
		} else {
			match++
		}
	} else {
		dst.GCPEndpointGroup = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AWSInterfaceEndpoint = nil
		dst.AzurePrivateEndpoint = nil
		dst.GCPEndpointGroup = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Endpoint)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Endpoint)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Endpoint) MarshalJSON() ([]byte, error) {
	if src.AWSInterfaceEndpoint != nil {
		return json.Marshal(&src.AWSInterfaceEndpoint)
	}

	if src.AzurePrivateEndpoint != nil {
		return json.Marshal(&src.AzurePrivateEndpoint)
	}

	if src.GCPEndpointGroup != nil {
		return json.Marshal(&src.GCPEndpointGroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Endpoint) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AWSInterfaceEndpoint != nil {
		return obj.AWSInterfaceEndpoint
	}

	if obj.AzurePrivateEndpoint != nil {
		return obj.AzurePrivateEndpoint
	}

	if obj.GCPEndpointGroup != nil {
		return obj.GCPEndpointGroup
	}

	// all schemas are nil
	return nil
}

type NullableEndpoint struct {
	value *Endpoint
	isSet bool
}

func (v NullableEndpoint) Get() *Endpoint {
	return v.value
}

func (v *NullableEndpoint) Set(val *Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpoint(val *Endpoint) *NullableEndpoint {
	return &NullableEndpoint{value: val, isSet: true}
}

func (v NullableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


