# KeyUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Purpose or explanation provided when someone created this organization API key. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**PrivateKey** | Pointer to **string** | Redacted private key returned for this organization API key. This key displays unredacted when first created. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public API key value set for the specified organization API key. | [optional] [readonly] 
**Roles** | Pointer to [**[]CloudRoleAssignment**](CloudRoleAssignment.md) | List that contains the roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the API key. | [optional] 

## Methods

### NewKeyUser

`func NewKeyUser() *KeyUser`

NewKeyUser instantiates a new KeyUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyUserWithDefaults

`func NewKeyUserWithDefaults() *KeyUser`

NewKeyUserWithDefaults instantiates a new KeyUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *KeyUser) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *KeyUser) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *KeyUser) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *KeyUser) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetId

`func (o *KeyUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *KeyUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *KeyUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *KeyUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *KeyUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KeyUser) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeyUser) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeyUser) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeyUser) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *KeyUser) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KeyUser) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KeyUser) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KeyUser) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRoles

`func (o *KeyUser) GetRoles() []CloudRoleAssignment`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *KeyUser) GetRolesOk() (*[]CloudRoleAssignment, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *KeyUser) SetRoles(v []CloudRoleAssignment)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *KeyUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


