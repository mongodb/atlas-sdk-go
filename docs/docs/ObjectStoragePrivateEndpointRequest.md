# ObjectStoragePrivateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider. | [optional] 
**RegionName** | Pointer to **string** | Cloud provider region of the S3 bucket that the Object Storage private endpoint accesses. For same-region endpoints, this is also the region where the VPC interface endpoint is deployed. | [optional] 
**VpcRegionName** | Pointer to **string** | Cloud provider region in which the VPC interface endpoint is deployed. Omit to deploy the interface endpoint in the same region as the S3 bucket (same-region endpoint). Set to a region different from &#x60;regionName&#x60; to create a cross-region endpoint. | [optional] 

## Methods

### NewObjectStoragePrivateEndpointRequest

`func NewObjectStoragePrivateEndpointRequest() *ObjectStoragePrivateEndpointRequest`

NewObjectStoragePrivateEndpointRequest instantiates a new ObjectStoragePrivateEndpointRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePrivateEndpointRequestWithDefaults

`func NewObjectStoragePrivateEndpointRequestWithDefaults() *ObjectStoragePrivateEndpointRequest`

NewObjectStoragePrivateEndpointRequestWithDefaults instantiates a new ObjectStoragePrivateEndpointRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ObjectStoragePrivateEndpointRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ObjectStoragePrivateEndpointRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ObjectStoragePrivateEndpointRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ObjectStoragePrivateEndpointRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetRegionName

`func (o *ObjectStoragePrivateEndpointRequest) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ObjectStoragePrivateEndpointRequest) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ObjectStoragePrivateEndpointRequest) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ObjectStoragePrivateEndpointRequest) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetVpcRegionName

`func (o *ObjectStoragePrivateEndpointRequest) GetVpcRegionName() string`

GetVpcRegionName returns the VpcRegionName field if non-nil, zero value otherwise.

### GetVpcRegionNameOk

`func (o *ObjectStoragePrivateEndpointRequest) GetVpcRegionNameOk() (*string, bool)`

GetVpcRegionNameOk returns a tuple with the VpcRegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRegionName

`func (o *ObjectStoragePrivateEndpointRequest) SetVpcRegionName(v string)`

SetVpcRegionName sets VpcRegionName field to given value.

### HasVpcRegionName

`func (o *ObjectStoragePrivateEndpointRequest) HasVpcRegionName() bool`

HasVpcRegionName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


