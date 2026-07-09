// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsModifyStreamProcessor A request to modify an existing stream processor.
type StreamsModifyStreamProcessor struct {
	// Flag that enables or disables failover for the stream processor.
	FailoverEnabled *bool `json:"failoverEnabled,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// New name for the stream processor.
	Name    *string                              `json:"name,omitempty"`
	Options *StreamsModifyStreamProcessorOptions `json:"options,omitempty"`
	// New pipeline for the stream processor.
	Pipeline *[]any `json:"pipeline,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsModifyStreamProcessor) MarshalJSON() ([]byte, error) {
	type noMethod StreamsModifyStreamProcessor
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsModifyStreamProcessor instantiates a new StreamsModifyStreamProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsModifyStreamProcessor() *StreamsModifyStreamProcessor {
	this := StreamsModifyStreamProcessor{}
	return &this
}

// NewStreamsModifyStreamProcessorWithDefaults instantiates a new StreamsModifyStreamProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsModifyStreamProcessorWithDefaults() *StreamsModifyStreamProcessor {
	this := StreamsModifyStreamProcessor{}
	return &this
}

// GetFailoverEnabled returns the FailoverEnabled field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessor) GetFailoverEnabled() bool {
	if o == nil || IsNil(o.FailoverEnabled) {
		var ret bool
		return ret
	}
	return *o.FailoverEnabled
}

// GetFailoverEnabledOk returns a tuple with the FailoverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessor) GetFailoverEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FailoverEnabled) {
		return nil, false
	}

	return o.FailoverEnabled, true
}

// HasFailoverEnabled returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessor) HasFailoverEnabled() bool {
	if o != nil && !IsNil(o.FailoverEnabled) {
		return true
	}

	return false
}

// SetFailoverEnabled gets a reference to the given bool and assigns it to the FailoverEnabled field.
func (o *StreamsModifyStreamProcessor) SetFailoverEnabled(v bool) {
	o.FailoverEnabled = &v
}

// SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled.
func (o *StreamsModifyStreamProcessor) SetFailoverEnabledNil() {
	o.FailoverEnabled = nil
	o.NullFields = append(o.NullFields, "FailoverEnabled")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessor) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessor) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessor) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsModifyStreamProcessor) SetLinks(v []Link) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StreamsModifyStreamProcessor) SetName(v string) {
	o.Name = &v
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *StreamsModifyStreamProcessor) SetNameNil() {
	o.Name = nil
	o.NullFields = append(o.NullFields, "Name")
}

// GetOptions returns the Options field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessor) GetOptions() StreamsModifyStreamProcessorOptions {
	if o == nil || IsNil(o.Options) {
		var ret StreamsModifyStreamProcessorOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessor) GetOptionsOk() (*StreamsModifyStreamProcessorOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}

	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessor) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given StreamsModifyStreamProcessorOptions and assigns it to the Options field.
func (o *StreamsModifyStreamProcessor) SetOptions(v StreamsModifyStreamProcessorOptions) {
	o.Options = &v
}

// SetOptionsNil sets Options to an explicit JSON null when marshaled.
func (o *StreamsModifyStreamProcessor) SetOptionsNil() {
	o.Options = nil
	o.NullFields = append(o.NullFields, "Options")
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise
func (o *StreamsModifyStreamProcessor) GetPipeline() []any {
	if o == nil || IsNil(o.Pipeline) {
		var ret []any
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsModifyStreamProcessor) GetPipelineOk() (*[]any, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}

	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *StreamsModifyStreamProcessor) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given []any and assigns it to the Pipeline field.
func (o *StreamsModifyStreamProcessor) SetPipeline(v []any) {
	o.Pipeline = &v
}
