# ServerlessConnectionStringsPrivateEndpointList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]ServerlessConnectionStringsPrivateEndpointItem**](ServerlessConnectionStringsPrivateEndpointItem.md) | List that contains the private endpoints through which you connect to MongoDB Cloud when you use &#x60;connectionStrings.privateEndpoint[n].srvConnectionString&#x60;. | [optional] [readonly] 
**SrvConnectionString** | Pointer to **string** | Private endpoint-aware connection string that uses the &#x60;mongodb+srv://&#x60; protocol to connect to MongoDB Cloud through a private endpoint. The &#x60;mongodb+srv&#x60; protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS). | [optional] [readonly] 
**Type** | Pointer to **string** | MongoDB process type to which your application connects. | [optional] [readonly] 

## Methods

### NewServerlessConnectionStringsPrivateEndpointList

`func NewServerlessConnectionStringsPrivateEndpointList() *ServerlessConnectionStringsPrivateEndpointList`

NewServerlessConnectionStringsPrivateEndpointList instantiates a new ServerlessConnectionStringsPrivateEndpointList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessConnectionStringsPrivateEndpointListWithDefaults

`func NewServerlessConnectionStringsPrivateEndpointListWithDefaults() *ServerlessConnectionStringsPrivateEndpointList`

NewServerlessConnectionStringsPrivateEndpointListWithDefaults instantiates a new ServerlessConnectionStringsPrivateEndpointList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *ServerlessConnectionStringsPrivateEndpointList) GetEndpoints() []ServerlessConnectionStringsPrivateEndpointItem`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ServerlessConnectionStringsPrivateEndpointList) GetEndpointsOk() (*[]ServerlessConnectionStringsPrivateEndpointItem, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ServerlessConnectionStringsPrivateEndpointList) SetEndpoints(v []ServerlessConnectionStringsPrivateEndpointItem)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ServerlessConnectionStringsPrivateEndpointList) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.
### GetSrvConnectionString

`func (o *ServerlessConnectionStringsPrivateEndpointList) GetSrvConnectionString() string`

GetSrvConnectionString returns the SrvConnectionString field if non-nil, zero value otherwise.

### GetSrvConnectionStringOk

`func (o *ServerlessConnectionStringsPrivateEndpointList) GetSrvConnectionStringOk() (*string, bool)`

GetSrvConnectionStringOk returns a tuple with the SrvConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrvConnectionString

`func (o *ServerlessConnectionStringsPrivateEndpointList) SetSrvConnectionString(v string)`

SetSrvConnectionString sets SrvConnectionString field to given value.

### HasSrvConnectionString

`func (o *ServerlessConnectionStringsPrivateEndpointList) HasSrvConnectionString() bool`

HasSrvConnectionString returns a boolean if a field has been set.
### GetType

`func (o *ServerlessConnectionStringsPrivateEndpointList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerlessConnectionStringsPrivateEndpointList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerlessConnectionStringsPrivateEndpointList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ServerlessConnectionStringsPrivateEndpointList) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


