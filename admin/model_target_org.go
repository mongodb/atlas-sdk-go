// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the TargetOrg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TargetOrg{}

// TargetOrg struct for TargetOrg
type TargetOrg struct {
	// Link token that contains all the information required to complete the link.
	LinkToken string `json:"linkToken"`
}

// NewTargetOrg instantiates a new TargetOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetOrg(linkToken string) *TargetOrg {
	this := TargetOrg{}
	this.LinkToken = linkToken
	return &this
}

// NewTargetOrgWithDefaults instantiates a new TargetOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetOrgWithDefaults() *TargetOrg {
	this := TargetOrg{}
	return &this
}

// GetLinkToken returns the LinkToken field value
func (o *TargetOrg) GetLinkToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkToken
}

// GetLinkTokenOk returns a tuple with the LinkToken field value
// and a boolean to check if the value has been set.
func (o *TargetOrg) GetLinkTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkToken, true
}

// SetLinkToken sets field value
func (o *TargetOrg) SetLinkToken(v string) {
	o.LinkToken = v
}

func (o TargetOrg) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o TargetOrg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["linkToken"] = o.LinkToken
	return toSerialize, nil
}

type NullableTargetOrg struct {
	value *TargetOrg
	isSet bool
}

func (v NullableTargetOrg) Get() *TargetOrg {
	return v.value
}

func (v *NullableTargetOrg) Set(val *TargetOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetOrg(val *TargetOrg) *NullableTargetOrg {
	return &NullableTargetOrg{value: val, isSet: true}
}

func (v NullableTargetOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
