// Code based on the AtlasAPI V2 OpenAPI file

package admin

// UnauthClusterCostEstimateRequest Minimal cluster description used by the unauthenticated cost-estimate endpoint. The server looks up the Cluster Starter Template matching the requested tier, applies its standard defaults, and prices the resulting configuration.
type UnauthClusterCostEstimateRequest struct {
	// Cloud service provider to estimate cost in.
	CloudProvider string `json:"cloudProvider"`
	// Cluster tier to estimate. Only a fixed set of representative tiers is supported on the unauthenticated endpoint (M0, FLEX, M10, M30); for other tiers, see https://www.mongodb.com/pricing or use the authenticated cost-estimate endpoints.
	InstanceSize string `json:"instanceSize"`
	// Cloud-provider region to estimate cost in. Region naming follows the chosen provider's convention (for example, US_EAST_1 for AWS).
	RegionName string `json:"regionName"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *UnauthClusterCostEstimateRequest) MarshalJSON() ([]byte, error) {
	type noMethod UnauthClusterCostEstimateRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewUnauthClusterCostEstimateRequest instantiates a new UnauthClusterCostEstimateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthClusterCostEstimateRequest(cloudProvider string, instanceSize string, regionName string) *UnauthClusterCostEstimateRequest {
	this := UnauthClusterCostEstimateRequest{}
	this.CloudProvider = cloudProvider
	this.InstanceSize = instanceSize
	this.RegionName = regionName
	return &this
}

// NewUnauthClusterCostEstimateRequestWithDefaults instantiates a new UnauthClusterCostEstimateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthClusterCostEstimateRequestWithDefaults() *UnauthClusterCostEstimateRequest {
	this := UnauthClusterCostEstimateRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *UnauthClusterCostEstimateRequest) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *UnauthClusterCostEstimateRequest) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetInstanceSize returns the InstanceSize field value
func (o *UnauthClusterCostEstimateRequest) GetInstanceSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceSize
}

// GetInstanceSizeOk returns a tuple with the InstanceSize field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateRequest) GetInstanceSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceSize, true
}

// SetInstanceSize sets field value
func (o *UnauthClusterCostEstimateRequest) SetInstanceSize(v string) {
	o.InstanceSize = v
}

// GetRegionName returns the RegionName field value
func (o *UnauthClusterCostEstimateRequest) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *UnauthClusterCostEstimateRequest) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *UnauthClusterCostEstimateRequest) SetRegionName(v string) {
	o.RegionName = v
}
