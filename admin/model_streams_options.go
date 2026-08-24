// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsOptions Optional configuration for the stream processor.
type StreamsOptions struct {
	Autoscaling *StreamsAutoscaling `json:"autoscaling,omitempty"`
	Dlq         *StreamsDLQ         `json:"dlq,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsOptions) MarshalJSON() ([]byte, error) {
	type noMethod StreamsOptions
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsOptions instantiates a new StreamsOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsOptions() *StreamsOptions {
	this := StreamsOptions{}
	return &this
}

// NewStreamsOptionsWithDefaults instantiates a new StreamsOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsOptionsWithDefaults() *StreamsOptions {
	this := StreamsOptions{}
	return &this
}

// GetAutoscaling returns the Autoscaling field value if set, zero value otherwise
func (o *StreamsOptions) GetAutoscaling() StreamsAutoscaling {
	if o == nil || IsNil(o.Autoscaling) {
		var ret StreamsAutoscaling
		return ret
	}
	return *o.Autoscaling
}

// GetAutoscalingOk returns a tuple with the Autoscaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsOptions) GetAutoscalingOk() (*StreamsAutoscaling, bool) {
	if o == nil || IsNil(o.Autoscaling) {
		return nil, false
	}

	return o.Autoscaling, true
}

// HasAutoscaling returns a boolean if a field has been set.
func (o *StreamsOptions) HasAutoscaling() bool {
	if o != nil && !IsNil(o.Autoscaling) {
		return true
	}

	return false
}

// SetAutoscaling gets a reference to the given StreamsAutoscaling and assigns it to the Autoscaling field.
func (o *StreamsOptions) SetAutoscaling(v StreamsAutoscaling) {
	o.Autoscaling = &v
	o.NullFields = removeNullField(o.NullFields, "Autoscaling")
}

// SetAutoscalingNil sets Autoscaling to an explicit JSON null when marshaled.
func (o *StreamsOptions) SetAutoscalingNil() {
	o.Autoscaling = nil
	o.NullFields = addNullField(o.NullFields, "Autoscaling")
}

// GetDlq returns the Dlq field value if set, zero value otherwise
func (o *StreamsOptions) GetDlq() StreamsDLQ {
	if o == nil || IsNil(o.Dlq) {
		var ret StreamsDLQ
		return ret
	}
	return *o.Dlq
}

// GetDlqOk returns a tuple with the Dlq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsOptions) GetDlqOk() (*StreamsDLQ, bool) {
	if o == nil || IsNil(o.Dlq) {
		return nil, false
	}

	return o.Dlq, true
}

// HasDlq returns a boolean if a field has been set.
func (o *StreamsOptions) HasDlq() bool {
	if o != nil && !IsNil(o.Dlq) {
		return true
	}

	return false
}

// SetDlq gets a reference to the given StreamsDLQ and assigns it to the Dlq field.
func (o *StreamsOptions) SetDlq(v StreamsDLQ) {
	o.Dlq = &v
	o.NullFields = removeNullField(o.NullFields, "Dlq")
}

// SetDlqNil sets Dlq to an explicit JSON null when marshaled.
func (o *StreamsOptions) SetDlqNil() {
	o.Dlq = nil
	o.NullFields = addNullField(o.NullFields, "Dlq")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsOptions) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsOptions) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsOptions) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsOptions) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsOptions) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}
