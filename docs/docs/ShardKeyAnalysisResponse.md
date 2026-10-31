# ShardKeyAnalysisResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvgDocSizeBytes** | Pointer to **int64** | Average size of one document in the collection, in bytes. | [optional] [readonly] 
**CorrelationCoefficient** | Pointer to **float64** | Correlation between the candidate shard key and the order in which documents were inserted, from &#x60;-1&#x60; to &#x60;1&#x60;. Values near &#x60;1&#x60; or &#x60;-1&#x60; indicate a monotonically changing key, which directs all inserts to one shard. This parameter is absent when the shard key has no supporting index or the collection is clustered, in which case MongoDB Cloud does not compute it. | [optional] [readonly] 
**CreatedAt** | **time.Time** | Date and time when this shard key analysis was requested (ISO 8601 format in UTC). This is when the analysis was submitted, not when it completed; the completion time is the &#x60;updatedAt&#x60; of the operation that produced it. | [readonly] 
**ExpiresAt** | **time.Time** | Date and time when MongoDB Cloud stops returning this shard key analysis (ISO 8601 format in UTC). Reads after this time return an HTTP 404 status. | [readonly] 
**Id** | **string** | Unique 24-hexadecimal digit string that identifies this shard key analysis. This is the same identifier as the &#x60;operationId&#x60; of the operation that produced it. | [readonly] 
**IsUnique** | Pointer to **bool** | Flag that indicates whether the index supporting the candidate shard key enforces a uniqueness constraint. | [optional] [readonly] 
**Monotonicity** | Pointer to **string** | Whether the shard key value increases or decreases with the order in which documents were inserted. A monotonically changing shard key directs all inserts to one shard. &#x60;UNKNOWN&#x60; indicates that MongoDB Cloud could not determine monotonicity, which happens when the shard key has no supporting index or the collection is clustered. | [optional] [readonly] 
**MostCommonValues** | Pointer to [**[]ShardKeyMostCommonValue**](ShardKeyMostCommonValue.md) | Shard key values that occur most frequently in the sampled documents, with how often each occurs. MongoDB Cloud returns this parameter only to callers that can read the collection&#39;s data, because it carries values out of the collection rather than statistics about them. | [optional] [readonly] 
**Namespace** | **string** | Combination of the database and collection that was analyzed, written as &#x60;&lt;database&gt;.&lt;collection&gt;&#x60;. | 
**Note** | Pointer to **string** | Guidance from MongoDB on how to interpret the metrics in this analysis. This parameter is absent when the analysis carries no guidance. | [optional] [readonly] 
**NumDistinctValues** | Pointer to **int64** | Number of distinct values that the candidate shard key takes across the sampled documents. A value close to &#x60;numDocsSampled&#x60; indicates high cardinality, which distributes documents well. | [optional] [readonly] 
**NumDocsSampled** | Pointer to **int64** | Number of documents that MongoDB Cloud sampled to compute these metrics. | [optional] [readonly] 
**NumDocsTotal** | Pointer to **int64** | Number of documents in the collection. | [optional] [readonly] 
**NumOrphanDocs** | Pointer to **int64** | Number of orphaned documents in the collection. This parameter is absent for collections that have none. | [optional] [readonly] 
**PercentageOfMultiShardReads** | Pointer to **float64** | Percentage of sampled reads that the candidate shard key would route to more than one shard but not to every shard. | [optional] [readonly] 
**PercentageOfMultiShardWrites** | Pointer to **float64** | Percentage of sampled writes that the candidate shard key would route to more than one shard but not to every shard. | [optional] [readonly] 
**PercentageOfScatterGatherReads** | Pointer to **float64** | Percentage of sampled reads that the candidate shard key would broadcast to every shard. Lower is better. | [optional] [readonly] 
**PercentageOfScatterGatherWrites** | Pointer to **float64** | Percentage of sampled writes that the candidate shard key would broadcast to every shard. Lower is better. | [optional] [readonly] 
**PercentageOfSingleShardReads** | Pointer to **float64** | Percentage of sampled reads that the candidate shard key would route to exactly one shard. Higher is better. | [optional] [readonly] 
**PercentageOfSingleShardWrites** | Pointer to **float64** | Percentage of sampled writes that the candidate shard key would route to exactly one shard. Higher is better. | [optional] [readonly] 
**RawOutput** | Pointer to **string** | Unmodified result of the &#x60;analyzeShardKey&#x60; command, as a JSON string. Its content tracks the MongoDB server version of the analyzed cluster rather than the version of this resource, so members outside the parameters above may appear or change without a new resource version. Do not parse it as a stable contract. &#x60;mostCommonValues&#x60; is removed from it for callers that cannot read the collection&#39;s data. This parameter is absent if MongoDB Cloud could not serialize the command result, which costs the raw view but not the analysis. | [optional] [readonly] 
**ReadSampleSize** | Pointer to [**ShardKeyReadSampleSize**](ShardKeyReadSampleSize.md) |  | [optional] 
**ShardKey** | Pointer to **[]map[string]string** | Candidate shard key that was analyzed, given as one object per shard key field in the order the fields form the key. This parameter is absent when MongoDB Cloud has no shard key recorded for the analysis. | [optional] 
**WriteSampleSize** | Pointer to [**ShardKeyWriteSampleSize**](ShardKeyWriteSampleSize.md) |  | [optional] 

