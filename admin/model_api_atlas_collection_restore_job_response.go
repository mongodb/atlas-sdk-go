// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// ApiAtlasCollectionRestoreJobResponse Collection restore job summary including the list of databases and collections in the restore scope (up to 100 items each).
type ApiAtlasCollectionRestoreJobResponse struct {
	// Suffix applied to restored collection names.
	CollectionSuffix *string `json:"collectionSuffix,omitempty"`
	// List of collections in the restore scope (up to 100 items).
	Collections *[]ApiAtlasRestoreNamespace `json:"collections,omitempty"`
	// Date and time when the restore job was created (ISO 8601 format in UTC).
	// Read only field.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Suffix applied to restored database names.
	DatabaseSuffix *string `json:"databaseSuffix,omitempty"`
	// List of databases in the restore scope (up to 100 items).
	Databases *[]ApiAtlasRestoreNamespace `json:"databases,omitempty"`
	// Error message when the job has failed or been canceled.
	// Read only field.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Date and time when the restore job finished (ISO 8601 format in UTC).
	// Read only field.
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the collection restore job.
	// Read only field.
	Id *string `json:"id,omitempty"`
	// Strategy for restoring indexes (all, none, or all except TTL).
	IndexStrategy *string `json:"indexStrategy,omitempty"`
	// Oplog increment for point-in-time restore.
	OplogInc *int `json:"oplogInc,omitempty"`
	// Oplog timestamp (seconds part) for point-in-time restore.
	OplogTs *int `json:"oplogTs,omitempty"`
	// Strategy for overwriting existing data (new if exists or overwrite if exists).
	OverwriteStrategy *string `json:"overwriteStrategy,omitempty"`
	// Point-in-time restore time in seconds since UNIX epoch.
	PointInTimeUtcSeconds *int `json:"pointInTimeUtcSeconds,omitempty"`
	// Number of documents restored so far across all supported collections.
	// Read only field.
	RestoredDocuments *int64 `json:"restoredDocuments,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the snapshot being restored.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// Current state of the collection restore job.
	// Read only field.
	State *string `json:"state,omitempty"`
	// Human-readable label that identifies the target cluster.
	TargetClusterName *string `json:"targetClusterName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the target group.
	TargetGroupId *string `json:"targetGroupId,omitempty"`
	// Total number of documents across all supported collections in the restore job. This value may initially reflect an estimate based on collection metadata and can change as accurate document counts become available during the restore.
	// Read only field.
	TotalDocuments *int64 `json:"totalDocuments,omitempty"`
}

// NewApiAtlasCollectionRestoreJobResponse instantiates a new ApiAtlasCollectionRestoreJobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCollectionRestoreJobResponse() *ApiAtlasCollectionRestoreJobResponse {
	this := ApiAtlasCollectionRestoreJobResponse{}
	return &this
}

// NewApiAtlasCollectionRestoreJobResponseWithDefaults instantiates a new ApiAtlasCollectionRestoreJobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCollectionRestoreJobResponseWithDefaults() *ApiAtlasCollectionRestoreJobResponse {
	this := ApiAtlasCollectionRestoreJobResponse{}
	return &this
}

// GetCollectionSuffix returns the CollectionSuffix field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetCollectionSuffix() string {
	if o == nil || IsNil(o.CollectionSuffix) {
		var ret string
		return ret
	}
	return *o.CollectionSuffix
}

// GetCollectionSuffixOk returns a tuple with the CollectionSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetCollectionSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.CollectionSuffix) {
		return nil, false
	}

	return o.CollectionSuffix, true
}

// HasCollectionSuffix returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasCollectionSuffix() bool {
	if o != nil && !IsNil(o.CollectionSuffix) {
		return true
	}

	return false
}

// SetCollectionSuffix gets a reference to the given string and assigns it to the CollectionSuffix field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetCollectionSuffix(v string) {
	o.CollectionSuffix = &v
}

