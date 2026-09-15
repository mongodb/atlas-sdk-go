// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OAuthConfigResponse OAuth 2.0 configuration returned for a metric integration. Secrets are never returned.
type OAuthConfigResponse struct {
	// How the client authenticates to the token endpoint.
	ClientAuthMethod string `json:"clientAuthMethod"`
	// OAuth 2.0 client identifier registered with the token endpoint.
	ClientId string `json:"clientId"`
	// OAuth 2.0 scopes requested on the token.
	Scopes         *[]string            `json:"scopes,omitempty"`
	SigningKeyInfo *OAuthSigningKeyInfo `json:"signingKeyInfo,omitempty"`
	// OAuth 2.0 token endpoint URL.
	TokenEndpoint string `json:"tokenEndpoint"`
	// Provider-specific parameters added to the token request.
	TokenRequestParams *map[string]string `json:"tokenRequestParams,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OAuthConfigResponse) MarshalJSON() ([]byte, error) {
	type noMethod OAuthConfigResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOAuthConfigResponse instantiates a new OAuthConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthConfigResponse(clientAuthMethod string, clientId string, tokenEndpoint string) *OAuthConfigResponse {
	this := OAuthConfigResponse{}
	this.ClientAuthMethod = clientAuthMethod
	this.ClientId = clientId
	this.TokenEndpoint = tokenEndpoint
	return &this
}

// NewOAuthConfigResponseWithDefaults instantiates a new OAuthConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthConfigResponseWithDefaults() *OAuthConfigResponse {
	this := OAuthConfigResponse{}
	return &this
}

// GetClientAuthMethod returns the ClientAuthMethod field value
func (o *OAuthConfigResponse) GetClientAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientAuthMethod
}

// GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field value
// and a boolean to check if the value has been set.
func (o *OAuthConfigResponse) GetClientAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAuthMethod, true
}

// SetClientAuthMethod sets field value
func (o *OAuthConfigResponse) SetClientAuthMethod(v string) {
	o.ClientAuthMethod = v
}

// GetClientId returns the ClientId field value
func (o *OAuthConfigResponse) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuthConfigResponse) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuthConfigResponse) SetClientId(v string) {
	o.ClientId = v
}

// GetScopes returns the Scopes field value if set, zero value otherwise
func (o *OAuthConfigResponse) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthConfigResponse) GetScopesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}

	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuthConfigResponse) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAuthConfigResponse) SetScopes(v []string) {
	o.Scopes = &v
	o.NullFields = removeNullField(o.NullFields, "Scopes")
}

// SetScopesNil sets Scopes to an explicit JSON null when marshaled.
func (o *OAuthConfigResponse) SetScopesNil() {
	o.Scopes = nil
	o.NullFields = addNullField(o.NullFields, "Scopes")
}

// GetSigningKeyInfo returns the SigningKeyInfo field value if set, zero value otherwise
func (o *OAuthConfigResponse) GetSigningKeyInfo() OAuthSigningKeyInfo {
	if o == nil || IsNil(o.SigningKeyInfo) {
		var ret OAuthSigningKeyInfo
		return ret
	}
	return *o.SigningKeyInfo
}

// GetSigningKeyInfoOk returns a tuple with the SigningKeyInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthConfigResponse) GetSigningKeyInfoOk() (*OAuthSigningKeyInfo, bool) {
	if o == nil || IsNil(o.SigningKeyInfo) {
		return nil, false
	}

	return o.SigningKeyInfo, true
}

// HasSigningKeyInfo returns a boolean if a field has been set.
func (o *OAuthConfigResponse) HasSigningKeyInfo() bool {
	if o != nil && !IsNil(o.SigningKeyInfo) {
		return true
	}

	return false
}

// SetSigningKeyInfo gets a reference to the given OAuthSigningKeyInfo and assigns it to the SigningKeyInfo field.
func (o *OAuthConfigResponse) SetSigningKeyInfo(v OAuthSigningKeyInfo) {
	o.SigningKeyInfo = &v
	o.NullFields = removeNullField(o.NullFields, "SigningKeyInfo")
}

// SetSigningKeyInfoNil sets SigningKeyInfo to an explicit JSON null when marshaled.
func (o *OAuthConfigResponse) SetSigningKeyInfoNil() {
	o.SigningKeyInfo = nil
	o.NullFields = addNullField(o.NullFields, "SigningKeyInfo")
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OAuthConfigResponse) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OAuthConfigResponse) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OAuthConfigResponse) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetTokenRequestParams returns the TokenRequestParams field value if set, zero value otherwise
func (o *OAuthConfigResponse) GetTokenRequestParams() map[string]string {
	if o == nil || IsNil(o.TokenRequestParams) {
		var ret map[string]string
		return ret
	}
	return *o.TokenRequestParams
}

// GetTokenRequestParamsOk returns a tuple with the TokenRequestParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthConfigResponse) GetTokenRequestParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TokenRequestParams) {
		return nil, false
	}

	return o.TokenRequestParams, true
}

// HasTokenRequestParams returns a boolean if a field has been set.
func (o *OAuthConfigResponse) HasTokenRequestParams() bool {
	if o != nil && !IsNil(o.TokenRequestParams) {
		return true
	}

	return false
}

// SetTokenRequestParams gets a reference to the given map[string]string and assigns it to the TokenRequestParams field.
func (o *OAuthConfigResponse) SetTokenRequestParams(v map[string]string) {
	o.TokenRequestParams = &v
	o.NullFields = removeNullField(o.NullFields, "TokenRequestParams")
}

// SetTokenRequestParamsNil sets TokenRequestParams to an explicit JSON null when marshaled.
func (o *OAuthConfigResponse) SetTokenRequestParamsNil() {
	o.TokenRequestParams = nil
	o.NullFields = addNullField(o.NullFields, "TokenRequestParams")
}
