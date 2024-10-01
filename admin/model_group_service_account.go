// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// GroupServiceAccount struct for GroupServiceAccount
type GroupServiceAccount struct {
	// The Client ID of the Service Account.
	ClientId *string `json:"clientId,omitempty"`
	// The date that the Service Account was created on.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Human readable description for the Service Account.
	Description *string `json:"description,omitempty"`
	// Human-readable name for the Service Account.
	Name *string `json:"name,omitempty"`
	// A list of Project roles associated with the Service Account.
	Roles *[]string `json:"roles,omitempty"`
	// A list of secrets associated with the specified Service Account.
	Secrets *[]ServiceAccountSecret `json:"secrets,omitempty"`
}

// NewGroupServiceAccount instantiates a new GroupServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupServiceAccount() *GroupServiceAccount {
	this := GroupServiceAccount{}
	return &this
}

// NewGroupServiceAccountWithDefaults instantiates a new GroupServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceAccountWithDefaults() *GroupServiceAccount {
	this := GroupServiceAccount{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise
func (o *GroupServiceAccount) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccount) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}

	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GroupServiceAccount) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *GroupServiceAccount) SetClientId(v string) {
	o.ClientId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *GroupServiceAccount) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupServiceAccount) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupServiceAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *GroupServiceAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupServiceAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *GroupServiceAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupServiceAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupServiceAccount) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *GroupServiceAccount) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccount) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupServiceAccount) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupServiceAccount) SetRoles(v []string) {
	o.Roles = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise
func (o *GroupServiceAccount) GetSecrets() []ServiceAccountSecret {
	if o == nil || IsNil(o.Secrets) {
		var ret []ServiceAccountSecret
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccount) GetSecretsOk() (*[]ServiceAccountSecret, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}

	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *GroupServiceAccount) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ServiceAccountSecret and assigns it to the Secrets field.
func (o *GroupServiceAccount) SetSecrets(v []ServiceAccountSecret) {
	o.Secrets = &v
}

func (o *GroupServiceAccount) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o *GroupServiceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
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
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}
