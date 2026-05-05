// Code based on the AtlasAPI V2 OpenAPI file

package admin

// StreamsTenantUpdatePreviewUpdateRequest Details to update a stream tenant, optionally accompanied by a processor status change (for example, triggering bulk regional failover).
type StreamsTenantUpdatePreviewUpdateRequest struct {
	// Human-readable label that identifies the cloud provider.
	CloudProvider   *string                 `json:"cloudProvider,omitempty"`
	ProcessorStatus *StreamsProcessorStatus `json:"processorStatus,omitempty"`
	// Name of the cloud provider region hosting Atlas Stream Processing.
	Region       *string       `json:"region,omitempty"`
	StreamConfig *StreamConfig `json:"streamConfig,omitempty"`
	// List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
	// Read only field.
	Links *[]Link `json:"links,omitempty"`
}

// NewStreamsTenantUpdatePreviewUpdateRequest instantiates a new StreamsTenantUpdatePreviewUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsTenantUpdatePreviewUpdateRequest() *StreamsTenantUpdatePreviewUpdateRequest {
	this := StreamsTenantUpdatePreviewUpdateRequest{}
	return &this
}

// NewStreamsTenantUpdatePreviewUpdateRequestWithDefaults instantiates a new StreamsTenantUpdatePreviewUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsTenantUpdatePreviewUpdateRequestWithDefaults() *StreamsTenantUpdatePreviewUpdateRequest {
	this := StreamsTenantUpdatePreviewUpdateRequest{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}

	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *StreamsTenantUpdatePreviewUpdateRequest) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetProcessorStatus returns the ProcessorStatus field value if set, zero value otherwise
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetProcessorStatus() StreamsProcessorStatus {
	if o == nil || IsNil(o.ProcessorStatus) {
		var ret StreamsProcessorStatus
		return ret
	}
	return *o.ProcessorStatus
}

// GetProcessorStatusOk returns a tuple with the ProcessorStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetProcessorStatusOk() (*StreamsProcessorStatus, bool) {
	if o == nil || IsNil(o.ProcessorStatus) {
		return nil, false
	}

	return o.ProcessorStatus, true
}

// HasProcessorStatus returns a boolean if a field has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) HasProcessorStatus() bool {
	if o != nil && !IsNil(o.ProcessorStatus) {
		return true
	}

	return false
}

// SetProcessorStatus gets a reference to the given StreamsProcessorStatus and assigns it to the ProcessorStatus field.
func (o *StreamsTenantUpdatePreviewUpdateRequest) SetProcessorStatus(v StreamsProcessorStatus) {
	o.ProcessorStatus = &v
}

// GetRegion returns the Region field value if set, zero value otherwise
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}

	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StreamsTenantUpdatePreviewUpdateRequest) SetRegion(v string) {
	o.Region = &v
}

// GetStreamConfig returns the StreamConfig field value if set, zero value otherwise
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetStreamConfig() StreamConfig {
	if o == nil || IsNil(o.StreamConfig) {
		var ret StreamConfig
		return ret
	}
	return *o.StreamConfig
}

// GetStreamConfigOk returns a tuple with the StreamConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetStreamConfigOk() (*StreamConfig, bool) {
	if o == nil || IsNil(o.StreamConfig) {
		return nil, false
	}

	return o.StreamConfig, true
}

// HasStreamConfig returns a boolean if a field has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) HasStreamConfig() bool {
	if o != nil && !IsNil(o.StreamConfig) {
		return true
	}

	return false
}

// SetStreamConfig gets a reference to the given StreamConfig and assigns it to the StreamConfig field.
func (o *StreamsTenantUpdatePreviewUpdateRequest) SetStreamConfig(v StreamConfig) {
	o.StreamConfig = &v
}

// GetLinks returns the Links field value if set, zero value otherwise
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetLinks() []Link {
	if o == nil || IsNil(o.Links) {
		var ret []Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) GetLinksOk() (*[]Link, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}

	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *StreamsTenantUpdatePreviewUpdateRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []Link and assigns it to the Links field.
func (o *StreamsTenantUpdatePreviewUpdateRequest) SetLinks(v []Link) {
	o.Links = &v
}
