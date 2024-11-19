# AWSAccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountId** | Pointer to **string** | The AWS Account ID. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**VpcId** | Pointer to **string** | The VPC ID. | [optional] 

## Methods

### NewAWSAccountDetails

`func NewAWSAccountDetails() *AWSAccountDetails`

NewAWSAccountDetails instantiates a new AWSAccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSAccountDetailsWithDefaults

`func NewAWSAccountDetailsWithDefaults() *AWSAccountDetails`

NewAWSAccountDetailsWithDefaults instantiates a new AWSAccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountId

`func (o *AWSAccountDetails) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *AWSAccountDetails) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *AWSAccountDetails) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.

### HasAwsAccountId

`func (o *AWSAccountDetails) HasAwsAccountId() bool`

HasAwsAccountId returns a boolean if a field has been set.
### GetLinks

`func (o *AWSAccountDetails) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AWSAccountDetails) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AWSAccountDetails) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AWSAccountDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetVpcId

`func (o *AWSAccountDetails) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AWSAccountDetails) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AWSAccountDetails) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AWSAccountDetails) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


