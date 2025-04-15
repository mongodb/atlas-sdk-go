# MdbAvailableVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | Cloud service provider on which MongoDB Cloud provisions the hosts. Set dedicated clusters to &#x60;AWS&#x60;, &#x60;GCP&#x60;, &#x60;AZURE&#x60; or &#x60;TENANT&#x60;. | [optional] 
**DefaultStatus** | Pointer to **string** | Whether the version is the current default for the Instance Size and Cloud Provider. | [optional] 
**InstanceSize** | Pointer to **string** | Instance size boundary to which your cluster can automatically scale. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Version** | Pointer to **string** | The MongoDB Major Version in question. | [optional] 

## Methods

### NewMdbAvailableVersion

`func NewMdbAvailableVersion() *MdbAvailableVersion`

NewMdbAvailableVersion instantiates a new MdbAvailableVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdbAvailableVersionWithDefaults

`func NewMdbAvailableVersionWithDefaults() *MdbAvailableVersion`

NewMdbAvailableVersionWithDefaults instantiates a new MdbAvailableVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *MdbAvailableVersion) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *MdbAvailableVersion) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *MdbAvailableVersion) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *MdbAvailableVersion) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.
### GetDefaultStatus

`func (o *MdbAvailableVersion) GetDefaultStatus() string`

GetDefaultStatus returns the DefaultStatus field if non-nil, zero value otherwise.

### GetDefaultStatusOk

`func (o *MdbAvailableVersion) GetDefaultStatusOk() (*string, bool)`

GetDefaultStatusOk returns a tuple with the DefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStatus

`func (o *MdbAvailableVersion) SetDefaultStatus(v string)`

SetDefaultStatus sets DefaultStatus field to given value.

### HasDefaultStatus

`func (o *MdbAvailableVersion) HasDefaultStatus() bool`

HasDefaultStatus returns a boolean if a field has been set.
### GetInstanceSize

`func (o *MdbAvailableVersion) GetInstanceSize() string`

GetInstanceSize returns the InstanceSize field if non-nil, zero value otherwise.

### GetInstanceSizeOk

`func (o *MdbAvailableVersion) GetInstanceSizeOk() (*string, bool)`

GetInstanceSizeOk returns a tuple with the InstanceSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSize

`func (o *MdbAvailableVersion) SetInstanceSize(v string)`

SetInstanceSize sets InstanceSize field to given value.

### HasInstanceSize

`func (o *MdbAvailableVersion) HasInstanceSize() bool`

HasInstanceSize returns a boolean if a field has been set.
### GetLinks

`func (o *MdbAvailableVersion) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MdbAvailableVersion) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MdbAvailableVersion) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MdbAvailableVersion) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetVersion

`func (o *MdbAvailableVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MdbAvailableVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MdbAvailableVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MdbAvailableVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


