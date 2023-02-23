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

// ListPeeringConnections200Response - struct for ListPeeringConnections200Response
type ListPeeringConnections200Response struct {
	PaginatedAWSPeerVpcView *PaginatedAWSPeerVpcView
	PaginatedAzurePeerNetworkView *PaginatedAzurePeerNetworkView
	PaginatedGCPPeerVpcView *PaginatedGCPPeerVpcView
}

// PaginatedAWSPeerVpcViewAsListPeeringConnections200Response is a convenience function that returns PaginatedAWSPeerVpcView wrapped in ListPeeringConnections200Response
func PaginatedAWSPeerVpcViewAsListPeeringConnections200Response(v *PaginatedAWSPeerVpcView) ListPeeringConnections200Response {
	return ListPeeringConnections200Response{
		PaginatedAWSPeerVpcView: v,
	}
}

// PaginatedAzurePeerNetworkViewAsListPeeringConnections200Response is a convenience function that returns PaginatedAzurePeerNetworkView wrapped in ListPeeringConnections200Response
func PaginatedAzurePeerNetworkViewAsListPeeringConnections200Response(v *PaginatedAzurePeerNetworkView) ListPeeringConnections200Response {
	return ListPeeringConnections200Response{
		PaginatedAzurePeerNetworkView: v,
	}
}

// PaginatedGCPPeerVpcViewAsListPeeringConnections200Response is a convenience function that returns PaginatedGCPPeerVpcView wrapped in ListPeeringConnections200Response
func PaginatedGCPPeerVpcViewAsListPeeringConnections200Response(v *PaginatedGCPPeerVpcView) ListPeeringConnections200Response {
	return ListPeeringConnections200Response{
		PaginatedGCPPeerVpcView: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListPeeringConnections200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PaginatedAWSPeerVpcView
	err = json.Unmarshal(data, &dst.PaginatedAWSPeerVpcView)
	if err == nil {
		jsonPaginatedAWSPeerVpcView, _ := json.Marshal(dst.PaginatedAWSPeerVpcView)
		if string(jsonPaginatedAWSPeerVpcView) == "{}" { // empty struct
			dst.PaginatedAWSPeerVpcView = nil
		} else {
			match++
		}
	} else {
		dst.PaginatedAWSPeerVpcView = nil
	}

	// try to unmarshal data into PaginatedAzurePeerNetworkView
	err = json.Unmarshal(data, &dst.PaginatedAzurePeerNetworkView)
	if err == nil {
		jsonPaginatedAzurePeerNetworkView, _ := json.Marshal(dst.PaginatedAzurePeerNetworkView)
		if string(jsonPaginatedAzurePeerNetworkView) == "{}" { // empty struct
			dst.PaginatedAzurePeerNetworkView = nil
		} else {
			match++
		}
	} else {
		dst.PaginatedAzurePeerNetworkView = nil
	}

	// try to unmarshal data into PaginatedGCPPeerVpcView
	err = json.Unmarshal(data, &dst.PaginatedGCPPeerVpcView)
	if err == nil {
		jsonPaginatedGCPPeerVpcView, _ := json.Marshal(dst.PaginatedGCPPeerVpcView)
		if string(jsonPaginatedGCPPeerVpcView) == "{}" { // empty struct
			dst.PaginatedGCPPeerVpcView = nil
		} else {
			match++
		}
	} else {
		dst.PaginatedGCPPeerVpcView = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PaginatedAWSPeerVpcView = nil
		dst.PaginatedAzurePeerNetworkView = nil
		dst.PaginatedGCPPeerVpcView = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListPeeringConnections200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListPeeringConnections200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListPeeringConnections200Response) MarshalJSON() ([]byte, error) {
	if src.PaginatedAWSPeerVpcView != nil {
		return json.Marshal(&src.PaginatedAWSPeerVpcView)
	}

	if src.PaginatedAzurePeerNetworkView != nil {
		return json.Marshal(&src.PaginatedAzurePeerNetworkView)
	}

	if src.PaginatedGCPPeerVpcView != nil {
		return json.Marshal(&src.PaginatedGCPPeerVpcView)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListPeeringConnections200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PaginatedAWSPeerVpcView != nil {
		return obj.PaginatedAWSPeerVpcView
	}

	if obj.PaginatedAzurePeerNetworkView != nil {
		return obj.PaginatedAzurePeerNetworkView
	}

	if obj.PaginatedGCPPeerVpcView != nil {
		return obj.PaginatedGCPPeerVpcView
	}

	// all schemas are nil
	return nil
}

type NullableListPeeringConnections200Response struct {
	value *ListPeeringConnections200Response
	isSet bool
}

func (v NullableListPeeringConnections200Response) Get() *ListPeeringConnections200Response {
	return v.value
}

func (v *NullableListPeeringConnections200Response) Set(val *ListPeeringConnections200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListPeeringConnections200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListPeeringConnections200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPeeringConnections200Response(val *ListPeeringConnections200Response) *NullableListPeeringConnections200Response {
	return &NullableListPeeringConnections200Response{value: val, isSet: true}
}

func (v NullableListPeeringConnections200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPeeringConnections200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


