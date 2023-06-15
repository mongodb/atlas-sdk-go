// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the GroupName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupName{}

// GroupName Request view to update the group name.
type GroupName struct {
	// Human-readable label that identifies the project included in the MongoDB Cloud organization.
	Name *string `json:"name,omitempty"`
}

// NewGroupName instantiates a new GroupName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupName() *GroupName {
	this := GroupName{}
	return &this
}

// NewGroupNameWithDefaults instantiates a new GroupName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupNameWithDefaults() *GroupName {
	this := GroupName{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupName) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupName) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupName) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupName) SetName(v string) {
	o.Name = &v
}

func (o GroupName) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGroupName struct {
	value *GroupName
	isSet bool
}

func (v NullableGroupName) Get() *GroupName {
	return v.value
}

func (v *NullableGroupName) Set(val *GroupName) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupName) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupName(val *GroupName) *NullableGroupName {
	return &NullableGroupName{value: val, isSet: true}
}

func (v NullableGroupName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
