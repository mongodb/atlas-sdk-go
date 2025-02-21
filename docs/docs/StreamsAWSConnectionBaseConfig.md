# StreamsAWSConnectionBaseConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**RoleArn** | Pointer to **string** | Amazon Resource Name (ARN) that identifies the Amazon Web Services (AWS) Identity and Access Management (IAM) role that MongoDB Cloud assumes when it accesses resources in your AWS account. | [optional] 

## Methods

### NewStreamsAWSConnectionBaseConfig

`func NewStreamsAWSConnectionBaseConfig() *StreamsAWSConnectionBaseConfig`

NewStreamsAWSConnectionBaseConfig instantiates a new StreamsAWSConnectionBaseConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsAWSConnectionBaseConfigWithDefaults

`func NewStreamsAWSConnectionBaseConfigWithDefaults() *StreamsAWSConnectionBaseConfig`

NewStreamsAWSConnectionBaseConfigWithDefaults instantiates a new StreamsAWSConnectionBaseConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsAWSConnectionBaseConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsAWSConnectionBaseConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsAWSConnectionBaseConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsAWSConnectionBaseConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRoleArn

`func (o *StreamsAWSConnectionBaseConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *StreamsAWSConnectionBaseConfig) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *StreamsAWSConnectionBaseConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *StreamsAWSConnectionBaseConfig) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


