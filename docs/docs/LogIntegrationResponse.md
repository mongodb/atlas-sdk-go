# LogIntegrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique 24-character hexadecimal digit string that identifies the log integration configuration. | 
**LogTypes** | **[]string** | Array of log types exported by this integration. | 
**Type** | **string** | Human-readable label that identifies the service to which you want to integrate with Atlas. The value must match the log integration type. This value cannot be modified after the integration is created. | 
**BucketName** | Pointer to **string** | Name of the bucket to store log files. | [optional] 
**IamRoleId** | Pointer to **string** | Unique 24-character hexadecimal string that identifies the AWS IAM role that Atlas uses to access the S3 bucket. | [optional] 
**KmsKey** | Pointer to **string** | AWS KMS key ID or ARN for server-side encryption (optional). If not provided, uses bucket default encryption settings. | [optional] 
**PrefixPath** | Pointer to **string** | Path prefix where the log files will be stored. Atlas will add further sub-directories based on the log type. | [optional] 
**UseLegacyPathStructure** | Pointer to **bool** | When true, uses the legacy daily-folder path structure compatible with Push-Based Log Export: &#x60;{prefix}/{cluster}/{hostname}/{logType}/{YYYY-MM-DD}/{timestamp}-{logType}.log&#x60;. When false (default), uses the flat timestamped structure: &#x60;{prefix}/{cluster}/{hostname}/{logType}/{timestamp}-{logType}.log&#x60;. | [optional] 
**ApiKey** | Pointer to **string** | API key for authentication. | [optional] 
**Region** | Pointer to **string** | Datadog site/region for log ingestion. Valid values: US1, US3, US5, EU, AP1, AP2, US1_FED. | [optional] 
**RoleId** | Pointer to **string** | Unique 24-character hexadecimal string that identifies the Atlas Cloud Provider Access role. | [optional] 
**OtelEndpoint** | Pointer to **string** | OpenTelemetry collector endpoint URL. | [optional] 
**OtelSuppliedHeaders** | Pointer to [**[]Header**](Header.md) | HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB. | [optional] 
**HecToken** | Pointer to **string** | HTTP Event Collector (HEC) token for authentication. | [optional] 
**HecUrl** | Pointer to **string** | HTTP Event Collector (HEC) endpoint URL. | [optional] 
**StorageAccountName** | Pointer to **string** | Storage account name where logs will be stored. | [optional] 
**StorageContainerName** | Pointer to **string** | Storage container name for log files. | [optional] 

## Methods

### NewLogIntegrationResponse

`func NewLogIntegrationResponse(id string, logTypes []string, type_ string, ) *LogIntegrationResponse`

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
### GetUseLegacyPathStructure

`func (o *LogIntegrationResponse) GetUseLegacyPathStructure() bool`

GetUseLegacyPathStructure returns the UseLegacyPathStructure field if non-nil, zero value otherwise.

### GetUseLegacyPathStructureOk

`func (o *LogIntegrationResponse) GetUseLegacyPathStructureOk() (*bool, bool)`

GetUseLegacyPathStructureOk returns a tuple with the UseLegacyPathStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyPathStructure

`func (o *LogIntegrationResponse) SetUseLegacyPathStructure(v bool)`

SetUseLegacyPathStructure sets UseLegacyPathStructure field to given value.

### HasUseLegacyPathStructure

`func (o *LogIntegrationResponse) HasUseLegacyPathStructure() bool`

HasUseLegacyPathStructure returns a boolean if a field has been set.
### GetApiKey

`func (o *LogIntegrationResponse) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *LogIntegrationResponse) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *LogIntegrationResponse) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *LogIntegrationResponse) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.
### GetRegion

`func (o *LogIntegrationResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogIntegrationResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogIntegrationResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogIntegrationResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.
### GetRoleId

`func (o *LogIntegrationResponse) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *LogIntegrationResponse) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *LogIntegrationResponse) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *LogIntegrationResponse) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.
### GetOtelEndpoint

`func (o *LogIntegrationResponse) GetOtelEndpoint() string`

GetOtelEndpoint returns the OtelEndpoint field if non-nil, zero value otherwise.

### GetOtelEndpointOk

`func (o *LogIntegrationResponse) GetOtelEndpointOk() (*string, bool)`

GetOtelEndpointOk returns a tuple with the OtelEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelEndpoint

`func (o *LogIntegrationResponse) SetOtelEndpoint(v string)`

SetOtelEndpoint sets OtelEndpoint field to given value.

### HasOtelEndpoint

`func (o *LogIntegrationResponse) HasOtelEndpoint() bool`

HasOtelEndpoint returns a boolean if a field has been set.
### GetOtelSuppliedHeaders

`func (o *LogIntegrationResponse) GetOtelSuppliedHeaders() []Header`

GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field if non-nil, zero value otherwise.

### GetOtelSuppliedHeadersOk

`func (o *LogIntegrationResponse) GetOtelSuppliedHeadersOk() (*[]Header, bool)`

GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelSuppliedHeaders

`func (o *LogIntegrationResponse) SetOtelSuppliedHeaders(v []Header)`

SetOtelSuppliedHeaders sets OtelSuppliedHeaders field to given value.

### HasOtelSuppliedHeaders

`func (o *LogIntegrationResponse) HasOtelSuppliedHeaders() bool`

HasOtelSuppliedHeaders returns a boolean if a field has been set.
### GetHecToken

`func (o *LogIntegrationResponse) GetHecToken() string`

GetHecToken returns the HecToken field if non-nil, zero value otherwise.

### GetHecTokenOk

`func (o *LogIntegrationResponse) GetHecTokenOk() (*string, bool)`

GetHecTokenOk returns a tuple with the HecToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHecToken

`func (o *LogIntegrationResponse) SetHecToken(v string)`

SetHecToken sets HecToken field to given value.

### HasHecToken

`func (o *LogIntegrationResponse) HasHecToken() bool`

HasHecToken returns a boolean if a field has been set.
### GetHecUrl

`func (o *LogIntegrationResponse) GetHecUrl() string`

GetHecUrl returns the HecUrl field if non-nil, zero value otherwise.

### GetHecUrlOk

`func (o *LogIntegrationResponse) GetHecUrlOk() (*string, bool)`

GetHecUrlOk returns a tuple with the HecUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHecUrl

`func (o *LogIntegrationResponse) SetHecUrl(v string)`

SetHecUrl sets HecUrl field to given value.

### HasHecUrl

`func (o *LogIntegrationResponse) HasHecUrl() bool`

HasHecUrl returns a boolean if a field has been set.
### GetStorageAccountName

`func (o *LogIntegrationResponse) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *LogIntegrationResponse) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *LogIntegrationResponse) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *LogIntegrationResponse) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.
### GetStorageContainerName

`func (o *LogIntegrationResponse) GetStorageContainerName() string`

GetStorageContainerName returns the StorageContainerName field if non-nil, zero value otherwise.

### GetStorageContainerNameOk

`func (o *LogIntegrationResponse) GetStorageContainerNameOk() (*string, bool)`

GetStorageContainerNameOk returns a tuple with the StorageContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerName

`func (o *LogIntegrationResponse) SetStorageContainerName(v string)`

SetStorageContainerName sets StorageContainerName field to given value.

### HasStorageContainerName

`func (o *LogIntegrationResponse) HasStorageContainerName() bool`

HasStorageContainerName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


