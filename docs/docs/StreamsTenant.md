# StreamsTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Connections** | Pointer to [**[]StreamsConnection**](StreamsConnection.md) | List of connections configured in the stream workspace. | [optional] [readonly] 
**DataProcessRegion** | Pointer to [**StreamsDataProcessRegion**](StreamsDataProcessRegion.md) |  | [optional] 
**FailoverRegions** | Pointer to [**[]StreamsDataProcessRegion**](StreamsDataProcessRegion.md) | List of failover regions configured for the stream workspace. | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Hostnames** | Pointer to **[]string** | List that contains the hostnames assigned to the stream workspace. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Label that identifies the stream workspace. | [optional] 
**SampleConnections** | Pointer to [**StreamsSampleConnections**](StreamsSampleConnections.md) |  | [optional] 
**StreamConfig** | Pointer to [**StreamConfig**](StreamConfig.md) |  | [optional] 

## Methods

### NewStreamsTenant

`func NewStreamsTenant() *StreamsTenant`

NewStreamsTenant instantiates a new StreamsTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamsTenantWithDefaults

`func NewStreamsTenantWithDefaults() *StreamsTenant`

NewStreamsTenantWithDefaults instantiates a new StreamsTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamsTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamsTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamsTenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamsTenant) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamsTenant) SetIdNil()`

SetIdNil sets Id to an explicit JSON null when marshaled, overriding any value previously set with SetId. Calling SetId again clears the null override.

### GetConnections

`func (o *StreamsTenant) GetConnections() []StreamsConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *StreamsTenant) GetConnectionsOk() (*[]StreamsConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *StreamsTenant) SetConnections(v []StreamsConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *StreamsTenant) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### SetConnectionsNil

`func (o *StreamsTenant) SetConnectionsNil()`

SetConnectionsNil sets Connections to an explicit JSON null when marshaled, overriding any value previously set with SetConnections. Calling SetConnections again clears the null override.

### GetDataProcessRegion

`func (o *StreamsTenant) GetDataProcessRegion() StreamsDataProcessRegion`

GetDataProcessRegion returns the DataProcessRegion field if non-nil, zero value otherwise.

### GetDataProcessRegionOk

`func (o *StreamsTenant) GetDataProcessRegionOk() (*StreamsDataProcessRegion, bool)`

GetDataProcessRegionOk returns a tuple with the DataProcessRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProcessRegion

`func (o *StreamsTenant) SetDataProcessRegion(v StreamsDataProcessRegion)`

SetDataProcessRegion sets DataProcessRegion field to given value.

### HasDataProcessRegion

`func (o *StreamsTenant) HasDataProcessRegion() bool`

HasDataProcessRegion returns a boolean if a field has been set.

### SetDataProcessRegionNil

`func (o *StreamsTenant) SetDataProcessRegionNil()`

SetDataProcessRegionNil sets DataProcessRegion to an explicit JSON null when marshaled, overriding any value previously set with SetDataProcessRegion. Calling SetDataProcessRegion again clears the null override.

### GetFailoverRegions

`func (o *StreamsTenant) GetFailoverRegions() []StreamsDataProcessRegion`

GetFailoverRegions returns the FailoverRegions field if non-nil, zero value otherwise.

### GetFailoverRegionsOk

`func (o *StreamsTenant) GetFailoverRegionsOk() (*[]StreamsDataProcessRegion, bool)`

GetFailoverRegionsOk returns a tuple with the FailoverRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverRegions

`func (o *StreamsTenant) SetFailoverRegions(v []StreamsDataProcessRegion)`

SetFailoverRegions sets FailoverRegions field to given value.

### HasFailoverRegions

`func (o *StreamsTenant) HasFailoverRegions() bool`

HasFailoverRegions returns a boolean if a field has been set.

### SetFailoverRegionsNil

`func (o *StreamsTenant) SetFailoverRegionsNil()`

SetFailoverRegionsNil sets FailoverRegions to an explicit JSON null when marshaled, overriding any value previously set with SetFailoverRegions. Calling SetFailoverRegions again clears the null override.

### GetGroupId

`func (o *StreamsTenant) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StreamsTenant) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StreamsTenant) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *StreamsTenant) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *StreamsTenant) SetGroupIdNil()`

SetGroupIdNil sets GroupId to an explicit JSON null when marshaled, overriding any value previously set with SetGroupId. Calling SetGroupId again clears the null override.

### GetHostnames

`func (o *StreamsTenant) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *StreamsTenant) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *StreamsTenant) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *StreamsTenant) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### SetHostnamesNil

`func (o *StreamsTenant) SetHostnamesNil()`

SetHostnamesNil sets Hostnames to an explicit JSON null when marshaled, overriding any value previously set with SetHostnames. Calling SetHostnames again clears the null override.

### GetLinks

`func (o *StreamsTenant) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StreamsTenant) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StreamsTenant) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StreamsTenant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *StreamsTenant) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.

### GetName

`func (o *StreamsTenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamsTenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamsTenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StreamsTenant) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StreamsTenant) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.

### GetSampleConnections

`func (o *StreamsTenant) GetSampleConnections() StreamsSampleConnections`

GetSampleConnections returns the SampleConnections field if non-nil, zero value otherwise.

### GetSampleConnectionsOk

`func (o *StreamsTenant) GetSampleConnectionsOk() (*StreamsSampleConnections, bool)`

GetSampleConnectionsOk returns a tuple with the SampleConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleConnections

`func (o *StreamsTenant) SetSampleConnections(v StreamsSampleConnections)`

SetSampleConnections sets SampleConnections field to given value.

### HasSampleConnections

`func (o *StreamsTenant) HasSampleConnections() bool`

HasSampleConnections returns a boolean if a field has been set.

### SetSampleConnectionsNil

`func (o *StreamsTenant) SetSampleConnectionsNil()`

SetSampleConnectionsNil sets SampleConnections to an explicit JSON null when marshaled, overriding any value previously set with SetSampleConnections. Calling SetSampleConnections again clears the null override.

### GetStreamConfig

`func (o *StreamsTenant) GetStreamConfig() StreamConfig`

GetStreamConfig returns the StreamConfig field if non-nil, zero value otherwise.

### GetStreamConfigOk

`func (o *StreamsTenant) GetStreamConfigOk() (*StreamConfig, bool)`

GetStreamConfigOk returns a tuple with the StreamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamConfig

`func (o *StreamsTenant) SetStreamConfig(v StreamConfig)`

SetStreamConfig sets StreamConfig field to given value.

### HasStreamConfig

`func (o *StreamsTenant) HasStreamConfig() bool`

HasStreamConfig returns a boolean if a field has been set.

### SetStreamConfigNil

`func (o *StreamsTenant) SetStreamConfigNil()`

SetStreamConfigNil sets StreamConfig to an explicit JSON null when marshaled, overriding any value previously set with SetStreamConfig. Calling SetStreamConfig again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


