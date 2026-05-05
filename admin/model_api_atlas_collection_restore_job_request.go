// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasCollectionRestoreJobRequest Request body for creating a collection restore job.
type ApiAtlasCollectionRestoreJobRequest struct {
	// Optional suffix applied to restored collection names.
	CollectionSuffix *string `json:"collectionSuffix,omitempty"`
	// List of collections to restore (up to 100 items).
	Collections *[]ApiAtlasRestoreNamespace `json:"collections,omitempty"`
	// Optional suffix applied to restored database names.
	DatabaseSuffix *string `json:"databaseSuffix,omitempty"`
	// List of databases to restore (up to 100 items).
	Databases *[]ApiAtlasRestoreNamespace `json:"databases,omitempty"`
	// Strategy for restoring indexes (all, none, or all except TTL).
	IndexStrategy string `json:"indexStrategy"`
	// Oplog increment for point-in-time restore.
	OplogInc *int `json:"oplogInc,omitempty"`
	// Oplog timestamp (seconds part) for point-in-time restore.
	OplogTs *int `json:"oplogTs,omitempty"`
	// Strategy for overwriting existing data (new if exists or overwrite if exists).
	OverwriteStrategy string `json:"overwriteStrategy"`
	// Point-in-time restore time in seconds since UNIX epoch.
	PointInTimeUtcSeconds *int `json:"pointInTimeUtcSeconds,omitempty"`
	// ID of the snapshot to restore.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// Target cluster name.
	TargetClusterName string `json:"targetClusterName"`
	// Unique 24-hexadecimal digit string that identifies the target group.
	TargetGroupId string `json:"targetGroupId"`
}

// NewApiAtlasCollectionRestoreJobRequest instantiates a new ApiAtlasCollectionRestoreJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCollectionRestoreJobRequest(indexStrategy string, overwriteStrategy string, targetClusterName string, targetGroupId string) *ApiAtlasCollectionRestoreJobRequest {
	this := ApiAtlasCollectionRestoreJobRequest{}
	this.IndexStrategy = indexStrategy
	this.OverwriteStrategy = overwriteStrategy
	this.TargetClusterName = targetClusterName
	this.TargetGroupId = targetGroupId
	return &this
}

// NewApiAtlasCollectionRestoreJobRequestWithDefaults instantiates a new ApiAtlasCollectionRestoreJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCollectionRestoreJobRequestWithDefaults() *ApiAtlasCollectionRestoreJobRequest {
	this := ApiAtlasCollectionRestoreJobRequest{}
	return &this
}

// GetCollectionSuffix returns the CollectionSuffix field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetCollectionSuffix() string {
	if o == nil || IsNil(o.CollectionSuffix) {
		var ret string
		return ret
	}
	return *o.CollectionSuffix
}

// GetCollectionSuffixOk returns a tuple with the CollectionSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetCollectionSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionSuffix) {
		return nil, false
	}

	return o.CollectionSuffix, true
}

// HasCollectionSuffix returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasCollectionSuffix() bool {
	if o != nil && !IsNil(o.CollectionSuffix) {
		return true
	}

	return false
}

// SetCollectionSuffix gets a reference to the given string and assigns it to the CollectionSuffix field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetCollectionSuffix(v string) {
	o.CollectionSuffix = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetCollections() []ApiAtlasRestoreNamespace {
	if o == nil || IsNil(o.Collections) {
		var ret []ApiAtlasRestoreNamespace
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetCollectionsOk() (*[]ApiAtlasRestoreNamespace, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}

	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []ApiAtlasRestoreNamespace and assigns it to the Collections field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetCollections(v []ApiAtlasRestoreNamespace) {
	o.Collections = &v
}

// GetDatabaseSuffix returns the DatabaseSuffix field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabaseSuffix() string {
	if o == nil || IsNil(o.DatabaseSuffix) {
		var ret string
		return ret
	}
	return *o.DatabaseSuffix
}

// GetDatabaseSuffixOk returns a tuple with the DatabaseSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabaseSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseSuffix) {
		return nil, false
	}

	return o.DatabaseSuffix, true
}

// HasDatabaseSuffix returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasDatabaseSuffix() bool {
	if o != nil && !IsNil(o.DatabaseSuffix) {
		return true
	}

	return false
}

