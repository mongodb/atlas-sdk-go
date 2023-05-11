# CreateGCPEndpointGroupRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointGroupName** | Pointer to **string** | Human-readable label that identifies a set of endpoints. | [optional] 
**Endpoints** | Pointer to [**[]CreateGCPForwardingRuleRequest**](CreateGCPForwardingRuleRequest.md) | List of individual private endpoints that comprise this endpoint group. | [optional] 
**GcpProjectId** | Pointer to **string** | Unique string that identifies the Google Cloud project in which you created the endpoints. | [optional] 

## Methods

### NewCreateGCPEndpointGroupRequestAllOf

`func NewCreateGCPEndpointGroupRequestAllOf() *CreateGCPEndpointGroupRequestAllOf`

NewCreateGCPEndpointGroupRequestAllOf instantiates a new CreateGCPEndpointGroupRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGCPEndpointGroupRequestAllOfWithDefaults

`func NewCreateGCPEndpointGroupRequestAllOfWithDefaults() *CreateGCPEndpointGroupRequestAllOf`

NewCreateGCPEndpointGroupRequestAllOfWithDefaults instantiates a new CreateGCPEndpointGroupRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointGroupName

`func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpointGroupName() string`

GetEndpointGroupName returns the EndpointGroupName field if non-nil, zero value otherwise.

### GetEndpointGroupNameOk

`func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpointGroupNameOk() (*string, bool)`

GetEndpointGroupNameOk returns a tuple with the EndpointGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupName

`func (o *CreateGCPEndpointGroupRequestAllOf) SetEndpointGroupName(v string)`

SetEndpointGroupName sets EndpointGroupName field to given value.

### HasEndpointGroupName

`func (o *CreateGCPEndpointGroupRequestAllOf) HasEndpointGroupName() bool`

HasEndpointGroupName returns a boolean if a field has been set.

### GetEndpoints

`func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpoints() []CreateGCPForwardingRuleRequest`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *CreateGCPEndpointGroupRequestAllOf) GetEndpointsOk() (*[]CreateGCPForwardingRuleRequest, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *CreateGCPEndpointGroupRequestAllOf) SetEndpoints(v []CreateGCPForwardingRuleRequest)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *CreateGCPEndpointGroupRequestAllOf) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *CreateGCPEndpointGroupRequestAllOf) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *CreateGCPEndpointGroupRequestAllOf) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *CreateGCPEndpointGroupRequestAllOf) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *CreateGCPEndpointGroupRequestAllOf) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


