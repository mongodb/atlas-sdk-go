// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasCollectionRestoreJobIndexStatus Overall index build status for a collection restore job.
type ApiAtlasCollectionRestoreJobIndexStatus struct {
	// Number of collections that failed to build indexes.
	FailedCollectionCount *int `json:"failedCollectionCount,omitempty"`
	// Index build state indicating the status of index creation during or after a restore operation.
	State *string `json:"state,omitempty"`
}

// NewApiAtlasCollectionRestoreJobIndexStatus instantiates a new ApiAtlasCollectionRestoreJobIndexStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCollectionRestoreJobIndexStatus() *ApiAtlasCollectionRestoreJobIndexStatus {
	this := ApiAtlasCollectionRestoreJobIndexStatus{}
	return &this
}

// NewApiAtlasCollectionRestoreJobIndexStatusWithDefaults instantiates a new ApiAtlasCollectionRestoreJobIndexStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCollectionRestoreJobIndexStatusWithDefaults() *ApiAtlasCollectionRestoreJobIndexStatus {
	this := ApiAtlasCollectionRestoreJobIndexStatus{}
	return &this
}

// GetFailedCollectionCount returns the FailedCollectionCount field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetFailedCollectionCount() int {
	if o == nil || IsNil(o.FailedCollectionCount) {
		var ret int
		return ret
	}
	return *o.FailedCollectionCount
}

// GetFailedCollectionCountOk returns a tuple with the FailedCollectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetFailedCollectionCountOk() (*int, bool) {
	if o == nil || IsNil(o.FailedCollectionCount) {
		return nil, false
	}

	return o.FailedCollectionCount, true
}

// HasFailedCollectionCount returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobIndexStatus) HasFailedCollectionCount() bool {
	if o != nil && !IsNil(o.FailedCollectionCount) {
		return true
	}

	return false
}

// SetFailedCollectionCount gets a reference to the given int and assigns it to the FailedCollectionCount field.
func (o *ApiAtlasCollectionRestoreJobIndexStatus) SetFailedCollectionCount(v int) {
	o.FailedCollectionCount = &v
}

// GetState returns the State field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreJobIndexStatus) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreJobIndexStatus) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiAtlasCollectionRestoreJobIndexStatus) SetState(v string) {
	o.State = &v
}
