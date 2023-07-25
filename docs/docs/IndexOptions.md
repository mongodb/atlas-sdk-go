# IndexOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var2dsphereIndexVersion** | Pointer to **int** | Index version number applied to the 2dsphere index. MongoDB 3.2 and later use version 3. Use this option to override the default version number. This option applies to the **2dsphere** index type only. | [optional] [default to 3]
**Background** | Pointer to **bool** | Flag that indicates whether MongoDB should build the index in the background. This applies to MongoDB databases running feature compatibility version 4.0 or earlier. MongoDB databases running FCV 4.2 or later build indexes using an optimized build process. This process holds the exclusive lock only at the beginning and end of the build process. The rest of the build process yields to interleaving read and write operations. MongoDB databases running FCV 4.2 or later ignore this option. This option applies to all index types. | [optional] [default to false]
**Bits** | Pointer to **int** | Number of precision applied to the stored geohash value of the location data. This option applies to the **2d** index type only. | [optional] [default to 26]
**BucketSize** | Pointer to **int** | Number of units within which to group the location values. You could group in the same bucket those location values within the specified number of units to each other. This option applies to the geoHaystack index type only.  MongoDB 5.0 removed geoHaystack Indexes and the &#x60;geoSearch&#x60; command. | [optional] 
**ColumnstoreProjection** | Pointer to **map[string]int** | The columnstoreProjection document allows to include or exclude subschemas schema. One cannot combine inclusion and exclusion statements. Accordingly, the &lt;value&gt; can be either of the following: 1 or true to include the field and recursively all fields it is a prefix of in the index 0 or false to exclude the field and recursively all fields it is a prefix of from the index. | [optional] 
**DefaultLanguage** | Pointer to **string** | Human language that determines the list of stop words and the rules for the stemmer and tokenizer. This option accepts the supported languages using its name in lowercase english or the ISO 639-2 code. If you set this parameter to &#x60;\&quot;none\&quot;&#x60;, then the text search uses simple tokenization with no list of stop words and no stemming. This option applies to the **text** index type only. | [optional] [default to "english"]
**ExpireAfterSeconds** | Pointer to **int** | Number of seconds that MongoDB retains documents in a Time To Live (TTL) index. | [optional] 
**Hidden** | Pointer to **bool** | Flag that determines whether the index is hidden from the query planner. A hidden index is not evaluated as part of the query plan selection. | [optional] [default to false]
**LanguageOverride** | Pointer to **string** | Human-readable label that identifies the document parameter that contains the override language for the document. This option applies to the **text** index type only. | [optional] [default to "language"]
**Max** | Pointer to **int** | Upper inclusive boundary to limit the longitude and latitude values. This option applies to the 2d index type only. | [optional] [default to 180]
**Min** | Pointer to **int** | Lower inclusive boundary to limit the longitude and latitude values. This option applies to the 2d index type only. | [optional] [default to -180]
**Name** | Pointer to **string** | Human-readable label that identifies this index. This option applies to all index types. | [optional] 
**PartialFilterExpression** | Pointer to **map[string]interface{}** | Rules that limit the documents that the index references to a filter expression. All MongoDB index types accept a **partialFilterExpression** option. **partialFilterExpression** can include following expressions:  - equality (&#x60;\&quot;parameter\&quot; : \&quot;value\&quot;&#x60; or using the &#x60;$eq&#x60; operator) - &#x60;\&quot;$exists\&quot;: true&#x60; , maximum: &#x60;$gt&#x60;, &#x60;$gte&#x60;, &#x60;$lt&#x60;, &#x60;$lte&#x60; comparisons - &#x60;$type&#x60; - &#x60;$and&#x60; (top-level only)  This option applies to all index types. | [optional] 
**Sparse** | Pointer to **bool** | Flag that indicates whether the index references documents that only have the specified parameter. These indexes use less space but behave differently in some situations like when sorting. The following index types default to sparse and ignore this option: &#x60;2dsphere&#x60;, &#x60;2d&#x60;, &#x60;geoHaystack&#x60;, &#x60;text&#x60;.  Compound indexes that includes one or more indexes with &#x60;2dsphere&#x60; keys alongside other key types, only the &#x60;2dsphere&#x60; index parameters determine which documents the index references. If you run MongoDB 3.2 or later, use partial indexes. This option applies to all index types. | [optional] [default to false]
**StorageEngine** | Pointer to **map[string]interface{}** | Storage engine set for the specific index. This value can be set only at creation. This option uses the following format: &#x60;\&quot;storageEngine\&quot; : { \&quot;&lt;storage-engine-name&gt;\&quot; : \&quot;&lt;options&gt;\&quot; }&#x60; MongoDB validates storage engine configuration options when creating indexes. To support replica sets with members with different storage engines, MongoDB logs these options to the oplog during replication. This option applies to all index types. | [optional] 
**TextIndexVersion** | Pointer to **int** | Version applied to this text index. MongoDB 3.2 and later use version &#x60;3&#x60;. Use this option to override the default version number. This option applies to the **text** index type only. | [optional] [default to 3]
**Unique** | Pointer to **bool** | Flag that indicates whether this index can accept insertion or update of documents when the index key value matches an existing index key value. Set &#x60;\&quot;unique\&quot; : true&#x60; to set this index as unique. You can&#39;t set a hashed index to be unique. This option applies to all index types. This option is unsupported for rolling indexes. | [optional] [default to false]
**Weights** | Pointer to **map[string]interface{}** | Relative importance to place upon provided index parameters. This object expresses this as key/value pairs of index parameter and weight to apply to that parameter. You can specify weights for some or all the indexed parameters. The weight must be an integer between 1 and 99,999. MongoDB 5.0 and later can apply **weights** to **text** indexes only. | [optional] 

## Methods

### NewIndexOptions

`func NewIndexOptions() *IndexOptions`

NewIndexOptions instantiates a new IndexOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndexOptionsWithDefaults

`func NewIndexOptionsWithDefaults() *IndexOptions`

NewIndexOptionsWithDefaults instantiates a new IndexOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar2dsphereIndexVersion

`func (o *IndexOptions) GetVar2dsphereIndexVersion() int`

GetVar2dsphereIndexVersion returns the Var2dsphereIndexVersion field if non-nil, zero value otherwise.

### GetVar2dsphereIndexVersionOk

`func (o *IndexOptions) GetVar2dsphereIndexVersionOk() (*int, bool)`

GetVar2dsphereIndexVersionOk returns a tuple with the Var2dsphereIndexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar2dsphereIndexVersion

`func (o *IndexOptions) SetVar2dsphereIndexVersion(v int)`

SetVar2dsphereIndexVersion sets Var2dsphereIndexVersion field to given value.

### HasVar2dsphereIndexVersion

`func (o *IndexOptions) HasVar2dsphereIndexVersion() bool`

HasVar2dsphereIndexVersion returns a boolean if a field has been set.
### GetBackground

`func (o *IndexOptions) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *IndexOptions) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *IndexOptions) SetBackground(v bool)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *IndexOptions) HasBackground() bool`

HasBackground returns a boolean if a field has been set.
### GetBits

`func (o *IndexOptions) GetBits() int`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *IndexOptions) GetBitsOk() (*int, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *IndexOptions) SetBits(v int)`

SetBits sets Bits field to given value.

### HasBits

`func (o *IndexOptions) HasBits() bool`

HasBits returns a boolean if a field has been set.
### GetBucketSize

`func (o *IndexOptions) GetBucketSize() int`

GetBucketSize returns the BucketSize field if non-nil, zero value otherwise.

### GetBucketSizeOk

`func (o *IndexOptions) GetBucketSizeOk() (*int, bool)`

GetBucketSizeOk returns a tuple with the BucketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketSize

`func (o *IndexOptions) SetBucketSize(v int)`

SetBucketSize sets BucketSize field to given value.

### HasBucketSize

`func (o *IndexOptions) HasBucketSize() bool`

HasBucketSize returns a boolean if a field has been set.
### GetColumnstoreProjection

`func (o *IndexOptions) GetColumnstoreProjection() map[string]int`

GetColumnstoreProjection returns the ColumnstoreProjection field if non-nil, zero value otherwise.

### GetColumnstoreProjectionOk

`func (o *IndexOptions) GetColumnstoreProjectionOk() (*map[string]int, bool)`

GetColumnstoreProjectionOk returns a tuple with the ColumnstoreProjection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnstoreProjection

`func (o *IndexOptions) SetColumnstoreProjection(v map[string]int)`

SetColumnstoreProjection sets ColumnstoreProjection field to given value.

### HasColumnstoreProjection

`func (o *IndexOptions) HasColumnstoreProjection() bool`

HasColumnstoreProjection returns a boolean if a field has been set.
### GetDefaultLanguage

`func (o *IndexOptions) GetDefaultLanguage() string`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *IndexOptions) GetDefaultLanguageOk() (*string, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *IndexOptions) SetDefaultLanguage(v string)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *IndexOptions) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.
### GetExpireAfterSeconds

`func (o *IndexOptions) GetExpireAfterSeconds() int`

GetExpireAfterSeconds returns the ExpireAfterSeconds field if non-nil, zero value otherwise.

### GetExpireAfterSecondsOk

`func (o *IndexOptions) GetExpireAfterSecondsOk() (*int, bool)`

GetExpireAfterSecondsOk returns a tuple with the ExpireAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAfterSeconds

`func (o *IndexOptions) SetExpireAfterSeconds(v int)`

SetExpireAfterSeconds sets ExpireAfterSeconds field to given value.

### HasExpireAfterSeconds

`func (o *IndexOptions) HasExpireAfterSeconds() bool`

HasExpireAfterSeconds returns a boolean if a field has been set.
### GetHidden

`func (o *IndexOptions) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *IndexOptions) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *IndexOptions) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *IndexOptions) HasHidden() bool`

HasHidden returns a boolean if a field has been set.
### GetLanguageOverride

`func (o *IndexOptions) GetLanguageOverride() string`

GetLanguageOverride returns the LanguageOverride field if non-nil, zero value otherwise.

### GetLanguageOverrideOk

`func (o *IndexOptions) GetLanguageOverrideOk() (*string, bool)`

GetLanguageOverrideOk returns a tuple with the LanguageOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageOverride

`func (o *IndexOptions) SetLanguageOverride(v string)`

SetLanguageOverride sets LanguageOverride field to given value.

### HasLanguageOverride

`func (o *IndexOptions) HasLanguageOverride() bool`

HasLanguageOverride returns a boolean if a field has been set.
### GetMax

`func (o *IndexOptions) GetMax() int`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *IndexOptions) GetMaxOk() (*int, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *IndexOptions) SetMax(v int)`