## Methods

### NewShardKeyAnalysisResponse

`func NewShardKeyAnalysisResponse(createdAt time.Time, expiresAt time.Time, id string, namespace string, ) *ShardKeyAnalysisResponse`

NewShardKeyAnalysisResponse instantiates a new ShardKeyAnalysisResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShardKeyAnalysisResponseWithDefaults

`func NewShardKeyAnalysisResponseWithDefaults() *ShardKeyAnalysisResponse`

NewShardKeyAnalysisResponseWithDefaults instantiates a new ShardKeyAnalysisResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvgDocSizeBytes

`func (o *ShardKeyAnalysisResponse) GetAvgDocSizeBytes() int64`

GetAvgDocSizeBytes returns the AvgDocSizeBytes field if non-nil, zero value otherwise.

### GetAvgDocSizeBytesOk

`func (o *ShardKeyAnalysisResponse) GetAvgDocSizeBytesOk() (*int64, bool)`

GetAvgDocSizeBytesOk returns a tuple with the AvgDocSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDocSizeBytes

`func (o *ShardKeyAnalysisResponse) SetAvgDocSizeBytes(v int64)`

SetAvgDocSizeBytes sets AvgDocSizeBytes field to given value.

### HasAvgDocSizeBytes

`func (o *ShardKeyAnalysisResponse) HasAvgDocSizeBytes() bool`

HasAvgDocSizeBytes returns a boolean if a field has been set.

### SetAvgDocSizeBytesNil

`func (o *ShardKeyAnalysisResponse) SetAvgDocSizeBytesNil()`

SetAvgDocSizeBytesNil sets AvgDocSizeBytes to an explicit JSON null when marshaled, overriding any value previously set with SetAvgDocSizeBytes. Calling SetAvgDocSizeBytes again clears the null override.

### GetCorrelationCoefficient

`func (o *ShardKeyAnalysisResponse) GetCorrelationCoefficient() float64`

GetCorrelationCoefficient returns the CorrelationCoefficient field if non-nil, zero value otherwise.

### GetCorrelationCoefficientOk

`func (o *ShardKeyAnalysisResponse) GetCorrelationCoefficientOk() (*float64, bool)`

GetCorrelationCoefficientOk returns a tuple with the CorrelationCoefficient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationCoefficient

`func (o *ShardKeyAnalysisResponse) SetCorrelationCoefficient(v float64)`

SetCorrelationCoefficient sets CorrelationCoefficient field to given value.

### HasCorrelationCoefficient

`func (o *ShardKeyAnalysisResponse) HasCorrelationCoefficient() bool`

