// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// X509CertificateUpdate struct for X509CertificateUpdate
type X509CertificateUpdate struct {
	// Certificate content.
	Content *string `json:"content,omitempty"`
	// Latest date that the certificate is valid. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	NotAfter *time.Time `json:"notAfter,omitempty"`
	// Earliest date that the certificate is valid. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	NotBefore *time.Time `json:"notBefore,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *X509CertificateUpdate) MarshalJSON() ([]byte, error) {
	type noMethod X509CertificateUpdate
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewX509CertificateUpdate instantiates a new X509CertificateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509CertificateUpdate() *X509CertificateUpdate {
	this := X509CertificateUpdate{}
	return &this
}

// NewX509CertificateUpdateWithDefaults instantiates a new X509CertificateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateUpdateWithDefaults() *X509CertificateUpdate {
	this := X509CertificateUpdate{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise
func (o *X509CertificateUpdate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateUpdate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}

	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *X509CertificateUpdate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *X509CertificateUpdate) SetContent(v string) {
	o.Content = &v
}

// SetContentNil sets Content to an explicit JSON null when marshaled.
func (o *X509CertificateUpdate) SetContentNil() {
	o.Content = nil
	o.NullFields = append(o.NullFields, "Content")
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise
func (o *X509CertificateUpdate) GetNotAfter() time.Time {
	if o == nil || IsNil(o.NotAfter) {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateUpdate) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotAfter) {
		return nil, false
	}

	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *X509CertificateUpdate) HasNotAfter() bool {
	if o != nil && !IsNil(o.NotAfter) {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *X509CertificateUpdate) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// SetNotAfterNil sets NotAfter to an explicit JSON null when marshaled.
func (o *X509CertificateUpdate) SetNotAfterNil() {
	o.NotAfter = nil
	o.NullFields = append(o.NullFields, "NotAfter")
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise
func (o *X509CertificateUpdate) GetNotBefore() time.Time {
	if o == nil || IsNil(o.NotBefore) {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *X509CertificateUpdate) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NotBefore) {
		return nil, false
	}

	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *X509CertificateUpdate) HasNotBefore() bool {
	if o != nil && !IsNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *X509CertificateUpdate) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// SetNotBeforeNil sets NotBefore to an explicit JSON null when marshaled.
func (o *X509CertificateUpdate) SetNotBeforeNil() {
	o.NotBefore = nil
	o.NullFields = append(o.NullFields, "NotBefore")
}
