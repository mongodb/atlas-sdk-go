// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// OAuthSigningKeyInfo Read-only metadata for the Atlas-managed signing key used by `PRIVATE_KEY_JWT`. Present only for that method. Register the `jwksUri` with your identity provider. Atlas rotates the underlying key without changing this URL.
type OAuthSigningKeyInfo struct {
	// Signing algorithm of the Atlas-managed key.
	// Read only field.
	Algorithm *string `json:"algorithm,omitempty"`
	// When the currently active signing key was created. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	// Read only field.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Public JWKS URL serving this integration's signing keys. Fixed for the lifetime of the integration.
	// Read only field.
	JwksUri *string `json:"jwksUri,omitempty"`
	// Key ID stamped on client assertions, the `SHA-1` thumbprint of the key certificate in uppercase hexadecimal. Changes when Atlas rotates the key.
	// Read only field.
	Kid *string `json:"kid,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *OAuthSigningKeyInfo) MarshalJSON() ([]byte, error) {
	type noMethod OAuthSigningKeyInfo
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewOAuthSigningKeyInfo instantiates a new OAuthSigningKeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthSigningKeyInfo() *OAuthSigningKeyInfo {
	this := OAuthSigningKeyInfo{}
	return &this
}

// NewOAuthSigningKeyInfoWithDefaults instantiates a new OAuthSigningKeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthSigningKeyInfoWithDefaults() *OAuthSigningKeyInfo {
	this := OAuthSigningKeyInfo{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise
func (o *OAuthSigningKeyInfo) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSigningKeyInfo) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}

	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *OAuthSigningKeyInfo) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *OAuthSigningKeyInfo) SetAlgorithm(v string) {
	o.Algorithm = &v
	o.NullFields = removeNullField(o.NullFields, "Algorithm")
}

// SetAlgorithmNil sets Algorithm to an explicit JSON null when marshaled.
func (o *OAuthSigningKeyInfo) SetAlgorithmNil() {
	o.Algorithm = nil
	o.NullFields = addNullField(o.NullFields, "Algorithm")
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *OAuthSigningKeyInfo) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSigningKeyInfo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OAuthSigningKeyInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OAuthSigningKeyInfo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
	o.NullFields = removeNullField(o.NullFields, "CreatedAt")
}

// SetCreatedAtNil sets CreatedAt to an explicit JSON null when marshaled.
func (o *OAuthSigningKeyInfo) SetCreatedAtNil() {
	o.CreatedAt = nil
	o.NullFields = addNullField(o.NullFields, "CreatedAt")
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise
func (o *OAuthSigningKeyInfo) GetJwksUri() string {
	if o == nil || IsNil(o.JwksUri) {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSigningKeyInfo) GetJwksUriOk() (*string, bool) {
	if o == nil || IsNil(o.JwksUri) {
		return nil, false
	}

	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *OAuthSigningKeyInfo) HasJwksUri() bool {
	if o != nil && !IsNil(o.JwksUri) {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *OAuthSigningKeyInfo) SetJwksUri(v string) {
	o.JwksUri = &v
	o.NullFields = removeNullField(o.NullFields, "JwksUri")
}

// SetJwksUriNil sets JwksUri to an explicit JSON null when marshaled.
func (o *OAuthSigningKeyInfo) SetJwksUriNil() {
	o.JwksUri = nil
	o.NullFields = addNullField(o.NullFields, "JwksUri")
}

// GetKid returns the Kid field value if set, zero value otherwise
func (o *OAuthSigningKeyInfo) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthSigningKeyInfo) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}

	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *OAuthSigningKeyInfo) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *OAuthSigningKeyInfo) SetKid(v string) {
	o.Kid = &v
	o.NullFields = removeNullField(o.NullFields, "Kid")
}

// SetKidNil sets Kid to an explicit JSON null when marshaled.
func (o *OAuthSigningKeyInfo) SetKidNil() {
	o.Kid = nil
	o.NullFields = addNullField(o.NullFields, "Kid")
}
