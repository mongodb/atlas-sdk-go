// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderAccessFeatureUsageDataLakeFeatureId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderAccessFeatureUsageDataLakeFeatureId{}

// CloudProviderAccessFeatureUsageDataLakeFeatureId Identifying characteristics about the data lake linked to this Amazon Web Services (AWS) Identity and Access Management (IAM) role.
type CloudProviderAccessFeatureUsageDataLakeFeatureId struct {
	// Unique 24-hexadecimal digit string that identifies your project.
	GroupId *string `json:"groupId,omitempty"`
	// Human-readable label that identifies the data lake.
	Name *string `json:"name,omitempty"`
}

// NewCloudProviderAccessFeatureUsageDataLakeFeatureId instantiates a new CloudProviderAccessFeatureUsageDataLakeFeatureId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderAccessFeatureUsageDataLakeFeatureId() *CloudProviderAccessFeatureUsageDataLakeFeatureId {
	this := CloudProviderAccessFeatureUsageDataLakeFeatureId{}
	return &this
}

// NewCloudProviderAccessFeatureUsageDataLakeFeatureIdWithDefaults instantiates a new CloudProviderAccessFeatureUsageDataLakeFeatureId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderAccessFeatureUsageDataLakeFeatureIdWithDefaults() *CloudProviderAccessFeatureUsageDataLakeFeatureId {
	this := CloudProviderAccessFeatureUsageDataLakeFeatureId{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudProviderAccessFeatureUsageDataLakeFeatureId) SetName(v string) {
	o.Name = &v
}

func (o CloudProviderAccessFeatureUsageDataLakeFeatureId) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderAccessFeatureUsageDataLakeFeatureId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCloudProviderAccessFeatureUsageDataLakeFeatureId struct {
	value *CloudProviderAccessFeatureUsageDataLakeFeatureId
	isSet bool
}

func (v NullableCloudProviderAccessFeatureUsageDataLakeFeatureId) Get() *CloudProviderAccessFeatureUsageDataLakeFeatureId {
	return v.value
}

func (v *NullableCloudProviderAccessFeatureUsageDataLakeFeatureId) Set(val *CloudProviderAccessFeatureUsageDataLakeFeatureId) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderAccessFeatureUsageDataLakeFeatureId) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderAccessFeatureUsageDataLakeFeatureId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderAccessFeatureUsageDataLakeFeatureId(val *CloudProviderAccessFeatureUsageDataLakeFeatureId) *NullableCloudProviderAccessFeatureUsageDataLakeFeatureId {
	return &NullableCloudProviderAccessFeatureUsageDataLakeFeatureId{value: val, isSet: true}
}

func (v NullableCloudProviderAccessFeatureUsageDataLakeFeatureId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderAccessFeatureUsageDataLakeFeatureId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
