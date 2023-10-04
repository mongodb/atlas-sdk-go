# ClusterDescriptionProcessArgs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChunkMigrationConcurrency** | Pointer to **int** | Number of threads on the source shard and the receiving shard for chunk migration. The number of threads should not exceed the half the total number of CPU cores in the sharded cluster. | [optional] [default to 1]
**DefaultReadConcern** | Pointer to **string** | Default level of acknowledgment requested from MongoDB for read operations set for this cluster.  MongoDB 4.4 clusters default to &#x60;available&#x60;. MongoDB 5.0 and later clusters default to &#x60;local&#x60;. | [optional] [default to "available"]
**DefaultWriteConcern** | Pointer to **string** | Default level of acknowledgment requested from MongoDB for write operations set for this cluster.  MongoDB 4.4 clusters default to &#x60;1&#x60;. MongoDB 5.0 and later clusters default to &#x60;majority&#x60;. | [optional] [default to "1"]
**FailIndexKeyTooLong** | Pointer to **bool** | Flag that indicates whether you can insert or update documents where all indexed entries don&#39;t exceed 1024 bytes. If you set this to false, [mongod](https://docs.mongodb.com/upcoming/reference/program/mongod/#mongodb-binary-bin.mongod) writes documents that exceed this limit but doesn&#39;t index them. This parameter has been removed as of [MongoDB 4.4](https://www.mongodb.com/docs/manual/reference/parameters/#mongodb-parameter-param.failIndexKeyTooLong). | [optional] [default to true]
**JavascriptEnabled** | Pointer to **bool** | Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript. | [optional] [default to true]
**MinimumEnabledTlsProtocol** | Pointer to **string** | Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version. | [optional] 
**NoTableScan** | Pointer to **bool** | Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results. | [optional] [default to false]
**OplogMinRetentionHours** | Pointer to **float64** | Minimum retention window for cluster&#39;s oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates. | [optional] 
**OplogSizeMB** | Pointer to **int** | Storage limit of cluster&#39;s oplog expressed in megabytes. A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates. | [optional] 
**SampleRefreshIntervalBIConnector** | Pointer to **int** | Interval in seconds at which the mongosqld process re-samples data to create its relational schema. | [optional] [default to 0]
**SampleSizeBIConnector** | Pointer to **int** | Number of documents per database to sample when gathering schema information. | [optional] [default to 1000]
**TransactionLifetimeLimitSeconds** | Pointer to **int64** | Lifetime, in seconds, of multi-document transactions. Atlas considers the transactions that exceed this limit as expired and so aborts them through a periodic cleanup process. | [optional] [default to 60]

## Methods

### NewClusterDescriptionProcessArgs

`func NewClusterDescriptionProcessArgs() *ClusterDescriptionProcessArgs`

NewClusterDescriptionProcessArgs instantiates a new ClusterDescriptionProcessArgs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDescriptionProcessArgsWithDefaults

`func NewClusterDescriptionProcessArgsWithDefaults() *ClusterDescriptionProcessArgs`

NewClusterDescriptionProcessArgsWithDefaults instantiates a new ClusterDescriptionProcessArgs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunkMigrationConcurrency

`func (o *ClusterDescriptionProcessArgs) GetChunkMigrationConcurrency() int`

GetChunkMigrationConcurrency returns the ChunkMigrationConcurrency field if non-nil, zero value otherwise.

### GetChunkMigrationConcurrencyOk

`func (o *ClusterDescriptionProcessArgs) GetChunkMigrationConcurrencyOk() (*int, bool)`

GetChunkMigrationConcurrencyOk returns a tuple with the ChunkMigrationConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkMigrationConcurrency

`func (o *ClusterDescriptionProcessArgs) SetChunkMigrationConcurrency(v int)`

SetChunkMigrationConcurrency sets ChunkMigrationConcurrency field to given value.

### HasChunkMigrationConcurrency

`func (o *ClusterDescriptionProcessArgs) HasChunkMigrationConcurrency() bool`

HasChunkMigrationConcurrency returns a boolean if a field has been set.
### GetDefaultReadConcern

`func (o *ClusterDescriptionProcessArgs) GetDefaultReadConcern() string`

GetDefaultReadConcern returns the DefaultReadConcern field if non-nil, zero value otherwise.

### GetDefaultReadConcernOk

`func (o *ClusterDescriptionProcessArgs) GetDefaultReadConcernOk() (*string, bool)`

GetDefaultReadConcernOk returns a tuple with the DefaultReadConcern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultReadConcern

`func (o *ClusterDescriptionProcessArgs) SetDefaultReadConcern(v string)`

