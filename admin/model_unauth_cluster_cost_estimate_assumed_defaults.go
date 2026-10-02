// Code based on the AtlasAPI V2 OpenAPI file

package admin

// UnauthClusterCostEstimateAssumedDefaults Standard defaults the server applied to produce a complete cluster configuration that could be priced. The reported cost reflects exactly this configuration. Omitted for free (M0) and Flex tiers, where the priced configuration is fully determined by the tier.
type UnauthClusterCostEstimateAssumedDefaults struct {
	// Topology the estimate assumes.
	// Read only field.
	ClusterType string `json:"clusterType"`
	// Disk size, in GB, the estimate assumes.
	// Read only field.
	DiskSizeGb float64 `json:"diskSizeGb"`
	// Number of electable nodes the estimate assumes.
	// Read only field.
	ElectableNodeCount int `json:"electableNodeCount"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *UnauthClusterCostEstimateAssumedDefaults) MarshalJSON() ([]byte, error) {
	type noMethod UnauthClusterCostEstimateAssumedDefaults
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewUnauthClusterCostEstimateAssumedDefaults instantiates a new UnauthClusterCostEstimateAssumedDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthClusterCostEstimateAssumedDefaults(clusterType string, diskSizeGb float64, electableNodeCount int) *UnauthClusterCostEstimateAssumedDefaults {
	this := UnauthClusterCostEstimateAssumedDefaults{}
	this.ClusterType = clusterType
	this.DiskSizeGb = diskSizeGb
	this.ElectableNodeCount = electableNodeCount
	return &this
}

// NewUnauthClusterCostEstimateAssumedDefaultsWithDefaults instantiates a new UnauthClusterCostEstimateAssumedDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthClusterCostEstimateAssumedDefaultsWithDefaults() *UnauthClusterCostEstimateAssumedDefaults {
	this := UnauthClusterCostEstimateAssumedDefaults{}
	return &this
}

// GetClusterType returns the ClusterType field value
func (o *UnauthClusterCostEstimateAssumedDefaults) GetClusterType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateAssumedDefaults) GetClusterTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterType, true
}

// SetClusterType sets field value
func (o *UnauthClusterCostEstimateAssumedDefaults) SetClusterType(v string) {
	o.ClusterType = v
}

// GetDiskSizeGb returns the DiskSizeGb field value
func (o *UnauthClusterCostEstimateAssumedDefaults) GetDiskSizeGb() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DiskSizeGb
}

// GetDiskSizeGbOk returns a tuple with the DiskSizeGb field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateAssumedDefaults) GetDiskSizeGbOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiskSizeGb, true
}

// SetDiskSizeGb sets field value
func (o *UnauthClusterCostEstimateAssumedDefaults) SetDiskSizeGb(v float64) {
	o.DiskSizeGb = v
}

// GetElectableNodeCount returns the ElectableNodeCount field value
func (o *UnauthClusterCostEstimateAssumedDefaults) GetElectableNodeCount() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.ElectableNodeCount
}

// GetElectableNodeCountOk returns a tuple with the ElectableNodeCount field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateAssumedDefaults) GetElectableNodeCountOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElectableNodeCount, true
}

// SetElectableNodeCount sets field value
func (o *UnauthClusterCostEstimateAssumedDefaults) SetElectableNodeCount(v int) {
	o.ElectableNodeCount = v
}
