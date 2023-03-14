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

// GetPeeringConnection200Response - struct for GetPeeringConnection200Response
type GetPeeringConnection200Response struct {
	AWSPeerVpc *AWSPeerVpc
	AzurePeerNetwork *AzurePeerNetwork
	GCPPeerVpc *GCPPeerVpc
}

// AWSPeerVpcAsGetPeeringConnection200Response is a convenience function that returns AWSPeerVpc wrapped in GetPeeringConnection200Response
func AWSPeerVpcAsGetPeeringConnection200Response(v *AWSPeerVpc) GetPeeringConnection200Response {
	return GetPeeringConnection200Response{
		AWSPeerVpc: v,
	}
}

// AzurePeerNetworkAsGetPeeringConnection200Response is a convenience function that returns AzurePeerNetwork wrapped in GetPeeringConnection200Response
func AzurePeerNetworkAsGetPeeringConnection200Response(v *AzurePeerNetwork) GetPeeringConnection200Response {
	return GetPeeringConnection200Response{
		AzurePeerNetwork: v,
	}
}

// GCPPeerVpcAsGetPeeringConnection200Response is a convenience function that returns GCPPeerVpc wrapped in GetPeeringConnection200Response
func GCPPeerVpcAsGetPeeringConnection200Response(v *GCPPeerVpc) GetPeeringConnection200Response {
	return GetPeeringConnection200Response{
		GCPPeerVpc: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetPeeringConnection200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AWSPeerVpc
	err = json.Unmarshal(data, &dst.AWSPeerVpc)
	if err == nil {
		jsonAWSPeerVpc, _ := json.Marshal(dst.AWSPeerVpc)
		if string(jsonAWSPeerVpc) == "{}" { // empty struct
			dst.AWSPeerVpc = nil
		} else {
			match++
		}
	} else {
		dst.AWSPeerVpc = nil
	}

	// try to unmarshal data into AzurePeerNetwork
	err = json.Unmarshal(data, &dst.AzurePeerNetwork)
	if err == nil {
		jsonAzurePeerNetwork, _ := json.Marshal(dst.AzurePeerNetwork)
		if string(jsonAzurePeerNetwork) == "{}" { // empty struct
			dst.AzurePeerNetwork = nil
		} else {
			match++
		}
	} else {
		dst.AzurePeerNetwork = nil
	}

	// try to unmarshal data into GCPPeerVpc
	err = json.Unmarshal(data, &dst.GCPPeerVpc)
	if err == nil {
		jsonGCPPeerVpc, _ := json.Marshal(dst.GCPPeerVpc)
		if string(jsonGCPPeerVpc) == "{}" { // empty struct
			dst.GCPPeerVpc = nil
		} else {
			match++
		}
	} else {
		dst.GCPPeerVpc = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AWSPeerVpc = nil
		dst.AzurePeerNetwork = nil
		dst.GCPPeerVpc = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetPeeringConnection200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetPeeringConnection200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetPeeringConnection200Response) MarshalJSON() ([]byte, error) {
	if src.AWSPeerVpc != nil {
		return json.Marshal(&src.AWSPeerVpc)
	}

	if src.AzurePeerNetwork != nil {
		return json.Marshal(&src.AzurePeerNetwork)
	}

	if src.GCPPeerVpc != nil {
		return json.Marshal(&src.GCPPeerVpc)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetPeeringConnection200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AWSPeerVpc != nil {
		return obj.AWSPeerVpc
	}

	if obj.AzurePeerNetwork != nil {
		return obj.AzurePeerNetwork
	}

	if obj.GCPPeerVpc != nil {
		return obj.GCPPeerVpc
	}

	// all schemas are nil
	return nil
}

type NullableGetPeeringConnection200Response struct {
	value *GetPeeringConnection200Response
	isSet bool
}

func (v NullableGetPeeringConnection200Response) Get() *GetPeeringConnection200Response {
	return v.value
}

func (v *NullableGetPeeringConnection200Response) Set(val *GetPeeringConnection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPeeringConnection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPeeringConnection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPeeringConnection200Response(val *GetPeeringConnection200Response) *NullableGetPeeringConnection200Response {
	return &NullableGetPeeringConnection200Response{value: val, isSet: true}
}

func (v NullableGetPeeringConnection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPeeringConnection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


