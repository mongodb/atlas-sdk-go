// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// FederationSamlIdentityProvider struct for FederationSamlIdentityProvider
type FederationSamlIdentityProvider struct {
	// URL that points to where to send the SAML response.
	AcsUrl *string `json:"acsUrl,omitempty"`
	// List that contains the domains associated with the identity provider.
	AssociatedDomains *[]string `json:"associatedDomains,omitempty"`
	// List that contains the connected organization configurations associated with the identity provider.
	AssociatedOrgs *[]ConnectedOrgConfig `json:"associatedOrgs,omitempty"`
	// Unique string that identifies the intended audience of the SAML assertion.
	AudienceUri *string `json:"audienceUri,omitempty"`
	// Date that the identity provider was created on. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The description of the identity provider.
	Description *string `json:"description,omitempty"`
	// Human-readable label that identifies the identity provider.
	DisplayName *string `json:"displayName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the identity provider.
	// Read only field.
	Id string `json:"id"`
	// String enum that indicates the type of the identity provider. Default is WORKFORCE.
	IdpType *string `json:"idpType,omitempty"`
	// Unique string that identifies the issuer of the SAML Assertion or OIDC metadata/discovery document URL.
	IssuerUri *string `json:"issuerUri,omitempty"`
	// Legacy 20-hexadecimal digit string that identifies the identity provider.
	OktaIdpId   string       `json:"oktaIdpId"`
	PemFileInfo *PemFileInfo `json:"pemFileInfo,omitempty"`
	// String enum that indicates the protocol of the identity provider. Either SAML or OIDC.
	Protocol *string `json:"protocol,omitempty"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request.
	RequestBinding *string `json:"requestBinding,omitempty"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty"`
	// Custom SSO Url for the identity provider.
	Slug *string `json:"slug,omitempty"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled *bool `json:"ssoDebugEnabled,omitempty"`
	// URL that points to the receiver of the SAML authentication request.
	SsoUrl *string `json:"ssoUrl,omitempty"`
	// String enum that indicates whether the identity provider is active.
	Status *string `json:"status,omitempty"`
	// Date that the identity provider was last updated on. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewFederationSamlIdentityProvider instantiates a new FederationSamlIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationSamlIdentityProvider(id string, oktaIdpId string) *FederationSamlIdentityProvider {
	this := FederationSamlIdentityProvider{}
	this.Id = id
	this.OktaIdpId = oktaIdpId
	return &this
}

// NewFederationSamlIdentityProviderWithDefaults instantiates a new FederationSamlIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationSamlIdentityProviderWithDefaults() *FederationSamlIdentityProvider {
	this := FederationSamlIdentityProvider{}
	return &this
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetAcsUrl() string {
	if o == nil || IsNil(o.AcsUrl) {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetAcsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AcsUrl) {
		return nil, false
	}

	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasAcsUrl() bool {
	if o != nil && !IsNil(o.AcsUrl) {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *FederationSamlIdentityProvider) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAssociatedDomains returns the AssociatedDomains field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetAssociatedDomains() []string {
	if o == nil || IsNil(o.AssociatedDomains) {
		var ret []string
		return ret
	}
	return *o.AssociatedDomains
}

// GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetAssociatedDomainsOk() (*[]string, bool) {
	if o == nil || IsNil(o.AssociatedDomains) {
		return nil, false
	}

	return o.AssociatedDomains, true
}

// HasAssociatedDomains returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasAssociatedDomains() bool {
	if o != nil && !IsNil(o.AssociatedDomains) {
		return true
	}

	return false
}

// SetAssociatedDomains gets a reference to the given []string and assigns it to the AssociatedDomains field.
func (o *FederationSamlIdentityProvider) SetAssociatedDomains(v []string) {
	o.AssociatedDomains = &v
}

// GetAssociatedOrgs returns the AssociatedOrgs field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetAssociatedOrgs() []ConnectedOrgConfig {
	if o == nil || IsNil(o.AssociatedOrgs) {
		var ret []ConnectedOrgConfig
		return ret
	}
	return *o.AssociatedOrgs
}

// GetAssociatedOrgsOk returns a tuple with the AssociatedOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetAssociatedOrgsOk() (*[]ConnectedOrgConfig, bool) {
	if o == nil || IsNil(o.AssociatedOrgs) {
		return nil, false
	}

	return o.AssociatedOrgs, true
}

// HasAssociatedOrgs returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasAssociatedOrgs() bool {
	if o != nil && !IsNil(o.AssociatedOrgs) {
		return true
	}

	return false
}

// SetAssociatedOrgs gets a reference to the given []ConnectedOrgConfig and assigns it to the AssociatedOrgs field.
func (o *FederationSamlIdentityProvider) SetAssociatedOrgs(v []ConnectedOrgConfig) {
	o.AssociatedOrgs = &v
}

// GetAudienceUri returns the AudienceUri field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetAudienceUri() string {
	if o == nil || IsNil(o.AudienceUri) {
		var ret string
		return ret
	}
	return *o.AudienceUri
}

// GetAudienceUriOk returns a tuple with the AudienceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetAudienceUriOk() (*string, bool) {
	if o == nil || IsNil(o.AudienceUri) {
		return nil, false
	}

	return o.AudienceUri, true
}

// HasAudienceUri returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasAudienceUri() bool {
	if o != nil && !IsNil(o.AudienceUri) {
		return true
	}

	return false
}

// SetAudienceUri gets a reference to the given string and assigns it to the AudienceUri field.
func (o *FederationSamlIdentityProvider) SetAudienceUri(v string) {
	o.AudienceUri = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FederationSamlIdentityProvider) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FederationSamlIdentityProvider) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}

	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FederationSamlIdentityProvider) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value
func (o *FederationSamlIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FederationSamlIdentityProvider) SetId(v string) {
	o.Id = v
}

// GetIdpType returns the IdpType field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetIdpType() string {
	if o == nil || IsNil(o.IdpType) {
		var ret string
		return ret
	}
	return *o.IdpType
}

// GetIdpTypeOk returns a tuple with the IdpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetIdpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdpType) {
		return nil, false
	}

	return o.IdpType, true
}

// HasIdpType returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasIdpType() bool {
	if o != nil && !IsNil(o.IdpType) {
		return true
	}

	return false
}

// SetIdpType gets a reference to the given string and assigns it to the IdpType field.
func (o *FederationSamlIdentityProvider) SetIdpType(v string) {
	o.IdpType = &v
}

// GetIssuerUri returns the IssuerUri field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetIssuerUri() string {
	if o == nil || IsNil(o.IssuerUri) {
		var ret string
		return ret
	}
	return *o.IssuerUri
}

// GetIssuerUriOk returns a tuple with the IssuerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetIssuerUriOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerUri) {
		return nil, false
	}

	return o.IssuerUri, true
}

// HasIssuerUri returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasIssuerUri() bool {
	if o != nil && !IsNil(o.IssuerUri) {
		return true
	}

	return false
}

// SetIssuerUri gets a reference to the given string and assigns it to the IssuerUri field.
func (o *FederationSamlIdentityProvider) SetIssuerUri(v string) {
	o.IssuerUri = &v
}

// GetOktaIdpId returns the OktaIdpId field value
func (o *FederationSamlIdentityProvider) GetOktaIdpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaIdpId
}

// GetOktaIdpIdOk returns a tuple with the OktaIdpId field value
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetOktaIdpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaIdpId, true
}

// SetOktaIdpId sets field value
func (o *FederationSamlIdentityProvider) SetOktaIdpId(v string) {
	o.OktaIdpId = v
}

// GetPemFileInfo returns the PemFileInfo field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetPemFileInfo() PemFileInfo {
	if o == nil || IsNil(o.PemFileInfo) {
		var ret PemFileInfo
		return ret
	}
	return *o.PemFileInfo
}

// GetPemFileInfoOk returns a tuple with the PemFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetPemFileInfoOk() (*PemFileInfo, bool) {
	if o == nil || IsNil(o.PemFileInfo) {
		return nil, false
	}

	return o.PemFileInfo, true
}

// HasPemFileInfo returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasPemFileInfo() bool {
	if o != nil && !IsNil(o.PemFileInfo) {
		return true
	}

	return false
}

// SetPemFileInfo gets a reference to the given PemFileInfo and assigns it to the PemFileInfo field.
func (o *FederationSamlIdentityProvider) SetPemFileInfo(v PemFileInfo) {
	o.PemFileInfo = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}

	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *FederationSamlIdentityProvider) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRequestBinding returns the RequestBinding field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetRequestBinding() string {
	if o == nil || IsNil(o.RequestBinding) {
		var ret string
		return ret
	}
	return *o.RequestBinding
}

// GetRequestBindingOk returns a tuple with the RequestBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetRequestBindingOk() (*string, bool) {
	if o == nil || IsNil(o.RequestBinding) {
		return nil, false
	}

	return o.RequestBinding, true
}

// HasRequestBinding returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasRequestBinding() bool {
	if o != nil && !IsNil(o.RequestBinding) {
		return true
	}

	return false
}

// SetRequestBinding gets a reference to the given string and assigns it to the RequestBinding field.
func (o *FederationSamlIdentityProvider) SetRequestBinding(v string) {
	o.RequestBinding = &v
}

// GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetResponseSignatureAlgorithm() string {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.ResponseSignatureAlgorithm
}

// GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetResponseSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		return nil, false
	}

	return o.ResponseSignatureAlgorithm, true
}

// HasResponseSignatureAlgorithm returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasResponseSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.ResponseSignatureAlgorithm) {
		return true
	}

	return false
}

// SetResponseSignatureAlgorithm gets a reference to the given string and assigns it to the ResponseSignatureAlgorithm field.
func (o *FederationSamlIdentityProvider) SetResponseSignatureAlgorithm(v string) {
	o.ResponseSignatureAlgorithm = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}

	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *FederationSamlIdentityProvider) SetSlug(v string) {
	o.Slug = &v
}

// GetSsoDebugEnabled returns the SsoDebugEnabled field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetSsoDebugEnabled() bool {
	if o == nil || IsNil(o.SsoDebugEnabled) {
		var ret bool
		return ret
	}
	return *o.SsoDebugEnabled
}

// GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetSsoDebugEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SsoDebugEnabled) {
		return nil, false
	}

	return o.SsoDebugEnabled, true
}

// HasSsoDebugEnabled returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasSsoDebugEnabled() bool {
	if o != nil && !IsNil(o.SsoDebugEnabled) {
		return true
	}

	return false
}

// SetSsoDebugEnabled gets a reference to the given bool and assigns it to the SsoDebugEnabled field.
func (o *FederationSamlIdentityProvider) SetSsoDebugEnabled(v bool) {
	o.SsoDebugEnabled = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}

	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *FederationSamlIdentityProvider) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FederationSamlIdentityProvider) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise
func (o *FederationSamlIdentityProvider) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProvider) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}

	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FederationSamlIdentityProvider) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FederationSamlIdentityProvider) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}
