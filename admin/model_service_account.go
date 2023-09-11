// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ServiceAccount struct for ServiceAccount
type ServiceAccount struct {
	AccessList []ServiceAccountAccessList `json:"accessList,omitempty"`
	// Description of the service account.
	Description *string `json:"description,omitempty"`
	// ID for the service account.
	Id *string `json:"id,omitempty"`
	// Name for service account.
	Name  *string                    `json:"name,omitempty"`
	Roles []ServiceAccountAccessRole `json:"roles,omitempty"`
	// List of secrets.
	Secrets []ServiceAccountSecret `json:"secrets,omitempty"`
}

// NewServiceAccount instantiates a new ServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccount() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// NewServiceAccountWithDefaults instantiates a new ServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountWithDefaults() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise
func (o *ServiceAccount) GetAccessList() []ServiceAccountAccessList {
	if o == nil || IsNil(o.AccessList) {
		var ret []ServiceAccountAccessList
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetAccessListOk() ([]ServiceAccountAccessList, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}

	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *ServiceAccount) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []ServiceAccountAccessList and assigns it to the AccessList field.
func (o *ServiceAccount) SetAccessList(v []ServiceAccountAccessList) {
	o.AccessList = v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *ServiceAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ServiceAccount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *ServiceAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccount) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *ServiceAccount) GetRoles() []ServiceAccountAccessRole {
	if o == nil || IsNil(o.Roles) {
		var ret []ServiceAccountAccessRole
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetRolesOk() ([]ServiceAccountAccessRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ServiceAccount) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []ServiceAccountAccessRole and assigns it to the Roles field.
func (o *ServiceAccount) SetRoles(v []ServiceAccountAccessRole) {
	o.Roles = v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise
func (o *ServiceAccount) GetSecrets() []ServiceAccountSecret {
	if o == nil || IsNil(o.Secrets) {
		var ret []ServiceAccountSecret
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetSecretsOk() ([]ServiceAccountSecret, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}

	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *ServiceAccount) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ServiceAccountSecret and assigns it to the Secrets field.
func (o *ServiceAccount) SetSecrets(v []ServiceAccountSecret) {
	o.Secrets = v
}

func (o ServiceAccount) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessList) {
		toSerialize["accessList"] = o.AccessList
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}
