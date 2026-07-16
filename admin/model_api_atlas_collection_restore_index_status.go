// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ApiAtlasCollectionRestoreIndexStatus Index build status for a collection within a restore job.
type ApiAtlasCollectionRestoreIndexStatus struct {
	// Error message if index build failed.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// List of index specifications that failed to build (up to 64 items).
	FailedIndexes *[]map[string]any `json:"failedIndexes,omitempty"`
	// Index build state indicating the status of index creation during or after a restore operation.
	State *string `json:"state,omitempty"`
}

// NewApiAtlasCollectionRestoreIndexStatus instantiates a new ApiAtlasCollectionRestoreIndexStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAtlasCollectionRestoreIndexStatus() *ApiAtlasCollectionRestoreIndexStatus {
	this := ApiAtlasCollectionRestoreIndexStatus{}
	return &this
}

// NewApiAtlasCollectionRestoreIndexStatusWithDefaults instantiates a new ApiAtlasCollectionRestoreIndexStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAtlasCollectionRestoreIndexStatusWithDefaults() *ApiAtlasCollectionRestoreIndexStatus {
	this := ApiAtlasCollectionRestoreIndexStatus{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreIndexStatus) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreIndexStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}

	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreIndexStatus) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiAtlasCollectionRestoreIndexStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFailedIndexes returns the FailedIndexes field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreIndexStatus) GetFailedIndexes() []map[string]any {
	if o == nil || IsNil(o.FailedIndexes) {
		var ret []map[string]any
		return ret
	}
	return *o.FailedIndexes
}

// GetFailedIndexesOk returns a tuple with the FailedIndexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreIndexStatus) GetFailedIndexesOk() (*[]map[string]any, bool) {
	if o == nil || IsNil(o.FailedIndexes) {
		return nil, false
	}

	return o.FailedIndexes, true
}

// HasFailedIndexes returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreIndexStatus) HasFailedIndexes() bool {
	if o != nil && !IsNil(o.FailedIndexes) {
		return true
	}

	return false
}

// SetFailedIndexes gets a reference to the given []map[string]any and assigns it to the FailedIndexes field.
func (o *ApiAtlasCollectionRestoreIndexStatus) SetFailedIndexes(v []map[string]any) {
	o.FailedIndexes = &v
}

// GetState returns the State field value if set, zero value otherwise
func (o *ApiAtlasCollectionRestoreIndexStatus) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAtlasCollectionRestoreIndexStatus) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}

	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiAtlasCollectionRestoreIndexStatus) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiAtlasCollectionRestoreIndexStatus) SetState(v string) {
	o.State = &v
}
