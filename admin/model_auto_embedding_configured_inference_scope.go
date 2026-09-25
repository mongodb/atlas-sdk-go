// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AutoEmbeddingConfiguredInferenceScope Configured cloud and geographic dimensions. Both values are AUTO when Atlas selects the scope.
type AutoEmbeddingConfiguredInferenceScope struct {
	// Configured cloud provider scope.
	Cloud string `json:"cloud"`
	// Configured geographic scope.
	Geography string `json:"geography"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AutoEmbeddingConfiguredInferenceScope) MarshalJSON() ([]byte, error) {
	type noMethod AutoEmbeddingConfiguredInferenceScope
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAutoEmbeddingConfiguredInferenceScope instantiates a new AutoEmbeddingConfiguredInferenceScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoEmbeddingConfiguredInferenceScope(cloud string, geography string) *AutoEmbeddingConfiguredInferenceScope {
	this := AutoEmbeddingConfiguredInferenceScope{}
	this.Cloud = cloud
	this.Geography = geography
	return &this
}

// NewAutoEmbeddingConfiguredInferenceScopeWithDefaults instantiates a new AutoEmbeddingConfiguredInferenceScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoEmbeddingConfiguredInferenceScopeWithDefaults() *AutoEmbeddingConfiguredInferenceScope {
	this := AutoEmbeddingConfiguredInferenceScope{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AutoEmbeddingConfiguredInferenceScope) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingConfiguredInferenceScope) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AutoEmbeddingConfiguredInferenceScope) SetCloud(v string) {
	o.Cloud = v
}

// GetGeography returns the Geography field value
func (o *AutoEmbeddingConfiguredInferenceScope) GetGeography() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Geography
}

// GetGeographyOk returns a tuple with the Geography field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingConfiguredInferenceScope) GetGeographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Geography, true
}

// SetGeography sets field value
func (o *AutoEmbeddingConfiguredInferenceScope) SetGeography(v string) {
	o.Geography = v
}
