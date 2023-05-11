# DBAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Human-readable label that identifies the privilege action. | 
**Resources** | Pointer to [**[]DBResource**](DBResource.md) | List of resources on which you grant the action. | [optional] 

## Methods

### NewDBAction

`func NewDBAction(action string, ) *DBAction`

NewDBAction instantiates a new DBAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBActionWithDefaults

`func NewDBActionWithDefaults() *DBAction`

NewDBActionWithDefaults instantiates a new DBAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *DBAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *DBAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *DBAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetResources

`func (o *DBAction) GetResources() []DBResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *DBAction) GetResourcesOk() (*[]DBResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *DBAction) SetResources(v []DBResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *DBAction) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


