// Code based on the AtlasAPI V2 OpenAPI file

package admin

// QuerySamplingCreateRequest Namespace on which to start query sampling.
type QuerySamplingCreateRequest struct {
	// Human-readable label that identifies the namespace (`<database>.<collection>`) on which to sample queries.
	Namespace string `json:"namespace"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *QuerySamplingCreateRequest) MarshalJSON() ([]byte, error) {
	type noMethod QuerySamplingCreateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewQuerySamplingCreateRequest instantiates a new QuerySamplingCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuerySamplingCreateRequest(namespace string) *QuerySamplingCreateRequest {
	this := QuerySamplingCreateRequest{}
	this.Namespace = namespace
	return &this
}

// NewQuerySamplingCreateRequestWithDefaults instantiates a new QuerySamplingCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuerySamplingCreateRequestWithDefaults() *QuerySamplingCreateRequest {
	this := QuerySamplingCreateRequest{}
	return &this
}

// GetNamespace returns the Namespace field value
func (o *QuerySamplingCreateRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *QuerySamplingCreateRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *QuerySamplingCreateRequest) SetNamespace(v string) {
	o.Namespace = v
}
