# MetricIntegrationUpdateRequest

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

### NewMetricIntegrationUpdateRequest

`func NewMetricIntegrationUpdateRequest(aggregationTemporality string, authType string, endpoint string, integrationType string, metricSelection []string, providerType string, ) *MetricIntegrationUpdateRequest`

NewMetricIntegrationUpdateRequest instantiates a new MetricIntegrationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricIntegrationUpdateRequestWithDefaults

`func NewMetricIntegrationUpdateRequestWithDefaults() *MetricIntegrationUpdateRequest`

NewMetricIntegrationUpdateRequestWithDefaults instantiates a new MetricIntegrationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationTemporality

`func (o *MetricIntegrationUpdateRequest) GetAggregationTemporality() string`

GetAggregationTemporality returns the AggregationTemporality field if non-nil, zero value otherwise.

### GetAggregationTemporalityOk

`func (o *MetricIntegrationUpdateRequest) GetAggregationTemporalityOk() (*string, bool)`

GetAggregationTemporalityOk returns a tuple with the AggregationTemporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTemporality

`func (o *MetricIntegrationUpdateRequest) SetAggregationTemporality(v string)`

SetAggregationTemporality sets AggregationTemporality field to given value.

### GetAuthType

`func (o *MetricIntegrationUpdateRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MetricIntegrationUpdateRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MetricIntegrationUpdateRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### GetEndpoint

`func (o *MetricIntegrationUpdateRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetricIntegrationUpdateRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetricIntegrationUpdateRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### GetHeaders

`func (o *MetricIntegrationUpdateRequest) GetHeaders() []Header`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *MetricIntegrationUpdateRequest) GetHeadersOk() (*[]Header, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *MetricIntegrationUpdateRequest) SetHeaders(v []Header)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *MetricIntegrationUpdateRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *MetricIntegrationUpdateRequest) SetHeadersNil()`

SetHeadersNil sets Headers to an explicit JSON null when marshaled, overriding any value previously set with SetHeaders. Calling SetHeaders again clears the null override.

### GetIntegrationType

`func (o *MetricIntegrationUpdateRequest) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *MetricIntegrationUpdateRequest) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *MetricIntegrationUpdateRequest) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### GetMetricSelection

`func (o *MetricIntegrationUpdateRequest) GetMetricSelection() []string`

GetMetricSelection returns the MetricSelection field if non-nil, zero value otherwise.

### GetMetricSelectionOk

`func (o *MetricIntegrationUpdateRequest) GetMetricSelectionOk() (*[]string, bool)`

GetMetricSelectionOk returns a tuple with the MetricSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelection

`func (o *MetricIntegrationUpdateRequest) SetMetricSelection(v []string)`

SetMetricSelection sets MetricSelection field to given value.

### GetProviderType

`func (o *MetricIntegrationUpdateRequest) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *MetricIntegrationUpdateRequest) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *MetricIntegrationUpdateRequest) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


