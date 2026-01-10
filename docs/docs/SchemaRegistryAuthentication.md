# SchemaRegistryAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Authentication type discriminator. Specifies the authentication mechanism for Confluent Schema Registry. | 
**Password** | Pointer to **string** | Password or Private Key for authentication. | [optional] 
**Username** | Pointer to **string** | Username or Public Key for authentication. | [optional] 

## Methods

### NewSchemaRegistryAuthentication

`func NewSchemaRegistryAuthentication(type_ string, ) *SchemaRegistryAuthentication`

NewSchemaRegistryAuthentication instantiates a new SchemaRegistryAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaRegistryAuthenticationWithDefaults

`func NewSchemaRegistryAuthenticationWithDefaults() *SchemaRegistryAuthentication`

NewSchemaRegistryAuthenticationWithDefaults instantiates a new SchemaRegistryAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaRegistryAuthentication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaRegistryAuthentication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaRegistryAuthentication) SetType(v string)`

SetType sets Type field to given value.

### GetPassword

`func (o *SchemaRegistryAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SchemaRegistryAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SchemaRegistryAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SchemaRegistryAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.
### GetUsername

`func (o *SchemaRegistryAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SchemaRegistryAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SchemaRegistryAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SchemaRegistryAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


