// Code based on the AtlasAPI V2 OpenAPI file

package admin

// GroupMigrationRequest struct for GroupMigrationRequest
type GroupMigrationRequest struct {
	// Unique 24-hexadecimal digit string that identifies the organization to move the specified project to.
	DestinationOrgId *string `json:"destinationOrgId,omitempty"`
	// Unique string that identifies the private part of the API Key used to verify access to the destination organization. This parameter is required only when you authenticate with Programmatic API Keys.
	DestinationOrgPrivateApiKey *string `json:"destinationOrgPrivateApiKey,omitempty"`
	// Unique string that identifies the public part of the API Key used to verify access to the destination organization. This parameter is required only when you authenticate with Programmatic API Keys.
	DestinationOrgPublicApiKey *string `json:"destinationOrgPublicApiKey,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *GroupMigrationRequest) MarshalJSON() ([]byte, error) {
	type noMethod GroupMigrationRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewGroupMigrationRequest instantiates a new GroupMigrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupMigrationRequest() *GroupMigrationRequest {
	this := GroupMigrationRequest{}
	return &this
}

// NewGroupMigrationRequestWithDefaults instantiates a new GroupMigrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupMigrationRequestWithDefaults() *GroupMigrationRequest {
	this := GroupMigrationRequest{}
	return &this
}

// GetDestinationOrgId returns the DestinationOrgId field value if set, zero value otherwise
func (o *GroupMigrationRequest) GetDestinationOrgId() string {
	if o == nil || IsNil(o.DestinationOrgId) {
		var ret string
		return ret
	}
	return *o.DestinationOrgId
}

// GetDestinationOrgIdOk returns a tuple with the DestinationOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMigrationRequest) GetDestinationOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationOrgId) {
		return nil, false
	}

	return o.DestinationOrgId, true
}

// HasDestinationOrgId returns a boolean if a field has been set.
func (o *GroupMigrationRequest) HasDestinationOrgId() bool {
	if o != nil && !IsNil(o.DestinationOrgId) {
		return true
	}

	return false
}

// SetDestinationOrgId gets a reference to the given string and assigns it to the DestinationOrgId field.
func (o *GroupMigrationRequest) SetDestinationOrgId(v string) {
	o.DestinationOrgId = &v
}

// SetDestinationOrgIdNil sets DestinationOrgId to an explicit JSON null when marshaled.
func (o *GroupMigrationRequest) SetDestinationOrgIdNil() {
	o.DestinationOrgId = nil
	o.NullFields = append(o.NullFields, "DestinationOrgId")
}

// GetDestinationOrgPrivateApiKey returns the DestinationOrgPrivateApiKey field value if set, zero value otherwise
func (o *GroupMigrationRequest) GetDestinationOrgPrivateApiKey() string {
	if o == nil || IsNil(o.DestinationOrgPrivateApiKey) {
		var ret string
		return ret
	}
	return *o.DestinationOrgPrivateApiKey
}

// GetDestinationOrgPrivateApiKeyOk returns a tuple with the DestinationOrgPrivateApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMigrationRequest) GetDestinationOrgPrivateApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationOrgPrivateApiKey) {
		return nil, false
	}

	return o.DestinationOrgPrivateApiKey, true
}

// HasDestinationOrgPrivateApiKey returns a boolean if a field has been set.
func (o *GroupMigrationRequest) HasDestinationOrgPrivateApiKey() bool {
	if o != nil && !IsNil(o.DestinationOrgPrivateApiKey) {
		return true
	}

	return false
}

// SetDestinationOrgPrivateApiKey gets a reference to the given string and assigns it to the DestinationOrgPrivateApiKey field.
func (o *GroupMigrationRequest) SetDestinationOrgPrivateApiKey(v string) {
	o.DestinationOrgPrivateApiKey = &v
}

// SetDestinationOrgPrivateApiKeyNil sets DestinationOrgPrivateApiKey to an explicit JSON null when marshaled.
func (o *GroupMigrationRequest) SetDestinationOrgPrivateApiKeyNil() {
	o.DestinationOrgPrivateApiKey = nil
	o.NullFields = append(o.NullFields, "DestinationOrgPrivateApiKey")
}

// GetDestinationOrgPublicApiKey returns the DestinationOrgPublicApiKey field value if set, zero value otherwise
func (o *GroupMigrationRequest) GetDestinationOrgPublicApiKey() string {
	if o == nil || IsNil(o.DestinationOrgPublicApiKey) {
		var ret string
		return ret
	}
	return *o.DestinationOrgPublicApiKey
}

// GetDestinationOrgPublicApiKeyOk returns a tuple with the DestinationOrgPublicApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupMigrationRequest) GetDestinationOrgPublicApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DestinationOrgPublicApiKey) {
		return nil, false
	}

	return o.DestinationOrgPublicApiKey, true
}

// HasDestinationOrgPublicApiKey returns a boolean if a field has been set.
func (o *GroupMigrationRequest) HasDestinationOrgPublicApiKey() bool {
	if o != nil && !IsNil(o.DestinationOrgPublicApiKey) {
		return true
	}

	return false
}

// SetDestinationOrgPublicApiKey gets a reference to the given string and assigns it to the DestinationOrgPublicApiKey field.
func (o *GroupMigrationRequest) SetDestinationOrgPublicApiKey(v string) {
	o.DestinationOrgPublicApiKey = &v
}

// SetDestinationOrgPublicApiKeyNil sets DestinationOrgPublicApiKey to an explicit JSON null when marshaled.
func (o *GroupMigrationRequest) SetDestinationOrgPublicApiKeyNil() {
	o.DestinationOrgPublicApiKey = nil
	o.NullFields = append(o.NullFields, "DestinationOrgPublicApiKey")
}
