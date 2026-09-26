# OrgLogIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogTypes** | **[]string** | Array of log types exported by this integration. | 
**Type** | **string** | Human-readable label that identifies the service to which you want to integrate with Atlas. The value must match the log integration type. This value cannot be modified after the integration is created. | 
**OtelEndpoint** | Pointer to **string** | OpenTelemetry collector endpoint URL. Must be HTTPS and not exceed 2048 characters. | [optional] 
**OtelSuppliedHeaders** | Pointer to [**[]OrgLogIntegrationHeader**](OrgLogIntegrationHeader.md) | HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB. | [optional] 

## Methods

### NewOrgLogIntegrationRequest

`func NewOrgLogIntegrationRequest(logTypes []string, type_ string, ) *OrgLogIntegrationRequest`

NewOrgLogIntegrationRequest instantiates a new OrgLogIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgLogIntegrationRequestWithDefaults

`func NewOrgLogIntegrationRequestWithDefaults() *OrgLogIntegrationRequest`

NewOrgLogIntegrationRequestWithDefaults instantiates a new OrgLogIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogTypes

`func (o *OrgLogIntegrationRequest) GetLogTypes() []string`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *OrgLogIntegrationRequest) GetLogTypesOk() (*[]string, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *OrgLogIntegrationRequest) SetLogTypes(v []string)`

SetLogTypes sets LogTypes field to given value.

### GetType

`func (o *OrgLogIntegrationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgLogIntegrationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgLogIntegrationRequest) SetType(v string)`

SetType sets Type field to given value.

### GetOtelEndpoint

`func (o *OrgLogIntegrationRequest) GetOtelEndpoint() string`

GetOtelEndpoint returns the OtelEndpoint field if non-nil, zero value otherwise.

### GetOtelEndpointOk

`func (o *OrgLogIntegrationRequest) GetOtelEndpointOk() (*string, bool)`

GetOtelEndpointOk returns a tuple with the OtelEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelEndpoint

`func (o *OrgLogIntegrationRequest) SetOtelEndpoint(v string)`

SetOtelEndpoint sets OtelEndpoint field to given value.

### HasOtelEndpoint

`func (o *OrgLogIntegrationRequest) HasOtelEndpoint() bool`

HasOtelEndpoint returns a boolean if a field has been set.

### SetOtelEndpointNil

`func (o *OrgLogIntegrationRequest) SetOtelEndpointNil()`

SetOtelEndpointNil sets OtelEndpoint to an explicit JSON null when marshaled, overriding any value previously set with SetOtelEndpoint. Calling SetOtelEndpoint again clears the null override.

### GetOtelSuppliedHeaders

`func (o *OrgLogIntegrationRequest) GetOtelSuppliedHeaders() []OrgLogIntegrationHeader`

GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field if non-nil, zero value otherwise.

### GetOtelSuppliedHeadersOk

`func (o *OrgLogIntegrationRequest) GetOtelSuppliedHeadersOk() (*[]OrgLogIntegrationHeader, bool)`

GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelSuppliedHeaders

`func (o *OrgLogIntegrationRequest) SetOtelSuppliedHeaders(v []OrgLogIntegrationHeader)`

SetOtelSuppliedHeaders sets OtelSuppliedHeaders field to given value.

### HasOtelSuppliedHeaders

`func (o *OrgLogIntegrationRequest) HasOtelSuppliedHeaders() bool`

HasOtelSuppliedHeaders returns a boolean if a field has been set.

### SetOtelSuppliedHeadersNil

`func (o *OrgLogIntegrationRequest) SetOtelSuppliedHeadersNil()`

SetOtelSuppliedHeadersNil sets OtelSuppliedHeaders to an explicit JSON null when marshaled, overriding any value previously set with SetOtelSuppliedHeaders. Calling SetOtelSuppliedHeaders again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


