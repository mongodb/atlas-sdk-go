// Code based on the AtlasAPI V2 OpenAPI file

package admin

// MetricIntegrationResponse Response schema for metric integration operations.
type MetricIntegrationResponse struct {
	// The temporality to send to the metric integration.
	AggregationTemporality string `json:"aggregationTemporality"`
	// Authentication method the integration uses when exporting metrics to the endpoint.
	AuthType string `json:"authType"`
	// OpenTelemetry collector endpoint URL.
	Endpoint string `json:"endpoint"`
	// HTTP headers for authentication and configuration. Values are redacted and never returned in plaintext.
	// Read only field.
	HeadersRedacted *[]RedactedHeader `json:"headersRedacted,omitempty"`
	// Type of metric integration. Identifies which protocol will be used for the integration.
	IntegrationType string `json:"integrationType"`
	// Unique identifier of the metric integration configuration.
	// Read only field.
	MetricIntegrationId string `json:"metricIntegrationId"`
	// Array of metric categories to export. Determines which types of metrics are sent to the integration.
	MetricSelection []string `json:"metricSelection"`
	// The provider type for the metric integration. Identifies the third-party service provider.
	ProviderType string `json:"providerType"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *MetricIntegrationResponse) MarshalJSON() ([]byte, error) {
	type noMethod MetricIntegrationResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewMetricIntegrationResponse instantiates a new MetricIntegrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricIntegrationResponse(aggregationTemporality string, authType string, endpoint string, integrationType string, metricIntegrationId string, metricSelection []string, providerType string) *MetricIntegrationResponse {
	this := MetricIntegrationResponse{}
	this.AggregationTemporality = aggregationTemporality
	this.AuthType = authType
	this.Endpoint = endpoint
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

// GetAuthType returns the AuthType field value
func (o *MetricIntegrationResponse) GetAuthType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetAuthTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *MetricIntegrationResponse) SetAuthType(v string) {
	o.AuthType = v
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

// GetHeadersRedacted returns the HeadersRedacted field value if set, zero value otherwise
func (o *MetricIntegrationResponse) GetHeadersRedacted() []RedactedHeader {
	if o == nil || IsNil(o.HeadersRedacted) {
		var ret []RedactedHeader
		return ret
	}
	return *o.HeadersRedacted
}

// GetHeadersRedactedOk returns a tuple with the HeadersRedacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIntegrationResponse) GetHeadersRedactedOk() (*[]RedactedHeader, bool) {
	if o == nil || IsNil(o.HeadersRedacted) {
		return nil, false
	}

	return o.HeadersRedacted, true
}

// HasHeadersRedacted returns a boolean if a field has been set.
func (o *MetricIntegrationResponse) HasHeadersRedacted() bool {
	if o != nil && !IsNil(o.HeadersRedacted) {
		return true
	}

	return false
}

// SetHeadersRedacted gets a reference to the given []RedactedHeader and assigns it to the HeadersRedacted field.
func (o *MetricIntegrationResponse) SetHeadersRedacted(v []RedactedHeader) {
	o.HeadersRedacted = &v
	o.NullFields = removeNullField(o.NullFields, "HeadersRedacted")
}

// SetHeadersRedactedNil sets HeadersRedacted to an explicit JSON null when marshaled.
func (o *MetricIntegrationResponse) SetHeadersRedactedNil() {
	o.HeadersRedacted = nil
	o.NullFields = addNullField(o.NullFields, "HeadersRedacted")
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
