# DataLakeAWSCloudProviderConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores. | [optional] [readonly] 
**IamAssumedRoleARN** | Pointer to **string** | Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores. | [optional] [readonly] 
**IamUserARN** | Pointer to **string** | Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores. | [optional] [readonly] 
**RoleId** | **string** | Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig. | 
**TestS3Bucket** | **string** | Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig. | 

## Methods

### NewDataLakeAWSCloudProviderConfig

`func NewDataLakeAWSCloudProviderConfig(roleId string, testS3Bucket string, ) *DataLakeAWSCloudProviderConfig`

NewDataLakeAWSCloudProviderConfig instantiates a new DataLakeAWSCloudProviderConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeAWSCloudProviderConfigWithDefaults

`func NewDataLakeAWSCloudProviderConfigWithDefaults() *DataLakeAWSCloudProviderConfig`

NewDataLakeAWSCloudProviderConfigWithDefaults instantiates a new DataLakeAWSCloudProviderConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *DataLakeAWSCloudProviderConfig) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *DataLakeAWSCloudProviderConfig) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *DataLakeAWSCloudProviderConfig) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *DataLakeAWSCloudProviderConfig) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIamAssumedRoleARN

`func (o *DataLakeAWSCloudProviderConfig) GetIamAssumedRoleARN() string`

GetIamAssumedRoleARN returns the IamAssumedRoleARN field if non-nil, zero value otherwise.

### GetIamAssumedRoleARNOk

`func (o *DataLakeAWSCloudProviderConfig) GetIamAssumedRoleARNOk() (*string, bool)`

GetIamAssumedRoleARNOk returns a tuple with the IamAssumedRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamAssumedRoleARN

`func (o *DataLakeAWSCloudProviderConfig) SetIamAssumedRoleARN(v string)`

SetIamAssumedRoleARN sets IamAssumedRoleARN field to given value.

### HasIamAssumedRoleARN

`func (o *DataLakeAWSCloudProviderConfig) HasIamAssumedRoleARN() bool`

HasIamAssumedRoleARN returns a boolean if a field has been set.

### GetIamUserARN

`func (o *DataLakeAWSCloudProviderConfig) GetIamUserARN() string`

GetIamUserARN returns the IamUserARN field if non-nil, zero value otherwise.

### GetIamUserARNOk

`func (o *DataLakeAWSCloudProviderConfig) GetIamUserARNOk() (*string, bool)`

GetIamUserARNOk returns a tuple with the IamUserARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamUserARN

`func (o *DataLakeAWSCloudProviderConfig) SetIamUserARN(v string)`

SetIamUserARN sets IamUserARN field to given value.

### HasIamUserARN

`func (o *DataLakeAWSCloudProviderConfig) HasIamUserARN() bool`

HasIamUserARN returns a boolean if a field has been set.

### GetRoleId

`func (o *DataLakeAWSCloudProviderConfig) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DataLakeAWSCloudProviderConfig) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DataLakeAWSCloudProviderConfig) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetTestS3Bucket

`func (o *DataLakeAWSCloudProviderConfig) GetTestS3Bucket() string`

GetTestS3Bucket returns the TestS3Bucket field if non-nil, zero value otherwise.

### GetTestS3BucketOk

`func (o *DataLakeAWSCloudProviderConfig) GetTestS3BucketOk() (*string, bool)`

GetTestS3BucketOk returns a tuple with the TestS3Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestS3Bucket

`func (o *DataLakeAWSCloudProviderConfig) SetTestS3Bucket(v string)`

SetTestS3Bucket sets TestS3Bucket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


