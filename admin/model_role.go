// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Role{}

// Role struct for Role
type Role struct {
	// Collection on which this role applies.
	CollectionName *string `json:"collectionName,omitempty"`
	// Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB.
	DatabaseName string `json:"databaseName"`
	// Human-readable label that identifies a group of privileges assigned to a database user. This value can either be a built-in role or a custom role.
	RoleName string `json:"roleName"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole(databaseName string, roleName string) *Role {
	this := Role{}
	this.DatabaseName = databaseName
	this.RoleName = roleName
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetCollectionName returns the CollectionName field value if set, zero value otherwise.
func (o *Role) GetCollectionName() string {
	if o == nil || IsNil(o.CollectionName) {
		var ret string
		return ret
	}
	return *o.CollectionName
}

// GetCollectionNameOk returns a tuple with the CollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionName) {
		return nil, false
	}
	return o.CollectionName, true
}

// HasCollectionName returns a boolean if a field has been set.
func (o *Role) HasCollectionName() bool {
	if o != nil && !IsNil(o.CollectionName) {
		return true
	}

	return false
}

// SetCollectionName gets a reference to the given string and assigns it to the CollectionName field.
func (o *Role) SetCollectionName(v string) {
	o.CollectionName = &v
}

// GetDatabaseName returns the DatabaseName field value
func (o *Role) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *Role) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *Role) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetRoleName returns the RoleName field value
func (o *Role) GetRoleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value
// and a boolean to check if the value has been set.
func (o *Role) GetRoleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleName, true
}

// SetRoleName sets field value
func (o *Role) SetRoleName(v string) {
	o.RoleName = v
}

func (o Role) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectionName) {
		toSerialize["collectionName"] = o.CollectionName
	}
	toSerialize["databaseName"] = o.DatabaseName
	toSerialize["roleName"] = o.RoleName
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
