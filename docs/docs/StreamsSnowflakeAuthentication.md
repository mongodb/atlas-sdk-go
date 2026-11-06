# StreamsSnowflakeAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | Snowflake account identifier. | 
**PrivateKey** | **string** | Private key of the service user&#39;s key pair. Never returned in responses; omit or leave blank on update to keep the stored key. | 
**Role** | **string** | Snowflake role the service user assumes. | 
**User** | **string** | Snowflake service user. | 

## Methods

### NewStreamsSnowflakeAuthentication

`func NewStreamsSnowflakeAuthentication(account string, privateKey string, role string, user string, ) *StreamsSnowflakeAuthentication`

NewStreamsSnowflakeAuthentication instantiates a new StreamsSnowflakeAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsSnowflakeAuthenticationWithDefaults

`func NewStreamsSnowflakeAuthenticationWithDefaults() *StreamsSnowflakeAuthentication`

NewStreamsSnowflakeAuthenticationWithDefaults instantiates a new StreamsSnowflakeAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *StreamsSnowflakeAuthentication) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *StreamsSnowflakeAuthentication) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *StreamsSnowflakeAuthentication) SetAccount(v string)`

SetAccount sets Account field to given value.

### GetPrivateKey

`func (o *StreamsSnowflakeAuthentication) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *StreamsSnowflakeAuthentication) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *StreamsSnowflakeAuthentication) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### GetRole

`func (o *StreamsSnowflakeAuthentication) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *StreamsSnowflakeAuthentication) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *StreamsSnowflakeAuthentication) SetRole(v string)`

SetRole sets Role field to given value.

### GetUser

`func (o *StreamsSnowflakeAuthentication) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *StreamsSnowflakeAuthentication) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *StreamsSnowflakeAuthentication) SetUser(v string)`

SetUser sets User field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


