# ClusterDescriptionProcessArgs20240805

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds** | Pointer to **int** | The minimum pre- and post-image retention time in seconds. | [optional] [default to -1]
**ChunkMigrationConcurrency** | Pointer to **int** | Number of threads on the source shard and the receiving shard for chunk migration. The number of threads should not exceed the half the total number of CPU cores in the sharded cluster. | [optional] 
**CustomOpensslCipherConfigTls12** | Pointer to **[]string** | The custom OpenSSL cipher suite list for TLS 1.2. This field is only valid when &#x60;tlsCipherConfigMode&#x60; is set to &#x60;CUSTOM&#x60;. | [optional] 
**CustomOpensslCipherConfigTls13** | Pointer to **[]string** | The custom OpenSSL cipher suite list for TLS 1.3. This field is only valid when &#x60;tlsCipherConfigMode&#x60; is set to &#x60;CUSTOM&#x60;. | [optional] 
**DefaultMaxTimeMS** | Pointer to **int** | Default time limit in milliseconds for individual read operations to complete. | [optional] 
**DefaultWriteConcern** | Pointer to **string** | Default level of acknowledgment requested from MongoDB for write operations when none is specified by the driver. | [optional] 
**JavascriptEnabled** | Pointer to **bool** | Flag that indicates whether the cluster allows execution of operations that perform server-side executions of JavaScript. When using 8.0+, we recommend disabling server-side JavaScript and using operators of aggregation pipeline as more performant alternative. | [optional] 
**MinimumEnabledTlsProtocol** | Pointer to **string** | Minimum Transport Layer Security (TLS) version that the cluster accepts for incoming connections. Clusters using TLS 1.0 or 1.1 should consider setting TLS 1.2 as the minimum TLS protocol version. | [optional] 
**NoTableScan** | Pointer to **bool** | Flag that indicates whether the cluster disables executing any query that requires a collection scan to return results. | [optional] 
**OplogMinRetentionHours** | Pointer to **float64** | Minimum retention window for cluster&#39;s oplog expressed in hours. A value of null indicates that the cluster uses the default minimum oplog window that MongoDB Cloud calculates. | [optional] 
**OplogSizeMB** | Pointer to **int** | Storage limit of cluster&#39;s oplog expressed in megabytes. A value of null indicates that the cluster uses the default oplog size that MongoDB Cloud calculates. | [optional] 
**QueryStatsLogVerbosity** | Pointer to **int** | May be set to 1 (disabled) or 3 (enabled). When set to 3, Atlas will include redacted and anonymized $queryStats output in MongoDB logs. $queryStats output does not contain literals or field values. Enabling this setting might impact the performance of your cluster. | [optional] 
**SampleRefreshIntervalBIConnector** | Pointer to **int** | Interval in seconds at which the mongosqld process re-samples data to create its relational schema. | [optional] [default to 0]
**SampleSizeBIConnector** | Pointer to **int** | Number of documents per database to sample when gathering schema information. | [optional] 
**TlsCipherConfigMode** | Pointer to **string** | The TLS cipher suite configuration mode. The default mode uses the default cipher suites. The custom mode allows you to specify custom cipher suites for both TLS 1.2 and TLS 1.3. | [optional] 
**TransactionLifetimeLimitSeconds** | Pointer to **int64** | Lifetime, in seconds, of multi-document transactions. Atlas considers the transactions that exceed this limit as expired and so aborts them through a periodic cleanup process. | [optional] 

## Methods

### NewClusterDescriptionProcessArgs20240805

`func NewClusterDescriptionProcessArgs20240805() *ClusterDescriptionProcessArgs20240805`

NewClusterDescriptionProcessArgs20240805 instantiates a new ClusterDescriptionProcessArgs20240805 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDescriptionProcessArgs20240805WithDefaults

`func NewClusterDescriptionProcessArgs20240805WithDefaults() *ClusterDescriptionProcessArgs20240805`

