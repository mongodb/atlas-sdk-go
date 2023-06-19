# GCPNetworkPeeringConnectionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud network container that contains the specified network peering connection. | 
**ErrorMessage** | Pointer to **string** | Details of the error returned when requesting a GCP network peering resource. The resource returns &#x60;null&#x60; if the request succeeded. | [optional] [readonly] 
**GcpProjectId** | **string** | Human-readable label that identifies the GCP project that contains the network that you want to peer with the MongoDB Cloud VPC. | 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the network peering connection. | [optional] [readonly] 
**NetworkName** | **string** | Human-readable label that identifies the network to peer with the MongoDB Cloud VPC. | 
**ProviderName** | Pointer to **string** | Cloud service provider that serves the requested network peering connection. | [optional] 
**Status** | Pointer to **string** | State of the network peering connection at the time you made the request. | [optional] [readonly] 

## Methods

### NewGCPNetworkPeeringConnectionSettings

`func NewGCPNetworkPeeringConnectionSettings(containerId string, gcpProjectId string, networkName string, ) *GCPNetworkPeeringConnectionSettings`

NewGCPNetworkPeeringConnectionSettings instantiates a new GCPNetworkPeeringConnectionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGCPNetworkPeeringConnectionSettingsWithDefaults

`func NewGCPNetworkPeeringConnectionSettingsWithDefaults() *GCPNetworkPeeringConnectionSettings`

NewGCPNetworkPeeringConnectionSettingsWithDefaults instantiates a new GCPNetworkPeeringConnectionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *GCPNetworkPeeringConnectionSettings) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GCPNetworkPeeringConnectionSettings) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GCPNetworkPeeringConnectionSettings) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetErrorMessage

`func (o *GCPNetworkPeeringConnectionSettings) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GCPNetworkPeeringConnectionSettings) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GCPNetworkPeeringConnectionSettings) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GCPNetworkPeeringConnectionSettings) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGcpProjectId

`func (o *GCPNetworkPeeringConnectionSettings) GetGcpProjectId() string`

GetGcpProjectId returns the GcpProjectId field if non-nil, zero value otherwise.

### GetGcpProjectIdOk

`func (o *GCPNetworkPeeringConnectionSettings) GetGcpProjectIdOk() (*string, bool)`

GetGcpProjectIdOk returns a tuple with the GcpProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectId

`func (o *GCPNetworkPeeringConnectionSettings) SetGcpProjectId(v string)`

SetGcpProjectId sets GcpProjectId field to given value.


### GetId

`func (o *GCPNetworkPeeringConnectionSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GCPNetworkPeeringConnectionSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GCPNetworkPeeringConnectionSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GCPNetworkPeeringConnectionSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkName

`func (o *GCPNetworkPeeringConnectionSettings) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *GCPNetworkPeeringConnectionSettings) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *GCPNetworkPeeringConnectionSettings) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.


### GetProviderName

`func (o *GCPNetworkPeeringConnectionSettings) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *GCPNetworkPeeringConnectionSettings) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *GCPNetworkPeeringConnectionSettings) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *GCPNetworkPeeringConnectionSettings) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetStatus

`func (o *GCPNetworkPeeringConnectionSettings) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GCPNetworkPeeringConnectionSettings) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GCPNetworkPeeringConnectionSettings) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GCPNetworkPeeringConnectionSettings) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


