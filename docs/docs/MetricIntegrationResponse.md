# MetricIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationTemporality** | **string** | The temporality to send to the metric integration. | 
**AuthType** | **string** | Authentication method the integration uses when exporting metrics to the endpoint. | 
**Endpoint** | **string** | OpenTelemetry collector endpoint URL. | 
**HeadersRedacted** | Pointer to [**[]RedactedHeader**](RedactedHeader.md) | HTTP headers for authentication and configuration. Values are redacted and never returned in plaintext. | [optional] [readonly] 
**IntegrationType** | **string** | Type of metric integration. Identifies which protocol will be used for the integration. | 
**MetricIntegrationId** | **string** | Unique identifier of the metric integration configuration. | [readonly] 
**MetricSelection** | **[]string** | Array of metric categories to export. Determines which types of metrics are sent to the integration. | 
**ProviderType** | **string** | The provider type for the metric integration. Identifies the third-party service provider. | 

## Methods

### NewMetricIntegrationResponse

`func NewMetricIntegrationResponse(aggregationTemporality string, authType string, endpoint string, integrationType string, metricIntegrationId string, metricSelection []string, providerType string, ) *MetricIntegrationResponse`

NewMetricIntegrationResponse instantiates a new MetricIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricIntegrationResponseWithDefaults

`func NewMetricIntegrationResponseWithDefaults() *MetricIntegrationResponse`

NewMetricIntegrationResponseWithDefaults instantiates a new MetricIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationTemporality

`func (o *MetricIntegrationResponse) GetAggregationTemporality() string`

GetAggregationTemporality returns the AggregationTemporality field if non-nil, zero value otherwise.

### GetAggregationTemporalityOk

`func (o *MetricIntegrationResponse) GetAggregationTemporalityOk() (*string, bool)`

GetAggregationTemporalityOk returns a tuple with the AggregationTemporality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTemporality

`func (o *MetricIntegrationResponse) SetAggregationTemporality(v string)`

SetAggregationTemporality sets AggregationTemporality field to given value.

### GetAuthType

`func (o *MetricIntegrationResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *MetricIntegrationResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *MetricIntegrationResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### GetEndpoint

`func (o *MetricIntegrationResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetricIntegrationResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetricIntegrationResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### GetHeadersRedacted

`func (o *MetricIntegrationResponse) GetHeadersRedacted() []RedactedHeader`

GetHeadersRedacted returns the HeadersRedacted field if non-nil, zero value otherwise.

### GetHeadersRedactedOk

`func (o *MetricIntegrationResponse) GetHeadersRedactedOk() (*[]RedactedHeader, bool)`

GetHeadersRedactedOk returns a tuple with the HeadersRedacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadersRedacted

`func (o *MetricIntegrationResponse) SetHeadersRedacted(v []RedactedHeader)`

SetHeadersRedacted sets HeadersRedacted field to given value.

### HasHeadersRedacted

`func (o *MetricIntegrationResponse) HasHeadersRedacted() bool`

HasHeadersRedacted returns a boolean if a field has been set.

### SetHeadersRedactedNil

`func (o *MetricIntegrationResponse) SetHeadersRedactedNil()`

SetHeadersRedactedNil sets HeadersRedacted to an explicit JSON null when marshaled, overriding any value previously set with SetHeadersRedacted. Calling SetHeadersRedacted again clears the null override.

### GetIntegrationType

`func (o *MetricIntegrationResponse) GetIntegrationType() string`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *MetricIntegrationResponse) GetIntegrationTypeOk() (*string, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *MetricIntegrationResponse) SetIntegrationType(v string)`

SetIntegrationType sets IntegrationType field to given value.

### GetMetricIntegrationId

`func (o *MetricIntegrationResponse) GetMetricIntegrationId() string`

GetMetricIntegrationId returns the MetricIntegrationId field if non-nil, zero value otherwise.

### GetMetricIntegrationIdOk

`func (o *MetricIntegrationResponse) GetMetricIntegrationIdOk() (*string, bool)`

GetMetricIntegrationIdOk returns a tuple with the MetricIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricIntegrationId

`func (o *MetricIntegrationResponse) SetMetricIntegrationId(v string)`

SetMetricIntegrationId sets MetricIntegrationId field to given value.

### GetMetricSelection

`func (o *MetricIntegrationResponse) GetMetricSelection() []string`

GetMetricSelection returns the MetricSelection field if non-nil, zero value otherwise.

### GetMetricSelectionOk

`func (o *MetricIntegrationResponse) GetMetricSelectionOk() (*[]string, bool)`

GetMetricSelectionOk returns a tuple with the MetricSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelection

`func (o *MetricIntegrationResponse) SetMetricSelection(v []string)`

SetMetricSelection sets MetricSelection field to given value.

### GetProviderType

`func (o *MetricIntegrationResponse) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *MetricIntegrationResponse) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *MetricIntegrationResponse) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


