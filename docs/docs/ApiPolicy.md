# ApiPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this backup policy. | [optional] 
**PolicyItems** | Pointer to [**[]ApiPolicyItem**](ApiPolicyItem.md) | List that contains the specifications for one policy. | [optional] 

## Methods

### NewApiPolicy

`func NewApiPolicy() *ApiPolicy`

NewApiPolicy instantiates a new ApiPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWithDefaults

`func NewApiPolicyWithDefaults() *ApiPolicy`

NewApiPolicyWithDefaults instantiates a new ApiPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPolicyItems

`func (o *ApiPolicy) GetPolicyItems() []ApiPolicyItem`

GetPolicyItems returns the PolicyItems field if non-nil, zero value otherwise.

### GetPolicyItemsOk

`func (o *ApiPolicy) GetPolicyItemsOk() (*[]ApiPolicyItem, bool)`

GetPolicyItemsOk returns a tuple with the PolicyItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyItems

`func (o *ApiPolicy) SetPolicyItems(v []ApiPolicyItem)`

SetPolicyItems sets PolicyItems field to given value.

### HasPolicyItems

`func (o *ApiPolicy) HasPolicyItems() bool`

HasPolicyItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


