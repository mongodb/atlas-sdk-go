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

// checks if the ServerlessBackupOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerlessBackupOptions{}

// ServerlessBackupOptions Group of settings that configure serverless backup.
type ServerlessBackupOptions struct {
	// Flag that indicates whether the serverless instance uses **Serverless Continuous Backup**.  If this parameter is `false`, the serverless instance uses **Basic Backup**.   | Option | Description |  |---|---|  | Serverless Continuous Backup | Atlas takes incremental [snapshots](https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#std-label-serverless-snapshots) of the data in your serverless instance every six hours and lets you restore the data from a selected point in time within the last 72 hours. Atlas also takes daily snapshots and retains these daily snapshots for 35 days. To learn more, see [Serverless Instance Costs](https://www.mongodb.com/docs/atlas/billing/serverless-instance-costs/#std-label-serverless-instance-costs). |  | Basic Backup | Atlas takes incremental [snapshots](https://www.mongodb.com/docs/atlas/backup/cloud-backup/overview/#std-label-serverless-snapshots) of the data in your serverless instance every six hours and retains only the two most recent snapshots. You can use this option for free. |
	ServerlessContinuousBackupEnabled *bool `json:"serverlessContinuousBackupEnabled,omitempty"`
}

// NewServerlessBackupOptions instantiates a new ServerlessBackupOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerlessBackupOptions() *ServerlessBackupOptions {
	this := ServerlessBackupOptions{}
	var serverlessContinuousBackupEnabled bool = true
	this.ServerlessContinuousBackupEnabled = &serverlessContinuousBackupEnabled
	return &this
}

// NewServerlessBackupOptionsWithDefaults instantiates a new ServerlessBackupOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerlessBackupOptionsWithDefaults() *ServerlessBackupOptions {
	this := ServerlessBackupOptions{}
	var serverlessContinuousBackupEnabled bool = true
	this.ServerlessContinuousBackupEnabled = &serverlessContinuousBackupEnabled
	return &this
}

// GetServerlessContinuousBackupEnabled returns the ServerlessContinuousBackupEnabled field value if set, zero value otherwise.
func (o *ServerlessBackupOptions) GetServerlessContinuousBackupEnabled() bool {
	if o == nil || IsNil(o.ServerlessContinuousBackupEnabled) {
		var ret bool
		return ret
	}
	return *o.ServerlessContinuousBackupEnabled
}

// GetServerlessContinuousBackupEnabledOk returns a tuple with the ServerlessContinuousBackupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerlessBackupOptions) GetServerlessContinuousBackupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ServerlessContinuousBackupEnabled) {
		return nil, false
	}
	return o.ServerlessContinuousBackupEnabled, true
}

// HasServerlessContinuousBackupEnabled returns a boolean if a field has been set.
func (o *ServerlessBackupOptions) HasServerlessContinuousBackupEnabled() bool {
	if o != nil && !IsNil(o.ServerlessContinuousBackupEnabled) {
		return true
	}

	return false
}

// SetServerlessContinuousBackupEnabled gets a reference to the given bool and assigns it to the ServerlessContinuousBackupEnabled field.
func (o *ServerlessBackupOptions) SetServerlessContinuousBackupEnabled(v bool) {
	o.ServerlessContinuousBackupEnabled = &v
}

func (o ServerlessBackupOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerlessBackupOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerlessContinuousBackupEnabled) {
		toSerialize["serverlessContinuousBackupEnabled"] = o.ServerlessContinuousBackupEnabled
	}
	return toSerialize, nil
}

type NullableServerlessBackupOptions struct {
	value *ServerlessBackupOptions
	isSet bool
}

func (v NullableServerlessBackupOptions) Get() *ServerlessBackupOptions {
	return v.value
}

func (v *NullableServerlessBackupOptions) Set(val *ServerlessBackupOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableServerlessBackupOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableServerlessBackupOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerlessBackupOptions(val *ServerlessBackupOptions) *NullableServerlessBackupOptions {
	return &NullableServerlessBackupOptions{value: val, isSet: true}
}

func (v NullableServerlessBackupOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerlessBackupOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


