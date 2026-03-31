# StreamsGCPConnectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**ServiceAccountId** | Pointer to **string** | Email address of the Google Cloud Platform (GCP) service account that Atlas Streams uses to connect to the GCP Pub/Sub resources. | [optional] 

## Methods

### NewStreamsGCPConnectionConfig

`func NewStreamsGCPConnectionConfig() *StreamsGCPConnectionConfig`

NewStreamsGCPConnectionConfig instantiates a new StreamsGCPConnectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsGCPConnectionConfigWithDefaults

`func NewStreamsGCPConnectionConfigWithDefaults() *StreamsGCPConnectionConfig`

NewStreamsGCPConnectionConfigWithDefaults instantiates a new StreamsGCPConnectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *StreamsGCPConnectionConfig) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsGCPConnectionConfig) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsGCPConnectionConfig) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsGCPConnectionConfig) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetServiceAccountId

`func (o *StreamsGCPConnectionConfig) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *StreamsGCPConnectionConfig) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *StreamsGCPConnectionConfig) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *StreamsGCPConnectionConfig) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


