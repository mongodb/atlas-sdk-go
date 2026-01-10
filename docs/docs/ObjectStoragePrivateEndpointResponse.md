# ObjectStoragePrivateEndpointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Human-readable label that identifies the cloud provider. | [optional] 
**ErrorMessage** | Pointer to **string** | Error message for failures associated with the Object Storage private endpoint. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**PrivateEndpointConnectionName** | Pointer to **string** | Connection name of the Private Endpoint. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region in which the Object Storage private endpoint is located. | [optional] 
**Status** | Pointer to **string** | State of the Object Storage private endpoint. | [optional] [readonly] 

## Methods

### NewObjectStoragePrivateEndpointResponse

`func NewObjectStoragePrivateEndpointResponse() *ObjectStoragePrivateEndpointResponse`

NewObjectStoragePrivateEndpointResponse instantiates a new ObjectStoragePrivateEndpointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectStoragePrivateEndpointResponseWithDefaults

`func NewObjectStoragePrivateEndpointResponseWithDefaults() *ObjectStoragePrivateEndpointResponse`

NewObjectStoragePrivateEndpointResponseWithDefaults instantiates a new ObjectStoragePrivateEndpointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *ObjectStoragePrivateEndpointResponse) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ObjectStoragePrivateEndpointResponse) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ObjectStoragePrivateEndpointResponse) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *ObjectStoragePrivateEndpointResponse) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetErrorMessage

`func (o *ObjectStoragePrivateEndpointResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ObjectStoragePrivateEndpointResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ObjectStoragePrivateEndpointResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ObjectStoragePrivateEndpointResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetId

`func (o *ObjectStoragePrivateEndpointResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectStoragePrivateEndpointResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectStoragePrivateEndpointResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObjectStoragePrivateEndpointResponse) HasId() bool`

HasId returns a boolean if a field has been set.
### GetPrivateEndpointConnectionName

`func (o *ObjectStoragePrivateEndpointResponse) GetPrivateEndpointConnectionName() string`

GetPrivateEndpointConnectionName returns the PrivateEndpointConnectionName field if non-nil, zero value otherwise.

### GetPrivateEndpointConnectionNameOk

`func (o *ObjectStoragePrivateEndpointResponse) GetPrivateEndpointConnectionNameOk() (*string, bool)`

GetPrivateEndpointConnectionNameOk returns a tuple with the PrivateEndpointConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpointConnectionName

`func (o *ObjectStoragePrivateEndpointResponse) SetPrivateEndpointConnectionName(v string)`

SetPrivateEndpointConnectionName sets PrivateEndpointConnectionName field to given value.

### HasPrivateEndpointConnectionName

`func (o *ObjectStoragePrivateEndpointResponse) HasPrivateEndpointConnectionName() bool`

HasPrivateEndpointConnectionName returns a boolean if a field has been set.
### GetRegionName

`func (o *ObjectStoragePrivateEndpointResponse) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *ObjectStoragePrivateEndpointResponse) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *ObjectStoragePrivateEndpointResponse) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *ObjectStoragePrivateEndpointResponse) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.
### GetStatus

`func (o *ObjectStoragePrivateEndpointResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ObjectStoragePrivateEndpointResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ObjectStoragePrivateEndpointResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ObjectStoragePrivateEndpointResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


