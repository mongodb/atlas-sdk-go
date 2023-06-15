// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the AccessListItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessListItem{}

// AccessListItem struct for AccessListItem
type AccessListItem struct {
	// Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that found in this project's access list.
	CidrBlock *string `json:"cidrBlock,omitempty"`
	// IP address included in the API access list.
	IpAddress string `json:"ipAddress"`
}

// NewAccessListItem instantiates a new AccessListItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessListItem(ipAddress string) *AccessListItem {
	this := AccessListItem{}
	this.IpAddress = ipAddress
	return &this
}

// NewAccessListItemWithDefaults instantiates a new AccessListItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessListItemWithDefaults() *AccessListItem {
	this := AccessListItem{}
	return &this
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise.
func (o *AccessListItem) GetCidrBlock() string {
	if o == nil || IsNil(o.CidrBlock) {
		var ret string
		return ret
	}
	return *o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessListItem) GetCidrBlockOk() (*string, bool) {
	if o == nil || IsNil(o.CidrBlock) {
		return nil, false
	}
	return o.CidrBlock, true
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *AccessListItem) HasCidrBlock() bool {
	if o != nil && !IsNil(o.CidrBlock) {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given string and assigns it to the CidrBlock field.
func (o *AccessListItem) SetCidrBlock(v string) {
	o.CidrBlock = &v
}

// GetIpAddress returns the IpAddress field value
func (o *AccessListItem) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *AccessListItem) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *AccessListItem) SetIpAddress(v string) {
	o.IpAddress = v
}

func (o AccessListItem) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o AccessListItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableAccessListItem struct {
	value *AccessListItem
	isSet bool
}

func (v NullableAccessListItem) Get() *AccessListItem {
	return v.value
}

func (v *NullableAccessListItem) Set(val *AccessListItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessListItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessListItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessListItem(val *AccessListItem) *NullableAccessListItem {
	return &NullableAccessListItem{value: val, isSet: true}
}

func (v NullableAccessListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessListItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
