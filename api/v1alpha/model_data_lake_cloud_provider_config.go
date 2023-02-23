/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"encoding/json"
)

// checks if the DataLakeCloudProviderConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataLakeCloudProviderConfig{}

// DataLakeCloudProviderConfig Cloud provider linked to this data lake.
type DataLakeCloudProviderConfig struct {
	Aws DataLakeAWSCloudProviderConfig `json:"aws"`
}

// NewDataLakeCloudProviderConfig instantiates a new DataLakeCloudProviderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataLakeCloudProviderConfig() *DataLakeCloudProviderConfig {
	this := DataLakeCloudProviderConfig{}
	return &this
}

// NewDataLakeCloudProviderConfigWithDefaults instantiates a new DataLakeCloudProviderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataLakeCloudProviderConfigWithDefaults() *DataLakeCloudProviderConfig {
	this := DataLakeCloudProviderConfig{}
	return &this
}

// GetAws returns the Aws field value
func (o *DataLakeCloudProviderConfig) GetAws() DataLakeAWSCloudProviderConfig {
	if o == nil {
		var ret DataLakeAWSCloudProviderConfig
		return ret
	}

	return o.Aws
}

// GetAwsOk returns a tuple with the Aws field value
// and a boolean to check if the value has been set.
func (o *DataLakeCloudProviderConfig) GetAwsOk() (*DataLakeAWSCloudProviderConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aws, true
}

// SetAws sets field value
func (o *DataLakeCloudProviderConfig) SetAws(v DataLakeAWSCloudProviderConfig) {
	o.Aws = v
}

func (o DataLakeCloudProviderConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataLakeCloudProviderConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aws"] = o.Aws
	return toSerialize, nil
}

type NullableDataLakeCloudProviderConfig struct {
	value *DataLakeCloudProviderConfig
	isSet bool
}

func (v NullableDataLakeCloudProviderConfig) Get() *DataLakeCloudProviderConfig {
	return v.value
}

func (v *NullableDataLakeCloudProviderConfig) Set(val *DataLakeCloudProviderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDataLakeCloudProviderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDataLakeCloudProviderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataLakeCloudProviderConfig(val *DataLakeCloudProviderConfig) *NullableDataLakeCloudProviderConfig {
	return &NullableDataLakeCloudProviderConfig{value: val, isSet: true}
}

func (v NullableDataLakeCloudProviderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataLakeCloudProviderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


