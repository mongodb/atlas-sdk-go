// Code based on the AtlasAPI V2 OpenAPI file

package admin

// DiskBackupCopySetting20240805 Copy setting item in the desired backup policy.
type DiskBackupCopySetting20240805 struct {
	// Human-readable label that identifies the cloud provider that stores the snapshot copy.
	CloudProvider *string `json:"cloudProvider,omitempty"`
	// List that describes which types of snapshots to copy.
	Frequencies *[]string `json:"frequencies,omitempty"`
	// Target region to copy snapshots belonging to zoneId. Please supply the 'Atlas Region' which can be found under [Cloud Providers](https://www.mongodb.com/docs/atlas/reference/cloud-providers/) 'regions' link.
	RegionName *string `json:"regionName,omitempty"`
	// Flag that indicates whether to copy the oplogs to the target region. You can use the oplogs to perform point-in-time restores.
	ShouldCopyOplogs *bool `json:"shouldCopyOplogs,omitempty"`
	// Unique 24-hexadecimal digit string that identifies the zone in a cluster. For global clusters, there can be multiple zones to choose from. For sharded clusters and replica set clusters, there is only one zone in the cluster. To find the Zone Id, do a GET request to Return One Cluster from One Project and consult the replicationSpecs array [Return One Cluster From One Project](#operation/getCluster).
	ZoneId string `json:"zoneId"`
}

// NewDiskBackupCopySetting20240805 instantiates a new DiskBackupCopySetting20240805 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskBackupCopySetting20240805(zoneId string) *DiskBackupCopySetting20240805 {
	this := DiskBackupCopySetting20240805{}
	this.ZoneId = zoneId
	return &this
}

// NewDiskBackupCopySetting20240805WithDefaults instantiates a new DiskBackupCopySetting20240805 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskBackupCopySetting20240805WithDefaults() *DiskBackupCopySetting20240805 {
	this := DiskBackupCopySetting20240805{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *DiskBackupCopySetting20240805) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting20240805) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *DiskBackupCopySetting20240805) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *DiskBackupCopySetting20240805) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetFrequencies returns the Frequencies field value if set, zero value otherwise
func (o *DiskBackupCopySetting20240805) GetFrequencies() []string {
	if o == nil || IsNil(o.Frequencies) {
		var ret []string
		return ret
	}
	return *o.Frequencies
}

// GetFrequenciesOk returns a tuple with the Frequencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting20240805) GetFrequenciesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Frequencies) {
		return nil, false
	}

	return o.Frequencies, true
}

// HasFrequencies returns a boolean if a field has been set.
func (o *DiskBackupCopySetting20240805) HasFrequencies() bool {
	if o != nil && !IsNil(o.Frequencies) {
		return true
	}

	return false
}

// SetFrequencies gets a reference to the given []string and assigns it to the Frequencies field.
func (o *DiskBackupCopySetting20240805) SetFrequencies(v []string) {
	o.Frequencies = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise
func (o *DiskBackupCopySetting20240805) GetRegionName() string {
	if o == nil || IsNil(o.RegionName) {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting20240805) GetRegionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RegionName) {
		return nil, false
	}

	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *DiskBackupCopySetting20240805) HasRegionName() bool {
	if o != nil && !IsNil(o.RegionName) {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *DiskBackupCopySetting20240805) SetRegionName(v string) {
	o.RegionName = &v
}

// GetShouldCopyOplogs returns the ShouldCopyOplogs field value if set, zero value otherwise
func (o *DiskBackupCopySetting20240805) GetShouldCopyOplogs() bool {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		var ret bool
		return ret
	}
	return *o.ShouldCopyOplogs
}

// GetShouldCopyOplogsOk returns a tuple with the ShouldCopyOplogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting20240805) GetShouldCopyOplogsOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldCopyOplogs) {
		return nil, false
	}

	return o.ShouldCopyOplogs, true
}

// HasShouldCopyOplogs returns a boolean if a field has been set.
func (o *DiskBackupCopySetting20240805) HasShouldCopyOplogs() bool {
	if o != nil && !IsNil(o.ShouldCopyOplogs) {
		return true
	}

	return false
}

// SetShouldCopyOplogs gets a reference to the given bool and assigns it to the ShouldCopyOplogs field.
func (o *DiskBackupCopySetting20240805) SetShouldCopyOplogs(v bool) {
	o.ShouldCopyOplogs = &v
}

// GetZoneId returns the ZoneId field value
func (o *DiskBackupCopySetting20240805) GetZoneId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *DiskBackupCopySetting20240805) GetZoneIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *DiskBackupCopySetting20240805) SetZoneId(v string) {
	o.ZoneId = v
}
