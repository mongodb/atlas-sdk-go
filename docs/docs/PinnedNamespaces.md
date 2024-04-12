# PinnedNamespaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the request cluster. | [optional] [readonly] 
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the request project. | [optional] [readonly] 
**PinnedNamespaces** | Pointer to **[]string** | List of all pinned namespaces. | [optional] [readonly] 

## Methods

### NewPinnedNamespaces

`func NewPinnedNamespaces() *PinnedNamespaces`

NewPinnedNamespaces instantiates a new PinnedNamespaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedNamespacesWithDefaults

`func NewPinnedNamespacesWithDefaults() *PinnedNamespaces`

NewPinnedNamespacesWithDefaults instantiates a new PinnedNamespaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *PinnedNamespaces) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PinnedNamespaces) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PinnedNamespaces) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *PinnedNamespaces) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.
### GetGroupId

`func (o *PinnedNamespaces) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PinnedNamespaces) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PinnedNamespaces) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PinnedNamespaces) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetPinnedNamespaces

`func (o *PinnedNamespaces) GetPinnedNamespaces() []string`

GetPinnedNamespaces returns the PinnedNamespaces field if non-nil, zero value otherwise.

### GetPinnedNamespacesOk

`func (o *PinnedNamespaces) GetPinnedNamespacesOk() (*[]string, bool)`

GetPinnedNamespacesOk returns a tuple with the PinnedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedNamespaces

`func (o *PinnedNamespaces) SetPinnedNamespaces(v []string)`

SetPinnedNamespaces sets PinnedNamespaces field to given value.

### HasPinnedNamespaces

`func (o *PinnedNamespaces) HasPinnedNamespaces() bool`

HasPinnedNamespaces returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


