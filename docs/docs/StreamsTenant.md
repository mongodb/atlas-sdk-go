# StreamsTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Connections** | Pointer to [**[]StreamsConnection**](StreamsConnection.md) | List of connections configured in the stream instance. | [optional] [readonly] 
**DataProcessRegion** | Pointer to [**StreamsDataProcessRegion**](StreamsDataProcessRegion.md) |  | [optional] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal character string that identifies the project. | [optional] [readonly] 
**Hostnames** | Pointer to **[]string** | List that contains the hostnames assigned to the stream instance. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the stream instance. | [optional] 

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


