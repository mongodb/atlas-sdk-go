// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AutoEmbeddingInferenceScopeUpdateRequestConfigured Cloud and geographic dimensions to configure atomically.
type AutoEmbeddingInferenceScopeUpdateRequestConfigured struct {
	// Cloud provider scope.
	Cloud string `json:"cloud"`
	// Geographic scope.
	Geography string `json:"geography"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) MarshalJSON() ([]byte, error) {
	type noMethod AutoEmbeddingInferenceScopeUpdateRequestConfigured
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAutoEmbeddingInferenceScopeUpdateRequestConfigured instantiates a new AutoEmbeddingInferenceScopeUpdateRequestConfigured object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoEmbeddingInferenceScopeUpdateRequestConfigured(cloud string, geography string) *AutoEmbeddingInferenceScopeUpdateRequestConfigured {
	this := AutoEmbeddingInferenceScopeUpdateRequestConfigured{}
	this.Cloud = cloud
	this.Geography = geography
	return &this
}

// NewAutoEmbeddingInferenceScopeUpdateRequestConfiguredWithDefaults instantiates a new AutoEmbeddingInferenceScopeUpdateRequestConfigured object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoEmbeddingInferenceScopeUpdateRequestConfiguredWithDefaults() *AutoEmbeddingInferenceScopeUpdateRequestConfigured {
	this := AutoEmbeddingInferenceScopeUpdateRequestConfigured{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) SetCloud(v string) {
	o.Cloud = v
}

// GetGeography returns the Geography field value
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) GetGeography() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Geography
}

// GetGeographyOk returns a tuple with the Geography field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) GetGeographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Geography, true
}

// SetGeography sets field value
func (o *AutoEmbeddingInferenceScopeUpdateRequestConfigured) SetGeography(v string) {
	o.Geography = v
}
