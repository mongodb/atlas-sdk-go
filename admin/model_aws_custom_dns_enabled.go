// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AWSCustomDNSEnabled struct for AWSCustomDNSEnabled
type AWSCustomDNSEnabled struct {
	// Flag that indicates whether the project's clusters deployed to Amazon Web Services (AWS) use a custom Domain Name System (DNS). When `\"enabled\": true`, connect to your cluster using Private IP for Peering connection strings.
	Enabled bool `json:"enabled"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AWSCustomDNSEnabled) MarshalJSON() ([]byte, error) {
	type noMethod AWSCustomDNSEnabled
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAWSCustomDNSEnabled instantiates a new AWSCustomDNSEnabled object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSCustomDNSEnabled(enabled bool) *AWSCustomDNSEnabled {
	this := AWSCustomDNSEnabled{}
	this.Enabled = enabled
	return &this
}

// NewAWSCustomDNSEnabledWithDefaults instantiates a new AWSCustomDNSEnabled object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSCustomDNSEnabledWithDefaults() *AWSCustomDNSEnabled {
	this := AWSCustomDNSEnabled{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *AWSCustomDNSEnabled) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AWSCustomDNSEnabled) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AWSCustomDNSEnabled) SetEnabled(v bool) {
	o.Enabled = v
}
