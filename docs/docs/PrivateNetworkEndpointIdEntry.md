# PrivateNetworkEndpointIdEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureLinkId** | Pointer to **string** | Link ID that identifies the Azure private endpoint connection. | [optional] 
**Comment** | Pointer to **string** | Human-readable string to associate with this private endpoint. | [optional] 
**CustomerEndpointDNSName** | Pointer to **string** | Human-readable label to identify customer&#39;s VPC endpoint DNS name. If defined, you must also specify a value for &#x60;region&#x60;. | [optional] 
**CustomerEndpointIPAddress** | Pointer to **string** | IP address used to connect to the Azure private endpoint. | [optional] 
**EndpointId** | **string** | Unique string that identifies the private endpoint. For AWS, this is a 22-character alphanumeric string in the format &#x60;vpce-&lt;17 hex characters&gt;&#x60;. For Azure, this is the full resource ID in the format &#x60;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{endpointName}&#x60;. | 
**ErrorMessage** | Pointer to **string** | Error message describing a failure approving the private endpoint request. | [optional] 
**Provider** | Pointer to **string** | Human-readable label that identifies the cloud service provider. Atlas Data Lake supports Amazon Web Services and Azure. | [optional] [default to "AWS"]
**Region** | Pointer to **string** | Human-readable label to identify the region of customer&#39;s VPC endpoint. If defined, you must also specify a value for &#x60;customerEndpointDNSName&#x60;. | [optional] 
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

### SetAzureLinkIdNil

`func (o *PrivateNetworkEndpointIdEntry) SetAzureLinkIdNil()`

SetAzureLinkIdNil sets AzureLinkId to an explicit JSON null when marshaled, overriding any value previously set with SetAzureLinkId. Calling SetAzureLinkId again clears the null override.

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

### SetCommentNil

`func (o *PrivateNetworkEndpointIdEntry) SetCommentNil()`

SetCommentNil sets Comment to an explicit JSON null when marshaled, overriding any value previously set with SetComment. Calling SetComment again clears the null override.

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

### SetCustomerEndpointDNSNameNil

`func (o *PrivateNetworkEndpointIdEntry) SetCustomerEndpointDNSNameNil()`

SetCustomerEndpointDNSNameNil sets CustomerEndpointDNSName to an explicit JSON null when marshaled, overriding any value previously set with SetCustomerEndpointDNSName. Calling SetCustomerEndpointDNSName again clears the null override.

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

### SetCustomerEndpointIPAddressNil

`func (o *PrivateNetworkEndpointIdEntry) SetCustomerEndpointIPAddressNil()`

SetCustomerEndpointIPAddressNil sets CustomerEndpointIPAddress to an explicit JSON null when marshaled, overriding any value previously set with SetCustomerEndpointIPAddress. Calling SetCustomerEndpointIPAddress again clears the null override.

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

### SetErrorMessageNil

`func (o *PrivateNetworkEndpointIdEntry) SetErrorMessageNil()`

SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled, overriding any value previously set with SetErrorMessage. Calling SetErrorMessage again clears the null override.

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

### SetProviderNil

`func (o *PrivateNetworkEndpointIdEntry) SetProviderNil()`

SetProviderNil sets Provider to an explicit JSON null when marshaled, overriding any value previously set with SetProvider. Calling SetProvider again clears the null override.

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

### SetRegionNil

`func (o *PrivateNetworkEndpointIdEntry) SetRegionNil()`

SetRegionNil sets Region to an explicit JSON null when marshaled, overriding any value previously set with SetRegion. Calling SetRegion again clears the null override.

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

### SetStatusNil

`func (o *PrivateNetworkEndpointIdEntry) SetStatusNil()`

SetStatusNil sets Status to an explicit JSON null when marshaled, overriding any value previously set with SetStatus. Calling SetStatus again clears the null override.

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

### SetTypeNil

`func (o *PrivateNetworkEndpointIdEntry) SetTypeNil()`

SetTypeNil sets Type to an explicit JSON null when marshaled, overriding any value previously set with SetType. Calling SetType again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


