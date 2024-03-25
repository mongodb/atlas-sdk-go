// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
	"time"
)

// OrgServiceAccountRequest struct for OrgServiceAccountRequest
type OrgServiceAccountRequest struct {
	// Human readable description for service account.
	Description string `json:"description"`
	// Human readable name for service account.
	Name string `json:"name"`
	// Project roles associated with Service account.
	RoleNames *[]string `json:"roleNames,omitempty"`
	// Secret expiration time.
	SecretExpiresAt time.Time `json:"secretExpiresAt"`
}

// NewOrgServiceAccountRequest instantiates a new OrgServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgServiceAccountRequest(description string, name string, secretExpiresAt time.Time) *OrgServiceAccountRequest {
	this := OrgServiceAccountRequest{}
	this.Description = description
	this.Name = name
	this.SecretExpiresAt = secretExpiresAt
	return &this
}

// NewOrgServiceAccountRequestWithDefaults instantiates a new OrgServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgServiceAccountRequestWithDefaults() *OrgServiceAccountRequest {
	this := OrgServiceAccountRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *OrgServiceAccountRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *OrgServiceAccountRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *OrgServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OrgServiceAccountRequest) SetName(v string) {
	o.Name = v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise
func (o *OrgServiceAccountRequest) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return *o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetRoleNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}

	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *OrgServiceAccountRequest) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *OrgServiceAccountRequest) SetRoleNames(v []string) {
	o.RoleNames = &v
}

// GetSecretExpiresAt returns the SecretExpiresAt field value
func (o *OrgServiceAccountRequest) GetSecretExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SecretExpiresAt
}

// GetSecretExpiresAtOk returns a tuple with the SecretExpiresAt field value
// and a boolean to check if the value has been set.
func (o *OrgServiceAccountRequest) GetSecretExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretExpiresAt, true
}

// SetSecretExpiresAt sets field value
func (o *OrgServiceAccountRequest) SetSecretExpiresAt(v time.Time) {
	o.SecretExpiresAt = v
}

func (o OrgServiceAccountRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o OrgServiceAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	if !IsNil(o.RoleNames) {
		toSerialize["roleNames"] = o.RoleNames
	}
	toSerialize["secretExpiresAt"] = o.SecretExpiresAt
	return toSerialize, nil
}
