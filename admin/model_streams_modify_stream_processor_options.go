// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsModifyStreamProcessorOptions Additional options for modifying a stream processor.
type StreamsModifyStreamProcessorOptions struct {
	Dlq *StreamsDLQ `json:"dlq,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// When true, the modified stream processor resumes from its last checkpoint.
	ResumeFromCheckpoint *bool `json:"resumeFromCheckpoint,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsModifyStreamProcessorOptions) MarshalJSON() ([]byte, error) {
	type noMethod StreamsModifyStreamProcessorOptions
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsModifyStreamProcessorOptions instantiates a new StreamsModifyStreamProcessorOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsModifyStreamProcessorOptions() *StreamsModifyStreamProcessorOptions {
	this := StreamsModifyStreamProcessorOptions{}
	return &this
}

// NewStreamsModifyStreamProcessorOptionsWithDefaults instantiates a new StreamsModifyStreamProcessorOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsModifyStreamProcessorOptionsWithDefaults() *StreamsModifyStreamProcessorOptions {
	this := StreamsModifyStreamProcessorOptions{}
	return &this
}

// GetDlq returns the Dlq field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessorOptions) GetDlq() StreamsDLQ {
	if o == nil || IsNil(o.Dlq) {
		var ret StreamsDLQ
		return ret
	}
	return *o.Dlq
}

// GetDlqOk returns a tuple with the Dlq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessorOptions) GetDlqOk() (*StreamsDLQ, bool) {
	if o == nil || IsNil(o.Dlq) {
		return nil, false
	}

	return o.Dlq, true
}

// HasDlq returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessorOptions) HasDlq() bool {
	if o != nil && !IsNil(o.Dlq) {
		return true
	}

	return false
}

// SetDlq gets a reference to the given StreamsDLQ and assigns it to the Dlq field.
func (o *StreamsModifyStreamProcessorOptions) SetDlq(v StreamsDLQ) {
	o.Dlq = &v
	o.NullFields = removeNullField(o.NullFields, "Dlq")
}

// SetDlqNil sets Dlq to an explicit JSON null when marshaled.
func (o *StreamsModifyStreamProcessorOptions) SetDlqNil() {
	o.Dlq = nil
	o.NullFields = addNullField(o.NullFields, "Dlq")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessorOptions) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessorOptions) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessorOptions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsModifyStreamProcessorOptions) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsModifyStreamProcessorOptions) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetResumeFromCheckpoint returns the ResumeFromCheckpoint field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessorOptions) GetResumeFromCheckpoint() bool {
	if o == nil || IsNil(o.ResumeFromCheckpoint) {
		var ret bool
		return ret
	}
	return *o.ResumeFromCheckpoint
}

// GetResumeFromCheckpointOk returns a tuple with the ResumeFromCheckpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessorOptions) GetResumeFromCheckpointOk() (*bool, bool) {
	if o == nil || IsNil(o.ResumeFromCheckpoint) {
		return nil, false
	}

	return o.ResumeFromCheckpoint, true
}

// HasResumeFromCheckpoint returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessorOptions) HasResumeFromCheckpoint() bool {
	if o != nil && !IsNil(o.ResumeFromCheckpoint) {
		return true
	}

	return false
}

// SetResumeFromCheckpoint gets a reference to the given bool and assigns it to the ResumeFromCheckpoint field.
func (o *StreamsModifyStreamProcessorOptions) SetResumeFromCheckpoint(v bool) {
	o.ResumeFromCheckpoint = &v
	o.NullFields = removeNullField(o.NullFields, "ResumeFromCheckpoint")
}

// SetResumeFromCheckpointNil sets ResumeFromCheckpoint to an explicit JSON null when marshaled.
func (o *StreamsModifyStreamProcessorOptions) SetResumeFromCheckpointNil() {
	o.ResumeFromCheckpoint = nil
	o.NullFields = addNullField(o.NullFields, "ResumeFromCheckpoint")
}
