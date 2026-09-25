// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterBackupSubscriptionResponse Paid-backup subscription state for a free (M0) cluster.
type ClusterBackupSubscriptionResponse struct {
	// Flag that indicates whether paid backups are enabled for the free (M0) cluster.
	Enabled bool `json:"enabled"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *ClusterBackupSubscriptionResponse) MarshalJSON() ([]byte, error) {
	type noMethod ClusterBackupSubscriptionResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewClusterBackupSubscriptionResponse instantiates a new ClusterBackupSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterBackupSubscriptionResponse(enabled bool) *ClusterBackupSubscriptionResponse {
	this := ClusterBackupSubscriptionResponse{}
	this.Enabled = enabled
	return &this
}

// NewClusterBackupSubscriptionResponseWithDefaults instantiates a new ClusterBackupSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterBackupSubscriptionResponseWithDefaults() *ClusterBackupSubscriptionResponse {
	this := ClusterBackupSubscriptionResponse{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *ClusterBackupSubscriptionResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ClusterBackupSubscriptionResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ClusterBackupSubscriptionResponse) SetEnabled(v bool) {
	o.Enabled = v
}
