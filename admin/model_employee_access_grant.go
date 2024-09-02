// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// EmployeeAccessGrant MongoDB employee granted access level and expiration for a cluster.
type EmployeeAccessGrant struct {
	// Expiration date for the employee access grant.
	Expiration time.Time `json:"expiration"`
	// Level of access to grant to MongoDB Employees.
	GrantType string `json:"grantType"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewEmployeeAccessGrant instantiates a new EmployeeAccessGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeAccessGrant(expiration time.Time, grantType string) *EmployeeAccessGrant {
	this := EmployeeAccessGrant{}
	this.Expiration = expiration
	this.GrantType = grantType
	return &this
}

// NewEmployeeAccessGrantWithDefaults instantiates a new EmployeeAccessGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeAccessGrantWithDefaults() *EmployeeAccessGrant {
	this := EmployeeAccessGrant{}
	return &this
}

// GetExpiration returns the Expiration field value
func (o *EmployeeAccessGrant) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *EmployeeAccessGrant) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
func (o *EmployeeAccessGrant) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetGrantType returns the GrantType field value
func (o *EmployeeAccessGrant) GetGrantType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value
// and a boolean to check if the value has been set.
func (o *EmployeeAccessGrant) GetGrantTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantType, true
}

// SetGrantType sets field value
func (o *EmployeeAccessGrant) SetGrantType(v string) {
	o.GrantType = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *EmployeeAccessGrant) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeAccessGrant) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EmployeeAccessGrant) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *EmployeeAccessGrant) SetLinks(v []Link) {
	o.Links = &v
}

func (o EmployeeAccessGrant) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o EmployeeAccessGrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expiration"] = o.Expiration
	toSerialize["grantType"] = o.GrantType
	return toSerialize, nil
}
