// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// PrivateLinkEndpoint - struct for PrivateLinkEndpoint
type PrivateLinkEndpoint struct {
	AWSInterfaceEndpoint    *AWSInterfaceEndpoint
	AzurePrivateEndpoint    *AzurePrivateEndpoint
	PrivateGCPEndpointGroup *PrivateGCPEndpointGroup
}

// AWSInterfaceEndpointAsPrivateLinkEndpoint is a convenience function that returns AWSInterfaceEndpoint wrapped in PrivateLinkEndpoint
func AWSInterfaceEndpointAsPrivateLinkEndpoint(v *AWSInterfaceEndpoint) PrivateLinkEndpoint {
	return PrivateLinkEndpoint{
		AWSInterfaceEndpoint: v,
	}
}

// AzurePrivateEndpointAsPrivateLinkEndpoint is a convenience function that returns AzurePrivateEndpoint wrapped in PrivateLinkEndpoint
func AzurePrivateEndpointAsPrivateLinkEndpoint(v *AzurePrivateEndpoint) PrivateLinkEndpoint {
	return PrivateLinkEndpoint{
		AzurePrivateEndpoint: v,
	}
}

// PrivateGCPEndpointGroupAsPrivateLinkEndpoint is a convenience function that returns PrivateGCPEndpointGroup wrapped in PrivateLinkEndpoint
func PrivateGCPEndpointGroupAsPrivateLinkEndpoint(v *PrivateGCPEndpointGroup) PrivateLinkEndpoint {
	return PrivateLinkEndpoint{
		PrivateGCPEndpointGroup: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PrivateLinkEndpoint) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["cloudProvider"] == "AWS" {
		// try to unmarshal JSON data into AWSInterfaceEndpoint
		err = json.Unmarshal(data, &dst.AWSInterfaceEndpoint)
		if err == nil {
			return nil // data stored in dst.AWSInterfaceEndpoint, return on the first match
		} else {
			dst.AWSInterfaceEndpoint = nil
			return fmt.Errorf("failed to unmarshal PrivateLinkEndpoint as AWSInterfaceEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWSInterfaceEndpoint'
	if jsonDict["cloudProvider"] == "AWSInterfaceEndpoint" {
		// try to unmarshal JSON data into AWSInterfaceEndpoint
		err = json.Unmarshal(data, &dst.AWSInterfaceEndpoint)
		if err == nil {
			return nil // data stored in dst.AWSInterfaceEndpoint, return on the first match
		} else {
			dst.AWSInterfaceEndpoint = nil
			return fmt.Errorf("failed to unmarshal PrivateLinkEndpoint as AWSInterfaceEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["cloudProvider"] == "AZURE" {
		// try to unmarshal JSON data into PrivateGCPEndpointGroup
		err = json.Unmarshal(data, &dst.PrivateGCPEndpointGroup)
		if err == nil {
			return nil // data stored in dst.PrivateGCPEndpointGroup, return on the first match
		} else {
			dst.PrivateGCPEndpointGroup = nil
			return fmt.Errorf("failed to unmarshal PrivateLinkEndpoint as PrivateGCPEndpointGroup: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzurePrivateEndpoint'
	if jsonDict["cloudProvider"] == "AzurePrivateEndpoint" {
		// try to unmarshal JSON data into AzurePrivateEndpoint
		err = json.Unmarshal(data, &dst.AzurePrivateEndpoint)
		if err == nil {
			return nil // data stored in dst.AzurePrivateEndpoint, return on the first match
		} else {
			dst.AzurePrivateEndpoint = nil
			return fmt.Errorf("failed to unmarshal PrivateLinkEndpoint as AzurePrivateEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCP'
	if jsonDict["cloudProvider"] == "GCP" {
		// try to unmarshal JSON data into AzurePrivateEndpoint
		err = json.Unmarshal(data, &dst.AzurePrivateEndpoint)
		if err == nil {
			return nil // data stored in dst.AzurePrivateEndpoint, return on the first match
		} else {
			dst.AzurePrivateEndpoint = nil
			return fmt.Errorf("failed to unmarshal PrivateLinkEndpoint as AzurePrivateEndpoint: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PrivateGCPEndpointGroup'
	if jsonDict["cloudProvider"] == "PrivateGCPEndpointGroup" {
		// try to unmarshal JSON data into PrivateGCPEndpointGroup
		err = json.Unmarshal(data, &dst.PrivateGCPEndpointGroup)
		if err == nil {
			return nil // data stored in dst.PrivateGCPEndpointGroup, return on the first match
		} else {
			dst.PrivateGCPEndpointGroup = nil
			return fmt.Errorf("failed to unmarshal PrivateLinkEndpoint as PrivateGCPEndpointGroup: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PrivateLinkEndpoint) MarshalJSON() ([]byte, error) {
	if src.AWSInterfaceEndpoint != nil {
		return json.Marshal(&src.AWSInterfaceEndpoint)
	}

	if src.AzurePrivateEndpoint != nil {
		return json.Marshal(&src.AzurePrivateEndpoint)
	}

	if src.PrivateGCPEndpointGroup != nil {
		return json.Marshal(&src.PrivateGCPEndpointGroup)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PrivateLinkEndpoint) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSInterfaceEndpoint != nil {
		return obj.AWSInterfaceEndpoint
	}

	if obj.AzurePrivateEndpoint != nil {
		return obj.AzurePrivateEndpoint
	}

	if obj.PrivateGCPEndpointGroup != nil {
		return obj.PrivateGCPEndpointGroup
	}

	// all schemas are nil
	return nil
}

type NullablePrivateLinkEndpoint struct {
	value *PrivateLinkEndpoint
	isSet bool
}

func (v NullablePrivateLinkEndpoint) Get() *PrivateLinkEndpoint {
	return v.value
}

func (v *NullablePrivateLinkEndpoint) Set(val *PrivateLinkEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateLinkEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateLinkEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateLinkEndpoint(val *PrivateLinkEndpoint) *NullablePrivateLinkEndpoint {
	return &NullablePrivateLinkEndpoint{value: val, isSet: true}
}

func (v NullablePrivateLinkEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateLinkEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
