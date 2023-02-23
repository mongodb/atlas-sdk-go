/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
	"fmt"
)

// EndpointService - struct for EndpointService
type EndpointService struct {
	AWSPrivateLinkConnection *AWSPrivateLinkConnection
	AzurePrivateLinkConnection *AzurePrivateLinkConnection
	GCPEndpointService *GCPEndpointService
}

// AWSPrivateLinkConnectionAsEndpointService is a convenience function that returns AWSPrivateLinkConnection wrapped in EndpointService
func AWSPrivateLinkConnectionAsEndpointService(v *AWSPrivateLinkConnection) EndpointService {
	return EndpointService{
		AWSPrivateLinkConnection: v,
	}
}

// AzurePrivateLinkConnectionAsEndpointService is a convenience function that returns AzurePrivateLinkConnection wrapped in EndpointService
func AzurePrivateLinkConnectionAsEndpointService(v *AzurePrivateLinkConnection) EndpointService {
	return EndpointService{
		AzurePrivateLinkConnection: v,
	}
}

// GCPEndpointServiceAsEndpointService is a convenience function that returns GCPEndpointService wrapped in EndpointService
func GCPEndpointServiceAsEndpointService(v *GCPEndpointService) EndpointService {
	return EndpointService{
		GCPEndpointService: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EndpointService) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSPrivateLinkConnection
	err = json.Unmarshal(data, &dst.AWSPrivateLinkConnection)
	if err == nil {
		jsonAWSPrivateLinkConnection, _ := json.Marshal(dst.AWSPrivateLinkConnection)
		if string(jsonAWSPrivateLinkConnection) == "{}" { // empty struct
			dst.AWSPrivateLinkConnection = nil
		} else {
			match++
		}
	} else {
		dst.AWSPrivateLinkConnection = nil
	}

	// try to unmarshal data into AzurePrivateLinkConnection
	err = json.Unmarshal(data, &dst.AzurePrivateLinkConnection)
	if err == nil {
		jsonAzurePrivateLinkConnection, _ := json.Marshal(dst.AzurePrivateLinkConnection)
		if string(jsonAzurePrivateLinkConnection) == "{}" { // empty struct
			dst.AzurePrivateLinkConnection = nil
		} else {
			match++
		}
	} else {
		dst.AzurePrivateLinkConnection = nil
	}

	// try to unmarshal data into GCPEndpointService
	err = json.Unmarshal(data, &dst.GCPEndpointService)
	if err == nil {
		jsonGCPEndpointService, _ := json.Marshal(dst.GCPEndpointService)
		if string(jsonGCPEndpointService) == "{}" { // empty struct
			dst.GCPEndpointService = nil
		} else {
			match++
		}
	} else {
		dst.GCPEndpointService = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AWSPrivateLinkConnection = nil
		dst.AzurePrivateLinkConnection = nil
		dst.GCPEndpointService = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EndpointService)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EndpointService)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EndpointService) MarshalJSON() ([]byte, error) {
	if src.AWSPrivateLinkConnection != nil {
		return json.Marshal(&src.AWSPrivateLinkConnection)
	}

	if src.AzurePrivateLinkConnection != nil {
		return json.Marshal(&src.AzurePrivateLinkConnection)
	}

	if src.GCPEndpointService != nil {
		return json.Marshal(&src.GCPEndpointService)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EndpointService) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AWSPrivateLinkConnection != nil {
		return obj.AWSPrivateLinkConnection
	}

	if obj.AzurePrivateLinkConnection != nil {
		return obj.AzurePrivateLinkConnection
	}

	if obj.GCPEndpointService != nil {
		return obj.GCPEndpointService
	}

	// all schemas are nil
	return nil
}

type NullableEndpointService struct {
	value *EndpointService
	isSet bool
}

func (v NullableEndpointService) Get() *EndpointService {
	return v.value
}

func (v *NullableEndpointService) Set(val *EndpointService) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointService) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointService(val *EndpointService) *NullableEndpointService {
	return &NullableEndpointService{value: val, isSet: true}
}

func (v NullableEndpointService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


