# ApiUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Purpose or explanation provided when someone created this organization API key. | [optional] 
**Id** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies this organization API key assigned to this project. | [optional] [readonly] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**PrivateKey** | Pointer to **string** | Redacted private key returned for this organization API key. This key displays unredacted when first created. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public API key value set for the specified organization API key. | [optional] [readonly] 
**Roles** | Pointer to [**[]RoleAssignment**](RoleAssignment.md) | List that contains the roles that the API key needs to have. All roles you provide must be valid for the specified project or organization. Each request must include a minimum of one valid role. The resource returns all project and organization roles assigned to the API key. | [optional] 

## Methods

### NewApiUser

`func NewApiUser() *ApiUser`

NewApiUser instantiates a new ApiUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserWithDefaults

`func NewApiUserWithDefaults() *ApiUser`

NewApiUserWithDefaults instantiates a new ApiUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *ApiUser) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ApiUser) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ApiUser) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ApiUser) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetId

`func (o *ApiUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *ApiUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPrivateKey

`func (o *ApiUser) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ApiUser) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ApiUser) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ApiUser) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *ApiUser) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ApiUser) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ApiUser) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ApiUser) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetRoles

`func (o *ApiUser) GetRoles() []RoleAssignment`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiUser) GetRolesOk() (*[]RoleAssignment, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiUser) SetRoles(v []RoleAssignment)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


