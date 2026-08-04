# LogIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**OtelEndpoint** | Pointer to **string** | OpenTelemetry collector endpoint URL. Must be HTTPS and not exceed 2048 characters. | [optional] 
**OtelSuppliedHeaders** | Pointer to [**[]Header**](Header.md) | HTTP headers for authentication and configuration. Maximum 10 headers, total size limit 2KB. | [optional] 
**HecToken** | Pointer to **string** | HTTP Event Collector (HEC) token for authentication. | [optional] 
**HecUrl** | Pointer to **string** | HTTP Event Collector (HEC) endpoint URL. | [optional] 
**StorageAccountName** | Pointer to **string** | Storage account name where logs will be stored. | [optional] 
**StorageContainerName** | Pointer to **string** | Storage container name for log files. | [optional] 

## Methods

### NewLogIntegrationRequest

`func NewLogIntegrationRequest(logTypes []string, type_ string, ) *LogIntegrationRequest`

NewLogIntegrationRequest instantiates a new LogIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogIntegrationRequestWithDefaults

`func NewLogIntegrationRequestWithDefaults() *LogIntegrationRequest`

NewLogIntegrationRequestWithDefaults instantiates a new LogIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogTypes

`func (o *LogIntegrationRequest) GetLogTypes() []string`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *LogIntegrationRequest) GetLogTypesOk() (*[]string, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *LogIntegrationRequest) SetLogTypes(v []string)`

SetLogTypes sets LogTypes field to given value.

### GetType

`func (o *LogIntegrationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogIntegrationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogIntegrationRequest) SetType(v string)`

SetType sets Type field to given value.

### GetBucketName

`func (o *LogIntegrationRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *LogIntegrationRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *LogIntegrationRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *LogIntegrationRequest) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### SetBucketNameNil

`func (o *LogIntegrationRequest) SetBucketNameNil()`

SetBucketNameNil sets BucketName to an explicit JSON null when marshaled, overriding any value previously set with SetBucketName. Calling SetBucketName again clears the null override.

### GetIamRoleId

`func (o *LogIntegrationRequest) GetIamRoleId() string`

GetIamRoleId returns the IamRoleId field if non-nil, zero value otherwise.

### GetIamRoleIdOk

`func (o *LogIntegrationRequest) GetIamRoleIdOk() (*string, bool)`

GetIamRoleIdOk returns a tuple with the IamRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRoleId

`func (o *LogIntegrationRequest) SetIamRoleId(v string)`

SetIamRoleId sets IamRoleId field to given value.

### HasIamRoleId

`func (o *LogIntegrationRequest) HasIamRoleId() bool`

HasIamRoleId returns a boolean if a field has been set.

### SetIamRoleIdNil

`func (o *LogIntegrationRequest) SetIamRoleIdNil()`

SetIamRoleIdNil sets IamRoleId to an explicit JSON null when marshaled, overriding any value previously set with SetIamRoleId. Calling SetIamRoleId again clears the null override.

### GetKmsKey

`func (o *LogIntegrationRequest) GetKmsKey() string`

GetKmsKey returns the KmsKey field if non-nil, zero value otherwise.

### GetKmsKeyOk

`func (o *LogIntegrationRequest) GetKmsKeyOk() (*string, bool)`

GetKmsKeyOk returns a tuple with the KmsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKey

`func (o *LogIntegrationRequest) SetKmsKey(v string)`

SetKmsKey sets KmsKey field to given value.

### HasKmsKey

`func (o *LogIntegrationRequest) HasKmsKey() bool`

HasKmsKey returns a boolean if a field has been set.

### SetKmsKeyNil

`func (o *LogIntegrationRequest) SetKmsKeyNil()`

SetKmsKeyNil sets KmsKey to an explicit JSON null when marshaled, overriding any value previously set with SetKmsKey. Calling SetKmsKey again clears the null override.

### GetPrefixPath

`func (o *LogIntegrationRequest) GetPrefixPath() string`

GetPrefixPath returns the PrefixPath field if non-nil, zero value otherwise.

### GetPrefixPathOk

`func (o *LogIntegrationRequest) GetPrefixPathOk() (*string, bool)`

GetPrefixPathOk returns a tuple with the PrefixPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixPath

`func (o *LogIntegrationRequest) SetPrefixPath(v string)`

SetPrefixPath sets PrefixPath field to given value.

