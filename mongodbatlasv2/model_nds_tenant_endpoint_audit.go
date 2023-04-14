/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"encoding/json"
	"time"
)

// checks if the NDSTenantEndpointAudit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NDSTenantEndpointAudit{}

// NDSTenantEndpointAudit NDS tenant endpoint audit indicates when atlas auto-scaling cluster tier up or down.
type NDSTenantEndpointAudit struct {
	// Unique 24-hexadecimal digit string that identifies the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **userId** parameter.
	ApiKeyId *string `json:"apiKeyId,omitempty"`
	// Date and time when this event occurred. This parameter expresses its value in the <a href=\"https://en.wikipedia.org/wiki/ISO_8601\" target=\"_blank\" rel=\"noopener noreferrer\">ISO 8601</a> timestamp format in UTC.
	Created time.Time `json:"created"`
	// Unique 24-hexadecimal digit string that identifies the endpoint associated with this event.
	EndpointId *string `json:"endpointId,omitempty"`
	EventTypeName NDSTenantEndpointAuditType `json:"eventTypeName"`
	// Unique 24-hexadecimal digit string that identifies the project in which the event occurred. The **eventId** identifies the specific event.
	GroupId *string `json:"groupId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the event.
	Id string `json:"id"`
	// Flag that indicates whether a MongoDB employee triggered the specified event.
	IsGlobalAdmin *bool `json:"isGlobalAdmin,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	Links []Link `json:"links,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the organization to which these events apply.
	OrgId *string `json:"orgId,omitempty"`
	// Unique identification string that the cloud provider uses to identify the private endpoint.
	ProviderEndpointId *string `json:"providerEndpointId,omitempty"`
	// Public part of the [API Key](https://dochub.mongodb.org/core/atlas-create-prog-api-key) that triggered the event. If this resource returns this parameter, it doesn't return the **username** parameter.
	PublicKey *string `json:"publicKey,omitempty"`
	Raw *Raw `json:"raw,omitempty"`
	// IPv4 or IPv6 address from which the user triggered this event.
	RemoteAddress *string `json:"remoteAddress,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the console user who triggered the event. If this resource returns this parameter, it doesn't return the **apiKeyId** parameter.
	UserId *string `json:"userId,omitempty"`
	// Email address for the user who triggered this event. If this resource returns this parameter, it doesn't return the **publicApiKey** parameter.
	Username *string `json:"username,omitempty"`
}

// NewNDSTenantEndpointAudit instantiates a new NDSTenantEndpointAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNDSTenantEndpointAudit(created time.Time, eventTypeName NDSTenantEndpointAuditType, id string) *NDSTenantEndpointAudit {
	this := NDSTenantEndpointAudit{}
	this.Created = created
	this.EventTypeName = eventTypeName
	this.Id = id
	return &this
}

// NewNDSTenantEndpointAuditWithDefaults instantiates a new NDSTenantEndpointAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNDSTenantEndpointAuditWithDefaults() *NDSTenantEndpointAudit {
	this := NDSTenantEndpointAudit{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetApiKeyId() string {
	if o == nil || IsNil(o.ApiKeyId) {
		var ret string
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetApiKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKeyId) {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasApiKeyId() bool {
	if o != nil && !IsNil(o.ApiKeyId) {
		return true
	}

	return false
}

// SetApiKeyId gets a reference to the given string and assigns it to the ApiKeyId field.
func (o *NDSTenantEndpointAudit) SetApiKeyId(v string) {
	o.ApiKeyId = &v
}

// GetCreated returns the Created field value
func (o *NDSTenantEndpointAudit) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *NDSTenantEndpointAudit) SetCreated(v time.Time) {
	o.Created = v
}

// GetEndpointId returns the EndpointId field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetEndpointId() string {
	if o == nil || IsNil(o.EndpointId) {
		var ret string
		return ret
	}
	return *o.EndpointId
}

// GetEndpointIdOk returns a tuple with the EndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointId) {
		return nil, false
	}
	return o.EndpointId, true
}

// HasEndpointId returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasEndpointId() bool {
	if o != nil && !IsNil(o.EndpointId) {
		return true
	}

	return false
}

// SetEndpointId gets a reference to the given string and assigns it to the EndpointId field.
func (o *NDSTenantEndpointAudit) SetEndpointId(v string) {
	o.EndpointId = &v
}

// GetEventTypeName returns the EventTypeName field value
func (o *NDSTenantEndpointAudit) GetEventTypeName() NDSTenantEndpointAuditType {
	if o == nil {
		var ret NDSTenantEndpointAuditType
		return ret
	}

	return o.EventTypeName
}

// GetEventTypeNameOk returns a tuple with the EventTypeName field value
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetEventTypeNameOk() (*NDSTenantEndpointAuditType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTypeName, true
}

// SetEventTypeName sets field value
func (o *NDSTenantEndpointAudit) SetEventTypeName(v NDSTenantEndpointAuditType) {
	o.EventTypeName = v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *NDSTenantEndpointAudit) SetGroupId(v string) {
	o.GroupId = &v
}

// GetId returns the Id field value
func (o *NDSTenantEndpointAudit) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NDSTenantEndpointAudit) SetId(v string) {
	o.Id = v
}

// GetIsGlobalAdmin returns the IsGlobalAdmin field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetIsGlobalAdmin() bool {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		var ret bool
		return ret
	}
	return *o.IsGlobalAdmin
}

// GetIsGlobalAdminOk returns a tuple with the IsGlobalAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetIsGlobalAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsGlobalAdmin) {
		return nil, false
	}
	return o.IsGlobalAdmin, true
}

// HasIsGlobalAdmin returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasIsGlobalAdmin() bool {
	if o != nil && !IsNil(o.IsGlobalAdmin) {
		return true
	}

	return false
}

// SetIsGlobalAdmin gets a reference to the given bool and assigns it to the IsGlobalAdmin field.
func (o *NDSTenantEndpointAudit) SetIsGlobalAdmin(v bool) {
	o.IsGlobalAdmin = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetLinksOk() ([]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *NDSTenantEndpointAudit) SetLinks(v []Link) {
	o.Links = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *NDSTenantEndpointAudit) SetOrgId(v string) {
	o.OrgId = &v
}

// GetProviderEndpointId returns the ProviderEndpointId field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetProviderEndpointId() string {
	if o == nil || IsNil(o.ProviderEndpointId) {
		var ret string
		return ret
	}
	return *o.ProviderEndpointId
}

// GetProviderEndpointIdOk returns a tuple with the ProviderEndpointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetProviderEndpointIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderEndpointId) {
		return nil, false
	}
	return o.ProviderEndpointId, true
}

// HasProviderEndpointId returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasProviderEndpointId() bool {
	if o != nil && !IsNil(o.ProviderEndpointId) {
		return true
	}

	return false
}

// SetProviderEndpointId gets a reference to the given string and assigns it to the ProviderEndpointId field.
func (o *NDSTenantEndpointAudit) SetProviderEndpointId(v string) {
	o.ProviderEndpointId = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *NDSTenantEndpointAudit) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetRaw() Raw {
	if o == nil || IsNil(o.Raw) {
		var ret Raw
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetRawOk() (*Raw, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given Raw and assigns it to the Raw field.
func (o *NDSTenantEndpointAudit) SetRaw(v Raw) {
	o.Raw = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetRemoteAddress() string {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret string
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetRemoteAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given string and assigns it to the RemoteAddress field.
func (o *NDSTenantEndpointAudit) SetRemoteAddress(v string) {
	o.RemoteAddress = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *NDSTenantEndpointAudit) SetUserId(v string) {
	o.UserId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *NDSTenantEndpointAudit) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NDSTenantEndpointAudit) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *NDSTenantEndpointAudit) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *NDSTenantEndpointAudit) SetUsername(v string) {
	o.Username = &v
}

func (o NDSTenantEndpointAudit) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o NDSTenantEndpointAudit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventTypeName"] = o.EventTypeName
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	return toSerialize, nil
}

type NullableNDSTenantEndpointAudit struct {
	value *NDSTenantEndpointAudit
	isSet bool
}

func (v NullableNDSTenantEndpointAudit) Get() *NDSTenantEndpointAudit {
	return v.value
}

func (v *NullableNDSTenantEndpointAudit) Set(val *NDSTenantEndpointAudit) {
	v.value = val
	v.isSet = true
}

func (v NullableNDSTenantEndpointAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableNDSTenantEndpointAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNDSTenantEndpointAudit(val *NDSTenantEndpointAudit) *NullableNDSTenantEndpointAudit {
	return &NullableNDSTenantEndpointAudit{value: val, isSet: true}
}

func (v NullableNDSTenantEndpointAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNDSTenantEndpointAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


