// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// CreateEndpointRequest - struct for CreateEndpointRequest
type CreateEndpointRequest struct {
	CreateAWSEndpointRequest      *CreateAWSEndpointRequest
	CreateAzureEndpointRequest    *CreateAzureEndpointRequest
	CreateGCPEndpointGroupRequest *CreateGCPEndpointGroupRequest
}

// CreateAWSEndpointRequestAsCreateEndpointRequest is a convenience function that returns CreateAWSEndpointRequest wrapped in CreateEndpointRequest
func CreateAWSEndpointRequestAsCreateEndpointRequest(v *CreateAWSEndpointRequest) CreateEndpointRequest {
	return CreateEndpointRequest{
		CreateAWSEndpointRequest: v,
	}
}

// CreateAzureEndpointRequestAsCreateEndpointRequest is a convenience function that returns CreateAzureEndpointRequest wrapped in CreateEndpointRequest
func CreateAzureEndpointRequestAsCreateEndpointRequest(v *CreateAzureEndpointRequest) CreateEndpointRequest {
	return CreateEndpointRequest{
		CreateAzureEndpointRequest: v,
	}
}

// CreateGCPEndpointGroupRequestAsCreateEndpointRequest is a convenience function that returns CreateGCPEndpointGroupRequest wrapped in CreateEndpointRequest
func CreateGCPEndpointGroupRequestAsCreateEndpointRequest(v *CreateGCPEndpointGroupRequest) CreateEndpointRequest {
	return CreateEndpointRequest{
		CreateGCPEndpointGroupRequest: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateEndpointRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateAWSEndpointRequest
	err = json.Unmarshal(data, &dst.CreateAWSEndpointRequest)
	if err == nil {
		jsonCreateAWSEndpointRequest, _ := json.Marshal(dst.CreateAWSEndpointRequest)
		if string(jsonCreateAWSEndpointRequest) == "{}" { // empty struct
			dst.CreateAWSEndpointRequest = nil
		} else {
			match++
		}
	} else {
		dst.CreateAWSEndpointRequest = nil
	}

	// try to unmarshal data into CreateAzureEndpointRequest
	err = json.Unmarshal(data, &dst.CreateAzureEndpointRequest)
	if err == nil {
		jsonCreateAzureEndpointRequest, _ := json.Marshal(dst.CreateAzureEndpointRequest)
		if string(jsonCreateAzureEndpointRequest) == "{}" { // empty struct
			dst.CreateAzureEndpointRequest = nil
		} else {
			match++
		}
	} else {
		dst.CreateAzureEndpointRequest = nil
	}

	// try to unmarshal data into CreateGCPEndpointGroupRequest
	err = json.Unmarshal(data, &dst.CreateGCPEndpointGroupRequest)
	if err == nil {
		jsonCreateGCPEndpointGroupRequest, _ := json.Marshal(dst.CreateGCPEndpointGroupRequest)
		if string(jsonCreateGCPEndpointGroupRequest) == "{}" { // empty struct
			dst.CreateGCPEndpointGroupRequest = nil
		} else {
			match++
		}
	} else {
		dst.CreateGCPEndpointGroupRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateAWSEndpointRequest = nil
		dst.CreateAzureEndpointRequest = nil
		dst.CreateGCPEndpointGroupRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateEndpointRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateEndpointRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateEndpointRequest) MarshalJSON() ([]byte, error) {
	if src.CreateAWSEndpointRequest != nil {
		return json.Marshal(&src.CreateAWSEndpointRequest)
	}

	if src.CreateAzureEndpointRequest != nil {
		return json.Marshal(&src.CreateAzureEndpointRequest)
	}

	if src.CreateGCPEndpointGroupRequest != nil {
		return json.Marshal(&src.CreateGCPEndpointGroupRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateEndpointRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CreateAWSEndpointRequest != nil {
		return obj.CreateAWSEndpointRequest
	}

	if obj.CreateAzureEndpointRequest != nil {
		return obj.CreateAzureEndpointRequest
	}

	if obj.CreateGCPEndpointGroupRequest != nil {
		return obj.CreateGCPEndpointGroupRequest
	}

	// all schemas are nil
	return nil
}

type NullableCreateEndpointRequest struct {
	value *CreateEndpointRequest
	isSet bool
}

func (v NullableCreateEndpointRequest) Get() *CreateEndpointRequest {
	return v.value
}

func (v *NullableCreateEndpointRequest) Set(val *CreateEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEndpointRequest(val *CreateEndpointRequest) *NullableCreateEndpointRequest {
	return &NullableCreateEndpointRequest{value: val, isSet: true}
}

func (v NullableCreateEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
