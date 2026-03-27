# AzureConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Region** | Pointer to **string** | Azure region where the storage account is located. | [optional] 
**ServicePrincipalId** | Pointer to **string** | Unique ID of the Azure Service Principal that has access to the storage account. | [optional] 
**StorageAccountName** | Pointer to **string** | Name of the Azure Storage Account to connect to. | [optional] 

## Methods

### NewAzureConnection

`func NewAzureConnection() *AzureConnection`

NewAzureConnection instantiates a new AzureConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureConnectionWithDefaults

`func NewAzureConnectionWithDefaults() *AzureConnection`

NewAzureConnectionWithDefaults instantiates a new AzureConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AzureConnection) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AzureConnection) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AzureConnection) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AzureConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRegion

`func (o *AzureConnection) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AzureConnection) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AzureConnection) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AzureConnection) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetServicePrincipalId

`func (o *AzureConnection) GetServicePrincipalId() string`

GetServicePrincipalId returns the ServicePrincipalId field if non-nil, zero value otherwise.

### GetServicePrincipalIdOk

`func (o *AzureConnection) GetServicePrincipalIdOk() (*string, bool)`

GetServicePrincipalIdOk returns a tuple with the ServicePrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipalId

`func (o *AzureConnection) SetServicePrincipalId(v string)`

SetServicePrincipalId sets ServicePrincipalId field to given value.

### HasServicePrincipalId

`func (o *AzureConnection) HasServicePrincipalId() bool`

HasServicePrincipalId returns a boolean if a field has been set.
### GetStorageAccountName

`func (o *AzureConnection) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *AzureConnection) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *AzureConnection) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *AzureConnection) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


