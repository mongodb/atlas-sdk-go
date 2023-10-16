// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// X509Certificate struct for X509Certificate
type X509Certificate struct {
	 
	NotAfter  *time.Time `json:"notAfter,omitempty"`
	NotBefore *time.Time `json:"notBefore,omitempty"`
}

// NewX509Certificate instantiates a new X509Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewX509Certificate() *X509Certificate {
	this := X509Certificate{}
	return &this
}

// NewX509CertificateWithDefaults instantiates a new X509Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewX509CertificateWithDefaults() *X509Certificate {
	this := X509Certificate{}
	return &this
}
