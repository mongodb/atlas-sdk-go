# EndpointService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | Cloud service provider that serves the requested endpoint service. | [readonly] 
**ErrorMessage** | Pointer to **string** | Error message returned when requesting private connection resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Private Endpoint Service. | [optional] [readonly] 
**RegionName** | Pointer to **string** | Cloud provider region that manages this Private Endpoint Service. | [optional] [readonly] 
**Status** | Pointer to **string** | State of the Private Endpoint Service connection when MongoDB Cloud received this request. | [optional] [readonly] 
**EndpointServiceName** | Pointer to **string** | Unique string that identifies the Amazon Web Services (AWS) PrivateLink endpoint service. MongoDB Cloud returns null while it creates the endpoint service. | [optional] [readonly] 
**InterfaceEndpoints** | Pointer to **[]string** | List of strings that identify private endpoint interfaces applied to the specified project. | [optional] [readonly] 
**SupportedRemoteRegions** | Pointer to **[]string** | List of regions that the endpoint service supports. Native cross region support is implemented for AWS only. | [optional] 
**PrivateEndpoints** | Pointer to **[]string** | List of private endpoints assigned to this Azure Private Link Service. | [optional] [readonly] 
**PrivateLinkServiceName** | Pointer to **string** | Unique string that identifies the Azure Private Link Service that MongoDB Cloud manages. | [optional] [readonly] 
**PrivateLinkServiceResourceId** | Pointer to **string** | Root-relative path that identifies of the Azure Private Link Service that MongoDB Cloud manages. Use this value to create a private endpoint connection to an Azure VNet. | [optional] [readonly] 
**EndpointGroupNames** | Pointer to **[]string** | List of Google Cloud network endpoint groups that corresponds to the Private Service Connect endpoint service. If this endpoint service uses PSC port-mapping, this field will only contain a list of one endpoint. | [optional] 
**PortMappingEnabled** | Pointer to **bool** | Flag that indicates whether this endpoint service uses PSC port-mapping. | [optional] 
**ServiceAttachmentNames** | Pointer to **[]string** | List of Uniform Resource Locators (URLs) that identifies endpoints that MongoDB Cloud can use to access one Google Cloud Service across a Google Cloud Virtual Private Connection (VPC) network. If this endpoint service uses PSC port-mapping, this field will only contain a list of one service attachment. | [optional] 

## Methods

### NewEndpointService

`func NewEndpointService(cloudProvider string, ) *EndpointService`

NewEndpointService instantiates a new EndpointService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointServiceWithDefaults

`func NewEndpointServiceWithDefaults() *EndpointService`

NewEndpointServiceWithDefaults instantiates a new EndpointService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *EndpointService) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EndpointService) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EndpointService) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### GetErrorMessage

`func (o *EndpointService) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EndpointService) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EndpointService) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EndpointService) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *EndpointService) SetErrorMessageNil()`

SetErrorMessageNil sets ErrorMessage to an explicit JSON null when marshaled, overriding any value previously set with SetErrorMessage. Calling SetErrorMessage again clears the null override.

### GetId

`func (o *EndpointService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointService) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *EndpointService) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetRegionName

`func (o *EndpointService) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *EndpointService) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *EndpointService) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### HasRegionName

`func (o *EndpointService) HasRegionName() bool`

HasRegionName returns a boolean if a field has been set.

### SetRegionNameNil

`func (o *EndpointService) SetRegionNameNil()`

SetRegionNameNil sets RegionName to an explicit JSON null when marshaled, overriding any value previously set with SetRegionName. Calling SetRegionName again clears the null override.

### GetStatus

`func (o *EndpointService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EndpointService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EndpointService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EndpointService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EndpointService) SetStatusNil()`

SetStatusNil sets Status to an explicit JSON null when marshaled, overriding any value previously set with SetStatus. Calling SetStatus again clears the null override.

### GetEndpointServiceName

`func (o *EndpointService) GetEndpointServiceName() string`

GetEndpointServiceName returns the EndpointServiceName field if non-nil, zero value otherwise.

### GetEndpointServiceNameOk

`func (o *EndpointService) GetEndpointServiceNameOk() (*string, bool)`

GetEndpointServiceNameOk returns a tuple with the EndpointServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointServiceName

`func (o *EndpointService) SetEndpointServiceName(v string)`

SetEndpointServiceName sets EndpointServiceName field to given value.

### HasEndpointServiceName

`func (o *EndpointService) HasEndpointServiceName() bool`

HasEndpointServiceName returns a boolean if a field has been set.

