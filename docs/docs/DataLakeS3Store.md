# DataLakeS3Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalStorageClasses** | Pointer to **[]string** | Collection of AWS S3 [storage classes](https://aws.amazon.com/s3/storage-classes/). Atlas Data Lake includes the files in these storage classes in the query results. | [optional] 
**Bucket** | Pointer to **string** | Human-readable label that identifies the AWS S3 bucket. This label must exactly match the name of an S3 bucket that the data lake can access with the configured AWS Identity and Access Management (IAM) credentials. | [optional] 
**Delimiter** | Pointer to **string** | The delimiter that separates **databases.[n].collections.[n].dataSources.[n].path** segments in the data store. MongoDB Cloud uses the delimiter to efficiently traverse S3 buckets with a hierarchical directory structure. You can specify any character supported by the S3 object keys as the delimiter. For example, you can specify an underscore (_) or a plus sign (+) or multiple characters, such as double underscores (__) as the delimiter. If omitted, defaults to &#x60;/&#x60;. | [optional] 
**IncludeTags** | Pointer to **bool** | Flag that indicates whether to use S3 tags on the files in the given path as additional partition attributes. If set to &#x60;true&#x60;, data lake adds the S3 tags as additional partition attributes and adds new top-level BSON elements associating each tag to each document. | [optional] [default to false]
**Prefix** | Pointer to **string** | Prefix that MongoDB Cloud applies when searching for files in the S3 bucket. The data store prepends the value of prefix to the **databases.[n].collections.[n].dataSources.[n].path** to create the full path for files to ingest. If omitted, MongoDB Cloud searches all files from the root of the S3 bucket. | [optional] 
**Public** | Pointer to **bool** | Flag that indicates whether the bucket is public. If set to &#x60;true&#x60;, MongoDB Cloud doesn&#39;t use the configured AWS Identity and Access Management (IAM) role to access the S3 bucket. If set to &#x60;false&#x60;, the configured AWS IAM role must include permissions to access the S3 bucket. | [optional] [default to false]
**Region** | Pointer to **string** |  Physical location where MongoDB Cloud deploys your AWS-hosted MongoDB cluster nodes. The region you choose can affect network latency for clients accessing your databases. When MongoDB Cloud deploys a dedicated cluster, it checks if a VPC or VPC connection exists for that provider and region. If not, MongoDB Cloud creates them as part of the deployment. MongoDB Cloud assigns the VPC a CIDR block. To limit a new VPC peering connection to one CIDR block and region, create the connection first. Deploy the cluster after the connection starts. | [optional] 
**Name** | Pointer to **string** | Human-readable label that identifies the data store. The **databases.[n].collections.[n].dataSources.[n].storeName** field references this values as part of the mapping configuration. To use MongoDB Cloud as a data store, the data lake requires a serverless instance or an &#x60;M10&#x60; or higher cluster. | [optional] 
**Provider** | **string** |  | 

## Methods

### NewDataLakeS3Store

`func NewDataLakeS3Store(provider string, ) *DataLakeS3Store`

NewDataLakeS3Store instantiates a new DataLakeS3Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeS3StoreWithDefaults

`func NewDataLakeS3StoreWithDefaults() *DataLakeS3Store`

NewDataLakeS3StoreWithDefaults instantiates a new DataLakeS3Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalStorageClasses

`func (o *DataLakeS3Store) GetAdditionalStorageClasses() []string`

GetAdditionalStorageClasses returns the AdditionalStorageClasses field if non-nil, zero value otherwise.

### GetAdditionalStorageClassesOk

`func (o *DataLakeS3Store) GetAdditionalStorageClassesOk() (*[]string, bool)`

GetAdditionalStorageClassesOk returns a tuple with the AdditionalStorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalStorageClasses

`func (o *DataLakeS3Store) SetAdditionalStorageClasses(v []string)`

SetAdditionalStorageClasses sets AdditionalStorageClasses field to given value.

### HasAdditionalStorageClasses

`func (o *DataLakeS3Store) HasAdditionalStorageClasses() bool`

HasAdditionalStorageClasses returns a boolean if a field has been set.

### GetBucket

`func (o *DataLakeS3Store) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DataLakeS3Store) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DataLakeS3Store) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *DataLakeS3Store) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetDelimiter

`func (o *DataLakeS3Store) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *DataLakeS3Store) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *DataLakeS3Store) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *DataLakeS3Store) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetIncludeTags

`func (o *DataLakeS3Store) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *DataLakeS3Store) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *DataLakeS3Store) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *DataLakeS3Store) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetPrefix

`func (o *DataLakeS3Store) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DataLakeS3Store) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DataLakeS3Store) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DataLakeS3Store) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPublic

`func (o *DataLakeS3Store) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *DataLakeS3Store) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *DataLakeS3Store) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *DataLakeS3Store) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRegion

`func (o *DataLakeS3Store) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DataLakeS3Store) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DataLakeS3Store) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DataLakeS3Store) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetName

`func (o *DataLakeS3Store) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeS3Store) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeS3Store) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeS3Store) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DataLakeS3Store) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DataLakeS3Store) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DataLakeS3Store) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


