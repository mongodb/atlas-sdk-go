// Code based on the AtlasAPI V2 OpenAPI file

package admin

// FlexProviderSettingsCreate20250101 Group of cloud provider settings that configure the provisioned MongoDB flex cluster.
type FlexProviderSettingsCreate20250101 struct {
	// Cloud service provider on which MongoDB Cloud provisioned the flex cluster.
	// Write only field.
	BackingProviderName string `json:"backingProviderName"`
	// Storage capacity available to the flex cluster expressed in gigabytes.
	// Read only field.
	DiskSizeGB *float64 `json:"diskSizeGB,omitempty"`
	// Human-readable label that identifies the cloud service provider.
	// Read only field.
	ProviderName *string `json:"providerName,omitempty"`
	// Human-readable label that identifies the geographic location of your MongoDB flex cluster. The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).
	// Write only field.
	RegionName string `json:"regionName"`
}

// NewFlexProviderSettingsCreate20250101 instantiates a new FlexProviderSettingsCreate20250101 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlexProviderSettingsCreate20250101(backingProviderName string, regionName string) *FlexProviderSettingsCreate20250101 {
	this := FlexProviderSettingsCreate20250101{}
	this.BackingProviderName = backingProviderName
	this.RegionName = regionName
	return &this
}

// NewFlexProviderSettingsCreate20250101WithDefaults instantiates a new FlexProviderSettingsCreate20250101 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlexProviderSettingsCreate20250101WithDefaults() *FlexProviderSettingsCreate20250101 {
	this := FlexProviderSettingsCreate20250101{}
	return &this
}

// GetBackingProviderName returns the BackingProviderName field value
func (o *FlexProviderSettingsCreate20250101) GetBackingProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BackingProviderName
}

// GetBackingProviderNameOk returns a tuple with the BackingProviderName field value
// and a boolean to check if the value has been set.
func (o *FlexProviderSettingsCreate20250101) GetBackingProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackingProviderName, true
}

// SetBackingProviderName sets field value
func (o *FlexProviderSettingsCreate20250101) SetBackingProviderName(v string) {
	o.BackingProviderName = v
}

// GetDiskSizeGB returns the DiskSizeGB field value if set, zero value otherwise
func (o *FlexProviderSettingsCreate20250101) GetDiskSizeGB() float64 {
	if o == nil || IsNil(o.DiskSizeGB) {
		var ret float64
		return ret
	}
	return *o.DiskSizeGB
}

// GetDiskSizeGBOk returns a tuple with the DiskSizeGB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexProviderSettingsCreate20250101) GetDiskSizeGBOk() (*float64, bool) {
	if o == nil || IsNil(o.DiskSizeGB) {
		return nil, false
	}

	return o.DiskSizeGB, true
}

// HasDiskSizeGB returns a boolean if a field has been set.
func (o *FlexProviderSettingsCreate20250101) HasDiskSizeGB() bool {
	if o != nil && !IsNil(o.DiskSizeGB) {
		return true
	}

	return false
}

// SetDiskSizeGB gets a reference to the given float64 and assigns it to the DiskSizeGB field.
func (o *FlexProviderSettingsCreate20250101) SetDiskSizeGB(v float64) {
	o.DiskSizeGB = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise
func (o *FlexProviderSettingsCreate20250101) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlexProviderSettingsCreate20250101) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}

	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *FlexProviderSettingsCreate20250101) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *FlexProviderSettingsCreate20250101) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetRegionName returns the RegionName field value
func (o *FlexProviderSettingsCreate20250101) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *FlexProviderSettingsCreate20250101) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *FlexProviderSettingsCreate20250101) SetRegionName(v string) {
	o.RegionName = v
}