### SetEndpointServiceNameNil

`func (o *EndpointService) SetEndpointServiceNameNil()`

SetEndpointServiceNameNil sets EndpointServiceName to an explicit JSON null when marshaled, overriding any value previously set with SetEndpointServiceName. Calling SetEndpointServiceName again clears the null override.

### GetInterfaceEndpoints

`func (o *EndpointService) GetInterfaceEndpoints() []string`

GetInterfaceEndpoints returns the InterfaceEndpoints field if non-nil, zero value otherwise.

### GetInterfaceEndpointsOk

`func (o *EndpointService) GetInterfaceEndpointsOk() (*[]string, bool)`

GetInterfaceEndpointsOk returns a tuple with the InterfaceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceEndpoints

`func (o *EndpointService) SetInterfaceEndpoints(v []string)`

SetInterfaceEndpoints sets InterfaceEndpoints field to given value.

### HasInterfaceEndpoints

`func (o *EndpointService) HasInterfaceEndpoints() bool`

HasInterfaceEndpoints returns a boolean if a field has been set.

### SetInterfaceEndpointsNil

`func (o *EndpointService) SetInterfaceEndpointsNil()`

SetInterfaceEndpointsNil sets InterfaceEndpoints to an explicit JSON null when marshaled, overriding any value previously set with SetInterfaceEndpoints. Calling SetInterfaceEndpoints again clears the null override.

### GetSupportedRemoteRegions

`func (o *EndpointService) GetSupportedRemoteRegions() []string`

GetSupportedRemoteRegions returns the SupportedRemoteRegions field if non-nil, zero value otherwise.

### GetSupportedRemoteRegionsOk

`func (o *EndpointService) GetSupportedRemoteRegionsOk() (*[]string, bool)`

GetSupportedRemoteRegionsOk returns a tuple with the SupportedRemoteRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedRemoteRegions

`func (o *EndpointService) SetSupportedRemoteRegions(v []string)`

SetSupportedRemoteRegions sets SupportedRemoteRegions field to given value.

### HasSupportedRemoteRegions

`func (o *EndpointService) HasSupportedRemoteRegions() bool`

HasSupportedRemoteRegions returns a boolean if a field has been set.

### SetSupportedRemoteRegionsNil

`func (o *EndpointService) SetSupportedRemoteRegionsNil()`

SetSupportedRemoteRegionsNil sets SupportedRemoteRegions to an explicit JSON null when marshaled, overriding any value previously set with SetSupportedRemoteRegions. Calling SetSupportedRemoteRegions again clears the null override.

### GetPrivateEndpoints

`func (o *EndpointService) GetPrivateEndpoints() []string`

GetPrivateEndpoints returns the PrivateEndpoints field if non-nil, zero value otherwise.

### GetPrivateEndpointsOk

`func (o *EndpointService) GetPrivateEndpointsOk() (*[]string, bool)`

GetPrivateEndpointsOk returns a tuple with the PrivateEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateEndpoints

`func (o *EndpointService) SetPrivateEndpoints(v []string)`

SetPrivateEndpoints sets PrivateEndpoints field to given value.

### HasPrivateEndpoints

`func (o *EndpointService) HasPrivateEndpoints() bool`

HasPrivateEndpoints returns a boolean if a field has been set.

### SetPrivateEndpointsNil

`func (o *EndpointService) SetPrivateEndpointsNil()`

SetPrivateEndpointsNil sets PrivateEndpoints to an explicit JSON null when marshaled, overriding any value previously set with SetPrivateEndpoints. Calling SetPrivateEndpoints again clears the null override.

### GetPrivateLinkServiceName

`func (o *EndpointService) GetPrivateLinkServiceName() string`

GetPrivateLinkServiceName returns the PrivateLinkServiceName field if non-nil, zero value otherwise.

### GetPrivateLinkServiceNameOk

`func (o *EndpointService) GetPrivateLinkServiceNameOk() (*string, bool)`

GetPrivateLinkServiceNameOk returns a tuple with the PrivateLinkServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceName

`func (o *EndpointService) SetPrivateLinkServiceName(v string)`

SetPrivateLinkServiceName sets PrivateLinkServiceName field to given value.

### HasPrivateLinkServiceName

`func (o *EndpointService) HasPrivateLinkServiceName() bool`

HasPrivateLinkServiceName returns a boolean if a field has been set.

### SetPrivateLinkServiceNameNil

`func (o *EndpointService) SetPrivateLinkServiceNameNil()`

