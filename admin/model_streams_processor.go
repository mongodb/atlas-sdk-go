// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsProcessor An atlas stream processor.
type StreamsProcessor struct {
	// Unique 24-hexadecimal character string that identifies the stream processor.
	// Read only field.
	Id *string `json:"_id,omitempty"`
	// Flag that enables or disables failover for the stream processor.
	FailoverEnabled *bool `json:"failoverEnabled,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Human-readable name of the stream processor.
	Name    *string         `json:"name,omitempty"`
	Options *StreamsOptions `json:"options,omitempty"`
	// Stream aggregation pipeline you want to apply to your streaming data.
	Pipeline *[]any `json:"pipeline,omitempty"`
	// Selected tier for the Stream Workspace. Configures Memory / VCPU allowances.
	Tier *string `json:"tier,omitempty"`
	// NullFields is a list of field names (e.g. "FieldName") to send as an explicit JSON null,
	// overriding the field's actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsProcessor) MarshalJSON() ([]byte, error) {
	type noMethod StreamsProcessor
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsProcessor instantiates a new StreamsProcessor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsProcessor() *StreamsProcessor {
	this := StreamsProcessor{}
	return &this
}

// NewStreamsProcessorWithDefaults instantiates a new StreamsProcessor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsProcessorWithDefaults() *StreamsProcessor {
	this := StreamsProcessor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise
func (o *StreamsProcessor) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}

	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StreamsProcessor) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StreamsProcessor) SetId(v string) {
	o.Id = &v
}

// SetIdNil sets Id to an explicit JSON null when marshaled.
func (o *StreamsProcessor) SetIdNil() {
	o.Id = nil
	o.NullFields = append(o.NullFields, "Id")
}

// GetFailoverEnabled returns the FailoverEnabled field value if set, zero value otherwise
func (o *StreamsProcessor) GetFailoverEnabled() bool {
	if o == nil || IsNil(o.FailoverEnabled) {
		var ret bool
		return ret
	}
	return *o.FailoverEnabled
}

// GetFailoverEnabledOk returns a tuple with the FailoverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetFailoverEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FailoverEnabled) {
		return nil, false
	}

	return o.FailoverEnabled, true
}

// HasFailoverEnabled returns a boolean if a field has been set.
func (o *StreamsProcessor) HasFailoverEnabled() bool {
	if o != nil && !IsNil(o.FailoverEnabled) {
		return true
	}

	return false
}

// SetFailoverEnabled gets a reference to the given bool and assigns it to the FailoverEnabled field.
func (o *StreamsProcessor) SetFailoverEnabled(v bool) {
	o.FailoverEnabled = &v
}

// SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled.
func (o *StreamsProcessor) SetFailoverEnabledNil() {
	o.FailoverEnabled = nil
	o.NullFields = append(o.NullFields, "FailoverEnabled")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsProcessor) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsProcessor) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsProcessor) SetLinks(v []Link) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise
func (o *StreamsProcessor) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}

	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StreamsProcessor) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StreamsProcessor) SetName(v string) {
	o.Name = &v
}

// SetNameNil sets Name to an explicit JSON null when marshaled.
func (o *StreamsProcessor) SetNameNil() {
	o.Name = nil
	o.NullFields = append(o.NullFields, "Name")
}

// GetOptions returns the Options field value if set, zero value otherwise
func (o *StreamsProcessor) GetOptions() StreamsOptions {
	if o == nil || IsNil(o.Options) {
		var ret StreamsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetOptionsOk() (*StreamsOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}

	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *StreamsProcessor) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given StreamsOptions and assigns it to the Options field.
func (o *StreamsProcessor) SetOptions(v StreamsOptions) {
	o.Options = &v
}

// SetOptionsNil sets Options to an explicit JSON null when marshaled.
func (o *StreamsProcessor) SetOptionsNil() {
	o.Options = nil
	o.NullFields = append(o.NullFields, "Options")
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise
func (o *StreamsProcessor) GetPipeline() []any {
	if o == nil || IsNil(o.Pipeline) {
		var ret []any
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetPipelineOk() (*[]any, bool) {
	if o == nil || IsNil(o.Pipeline) {
		return nil, false
	}

	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *StreamsProcessor) HasPipeline() bool {
	if o != nil && !IsNil(o.Pipeline) {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given []any and assigns it to the Pipeline field.
func (o *StreamsProcessor) SetPipeline(v []any) {
	o.Pipeline = &v
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *StreamsProcessor) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessor) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *StreamsProcessor) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *StreamsProcessor) SetTier(v string) {
	o.Tier = &v
}

// SetTierNil sets Tier to an explicit JSON null when marshaled.
func (o *StreamsProcessor) SetTierNil() {
	o.Tier = nil
	o.NullFields = append(o.NullFields, "Tier")
}
