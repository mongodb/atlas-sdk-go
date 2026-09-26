# OrgLogIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-character hexadecimal digit string that identifies the log integration configuration. | [readonly] 
**LogTypes** | **[]string** | Array of log types exported by this integration. | 
**Type** | **string** | Human-readable label that identifies the service to which you want to integrate with Atlas. The value must match the log integration type. This value cannot be modified after the integration is created. | 
**OtelEndpoint** | Pointer to **string** | OpenTelemetry collector endpoint URL. | [optional] 
**OtelSuppliedHeaders** | Pointer to [**[]OrgLogIntegrationHeader**](OrgLogIntegrationHeader.md) | HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB. Values are redacted. | [optional] 

## Methods

### NewOrgLogIntegrationResponse

`func NewOrgLogIntegrationResponse(id string, logTypes []string, type_ string, ) *OrgLogIntegrationResponse`

NewOrgLogIntegrationResponse instantiates a new OrgLogIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgLogIntegrationResponseWithDefaults

`func NewOrgLogIntegrationResponseWithDefaults() *OrgLogIntegrationResponse`

NewOrgLogIntegrationResponseWithDefaults instantiates a new OrgLogIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrgLogIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrgLogIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrgLogIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.

### GetLogTypes

`func (o *OrgLogIntegrationResponse) GetLogTypes() []string`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *OrgLogIntegrationResponse) GetLogTypesOk() (*[]string, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *OrgLogIntegrationResponse) SetLogTypes(v []string)`

SetLogTypes sets LogTypes field to given value.

### GetType

`func (o *OrgLogIntegrationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrgLogIntegrationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrgLogIntegrationResponse) SetType(v string)`

SetType sets Type field to given value.

### GetOtelEndpoint

`func (o *OrgLogIntegrationResponse) GetOtelEndpoint() string`

GetOtelEndpoint returns the OtelEndpoint field if non-nil, zero value otherwise.

### GetOtelEndpointOk

`func (o *OrgLogIntegrationResponse) GetOtelEndpointOk() (*string, bool)`

GetOtelEndpointOk returns a tuple with the OtelEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelEndpoint

`func (o *OrgLogIntegrationResponse) SetOtelEndpoint(v string)`

SetOtelEndpoint sets OtelEndpoint field to given value.

### HasOtelEndpoint

`func (o *OrgLogIntegrationResponse) HasOtelEndpoint() bool`

HasOtelEndpoint returns a boolean if a field has been set.

### SetOtelEndpointNil

`func (o *OrgLogIntegrationResponse) SetOtelEndpointNil()`

SetOtelEndpointNil sets OtelEndpoint to an explicit JSON null when marshaled, overriding any value previously set with SetOtelEndpoint. Calling SetOtelEndpoint again clears the null override.

### GetOtelSuppliedHeaders

`func (o *OrgLogIntegrationResponse) GetOtelSuppliedHeaders() []OrgLogIntegrationHeader`

GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field if non-nil, zero value otherwise.

### GetOtelSuppliedHeadersOk

`func (o *OrgLogIntegrationResponse) GetOtelSuppliedHeadersOk() (*[]OrgLogIntegrationHeader, bool)`

GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelSuppliedHeaders

`func (o *OrgLogIntegrationResponse) SetOtelSuppliedHeaders(v []OrgLogIntegrationHeader)`

SetOtelSuppliedHeaders sets OtelSuppliedHeaders field to given value.

### HasOtelSuppliedHeaders

`func (o *OrgLogIntegrationResponse) HasOtelSuppliedHeaders() bool`

HasOtelSuppliedHeaders returns a boolean if a field has been set.

### SetOtelSuppliedHeadersNil

`func (o *OrgLogIntegrationResponse) SetOtelSuppliedHeadersNil()`

SetOtelSuppliedHeadersNil sets OtelSuppliedHeaders to an explicit JSON null when marshaled, overriding any value previously set with SetOtelSuppliedHeaders. Calling SetOtelSuppliedHeaders again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


