// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// OrgServiceAccount struct for OrgServiceAccount
type OrgServiceAccount struct {
	// Service account creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description of the service account.
	Description *string `json:"description,omitempty"`
	// ID for the service account.
	Id *string `json:"id,omitempty"`
	// Name for service account.
	Name *string `json:"name,omitempty"`
	// Organization roles assigned to the Service Account.
	RoleNames *[]string `json:"roleNames,omitempty"`
	// List of secrets.
	Secrets *[]ServiceAccountSecret `json:"secrets,omitempty"`
}

// NewOrgServiceAccount instantiates a new OrgServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgServiceAccount() *OrgServiceAccount {
	this := OrgServiceAccount{}
	return &this
}

// NewOrgServiceAccountWithDefaults instantiates a new OrgServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgServiceAccountWithDefaults() *OrgServiceAccount {
	this := OrgServiceAccount{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *OrgServiceAccount) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccount) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrgServiceAccount) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrgServiceAccount) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *OrgServiceAccount) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OrgServiceAccount) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OrgServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *OrgServiceAccount) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccount) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgServiceAccount) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgServiceAccount) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *OrgServiceAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrgServiceAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrgServiceAccount) SetName(v string) {
	o.Name = &v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise
func (o *OrgServiceAccount) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return *o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccount) GetRoleNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}

	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *OrgServiceAccount) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *OrgServiceAccount) SetRoleNames(v []string) {
	o.RoleNames = &v
}

// GetSecrets returns the Secrets field value if set, zero value otherwise
func (o *OrgServiceAccount) GetSecrets() []ServiceAccountSecret {
	if o == nil || IsNil(o.Secrets) {
		var ret []ServiceAccountSecret
		return ret
	}
	return *o.Secrets
}

// GetSecretsOk returns a tuple with the Secrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccount) GetSecretsOk() (*[]ServiceAccountSecret, bool) {
	if o == nil || IsNil(o.Secrets) {
		return nil, false
	}

	return o.Secrets, true
}

// HasSecrets returns a boolean if a field has been set.
func (o *OrgServiceAccount) HasSecrets() bool {
	if o != nil && !IsNil(o.Secrets) {
		return true
	}

	return false
}

// SetSecrets gets a reference to the given []ServiceAccountSecret and assigns it to the Secrets field.
func (o *OrgServiceAccount) SetSecrets(v []ServiceAccountSecret) {
	o.Secrets = &v
}

func (o OrgServiceAccount) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OrgServiceAccount) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RoleNames) {
		toSerialize["roleNames"] = o.RoleNames
	}
	if !IsNil(o.Secrets) {
		toSerialize["secrets"] = o.Secrets
	}
	return toSerialize, nil
}
