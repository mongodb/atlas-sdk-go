# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessList** | Pointer to [**[]AccessListItem**](AccessListItem.md) | List of network addresses granted access to this API using this API key. | [optional] [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies this organization API key. | [readonly] 
**PublicKey** | **string** | Public API key value set for the specified organization API key. | [readonly] 
**Roles** | Pointer to [**[]RoleAssignment**](RoleAssignment.md) | List that contains roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the Cloud user. | [optional] [readonly] 

## Methods

### NewKey

`func NewKey(id string, publicKey string, ) *Key`

NewKey instantiates a new Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyWithDefaults

`func NewKeyWithDefaults() *Key`

NewKeyWithDefaults instantiates a new Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *Key) GetAccessList() []AccessListItem`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *Key) GetAccessListOk() (*[]AccessListItem, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *Key) SetAccessList(v []AccessListItem)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *Key) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetId

`func (o *Key) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Key) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Key) SetId(v string)`

SetId sets Id field to given value.


### GetPublicKey

`func (o *Key) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Key) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Key) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetRoles

`func (o *Key) GetRoles() []RoleAssignment`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Key) GetRolesOk() (*[]RoleAssignment, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Key) SetRoles(v []RoleAssignment)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Key) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


