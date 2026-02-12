# CloudDatabaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIAMType** | Pointer to **string** | Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role. | [optional] [default to "NONE"]
**DatabaseName** | **string** | The database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. If the user authenticates with AWS IAM, x.509, LDAP, or OIDC Workload this value should be &#x60;$external&#x60;. If the user authenticates with SCRAM-SHA or OIDC Workforce, this value should be &#x60;admin&#x60;. | [default to "admin"]
**DeleteAfterDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request. | [optional] 
**Description** | Pointer to **string** | Description of this database user. | [optional] 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | 
**Labels** | Pointer to [**[]ComponentLabel**](ComponentLabel.md) | List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console. | [optional] 
**LdapAuthType** | Pointer to **string** | Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host. | [optional] [default to "NONE"]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**OidcAuthType** | Pointer to **string** | Human-readable label that indicates whether the new database user or group authenticates with OIDC federated authentication. To create a federated authentication user, specify the value of USER in this field. To create a federated authentication group, specify the value of &#x60;IDP_GROUP&#x60; in this field. | [optional] [default to "NONE"]
**Password** | Pointer to **string** | Alphanumeric string that authenticates this database user against the database specified in &#x60;databaseName&#x60;. To authenticate with SCRAM-SHA, you must specify this parameter. This parameter doesn&#39;t appear in this response. | [optional] 
**Roles** | Pointer to [**[]DatabaseUserRole**](DatabaseUserRole.md) | List that provides the pairings of one role with one applicable database. | [optional] 
**Scopes** | Pointer to [**[]UserScope**](UserScope.md) | List that contains clusters, MongoDB Atlas Data Lakes, and MongoDB Atlas Streams Workspaces that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters, MongoDB Atlas Data Lakes, and MongoDB Atlas Streams Workspaces in the project. | [optional] 
**Username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | &#x60;awsIAMType&#x60; | &#x60;ROLE&#x60; | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | &#x60;awsIAMType&#x60; | &#x60;USER&#x60; | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | &#x60;x509Type&#x60; | &#x60;CUSTOMER&#x60; | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | &#x60;x509Type&#x60; | &#x60;MANAGED&#x60; | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | &#x60;ldapAuthType&#x60; | &#x60;USER&#x60; | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | &#x60;ldapAuthType&#x60; | &#x60;GROUP&#x60; | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | OIDC Workforce | &#x60;oidcAuthType&#x60; | &#x60;IDP_GROUP&#x60; | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP group name | | OIDC Workload | &#x60;oidcAuthType&#x60; | &#x60;USER&#x60; | Atlas OIDC IdP ID (found in federation settings), followed by a &#39;/&#39;, followed by the IdP user name | | SCRAM-SHA | &#x60;awsIAMType&#x60;, &#x60;x509Type&#x60;, &#x60;ldapAuthType&#x60;, &#x60;oidcAuthType&#x60; | &#x60;NONE&#x60; | Alphanumeric string |  | 
**X509Type** | Pointer to **string** | X.509 method that MongoDB Cloud uses to authenticate the database user.  - For application-managed X.509, specify &#x60;MANAGED&#x60;. - For self-managed X.509, specify &#x60;CUSTOMER&#x60;.  Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database. | [optional] [default to "NONE"]

## Methods

### NewCloudDatabaseUser

`func NewCloudDatabaseUser(databaseName string, groupId string, username string, ) *CloudDatabaseUser`

NewCloudDatabaseUser instantiates a new CloudDatabaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudDatabaseUserWithDefaults

`func NewCloudDatabaseUserWithDefaults() *CloudDatabaseUser`

NewCloudDatabaseUserWithDefaults instantiates a new CloudDatabaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIAMType

`func (o *CloudDatabaseUser) GetAwsIAMType() string`

GetAwsIAMType returns the AwsIAMType field if non-nil, zero value otherwise.

### GetAwsIAMTypeOk

`func (o *CloudDatabaseUser) GetAwsIAMTypeOk() (*string, bool)`

GetAwsIAMTypeOk returns a tuple with the AwsIAMType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIAMType

`func (o *CloudDatabaseUser) SetAwsIAMType(v string)`

