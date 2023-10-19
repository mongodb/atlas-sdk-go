# GroupIPAddresses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the MongoDB Cloud project. | [optional] [readonly] 
**Services** | Pointer to [**GroupService**](GroupService.md) |  | [optional] 

## Methods

### NewGroupIPAddresses

`func NewGroupIPAddresses() *GroupIPAddresses`

NewGroupIPAddresses instantiates a new GroupIPAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupIPAddressesWithDefaults

`func NewGroupIPAddressesWithDefaults() *GroupIPAddresses`

NewGroupIPAddressesWithDefaults instantiates a new GroupIPAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupIPAddresses) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupIPAddresses) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupIPAddresses) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupIPAddresses) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.
### GetServices

`func (o *GroupIPAddresses) GetServices() GroupService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GroupIPAddresses) GetServicesOk() (*GroupService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GroupIPAddresses) SetServices(v GroupService)`

SetServices sets Services field to given value.

### HasServices

`func (o *GroupIPAddresses) HasServices() bool`

HasServices returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


