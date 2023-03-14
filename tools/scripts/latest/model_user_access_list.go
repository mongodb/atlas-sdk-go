/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
	"time"
)

// checks if the UserAccessList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAccessList{}

// UserAccessList struct for UserAccessList
type UserAccessList struct {
	// Range of network addresses that you want to add to the access list for the API key. This parameter requires the range to be expressed in classless inter-domain routing (CIDR) notation of Internet Protocol version 4 or version 6 addresses. You can set a value for this parameter or **ipAddress** but not both in the same request.
	CidrBlock *string `json:"cidrBlock,omitempty"`
	// Total number of requests that have originated from the Internet Protocol (IP) address given as the value of the *lastUsedAddress* parameter.
	Count *int32 `json:"count,omitempty"`
	// Date and time when someone added the network addresses to the specified API access list. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	Created *time.Time `json:"created,omitempty"`
	// Network address that you want to add to the access list for the API key. This parameter requires the address to be expressed as one Internet Protocol version 4 or version 6 address. You can set a value for this parameter or **cidrBlock** but not both in the same request.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Date and time when MongoDB Cloud received the most recent request that originated from this Internet Protocol version 4 or version 6 address. The resource returns this parameter when at least one request has originated from this IP address. MongoDB Cloud updates this parameter each time a client accesses the permitted resource. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	LastUsed *time.Time `json:"lastUsed,omitempty"`
	// Network address that issued the most recent request to the API. This parameter requires the address to be expressed as one Internet Protocol version 4 or version 6 address. The resource returns this parameter after this IP address made at least one request.
	LastUsedAddress *string `json:"lastUsedAddress,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
}

// NewUserAccessList instantiates a new UserAccessList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAccessList() *UserAccessList {
	this := UserAccessList{}
	return &this
}

// NewUserAccessListWithDefaults instantiates a new UserAccessList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAccessListWithDefaults() *UserAccessList {
	this := UserAccessList{}
	return &this
}

// GetCidrBlock returns the CidrBlock field value if set, zero value otherwise.
func (o *UserAccessList) GetCidrBlock() string {
	if o == nil || IsNil(o.CidrBlock) {
		var ret string
		return ret
	}
	return *o.CidrBlock
}

// GetCidrBlockOk returns a tuple with the CidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetCidrBlockOk() (*string, bool) {
	if o == nil || IsNil(o.CidrBlock) {
		return nil, false
	}
	return o.CidrBlock, true
}

// HasCidrBlock returns a boolean if a field has been set.
func (o *UserAccessList) HasCidrBlock() bool {
	if o != nil && !IsNil(o.CidrBlock) {
		return true
	}

	return false
}

// SetCidrBlock gets a reference to the given string and assigns it to the CidrBlock field.
func (o *UserAccessList) SetCidrBlock(v string) {
	o.CidrBlock = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UserAccessList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UserAccessList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *UserAccessList) SetCount(v int32) {
	o.Count = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserAccessList) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserAccessList) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserAccessList) SetCreated(v time.Time) {
	o.Created = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *UserAccessList) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *UserAccessList) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *UserAccessList) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *UserAccessList) GetLastUsed() time.Time {
	if o == nil || IsNil(o.LastUsed) {
		var ret time.Time
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetLastUsedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUsed) {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *UserAccessList) HasLastUsed() bool {
	if o != nil && !IsNil(o.LastUsed) {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given time.Time and assigns it to the LastUsed field.
func (o *UserAccessList) SetLastUsed(v time.Time) {
	o.LastUsed = &v
}

// GetLastUsedAddress returns the LastUsedAddress field value if set, zero value otherwise.
func (o *UserAccessList) GetLastUsedAddress() string {
	if o == nil || IsNil(o.LastUsedAddress) {
		var ret string
		return ret
	}
	return *o.LastUsedAddress
}

// GetLastUsedAddressOk returns a tuple with the LastUsedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetLastUsedAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsedAddress) {
		return nil, false
	}
	return o.LastUsedAddress, true
}

// HasLastUsedAddress returns a boolean if a field has been set.
func (o *UserAccessList) HasLastUsedAddress() bool {
	if o != nil && !IsNil(o.LastUsedAddress) {
		return true
	}

	return false
}

// SetLastUsedAddress gets a reference to the given string and assigns it to the LastUsedAddress field.
func (o *UserAccessList) SetLastUsedAddress(v string) {
	o.LastUsedAddress = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserAccessList) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAccessList) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserAccessList) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *UserAccessList) SetLinks(v []Link) {
	o.Links = v
}

func (o UserAccessList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAccessList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CidrBlock) {
		toSerialize["cidrBlock"] = o.CidrBlock
	}
	// skip: count is readOnly
	// skip: created is readOnly
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	// skip: lastUsed is readOnly
	// skip: lastUsedAddress is readOnly
	// skip: links is readOnly
	return toSerialize, nil
}

type NullableUserAccessList struct {
	value *UserAccessList
	isSet bool
}

func (v NullableUserAccessList) Get() *UserAccessList {
	return v.value
}

func (v *NullableUserAccessList) Set(val *UserAccessList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAccessList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAccessList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAccessList(val *UserAccessList) *NullableUserAccessList {
	return &NullableUserAccessList{value: val, isSet: true}
}

func (v NullableUserAccessList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAccessList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


