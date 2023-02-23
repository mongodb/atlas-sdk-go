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

// checks if the CloudProviderAccessFeatureUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessFeatureUsage{}

// CloudProviderAccessFeatureUsage MongoDB Cloud features associated with this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
type CloudProviderAccessFeatureUsage struct {
	// Human-readable label that describes one MongoDB Cloud feature linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	FeatureType *string `json:"featureType,omitempty"`
}

// NewCloudProviderAccessFeatureUsage instantiates a new CloudProviderAccessFeatureUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessFeatureUsage() *CloudProviderAccessFeatureUsage {
	this := CloudProviderAccessFeatureUsage{}
	return &this
}

// NewCloudProviderAccessFeatureUsageWithDefaults instantiates a new CloudProviderAccessFeatureUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessFeatureUsageWithDefaults() *CloudProviderAccessFeatureUsage {
	this := CloudProviderAccessFeatureUsage{}
	return &this
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *CloudProviderAccessFeatureUsage) GetFeatureType() string {
	if o == nil || IsNil(o.FeatureType) {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsage) GetFeatureTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureType) {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsage) HasFeatureType() bool {
	if o != nil && !IsNil(o.FeatureType) {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *CloudProviderAccessFeatureUsage) SetFeatureType(v string) {
	o.FeatureType = &v
}

func (o CloudProviderAccessFeatureUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudProviderAccessFeatureUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: featureType is readOnly
	return toSerialize, nil
}

type NullableCloudProviderAccessFeatureUsage struct {
	value *CloudProviderAccessFeatureUsage
	isSet bool
}

func (v NullableCloudProviderAccessFeatureUsage) Get() *CloudProviderAccessFeatureUsage {
	return v.value
}

func (v *NullableCloudProviderAccessFeatureUsage) Set(val *CloudProviderAccessFeatureUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessFeatureUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessFeatureUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessFeatureUsage(val *CloudProviderAccessFeatureUsage) *NullableCloudProviderAccessFeatureUsage {
	return &NullableCloudProviderAccessFeatureUsage{value: val, isSet: true}
}

func (v NullableCloudProviderAccessFeatureUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessFeatureUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


