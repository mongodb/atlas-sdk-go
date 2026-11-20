// Code based on the AtlasAPI V2 OpenAPI file

package admin

// AutoEmbeddingEffectiveInferenceScope Cloud and geographic dimensions that Atlas applies to this cluster.
type AutoEmbeddingEffectiveInferenceScope struct {
	// Effective cloud provider scope.
	// Read only field.
	Cloud string `json:"cloud"`
	// Effective geographic scope.
	// Read only field.
	Geography string `json:"geography"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *AutoEmbeddingEffectiveInferenceScope) MarshalJSON() ([]byte, error) {
	type noMethod AutoEmbeddingEffectiveInferenceScope
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewAutoEmbeddingEffectiveInferenceScope instantiates a new AutoEmbeddingEffectiveInferenceScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoEmbeddingEffectiveInferenceScope(cloud string, geography string) *AutoEmbeddingEffectiveInferenceScope {
	this := AutoEmbeddingEffectiveInferenceScope{}
	this.Cloud = cloud
	this.Geography = geography
	return &this
}

// NewAutoEmbeddingEffectiveInferenceScopeWithDefaults instantiates a new AutoEmbeddingEffectiveInferenceScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoEmbeddingEffectiveInferenceScopeWithDefaults() *AutoEmbeddingEffectiveInferenceScope {
	this := AutoEmbeddingEffectiveInferenceScope{}
	return &this
}

// GetCloud returns the Cloud field value
func (o *AutoEmbeddingEffectiveInferenceScope) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingEffectiveInferenceScope) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value
func (o *AutoEmbeddingEffectiveInferenceScope) SetCloud(v string) {
	o.Cloud = v
}

// GetGeography returns the Geography field value
func (o *AutoEmbeddingEffectiveInferenceScope) GetGeography() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Geography
}

// GetGeographyOk returns a tuple with the Geography field value
// and a boolean to check if the value has been set.
func (o *AutoEmbeddingEffectiveInferenceScope) GetGeographyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Geography, true
}

// SetGeography sets field value
func (o *AutoEmbeddingEffectiveInferenceScope) SetGeography(v string) {
	o.Geography = v
}
