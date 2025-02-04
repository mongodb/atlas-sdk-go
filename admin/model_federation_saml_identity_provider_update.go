// Code based on the AtlasAPI V2 OpenAPI file

package admin

// FederationSamlIdentityProviderUpdate struct for FederationSamlIdentityProviderUpdate
type FederationSamlIdentityProviderUpdate struct {
	// List that contains the domains associated with the identity provider.
	AssociatedDomains *[]string `json:"associatedDomains,omitempty"`
	// The description of the identity provider.
	Description *string `json:"description,omitempty"`
	// Human-readable label that identifies the identity provider.
	DisplayName *string `json:"displayName,omitempty"`
	// String enum that indicates the type of the identity provider. Default is WORKFORCE.
	IdpType *string `json:"idpType,omitempty"`
	// Unique string that identifies the issuer of the SAML Assertion or OIDC metadata/discovery document URL.
	IssuerUri   *string            `json:"issuerUri,omitempty"`
	PemFileInfo *PemFileInfoUpdate `json:"pemFileInfo,omitempty"`
	// String enum that indicates the protocol of the identity provider. Either SAML or OIDC.
	Protocol *string `json:"protocol,omitempty"`
	// SAML Authentication Request Protocol HTTP method binding (POST or REDIRECT) that Federated Authentication uses to send the authentication request.
	RequestBinding *string `json:"requestBinding,omitempty"`
	// Signature algorithm that Federated Authentication uses to encrypt the identity provider signature.
	ResponseSignatureAlgorithm *string `json:"responseSignatureAlgorithm,omitempty"`
	// Custom SSO Url for the identity provider.
	Slug *string `json:"slug,omitempty"`
	// Flag that indicates whether the identity provider has SSO debug enabled.
	SsoDebugEnabled bool `json:"ssoDebugEnabled"`
	// URL that points to the receiver of the SAML authentication request.
	SsoUrl *string `json:"ssoUrl,omitempty"`
	// String enum that indicates whether the identity provider is active.
	Status *string `json:"status,omitempty"`
}

// NewFederationSamlIdentityProviderUpdate instantiates a new FederationSamlIdentityProviderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederationSamlIdentityProviderUpdate(ssoDebugEnabled bool) *FederationSamlIdentityProviderUpdate {
	this := FederationSamlIdentityProviderUpdate{}
	this.SsoDebugEnabled = ssoDebugEnabled
	return &this
}

// NewFederationSamlIdentityProviderUpdateWithDefaults instantiates a new FederationSamlIdentityProviderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederationSamlIdentityProviderUpdateWithDefaults() *FederationSamlIdentityProviderUpdate {
	this := FederationSamlIdentityProviderUpdate{}
	return &this
}

// GetAssociatedDomains returns the AssociatedDomains field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetAssociatedDomains() []string {
	if o == nil || IsNil(o.AssociatedDomains) {
		var ret []string
		return ret
	}
	return *o.AssociatedDomains
}

// GetAssociatedDomainsOk returns a tuple with the AssociatedDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetAssociatedDomainsOk() (*[]string, bool) {
	if o == nil || IsNil(o.AssociatedDomains) {
		return nil, false
	}

	return o.AssociatedDomains, true
}

// HasAssociatedDomains returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasAssociatedDomains() bool {
	if o != nil && !IsNil(o.AssociatedDomains) {
		return true
	}

	return false
}

// SetAssociatedDomains gets a reference to the given []string and assigns it to the AssociatedDomains field.
func (o *FederationSamlIdentityProviderUpdate) SetAssociatedDomains(v []string) {
	o.AssociatedDomains = &v
}

// GetDescription returns the Description field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}

	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FederationSamlIdentityProviderUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}

	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FederationSamlIdentityProviderUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIdpType returns the IdpType field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetIdpType() string {
	if o == nil || IsNil(o.IdpType) {
		var ret string
		return ret
	}
	return *o.IdpType
}

// GetIdpTypeOk returns a tuple with the IdpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetIdpTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IdpType) {
		return nil, false
	}

	return o.IdpType, true
}

// HasIdpType returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasIdpType() bool {
	if o != nil && !IsNil(o.IdpType) {
		return true
	}

	return false
}

// SetIdpType gets a reference to the given string and assigns it to the IdpType field.
func (o *FederationSamlIdentityProviderUpdate) SetIdpType(v string) {
	o.IdpType = &v
}

// GetIssuerUri returns the IssuerUri field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetIssuerUri() string {
	if o == nil || IsNil(o.IssuerUri) {
		var ret string
		return ret
	}
	return *o.IssuerUri
}

