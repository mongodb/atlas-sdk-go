// Code based on the AtlasAPI V2 OpenAPI file

package admin

// PemFileInfoUpdate PEM file information for the identity provider's current certificates.
type PemFileInfoUpdate struct {
	// List of certificates in the file.
	Certificates *[]X509CertificateUpdate `json:"certificates,omitempty"`
	// Human-readable label given to the file.
	FileName *string `json:"fileName,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *PemFileInfoUpdate) MarshalJSON() ([]byte, error) {
	type noMethod PemFileInfoUpdate
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewPemFileInfoUpdate instantiates a new PemFileInfoUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPemFileInfoUpdate() *PemFileInfoUpdate {
	this := PemFileInfoUpdate{}
	return &this
}

// NewPemFileInfoUpdateWithDefaults instantiates a new PemFileInfoUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPemFileInfoUpdateWithDefaults() *PemFileInfoUpdate {
	this := PemFileInfoUpdate{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise
func (o *PemFileInfoUpdate) GetCertificates() []X509CertificateUpdate {
	if o == nil || IsNil(o.Certificates) {
		var ret []X509CertificateUpdate
		return ret
	}
	return *o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PemFileInfoUpdate) GetCertificatesOk() (*[]X509CertificateUpdate, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}

	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *PemFileInfoUpdate) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []X509CertificateUpdate and assigns it to the Certificates field.
func (o *PemFileInfoUpdate) SetCertificates(v []X509CertificateUpdate) {
	o.Certificates = &v
	o.NullFields = removeNullField(o.NullFields, "Certificates")
}

// SetCertificatesNil sets Certificates to an explicit JSON null when marshaled.
func (o *PemFileInfoUpdate) SetCertificatesNil() {
	o.Certificates = nil
	o.NullFields = addNullField(o.NullFields, "Certificates")
}

// GetFileName returns the FileName field value if set, zero value otherwise
func (o *PemFileInfoUpdate) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PemFileInfoUpdate) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}

	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *PemFileInfoUpdate) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *PemFileInfoUpdate) SetFileName(v string) {
	o.FileName = &v
	o.NullFields = removeNullField(o.NullFields, "FileName")
}

// SetFileNameNil sets FileName to an explicit JSON null when marshaled.
func (o *PemFileInfoUpdate) SetFileNameNil() {
	o.FileName = nil
	o.NullFields = addNullField(o.NullFields, "FileName")
}