SetAwsIAMType sets AwsIAMType field to given value.

### HasAwsIAMType

`func (o *CloudDatabaseUser) HasAwsIAMType() bool`

HasAwsIAMType returns a boolean if a field has been set.
### GetDatabaseName

`func (o *CloudDatabaseUser) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *CloudDatabaseUser) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *CloudDatabaseUser) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### GetDeleteAfterDate

`func (o *CloudDatabaseUser) GetDeleteAfterDate() time.Time`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *CloudDatabaseUser) GetDeleteAfterDateOk() (*time.Time, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *CloudDatabaseUser) SetDeleteAfterDate(v time.Time)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *CloudDatabaseUser) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.
### GetDescription

`func (o *CloudDatabaseUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudDatabaseUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudDatabaseUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudDatabaseUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.
### GetGroupId

`func (o *CloudDatabaseUser) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CloudDatabaseUser) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CloudDatabaseUser) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### GetLabels

`func (o *CloudDatabaseUser) GetLabels() []ComponentLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CloudDatabaseUser) GetLabelsOk() (*[]ComponentLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CloudDatabaseUser) SetLabels(v []ComponentLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CloudDatabaseUser) HasLabels() bool`

HasLabels returns a boolean if a field has been set.
### GetLdapAuthType

`func (o *CloudDatabaseUser) GetLdapAuthType() string`

GetLdapAuthType returns the LdapAuthType field if non-nil, zero value otherwise.

### GetLdapAuthTypeOk

`func (o *CloudDatabaseUser) GetLdapAuthTypeOk() (*string, bool)`

GetLdapAuthTypeOk returns a tuple with the LdapAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAuthType

`func (o *CloudDatabaseUser) SetLdapAuthType(v string)`

SetLdapAuthType sets LdapAuthType field to given value.

### HasLdapAuthType

`func (o *CloudDatabaseUser) HasLdapAuthType() bool`

HasLdapAuthType returns a boolean if a field has been set.
### GetLinks

`func (o *CloudDatabaseUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CloudDatabaseUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CloudDatabaseUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CloudDatabaseUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.
### GetOidcAuthType

`func (o *CloudDatabaseUser) GetOidcAuthType() string`

GetOidcAuthType returns the OidcAuthType field if non-nil, zero value otherwise.

### GetOidcAuthTypeOk

`func (o *CloudDatabaseUser) GetOidcAuthTypeOk() (*string, bool)`

GetOidcAuthTypeOk returns a tuple with the OidcAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthType

`func (o *CloudDatabaseUser) SetOidcAuthType(v string)`

SetOidcAuthType sets OidcAuthType field to given value.

### HasOidcAuthType

`func (o *CloudDatabaseUser) HasOidcAuthType() bool`

HasOidcAuthType returns a boolean if a field has been set.
### GetPassword

`func (o *CloudDatabaseUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CloudDatabaseUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CloudDatabaseUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CloudDatabaseUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.
### GetRoles

`func (o *CloudDatabaseUser) GetRoles() []DatabaseUserRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CloudDatabaseUser) GetRolesOk() (*[]DatabaseUserRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CloudDatabaseUser) SetRoles(v []DatabaseUserRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CloudDatabaseUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.
### GetScopes

`func (o *CloudDatabaseUser) GetScopes() []UserScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CloudDatabaseUser) GetScopesOk() (*[]UserScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CloudDatabaseUser) SetScopes(v []UserScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CloudDatabaseUser) HasScopes() bool`

HasScopes returns a boolean if a field has been set.
### GetUsername

`func (o *CloudDatabaseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CloudDatabaseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CloudDatabaseUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### GetX509Type

`func (o *CloudDatabaseUser) GetX509Type() string`

GetX509Type returns the X509Type field if non-nil, zero value otherwise.

### GetX509TypeOk

`func (o *CloudDatabaseUser) GetX509TypeOk() (*string, bool)`

GetX509TypeOk returns a tuple with the X509Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Type

`func (o *CloudDatabaseUser) SetX509Type(v string)`

SetX509Type sets X509Type field to given value.

### HasX509Type

`func (o *CloudDatabaseUser) HasX509Type() bool`

HasX509Type returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


