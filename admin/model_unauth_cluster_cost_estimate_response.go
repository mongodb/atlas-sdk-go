// Code based on the AtlasAPI V2 OpenAPI file

package admin

// UnauthClusterCostEstimateResponse Estimated hourly and monthly compute cost for a cluster of the requested tier in the requested region. Echoes the request inputs and, for dedicated tiers, the standard defaults the server applied.
type UnauthClusterCostEstimateResponse struct {
	AssumedDefaults *UnauthClusterCostEstimateAssumedDefaults `json:"assumedDefaults,omitempty"`
	// Cloud service provider to estimate cost in.
	// Read only field.
	CloudProvider string              `json:"cloudProvider"`
	CostEstimate  ClusterCostEstimate `json:"costEstimate"`
	// Cluster tier to estimate. Only a fixed set of representative tiers is supported on the unauthenticated endpoint; for other tiers, use the authenticated cost-estimate endpoints.
	// Read only field.
	InstanceSize string `json:"instanceSize"`
	// Cloud-provider region the estimate was produced for.
	// Read only field.
	RegionName string `json:"regionName"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *UnauthClusterCostEstimateResponse) MarshalJSON() ([]byte, error) {
	type noMethod UnauthClusterCostEstimateResponse
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewUnauthClusterCostEstimateResponse instantiates a new UnauthClusterCostEstimateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthClusterCostEstimateResponse(cloudProvider string, costEstimate ClusterCostEstimate, instanceSize string, regionName string) *UnauthClusterCostEstimateResponse {
	this := UnauthClusterCostEstimateResponse{}
	this.CloudProvider = cloudProvider
	this.CostEstimate = costEstimate
	this.InstanceSize = instanceSize
	this.RegionName = regionName
	return &this
}

// NewUnauthClusterCostEstimateResponseWithDefaults instantiates a new UnauthClusterCostEstimateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthClusterCostEstimateResponseWithDefaults() *UnauthClusterCostEstimateResponse {
	this := UnauthClusterCostEstimateResponse{}
	return &this
}

// GetAssumedDefaults returns the AssumedDefaults field value if set, zero value otherwise
func (o *UnauthClusterCostEstimateResponse) GetAssumedDefaults() UnauthClusterCostEstimateAssumedDefaults {
	if o == nil || IsNil(o.AssumedDefaults) {
		var ret UnauthClusterCostEstimateAssumedDefaults
		return ret
	}
	return *o.AssumedDefaults
}

// GetAssumedDefaultsOk returns a tuple with the AssumedDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateResponse) GetAssumedDefaultsOk() (*UnauthClusterCostEstimateAssumedDefaults, bool) {
	if o == nil || IsNil(o.AssumedDefaults) {
		return nil, false
	}

	return o.AssumedDefaults, true
}

// HasAssumedDefaults returns a boolean if a field has been set.
func (o *UnauthClusterCostEstimateResponse) HasAssumedDefaults() bool {
	if o != nil && !IsNil(o.AssumedDefaults) {
		return true
	}

	return false
}

// SetAssumedDefaults gets a reference to the given UnauthClusterCostEstimateAssumedDefaults and assigns it to the AssumedDefaults field.
func (o *UnauthClusterCostEstimateResponse) SetAssumedDefaults(v UnauthClusterCostEstimateAssumedDefaults) {
	o.AssumedDefaults = &v
	o.NullFields = removeNullField(o.NullFields, "AssumedDefaults")
}

// SetAssumedDefaultsNil sets AssumedDefaults to an explicit JSON null when marshaled.
func (o *UnauthClusterCostEstimateResponse) SetAssumedDefaultsNil() {
	o.AssumedDefaults = nil
	o.NullFields = addNullField(o.NullFields, "AssumedDefaults")
}

// GetCloudProvider returns the CloudProvider field value
func (o *UnauthClusterCostEstimateResponse) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateResponse) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *UnauthClusterCostEstimateResponse) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCostEstimate returns the CostEstimate field value
func (o *UnauthClusterCostEstimateResponse) GetCostEstimate() ClusterCostEstimate {
	if o == nil {
		var ret ClusterCostEstimate
		return ret
	}

	return o.CostEstimate
}

// GetCostEstimateOk returns a tuple with the CostEstimate field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateResponse) GetCostEstimateOk() (*ClusterCostEstimate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostEstimate, true
}

// SetCostEstimate sets field value
func (o *UnauthClusterCostEstimateResponse) SetCostEstimate(v ClusterCostEstimate) {
	o.CostEstimate = v
}

// GetInstanceSize returns the InstanceSize field value
func (o *UnauthClusterCostEstimateResponse) GetInstanceSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateResponse) GetInstanceSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceSize, true
}

// SetInstanceSize sets field value
func (o *UnauthClusterCostEstimateResponse) SetInstanceSize(v string) {
	o.InstanceSize = v
}

// GetRegionName returns the RegionName field value
func (o *UnauthClusterCostEstimateResponse) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateResponse) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *UnauthClusterCostEstimateResponse) SetRegionName(v string) {
	o.RegionName = v
}
