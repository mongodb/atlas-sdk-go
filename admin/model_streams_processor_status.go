// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsProcessorStatus Desired status change to apply to a tenant's stream processors.
type StreamsProcessorStatus struct {
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Strategy for the processor: GRACEFUL - attempt to stop the processor, error if processor cannot be stopped. if stop was successful, start the processor in the new region with the latest checkpoint.  FORCED - attempt to stop the processor, proceed to starting the processor in the new region with checkpoints disabled regardless of whether or not the stop succeeds.
	Mode *string `json:"mode,omitempty"`
	// Name of the region against which to apply the status change. Required when `status` is `FAILED_OVER`; optional otherwise.
	Region *string `json:"region,omitempty"`
	// Represents the desired action to apply to stream processors within a workspace, such as starting all processors, stopping all processors, or performing a bulk regional failover.
	Status string `json:"status"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsProcessorStatus) MarshalJSON() ([]byte, error) {
	type noMethod StreamsProcessorStatus
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsProcessorStatus instantiates a new StreamsProcessorStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsProcessorStatus(status string) *StreamsProcessorStatus {
	this := StreamsProcessorStatus{}
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
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsProcessorStatus) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
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
	o.NullFields = removeNullField(o.NullFields, "Mode")
}

// SetModeNil sets Mode to an explicit JSON null when marshaled.
func (o *StreamsProcessorStatus) SetModeNil() {
	o.Mode = nil
	o.NullFields = addNullField(o.NullFields, "Mode")
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsProcessorStatus) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorStatus) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsProcessorStatus) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsProcessorStatus) SetRegion(v string) {
	o.Region = &v
	o.NullFields = removeNullField(o.NullFields, "Region")
}

// SetRegionNil sets Region to an explicit JSON null when marshaled.
func (o *StreamsProcessorStatus) SetRegionNil() {
	o.Region = nil
	o.NullFields = addNullField(o.NullFields, "Region")
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
