# DataLakeStoreSettings

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
**ClusterId** | **string** | ID of the Cluster the Online Archive belongs to. | 
**ClusterName** | **string** | Human-readable label of the MongoDB Cloud cluster on which the store is based. | 
**ProjectId** | **string** | Unique 24-hexadecimal digit string that identifies the project. | [readonly] 
**ReadPreference** | Pointer to [**DataLakeAtlasStoreReadPreference**](DataLakeAtlasStoreReadPreference.md) |  | [optional] 
**AllowInsecure** | Pointer to **bool** | Flag that validates the scheme in the specified URLs. If &#x60;true&#x60;, allows insecure &#x60;HTTP&#x60; scheme, doesn&#39;t verify the server&#39;s certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If &#x60;false&#x60;, allows secure &#x60;HTTPS&#x60; scheme only. | [optional] [default to false]
**DefaultFormat** | Pointer to **string** | Default format that Data Lake assumes if it encounters a file without an extension while searching the &#x60;storeName&#x60;. If omitted, Data Lake attempts to detect the file type by processing a few bytes of the file. The specified format only applies to the URLs specified in the **databases.[n].collections.[n].dataSources** object. | [optional] 
**Urls** | Pointer to **[]string** | Comma-separated list of publicly accessible HTTP URLs where data is stored. You can&#39;t specify URLs that require authentication. | [optional] 

## Methods

### NewDataLakeStoreSettings

`func NewDataLakeStoreSettings(provider string, clusterId string, clusterName string, projectId string, ) *DataLakeStoreSettings`

NewDataLakeStoreSettings instantiates a new DataLakeStoreSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLakeStoreSettingsWithDefaults

`func NewDataLakeStoreSettingsWithDefaults() *DataLakeStoreSettings`

NewDataLakeStoreSettingsWithDefaults instantiates a new DataLakeStoreSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalStorageClasses

`func (o *DataLakeStoreSettings) GetAdditionalStorageClasses() []string`

GetAdditionalStorageClasses returns the AdditionalStorageClasses field if non-nil, zero value otherwise.

### GetAdditionalStorageClassesOk

`func (o *DataLakeStoreSettings) GetAdditionalStorageClassesOk() (*[]string, bool)`

GetAdditionalStorageClassesOk returns a tuple with the AdditionalStorageClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalStorageClasses

`func (o *DataLakeStoreSettings) SetAdditionalStorageClasses(v []string)`

SetAdditionalStorageClasses sets AdditionalStorageClasses field to given value.

### HasAdditionalStorageClasses

`func (o *DataLakeStoreSettings) HasAdditionalStorageClasses() bool`

HasAdditionalStorageClasses returns a boolean if a field has been set.

### GetBucket

`func (o *DataLakeStoreSettings) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *DataLakeStoreSettings) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *DataLakeStoreSettings) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *DataLakeStoreSettings) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetDelimiter

`func (o *DataLakeStoreSettings) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *DataLakeStoreSettings) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *DataLakeStoreSettings) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *DataLakeStoreSettings) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetIncludeTags

`func (o *DataLakeStoreSettings) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *DataLakeStoreSettings) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *DataLakeStoreSettings) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *DataLakeStoreSettings) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetPrefix

`func (o *DataLakeStoreSettings) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DataLakeStoreSettings) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DataLakeStoreSettings) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DataLakeStoreSettings) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPublic

`func (o *DataLakeStoreSettings) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *DataLakeStoreSettings) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *DataLakeStoreSettings) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *DataLakeStoreSettings) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRegion

`func (o *DataLakeStoreSettings) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DataLakeStoreSettings) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DataLakeStoreSettings) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DataLakeStoreSettings) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetName

`func (o *DataLakeStoreSettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataLakeStoreSettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataLakeStoreSettings) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataLakeStoreSettings) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvider

`func (o *DataLakeStoreSettings) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *DataLakeStoreSettings) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *DataLakeStoreSettings) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetClusterId

`func (o *DataLakeStoreSettings) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DataLakeStoreSettings) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DataLakeStoreSettings) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *DataLakeStoreSettings) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *DataLakeStoreSettings) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *DataLakeStoreSettings) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetProjectId

`func (o *DataLakeStoreSettings) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DataLakeStoreSettings) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DataLakeStoreSettings) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetReadPreference

`func (o *DataLakeStoreSettings) GetReadPreference() DataLakeAtlasStoreReadPreference`

GetReadPreference returns the ReadPreference field if non-nil, zero value otherwise.

### GetReadPreferenceOk

`func (o *DataLakeStoreSettings) GetReadPreferenceOk() (*DataLakeAtlasStoreReadPreference, bool)`

GetReadPreferenceOk returns a tuple with the ReadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadPreference

`func (o *DataLakeStoreSettings) SetReadPreference(v DataLakeAtlasStoreReadPreference)`

SetReadPreference sets ReadPreference field to given value.

### HasReadPreference

`func (o *DataLakeStoreSettings) HasReadPreference() bool`

HasReadPreference returns a boolean if a field has been set.

### GetAllowInsecure

`func (o *DataLakeStoreSettings) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *DataLakeStoreSettings) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *DataLakeStoreSettings) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *DataLakeStoreSettings) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetDefaultFormat

`func (o *DataLakeStoreSettings) GetDefaultFormat() string`

GetDefaultFormat returns the DefaultFormat field if non-nil, zero value otherwise.

### GetDefaultFormatOk

`func (o *DataLakeStoreSettings) GetDefaultFormatOk() (*string, bool)`

GetDefaultFormatOk returns a tuple with the DefaultFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormat

`func (o *DataLakeStoreSettings) SetDefaultFormat(v string)`

SetDefaultFormat sets DefaultFormat field to given value.

### HasDefaultFormat

`func (o *DataLakeStoreSettings) HasDefaultFormat() bool`

HasDefaultFormat returns a boolean if a field has been set.

### GetUrls

`func (o *DataLakeStoreSettings) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *DataLakeStoreSettings) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *DataLakeStoreSettings) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *DataLakeStoreSettings) HasUrls() bool`

HasUrls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