SetPrivateLinkServiceNameNil sets PrivateLinkServiceName to an explicit JSON null when marshaled, overriding any value previously set with SetPrivateLinkServiceName. Calling SetPrivateLinkServiceName again clears the null override.

### GetPrivateLinkServiceResourceId

`func (o *EndpointService) GetPrivateLinkServiceResourceId() string`

GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field if non-nil, zero value otherwise.

### GetPrivateLinkServiceResourceIdOk

`func (o *EndpointService) GetPrivateLinkServiceResourceIdOk() (*string, bool)`

GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLinkServiceResourceId

`func (o *EndpointService) SetPrivateLinkServiceResourceId(v string)`

SetPrivateLinkServiceResourceId sets PrivateLinkServiceResourceId field to given value.

### HasPrivateLinkServiceResourceId

`func (o *EndpointService) HasPrivateLinkServiceResourceId() bool`

HasPrivateLinkServiceResourceId returns a boolean if a field has been set.

### SetPrivateLinkServiceResourceIdNil

`func (o *EndpointService) SetPrivateLinkServiceResourceIdNil()`

SetPrivateLinkServiceResourceIdNil sets PrivateLinkServiceResourceId to an explicit JSON null when marshaled, overriding any value previously set with SetPrivateLinkServiceResourceId. Calling SetPrivateLinkServiceResourceId again clears the null override.

### GetEndpointGroupNames

`func (o *EndpointService) GetEndpointGroupNames() []string`

GetEndpointGroupNames returns the EndpointGroupNames field if non-nil, zero value otherwise.

### GetEndpointGroupNamesOk

`func (o *EndpointService) GetEndpointGroupNamesOk() (*[]string, bool)`

GetEndpointGroupNamesOk returns a tuple with the EndpointGroupNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroupNames

`func (o *EndpointService) SetEndpointGroupNames(v []string)`

SetEndpointGroupNames sets EndpointGroupNames field to given value.

### HasEndpointGroupNames

`func (o *EndpointService) HasEndpointGroupNames() bool`

HasEndpointGroupNames returns a boolean if a field has been set.

### SetEndpointGroupNamesNil

`func (o *EndpointService) SetEndpointGroupNamesNil()`

SetEndpointGroupNamesNil sets EndpointGroupNames to an explicit JSON null when marshaled, overriding any value previously set with SetEndpointGroupNames. Calling SetEndpointGroupNames again clears the null override.

### GetPortMappingEnabled

`func (o *EndpointService) GetPortMappingEnabled() bool`

GetPortMappingEnabled returns the PortMappingEnabled field if non-nil, zero value otherwise.

### GetPortMappingEnabledOk

`func (o *EndpointService) GetPortMappingEnabledOk() (*bool, bool)`

GetPortMappingEnabledOk returns a tuple with the PortMappingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMappingEnabled

`func (o *EndpointService) SetPortMappingEnabled(v bool)`

SetPortMappingEnabled sets PortMappingEnabled field to given value.

### HasPortMappingEnabled

`func (o *EndpointService) HasPortMappingEnabled() bool`

HasPortMappingEnabled returns a boolean if a field has been set.

### SetPortMappingEnabledNil

`func (o *EndpointService) SetPortMappingEnabledNil()`

SetPortMappingEnabledNil sets PortMappingEnabled to an explicit JSON null when marshaled, overriding any value previously set with SetPortMappingEnabled. Calling SetPortMappingEnabled again clears the null override.

### GetServiceAttachmentNames

`func (o *EndpointService) GetServiceAttachmentNames() []string`

GetServiceAttachmentNames returns the ServiceAttachmentNames field if non-nil, zero value otherwise.

### GetServiceAttachmentNamesOk

`func (o *EndpointService) GetServiceAttachmentNamesOk() (*[]string, bool)`

GetServiceAttachmentNamesOk returns a tuple with the ServiceAttachmentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachmentNames

`func (o *EndpointService) SetServiceAttachmentNames(v []string)`

SetServiceAttachmentNames sets ServiceAttachmentNames field to given value.

### HasServiceAttachmentNames

`func (o *EndpointService) HasServiceAttachmentNames() bool`

HasServiceAttachmentNames returns a boolean if a field has been set.

### SetServiceAttachmentNamesNil

`func (o *EndpointService) SetServiceAttachmentNamesNil()`

SetServiceAttachmentNamesNil sets ServiceAttachmentNames to an explicit JSON null when marshaled, overriding any value previously set with SetServiceAttachmentNames. Calling SetServiceAttachmentNames again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


