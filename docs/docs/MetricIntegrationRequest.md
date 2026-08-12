# MetricIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationTemporality** | **string** | The temporality to send to the metric integration. | 
**AuthType** | **string** | Authentication method the integration uses when exporting metrics to the endpoint. &#x60;HEADER&#x60; authenticates with the static HTTP headers provided in the &#x60;headers&#x60; field, which must be set when this value is used. | 
**Endpoint** | **string** | OpenTelemetry collector endpoint URL. Must use HTTPS. | 
**Headers** | Pointer to [**[]Header**](Header.md) | HTTP headers for authentication and configuration. Total size limit 2KB. Required when &#x60;authType&#x60; is &#x60;HEADER&#x60;. | [optional] 
**IntegrationType** | **string** | Type of metric integration. Identifies which protocol will be used for the integration. This value cannot be modified after the integration is created. | 
**MetricSelection** | **[]string** | Array of metric categories to export. Determines which types of metrics are sent to the integration. | 
**ProviderType** | **string** | The provider type for the metric integration. Identifies the third-party service provider. | 

## Methods

### NewMetricIntegrationRequest

`func NewMetricIntegrationRequest(aggregationTemporality string, authType string, endpoint string, integrationType string, metricSelection []string, providerType string, ) *MetricIntegrationRequest`

NewMetricIntegrationRequest instantiates a new MetricIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricIntegrationRequestWithDefaults

`func NewMetricIntegrationRequestWithDefaults() *MetricIntegrationRequest`

NewMetricIntegrationRequestWithDefaults instantiates a new MetricIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationTemporality

`func (o *MetricIntegrationRequest) GetAggregationTemporality() string`

GetAggregationTemporality returns the AggregationTemporality field if non-nil, zero value otherwise.

### GetAggregationTemporalityOk

`func (o *MetricIntegrationRequest) GetAggregationTemporalityOk() (*string, bool)`

GetAggregationTemporalityOk returns a tuple with the AggregationTemporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTemporality

`func (o *MetricIntegrationRequest) SetAggregationTemporality(v string)`

SetAggregationTemporality sets AggregationTemporality field to given value.

### GetAuthType

`func (o *MetricIntegrationRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MetricIntegrationRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MetricIntegrationRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### GetEndpoint

`func (o *MetricIntegrationRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetricIntegrationRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetricIntegrationRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### GetHeaders

`func (o *MetricIntegrationRequest) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MetricIntegrationRequest) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MetricIntegrationRequest) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MetricIntegrationRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *MetricIntegrationRequest) SetHeadersNil()`

SetHeadersNil sets Headers to an explicit JSON null when marshaled, overriding any value previously set with SetHeaders. Calling SetHeaders again clears the null override.

### GetIntegrationType

`func (o *MetricIntegrationRequest) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *MetricIntegrationRequest) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *MetricIntegrationRequest) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### GetMetricSelection

`func (o *MetricIntegrationRequest) GetMetricSelection() []string`

GetMetricSelection returns the MetricSelection field if non-nil, zero value otherwise.

### GetMetricSelectionOk

`func (o *MetricIntegrationRequest) GetMetricSelectionOk() (*[]string, bool)`

GetMetricSelectionOk returns a tuple with the MetricSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelection

`func (o *MetricIntegrationRequest) SetMetricSelection(v []string)`

SetMetricSelection sets MetricSelection field to given value.

### GetProviderType

`func (o *MetricIntegrationRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *MetricIntegrationRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *MetricIntegrationRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


