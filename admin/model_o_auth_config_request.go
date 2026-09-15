// Code based on the AtlasAPI V2 OpenAPI file

package admin

// OAuthConfigRequest OAuth 2.0 client credentials configuration. Required when `authType` is `OAUTH2`. Secrets are never returned.
type OAuthConfigRequest struct {
	// How the client authenticates to the token endpoint. `CLIENT_SECRET` sends a shared secret. `PRIVATE_KEY_JWT` signs a client assertion with an Atlas-generated, Atlas-managed key. Register the returned JWKS URL with your identity provider.
	ClientAuthMethod string `json:"clientAuthMethod"`
	// OAuth 2.0 client identifier registered with the token endpoint.
	ClientId string `json:"clientId"`
	// Shared client secret. Required when `clientAuthMethod` is `CLIENT_SECRET`, and rejected for `PRIVATE_KEY_JWT`. Encrypted at rest and never returned.
	// Write only field.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// Optional OAuth 2.0 scopes requested on the token, sent as a space delimited `scope` parameter. Applies to both client authentication methods.
	Scopes *[]string `json:"scopes,omitempty"`
	// OAuth 2.0 token endpoint URL. Must use HTTPS.
	TokenEndpoint string `json:"tokenEndpoint"`
	// Optional provider-specific parameters added to the token request, for example a resource indicator. Applies to both client authentication methods.
	TokenRequestParams *map[string]string `json:"tokenRequestParams,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OAuthConfigRequest) MarshalJSON() ([]byte, error) {
	type noMethod OAuthConfigRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOAuthConfigRequest instantiates a new OAuthConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthConfigRequest(clientAuthMethod string, clientId string, tokenEndpoint string) *OAuthConfigRequest {
	this := OAuthConfigRequest{}
	this.ClientAuthMethod = clientAuthMethod
	this.ClientId = clientId
	this.TokenEndpoint = tokenEndpoint
	return &this
}

// NewOAuthConfigRequestWithDefaults instantiates a new OAuthConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthConfigRequestWithDefaults() *OAuthConfigRequest {
	this := OAuthConfigRequest{}
	return &this
}

// GetClientAuthMethod returns the ClientAuthMethod field value
func (o *OAuthConfigRequest) GetClientAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientAuthMethod
}

// GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field value
// and a boolean to check if the value has been set.
func (o *OAuthConfigRequest) GetClientAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientAuthMethod, true
}

// SetClientAuthMethod sets field value
func (o *OAuthConfigRequest) SetClientAuthMethod(v string) {
	o.ClientAuthMethod = v
}

// GetClientId returns the ClientId field value
func (o *OAuthConfigRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuthConfigRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuthConfigRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise
func (o *OAuthConfigRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthConfigRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}

	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuthConfigRequest) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *OAuthConfigRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
	o.NullFields = removeNullField(o.NullFields, "ClientSecret")
}

// SetClientSecretNil sets ClientSecret to an explicit JSON null when marshaled.
func (o *OAuthConfigRequest) SetClientSecretNil() {
	o.ClientSecret = nil
	o.NullFields = addNullField(o.NullFields, "ClientSecret")
}

// GetScopes returns the Scopes field value if set, zero value otherwise
func (o *OAuthConfigRequest) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthConfigRequest) GetScopesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}

	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *OAuthConfigRequest) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *OAuthConfigRequest) SetScopes(v []string) {
	o.Scopes = &v
	o.NullFields = removeNullField(o.NullFields, "Scopes")
}

// SetScopesNil sets Scopes to an explicit JSON null when marshaled.
func (o *OAuthConfigRequest) SetScopesNil() {
	o.Scopes = nil
	o.NullFields = addNullField(o.NullFields, "Scopes")
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *OAuthConfigRequest) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *OAuthConfigRequest) GetTokenEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *OAuthConfigRequest) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetTokenRequestParams returns the TokenRequestParams field value if set, zero value otherwise
func (o *OAuthConfigRequest) GetTokenRequestParams() map[string]string {
	if o == nil || IsNil(o.TokenRequestParams) {
		var ret map[string]string
		return ret
	}
	return *o.TokenRequestParams
}

// GetTokenRequestParamsOk returns a tuple with the TokenRequestParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthConfigRequest) GetTokenRequestParamsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TokenRequestParams) {
		return nil, false
	}

	return o.TokenRequestParams, true
}

// HasTokenRequestParams returns a boolean if a field has been set.
func (o *OAuthConfigRequest) HasTokenRequestParams() bool {
	if o != nil && !IsNil(o.TokenRequestParams) {
		return true
	}

	return false
}

// SetTokenRequestParams gets a reference to the given map[string]string and assigns it to the TokenRequestParams field.
func (o *OAuthConfigRequest) SetTokenRequestParams(v map[string]string) {
	o.TokenRequestParams = &v
	o.NullFields = removeNullField(o.NullFields, "TokenRequestParams")
}

// SetTokenRequestParamsNil sets TokenRequestParams to an explicit JSON null when marshaled.
func (o *OAuthConfigRequest) SetTokenRequestParamsNil() {
	o.TokenRequestParams = nil
	o.NullFields = addNullField(o.NullFields, "TokenRequestParams")
}
