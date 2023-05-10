// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the InheritedRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InheritedRole{}

// InheritedRole Role inherited from another context for this database user.
type InheritedRole struct {
	// Human-readable label that identifies the database on which someone grants the action to one MongoDB user.
	Db string `json:"db"`
	// Human-readable label that identifies the role inherited. Set this value to `admin` for every role except `read` or `readWrite`.
	Role string `json:"role"`
}

// NewInheritedRole instantiates a new InheritedRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInheritedRole(db string, role string) *InheritedRole {
	this := InheritedRole{}
	this.Db = db
	this.Role = role
	return &this
}

// NewInheritedRoleWithDefaults instantiates a new InheritedRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInheritedRoleWithDefaults() *InheritedRole {
	this := InheritedRole{}
	return &this
}

// GetDb returns the Db field value
func (o *InheritedRole) GetDb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Db
}

// GetDbOk returns a tuple with the Db field value
// and a boolean to check if the value has been set.
func (o *InheritedRole) GetDbOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Db, true
}

// SetDb sets field value
func (o *InheritedRole) SetDb(v string) {
	o.Db = v
}

// GetRole returns the Role field value
func (o *InheritedRole) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *InheritedRole) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *InheritedRole) SetRole(v string) {
	o.Role = v
}

func (o InheritedRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o InheritedRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["db"] = o.Db
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

type NullableInheritedRole struct {
	value *InheritedRole
	isSet bool
}

func (v NullableInheritedRole) Get() *InheritedRole {
	return v.value
}

func (v *NullableInheritedRole) Set(val *InheritedRole) {
	v.value = val
	v.isSet = true
}

func (v NullableInheritedRole) IsSet() bool {
	return v.isSet
}

func (v *NullableInheritedRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInheritedRole(val *InheritedRole) *NullableInheritedRole {
	return &NullableInheritedRole{value: val, isSet: true}
}

func (v NullableInheritedRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInheritedRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
