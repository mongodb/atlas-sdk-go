// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgDelegationSettingsResponse struct for OrgDelegationSettingsResponse
type OrgDelegationSettingsResponse struct {
	// Policy that controls how MCP (Model Context Protocol) delegated access is permitted within this organization. Possible values are `DISALLOWED`, `READ_ONLY`, and `READ_WRITE`. Defaults to `DISALLOWED`.
	DelegatedMcpAccess *string `json:"delegatedMcpAccess,omitempty"`
	// Policy that controls whether partner delegated access is permitted within this organization. Possible values are `DISALLOWED` and `READ_WRITE`. Defaults to `DISALLOWED`.
	DelegatedPartnerAccess *string `json:"delegatedPartnerAccess,omitempty"`
	// Maximum number of seconds a refresh token may be idle before it expires. When not set, the system default applies.
	IdleRefreshTokenLifetime *int64 `json:"idleRefreshTokenLifetime,omitempty"`
	// Maximum lifetime of a refresh token in seconds, regardless of activity. When not set, the system default applies.
	MaximumRefreshTokenLifetime *int64 `json:"maximumRefreshTokenLifetime,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OrgDelegationSettingsResponse) MarshalJSON() ([]byte, error) {
	type noMethod OrgDelegationSettingsResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOrgDelegationSettingsResponse instantiates a new OrgDelegationSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgDelegationSettingsResponse() *OrgDelegationSettingsResponse {
	this := OrgDelegationSettingsResponse{}
	return &this
}

// NewOrgDelegationSettingsResponseWithDefaults instantiates a new OrgDelegationSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgDelegationSettingsResponseWithDefaults() *OrgDelegationSettingsResponse {
	this := OrgDelegationSettingsResponse{}
	return &this
}

// GetDelegatedMcpAccess returns the DelegatedMcpAccess field value if set, zero value otherwise
func (o *OrgDelegationSettingsResponse) GetDelegatedMcpAccess() string {
	if o == nil || IsNil(o.DelegatedMcpAccess) {
		var ret string
		return ret
	}
	return *o.DelegatedMcpAccess
}

// GetDelegatedMcpAccessOk returns a tuple with the DelegatedMcpAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsResponse) GetDelegatedMcpAccessOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedMcpAccess) {
		return nil, false
	}

	return o.DelegatedMcpAccess, true
}

// HasDelegatedMcpAccess returns a boolean if a field has been set.
func (o *OrgDelegationSettingsResponse) HasDelegatedMcpAccess() bool {
	if o != nil && !IsNil(o.DelegatedMcpAccess) {
		return true
	}

	return false
}

// SetDelegatedMcpAccess gets a reference to the given string and assigns it to the DelegatedMcpAccess field.
func (o *OrgDelegationSettingsResponse) SetDelegatedMcpAccess(v string) {
	o.DelegatedMcpAccess = &v
	o.NullFields = removeNullField(o.NullFields, "DelegatedMcpAccess")
}

// SetDelegatedMcpAccessNil sets DelegatedMcpAccess to an explicit JSON null when marshaled.
func (o *OrgDelegationSettingsResponse) SetDelegatedMcpAccessNil() {
	o.DelegatedMcpAccess = nil
	o.NullFields = addNullField(o.NullFields, "DelegatedMcpAccess")
}

// GetDelegatedPartnerAccess returns the DelegatedPartnerAccess field value if set, zero value otherwise
func (o *OrgDelegationSettingsResponse) GetDelegatedPartnerAccess() string {
	if o == nil || IsNil(o.DelegatedPartnerAccess) {
		var ret string
		return ret
	}
	return *o.DelegatedPartnerAccess
}

// GetDelegatedPartnerAccessOk returns a tuple with the DelegatedPartnerAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsResponse) GetDelegatedPartnerAccessOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedPartnerAccess) {
		return nil, false
	}

	return o.DelegatedPartnerAccess, true
}

// HasDelegatedPartnerAccess returns a boolean if a field has been set.
func (o *OrgDelegationSettingsResponse) HasDelegatedPartnerAccess() bool {
	if o != nil && !IsNil(o.DelegatedPartnerAccess) {
		return true
	}

	return false
}

// SetDelegatedPartnerAccess gets a reference to the given string and assigns it to the DelegatedPartnerAccess field.
func (o *OrgDelegationSettingsResponse) SetDelegatedPartnerAccess(v string) {
	o.DelegatedPartnerAccess = &v
	o.NullFields = removeNullField(o.NullFields, "DelegatedPartnerAccess")
}

// SetDelegatedPartnerAccessNil sets DelegatedPartnerAccess to an explicit JSON null when marshaled.
func (o *OrgDelegationSettingsResponse) SetDelegatedPartnerAccessNil() {
	o.DelegatedPartnerAccess = nil
	o.NullFields = addNullField(o.NullFields, "DelegatedPartnerAccess")
}

// GetIdleRefreshTokenLifetime returns the IdleRefreshTokenLifetime field value if set, zero value otherwise
func (o *OrgDelegationSettingsResponse) GetIdleRefreshTokenLifetime() int64 {
	if o == nil || IsNil(o.IdleRefreshTokenLifetime) {
		var ret int64
		return ret
	}
	return *o.IdleRefreshTokenLifetime
}

// GetIdleRefreshTokenLifetimeOk returns a tuple with the IdleRefreshTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsResponse) GetIdleRefreshTokenLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.IdleRefreshTokenLifetime) {
		return nil, false
	}

	return o.IdleRefreshTokenLifetime, true
}

// HasIdleRefreshTokenLifetime returns a boolean if a field has been set.
func (o *OrgDelegationSettingsResponse) HasIdleRefreshTokenLifetime() bool {
	if o != nil && !IsNil(o.IdleRefreshTokenLifetime) {
		return true
	}

	return false
}

// SetIdleRefreshTokenLifetime gets a reference to the given int64 and assigns it to the IdleRefreshTokenLifetime field.
func (o *OrgDelegationSettingsResponse) SetIdleRefreshTokenLifetime(v int64) {
	o.IdleRefreshTokenLifetime = &v
	o.NullFields = removeNullField(o.NullFields, "IdleRefreshTokenLifetime")
}

// SetIdleRefreshTokenLifetimeNil sets IdleRefreshTokenLifetime to an explicit JSON null when marshaled.
func (o *OrgDelegationSettingsResponse) SetIdleRefreshTokenLifetimeNil() {
	o.IdleRefreshTokenLifetime = nil
	o.NullFields = addNullField(o.NullFields, "IdleRefreshTokenLifetime")
}

// GetMaximumRefreshTokenLifetime returns the MaximumRefreshTokenLifetime field value if set, zero value otherwise
func (o *OrgDelegationSettingsResponse) GetMaximumRefreshTokenLifetime() int64 {
	if o == nil || IsNil(o.MaximumRefreshTokenLifetime) {
		var ret int64
		return ret
	}
	return *o.MaximumRefreshTokenLifetime
}

// GetMaximumRefreshTokenLifetimeOk returns a tuple with the MaximumRefreshTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsResponse) GetMaximumRefreshTokenLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumRefreshTokenLifetime) {
		return nil, false
	}

	return o.MaximumRefreshTokenLifetime, true
}

// HasMaximumRefreshTokenLifetime returns a boolean if a field has been set.
func (o *OrgDelegationSettingsResponse) HasMaximumRefreshTokenLifetime() bool {
	if o != nil && !IsNil(o.MaximumRefreshTokenLifetime) {
		return true
	}

	return false
}

// SetMaximumRefreshTokenLifetime gets a reference to the given int64 and assigns it to the MaximumRefreshTokenLifetime field.
func (o *OrgDelegationSettingsResponse) SetMaximumRefreshTokenLifetime(v int64) {
	o.MaximumRefreshTokenLifetime = &v
	o.NullFields = removeNullField(o.NullFields, "MaximumRefreshTokenLifetime")
}

// SetMaximumRefreshTokenLifetimeNil sets MaximumRefreshTokenLifetime to an explicit JSON null when marshaled.
func (o *OrgDelegationSettingsResponse) SetMaximumRefreshTokenLifetimeNil() {
	o.MaximumRefreshTokenLifetime = nil
	o.NullFields = addNullField(o.NullFields, "MaximumRefreshTokenLifetime")
}
