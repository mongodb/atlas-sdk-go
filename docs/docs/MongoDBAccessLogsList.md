# MongoDBAccessLogsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLogs** | Pointer to [**[]MongoDBAccessLogs**](MongoDBAccessLogs.md) | Authentication attempt, one per object, made against the cluster. | [optional] [readonly] 

## Methods

### NewMongoDBAccessLogsList

`func NewMongoDBAccessLogsList() *MongoDBAccessLogsList`

NewMongoDBAccessLogsList instantiates a new MongoDBAccessLogsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBAccessLogsListWithDefaults

`func NewMongoDBAccessLogsListWithDefaults() *MongoDBAccessLogsList`

NewMongoDBAccessLogsListWithDefaults instantiates a new MongoDBAccessLogsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLogs

`func (o *MongoDBAccessLogsList) GetAccessLogs() []MongoDBAccessLogs`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *MongoDBAccessLogsList) GetAccessLogsOk() (*[]MongoDBAccessLogs, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *MongoDBAccessLogsList) SetAccessLogs(v []MongoDBAccessLogs)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *MongoDBAccessLogsList) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


