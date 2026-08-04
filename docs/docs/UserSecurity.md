# UserSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerX509** | Pointer to [**DBUserTLSX509Settings**](DBUserTLSX509Settings.md) |  | [optional] 
**Ldap** | Pointer to [**LDAPSecuritySettings**](LDAPSecuritySettings.md) |  | [optional] 
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

`func (o *UserSecurity) GetCustomerX509() DBUserTLSX509Settings`

GetCustomerX509 returns the CustomerX509 field if non-nil, zero value otherwise.

### GetCustomerX509Ok

`func (o *UserSecurity) GetCustomerX509Ok() (*DBUserTLSX509Settings, bool)`

GetCustomerX509Ok returns a tuple with the CustomerX509 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerX509

`func (o *UserSecurity) SetCustomerX509(v DBUserTLSX509Settings)`

SetCustomerX509 sets CustomerX509 field to given value.

### HasCustomerX509

`func (o *UserSecurity) HasCustomerX509() bool`

HasCustomerX509 returns a boolean if a field has been set.

### SetCustomerX509Nil

`func (o *UserSecurity) SetCustomerX509Nil()`

SetCustomerX509Nil sets CustomerX509 to an explicit JSON null when marshaled, overriding any value previously set with SetCustomerX509. Calling SetCustomerX509 again clears the null override.

### GetLdap

`func (o *UserSecurity) GetLdap() LDAPSecuritySettings`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *UserSecurity) GetLdapOk() (*LDAPSecuritySettings, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *UserSecurity) SetLdap(v LDAPSecuritySettings)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *UserSecurity) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### SetLdapNil

`func (o *UserSecurity) SetLdapNil()`

SetLdapNil sets Ldap to an explicit JSON null when marshaled, overriding any value previously set with SetLdap. Calling SetLdap again clears the null override.

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

### SetLinksNil

`func (o *UserSecurity) SetLinksNil()`

SetLinksNil sets Links to an explicit JSON null when marshaled, overriding any value previously set with SetLinks. Calling SetLinks again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


