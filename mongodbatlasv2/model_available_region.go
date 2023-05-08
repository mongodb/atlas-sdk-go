/*
API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
)

// checks if the AvailableRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailableRegion{}

// AvailableRegion List of regions that this cloud provider supports for this instance size.
type AvailableRegion struct {
	// Flag that indicates whether the cloud provider sets this region as its default. AWS defaults to US_EAST_1, GCP defaults to CENTRAL_US, and AZURE defaults to US_WEST_2.
	Default *bool `json:"default,omitempty"`
	// Human-readable label that identifies the supported region.
	Name *string `json:"name,omitempty"`
}

// NewAvailableRegion instantiates a new AvailableRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailableRegion() *AvailableRegion {
	this := AvailableRegion{}
	return &this
}

// NewAvailableRegionWithDefaults instantiates a new AvailableRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailableRegionWithDefaults() *AvailableRegion {
	this := AvailableRegion{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *AvailableRegion) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableRegion) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *AvailableRegion) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *AvailableRegion) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AvailableRegion) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailableRegion) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AvailableRegion) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AvailableRegion) SetName(v string) {
	o.Name = &v
}

func (o AvailableRegion) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AvailableRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAvailableRegion struct {
	value *AvailableRegion
	isSet bool
}

func (v NullableAvailableRegion) Get() *AvailableRegion {
	return v.value
}

func (v *NullableAvailableRegion) Set(val *AvailableRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailableRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailableRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailableRegion(val *AvailableRegion) *NullableAvailableRegion {
	return &NullableAvailableRegion{value: val, isSet: true}
}

func (v NullableAvailableRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailableRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


