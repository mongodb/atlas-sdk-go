# CollStatsRankedNamespaces

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the request project. | [optional] [readonly] 
**IdentifierId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the request process. | [optional] [readonly] 
**RankedNamespaces** | Pointer to **[]string** | Ordered list of the hottest namespaces, highest value first. | [optional] [readonly] 

## Methods

### NewCollStatsRankedNamespaces

`func NewCollStatsRankedNamespaces() *CollStatsRankedNamespaces`

NewCollStatsRankedNamespaces instantiates a new CollStatsRankedNamespaces object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollStatsRankedNamespacesWithDefaults

`func NewCollStatsRankedNamespacesWithDefaults() *CollStatsRankedNamespaces`

NewCollStatsRankedNamespacesWithDefaults instantiates a new CollStatsRankedNamespaces object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *CollStatsRankedNamespaces) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CollStatsRankedNamespaces) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CollStatsRankedNamespaces) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *CollStatsRankedNamespaces) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetIdentifierId

`func (o *CollStatsRankedNamespaces) GetIdentifierId() string`

GetIdentifierId returns the IdentifierId field if non-nil, zero value otherwise.

### GetIdentifierIdOk

`func (o *CollStatsRankedNamespaces) GetIdentifierIdOk() (*string, bool)`

GetIdentifierIdOk returns a tuple with the IdentifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierId

`func (o *CollStatsRankedNamespaces) SetIdentifierId(v string)`

SetIdentifierId sets IdentifierId field to given value.

### HasIdentifierId

`func (o *CollStatsRankedNamespaces) HasIdentifierId() bool`

HasIdentifierId returns a boolean if a field has been set.
### GetRankedNamespaces

`func (o *CollStatsRankedNamespaces) GetRankedNamespaces() []string`

GetRankedNamespaces returns the RankedNamespaces field if non-nil, zero value otherwise.

### GetRankedNamespacesOk

`func (o *CollStatsRankedNamespaces) GetRankedNamespacesOk() (*[]string, bool)`

GetRankedNamespacesOk returns a tuple with the RankedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedNamespaces

`func (o *CollStatsRankedNamespaces) SetRankedNamespaces(v []string)`

SetRankedNamespaces sets RankedNamespaces field to given value.

### HasRankedNamespaces

`func (o *CollStatsRankedNamespaces) HasRankedNamespaces() bool`

HasRankedNamespaces returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