SetMax sets Max field to given value.

### HasMax

`func (o *IndexOptions) HasMax() bool`

HasMax returns a boolean if a field has been set.
### GetMin

`func (o *IndexOptions) GetMin() int`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *IndexOptions) GetMinOk() (*int, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *IndexOptions) SetMin(v int)`

SetMin sets Min field to given value.

### HasMin

`func (o *IndexOptions) HasMin() bool`

HasMin returns a boolean if a field has been set.
### GetName

`func (o *IndexOptions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IndexOptions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IndexOptions) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IndexOptions) HasName() bool`

HasName returns a boolean if a field has been set.
### GetPartialFilterExpression

`func (o *IndexOptions) GetPartialFilterExpression() map[string]interface{}`

GetPartialFilterExpression returns the PartialFilterExpression field if non-nil, zero value otherwise.

### GetPartialFilterExpressionOk

`func (o *IndexOptions) GetPartialFilterExpressionOk() (*map[string]interface{}, bool)`

GetPartialFilterExpressionOk returns a tuple with the PartialFilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFilterExpression

`func (o *IndexOptions) SetPartialFilterExpression(v map[string]interface{})`

SetPartialFilterExpression sets PartialFilterExpression field to given value.

### HasPartialFilterExpression

`func (o *IndexOptions) HasPartialFilterExpression() bool`

HasPartialFilterExpression returns a boolean if a field has been set.
### GetSparse

`func (o *IndexOptions) GetSparse() bool`

GetSparse returns the Sparse field if non-nil, zero value otherwise.

### GetSparseOk

`func (o *IndexOptions) GetSparseOk() (*bool, bool)`

GetSparseOk returns a tuple with the Sparse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparse

`func (o *IndexOptions) SetSparse(v bool)`

SetSparse sets Sparse field to given value.

### HasSparse

`func (o *IndexOptions) HasSparse() bool`

HasSparse returns a boolean if a field has been set.
### GetStorageEngine

`func (o *IndexOptions) GetStorageEngine() map[string]interface{}`

GetStorageEngine returns the StorageEngine field if non-nil, zero value otherwise.

### GetStorageEngineOk

`func (o *IndexOptions) GetStorageEngineOk() (*map[string]interface{}, bool)`

GetStorageEngineOk returns a tuple with the StorageEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageEngine

`func (o *IndexOptions) SetStorageEngine(v map[string]interface{})`

SetStorageEngine sets StorageEngine field to given value.

### HasStorageEngine

`func (o *IndexOptions) HasStorageEngine() bool`

HasStorageEngine returns a boolean if a field has been set.
### GetTextIndexVersion

`func (o *IndexOptions) GetTextIndexVersion() int`

GetTextIndexVersion returns the TextIndexVersion field if non-nil, zero value otherwise.

### GetTextIndexVersionOk

`func (o *IndexOptions) GetTextIndexVersionOk() (*int, bool)`

GetTextIndexVersionOk returns a tuple with the TextIndexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextIndexVersion

`func (o *IndexOptions) SetTextIndexVersion(v int)`

SetTextIndexVersion sets TextIndexVersion field to given value.

### HasTextIndexVersion

`func (o *IndexOptions) HasTextIndexVersion() bool`

HasTextIndexVersion returns a boolean if a field has been set.
### GetUnique

`func (o *IndexOptions) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *IndexOptions) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *IndexOptions) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *IndexOptions) HasUnique() bool`

HasUnique returns a boolean if a field has been set.
### GetWeights

`func (o *IndexOptions) GetWeights() map[string]interface{}`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *IndexOptions) GetWeightsOk() (*map[string]interface{}, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *IndexOptions) SetWeights(v map[string]interface{})`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *IndexOptions) HasWeights() bool`

HasWeights returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


