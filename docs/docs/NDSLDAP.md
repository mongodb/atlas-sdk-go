# NDSLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationEnabled** | Pointer to **bool** | Flag that indicates whether users can authenticate using an Lightweight Directory Access Protocol (LDAP) host. | [optional] 
**AuthorizationEnabled** | Pointer to **bool** | Flag that indicates whether users can authorize access to MongoDB Cloud resources using an Lightweight Directory Access Protocol (LDAP) host. | [optional] 
**AuthzQueryTemplate** | Pointer to **string** | Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user. MongoDB Cloud uses this parameter only for user authorization. Use the &#x60;{USER}&#x60; placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query according to [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516). | [optional] [default to "{USER}?memberOf?base"]
**BindPassword** | Pointer to **string** | Password that MongoDB Cloud uses to authenticate the **bindUsername**. | [optional] 
**BindUsername** | Pointer to **string** | Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. LDAP distinguished names must be formatted according to RFC 2253. | [optional] 
**CaCertificate** | Pointer to **string** | Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: &#x60;\&quot;caCertificate\&quot;: \&quot;\&quot;&#x60;. | [optional] 
**Hostname** | Pointer to **string** | Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster. | [optional] 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Port** | Pointer to **int** | Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections. | [optional] [default to 636]
**UserToDNMapping** | Pointer to [**[]NDSUserToDNMapping**](NDSUserToDNMapping.md) | User-to-Distinguished Name (DN) map that MongoDB Cloud uses to transform a Lightweight Directory Access Protocol (LDAP) username into an LDAP DN. | [optional] 

## Methods

### NewNDSLDAP

`func NewNDSLDAP() *NDSLDAP`

NewNDSLDAP instantiates a new NDSLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNDSLDAPWithDefaults

`func NewNDSLDAPWithDefaults() *NDSLDAP`

NewNDSLDAPWithDefaults instantiates a new NDSLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationEnabled

`func (o *NDSLDAP) GetAuthenticationEnabled() bool`

GetAuthenticationEnabled returns the AuthenticationEnabled field if non-nil, zero value otherwise.

### GetAuthenticationEnabledOk

`func (o *NDSLDAP) GetAuthenticationEnabledOk() (*bool, bool)`

GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationEnabled

`func (o *NDSLDAP) SetAuthenticationEnabled(v bool)`

SetAuthenticationEnabled sets AuthenticationEnabled field to given value.

### HasAuthenticationEnabled

`func (o *NDSLDAP) HasAuthenticationEnabled() bool`

HasAuthenticationEnabled returns a boolean if a field has been set.

### GetAuthorizationEnabled

`func (o *NDSLDAP) GetAuthorizationEnabled() bool`

GetAuthorizationEnabled returns the AuthorizationEnabled field if non-nil, zero value otherwise.

### GetAuthorizationEnabledOk

`func (o *NDSLDAP) GetAuthorizationEnabledOk() (*bool, bool)`

GetAuthorizationEnabledOk returns a tuple with the AuthorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEnabled

`func (o *NDSLDAP) SetAuthorizationEnabled(v bool)`

SetAuthorizationEnabled sets AuthorizationEnabled field to given value.

### HasAuthorizationEnabled

`func (o *NDSLDAP) HasAuthorizationEnabled() bool`

HasAuthorizationEnabled returns a boolean if a field has been set.

### GetAuthzQueryTemplate

`func (o *NDSLDAP) GetAuthzQueryTemplate() string`

GetAuthzQueryTemplate returns the AuthzQueryTemplate field if non-nil, zero value otherwise.

### GetAuthzQueryTemplateOk

`func (o *NDSLDAP) GetAuthzQueryTemplateOk() (*string, bool)`

GetAuthzQueryTemplateOk returns a tuple with the AuthzQueryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzQueryTemplate

`func (o *NDSLDAP) SetAuthzQueryTemplate(v string)`

SetAuthzQueryTemplate sets AuthzQueryTemplate field to given value.

### HasAuthzQueryTemplate

`func (o *NDSLDAP) HasAuthzQueryTemplate() bool`

HasAuthzQueryTemplate returns a boolean if a field has been set.

### GetBindPassword

`func (o *NDSLDAP) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *NDSLDAP) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *NDSLDAP) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.

### HasBindPassword

`func (o *NDSLDAP) HasBindPassword() bool`

HasBindPassword returns a boolean if a field has been set.

### GetBindUsername

`func (o *NDSLDAP) GetBindUsername() string`

GetBindUsername returns the BindUsername field if non-nil, zero value otherwise.

### GetBindUsernameOk

`func (o *NDSLDAP) GetBindUsernameOk() (*string, bool)`

GetBindUsernameOk returns a tuple with the BindUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindUsername

`func (o *NDSLDAP) SetBindUsername(v string)`

SetBindUsername sets BindUsername field to given value.

### HasBindUsername

`func (o *NDSLDAP) HasBindUsername() bool`

HasBindUsername returns a boolean if a field has been set.

### GetCaCertificate

`func (o *NDSLDAP) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *NDSLDAP) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *NDSLDAP) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *NDSLDAP) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetHostname

`func (o *NDSLDAP) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NDSLDAP) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NDSLDAP) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *NDSLDAP) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLinks

`func (o *NDSLDAP) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NDSLDAP) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NDSLDAP) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NDSLDAP) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPort

`func (o *NDSLDAP) GetPort() int`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NDSLDAP) GetPortOk() (*int, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NDSLDAP) SetPort(v int)`

SetPort sets Port field to given value.

### HasPort

`func (o *NDSLDAP) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUserToDNMapping

`func (o *NDSLDAP) GetUserToDNMapping() []NDSUserToDNMapping`

GetUserToDNMapping returns the UserToDNMapping field if non-nil, zero value otherwise.

### GetUserToDNMappingOk

`func (o *NDSLDAP) GetUserToDNMappingOk() (*[]NDSUserToDNMapping, bool)`

GetUserToDNMappingOk returns a tuple with the UserToDNMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserToDNMapping

`func (o *NDSLDAP) SetUserToDNMapping(v []NDSUserToDNMapping)`

SetUserToDNMapping sets UserToDNMapping field to given value.

### HasUserToDNMapping

`func (o *NDSLDAP) HasUserToDNMapping() bool`

HasUserToDNMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


