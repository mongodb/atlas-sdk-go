# NDSLDAPVerifyConnectivityJobRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthzQueryTemplate** | Pointer to **string** | Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud applies to create an LDAP query to return the LDAP groups associated with the authenticated MongoDB user. MongoDB Cloud uses this parameter only for user authorization.  Use the &#x60;{USER}&#x60; placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query per [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516). | [optional] [default to "{USER}?memberOf?base"]
**BindPassword** | **string** | Password that MongoDB Cloud uses to authenticate the **bindUsername**. | 
**BindUsername** | **string** | Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. LDAP distinguished names must be formatted according to RFC 2253. | 
**CaCertificate** | Pointer to **string** | Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: &#x60;\&quot;caCertificate\&quot;: \&quot;\&quot;&#x60;. | [optional] 
**Hostname** | **string** | Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster. | 
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Port** | **int** | IANA port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections. | [default to 636]

## Methods

### NewNDSLDAPVerifyConnectivityJobRequestParams

`func NewNDSLDAPVerifyConnectivityJobRequestParams(bindPassword string, bindUsername string, hostname string, port int, ) *NDSLDAPVerifyConnectivityJobRequestParams`

NewNDSLDAPVerifyConnectivityJobRequestParams instantiates a new NDSLDAPVerifyConnectivityJobRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNDSLDAPVerifyConnectivityJobRequestParamsWithDefaults

`func NewNDSLDAPVerifyConnectivityJobRequestParamsWithDefaults() *NDSLDAPVerifyConnectivityJobRequestParams`

NewNDSLDAPVerifyConnectivityJobRequestParamsWithDefaults instantiates a new NDSLDAPVerifyConnectivityJobRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthzQueryTemplate

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetAuthzQueryTemplate() string`

GetAuthzQueryTemplate returns the AuthzQueryTemplate field if non-nil, zero value otherwise.

### GetAuthzQueryTemplateOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetAuthzQueryTemplateOk() (*string, bool)`

GetAuthzQueryTemplateOk returns a tuple with the AuthzQueryTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzQueryTemplate

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetAuthzQueryTemplate(v string)`

SetAuthzQueryTemplate sets AuthzQueryTemplate field to given value.

### HasAuthzQueryTemplate

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) HasAuthzQueryTemplate() bool`

HasAuthzQueryTemplate returns a boolean if a field has been set.

### GetBindPassword

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetBindUsername

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindUsername() string`

GetBindUsername returns the BindUsername field if non-nil, zero value otherwise.

### GetBindUsernameOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetBindUsernameOk() (*string, bool)`

GetBindUsernameOk returns a tuple with the BindUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindUsername

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetBindUsername(v string)`

SetBindUsername sets BindUsername field to given value.


### GetCaCertificate

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetHostname

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetLinks

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPort

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetPort() int`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) GetPortOk() (*int, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NDSLDAPVerifyConnectivityJobRequestParams) SetPort(v int)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


