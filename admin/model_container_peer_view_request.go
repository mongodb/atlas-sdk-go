// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// ContainerPeerViewRequest - Collection of settings that configures the network connection for a virtual private connection.
type ContainerPeerViewRequest struct {
	AWSPeerVpcRequest       *AWSPeerVpcRequest
	AzurePeerNetworkRequest *AzurePeerNetworkRequest
	GCPPeerVpcRequest       *GCPPeerVpcRequest
}

// AWSPeerVpcRequestAsContainerPeerViewRequest is a convenience function that returns AWSPeerVpcRequest wrapped in ContainerPeerViewRequest
func AWSPeerVpcRequestAsContainerPeerViewRequest(v *AWSPeerVpcRequest) ContainerPeerViewRequest {
	return ContainerPeerViewRequest{
		AWSPeerVpcRequest: v,
	}
}

// AzurePeerNetworkRequestAsContainerPeerViewRequest is a convenience function that returns AzurePeerNetworkRequest wrapped in ContainerPeerViewRequest
func AzurePeerNetworkRequestAsContainerPeerViewRequest(v *AzurePeerNetworkRequest) ContainerPeerViewRequest {
	return ContainerPeerViewRequest{
		AzurePeerNetworkRequest: v,
	}
}

// GCPPeerVpcRequestAsContainerPeerViewRequest is a convenience function that returns GCPPeerVpcRequest wrapped in ContainerPeerViewRequest
func GCPPeerVpcRequestAsContainerPeerViewRequest(v *GCPPeerVpcRequest) ContainerPeerViewRequest {
	return ContainerPeerViewRequest{
		GCPPeerVpcRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ContainerPeerViewRequest) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into AWSPeerVpcRequest
		err = json.Unmarshal(data, &dst.AWSPeerVpcRequest)
		if err == nil {
			return nil // data stored in dst.AWSPeerVpcRequest, return on the first match
		} else {
			dst.AWSPeerVpcRequest = nil
			return fmt.Errorf("failed to unmarshal ContainerPeerViewRequest as AWSPeerVpcRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWSPeerVpcRequest'
	if jsonDict["providerName"] == "AWSPeerVpcRequest" {
		// try to unmarshal JSON data into AWSPeerVpcRequest
		err = json.Unmarshal(data, &dst.AWSPeerVpcRequest)
		if err == nil {
			return nil // data stored in dst.AWSPeerVpcRequest, return on the first match
		} else {
			dst.AWSPeerVpcRequest = nil
			return fmt.Errorf("failed to unmarshal ContainerPeerViewRequest as AWSPeerVpcRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into AzurePeerNetworkRequest
		err = json.Unmarshal(data, &dst.AzurePeerNetworkRequest)
		if err == nil {
			return nil // data stored in dst.AzurePeerNetworkRequest, return on the first match
		} else {
			dst.AzurePeerNetworkRequest = nil
			return fmt.Errorf("failed to unmarshal ContainerPeerViewRequest as AzurePeerNetworkRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzurePeerNetworkRequest'
	if jsonDict["providerName"] == "AzurePeerNetworkRequest" {
		// try to unmarshal JSON data into AzurePeerNetworkRequest
		err = json.Unmarshal(data, &dst.AzurePeerNetworkRequest)
		if err == nil {
			return nil // data stored in dst.AzurePeerNetworkRequest, return on the first match
		} else {
			dst.AzurePeerNetworkRequest = nil
			return fmt.Errorf("failed to unmarshal ContainerPeerViewRequest as AzurePeerNetworkRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCP'
	if jsonDict["providerName"] == "GCP" {
		// try to unmarshal JSON data into GCPPeerVpcRequest
		err = json.Unmarshal(data, &dst.GCPPeerVpcRequest)
		if err == nil {
			return nil // data stored in dst.GCPPeerVpcRequest, return on the first match
		} else {
			dst.GCPPeerVpcRequest = nil
			return fmt.Errorf("failed to unmarshal ContainerPeerViewRequest as GCPPeerVpcRequest: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCPPeerVpcRequest'
	if jsonDict["providerName"] == "GCPPeerVpcRequest" {
		// try to unmarshal JSON data into GCPPeerVpcRequest
		err = json.Unmarshal(data, &dst.GCPPeerVpcRequest)
		if err == nil {
			return nil // data stored in dst.GCPPeerVpcRequest, return on the first match
		} else {
			dst.GCPPeerVpcRequest = nil
			return fmt.Errorf("failed to unmarshal ContainerPeerViewRequest as GCPPeerVpcRequest: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ContainerPeerViewRequest) MarshalJSON() ([]byte, error) {
	if src.AWSPeerVpcRequest != nil {
		return json.Marshal(&src.AWSPeerVpcRequest)
	}

	if src.AzurePeerNetworkRequest != nil {
		return json.Marshal(&src.AzurePeerNetworkRequest)
	}

	if src.GCPPeerVpcRequest != nil {
		return json.Marshal(&src.GCPPeerVpcRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ContainerPeerViewRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSPeerVpcRequest != nil {
		return obj.AWSPeerVpcRequest
	}

	if obj.AzurePeerNetworkRequest != nil {
		return obj.AzurePeerNetworkRequest
	}

	if obj.GCPPeerVpcRequest != nil {
		return obj.GCPPeerVpcRequest
	}

	// all schemas are nil
	return nil
}

type NullableContainerPeerViewRequest struct {
	value *ContainerPeerViewRequest
	isSet bool
}

func (v NullableContainerPeerViewRequest) Get() *ContainerPeerViewRequest {
	return v.value
}

func (v *NullableContainerPeerViewRequest) Set(val *ContainerPeerViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerPeerViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerPeerViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerPeerViewRequest(val *ContainerPeerViewRequest) *NullableContainerPeerViewRequest {
	return &NullableContainerPeerViewRequest{value: val, isSet: true}
}

func (v NullableContainerPeerViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerPeerViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
