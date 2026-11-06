# StreamsSnowflakeConnectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | [**StreamsSnowflakeAuthentication**](StreamsSnowflakeAuthentication.md) |  | 

## Methods

### NewStreamsSnowflakeConnectionConfig

`func NewStreamsSnowflakeConnectionConfig(authentication StreamsSnowflakeAuthentication, ) *StreamsSnowflakeConnectionConfig`

NewStreamsSnowflakeConnectionConfig instantiates a new StreamsSnowflakeConnectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsSnowflakeConnectionConfigWithDefaults

`func NewStreamsSnowflakeConnectionConfigWithDefaults() *StreamsSnowflakeConnectionConfig`

NewStreamsSnowflakeConnectionConfigWithDefaults instantiates a new StreamsSnowflakeConnectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthentication

`func (o *StreamsSnowflakeConnectionConfig) GetAuthentication() StreamsSnowflakeAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *StreamsSnowflakeConnectionConfig) GetAuthenticationOk() (*StreamsSnowflakeAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *StreamsSnowflakeConnectionConfig) SetAuthentication(v StreamsSnowflakeAuthentication)`

SetAuthentication sets Authentication field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