HasCorrelationCoefficient returns a boolean if a field has been set.

### SetCorrelationCoefficientNil

`func (o *ShardKeyAnalysisResponse) SetCorrelationCoefficientNil()`

SetCorrelationCoefficientNil sets CorrelationCoefficient to an explicit JSON null when marshaled, overriding any value previously set with SetCorrelationCoefficient. Calling SetCorrelationCoefficient again clears the null override.

### GetCreatedAt

`func (o *ShardKeyAnalysisResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ShardKeyAnalysisResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ShardKeyAnalysisResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetExpiresAt

`func (o *ShardKeyAnalysisResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ShardKeyAnalysisResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ShardKeyAnalysisResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### GetId

`func (o *ShardKeyAnalysisResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShardKeyAnalysisResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShardKeyAnalysisResponse) SetId(v string)`

SetId sets Id field to given value.

### GetIsUnique

`func (o *ShardKeyAnalysisResponse) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *ShardKeyAnalysisResponse) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *ShardKeyAnalysisResponse) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.

### HasIsUnique

`func (o *ShardKeyAnalysisResponse) HasIsUnique() bool`

HasIsUnique returns a boolean if a field has been set.

### SetIsUniqueNil

`func (o *ShardKeyAnalysisResponse) SetIsUniqueNil()`

SetIsUniqueNil sets IsUnique to an explicit JSON null when marshaled, overriding any value previously set with SetIsUnique. Calling SetIsUnique again clears the null override.

### GetMonotonicity

`func (o *ShardKeyAnalysisResponse) GetMonotonicity() string`

GetMonotonicity returns the Monotonicity field if non-nil, zero value otherwise.

### GetMonotonicityOk

`func (o *ShardKeyAnalysisResponse) GetMonotonicityOk() (*string, bool)`

GetMonotonicityOk returns a tuple with the Monotonicity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonotonicity

`func (o *ShardKeyAnalysisResponse) SetMonotonicity(v string)`

SetMonotonicity sets Monotonicity field to given value.

### HasMonotonicity

`func (o *ShardKeyAnalysisResponse) HasMonotonicity() bool`

HasMonotonicity returns a boolean if a field has been set.

### SetMonotonicityNil

`func (o *ShardKeyAnalysisResponse) SetMonotonicityNil()`

SetMonotonicityNil sets Monotonicity to an explicit JSON null when marshaled, overriding any value previously set with SetMonotonicity. Calling SetMonotonicity again clears the null override.

### GetMostCommonValues

`func (o *ShardKeyAnalysisResponse) GetMostCommonValues() []ShardKeyMostCommonValue`

GetMostCommonValues returns the MostCommonValues field if non-nil, zero value otherwise.

### GetMostCommonValuesOk

`func (o *ShardKeyAnalysisResponse) GetMostCommonValuesOk() (*[]ShardKeyMostCommonValue, bool)`

GetMostCommonValuesOk returns a tuple with the MostCommonValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostCommonValues

`func (o *ShardKeyAnalysisResponse) SetMostCommonValues(v []ShardKeyMostCommonValue)`

SetMostCommonValues sets MostCommonValues field to given value.

### HasMostCommonValues

`func (o *ShardKeyAnalysisResponse) HasMostCommonValues() bool`

HasMostCommonValues returns a boolean if a field has been set.

### SetMostCommonValuesNil

`func (o *ShardKeyAnalysisResponse) SetMostCommonValuesNil()`

SetMostCommonValuesNil sets MostCommonValues to an explicit JSON null when marshaled, overriding any value previously set with SetMostCommonValues. Calling SetMostCommonValues again clears the null override.

### GetNamespace

`func (o *ShardKeyAnalysisResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ShardKeyAnalysisResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ShardKeyAnalysisResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### GetNote

`func (o *ShardKeyAnalysisResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ShardKeyAnalysisResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ShardKeyAnalysisResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ShardKeyAnalysisResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ShardKeyAnalysisResponse) SetNoteNil()`

SetNoteNil sets Note to an explicit JSON null when marshaled, overriding any value previously set with SetNote. Calling SetNote again clears the null override.

### GetNumDistinctValues

`func (o *ShardKeyAnalysisResponse) GetNumDistinctValues() int64`

GetNumDistinctValues returns the NumDistinctValues field if non-nil, zero value otherwise.

### GetNumDistinctValuesOk

`func (o *ShardKeyAnalysisResponse) GetNumDistinctValuesOk() (*int64, bool)`

GetNumDistinctValuesOk returns a tuple with the NumDistinctValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDistinctValues

`func (o *ShardKeyAnalysisResponse) SetNumDistinctValues(v int64)`

SetNumDistinctValues sets NumDistinctValues field to given value.

### HasNumDistinctValues

`func (o *ShardKeyAnalysisResponse) HasNumDistinctValues() bool`

HasNumDistinctValues returns a boolean if a field has been set.

### SetNumDistinctValuesNil

`func (o *ShardKeyAnalysisResponse) SetNumDistinctValuesNil()`

SetNumDistinctValuesNil sets NumDistinctValues to an explicit JSON null when marshaled, overriding any value previously set with SetNumDistinctValues. Calling SetNumDistinctValues again clears the null override.

### GetNumDocsSampled

`func (o *ShardKeyAnalysisResponse) GetNumDocsSampled() int64`

GetNumDocsSampled returns the NumDocsSampled field if non-nil, zero value otherwise.

### GetNumDocsSampledOk

`func (o *ShardKeyAnalysisResponse) GetNumDocsSampledOk() (*int64, bool)`

GetNumDocsSampledOk returns a tuple with the NumDocsSampled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDocsSampled

`func (o *ShardKeyAnalysisResponse) SetNumDocsSampled(v int64)`

SetNumDocsSampled sets NumDocsSampled field to given value.

### HasNumDocsSampled

`func (o *ShardKeyAnalysisResponse) HasNumDocsSampled() bool`

HasNumDocsSampled returns a boolean if a field has been set.

### SetNumDocsSampledNil

`func (o *ShardKeyAnalysisResponse) SetNumDocsSampledNil()`

SetNumDocsSampledNil sets NumDocsSampled to an explicit JSON null when marshaled, overriding any value previously set with SetNumDocsSampled. Calling SetNumDocsSampled again clears the null override.

### GetNumDocsTotal

`func (o *ShardKeyAnalysisResponse) GetNumDocsTotal() int64`

GetNumDocsTotal returns the NumDocsTotal field if non-nil, zero value otherwise.

### GetNumDocsTotalOk

`func (o *ShardKeyAnalysisResponse) GetNumDocsTotalOk() (*int64, bool)`

GetNumDocsTotalOk returns a tuple with the NumDocsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDocsTotal

`func (o *ShardKeyAnalysisResponse) SetNumDocsTotal(v int64)`

SetNumDocsTotal sets NumDocsTotal field to given value.

### HasNumDocsTotal

`func (o *ShardKeyAnalysisResponse) HasNumDocsTotal() bool`

HasNumDocsTotal returns a boolean if a field has been set.

### SetNumDocsTotalNil

`func (o *ShardKeyAnalysisResponse) SetNumDocsTotalNil()`

SetNumDocsTotalNil sets NumDocsTotal to an explicit JSON null when marshaled, overriding any value previously set with SetNumDocsTotal. Calling SetNumDocsTotal again clears the null override.

### GetNumOrphanDocs

`func (o *ShardKeyAnalysisResponse) GetNumOrphanDocs() int64`

GetNumOrphanDocs returns the NumOrphanDocs field if non-nil, zero value otherwise.

### GetNumOrphanDocsOk

`func (o *ShardKeyAnalysisResponse) GetNumOrphanDocsOk() (*int64, bool)`

GetNumOrphanDocsOk returns a tuple with the NumOrphanDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOrphanDocs

`func (o *ShardKeyAnalysisResponse) SetNumOrphanDocs(v int64)`

SetNumOrphanDocs sets NumOrphanDocs field to given value.

### HasNumOrphanDocs

`func (o *ShardKeyAnalysisResponse) HasNumOrphanDocs() bool`

HasNumOrphanDocs returns a boolean if a field has been set.

### SetNumOrphanDocsNil

`func (o *ShardKeyAnalysisResponse) SetNumOrphanDocsNil()`

SetNumOrphanDocsNil sets NumOrphanDocs to an explicit JSON null when marshaled, overriding any value previously set with SetNumOrphanDocs. Calling SetNumOrphanDocs again clears the null override.

### GetPercentageOfMultiShardReads

`func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardReads() float64`

GetPercentageOfMultiShardReads returns the PercentageOfMultiShardReads field if non-nil, zero value otherwise.

### GetPercentageOfMultiShardReadsOk

`func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardReadsOk() (*float64, bool)`

GetPercentageOfMultiShardReadsOk returns a tuple with the PercentageOfMultiShardReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfMultiShardReads

`func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardReads(v float64)`

SetPercentageOfMultiShardReads sets PercentageOfMultiShardReads field to given value.

### HasPercentageOfMultiShardReads

`func (o *ShardKeyAnalysisResponse) HasPercentageOfMultiShardReads() bool`

HasPercentageOfMultiShardReads returns a boolean if a field has been set.

### SetPercentageOfMultiShardReadsNil

`func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardReadsNil()`

SetPercentageOfMultiShardReadsNil sets PercentageOfMultiShardReads to an explicit JSON null when marshaled, overriding any value previously set with SetPercentageOfMultiShardReads. Calling SetPercentageOfMultiShardReads again clears the null override.

### GetPercentageOfMultiShardWrites

`func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardWrites() float64`

GetPercentageOfMultiShardWrites returns the PercentageOfMultiShardWrites field if non-nil, zero value otherwise.

### GetPercentageOfMultiShardWritesOk

`func (o *ShardKeyAnalysisResponse) GetPercentageOfMultiShardWritesOk() (*float64, bool)`

GetPercentageOfMultiShardWritesOk returns a tuple with the PercentageOfMultiShardWrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfMultiShardWrites

`func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardWrites(v float64)`

SetPercentageOfMultiShardWrites sets PercentageOfMultiShardWrites field to given value.

### HasPercentageOfMultiShardWrites

`func (o *ShardKeyAnalysisResponse) HasPercentageOfMultiShardWrites() bool`

HasPercentageOfMultiShardWrites returns a boolean if a field has been set.

### SetPercentageOfMultiShardWritesNil

`func (o *ShardKeyAnalysisResponse) SetPercentageOfMultiShardWritesNil()`

SetPercentageOfMultiShardWritesNil sets PercentageOfMultiShardWrites to an explicit JSON null when marshaled, overriding any value previously set with SetPercentageOfMultiShardWrites. Calling SetPercentageOfMultiShardWrites again clears the null override.

### GetPercentageOfScatterGatherReads

`func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherReads() float64`

GetPercentageOfScatterGatherReads returns the PercentageOfScatterGatherReads field if non-nil, zero value otherwise.

### GetPercentageOfScatterGatherReadsOk

`func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherReadsOk() (*float64, bool)`

GetPercentageOfScatterGatherReadsOk returns a tuple with the PercentageOfScatterGatherReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfScatterGatherReads

`func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherReads(v float64)`

SetPercentageOfScatterGatherReads sets PercentageOfScatterGatherReads field to given value.

### HasPercentageOfScatterGatherReads

`func (o *ShardKeyAnalysisResponse) HasPercentageOfScatterGatherReads() bool`

HasPercentageOfScatterGatherReads returns a boolean if a field has been set.

### SetPercentageOfScatterGatherReadsNil

`func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherReadsNil()`

SetPercentageOfScatterGatherReadsNil sets PercentageOfScatterGatherReads to an explicit JSON null when marshaled, overriding any value previously set with SetPercentageOfScatterGatherReads. Calling SetPercentageOfScatterGatherReads again clears the null override.

### GetPercentageOfScatterGatherWrites

`func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherWrites() float64`

GetPercentageOfScatterGatherWrites returns the PercentageOfScatterGatherWrites field if non-nil, zero value otherwise.

### GetPercentageOfScatterGatherWritesOk

`func (o *ShardKeyAnalysisResponse) GetPercentageOfScatterGatherWritesOk() (*float64, bool)`

GetPercentageOfScatterGatherWritesOk returns a tuple with the PercentageOfScatterGatherWrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfScatterGatherWrites

`func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherWrites(v float64)`

SetPercentageOfScatterGatherWrites sets PercentageOfScatterGatherWrites field to given value.

### HasPercentageOfScatterGatherWrites

`func (o *ShardKeyAnalysisResponse) HasPercentageOfScatterGatherWrites() bool`

HasPercentageOfScatterGatherWrites returns a boolean if a field has been set.

### SetPercentageOfScatterGatherWritesNil

`func (o *ShardKeyAnalysisResponse) SetPercentageOfScatterGatherWritesNil()`

SetPercentageOfScatterGatherWritesNil sets PercentageOfScatterGatherWrites to an explicit JSON null when marshaled, overriding any value previously set with SetPercentageOfScatterGatherWrites. Calling SetPercentageOfScatterGatherWrites again clears the null override.

### GetPercentageOfSingleShardReads

`func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardReads() float64`

GetPercentageOfSingleShardReads returns the PercentageOfSingleShardReads field if non-nil, zero value otherwise.

### GetPercentageOfSingleShardReadsOk

`func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardReadsOk() (*float64, bool)`

GetPercentageOfSingleShardReadsOk returns a tuple with the PercentageOfSingleShardReads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfSingleShardReads

`func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardReads(v float64)`

SetPercentageOfSingleShardReads sets PercentageOfSingleShardReads field to given value.

### HasPercentageOfSingleShardReads

`func (o *ShardKeyAnalysisResponse) HasPercentageOfSingleShardReads() bool`

HasPercentageOfSingleShardReads returns a boolean if a field has been set.

### SetPercentageOfSingleShardReadsNil

`func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardReadsNil()`

SetPercentageOfSingleShardReadsNil sets PercentageOfSingleShardReads to an explicit JSON null when marshaled, overriding any value previously set with SetPercentageOfSingleShardReads. Calling SetPercentageOfSingleShardReads again clears the null override.

### GetPercentageOfSingleShardWrites

`func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardWrites() float64`

GetPercentageOfSingleShardWrites returns the PercentageOfSingleShardWrites field if non-nil, zero value otherwise.

### GetPercentageOfSingleShardWritesOk

`func (o *ShardKeyAnalysisResponse) GetPercentageOfSingleShardWritesOk() (*float64, bool)`

GetPercentageOfSingleShardWritesOk returns a tuple with the PercentageOfSingleShardWrites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfSingleShardWrites

`func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardWrites(v float64)`

SetPercentageOfSingleShardWrites sets PercentageOfSingleShardWrites field to given value.

### HasPercentageOfSingleShardWrites

`func (o *ShardKeyAnalysisResponse) HasPercentageOfSingleShardWrites() bool`

HasPercentageOfSingleShardWrites returns a boolean if a field has been set.

### SetPercentageOfSingleShardWritesNil

`func (o *ShardKeyAnalysisResponse) SetPercentageOfSingleShardWritesNil()`

SetPercentageOfSingleShardWritesNil sets PercentageOfSingleShardWrites to an explicit JSON null when marshaled, overriding any value previously set with SetPercentageOfSingleShardWrites. Calling SetPercentageOfSingleShardWrites again clears the null override.

### GetRawOutput

`func (o *ShardKeyAnalysisResponse) GetRawOutput() string`

GetRawOutput returns the RawOutput field if non-nil, zero value otherwise.

### GetRawOutputOk

`func (o *ShardKeyAnalysisResponse) GetRawOutputOk() (*string, bool)`

GetRawOutputOk returns a tuple with the RawOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawOutput

`func (o *ShardKeyAnalysisResponse) SetRawOutput(v string)`

SetRawOutput sets RawOutput field to given value.

### HasRawOutput

`func (o *ShardKeyAnalysisResponse) HasRawOutput() bool`

HasRawOutput returns a boolean if a field has been set.

### SetRawOutputNil

`func (o *ShardKeyAnalysisResponse) SetRawOutputNil()`

SetRawOutputNil sets RawOutput to an explicit JSON null when marshaled, overriding any value previously set with SetRawOutput. Calling SetRawOutput again clears the null override.

### GetReadSampleSize

`func (o *ShardKeyAnalysisResponse) GetReadSampleSize() ShardKeyReadSampleSize`

GetReadSampleSize returns the ReadSampleSize field if non-nil, zero value otherwise.

### GetReadSampleSizeOk

`func (o *ShardKeyAnalysisResponse) GetReadSampleSizeOk() (*ShardKeyReadSampleSize, bool)`

GetReadSampleSizeOk returns a tuple with the ReadSampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadSampleSize

`func (o *ShardKeyAnalysisResponse) SetReadSampleSize(v ShardKeyReadSampleSize)`

SetReadSampleSize sets ReadSampleSize field to given value.

### HasReadSampleSize

`func (o *ShardKeyAnalysisResponse) HasReadSampleSize() bool`

HasReadSampleSize returns a boolean if a field has been set.

### SetReadSampleSizeNil

`func (o *ShardKeyAnalysisResponse) SetReadSampleSizeNil()`

SetReadSampleSizeNil sets ReadSampleSize to an explicit JSON null when marshaled, overriding any value previously set with SetReadSampleSize. Calling SetReadSampleSize again clears the null override.

### GetShardKey

`func (o *ShardKeyAnalysisResponse) GetShardKey() []map[string]string`

GetShardKey returns the ShardKey field if non-nil, zero value otherwise.

### GetShardKeyOk

`func (o *ShardKeyAnalysisResponse) GetShardKeyOk() (*[]map[string]string, bool)`

GetShardKeyOk returns a tuple with the ShardKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardKey

`func (o *ShardKeyAnalysisResponse) SetShardKey(v []map[string]string)`

SetShardKey sets ShardKey field to given value.

### HasShardKey

`func (o *ShardKeyAnalysisResponse) HasShardKey() bool`

HasShardKey returns a boolean if a field has been set.

### SetShardKeyNil

`func (o *ShardKeyAnalysisResponse) SetShardKeyNil()`

SetShardKeyNil sets ShardKey to an explicit JSON null when marshaled, overriding any value previously set with SetShardKey. Calling SetShardKey again clears the null override.

### GetWriteSampleSize

`func (o *ShardKeyAnalysisResponse) GetWriteSampleSize() ShardKeyWriteSampleSize`

GetWriteSampleSize returns the WriteSampleSize field if non-nil, zero value otherwise.

### GetWriteSampleSizeOk

`func (o *ShardKeyAnalysisResponse) GetWriteSampleSizeOk() (*ShardKeyWriteSampleSize, bool)`

GetWriteSampleSizeOk returns a tuple with the WriteSampleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteSampleSize

`func (o *ShardKeyAnalysisResponse) SetWriteSampleSize(v ShardKeyWriteSampleSize)`

SetWriteSampleSize sets WriteSampleSize field to given value.

### HasWriteSampleSize

`func (o *ShardKeyAnalysisResponse) HasWriteSampleSize() bool`

HasWriteSampleSize returns a boolean if a field has been set.

### SetWriteSampleSizeNil

`func (o *ShardKeyAnalysisResponse) SetWriteSampleSizeNil()`

SetWriteSampleSizeNil sets WriteSampleSize to an explicit JSON null when marshaled, overriding any value previously set with SetWriteSampleSize. Calling SetWriteSampleSize again clears the null override.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


