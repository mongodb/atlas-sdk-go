// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MetricIntegrationResponse Response schema for metric integration operations.
type MetricIntegrationResponse struct {
	// The temporality to send to the metric integration.
	AggregationTemporality string `json:"aggregationTemporality"`
	// OpenTelemetry collector endpoint URL.
	Endpoint string `json:"endpoint"`
	// HTTP headers for authentication and configuration. Header values are redacted in responses.
	Headers []Header `json:"headers"`
	// Type of metric integration. Identifies which protocol will be used for the integration.
	IntegrationType string `json:"integrationType"`
	// Unique hexadecimal digit string that identifies the metric integration configuration.
	// Read only field.
	MetricIntegrationId string `json:"metricIntegrationId"`
	// Array of metric categories to export. Determines which types of metrics are sent to the integration.
	MetricSelection []string `json:"metricSelection"`
	// The provider type for the metric integration. Identifies the third-party service provider.
	ProviderType string `json:"providerType"`
}

// NewMetricIntegrationResponse instantiates a new MetricIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricIntegrationResponse(aggregationTemporality string, endpoint string, headers []Header, integrationType string, metricIntegrationId string, metricSelection []string, providerType string) *MetricIntegrationResponse {
	this := MetricIntegrationResponse{}
	this.AggregationTemporality = aggregationTemporality
	this.Endpoint = endpoint
	this.Headers = headers
	this.IntegrationType = integrationType
	this.MetricIntegrationId = metricIntegrationId
	this.MetricSelection = metricSelection
	this.ProviderType = providerType
	return &this
}

// NewMetricIntegrationResponseWithDefaults instantiates a new MetricIntegrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricIntegrationResponseWithDefaults() *MetricIntegrationResponse {
	this := MetricIntegrationResponse{}
	return &this
}

// GetAggregationTemporality returns the AggregationTemporality field value
func (o *MetricIntegrationResponse) GetAggregationTemporality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AggregationTemporality
}

// GetAggregationTemporalityOk returns a tuple with the AggregationTemporality field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetAggregationTemporalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationTemporality, true
}

// SetAggregationTemporality sets field value
func (o *MetricIntegrationResponse) SetAggregationTemporality(v string) {
	o.AggregationTemporality = v
}

// GetEndpoint returns the Endpoint field value
func (o *MetricIntegrationResponse) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *MetricIntegrationResponse) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetHeaders returns the Headers field value
func (o *MetricIntegrationResponse) GetHeaders() []Header {
	if o == nil {
		var ret []Header
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetHeadersOk() (*[]Header, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *MetricIntegrationResponse) SetHeaders(v []Header) {
	o.Headers = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *MetricIntegrationResponse) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value
func (o *MetricIntegrationResponse) SetIntegrationType(v string) {
	o.IntegrationType = v
}

// GetMetricIntegrationId returns the MetricIntegrationId field value
func (o *MetricIntegrationResponse) GetMetricIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricIntegrationId
}

// GetMetricIntegrationIdOk returns a tuple with the MetricIntegrationId field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetMetricIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricIntegrationId, true
}

// SetMetricIntegrationId sets field value
func (o *MetricIntegrationResponse) SetMetricIntegrationId(v string) {
	o.MetricIntegrationId = v
}

// GetMetricSelection returns the MetricSelection field value
func (o *MetricIntegrationResponse) GetMetricSelection() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MetricSelection
}

// GetMetricSelectionOk returns a tuple with the MetricSelection field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetMetricSelectionOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricSelection, true
}

// SetMetricSelection sets field value
func (o *MetricIntegrationResponse) SetMetricSelection(v []string) {
	o.MetricSelection = v
}

// GetProviderType returns the ProviderType field value
func (o *MetricIntegrationResponse) GetProviderType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetProviderTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderType, true
}

// SetProviderType sets field value
func (o *MetricIntegrationResponse) SetProviderType(v string) {
	o.ProviderType = v
}
