// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// CloudProviderAccessRole - Cloud provider access role.
type CloudProviderAccessRole struct {
	CloudProviderAccessAWSIAMRole            *CloudProviderAccessAWSIAMRole
	CloudProviderAccessAzureServicePrincipal *CloudProviderAccessAzureServicePrincipal
}

// CloudProviderAccessAWSIAMRoleAsCloudProviderAccessRole is a convenience function that returns CloudProviderAccessAWSIAMRole wrapped in CloudProviderAccessRole
func CloudProviderAccessAWSIAMRoleAsCloudProviderAccessRole(v *CloudProviderAccessAWSIAMRole) CloudProviderAccessRole {
	return CloudProviderAccessRole{
		CloudProviderAccessAWSIAMRole: v,
	}
}

// CloudProviderAccessAzureServicePrincipalAsCloudProviderAccessRole is a convenience function that returns CloudProviderAccessAzureServicePrincipal wrapped in CloudProviderAccessRole
func CloudProviderAccessAzureServicePrincipalAsCloudProviderAccessRole(v *CloudProviderAccessAzureServicePrincipal) CloudProviderAccessRole {
	return CloudProviderAccessRole{
		CloudProviderAccessAzureServicePrincipal: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CloudProviderAccessRole) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into CloudProviderAccessAWSIAMRole
		err = json.Unmarshal(data, &dst.CloudProviderAccessAWSIAMRole)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessAWSIAMRole, return on the first match
		} else {
			dst.CloudProviderAccessAWSIAMRole = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessRole as CloudProviderAccessAWSIAMRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into CloudProviderAccessAzureServicePrincipal
		err = json.Unmarshal(data, &dst.CloudProviderAccessAzureServicePrincipal)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessAzureServicePrincipal, return on the first match
		} else {
			dst.CloudProviderAccessAzureServicePrincipal = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessRole as CloudProviderAccessAzureServicePrincipal: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudProviderAccessAWSIAMRole'
	if jsonDict["providerName"] == "CloudProviderAccessAWSIAMRole" {
		// try to unmarshal JSON data into CloudProviderAccessAWSIAMRole
		err = json.Unmarshal(data, &dst.CloudProviderAccessAWSIAMRole)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessAWSIAMRole, return on the first match
		} else {
			dst.CloudProviderAccessAWSIAMRole = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessRole as CloudProviderAccessAWSIAMRole: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudProviderAccessAzureServicePrincipal'
	if jsonDict["providerName"] == "CloudProviderAccessAzureServicePrincipal" {
		// try to unmarshal JSON data into CloudProviderAccessAzureServicePrincipal
		err = json.Unmarshal(data, &dst.CloudProviderAccessAzureServicePrincipal)
		if err == nil {
			return nil // data stored in dst.CloudProviderAccessAzureServicePrincipal, return on the first match
		} else {
			dst.CloudProviderAccessAzureServicePrincipal = nil
			return fmt.Errorf("failed to unmarshal CloudProviderAccessRole as CloudProviderAccessAzureServicePrincipal: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CloudProviderAccessRole) MarshalJSON() ([]byte, error) {
	if src.CloudProviderAccessAWSIAMRole != nil {
		return json.Marshal(&src.CloudProviderAccessAWSIAMRole)
	}

	if src.CloudProviderAccessAzureServicePrincipal != nil {
		return json.Marshal(&src.CloudProviderAccessAzureServicePrincipal)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CloudProviderAccessRole) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.CloudProviderAccessAWSIAMRole != nil {
		return obj.CloudProviderAccessAWSIAMRole
	}

	if obj.CloudProviderAccessAzureServicePrincipal != nil {
		return obj.CloudProviderAccessAzureServicePrincipal
	}

	// all schemas are nil
	return nil
}

type NullableCloudProviderAccessRole struct {
	value *CloudProviderAccessRole
	isSet bool
}

func (v NullableCloudProviderAccessRole) Get() *CloudProviderAccessRole {
	return v.value
}

func (v *NullableCloudProviderAccessRole) Set(val *CloudProviderAccessRole) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessRole) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessRole(val *CloudProviderAccessRole) *NullableCloudProviderAccessRole {
	return &NullableCloudProviderAccessRole{value: val, isSet: true}
}

func (v NullableCloudProviderAccessRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