// GetIssuerUriOk returns a tuple with the IssuerUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetIssuerUriOk() (*string, bool) {
	if o == nil || IsNil(o.IssuerUri) {
		return nil, false
	}

	return o.IssuerUri, true
}

// HasIssuerUri returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasIssuerUri() bool {
	if o != nil && !IsNil(o.IssuerUri) {
		return true
	}

	return false
}

// SetIssuerUri gets a reference to the given string and assigns it to the IssuerUri field.
func (o *FederationSamlIdentityProviderUpdate) SetIssuerUri(v string) {
	o.IssuerUri = &v
}

// GetPemFileInfo returns the PemFileInfo field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetPemFileInfo() PemFileInfoUpdate {
	if o == nil || IsNil(o.PemFileInfo) {
		var ret PemFileInfoUpdate
		return ret
	}
	return *o.PemFileInfo
}

// GetPemFileInfoOk returns a tuple with the PemFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetPemFileInfoOk() (*PemFileInfoUpdate, bool) {
	if o == nil || IsNil(o.PemFileInfo) {
		return nil, false
	}

	return o.PemFileInfo, true
}

// HasPemFileInfo returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasPemFileInfo() bool {
	if o != nil && !IsNil(o.PemFileInfo) {
		return true
	}

	return false
}

// SetPemFileInfo gets a reference to the given PemFileInfoUpdate and assigns it to the PemFileInfo field.
func (o *FederationSamlIdentityProviderUpdate) SetPemFileInfo(v PemFileInfoUpdate) {
	o.PemFileInfo = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}

	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *FederationSamlIdentityProviderUpdate) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRequestBinding returns the RequestBinding field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetRequestBinding() string {
	if o == nil || IsNil(o.RequestBinding) {
		var ret string
		return ret
	}
	return *o.RequestBinding
}

// GetRequestBindingOk returns a tuple with the RequestBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetRequestBindingOk() (*string, bool) {
	if o == nil || IsNil(o.RequestBinding) {
		return nil, false
	}

	return o.RequestBinding, true
}

// HasRequestBinding returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasRequestBinding() bool {
	if o != nil && !IsNil(o.RequestBinding) {
		return true
	}

	return false
}

// SetRequestBinding gets a reference to the given string and assigns it to the RequestBinding field.
func (o *FederationSamlIdentityProviderUpdate) SetRequestBinding(v string) {
	o.RequestBinding = &v
}

// GetResponseSignatureAlgorithm returns the ResponseSignatureAlgorithm field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetResponseSignatureAlgorithm() string {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		var ret string
		return ret
	}
	return *o.ResponseSignatureAlgorithm
}

// GetResponseSignatureAlgorithmOk returns a tuple with the ResponseSignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetResponseSignatureAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseSignatureAlgorithm) {
		return nil, false
	}

	return o.ResponseSignatureAlgorithm, true
}

// HasResponseSignatureAlgorithm returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasResponseSignatureAlgorithm() bool {
	if o != nil && !IsNil(o.ResponseSignatureAlgorithm) {
		return true
	}

	return false
}

// SetResponseSignatureAlgorithm gets a reference to the given string and assigns it to the ResponseSignatureAlgorithm field.
func (o *FederationSamlIdentityProviderUpdate) SetResponseSignatureAlgorithm(v string) {
	o.ResponseSignatureAlgorithm = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}

	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *FederationSamlIdentityProviderUpdate) SetSlug(v string) {
	o.Slug = &v
}

// GetSsoDebugEnabled returns the SsoDebugEnabled field value
func (o *FederationSamlIdentityProviderUpdate) GetSsoDebugEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SsoDebugEnabled
}

// GetSsoDebugEnabledOk returns a tuple with the SsoDebugEnabled field value
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetSsoDebugEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SsoDebugEnabled, true
}

// SetSsoDebugEnabled sets field value
func (o *FederationSamlIdentityProviderUpdate) SetSsoDebugEnabled(v bool) {
	o.SsoDebugEnabled = v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetSsoUrl() string {
	if o == nil || IsNil(o.SsoUrl) {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetSsoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SsoUrl) {
		return nil, false
	}

	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasSsoUrl() bool {
	if o != nil && !IsNil(o.SsoUrl) {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *FederationSamlIdentityProviderUpdate) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise
func (o *FederationSamlIdentityProviderUpdate) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederationSamlIdentityProviderUpdate) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}

	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FederationSamlIdentityProviderUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FederationSamlIdentityProviderUpdate) SetStatus(v string) {
	o.Status = &v
}
