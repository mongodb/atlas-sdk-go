// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsProcessorStatus Desired status change to apply to a tenant's stream processors.
type StreamsProcessorStatus struct {
	// Strategy for the processor: GRACEFUL - attempt to stop the processor, error if processor cannot be stopped. if stop was successful, start the processor in the new region with the latest checkpoint.  FORCED - attempt to stop the processor, proceed to starting the processor in the new region with checkpoints disabled regardless of whether or not the stop succeeds.
	Mode *string `json:"mode,omitempty"`
	// Name of the region against which to apply the status change.
	Region string `json:"region"`
	// Represents the desired action to apply to stream processors within a workspace, such as starting all processors.
	Status string `json:"status"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewStreamsProcessorStatus instantiates a new StreamsProcessorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsProcessorStatus(region string, status string) *StreamsProcessorStatus {
	this := StreamsProcessorStatus{}
	this.Region = region
	this.Status = status
	return &this
}

// NewStreamsProcessorStatusWithDefaults instantiates a new StreamsProcessorStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsProcessorStatusWithDefaults() *StreamsProcessorStatus {
	this := StreamsProcessorStatus{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise
func (o *StreamsProcessorStatus) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorStatus) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}

	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *StreamsProcessorStatus) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *StreamsProcessorStatus) SetMode(v string) {
	o.Mode = &v
}

// GetRegion returns the Region field value
func (o *StreamsProcessorStatus) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorStatus) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *StreamsProcessorStatus) SetRegion(v string) {
	o.Region = v
}

// GetStatus returns the Status field value
func (o *StreamsProcessorStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *StreamsProcessorStatus) SetStatus(v string) {
	o.Status = v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsProcessorStatus) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorStatus) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsProcessorStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsProcessorStatus) SetLinks(v []Link) {
	o.Links = &v
}
