// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the CloudProviderRegions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudProviderRegions{}

// CloudProviderRegions struct for CloudProviderRegions
type CloudProviderRegions struct {
	// List of instances sizes that this cloud provider supports.
	InstanceSizes []string `json:"instanceSizes,omitempty"`
	// Human-readable label that identifies the Cloud provider.
	Provider *string `json:"provider,omitempty"`
}

// NewCloudProviderRegions instantiates a new CloudProviderRegions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudProviderRegions() *CloudProviderRegions {
	this := CloudProviderRegions{}
	return &this
}

// NewCloudProviderRegionsWithDefaults instantiates a new CloudProviderRegions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudProviderRegionsWithDefaults() *CloudProviderRegions {
	this := CloudProviderRegions{}
	return &this
}

// GetInstanceSizes returns the InstanceSizes field value if set, zero value otherwise.
func (o *CloudProviderRegions) GetInstanceSizes() []string {
	if o == nil || IsNil(o.InstanceSizes) {
		var ret []string
		return ret
	}
	return o.InstanceSizes
}

// GetInstanceSizesOk returns a tuple with the InstanceSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderRegions) GetInstanceSizesOk() ([]string, bool) {
	if o == nil || IsNil(o.InstanceSizes) {
		return nil, false
	}
	return o.InstanceSizes, true
}

// HasInstanceSizes returns a boolean if a field has been set.
func (o *CloudProviderRegions) HasInstanceSizes() bool {
	if o != nil && !IsNil(o.InstanceSizes) {
		return true
	}

	return false
}

// SetInstanceSizes gets a reference to the given []string and assigns it to the InstanceSizes field.
func (o *CloudProviderRegions) SetInstanceSizes(v []string) {
	o.InstanceSizes = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CloudProviderRegions) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudProviderRegions) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CloudProviderRegions) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CloudProviderRegions) SetProvider(v string) {
	o.Provider = &v
}

func (o CloudProviderRegions) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o CloudProviderRegions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	return toSerialize, nil
}

type NullableCloudProviderRegions struct {
	value *CloudProviderRegions
	isSet bool
}

func (v NullableCloudProviderRegions) Get() *CloudProviderRegions {
	return v.value
}

func (v *NullableCloudProviderRegions) Set(val *CloudProviderRegions) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudProviderRegions) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudProviderRegions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudProviderRegions(val *CloudProviderRegions) *NullableCloudProviderRegions {
	return &NullableCloudProviderRegions{value: val, isSet: true}
}

func (v NullableCloudProviderRegions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudProviderRegions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