### HasPrefixPath

`func (o *LogIntegrationRequest) HasPrefixPath() bool`

HasPrefixPath returns a boolean if a field has been set.

### SetPrefixPathNil

`func (o *LogIntegrationRequest) SetPrefixPathNil()`

SetPrefixPathNil sets PrefixPath to an explicit JSON null when marshaled, overriding any value previously set with SetPrefixPath. Calling SetPrefixPath again clears the null override.

### GetUseLegacyPathStructure

`func (o *LogIntegrationRequest) GetUseLegacyPathStructure() bool`

GetUseLegacyPathStructure returns the UseLegacyPathStructure field if non-nil, zero value otherwise.

### GetUseLegacyPathStructureOk

`func (o *LogIntegrationRequest) GetUseLegacyPathStructureOk() (*bool, bool)`

GetUseLegacyPathStructureOk returns a tuple with the UseLegacyPathStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyPathStructure

`func (o *LogIntegrationRequest) SetUseLegacyPathStructure(v bool)`

SetUseLegacyPathStructure sets UseLegacyPathStructure field to given value.

### HasUseLegacyPathStructure

`func (o *LogIntegrationRequest) HasUseLegacyPathStructure() bool`

HasUseLegacyPathStructure returns a boolean if a field has been set.

### SetUseLegacyPathStructureNil

`func (o *LogIntegrationRequest) SetUseLegacyPathStructureNil()`

SetUseLegacyPathStructureNil sets UseLegacyPathStructure to an explicit JSON null when marshaled, overriding any value previously set with SetUseLegacyPathStructure. Calling SetUseLegacyPathStructure again clears the null override.

### GetApiKey

`func (o *LogIntegrationRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *LogIntegrationRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *LogIntegrationRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *LogIntegrationRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *LogIntegrationRequest) SetApiKeyNil()`

SetApiKeyNil sets ApiKey to an explicit JSON null when marshaled, overriding any value previously set with SetApiKey. Calling SetApiKey again clears the null override.

### GetRegion

`func (o *LogIntegrationRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogIntegrationRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogIntegrationRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogIntegrationRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *LogIntegrationRequest) SetRegionNil()`

SetRegionNil sets Region to an explicit JSON null when marshaled, overriding any value previously set with SetRegion. Calling SetRegion again clears the null override.

### GetRoleId

