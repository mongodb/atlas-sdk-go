# ApiPrivateDownloadDeliveryUrl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryUrl** | Pointer to **string** | One Uniform Resource Locator that points to the compressed snapshot files for manual download. | [optional] 
**EndpointId** | Pointer to **string** | Unique 22-character alphanumeric string that identifies the private endpoint. | [optional] 

## Methods

### NewApiPrivateDownloadDeliveryUrl

`func NewApiPrivateDownloadDeliveryUrl() *ApiPrivateDownloadDeliveryUrl`

NewApiPrivateDownloadDeliveryUrl instantiates a new ApiPrivateDownloadDeliveryUrl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPrivateDownloadDeliveryUrlWithDefaults

`func NewApiPrivateDownloadDeliveryUrlWithDefaults() *ApiPrivateDownloadDeliveryUrl`

NewApiPrivateDownloadDeliveryUrlWithDefaults instantiates a new ApiPrivateDownloadDeliveryUrl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryUrl

`func (o *ApiPrivateDownloadDeliveryUrl) GetDeliveryUrl() string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *ApiPrivateDownloadDeliveryUrl) GetDeliveryUrlOk() (*string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *ApiPrivateDownloadDeliveryUrl) SetDeliveryUrl(v string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *ApiPrivateDownloadDeliveryUrl) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.
### GetEndpointId

`func (o *ApiPrivateDownloadDeliveryUrl) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *ApiPrivateDownloadDeliveryUrl) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *ApiPrivateDownloadDeliveryUrl) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### HasEndpointId

`func (o *ApiPrivateDownloadDeliveryUrl) HasEndpointId() bool`

HasEndpointId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


