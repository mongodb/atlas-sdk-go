// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MetricIntegrationUpdateRequest Request schema for updating a metric integration.
type MetricIntegrationUpdateRequest struct {
	// The temporality to send to the metric integration.
	AggregationTemporality string `json:"aggregationTemporality"`
	// OpenTelemetry collector endpoint URL. Must use HTTPS.
	Endpoint string `json:"endpoint"`
	// HTTP headers for authentication and configuration. Total size limit 2KB.
	Headers []Header `json:"headers"`
	// Type of metric integration. Identifies which protocol will be used for the integration. This value cannot be modified after the integration is created.
	IntegrationType string `json:"integrationType"`
	// Array of metric categories to export. Determines which types of metrics are sent to the integration.
	MetricSelection []string `json:"metricSelection"`
	// The provider type for the metric integration. Identifies the third-party service provider.
	ProviderType string `json:"providerType"`
}

// NewMetricIntegrationUpdateRequest instantiates a new MetricIntegrationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricIntegrationUpdateRequest(aggregationTemporality string, endpoint string, headers []Header, integrationType string, metricSelection []string, providerType string) *MetricIntegrationUpdateRequest {
	this := MetricIntegrationUpdateRequest{}
	this.AggregationTemporality = aggregationTemporality
	this.Endpoint = endpoint
	this.Headers = headers
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

// GetHeaders returns the Headers field value
func (o *MetricIntegrationUpdateRequest) GetHeaders() []Header {
	if o == nil {
		var ret []Header
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationUpdateRequest) GetHeadersOk() (*[]Header, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *MetricIntegrationUpdateRequest) SetHeaders(v []Header) {
	o.Headers = v
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
