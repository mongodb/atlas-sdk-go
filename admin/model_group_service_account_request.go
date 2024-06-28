// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupServiceAccountRequest struct for GroupServiceAccountRequest
type GroupServiceAccountRequest struct {
	// Human readable description for the service account.
	Description string `json:"description"`
	// Human-readable name for the service account. The name is modifiable and does not have to be unique.
	Name string `json:"name"`
	// Project roles associated with the service account.
	Roles *[]string `json:"roles,omitempty"`
	// Secret expiration time.
	SecretExpiresAfterHours int `json:"secretExpiresAfterHours"`
}

// NewGroupServiceAccountRequest instantiates a new GroupServiceAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupServiceAccountRequest(description string, name string, secretExpiresAfterHours int) *GroupServiceAccountRequest {
	this := GroupServiceAccountRequest{}
	this.Description = description
	this.Name = name
	this.SecretExpiresAfterHours = secretExpiresAfterHours
	return &this
}

// NewGroupServiceAccountRequestWithDefaults instantiates a new GroupServiceAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceAccountRequestWithDefaults() *GroupServiceAccountRequest {
	this := GroupServiceAccountRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *GroupServiceAccountRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *GroupServiceAccountRequest) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *GroupServiceAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupServiceAccountRequest) SetName(v string) {
	o.Name = v
}

// GetRoles returns the Roles field value if set, zero value otherwise
func (o *GroupServiceAccountRequest) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetRolesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}

	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *GroupServiceAccountRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *GroupServiceAccountRequest) SetRoles(v []string) {
	o.Roles = &v
}

// GetSecretExpiresAfterHours returns the SecretExpiresAfterHours field value
func (o *GroupServiceAccountRequest) GetSecretExpiresAfterHours() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.SecretExpiresAfterHours
}

// GetSecretExpiresAfterHoursOk returns a tuple with the SecretExpiresAfterHours field value
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountRequest) GetSecretExpiresAfterHoursOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretExpiresAfterHours, true
}

// SetSecretExpiresAfterHours sets field value
func (o *GroupServiceAccountRequest) SetSecretExpiresAfterHours(v int) {
	o.SecretExpiresAfterHours = v
}

func (o GroupServiceAccountRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupServiceAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	toSerialize["secretExpiresAfterHours"] = o.SecretExpiresAfterHours
	return toSerialize, nil
}