SetDefaultReadConcern sets DefaultReadConcern field to given value.

### HasDefaultReadConcern

`func (o *ClusterDescriptionProcessArgs) HasDefaultReadConcern() bool`

HasDefaultReadConcern returns a boolean if a field has been set.
### GetDefaultWriteConcern

`func (o *ClusterDescriptionProcessArgs) GetDefaultWriteConcern() string`

GetDefaultWriteConcern returns the DefaultWriteConcern field if non-nil, zero value otherwise.

### GetDefaultWriteConcernOk

`func (o *ClusterDescriptionProcessArgs) GetDefaultWriteConcernOk() (*string, bool)`

GetDefaultWriteConcernOk returns a tuple with the DefaultWriteConcern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWriteConcern

`func (o *ClusterDescriptionProcessArgs) SetDefaultWriteConcern(v string)`

SetDefaultWriteConcern sets DefaultWriteConcern field to given value.

### HasDefaultWriteConcern

`func (o *ClusterDescriptionProcessArgs) HasDefaultWriteConcern() bool`

HasDefaultWriteConcern returns a boolean if a field has been set.
### GetFailIndexKeyTooLong

`func (o *ClusterDescriptionProcessArgs) GetFailIndexKeyTooLong() bool`

GetFailIndexKeyTooLong returns the FailIndexKeyTooLong field if non-nil, zero value otherwise.

### GetFailIndexKeyTooLongOk

`func (o *ClusterDescriptionProcessArgs) GetFailIndexKeyTooLongOk() (*bool, bool)`

GetFailIndexKeyTooLongOk returns a tuple with the FailIndexKeyTooLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailIndexKeyTooLong

`func (o *ClusterDescriptionProcessArgs) SetFailIndexKeyTooLong(v bool)`

SetFailIndexKeyTooLong sets FailIndexKeyTooLong field to given value.

### HasFailIndexKeyTooLong

`func (o *ClusterDescriptionProcessArgs) HasFailIndexKeyTooLong() bool`

HasFailIndexKeyTooLong returns a boolean if a field has been set.
### GetJavascriptEnabled

`func (o *ClusterDescriptionProcessArgs) GetJavascriptEnabled() bool`

GetJavascriptEnabled returns the JavascriptEnabled field if non-nil, zero value otherwise.

### GetJavascriptEnabledOk

`func (o *ClusterDescriptionProcessArgs) GetJavascriptEnabledOk() (*bool, bool)`

GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptEnabled

`func (o *ClusterDescriptionProcessArgs) SetJavascriptEnabled(v bool)`

SetJavascriptEnabled sets JavascriptEnabled field to given value.

### HasJavascriptEnabled

`func (o *ClusterDescriptionProcessArgs) HasJavascriptEnabled() bool`

HasJavascriptEnabled returns a boolean if a field has been set.
### GetMinimumEnabledTlsProtocol

`func (o *ClusterDescriptionProcessArgs) GetMinimumEnabledTlsProtocol() string`

GetMinimumEnabledTlsProtocol returns the MinimumEnabledTlsProtocol field if non-nil, zero value otherwise.

### GetMinimumEnabledTlsProtocolOk

`func (o *ClusterDescriptionProcessArgs) GetMinimumEnabledTlsProtocolOk() (*string, bool)`

GetMinimumEnabledTlsProtocolOk returns a tuple with the MinimumEnabledTlsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumEnabledTlsProtocol

`func (o *ClusterDescriptionProcessArgs) SetMinimumEnabledTlsProtocol(v string)`

SetMinimumEnabledTlsProtocol sets MinimumEnabledTlsProtocol field to given value.

### HasMinimumEnabledTlsProtocol

`func (o *ClusterDescriptionProcessArgs) HasMinimumEnabledTlsProtocol() bool`

HasMinimumEnabledTlsProtocol returns a boolean if a field has been set.
### GetNoTableScan

`func (o *ClusterDescriptionProcessArgs) GetNoTableScan() bool`

GetNoTableScan returns the NoTableScan field if non-nil, zero value otherwise.

### GetNoTableScanOk

`func (o *ClusterDescriptionProcessArgs) GetNoTableScanOk() (*bool, bool)`

GetNoTableScanOk returns a tuple with the NoTableScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTableScan

`func (o *ClusterDescriptionProcessArgs) SetNoTableScan(v bool)`

SetNoTableScan sets NoTableScan field to given value.

### HasNoTableScan

`func (o *ClusterDescriptionProcessArgs) HasNoTableScan() bool`

HasNoTableScan returns a boolean if a field has been set.
### GetOplogMinRetentionHours

