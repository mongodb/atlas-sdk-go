// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// NetworkContainerPeer - struct for NetworkContainerPeer
type NetworkContainerPeer struct {
	AWSPeerVpc       *AWSPeerVpc
	AzurePeerNetwork *AzurePeerNetwork
	GCPPeerVpc       *GCPPeerVpc
}

// AWSPeerVpcAsNetworkContainerPeer is a convenience function that returns AWSPeerVpc wrapped in NetworkContainerPeer
func AWSPeerVpcAsNetworkContainerPeer(v *AWSPeerVpc) NetworkContainerPeer {
	return NetworkContainerPeer{
		AWSPeerVpc: v,
	}
}

// AzurePeerNetworkAsNetworkContainerPeer is a convenience function that returns AzurePeerNetwork wrapped in NetworkContainerPeer
func AzurePeerNetworkAsNetworkContainerPeer(v *AzurePeerNetwork) NetworkContainerPeer {
	return NetworkContainerPeer{
		AzurePeerNetwork: v,
	}
}

// GCPPeerVpcAsNetworkContainerPeer is a convenience function that returns GCPPeerVpc wrapped in NetworkContainerPeer
func GCPPeerVpcAsNetworkContainerPeer(v *GCPPeerVpc) NetworkContainerPeer {
	return NetworkContainerPeer{
		GCPPeerVpc: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkContainerPeer) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into AWSPeerVpc
		err = json.Unmarshal(data, &dst.AWSPeerVpc)
		if err == nil {
			return nil // data stored in dst.AWSPeerVpc, return on the first match
		} else {
			dst.AWSPeerVpc = nil
			return fmt.Errorf("failed to unmarshal NetworkContainerPeer as AWSPeerVpc: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWSPeerVpc'
	if jsonDict["providerName"] == "AWSPeerVpc" {
		// try to unmarshal JSON data into AWSPeerVpc
		err = json.Unmarshal(data, &dst.AWSPeerVpc)
		if err == nil {
			return nil // data stored in dst.AWSPeerVpc, return on the first match
		} else {
			dst.AWSPeerVpc = nil
			return fmt.Errorf("failed to unmarshal NetworkContainerPeer as AWSPeerVpc: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into AzurePeerNetwork
		err = json.Unmarshal(data, &dst.AzurePeerNetwork)
		if err == nil {
			return nil // data stored in dst.AzurePeerNetwork, return on the first match
		} else {
			dst.AzurePeerNetwork = nil
			return fmt.Errorf("failed to unmarshal NetworkContainerPeer as AzurePeerNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzurePeerNetwork'
	if jsonDict["providerName"] == "AzurePeerNetwork" {
		// try to unmarshal JSON data into AzurePeerNetwork
		err = json.Unmarshal(data, &dst.AzurePeerNetwork)
		if err == nil {
			return nil // data stored in dst.AzurePeerNetwork, return on the first match
		} else {
			dst.AzurePeerNetwork = nil
			return fmt.Errorf("failed to unmarshal NetworkContainerPeer as AzurePeerNetwork: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCP'
	if jsonDict["providerName"] == "GCP" {
		// try to unmarshal JSON data into GCPPeerVpc
		err = json.Unmarshal(data, &dst.GCPPeerVpc)
		if err == nil {
			return nil // data stored in dst.GCPPeerVpc, return on the first match
		} else {
			dst.GCPPeerVpc = nil
			return fmt.Errorf("failed to unmarshal NetworkContainerPeer as GCPPeerVpc: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCPPeerVpc'
	if jsonDict["providerName"] == "GCPPeerVpc" {
		// try to unmarshal JSON data into GCPPeerVpc
		err = json.Unmarshal(data, &dst.GCPPeerVpc)
		if err == nil {
			return nil // data stored in dst.GCPPeerVpc, return on the first match
		} else {
			dst.GCPPeerVpc = nil
			return fmt.Errorf("failed to unmarshal NetworkContainerPeer as GCPPeerVpc: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkContainerPeer) MarshalJSON() ([]byte, error) {
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
func (obj *NetworkContainerPeer) GetActualInstance() interface{} {
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

type NullableNetworkContainerPeer struct {
	value *NetworkContainerPeer
	isSet bool
}

func (v NullableNetworkContainerPeer) Get() *NetworkContainerPeer {
	return v.value
}

func (v *NullableNetworkContainerPeer) Set(val *NetworkContainerPeer) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkContainerPeer) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkContainerPeer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkContainerPeer(val *NetworkContainerPeer) *NullableNetworkContainerPeer {
	return &NullableNetworkContainerPeer{value: val, isSet: true}
}

func (v NullableNetworkContainerPeer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkContainerPeer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
