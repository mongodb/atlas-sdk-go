// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the UserScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserScope{}

// UserScope Range of resources available to this database user.
type UserScope struct {
	// Human-readable label that identifies the cluster or MongoDB Atlas Data Lake that this database user can access.
	Name string `json:"name"`
	// Category of resource that this database user can access.
	Type string `json:"type"`
}

// NewUserScope instantiates a new UserScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserScope(name string, type_ string) *UserScope {
	this := UserScope{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewUserScopeWithDefaults instantiates a new UserScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserScopeWithDefaults() *UserScope {
	this := UserScope{}
	return &this
}

// GetName returns the Name field value
func (o *UserScope) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UserScope) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UserScope) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *UserScope) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserScope) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserScope) SetType(v string) {
	o.Type = v
}

func (o UserScope) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o UserScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableUserScope struct {
	value *UserScope
	isSet bool
}

func (v NullableUserScope) Get() *UserScope {
	return v.value
}

func (v *NullableUserScope) Set(val *UserScope) {
	v.value = val
	v.isSet = true
}

func (v NullableUserScope) IsSet() bool {
	return v.isSet
}

func (v *NullableUserScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserScope(val *UserScope) *NullableUserScope {
	return &NullableUserScope{value: val, isSet: true}
}

func (v NullableUserScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
