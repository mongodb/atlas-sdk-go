// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"time"
)

// StreamsStartStreamProcessorWith A request to start a stream processor.
type StreamsStartStreamProcessorWith struct {
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
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsStartStreamProcessorWith) MarshalJSON() ([]byte, error) {
	type noMethod StreamsStartStreamProcessorWith
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsStartStreamProcessorWith instantiates a new StreamsStartStreamProcessorWith object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsStartStreamProcessorWith() *StreamsStartStreamProcessorWith {
	this := StreamsStartStreamProcessorWith{}
	return &this
}

// NewStreamsStartStreamProcessorWithWithDefaults instantiates a new StreamsStartStreamProcessorWith object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsStartStreamProcessorWithWithDefaults() *StreamsStartStreamProcessorWith {
	this := StreamsStartStreamProcessorWith{}
	return &this
}

// GetFailover returns the Failover field value if set, zero value otherwise
func (o *StreamsStartStreamProcessorWith) GetFailover() StreamsStartProcessorFailover {
	if o == nil || IsNil(o.Failover) {
		var ret StreamsStartProcessorFailover
		return ret
	}
	return *o.Failover
}

// GetFailoverOk returns a tuple with the Failover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartStreamProcessorWith) GetFailoverOk() (*StreamsStartProcessorFailover, bool) {
	if o == nil || IsNil(o.Failover) {
		return nil, false
	}

	return o.Failover, true
}

// HasFailover returns a boolean if a field has been set.
func (o *StreamsStartStreamProcessorWith) HasFailover() bool {
	if o != nil && !IsNil(o.Failover) {
		return true
	}

	return false
}

// SetFailover gets a reference to the given StreamsStartProcessorFailover and assigns it to the Failover field.
func (o *StreamsStartStreamProcessorWith) SetFailover(v StreamsStartProcessorFailover) {
	o.Failover = &v
	o.NullFields = removeNullField(o.NullFields, "Failover")
}

// SetFailoverNil sets Failover to an explicit JSON null when marshaled.
func (o *StreamsStartStreamProcessorWith) SetFailoverNil() {
	o.Failover = nil
	o.NullFields = addNullField(o.NullFields, "Failover")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsStartStreamProcessorWith) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartStreamProcessorWith) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsStartStreamProcessorWith) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsStartStreamProcessorWith) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsStartStreamProcessorWith) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetResumeFromCheckpoint returns the ResumeFromCheckpoint field value if set, zero value otherwise
func (o *StreamsStartStreamProcessorWith) GetResumeFromCheckpoint() bool {
	if o == nil || IsNil(o.ResumeFromCheckpoint) {
		var ret bool
		return ret
	}
	return *o.ResumeFromCheckpoint
}

// GetResumeFromCheckpointOk returns a tuple with the ResumeFromCheckpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartStreamProcessorWith) GetResumeFromCheckpointOk() (*bool, bool) {
	if o == nil || IsNil(o.ResumeFromCheckpoint) {
		return nil, false
	}

	return o.ResumeFromCheckpoint, true
}

// HasResumeFromCheckpoint returns a boolean if a field has been set.
func (o *StreamsStartStreamProcessorWith) HasResumeFromCheckpoint() bool {
	if o != nil && !IsNil(o.ResumeFromCheckpoint) {
		return true
	}

	return false
}

// SetResumeFromCheckpoint gets a reference to the given bool and assigns it to the ResumeFromCheckpoint field.
func (o *StreamsStartStreamProcessorWith) SetResumeFromCheckpoint(v bool) {
	o.ResumeFromCheckpoint = &v
	o.NullFields = removeNullField(o.NullFields, "ResumeFromCheckpoint")
}

// SetResumeFromCheckpointNil sets ResumeFromCheckpoint to an explicit JSON null when marshaled.
func (o *StreamsStartStreamProcessorWith) SetResumeFromCheckpointNil() {
	o.ResumeFromCheckpoint = nil
	o.NullFields = addNullField(o.NullFields, "ResumeFromCheckpoint")
}

// GetStartAtOperationTime returns the StartAtOperationTime field value if set, zero value otherwise
func (o *StreamsStartStreamProcessorWith) GetStartAtOperationTime() time.Time {
	if o == nil || IsNil(o.StartAtOperationTime) {
		var ret time.Time
		return ret
	}
	return *o.StartAtOperationTime
}

// GetStartAtOperationTimeOk returns a tuple with the StartAtOperationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartStreamProcessorWith) GetStartAtOperationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartAtOperationTime) {
		return nil, false
	}

	return o.StartAtOperationTime, true
}

// HasStartAtOperationTime returns a boolean if a field has been set.
func (o *StreamsStartStreamProcessorWith) HasStartAtOperationTime() bool {
	if o != nil && !IsNil(o.StartAtOperationTime) {
		return true
	}

	return false
}

// SetStartAtOperationTime gets a reference to the given time.Time and assigns it to the StartAtOperationTime field.
func (o *StreamsStartStreamProcessorWith) SetStartAtOperationTime(v time.Time) {
	o.StartAtOperationTime = &v
	o.NullFields = removeNullField(o.NullFields, "StartAtOperationTime")
}

// SetStartAtOperationTimeNil sets StartAtOperationTime to an explicit JSON null when marshaled.
func (o *StreamsStartStreamProcessorWith) SetStartAtOperationTimeNil() {
	o.StartAtOperationTime = nil
	o.NullFields = addNullField(o.NullFields, "StartAtOperationTime")
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *StreamsStartStreamProcessorWith) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStartStreamProcessorWith) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *StreamsStartStreamProcessorWith) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *StreamsStartStreamProcessorWith) SetTier(v string) {
	o.Tier = &v
	o.NullFields = removeNullField(o.NullFields, "Tier")
}

// SetTierNil sets Tier to an explicit JSON null when marshaled.
func (o *StreamsStartStreamProcessorWith) SetTierNil() {
	o.Tier = nil
	o.NullFields = addNullField(o.NullFields, "Tier")
}
