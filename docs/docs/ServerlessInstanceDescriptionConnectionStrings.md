# ServerlessInstanceDescriptionConnectionStrings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateEndpoint** | Pointer to [**[]ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint**](ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint.md) | List of private endpoint-aware connection strings that you can use to connect to this serverless instance through a private endpoint. This parameter returns only if you created a private endpoint for this serverless instance and it is AVAILABLE. | [optional] [readonly] 
**StandardSrv** | Pointer to **string** | Public connection string that you can use to connect to this serverless instance. This connection string uses the &#x60;mongodb+srv://&#x60; protocol. | [optional] [readonly] 

## Methods

### NewServerlessInstanceDescriptionConnectionStrings

`func NewServerlessInstanceDescriptionConnectionStrings() *ServerlessInstanceDescriptionConnectionStrings`

NewServerlessInstanceDescriptionConnectionStrings instantiates a new ServerlessInstanceDescriptionConnectionStrings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessInstanceDescriptionConnectionStringsWithDefaults

`func NewServerlessInstanceDescriptionConnectionStringsWithDefaults() *ServerlessInstanceDescriptionConnectionStrings`

NewServerlessInstanceDescriptionConnectionStringsWithDefaults instantiates a new ServerlessInstanceDescriptionConnectionStrings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateEndpoint

`func (o *ServerlessInstanceDescriptionConnectionStrings) GetPrivateEndpoint() []ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint`

GetPrivateEndpoint returns the PrivateEndpoint field if non-nil, zero value otherwise.

### GetPrivateEndpointOk

`func (o *ServerlessInstanceDescriptionConnectionStrings) GetPrivateEndpointOk() (*[]ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint, bool)`

GetPrivateEndpointOk returns a tuple with the PrivateEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoint

`func (o *ServerlessInstanceDescriptionConnectionStrings) SetPrivateEndpoint(v []ServerlessInstanceDescriptionConnectionStringsPrivateEndpoint)`

SetPrivateEndpoint sets PrivateEndpoint field to given value.

### HasPrivateEndpoint

`func (o *ServerlessInstanceDescriptionConnectionStrings) HasPrivateEndpoint() bool`

HasPrivateEndpoint returns a boolean if a field has been set.

### GetStandardSrv

`func (o *ServerlessInstanceDescriptionConnectionStrings) GetStandardSrv() string`

GetStandardSrv returns the StandardSrv field if non-nil, zero value otherwise.

### GetStandardSrvOk

`func (o *ServerlessInstanceDescriptionConnectionStrings) GetStandardSrvOk() (*string, bool)`

GetStandardSrvOk returns a tuple with the StandardSrv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardSrv

`func (o *ServerlessInstanceDescriptionConnectionStrings) SetStandardSrv(v string)`

SetStandardSrv sets StandardSrv field to given value.

### HasStandardSrv

`func (o *ServerlessInstanceDescriptionConnectionStrings) HasStandardSrv() bool`

HasStandardSrv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


