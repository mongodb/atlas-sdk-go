// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// ServerlessTenantEndpointUpdate - Update view for a serverless tenant endpoint.
type ServerlessTenantEndpointUpdate struct {
	ServerlessAWSTenantEndpointUpdate   *ServerlessAWSTenantEndpointUpdate
	ServerlessAzureTenantEndpointUpdate *ServerlessAzureTenantEndpointUpdate
}

// ServerlessAWSTenantEndpointUpdateAsServerlessTenantEndpointUpdate is a convenience function that returns ServerlessAWSTenantEndpointUpdate wrapped in ServerlessTenantEndpointUpdate
func ServerlessAWSTenantEndpointUpdateAsServerlessTenantEndpointUpdate(v *ServerlessAWSTenantEndpointUpdate) ServerlessTenantEndpointUpdate {
	return ServerlessTenantEndpointUpdate{
		ServerlessAWSTenantEndpointUpdate: v,
	}
}

// ServerlessAzureTenantEndpointUpdateAsServerlessTenantEndpointUpdate is a convenience function that returns ServerlessAzureTenantEndpointUpdate wrapped in ServerlessTenantEndpointUpdate
func ServerlessAzureTenantEndpointUpdateAsServerlessTenantEndpointUpdate(v *ServerlessAzureTenantEndpointUpdate) ServerlessTenantEndpointUpdate {
	return ServerlessTenantEndpointUpdate{
		ServerlessAzureTenantEndpointUpdate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServerlessTenantEndpointUpdate) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into ServerlessAWSTenantEndpointUpdate
		err = json.Unmarshal(data, &dst.ServerlessAWSTenantEndpointUpdate)
		if err == nil {
			return nil // data stored in dst.ServerlessAWSTenantEndpointUpdate, return on the first match
		} else {
			dst.ServerlessAWSTenantEndpointUpdate = nil
			return fmt.Errorf("failed to unmarshal ServerlessTenantEndpointUpdate as ServerlessAWSTenantEndpointUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into ServerlessAzureTenantEndpointUpdate
		err = json.Unmarshal(data, &dst.ServerlessAzureTenantEndpointUpdate)
		if err == nil {
			return nil // data stored in dst.ServerlessAzureTenantEndpointUpdate, return on the first match
		} else {
			dst.ServerlessAzureTenantEndpointUpdate = nil
			return fmt.Errorf("failed to unmarshal ServerlessTenantEndpointUpdate as ServerlessAzureTenantEndpointUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ServerlessAWSTenantEndpointUpdate'
	if jsonDict["providerName"] == "ServerlessAWSTenantEndpointUpdate" {
		// try to unmarshal JSON data into ServerlessAWSTenantEndpointUpdate
		err = json.Unmarshal(data, &dst.ServerlessAWSTenantEndpointUpdate)
		if err == nil {
			return nil // data stored in dst.ServerlessAWSTenantEndpointUpdate, return on the first match
		} else {
			dst.ServerlessAWSTenantEndpointUpdate = nil
			return fmt.Errorf("failed to unmarshal ServerlessTenantEndpointUpdate as ServerlessAWSTenantEndpointUpdate: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ServerlessAzureTenantEndpointUpdate'
	if jsonDict["providerName"] == "ServerlessAzureTenantEndpointUpdate" {
		// try to unmarshal JSON data into ServerlessAzureTenantEndpointUpdate
		err = json.Unmarshal(data, &dst.ServerlessAzureTenantEndpointUpdate)
		if err == nil {
			return nil // data stored in dst.ServerlessAzureTenantEndpointUpdate, return on the first match
		} else {
			dst.ServerlessAzureTenantEndpointUpdate = nil
			return fmt.Errorf("failed to unmarshal ServerlessTenantEndpointUpdate as ServerlessAzureTenantEndpointUpdate: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServerlessTenantEndpointUpdate) MarshalJSON() ([]byte, error) {
	if src.ServerlessAWSTenantEndpointUpdate != nil {
		return json.Marshal(&src.ServerlessAWSTenantEndpointUpdate)
	}

	if src.ServerlessAzureTenantEndpointUpdate != nil {
		return json.Marshal(&src.ServerlessAzureTenantEndpointUpdate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServerlessTenantEndpointUpdate) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ServerlessAWSTenantEndpointUpdate != nil {
		return obj.ServerlessAWSTenantEndpointUpdate
	}

	if obj.ServerlessAzureTenantEndpointUpdate != nil {
		return obj.ServerlessAzureTenantEndpointUpdate
	}

	// all schemas are nil
	return nil
}

type NullableServerlessTenantEndpointUpdate struct {
	value *ServerlessTenantEndpointUpdate
	isSet bool
}

func (v NullableServerlessTenantEndpointUpdate) Get() *ServerlessTenantEndpointUpdate {
	return v.value
}

func (v *NullableServerlessTenantEndpointUpdate) Set(val *ServerlessTenantEndpointUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessTenantEndpointUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessTenantEndpointUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessTenantEndpointUpdate(val *ServerlessTenantEndpointUpdate) *NullableServerlessTenantEndpointUpdate {
	return &NullableServerlessTenantEndpointUpdate{value: val, isSet: true}
}

func (v NullableServerlessTenantEndpointUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessTenantEndpointUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
