// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
	"fmt"
)

// RegionConfig - Cloud service provider on which MongoDB Cloud provisions the hosts.
type RegionConfig struct {
	AWSRegionConfig    *AWSRegionConfig
	AzureRegionConfig  *AzureRegionConfig
	GCPRegionConfig    *GCPRegionConfig
	TenantRegionConfig *TenantRegionConfig
}

// AWSRegionConfigAsRegionConfig is a convenience function that returns AWSRegionConfig wrapped in RegionConfig
func AWSRegionConfigAsRegionConfig(v *AWSRegionConfig) RegionConfig {
	return RegionConfig{
		AWSRegionConfig: v,
	}
}

// AzureRegionConfigAsRegionConfig is a convenience function that returns AzureRegionConfig wrapped in RegionConfig
func AzureRegionConfigAsRegionConfig(v *AzureRegionConfig) RegionConfig {
	return RegionConfig{
		AzureRegionConfig: v,
	}
}

// GCPRegionConfigAsRegionConfig is a convenience function that returns GCPRegionConfig wrapped in RegionConfig
func GCPRegionConfigAsRegionConfig(v *GCPRegionConfig) RegionConfig {
	return RegionConfig{
		GCPRegionConfig: v,
	}
}

// TenantRegionConfigAsRegionConfig is a convenience function that returns TenantRegionConfig wrapped in RegionConfig
func TenantRegionConfigAsRegionConfig(v *TenantRegionConfig) RegionConfig {
	return RegionConfig{
		TenantRegionConfig: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RegionConfig) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into AWSRegionConfig
		err = json.Unmarshal(data, &dst.AWSRegionConfig)
		if err == nil {
			return nil // data stored in dst.AWSRegionConfig, return on the first match
		} else {
			dst.AWSRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as AWSRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWSRegionConfig'
	if jsonDict["providerName"] == "AWSRegionConfig" {
		// try to unmarshal JSON data into AWSRegionConfig
		err = json.Unmarshal(data, &dst.AWSRegionConfig)
		if err == nil {
			return nil // data stored in dst.AWSRegionConfig, return on the first match
		} else {
			dst.AWSRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as AWSRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into AzureRegionConfig
		err = json.Unmarshal(data, &dst.AzureRegionConfig)
		if err == nil {
			return nil // data stored in dst.AzureRegionConfig, return on the first match
		} else {
			dst.AzureRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as AzureRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzureRegionConfig'
	if jsonDict["providerName"] == "AzureRegionConfig" {
		// try to unmarshal JSON data into AzureRegionConfig
		err = json.Unmarshal(data, &dst.AzureRegionConfig)
		if err == nil {
			return nil // data stored in dst.AzureRegionConfig, return on the first match
		} else {
			dst.AzureRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as AzureRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCP'
	if jsonDict["providerName"] == "GCP" {
		// try to unmarshal JSON data into GCPRegionConfig
		err = json.Unmarshal(data, &dst.GCPRegionConfig)
		if err == nil {
			return nil // data stored in dst.GCPRegionConfig, return on the first match
		} else {
			dst.GCPRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as GCPRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCPRegionConfig'
	if jsonDict["providerName"] == "GCPRegionConfig" {
		// try to unmarshal JSON data into GCPRegionConfig
		err = json.Unmarshal(data, &dst.GCPRegionConfig)
		if err == nil {
			return nil // data stored in dst.GCPRegionConfig, return on the first match
		} else {
			dst.GCPRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as GCPRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TENANT'
	if jsonDict["providerName"] == "TENANT" {
		// try to unmarshal JSON data into TenantRegionConfig
		err = json.Unmarshal(data, &dst.TenantRegionConfig)
		if err == nil {
			return nil // data stored in dst.TenantRegionConfig, return on the first match
		} else {
			dst.TenantRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as TenantRegionConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'TenantRegionConfig'
	if jsonDict["providerName"] == "TenantRegionConfig" {
		// try to unmarshal JSON data into TenantRegionConfig
		err = json.Unmarshal(data, &dst.TenantRegionConfig)
		if err == nil {
			return nil // data stored in dst.TenantRegionConfig, return on the first match
		} else {
			dst.TenantRegionConfig = nil
			return fmt.Errorf("failed to unmarshal RegionConfig as TenantRegionConfig: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RegionConfig) MarshalJSON() ([]byte, error) {
	if src.AWSRegionConfig != nil {
		return json.Marshal(&src.AWSRegionConfig)
	}

	if src.AzureRegionConfig != nil {
		return json.Marshal(&src.AzureRegionConfig)
	}

	if src.GCPRegionConfig != nil {
		return json.Marshal(&src.GCPRegionConfig)
	}

	if src.TenantRegionConfig != nil {
		return json.Marshal(&src.TenantRegionConfig)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RegionConfig) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AWSRegionConfig != nil {
		return obj.AWSRegionConfig
	}

	if obj.AzureRegionConfig != nil {
		return obj.AzureRegionConfig
	}

	if obj.GCPRegionConfig != nil {
		return obj.GCPRegionConfig
	}

	if obj.TenantRegionConfig != nil {
		return obj.TenantRegionConfig
	}

	// all schemas are nil
	return nil
}

type NullableRegionConfig struct {
	value *RegionConfig
	isSet bool
}

func (v NullableRegionConfig) Get() *RegionConfig {
	return v.value
}

func (v *NullableRegionConfig) Set(val *RegionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionConfig(val *RegionConfig) *NullableRegionConfig {
	return &NullableRegionConfig{value: val, isSet: true}
}

func (v NullableRegionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