`func (o *LogIntegrationRequest) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *LogIntegrationRequest) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *LogIntegrationRequest) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *LogIntegrationRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.

### SetRoleIdNil

`func (o *LogIntegrationRequest) SetRoleIdNil()`

SetRoleIdNil sets RoleId to an explicit JSON null when marshaled, overriding any value previously set with SetRoleId. Calling SetRoleId again clears the null override.

### GetOtelEndpoint

`func (o *LogIntegrationRequest) GetOtelEndpoint() string`

GetOtelEndpoint returns the OtelEndpoint field if non-nil, zero value otherwise.

### GetOtelEndpointOk

`func (o *LogIntegrationRequest) GetOtelEndpointOk() (*string, bool)`

GetOtelEndpointOk returns a tuple with the OtelEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelEndpoint

`func (o *LogIntegrationRequest) SetOtelEndpoint(v string)`

SetOtelEndpoint sets OtelEndpoint field to given value.

### HasOtelEndpoint

`func (o *LogIntegrationRequest) HasOtelEndpoint() bool`

HasOtelEndpoint returns a boolean if a field has been set.

### SetOtelEndpointNil

`func (o *LogIntegrationRequest) SetOtelEndpointNil()`

SetOtelEndpointNil sets OtelEndpoint to an explicit JSON null when marshaled, overriding any value previously set with SetOtelEndpoint. Calling SetOtelEndpoint again clears the null override.

### GetOtelSuppliedHeaders

`func (o *LogIntegrationRequest) GetOtelSuppliedHeaders() []Header`

GetOtelSuppliedHeaders returns the OtelSuppliedHeaders field if non-nil, zero value otherwise.

### GetOtelSuppliedHeadersOk

`func (o *LogIntegrationRequest) GetOtelSuppliedHeadersOk() (*[]Header, bool)`

GetOtelSuppliedHeadersOk returns a tuple with the OtelSuppliedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelSuppliedHeaders

`func (o *LogIntegrationRequest) SetOtelSuppliedHeaders(v []Header)`

SetOtelSuppliedHeaders sets OtelSuppliedHeaders field to given value.

### HasOtelSuppliedHeaders

`func (o *LogIntegrationRequest) HasOtelSuppliedHeaders() bool`

HasOtelSuppliedHeaders returns a boolean if a field has been set.

### SetOtelSuppliedHeadersNil

`func (o *LogIntegrationRequest) SetOtelSuppliedHeadersNil()`

SetOtelSuppliedHeadersNil sets OtelSuppliedHeaders to an explicit JSON null when marshaled, overriding any value previously set with SetOtelSuppliedHeaders. Calling SetOtelSuppliedHeaders again clears the null override.

### GetHecToken

`func (o *LogIntegrationRequest) GetHecToken() string`

GetHecToken returns the HecToken field if non-nil, zero value otherwise.

### GetHecTokenOk

`func (o *LogIntegrationRequest) GetHecTokenOk() (*string, bool)`

GetHecTokenOk returns a tuple with the HecToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHecToken

`func (o *LogIntegrationRequest) SetHecToken(v string)`

SetHecToken sets HecToken field to given value.

### HasHecToken

`func (o *LogIntegrationRequest) HasHecToken() bool`

HasHecToken returns a boolean if a field has been set.

### SetHecTokenNil

`func (o *LogIntegrationRequest) SetHecTokenNil()`

SetHecTokenNil sets HecToken to an explicit JSON null when marshaled, overriding any value previously set with SetHecToken. Calling SetHecToken again clears the null override.

### GetHecUrl

`func (o *LogIntegrationRequest) GetHecUrl() string`

GetHecUrl returns the HecUrl field if non-nil, zero value otherwise.

### GetHecUrlOk

`func (o *LogIntegrationRequest) GetHecUrlOk() (*string, bool)`

GetHecUrlOk returns a tuple with the HecUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHecUrl

`func (o *LogIntegrationRequest) SetHecUrl(v string)`

SetHecUrl sets HecUrl field to given value.

### HasHecUrl

`func (o *LogIntegrationRequest) HasHecUrl() bool`

HasHecUrl returns a boolean if a field has been set.

### SetHecUrlNil

`func (o *LogIntegrationRequest) SetHecUrlNil()`

SetHecUrlNil sets HecUrl to an explicit JSON null when marshaled, overriding any value previously set with SetHecUrl. Calling SetHecUrl again clears the null override.

### GetStorageAccountName

`func (o *LogIntegrationRequest) GetStorageAccountName() string`

GetStorageAccountName returns the StorageAccountName field if non-nil, zero value otherwise.

### GetStorageAccountNameOk

`func (o *LogIntegrationRequest) GetStorageAccountNameOk() (*string, bool)`

GetStorageAccountNameOk returns a tuple with the StorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountName

`func (o *LogIntegrationRequest) SetStorageAccountName(v string)`

SetStorageAccountName sets StorageAccountName field to given value.

### HasStorageAccountName

`func (o *LogIntegrationRequest) HasStorageAccountName() bool`

HasStorageAccountName returns a boolean if a field has been set.

### SetStorageAccountNameNil

`func (o *LogIntegrationRequest) SetStorageAccountNameNil()`

SetStorageAccountNameNil sets StorageAccountName to an explicit JSON null when marshaled, overriding any value previously set with SetStorageAccountName. Calling SetStorageAccountName again clears the null override.

### GetStorageContainerName

`func (o *LogIntegrationRequest) GetStorageContainerName() string`

GetStorageContainerName returns the StorageContainerName field if non-nil, zero value otherwise.

### GetStorageContainerNameOk

`func (o *LogIntegrationRequest) GetStorageContainerNameOk() (*string, bool)`

GetStorageContainerNameOk returns a tuple with the StorageContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageContainerName

`func (o *LogIntegrationRequest) SetStorageContainerName(v string)`

SetStorageContainerName sets StorageContainerName field to given value.

### HasStorageContainerName

`func (o *LogIntegrationRequest) HasStorageContainerName() bool`

HasStorageContainerName returns a boolean if a field has been set.

### SetStorageContainerNameNil

`func (o *LogIntegrationRequest) SetStorageContainerNameNil()`

SetStorageContainerNameNil sets StorageContainerName to an explicit JSON null when marshaled, overriding any value previously set with SetStorageContainerName. Calling SetStorageContainerName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