NewClusterDescriptionProcessArgs20240805WithDefaults instantiates a new ClusterDescriptionProcessArgs20240805 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds

`func (o *ClusterDescriptionProcessArgs20240805) GetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() int`

GetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds returns the ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds field if non-nil, zero value otherwise.

### GetChangeStreamOptionsPreAndPostImagesExpireAfterSecondsOk

`func (o *ClusterDescriptionProcessArgs20240805) GetChangeStreamOptionsPreAndPostImagesExpireAfterSecondsOk() (*int, bool)`

GetChangeStreamOptionsPreAndPostImagesExpireAfterSecondsOk returns a tuple with the ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds

`func (o *ClusterDescriptionProcessArgs20240805) SetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds(v int)`

SetChangeStreamOptionsPreAndPostImagesExpireAfterSeconds sets ChangeStreamOptionsPreAndPostImagesExpireAfterSeconds field to given value.

### HasChangeStreamOptionsPreAndPostImagesExpireAfterSeconds

`func (o *ClusterDescriptionProcessArgs20240805) HasChangeStreamOptionsPreAndPostImagesExpireAfterSeconds() bool`

HasChangeStreamOptionsPreAndPostImagesExpireAfterSeconds returns a boolean if a field has been set.
### GetChunkMigrationConcurrency

`func (o *ClusterDescriptionProcessArgs20240805) GetChunkMigrationConcurrency() int`

GetChunkMigrationConcurrency returns the ChunkMigrationConcurrency field if non-nil, zero value otherwise.

### GetChunkMigrationConcurrencyOk

`func (o *ClusterDescriptionProcessArgs20240805) GetChunkMigrationConcurrencyOk() (*int, bool)`

GetChunkMigrationConcurrencyOk returns a tuple with the ChunkMigrationConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkMigrationConcurrency

`func (o *ClusterDescriptionProcessArgs20240805) SetChunkMigrationConcurrency(v int)`

SetChunkMigrationConcurrency sets ChunkMigrationConcurrency field to given value.

### HasChunkMigrationConcurrency

`func (o *ClusterDescriptionProcessArgs20240805) HasChunkMigrationConcurrency() bool`

HasChunkMigrationConcurrency returns a boolean if a field has been set.
### GetCustomOpensslCipherConfigTls12

`func (o *ClusterDescriptionProcessArgs20240805) GetCustomOpensslCipherConfigTls12() []string`

GetCustomOpensslCipherConfigTls12 returns the CustomOpensslCipherConfigTls12 field if non-nil, zero value otherwise.

### GetCustomOpensslCipherConfigTls12Ok

`func (o *ClusterDescriptionProcessArgs20240805) GetCustomOpensslCipherConfigTls12Ok() (*[]string, bool)`

GetCustomOpensslCipherConfigTls12Ok returns a tuple with the CustomOpensslCipherConfigTls12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOpensslCipherConfigTls12

`func (o *ClusterDescriptionProcessArgs20240805) SetCustomOpensslCipherConfigTls12(v []string)`

SetCustomOpensslCipherConfigTls12 sets CustomOpensslCipherConfigTls12 field to given value.

### HasCustomOpensslCipherConfigTls12

`func (o *ClusterDescriptionProcessArgs20240805) HasCustomOpensslCipherConfigTls12() bool`

HasCustomOpensslCipherConfigTls12 returns a boolean if a field has been set.
### GetCustomOpensslCipherConfigTls13

`func (o *ClusterDescriptionProcessArgs20240805) GetCustomOpensslCipherConfigTls13() []string`

GetCustomOpensslCipherConfigTls13 returns the CustomOpensslCipherConfigTls13 field if non-nil, zero value otherwise.

### GetCustomOpensslCipherConfigTls13Ok

`func (o *ClusterDescriptionProcessArgs20240805) GetCustomOpensslCipherConfigTls13Ok() (*[]string, bool)`

GetCustomOpensslCipherConfigTls13Ok returns a tuple with the CustomOpensslCipherConfigTls13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOpensslCipherConfigTls13

