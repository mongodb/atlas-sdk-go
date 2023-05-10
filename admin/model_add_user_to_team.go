// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

import (
	"encoding/json"
)

// checks if the AddUserToTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserToTeam{}

// AddUserToTeam struct for AddUserToTeam
type AddUserToTeam struct {
	// Unique 24-hexadecimal digit string that identifies the MongoDB Cloud user.
	Id string `json:"id"`
}

// NewAddUserToTeam instantiates a new AddUserToTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserToTeam(id string) *AddUserToTeam {
	this := AddUserToTeam{}
	this.Id = id
	return &this
}

// NewAddUserToTeamWithDefaults instantiates a new AddUserToTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserToTeamWithDefaults() *AddUserToTeam {
	this := AddUserToTeam{}
	return &this
}

// GetId returns the Id field value
func (o *AddUserToTeam) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AddUserToTeam) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AddUserToTeam) SetId(v string) {
	o.Id = v
}

func (o AddUserToTeam) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AddUserToTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAddUserToTeam struct {
	value *AddUserToTeam
	isSet bool
}

func (v NullableAddUserToTeam) Get() *AddUserToTeam {
	return v.value
}

func (v *NullableAddUserToTeam) Set(val *AddUserToTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserToTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserToTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserToTeam(val *AddUserToTeam) *NullableAddUserToTeam {
	return &NullableAddUserToTeam{value: val, isSet: true}
}

func (v NullableAddUserToTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserToTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
