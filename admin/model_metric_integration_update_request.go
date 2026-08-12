// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MetricIntegrationUpdateRequest Request schema for updating a metric integration.
type MetricIntegrationUpdateRequest struct {
	// The temporality to send to the metric integration.
	AggregationTemporality string `json:"aggregationTemporality"`
	// Authentication method the integration uses when exporting metrics to the endpoint. `HEADER` authenticates with the static HTTP headers provided in the `headers` field, which must be set when this value is used.
	AuthType string `json:"authType"`
	// OpenTelemetry collector endpoint URL. Must use HTTPS.
	Endpoint string `json:"endpoint"`
	// HTTP headers for authentication and configuration. Total size limit 2KB. Required when `authType` is `HEADER`.
	// Write only field.
	Headers *[]Header `json:"headers,omitempty"`
	// Type of metric integration. Identifies which protocol will be used for the integration. This value cannot be modified after the integration is created.
	IntegrationType string `json:"integrationType"`
	// Array of metric categories to export. Determines which types of metrics are sent to the integration.
	MetricSelection []string `json:"metricSelection"`
	// The provider type for the metric integration. Identifies the third-party service provider.
	ProviderType string `json:"providerType"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MetricIntegrationUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod MetricIntegrationUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMetricIntegrationUpdateRequest instantiates a new MetricIntegrationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricIntegrationUpdateRequest(aggregationTemporality string, authType string, endpoint string, integrationType string, metricSelection []string, providerType string) *MetricIntegrationUpdateRequest {
	this := MetricIntegrationUpdateRequest{}
	this.AggregationTemporality = aggregationTemporality
	this.AuthType = authType
	this.Endpoint = endpoint
	this.IntegrationType = integrationType
	this.MetricSelection = metricSelection
	this.ProviderType = providerType
	return &this
}

// NewMetricIntegrationUpdateRequestWithDefaults instantiates a new MetricIntegrationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricIntegrationUpdateRequestWithDefaults() *MetricIntegrationUpdateRequest {
	this := MetricIntegrationUpdateRequest{}
	return &this
}

// GetAggregationTemporality returns the AggregationTemporality field value
func (o *MetricIntegrationUpdateRequest) GetAggregationTemporality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AggregationTemporality
}

// GetAggregationTemporalityOk returns a tuple with the AggregationTemporality field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetAggregationTemporalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationTemporality, true
}

// SetAggregationTemporality sets field value
func (o *MetricIntegrationUpdateRequest) SetAggregationTemporality(v string) {
	o.AggregationTemporality = v
}

// GetAuthType returns the AuthType field value
func (o *MetricIntegrationUpdateRequest) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *MetricIntegrationUpdateRequest) SetAuthType(v string) {
	o.AuthType = v
}

// GetEndpoint returns the Endpoint field value
func (o *MetricIntegrationUpdateRequest) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *MetricIntegrationUpdateRequest) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise
func (o *MetricIntegrationUpdateRequest) GetHeaders() []Header {
	if o == nil || IsNil(o.Headers) {
		var ret []Header
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetHeadersOk() (*[]Header, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}

	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *MetricIntegrationUpdateRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []Header and assigns it to the Headers field.
func (o *MetricIntegrationUpdateRequest) SetHeaders(v []Header) {
	o.Headers = &v
	o.NullFields = removeNullField(o.NullFields, "Headers")
}

// SetHeadersNil sets Headers to an explicit JSON null when marshaled.
func (o *MetricIntegrationUpdateRequest) SetHeadersNil() {
	o.Headers = nil
	o.NullFields = addNullField(o.NullFields, "Headers")
}

// GetIntegrationType returns the IntegrationType field value
func (o *MetricIntegrationUpdateRequest) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *MetricIntegrationUpdateRequest) SetIntegrationType(v string) {
	o.IntegrationType = v
}

// GetMetricSelection returns the MetricSelection field value
func (o *MetricIntegrationUpdateRequest) GetMetricSelection() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MetricSelection
}

// GetMetricSelectionOk returns a tuple with the MetricSelection field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetMetricSelectionOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricSelection, true
}

// SetMetricSelection sets field value
func (o *MetricIntegrationUpdateRequest) SetMetricSelection(v []string) {
	o.MetricSelection = v
}

// GetProviderType returns the ProviderType field value
func (o *MetricIntegrationUpdateRequest) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *MetricIntegrationUpdateRequest) SetProviderType(v string) {
	o.ProviderType = v
}
