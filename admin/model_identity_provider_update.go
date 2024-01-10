// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// IdentityProviderUpdate struct for IdentityProviderUpdate
type IdentityProviderUpdate struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains *[]string `json:"associatedDomains,omitempty"`
	// The description for the identity provider.
	Description *string `json:"description,omitempty"`
	// Human-readable label that identifies the identity provider.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique string that identifies the issuer of the SAML Assertion.
	IssuerUri *string `json:"issuerUri,omitempty"`
	// The protocol for the identity provider.
	Protocol *string `json:"protocol,omitempty"`
	// Audience claim for the identity provider.
	AudienceClaim *[]string `json:"audienceClaim,omitempty"`
	// Client ID for the identity provider.
	ClientId *string `json:"clientId,omitempty"`
	// Groups claim for the identity provider.
	GroupsClaim *string `json:"groupsClaim,omitempty"`
	// Requested scopes for the identity provider.
	RequestedScopes *[]string `json:"requestedScopes,omitempty"`
	// User claim for the identity provider.
	UserClaim   *string            `json:"userClaim,omitempty"`
	PemFileInfo *PemFileInfoUpdate `json:"pemFileInfo,omitempty"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request.
	RequestBinding *string `json:"requestBinding,omitempty"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty"`
	// Custom SSO Url for identity provider.
	Slug *string `json:"slug,omitempty"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled *bool `json:"ssoDebugEnabled,omitempty"`
	// Unique string that identifies the intended audience of the SAML assertion.
	SsoUrl *string `json:"ssoUrl,omitempty"`
	// String enum that indicates whether the identity provider is active.
	Status *string `json:"status,omitempty"`
}

// NewIdentityProviderUpdate instantiates a new IdentityProviderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderUpdate() *IdentityProviderUpdate {
	this := IdentityProviderUpdate{}
	return &this
}

// NewIdentityProviderUpdateWithDefaults instantiates a new IdentityProviderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderUpdateWithDefaults() *IdentityProviderUpdate {
	this := IdentityProviderUpdate{}
	return &this
}

// GetAssociatedDomains returns the AssociatedDomains field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetAssociatedDomains() []string {
	if o == nil || IsNil(o.AssociatedDomains) {
		var ret []string
		return ret
	}
	return *o.AssociatedDomains
}

// GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetAssociatedDomainsOk() (*[]string, bool) {
	if o == nil || IsNil(o.AssociatedDomains) {
		return nil, false
	}

	return o.AssociatedDomains, true
}

// HasAssociatedDomains returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasAssociatedDomains() bool {
	if o != nil && !IsNil(o.AssociatedDomains) {
		return true
	}

	return false
}

// SetAssociatedDomains gets a reference to the given []string and assigns it to the AssociatedDomains field.
func (o *IdentityProviderUpdate) SetAssociatedDomains(v []string) {
	o.AssociatedDomains = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityProviderUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}

	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentityProviderUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIssuerUri returns the IssuerUri field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetIssuerUri() string {
	if o == nil || IsNil(o.IssuerUri) {
		var ret string
		return ret
	}
	return *o.IssuerUri
}

// GetIssuerUriOk returns a tuple with the IssuerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetIssuerUriOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerUri) {
		return nil, false
	}

	return o.IssuerUri, true
}

// HasIssuerUri returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasIssuerUri() bool {
	if o != nil && !IsNil(o.IssuerUri) {
		return true
	}

	return false
}

// SetIssuerUri gets a reference to the given string and assigns it to the IssuerUri field.
func (o *IdentityProviderUpdate) SetIssuerUri(v string) {
	o.IssuerUri = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}

	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *IdentityProviderUpdate) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAudienceClaim returns the AudienceClaim field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetAudienceClaim() []string {
	if o == nil || IsNil(o.AudienceClaim) {
		var ret []string
		return ret
	}
	return *o.AudienceClaim
}

// GetAudienceClaimOk returns a tuple with the AudienceClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetAudienceClaimOk() (*[]string, bool) {
	if o == nil || IsNil(o.AudienceClaim) {
		return nil, false
	}

	return o.AudienceClaim, true
}

// HasAudienceClaim returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasAudienceClaim() bool {
	if o != nil && !IsNil(o.AudienceClaim) {
		return true
	}

	return false
}

// SetAudienceClaim gets a reference to the given []string and assigns it to the AudienceClaim field.
func (o *IdentityProviderUpdate) SetAudienceClaim(v []string) {
	o.AudienceClaim = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}

	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *IdentityProviderUpdate) SetClientId(v string) {
	o.ClientId = &v
}

// GetGroupsClaim returns the GroupsClaim field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetGroupsClaim() string {
	if o == nil || IsNil(o.GroupsClaim) {
		var ret string
		return ret
	}
	return *o.GroupsClaim
}

// GetGroupsClaimOk returns a tuple with the GroupsClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetGroupsClaimOk() (*string, bool) {
	if o == nil || IsNil(o.GroupsClaim) {
		return nil, false
	}

	return o.GroupsClaim, true
}

// HasGroupsClaim returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasGroupsClaim() bool {
	if o != nil && !IsNil(o.GroupsClaim) {
		return true
	}

	return false
}

// SetGroupsClaim gets a reference to the given string and assigns it to the GroupsClaim field.
func (o *IdentityProviderUpdate) SetGroupsClaim(v string) {
	o.GroupsClaim = &v
}

// GetRequestedScopes returns the RequestedScopes field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetRequestedScopes() []string {
	if o == nil || IsNil(o.RequestedScopes) {
		var ret []string
		return ret
	}
	return *o.RequestedScopes
}

// GetRequestedScopesOk returns a tuple with the RequestedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetRequestedScopesOk() (*[]string, bool) {
	if o == nil || IsNil(o.RequestedScopes) {
		return nil, false
	}

	return o.RequestedScopes, true
}

// HasRequestedScopes returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasRequestedScopes() bool {
	if o != nil && !IsNil(o.RequestedScopes) {
		return true
	}

	return false
}

// SetRequestedScopes gets a reference to the given []string and assigns it to the RequestedScopes field.
func (o *IdentityProviderUpdate) SetRequestedScopes(v []string) {
	o.RequestedScopes = &v
}

// GetUserClaim returns the UserClaim field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetUserClaim() string {
	if o == nil || IsNil(o.UserClaim) {
		var ret string
		return ret
	}
	return *o.UserClaim
}

// GetUserClaimOk returns a tuple with the UserClaim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetUserClaimOk() (*string, bool) {
	if o == nil || IsNil(o.UserClaim) {
		return nil, false
	}

	return o.UserClaim, true
}

// HasUserClaim returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasUserClaim() bool {
	if o != nil && !IsNil(o.UserClaim) {
		return true
	}

	return false
}

// SetUserClaim gets a reference to the given string and assigns it to the UserClaim field.
func (o *IdentityProviderUpdate) SetUserClaim(v string) {
	o.UserClaim = &v
}

// GetPemFileInfo returns the PemFileInfo field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetPemFileInfo() PemFileInfoUpdate {
	if o == nil || IsNil(o.PemFileInfo) {
		var ret PemFileInfoUpdate
		return ret
	}
	return *o.PemFileInfo
}

// GetPemFileInfoOk returns a tuple with the PemFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetPemFileInfoOk() (*PemFileInfoUpdate, bool) {
	if o == nil || IsNil(o.PemFileInfo) {
		return nil, false
	}

	return o.PemFileInfo, true
}

// HasPemFileInfo returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasPemFileInfo() bool {
	if o != nil && !IsNil(o.PemFileInfo) {
		return true
	}

	return false
}

// SetPemFileInfo gets a reference to the given PemFileInfoUpdate and assigns it to the PemFileInfo field.
func (o *IdentityProviderUpdate) SetPemFileInfo(v PemFileInfoUpdate) {
	o.PemFileInfo = &v
}

// GetRequestBinding returns the RequestBinding field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetRequestBinding() string {
	if o == nil || IsNil(o.RequestBinding) {
		var ret string
		return ret
	}
	return *o.RequestBinding
}

// GetRequestBindingOk returns a tuple with the RequestBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetRequestBindingOk() (*string, bool) {
	if o == nil || IsNil(o.RequestBinding) {
		return nil, false
	}

	return o.RequestBinding, true
}

// HasRequestBinding returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasRequestBinding() bool {
	if o != nil && !IsNil(o.RequestBinding) {
		return true
	}

	return false
}

// SetRequestBinding gets a reference to the given string and assigns it to the RequestBinding field.
func (o *IdentityProviderUpdate) SetRequestBinding(v string) {
	o.RequestBinding = &v
}

// GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetResponseSignatureAlgorithm() string {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.ResponseSignatureAlgorithm
}

// GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetResponseSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		return nil, false
	}

	return o.ResponseSignatureAlgorithm, true
}

// HasResponseSignatureAlgorithm returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasResponseSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.ResponseSignatureAlgorithm) {
		return true
	}

	return false
}

// SetResponseSignatureAlgorithm gets a reference to the given string and assigns it to the ResponseSignatureAlgorithm field.
func (o *IdentityProviderUpdate) SetResponseSignatureAlgorithm(v string) {
	o.ResponseSignatureAlgorithm = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}

	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *IdentityProviderUpdate) SetSlug(v string) {
	o.Slug = &v
}

// GetSsoDebugEnabled returns the SsoDebugEnabled field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetSsoDebugEnabled() bool {
	if o == nil || IsNil(o.SsoDebugEnabled) {
		var ret bool
		return ret
	}
	return *o.SsoDebugEnabled
}

// GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetSsoDebugEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SsoDebugEnabled) {
		return nil, false
	}

	return o.SsoDebugEnabled, true
}

// HasSsoDebugEnabled returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasSsoDebugEnabled() bool {
	if o != nil && !IsNil(o.SsoDebugEnabled) {
		return true
	}

	return false
}

// SetSsoDebugEnabled gets a reference to the given bool and assigns it to the SsoDebugEnabled field.
func (o *IdentityProviderUpdate) SetSsoDebugEnabled(v bool) {
	o.SsoDebugEnabled = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}

	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *IdentityProviderUpdate) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *IdentityProviderUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentityProviderUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentityProviderUpdate) SetStatus(v string) {
	o.Status = &v
}

func (o IdentityProviderUpdate) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o IdentityProviderUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedDomains) {
		toSerialize["associatedDomains"] = o.AssociatedDomains
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IssuerUri) {
		toSerialize["issuerUri"] = o.IssuerUri
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.AudienceClaim) {
		toSerialize["audienceClaim"] = o.AudienceClaim
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.GroupsClaim) {
		toSerialize["groupsClaim"] = o.GroupsClaim
	}
	if !IsNil(o.RequestedScopes) {
		toSerialize["requestedScopes"] = o.RequestedScopes
	}
	if !IsNil(o.UserClaim) {
		toSerialize["userClaim"] = o.UserClaim
	}
	if !IsNil(o.PemFileInfo) {
		toSerialize["pemFileInfo"] = o.PemFileInfo
	}
	if !IsNil(o.RequestBinding) {
		toSerialize["requestBinding"] = o.RequestBinding
	}
	if !IsNil(o.ResponseSignatureAlgorithm) {
		toSerialize["responseSignatureAlgorithm"] = o.ResponseSignatureAlgorithm
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.SsoDebugEnabled) {
		toSerialize["ssoDebugEnabled"] = o.SsoDebugEnabled
	}
	if !IsNil(o.SsoUrl) {
		toSerialize["ssoUrl"] = o.SsoUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}
