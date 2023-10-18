// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServiceAccountAccessRole Service account access roles.
type ServiceAccountAccessRole struct {
	Groups []ServiceAccountProjectAccessRole `json:"groups,omitempty"`
	Org    []string                          `json:"org,omitempty"`
}

// NewServiceAccountAccessRole instantiates a new ServiceAccountAccessRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountAccessRole() *ServiceAccountAccessRole {
	this := ServiceAccountAccessRole{}
	return &this
}

// NewServiceAccountAccessRoleWithDefaults instantiates a new ServiceAccountAccessRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountAccessRoleWithDefaults() *ServiceAccountAccessRole {
	this := ServiceAccountAccessRole{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise
func (o *ServiceAccountAccessRole) GetGroups() []ServiceAccountProjectAccessRole {
	if o == nil || IsNil(o.Groups) {
		var ret []ServiceAccountProjectAccessRole
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessRole) GetGroupsOk() ([]ServiceAccountProjectAccessRole, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}

	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ServiceAccountAccessRole) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []ServiceAccountProjectAccessRole and assigns it to the Groups field.
func (o *ServiceAccountAccessRole) SetGroups(v []ServiceAccountProjectAccessRole) {
	o.Groups = v
}

// GetOrg returns the Org field value if set, zero value otherwise
func (o *ServiceAccountAccessRole) GetOrg() []string {
	if o == nil || IsNil(o.Org) {
		var ret []string
		return ret
	}
	return o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountAccessRole) GetOrgOk() ([]string, bool) {
	if o == nil || IsNil(o.Org) {
		return nil, false
	}

	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *ServiceAccountAccessRole) HasOrg() bool {
	if o != nil && !IsNil(o.Org) {
		return true
	}

	return false
}

// SetOrg gets a reference to the given []string and assigns it to the Org field.
func (o *ServiceAccountAccessRole) SetOrg(v []string) {
	o.Org = v
}

func (o ServiceAccountAccessRole) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountAccessRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Org) {
		toSerialize["org"] = o.Org
	}
	return toSerialize, nil
}
