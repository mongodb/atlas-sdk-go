// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the KeyUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyUser{}

// KeyUser struct for KeyUser
type KeyUser struct {
	// Purpose or explanation provided when someone created this organization API key.
	Desc *string `json:"desc,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project.
	Id *string `json:"id,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Redacted private key returned for this organization API key. This key displays unredacted when first created.
	PrivateKey *string `json:"privateKey,omitempty"`
	// Public API key value set for the specified organization API key.
	PublicKey *string `json:"publicKey,omitempty"`
	// List that contains the roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the API key.
	Roles []CloudRoleAssignment `json:"roles,omitempty"`
}

// NewKeyUser instantiates a new KeyUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyUser() *KeyUser {
	this := KeyUser{}
	return &this
}

// NewKeyUserWithDefaults instantiates a new KeyUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyUserWithDefaults() *KeyUser {
	this := KeyUser{}
	return &this
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *KeyUser) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUser) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *KeyUser) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *KeyUser) SetDesc(v string) {
	o.Desc = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KeyUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KeyUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KeyUser) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *KeyUser) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUser) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *KeyUser) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *KeyUser) SetLinks(v []Link) {
	o.Links = v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *KeyUser) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUser) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *KeyUser) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *KeyUser) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *KeyUser) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUser) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *KeyUser) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *KeyUser) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *KeyUser) GetRoles() []CloudRoleAssignment {
	if o == nil || IsNil(o.Roles) {
		var ret []CloudRoleAssignment
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUser) GetRolesOk() ([]CloudRoleAssignment, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *KeyUser) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []CloudRoleAssignment and assigns it to the Roles field.
func (o *KeyUser) SetRoles(v []CloudRoleAssignment) {
	o.Roles = v
}

func (o KeyUser) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o KeyUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableKeyUser struct {
	value *KeyUser
	isSet bool
}

func (v NullableKeyUser) Get() *KeyUser {
	return v.value
}

func (v *NullableKeyUser) Set(val *KeyUser) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyUser) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyUser(val *KeyUser) *NullableKeyUser {
	return &NullableKeyUser{value: val, isSet: true}
}

func (v NullableKeyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
