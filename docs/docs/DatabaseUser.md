# DatabaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsIAMType** | Pointer to **string** | Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user&#39;s role. | [optional] [default to "NONE"]
**DatabaseName** | **string** | Database against which the database user authenticates. Database users must provide both a username and authentication database to log into MongoDB. | [default to "admin"]
**DeleteAfterDate** | Pointer to **time.Time** | Date and time when MongoDB Cloud deletes the user. This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request. | [optional] 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | 
**Labels** | Pointer to [**[]NDSLabel**](NDSLabel.md) | List that contains the key-value pairs for tagging and categorizing the MongoDB database user. The labels that you define do not appear in the console. | [optional] 
**LdapAuthType** | Pointer to **string** | Part of the Lightweight Directory Access Protocol (LDAP) record that the database uses to authenticate this database user on the LDAP host. | [optional] [default to "NONE"]
**Links** | Pointer to [**[]Link**](Link.md) | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships. | [optional] [readonly] 
**Password** | Pointer to **string** | Alphanumeric string that authenticates this database user against the database specified in &#x60;databaseName&#x60;. To authenticate with SCRAM-SHA, you must specify this parameter. This parameter doesn&#39;t appear in this response. | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) | List that provides the pairings of one role with one applicable database. | [optional] 
**Scopes** | Pointer to [**[]UserScope**](UserScope.md) | List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project. | [optional] 
**Username** | **string** | Human-readable label that represents the user that authenticates to MongoDB. The format of this label depends on the method of authentication:  | Authentication Method | Parameter Needed | Parameter Value | username Format | |---|---|---|---| | AWS IAM | awsType | ROLE | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | AWS IAM | awsType | USER | &lt;abbr title&#x3D;\&quot;Amazon Resource Name\&quot;&gt;ARN&lt;/abbr&gt; | | x.509 | x509Type | CUSTOMER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | x.509 | x509Type | MANAGED | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | USER | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | LDAP | ldapAuthType | GROUP | [RFC 2253](https://tools.ietf.org/html/2253) Distinguished Name | | SCRAM-SHA | awsType, x509Type, ldapAuthType | NONE | Alphanumeric string |  | 
**X509Type** | Pointer to **string** | X.509 method that MongoDB Cloud uses to authenticate the database user.  - For application-managed X.509, specify &#x60;MANAGED&#x60;. - For self-managed X.509, specify &#x60;CUSTOMER&#x60;.  Users created with the &#x60;CUSTOMER&#x60; method require a Common Name (CN) in the **username** parameter. You must create externally authenticated users on the &#x60;$external&#x60; database. | [optional] [default to "NONE"]

## Methods

### NewDatabaseUser

`func NewDatabaseUser(databaseName string, groupId string, username string, ) *DatabaseUser`

NewDatabaseUser instantiates a new DatabaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUserWithDefaults

`func NewDatabaseUserWithDefaults() *DatabaseUser`

NewDatabaseUserWithDefaults instantiates a new DatabaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsIAMType

`func (o *DatabaseUser) GetAwsIAMType() string`

GetAwsIAMType returns the AwsIAMType field if non-nil, zero value otherwise.

### GetAwsIAMTypeOk

`func (o *DatabaseUser) GetAwsIAMTypeOk() (*string, bool)`

GetAwsIAMTypeOk returns a tuple with the AwsIAMType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIAMType

`func (o *DatabaseUser) SetAwsIAMType(v string)`

SetAwsIAMType sets AwsIAMType field to given value.

### HasAwsIAMType

`func (o *DatabaseUser) HasAwsIAMType() bool`

HasAwsIAMType returns a boolean if a field has been set.

### GetDatabaseName

`func (o *DatabaseUser) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *DatabaseUser) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *DatabaseUser) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDeleteAfterDate

`func (o *DatabaseUser) GetDeleteAfterDate() time.Time`

GetDeleteAfterDate returns the DeleteAfterDate field if non-nil, zero value otherwise.

### GetDeleteAfterDateOk

`func (o *DatabaseUser) GetDeleteAfterDateOk() (*time.Time, bool)`

GetDeleteAfterDateOk returns a tuple with the DeleteAfterDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAfterDate

`func (o *DatabaseUser) SetDeleteAfterDate(v time.Time)`

SetDeleteAfterDate sets DeleteAfterDate field to given value.

### HasDeleteAfterDate

`func (o *DatabaseUser) HasDeleteAfterDate() bool`

HasDeleteAfterDate returns a boolean if a field has been set.

### GetGroupId

`func (o *DatabaseUser) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DatabaseUser) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DatabaseUser) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLabels

`func (o *DatabaseUser) GetLabels() []NDSLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *DatabaseUser) GetLabelsOk() (*[]NDSLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *DatabaseUser) SetLabels(v []NDSLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *DatabaseUser) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLdapAuthType

`func (o *DatabaseUser) GetLdapAuthType() string`

GetLdapAuthType returns the LdapAuthType field if non-nil, zero value otherwise.

### GetLdapAuthTypeOk

`func (o *DatabaseUser) GetLdapAuthTypeOk() (*string, bool)`

GetLdapAuthTypeOk returns a tuple with the LdapAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAuthType

`func (o *DatabaseUser) SetLdapAuthType(v string)`

SetLdapAuthType sets LdapAuthType field to given value.

### HasLdapAuthType

`func (o *DatabaseUser) HasLdapAuthType() bool`

HasLdapAuthType returns a boolean if a field has been set.

### GetLinks

`func (o *DatabaseUser) GetLinks() []Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DatabaseUser) GetLinksOk() (*[]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DatabaseUser) SetLinks(v []Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DatabaseUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPassword

`func (o *DatabaseUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DatabaseUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *DatabaseUser) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DatabaseUser) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DatabaseUser) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DatabaseUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScopes

`func (o *DatabaseUser) GetScopes() []UserScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DatabaseUser) GetScopesOk() (*[]UserScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DatabaseUser) SetScopes(v []UserScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DatabaseUser) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetUsername

`func (o *DatabaseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatabaseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatabaseUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetX509Type

`func (o *DatabaseUser) GetX509Type() string`

GetX509Type returns the X509Type field if non-nil, zero value otherwise.

### GetX509TypeOk

`func (o *DatabaseUser) GetX509TypeOk() (*string, bool)`

GetX509TypeOk returns a tuple with the X509Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509Type

`func (o *DatabaseUser) SetX509Type(v string)`

SetX509Type sets X509Type field to given value.

### HasX509Type

`func (o *DatabaseUser) HasX509Type() bool`

HasX509Type returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


