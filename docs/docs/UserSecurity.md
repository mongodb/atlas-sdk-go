# UserSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerX509** | Pointer to [**CustomerX509**](CustomerX509.md) |  | [optional] 
**Ldap** | Pointer to [**NDSLDAP**](NDSLDAP.md) |  | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 

## Methods

### NewUserSecurity

`func NewUserSecurity() *UserSecurity`

NewUserSecurity instantiates a new UserSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSecurityWithDefaults

`func NewUserSecurityWithDefaults() *UserSecurity`

NewUserSecurityWithDefaults instantiates a new UserSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerX509

`func (o *UserSecurity) GetCustomerX509() CustomerX509`

GetCustomerX509 returns the CustomerX509 field if non-nil, zero value otherwise.

### GetCustomerX509Ok

`func (o *UserSecurity) GetCustomerX509Ok() (*CustomerX509, bool)`

GetCustomerX509Ok returns a tuple with the CustomerX509 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerX509

`func (o *UserSecurity) SetCustomerX509(v CustomerX509)`

SetCustomerX509 sets CustomerX509 field to given value.

### HasCustomerX509

`func (o *UserSecurity) HasCustomerX509() bool`

HasCustomerX509 returns a boolean if a field has been set.

### GetLdap

`func (o *UserSecurity) GetLdap() NDSLDAP`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *UserSecurity) GetLdapOk() (*NDSLDAP, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *UserSecurity) SetLdap(v NDSLDAP)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *UserSecurity) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetLinks

`func (o *UserSecurity) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserSecurity) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserSecurity) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserSecurity) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


