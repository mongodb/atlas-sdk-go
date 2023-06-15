# EncryptionAtRest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKms** | Pointer to [**AWSKMSConfiguration**](AWSKMSConfiguration.md) |  | [optional] 
**AzureKeyVault** | Pointer to [**AzureKeyVault**](AzureKeyVault.md) |  | [optional] 
**GoogleCloudKms** | Pointer to [**GoogleCloudKMS**](GoogleCloudKMS.md) |  | [optional] 

## Methods

### NewEncryptionAtRest

`func NewEncryptionAtRest() *EncryptionAtRest`

NewEncryptionAtRest instantiates a new EncryptionAtRest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptionAtRestWithDefaults

`func NewEncryptionAtRestWithDefaults() *EncryptionAtRest`

NewEncryptionAtRestWithDefaults instantiates a new EncryptionAtRest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKms

`func (o *EncryptionAtRest) GetAwsKms() AWSKMSConfiguration`

GetAwsKms returns the AwsKms field if non-nil, zero value otherwise.

### GetAwsKmsOk

`func (o *EncryptionAtRest) GetAwsKmsOk() (*AWSKMSConfiguration, bool)`

GetAwsKmsOk returns a tuple with the AwsKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKms

`func (o *EncryptionAtRest) SetAwsKms(v AWSKMSConfiguration)`

SetAwsKms sets AwsKms field to given value.

### HasAwsKms

`func (o *EncryptionAtRest) HasAwsKms() bool`

HasAwsKms returns a boolean if a field has been set.

### GetAzureKeyVault

`func (o *EncryptionAtRest) GetAzureKeyVault() AzureKeyVault`

GetAzureKeyVault returns the AzureKeyVault field if non-nil, zero value otherwise.

### GetAzureKeyVaultOk

`func (o *EncryptionAtRest) GetAzureKeyVaultOk() (*AzureKeyVault, bool)`

GetAzureKeyVaultOk returns a tuple with the AzureKeyVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKeyVault

`func (o *EncryptionAtRest) SetAzureKeyVault(v AzureKeyVault)`

SetAzureKeyVault sets AzureKeyVault field to given value.

### HasAzureKeyVault

`func (o *EncryptionAtRest) HasAzureKeyVault() bool`

HasAzureKeyVault returns a boolean if a field has been set.

### GetGoogleCloudKms

`func (o *EncryptionAtRest) GetGoogleCloudKms() GoogleCloudKMS`

GetGoogleCloudKms returns the GoogleCloudKms field if non-nil, zero value otherwise.

### GetGoogleCloudKmsOk

`func (o *EncryptionAtRest) GetGoogleCloudKmsOk() (*GoogleCloudKMS, bool)`

GetGoogleCloudKmsOk returns a tuple with the GoogleCloudKms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleCloudKms

`func (o *EncryptionAtRest) SetGoogleCloudKms(v GoogleCloudKMS)`

SetGoogleCloudKms sets GoogleCloudKms field to given value.

### HasGoogleCloudKms

`func (o *EncryptionAtRest) HasGoogleCloudKms() bool`

HasGoogleCloudKms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


