// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServiceAccountUpdateRequest struct for ServiceAccountUpdateRequest
type ServiceAccountUpdateRequest struct {
	// Service account access list.
	AccessList []ServiceAccountAccessList `json:"accessList,omitempty"`
	// Description of the service account client.
	Description *string `json:"description,omitempty"`
	// Client name for service account.
	Name *string `json:"name,omitempty"`
	// Service account access roles.
	Roles []ServiceAccountAccessRole `json:"roles,omitempty"`
}

// NewServiceAccountUpdateRequest instantiates a new ServiceAccountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountUpdateRequest() *ServiceAccountUpdateRequest {
	this := ServiceAccountUpdateRequest{}
	return &this
}

// NewServiceAccountUpdateRequestWithDefaults instantiates a new ServiceAccountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountUpdateRequestWithDefaults() *ServiceAccountUpdateRequest {
	this := ServiceAccountUpdateRequest{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise
func (o *ServiceAccountUpdateRequest) GetAccessList() []ServiceAccountAccessList {
	if o == nil || IsNil(o.AccessList) {
		var ret []ServiceAccountAccessList
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetAccessListOk() ([]ServiceAccountAccessList, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}

	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *ServiceAccountUpdateRequest) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []ServiceAccountAccessList and assigns it to the AccessList field.
func (o *ServiceAccountUpdateRequest) SetAccessList(v []ServiceAccountAccessList) {
	o.AccessList = v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *ServiceAccountUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *ServiceAccountUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *ServiceAccountUpdateRequest) GetRoles() []ServiceAccountAccessRole {
	if o == nil || IsNil(o.Roles) {
		var ret []ServiceAccountAccessRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountUpdateRequest) GetRolesOk() ([]ServiceAccountAccessRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ServiceAccountUpdateRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ServiceAccountAccessRole and assigns it to the Roles field.
func (o *ServiceAccountUpdateRequest) SetRoles(v []ServiceAccountAccessRole) {
	o.Roles = v
}

func (o ServiceAccountUpdateRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessList) {
		toSerialize["accessList"] = o.AccessList
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}
