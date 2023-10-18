// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// ServiceAccountRequest struct for ServiceAccountRequest
type ServiceAccountRequest struct {
	AccessList []ServiceAccountAccessList `json:"accessList,omitempty"`
	// Human readable description for service account.
	Description string `json:"description"`
	// Human readable name for service account.
	Name  string                   `json:"name"`
	Roles ServiceAccountAccessRole `json:"roles"`
	// Secret expiration time.
	SecretExpiresAt time.Time `json:"secretExpiresAt"`
}

// NewServiceAccountRequest instantiates a new ServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountRequest(description string, name string, roles ServiceAccountAccessRole, secretExpiresAt time.Time) *ServiceAccountRequest {
	this := ServiceAccountRequest{}
	this.Description = description
	this.Name = name
	this.Roles = roles
	this.SecretExpiresAt = secretExpiresAt
	return &this
}

// NewServiceAccountRequestWithDefaults instantiates a new ServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountRequestWithDefaults() *ServiceAccountRequest {
	this := ServiceAccountRequest{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise
func (o *ServiceAccountRequest) GetAccessList() []ServiceAccountAccessList {
	if o == nil || IsNil(o.AccessList) {
		var ret []ServiceAccountAccessList
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetAccessListOk() ([]ServiceAccountAccessList, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}

	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *ServiceAccountRequest) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []ServiceAccountAccessList and assigns it to the AccessList field.
func (o *ServiceAccountRequest) SetAccessList(v []ServiceAccountAccessList) {
	o.AccessList = v
}

// GetDescription returns the Description field value
func (o *ServiceAccountRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ServiceAccountRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *ServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceAccountRequest) SetName(v string) {
	o.Name = v
}

// GetRoles returns the Roles field value
func (o *ServiceAccountRequest) GetRoles() ServiceAccountAccessRole {
	if o == nil {
		var ret ServiceAccountAccessRole
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetRolesOk() (*ServiceAccountAccessRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Roles, true
}

// SetRoles sets field value
func (o *ServiceAccountRequest) SetRoles(v ServiceAccountAccessRole) {
	o.Roles = v
}

// GetSecretExpiresAt returns the SecretExpiresAt field value
func (o *ServiceAccountRequest) GetSecretExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SecretExpiresAt
}

// GetSecretExpiresAtOk returns a tuple with the SecretExpiresAt field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountRequest) GetSecretExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretExpiresAt, true
}

// SetSecretExpiresAt sets field value
func (o *ServiceAccountRequest) SetSecretExpiresAt(v time.Time) {
	o.SecretExpiresAt = v
}

func (o ServiceAccountRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessList) {
		toSerialize["accessList"] = o.AccessList
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["roles"] = o.Roles
	toSerialize["secretExpiresAt"] = o.SecretExpiresAt
	return toSerialize, nil
}
