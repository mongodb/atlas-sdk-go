# AccountDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountId** | Pointer to **string** | The AWS Account ID. | [optional] 
**CidrBlock** | Pointer to **string** | The VPC CIDR Block. | [optional] 
**CloudProvider** | Pointer to **string** | Cloud provider. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**VpcId** | Pointer to **string** | The VPC ID. | [optional] 
**AzureSubscriptionId** | Pointer to **string** | The Azure Subscription ID. | [optional] 
**VirtualNetworkName** | Pointer to **string** | The name of the virtual network. | [optional] 
**GcpProjectId** | Pointer to **string** | The GCP Project ID. | [optional] 
**VpcNetworkName** | Pointer to **string** | The name of the VPC network. | [optional] 

## Methods

### NewAccountDetails

`func NewAccountDetails() *AccountDetails`

NewAccountDetails instantiates a new AccountDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountDetailsWithDefaults

`func NewAccountDetailsWithDefaults() *AccountDetails`

NewAccountDetailsWithDefaults instantiates a new AccountDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountId

`func (o *AccountDetails) GetAwsAccountId() string`

GetAwsAccountId returns the AwsAccountId field if non-nil, zero value otherwise.

### GetAwsAccountIdOk

`func (o *AccountDetails) GetAwsAccountIdOk() (*string, bool)`

GetAwsAccountIdOk returns a tuple with the AwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountId

`func (o *AccountDetails) SetAwsAccountId(v string)`

SetAwsAccountId sets AwsAccountId field to given value.

### HasAwsAccountId

`func (o *AccountDetails) HasAwsAccountId() bool`

HasAwsAccountId returns a boolean if a field has been set.
### GetCidrBlock

`func (o *AccountDetails) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AccountDetails) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AccountDetails) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AccountDetails) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.
### GetCloudProvider

`func (o *AccountDetails) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AccountDetails) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AccountDetails) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AccountDetails) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetLinks

`func (o *AccountDetails) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AccountDetails) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AccountDetails) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AccountDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetVpcId

`func (o *AccountDetails) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AccountDetails) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AccountDetails) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *AccountDetails) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.
### GetAzureSubscriptionId

`func (o *AccountDetails) GetAzureSubscriptionId() string`

GetAzureSubscriptionId returns the AzureSubscriptionId field if non-nil, zero value otherwise.

### GetAzureSubscriptionIdOk

`func (o *AccountDetails) GetAzureSubscriptionIdOk() (*string, bool)`

GetAzureSubscriptionIdOk returns a tuple with the AzureSubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionId

`func (o *AccountDetails) SetAzureSubscriptionId(v string)`

SetAzureSubscriptionId sets AzureSubscriptionId field to given value.

### HasAzureSubscriptionId

`func (o *AccountDetails) HasAzureSubscriptionId() bool`

HasAzureSubscriptionId returns a boolean if a field has been set.
### GetVirtualNetworkName

`func (o *AccountDetails) GetVirtualNetworkName() string`

GetVirtualNetworkName returns the VirtualNetworkName field if non-nil, zero value otherwise.

### GetVirtualNetworkNameOk

`func (o *AccountDetails) GetVirtualNetworkNameOk() (*string, bool)`

GetVirtualNetworkNameOk returns a tuple with the VirtualNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNetworkName

`func (o *AccountDetails) SetVirtualNetworkName(v string)`

SetVirtualNetworkName sets VirtualNetworkName field to given value.

### HasVirtualNetworkName

`func (o *AccountDetails) HasVirtualNetworkName() bool`

HasVirtualNetworkName returns a boolean if a field has been set.
### GetGcpProjectId

`func (o *AccountDetails) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *AccountDetails) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *AccountDetails) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.

### HasGcpProjectId

`func (o *AccountDetails) HasGcpProjectId() bool`

HasGcpProjectId returns a boolean if a field has been set.
### GetVpcNetworkName

`func (o *AccountDetails) GetVpcNetworkName() string`

GetVpcNetworkName returns the VpcNetworkName field if non-nil, zero value otherwise.

### GetVpcNetworkNameOk

`func (o *AccountDetails) GetVpcNetworkNameOk() (*string, bool)`

GetVpcNetworkNameOk returns a tuple with the VpcNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcNetworkName

`func (o *AccountDetails) SetVpcNetworkName(v string)`

SetVpcNetworkName sets VpcNetworkName field to given value.

### HasVpcNetworkName

`func (o *AccountDetails) HasVpcNetworkName() bool`

HasVpcNetworkName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


