// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderAccessEncryptionAtRestFeatureUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessEncryptionAtRestFeatureUsage{}

// CloudProviderAccessEncryptionAtRestFeatureUsage Details that describe the Key Management Service (KMS) linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
type CloudProviderAccessEncryptionAtRestFeatureUsage struct {
	// Object that contains the identifying characteristics of the Amazon Web Services (AWS) Key Management Service (KMS). This field always returns a null value.
	FeatureId map[string]interface{} `json:"featureId,omitempty"`
	// Human-readable label that describes one MongoDB Cloud feature linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
	FeatureType *string `json:"featureType,omitempty"`
}

// NewCloudProviderAccessEncryptionAtRestFeatureUsage instantiates a new CloudProviderAccessEncryptionAtRestFeatureUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessEncryptionAtRestFeatureUsage() *CloudProviderAccessEncryptionAtRestFeatureUsage {
	this := CloudProviderAccessEncryptionAtRestFeatureUsage{}
	return &this
}

// NewCloudProviderAccessEncryptionAtRestFeatureUsageWithDefaults instantiates a new CloudProviderAccessEncryptionAtRestFeatureUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessEncryptionAtRestFeatureUsageWithDefaults() *CloudProviderAccessEncryptionAtRestFeatureUsage {
	this := CloudProviderAccessEncryptionAtRestFeatureUsage{}
	return &this
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureId() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureIdOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return map[string]interface{}{}, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) HasFeatureId() bool {
	if o != nil && IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given map[string]interface{} and assigns it to the FeatureId field.
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) SetFeatureId(v map[string]interface{}) {
	o.FeatureId = v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureType() string {
	if o == nil || IsNil(o.FeatureType) {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) GetFeatureTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FeatureType) {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) HasFeatureType() bool {
	if o != nil && !IsNil(o.FeatureType) {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *CloudProviderAccessEncryptionAtRestFeatureUsage) SetFeatureType(v string) {
	o.FeatureType = &v
}

func (o CloudProviderAccessEncryptionAtRestFeatureUsage) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessEncryptionAtRestFeatureUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	return toSerialize, nil
}

type NullableCloudProviderAccessEncryptionAtRestFeatureUsage struct {
	value *CloudProviderAccessEncryptionAtRestFeatureUsage
	isSet bool
}

func (v NullableCloudProviderAccessEncryptionAtRestFeatureUsage) Get() *CloudProviderAccessEncryptionAtRestFeatureUsage {
	return v.value
}

func (v *NullableCloudProviderAccessEncryptionAtRestFeatureUsage) Set(val *CloudProviderAccessEncryptionAtRestFeatureUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessEncryptionAtRestFeatureUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessEncryptionAtRestFeatureUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessEncryptionAtRestFeatureUsage(val *CloudProviderAccessEncryptionAtRestFeatureUsage) *NullableCloudProviderAccessEncryptionAtRestFeatureUsage {
	return &NullableCloudProviderAccessEncryptionAtRestFeatureUsage{value: val, isSet: true}
}

func (v NullableCloudProviderAccessEncryptionAtRestFeatureUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessEncryptionAtRestFeatureUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
