# ClusterDescriptionConnectionStringsPrivateEndpointEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointId** | Pointer to **string** | Unique string that the cloud provider uses to identify the private endpoint. | [optional] [readonly] 
**ProviderName** | Pointer to **string** | Cloud provider in which MongoDB Cloud deploys the private endpoint. | [optional] [readonly] 
**Region** | Pointer to **string** | Region where the private endpoint is deployed. | [optional] [readonly] 

## Methods

### NewClusterDescriptionConnectionStringsPrivateEndpointEndpoint

`func NewClusterDescriptionConnectionStringsPrivateEndpointEndpoint() *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint`

NewClusterDescriptionConnectionStringsPrivateEndpointEndpoint instantiates a new ClusterDescriptionConnectionStringsPrivateEndpointEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDescriptionConnectionStringsPrivateEndpointEndpointWithDefaults

`func NewClusterDescriptionConnectionStringsPrivateEndpointEndpointWithDefaults() *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint`

NewClusterDescriptionConnectionStringsPrivateEndpointEndpointWithDefaults instantiates a new ClusterDescriptionConnectionStringsPrivateEndpointEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointId

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

### GetProviderName

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetRegion

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ClusterDescriptionConnectionStringsPrivateEndpointEndpoint) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


