// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the Key type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Key{}

// Key Details contained in one API key.
type Key struct {
	// List of network addresses granted access to this API using this API key.
	AccessList []AccessListItem `json:"accessList,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization API key.
	Id string `json:"id"`
	// Public API key value set for the specified organization API key.
	PublicKey string `json:"publicKey"`
	// List that contains roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the Cloud user.
	Roles []CloudRoleAssignment `json:"roles,omitempty"`
}

// NewKey instantiates a new Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKey(id string, publicKey string) *Key {
	this := Key{}
	this.Id = id
	this.PublicKey = publicKey
	return &this
}

// NewKeyWithDefaults instantiates a new Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyWithDefaults() *Key {
	this := Key{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *Key) GetAccessList() []AccessListItem {
	if o == nil || IsNil(o.AccessList) {
		var ret []AccessListItem
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetAccessListOk() ([]AccessListItem, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *Key) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []AccessListItem and assigns it to the AccessList field.
func (o *Key) SetAccessList(v []AccessListItem) {
	o.AccessList = v
}

// GetId returns the Id field value
func (o *Key) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Key) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Key) SetId(v string) {
	o.Id = v
}

// GetPublicKey returns the PublicKey field value
func (o *Key) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *Key) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *Key) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *Key) GetRoles() []CloudRoleAssignment {
	if o == nil || IsNil(o.Roles) {
		var ret []CloudRoleAssignment
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetRolesOk() ([]CloudRoleAssignment, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *Key) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CloudRoleAssignment and assigns it to the Roles field.
func (o *Key) SetRoles(v []CloudRoleAssignment) {
	o.Roles = v
}

func (o Key) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Key) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableKey struct {
	value *Key
	isSet bool
}

func (v NullableKey) Get() *Key {
	return v.value
}

func (v *NullableKey) Set(val *Key) {
	v.value = val
	v.isSet = true
}

func (v NullableKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKey(val *Key) *NullableKey {
	return &NullableKey{value: val, isSet: true}
}

func (v NullableKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