// GetCollections returns the Collections field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetCollections() []ApiAtlasRestoreNamespace {
	if o == nil || IsNil(o.Collections) {
		var ret []ApiAtlasRestoreNamespace
		return ret
	}
	return *o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetCollectionsOk() (*[]ApiAtlasRestoreNamespace, bool) {
	if o == nil || IsNil(o.Collections) {
		return nil, false
	}

	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasCollections() bool {
	if o != nil && !IsNil(o.Collections) {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []ApiAtlasRestoreNamespace and assigns it to the Collections field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetCollections(v []ApiAtlasRestoreNamespace) {
	o.Collections = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}

	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDatabaseSuffix returns the DatabaseSuffix field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabaseSuffix() string {
	if o == nil || IsNil(o.DatabaseSuffix) {
		var ret string
		return ret
	}
	return *o.DatabaseSuffix
}

// GetDatabaseSuffixOk returns a tuple with the DatabaseSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabaseSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseSuffix) {
		return nil, false
	}

	return o.DatabaseSuffix, true
}

// HasDatabaseSuffix returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasDatabaseSuffix() bool {
	if o != nil && !IsNil(o.DatabaseSuffix) {
		return true
	}

	return false
}

// SetDatabaseSuffix gets a reference to the given string and assigns it to the DatabaseSuffix field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetDatabaseSuffix(v string) {
	o.DatabaseSuffix = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabases() []ApiAtlasRestoreNamespace {
	if o == nil || IsNil(o.Databases) {
		var ret []ApiAtlasRestoreNamespace
		return ret
	}
	return *o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetDatabasesOk() (*[]ApiAtlasRestoreNamespace, bool) {
	if o == nil || IsNil(o.Databases) {
		return nil, false
	}

	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasDatabases() bool {
	if o != nil && !IsNil(o.Databases) {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []ApiAtlasRestoreNamespace and assigns it to the Databases field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetDatabases(v []ApiAtlasRestoreNamespace) {
	o.Databases = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetFinishedAt() time.Time {
	if o == nil || IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}

	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetId(v string) {
	o.Id = &v
}

// GetIndexStrategy returns the IndexStrategy field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetIndexStrategy() string {
	if o == nil || IsNil(o.IndexStrategy) {
		var ret string
		return ret
	}
	return *o.IndexStrategy
}

// GetIndexStrategyOk returns a tuple with the IndexStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetIndexStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.IndexStrategy) {
		return nil, false
	}

	return o.IndexStrategy, true
}

// HasIndexStrategy returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasIndexStrategy() bool {
	if o != nil && !IsNil(o.IndexStrategy) {
		return true
	}

	return false
}

// SetIndexStrategy gets a reference to the given string and assigns it to the IndexStrategy field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetIndexStrategy(v string) {
	o.IndexStrategy = &v
}

// GetOplogInc returns the OplogInc field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogInc() int {
	if o == nil || IsNil(o.OplogInc) {
		var ret int
		return ret
	}
	return *o.OplogInc
}

// GetOplogIncOk returns a tuple with the OplogInc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogIncOk() (*int, bool) {
	if o == nil || IsNil(o.OplogInc) {
		return nil, false
	}

	return o.OplogInc, true
}

// HasOplogInc returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasOplogInc() bool {
	if o != nil && !IsNil(o.OplogInc) {
		return true
	}

	return false
}

// SetOplogInc gets a reference to the given int and assigns it to the OplogInc field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetOplogInc(v int) {
	o.OplogInc = &v
}

// GetOplogTs returns the OplogTs field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogTs() int {
	if o == nil || IsNil(o.OplogTs) {
		var ret int
		return ret
	}
	return *o.OplogTs
}

// GetOplogTsOk returns a tuple with the OplogTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetOplogTsOk() (*int, bool) {
	if o == nil || IsNil(o.OplogTs) {
		return nil, false
	}

	return o.OplogTs, true
}

// HasOplogTs returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasOplogTs() bool {
	if o != nil && !IsNil(o.OplogTs) {
		return true
	}

	return false
}

// SetOplogTs gets a reference to the given int and assigns it to the OplogTs field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetOplogTs(v int) {
	o.OplogTs = &v
}

// GetOverwriteStrategy returns the OverwriteStrategy field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetOverwriteStrategy() string {
	if o == nil || IsNil(o.OverwriteStrategy) {
		var ret string
		return ret
	}
	return *o.OverwriteStrategy
}

// GetOverwriteStrategyOk returns a tuple with the OverwriteStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetOverwriteStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.OverwriteStrategy) {
		return nil, false
	}

	return o.OverwriteStrategy, true
}

// HasOverwriteStrategy returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasOverwriteStrategy() bool {
	if o != nil && !IsNil(o.OverwriteStrategy) {
		return true
	}

	return false
}