`func (o *ClusterDescriptionProcessArgs) GetOplogMinRetentionHours() float64`

GetOplogMinRetentionHours returns the OplogMinRetentionHours field if non-nil, zero value otherwise.

### GetOplogMinRetentionHoursOk

`func (o *ClusterDescriptionProcessArgs) GetOplogMinRetentionHoursOk() (*float64, bool)`

GetOplogMinRetentionHoursOk returns a tuple with the OplogMinRetentionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogMinRetentionHours

`func (o *ClusterDescriptionProcessArgs) SetOplogMinRetentionHours(v float64)`

SetOplogMinRetentionHours sets OplogMinRetentionHours field to given value.

### HasOplogMinRetentionHours

`func (o *ClusterDescriptionProcessArgs) HasOplogMinRetentionHours() bool`

HasOplogMinRetentionHours returns a boolean if a field has been set.
### GetOplogSizeMB

`func (o *ClusterDescriptionProcessArgs) GetOplogSizeMB() int`

GetOplogSizeMB returns the OplogSizeMB field if non-nil, zero value otherwise.

### GetOplogSizeMBOk

`func (o *ClusterDescriptionProcessArgs) GetOplogSizeMBOk() (*int, bool)`

GetOplogSizeMBOk returns a tuple with the OplogSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogSizeMB

`func (o *ClusterDescriptionProcessArgs) SetOplogSizeMB(v int)`

SetOplogSizeMB sets OplogSizeMB field to given value.

### HasOplogSizeMB

`func (o *ClusterDescriptionProcessArgs) HasOplogSizeMB() bool`

HasOplogSizeMB returns a boolean if a field has been set.
### GetSampleRefreshIntervalBIConnector

`func (o *ClusterDescriptionProcessArgs) GetSampleRefreshIntervalBIConnector() int`

GetSampleRefreshIntervalBIConnector returns the SampleRefreshIntervalBIConnector field if non-nil, zero value otherwise.

### GetSampleRefreshIntervalBIConnectorOk

`func (o *ClusterDescriptionProcessArgs) GetSampleRefreshIntervalBIConnectorOk() (*int, bool)`

GetSampleRefreshIntervalBIConnectorOk returns a tuple with the SampleRefreshIntervalBIConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRefreshIntervalBIConnector

`func (o *ClusterDescriptionProcessArgs) SetSampleRefreshIntervalBIConnector(v int)`

SetSampleRefreshIntervalBIConnector sets SampleRefreshIntervalBIConnector field to given value.

### HasSampleRefreshIntervalBIConnector

`func (o *ClusterDescriptionProcessArgs) HasSampleRefreshIntervalBIConnector() bool`

HasSampleRefreshIntervalBIConnector returns a boolean if a field has been set.
### GetSampleSizeBIConnector

`func (o *ClusterDescriptionProcessArgs) GetSampleSizeBIConnector() int`

GetSampleSizeBIConnector returns the SampleSizeBIConnector field if non-nil, zero value otherwise.

### GetSampleSizeBIConnectorOk

`func (o *ClusterDescriptionProcessArgs) GetSampleSizeBIConnectorOk() (*int, bool)`

GetSampleSizeBIConnectorOk returns a tuple with the SampleSizeBIConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSizeBIConnector

`func (o *ClusterDescriptionProcessArgs) SetSampleSizeBIConnector(v int)`

SetSampleSizeBIConnector sets SampleSizeBIConnector field to given value.

### HasSampleSizeBIConnector

`func (o *ClusterDescriptionProcessArgs) HasSampleSizeBIConnector() bool`

HasSampleSizeBIConnector returns a boolean if a field has been set.
### GetTransactionLifetimeLimitSeconds

`func (o *ClusterDescriptionProcessArgs) GetTransactionLifetimeLimitSeconds() int64`

GetTransactionLifetimeLimitSeconds returns the TransactionLifetimeLimitSeconds field if non-nil, zero value otherwise.

### GetTransactionLifetimeLimitSecondsOk

`func (o *ClusterDescriptionProcessArgs) GetTransactionLifetimeLimitSecondsOk() (*int64, bool)`

GetTransactionLifetimeLimitSecondsOk returns a tuple with the TransactionLifetimeLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLifetimeLimitSeconds

`func (o *ClusterDescriptionProcessArgs) SetTransactionLifetimeLimitSeconds(v int64)`

SetTransactionLifetimeLimitSeconds sets TransactionLifetimeLimitSeconds field to given value.

### HasTransactionLifetimeLimitSeconds

`func (o *ClusterDescriptionProcessArgs) HasTransactionLifetimeLimitSeconds() bool`

HasTransactionLifetimeLimitSeconds returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


