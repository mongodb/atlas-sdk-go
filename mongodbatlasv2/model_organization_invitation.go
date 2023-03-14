/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package latest

import (
	"encoding/json"
	"time"
)

// checks if the OrganizationInvitation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationInvitation{}

// OrganizationInvitation struct for OrganizationInvitation
type OrganizationInvitation struct {
	// Date and time when MongoDB Cloud sent the invitation. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Date and time when the invitation from MongoDB Cloud expires. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Unique 24-hexadecimal digit string that identifies this organization.
	Id *string `json:"id,omitempty"`
	// Email address of the MongoDB Cloud user who sent the invitation to join the organization.
	InviterUsername *string `json:"inviterUsername,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization.
	OrgId *string `json:"orgId,omitempty"`
	// Human-readable label that identifies this organization.
	OrgName string `json:"orgName"`
	// One or more organization or project level roles to assign to the MongoDB Cloud user.
	Roles []string `json:"roles,omitempty"`
	// List of unique 24-hexadecimal digit strings that identifies each team.
	TeamIds []string `json:"teamIds,omitempty"`
	// Email address of the MongoDB Cloud user invited to join the organization.
	Username *string `json:"username,omitempty"`
}

// NewOrganizationInvitation instantiates a new OrganizationInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitation(orgName string) *OrganizationInvitation {
	this := OrganizationInvitation{}
	this.OrgName = orgName
	return &this
}

// NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationWithDefaults() *OrganizationInvitation {
	this := OrganizationInvitation{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrganizationInvitation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OrganizationInvitation) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationInvitation) SetId(v string) {
	o.Id = &v
}

// GetInviterUsername returns the InviterUsername field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetInviterUsername() string {
	if o == nil || IsNil(o.InviterUsername) {
		var ret string
		return ret
	}
	return *o.InviterUsername
}

// GetInviterUsernameOk returns a tuple with the InviterUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetInviterUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.InviterUsername) {
		return nil, false
	}
	return o.InviterUsername, true
}

// HasInviterUsername returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasInviterUsername() bool {
	if o != nil && !IsNil(o.InviterUsername) {
		return true
	}

	return false
}

// SetInviterUsername gets a reference to the given string and assigns it to the InviterUsername field.
func (o *OrganizationInvitation) SetInviterUsername(v string) {
	o.InviterUsername = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *OrganizationInvitation) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrganizationInvitation) SetOrgId(v string) {
	o.OrgId = &v
}

// GetOrgName returns the OrgName field value
func (o *OrganizationInvitation) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *OrganizationInvitation) SetOrgName(v string) {
	o.OrgName = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetRoles() []string {
	if o == nil || IsNil(o.Roles) {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *OrganizationInvitation) SetRoles(v []string) {
	o.Roles = v
}

// GetTeamIds returns the TeamIds field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetTeamIds() []string {
	if o == nil || IsNil(o.TeamIds) {
		var ret []string
		return ret
	}
	return o.TeamIds
}

// GetTeamIdsOk returns a tuple with the TeamIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetTeamIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.TeamIds) {
		return nil, false
	}
	return o.TeamIds, true
}

// HasTeamIds returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasTeamIds() bool {
	if o != nil && !IsNil(o.TeamIds) {
		return true
	}

	return false
}

// SetTeamIds gets a reference to the given []string and assigns it to the TeamIds field.
func (o *OrganizationInvitation) SetTeamIds(v []string) {
	o.TeamIds = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *OrganizationInvitation) SetUsername(v string) {
	o.Username = &v
}

func (o OrganizationInvitation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationInvitation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: createdAt is readOnly
	// skip: expiresAt is readOnly
	// skip: id is readOnly
	// skip: inviterUsername is readOnly
	// skip: links is readOnly
	// skip: orgId is readOnly
	toSerialize["orgName"] = o.OrgName
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	// skip: teamIds is readOnly
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableOrganizationInvitation struct {
	value *OrganizationInvitation
	isSet bool
}

func (v NullableOrganizationInvitation) Get() *OrganizationInvitation {
	return v.value
}

func (v *NullableOrganizationInvitation) Set(val *OrganizationInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitation(val *OrganizationInvitation) *NullableOrganizationInvitation {
	return &NullableOrganizationInvitation{value: val, isSet: true}
}

func (v NullableOrganizationInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


