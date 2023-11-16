# PrivateEndpointHostname

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Human-readable label that identifies the hostname. | [optional] [readonly] 
**PrivateEndpoint** | Pointer to **string** | Human-readable label that identifies private endpoint. | [optional] [readonly] 

## Methods

### NewPrivateEndpointHostname

`func NewPrivateEndpointHostname() *PrivateEndpointHostname`

NewPrivateEndpointHostname instantiates a new PrivateEndpointHostname object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateEndpointHostnameWithDefaults

`func NewPrivateEndpointHostnameWithDefaults() *PrivateEndpointHostname`

NewPrivateEndpointHostnameWithDefaults instantiates a new PrivateEndpointHostname object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *PrivateEndpointHostname) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *PrivateEndpointHostname) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *PrivateEndpointHostname) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *PrivateEndpointHostname) HasHostname() bool`

HasHostname returns a boolean if a field has been set.
### GetPrivateEndpoint

`func (o *PrivateEndpointHostname) GetPrivateEndpoint() string`

GetPrivateEndpoint returns the PrivateEndpoint field if non-nil, zero value otherwise.

### GetPrivateEndpointOk

`func (o *PrivateEndpointHostname) GetPrivateEndpointOk() (*string, bool)`

GetPrivateEndpointOk returns a tuple with the PrivateEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoint

`func (o *PrivateEndpointHostname) SetPrivateEndpoint(v string)`

SetPrivateEndpoint sets PrivateEndpoint field to given value.

### HasPrivateEndpoint

`func (o *PrivateEndpointHostname) HasPrivateEndpoint() bool`

HasPrivateEndpoint returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


