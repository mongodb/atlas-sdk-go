# AWSCustomDNSEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Flag that indicates whether the project&#39;s clusters deployed to Amazon Web Services (AWS) use a custom Domain Name System (DNS). When &#x60;\&quot;enabled\&quot;: true&#x60;, connect to your cluster using Private IP for Peering connection strings. | 

## Methods

### NewAWSCustomDNSEnabled

`func NewAWSCustomDNSEnabled(enabled bool, ) *AWSCustomDNSEnabled`

NewAWSCustomDNSEnabled instantiates a new AWSCustomDNSEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCustomDNSEnabledWithDefaults

`func NewAWSCustomDNSEnabledWithDefaults() *AWSCustomDNSEnabled`

NewAWSCustomDNSEnabledWithDefaults instantiates a new AWSCustomDNSEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AWSCustomDNSEnabled) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AWSCustomDNSEnabled) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AWSCustomDNSEnabled) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


