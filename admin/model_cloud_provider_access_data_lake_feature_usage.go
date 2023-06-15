// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderAccessDataLakeFeatureUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessDataLakeFeatureUsage{}

// CloudProviderAccessDataLakeFeatureUsage Details that describe the Atlas Data Lakes linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
type CloudProviderAccessDataLakeFeatureUsage struct {
	FeatureId *CloudProviderAccessFeatureUsageDataLakeFeatureId `json:"featureId,omitempty"`
	// Human-readable label that describes one MongoDB Cloud feature linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	FeatureType *string `json:"featureType,omitempty"`
}

// NewCloudProviderAccessDataLakeFeatureUsage instantiates a new CloudProviderAccessDataLakeFeatureUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessDataLakeFeatureUsage() *CloudProviderAccessDataLakeFeatureUsage {
	this := CloudProviderAccessDataLakeFeatureUsage{}
	return &this
}

// NewCloudProviderAccessDataLakeFeatureUsageWithDefaults instantiates a new CloudProviderAccessDataLakeFeatureUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessDataLakeFeatureUsageWithDefaults() *CloudProviderAccessDataLakeFeatureUsage {
	this := CloudProviderAccessDataLakeFeatureUsage{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *CloudProviderAccessDataLakeFeatureUsage) GetFeatureId() CloudProviderAccessFeatureUsageDataLakeFeatureId {
	if o == nil || IsNil(o.FeatureId) {
		var ret CloudProviderAccessFeatureUsageDataLakeFeatureId
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessDataLakeFeatureUsage) GetFeatureIdOk() (*CloudProviderAccessFeatureUsageDataLakeFeatureId, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CloudProviderAccessDataLakeFeatureUsage) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given CloudProviderAccessFeatureUsageDataLakeFeatureId and assigns it to the FeatureId field.
func (o *CloudProviderAccessDataLakeFeatureUsage) SetFeatureId(v CloudProviderAccessFeatureUsageDataLakeFeatureId) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *CloudProviderAccessDataLakeFeatureUsage) GetFeatureType() string {
	if o == nil || IsNil(o.FeatureType) {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessDataLakeFeatureUsage) GetFeatureTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureType) {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *CloudProviderAccessDataLakeFeatureUsage) HasFeatureType() bool {
	if o != nil && !IsNil(o.FeatureType) {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *CloudProviderAccessDataLakeFeatureUsage) SetFeatureType(v string) {
	o.FeatureType = &v
}

func (o CloudProviderAccessDataLakeFeatureUsage) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessDataLakeFeatureUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureId) {
		toSerialize["featureId"] = o.FeatureId
	}
	return toSerialize, nil
}

type NullableCloudProviderAccessDataLakeFeatureUsage struct {
	value *CloudProviderAccessDataLakeFeatureUsage
	isSet bool
}

func (v NullableCloudProviderAccessDataLakeFeatureUsage) Get() *CloudProviderAccessDataLakeFeatureUsage {
	return v.value
}

func (v *NullableCloudProviderAccessDataLakeFeatureUsage) Set(val *CloudProviderAccessDataLakeFeatureUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessDataLakeFeatureUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessDataLakeFeatureUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessDataLakeFeatureUsage(val *CloudProviderAccessDataLakeFeatureUsage) *NullableCloudProviderAccessDataLakeFeatureUsage {
	return &NullableCloudProviderAccessDataLakeFeatureUsage{value: val, isSet: true}
}

func (v NullableCloudProviderAccessDataLakeFeatureUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessDataLakeFeatureUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
