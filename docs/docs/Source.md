# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificatePath** | Pointer to **string** | Path to the CA certificate that signed SSL certificates use to authenticate to the source cluster. | [optional] 
**ClusterName** | **string** | Label that identifies the source cluster name. | 
**GroupId** | **string** | Unique 24-hexadecimal digit string that identifies the source project. | 
**ManagedAuthentication** | **bool** | Flag that indicates whether MongoDB Automation manages authentication to the source cluster. If true, do not provide values for username and password. | 
**Password** | Pointer to **string** | Password that authenticates the username to the source cluster. | [optional] 
**Ssl** | **bool** | Flag that indicates whether you have SSL enabled. | 
**Username** | Pointer to **string** | Label that identifies the SCRAM-SHA user that connects to the source cluster. | [optional] 

## Methods

### NewSource

`func NewSource(clusterName string, groupId string, managedAuthentication bool, ssl bool, ) *Source`

NewSource instantiates a new Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceWithDefaults

`func NewSourceWithDefaults() *Source`

NewSourceWithDefaults instantiates a new Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificatePath

`func (o *Source) GetCaCertificatePath() string`

GetCaCertificatePath returns the CaCertificatePath field if non-nil, zero value otherwise.

### GetCaCertificatePathOk

`func (o *Source) GetCaCertificatePathOk() (*string, bool)`

GetCaCertificatePathOk returns a tuple with the CaCertificatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificatePath

`func (o *Source) SetCaCertificatePath(v string)`

SetCaCertificatePath sets CaCertificatePath field to given value.

### HasCaCertificatePath

`func (o *Source) HasCaCertificatePath() bool`

HasCaCertificatePath returns a boolean if a field has been set.

### GetClusterName

`func (o *Source) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Source) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Source) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetGroupId

`func (o *Source) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Source) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Source) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetManagedAuthentication

`func (o *Source) GetManagedAuthentication() bool`

GetManagedAuthentication returns the ManagedAuthentication field if non-nil, zero value otherwise.

### GetManagedAuthenticationOk

`func (o *Source) GetManagedAuthenticationOk() (*bool, bool)`

GetManagedAuthenticationOk returns a tuple with the ManagedAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedAuthentication

`func (o *Source) SetManagedAuthentication(v bool)`

SetManagedAuthentication sets ManagedAuthentication field to given value.


### GetPassword

`func (o *Source) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Source) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Source) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Source) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSsl

`func (o *Source) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *Source) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *Source) SetSsl(v bool)`

SetSsl sets Ssl field to given value.


### GetUsername

`func (o *Source) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Source) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Source) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Source) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


