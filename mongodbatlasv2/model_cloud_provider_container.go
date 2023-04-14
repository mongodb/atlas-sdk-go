/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"fmt"
)

// CloudProviderContainer - Collection of settings that configures the network container for a virtual private connection on Amazon Web Services.
type CloudProviderContainer struct {
	AWSCloudProviderContainer *AWSCloudProviderContainer
	AzureCloudProviderContainer *AzureCloudProviderContainer
	GCPCloudProviderContainer *GCPCloudProviderContainer
}

// AWSCloudProviderContainerAsCloudProviderContainer is a convenience function that returns AWSCloudProviderContainer wrapped in CloudProviderContainer
func AWSCloudProviderContainerAsCloudProviderContainer(v *AWSCloudProviderContainer) CloudProviderContainer {
	return CloudProviderContainer{
		AWSCloudProviderContainer: v,
	}
}

// AzureCloudProviderContainerAsCloudProviderContainer is a convenience function that returns AzureCloudProviderContainer wrapped in CloudProviderContainer
func AzureCloudProviderContainerAsCloudProviderContainer(v *AzureCloudProviderContainer) CloudProviderContainer {
	return CloudProviderContainer{
		AzureCloudProviderContainer: v,
	}
}

// GCPCloudProviderContainerAsCloudProviderContainer is a convenience function that returns GCPCloudProviderContainer wrapped in CloudProviderContainer
func GCPCloudProviderContainerAsCloudProviderContainer(v *GCPCloudProviderContainer) CloudProviderContainer {
	return CloudProviderContainer{
		GCPCloudProviderContainer: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CloudProviderContainer) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'AWS'
	if jsonDict["providerName"] == "AWS" {
		// try to unmarshal JSON data into AWSCloudProviderContainer
		err = json.Unmarshal(data, &dst.AWSCloudProviderContainer)
		if err == nil {
			return nil // data stored in dst.AWSCloudProviderContainer, return on the first match
		} else {
			dst.AWSCloudProviderContainer = nil
			return fmt.Errorf("failed to unmarshal CloudProviderContainer as AWSCloudProviderContainer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AWSCloudProviderContainer'
	if jsonDict["providerName"] == "AWSCloudProviderContainer" {
		// try to unmarshal JSON data into AWSCloudProviderContainer
		err = json.Unmarshal(data, &dst.AWSCloudProviderContainer)
		if err == nil {
			return nil // data stored in dst.AWSCloudProviderContainer, return on the first match
		} else {
			dst.AWSCloudProviderContainer = nil
			return fmt.Errorf("failed to unmarshal CloudProviderContainer as AWSCloudProviderContainer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AZURE'
	if jsonDict["providerName"] == "AZURE" {
		// try to unmarshal JSON data into AzureCloudProviderContainer
		err = json.Unmarshal(data, &dst.AzureCloudProviderContainer)
		if err == nil {
			return nil // data stored in dst.AzureCloudProviderContainer, return on the first match
		} else {
			dst.AzureCloudProviderContainer = nil
			return fmt.Errorf("failed to unmarshal CloudProviderContainer as AzureCloudProviderContainer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'AzureCloudProviderContainer'
	if jsonDict["providerName"] == "AzureCloudProviderContainer" {
		// try to unmarshal JSON data into AzureCloudProviderContainer
		err = json.Unmarshal(data, &dst.AzureCloudProviderContainer)
		if err == nil {
			return nil // data stored in dst.AzureCloudProviderContainer, return on the first match
		} else {
			dst.AzureCloudProviderContainer = nil
			return fmt.Errorf("failed to unmarshal CloudProviderContainer as AzureCloudProviderContainer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCP'
	if jsonDict["providerName"] == "GCP" {
		// try to unmarshal JSON data into GCPCloudProviderContainer
		err = json.Unmarshal(data, &dst.GCPCloudProviderContainer)
		if err == nil {
			return nil // data stored in dst.GCPCloudProviderContainer, return on the first match
		} else {
			dst.GCPCloudProviderContainer = nil
			return fmt.Errorf("failed to unmarshal CloudProviderContainer as GCPCloudProviderContainer: %s", err.Error())
		}
	}

	// check if the discriminator value is 'GCPCloudProviderContainer'
	if jsonDict["providerName"] == "GCPCloudProviderContainer" {
		// try to unmarshal JSON data into GCPCloudProviderContainer
		err = json.Unmarshal(data, &dst.GCPCloudProviderContainer)
		if err == nil {
			return nil // data stored in dst.GCPCloudProviderContainer, return on the first match
		} else {
			dst.GCPCloudProviderContainer = nil
			return fmt.Errorf("failed to unmarshal CloudProviderContainer as GCPCloudProviderContainer: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CloudProviderContainer) MarshalJSON() ([]byte, error) {
	if src.AWSCloudProviderContainer != nil {
		return json.Marshal(&src.AWSCloudProviderContainer)
	}

	if src.AzureCloudProviderContainer != nil {
		return json.Marshal(&src.AzureCloudProviderContainer)
	}

	if src.GCPCloudProviderContainer != nil {
		return json.Marshal(&src.GCPCloudProviderContainer)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CloudProviderContainer) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AWSCloudProviderContainer != nil {
		return obj.AWSCloudProviderContainer
	}

	if obj.AzureCloudProviderContainer != nil {
		return obj.AzureCloudProviderContainer
	}

	if obj.GCPCloudProviderContainer != nil {
		return obj.GCPCloudProviderContainer
	}

	// all schemas are nil
	return nil
}

type NullableCloudProviderContainer struct {
	value *CloudProviderContainer
	isSet bool
}

func (v NullableCloudProviderContainer) Get() *CloudProviderContainer {
	return v.value
}

func (v *NullableCloudProviderContainer) Set(val *CloudProviderContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderContainer(val *CloudProviderContainer) *NullableCloudProviderContainer {
	return &NullableCloudProviderContainer{value: val, isSet: true}
}

func (v NullableCloudProviderContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


