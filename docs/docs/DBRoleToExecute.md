# DBRoleToExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Role** | Pointer to **string** | The name of the role to use. Can be a built in role or a custom role. | [optional] 
**Type** | Pointer to **string** | Type of the DB role. Can be either Built In or Custom. | [optional] 

## Methods

### NewDBRoleToExecute

`func NewDBRoleToExecute() *DBRoleToExecute`

NewDBRoleToExecute instantiates a new DBRoleToExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBRoleToExecuteWithDefaults

`func NewDBRoleToExecuteWithDefaults() *DBRoleToExecute`

NewDBRoleToExecuteWithDefaults instantiates a new DBRoleToExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DBRoleToExecute) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DBRoleToExecute) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DBRoleToExecute) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DBRoleToExecute) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetRole

`func (o *DBRoleToExecute) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DBRoleToExecute) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DBRoleToExecute) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DBRoleToExecute) HasRole() bool`

HasRole returns a boolean if a field has been set.
### GetType

`func (o *DBRoleToExecute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DBRoleToExecute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DBRoleToExecute) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DBRoleToExecute) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


