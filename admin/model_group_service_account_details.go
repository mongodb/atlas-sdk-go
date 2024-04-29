// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// GroupServiceAccountDetails struct for GroupServiceAccountDetails
type GroupServiceAccountDetails struct {
	// Timestamp representing creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description of the service account.
	Description *string `json:"description,omitempty"`
	// Service account client id.
	Id *string `json:"id,omitempty"`
	// Service account name.
	Name *string `json:"name,omitempty"`
	// Roles assigned to the Service Account group.
	Roles *[]string `json:"roles,omitempty"`
	// List of Client Secret Public details.
	Secrets *[]ServiceAccountSecretDetails `json:"secrets,omitempty"`
}

// NewGroupServiceAccountDetails instantiates a new GroupServiceAccountDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupServiceAccountDetails() *GroupServiceAccountDetails {
	this := GroupServiceAccountDetails{}
	return &this
}

// NewGroupServiceAccountDetailsWithDefaults instantiates a new GroupServiceAccountDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceAccountDetailsWithDefaults() *GroupServiceAccountDetails {
	this := GroupServiceAccountDetails{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *GroupServiceAccountDetails) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GroupServiceAccountDetails) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GroupServiceAccountDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *GroupServiceAccountDetails) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountDetails) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupServiceAccountDetails) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupServiceAccountDetails) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *GroupServiceAccountDetails) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountDetails) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupServiceAccountDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupServiceAccountDetails) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *GroupServiceAccountDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupServiceAccountDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupServiceAccountDetails) SetName(v string) {
	o.Name = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *GroupServiceAccountDetails) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountDetails) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupServiceAccountDetails) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupServiceAccountDetails) SetRoles(v []string) {
	o.Roles = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise
func (o *GroupServiceAccountDetails) GetSecrets() []ServiceAccountSecretDetails {
	if o == nil || IsNil(o.Secrets) {
		var ret []ServiceAccountSecretDetails
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountDetails) GetSecretsOk() (*[]ServiceAccountSecretDetails, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}

	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *GroupServiceAccountDetails) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ServiceAccountSecretDetails and assigns it to the Secrets field.
func (o *GroupServiceAccountDetails) SetSecrets(v []ServiceAccountSecretDetails) {
	o.Secrets = &v
}

func (o GroupServiceAccountDetails) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupServiceAccountDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
