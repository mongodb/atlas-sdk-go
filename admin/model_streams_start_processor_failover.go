// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsStartProcessorFailover Failover options for starting a stream processor.
type StreamsStartProcessorFailover struct {
	// If true, simulates the operation without making any changes.
	DryRun *bool `json:"dryRun,omitempty"`
	// Strategy for the processor: GRACEFUL - attempt to stop the processor, error if processor cannot be stopped. if stop was successful, start the processor in the new region with the latest checkpoint.  FORCED - attempt to stop the processor, proceed to starting the processor in the new region with checkpoints disabled regardless of whether or not the stop succeeds.
	Mode *string `json:"mode,omitempty"`
	// Cloud provider region where the stream processor should be started in failover mode. The region must be a valid region for the stream processor's cloud provider and must be included in the tenant's configured failover regions, or it may be the tenant's default (primary) region.
	Region string `json:"region"`
}

// NewStreamsStartProcessorFailover instantiates a new StreamsStartProcessorFailover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsStartProcessorFailover(region string) *StreamsStartProcessorFailover {
	this := StreamsStartProcessorFailover{}
	this.Region = region
	return &this
}

// NewStreamsStartProcessorFailoverWithDefaults instantiates a new StreamsStartProcessorFailover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsStartProcessorFailoverWithDefaults() *StreamsStartProcessorFailover {
	this := StreamsStartProcessorFailover{}
	return &this
}

// GetDryRun returns the DryRun field value if set, zero value otherwise
func (o *StreamsStartProcessorFailover) GetDryRun() bool {
	if o == nil || IsNil(o.DryRun) {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartProcessorFailover) GetDryRunOk() (*bool, bool) {
	if o == nil || IsNil(o.DryRun) {
		return nil, false
	}

	return o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *StreamsStartProcessorFailover) HasDryRun() bool {
	if o != nil && !IsNil(o.DryRun) {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *StreamsStartProcessorFailover) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetMode returns the Mode field value if set, zero value otherwise
func (o *StreamsStartProcessorFailover) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartProcessorFailover) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}

	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *StreamsStartProcessorFailover) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *StreamsStartProcessorFailover) SetMode(v string) {
	o.Mode = &v
}

// GetRegion returns the Region field value
func (o *StreamsStartProcessorFailover) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *StreamsStartProcessorFailover) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *StreamsStartProcessorFailover) SetRegion(v string) {
	o.Region = v
}