// SetOverwriteStrategy gets a reference to the given string and assigns it to the OverwriteStrategy field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetOverwriteStrategy(v string) {
	o.OverwriteStrategy = &v
}

// GetPointInTimeUtcSeconds returns the PointInTimeUtcSeconds field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetPointInTimeUtcSeconds() int {
	if o == nil || IsNil(o.PointInTimeUtcSeconds) {
		var ret int
		return ret
	}
	return *o.PointInTimeUtcSeconds
}

// GetPointInTimeUtcSecondsOk returns a tuple with the PointInTimeUtcSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetPointInTimeUtcSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.PointInTimeUtcSeconds) {
		return nil, false
	}

	return o.PointInTimeUtcSeconds, true
}

// HasPointInTimeUtcSeconds returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasPointInTimeUtcSeconds() bool {
	if o != nil && !IsNil(o.PointInTimeUtcSeconds) {
		return true
	}

	return false
}

// SetPointInTimeUtcSeconds gets a reference to the given int and assigns it to the PointInTimeUtcSeconds field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetPointInTimeUtcSeconds(v int) {
	o.PointInTimeUtcSeconds = &v
}

// GetRestoredDocuments returns the RestoredDocuments field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetRestoredDocuments() int64 {
	if o == nil || IsNil(o.RestoredDocuments) {
		var ret int64
		return ret
	}
	return *o.RestoredDocuments
}

// GetRestoredDocumentsOk returns a tuple with the RestoredDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetRestoredDocumentsOk() (*int64, bool) {
	if o == nil || IsNil(o.RestoredDocuments) {
		return nil, false
	}

	return o.RestoredDocuments, true
}

// HasRestoredDocuments returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasRestoredDocuments() bool {
	if o != nil && !IsNil(o.RestoredDocuments) {
		return true
	}

	return false
}

// SetRestoredDocuments gets a reference to the given int64 and assigns it to the RestoredDocuments field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetRestoredDocuments(v int64) {
	o.RestoredDocuments = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}

	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetState returns the State field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetState(v string) {
	o.State = &v
}

// GetTargetClusterName returns the TargetClusterName field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetClusterName() string {
	if o == nil || IsNil(o.TargetClusterName) {
		var ret string
		return ret
	}
	return *o.TargetClusterName
}

// GetTargetClusterNameOk returns a tuple with the TargetClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetClusterNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetClusterName) {
		return nil, false
	}

	return o.TargetClusterName, true
}

// HasTargetClusterName returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasTargetClusterName() bool {
	if o != nil && !IsNil(o.TargetClusterName) {
		return true
	}

	return false
}

// SetTargetClusterName gets a reference to the given string and assigns it to the TargetClusterName field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetTargetClusterName(v string) {
	o.TargetClusterName = &v
}

// GetTargetGroupId returns the TargetGroupId field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetGroupId() string {
	if o == nil || IsNil(o.TargetGroupId) {
		var ret string
		return ret
	}
	return *o.TargetGroupId
}

// GetTargetGroupIdOk returns a tuple with the TargetGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetTargetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetGroupId) {
		return nil, false
	}

	return o.TargetGroupId, true
}

// HasTargetGroupId returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasTargetGroupId() bool {
	if o != nil && !IsNil(o.TargetGroupId) {
		return true
	}

	return false
}

// SetTargetGroupId gets a reference to the given string and assigns it to the TargetGroupId field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetTargetGroupId(v string) {
	o.TargetGroupId = &v
}

// GetTotalDocuments returns the TotalDocuments field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobResponse) GetTotalDocuments() int64 {
	if o == nil || IsNil(o.TotalDocuments) {
		var ret int64
		return ret
	}
	return *o.TotalDocuments
}

// GetTotalDocumentsOk returns a tuple with the TotalDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) GetTotalDocumentsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalDocuments) {
		return nil, false
	}

	return o.TotalDocuments, true
}

// HasTotalDocuments returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobResponse) HasTotalDocuments() bool {
	if o != nil && !IsNil(o.TotalDocuments) {
		return true
	}

	return false
}

// SetTotalDocuments gets a reference to the given int64 and assigns it to the TotalDocuments field.
func (o *ApiAtlasCollectionRestoreJobResponse) SetTotalDocuments(v int64) {
	o.TotalDocuments = &v
}