// SetDatabaseSuffix gets a reference to the given string and assigns it to the DatabaseSuffix field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetDatabaseSuffix(v string) {
	o.DatabaseSuffix = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabases() []ApiAtlasRestoreNamespace {
	if o == nil || IsNil(o.Databases) {
		var ret []ApiAtlasRestoreNamespace
		return ret
	}
	return *o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetDatabasesOk() (*[]ApiAtlasRestoreNamespace, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}

	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []ApiAtlasRestoreNamespace and assigns it to the Databases field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetDatabases(v []ApiAtlasRestoreNamespace) {
	o.Databases = &v
}

// GetIndexStrategy returns the IndexStrategy field value
func (o *ApiAtlasCollectionRestoreJobRequest) GetIndexStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexStrategy
}

// GetIndexStrategyOk returns a tuple with the IndexStrategy field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetIndexStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexStrategy, true
}

// SetIndexStrategy sets field value
func (o *ApiAtlasCollectionRestoreJobRequest) SetIndexStrategy(v string) {
	o.IndexStrategy = v
}

// GetOplogInc returns the OplogInc field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogInc() int {
	if o == nil || IsNil(o.OplogInc) {
		var ret int
		return ret
	}
	return *o.OplogInc
}

// GetOplogIncOk returns a tuple with the OplogInc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogIncOk() (*int, bool) {
	if o == nil || IsNil(o.OplogInc) {
		return nil, false
	}

	return o.OplogInc, true
}

// HasOplogInc returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasOplogInc() bool {
	if o != nil && !IsNil(o.OplogInc) {
		return true
	}

	return false
}

// SetOplogInc gets a reference to the given int and assigns it to the OplogInc field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetOplogInc(v int) {
	o.OplogInc = &v
}

// GetOplogTs returns the OplogTs field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogTs() int {
	if o == nil || IsNil(o.OplogTs) {
		var ret int
		return ret
	}
	return *o.OplogTs
}

// GetOplogTsOk returns a tuple with the OplogTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetOplogTsOk() (*int, bool) {
	if o == nil || IsNil(o.OplogTs) {
		return nil, false
	}

	return o.OplogTs, true
}

// HasOplogTs returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasOplogTs() bool {
	if o != nil && !IsNil(o.OplogTs) {
		return true
	}

	return false
}

// SetOplogTs gets a reference to the given int and assigns it to the OplogTs field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetOplogTs(v int) {
	o.OplogTs = &v
}

// GetOverwriteStrategy returns the OverwriteStrategy field value
func (o *ApiAtlasCollectionRestoreJobRequest) GetOverwriteStrategy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OverwriteStrategy
}

// GetOverwriteStrategyOk returns a tuple with the OverwriteStrategy field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetOverwriteStrategyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverwriteStrategy, true
}

// SetOverwriteStrategy sets field value
func (o *ApiAtlasCollectionRestoreJobRequest) SetOverwriteStrategy(v string) {
	o.OverwriteStrategy = v
}

// GetPointInTimeUtcSeconds returns the PointInTimeUtcSeconds field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetPointInTimeUtcSeconds() int {
	if o == nil || IsNil(o.PointInTimeUtcSeconds) {
		var ret int
		return ret
	}
	return *o.PointInTimeUtcSeconds
}

// GetPointInTimeUtcSecondsOk returns a tuple with the PointInTimeUtcSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetPointInTimeUtcSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.PointInTimeUtcSeconds) {
		return nil, false
	}

	return o.PointInTimeUtcSeconds, true
}

// HasPointInTimeUtcSeconds returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasPointInTimeUtcSeconds() bool {
	if o != nil && !IsNil(o.PointInTimeUtcSeconds) {
		return true
	}

	return false
}

// SetPointInTimeUtcSeconds gets a reference to the given int and assigns it to the PointInTimeUtcSeconds field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetPointInTimeUtcSeconds(v int) {
	o.PointInTimeUtcSeconds = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobRequest) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}

	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ApiAtlasCollectionRestoreJobRequest) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetTargetClusterName returns the TargetClusterName field value
func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetClusterName, true
}

// SetTargetClusterName sets field value
func (o *ApiAtlasCollectionRestoreJobRequest) SetTargetClusterName(v string) {
	o.TargetClusterName = v
}

// GetTargetGroupId returns the TargetGroupId field value
func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobRequest) GetTargetGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetGroupId, true
}

// SetTargetGroupId sets field value
func (o *ApiAtlasCollectionRestoreJobRequest) SetTargetGroupId(v string) {
	o.TargetGroupId = v
}
