// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterBackupSubscriptionUpdateRequest Paid-backup subscription state to apply to a free (M0) cluster. Omit enabled to leave the subscription unchanged.
type ClusterBackupSubscriptionUpdateRequest struct {
	// Flag that indicates whether paid backups are enabled for the free (M0) cluster. Omit this field to leave the subscription unchanged.
	Enabled *bool `json:"enabled,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ClusterBackupSubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	type noMethod ClusterBackupSubscriptionUpdateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewClusterBackupSubscriptionUpdateRequest instantiates a new ClusterBackupSubscriptionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterBackupSubscriptionUpdateRequest() *ClusterBackupSubscriptionUpdateRequest {
	this := ClusterBackupSubscriptionUpdateRequest{}
	return &this
}

// NewClusterBackupSubscriptionUpdateRequestWithDefaults instantiates a new ClusterBackupSubscriptionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterBackupSubscriptionUpdateRequestWithDefaults() *ClusterBackupSubscriptionUpdateRequest {
	this := ClusterBackupSubscriptionUpdateRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *ClusterBackupSubscriptionUpdateRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterBackupSubscriptionUpdateRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ClusterBackupSubscriptionUpdateRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ClusterBackupSubscriptionUpdateRequest) SetEnabled(v bool) {
	o.Enabled = &v
	o.NullFields = removeNullField(o.NullFields, "Enabled")
}

// SetEnabledNil sets Enabled to an explicit JSON null when marshaled.
func (o *ClusterBackupSubscriptionUpdateRequest) SetEnabledNil() {
	o.Enabled = nil
	o.NullFields = addNullField(o.NullFields, "Enabled")
}
