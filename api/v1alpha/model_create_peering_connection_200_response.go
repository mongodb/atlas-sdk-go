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

// CreatePeeringConnection200Response - struct for CreatePeeringConnection200Response
type CreatePeeringConnection200Response struct {
	AWSPeerVpc *AWSPeerVpc
	AzurePeerNetwork *AzurePeerNetwork
	GCPPeerVpc *GCPPeerVpc
}

// AWSPeerVpcAsCreatePeeringConnection200Response is a convenience function that returns AWSPeerVpc wrapped in CreatePeeringConnection200Response
func AWSPeerVpcAsCreatePeeringConnection200Response(v *AWSPeerVpc) CreatePeeringConnection200Response {
	return CreatePeeringConnection200Response{
		AWSPeerVpc: v,
	}
}

// AzurePeerNetworkAsCreatePeeringConnection200Response is a convenience function that returns AzurePeerNetwork wrapped in CreatePeeringConnection200Response
func AzurePeerNetworkAsCreatePeeringConnection200Response(v *AzurePeerNetwork) CreatePeeringConnection200Response {
	return CreatePeeringConnection200Response{
		AzurePeerNetwork: v,
	}
}

// GCPPeerVpcAsCreatePeeringConnection200Response is a convenience function that returns GCPPeerVpc wrapped in CreatePeeringConnection200Response
func GCPPeerVpcAsCreatePeeringConnection200Response(v *GCPPeerVpc) CreatePeeringConnection200Response {
	return CreatePeeringConnection200Response{
		GCPPeerVpc: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreatePeeringConnection200Response) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(CreatePeeringConnection200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreatePeeringConnection200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreatePeeringConnection200Response) MarshalJSON() ([]byte, error) {
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
func (obj *CreatePeeringConnection200Response) GetActualInstance() (interface{}) {
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

type NullableCreatePeeringConnection200Response struct {
	value *CreatePeeringConnection200Response
	isSet bool
}

func (v NullableCreatePeeringConnection200Response) Get() *CreatePeeringConnection200Response {
	return v.value
}

func (v *NullableCreatePeeringConnection200Response) Set(val *CreatePeeringConnection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePeeringConnection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePeeringConnection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePeeringConnection200Response(val *CreatePeeringConnection200Response) *NullableCreatePeeringConnection200Response {
	return &NullableCreatePeeringConnection200Response{value: val, isSet: true}
}

func (v NullableCreatePeeringConnection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePeeringConnection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


