# MongoDBAccessLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResult** | Pointer to **bool** | Flag that indicates whether the response should return successful authentication attempts only. | [optional] 
**AuthSource** | Pointer to **string** | Database against which someone attempted to authenticate. | [optional] [readonly] 
**FailureReason** | Pointer to **string** | Reason that the authentication failed. Null if authentication succeeded. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Human-readable label that identifies the hostname of the target node that received the authentication attempt. | [optional] [readonly] 
**IpAddress** | Pointer to **string** | Internet Protocol address that attempted to authenticate with the database. | [optional] [readonly] 
**LogLine** | Pointer to **string** | Text of the host log concerning the authentication attempt. | [optional] [readonly] 
**Timestamp** | Pointer to **string** | Date and time when someone made this authentication attempt. MongoDB Cloud represents this timestamp in ISO 8601 format in UTC. | [optional] [readonly] 
**Username** | Pointer to **string** | Username used to authenticate against the database. | [optional] [readonly] 

## Methods

### NewMongoDBAccessLogs

`func NewMongoDBAccessLogs() *MongoDBAccessLogs`

NewMongoDBAccessLogs instantiates a new MongoDBAccessLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBAccessLogsWithDefaults

`func NewMongoDBAccessLogsWithDefaults() *MongoDBAccessLogs`

NewMongoDBAccessLogsWithDefaults instantiates a new MongoDBAccessLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResult

`func (o *MongoDBAccessLogs) GetAuthResult() bool`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *MongoDBAccessLogs) GetAuthResultOk() (*bool, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *MongoDBAccessLogs) SetAuthResult(v bool)`

SetAuthResult sets AuthResult field to given value.

### HasAuthResult

`func (o *MongoDBAccessLogs) HasAuthResult() bool`

HasAuthResult returns a boolean if a field has been set.
### GetAuthSource

`func (o *MongoDBAccessLogs) GetAuthSource() string`

GetAuthSource returns the AuthSource field if non-nil, zero value otherwise.

### GetAuthSourceOk

`func (o *MongoDBAccessLogs) GetAuthSourceOk() (*string, bool)`

GetAuthSourceOk returns a tuple with the AuthSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSource

`func (o *MongoDBAccessLogs) SetAuthSource(v string)`

SetAuthSource sets AuthSource field to given value.

### HasAuthSource

`func (o *MongoDBAccessLogs) HasAuthSource() bool`

HasAuthSource returns a boolean if a field has been set.
### GetFailureReason

`func (o *MongoDBAccessLogs) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MongoDBAccessLogs) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MongoDBAccessLogs) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MongoDBAccessLogs) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.
### GetGroupId

`func (o *MongoDBAccessLogs) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MongoDBAccessLogs) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MongoDBAccessLogs) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MongoDBAccessLogs) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetHostname

`func (o *MongoDBAccessLogs) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *MongoDBAccessLogs) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *MongoDBAccessLogs) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *MongoDBAccessLogs) HasHostname() bool`

HasHostname returns a boolean if a field has been set.
### GetIpAddress

`func (o *MongoDBAccessLogs) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *MongoDBAccessLogs) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *MongoDBAccessLogs) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *MongoDBAccessLogs) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.
### GetLogLine

`func (o *MongoDBAccessLogs) GetLogLine() string`

GetLogLine returns the LogLine field if non-nil, zero value otherwise.

### GetLogLineOk

`func (o *MongoDBAccessLogs) GetLogLineOk() (*string, bool)`

GetLogLineOk returns a tuple with the LogLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogLine

`func (o *MongoDBAccessLogs) SetLogLine(v string)`

SetLogLine sets LogLine field to given value.

### HasLogLine

`func (o *MongoDBAccessLogs) HasLogLine() bool`

HasLogLine returns a boolean if a field has been set.
### GetTimestamp

`func (o *MongoDBAccessLogs) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MongoDBAccessLogs) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MongoDBAccessLogs) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MongoDBAccessLogs) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.
### GetUsername

`func (o *MongoDBAccessLogs) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MongoDBAccessLogs) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MongoDBAccessLogs) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MongoDBAccessLogs) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


