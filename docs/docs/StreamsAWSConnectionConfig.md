# StreamsAWSConnectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account. | [optional] 
**TestBucket** | Pointer to **string** | The name of an S3 bucket used to check authorization of the passed-in IAM role ARN. | [optional] 

## Methods

### NewStreamsAWSConnectionConfig

`func NewStreamsAWSConnectionConfig() *StreamsAWSConnectionConfig`

NewStreamsAWSConnectionConfig instantiates a new StreamsAWSConnectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsAWSConnectionConfigWithDefaults

`func NewStreamsAWSConnectionConfigWithDefaults() *StreamsAWSConnectionConfig`

NewStreamsAWSConnectionConfigWithDefaults instantiates a new StreamsAWSConnectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsAWSConnectionConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsAWSConnectionConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsAWSConnectionConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsAWSConnectionConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRoleArn

`func (o *StreamsAWSConnectionConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *StreamsAWSConnectionConfig) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *StreamsAWSConnectionConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *StreamsAWSConnectionConfig) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.
### GetTestBucket

`func (o *StreamsAWSConnectionConfig) GetTestBucket() string`

GetTestBucket returns the TestBucket field if non-nil, zero value otherwise.

### GetTestBucketOk

`func (o *StreamsAWSConnectionConfig) GetTestBucketOk() (*string, bool)`

GetTestBucketOk returns a tuple with the TestBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestBucket

`func (o *StreamsAWSConnectionConfig) SetTestBucket(v string)`

SetTestBucket sets TestBucket field to given value.

### HasTestBucket

`func (o *StreamsAWSConnectionConfig) HasTestBucket() bool`

HasTestBucket returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


