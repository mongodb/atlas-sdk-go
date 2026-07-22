// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsGCPConnectionConfig GCP-specific configuration for the connection.
type StreamsGCPConnectionConfig struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Email address of the Google Cloud Platform (GCP) service account that Atlas Streams uses to connect to the GCP Pub/Sub resources.
	ServiceAccountId *string `json:"serviceAccountId,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsGCPConnectionConfig) MarshalJSON() ([]byte, error) {
	type noMethod StreamsGCPConnectionConfig
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsGCPConnectionConfig instantiates a new StreamsGCPConnectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsGCPConnectionConfig() *StreamsGCPConnectionConfig {
	this := StreamsGCPConnectionConfig{}
	return &this
}

// NewStreamsGCPConnectionConfigWithDefaults instantiates a new StreamsGCPConnectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsGCPConnectionConfigWithDefaults() *StreamsGCPConnectionConfig {
	this := StreamsGCPConnectionConfig{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsGCPConnectionConfig) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsGCPConnectionConfig) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsGCPConnectionConfig) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsGCPConnectionConfig) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsGCPConnectionConfig) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetServiceAccountId returns the ServiceAccountId field value if set, zero value otherwise
func (o *StreamsGCPConnectionConfig) GetServiceAccountId() string {
	if o == nil || IsNil(o.ServiceAccountId) {
		var ret string
		return ret
	}
	return *o.ServiceAccountId
}

// GetServiceAccountIdOk returns a tuple with the ServiceAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsGCPConnectionConfig) GetServiceAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountId) {
		return nil, false
	}

	return o.ServiceAccountId, true
}

// HasServiceAccountId returns a boolean if a field has been set.
func (o *StreamsGCPConnectionConfig) HasServiceAccountId() bool {
	if o != nil && !IsNil(o.ServiceAccountId) {
		return true
	}

	return false
}

// SetServiceAccountId gets a reference to the given string and assigns it to the ServiceAccountId field.
func (o *StreamsGCPConnectionConfig) SetServiceAccountId(v string) {
	o.ServiceAccountId = &v
	o.NullFields = removeNullField(o.NullFields, "ServiceAccountId")
}

// SetServiceAccountIdNil sets ServiceAccountId to an explicit JSON null when marshaled.
func (o *StreamsGCPConnectionConfig) SetServiceAccountIdNil() {
	o.ServiceAccountId = nil
	o.NullFields = addNullField(o.NullFields, "ServiceAccountId")
}