`func (o *ClusterDescriptionProcessArgs20240805) SetCustomOpensslCipherConfigTls13(v []string)`

SetCustomOpensslCipherConfigTls13 sets CustomOpensslCipherConfigTls13 field to given value.

### HasCustomOpensslCipherConfigTls13

`func (o *ClusterDescriptionProcessArgs20240805) HasCustomOpensslCipherConfigTls13() bool`

HasCustomOpensslCipherConfigTls13 returns a boolean if a field has been set.
### GetDefaultMaxTimeMS

`func (o *ClusterDescriptionProcessArgs20240805) GetDefaultMaxTimeMS() int`

GetDefaultMaxTimeMS returns the DefaultMaxTimeMS field if non-nil, zero value otherwise.

### GetDefaultMaxTimeMSOk

`func (o *ClusterDescriptionProcessArgs20240805) GetDefaultMaxTimeMSOk() (*int, bool)`

GetDefaultMaxTimeMSOk returns a tuple with the DefaultMaxTimeMS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMaxTimeMS

`func (o *ClusterDescriptionProcessArgs20240805) SetDefaultMaxTimeMS(v int)`

SetDefaultMaxTimeMS sets DefaultMaxTimeMS field to given value.

### HasDefaultMaxTimeMS

`func (o *ClusterDescriptionProcessArgs20240805) HasDefaultMaxTimeMS() bool`

HasDefaultMaxTimeMS returns a boolean if a field has been set.
### GetDefaultWriteConcern

`func (o *ClusterDescriptionProcessArgs20240805) GetDefaultWriteConcern() string`

GetDefaultWriteConcern returns the DefaultWriteConcern field if non-nil, zero value otherwise.

### GetDefaultWriteConcernOk

`func (o *ClusterDescriptionProcessArgs20240805) GetDefaultWriteConcernOk() (*string, bool)`

GetDefaultWriteConcernOk returns a tuple with the DefaultWriteConcern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWriteConcern

`func (o *ClusterDescriptionProcessArgs20240805) SetDefaultWriteConcern(v string)`

SetDefaultWriteConcern sets DefaultWriteConcern field to given value.

### HasDefaultWriteConcern

`func (o *ClusterDescriptionProcessArgs20240805) HasDefaultWriteConcern() bool`

HasDefaultWriteConcern returns a boolean if a field has been set.
### GetJavascriptEnabled

`func (o *ClusterDescriptionProcessArgs20240805) GetJavascriptEnabled() bool`

GetJavascriptEnabled returns the JavascriptEnabled field if non-nil, zero value otherwise.

### GetJavascriptEnabledOk

`func (o *ClusterDescriptionProcessArgs20240805) GetJavascriptEnabledOk() (*bool, bool)`

GetJavascriptEnabledOk returns a tuple with the JavascriptEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptEnabled

`func (o *ClusterDescriptionProcessArgs20240805) SetJavascriptEnabled(v bool)`

SetJavascriptEnabled sets JavascriptEnabled field to given value.

### HasJavascriptEnabled

`func (o *ClusterDescriptionProcessArgs20240805) HasJavascriptEnabled() bool`

HasJavascriptEnabled returns a boolean if a field has been set.
### GetMinimumEnabledTlsProtocol

`func (o *ClusterDescriptionProcessArgs20240805) GetMinimumEnabledTlsProtocol() string`

GetMinimumEnabledTlsProtocol returns the MinimumEnabledTlsProtocol field if non-nil, zero value otherwise.

### GetMinimumEnabledTlsProtocolOk

`func (o *ClusterDescriptionProcessArgs20240805) GetMinimumEnabledTlsProtocolOk() (*string, bool)`

GetMinimumEnabledTlsProtocolOk returns a tuple with the MinimumEnabledTlsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumEnabledTlsProtocol

`func (o *ClusterDescriptionProcessArgs20240805) SetMinimumEnabledTlsProtocol(v string)`

