// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsProcessorWithStats An atlas stream processor with optional stats.
type StreamsProcessorWithStats struct {
	// Unique 24-hexadecimal character string that identifies the stream processor.
	// Read only field.
	Id string `json:"_id"`
	// Flag that indicates whether the stream processor is eligible for failover.
	// Read only field.
	EligibleForFailover *bool `json:"eligibleForFailover,omitempty"`
	// Flag that enables or disables failover for the stream processor.
	// Read only field.
	FailoverEnabled *bool `json:"failoverEnabled,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
	// Human-readable name of the stream processor.
	// Read only field.
	Name    string          `json:"name"`
	Options *StreamsOptions `json:"options,omitempty"`
	// Stream aggregation pipeline you want to apply to your streaming data.
	// Read only field.
	Pipeline []any `json:"pipeline"`
	// The state of the stream processor. Commonly occurring states are 'CREATED', 'STARTED', 'STOPPED' and 'FAILED'.
	// Read only field.
	State string `json:"state"`
	// The stats associated with the stream processor.
	// Read only field.
	Stats any `json:"stats,omitempty"`
	// Selected tier for the Stream Workspace. Configures Memory / VCPU allowances.
	Tier *string `json:"tier,omitempty"`
	// NullFields is an internal field that is never sent as part of the payload (see the `json:"-"` tag below).
	// It holds a list of field names (e.g. "FieldName") to send as an explicit JSON null instead of their actual value.
	NullFields []string `json:"-"`
}

// MarshalJSON honors NullFields, in addition to the regular struct tags.
func (o *StreamsProcessorWithStats) MarshalJSON() ([]byte, error) {
	type noMethod StreamsProcessorWithStats
	return marshalWithNullFields(noMethod(*o), o.NullFields)
}

// NewStreamsProcessorWithStats instantiates a new StreamsProcessorWithStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsProcessorWithStats(id string, name string, pipeline []any, state string) *StreamsProcessorWithStats {
	this := StreamsProcessorWithStats{}
	this.Id = id
	this.Name = name
	this.Pipeline = pipeline
	this.State = state
	return &this
}

// NewStreamsProcessorWithStatsWithDefaults instantiates a new StreamsProcessorWithStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsProcessorWithStatsWithDefaults() *StreamsProcessorWithStats {
	this := StreamsProcessorWithStats{}
	return &this
}

// GetId returns the Id field value
func (o *StreamsProcessorWithStats) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StreamsProcessorWithStats) SetId(v string) {
	o.Id = v
}

// GetEligibleForFailover returns the EligibleForFailover field value if set, zero value otherwise
func (o *StreamsProcessorWithStats) GetEligibleForFailover() bool {
	if o == nil || IsNil(o.EligibleForFailover) {
		var ret bool
		return ret
	}
	return *o.EligibleForFailover
}

// GetEligibleForFailoverOk returns a tuple with the EligibleForFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetEligibleForFailoverOk() (*bool, bool) {
	if o == nil || IsNil(o.EligibleForFailover) {
		return nil, false
	}

	return o.EligibleForFailover, true
}

// HasEligibleForFailover returns a boolean if a field has been set.
func (o *StreamsProcessorWithStats) HasEligibleForFailover() bool {
	if o != nil && !IsNil(o.EligibleForFailover) {
		return true
	}

	return false
}

// SetEligibleForFailover gets a reference to the given bool and assigns it to the EligibleForFailover field.
func (o *StreamsProcessorWithStats) SetEligibleForFailover(v bool) {
	o.EligibleForFailover = &v
	o.NullFields = removeNullField(o.NullFields, "EligibleForFailover")
}

// SetEligibleForFailoverNil sets EligibleForFailover to an explicit JSON null when marshaled.
func (o *StreamsProcessorWithStats) SetEligibleForFailoverNil() {
	o.EligibleForFailover = nil
	o.NullFields = addNullField(o.NullFields, "EligibleForFailover")
}

// GetFailoverEnabled returns the FailoverEnabled field value if set, zero value otherwise
func (o *StreamsProcessorWithStats) GetFailoverEnabled() bool {
	if o == nil || IsNil(o.FailoverEnabled) {
		var ret bool
		return ret
	}
	return *o.FailoverEnabled
}

// GetFailoverEnabledOk returns a tuple with the FailoverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetFailoverEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FailoverEnabled) {
		return nil, false
	}

	return o.FailoverEnabled, true
}

// HasFailoverEnabled returns a boolean if a field has been set.
func (o *StreamsProcessorWithStats) HasFailoverEnabled() bool {
	if o != nil && !IsNil(o.FailoverEnabled) {
		return true
	}

	return false
}

// SetFailoverEnabled gets a reference to the given bool and assigns it to the FailoverEnabled field.
func (o *StreamsProcessorWithStats) SetFailoverEnabled(v bool) {
	o.FailoverEnabled = &v
	o.NullFields = removeNullField(o.NullFields, "FailoverEnabled")
}

// SetFailoverEnabledNil sets FailoverEnabled to an explicit JSON null when marshaled.
func (o *StreamsProcessorWithStats) SetFailoverEnabledNil() {
	o.FailoverEnabled = nil
	o.NullFields = addNullField(o.NullFields, "FailoverEnabled")
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsProcessorWithStats) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsProcessorWithStats) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsProcessorWithStats) SetLinks(v []Link) {
	o.Links = &v
	o.NullFields = removeNullField(o.NullFields, "Links")
}

// SetLinksNil sets Links to an explicit JSON null when marshaled.
func (o *StreamsProcessorWithStats) SetLinksNil() {
	o.Links = nil
	o.NullFields = addNullField(o.NullFields, "Links")
}

// GetName returns the Name field value
func (o *StreamsProcessorWithStats) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StreamsProcessorWithStats) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise
func (o *StreamsProcessorWithStats) GetOptions() StreamsOptions {
	if o == nil || IsNil(o.Options) {
		var ret StreamsOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetOptionsOk() (*StreamsOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}

	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *StreamsProcessorWithStats) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given StreamsOptions and assigns it to the Options field.
func (o *StreamsProcessorWithStats) SetOptions(v StreamsOptions) {
	o.Options = &v
	o.NullFields = removeNullField(o.NullFields, "Options")
}

// SetOptionsNil sets Options to an explicit JSON null when marshaled.
func (o *StreamsProcessorWithStats) SetOptionsNil() {
	o.Options = nil
	o.NullFields = addNullField(o.NullFields, "Options")
}

// GetPipeline returns the Pipeline field value
func (o *StreamsProcessorWithStats) GetPipeline() []any {
	if o == nil {
		var ret []any
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetPipelineOk() (*[]any, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *StreamsProcessorWithStats) SetPipeline(v []any) {
	o.Pipeline = v
}

// GetState returns the State field value
func (o *StreamsProcessorWithStats) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *StreamsProcessorWithStats) SetState(v string) {
	o.State = v
}

// GetStats returns the Stats field value if set, zero value otherwise
func (o *StreamsProcessorWithStats) GetStats() any {
	if o == nil || IsNil(o.Stats) {
		var ret any
		return ret
	}
	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetStatsOk() (any, bool) {
	if o == nil || IsNil(o.Stats) {
		var ret any
		return ret, false
	}

	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *StreamsProcessorWithStats) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given any and assigns it to the Stats field.
func (o *StreamsProcessorWithStats) SetStats(v any) {
	o.Stats = v
	o.NullFields = removeNullField(o.NullFields, "Stats")
}

// SetStatsNil sets Stats to an explicit JSON null when marshaled.
func (o *StreamsProcessorWithStats) SetStatsNil() {
	o.Stats = nil
	o.NullFields = addNullField(o.NullFields, "Stats")
}

// GetTier returns the Tier field value if set, zero value otherwise
func (o *StreamsProcessorWithStats) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsProcessorWithStats) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}

	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *StreamsProcessorWithStats) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *StreamsProcessorWithStats) SetTier(v string) {
	o.Tier = &v
	o.NullFields = removeNullField(o.NullFields, "Tier")
}

// SetTierNil sets Tier to an explicit JSON null when marshaled.
func (o *StreamsProcessorWithStats) SetTierNil() {
	o.Tier = nil
	o.NullFields = addNullField(o.NullFields, "Tier")
}
