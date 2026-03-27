# ApiAtlasModifyEndpointServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Human-readable label that identifies the cloud service provider for the private endpoint service which you want to update. | 
**SupportedRemoteRegions** | Pointer to **[]string** | List of regions that the endpoint service supports. Native cross region support is implemented for AWS only. | [optional] 

## Methods

### NewApiAtlasModifyEndpointServiceRequest

`func NewApiAtlasModifyEndpointServiceRequest(cloudProvider string, ) *ApiAtlasModifyEndpointServiceRequest`

NewApiAtlasModifyEndpointServiceRequest instantiates a new ApiAtlasModifyEndpointServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAtlasModifyEndpointServiceRequestWithDefaults

`func NewApiAtlasModifyEndpointServiceRequestWithDefaults() *ApiAtlasModifyEndpointServiceRequest`

NewApiAtlasModifyEndpointServiceRequestWithDefaults instantiates a new ApiAtlasModifyEndpointServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ApiAtlasModifyEndpointServiceRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ApiAtlasModifyEndpointServiceRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ApiAtlasModifyEndpointServiceRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetSupportedRemoteRegions

`func (o *ApiAtlasModifyEndpointServiceRequest) GetSupportedRemoteRegions() []string`

GetSupportedRemoteRegions returns the SupportedRemoteRegions field if non-nil, zero value otherwise.

### GetSupportedRemoteRegionsOk

`func (o *ApiAtlasModifyEndpointServiceRequest) GetSupportedRemoteRegionsOk() (*[]string, bool)`

GetSupportedRemoteRegionsOk returns a tuple with the SupportedRemoteRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRemoteRegions

`func (o *ApiAtlasModifyEndpointServiceRequest) SetSupportedRemoteRegions(v []string)`

SetSupportedRemoteRegions sets SupportedRemoteRegions field to given value.

### HasSupportedRemoteRegions

`func (o *ApiAtlasModifyEndpointServiceRequest) HasSupportedRemoteRegions() bool`

HasSupportedRemoteRegions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


