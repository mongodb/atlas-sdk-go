# S3LogIntegrationRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | Human-readable label that identifies the S3 bucket name for storing log files. | [optional] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the AWS IAM role that MongoDB Cloud uses to access your S3 bucket. | [optional] 
**KmsKey** | Pointer to **string** | AWS KMS key ID or ARN for server-side encryption (optional). If not provided, uses bucket default encryption settings. | [optional] 
**LogTypes** | Pointer to **[]string** | Array of log types to export to S3. Valid values: MONGOD, MONGOS, MONGOD_AUDIT, MONGOS_AUDIT. | [optional] 
**PrefixPath** | Pointer to **string** | S3 directory path prefix where the log files will be stored. MongoDB Cloud will add further sub-directories based on the log type. | [optional] 
**Type** | Pointer to **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the log integration type. | [optional] 

## Methods

### NewS3LogIntegrationRequestAllOf

`func NewS3LogIntegrationRequestAllOf() *S3LogIntegrationRequestAllOf`

NewS3LogIntegrationRequestAllOf instantiates a new S3LogIntegrationRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3LogIntegrationRequestAllOfWithDefaults

`func NewS3LogIntegrationRequestAllOfWithDefaults() *S3LogIntegrationRequestAllOf`

NewS3LogIntegrationRequestAllOfWithDefaults instantiates a new S3LogIntegrationRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *S3LogIntegrationRequestAllOf) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *S3LogIntegrationRequestAllOf) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *S3LogIntegrationRequestAllOf) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *S3LogIntegrationRequestAllOf) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.
### GetIamRoleId

`func (o *S3LogIntegrationRequestAllOf) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *S3LogIntegrationRequestAllOf) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *S3LogIntegrationRequestAllOf) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *S3LogIntegrationRequestAllOf) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.
### GetKmsKey

`func (o *S3LogIntegrationRequestAllOf) GetKmsKey() string`

GetKmsKey returns the KmsKey field if non-nil, zero value otherwise.

### GetKmsKeyOk

`func (o *S3LogIntegrationRequestAllOf) GetKmsKeyOk() (*string, bool)`

GetKmsKeyOk returns a tuple with the KmsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKey

`func (o *S3LogIntegrationRequestAllOf) SetKmsKey(v string)`

SetKmsKey sets KmsKey field to given value.

### HasKmsKey

`func (o *S3LogIntegrationRequestAllOf) HasKmsKey() bool`

HasKmsKey returns a boolean if a field has been set.
### GetLogTypes

`func (o *S3LogIntegrationRequestAllOf) GetLogTypes() []string`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *S3LogIntegrationRequestAllOf) GetLogTypesOk() (*[]string, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *S3LogIntegrationRequestAllOf) SetLogTypes(v []string)`

SetLogTypes sets LogTypes field to given value.

### HasLogTypes

`func (o *S3LogIntegrationRequestAllOf) HasLogTypes() bool`

HasLogTypes returns a boolean if a field has been set.
### GetPrefixPath

`func (o *S3LogIntegrationRequestAllOf) GetPrefixPath() string`

GetPrefixPath returns the PrefixPath field if non-nil, zero value otherwise.

### GetPrefixPathOk

`func (o *S3LogIntegrationRequestAllOf) GetPrefixPathOk() (*string, bool)`

GetPrefixPathOk returns a tuple with the PrefixPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPath

`func (o *S3LogIntegrationRequestAllOf) SetPrefixPath(v string)`

SetPrefixPath sets PrefixPath field to given value.

### HasPrefixPath

`func (o *S3LogIntegrationRequestAllOf) HasPrefixPath() bool`

HasPrefixPath returns a boolean if a field has been set.
### GetType

`func (o *S3LogIntegrationRequestAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *S3LogIntegrationRequestAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *S3LogIntegrationRequestAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *S3LogIntegrationRequestAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


