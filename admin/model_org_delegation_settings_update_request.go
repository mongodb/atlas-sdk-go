// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgDelegationSettingsUpdateRequest struct for OrgDelegationSettingsUpdateRequest
type OrgDelegationSettingsUpdateRequest struct {
	// Policy that controls how MCP (Model Context Protocol) delegated access is permitted within this organization. Possible values are `DISALLOWED`, `READ_ONLY`, and `READ_WRITE`. Defaults to `DISALLOWED`.
	DelegatedMcpAccess *string `json:"delegatedMcpAccess,omitempty"`
	// Policy that controls whether partner delegated access is permitted within this organization. Possible values are `DISALLOWED` and `READ_WRITE`. Defaults to `DISALLOWED`.
	DelegatedPartnerAccess *string `json:"delegatedPartnerAccess,omitempty"`
	// Maximum number of seconds a refresh token may be idle before it expires. Omit to leave unchanged; set to null to reset to the system default. Must be between 1 and 31536000 (1 year) when provided.
	IdleRefreshTokenLifetime *int64 `json:"idleRefreshTokenLifetime,omitempty"`
	// Maximum lifetime of a refresh token in seconds, regardless of activity. Omit to leave unchanged; set to null to reset to the system default. Must be between 1 and 31536000 (1 year) when provided.
	MaximumRefreshTokenLifetime *int64 `json:"maximumRefreshTokenLifetime,omitempty"`
}

// NewOrgDelegationSettingsUpdateRequest instantiates a new OrgDelegationSettingsUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgDelegationSettingsUpdateRequest() *OrgDelegationSettingsUpdateRequest {
	this := OrgDelegationSettingsUpdateRequest{}
	return &this
}

// NewOrgDelegationSettingsUpdateRequestWithDefaults instantiates a new OrgDelegationSettingsUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgDelegationSettingsUpdateRequestWithDefaults() *OrgDelegationSettingsUpdateRequest {
	this := OrgDelegationSettingsUpdateRequest{}
	return &this
}

// GetDelegatedMcpAccess returns the DelegatedMcpAccess field value if set, zero value otherwise
func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedMcpAccess() string {
	if o == nil || IsNil(o.DelegatedMcpAccess) {
		var ret string
		return ret
	}
	return *o.DelegatedMcpAccess
}

// GetDelegatedMcpAccessOk returns a tuple with the DelegatedMcpAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedMcpAccessOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedMcpAccess) {
		return nil, false
	}

	return o.DelegatedMcpAccess, true
}

// HasDelegatedMcpAccess returns a boolean if a field has been set.
func (o *OrgDelegationSettingsUpdateRequest) HasDelegatedMcpAccess() bool {
	if o != nil && !IsNil(o.DelegatedMcpAccess) {
		return true
	}

	return false
}

// SetDelegatedMcpAccess gets a reference to the given string and assigns it to the DelegatedMcpAccess field.
func (o *OrgDelegationSettingsUpdateRequest) SetDelegatedMcpAccess(v string) {
	o.DelegatedMcpAccess = &v
}

// GetDelegatedPartnerAccess returns the DelegatedPartnerAccess field value if set, zero value otherwise
func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedPartnerAccess() string {
	if o == nil || IsNil(o.DelegatedPartnerAccess) {
		var ret string
		return ret
	}
	return *o.DelegatedPartnerAccess
}

// GetDelegatedPartnerAccessOk returns a tuple with the DelegatedPartnerAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsUpdateRequest) GetDelegatedPartnerAccessOk() (*string, bool) {
	if o == nil || IsNil(o.DelegatedPartnerAccess) {
		return nil, false
	}

	return o.DelegatedPartnerAccess, true
}

// HasDelegatedPartnerAccess returns a boolean if a field has been set.
func (o *OrgDelegationSettingsUpdateRequest) HasDelegatedPartnerAccess() bool {
	if o != nil && !IsNil(o.DelegatedPartnerAccess) {
		return true
	}

	return false
}

// SetDelegatedPartnerAccess gets a reference to the given string and assigns it to the DelegatedPartnerAccess field.
func (o *OrgDelegationSettingsUpdateRequest) SetDelegatedPartnerAccess(v string) {
	o.DelegatedPartnerAccess = &v
}

// GetIdleRefreshTokenLifetime returns the IdleRefreshTokenLifetime field value if set, zero value otherwise
func (o *OrgDelegationSettingsUpdateRequest) GetIdleRefreshTokenLifetime() int64 {
	if o == nil || IsNil(o.IdleRefreshTokenLifetime) {
		var ret int64
		return ret
	}
	return *o.IdleRefreshTokenLifetime
}

// GetIdleRefreshTokenLifetimeOk returns a tuple with the IdleRefreshTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsUpdateRequest) GetIdleRefreshTokenLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.IdleRefreshTokenLifetime) {
		return nil, false
	}

	return o.IdleRefreshTokenLifetime, true
}

// HasIdleRefreshTokenLifetime returns a boolean if a field has been set.
func (o *OrgDelegationSettingsUpdateRequest) HasIdleRefreshTokenLifetime() bool {
	if o != nil && !IsNil(o.IdleRefreshTokenLifetime) {
		return true
	}

	return false
}

// SetIdleRefreshTokenLifetime gets a reference to the given int64 and assigns it to the IdleRefreshTokenLifetime field.
func (o *OrgDelegationSettingsUpdateRequest) SetIdleRefreshTokenLifetime(v int64) {
	o.IdleRefreshTokenLifetime = &v
}

// GetMaximumRefreshTokenLifetime returns the MaximumRefreshTokenLifetime field value if set, zero value otherwise
func (o *OrgDelegationSettingsUpdateRequest) GetMaximumRefreshTokenLifetime() int64 {
	if o == nil || IsNil(o.MaximumRefreshTokenLifetime) {
		var ret int64
		return ret
	}
	return *o.MaximumRefreshTokenLifetime
}

// GetMaximumRefreshTokenLifetimeOk returns a tuple with the MaximumRefreshTokenLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgDelegationSettingsUpdateRequest) GetMaximumRefreshTokenLifetimeOk() (*int64, bool) {
	if o == nil || IsNil(o.MaximumRefreshTokenLifetime) {
		return nil, false
	}

	return o.MaximumRefreshTokenLifetime, true
}

// HasMaximumRefreshTokenLifetime returns a boolean if a field has been set.
func (o *OrgDelegationSettingsUpdateRequest) HasMaximumRefreshTokenLifetime() bool {
	if o != nil && !IsNil(o.MaximumRefreshTokenLifetime) {
		return true
	}

	return false
}

// SetMaximumRefreshTokenLifetime gets a reference to the given int64 and assigns it to the MaximumRefreshTokenLifetime field.
func (o *OrgDelegationSettingsUpdateRequest) SetMaximumRefreshTokenLifetime(v int64) {
	o.MaximumRefreshTokenLifetime = &v
}
