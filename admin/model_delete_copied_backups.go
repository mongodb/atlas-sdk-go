// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DeleteCopiedBackups Deleted copy setting whose backup copies need to also be deleted.
type DeleteCopiedBackups struct {
	// Human-readable label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete.
	// Write only field.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// Target region for the deleted copy setting whose backup copies you want to delete. Please supply the 'Atlas Region'.
	// Write only field.
	RegionName *string `json:"regionName,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica setclusters, there is only one zone in the cluster. To find the Replication Spec Id, do a GET request to Return One Cluster from One Project and consult the replicationSpecs array.
	// Write only field.
	ReplicationSpecId *string `json:"replicationSpecId,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Zone Id, do a GET request to Return One Cluster from One Project and consult the replicationSpecs array.
	// Write only field.
	ZoneId *string `json:"zoneId,omitempty"`
}

// NewDeleteCopiedBackups instantiates a new DeleteCopiedBackups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCopiedBackups() *DeleteCopiedBackups {
	this := DeleteCopiedBackups{}
	return &this
}

// NewDeleteCopiedBackupsWithDefaults instantiates a new DeleteCopiedBackups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCopiedBackupsWithDefaults() *DeleteCopiedBackups {
	this := DeleteCopiedBackups{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *DeleteCopiedBackups) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DeleteCopiedBackups) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *DeleteCopiedBackups) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *DeleteCopiedBackups) SetRegionName(v string) {
	o.RegionName = &v
}

// GetReplicationSpecId returns the ReplicationSpecId field value if set, zero value otherwise
func (o *DeleteCopiedBackups) GetReplicationSpecId() string {
	if o == nil || IsNil(o.ReplicationSpecId) {
		var ret string
		return ret
	}
	return *o.ReplicationSpecId
}

// GetReplicationSpecIdOk returns a tuple with the ReplicationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetReplicationSpecIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationSpecId) {
		return nil, false
	}

	return o.ReplicationSpecId, true
}

// HasReplicationSpecId returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasReplicationSpecId() bool {
	if o != nil && !IsNil(o.ReplicationSpecId) {
		return true
	}

	return false
}

// SetReplicationSpecId gets a reference to the given string and assigns it to the ReplicationSpecId field.
func (o *DeleteCopiedBackups) SetReplicationSpecId(v string) {
	o.ReplicationSpecId = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise
func (o *DeleteCopiedBackups) GetZoneId() string {
	if o == nil || IsNil(o.ZoneId) {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteCopiedBackups) GetZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}

	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *DeleteCopiedBackups) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *DeleteCopiedBackups) SetZoneId(v string) {
	o.ZoneId = &v
}
