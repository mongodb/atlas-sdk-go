# BaseCloudProviderAccessFeatureUsageFeatureId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies your project. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable label that identifies the data lake. | [optional] 
**ConfigId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the Atlas log integration configuration. | [optional] [readonly] 
**SinkName** | Pointer to **string** | Human-readable label that identifies the sink, such as an S3/GCS bucket or Azure container name. | [optional] [readonly] 
**ExportBucketId** | Pointer to **string** | Unique 24-hexadecimal digit string that identifies the AWS S3 bucket to which you export your snapshots. | [optional] [readonly] 
**BucketName** | Pointer to **string** | Name of the AWS S3 bucket to which your logs will be exported to. | [optional] [readonly] 

## Methods

### NewBaseCloudProviderAccessFeatureUsageFeatureId

`func NewBaseCloudProviderAccessFeatureUsageFeatureId() *BaseCloudProviderAccessFeatureUsageFeatureId`

NewBaseCloudProviderAccessFeatureUsageFeatureId instantiates a new BaseCloudProviderAccessFeatureUsageFeatureId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseCloudProviderAccessFeatureUsageFeatureIdWithDefaults

`func NewBaseCloudProviderAccessFeatureUsageFeatureIdWithDefaults() *BaseCloudProviderAccessFeatureUsageFeatureId`

NewBaseCloudProviderAccessFeatureUsageFeatureIdWithDefaults instantiates a new BaseCloudProviderAccessFeatureUsageFeatureId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetGroupIdNil()`

SetGroupIdNil sets GroupId to an explicit JSON null when marshaled, overriding any value previously set with SetGroupId. Calling SetGroupId again clears the null override.

### GetName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetNameNil()`

SetNameNil sets Name to an explicit JSON null when marshaled, overriding any value previously set with SetName. Calling SetName again clears the null override.

### GetConfigId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### SetConfigIdNil

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetConfigIdNil()`

SetConfigIdNil sets ConfigId to an explicit JSON null when marshaled, overriding any value previously set with SetConfigId. Calling SetConfigId again clears the null override.

### GetSinkName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetSinkName() string`

GetSinkName returns the SinkName field if non-nil, zero value otherwise.

### GetSinkNameOk

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetSinkNameOk() (*string, bool)`

GetSinkNameOk returns a tuple with the SinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinkName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetSinkName(v string)`

SetSinkName sets SinkName field to given value.

### HasSinkName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasSinkName() bool`

HasSinkName returns a boolean if a field has been set.

### SetSinkNameNil

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetSinkNameNil()`

SetSinkNameNil sets SinkName to an explicit JSON null when marshaled, overriding any value previously set with SetSinkName. Calling SetSinkName again clears the null override.

### GetExportBucketId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetExportBucketId() string`

GetExportBucketId returns the ExportBucketId field if non-nil, zero value otherwise.

### GetExportBucketIdOk

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetExportBucketIdOk() (*string, bool)`

GetExportBucketIdOk returns a tuple with the ExportBucketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportBucketId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetExportBucketId(v string)`

SetExportBucketId sets ExportBucketId field to given value.

### HasExportBucketId

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasExportBucketId() bool`

HasExportBucketId returns a boolean if a field has been set.

### SetExportBucketIdNil

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetExportBucketIdNil()`

SetExportBucketIdNil sets ExportBucketId to an explicit JSON null when marshaled, overriding any value previously set with SetExportBucketId. Calling SetExportBucketId again clears the null override.

### GetBucketName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### SetBucketNameNil

`func (o *BaseCloudProviderAccessFeatureUsageFeatureId) SetBucketNameNil()`

SetBucketNameNil sets BucketName to an explicit JSON null when marshaled, overriding any value previously set with SetBucketName. Calling SetBucketName again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


