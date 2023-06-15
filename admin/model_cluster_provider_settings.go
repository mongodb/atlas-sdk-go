// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"fmt"
)

// ClusterProviderSettings - Group of cloud provider settings that configure the provisioned MongoDB hosts.
type ClusterProviderSettings struct {
	AWSCloudProviderSettings    *AWSCloudProviderSettings
	AzureCloudProviderSettings  *AzureCloudProviderSettings
	CloudGCPProviderSettings    *CloudGCPProviderSettings
	ClusterFreeProviderSettings *ClusterFreeProviderSettings
}

// AWSCloudProviderSettingsAsClusterProviderSettings is a convenience function that returns AWSCloudProviderSettings wrapped in ClusterProviderSettings
func AWSCloudProviderSettingsAsClusterProviderSettings(v *AWSCloudProviderSettings) ClusterProviderSettings {
	return ClusterProviderSettings{
		AWSCloudProviderSettings: v,
	}
}

// AzureCloudProviderSettingsAsClusterProviderSettings is a convenience function that returns AzureCloudProviderSettings wrapped in ClusterProviderSettings
func AzureCloudProviderSettingsAsClusterProviderSettings(v *AzureCloudProviderSettings) ClusterProviderSettings {
	return ClusterProviderSettings{
		AzureCloudProviderSettings: v,
	}
}

// CloudGCPProviderSettingsAsClusterProviderSettings is a convenience function that returns CloudGCPProviderSettings wrapped in ClusterProviderSettings
func CloudGCPProviderSettingsAsClusterProviderSettings(v *CloudGCPProviderSettings) ClusterProviderSettings {
	return ClusterProviderSettings{
		CloudGCPProviderSettings: v,
	}
}

// ClusterFreeProviderSettingsAsClusterProviderSettings is a convenience function that returns ClusterFreeProviderSettings wrapped in ClusterProviderSettings
func ClusterFreeProviderSettingsAsClusterProviderSettings(v *ClusterFreeProviderSettings) ClusterProviderSettings {
	return ClusterProviderSettings{
		ClusterFreeProviderSettings: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ClusterProviderSettings) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into AWSCloudProviderSettings
		err = json.Unmarshal(data, &dst.AWSCloudProviderSettings)
		if err == nil {
			return nil // data stored in dst.AWSCloudProviderSettings, return on the first match
		} else {
			dst.AWSCloudProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as AWSCloudProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWSCloudProviderSettings'
	if jsonDict["providerName"] == "AWSCloudProviderSettings" {
		// try to unmarshal JSON data into AWSCloudProviderSettings
		err = json.Unmarshal(data, &dst.AWSCloudProviderSettings)
		if err == nil {
			return nil // data stored in dst.AWSCloudProviderSettings, return on the first match
		} else {
			dst.AWSCloudProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as AWSCloudProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into AzureCloudProviderSettings
		err = json.Unmarshal(data, &dst.AzureCloudProviderSettings)
		if err == nil {
			return nil // data stored in dst.AzureCloudProviderSettings, return on the first match
		} else {
			dst.AzureCloudProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as AzureCloudProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzureCloudProviderSettings'
	if jsonDict["providerName"] == "AzureCloudProviderSettings" {
		// try to unmarshal JSON data into AzureCloudProviderSettings
		err = json.Unmarshal(data, &dst.AzureCloudProviderSettings)
		if err == nil {
			return nil // data stored in dst.AzureCloudProviderSettings, return on the first match
		} else {
			dst.AzureCloudProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as AzureCloudProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudGCPProviderSettings'
	if jsonDict["providerName"] == "CloudGCPProviderSettings" {
		// try to unmarshal JSON data into CloudGCPProviderSettings
		err = json.Unmarshal(data, &dst.CloudGCPProviderSettings)
		if err == nil {
			return nil // data stored in dst.CloudGCPProviderSettings, return on the first match
		} else {
			dst.CloudGCPProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as CloudGCPProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ClusterFreeProviderSettings'
	if jsonDict["providerName"] == "ClusterFreeProviderSettings" {
		// try to unmarshal JSON data into ClusterFreeProviderSettings
		err = json.Unmarshal(data, &dst.ClusterFreeProviderSettings)
		if err == nil {
			return nil // data stored in dst.ClusterFreeProviderSettings, return on the first match
		} else {
			dst.ClusterFreeProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as ClusterFreeProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCP'
	if jsonDict["providerName"] == "GCP" {
		// try to unmarshal JSON data into CloudGCPProviderSettings
		err = json.Unmarshal(data, &dst.CloudGCPProviderSettings)
		if err == nil {
			return nil // data stored in dst.CloudGCPProviderSettings, return on the first match
		} else {
			dst.CloudGCPProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as CloudGCPProviderSettings: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TENANT'
	if jsonDict["providerName"] == "TENANT" {
		// try to unmarshal JSON data into ClusterFreeProviderSettings
		err = json.Unmarshal(data, &dst.ClusterFreeProviderSettings)
		if err == nil {
			return nil // data stored in dst.ClusterFreeProviderSettings, return on the first match
		} else {
			dst.ClusterFreeProviderSettings = nil
			return fmt.Errorf("failed to unmarshal ClusterProviderSettings as ClusterFreeProviderSettings: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ClusterProviderSettings) MarshalJSON() ([]byte, error) {
	if src.AWSCloudProviderSettings != nil {
		return json.Marshal(&src.AWSCloudProviderSettings)
	}

	if src.AzureCloudProviderSettings != nil {
		return json.Marshal(&src.AzureCloudProviderSettings)
	}

	if src.CloudGCPProviderSettings != nil {
		return json.Marshal(&src.CloudGCPProviderSettings)
	}

	if src.ClusterFreeProviderSettings != nil {
		return json.Marshal(&src.ClusterFreeProviderSettings)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ClusterProviderSettings) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSCloudProviderSettings != nil {
		return obj.AWSCloudProviderSettings
	}

	if obj.AzureCloudProviderSettings != nil {
		return obj.AzureCloudProviderSettings
	}

	if obj.CloudGCPProviderSettings != nil {
		return obj.CloudGCPProviderSettings
	}

	if obj.ClusterFreeProviderSettings != nil {
		return obj.ClusterFreeProviderSettings
	}

	// all schemas are nil
	return nil
}

type NullableClusterProviderSettings struct {
	value *ClusterProviderSettings
	isSet bool
}

func (v NullableClusterProviderSettings) Get() *ClusterProviderSettings {
	return v.value
}

func (v *NullableClusterProviderSettings) Set(val *ClusterProviderSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterProviderSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterProviderSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterProviderSettings(val *ClusterProviderSettings) *NullableClusterProviderSettings {
	return &NullableClusterProviderSettings{value: val, isSet: true}
}

func (v NullableClusterProviderSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterProviderSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
