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

// checks if the CloudProviderAccessExportSnapshotFeatureUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessExportSnapshotFeatureUsage{}

// CloudProviderAccessExportSnapshotFeatureUsage Details that describe the Amazon Web Services (AWS) Simple Storage Service (S3) export buckets linked to this AWS Identity and Access Management (IAM) role.
type CloudProviderAccessExportSnapshotFeatureUsage struct {
	FeatureId *CloudProviderAccessFeatureUsageExportSnapshotFeatureId `json:"featureId,omitempty"`
}

// NewCloudProviderAccessExportSnapshotFeatureUsage instantiates a new CloudProviderAccessExportSnapshotFeatureUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessExportSnapshotFeatureUsage() *CloudProviderAccessExportSnapshotFeatureUsage {
	this := CloudProviderAccessExportSnapshotFeatureUsage{}
	return &this
}

// NewCloudProviderAccessExportSnapshotFeatureUsageWithDefaults instantiates a new CloudProviderAccessExportSnapshotFeatureUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessExportSnapshotFeatureUsageWithDefaults() *CloudProviderAccessExportSnapshotFeatureUsage {
	this := CloudProviderAccessExportSnapshotFeatureUsage{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CloudProviderAccessExportSnapshotFeatureUsage) GetFeatureId() CloudProviderAccessFeatureUsageExportSnapshotFeatureId {
	if o == nil || IsNil(o.FeatureId) {
		var ret CloudProviderAccessFeatureUsageExportSnapshotFeatureId
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessExportSnapshotFeatureUsage) GetFeatureIdOk() (*CloudProviderAccessFeatureUsageExportSnapshotFeatureId, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CloudProviderAccessExportSnapshotFeatureUsage) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given CloudProviderAccessFeatureUsageExportSnapshotFeatureId and assigns it to the FeatureId field.
func (o *CloudProviderAccessExportSnapshotFeatureUsage) SetFeatureId(v CloudProviderAccessFeatureUsageExportSnapshotFeatureId) {
	o.FeatureId = &v
}

func (o CloudProviderAccessExportSnapshotFeatureUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProviderAccessExportSnapshotFeatureUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["featureId"] = o.FeatureId
	}
	return toSerialize, nil
}

type NullableCloudProviderAccessExportSnapshotFeatureUsage struct {
	value *CloudProviderAccessExportSnapshotFeatureUsage
	isSet bool
}

func (v NullableCloudProviderAccessExportSnapshotFeatureUsage) Get() *CloudProviderAccessExportSnapshotFeatureUsage {
	return v.value
}

func (v *NullableCloudProviderAccessExportSnapshotFeatureUsage) Set(val *CloudProviderAccessExportSnapshotFeatureUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessExportSnapshotFeatureUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessExportSnapshotFeatureUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessExportSnapshotFeatureUsage(val *CloudProviderAccessExportSnapshotFeatureUsage) *NullableCloudProviderAccessExportSnapshotFeatureUsage {
	return &NullableCloudProviderAccessExportSnapshotFeatureUsage{value: val, isSet: true}
}

func (v NullableCloudProviderAccessExportSnapshotFeatureUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessExportSnapshotFeatureUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


