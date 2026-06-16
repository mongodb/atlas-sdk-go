// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OrgDelegationSettingsUpdateRequest struct for OrgDelegationSettingsUpdateRequest
type OrgDelegationSettingsUpdateRequest struct {
	// Policy that controls how MCP (Model Context Protocol) delegated access is permitted within this organization. Possible values are `DISALLOWED`, `READ_ONLY`, and `READ_WRITE`. Defaults to `DISALLOWED`.
	DelegatedMcpAccess *string `json:"delegatedMcpAccess,omitempty"`
	// Policy that controls whether partner delegated access is permitted within this organization. Possible values are `DISALLOWED` and `READ_WRITE`. Defaults to `DISALLOWED`.
	DelegatedPartnerAccess *string `json:"delegatedPartnerAccess,omitempty"`
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
