# CloudProviderEndpointServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Human-readable label that identifies the cloud service provider for which you want to create the private endpoint service. | 
**Region** | **string** | Cloud provider region in which you want to create the private endpoint service. Regions accepted as values differ for [Amazon Web Services](https://docs.atlas.mongodb.com/reference/amazon-aws/), [Google Cloud Platform](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Microsoft Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/). | 

## Methods

### NewCloudProviderEndpointServiceRequest

`func NewCloudProviderEndpointServiceRequest(providerName string, region string, ) *CloudProviderEndpointServiceRequest`

NewCloudProviderEndpointServiceRequest instantiates a new CloudProviderEndpointServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderEndpointServiceRequestWithDefaults

`func NewCloudProviderEndpointServiceRequestWithDefaults() *CloudProviderEndpointServiceRequest`

NewCloudProviderEndpointServiceRequestWithDefaults instantiates a new CloudProviderEndpointServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderName

`func (o *CloudProviderEndpointServiceRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CloudProviderEndpointServiceRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CloudProviderEndpointServiceRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### GetRegion

`func (o *CloudProviderEndpointServiceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudProviderEndpointServiceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudProviderEndpointServiceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


