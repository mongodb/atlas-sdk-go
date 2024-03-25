// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// GroupServiceAccountUpdateRequest struct for GroupServiceAccountUpdateRequest
type GroupServiceAccountUpdateRequest struct {
	// Description of the service account client.
	Description *string `json:"description,omitempty"`
	// Client name for service account.
	Name *string `json:"name,omitempty"`
	// Project roles associated with Service account.
	RoleNames *[]string `json:"roleNames,omitempty"`
}

// NewGroupServiceAccountUpdateRequest instantiates a new GroupServiceAccountUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupServiceAccountUpdateRequest() *GroupServiceAccountUpdateRequest {
	this := GroupServiceAccountUpdateRequest{}
	return &this
}

// NewGroupServiceAccountUpdateRequestWithDefaults instantiates a new GroupServiceAccountUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupServiceAccountUpdateRequestWithDefaults() *GroupServiceAccountUpdateRequest {
	this := GroupServiceAccountUpdateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *GroupServiceAccountUpdateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupServiceAccountUpdateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupServiceAccountUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *GroupServiceAccountUpdateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountUpdateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupServiceAccountUpdateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupServiceAccountUpdateRequest) SetName(v string) {
	o.Name = &v
}

// GetRoleNames returns the RoleNames field value if set, zero value otherwise
func (o *GroupServiceAccountUpdateRequest) GetRoleNames() []string {
	if o == nil || IsNil(o.RoleNames) {
		var ret []string
		return ret
	}
	return *o.RoleNames
}

// GetRoleNamesOk returns a tuple with the RoleNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupServiceAccountUpdateRequest) GetRoleNamesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RoleNames) {
		return nil, false
	}

	return o.RoleNames, true
}

// HasRoleNames returns a boolean if a field has been set.
func (o *GroupServiceAccountUpdateRequest) HasRoleNames() bool {
	if o != nil && !IsNil(o.RoleNames) {
		return true
	}

	return false
}

// SetRoleNames gets a reference to the given []string and assigns it to the RoleNames field.
func (o *GroupServiceAccountUpdateRequest) SetRoleNames(v []string) {
	o.RoleNames = &v
}

func (o GroupServiceAccountUpdateRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o GroupServiceAccountUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RoleNames) {
		toSerialize["roleNames"] = o.RoleNames
	}
	return toSerialize, nil
}