SetMinimumEnabledTlsProtocol sets MinimumEnabledTlsProtocol field to given value.

### HasMinimumEnabledTlsProtocol

`func (o *ClusterDescriptionProcessArgs20240805) HasMinimumEnabledTlsProtocol() bool`

HasMinimumEnabledTlsProtocol returns a boolean if a field has been set.
### GetNoTableScan

`func (o *ClusterDescriptionProcessArgs20240805) GetNoTableScan() bool`

GetNoTableScan returns the NoTableScan field if non-nil, zero value otherwise.

### GetNoTableScanOk

`func (o *ClusterDescriptionProcessArgs20240805) GetNoTableScanOk() (*bool, bool)`

GetNoTableScanOk returns a tuple with the NoTableScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTableScan

`func (o *ClusterDescriptionProcessArgs20240805) SetNoTableScan(v bool)`

SetNoTableScan sets NoTableScan field to given value.

### HasNoTableScan

`func (o *ClusterDescriptionProcessArgs20240805) HasNoTableScan() bool`

HasNoTableScan returns a boolean if a field has been set.
### GetOplogMinRetentionHours

`func (o *ClusterDescriptionProcessArgs20240805) GetOplogMinRetentionHours() float64`

GetOplogMinRetentionHours returns the OplogMinRetentionHours field if non-nil, zero value otherwise.

### GetOplogMinRetentionHoursOk

`func (o *ClusterDescriptionProcessArgs20240805) GetOplogMinRetentionHoursOk() (*float64, bool)`

GetOplogMinRetentionHoursOk returns a tuple with the OplogMinRetentionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogMinRetentionHours

`func (o *ClusterDescriptionProcessArgs20240805) SetOplogMinRetentionHours(v float64)`

SetOplogMinRetentionHours sets OplogMinRetentionHours field to given value.

### HasOplogMinRetentionHours

`func (o *ClusterDescriptionProcessArgs20240805) HasOplogMinRetentionHours() bool`

HasOplogMinRetentionHours returns a boolean if a field has been set.
### GetOplogSizeMB

`func (o *ClusterDescriptionProcessArgs20240805) GetOplogSizeMB() int`

GetOplogSizeMB returns the OplogSizeMB field if non-nil, zero value otherwise.

### GetOplogSizeMBOk

`func (o *ClusterDescriptionProcessArgs20240805) GetOplogSizeMBOk() (*int, bool)`

GetOplogSizeMBOk returns a tuple with the OplogSizeMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOplogSizeMB

`func (o *ClusterDescriptionProcessArgs20240805) SetOplogSizeMB(v int)`

SetOplogSizeMB sets OplogSizeMB field to given value.

### HasOplogSizeMB

`func (o *ClusterDescriptionProcessArgs20240805) HasOplogSizeMB() bool`

HasOplogSizeMB returns a boolean if a field has been set.
### GetQueryStatsLogVerbosity

`func (o *ClusterDescriptionProcessArgs20240805) GetQueryStatsLogVerbosity() int`

GetQueryStatsLogVerbosity returns the QueryStatsLogVerbosity field if non-nil, zero value otherwise.

### GetQueryStatsLogVerbosityOk

`func (o *ClusterDescriptionProcessArgs20240805) GetQueryStatsLogVerbosityOk() (*int, bool)`

GetQueryStatsLogVerbosityOk returns a tuple with the QueryStatsLogVerbosity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryStatsLogVerbosity

`func (o *ClusterDescriptionProcessArgs20240805) SetQueryStatsLogVerbosity(v int)`

SetQueryStatsLogVerbosity sets QueryStatsLogVerbosity field to given value.

### HasQueryStatsLogVerbosity

`func (o *ClusterDescriptionProcessArgs20240805) HasQueryStatsLogVerbosity() bool`

HasQueryStatsLogVerbosity returns a boolean if a field has been set.
### GetSampleRefreshIntervalBIConnector

