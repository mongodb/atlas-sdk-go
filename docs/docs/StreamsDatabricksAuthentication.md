# StreamsDatabricksAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | OAuth client ID of the Databricks service principal. | 
**ClientSecret** | **string** | OAuth client secret of the Databricks service principal. Never returned in responses; omit or leave blank on update to keep the stored secret. | 

## Methods

### NewStreamsDatabricksAuthentication

`func NewStreamsDatabricksAuthentication(clientId string, clientSecret string, ) *StreamsDatabricksAuthentication`

NewStreamsDatabricksAuthentication instantiates a new StreamsDatabricksAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsDatabricksAuthenticationWithDefaults

`func NewStreamsDatabricksAuthenticationWithDefaults() *StreamsDatabricksAuthentication`

NewStreamsDatabricksAuthenticationWithDefaults instantiates a new StreamsDatabricksAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *StreamsDatabricksAuthentication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *StreamsDatabricksAuthentication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *StreamsDatabricksAuthentication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### GetClientSecret

`func (o *StreamsDatabricksAuthentication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *StreamsDatabricksAuthentication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *StreamsDatabricksAuthentication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


