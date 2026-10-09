// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsProcessorRequest Details to create a stream processor.
type StreamsProcessorRequest struct {
	// Flag that enables or disables failover for the stream processor.
	FailoverEnabled *bool `json:"failoverEnabled,omitempty"`
	// Human-readable name of the stream processor.
	Name    string          `json:"name"`
	Options *StreamsOptions `json:"options,omitempty"`
	// Stream aggregation pipeline you want to apply to your streaming data.
	Pipeline []any `json:"pipeline"`
	// Baseline processing tier requested by the user.
	Tier *string `json:"tier,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsProcessorRequest) MarshalJSON() ([]byte, error) {
	type noMethod StreamsProcessorRequest
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsProcessorRequest instantiates a new StreamsProcessorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsProcessorRequest(name string, pipeline []any) *StreamsProcessorRequest {
	this := StreamsProcessorRequest{}
	this.Name = name
	this.Pipeline = pipeline
	return &this
}

// NewStreamsProcessorRequestWithDefaults instantiates a new StreamsProcessorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsProcessorRequestWithDefaults() *StreamsProcessorRequest {
	this := StreamsProcessorRequest{}
	return &this
}

// GetFailoverEnabled returns the FailoverEnabled field value if set, zero value otherwise
func (o *StreamsProcessorRequest) GetFailoverEnabled() bool {
	if o == nil || IsNil(o.FailoverEnabled) {
		var ret bool
		return ret
	}
	return *o.FailoverEnabled
}

// GetFailoverEnabledOk returns a tuple with the FailoverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorRequest) GetFailoverEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FailoverEnabled) {
		return nil, false
	}

	return o.FailoverEnabled, true
}

// HasFailoverEnabled returns a boolean if a field has been set.
func (o *StreamsProcessorRequest) HasFailoverEnabled() bool {
	if o != nil && !IsNil(o.FailoverEnabled) {
		return true
	}

	return false
}

// SetFailoverEnabled gets a reference to the given bool and assigns it to the FailoverEnabled field.
func (o *StreamsProcessorRequest) SetFailoverEnabled(v bool) {
	o.FailoverEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "FailoverEnabled")
}

// SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled.
func (o *StreamsProcessorRequest) SetFailoverEnabledNil() {
	o.FailoverEnabled = nil
	o.NullFields = addNullField(o.NullFields, "FailoverEnabled")
}

// GetName returns the Name field value
func (o *StreamsProcessorRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StreamsProcessorRequest) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise
func (o *StreamsProcessorRequest) GetOptions() StreamsOptions {
	if o == nil || IsNil(o.Options) {
		var ret StreamsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorRequest) GetOptionsOk() (*StreamsOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}

	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *StreamsProcessorRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given StreamsOptions and assigns it to the Options field.
func (o *StreamsProcessorRequest) SetOptions(v StreamsOptions) {
	o.Options = &v
	o.NullFields = removeNullField(o.NullFields, "Options")
}

// SetOptionsNil sets Options to an explicit JSON null when marshaled.
func (o *StreamsProcessorRequest) SetOptionsNil() {
	o.Options = nil
	o.NullFields = addNullField(o.NullFields, "Options")
}

// GetPipeline returns the Pipeline field value
func (o *StreamsProcessorRequest) GetPipeline() []any {
	if o == nil {
		var ret []any
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorRequest) GetPipelineOk() (*[]any, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *StreamsProcessorRequest) SetPipeline(v []any) {
	o.Pipeline = v
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *StreamsProcessorRequest) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorRequest) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *StreamsProcessorRequest) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *StreamsProcessorRequest) SetTier(v string) {
	o.Tier = &v
	o.NullFields = removeNullField(o.NullFields, "Tier")
}

// SetTierNil sets Tier to an explicit JSON null when marshaled.
func (o *StreamsProcessorRequest) SetTierNil() {
	o.Tier = nil
	o.NullFields = addNullField(o.NullFields, "Tier")
}