`func (o *ClusterDescriptionProcessArgs20240805) GetSampleRefreshIntervalBIConnector() int`

GetSampleRefreshIntervalBIConnector returns the SampleRefreshIntervalBIConnector field if non-nil, zero value otherwise.

### GetSampleRefreshIntervalBIConnectorOk

`func (o *ClusterDescriptionProcessArgs20240805) GetSampleRefreshIntervalBIConnectorOk() (*int, bool)`

GetSampleRefreshIntervalBIConnectorOk returns a tuple with the SampleRefreshIntervalBIConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRefreshIntervalBIConnector

`func (o *ClusterDescriptionProcessArgs20240805) SetSampleRefreshIntervalBIConnector(v int)`

SetSampleRefreshIntervalBIConnector sets SampleRefreshIntervalBIConnector field to given value.

### HasSampleRefreshIntervalBIConnector

`func (o *ClusterDescriptionProcessArgs20240805) HasSampleRefreshIntervalBIConnector() bool`

HasSampleRefreshIntervalBIConnector returns a boolean if a field has been set.
### GetSampleSizeBIConnector

`func (o *ClusterDescriptionProcessArgs20240805) GetSampleSizeBIConnector() int`

GetSampleSizeBIConnector returns the SampleSizeBIConnector field if non-nil, zero value otherwise.

### GetSampleSizeBIConnectorOk

`func (o *ClusterDescriptionProcessArgs20240805) GetSampleSizeBIConnectorOk() (*int, bool)`

GetSampleSizeBIConnectorOk returns a tuple with the SampleSizeBIConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleSizeBIConnector

`func (o *ClusterDescriptionProcessArgs20240805) SetSampleSizeBIConnector(v int)`

SetSampleSizeBIConnector sets SampleSizeBIConnector field to given value.

### HasSampleSizeBIConnector

`func (o *ClusterDescriptionProcessArgs20240805) HasSampleSizeBIConnector() bool`

HasSampleSizeBIConnector returns a boolean if a field has been set.
### GetTlsCipherConfigMode

`func (o *ClusterDescriptionProcessArgs20240805) GetTlsCipherConfigMode() string`

GetTlsCipherConfigMode returns the TlsCipherConfigMode field if non-nil, zero value otherwise.

### GetTlsCipherConfigModeOk

`func (o *ClusterDescriptionProcessArgs20240805) GetTlsCipherConfigModeOk() (*string, bool)`

GetTlsCipherConfigModeOk returns a tuple with the TlsCipherConfigMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherConfigMode

`func (o *ClusterDescriptionProcessArgs20240805) SetTlsCipherConfigMode(v string)`

SetTlsCipherConfigMode sets TlsCipherConfigMode field to given value.

### HasTlsCipherConfigMode

`func (o *ClusterDescriptionProcessArgs20240805) HasTlsCipherConfigMode() bool`

HasTlsCipherConfigMode returns a boolean if a field has been set.
### GetTransactionLifetimeLimitSeconds

`func (o *ClusterDescriptionProcessArgs20240805) GetTransactionLifetimeLimitSeconds() int64`

GetTransactionLifetimeLimitSeconds returns the TransactionLifetimeLimitSeconds field if non-nil, zero value otherwise.

### GetTransactionLifetimeLimitSecondsOk

`func (o *ClusterDescriptionProcessArgs20240805) GetTransactionLifetimeLimitSecondsOk() (*int64, bool)`

GetTransactionLifetimeLimitSecondsOk returns a tuple with the TransactionLifetimeLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionLifetimeLimitSeconds

`func (o *ClusterDescriptionProcessArgs20240805) SetTransactionLifetimeLimitSeconds(v int64)`

SetTransactionLifetimeLimitSeconds sets TransactionLifetimeLimitSeconds field to given value.

### HasTransactionLifetimeLimitSeconds

`func (o *ClusterDescriptionProcessArgs20240805) HasTransactionLifetimeLimitSeconds() bool`

HasTransactionLifetimeLimitSeconds returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


