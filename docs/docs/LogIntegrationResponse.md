# LogIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-character hexadecimal digit string that identifies the log integration configuration. | 
**Type** | **string** | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the log integration type. | 
**BucketName** | Pointer to **string** | Human-readable label that identifies the S3 bucket name for storing log files. | [optional] 
**IamRoleId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the AWS IAM role that MongoDB Cloud uses to access your S3 bucket. | [optional] 
**KmsKey** | Pointer to **string** | AWS KMS key ID or ARN for server-side encryption (optional). If not provided, uses bucket default encryption settings. | [optional] 
**LogTypes** | Pointer to **[]string** | Array of log types to export to S3. Valid values: MONGOD, MONGOS, MONGOD_AUDIT, MONGOS_AUDIT. | [optional] 
**PrefixPath** | Pointer to **string** | S3 directory path prefix where the log files will be stored. MongoDB Cloud will add further sub-directories based on the log type. | [optional] 

## Methods

### NewLogIntegrationResponse

`func NewLogIntegrationResponse(id string, type_ string, ) *LogIntegrationResponse`

NewLogIntegrationResponse instantiates a new LogIntegrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogIntegrationResponseWithDefaults

`func NewLogIntegrationResponseWithDefaults() *LogIntegrationResponse`

NewLogIntegrationResponseWithDefaults instantiates a new LogIntegrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LogIntegrationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogIntegrationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogIntegrationResponse) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *LogIntegrationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogIntegrationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogIntegrationResponse) SetType(v string)`

SetType sets Type field to given value.

### GetBucketName

`func (o *LogIntegrationResponse) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LogIntegrationResponse) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LogIntegrationResponse) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LogIntegrationResponse) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.
### GetIamRoleId

`func (o *LogIntegrationResponse) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *LogIntegrationResponse) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *LogIntegrationResponse) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *LogIntegrationResponse) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.
### GetKmsKey

`func (o *LogIntegrationResponse) GetKmsKey() string`

GetKmsKey returns the KmsKey field if non-nil, zero value otherwise.

### GetKmsKeyOk

`func (o *LogIntegrationResponse) GetKmsKeyOk() (*string, bool)`

GetKmsKeyOk returns a tuple with the KmsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKey

`func (o *LogIntegrationResponse) SetKmsKey(v string)`

SetKmsKey sets KmsKey field to given value.

### HasKmsKey

`func (o *LogIntegrationResponse) HasKmsKey() bool`

HasKmsKey returns a boolean if a field has been set.
### GetLogTypes

`func (o *LogIntegrationResponse) GetLogTypes() []string`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *LogIntegrationResponse) GetLogTypesOk() (*[]string, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *LogIntegrationResponse) SetLogTypes(v []string)`

SetLogTypes sets LogTypes field to given value.

### HasLogTypes

`func (o *LogIntegrationResponse) HasLogTypes() bool`

HasLogTypes returns a boolean if a field has been set.
### GetPrefixPath

`func (o *LogIntegrationResponse) GetPrefixPath() string`

GetPrefixPath returns the PrefixPath field if non-nil, zero value otherwise.

### GetPrefixPathOk

`func (o *LogIntegrationResponse) GetPrefixPathOk() (*string, bool)`

GetPrefixPathOk returns a tuple with the PrefixPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPath

`func (o *LogIntegrationResponse) SetPrefixPath(v string)`

SetPrefixPath sets PrefixPath field to given value.

### HasPrefixPath

`func (o *LogIntegrationResponse) HasPrefixPath() bool`

HasPrefixPath returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


