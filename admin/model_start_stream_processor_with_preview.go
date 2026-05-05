// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// StartStreamProcessorWithPreview A request to start a stream processor.
type StartStreamProcessorWithPreview struct {
	Failover *StreamsStartProcessorFailover `json:"failover,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// When true or not specified, the stream processor resumes from its last checkpoint. When false, the stream processor starts fresh.
	ResumeFromCheckpoint *bool `json:"resumeFromCheckpoint,omitempty"`
	// The operation time after which the change stream source should begin reporting. This parameter expresses its value in the ISO 8601 timestamp format in UTC.
	StartAtOperationTime *time.Time `json:"startAtOperationTime,omitempty"`
	// Selected tier for the Stream Workspace. Configures Memory / VCPU allowances.
	Tier *string `json:"tier,omitempty"`
}

// NewStartStreamProcessorWithPreview instantiates a new StartStreamProcessorWithPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartStreamProcessorWithPreview() *StartStreamProcessorWithPreview {
	this := StartStreamProcessorWithPreview{}
	return &this
}

// NewStartStreamProcessorWithPreviewWithDefaults instantiates a new StartStreamProcessorWithPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartStreamProcessorWithPreviewWithDefaults() *StartStreamProcessorWithPreview {
	this := StartStreamProcessorWithPreview{}
	return &this
}

// GetFailover returns the Failover field value if set, zero value otherwise
func (o *StartStreamProcessorWithPreview) GetFailover() StreamsStartProcessorFailover {
	if o == nil || IsNil(o.Failover) {
		var ret StreamsStartProcessorFailover
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartStreamProcessorWithPreview) GetFailoverOk() (*StreamsStartProcessorFailover, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}

	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StartStreamProcessorWithPreview) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given StreamsStartProcessorFailover and assigns it to the Failover field.
func (o *StartStreamProcessorWithPreview) SetFailover(v StreamsStartProcessorFailover) {
	o.Failover = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StartStreamProcessorWithPreview) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartStreamProcessorWithPreview) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StartStreamProcessorWithPreview) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StartStreamProcessorWithPreview) SetLinks(v []Link) {
	o.Links = &v
}

// GetResumeFromCheckpoint returns the ResumeFromCheckpoint field value if set, zero value otherwise
func (o *StartStreamProcessorWithPreview) GetResumeFromCheckpoint() bool {
	if o == nil || IsNil(o.ResumeFromCheckpoint) {
		var ret bool
		return ret
	}
	return *o.ResumeFromCheckpoint
}

// GetResumeFromCheckpointOk returns a tuple with the ResumeFromCheckpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartStreamProcessorWithPreview) GetResumeFromCheckpointOk() (*bool, bool) {
	if o == nil || IsNil(o.ResumeFromCheckpoint) {
		return nil, false
	}

	return o.ResumeFromCheckpoint, true
}

// HasResumeFromCheckpoint returns a boolean if a field has been set.
func (o *StartStreamProcessorWithPreview) HasResumeFromCheckpoint() bool {
	if o != nil && !IsNil(o.ResumeFromCheckpoint) {
		return true
	}

	return false
}

// SetResumeFromCheckpoint gets a reference to the given bool and assigns it to the ResumeFromCheckpoint field.
func (o *StartStreamProcessorWithPreview) SetResumeFromCheckpoint(v bool) {
	o.ResumeFromCheckpoint = &v
}

// GetStartAtOperationTime returns the StartAtOperationTime field value if set, zero value otherwise
func (o *StartStreamProcessorWithPreview) GetStartAtOperationTime() time.Time {
	if o == nil || IsNil(o.StartAtOperationTime) {
		var ret time.Time
		return ret
	}
	return *o.StartAtOperationTime
}

// GetStartAtOperationTimeOk returns a tuple with the StartAtOperationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartStreamProcessorWithPreview) GetStartAtOperationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartAtOperationTime) {
		return nil, false
	}

	return o.StartAtOperationTime, true
}

// HasStartAtOperationTime returns a boolean if a field has been set.
func (o *StartStreamProcessorWithPreview) HasStartAtOperationTime() bool {
	if o != nil && !IsNil(o.StartAtOperationTime) {
		return true
	}

	return false
}

// SetStartAtOperationTime gets a reference to the given time.Time and assigns it to the StartAtOperationTime field.
func (o *StartStreamProcessorWithPreview) SetStartAtOperationTime(v time.Time) {
	o.StartAtOperationTime = &v
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *StartStreamProcessorWithPreview) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartStreamProcessorWithPreview) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *StartStreamProcessorWithPreview) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *StartStreamProcessorWithPreview) SetTier(v string) {
	o.Tier = &v
}
