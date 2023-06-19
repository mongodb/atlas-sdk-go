// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the AvailableCloudProviderRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableCloudProviderRegion{}

// AvailableCloudProviderRegion List of regions that this cloud provider supports for this instance size.
type AvailableCloudProviderRegion struct {
	// Flag that indicates whether the cloud provider sets this region as its default. AWS defaults to US_EAST_1, GCP defaults to CENTRAL_US, and AZURE defaults to US_WEST_2.
	Default *bool `json:"default,omitempty"`
	// Human-readable label that identifies the supported region.
	Name *string `json:"name,omitempty"`
}

// NewAvailableCloudProviderRegion instantiates a new AvailableCloudProviderRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableCloudProviderRegion() *AvailableCloudProviderRegion {
	this := AvailableCloudProviderRegion{}
	return &this
}

// NewAvailableCloudProviderRegionWithDefaults instantiates a new AvailableCloudProviderRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableCloudProviderRegionWithDefaults() *AvailableCloudProviderRegion {
	this := AvailableCloudProviderRegion{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *AvailableCloudProviderRegion) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableCloudProviderRegion) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *AvailableCloudProviderRegion) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *AvailableCloudProviderRegion) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AvailableCloudProviderRegion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableCloudProviderRegion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AvailableCloudProviderRegion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AvailableCloudProviderRegion) SetName(v string) {
	o.Name = &v
}

func (o AvailableCloudProviderRegion) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AvailableCloudProviderRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAvailableCloudProviderRegion struct {
	value *AvailableCloudProviderRegion
	isSet bool
}

func (v NullableAvailableCloudProviderRegion) Get() *AvailableCloudProviderRegion {
	return v.value
}

func (v *NullableAvailableCloudProviderRegion) Set(val *AvailableCloudProviderRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableCloudProviderRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableCloudProviderRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableCloudProviderRegion(val *AvailableCloudProviderRegion) *NullableAvailableCloudProviderRegion {
	return &NullableAvailableCloudProviderRegion{value: val, isSet: true}
}

func (v NullableAvailableCloudProviderRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableCloudProviderRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
