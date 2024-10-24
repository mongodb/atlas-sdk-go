# PrivateNetworkEndpointIdEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureConnectionName** | Pointer to **string** | Connection name that identifies the Azure private endpoint connection. | [optional] 
**AzureLinkId** | Pointer to **string** | Link ID that identifies the Azure private endpoint connection. | [optional] 
**Comment** | Pointer to **string** | Human-readable string to associate with this private endpoint. | [optional] 
**CustomerEndpointDNSName** | Pointer to **string** | Human-readable label to identify customer&#39;s VPC endpoint DNS name. | [optional] 
**CustomerEndpointIPAddress** | Pointer to **string** | IP address used to connect to the Azure private endpoint. | [optional] 
**EndpointId** | **string** | Unique 22-character alphanumeric string that identifies the private endpoint. | 
**ErrorMessage** | Pointer to **string** | Error message describing a failure approving the private endpoint request. | [optional] 
**Provider** | Pointer to **string** | Human-readable label that identifies the cloud service provider. Atlas Data Lake supports Amazon Web Services only. | [optional] [default to "AWS"]
**Region** | Pointer to **string** | Human-readable label to identify the region of customer&#39;s VPC endpoint. | [optional] 
**Status** | Pointer to **string** | Status of the private endpoint connection request. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies the resource type associated with this private endpoint. | [optional] [default to "DATA_LAKE"]

## Methods

### NewPrivateNetworkEndpointIdEntry

`func NewPrivateNetworkEndpointIdEntry(endpointId string, ) *PrivateNetworkEndpointIdEntry`

NewPrivateNetworkEndpointIdEntry instantiates a new PrivateNetworkEndpointIdEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateNetworkEndpointIdEntryWithDefaults

`func NewPrivateNetworkEndpointIdEntryWithDefaults() *PrivateNetworkEndpointIdEntry`

NewPrivateNetworkEndpointIdEntryWithDefaults instantiates a new PrivateNetworkEndpointIdEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureConnectionName

`func (o *PrivateNetworkEndpointIdEntry) GetAzureConnectionName() string`

GetAzureConnectionName returns the AzureConnectionName field if non-nil, zero value otherwise.

### GetAzureConnectionNameOk

`func (o *PrivateNetworkEndpointIdEntry) GetAzureConnectionNameOk() (*string, bool)`

GetAzureConnectionNameOk returns a tuple with the AzureConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureConnectionName

`func (o *PrivateNetworkEndpointIdEntry) SetAzureConnectionName(v string)`

SetAzureConnectionName sets AzureConnectionName field to given value.

### HasAzureConnectionName

`func (o *PrivateNetworkEndpointIdEntry) HasAzureConnectionName() bool`

HasAzureConnectionName returns a boolean if a field has been set.
### GetAzureLinkId

`func (o *PrivateNetworkEndpointIdEntry) GetAzureLinkId() string`

GetAzureLinkId returns the AzureLinkId field if non-nil, zero value otherwise.

### GetAzureLinkIdOk

`func (o *PrivateNetworkEndpointIdEntry) GetAzureLinkIdOk() (*string, bool)`

GetAzureLinkIdOk returns a tuple with the AzureLinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureLinkId

`func (o *PrivateNetworkEndpointIdEntry) SetAzureLinkId(v string)`

SetAzureLinkId sets AzureLinkId field to given value.

### HasAzureLinkId

`func (o *PrivateNetworkEndpointIdEntry) HasAzureLinkId() bool`

HasAzureLinkId returns a boolean if a field has been set.
### GetComment

`func (o *PrivateNetworkEndpointIdEntry) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PrivateNetworkEndpointIdEntry) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PrivateNetworkEndpointIdEntry) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PrivateNetworkEndpointIdEntry) HasComment() bool`

HasComment returns a boolean if a field has been set.
### GetCustomerEndpointDNSName

`func (o *PrivateNetworkEndpointIdEntry) GetCustomerEndpointDNSName() string`

GetCustomerEndpointDNSName returns the CustomerEndpointDNSName field if non-nil, zero value otherwise.

### GetCustomerEndpointDNSNameOk

`func (o *PrivateNetworkEndpointIdEntry) GetCustomerEndpointDNSNameOk() (*string, bool)`

GetCustomerEndpointDNSNameOk returns a tuple with the CustomerEndpointDNSName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEndpointDNSName

`func (o *PrivateNetworkEndpointIdEntry) SetCustomerEndpointDNSName(v string)`

SetCustomerEndpointDNSName sets CustomerEndpointDNSName field to given value.

### HasCustomerEndpointDNSName

`func (o *PrivateNetworkEndpointIdEntry) HasCustomerEndpointDNSName() bool`

HasCustomerEndpointDNSName returns a boolean if a field has been set.
### GetCustomerEndpointIPAddress

`func (o *PrivateNetworkEndpointIdEntry) GetCustomerEndpointIPAddress() string`

GetCustomerEndpointIPAddress returns the CustomerEndpointIPAddress field if non-nil, zero value otherwise.

### GetCustomerEndpointIPAddressOk

`func (o *PrivateNetworkEndpointIdEntry) GetCustomerEndpointIPAddressOk() (*string, bool)`

GetCustomerEndpointIPAddressOk returns a tuple with the CustomerEndpointIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEndpointIPAddress

`func (o *PrivateNetworkEndpointIdEntry) SetCustomerEndpointIPAddress(v string)`

SetCustomerEndpointIPAddress sets CustomerEndpointIPAddress field to given value.

### HasCustomerEndpointIPAddress

`func (o *PrivateNetworkEndpointIdEntry) HasCustomerEndpointIPAddress() bool`

HasCustomerEndpointIPAddress returns a boolean if a field has been set.
### GetEndpointId

`func (o *PrivateNetworkEndpointIdEntry) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### GetEndpointIdOk

`func (o *PrivateNetworkEndpointIdEntry) GetEndpointIdOk() (*string, bool)`

GetEndpointIdOk returns a tuple with the EndpointId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointId

`func (o *PrivateNetworkEndpointIdEntry) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### GetErrorMessage

`func (o *PrivateNetworkEndpointIdEntry) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PrivateNetworkEndpointIdEntry) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PrivateNetworkEndpointIdEntry) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PrivateNetworkEndpointIdEntry) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.
### GetProvider

`func (o *PrivateNetworkEndpointIdEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PrivateNetworkEndpointIdEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PrivateNetworkEndpointIdEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PrivateNetworkEndpointIdEntry) HasProvider() bool`

HasProvider returns a boolean if a field has been set.
### GetRegion

`func (o *PrivateNetworkEndpointIdEntry) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PrivateNetworkEndpointIdEntry) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PrivateNetworkEndpointIdEntry) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PrivateNetworkEndpointIdEntry) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetStatus

`func (o *PrivateNetworkEndpointIdEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateNetworkEndpointIdEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateNetworkEndpointIdEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateNetworkEndpointIdEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.
### GetType

`func (o *PrivateNetworkEndpointIdEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrivateNetworkEndpointIdEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrivateNetworkEndpointIdEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PrivateNetworkEndpointIdEntry) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


