// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// ServiceAccountDetails struct for ServiceAccountDetails
type ServiceAccountDetails struct {
	// Service account access list.
	AccessList []ServiceAccountAccessList `json:"accessList,omitempty"`
	// Timestamp representing creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description of the service account.
	Description *string `json:"description,omitempty"`
	// Service account client id.
	Id *string `json:"id,omitempty"`
	// Service account name.
	Name    *string                       `json:"name,omitempty"`
	Roles   *ServiceAccountAccessRole     `json:"roles,omitempty"`
	Secrets []ServiceAccountSecretDetails `json:"secrets,omitempty"`
}

// NewServiceAccountDetails instantiates a new ServiceAccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountDetails() *ServiceAccountDetails {
	this := ServiceAccountDetails{}
	return &this
}

// NewServiceAccountDetailsWithDefaults instantiates a new ServiceAccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountDetailsWithDefaults() *ServiceAccountDetails {
	this := ServiceAccountDetails{}
	return &this
}

// GetAccessList returns the AccessList field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetAccessList() []ServiceAccountAccessList {
	if o == nil || IsNil(o.AccessList) {
		var ret []ServiceAccountAccessList
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetAccessListOk() ([]ServiceAccountAccessList, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}

	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []ServiceAccountAccessList and assigns it to the AccessList field.
func (o *ServiceAccountDetails) SetAccessList(v []ServiceAccountAccessList) {
	o.AccessList = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ServiceAccountDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccountDetails) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccountDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceAccountDetails) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetRoles() ServiceAccountAccessRole {
	if o == nil || IsNil(o.Roles) {
		var ret ServiceAccountAccessRole
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetRolesOk() (*ServiceAccountAccessRole, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given ServiceAccountAccessRole and assigns it to the Roles field.
func (o *ServiceAccountDetails) SetRoles(v ServiceAccountAccessRole) {
	o.Roles = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise
func (o *ServiceAccountDetails) GetSecrets() []ServiceAccountSecretDetails {
	if o == nil || IsNil(o.Secrets) {
		var ret []ServiceAccountSecretDetails
		return ret
	}
	return o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetails) GetSecretsOk() ([]ServiceAccountSecretDetails, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}

	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *ServiceAccountDetails) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ServiceAccountSecretDetails and assigns it to the Secrets field.
func (o *ServiceAccountDetails) SetSecrets(v []ServiceAccountSecretDetails) {
	o.Secrets = v
}

func (o ServiceAccountDetails) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ServiceAccountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessList) {
		toSerialize["accessList"] = o.AccessList
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
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
